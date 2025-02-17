package crypto

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/arifai/zenith/pkg/errormessage"
	"golang.org/x/crypto/argon2"
	"log"
	"strings"
)

// Argon2IdHash is the struct for GenerateHash and VerifyHash
type Argon2IdHash struct {
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
	SaltLen uint32
}

// DefaultArgon2IDHash is an instance of crypto.Argon2IdHash used to configure Argon2ID password hashing with specified time,
// memory, threads, key length, and salt length. Reference: https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
var DefaultArgon2IDHash = &Argon2IdHash{Time: 2, Memory: 19 * 1024, Threads: 1, KeyLen: 32, SaltLen: 32}

// Argon2Version is the version of the argon2 algorithm
const Argon2Version = argon2.Version

// GenerateHash generates a password hash using the Argon2ID algorithm and given salt. Returns the encoded hash or an error.
func (a *Argon2IdHash) GenerateHash(password, salt []byte) (string, error) {
	if err := validateSaltLength(salt, a.SaltLen); err != nil {
		return "", err
	}

	if len(salt) == 0 {
		var err error
		salt, err = generateBytes(a.SaltLen)
		if err != nil {
			return "", err
		}
	}

	hash := argon2.IDKey(password, salt, a.Time, a.Memory, a.Threads, a.KeyLen)

	return encodeHashComponents(salt, hash, a), nil
}

func validateSaltLength(salt []byte, expectedSaltLen uint32) error {
	if len(salt) > 0 && uint32(len(salt)) != expectedSaltLen {
		log.Printf("salt length is incorrect: expected %d bytes, got %d bytes", expectedSaltLen, len(salt))
		return errormessage.ErrInvalidSaltLength
	}
	return nil
}

func encodeHashComponents(salt, hash []byte, a *Argon2IdHash) string {
	base64Salt := base64.StdEncoding.EncodeToString(salt)
	base64Hash := base64.StdEncoding.EncodeToString(hash)
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", Argon2Version, a.Memory, a.Time, a.Threads, base64Salt, base64Hash)
}

// VerifyHash compares a password with its encoded hash to check for validity.
// It returns true if the password matches the hash, false otherwise. If an error occurs during the process,
// it returns false and the error.
func VerifyHash(password, encodedHash string) (bool, error) {
	a, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, a.Time, a.Memory, a.Threads, a.KeyLen)
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}

	return false, nil
}

// generateBytes generates a slice of random bytes with the specified length.
// Returns the generated bytes or an error if any.
func generateBytes(length uint32) ([]byte, error) {
	secret := make([]byte, length)
	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}
	return secret, nil
}

// decodeHash decodes an encoded Argon2ID hash string into its components: Argon2IdHash struct, salt, and hash.
// It returns an error if the encoded string is invalid or incompatible.
func decodeHash(encodedHash string) (a *Argon2IdHash, salt, hash []byte, err error) {
	value := strings.Split(encodedHash, "$")
	if len(value) != 6 {
		return nil, nil, nil, errormessage.ErrInvalidEncodedHash
	}

	var version int
	if _, err = fmt.Sscanf(value[2], "v=%d", &version); err != nil {
		return nil, nil, nil, errormessage.ErrIncompatibleArgon2Version
	}

	a = &Argon2IdHash{}
	if _, err = fmt.Sscanf(value[3], "m=%d,t=%d,p=%d", &a.Memory, &a.Time, &a.Threads); err != nil {
		return nil, nil, nil, err
	}

	if salt, err = base64.StdEncoding.DecodeString(value[4]); err != nil {
		return nil, nil, nil, err
	}
	a.SaltLen = uint32(len(salt))

	if hash, err = base64.StdEncoding.DecodeString(value[5]); err != nil {
		return nil, nil, nil, err
	}
	a.KeyLen = uint32(len(hash))

	return a, salt, hash, nil
}
