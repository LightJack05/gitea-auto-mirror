package crypto

// Argon2idPasswordHash represents an Argon2id hashed password with its parameters
type Argon2idPasswordHash struct {
	Version     int
	Salt        []byte
	Hash        []byte
	Time        uint32
	Memory      uint32
	Parallelism uint8
}
