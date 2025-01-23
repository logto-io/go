# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignUserRoles**](UsersAPI.md#AssignUserRoles) | **Post** /api/users/{userId}/roles | Assign roles to user
[**CreateUser**](UsersAPI.md#CreateUser) | **Post** /api/users | Create user
[**CreateUserIdentity**](UsersAPI.md#CreateUserIdentity) | **Post** /api/users/{userId}/identities | Link social identity to user
[**CreateUserMfaVerification**](UsersAPI.md#CreateUserMfaVerification) | **Post** /api/users/{userId}/mfa-verifications | Create an MFA verification for a user
[**CreateUserPersonalAccessToken**](UsersAPI.md#CreateUserPersonalAccessToken) | **Post** /api/users/{userId}/personal-access-tokens | Add personal access token
[**DeleteUser**](UsersAPI.md#DeleteUser) | **Delete** /api/users/{userId} | Delete user
[**DeleteUserIdentity**](UsersAPI.md#DeleteUserIdentity) | **Delete** /api/users/{userId}/identities/{target} | Delete social identity from user
[**DeleteUserMfaVerification**](UsersAPI.md#DeleteUserMfaVerification) | **Delete** /api/users/{userId}/mfa-verifications/{verificationId} | Delete an MFA verification for a user
[**DeleteUserPersonalAccessToken**](UsersAPI.md#DeleteUserPersonalAccessToken) | **Delete** /api/users/{userId}/personal-access-tokens/{name} | Delete personal access token
[**DeleteUserRole**](UsersAPI.md#DeleteUserRole) | **Delete** /api/users/{userId}/roles/{roleId} | Remove role from user
[**GetUser**](UsersAPI.md#GetUser) | **Get** /api/users/{userId} | Get user
[**GetUserHasPassword**](UsersAPI.md#GetUserHasPassword) | **Get** /api/users/{userId}/has-password | Check if user has password
[**ListUserCustomData**](UsersAPI.md#ListUserCustomData) | **Get** /api/users/{userId}/custom-data | Get user custom data
[**ListUserMfaVerifications**](UsersAPI.md#ListUserMfaVerifications) | **Get** /api/users/{userId}/mfa-verifications | Get user&#39;s MFA verifications
[**ListUserOrganizations**](UsersAPI.md#ListUserOrganizations) | **Get** /api/users/{userId}/organizations | Get organizations for a user
[**ListUserPersonalAccessTokens**](UsersAPI.md#ListUserPersonalAccessTokens) | **Get** /api/users/{userId}/personal-access-tokens | Get personal access tokens
[**ListUserRoles**](UsersAPI.md#ListUserRoles) | **Get** /api/users/{userId}/roles | Get roles for user
[**ListUsers**](UsersAPI.md#ListUsers) | **Get** /api/users | Get users
[**ReplaceUserIdentity**](UsersAPI.md#ReplaceUserIdentity) | **Put** /api/users/{userId}/identities/{target} | Update social identity of user
[**ReplaceUserRoles**](UsersAPI.md#ReplaceUserRoles) | **Put** /api/users/{userId}/roles | Update roles for user
[**UpdateUser**](UsersAPI.md#UpdateUser) | **Patch** /api/users/{userId} | Update user
[**UpdateUserCustomData**](UsersAPI.md#UpdateUserCustomData) | **Patch** /api/users/{userId}/custom-data | Update user custom data
[**UpdateUserIsSuspended**](UsersAPI.md#UpdateUserIsSuspended) | **Patch** /api/users/{userId}/is-suspended | Update user suspension status
[**UpdateUserPassword**](UsersAPI.md#UpdateUserPassword) | **Patch** /api/users/{userId}/password | Update user password
[**UpdateUserPersonalAccessToken**](UsersAPI.md#UpdateUserPersonalAccessToken) | **Patch** /api/users/{userId}/personal-access-tokens/{name} | Update personal access token
[**UpdateUserProfile**](UsersAPI.md#UpdateUserProfile) | **Patch** /api/users/{userId}/profile | Update user profile
[**VerifyUserPassword**](UsersAPI.md#VerifyUserPassword) | **Post** /api/users/{userId}/password/verify | Verify user password



## AssignUserRoles

> AssignUserRoles(ctx, userId).AssignApplicationRolesRequest(assignApplicationRolesRequest).Execute()

Assign roles to user



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
	userId := "userId_example" // string | The unique identifier of the user.
	assignApplicationRolesRequest := *openapiclient.NewAssignApplicationRolesRequest([]string{"RoleIds_example"}) // AssignApplicationRolesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AssignUserRoles(context.Background(), userId).AssignApplicationRolesRequest(assignApplicationRolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AssignUserRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignUserRolesRequest struct via the builder pattern


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


## CreateUser

> UpdateUser200Response CreateUser(ctx).CreateUserRequest(createUserRequest).Execute()

Create user



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
	createUserRequest := *openapiclient.NewCreateUserRequest() // CreateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CreateUser(context.Background()).CreateUserRequest(createUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: UpdateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserRequest** | [**CreateUserRequest**](CreateUserRequest.md) |  | 

### Return type

[**UpdateUser200Response**](UpdateUser200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserIdentity

> map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue CreateUserIdentity(ctx, userId).CreateUserIdentityRequest(createUserIdentityRequest).Execute()

Link social identity to user



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
	userId := "userId_example" // string | The unique identifier of the user.
	createUserIdentityRequest := *openapiclient.NewCreateUserIdentityRequest("ConnectorId_example", map[string]interface{}{"key": interface{}({})}) // CreateUserIdentityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CreateUserIdentity(context.Background(), userId).CreateUserIdentityRequest(createUserIdentityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUserIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserIdentity`: map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUserIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUserIdentityRequest** | [**CreateUserIdentityRequest**](CreateUserIdentityRequest.md) |  | 

### Return type

[**map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue**](GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserMfaVerification

> CreateUserMfaVerification200Response CreateUserMfaVerification(ctx, userId).CreateUserMfaVerificationRequest(createUserMfaVerificationRequest).Execute()

Create an MFA verification for a user



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
	userId := "userId_example" // string | The unique identifier of the user.
	createUserMfaVerificationRequest := openapiclient.CreateUserMfaVerification_request{CreateUserMfaVerificationRequestOneOf: openapiclient.NewCreateUserMfaVerificationRequestOneOf("Type_example")} // CreateUserMfaVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CreateUserMfaVerification(context.Background(), userId).CreateUserMfaVerificationRequest(createUserMfaVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUserMfaVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserMfaVerification`: CreateUserMfaVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUserMfaVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserMfaVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUserMfaVerificationRequest** | [**CreateUserMfaVerificationRequest**](CreateUserMfaVerificationRequest.md) |  | 

### Return type

[**CreateUserMfaVerification200Response**](CreateUserMfaVerification200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserPersonalAccessToken

> ListUserPersonalAccessTokens200ResponseInner CreateUserPersonalAccessToken(ctx, userId).CreateUserPersonalAccessTokenRequest(createUserPersonalAccessTokenRequest).Execute()

Add personal access token



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
	userId := "userId_example" // string | The unique identifier of the user.
	createUserPersonalAccessTokenRequest := *openapiclient.NewCreateUserPersonalAccessTokenRequest("Name_example") // CreateUserPersonalAccessTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CreateUserPersonalAccessToken(context.Background(), userId).CreateUserPersonalAccessTokenRequest(createUserPersonalAccessTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUserPersonalAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserPersonalAccessToken`: ListUserPersonalAccessTokens200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUserPersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserPersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUserPersonalAccessTokenRequest** | [**CreateUserPersonalAccessTokenRequest**](CreateUserPersonalAccessTokenRequest.md) |  | 

### Return type

[**ListUserPersonalAccessTokens200ResponseInner**](ListUserPersonalAccessTokens200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId).Execute()

Delete user



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
	userId := "userId_example" // string | The unique identifier of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.DeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## DeleteUserIdentity

> UpdateUser200Response DeleteUserIdentity(ctx, userId, target).Execute()

Delete social identity from user



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
	userId := "userId_example" // string | The unique identifier of the user.
	target := "target_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.DeleteUserIdentity(context.Background(), userId, target).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserIdentity`: UpdateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeleteUserIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 
**target** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UpdateUser200Response**](UpdateUser200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserMfaVerification

> DeleteUserMfaVerification(ctx, userId, verificationId).Execute()

Delete an MFA verification for a user



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
	userId := "userId_example" // string | The unique identifier of the user.
	verificationId := "verificationId_example" // string | The unique identifier of the verification.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.DeleteUserMfaVerification(context.Background(), userId, verificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserMfaVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 
**verificationId** | **string** | The unique identifier of the verification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserMfaVerificationRequest struct via the builder pattern


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


## DeleteUserPersonalAccessToken

> DeleteUserPersonalAccessToken(ctx, userId, name).Execute()

Delete personal access token



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
	userId := "userId_example" // string | The unique identifier of the user.
	name := "name_example" // string | The name of the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.DeleteUserPersonalAccessToken(context.Background(), userId, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserPersonalAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 
**name** | **string** | The name of the token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserPersonalAccessTokenRequest struct via the builder pattern


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


## DeleteUserRole

> DeleteUserRole(ctx, userId, roleId).Execute()

Remove role from user



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
	userId := "userId_example" // string | The unique identifier of the user.
	roleId := "roleId_example" // string | The unique identifier of the role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.DeleteUserRole(context.Background(), userId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 
**roleId** | **string** | The unique identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRoleRequest struct via the builder pattern


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


## GetUser

> GetUser200Response GetUser(ctx, userId).IncludeSsoIdentities(includeSsoIdentities).Execute()

Get user



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
	userId := "userId_example" // string | The unique identifier of the user.
	includeSsoIdentities := "includeSsoIdentities_example" // string | If it's provided with a truthy value (`true`, `1`, `yes`), each user in the response will include a `ssoIdentities` property containing a list of SSO identities associated with the user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUser(context.Background(), userId).IncludeSsoIdentities(includeSsoIdentities).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: GetUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSsoIdentities** | **string** | If it&#39;s provided with a truthy value (&#x60;true&#x60;, &#x60;1&#x60;, &#x60;yes&#x60;), each user in the response will include a &#x60;ssoIdentities&#x60; property containing a list of SSO identities associated with the user. | 

### Return type

[**GetUser200Response**](GetUser200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserHasPassword

> GetUserHasPassword200Response GetUserHasPassword(ctx, userId).Execute()

Check if user has password



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
	userId := "userId_example" // string | The unique identifier of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserHasPassword(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserHasPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserHasPassword`: GetUserHasPassword200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserHasPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserHasPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserHasPassword200Response**](GetUserHasPassword200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserCustomData

> map[string]interface{} ListUserCustomData(ctx, userId).Execute()

Get user custom data



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
	userId := "userId_example" // string | The unique identifier of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ListUserCustomData(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserCustomData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserCustomData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserCustomData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserCustomDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserMfaVerifications

> []ListUserMfaVerifications200ResponseInner ListUserMfaVerifications(ctx, userId).Execute()

Get user's MFA verifications



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
	userId := "userId_example" // string | The unique identifier of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ListUserMfaVerifications(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserMfaVerifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserMfaVerifications`: []ListUserMfaVerifications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserMfaVerifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserMfaVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListUserMfaVerifications200ResponseInner**](ListUserMfaVerifications200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserOrganizations

> []ListApplicationOrganizations200ResponseInner ListUserOrganizations(ctx, userId).Execute()

Get organizations for a user



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
	userId := "userId_example" // string | The unique identifier of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ListUserOrganizations(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserOrganizations`: []ListApplicationOrganizations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserOrganizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListUserPersonalAccessTokens

> []ListUserPersonalAccessTokens200ResponseInner ListUserPersonalAccessTokens(ctx, userId).Execute()

Get personal access tokens



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
	userId := "userId_example" // string | The unique identifier of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ListUserPersonalAccessTokens(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserPersonalAccessTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserPersonalAccessTokens`: []ListUserPersonalAccessTokens200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserPersonalAccessTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserPersonalAccessTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListUserPersonalAccessTokens200ResponseInner**](ListUserPersonalAccessTokens200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserRoles

> []ListApplicationRoles200ResponseInner ListUserRoles(ctx, userId).Page(page).PageSize(pageSize).SearchParams(searchParams).Execute()

Get roles for user



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
	userId := "userId_example" // string | The unique identifier of the user.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)
	searchParams := map[string]string{"key": "Inner_example"} // map[string]string | Search query parameters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ListUserRoles(context.Background(), userId).Page(page).PageSize(pageSize).SearchParams(searchParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserRoles`: []ListApplicationRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserRolesRequest struct via the builder pattern


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


## ListUsers

> []UpdateUser200Response ListUsers(ctx).Page(page).PageSize(pageSize).SearchParams(searchParams).Execute()

Get users



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
	searchParams := map[string]string{"key": "Inner_example"} // map[string]string | Search query parameters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ListUsers(context.Background()).Page(page).PageSize(pageSize).SearchParams(searchParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: []UpdateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]
 **searchParams** | **map[string]string** | Search query parameters. | 

### Return type

[**[]UpdateUser200Response**](UpdateUser200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceUserIdentity

> map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue ReplaceUserIdentity(ctx, userId, target).ReplaceUserIdentityRequest(replaceUserIdentityRequest).Execute()

Update social identity of user



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
	userId := "userId_example" // string | The unique identifier of the user.
	target := "target_example" // string | 
	replaceUserIdentityRequest := *openapiclient.NewReplaceUserIdentityRequest("UserId_example") // ReplaceUserIdentityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ReplaceUserIdentity(context.Background(), userId, target).ReplaceUserIdentityRequest(replaceUserIdentityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ReplaceUserIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceUserIdentity`: map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ReplaceUserIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 
**target** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceUserIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceUserIdentityRequest** | [**ReplaceUserIdentityRequest**](ReplaceUserIdentityRequest.md) |  | 

### Return type

[**map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue**](GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceUserRoles

> ReplaceUserRoles(ctx, userId).AssignApplicationRolesRequest(assignApplicationRolesRequest).Execute()

Update roles for user



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
	userId := "userId_example" // string | The unique identifier of the user.
	assignApplicationRolesRequest := *openapiclient.NewAssignApplicationRolesRequest([]string{"RoleIds_example"}) // AssignApplicationRolesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.ReplaceUserRoles(context.Background(), userId).AssignApplicationRolesRequest(assignApplicationRolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ReplaceUserRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceUserRolesRequest struct via the builder pattern


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


## UpdateUser

> UpdateUser200Response UpdateUser(ctx, userId).UpdateUserRequest(updateUserRequest).Execute()

Update user



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
	userId := "userId_example" // string | The unique identifier of the user.
	updateUserRequest := *openapiclient.NewUpdateUserRequest() // UpdateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUser(context.Background(), userId).UpdateUserRequest(updateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: UpdateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserRequest** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 

### Return type

[**UpdateUser200Response**](UpdateUser200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserCustomData

> map[string]interface{} UpdateUserCustomData(ctx, userId).UpdateUserCustomDataRequest(updateUserCustomDataRequest).Execute()

Update user custom data



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
	userId := "userId_example" // string | The unique identifier of the user.
	updateUserCustomDataRequest := *openapiclient.NewUpdateUserCustomDataRequest(map[string]interface{}(123)) // UpdateUserCustomDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserCustomData(context.Background(), userId).UpdateUserCustomDataRequest(updateUserCustomDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserCustomData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserCustomData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserCustomData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserCustomDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserCustomDataRequest** | [**UpdateUserCustomDataRequest**](UpdateUserCustomDataRequest.md) |  | 

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


## UpdateUserIsSuspended

> UpdateUser200Response UpdateUserIsSuspended(ctx, userId).UpdateUserIsSuspendedRequest(updateUserIsSuspendedRequest).Execute()

Update user suspension status



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
	userId := "userId_example" // string | The unique identifier of the user.
	updateUserIsSuspendedRequest := *openapiclient.NewUpdateUserIsSuspendedRequest(false) // UpdateUserIsSuspendedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserIsSuspended(context.Background(), userId).UpdateUserIsSuspendedRequest(updateUserIsSuspendedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserIsSuspended``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserIsSuspended`: UpdateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserIsSuspended`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserIsSuspendedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserIsSuspendedRequest** | [**UpdateUserIsSuspendedRequest**](UpdateUserIsSuspendedRequest.md) |  | 

### Return type

[**UpdateUser200Response**](UpdateUser200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPassword

> UpdateUser200Response UpdateUserPassword(ctx, userId).UpdateUserPasswordRequest(updateUserPasswordRequest).Execute()

Update user password



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
	userId := "userId_example" // string | The unique identifier of the user.
	updateUserPasswordRequest := *openapiclient.NewUpdateUserPasswordRequest("Password_example") // UpdateUserPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserPassword(context.Background(), userId).UpdateUserPasswordRequest(updateUserPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserPassword`: UpdateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserPasswordRequest** | [**UpdateUserPasswordRequest**](UpdateUserPasswordRequest.md) |  | 

### Return type

[**UpdateUser200Response**](UpdateUser200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPersonalAccessToken

> ListUserPersonalAccessTokens200ResponseInner UpdateUserPersonalAccessToken(ctx, userId, name).UpdateUserPersonalAccessTokenRequest(updateUserPersonalAccessTokenRequest).Execute()

Update personal access token



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
	userId := "userId_example" // string | The unique identifier of the user.
	name := "name_example" // string | The name of the token.
	updateUserPersonalAccessTokenRequest := *openapiclient.NewUpdateUserPersonalAccessTokenRequest("Name_example") // UpdateUserPersonalAccessTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserPersonalAccessToken(context.Background(), userId, name).UpdateUserPersonalAccessTokenRequest(updateUserPersonalAccessTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserPersonalAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserPersonalAccessToken`: ListUserPersonalAccessTokens200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserPersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 
**name** | **string** | The name of the token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateUserPersonalAccessTokenRequest** | [**UpdateUserPersonalAccessTokenRequest**](UpdateUserPersonalAccessTokenRequest.md) |  | 

### Return type

[**ListUserPersonalAccessTokens200ResponseInner**](ListUserPersonalAccessTokens200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserProfile

> GetJwtCustomizer200ResponseOneOfContextSampleUserProfile UpdateUserProfile(ctx, userId).UpdateUserProfileRequest(updateUserProfileRequest).Execute()

Update user profile



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
	userId := "userId_example" // string | The unique identifier of the user.
	updateUserProfileRequest := *openapiclient.NewUpdateUserProfileRequest(*openapiclient.NewUpdateUserProfileRequestProfile()) // UpdateUserProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserProfile(context.Background(), userId).UpdateUserProfileRequest(updateUserProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserProfile`: GetJwtCustomizer200ResponseOneOfContextSampleUserProfile
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserProfileRequest** | [**UpdateUserProfileRequest**](UpdateUserProfileRequest.md) |  | 

### Return type

[**GetJwtCustomizer200ResponseOneOfContextSampleUserProfile**](GetJwtCustomizer200ResponseOneOfContextSampleUserProfile.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyUserPassword

> VerifyUserPassword(ctx, userId).VerifyUserPasswordRequest(verifyUserPasswordRequest).Execute()

Verify user password



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
	userId := "userId_example" // string | The unique identifier of the user.
	verifyUserPasswordRequest := *openapiclient.NewVerifyUserPasswordRequest("Password_example") // VerifyUserPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.VerifyUserPassword(context.Background(), userId).VerifyUserPasswordRequest(verifyUserPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.VerifyUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyUserPasswordRequest** | [**VerifyUserPasswordRequest**](VerifyUserPasswordRequest.md) |  | 

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

