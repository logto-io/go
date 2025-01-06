# \ResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResource**](ResourcesAPI.md#CreateResource) | **Post** /api/resources | Create an API resource
[**CreateResourceScope**](ResourcesAPI.md#CreateResourceScope) | **Post** /api/resources/{resourceId}/scopes | Create API resource scope
[**DeleteResource**](ResourcesAPI.md#DeleteResource) | **Delete** /api/resources/{id} | Delete API resource
[**DeleteResourceScope**](ResourcesAPI.md#DeleteResourceScope) | **Delete** /api/resources/{resourceId}/scopes/{scopeId} | Delete API resource scope
[**GetResource**](ResourcesAPI.md#GetResource) | **Get** /api/resources/{id} | Get API resource
[**ListResourceScopes**](ResourcesAPI.md#ListResourceScopes) | **Get** /api/resources/{resourceId}/scopes | Get API resource scopes
[**ListResources**](ResourcesAPI.md#ListResources) | **Get** /api/resources | Get API resources
[**UpdateResource**](ResourcesAPI.md#UpdateResource) | **Patch** /api/resources/{id} | Update API resource
[**UpdateResourceIsDefault**](ResourcesAPI.md#UpdateResourceIsDefault) | **Patch** /api/resources/{id}/is-default | Set API resource as default
[**UpdateResourceScope**](ResourcesAPI.md#UpdateResourceScope) | **Patch** /api/resources/{resourceId}/scopes/{scopeId} | Update API resource scope



## CreateResource

> ListResources200ResponseInner CreateResource(ctx).CreateResourceRequest(createResourceRequest).Execute()

Create an API resource



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
	createResourceRequest := *openapiclient.NewCreateResourceRequest("Name_example", "Indicator_example") // CreateResourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.CreateResource(context.Background()).CreateResourceRequest(createResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.CreateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResource`: ListResources200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.CreateResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createResourceRequest** | [**CreateResourceRequest**](CreateResourceRequest.md) |  | 

### Return type

[**ListResources200ResponseInner**](ListResources200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateResourceScope

> ListResources200ResponseInnerScopesInner CreateResourceScope(ctx, resourceId).CreateResourceScopeRequest(createResourceScopeRequest).Execute()

Create API resource scope



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
	resourceId := "resourceId_example" // string | The unique identifier of the resource.
	createResourceScopeRequest := *openapiclient.NewCreateResourceScopeRequest("Name_example") // CreateResourceScopeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.CreateResourceScope(context.Background(), resourceId).CreateResourceScopeRequest(createResourceScopeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.CreateResourceScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResourceScope`: ListResources200ResponseInnerScopesInner
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.CreateResourceScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createResourceScopeRequest** | [**CreateResourceScopeRequest**](CreateResourceScopeRequest.md) |  | 

### Return type

[**ListResources200ResponseInnerScopesInner**](ListResources200ResponseInnerScopesInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResource

> DeleteResource(ctx, id).Execute()

Delete API resource



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
	id := "id_example" // string | The unique identifier of the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcesAPI.DeleteResource(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.DeleteResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceRequest struct via the builder pattern


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


## DeleteResourceScope

> DeleteResourceScope(ctx, resourceId, scopeId).Execute()

Delete API resource scope



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
	resourceId := "resourceId_example" // string | The unique identifier of the resource.
	scopeId := "scopeId_example" // string | The unique identifier of the scope.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcesAPI.DeleteResourceScope(context.Background(), resourceId, scopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.DeleteResourceScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The unique identifier of the resource. | 
**scopeId** | **string** | The unique identifier of the scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceScopeRequest struct via the builder pattern


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


## GetResource

> GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource GetResource(ctx, id).Execute()

Get API resource



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
	id := "id_example" // string | The unique identifier of the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.GetResource(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.GetResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResource`: GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.GetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource**](GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceScopes

> []ListResources200ResponseInnerScopesInner ListResourceScopes(ctx, resourceId).Page(page).PageSize(pageSize).SearchParams(searchParams).Execute()

Get API resource scopes



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
	resourceId := "resourceId_example" // string | The unique identifier of the resource.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)
	searchParams := map[string]string{"key": "Inner_example"} // map[string]string | Search query parameters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.ListResourceScopes(context.Background(), resourceId).Page(page).PageSize(pageSize).SearchParams(searchParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.ListResourceScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceScopes`: []ListResources200ResponseInnerScopesInner
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.ListResourceScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResourceScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]
 **searchParams** | **map[string]string** | Search query parameters. | 

### Return type

[**[]ListResources200ResponseInnerScopesInner**](ListResources200ResponseInnerScopesInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResources

> []ListResources200ResponseInner ListResources(ctx).IncludeScopes(includeScopes).Page(page).PageSize(pageSize).Execute()

Get API resources



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
	includeScopes := "includeScopes_example" // string | If it's provided with a truthy value (`true`, `1`, `yes`), the scopes of each resource will be included in the response. (optional)
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.ListResources(context.Background()).IncludeScopes(includeScopes).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.ListResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResources`: []ListResources200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.ListResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeScopes** | **string** | If it&#39;s provided with a truthy value (&#x60;true&#x60;, &#x60;1&#x60;, &#x60;yes&#x60;), the scopes of each resource will be included in the response. | 
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListResources200ResponseInner**](ListResources200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResource

> GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource UpdateResource(ctx, id).UpdateResourceRequest(updateResourceRequest).Execute()

Update API resource



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
	id := "id_example" // string | The unique identifier of the resource.
	updateResourceRequest := *openapiclient.NewUpdateResourceRequest() // UpdateResourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.UpdateResource(context.Background(), id).UpdateResourceRequest(updateResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.UpdateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResource`: GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.UpdateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateResourceRequest** | [**UpdateResourceRequest**](UpdateResourceRequest.md) |  | 

### Return type

[**GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource**](GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceIsDefault

> GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource UpdateResourceIsDefault(ctx, id).UpdateResourceIsDefaultRequest(updateResourceIsDefaultRequest).Execute()

Set API resource as default



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
	id := "id_example" // string | The unique identifier of the resource.
	updateResourceIsDefaultRequest := *openapiclient.NewUpdateResourceIsDefaultRequest(false) // UpdateResourceIsDefaultRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.UpdateResourceIsDefault(context.Background(), id).UpdateResourceIsDefaultRequest(updateResourceIsDefaultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.UpdateResourceIsDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResourceIsDefault`: GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.UpdateResourceIsDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceIsDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateResourceIsDefaultRequest** | [**UpdateResourceIsDefaultRequest**](UpdateResourceIsDefaultRequest.md) |  | 

### Return type

[**GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource**](GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceScope

> ListResources200ResponseInnerScopesInner UpdateResourceScope(ctx, resourceId, scopeId).UpdateResourceScopeRequest(updateResourceScopeRequest).Execute()

Update API resource scope



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
	resourceId := "resourceId_example" // string | The unique identifier of the resource.
	scopeId := "scopeId_example" // string | The unique identifier of the scope.
	updateResourceScopeRequest := *openapiclient.NewUpdateResourceScopeRequest() // UpdateResourceScopeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.UpdateResourceScope(context.Background(), resourceId, scopeId).UpdateResourceScopeRequest(updateResourceScopeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.UpdateResourceScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResourceScope`: ListResources200ResponseInnerScopesInner
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.UpdateResourceScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The unique identifier of the resource. | 
**scopeId** | **string** | The unique identifier of the scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateResourceScopeRequest** | [**UpdateResourceScopeRequest**](UpdateResourceScopeRequest.md) |  | 

### Return type

[**ListResources200ResponseInnerScopesInner**](ListResources200ResponseInnerScopesInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

