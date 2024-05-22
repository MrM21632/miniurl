package uidgen

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var ErrEmptyEnvVar = errors.New("getenv: Specified environment variable is empty or missing")

func GetenvStr(key string) (string, error) {
	result := os.Getenv(key)
	if result == "" {
		return result, ErrEmptyEnvVar
	}

	return result, nil
}

func GetenvInteger(key string) (uint64, error) {
	str, err := GetenvStr(key)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func GetenvBoolean(key string) (bool, error) {
	str, err := GetenvStr(key)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}

	return result, nil
}

// Epoch is configurable using the UIDGEN_EPOCH_MS environment variable. This allows for quick, reconfigurable
// deployemnts of the UID generation service, especially when scale is required.
//
// If the epoch is not found, we default to the epoch used by Twitter's original Snowflake algorithm.
func GetEpoch() uint64 {
	var epoch uint64
	var err error
	if epoch, err = GetenvInteger("UIDGEN_EPOCH_MS"); err != nil {
		fmt.Println("get epoch failed: envvar UIDGEN_EPOCH_MS not found, using default")
		return 1288834974657
	}

	return epoch
}

// Server ID is configurable using the UIDGEN_NODE_ID environment variable. This allows for quick, reconfigurable
// deployments of the UID generation service, especially when scale is required.
//
// If the server ID is not found, we default to 0.
func GetServerId() uint64 {
	var serverId uint64
	var err error
	if serverId, err = GetenvInteger("UIDGEN_NODE_ID"); err != nil {
		fmt.Println("get server id failed: envvar UIDGEN_NODE_ID not found, using default")
		return 0
	}

	return serverId
}
