package core

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
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
	rand.Read(token)
	return base64.RawURLEncoding.EncodeToString(token)
}
