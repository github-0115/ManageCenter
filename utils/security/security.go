package security

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

var (
	method     = "pbkdf2:sha1"
	saltLength = 8
	iterations = 1000
	SALT_CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func genSalt() string {
	var bytes = make([]byte, saltLength)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = SALT_CHARS[v%byte(len(SALT_CHARS))]
	}
	return string(bytes)
}

func hashInternal(salt string, password string) string {
	hash := pbkdf2.Key([]byte(password), []byte(salt), iterations, 20, sha1.New)
	return hex.EncodeToString(hash)
}

func GeneratePasswordHash(password string) string {
	salt := genSalt()
	hash := hashInternal(salt, password)
	return fmt.Sprintf("pbkdf2:sha1:%v$%s$%s", iterations, salt, hash)
}

func CheckPasswordHash(password string, hash string) bool {
	if strings.Count(hash, "$") < 2 {
		return false
	}
	pwd_hash_list := strings.Split(hash, "$")
	return pwd_hash_list[2] == hashInternal(pwd_hash_list[1], password)
}
