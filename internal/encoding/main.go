package encoding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

// HMACSHA256 ...
func HMACSHA256(data []byte, secret string) string {
	hm := hmac.New(sha256.New, []byte(secret))
	hm.Write(data)
	return hex.EncodeToString(hm.Sum(nil))
}

// Hash will take in identifiers and an IP address and use it to return a hashed string
func Hash(identifiers interface{}, ip string) (string, error) {
	// Encode identifiers into byte array
	hashArray, err := json.Marshal(identifiers)
	if err != nil {
		return "", err
	}

	// Encode byte array into base64 into hmac sha256
	hash := base64.StdEncoding.EncodeToString(hashArray)
	hash = HMACSHA256([]byte(hash), ip)

	return hash, nil
}
