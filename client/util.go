package client

import (
	"net/http"
	"sort"
	"strings"

	"gopkg.in/square/go-jose.v2/jwt"
)

func GetOriginRequestUrl(request *http.Request) string {
	return getRequestProtocol(request) + "://" + request.Host + request.RequestURI
}

func getRequestProtocol(request *http.Request) string {
	if request.TLS != nil {
		return "https"
	}
	proto := request.Header.Get("X-Forwarded-Proto")
	if proto != "" {
		extractedProto := strings.Split(proto, ",")[0]
		return strings.ToLower(strings.Trim(extractedProto, " "))
	}
	return "http"
}

func buildAccessTokenKey(scopes []string, resource string, organizationId string) string {
	sort.Strings(scopes)
	scopesPart := strings.Join(scopes, " ")

	organizationPart := ""
	if organizationId != "" {
		organizationPart = "#" + organizationId
	}

	return scopesPart + "@" + resource + organizationPart
}

func getResourceFromAccessToken(accessToken string) string {
	jwtObject, parseToJwtErr := jwt.ParseSigned(accessToken)
	if parseToJwtErr != nil {
		return ""
	}

	type audContainedClaims struct {
		Aud string `json:"aud"`
	}

	var audClaim audContainedClaims
	claimsErr := jwtObject.UnsafeClaimsWithoutVerification(&audClaim)
	if claimsErr != nil {
		return ""
	}

	return audClaim.Aud
}
