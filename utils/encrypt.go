package utils

import (
	"crypto/sha1"
	"encoding/base64"
)

func HashSHA256(stringToHash string) string {
	hasher := sha1.New()
	hasher.Write([]byte(stringToHash))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
