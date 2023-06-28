package hash

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

const (
	saltLength  = 16
	memory      = 65536
	iterations  = 3
	parallelism = 2
	keyLength   = 32
)

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

func generateRandomBytes(n uint32) ([]byte, error) {
	arr := make([]byte, n)
	_, err := rand.Read(arr)

	if err != nil {
		return nil, err
	}

	return arr, nil
}

func HashPassword(password string) (string, error) {
	salt, err := generateRandomBytes(saltLength)

	if err != nil {
		return "", err
	}

	hash := argon2.Key([]byte(password), salt, iterations, memory, parallelism, keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, iterations, parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func VerifyPassword(plainPassword string, hashedPassword string) (bool, error) {
	salt, hash, err := decodeHashedPassword(hashedPassword)

	if err != nil {
		return false, err
	}

	otherHash := argon2.Key([]byte(plainPassword), salt, iterations, memory, parallelism, keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}

	return false, nil
}

func decodeHashedPassword(str string) ([]byte, []byte, error) {
	parts := strings.Split(str, "$")

	if len(parts) != 6 {
		return nil, nil, ErrInvalidHash
	}

	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)

	if err != nil {
		return nil, nil, err
	}

	if version != argon2.Version {
		return nil, nil, ErrIncompatibleVersion
	}

	var (
		mem, iter, parallel int
	)
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &mem, &iter, &parallel)

	if err != nil {
		return nil, nil, err
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(parts[4])

	if err != nil {
		return nil, nil, err
	}

	hash, err := base64.RawStdEncoding.Strict().DecodeString(parts[5])
	if err != nil {
		return nil, nil, err
	}

	return salt, hash, nil
}
