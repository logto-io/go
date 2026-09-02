package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

func generateTestTokenBySigningKey(keyId string, signingKey jose.SigningKey, claims any) (string, error) {
	signingKeyOptions := jose.SignerOptions{}
	signingKeyOptions.WithType("JWT")
	signingKeyOptions.WithHeader("kid", keyId)

	rsaSigner, createSignerError := jose.NewSigner(signingKey, &signingKeyOptions)
	if createSignerError != nil {
		return "", createSignerError
	}

	builder := jwt.Signed(rsaSigner)

	token, buildTokenError := builder.Claims(claims).Serialize()

	if buildTokenError != nil {
		return "", buildTokenError
	}

	return token, nil
}

func generateRsaSigningKeyAndJwks(signingKeyId string) (jose.SigningKey, jose.JSONWebKeySet, error) {
	rsaPrivateKey, generateKeyError := rsa.GenerateKey(rand.Reader, 2048)

	if generateKeyError != nil {
		return jose.SigningKey{}, jose.JSONWebKeySet{}, generateKeyError
	}

	signingKey := jose.SigningKey{Algorithm: jose.RS256, Key: rsaPrivateKey}

	jwks := jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			{
				Key:       &rsaPrivateKey.PublicKey,
				Algorithm: string(jose.RS256),
				KeyID:     signingKeyId,
			},
		},
	}

	return signingKey, jwks, nil
}

func generateEcdsaSigningKeyAndJwks(signingKeyId string) (jose.SigningKey, jose.JSONWebKeySet, error) {
	ecdsaPrivateKey, generateKeyError := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if generateKeyError != nil {
		return jose.SigningKey{}, jose.JSONWebKeySet{}, generateKeyError
	}

	signingKey := jose.SigningKey{Algorithm: jose.ES512, Key: ecdsaPrivateKey}

	jwks := jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			{
				Key:       &ecdsaPrivateKey.PublicKey,
				Algorithm: string(jose.ES512),
				KeyID:     signingKeyId,
			},
		},
	}

	return signingKey, jwks, nil
}

func generateTestTokenAndCorrespondJwks(
	claims any,
	signingKeyAndJwksGenerator func(string) (jose.SigningKey, jose.JSONWebKeySet, error),
) (string, jose.JSONWebKeySet, error) {
	signingKeyId := "test-signing-key-id"
	signingKey, jwks, generateKeyError := signingKeyAndJwksGenerator(signingKeyId)
	if generateKeyError != nil {
		return "", jose.JSONWebKeySet{}, generateKeyError
	}

	token, generateTokenError := generateTestTokenBySigningKey(signingKeyId, signingKey, claims)
	if generateTokenError != nil {
		return "", jose.JSONWebKeySet{}, generateTokenError
	}

	return token, jwks, nil
}

func generateRsaSigningTestTokenAndCorrespondJwks(claims any) (string, jose.JSONWebKeySet, error) {
	return generateTestTokenAndCorrespondJwks(claims, generateRsaSigningKeyAndJwks)
}

func generateEcdsaSigningTestTokenAndCorrespondJwks(claims any) (string, jose.JSONWebKeySet, error) {
	return generateTestTokenAndCorrespondJwks(claims, generateEcdsaSigningKeyAndJwks)
}
