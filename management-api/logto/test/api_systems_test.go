/*
Logto API references

Testing SystemsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package logto

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/logto-io/go/management-api/logto"
)

func Test_logto_SystemsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SystemsAPIService GetSystemApplicationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SystemsAPI.GetSystemApplicationConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
