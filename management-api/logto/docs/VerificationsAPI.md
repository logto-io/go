# \VerificationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVerificationByPassword**](VerificationsAPI.md#CreateVerificationByPassword) | **Post** /api/verifications/password | Create a record by password
[**CreateVerificationBySocial**](VerificationsAPI.md#CreateVerificationBySocial) | **Post** /api/verifications/social | Create a social verification record
[**CreateVerificationByVerificationCode**](VerificationsAPI.md#CreateVerificationByVerificationCode) | **Post** /api/verifications/verification-code | Create a record by verification code
[**VerifyVerificationBySocial**](VerificationsAPI.md#VerifyVerificationBySocial) | **Post** /api/verifications/social/verify | Verify a social verification record
[**VerifyVerificationByVerificationCode**](VerificationsAPI.md#VerifyVerificationByVerificationCode) | **Post** /api/verifications/verification-code/verify | Verify verification code



## CreateVerificationByPassword

> CreateVerificationByPassword201Response CreateVerificationByPassword(ctx).CreateVerificationByPasswordRequest(createVerificationByPasswordRequest).Execute()

Create a record by password



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
	createVerificationByPasswordRequest := *openapiclient.NewCreateVerificationByPasswordRequest("Password_example") // CreateVerificationByPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.CreateVerificationByPassword(context.Background()).CreateVerificationByPasswordRequest(createVerificationByPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.CreateVerificationByPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVerificationByPassword`: CreateVerificationByPassword201Response
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.CreateVerificationByPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerificationByPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVerificationByPasswordRequest** | [**CreateVerificationByPasswordRequest**](CreateVerificationByPasswordRequest.md) |  | 

### Return type

[**CreateVerificationByPassword201Response**](CreateVerificationByPassword201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVerificationBySocial

> CreateVerificationBySocial201Response CreateVerificationBySocial(ctx).CreateVerificationBySocialRequest(createVerificationBySocialRequest).Execute()

Create a social verification record



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
	createVerificationBySocialRequest := *openapiclient.NewCreateVerificationBySocialRequest("State_example", "RedirectUri_example", "ConnectorId_example") // CreateVerificationBySocialRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.CreateVerificationBySocial(context.Background()).CreateVerificationBySocialRequest(createVerificationBySocialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.CreateVerificationBySocial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVerificationBySocial`: CreateVerificationBySocial201Response
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.CreateVerificationBySocial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerificationBySocialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVerificationBySocialRequest** | [**CreateVerificationBySocialRequest**](CreateVerificationBySocialRequest.md) |  | 

### Return type

[**CreateVerificationBySocial201Response**](CreateVerificationBySocial201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVerificationByVerificationCode

> CreateVerificationByPassword201Response CreateVerificationByVerificationCode(ctx).CreateVerificationByVerificationCodeRequest(createVerificationByVerificationCodeRequest).Execute()

Create a record by verification code



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
	createVerificationByVerificationCodeRequest := *openapiclient.NewCreateVerificationByVerificationCodeRequest(*openapiclient.NewCreateAndSendVerificationCodeRequestIdentifier("Type_example", "Value_example")) // CreateVerificationByVerificationCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.CreateVerificationByVerificationCode(context.Background()).CreateVerificationByVerificationCodeRequest(createVerificationByVerificationCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.CreateVerificationByVerificationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVerificationByVerificationCode`: CreateVerificationByPassword201Response
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.CreateVerificationByVerificationCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerificationByVerificationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVerificationByVerificationCodeRequest** | [**CreateVerificationByVerificationCodeRequest**](CreateVerificationByVerificationCodeRequest.md) |  | 

### Return type

[**CreateVerificationByPassword201Response**](CreateVerificationByPassword201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyVerificationBySocial

> VerifyVerificationByVerificationCode200Response VerifyVerificationBySocial(ctx).VerifyVerificationBySocialRequest(verifyVerificationBySocialRequest).Execute()

Verify a social verification record



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
	verifyVerificationBySocialRequest := *openapiclient.NewVerifyVerificationBySocialRequest(map[string]interface{}(123), "VerificationRecordId_example") // VerifyVerificationBySocialRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.VerifyVerificationBySocial(context.Background()).VerifyVerificationBySocialRequest(verifyVerificationBySocialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.VerifyVerificationBySocial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyVerificationBySocial`: VerifyVerificationByVerificationCode200Response
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.VerifyVerificationBySocial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyVerificationBySocialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyVerificationBySocialRequest** | [**VerifyVerificationBySocialRequest**](VerifyVerificationBySocialRequest.md) |  | 

### Return type

[**VerifyVerificationByVerificationCode200Response**](VerifyVerificationByVerificationCode200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyVerificationByVerificationCode

> VerifyVerificationByVerificationCode200Response VerifyVerificationByVerificationCode(ctx).VerifyVerificationCodeVerificationRequest(verifyVerificationCodeVerificationRequest).Execute()

Verify verification code



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
	verifyVerificationCodeVerificationRequest := *openapiclient.NewVerifyVerificationCodeVerificationRequest(*openapiclient.NewVerifyVerificationCodeVerificationRequestIdentifier("Type_example", "Value_example"), "VerificationId_example", "Code_example") // VerifyVerificationCodeVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.VerifyVerificationByVerificationCode(context.Background()).VerifyVerificationCodeVerificationRequest(verifyVerificationCodeVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.VerifyVerificationByVerificationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyVerificationByVerificationCode`: VerifyVerificationByVerificationCode200Response
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.VerifyVerificationByVerificationCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyVerificationByVerificationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyVerificationCodeVerificationRequest** | [**VerifyVerificationCodeVerificationRequest**](VerifyVerificationCodeVerificationRequest.md) |  | 

### Return type

[**VerifyVerificationByVerificationCode200Response**](VerifyVerificationByVerificationCode200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

