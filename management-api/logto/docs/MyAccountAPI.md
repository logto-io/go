# \MyAccountAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserIdentities**](MyAccountAPI.md#AddUserIdentities) | **Post** /api/my-account/identities | Add a user identity
[**DeleteIdentity**](MyAccountAPI.md#DeleteIdentity) | **Delete** /api/my-account/identities/{target} | Delete a user identity
[**DeletePrimaryEmail**](MyAccountAPI.md#DeletePrimaryEmail) | **Delete** /api/my-account/primary-email | Delete primary email
[**DeletePrimaryPhone**](MyAccountAPI.md#DeletePrimaryPhone) | **Delete** /api/my-account/primary-phone | Delete primary phone
[**GetProfile**](MyAccountAPI.md#GetProfile) | **Get** /api/my-account | Get profile
[**UpdateOtherProfile**](MyAccountAPI.md#UpdateOtherProfile) | **Patch** /api/my-account/profile | Update other profile
[**UpdatePassword**](MyAccountAPI.md#UpdatePassword) | **Post** /api/my-account/password | Update password
[**UpdatePrimaryEmail**](MyAccountAPI.md#UpdatePrimaryEmail) | **Post** /api/my-account/primary-email | Update primary email
[**UpdatePrimaryPhone**](MyAccountAPI.md#UpdatePrimaryPhone) | **Post** /api/my-account/primary-phone | Update primary phone
[**UpdateProfile**](MyAccountAPI.md#UpdateProfile) | **Patch** /api/my-account | Update profile



## AddUserIdentities

> AddUserIdentities(ctx).AddUserIdentitiesRequest(addUserIdentitiesRequest).Execute()

Add a user identity



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
	addUserIdentitiesRequest := *openapiclient.NewAddUserIdentitiesRequest("NewIdentifierVerificationRecordId_example") // AddUserIdentitiesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyAccountAPI.AddUserIdentities(context.Background()).AddUserIdentitiesRequest(addUserIdentitiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.AddUserIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addUserIdentitiesRequest** | [**AddUserIdentitiesRequest**](AddUserIdentitiesRequest.md) |  | 

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


## DeleteIdentity

> DeleteIdentity(ctx, target).Execute()

Delete a user identity



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
	target := "target_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyAccountAPI.DeleteIdentity(context.Background(), target).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.DeleteIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityRequest struct via the builder pattern


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


## DeletePrimaryEmail

> DeletePrimaryEmail(ctx).Execute()

Delete primary email



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
	r, err := apiClient.MyAccountAPI.DeletePrimaryEmail(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.DeletePrimaryEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrimaryEmailRequest struct via the builder pattern


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


## DeletePrimaryPhone

> DeletePrimaryPhone(ctx).Execute()

Delete primary phone



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
	r, err := apiClient.MyAccountAPI.DeletePrimaryPhone(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.DeletePrimaryPhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrimaryPhoneRequest struct via the builder pattern


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


## GetProfile

> GetProfile200Response GetProfile(ctx).Execute()

Get profile



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
	resp, r, err := apiClient.MyAccountAPI.GetProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.GetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfile`: GetProfile200Response
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.GetProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


### Return type

[**GetProfile200Response**](GetProfile200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOtherProfile

> GetJwtCustomizer200ResponseOneOfContextSampleUserProfile UpdateOtherProfile(ctx).UpdateOtherProfileRequest(updateOtherProfileRequest).Execute()

Update other profile



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
	updateOtherProfileRequest := *openapiclient.NewUpdateOtherProfileRequest() // UpdateOtherProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountAPI.UpdateOtherProfile(context.Background()).UpdateOtherProfileRequest(updateOtherProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.UpdateOtherProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOtherProfile`: GetJwtCustomizer200ResponseOneOfContextSampleUserProfile
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.UpdateOtherProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOtherProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOtherProfileRequest** | [**UpdateOtherProfileRequest**](UpdateOtherProfileRequest.md) |  | 

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


## UpdatePassword

> UpdatePassword(ctx).UpdatePasswordRequest(updatePasswordRequest).Execute()

Update password



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
	updatePasswordRequest := *openapiclient.NewUpdatePasswordRequest("Password_example") // UpdatePasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyAccountAPI.UpdatePassword(context.Background()).UpdatePasswordRequest(updatePasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.UpdatePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePasswordRequest** | [**UpdatePasswordRequest**](UpdatePasswordRequest.md) |  | 

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


## UpdatePrimaryEmail

> UpdatePrimaryEmail(ctx).UpdatePrimaryEmailRequest(updatePrimaryEmailRequest).Execute()

Update primary email



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
	updatePrimaryEmailRequest := *openapiclient.NewUpdatePrimaryEmailRequest("Email_example", "NewIdentifierVerificationRecordId_example") // UpdatePrimaryEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyAccountAPI.UpdatePrimaryEmail(context.Background()).UpdatePrimaryEmailRequest(updatePrimaryEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.UpdatePrimaryEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrimaryEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePrimaryEmailRequest** | [**UpdatePrimaryEmailRequest**](UpdatePrimaryEmailRequest.md) |  | 

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


## UpdatePrimaryPhone

> UpdatePrimaryPhone(ctx).UpdatePrimaryPhoneRequest(updatePrimaryPhoneRequest).Execute()

Update primary phone



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
	updatePrimaryPhoneRequest := *openapiclient.NewUpdatePrimaryPhoneRequest("Phone_example", "NewIdentifierVerificationRecordId_example") // UpdatePrimaryPhoneRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyAccountAPI.UpdatePrimaryPhone(context.Background()).UpdatePrimaryPhoneRequest(updatePrimaryPhoneRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.UpdatePrimaryPhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrimaryPhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePrimaryPhoneRequest** | [**UpdatePrimaryPhoneRequest**](UpdatePrimaryPhoneRequest.md) |  | 

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


## UpdateProfile

> GetProfile200Response UpdateProfile(ctx).UpdateProfileRequest(updateProfileRequest).Execute()

Update profile



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
	updateProfileRequest := *openapiclient.NewUpdateProfileRequest() // UpdateProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountAPI.UpdateProfile(context.Background()).UpdateProfileRequest(updateProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountAPI.UpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProfile`: GetProfile200Response
	fmt.Fprintf(os.Stdout, "Response from `MyAccountAPI.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateProfileRequest** | [**UpdateProfileRequest**](UpdateProfileRequest.md) |  | 

### Return type

[**GetProfile200Response**](GetProfile200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

