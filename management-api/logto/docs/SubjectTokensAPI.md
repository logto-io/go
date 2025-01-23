# \SubjectTokensAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubjectToken**](SubjectTokensAPI.md#CreateSubjectToken) | **Post** /api/subject-tokens | Create a new subject token.



## CreateSubjectToken

> CreateSubjectToken201Response CreateSubjectToken(ctx).CreateSubjectTokenRequest(createSubjectTokenRequest).Execute()

Create a new subject token.



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
	createSubjectTokenRequest := *openapiclient.NewCreateSubjectTokenRequest("UserId_example") // CreateSubjectTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectTokensAPI.CreateSubjectToken(context.Background()).CreateSubjectTokenRequest(createSubjectTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectTokensAPI.CreateSubjectToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubjectToken`: CreateSubjectToken201Response
	fmt.Fprintf(os.Stdout, "Response from `SubjectTokensAPI.CreateSubjectToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubjectTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubjectTokenRequest** | [**CreateSubjectTokenRequest**](CreateSubjectTokenRequest.md) |  | 

### Return type

[**CreateSubjectToken201Response**](CreateSubjectToken201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

