package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"reflect"
)

// GetIPFromRequest ...
func GetIPFromRequest(rq *http.Request) string {
	forwarded := rq.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return rq.RemoteAddr
}

// GetArrayFromStruct ...
func GetArrayFromStruct(itf interface{}) []interface{} {
	values := reflect.ValueOf(itf)
	ary := make([]interface{}, values.NumField())
	for i := 0; i < values.NumField(); i++ {
		ary[i] = values.Field(i).Interface()
	}
	return ary
}

// GetHashFromArray takes an array of strings and returns a unique hash
func GetHashFromArray(identifiers []interface{}) (*string, error) {
	mIdentifiers, err := json.Marshal(identifiers)
	if err != nil {
		return nil, err
	}

	hash := base64.StdEncoding.EncodeToString(mIdentifiers)

	return &hash, nil
}

// GetArrayHashFromStruct ...
func GetArrayHashFromStruct(identifier interface{}) (*string, error) {
	array := GetArrayFromStruct(identifier)
	return GetHashFromArray(array)
}

// GetJSONHashFromStruct ...
func GetJSONHashFromStruct(identifiers interface{}) (*string, error) {
	mIdentifiers, err := json.Marshal(identifiers)
	if err != nil {
		return nil, err
	}

	hash := base64.StdEncoding.EncodeToString(mIdentifiers)

	return &hash, nil
}

// EncodeArray takes an array and returns it encoded using sha256
func EncodeArray(identifiers []interface{}) ([]byte, error) {
	sha := sha256.New()

	for _, value := range identifiers {
		v, err := getBytes(value)
		if err != nil {
			return nil, err
		}

		sha.Write(v)
	}

	return sha.Sum(nil), nil
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}

// EncodeStruct takes a struct and returns it encoded using sha256
func EncodeStruct(identifiers interface{}) ([]byte, error) {
	array := GetArrayFromStruct(identifiers)
	return EncodeArray(array)
}

// HMACSHA256 ...
func HMACSHA256(data []byte, secret string) string {
	hm := hmac.New(sha256.New, []byte(secret))
	hm.Write(data)
	return hex.EncodeToString(hm.Sum(nil))
}
