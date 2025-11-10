package crypto

type Argon2idPasswordHash struct {
	Version     int
	Salt        []byte
	Hash        []byte
	Time        uint32
	Memory      uint32
	Parallelism uint8
}
