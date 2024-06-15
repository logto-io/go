package core

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

var (
	defaultRandomStringLength uint32 = 64
)

func GenerateCodeVerifier() string {
	return generateRandomString(defaultRandomStringLength)
}

func GenerateCodeChallenge(codeVerifier string) string {
	sha := sha256.New()
	sha.Write([]byte(codeVerifier))
	return base64.RawURLEncoding.EncodeToString(sha.Sum(nil))
}

func GenerateState() string {
	return generateRandomString(defaultRandomStringLength)
}

func generateRandomString(length uint32) string {
	token := make([]byte, length)
	if _, error := rand.Read(token); error != nil {
		// This should never happen
		panic(fmt.Sprintf("Failed to generate random string: %v", error))
	}
	return base64.RawURLEncoding.EncodeToString(token)
}
