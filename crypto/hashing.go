package crypto

import (
	"crypto/rand"
	"hash/crc64"
    "log"
    "fmt"
)

var DefaultSaltLength int = 10

func GenerateSalt(length int) (string, error) {
    salt := make([]byte, length)
    _, err := rand.Read(salt)
    if err != nil {
        log.Print(err)
        return "", err
    }
    return string(salt), nil
}

func HashPassword(password string) (hashedPassword string) {
	crcTable := crc64.MakeTable(crc64.ECMA)
	hashedPassword = fmt.Sprintf("%d", crc64.Checksum([]byte(password), crcTable))
    return
}