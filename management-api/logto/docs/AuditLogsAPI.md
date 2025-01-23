# \AuditLogsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLog**](AuditLogsAPI.md#GetLog) | **Get** /api/logs/{id} | Get log
[**ListLogs**](AuditLogsAPI.md#ListLogs) | **Get** /api/logs | Get logs



## GetLog

> ListLogs200ResponseInner GetLog(ctx, id).Execute()

Get log



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
	id := "id_example" // string | The unique identifier of the log.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.GetLog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLog`: ListLogs200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.GetLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the log. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListLogs200ResponseInner**](ListLogs200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogs

> []ListLogs200ResponseInner ListLogs(ctx).UserId(userId).ApplicationId(applicationId).LogKey(logKey).Page(page).PageSize(pageSize).Execute()

Get logs



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
	userId := "userId_example" // string | Filter logs by user ID. (optional)
	applicationId := "applicationId_example" // string | Filter logs by application ID. (optional)
	logKey := "logKey_example" // string | Filter logs by log key. (optional)
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.ListLogs(context.Background()).UserId(userId).ApplicationId(applicationId).LogKey(logKey).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.ListLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLogs`: []ListLogs200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.ListLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Filter logs by user ID. | 
 **applicationId** | **string** | Filter logs by application ID. | 
 **logKey** | **string** | Filter logs by log key. | 
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListLogs200ResponseInner**](ListLogs200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

