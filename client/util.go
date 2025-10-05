package client

import (
	"net/http"
	"sort"
	"strings"

	"github.com/logto-io/go/v2/core"
)

func GetOriginRequestUrl(request *http.Request) string {
	return getRequestProtocol(request) + "://" + request.Host + request.RequestURI
}

func getForwaredRequestUrl(request *http.Request) string {
	proto := getRequestProtocol(request)
	host := getForwaredRequestHost(request)
	uri := getForwaredRequestRequestUri(request)
	return proto + "://" + host + uri
}
func getForwaredRequestHost(request *http.Request) string {
	host := request.Header.Get("X-Forwarded-Host")
	if host != "" {
		return host
	}
	return request.Host
}
func getForwaredRequestRequestUri(request *http.Request) string {
	uri := request.Header.Get("X-Forwarded-Url")
	if uri != "" {
		return uri
	}
	return request.RequestURI
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
	jwtObject, parseToJwtErr := core.ParseSignedJwt(accessToken)
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
