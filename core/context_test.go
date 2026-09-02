package core

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type roundTripperFunc func(request *http.Request) (*http.Response, error)

func (roundTrip roundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return roundTrip(request)
}

type testContextKey struct{}

func TestContextVariantsShouldBindContextToRequest(t *testing.T) {
	ctx := context.WithValue(context.Background(), testContextKey{}, "value")

	var requestContext context.Context
	client := &http.Client{
		Transport: roundTripperFunc(func(request *http.Request) (*http.Response, error) {
			requestContext = request.Context()
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{},
				Body:       io.NopCloser(strings.NewReader(`{}`)),
			}, nil
		}),
	}

	testCases := []struct {
		name string
		call func() error
	}{
		{
			name: "FetchOidcConfigContext",
			call: func() error {
				_, err := FetchOidcConfigContext(ctx, client, "https://example.com/oidc/.well-known/openid-configuration")
				return err
			},
		},
		{
			name: "FetchJwksContext",
			call: func() error {
				_, err := FetchJwksContext(ctx, client, "https://example.com/oidc/jwks")
				return err
			},
		},
		{
			name: "FetchTokenByAuthorizationCodeContext",
			call: func() error {
				_, err := FetchTokenByAuthorizationCodeContext(ctx, client, &FetchTokenByAuthorizationCodeOptions{
					TokenEndpoint: "https://example.com/oidc/token",
				})
				return err
			},
		},
		{
			name: "FetchTokenByRefreshTokenContext",
			call: func() error {
				_, err := FetchTokenByRefreshTokenContext(ctx, client, &FetchTokenByRefreshTokenOptions{
					TokenEndpoint: "https://example.com/oidc/token",
				})
				return err
			},
		},
		{
			name: "FetchUserInfoContext",
			call: func() error {
				_, err := FetchUserInfoContext(ctx, client, "https://example.com/oidc/me", "accessToken")
				return err
			},
		},
		{
			name: "RevokeContext",
			call: func() error {
				return RevokeContext(ctx, client, &RevocationOptions{
					RevocationEndpoint: "https://example.com/oidc/revoke",
				})
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			requestContext = nil

			assert.Nil(t, testCase.call())

			if assert.NotNil(t, requestContext) {
				assert.Equal(t, "value", requestContext.Value(testContextKey{}))
			}
		})
	}
}

func TestContextVariantsShouldAbortRequestWhenContextIsCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, fetchOidcConfigErr := FetchOidcConfigContext(ctx, &http.Client{}, "https://example.com/oidc/.well-known/openid-configuration")

	assert.ErrorIs(t, fetchOidcConfigErr, context.Canceled)
}
