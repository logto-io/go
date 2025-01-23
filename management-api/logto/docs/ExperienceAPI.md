# \ExperienceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserProfile**](ExperienceAPI.md#AddUserProfile) | **Post** /api/experience/profile | Add user profile
[**BindMfaVerification**](ExperienceAPI.md#BindMfaVerification) | **Post** /api/experience/profile/mfa | Bind MFA verification by verificationId
[**CreateAndSendVerificationCode**](ExperienceAPI.md#CreateAndSendVerificationCode) | **Post** /api/experience/verification/verification-code | Create and send verification code
[**CreateEnterpriseSsoVerification**](ExperienceAPI.md#CreateEnterpriseSsoVerification) | **Post** /api/experience/verification/sso/{connectorId}/authorization-uri | Create enterprise SSO verification
[**CreateNewPasswordIdentityVerification**](ExperienceAPI.md#CreateNewPasswordIdentityVerification) | **Post** /api/experience/verification/new-password-identity | Create new password identity verification
[**CreatePasswordVerification**](ExperienceAPI.md#CreatePasswordVerification) | **Post** /api/experience/verification/password | Create password verification record
[**CreateSocialVerification**](ExperienceAPI.md#CreateSocialVerification) | **Post** /api/experience/verification/social/{connectorId}/authorization-uri | Create social verification
[**CreateTotpSecret**](ExperienceAPI.md#CreateTotpSecret) | **Post** /api/experience/verification/totp/secret | Create TOTP secret
[**CreateWebAuthnAuthenticationVerification**](ExperienceAPI.md#CreateWebAuthnAuthenticationVerification) | **Post** /api/experience/verification/web-authn/authentication | Create WebAuthn authentication verification
[**CreateWebAuthnRegistrationVerification**](ExperienceAPI.md#CreateWebAuthnRegistrationVerification) | **Post** /api/experience/verification/web-authn/registration | Create WebAuthn registration verification
[**GenerateBackupCodes**](ExperienceAPI.md#GenerateBackupCodes) | **Post** /api/experience/verification/backup-code/generate | Generate backup codes
[**GetEnabledSsoConnectors**](ExperienceAPI.md#GetEnabledSsoConnectors) | **Get** /api/experience/sso-connectors | Get enabled SSO connectors by the given email&#39;s domain
[**IdentifyUser**](ExperienceAPI.md#IdentifyUser) | **Post** /api/experience/identification | Identify user for the current interaction
[**InitInteraction**](ExperienceAPI.md#InitInteraction) | **Put** /api/experience | Init new interaction
[**ResetUserPassword**](ExperienceAPI.md#ResetUserPassword) | **Put** /api/experience/profile/password | Reset user password
[**SkipMfaBindingFlow**](ExperienceAPI.md#SkipMfaBindingFlow) | **Post** /api/experience/profile/mfa/mfa-skipped | Skip MFA binding flow
[**SubmitInteraction**](ExperienceAPI.md#SubmitInteraction) | **Post** /api/experience/submit | Submit interaction
[**UpdateInteractionEvent**](ExperienceAPI.md#UpdateInteractionEvent) | **Put** /api/experience/interaction-event | Update interaction event
[**VerifyBackupCode**](ExperienceAPI.md#VerifyBackupCode) | **Post** /api/experience/verification/backup-code/verify | Verify backup code
[**VerifyEnterpriseSsoVerification**](ExperienceAPI.md#VerifyEnterpriseSsoVerification) | **Post** /api/experience/verification/sso/{connectorId}/verify | Verify enterprise SSO verification
[**VerifySocialVerification**](ExperienceAPI.md#VerifySocialVerification) | **Post** /api/experience/verification/social/{connectorId}/verify | Verify social verification
[**VerifyTotpVerification**](ExperienceAPI.md#VerifyTotpVerification) | **Post** /api/experience/verification/totp/verify | Verify TOTP verification
[**VerifyVerificationCodeVerification**](ExperienceAPI.md#VerifyVerificationCodeVerification) | **Post** /api/experience/verification/verification-code/verify | Verify verification code
[**VerifyWebAuthnAuthenticationVerification**](ExperienceAPI.md#VerifyWebAuthnAuthenticationVerification) | **Post** /api/experience/verification/web-authn/authentication/verify | Verify WebAuthn authentication verification
[**VerifyWebAuthnRegistrationVerification**](ExperienceAPI.md#VerifyWebAuthnRegistrationVerification) | **Post** /api/experience/verification/web-authn/registration/verify | Verify WebAuthn registration verification



## AddUserProfile

> AddUserProfile(ctx).AddUserProfileRequest(addUserProfileRequest).Execute()

Add user profile



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
	addUserProfileRequest := openapiclient.AddUserProfile_request{AddUserProfileRequestOneOf: openapiclient.NewAddUserProfileRequestOneOf("Type_example", "Value_example")} // AddUserProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperienceAPI.AddUserProfile(context.Background()).AddUserProfileRequest(addUserProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.AddUserProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addUserProfileRequest** | [**AddUserProfileRequest**](AddUserProfileRequest.md) |  | 

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


## BindMfaVerification

> BindMfaVerification(ctx).BindMfaVerificationRequest(bindMfaVerificationRequest).Execute()

Bind MFA verification by verificationId



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
	bindMfaVerificationRequest := *openapiclient.NewBindMfaVerificationRequest("Type_example", "VerificationId_example") // BindMfaVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperienceAPI.BindMfaVerification(context.Background()).BindMfaVerificationRequest(bindMfaVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.BindMfaVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBindMfaVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindMfaVerificationRequest** | [**BindMfaVerificationRequest**](BindMfaVerificationRequest.md) |  | 

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


## CreateAndSendVerificationCode

> CreateAndSendVerificationCode200Response CreateAndSendVerificationCode(ctx).CreateAndSendVerificationCodeRequest(createAndSendVerificationCodeRequest).Execute()

Create and send verification code



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
	createAndSendVerificationCodeRequest := *openapiclient.NewCreateAndSendVerificationCodeRequest(*openapiclient.NewCreateAndSendVerificationCodeRequestIdentifier("Type_example", "Value_example"), "InteractionEvent_example") // CreateAndSendVerificationCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.CreateAndSendVerificationCode(context.Background()).CreateAndSendVerificationCodeRequest(createAndSendVerificationCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.CreateAndSendVerificationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAndSendVerificationCode`: CreateAndSendVerificationCode200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.CreateAndSendVerificationCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAndSendVerificationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAndSendVerificationCodeRequest** | [**CreateAndSendVerificationCodeRequest**](CreateAndSendVerificationCodeRequest.md) |  | 

### Return type

[**CreateAndSendVerificationCode200Response**](CreateAndSendVerificationCode200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnterpriseSsoVerification

> CreateEnterpriseSsoVerification200Response CreateEnterpriseSsoVerification(ctx, connectorId).CreateEnterpriseSsoVerificationRequest(createEnterpriseSsoVerificationRequest).Execute()

Create enterprise SSO verification



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
	createEnterpriseSsoVerificationRequest := *openapiclient.NewCreateEnterpriseSsoVerificationRequest("State_example", "RedirectUri_example") // CreateEnterpriseSsoVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.CreateEnterpriseSsoVerification(context.Background(), connectorId).CreateEnterpriseSsoVerificationRequest(createEnterpriseSsoVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.CreateEnterpriseSsoVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnterpriseSsoVerification`: CreateEnterpriseSsoVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.CreateEnterpriseSsoVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnterpriseSsoVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEnterpriseSsoVerificationRequest** | [**CreateEnterpriseSsoVerificationRequest**](CreateEnterpriseSsoVerificationRequest.md) |  | 

### Return type

[**CreateEnterpriseSsoVerification200Response**](CreateEnterpriseSsoVerification200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewPasswordIdentityVerification

> CreateNewPasswordIdentityVerification200Response CreateNewPasswordIdentityVerification(ctx).CreateNewPasswordIdentityVerificationRequest(createNewPasswordIdentityVerificationRequest).Execute()

Create new password identity verification



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
	createNewPasswordIdentityVerificationRequest := *openapiclient.NewCreateNewPasswordIdentityVerificationRequest(*openapiclient.NewCreateNewPasswordIdentityVerificationRequestIdentifier("Type_example", "Value_example"), "Password_example") // CreateNewPasswordIdentityVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.CreateNewPasswordIdentityVerification(context.Background()).CreateNewPasswordIdentityVerificationRequest(createNewPasswordIdentityVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.CreateNewPasswordIdentityVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewPasswordIdentityVerification`: CreateNewPasswordIdentityVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.CreateNewPasswordIdentityVerification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewPasswordIdentityVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNewPasswordIdentityVerificationRequest** | [**CreateNewPasswordIdentityVerificationRequest**](CreateNewPasswordIdentityVerificationRequest.md) |  | 

### Return type

[**CreateNewPasswordIdentityVerification200Response**](CreateNewPasswordIdentityVerification200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePasswordVerification

> CreatePasswordVerification200Response CreatePasswordVerification(ctx).CreatePasswordVerificationRequest(createPasswordVerificationRequest).Execute()

Create password verification record



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
	createPasswordVerificationRequest := *openapiclient.NewCreatePasswordVerificationRequest(*openapiclient.NewCreatePasswordVerificationRequestIdentifier("Type_example", "Value_example"), "Password_example") // CreatePasswordVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.CreatePasswordVerification(context.Background()).CreatePasswordVerificationRequest(createPasswordVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.CreatePasswordVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePasswordVerification`: CreatePasswordVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.CreatePasswordVerification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePasswordVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPasswordVerificationRequest** | [**CreatePasswordVerificationRequest**](CreatePasswordVerificationRequest.md) |  | 

### Return type

[**CreatePasswordVerification200Response**](CreatePasswordVerification200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSocialVerification

> CreateSocialVerification200Response CreateSocialVerification(ctx, connectorId).CreateSocialVerificationRequest(createSocialVerificationRequest).Execute()

Create social verification



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
	createSocialVerificationRequest := *openapiclient.NewCreateSocialVerificationRequest("State_example", "RedirectUri_example") // CreateSocialVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.CreateSocialVerification(context.Background(), connectorId).CreateSocialVerificationRequest(createSocialVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.CreateSocialVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSocialVerification`: CreateSocialVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.CreateSocialVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSocialVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSocialVerificationRequest** | [**CreateSocialVerificationRequest**](CreateSocialVerificationRequest.md) |  | 

### Return type

[**CreateSocialVerification200Response**](CreateSocialVerification200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTotpSecret

> CreateTotpSecret200Response CreateTotpSecret(ctx).Execute()

Create TOTP secret



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
	resp, r, err := apiClient.ExperienceAPI.CreateTotpSecret(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.CreateTotpSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTotpSecret`: CreateTotpSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.CreateTotpSecret`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTotpSecretRequest struct via the builder pattern


### Return type

[**CreateTotpSecret200Response**](CreateTotpSecret200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebAuthnAuthenticationVerification

> CreateWebAuthnAuthenticationVerification200Response CreateWebAuthnAuthenticationVerification(ctx).Execute()

Create WebAuthn authentication verification



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
	resp, r, err := apiClient.ExperienceAPI.CreateWebAuthnAuthenticationVerification(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.CreateWebAuthnAuthenticationVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebAuthnAuthenticationVerification`: CreateWebAuthnAuthenticationVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.CreateWebAuthnAuthenticationVerification`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebAuthnAuthenticationVerificationRequest struct via the builder pattern


### Return type

[**CreateWebAuthnAuthenticationVerification200Response**](CreateWebAuthnAuthenticationVerification200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebAuthnRegistrationVerification

> CreateWebAuthnRegistrationVerification200Response CreateWebAuthnRegistrationVerification(ctx).Execute()

Create WebAuthn registration verification



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
	resp, r, err := apiClient.ExperienceAPI.CreateWebAuthnRegistrationVerification(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.CreateWebAuthnRegistrationVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebAuthnRegistrationVerification`: CreateWebAuthnRegistrationVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.CreateWebAuthnRegistrationVerification`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebAuthnRegistrationVerificationRequest struct via the builder pattern


### Return type

[**CreateWebAuthnRegistrationVerification200Response**](CreateWebAuthnRegistrationVerification200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateBackupCodes

> GenerateBackupCodes200Response GenerateBackupCodes(ctx).Execute()

Generate backup codes



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
	resp, r, err := apiClient.ExperienceAPI.GenerateBackupCodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.GenerateBackupCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateBackupCodes`: GenerateBackupCodes200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.GenerateBackupCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateBackupCodesRequest struct via the builder pattern


### Return type

[**GenerateBackupCodes200Response**](GenerateBackupCodes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnabledSsoConnectors

> GetEnabledSsoConnectors200Response GetEnabledSsoConnectors(ctx).Email(email).Execute()

Get enabled SSO connectors by the given email's domain



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
	email := "email_example" // string | The email address to find the enabled SSO connectors.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.GetEnabledSsoConnectors(context.Background()).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.GetEnabledSsoConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnabledSsoConnectors`: GetEnabledSsoConnectors200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.GetEnabledSsoConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnabledSsoConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | The email address to find the enabled SSO connectors. | 

### Return type

[**GetEnabledSsoConnectors200Response**](GetEnabledSsoConnectors200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentifyUser

> IdentifyUser(ctx).IdentifyUserRequest(identifyUserRequest).Execute()

Identify user for the current interaction



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
	identifyUserRequest := *openapiclient.NewIdentifyUserRequest() // IdentifyUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperienceAPI.IdentifyUser(context.Background()).IdentifyUserRequest(identifyUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.IdentifyUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentifyUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifyUserRequest** | [**IdentifyUserRequest**](IdentifyUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitInteraction

> InitInteraction(ctx).InitInteractionRequest(initInteractionRequest).Execute()

Init new interaction



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
	initInteractionRequest := *openapiclient.NewInitInteractionRequest("InteractionEvent_example") // InitInteractionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperienceAPI.InitInteraction(context.Background()).InitInteractionRequest(initInteractionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.InitInteraction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitInteractionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **initInteractionRequest** | [**InitInteractionRequest**](InitInteractionRequest.md) |  | 

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


## ResetUserPassword

> ResetUserPassword(ctx).ResetUserPasswordRequest(resetUserPasswordRequest).Execute()

Reset user password



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
	resetUserPasswordRequest := *openapiclient.NewResetUserPasswordRequest("Password_example") // ResetUserPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperienceAPI.ResetUserPassword(context.Background()).ResetUserPasswordRequest(resetUserPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.ResetUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetUserPasswordRequest** | [**ResetUserPasswordRequest**](ResetUserPasswordRequest.md) |  | 

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


## SkipMfaBindingFlow

> SkipMfaBindingFlow(ctx).Execute()

Skip MFA binding flow



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
	r, err := apiClient.ExperienceAPI.SkipMfaBindingFlow(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.SkipMfaBindingFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSkipMfaBindingFlowRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitInteraction

> SubmitInteraction200Response SubmitInteraction(ctx).Execute()

Submit interaction



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
	resp, r, err := apiClient.ExperienceAPI.SubmitInteraction(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.SubmitInteraction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitInteraction`: SubmitInteraction200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.SubmitInteraction`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitInteractionRequest struct via the builder pattern


### Return type

[**SubmitInteraction200Response**](SubmitInteraction200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInteractionEvent

> UpdateInteractionEvent(ctx).UpdateInteractionEventRequest(updateInteractionEventRequest).Execute()

Update interaction event



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
	updateInteractionEventRequest := *openapiclient.NewUpdateInteractionEventRequest("InteractionEvent_example") // UpdateInteractionEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperienceAPI.UpdateInteractionEvent(context.Background()).UpdateInteractionEventRequest(updateInteractionEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.UpdateInteractionEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInteractionEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateInteractionEventRequest** | [**UpdateInteractionEventRequest**](UpdateInteractionEventRequest.md) |  | 

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


## VerifyBackupCode

> VerifyBackupCode200Response VerifyBackupCode(ctx).VerifyBackupCodeRequest(verifyBackupCodeRequest).Execute()

Verify backup code



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
	verifyBackupCodeRequest := *openapiclient.NewVerifyBackupCodeRequest("Code_example") // VerifyBackupCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.VerifyBackupCode(context.Background()).VerifyBackupCodeRequest(verifyBackupCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.VerifyBackupCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyBackupCode`: VerifyBackupCode200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.VerifyBackupCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyBackupCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyBackupCodeRequest** | [**VerifyBackupCodeRequest**](VerifyBackupCodeRequest.md) |  | 

### Return type

[**VerifyBackupCode200Response**](VerifyBackupCode200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEnterpriseSsoVerification

> VerifyEnterpriseSsoVerification200Response VerifyEnterpriseSsoVerification(ctx, connectorId).VerifyEnterpriseSsoVerificationRequest(verifyEnterpriseSsoVerificationRequest).Execute()

Verify enterprise SSO verification



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
	verifyEnterpriseSsoVerificationRequest := *openapiclient.NewVerifyEnterpriseSsoVerificationRequest(map[string]interface{}(123), "VerificationId_example") // VerifyEnterpriseSsoVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.VerifyEnterpriseSsoVerification(context.Background(), connectorId).VerifyEnterpriseSsoVerificationRequest(verifyEnterpriseSsoVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.VerifyEnterpriseSsoVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyEnterpriseSsoVerification`: VerifyEnterpriseSsoVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.VerifyEnterpriseSsoVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyEnterpriseSsoVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyEnterpriseSsoVerificationRequest** | [**VerifyEnterpriseSsoVerificationRequest**](VerifyEnterpriseSsoVerificationRequest.md) |  | 

### Return type

[**VerifyEnterpriseSsoVerification200Response**](VerifyEnterpriseSsoVerification200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifySocialVerification

> VerifySocialVerification200Response VerifySocialVerification(ctx, connectorId).VerifySocialVerificationRequest(verifySocialVerificationRequest).Execute()

Verify social verification



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
	verifySocialVerificationRequest := *openapiclient.NewVerifySocialVerificationRequest(map[string]interface{}(123)) // VerifySocialVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.VerifySocialVerification(context.Background(), connectorId).VerifySocialVerificationRequest(verifySocialVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.VerifySocialVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifySocialVerification`: VerifySocialVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.VerifySocialVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifySocialVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifySocialVerificationRequest** | [**VerifySocialVerificationRequest**](VerifySocialVerificationRequest.md) |  | 

### Return type

[**VerifySocialVerification200Response**](VerifySocialVerification200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyTotpVerification

> VerifyTotpVerification200Response VerifyTotpVerification(ctx).VerifyTotpVerificationRequest(verifyTotpVerificationRequest).Execute()

Verify TOTP verification



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
	verifyTotpVerificationRequest := *openapiclient.NewVerifyTotpVerificationRequest("Code_example") // VerifyTotpVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.VerifyTotpVerification(context.Background()).VerifyTotpVerificationRequest(verifyTotpVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.VerifyTotpVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyTotpVerification`: VerifyTotpVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.VerifyTotpVerification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyTotpVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyTotpVerificationRequest** | [**VerifyTotpVerificationRequest**](VerifyTotpVerificationRequest.md) |  | 

### Return type

[**VerifyTotpVerification200Response**](VerifyTotpVerification200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyVerificationCodeVerification

> VerifyVerificationCodeVerification200Response VerifyVerificationCodeVerification(ctx).VerifyVerificationCodeVerificationRequest(verifyVerificationCodeVerificationRequest).Execute()

Verify verification code



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
	verifyVerificationCodeVerificationRequest := *openapiclient.NewVerifyVerificationCodeVerificationRequest(*openapiclient.NewVerifyVerificationCodeVerificationRequestIdentifier("Type_example", "Value_example"), "VerificationId_example", "Code_example") // VerifyVerificationCodeVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.VerifyVerificationCodeVerification(context.Background()).VerifyVerificationCodeVerificationRequest(verifyVerificationCodeVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.VerifyVerificationCodeVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyVerificationCodeVerification`: VerifyVerificationCodeVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.VerifyVerificationCodeVerification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyVerificationCodeVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyVerificationCodeVerificationRequest** | [**VerifyVerificationCodeVerificationRequest**](VerifyVerificationCodeVerificationRequest.md) |  | 

### Return type

[**VerifyVerificationCodeVerification200Response**](VerifyVerificationCodeVerification200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyWebAuthnAuthenticationVerification

> VerifyWebAuthnAuthenticationVerification200Response VerifyWebAuthnAuthenticationVerification(ctx).VerifyWebAuthnAuthenticationVerificationRequest(verifyWebAuthnAuthenticationVerificationRequest).Execute()

Verify WebAuthn authentication verification



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
	verifyWebAuthnAuthenticationVerificationRequest := *openapiclient.NewVerifyWebAuthnAuthenticationVerificationRequest("VerificationId_example", *openapiclient.NewVerifyWebAuthnAuthenticationVerificationRequestPayload("Type_example", "Id_example", "RawId_example", *openapiclient.NewVerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults(), *openapiclient.NewVerifyWebAuthnAuthenticationVerificationRequestPayloadResponse("ClientDataJSON_example", "AuthenticatorData_example", "Signature_example"))) // VerifyWebAuthnAuthenticationVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.VerifyWebAuthnAuthenticationVerification(context.Background()).VerifyWebAuthnAuthenticationVerificationRequest(verifyWebAuthnAuthenticationVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.VerifyWebAuthnAuthenticationVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyWebAuthnAuthenticationVerification`: VerifyWebAuthnAuthenticationVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.VerifyWebAuthnAuthenticationVerification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyWebAuthnAuthenticationVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyWebAuthnAuthenticationVerificationRequest** | [**VerifyWebAuthnAuthenticationVerificationRequest**](VerifyWebAuthnAuthenticationVerificationRequest.md) |  | 

### Return type

[**VerifyWebAuthnAuthenticationVerification200Response**](VerifyWebAuthnAuthenticationVerification200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyWebAuthnRegistrationVerification

> VerifyWebAuthnRegistrationVerification200Response VerifyWebAuthnRegistrationVerification(ctx).VerifyWebAuthnRegistrationVerificationRequest(verifyWebAuthnRegistrationVerificationRequest).Execute()

Verify WebAuthn registration verification



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
	verifyWebAuthnRegistrationVerificationRequest := *openapiclient.NewVerifyWebAuthnRegistrationVerificationRequest("VerificationId_example", *openapiclient.NewVerifyWebAuthnRegistrationVerificationRequestPayload("Type_example", "Id_example", "RawId_example", *openapiclient.NewVerifyWebAuthnRegistrationVerificationRequestPayloadResponse("ClientDataJSON_example", "AttestationObject_example"), *openapiclient.NewVerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults())) // VerifyWebAuthnRegistrationVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperienceAPI.VerifyWebAuthnRegistrationVerification(context.Background()).VerifyWebAuthnRegistrationVerificationRequest(verifyWebAuthnRegistrationVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperienceAPI.VerifyWebAuthnRegistrationVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyWebAuthnRegistrationVerification`: VerifyWebAuthnRegistrationVerification200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperienceAPI.VerifyWebAuthnRegistrationVerification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyWebAuthnRegistrationVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyWebAuthnRegistrationVerificationRequest** | [**VerifyWebAuthnRegistrationVerificationRequest**](VerifyWebAuthnRegistrationVerificationRequest.md) |  | 

### Return type

[**VerifyWebAuthnRegistrationVerification200Response**](VerifyWebAuthnRegistrationVerification200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

