package crypto

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/argon2"
)

// HashPassword represents an Argon2id hashed password with its parameters
func HashPassword(password string, memory uint32, time uint32, parallelism uint8, salt []byte, hashLength uint32) *Argon2idPasswordHash {
	hash := argon2.IDKey([]byte(password), salt, time, memory, parallelism, hashLength)
	return &Argon2idPasswordHash{
		Memory:      memory,
		Time:        time,
		Parallelism: parallelism,
		Salt:        salt,
		Hash:        hash,
		Version:     argon2.Version,
	}
}

// ParseHash parses an Argon2id hash string into an Argon2idPasswordHash struct
func ParseHash(hash string) (*Argon2idPasswordHash, error) {
	segments := strings.Split(hash, "$")
	if len(segments) != 6 {
		return nil, fmt.Errorf("invalid hash format: expected 6 hash segments, found %d", len(segments))
	}
	if segments[1] != "argon2id" {
		return nil, fmt.Errorf("unsupported hash type: %s", segments[1])
	}
	var argon2Hash Argon2idPasswordHash
	fmt.Sscanf(segments[2], "v=%d", &argon2Hash.Version)
	if(argon2Hash.Version != argon2.Version){
		return nil, fmt.Errorf("unsupported or invalid argon2 version: %d", argon2Hash.Version)
	}

	fmt.Sscanf(segments[3], "m=%d,t=%d,p=%d", &argon2Hash.Memory, &argon2Hash.Time, &argon2Hash.Parallelism)
	if(argon2Hash.Memory == 0 || argon2Hash.Time == 0 || argon2Hash.Parallelism == 0){
		return nil, fmt.Errorf("cannot parse hash with zero values: one or more parameters were 0: m=%d,t=%d,p=%d", argon2Hash.Memory, argon2Hash.Time, argon2Hash.Parallelism)
	}

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

// CompareHashes compares two Argon2idPasswordHash structs for equality
func CompareHashes(hash1, hash2 *Argon2idPasswordHash) bool {
	if hash1.Memory != hash2.Memory || hash1.Time != hash2.Time || hash1.Parallelism != hash2.Parallelism || hash1.Version != hash2.Version {
		return false
	}
	if !bytes.Equal(hash1.Salt, hash2.Salt) || !bytes.Equal(hash1.Hash, hash2.Hash) {
		return false
	}
	return true
}

// decodeHashBase64 decodes a base64 encoded string, trying both standard and raw encodings
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
