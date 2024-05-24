package miniurl

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"log"
	"miniurl/miniurl/models"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jxskiss/base62"
)

func GetUniqueId() (uint64, error) {
	url := os.Getenv("UIDGEN_ADDRESS")
	resp, err := http.Post(url+"/generate", "application/json", nil)
	if err != nil {
		log.Println("Error occurred when calling uidgen: " + err.Error())
		return 0, err
	}
	if resp == nil {
		log.Println("Missing response body from uidgen")
		return 0, err
	}

	var result models.GetUniqueIdResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Invalid response body from uidgen")
		return 0, err
	}

	return strconv.ParseUint(result.UniqueID, 10, 64)
}

func ComputeChecksum(value uint64) string {
	salt := make([]byte, 8)
	rand.Read(salt)

	val_arr := make([]byte, 8)
	binary.LittleEndian.PutUint64(val_arr, value)

	result := sha256.Sum256(append(val_arr[:], salt[:]...))
	return base62.EncodeToString(result[:])
}

func CreateUrlRecord(c *gin.Context) {
	var input models.CreateURLInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if found, _ := CheckForLongUrlInDatabase(input.OriginalURL); found {
		c.JSON(http.StatusConflict, gin.H{"message": "Record with URL " + input.OriginalURL + " already exists"})
		return
	}

	new_uid, err := GetUniqueId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// Keep trying to generate a new shortened_url until we get a unique one
	var shortened_url string
	for {
		encoded_checksum := ComputeChecksum(new_uid)
		shortened_url = encoded_checksum[:8]
		if found, _ := CheckForShortUrlInDatabase(shortened_url); !found {
			break
		}
	}

	result, err := WriteNewRecordToDatabase(input.OriginalURL, shortened_url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"shortened_url": result})
}
