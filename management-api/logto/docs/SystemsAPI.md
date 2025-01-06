# \SystemsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemApplicationConfig**](SystemsAPI.md#GetSystemApplicationConfig) | **Get** /api/systems/application | Get the application constants.



## GetSystemApplicationConfig

> GetSystemApplicationConfig200Response GetSystemApplicationConfig(ctx).Execute()

Get the application constants.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/logto-io/go/management-api/logto"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemsAPI.GetSystemApplicationConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemsAPI.GetSystemApplicationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemApplicationConfig`: GetSystemApplicationConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemsAPI.GetSystemApplicationConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemApplicationConfigRequest struct via the builder pattern


### Return type

[**GetSystemApplicationConfig200Response**](GetSystemApplicationConfig200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

