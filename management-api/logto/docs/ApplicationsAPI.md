# \ApplicationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignApplicationRoles**](ApplicationsAPI.md#AssignApplicationRoles) | **Post** /api/applications/{applicationId}/roles | Assign API resource roles to application
[**CreateApplication**](ApplicationsAPI.md#CreateApplication) | **Post** /api/applications | Create an application
[**CreateApplicationProtectedAppMetadataCustomDomain**](ApplicationsAPI.md#CreateApplicationProtectedAppMetadataCustomDomain) | **Post** /api/applications/{id}/protected-app-metadata/custom-domains | Add a custom domain to the application.
[**CreateApplicationSecret**](ApplicationsAPI.md#CreateApplicationSecret) | **Post** /api/applications/{id}/secrets | Add application secret
[**CreateApplicationUserConsentOrganization**](ApplicationsAPI.md#CreateApplicationUserConsentOrganization) | **Post** /api/applications/{id}/users/{userId}/consent-organizations | Grant a list of organization access of a user for a application.
[**CreateApplicationUserConsentScope**](ApplicationsAPI.md#CreateApplicationUserConsentScope) | **Post** /api/applications/{applicationId}/user-consent-scopes | Assign user consent scopes to application.
[**DeleteApplication**](ApplicationsAPI.md#DeleteApplication) | **Delete** /api/applications/{id} | Delete application
[**DeleteApplicationLegacySecret**](ApplicationsAPI.md#DeleteApplicationLegacySecret) | **Delete** /api/applications/{id}/legacy-secret | Delete application legacy secret
[**DeleteApplicationProtectedAppMetadataCustomDomain**](ApplicationsAPI.md#DeleteApplicationProtectedAppMetadataCustomDomain) | **Delete** /api/applications/{id}/protected-app-metadata/custom-domains/{domain} | Remove custom domain.
[**DeleteApplicationRole**](ApplicationsAPI.md#DeleteApplicationRole) | **Delete** /api/applications/{applicationId}/roles/{roleId} | Remove a API resource role from application
[**DeleteApplicationSecret**](ApplicationsAPI.md#DeleteApplicationSecret) | **Delete** /api/applications/{id}/secrets/{name} | Delete application secret
[**DeleteApplicationUserConsentOrganization**](ApplicationsAPI.md#DeleteApplicationUserConsentOrganization) | **Delete** /api/applications/{id}/users/{userId}/consent-organizations/{organizationId} | Revoke a user&#39;s access to an organization for a application.
[**DeleteApplicationUserConsentScope**](ApplicationsAPI.md#DeleteApplicationUserConsentScope) | **Delete** /api/applications/{applicationId}/user-consent-scopes/{scopeType}/{scopeId} | Remove user consent scope from application.
[**GetApplication**](ApplicationsAPI.md#GetApplication) | **Get** /api/applications/{id} | Get application
[**GetApplicationSignInExperience**](ApplicationsAPI.md#GetApplicationSignInExperience) | **Get** /api/applications/{applicationId}/sign-in-experience | Get the application level sign-in experience
[**ListApplicationOrganizations**](ApplicationsAPI.md#ListApplicationOrganizations) | **Get** /api/applications/{id}/organizations | Get application organizations
[**ListApplicationProtectedAppMetadataCustomDomains**](ApplicationsAPI.md#ListApplicationProtectedAppMetadataCustomDomains) | **Get** /api/applications/{id}/protected-app-metadata/custom-domains | Get application custom domains.
[**ListApplicationRoles**](ApplicationsAPI.md#ListApplicationRoles) | **Get** /api/applications/{applicationId}/roles | Get application API resource roles
[**ListApplicationSecrets**](ApplicationsAPI.md#ListApplicationSecrets) | **Get** /api/applications/{id}/secrets | Get application secrets
[**ListApplicationUserConsentOrganizations**](ApplicationsAPI.md#ListApplicationUserConsentOrganizations) | **Get** /api/applications/{id}/users/{userId}/consent-organizations | List all the user consented organizations of a application.
[**ListApplicationUserConsentScopes**](ApplicationsAPI.md#ListApplicationUserConsentScopes) | **Get** /api/applications/{applicationId}/user-consent-scopes | List all the user consent scopes of an application.
[**ListApplications**](ApplicationsAPI.md#ListApplications) | **Get** /api/applications | Get applications
[**ReplaceApplicationRoles**](ApplicationsAPI.md#ReplaceApplicationRoles) | **Put** /api/applications/{applicationId}/roles | Update API resource roles for application
[**ReplaceApplicationSignInExperience**](ApplicationsAPI.md#ReplaceApplicationSignInExperience) | **Put** /api/applications/{applicationId}/sign-in-experience | Update application level sign-in experience
[**ReplaceApplicationUserConsentOrganizations**](ApplicationsAPI.md#ReplaceApplicationUserConsentOrganizations) | **Put** /api/applications/{id}/users/{userId}/consent-organizations | Grant a list of organization access of a user for a application.
[**UpdateApplication**](ApplicationsAPI.md#UpdateApplication) | **Patch** /api/applications/{id} | Update application
[**UpdateApplicationCustomData**](ApplicationsAPI.md#UpdateApplicationCustomData) | **Patch** /api/applications/{applicationId}/custom-data | Update application custom data
[**UpdateApplicationSecret**](ApplicationsAPI.md#UpdateApplicationSecret) | **Patch** /api/applications/{id}/secrets/{name} | Update application secret



## AssignApplicationRoles

> AssignApplicationRoles(ctx, applicationId).AssignApplicationRolesRequest(assignApplicationRolesRequest).Execute()

Assign API resource roles to application



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
	applicationId := "applicationId_example" // string | The unique identifier of the application.
	assignApplicationRolesRequest := *openapiclient.NewAssignApplicationRolesRequest([]string{"RoleIds_example"}) // AssignApplicationRolesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.AssignApplicationRoles(context.Background(), applicationId).AssignApplicationRolesRequest(assignApplicationRolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.AssignApplicationRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignApplicationRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignApplicationRolesRequest** | [**AssignApplicationRolesRequest**](AssignApplicationRolesRequest.md) |  | 

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


## CreateApplication

> ListApplications200ResponseInner CreateApplication(ctx).CreateApplicationRequest(createApplicationRequest).Execute()

Create an application



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
	createApplicationRequest := *openapiclient.NewCreateApplicationRequest("Name_example", "Type_example") // CreateApplicationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.CreateApplication(context.Background()).CreateApplicationRequest(createApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CreateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplication`: ListApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.CreateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApplicationRequest** | [**CreateApplicationRequest**](CreateApplicationRequest.md) |  | 

### Return type

[**ListApplications200ResponseInner**](ListApplications200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationProtectedAppMetadataCustomDomain

> CreateApplicationProtectedAppMetadataCustomDomain(ctx, id).CreateApplicationProtectedAppMetadataCustomDomainRequest(createApplicationProtectedAppMetadataCustomDomainRequest).Execute()

Add a custom domain to the application.



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
	id := "id_example" // string | The unique identifier of the application.
	createApplicationProtectedAppMetadataCustomDomainRequest := *openapiclient.NewCreateApplicationProtectedAppMetadataCustomDomainRequest("Domain_example") // CreateApplicationProtectedAppMetadataCustomDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.CreateApplicationProtectedAppMetadataCustomDomain(context.Background(), id).CreateApplicationProtectedAppMetadataCustomDomainRequest(createApplicationProtectedAppMetadataCustomDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CreateApplicationProtectedAppMetadataCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationProtectedAppMetadataCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createApplicationProtectedAppMetadataCustomDomainRequest** | [**CreateApplicationProtectedAppMetadataCustomDomainRequest**](CreateApplicationProtectedAppMetadataCustomDomainRequest.md) |  | 

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


## CreateApplicationSecret

> ListApplicationSecrets200ResponseInner CreateApplicationSecret(ctx, id).CreateApplicationSecretRequest(createApplicationSecretRequest).Execute()

Add application secret



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
	id := "id_example" // string | The unique identifier of the application.
	createApplicationSecretRequest := *openapiclient.NewCreateApplicationSecretRequest("Name_example") // CreateApplicationSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.CreateApplicationSecret(context.Background(), id).CreateApplicationSecretRequest(createApplicationSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CreateApplicationSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationSecret`: ListApplicationSecrets200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.CreateApplicationSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createApplicationSecretRequest** | [**CreateApplicationSecretRequest**](CreateApplicationSecretRequest.md) |  | 

### Return type

[**ListApplicationSecrets200ResponseInner**](ListApplicationSecrets200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationUserConsentOrganization

> CreateApplicationUserConsentOrganization(ctx, id, userId).CreateApplicationUserConsentOrganizationRequest(createApplicationUserConsentOrganizationRequest).Execute()

Grant a list of organization access of a user for a application.



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
	id := "id_example" // string | The unique identifier of the application.
	userId := "userId_example" // string | The unique identifier of the user.
	createApplicationUserConsentOrganizationRequest := *openapiclient.NewCreateApplicationUserConsentOrganizationRequest([]string{"OrganizationIds_example"}) // CreateApplicationUserConsentOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.CreateApplicationUserConsentOrganization(context.Background(), id, userId).CreateApplicationUserConsentOrganizationRequest(createApplicationUserConsentOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CreateApplicationUserConsentOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationUserConsentOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createApplicationUserConsentOrganizationRequest** | [**CreateApplicationUserConsentOrganizationRequest**](CreateApplicationUserConsentOrganizationRequest.md) |  | 

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


## CreateApplicationUserConsentScope

> CreateApplicationUserConsentScope(ctx, applicationId).CreateApplicationUserConsentScopeRequest(createApplicationUserConsentScopeRequest).Execute()

Assign user consent scopes to application.



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
	applicationId := "applicationId_example" // string | The unique identifier of the application.
	createApplicationUserConsentScopeRequest := *openapiclient.NewCreateApplicationUserConsentScopeRequest() // CreateApplicationUserConsentScopeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.CreateApplicationUserConsentScope(context.Background(), applicationId).CreateApplicationUserConsentScopeRequest(createApplicationUserConsentScopeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CreateApplicationUserConsentScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationUserConsentScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createApplicationUserConsentScopeRequest** | [**CreateApplicationUserConsentScopeRequest**](CreateApplicationUserConsentScopeRequest.md) |  | 

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


## DeleteApplication

> DeleteApplication(ctx, id).Execute()

Delete application



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
	id := "id_example" // string | The unique identifier of the application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.DeleteApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## DeleteApplicationLegacySecret

> DeleteApplicationLegacySecret200Response DeleteApplicationLegacySecret(ctx, id).Execute()

Delete application legacy secret



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
	id := "id_example" // string | The unique identifier of the application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.DeleteApplicationLegacySecret(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplicationLegacySecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApplicationLegacySecret`: DeleteApplicationLegacySecret200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.DeleteApplicationLegacySecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationLegacySecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteApplicationLegacySecret200Response**](DeleteApplicationLegacySecret200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationProtectedAppMetadataCustomDomain

> DeleteApplicationProtectedAppMetadataCustomDomain(ctx, id, domain).Execute()

Remove custom domain.



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
	id := "id_example" // string | The unique identifier of the application.
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.DeleteApplicationProtectedAppMetadataCustomDomain(context.Background(), id, domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplicationProtectedAppMetadataCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationProtectedAppMetadataCustomDomainRequest struct via the builder pattern


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


## DeleteApplicationRole

> DeleteApplicationRole(ctx, applicationId, roleId).Execute()

Remove a API resource role from application



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
	applicationId := "applicationId_example" // string | The unique identifier of the application.
	roleId := "roleId_example" // string | The unique identifier of the role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.DeleteApplicationRole(context.Background(), applicationId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplicationRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The unique identifier of the application. | 
**roleId** | **string** | The unique identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRoleRequest struct via the builder pattern


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


## DeleteApplicationSecret

> DeleteApplicationSecret(ctx, id, name).Execute()

Delete application secret



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
	id := "id_example" // string | The unique identifier of the application.
	name := "name_example" // string | The name of the secret.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.DeleteApplicationSecret(context.Background(), id, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplicationSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 
**name** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationSecretRequest struct via the builder pattern


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


## DeleteApplicationUserConsentOrganization

> DeleteApplicationUserConsentOrganization(ctx, id, userId, organizationId).Execute()

Revoke a user's access to an organization for a application.



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
	id := "id_example" // string | The unique identifier of the application.
	userId := "userId_example" // string | The unique identifier of the user.
	organizationId := "organizationId_example" // string | The unique identifier of the organization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.DeleteApplicationUserConsentOrganization(context.Background(), id, userId, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplicationUserConsentOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 
**userId** | **string** | The unique identifier of the user. | 
**organizationId** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationUserConsentOrganizationRequest struct via the builder pattern


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


## DeleteApplicationUserConsentScope

> DeleteApplicationUserConsentScope(ctx, applicationId, scopeType, scopeId).Execute()

Remove user consent scope from application.



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
	applicationId := "applicationId_example" // string | The unique identifier of the application.
	scopeType := "scopeType_example" // string | 
	scopeId := "scopeId_example" // string | The unique identifier of the scope.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.DeleteApplicationUserConsentScope(context.Background(), applicationId, scopeType, scopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplicationUserConsentScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The unique identifier of the application. | 
**scopeType** | **string** |  | 
**scopeId** | **string** | The unique identifier of the scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationUserConsentScopeRequest struct via the builder pattern


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


## GetApplication

> GetApplication200Response GetApplication(ctx, id).Execute()

Get application



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
	id := "id_example" // string | The unique identifier of the application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.GetApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplication`: GetApplication200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetApplication200Response**](GetApplication200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationSignInExperience

> GetApplicationSignInExperience200Response GetApplicationSignInExperience(ctx, applicationId).Execute()

Get the application level sign-in experience



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
	applicationId := "applicationId_example" // string | The unique identifier of the application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.GetApplicationSignInExperience(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplicationSignInExperience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationSignInExperience`: GetApplicationSignInExperience200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplicationSignInExperience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationSignInExperienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetApplicationSignInExperience200Response**](GetApplicationSignInExperience200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationOrganizations

> []ListApplicationOrganizations200ResponseInner ListApplicationOrganizations(ctx, id).Page(page).PageSize(pageSize).Execute()

Get application organizations



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
	id := "id_example" // string | The unique identifier of the application.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.ListApplicationOrganizations(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ListApplicationOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationOrganizations`: []ListApplicationOrganizations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.ListApplicationOrganizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListApplicationOrganizations200ResponseInner**](ListApplicationOrganizations200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationProtectedAppMetadataCustomDomains

> []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner ListApplicationProtectedAppMetadataCustomDomains(ctx, id).Execute()

Get application custom domains.



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
	id := "id_example" // string | The unique identifier of the application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.ListApplicationProtectedAppMetadataCustomDomains(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ListApplicationProtectedAppMetadataCustomDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationProtectedAppMetadataCustomDomains`: []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.ListApplicationProtectedAppMetadataCustomDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationProtectedAppMetadataCustomDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner**](ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationRoles

> []ListApplicationRoles200ResponseInner ListApplicationRoles(ctx, applicationId).Page(page).PageSize(pageSize).SearchParams(searchParams).Execute()

Get application API resource roles



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
	applicationId := "applicationId_example" // string | The unique identifier of the application.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)
	searchParams := map[string]string{"key": "Inner_example"} // map[string]string | Search query parameters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.ListApplicationRoles(context.Background(), applicationId).Page(page).PageSize(pageSize).SearchParams(searchParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ListApplicationRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationRoles`: []ListApplicationRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.ListApplicationRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]
 **searchParams** | **map[string]string** | Search query parameters. | 

### Return type

[**[]ListApplicationRoles200ResponseInner**](ListApplicationRoles200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationSecrets

> []ListApplicationSecrets200ResponseInner ListApplicationSecrets(ctx, id).Execute()

Get application secrets



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
	id := "id_example" // string | The unique identifier of the application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.ListApplicationSecrets(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ListApplicationSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationSecrets`: []ListApplicationSecrets200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.ListApplicationSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListApplicationSecrets200ResponseInner**](ListApplicationSecrets200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationUserConsentOrganizations

> ListApplicationUserConsentOrganizations200Response ListApplicationUserConsentOrganizations(ctx, id, userId).Page(page).PageSize(pageSize).Execute()

List all the user consented organizations of a application.



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
	id := "id_example" // string | The unique identifier of the application.
	userId := "userId_example" // string | The unique identifier of the user.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.ListApplicationUserConsentOrganizations(context.Background(), id, userId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ListApplicationUserConsentOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationUserConsentOrganizations`: ListApplicationUserConsentOrganizations200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.ListApplicationUserConsentOrganizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationUserConsentOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**ListApplicationUserConsentOrganizations200Response**](ListApplicationUserConsentOrganizations200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationUserConsentScopes

> ListApplicationUserConsentScopes200Response ListApplicationUserConsentScopes(ctx, applicationId).Execute()

List all the user consent scopes of an application.



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
	applicationId := "applicationId_example" // string | The unique identifier of the application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.ListApplicationUserConsentScopes(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ListApplicationUserConsentScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationUserConsentScopes`: ListApplicationUserConsentScopes200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.ListApplicationUserConsentScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationUserConsentScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListApplicationUserConsentScopes200Response**](ListApplicationUserConsentScopes200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplications

> []ListApplications200ResponseInner ListApplications(ctx).Types(types).ExcludeRoleId(excludeRoleId).ExcludeOrganizationId(excludeOrganizationId).IsThirdParty(isThirdParty).Page(page).PageSize(pageSize).SearchParams(searchParams).Execute()

Get applications



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
	types := openapiclient.ListApplications_types_parameter{ArrayOfString: new([]string)} // ListApplicationsTypesParameter | An array of application types to filter applications. (optional)
	excludeRoleId := "excludeRoleId_example" // string |  (optional)
	excludeOrganizationId := "excludeOrganizationId_example" // string |  (optional)
	isThirdParty := openapiclient.ListApplications_isThirdParty_parameter{String: new(string)} // ListApplicationsIsThirdPartyParameter |  (optional)
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)
	searchParams := map[string]string{"key": "Inner_example"} // map[string]string | Search query parameters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.ListApplications(context.Background()).Types(types).ExcludeRoleId(excludeRoleId).ExcludeOrganizationId(excludeOrganizationId).IsThirdParty(isThirdParty).Page(page).PageSize(pageSize).SearchParams(searchParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ListApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplications`: []ListApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.ListApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **types** | [**ListApplicationsTypesParameter**](ListApplicationsTypesParameter.md) | An array of application types to filter applications. | 
 **excludeRoleId** | **string** |  | 
 **excludeOrganizationId** | **string** |  | 
 **isThirdParty** | [**ListApplicationsIsThirdPartyParameter**](ListApplicationsIsThirdPartyParameter.md) |  | 
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]
 **searchParams** | **map[string]string** | Search query parameters. | 

### Return type

[**[]ListApplications200ResponseInner**](ListApplications200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceApplicationRoles

> ReplaceApplicationRoles(ctx, applicationId).ReplaceApplicationRolesRequest(replaceApplicationRolesRequest).Execute()

Update API resource roles for application



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
	applicationId := "applicationId_example" // string | The unique identifier of the application.
	replaceApplicationRolesRequest := *openapiclient.NewReplaceApplicationRolesRequest([]string{"RoleIds_example"}) // ReplaceApplicationRolesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.ReplaceApplicationRoles(context.Background(), applicationId).ReplaceApplicationRolesRequest(replaceApplicationRolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ReplaceApplicationRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceApplicationRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceApplicationRolesRequest** | [**ReplaceApplicationRolesRequest**](ReplaceApplicationRolesRequest.md) |  | 

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


## ReplaceApplicationSignInExperience

> GetApplicationSignInExperience200Response ReplaceApplicationSignInExperience(ctx, applicationId).ReplaceApplicationSignInExperienceRequest(replaceApplicationSignInExperienceRequest).Execute()

Update application level sign-in experience



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
	applicationId := "applicationId_example" // string | The unique identifier of the application.
	replaceApplicationSignInExperienceRequest := *openapiclient.NewReplaceApplicationSignInExperienceRequest(openapiclient.ReplaceApplicationSignInExperience_request_termsOfUseUrl{String: new(string)}, openapiclient.ReplaceApplicationSignInExperience_request_termsOfUseUrl{String: new(string)}) // ReplaceApplicationSignInExperienceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.ReplaceApplicationSignInExperience(context.Background(), applicationId).ReplaceApplicationSignInExperienceRequest(replaceApplicationSignInExperienceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ReplaceApplicationSignInExperience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceApplicationSignInExperience`: GetApplicationSignInExperience200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.ReplaceApplicationSignInExperience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceApplicationSignInExperienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceApplicationSignInExperienceRequest** | [**ReplaceApplicationSignInExperienceRequest**](ReplaceApplicationSignInExperienceRequest.md) |  | 

### Return type

[**GetApplicationSignInExperience200Response**](GetApplicationSignInExperience200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceApplicationUserConsentOrganizations

> ReplaceApplicationUserConsentOrganizations(ctx, id, userId).ReplaceApplicationUserConsentOrganizationsRequest(replaceApplicationUserConsentOrganizationsRequest).Execute()

Grant a list of organization access of a user for a application.



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
	id := "id_example" // string | The unique identifier of the application.
	userId := "userId_example" // string | The unique identifier of the user.
	replaceApplicationUserConsentOrganizationsRequest := *openapiclient.NewReplaceApplicationUserConsentOrganizationsRequest([]string{"OrganizationIds_example"}) // ReplaceApplicationUserConsentOrganizationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.ReplaceApplicationUserConsentOrganizations(context.Background(), id, userId).ReplaceApplicationUserConsentOrganizationsRequest(replaceApplicationUserConsentOrganizationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.ReplaceApplicationUserConsentOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceApplicationUserConsentOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceApplicationUserConsentOrganizationsRequest** | [**ReplaceApplicationUserConsentOrganizationsRequest**](ReplaceApplicationUserConsentOrganizationsRequest.md) |  | 

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


## UpdateApplication

> ListApplications200ResponseInner UpdateApplication(ctx, id).UpdateApplicationRequest(updateApplicationRequest).Execute()

Update application



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
	id := "id_example" // string | The unique identifier of the application.
	updateApplicationRequest := *openapiclient.NewUpdateApplicationRequest() // UpdateApplicationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.UpdateApplication(context.Background(), id).UpdateApplicationRequest(updateApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.UpdateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplication`: ListApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateApplicationRequest** | [**UpdateApplicationRequest**](UpdateApplicationRequest.md) |  | 

### Return type

[**ListApplications200ResponseInner**](ListApplications200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationCustomData

> map[string]interface{} UpdateApplicationCustomData(ctx, applicationId).Body(body).Execute()

Update application custom data



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
	applicationId := "applicationId_example" // string | The unique identifier of the application.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.UpdateApplicationCustomData(context.Background(), applicationId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.UpdateApplicationCustomData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationCustomData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.UpdateApplicationCustomData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationCustomDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationSecret

> ListApplicationSecrets200ResponseInner UpdateApplicationSecret(ctx, id, name).UpdateApplicationSecretRequest(updateApplicationSecretRequest).Execute()

Update application secret



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
	id := "id_example" // string | The unique identifier of the application.
	name := "name_example" // string | The name of the secret.
	updateApplicationSecretRequest := *openapiclient.NewUpdateApplicationSecretRequest("Name_example") // UpdateApplicationSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.UpdateApplicationSecret(context.Background(), id, name).UpdateApplicationSecretRequest(updateApplicationSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.UpdateApplicationSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationSecret`: ListApplicationSecrets200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.UpdateApplicationSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the application. | 
**name** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateApplicationSecretRequest** | [**UpdateApplicationSecretRequest**](UpdateApplicationSecretRequest.md) |  | 

### Return type

[**ListApplicationSecrets200ResponseInner**](ListApplicationSecrets200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

