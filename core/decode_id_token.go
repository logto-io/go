package core

func DecodeIdToken(token string) (IdTokenClaims, error) {
	jwtObject, err := ParseSignedJwt(token)

	if err != nil {
		return IdTokenClaims{}, err
	}

	var idTokenClaims IdTokenClaims
	claimsErr := jwtObject.UnsafeClaimsWithoutVerification(&idTokenClaims)

	if claimsErr != nil {
		return IdTokenClaims{}, claimsErr
	}

	return idTokenClaims, nil
}
