package crypto

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

func ParseHash(hash string) (*Argon2idPasswordHash, error) {
	segments := strings.Split(hash, "$")
	if len(segments) != 6 {
		return nil, fmt.Errorf("invalid hash format")
	}
	if segments[1] != "argon2id" {
		return nil, fmt.Errorf("unsupported hash type: %s", segments[1])
	}
	var argon2Hash Argon2idPasswordHash
	fmt.Sscanf(segments[2], "v=%d", &argon2Hash.Version)

	fmt.Sscanf(segments[3], "m=%d,t=%d,p=%d", &argon2Hash.Memory, &argon2Hash.Time, &argon2Hash.Parallelism)

	saltBytes, err := decodeHashBase64(segments[4])
	if err != nil {
		return nil, fmt.Errorf("invalid salt encoding: %v", err)
	}
	argon2Hash.Salt = saltBytes

	hashBytes, err := decodeHashBase64(segments[5])
	if err != nil {
		return nil, fmt.Errorf("invalid hash encoding: %v", err)
	}
	argon2Hash.Hash = hashBytes
	return &argon2Hash, nil
}

func decodeHashBase64(encoded string) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(encoded)
	if err == nil {
		log.Println("Info: Decoded base64 string using StdEncoding (With padding)")
		return bytes, nil
	}
	bytes, err = base64.RawStdEncoding.DecodeString(encoded)
	if err == nil {
		log.Println("Info: Decoded base64 string using RawStdEncoding (Without padding) instead of StdEncoding")
		return bytes, nil
	}
	return nil, fmt.Errorf("Failed to decode base64 string, tried with and without padding: %v", err)
}
