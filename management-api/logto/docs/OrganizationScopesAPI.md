# \OrganizationScopesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationScope**](OrganizationScopesAPI.md#CreateOrganizationScope) | **Post** /api/organization-scopes | Create an organization scope
[**DeleteOrganizationScope**](OrganizationScopesAPI.md#DeleteOrganizationScope) | **Delete** /api/organization-scopes/{id} | Delete organization scope
[**GetOrganizationScope**](OrganizationScopesAPI.md#GetOrganizationScope) | **Get** /api/organization-scopes/{id} | Get organization scope
[**ListOrganizationScopes**](OrganizationScopesAPI.md#ListOrganizationScopes) | **Get** /api/organization-scopes | Get organization scopes
[**UpdateOrganizationScope**](OrganizationScopesAPI.md#UpdateOrganizationScope) | **Patch** /api/organization-scopes/{id} | Update organization scope



## CreateOrganizationScope

> ListOrganizationRoleScopes200ResponseInner CreateOrganizationScope(ctx).CreateOrganizationScopeRequest(createOrganizationScopeRequest).Execute()

Create an organization scope



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
	createOrganizationScopeRequest := *openapiclient.NewCreateOrganizationScopeRequest("Name_example") // CreateOrganizationScopeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationScopesAPI.CreateOrganizationScope(context.Background()).CreateOrganizationScopeRequest(createOrganizationScopeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationScopesAPI.CreateOrganizationScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationScope`: ListOrganizationRoleScopes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationScopesAPI.CreateOrganizationScope`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganizationScopeRequest** | [**CreateOrganizationScopeRequest**](CreateOrganizationScopeRequest.md) |  | 

### Return type

[**ListOrganizationRoleScopes200ResponseInner**](ListOrganizationRoleScopes200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationScope

> DeleteOrganizationScope(ctx, id).Execute()

Delete organization scope



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
	id := "id_example" // string | The unique identifier of the organization scope.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationScopesAPI.DeleteOrganizationScope(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationScopesAPI.DeleteOrganizationScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationScopeRequest struct via the builder pattern


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


## GetOrganizationScope

> ListOrganizationRoleScopes200ResponseInner GetOrganizationScope(ctx, id).Execute()

Get organization scope



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
	id := "id_example" // string | The unique identifier of the organization scope.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationScopesAPI.GetOrganizationScope(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationScopesAPI.GetOrganizationScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationScope`: ListOrganizationRoleScopes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationScopesAPI.GetOrganizationScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListOrganizationRoleScopes200ResponseInner**](ListOrganizationRoleScopes200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationScopes

> []ListOrganizationRoleScopes200ResponseInner ListOrganizationScopes(ctx).Q(q).Page(page).PageSize(pageSize).Execute()

Get organization scopes



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
	q := "q_example" // string |  (optional)
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationScopesAPI.ListOrganizationScopes(context.Background()).Q(q).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationScopesAPI.ListOrganizationScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationScopes`: []ListOrganizationRoleScopes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationScopesAPI.ListOrganizationScopes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListOrganizationRoleScopes200ResponseInner**](ListOrganizationRoleScopes200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationScope

> ListOrganizationRoleScopes200ResponseInner UpdateOrganizationScope(ctx, id).UpdateOrganizationScopeRequest(updateOrganizationScopeRequest).Execute()

Update organization scope



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
	id := "id_example" // string | The unique identifier of the organization scope.
	updateOrganizationScopeRequest := *openapiclient.NewUpdateOrganizationScopeRequest() // UpdateOrganizationScopeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationScopesAPI.UpdateOrganizationScope(context.Background(), id).UpdateOrganizationScopeRequest(updateOrganizationScopeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationScopesAPI.UpdateOrganizationScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationScope`: ListOrganizationRoleScopes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationScopesAPI.UpdateOrganizationScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationScopeRequest** | [**UpdateOrganizationScopeRequest**](UpdateOrganizationScopeRequest.md) |  | 

### Return type

[**ListOrganizationRoleScopes200ResponseInner**](ListOrganizationRoleScopes200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

