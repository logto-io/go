# \OrganizationRolesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationRole**](OrganizationRolesAPI.md#CreateOrganizationRole) | **Post** /api/organization-roles | Create an organization role
[**CreateOrganizationRoleResourceScope**](OrganizationRolesAPI.md#CreateOrganizationRoleResourceScope) | **Post** /api/organization-roles/{id}/resource-scopes | Assign resource scopes to organization role
[**CreateOrganizationRoleScope**](OrganizationRolesAPI.md#CreateOrganizationRoleScope) | **Post** /api/organization-roles/{id}/scopes | Assign organization scopes to organization role
[**DeleteOrganizationRole**](OrganizationRolesAPI.md#DeleteOrganizationRole) | **Delete** /api/organization-roles/{id} | Delete organization role
[**DeleteOrganizationRoleResourceScope**](OrganizationRolesAPI.md#DeleteOrganizationRoleResourceScope) | **Delete** /api/organization-roles/{id}/resource-scopes/{scopeId} | Remove resource scope
[**DeleteOrganizationRoleScope**](OrganizationRolesAPI.md#DeleteOrganizationRoleScope) | **Delete** /api/organization-roles/{id}/scopes/{organizationScopeId} | Remove organization scope
[**GetOrganizationRole**](OrganizationRolesAPI.md#GetOrganizationRole) | **Get** /api/organization-roles/{id} | Get organization role
[**ListOrganizationRoleResourceScopes**](OrganizationRolesAPI.md#ListOrganizationRoleResourceScopes) | **Get** /api/organization-roles/{id}/resource-scopes | Get organization role resource scopes
[**ListOrganizationRoleScopes**](OrganizationRolesAPI.md#ListOrganizationRoleScopes) | **Get** /api/organization-roles/{id}/scopes | Get organization role scopes
[**ListOrganizationRoles**](OrganizationRolesAPI.md#ListOrganizationRoles) | **Get** /api/organization-roles | Get organization roles
[**ReplaceOrganizationRoleResourceScopes**](OrganizationRolesAPI.md#ReplaceOrganizationRoleResourceScopes) | **Put** /api/organization-roles/{id}/resource-scopes | Replace resource scopes for organization role
[**ReplaceOrganizationRoleScopes**](OrganizationRolesAPI.md#ReplaceOrganizationRoleScopes) | **Put** /api/organization-roles/{id}/scopes | Replace organization scopes for organization role
[**UpdateOrganizationRole**](OrganizationRolesAPI.md#UpdateOrganizationRole) | **Patch** /api/organization-roles/{id} | Update organization role



## CreateOrganizationRole

> GetOrganizationRole200Response CreateOrganizationRole(ctx).CreateOrganizationRoleRequest(createOrganizationRoleRequest).Execute()

Create an organization role



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
	createOrganizationRoleRequest := *openapiclient.NewCreateOrganizationRoleRequest("Name_example", []string{"OrganizationScopeIds_example"}, []string{"ResourceScopeIds_example"}) // CreateOrganizationRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationRolesAPI.CreateOrganizationRole(context.Background()).CreateOrganizationRoleRequest(createOrganizationRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.CreateOrganizationRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationRole`: GetOrganizationRole200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationRolesAPI.CreateOrganizationRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganizationRoleRequest** | [**CreateOrganizationRoleRequest**](CreateOrganizationRoleRequest.md) |  | 

### Return type

[**GetOrganizationRole200Response**](GetOrganizationRole200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationRoleResourceScope

> CreateOrganizationRoleResourceScope(ctx, id).CreateOrganizationRoleResourceScopeRequest(createOrganizationRoleResourceScopeRequest).Execute()

Assign resource scopes to organization role



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
	id := "id_example" // string | The unique identifier of the organization role.
	createOrganizationRoleResourceScopeRequest := *openapiclient.NewCreateOrganizationRoleResourceScopeRequest([]string{"ScopeIds_example"}) // CreateOrganizationRoleResourceScopeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationRolesAPI.CreateOrganizationRoleResourceScope(context.Background(), id).CreateOrganizationRoleResourceScopeRequest(createOrganizationRoleResourceScopeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.CreateOrganizationRoleResourceScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRoleResourceScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationRoleResourceScopeRequest** | [**CreateOrganizationRoleResourceScopeRequest**](CreateOrganizationRoleResourceScopeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationRoleScope

> CreateOrganizationRoleScope(ctx, id).CreateOrganizationRoleScopeRequest(createOrganizationRoleScopeRequest).Execute()

Assign organization scopes to organization role



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
	id := "id_example" // string | The unique identifier of the organization role.
	createOrganizationRoleScopeRequest := *openapiclient.NewCreateOrganizationRoleScopeRequest([]string{"OrganizationScopeIds_example"}) // CreateOrganizationRoleScopeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationRolesAPI.CreateOrganizationRoleScope(context.Background(), id).CreateOrganizationRoleScopeRequest(createOrganizationRoleScopeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.CreateOrganizationRoleScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRoleScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationRoleScopeRequest** | [**CreateOrganizationRoleScopeRequest**](CreateOrganizationRoleScopeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationRole

> DeleteOrganizationRole(ctx, id).Execute()

Delete organization role



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
	id := "id_example" // string | The unique identifier of the organization role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationRolesAPI.DeleteOrganizationRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.DeleteOrganizationRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRoleRequest struct via the builder pattern


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


## DeleteOrganizationRoleResourceScope

> DeleteOrganizationRoleResourceScope(ctx, id, scopeId).Execute()

Remove resource scope



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
	id := "id_example" // string | The unique identifier of the organization role.
	scopeId := "scopeId_example" // string | The unique identifier of the scope.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationRolesAPI.DeleteOrganizationRoleResourceScope(context.Background(), id, scopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.DeleteOrganizationRoleResourceScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization role. | 
**scopeId** | **string** | The unique identifier of the scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRoleResourceScopeRequest struct via the builder pattern


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


## DeleteOrganizationRoleScope

> DeleteOrganizationRoleScope(ctx, id, organizationScopeId).Execute()

Remove organization scope



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
	id := "id_example" // string | The unique identifier of the organization role.
	organizationScopeId := "organizationScopeId_example" // string | The unique identifier of the organization scope.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationRolesAPI.DeleteOrganizationRoleScope(context.Background(), id, organizationScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.DeleteOrganizationRoleScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization role. | 
**organizationScopeId** | **string** | The unique identifier of the organization scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRoleScopeRequest struct via the builder pattern


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


## GetOrganizationRole

> GetOrganizationRole200Response GetOrganizationRole(ctx, id).Execute()

Get organization role



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
	id := "id_example" // string | The unique identifier of the organization role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationRolesAPI.GetOrganizationRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.GetOrganizationRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationRole`: GetOrganizationRole200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationRolesAPI.GetOrganizationRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationRole200Response**](GetOrganizationRole200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationRoleResourceScopes

> []ListResources200ResponseInnerScopesInner ListOrganizationRoleResourceScopes(ctx, id).Page(page).PageSize(pageSize).Execute()

Get organization role resource scopes



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
	id := "id_example" // string | The unique identifier of the organization role.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationRolesAPI.ListOrganizationRoleResourceScopes(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.ListOrganizationRoleResourceScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationRoleResourceScopes`: []ListResources200ResponseInnerScopesInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationRolesAPI.ListOrganizationRoleResourceScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationRoleResourceScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

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


## ListOrganizationRoleScopes

> []ListOrganizationRoleScopes200ResponseInner ListOrganizationRoleScopes(ctx, id).Page(page).PageSize(pageSize).Execute()

Get organization role scopes



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
	id := "id_example" // string | The unique identifier of the organization role.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationRolesAPI.ListOrganizationRoleScopes(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.ListOrganizationRoleScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationRoleScopes`: []ListOrganizationRoleScopes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationRolesAPI.ListOrganizationRoleScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationRoleScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## ListOrganizationRoles

> []ListOrganizationRoles200ResponseInner ListOrganizationRoles(ctx).Q(q).Page(page).PageSize(pageSize).Execute()

Get organization roles



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
	resp, r, err := apiClient.OrganizationRolesAPI.ListOrganizationRoles(context.Background()).Q(q).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.ListOrganizationRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationRoles`: []ListOrganizationRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationRolesAPI.ListOrganizationRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListOrganizationRoles200ResponseInner**](ListOrganizationRoles200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOrganizationRoleResourceScopes

> ReplaceOrganizationRoleResourceScopes(ctx, id).ReplaceOrganizationRoleResourceScopesRequest(replaceOrganizationRoleResourceScopesRequest).Execute()

Replace resource scopes for organization role



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
	id := "id_example" // string | The unique identifier of the organization role.
	replaceOrganizationRoleResourceScopesRequest := *openapiclient.NewReplaceOrganizationRoleResourceScopesRequest([]string{"ScopeIds_example"}) // ReplaceOrganizationRoleResourceScopesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationRolesAPI.ReplaceOrganizationRoleResourceScopes(context.Background(), id).ReplaceOrganizationRoleResourceScopesRequest(replaceOrganizationRoleResourceScopesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.ReplaceOrganizationRoleResourceScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrganizationRoleResourceScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceOrganizationRoleResourceScopesRequest** | [**ReplaceOrganizationRoleResourceScopesRequest**](ReplaceOrganizationRoleResourceScopesRequest.md) |  | 

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


## ReplaceOrganizationRoleScopes

> ReplaceOrganizationRoleScopes(ctx, id).ReplaceOrganizationRoleScopesRequest(replaceOrganizationRoleScopesRequest).Execute()

Replace organization scopes for organization role



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
	id := "id_example" // string | The unique identifier of the organization role.
	replaceOrganizationRoleScopesRequest := *openapiclient.NewReplaceOrganizationRoleScopesRequest([]string{"OrganizationScopeIds_example"}) // ReplaceOrganizationRoleScopesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationRolesAPI.ReplaceOrganizationRoleScopes(context.Background(), id).ReplaceOrganizationRoleScopesRequest(replaceOrganizationRoleScopesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.ReplaceOrganizationRoleScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrganizationRoleScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceOrganizationRoleScopesRequest** | [**ReplaceOrganizationRoleScopesRequest**](ReplaceOrganizationRoleScopesRequest.md) |  | 

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


## UpdateOrganizationRole

> GetOrganizationRole200Response UpdateOrganizationRole(ctx, id).UpdateOrganizationRoleRequest(updateOrganizationRoleRequest).Execute()

Update organization role



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
	id := "id_example" // string | The unique identifier of the organization role.
	updateOrganizationRoleRequest := *openapiclient.NewUpdateOrganizationRoleRequest() // UpdateOrganizationRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationRolesAPI.UpdateOrganizationRole(context.Background(), id).UpdateOrganizationRoleRequest(updateOrganizationRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.UpdateOrganizationRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationRole`: GetOrganizationRole200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationRolesAPI.UpdateOrganizationRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationRoleRequest** | [**UpdateOrganizationRoleRequest**](UpdateOrganizationRoleRequest.md) |  | 

### Return type

[**GetOrganizationRole200Response**](GetOrganizationRole200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

