# \ConnectorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnector**](ConnectorsAPI.md#CreateConnector) | **Post** /api/connectors | Create connector
[**CreateConnectorAuthorizationUri**](ConnectorsAPI.md#CreateConnectorAuthorizationUri) | **Post** /api/connectors/{connectorId}/authorization-uri | Get connector&#39;s authorization URI
[**CreateConnectorTest**](ConnectorsAPI.md#CreateConnectorTest) | **Post** /api/connectors/{factoryId}/test | Test passwordless connector
[**DeleteConnector**](ConnectorsAPI.md#DeleteConnector) | **Delete** /api/connectors/{id} | Delete connector
[**GetConnector**](ConnectorsAPI.md#GetConnector) | **Get** /api/connectors/{id} | Get connector
[**ListConnectors**](ConnectorsAPI.md#ListConnectors) | **Get** /api/connectors | Get connectors
[**UpdateConnector**](ConnectorsAPI.md#UpdateConnector) | **Patch** /api/connectors/{id} | Update connector



## CreateConnector

> ListConnectors200ResponseInner CreateConnector(ctx).CreateConnectorRequest(createConnectorRequest).Execute()

Create connector



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
	createConnectorRequest := *openapiclient.NewCreateConnectorRequest("ConnectorId_example") // CreateConnectorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.CreateConnector(context.Background()).CreateConnectorRequest(createConnectorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CreateConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnector`: ListConnectors200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.CreateConnector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConnectorRequest** | [**CreateConnectorRequest**](CreateConnectorRequest.md) |  | 

### Return type

[**ListConnectors200ResponseInner**](ListConnectors200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectorAuthorizationUri

> CreateConnectorAuthorizationUri200Response CreateConnectorAuthorizationUri(ctx, connectorId).CreateConnectorAuthorizationUriRequest(createConnectorAuthorizationUriRequest).Execute()

Get connector's authorization URI



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
	connectorId := "connectorId_example" // string | The unique identifier of the connector.
	createConnectorAuthorizationUriRequest := *openapiclient.NewCreateConnectorAuthorizationUriRequest("State_example", "RedirectUri_example") // CreateConnectorAuthorizationUriRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.CreateConnectorAuthorizationUri(context.Background(), connectorId).CreateConnectorAuthorizationUriRequest(createConnectorAuthorizationUriRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CreateConnectorAuthorizationUri``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectorAuthorizationUri`: CreateConnectorAuthorizationUri200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.CreateConnectorAuthorizationUri`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorAuthorizationUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createConnectorAuthorizationUriRequest** | [**CreateConnectorAuthorizationUriRequest**](CreateConnectorAuthorizationUriRequest.md) |  | 

### Return type

[**CreateConnectorAuthorizationUri200Response**](CreateConnectorAuthorizationUri200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectorTest

> CreateConnectorTest(ctx, factoryId).CreateConnectorTestRequest(createConnectorTestRequest).Execute()

Test passwordless connector



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
	factoryId := "factoryId_example" // string | The unique identifier of the factory.
	createConnectorTestRequest := *openapiclient.NewCreateConnectorTestRequest(map[string]interface{}(123)) // CreateConnectorTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorsAPI.CreateConnectorTest(context.Background(), factoryId).CreateConnectorTestRequest(createConnectorTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CreateConnectorTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**factoryId** | **string** | The unique identifier of the factory. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createConnectorTestRequest** | [**CreateConnectorTestRequest**](CreateConnectorTestRequest.md) |  | 

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


## DeleteConnector

> DeleteConnector(ctx, id).Execute()

Delete connector



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
	id := "id_example" // string | The unique identifier of the connector.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorsAPI.DeleteConnector(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.DeleteConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorRequest struct via the builder pattern


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


## GetConnector

> ListConnectors200ResponseInner GetConnector(ctx, id).Execute()

Get connector



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
	id := "id_example" // string | The unique identifier of the connector.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.GetConnector(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.GetConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnector`: ListConnectors200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.GetConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListConnectors200ResponseInner**](ListConnectors200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectors

> []ListConnectors200ResponseInner ListConnectors(ctx).Target(target).Execute()

Get connectors



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
	target := "target_example" // string | Filter connectors by target. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.ListConnectors(context.Background()).Target(target).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.ListConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnectors`: []ListConnectors200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.ListConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **target** | **string** | Filter connectors by target. | 

### Return type

[**[]ListConnectors200ResponseInner**](ListConnectors200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnector

> ListConnectors200ResponseInner UpdateConnector(ctx, id).UpdateConnectorRequest(updateConnectorRequest).Execute()

Update connector



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
	id := "id_example" // string | The unique identifier of the connector.
	updateConnectorRequest := *openapiclient.NewUpdateConnectorRequest() // UpdateConnectorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.UpdateConnector(context.Background(), id).UpdateConnectorRequest(updateConnectorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.UpdateConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConnector`: ListConnectors200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.UpdateConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateConnectorRequest** | [**UpdateConnectorRequest**](UpdateConnectorRequest.md) |  | 

### Return type

[**ListConnectors200ResponseInner**](ListConnectors200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

