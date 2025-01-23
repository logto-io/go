# \SSOConnectorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSsoConnector**](SSOConnectorsAPI.md#CreateSsoConnector) | **Post** /api/sso-connectors | Create SSO connector
[**DeleteSsoConnector**](SSOConnectorsAPI.md#DeleteSsoConnector) | **Delete** /api/sso-connectors/{id} | Delete SSO connector
[**GetSsoConnector**](SSOConnectorsAPI.md#GetSsoConnector) | **Get** /api/sso-connectors/{id} | Get SSO connector
[**ListSsoConnectors**](SSOConnectorsAPI.md#ListSsoConnectors) | **Get** /api/sso-connectors | List SSO connectors
[**UpdateSsoConnector**](SSOConnectorsAPI.md#UpdateSsoConnector) | **Patch** /api/sso-connectors/{id} | Update SSO connector



## CreateSsoConnector

> ListOrganizationJitSsoConnectors200ResponseInner CreateSsoConnector(ctx).CreateSsoConnectorRequest(createSsoConnectorRequest).Execute()

Create SSO connector



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
	createSsoConnectorRequest := *openapiclient.NewCreateSsoConnectorRequest("ProviderName_example", "ConnectorName_example") // CreateSsoConnectorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOConnectorsAPI.CreateSsoConnector(context.Background()).CreateSsoConnectorRequest(createSsoConnectorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOConnectorsAPI.CreateSsoConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSsoConnector`: ListOrganizationJitSsoConnectors200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SSOConnectorsAPI.CreateSsoConnector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSsoConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSsoConnectorRequest** | [**CreateSsoConnectorRequest**](CreateSsoConnectorRequest.md) |  | 

### Return type

[**ListOrganizationJitSsoConnectors200ResponseInner**](ListOrganizationJitSsoConnectors200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSsoConnector

> DeleteSsoConnector(ctx, id).Execute()

Delete SSO connector



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
	id := "id_example" // string | The unique identifier of the sso connector.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SSOConnectorsAPI.DeleteSsoConnector(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOConnectorsAPI.DeleteSsoConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the sso connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSsoConnectorRequest struct via the builder pattern


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


## GetSsoConnector

> ListSsoConnectors200ResponseInner GetSsoConnector(ctx, id).Execute()

Get SSO connector



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
	id := "id_example" // string | The unique identifier of the sso connector.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOConnectorsAPI.GetSsoConnector(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOConnectorsAPI.GetSsoConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsoConnector`: ListSsoConnectors200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SSOConnectorsAPI.GetSsoConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the sso connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsoConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSsoConnectors200ResponseInner**](ListSsoConnectors200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSsoConnectors

> []ListSsoConnectors200ResponseInner ListSsoConnectors(ctx).Page(page).PageSize(pageSize).Execute()

List SSO connectors



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
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOConnectorsAPI.ListSsoConnectors(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOConnectorsAPI.ListSsoConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSsoConnectors`: []ListSsoConnectors200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SSOConnectorsAPI.ListSsoConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSsoConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListSsoConnectors200ResponseInner**](ListSsoConnectors200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSsoConnector

> ListSsoConnectors200ResponseInner UpdateSsoConnector(ctx, id).UpdateSsoConnectorRequest(updateSsoConnectorRequest).Execute()

Update SSO connector



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
	id := "id_example" // string | The unique identifier of the sso connector.
	updateSsoConnectorRequest := *openapiclient.NewUpdateSsoConnectorRequest() // UpdateSsoConnectorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOConnectorsAPI.UpdateSsoConnector(context.Background(), id).UpdateSsoConnectorRequest(updateSsoConnectorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOConnectorsAPI.UpdateSsoConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSsoConnector`: ListSsoConnectors200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SSOConnectorsAPI.UpdateSsoConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the sso connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsoConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSsoConnectorRequest** | [**UpdateSsoConnectorRequest**](UpdateSsoConnectorRequest.md) |  | 

### Return type

[**ListSsoConnectors200ResponseInner**](ListSsoConnectors200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

