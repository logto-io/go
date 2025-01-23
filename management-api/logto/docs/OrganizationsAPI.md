# \OrganizationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrganizationApplications**](OrganizationsAPI.md#AddOrganizationApplications) | **Post** /api/organizations/{id}/applications | Add organization application
[**AddOrganizationUsers**](OrganizationsAPI.md#AddOrganizationUsers) | **Post** /api/organizations/{id}/users | Add user members to organization
[**AssignOrganizationRolesToApplication**](OrganizationsAPI.md#AssignOrganizationRolesToApplication) | **Post** /api/organizations/{id}/applications/{applicationId}/roles | Add organization application role
[**AssignOrganizationRolesToApplications**](OrganizationsAPI.md#AssignOrganizationRolesToApplications) | **Post** /api/organizations/{id}/applications/roles | Assign roles to applications in an organization
[**AssignOrganizationRolesToUser**](OrganizationsAPI.md#AssignOrganizationRolesToUser) | **Post** /api/organizations/{id}/users/{userId}/roles | Assign roles to a user in an organization
[**AssignOrganizationRolesToUsers**](OrganizationsAPI.md#AssignOrganizationRolesToUsers) | **Post** /api/organizations/{id}/users/roles | Assign roles to organization user members
[**CreateOrganization**](OrganizationsAPI.md#CreateOrganization) | **Post** /api/organizations | Create an organization
[**CreateOrganizationJitEmailDomain**](OrganizationsAPI.md#CreateOrganizationJitEmailDomain) | **Post** /api/organizations/{id}/jit/email-domains | Add organization JIT email domain
[**CreateOrganizationJitRole**](OrganizationsAPI.md#CreateOrganizationJitRole) | **Post** /api/organizations/{id}/jit/roles | Add organization JIT default roles
[**CreateOrganizationJitSsoConnector**](OrganizationsAPI.md#CreateOrganizationJitSsoConnector) | **Post** /api/organizations/{id}/jit/sso-connectors | Add organization JIT SSO connectors
[**DeleteOrganization**](OrganizationsAPI.md#DeleteOrganization) | **Delete** /api/organizations/{id} | Delete organization
[**DeleteOrganizationApplication**](OrganizationsAPI.md#DeleteOrganizationApplication) | **Delete** /api/organizations/{id}/applications/{applicationId} | Remove organization application
[**DeleteOrganizationApplicationRole**](OrganizationsAPI.md#DeleteOrganizationApplicationRole) | **Delete** /api/organizations/{id}/applications/{applicationId}/roles/{organizationRoleId} | Remove organization application role
[**DeleteOrganizationJitEmailDomain**](OrganizationsAPI.md#DeleteOrganizationJitEmailDomain) | **Delete** /api/organizations/{id}/jit/email-domains/{emailDomain} | Remove organization JIT email domain
[**DeleteOrganizationJitRole**](OrganizationsAPI.md#DeleteOrganizationJitRole) | **Delete** /api/organizations/{id}/jit/roles/{organizationRoleId} | Remove organization JIT default role
[**DeleteOrganizationJitSsoConnector**](OrganizationsAPI.md#DeleteOrganizationJitSsoConnector) | **Delete** /api/organizations/{id}/jit/sso-connectors/{ssoConnectorId} | Remove organization JIT SSO connector
[**DeleteOrganizationUser**](OrganizationsAPI.md#DeleteOrganizationUser) | **Delete** /api/organizations/{id}/users/{userId} | Remove user member from organization
[**DeleteOrganizationUserRole**](OrganizationsAPI.md#DeleteOrganizationUserRole) | **Delete** /api/organizations/{id}/users/{userId}/roles/{organizationRoleId} | Remove a role from a user in an organization
[**GetOrganization**](OrganizationsAPI.md#GetOrganization) | **Get** /api/organizations/{id} | Get organization
[**ListOrganizationApplicationRoles**](OrganizationsAPI.md#ListOrganizationApplicationRoles) | **Get** /api/organizations/{id}/applications/{applicationId}/roles | Get organization application roles
[**ListOrganizationApplications**](OrganizationsAPI.md#ListOrganizationApplications) | **Get** /api/organizations/{id}/applications | Get organization applications
[**ListOrganizationJitEmailDomains**](OrganizationsAPI.md#ListOrganizationJitEmailDomains) | **Get** /api/organizations/{id}/jit/email-domains | Get organization JIT email domains
[**ListOrganizationJitRoles**](OrganizationsAPI.md#ListOrganizationJitRoles) | **Get** /api/organizations/{id}/jit/roles | Get organization JIT default roles
[**ListOrganizationJitSsoConnectors**](OrganizationsAPI.md#ListOrganizationJitSsoConnectors) | **Get** /api/organizations/{id}/jit/sso-connectors | Get organization JIT SSO connectors
[**ListOrganizationUserRoles**](OrganizationsAPI.md#ListOrganizationUserRoles) | **Get** /api/organizations/{id}/users/{userId}/roles | Get roles for a user in an organization
[**ListOrganizationUserScopes**](OrganizationsAPI.md#ListOrganizationUserScopes) | **Get** /api/organizations/{id}/users/{userId}/scopes | Get scopes for a user in an organization tailored by the organization roles
[**ListOrganizationUsers**](OrganizationsAPI.md#ListOrganizationUsers) | **Get** /api/organizations/{id}/users | Get organization user members
[**ListOrganizations**](OrganizationsAPI.md#ListOrganizations) | **Get** /api/organizations | Get organizations
[**ReplaceOrganizationApplicationRoles**](OrganizationsAPI.md#ReplaceOrganizationApplicationRoles) | **Put** /api/organizations/{id}/applications/{applicationId}/roles | Replace organization application roles
[**ReplaceOrganizationApplications**](OrganizationsAPI.md#ReplaceOrganizationApplications) | **Put** /api/organizations/{id}/applications | Replace organization applications
[**ReplaceOrganizationJitEmailDomains**](OrganizationsAPI.md#ReplaceOrganizationJitEmailDomains) | **Put** /api/organizations/{id}/jit/email-domains | Replace organization JIT email domains
[**ReplaceOrganizationJitRoles**](OrganizationsAPI.md#ReplaceOrganizationJitRoles) | **Put** /api/organizations/{id}/jit/roles | Replace organization JIT default roles
[**ReplaceOrganizationJitSsoConnectors**](OrganizationsAPI.md#ReplaceOrganizationJitSsoConnectors) | **Put** /api/organizations/{id}/jit/sso-connectors | Replace organization JIT SSO connectors
[**ReplaceOrganizationUserRoles**](OrganizationsAPI.md#ReplaceOrganizationUserRoles) | **Put** /api/organizations/{id}/users/{userId}/roles | Update roles for a user in an organization
[**ReplaceOrganizationUsers**](OrganizationsAPI.md#ReplaceOrganizationUsers) | **Put** /api/organizations/{id}/users | Replace organization user members
[**UpdateOrganization**](OrganizationsAPI.md#UpdateOrganization) | **Patch** /api/organizations/{id} | Update organization



## AddOrganizationApplications

> AddOrganizationApplications(ctx, id).AddOrganizationApplicationsRequest(addOrganizationApplicationsRequest).Execute()

Add organization application



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
	id := "id_example" // string | The unique identifier of the organization.
	addOrganizationApplicationsRequest := *openapiclient.NewAddOrganizationApplicationsRequest([]string{"ApplicationIds_example"}) // AddOrganizationApplicationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AddOrganizationApplications(context.Background(), id).AddOrganizationApplicationsRequest(addOrganizationApplicationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AddOrganizationApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrganizationApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addOrganizationApplicationsRequest** | [**AddOrganizationApplicationsRequest**](AddOrganizationApplicationsRequest.md) |  | 

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


## AddOrganizationUsers

> AddOrganizationUsers(ctx, id).AddOrganizationUsersRequest(addOrganizationUsersRequest).Execute()

Add user members to organization



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
	id := "id_example" // string | The unique identifier of the organization.
	addOrganizationUsersRequest := *openapiclient.NewAddOrganizationUsersRequest([]string{"UserIds_example"}) // AddOrganizationUsersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AddOrganizationUsers(context.Background(), id).AddOrganizationUsersRequest(addOrganizationUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AddOrganizationUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrganizationUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addOrganizationUsersRequest** | [**AddOrganizationUsersRequest**](AddOrganizationUsersRequest.md) |  | 

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


## AssignOrganizationRolesToApplication

> AssignOrganizationRolesToApplication(ctx, id, applicationId).AssignOrganizationRolesToApplicationRequest(assignOrganizationRolesToApplicationRequest).Execute()

Add organization application role



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
	id := "id_example" // string | The unique identifier of the organization.
	applicationId := "applicationId_example" // string | The unique identifier of the application.
	assignOrganizationRolesToApplicationRequest := *openapiclient.NewAssignOrganizationRolesToApplicationRequest([]string{"OrganizationRoleIds_example"}) // AssignOrganizationRolesToApplicationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AssignOrganizationRolesToApplication(context.Background(), id, applicationId).AssignOrganizationRolesToApplicationRequest(assignOrganizationRolesToApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AssignOrganizationRolesToApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOrganizationRolesToApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assignOrganizationRolesToApplicationRequest** | [**AssignOrganizationRolesToApplicationRequest**](AssignOrganizationRolesToApplicationRequest.md) |  | 

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


## AssignOrganizationRolesToApplications

> AssignOrganizationRolesToApplications(ctx, id).AssignOrganizationRolesToApplicationsRequest(assignOrganizationRolesToApplicationsRequest).Execute()

Assign roles to applications in an organization



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
	id := "id_example" // string | The unique identifier of the organization.
	assignOrganizationRolesToApplicationsRequest := *openapiclient.NewAssignOrganizationRolesToApplicationsRequest([]string{"ApplicationIds_example"}, []string{"OrganizationRoleIds_example"}) // AssignOrganizationRolesToApplicationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AssignOrganizationRolesToApplications(context.Background(), id).AssignOrganizationRolesToApplicationsRequest(assignOrganizationRolesToApplicationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AssignOrganizationRolesToApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOrganizationRolesToApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignOrganizationRolesToApplicationsRequest** | [**AssignOrganizationRolesToApplicationsRequest**](AssignOrganizationRolesToApplicationsRequest.md) |  | 

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


## AssignOrganizationRolesToUser

> AssignOrganizationRolesToUser(ctx, id, userId).AssignOrganizationRolesToUserRequest(assignOrganizationRolesToUserRequest).Execute()

Assign roles to a user in an organization



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
	id := "id_example" // string | The unique identifier of the organization.
	userId := "userId_example" // string | The unique identifier of the user.
	assignOrganizationRolesToUserRequest := *openapiclient.NewAssignOrganizationRolesToUserRequest([]string{"OrganizationRoleIds_example"}) // AssignOrganizationRolesToUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AssignOrganizationRolesToUser(context.Background(), id, userId).AssignOrganizationRolesToUserRequest(assignOrganizationRolesToUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AssignOrganizationRolesToUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOrganizationRolesToUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assignOrganizationRolesToUserRequest** | [**AssignOrganizationRolesToUserRequest**](AssignOrganizationRolesToUserRequest.md) |  | 

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


## AssignOrganizationRolesToUsers

> AssignOrganizationRolesToUsers(ctx, id).AssignOrganizationRolesToUsersRequest(assignOrganizationRolesToUsersRequest).Execute()

Assign roles to organization user members



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
	id := "id_example" // string | The unique identifier of the organization.
	assignOrganizationRolesToUsersRequest := *openapiclient.NewAssignOrganizationRolesToUsersRequest([]string{"UserIds_example"}, []string{"OrganizationRoleIds_example"}) // AssignOrganizationRolesToUsersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AssignOrganizationRolesToUsers(context.Background(), id).AssignOrganizationRolesToUsersRequest(assignOrganizationRolesToUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AssignOrganizationRolesToUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOrganizationRolesToUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignOrganizationRolesToUsersRequest** | [**AssignOrganizationRolesToUsersRequest**](AssignOrganizationRolesToUsersRequest.md) |  | 

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


## CreateOrganization

> ListApplicationUserConsentOrganizations200ResponseOrganizationsInner CreateOrganization(ctx).CreateOrganizationRequest(createOrganizationRequest).Execute()

Create an organization



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
	createOrganizationRequest := *openapiclient.NewCreateOrganizationRequest("Name_example") // CreateOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateOrganization(context.Background()).CreateOrganizationRequest(createOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganization`: ListApplicationUserConsentOrganizations200ResponseOrganizationsInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganizationRequest** | [**CreateOrganizationRequest**](CreateOrganizationRequest.md) |  | 

### Return type

[**ListApplicationUserConsentOrganizations200ResponseOrganizationsInner**](ListApplicationUserConsentOrganizations200ResponseOrganizationsInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationJitEmailDomain

> ListOrganizationJitEmailDomains200ResponseInner CreateOrganizationJitEmailDomain(ctx, id).CreateOrganizationJitEmailDomainRequest(createOrganizationJitEmailDomainRequest).Execute()

Add organization JIT email domain



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
	id := "id_example" // string | The unique identifier of the organization.
	createOrganizationJitEmailDomainRequest := *openapiclient.NewCreateOrganizationJitEmailDomainRequest("EmailDomain_example") // CreateOrganizationJitEmailDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateOrganizationJitEmailDomain(context.Background(), id).CreateOrganizationJitEmailDomainRequest(createOrganizationJitEmailDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateOrganizationJitEmailDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationJitEmailDomain`: ListOrganizationJitEmailDomains200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateOrganizationJitEmailDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationJitEmailDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationJitEmailDomainRequest** | [**CreateOrganizationJitEmailDomainRequest**](CreateOrganizationJitEmailDomainRequest.md) |  | 

### Return type

[**ListOrganizationJitEmailDomains200ResponseInner**](ListOrganizationJitEmailDomains200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationJitRole

> CreateOrganizationJitRole(ctx, id).CreateOrganizationJitRoleRequest(createOrganizationJitRoleRequest).Execute()

Add organization JIT default roles



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
	id := "id_example" // string | The unique identifier of the organization.
	createOrganizationJitRoleRequest := *openapiclient.NewCreateOrganizationJitRoleRequest([]string{"OrganizationRoleIds_example"}) // CreateOrganizationJitRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.CreateOrganizationJitRole(context.Background(), id).CreateOrganizationJitRoleRequest(createOrganizationJitRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateOrganizationJitRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationJitRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationJitRoleRequest** | [**CreateOrganizationJitRoleRequest**](CreateOrganizationJitRoleRequest.md) |  | 

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


## CreateOrganizationJitSsoConnector

> CreateOrganizationJitSsoConnector(ctx, id).CreateOrganizationJitSsoConnectorRequest(createOrganizationJitSsoConnectorRequest).Execute()

Add organization JIT SSO connectors



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
	id := "id_example" // string | The unique identifier of the organization.
	createOrganizationJitSsoConnectorRequest := *openapiclient.NewCreateOrganizationJitSsoConnectorRequest([]string{"SsoConnectorIds_example"}) // CreateOrganizationJitSsoConnectorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.CreateOrganizationJitSsoConnector(context.Background(), id).CreateOrganizationJitSsoConnectorRequest(createOrganizationJitSsoConnectorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateOrganizationJitSsoConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationJitSsoConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationJitSsoConnectorRequest** | [**CreateOrganizationJitSsoConnectorRequest**](CreateOrganizationJitSsoConnectorRequest.md) |  | 

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


## DeleteOrganization

> DeleteOrganization(ctx, id).Execute()

Delete organization



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
	id := "id_example" // string | The unique identifier of the organization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganization(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


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


## DeleteOrganizationApplication

> DeleteOrganizationApplication(ctx, id, applicationId).Execute()

Remove organization application



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
	id := "id_example" // string | The unique identifier of the organization.
	applicationId := "applicationId_example" // string | The unique identifier of the application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganizationApplication(context.Background(), id, applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganizationApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationApplicationRequest struct via the builder pattern


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


## DeleteOrganizationApplicationRole

> DeleteOrganizationApplicationRole(ctx, id, applicationId, organizationRoleId).Execute()

Remove organization application role



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
	id := "id_example" // string | The unique identifier of the organization.
	applicationId := "applicationId_example" // string | The unique identifier of the application.
	organizationRoleId := "organizationRoleId_example" // string | The unique identifier of the organization role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganizationApplicationRole(context.Background(), id, applicationId, organizationRoleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganizationApplicationRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**applicationId** | **string** | The unique identifier of the application. | 
**organizationRoleId** | **string** | The unique identifier of the organization role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationApplicationRoleRequest struct via the builder pattern


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


## DeleteOrganizationJitEmailDomain

> DeleteOrganizationJitEmailDomain(ctx, id, emailDomain).Execute()

Remove organization JIT email domain



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
	id := "id_example" // string | The unique identifier of the organization.
	emailDomain := "emailDomain_example" // string | The email domain to remove.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganizationJitEmailDomain(context.Background(), id, emailDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganizationJitEmailDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**emailDomain** | **string** | The email domain to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationJitEmailDomainRequest struct via the builder pattern


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


## DeleteOrganizationJitRole

> DeleteOrganizationJitRole(ctx, id, organizationRoleId).Execute()

Remove organization JIT default role



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
	id := "id_example" // string | The unique identifier of the organization.
	organizationRoleId := "organizationRoleId_example" // string | The unique identifier of the organization role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganizationJitRole(context.Background(), id, organizationRoleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganizationJitRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**organizationRoleId** | **string** | The unique identifier of the organization role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationJitRoleRequest struct via the builder pattern


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


## DeleteOrganizationJitSsoConnector

> DeleteOrganizationJitSsoConnector(ctx, id, ssoConnectorId).Execute()

Remove organization JIT SSO connector



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
	id := "id_example" // string | The unique identifier of the organization.
	ssoConnectorId := "ssoConnectorId_example" // string | The unique identifier of the sso connector.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganizationJitSsoConnector(context.Background(), id, ssoConnectorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganizationJitSsoConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**ssoConnectorId** | **string** | The unique identifier of the sso connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationJitSsoConnectorRequest struct via the builder pattern


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


## DeleteOrganizationUser

> DeleteOrganizationUser(ctx, id, userId).Execute()

Remove user member from organization



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
	id := "id_example" // string | The unique identifier of the organization.
	userId := "userId_example" // string | The unique identifier of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganizationUser(context.Background(), id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganizationUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationUserRequest struct via the builder pattern


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


## DeleteOrganizationUserRole

> DeleteOrganizationUserRole(ctx, id, userId, organizationRoleId).Execute()

Remove a role from a user in an organization



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
	id := "id_example" // string | The unique identifier of the organization.
	userId := "userId_example" // string | The unique identifier of the user.
	organizationRoleId := "organizationRoleId_example" // string | The unique identifier of the organization role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganizationUserRole(context.Background(), id, userId, organizationRoleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganizationUserRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**userId** | **string** | The unique identifier of the user. | 
**organizationRoleId** | **string** | The unique identifier of the organization role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationUserRoleRequest struct via the builder pattern


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


## GetOrganization

> ListApplicationUserConsentOrganizations200ResponseOrganizationsInner GetOrganization(ctx, id).Execute()

Get organization



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
	id := "id_example" // string | The unique identifier of the organization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganization(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganization`: ListApplicationUserConsentOrganizations200ResponseOrganizationsInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListApplicationUserConsentOrganizations200ResponseOrganizationsInner**](ListApplicationUserConsentOrganizations200ResponseOrganizationsInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationApplicationRoles

> []GetOrganizationRole200Response ListOrganizationApplicationRoles(ctx, id, applicationId).Page(page).PageSize(pageSize).Execute()

Get organization application roles



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
	id := "id_example" // string | The unique identifier of the organization.
	applicationId := "applicationId_example" // string | The unique identifier of the application.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListOrganizationApplicationRoles(context.Background(), id, applicationId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListOrganizationApplicationRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationApplicationRoles`: []GetOrganizationRole200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListOrganizationApplicationRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationApplicationRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]GetOrganizationRole200Response**](GetOrganizationRole200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationApplications

> []ListOrganizationApplications200ResponseInner ListOrganizationApplications(ctx, id).Q(q).Page(page).PageSize(pageSize).Execute()

Get organization applications



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
	id := "id_example" // string | The unique identifier of the organization.
	q := "q_example" // string |  (optional)
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListOrganizationApplications(context.Background(), id).Q(q).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListOrganizationApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationApplications`: []ListOrganizationApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListOrganizationApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** |  | 
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListOrganizationApplications200ResponseInner**](ListOrganizationApplications200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationJitEmailDomains

> []ListOrganizationJitEmailDomains200ResponseInner ListOrganizationJitEmailDomains(ctx, id).Page(page).PageSize(pageSize).Execute()

Get organization JIT email domains



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
	id := "id_example" // string | The unique identifier of the organization.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListOrganizationJitEmailDomains(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListOrganizationJitEmailDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationJitEmailDomains`: []ListOrganizationJitEmailDomains200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListOrganizationJitEmailDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationJitEmailDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListOrganizationJitEmailDomains200ResponseInner**](ListOrganizationJitEmailDomains200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationJitRoles

> []GetOrganizationRole200Response ListOrganizationJitRoles(ctx, id).Page(page).PageSize(pageSize).Execute()

Get organization JIT default roles



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
	id := "id_example" // string | The unique identifier of the organization.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListOrganizationJitRoles(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListOrganizationJitRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationJitRoles`: []GetOrganizationRole200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListOrganizationJitRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationJitRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]GetOrganizationRole200Response**](GetOrganizationRole200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationJitSsoConnectors

> []ListOrganizationJitSsoConnectors200ResponseInner ListOrganizationJitSsoConnectors(ctx, id).Page(page).PageSize(pageSize).Execute()

Get organization JIT SSO connectors



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
	id := "id_example" // string | The unique identifier of the organization.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListOrganizationJitSsoConnectors(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListOrganizationJitSsoConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationJitSsoConnectors`: []ListOrganizationJitSsoConnectors200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListOrganizationJitSsoConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationJitSsoConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListOrganizationJitSsoConnectors200ResponseInner**](ListOrganizationJitSsoConnectors200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationUserRoles

> []GetOrganizationRole200Response ListOrganizationUserRoles(ctx, id, userId).Page(page).PageSize(pageSize).Execute()

Get roles for a user in an organization



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
	id := "id_example" // string | The unique identifier of the organization.
	userId := "userId_example" // string | The unique identifier of the user.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListOrganizationUserRoles(context.Background(), id, userId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListOrganizationUserRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationUserRoles`: []GetOrganizationRole200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListOrganizationUserRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]GetOrganizationRole200Response**](GetOrganizationRole200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationUserScopes

> []ListOrganizationRoleScopes200ResponseInner ListOrganizationUserScopes(ctx, id, userId).Execute()

Get scopes for a user in an organization tailored by the organization roles



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
	id := "id_example" // string | The unique identifier of the organization.
	userId := "userId_example" // string | The unique identifier of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListOrganizationUserScopes(context.Background(), id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListOrganizationUserScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationUserScopes`: []ListOrganizationRoleScopes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListOrganizationUserScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationUserScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ListOrganizationUsers

> []ListOrganizationUsers200ResponseInner ListOrganizationUsers(ctx, id).Q(q).Page(page).PageSize(pageSize).Execute()

Get organization user members



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
	id := "id_example" // string | The unique identifier of the organization.
	q := "q_example" // string | The query to filter users. It will match multiple fields of users, including ID, name, username, email, and phone number.  If not provided, all users will be returned. (optional)
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListOrganizationUsers(context.Background(), id).Q(q).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListOrganizationUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationUsers`: []ListOrganizationUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListOrganizationUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | The query to filter users. It will match multiple fields of users, including ID, name, username, email, and phone number.  If not provided, all users will be returned. | 
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListOrganizationUsers200ResponseInner**](ListOrganizationUsers200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizations

> []ListOrganizations200ResponseInner ListOrganizations(ctx).Q(q).ShowFeatured(showFeatured).Page(page).PageSize(pageSize).Execute()

Get organizations



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
	q := "q_example" // string | The query to filter organizations. It can be a partial ID or name.  If not provided, all organizations will be returned. (optional)
	showFeatured := "showFeatured_example" // string | Whether to show featured users in the organization. Featured users are randomly selected from the organization members.  If not provided, `featuredUsers` will not be included in the response. (optional)
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListOrganizations(context.Background()).Q(q).ShowFeatured(showFeatured).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizations`: []ListOrganizations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The query to filter organizations. It can be a partial ID or name.  If not provided, all organizations will be returned. | 
 **showFeatured** | **string** | Whether to show featured users in the organization. Featured users are randomly selected from the organization members.  If not provided, &#x60;featuredUsers&#x60; will not be included in the response. | 
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListOrganizations200ResponseInner**](ListOrganizations200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOrganizationApplicationRoles

> ReplaceOrganizationApplicationRoles(ctx, id, applicationId).ReplaceOrganizationApplicationRolesRequest(replaceOrganizationApplicationRolesRequest).Execute()

Replace organization application roles



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
	id := "id_example" // string | The unique identifier of the organization.
	applicationId := "applicationId_example" // string | The unique identifier of the application.
	replaceOrganizationApplicationRolesRequest := *openapiclient.NewReplaceOrganizationApplicationRolesRequest([]string{"OrganizationRoleIds_example"}) // ReplaceOrganizationApplicationRolesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.ReplaceOrganizationApplicationRoles(context.Background(), id, applicationId).ReplaceOrganizationApplicationRolesRequest(replaceOrganizationApplicationRolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ReplaceOrganizationApplicationRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrganizationApplicationRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceOrganizationApplicationRolesRequest** | [**ReplaceOrganizationApplicationRolesRequest**](ReplaceOrganizationApplicationRolesRequest.md) |  | 

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


## ReplaceOrganizationApplications

> ReplaceOrganizationApplications(ctx, id).ReplaceOrganizationApplicationsRequest(replaceOrganizationApplicationsRequest).Execute()

Replace organization applications



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
	id := "id_example" // string | The unique identifier of the organization.
	replaceOrganizationApplicationsRequest := *openapiclient.NewReplaceOrganizationApplicationsRequest([]string{"ApplicationIds_example"}) // ReplaceOrganizationApplicationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.ReplaceOrganizationApplications(context.Background(), id).ReplaceOrganizationApplicationsRequest(replaceOrganizationApplicationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ReplaceOrganizationApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrganizationApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceOrganizationApplicationsRequest** | [**ReplaceOrganizationApplicationsRequest**](ReplaceOrganizationApplicationsRequest.md) |  | 

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


## ReplaceOrganizationJitEmailDomains

> ReplaceOrganizationJitEmailDomains(ctx, id).ReplaceOrganizationJitEmailDomainsRequest(replaceOrganizationJitEmailDomainsRequest).Execute()

Replace organization JIT email domains



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
	id := "id_example" // string | The unique identifier of the organization.
	replaceOrganizationJitEmailDomainsRequest := *openapiclient.NewReplaceOrganizationJitEmailDomainsRequest([]string{"EmailDomains_example"}) // ReplaceOrganizationJitEmailDomainsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.ReplaceOrganizationJitEmailDomains(context.Background(), id).ReplaceOrganizationJitEmailDomainsRequest(replaceOrganizationJitEmailDomainsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ReplaceOrganizationJitEmailDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrganizationJitEmailDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceOrganizationJitEmailDomainsRequest** | [**ReplaceOrganizationJitEmailDomainsRequest**](ReplaceOrganizationJitEmailDomainsRequest.md) |  | 

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


## ReplaceOrganizationJitRoles

> ReplaceOrganizationJitRoles(ctx, id).ReplaceOrganizationJitRolesRequest(replaceOrganizationJitRolesRequest).Execute()

Replace organization JIT default roles



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
	id := "id_example" // string | The unique identifier of the organization.
	replaceOrganizationJitRolesRequest := *openapiclient.NewReplaceOrganizationJitRolesRequest([]string{"OrganizationRoleIds_example"}) // ReplaceOrganizationJitRolesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.ReplaceOrganizationJitRoles(context.Background(), id).ReplaceOrganizationJitRolesRequest(replaceOrganizationJitRolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ReplaceOrganizationJitRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrganizationJitRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceOrganizationJitRolesRequest** | [**ReplaceOrganizationJitRolesRequest**](ReplaceOrganizationJitRolesRequest.md) |  | 

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


## ReplaceOrganizationJitSsoConnectors

> ReplaceOrganizationJitSsoConnectors(ctx, id).ReplaceOrganizationJitSsoConnectorsRequest(replaceOrganizationJitSsoConnectorsRequest).Execute()

Replace organization JIT SSO connectors



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
	id := "id_example" // string | The unique identifier of the organization.
	replaceOrganizationJitSsoConnectorsRequest := *openapiclient.NewReplaceOrganizationJitSsoConnectorsRequest([]string{"SsoConnectorIds_example"}) // ReplaceOrganizationJitSsoConnectorsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.ReplaceOrganizationJitSsoConnectors(context.Background(), id).ReplaceOrganizationJitSsoConnectorsRequest(replaceOrganizationJitSsoConnectorsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ReplaceOrganizationJitSsoConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrganizationJitSsoConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceOrganizationJitSsoConnectorsRequest** | [**ReplaceOrganizationJitSsoConnectorsRequest**](ReplaceOrganizationJitSsoConnectorsRequest.md) |  | 

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


## ReplaceOrganizationUserRoles

> ReplaceOrganizationUserRoles(ctx, id, userId).ReplaceOrganizationUserRolesRequest(replaceOrganizationUserRolesRequest).Execute()

Update roles for a user in an organization



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
	id := "id_example" // string | The unique identifier of the organization.
	userId := "userId_example" // string | The unique identifier of the user.
	replaceOrganizationUserRolesRequest := *openapiclient.NewReplaceOrganizationUserRolesRequest([]string{"OrganizationRoleIds_example"}) // ReplaceOrganizationUserRolesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.ReplaceOrganizationUserRoles(context.Background(), id, userId).ReplaceOrganizationUserRolesRequest(replaceOrganizationUserRolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ReplaceOrganizationUserRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrganizationUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceOrganizationUserRolesRequest** | [**ReplaceOrganizationUserRolesRequest**](ReplaceOrganizationUserRolesRequest.md) |  | 

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


## ReplaceOrganizationUsers

> ReplaceOrganizationUsers(ctx, id).ReplaceOrganizationUsersRequest(replaceOrganizationUsersRequest).Execute()

Replace organization user members



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
	id := "id_example" // string | The unique identifier of the organization.
	replaceOrganizationUsersRequest := *openapiclient.NewReplaceOrganizationUsersRequest([]string{"UserIds_example"}) // ReplaceOrganizationUsersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.ReplaceOrganizationUsers(context.Background(), id).ReplaceOrganizationUsersRequest(replaceOrganizationUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ReplaceOrganizationUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrganizationUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceOrganizationUsersRequest** | [**ReplaceOrganizationUsersRequest**](ReplaceOrganizationUsersRequest.md) |  | 

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


## UpdateOrganization

> ListApplicationUserConsentOrganizations200ResponseOrganizationsInner UpdateOrganization(ctx, id).UpdateOrganizationRequest(updateOrganizationRequest).Execute()

Update organization



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
	id := "id_example" // string | The unique identifier of the organization.
	updateOrganizationRequest := *openapiclient.NewUpdateOrganizationRequest() // UpdateOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.UpdateOrganization(context.Background(), id).UpdateOrganizationRequest(updateOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.UpdateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganization`: ListApplicationUserConsentOrganizations200ResponseOrganizationsInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationRequest** | [**UpdateOrganizationRequest**](UpdateOrganizationRequest.md) |  | 

### Return type

[**ListApplicationUserConsentOrganizations200ResponseOrganizationsInner**](ListApplicationUserConsentOrganizations200ResponseOrganizationsInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

