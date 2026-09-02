package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizedShouldTrimTrailingSlashOfBaseUrl(t *testing.T) {
	logtoConfig := LogtoConfig{BaseUrl: "https://my-app.com/"}

	logtoConfig.normalized()

	assert.Equal(t, "https://my-app.com", logtoConfig.BaseUrl)
}
