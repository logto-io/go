# \ConnectorFactoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnectorFactory**](ConnectorFactoriesAPI.md#GetConnectorFactory) | **Get** /api/connector-factories/{id} | Get connector factory
[**ListConnectorFactories**](ConnectorFactoriesAPI.md#ListConnectorFactories) | **Get** /api/connector-factories | Get connector factories



## GetConnectorFactory

> ListConnectorFactories200ResponseInner GetConnectorFactory(ctx, id).Execute()

Get connector factory



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
	id := "id_example" // string | The unique identifier of the connector factory.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorFactoriesAPI.GetConnectorFactory(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorFactoriesAPI.GetConnectorFactory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorFactory`: ListConnectorFactories200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConnectorFactoriesAPI.GetConnectorFactory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the connector factory. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorFactoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListConnectorFactories200ResponseInner**](ListConnectorFactories200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectorFactories

> []ListConnectorFactories200ResponseInner ListConnectorFactories(ctx).Execute()

Get connector factories



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
	resp, r, err := apiClient.ConnectorFactoriesAPI.ListConnectorFactories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorFactoriesAPI.ListConnectorFactories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnectorFactories`: []ListConnectorFactories200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConnectorFactoriesAPI.ListConnectorFactories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorFactoriesRequest struct via the builder pattern


### Return type

[**[]ListConnectorFactories200ResponseInner**](ListConnectorFactories200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

