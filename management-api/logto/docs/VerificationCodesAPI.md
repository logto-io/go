# \VerificationCodesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVerificationCode**](VerificationCodesAPI.md#CreateVerificationCode) | **Post** /api/verification-codes | Request and send a verification code
[**VerifyVerificationCode**](VerificationCodesAPI.md#VerifyVerificationCode) | **Post** /api/verification-codes/verify | Verify a verification code



## CreateVerificationCode

> CreateVerificationCode(ctx).CreateVerificationCodeRequest(createVerificationCodeRequest).Execute()

Request and send a verification code



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
	createVerificationCodeRequest := openapiclient.CreateVerificationCode_request{CreateVerificationCodeRequestOneOf: openapiclient.NewCreateVerificationCodeRequestOneOf("Email_example")} // CreateVerificationCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VerificationCodesAPI.CreateVerificationCode(context.Background()).CreateVerificationCodeRequest(createVerificationCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationCodesAPI.CreateVerificationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerificationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVerificationCodeRequest** | [**CreateVerificationCodeRequest**](CreateVerificationCodeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyVerificationCode

> VerifyVerificationCode(ctx).VerifyVerificationCodeRequest(verifyVerificationCodeRequest).Execute()

Verify a verification code



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
	verifyVerificationCodeRequest := openapiclient.VerifyVerificationCode_request{VerifyVerificationCodeRequestOneOf: openapiclient.NewVerifyVerificationCodeRequestOneOf("Email_example", "VerificationCode_example")} // VerifyVerificationCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VerificationCodesAPI.VerifyVerificationCode(context.Background()).VerifyVerificationCodeRequest(verifyVerificationCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationCodesAPI.VerifyVerificationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyVerificationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyVerificationCodeRequest** | [**VerifyVerificationCodeRequest**](VerifyVerificationCodeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

