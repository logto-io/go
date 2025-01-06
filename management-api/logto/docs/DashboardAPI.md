# \DashboardAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveUserCounts**](DashboardAPI.md#GetActiveUserCounts) | **Get** /api/dashboard/users/active | Get active user data
[**GetNewUserCounts**](DashboardAPI.md#GetNewUserCounts) | **Get** /api/dashboard/users/new | Get new user count
[**GetTotalUserCount**](DashboardAPI.md#GetTotalUserCount) | **Get** /api/dashboard/users/total | Get total user count



## GetActiveUserCounts

> GetActiveUserCounts200Response GetActiveUserCounts(ctx).Date(date).Execute()

Get active user data



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
	date := "date_example" // string | The date to get active user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetActiveUserCounts(context.Background()).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetActiveUserCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveUserCounts`: GetActiveUserCounts200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetActiveUserCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveUserCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** | The date to get active user data. | 

### Return type

[**GetActiveUserCounts200Response**](GetActiveUserCounts200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewUserCounts

> GetNewUserCounts200Response GetNewUserCounts(ctx).Execute()

Get new user count



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
	resp, r, err := apiClient.DashboardAPI.GetNewUserCounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetNewUserCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewUserCounts`: GetNewUserCounts200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetNewUserCounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewUserCountsRequest struct via the builder pattern


### Return type

[**GetNewUserCounts200Response**](GetNewUserCounts200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotalUserCount

> GetTotalUserCount200Response GetTotalUserCount(ctx).Execute()

Get total user count



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
	resp, r, err := apiClient.DashboardAPI.GetTotalUserCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetTotalUserCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTotalUserCount`: GetTotalUserCount200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetTotalUserCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTotalUserCountRequest struct via the builder pattern


### Return type

[**GetTotalUserCount200Response**](GetTotalUserCount200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

