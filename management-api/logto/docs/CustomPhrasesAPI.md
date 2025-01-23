# \CustomPhrasesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCustomPhrase**](CustomPhrasesAPI.md#DeleteCustomPhrase) | **Delete** /api/custom-phrases/{languageTag} | Delete custom phrase
[**GetCustomPhrase**](CustomPhrasesAPI.md#GetCustomPhrase) | **Get** /api/custom-phrases/{languageTag} | Get custom phrases
[**ListCustomPhrases**](CustomPhrasesAPI.md#ListCustomPhrases) | **Get** /api/custom-phrases | Get all custom phrases
[**ReplaceCustomPhrase**](CustomPhrasesAPI.md#ReplaceCustomPhrase) | **Put** /api/custom-phrases/{languageTag} | Upsert custom phrases



## DeleteCustomPhrase

> DeleteCustomPhrase(ctx, languageTag).Execute()

Delete custom phrase



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
	languageTag := "languageTag_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomPhrasesAPI.DeleteCustomPhrase(context.Background(), languageTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPhrasesAPI.DeleteCustomPhrase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageTag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomPhraseRequest struct via the builder pattern


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


## GetCustomPhrase

> ListCustomPhrases200ResponseInner GetCustomPhrase(ctx, languageTag).Execute()

Get custom phrases



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
	languageTag := "languageTag_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomPhrasesAPI.GetCustomPhrase(context.Background(), languageTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPhrasesAPI.GetCustomPhrase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomPhrase`: ListCustomPhrases200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CustomPhrasesAPI.GetCustomPhrase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageTag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomPhraseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListCustomPhrases200ResponseInner**](ListCustomPhrases200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomPhrases

> []ListCustomPhrases200ResponseInner ListCustomPhrases(ctx).Execute()

Get all custom phrases



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
	resp, r, err := apiClient.CustomPhrasesAPI.ListCustomPhrases(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPhrasesAPI.ListCustomPhrases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomPhrases`: []ListCustomPhrases200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CustomPhrasesAPI.ListCustomPhrases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomPhrasesRequest struct via the builder pattern


### Return type

[**[]ListCustomPhrases200ResponseInner**](ListCustomPhrases200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCustomPhrase

> ListCustomPhrases200ResponseInner ReplaceCustomPhrase(ctx, languageTag).TranslationObject(translationObject).Execute()

Upsert custom phrases



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
	languageTag := "languageTag_example" // string | 
	translationObject := *openapiclient.NewTranslationObject() // TranslationObject | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomPhrasesAPI.ReplaceCustomPhrase(context.Background(), languageTag).TranslationObject(translationObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPhrasesAPI.ReplaceCustomPhrase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceCustomPhrase`: ListCustomPhrases200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CustomPhrasesAPI.ReplaceCustomPhrase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageTag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCustomPhraseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **translationObject** | [**TranslationObject**](TranslationObject.md) |  | 

### Return type

[**ListCustomPhrases200ResponseInner**](ListCustomPhrases200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

