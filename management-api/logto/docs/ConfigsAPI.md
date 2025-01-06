# \ConfigsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteJwtCustomizer**](ConfigsAPI.md#DeleteJwtCustomizer) | **Delete** /api/configs/jwt-customizer/{tokenTypePath} | Delete JWT customizer
[**DeleteOidcKey**](ConfigsAPI.md#DeleteOidcKey) | **Delete** /api/configs/oidc/{keyType}/{keyId} | Delete OIDC key
[**GetAdminConsoleConfig**](ConfigsAPI.md#GetAdminConsoleConfig) | **Get** /api/configs/admin-console | Get admin console config
[**GetJwtCustomizer**](ConfigsAPI.md#GetJwtCustomizer) | **Get** /api/configs/jwt-customizer/{tokenTypePath} | Get JWT customizer
[**GetOidcKeys**](ConfigsAPI.md#GetOidcKeys) | **Get** /api/configs/oidc/{keyType} | Get OIDC keys
[**ListJwtCustomizers**](ConfigsAPI.md#ListJwtCustomizers) | **Get** /api/configs/jwt-customizer | Get all JWT customizers
[**RotateOidcKeys**](ConfigsAPI.md#RotateOidcKeys) | **Post** /api/configs/oidc/{keyType}/rotate | Rotate OIDC keys
[**TestJwtCustomizer**](ConfigsAPI.md#TestJwtCustomizer) | **Post** /api/configs/jwt-customizer/test | Test JWT customizer
[**UpdateAdminConsoleConfig**](ConfigsAPI.md#UpdateAdminConsoleConfig) | **Patch** /api/configs/admin-console | Update admin console config
[**UpdateJwtCustomizer**](ConfigsAPI.md#UpdateJwtCustomizer) | **Patch** /api/configs/jwt-customizer/{tokenTypePath} | Update JWT customizer
[**UpsertJwtCustomizer**](ConfigsAPI.md#UpsertJwtCustomizer) | **Put** /api/configs/jwt-customizer/{tokenTypePath} | Create or update JWT customizer



## DeleteJwtCustomizer

> DeleteJwtCustomizer(ctx, tokenTypePath).Execute()

Delete JWT customizer



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
	tokenTypePath := "tokenTypePath_example" // string | The token type path to delete the JWT customizer for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigsAPI.DeleteJwtCustomizer(context.Background(), tokenTypePath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.DeleteJwtCustomizer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenTypePath** | **string** | The token type path to delete the JWT customizer for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJwtCustomizerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOidcKey

> DeleteOidcKey(ctx, keyType, keyId).Execute()

Delete OIDC key



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
	keyType := "keyType_example" // string | Private keys are used to sign OIDC JWTs. Cookie keys are used to sign OIDC cookies. For clients, they do not need to know private keys to verify OIDC JWTs; they can use public keys from the JWKS endpoint instead.
	keyId := "keyId_example" // string | The unique identifier of the key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigsAPI.DeleteOidcKey(context.Background(), keyType, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.DeleteOidcKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyType** | **string** | Private keys are used to sign OIDC JWTs. Cookie keys are used to sign OIDC cookies. For clients, they do not need to know private keys to verify OIDC JWTs; they can use public keys from the JWKS endpoint instead. | 
**keyId** | **string** | The unique identifier of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOidcKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminConsoleConfig

> GetAdminConsoleConfig200Response GetAdminConsoleConfig(ctx).Execute()

Get admin console config



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
	resp, r, err := apiClient.ConfigsAPI.GetAdminConsoleConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetAdminConsoleConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminConsoleConfig`: GetAdminConsoleConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetAdminConsoleConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminConsoleConfigRequest struct via the builder pattern


### Return type

[**GetAdminConsoleConfig200Response**](GetAdminConsoleConfig200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJwtCustomizer

> GetJwtCustomizer200Response GetJwtCustomizer(ctx, tokenTypePath).Execute()

Get JWT customizer



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
	tokenTypePath := "tokenTypePath_example" // string | The token type to get the JWT customizer for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.GetJwtCustomizer(context.Background(), tokenTypePath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetJwtCustomizer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJwtCustomizer`: GetJwtCustomizer200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetJwtCustomizer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenTypePath** | **string** | The token type to get the JWT customizer for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJwtCustomizerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetJwtCustomizer200Response**](GetJwtCustomizer200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOidcKeys

> []GetOidcKeys200ResponseInner GetOidcKeys(ctx, keyType).Execute()

Get OIDC keys



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
	keyType := "keyType_example" // string | Private keys are used to sign OIDC JWTs. Cookie keys are used to sign OIDC cookies. For clients, they do not need to know private keys to verify OIDC JWTs; they can use public keys from the JWKS endpoint instead.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.GetOidcKeys(context.Background(), keyType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetOidcKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOidcKeys`: []GetOidcKeys200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetOidcKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyType** | **string** | Private keys are used to sign OIDC JWTs. Cookie keys are used to sign OIDC cookies. For clients, they do not need to know private keys to verify OIDC JWTs; they can use public keys from the JWKS endpoint instead. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOidcKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOidcKeys200ResponseInner**](GetOidcKeys200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJwtCustomizers

> []ListJwtCustomizers200ResponseInner ListJwtCustomizers(ctx).Execute()

Get all JWT customizers



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
	resp, r, err := apiClient.ConfigsAPI.ListJwtCustomizers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.ListJwtCustomizers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListJwtCustomizers`: []ListJwtCustomizers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.ListJwtCustomizers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListJwtCustomizersRequest struct via the builder pattern


### Return type

[**[]ListJwtCustomizers200ResponseInner**](ListJwtCustomizers200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateOidcKeys

> []GetOidcKeys200ResponseInner RotateOidcKeys(ctx, keyType).RotateOidcKeysRequest(rotateOidcKeysRequest).Execute()

Rotate OIDC keys



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
	keyType := "keyType_example" // string | Private keys are used to sign OIDC JWTs. Cookie keys are used to sign OIDC cookies. For clients, they do not need to know private keys to verify OIDC JWTs; they can use public keys from the JWKS endpoint instead.
	rotateOidcKeysRequest := *openapiclient.NewRotateOidcKeysRequest() // RotateOidcKeysRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.RotateOidcKeys(context.Background(), keyType).RotateOidcKeysRequest(rotateOidcKeysRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.RotateOidcKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateOidcKeys`: []GetOidcKeys200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.RotateOidcKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyType** | **string** | Private keys are used to sign OIDC JWTs. Cookie keys are used to sign OIDC cookies. For clients, they do not need to know private keys to verify OIDC JWTs; they can use public keys from the JWKS endpoint instead. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateOidcKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rotateOidcKeysRequest** | [**RotateOidcKeysRequest**](RotateOidcKeysRequest.md) |  | 

### Return type

[**[]GetOidcKeys200ResponseInner**](GetOidcKeys200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestJwtCustomizer

> map[string]interface{} TestJwtCustomizer(ctx).TestJwtCustomizerRequest(testJwtCustomizerRequest).Execute()

Test JWT customizer



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
	testJwtCustomizerRequest := openapiclient.TestJwtCustomizer_request{TestJwtCustomizerRequestOneOf: openapiclient.NewTestJwtCustomizerRequestOneOf("TokenType_example", "Script_example", *openapiclient.NewGetJwtCustomizer200ResponseOneOfTokenSample(), *openapiclient.NewGetJwtCustomizer200ResponseOneOfContextSample(*openapiclient.NewGetJwtCustomizer200ResponseOneOfContextSampleUser()))} // TestJwtCustomizerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.TestJwtCustomizer(context.Background()).TestJwtCustomizerRequest(testJwtCustomizerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.TestJwtCustomizer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestJwtCustomizer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.TestJwtCustomizer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestJwtCustomizerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testJwtCustomizerRequest** | [**TestJwtCustomizerRequest**](TestJwtCustomizerRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdminConsoleConfig

> GetAdminConsoleConfig200Response UpdateAdminConsoleConfig(ctx).UpdateAdminConsoleConfigRequest(updateAdminConsoleConfigRequest).Execute()

Update admin console config



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
	updateAdminConsoleConfigRequest := *openapiclient.NewUpdateAdminConsoleConfigRequest() // UpdateAdminConsoleConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.UpdateAdminConsoleConfig(context.Background()).UpdateAdminConsoleConfigRequest(updateAdminConsoleConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.UpdateAdminConsoleConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdminConsoleConfig`: GetAdminConsoleConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.UpdateAdminConsoleConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdminConsoleConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAdminConsoleConfigRequest** | [**UpdateAdminConsoleConfigRequest**](UpdateAdminConsoleConfigRequest.md) |  | 

### Return type

[**GetAdminConsoleConfig200Response**](GetAdminConsoleConfig200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateJwtCustomizer

> GetJwtCustomizer200Response UpdateJwtCustomizer(ctx, tokenTypePath).UpsertJwtCustomizerRequest(upsertJwtCustomizerRequest).Execute()

Update JWT customizer



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
	tokenTypePath := "tokenTypePath_example" // string | The token type to update a JWT customizer for.
	upsertJwtCustomizerRequest := *openapiclient.NewUpsertJwtCustomizerRequest() // UpsertJwtCustomizerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.UpdateJwtCustomizer(context.Background(), tokenTypePath).UpsertJwtCustomizerRequest(upsertJwtCustomizerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.UpdateJwtCustomizer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateJwtCustomizer`: GetJwtCustomizer200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.UpdateJwtCustomizer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenTypePath** | **string** | The token type to update a JWT customizer for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJwtCustomizerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsertJwtCustomizerRequest** | [**UpsertJwtCustomizerRequest**](UpsertJwtCustomizerRequest.md) |  | 

### Return type

[**GetJwtCustomizer200Response**](GetJwtCustomizer200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertJwtCustomizer

> GetJwtCustomizer200Response UpsertJwtCustomizer(ctx, tokenTypePath).UpsertJwtCustomizerRequest(upsertJwtCustomizerRequest).Execute()

Create or update JWT customizer



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
	tokenTypePath := "tokenTypePath_example" // string | The token type to create a JWT customizer for.
	upsertJwtCustomizerRequest := *openapiclient.NewUpsertJwtCustomizerRequest() // UpsertJwtCustomizerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.UpsertJwtCustomizer(context.Background(), tokenTypePath).UpsertJwtCustomizerRequest(upsertJwtCustomizerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.UpsertJwtCustomizer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertJwtCustomizer`: GetJwtCustomizer200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.UpsertJwtCustomizer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenTypePath** | **string** | The token type to create a JWT customizer for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertJwtCustomizerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsertJwtCustomizerRequest** | [**UpsertJwtCustomizerRequest**](UpsertJwtCustomizerRequest.md) |  | 

### Return type

[**GetJwtCustomizer200Response**](GetJwtCustomizer200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

