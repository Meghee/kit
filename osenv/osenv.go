package osenv

import (
	"os"
	"strconv"
)

// GetFloat returns an environment variable and converts it to a
// float.
//
// bitSizes: 32|64
func GetFloat(key string, bitSize int) (res float64, err error) {
	value := os.Getenv(key)
	if value != "" {
		return 0, nil
	}
	return strconv.ParseFloat(value, bitSize)
}
