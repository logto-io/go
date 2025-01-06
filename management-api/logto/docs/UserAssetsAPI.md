# \UserAssetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserAsset**](UserAssetsAPI.md#CreateUserAsset) | **Post** /api/user-assets | Upload asset
[**GetUserAssetServiceStatus**](UserAssetsAPI.md#GetUserAssetServiceStatus) | **Get** /api/user-assets/service-status | Get service status



## CreateUserAsset

> CreateUserAsset200Response CreateUserAsset(ctx).Execute()

Upload asset



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
	resp, r, err := apiClient.UserAssetsAPI.CreateUserAsset(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAssetsAPI.CreateUserAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserAsset`: CreateUserAsset200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAssetsAPI.CreateUserAsset`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserAssetRequest struct via the builder pattern


### Return type

[**CreateUserAsset200Response**](CreateUserAsset200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAssetServiceStatus

> GetUserAssetServiceStatus200Response GetUserAssetServiceStatus(ctx).Execute()

Get service status



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
	resp, r, err := apiClient.UserAssetsAPI.GetUserAssetServiceStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAssetsAPI.GetUserAssetServiceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAssetServiceStatus`: GetUserAssetServiceStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAssetsAPI.GetUserAssetServiceStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAssetServiceStatusRequest struct via the builder pattern


### Return type

[**GetUserAssetServiceStatus200Response**](GetUserAssetServiceStatus200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

