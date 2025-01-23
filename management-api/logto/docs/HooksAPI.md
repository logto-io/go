# \HooksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHook**](HooksAPI.md#CreateHook) | **Post** /api/hooks | Create a hook
[**CreateHookTest**](HooksAPI.md#CreateHookTest) | **Post** /api/hooks/{id}/test | Test hook
[**DeleteHook**](HooksAPI.md#DeleteHook) | **Delete** /api/hooks/{id} | Delete hook
[**GetHook**](HooksAPI.md#GetHook) | **Get** /api/hooks/{id} | Get hook
[**ListHookRecentLogs**](HooksAPI.md#ListHookRecentLogs) | **Get** /api/hooks/{id}/recent-logs | Get recent logs for a hook
[**ListHooks**](HooksAPI.md#ListHooks) | **Get** /api/hooks | Get hooks
[**UpdateHook**](HooksAPI.md#UpdateHook) | **Patch** /api/hooks/{id} | Update hook
[**UpdateHookSigningKey**](HooksAPI.md#UpdateHookSigningKey) | **Patch** /api/hooks/{id}/signing-key | Update signing key for a hook



## CreateHook

> CreateHook201Response CreateHook(ctx).CreateHookRequest(createHookRequest).Execute()

Create a hook



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
	createHookRequest := *openapiclient.NewCreateHookRequest(*openapiclient.NewCreateHookRequestConfig("Url_example")) // CreateHookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HooksAPI.CreateHook(context.Background()).CreateHookRequest(createHookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HooksAPI.CreateHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHook`: CreateHook201Response
	fmt.Fprintf(os.Stdout, "Response from `HooksAPI.CreateHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createHookRequest** | [**CreateHookRequest**](CreateHookRequest.md) |  | 

### Return type

[**CreateHook201Response**](CreateHook201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHookTest

> CreateHookTest(ctx, id).CreateHookTestRequest(createHookTestRequest).Execute()

Test hook



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
	id := "id_example" // string | The unique identifier of the hook.
	createHookTestRequest := *openapiclient.NewCreateHookTestRequest([]string{"Events_example"}, *openapiclient.NewCreateHookTestRequestConfig("Url_example")) // CreateHookTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HooksAPI.CreateHookTest(context.Background(), id).CreateHookTestRequest(createHookTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HooksAPI.CreateHookTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHookTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createHookTestRequest** | [**CreateHookTestRequest**](CreateHookTestRequest.md) |  | 

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


## DeleteHook

> DeleteHook(ctx, id).Execute()

Delete hook



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
	id := "id_example" // string | The unique identifier of the hook.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HooksAPI.DeleteHook(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HooksAPI.DeleteHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHookRequest struct via the builder pattern


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


## GetHook

> ListHooks200ResponseInner GetHook(ctx, id).IncludeExecutionStats(includeExecutionStats).Execute()

Get hook



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
	id := "id_example" // string | The unique identifier of the hook.
	includeExecutionStats := "includeExecutionStats_example" // string | Whether to include execution stats in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HooksAPI.GetHook(context.Background(), id).IncludeExecutionStats(includeExecutionStats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HooksAPI.GetHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHook`: ListHooks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `HooksAPI.GetHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeExecutionStats** | **string** | Whether to include execution stats in the response. | 

### Return type

[**ListHooks200ResponseInner**](ListHooks200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHookRecentLogs

> []ListHookRecentLogs200ResponseInner ListHookRecentLogs(ctx, id).LogKey(logKey).Page(page).PageSize(pageSize).Execute()

Get recent logs for a hook



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
	id := "id_example" // string | The unique identifier of the hook.
	logKey := "logKey_example" // string | The log key to filter logs. (optional)
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HooksAPI.ListHookRecentLogs(context.Background(), id).LogKey(logKey).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HooksAPI.ListHookRecentLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHookRecentLogs`: []ListHookRecentLogs200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `HooksAPI.ListHookRecentLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHookRecentLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logKey** | **string** | The log key to filter logs. | 
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListHookRecentLogs200ResponseInner**](ListHookRecentLogs200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHooks

> []ListHooks200ResponseInner ListHooks(ctx).IncludeExecutionStats(includeExecutionStats).Page(page).PageSize(pageSize).Execute()

Get hooks



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
	includeExecutionStats := "includeExecutionStats_example" // string | Whether to include execution stats in the response. (optional)
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HooksAPI.ListHooks(context.Background()).IncludeExecutionStats(includeExecutionStats).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HooksAPI.ListHooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHooks`: []ListHooks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `HooksAPI.ListHooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeExecutionStats** | **string** | Whether to include execution stats in the response. | 
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListHooks200ResponseInner**](ListHooks200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHook

> CreateHook201Response UpdateHook(ctx, id).UpdateHookRequest(updateHookRequest).Execute()

Update hook



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
	id := "id_example" // string | The unique identifier of the hook.
	updateHookRequest := *openapiclient.NewUpdateHookRequest() // UpdateHookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HooksAPI.UpdateHook(context.Background(), id).UpdateHookRequest(updateHookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HooksAPI.UpdateHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHook`: CreateHook201Response
	fmt.Fprintf(os.Stdout, "Response from `HooksAPI.UpdateHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateHookRequest** | [**UpdateHookRequest**](UpdateHookRequest.md) |  | 

### Return type

[**CreateHook201Response**](CreateHook201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHookSigningKey

> CreateHook201Response UpdateHookSigningKey(ctx, id).Execute()

Update signing key for a hook



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
	id := "id_example" // string | The unique identifier of the hook.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HooksAPI.UpdateHookSigningKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HooksAPI.UpdateHookSigningKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHookSigningKey`: CreateHook201Response
	fmt.Fprintf(os.Stdout, "Response from `HooksAPI.UpdateHookSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHookSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateHook201Response**](CreateHook201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

