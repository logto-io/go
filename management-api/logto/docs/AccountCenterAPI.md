# \AccountCenterAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountCenterSettings**](AccountCenterAPI.md#GetAccountCenterSettings) | **Get** /api/account-center | Get account center settings
[**UpdateAccountCenterSettings**](AccountCenterAPI.md#UpdateAccountCenterSettings) | **Patch** /api/account-center | Update account center settings



## GetAccountCenterSettings

> GetAccountCenterSettings200Response GetAccountCenterSettings(ctx).Execute()

Get account center settings



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
	resp, r, err := apiClient.AccountCenterAPI.GetAccountCenterSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountCenterAPI.GetAccountCenterSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountCenterSettings`: GetAccountCenterSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountCenterAPI.GetAccountCenterSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCenterSettingsRequest struct via the builder pattern


### Return type

[**GetAccountCenterSettings200Response**](GetAccountCenterSettings200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountCenterSettings

> GetAccountCenterSettings200Response UpdateAccountCenterSettings(ctx).UpdateAccountCenterSettingsRequest(updateAccountCenterSettingsRequest).Execute()

Update account center settings



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
	updateAccountCenterSettingsRequest := *openapiclient.NewUpdateAccountCenterSettingsRequest() // UpdateAccountCenterSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountCenterAPI.UpdateAccountCenterSettings(context.Background()).UpdateAccountCenterSettingsRequest(updateAccountCenterSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountCenterAPI.UpdateAccountCenterSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountCenterSettings`: GetAccountCenterSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountCenterAPI.UpdateAccountCenterSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountCenterSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAccountCenterSettingsRequest** | [**UpdateAccountCenterSettingsRequest**](UpdateAccountCenterSettingsRequest.md) |  | 

### Return type

[**GetAccountCenterSettings200Response**](GetAccountCenterSettings200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

