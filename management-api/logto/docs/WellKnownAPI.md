# \WellKnownAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSignInExperienceConfig**](WellKnownAPI.md#GetSignInExperienceConfig) | **Get** /api/.well-known/sign-in-exp | Get full sign-in experience
[**GetSignInExperiencePhrases**](WellKnownAPI.md#GetSignInExperiencePhrases) | **Get** /api/.well-known/phrases | Get localized phrases
[**GetWellKnownExperience**](WellKnownAPI.md#GetWellKnownExperience) | **Get** /api/.well-known/experience | Get full sign-in experience
[**GetWellKnownExperienceOpenapiJson**](WellKnownAPI.md#GetWellKnownExperienceOpenapiJson) | **Get** /api/.well-known/experience.openapi.json | Get Experience API swagger JSON
[**GetWellKnownManagementOpenapiJson**](WellKnownAPI.md#GetWellKnownManagementOpenapiJson) | **Get** /api/.well-known/management.openapi.json | Get Management API swagger JSON
[**GetWellKnownUserOpenapiJson**](WellKnownAPI.md#GetWellKnownUserOpenapiJson) | **Get** /api/.well-known/user.openapi.json | Get User API swagger JSON



## GetSignInExperienceConfig

> GetSignInExperienceConfig200Response GetSignInExperienceConfig(ctx).OrganizationId(organizationId).AppId(appId).Execute()

Get full sign-in experience



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
	organizationId := "organizationId_example" // string |  (optional)
	appId := "appId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WellKnownAPI.GetSignInExperienceConfig(context.Background()).OrganizationId(organizationId).AppId(appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WellKnownAPI.GetSignInExperienceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignInExperienceConfig`: GetSignInExperienceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `WellKnownAPI.GetSignInExperienceConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignInExperienceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** |  | 
 **appId** | **string** |  | 

### Return type

[**GetSignInExperienceConfig200Response**](GetSignInExperienceConfig200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignInExperiencePhrases

> map[string]GetSignInExperiencePhrases200ResponseValue GetSignInExperiencePhrases(ctx).Lng(lng).Execute()

Get localized phrases



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
	lng := "lng_example" // string | The language tag for localization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WellKnownAPI.GetSignInExperiencePhrases(context.Background()).Lng(lng).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WellKnownAPI.GetSignInExperiencePhrases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignInExperiencePhrases`: map[string]GetSignInExperiencePhrases200ResponseValue
	fmt.Fprintf(os.Stdout, "Response from `WellKnownAPI.GetSignInExperiencePhrases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignInExperiencePhrasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lng** | **string** | The language tag for localization. | 

### Return type

[**map[string]GetSignInExperiencePhrases200ResponseValue**](GetSignInExperiencePhrases200ResponseValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWellKnownExperience

> GetSignInExperienceConfig200Response GetWellKnownExperience(ctx).OrganizationId(organizationId).AppId(appId).Execute()

Get full sign-in experience



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
	organizationId := "organizationId_example" // string |  (optional)
	appId := "appId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WellKnownAPI.GetWellKnownExperience(context.Background()).OrganizationId(organizationId).AppId(appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WellKnownAPI.GetWellKnownExperience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWellKnownExperience`: GetSignInExperienceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `WellKnownAPI.GetWellKnownExperience`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWellKnownExperienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** |  | 
 **appId** | **string** |  | 

### Return type

[**GetSignInExperienceConfig200Response**](GetSignInExperienceConfig200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWellKnownExperienceOpenapiJson

> GetWellKnownExperienceOpenapiJson(ctx).Execute()

Get Experience API swagger JSON



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
	r, err := apiClient.WellKnownAPI.GetWellKnownExperienceOpenapiJson(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WellKnownAPI.GetWellKnownExperienceOpenapiJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWellKnownExperienceOpenapiJsonRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWellKnownManagementOpenapiJson

> GetWellKnownManagementOpenapiJson(ctx).Execute()

Get Management API swagger JSON



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
	r, err := apiClient.WellKnownAPI.GetWellKnownManagementOpenapiJson(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WellKnownAPI.GetWellKnownManagementOpenapiJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWellKnownManagementOpenapiJsonRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWellKnownUserOpenapiJson

> GetWellKnownUserOpenapiJson(ctx).Execute()

Get User API swagger JSON



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
	r, err := apiClient.WellKnownAPI.GetWellKnownUserOpenapiJson(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WellKnownAPI.GetWellKnownUserOpenapiJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWellKnownUserOpenapiJsonRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

