package main

import (
	"fmt"
	"net/http"
	"strconv"
	"uidgen/uidgen"

	"github.com/gin-gonic/gin"
)

func main() {
	node, err := uidgen.InitializeNode()
	if err != nil {
		fmt.Printf("Encountered error while initializing node: %s", err.Error())
		return
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.POST(
		"/generate",
		func(c *gin.Context) {
			id := node.GeanerateId()
			c.JSON(http.StatusOK, gin.H{"uid": strconv.FormatUint(uint64(id), 10)})
		},
	)

	r.Run(":8080")
}
