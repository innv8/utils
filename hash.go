package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// GenerateSalt
// Generates salt to be used in hashing
func GenerateSalt() string {
	start := time.Now()
	defer func() {
		LogINFO("generating salt took %v", time.Since(start))
	}()
	var charset = []rune("1234567890qwertyuiopasdfghjklzxcvbnmASDFGHJKLZXCVBNMQWERTYUIOP")
	var builder strings.Builder
	charsLen := int64(len(charset))
	for i := 0; i < 32; i++ {
		temp, _ := rand.Int(rand.Reader, big.NewInt(charsLen))
		index := temp.Int64()
		builder.WriteRune(charset[index])
		//	builder.WriteRune(charset[rand.Intn(charsLen)])
	}
	return builder.String()
}

// HashPassword
// Hash the password
func HashPassword(password, salt string) string {
	start := time.Now()
	defer func() {
		LogINFO("hashing took %v", time.Since(start))
	}()
	p := password + salt
	bytes, _ := bcrypt.GenerateFromPassword([]byte(p), 14)
	return string(bytes)
}

// CheckPassword
// Validates the password
func CheckPassword(password, salt, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}
