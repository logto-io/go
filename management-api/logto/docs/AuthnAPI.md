# \AuthnAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssertSaml**](AuthnAPI.md#AssertSaml) | **Post** /api/authn/saml/{connectorId} | SAML ACS endpoint (social)
[**AssertSingleSignOnSaml**](AuthnAPI.md#AssertSingleSignOnSaml) | **Post** /api/authn/single-sign-on/saml/{connectorId} | SAML ACS endpoint (SSO)
[**GetHasuraAuth**](AuthnAPI.md#GetHasuraAuth) | **Get** /api/authn/hasura | Hasura auth hook endpoint



## AssertSaml

> AssertSaml(ctx, connectorId).Body(body).Execute()

SAML ACS endpoint (social)



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthnAPI.AssertSaml(context.Background(), connectorId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthnAPI.AssertSaml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssertSamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssertSingleSignOnSaml

> AssertSingleSignOnSaml(ctx, connectorId).AssertSingleSignOnSamlRequest(assertSingleSignOnSamlRequest).Execute()

SAML ACS endpoint (SSO)



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
	assertSingleSignOnSamlRequest := *openapiclient.NewAssertSingleSignOnSamlRequest("SAMLResponse_example") // AssertSingleSignOnSamlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthnAPI.AssertSingleSignOnSaml(context.Background(), connectorId).AssertSingleSignOnSamlRequest(assertSingleSignOnSamlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthnAPI.AssertSingleSignOnSaml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssertSingleSignOnSamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assertSingleSignOnSamlRequest** | [**AssertSingleSignOnSamlRequest**](AssertSingleSignOnSamlRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHasuraAuth

> GetHasuraAuth200Response GetHasuraAuth(ctx).Resource(resource).UnauthorizedRole(unauthorizedRole).Execute()

Hasura auth hook endpoint



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
	resource := "resource_example" // string | 
	unauthorizedRole := "unauthorizedRole_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthnAPI.GetHasuraAuth(context.Background()).Resource(resource).UnauthorizedRole(unauthorizedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthnAPI.GetHasuraAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHasuraAuth`: GetHasuraAuth200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthnAPI.GetHasuraAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHasuraAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resource** | **string** |  | 
 **unauthorizedRole** | **string** |  | 

### Return type

[**GetHasuraAuth200Response**](GetHasuraAuth200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

