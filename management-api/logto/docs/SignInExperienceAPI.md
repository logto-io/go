# \SignInExperienceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckPasswordWithDefaultSignInExperience**](SignInExperienceAPI.md#CheckPasswordWithDefaultSignInExperience) | **Post** /api/sign-in-exp/default/check-password | Check if a password meets the password policy
[**GetSignInExp**](SignInExperienceAPI.md#GetSignInExp) | **Get** /api/sign-in-exp | Get default sign-in experience settings
[**UpdateSignInExp**](SignInExperienceAPI.md#UpdateSignInExp) | **Patch** /api/sign-in-exp | Update default sign-in experience settings
[**UploadCustomUiAssets**](SignInExperienceAPI.md#UploadCustomUiAssets) | **Post** /api/sign-in-exp/default/custom-ui-assets | Upload custom UI assets



## CheckPasswordWithDefaultSignInExperience

> CheckPasswordWithDefaultSignInExperience200Response CheckPasswordWithDefaultSignInExperience(ctx).CheckPasswordWithDefaultSignInExperienceRequest(checkPasswordWithDefaultSignInExperienceRequest).Execute()

Check if a password meets the password policy



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
	checkPasswordWithDefaultSignInExperienceRequest := *openapiclient.NewCheckPasswordWithDefaultSignInExperienceRequest("Password_example") // CheckPasswordWithDefaultSignInExperienceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignInExperienceAPI.CheckPasswordWithDefaultSignInExperience(context.Background()).CheckPasswordWithDefaultSignInExperienceRequest(checkPasswordWithDefaultSignInExperienceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignInExperienceAPI.CheckPasswordWithDefaultSignInExperience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckPasswordWithDefaultSignInExperience`: CheckPasswordWithDefaultSignInExperience200Response
	fmt.Fprintf(os.Stdout, "Response from `SignInExperienceAPI.CheckPasswordWithDefaultSignInExperience`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckPasswordWithDefaultSignInExperienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkPasswordWithDefaultSignInExperienceRequest** | [**CheckPasswordWithDefaultSignInExperienceRequest**](CheckPasswordWithDefaultSignInExperienceRequest.md) |  | 

### Return type

[**CheckPasswordWithDefaultSignInExperience200Response**](CheckPasswordWithDefaultSignInExperience200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignInExp

> GetSignInExp200Response GetSignInExp(ctx).Execute()

Get default sign-in experience settings



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
	resp, r, err := apiClient.SignInExperienceAPI.GetSignInExp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignInExperienceAPI.GetSignInExp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignInExp`: GetSignInExp200Response
	fmt.Fprintf(os.Stdout, "Response from `SignInExperienceAPI.GetSignInExp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignInExpRequest struct via the builder pattern


### Return type

[**GetSignInExp200Response**](GetSignInExp200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSignInExp

> UpdateSignInExp200Response UpdateSignInExp(ctx).UpdateSignInExpRequest(updateSignInExpRequest).RemoveUnusedDemoSocialConnector(removeUnusedDemoSocialConnector).Execute()

Update default sign-in experience settings



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
	updateSignInExpRequest := *openapiclient.NewUpdateSignInExpRequest() // UpdateSignInExpRequest | 
	removeUnusedDemoSocialConnector := "removeUnusedDemoSocialConnector_example" // string | Whether to remove unused demo social connectors. (These demo social connectors are only used during cloud user onboarding) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignInExperienceAPI.UpdateSignInExp(context.Background()).UpdateSignInExpRequest(updateSignInExpRequest).RemoveUnusedDemoSocialConnector(removeUnusedDemoSocialConnector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignInExperienceAPI.UpdateSignInExp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSignInExp`: UpdateSignInExp200Response
	fmt.Fprintf(os.Stdout, "Response from `SignInExperienceAPI.UpdateSignInExp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSignInExpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSignInExpRequest** | [**UpdateSignInExpRequest**](UpdateSignInExpRequest.md) |  | 
 **removeUnusedDemoSocialConnector** | **string** | Whether to remove unused demo social connectors. (These demo social connectors are only used during cloud user onboarding) | 

### Return type

[**UpdateSignInExp200Response**](UpdateSignInExp200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCustomUiAssets

> UploadCustomUiAssets200Response UploadCustomUiAssets(ctx).File(file).Execute()

Upload custom UI assets



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
	file := TODO // interface{} | The zip file containing custom web assets such as HTML, CSS, and JavaScript files. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignInExperienceAPI.UploadCustomUiAssets(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignInExperienceAPI.UploadCustomUiAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadCustomUiAssets`: UploadCustomUiAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `SignInExperienceAPI.UploadCustomUiAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadCustomUiAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | [**interface{}**](interface{}.md) | The zip file containing custom web assets such as HTML, CSS, and JavaScript files. | 

### Return type

[**UploadCustomUiAssets200Response**](UploadCustomUiAssets200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

