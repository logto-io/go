# \InteractionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiInteractionBindMfaPost**](InteractionAPI.md#ApiInteractionBindMfaPost) | **Post** /api/interaction/bind-mfa | 
[**ApiInteractionConsentGet**](InteractionAPI.md#ApiInteractionConsentGet) | **Get** /api/interaction/consent | 
[**ApiInteractionConsentPost**](InteractionAPI.md#ApiInteractionConsentPost) | **Post** /api/interaction/consent | 
[**ApiInteractionDelete**](InteractionAPI.md#ApiInteractionDelete) | **Delete** /api/interaction | 
[**ApiInteractionEventPut**](InteractionAPI.md#ApiInteractionEventPut) | **Put** /api/interaction/event | 
[**ApiInteractionIdentifiersPatch**](InteractionAPI.md#ApiInteractionIdentifiersPatch) | **Patch** /api/interaction/identifiers | 
[**ApiInteractionMfaPut**](InteractionAPI.md#ApiInteractionMfaPut) | **Put** /api/interaction/mfa | 
[**ApiInteractionMfaSkippedPut**](InteractionAPI.md#ApiInteractionMfaSkippedPut) | **Put** /api/interaction/mfa-skipped | 
[**ApiInteractionProfileDelete**](InteractionAPI.md#ApiInteractionProfileDelete) | **Delete** /api/interaction/profile | 
[**ApiInteractionProfilePatch**](InteractionAPI.md#ApiInteractionProfilePatch) | **Patch** /api/interaction/profile | 
[**ApiInteractionProfilePut**](InteractionAPI.md#ApiInteractionProfilePut) | **Put** /api/interaction/profile | 
[**ApiInteractionPut**](InteractionAPI.md#ApiInteractionPut) | **Put** /api/interaction | 
[**ApiInteractionSingleSignOnConnectorIdAuthenticationPost**](InteractionAPI.md#ApiInteractionSingleSignOnConnectorIdAuthenticationPost) | **Post** /api/interaction/single-sign-on/{connectorId}/authentication | 
[**ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPost**](InteractionAPI.md#ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPost) | **Post** /api/interaction/single-sign-on/{connectorId}/authorization-url | 
[**ApiInteractionSingleSignOnConnectorIdRegistrationPost**](InteractionAPI.md#ApiInteractionSingleSignOnConnectorIdRegistrationPost) | **Post** /api/interaction/single-sign-on/{connectorId}/registration | 
[**ApiInteractionSingleSignOnConnectorsGet**](InteractionAPI.md#ApiInteractionSingleSignOnConnectorsGet) | **Get** /api/interaction/single-sign-on/connectors | 
[**ApiInteractionSubmitPost**](InteractionAPI.md#ApiInteractionSubmitPost) | **Post** /api/interaction/submit | 
[**ApiInteractionVerificationSocialAuthorizationUriPost**](InteractionAPI.md#ApiInteractionVerificationSocialAuthorizationUriPost) | **Post** /api/interaction/verification/social-authorization-uri | 
[**ApiInteractionVerificationTotpPost**](InteractionAPI.md#ApiInteractionVerificationTotpPost) | **Post** /api/interaction/verification/totp | 
[**ApiInteractionVerificationVerificationCodePost**](InteractionAPI.md#ApiInteractionVerificationVerificationCodePost) | **Post** /api/interaction/verification/verification-code | 
[**ApiInteractionVerificationWebauthnAuthenticationPost**](InteractionAPI.md#ApiInteractionVerificationWebauthnAuthenticationPost) | **Post** /api/interaction/verification/webauthn-authentication | 
[**ApiInteractionVerificationWebauthnRegistrationPost**](InteractionAPI.md#ApiInteractionVerificationWebauthnRegistrationPost) | **Post** /api/interaction/verification/webauthn-registration | 



## ApiInteractionBindMfaPost

> ApiInteractionBindMfaPost(ctx).ApiInteractionBindMfaPostRequest(apiInteractionBindMfaPostRequest).Execute()



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
	apiInteractionBindMfaPostRequest := openapiclient._api_interaction_bind_mfa_post_request{ApiInteractionBindMfaPostRequestOneOf: openapiclient.NewApiInteractionBindMfaPostRequestOneOf("Type_example", "Code_example")} // ApiInteractionBindMfaPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InteractionAPI.ApiInteractionBindMfaPost(context.Background()).ApiInteractionBindMfaPostRequest(apiInteractionBindMfaPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionBindMfaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionBindMfaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiInteractionBindMfaPostRequest** | [**ApiInteractionBindMfaPostRequest**](ApiInteractionBindMfaPostRequest.md) |  | 

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


## ApiInteractionConsentGet

> ApiInteractionConsentGet200Response ApiInteractionConsentGet(ctx).Execute()



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
	resp, r, err := apiClient.InteractionAPI.ApiInteractionConsentGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionConsentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiInteractionConsentGet`: ApiInteractionConsentGet200Response
	fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.ApiInteractionConsentGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionConsentGetRequest struct via the builder pattern


### Return type

[**ApiInteractionConsentGet200Response**](ApiInteractionConsentGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiInteractionConsentPost

> ApiInteractionConsentPost(ctx).ApiInteractionConsentPostRequest(apiInteractionConsentPostRequest).Execute()



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
	apiInteractionConsentPostRequest := *openapiclient.NewApiInteractionConsentPostRequest() // ApiInteractionConsentPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InteractionAPI.ApiInteractionConsentPost(context.Background()).ApiInteractionConsentPostRequest(apiInteractionConsentPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionConsentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionConsentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiInteractionConsentPostRequest** | [**ApiInteractionConsentPostRequest**](ApiInteractionConsentPostRequest.md) |  | 

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


## ApiInteractionDelete

> ApiInteractionDelete(ctx).Execute()



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
	r, err := apiClient.InteractionAPI.ApiInteractionDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionDeleteRequest struct via the builder pattern


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


## ApiInteractionEventPut

> ApiInteractionEventPut(ctx).ApiInteractionEventPutRequest(apiInteractionEventPutRequest).Execute()



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
	apiInteractionEventPutRequest := *openapiclient.NewApiInteractionEventPutRequest("Event_example") // ApiInteractionEventPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InteractionAPI.ApiInteractionEventPut(context.Background()).ApiInteractionEventPutRequest(apiInteractionEventPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionEventPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionEventPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiInteractionEventPutRequest** | [**ApiInteractionEventPutRequest**](ApiInteractionEventPutRequest.md) |  | 

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


## ApiInteractionIdentifiersPatch

> ApiInteractionIdentifiersPatch(ctx).ApiInteractionPutRequestIdentifier(apiInteractionPutRequestIdentifier).Execute()



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
	apiInteractionPutRequestIdentifier := openapiclient._api_interaction_put_request_identifier{ApiInteractionPutRequestIdentifierOneOf: openapiclient.NewApiInteractionPutRequestIdentifierOneOf("Username_example", "Password_example")} // ApiInteractionPutRequestIdentifier | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InteractionAPI.ApiInteractionIdentifiersPatch(context.Background()).ApiInteractionPutRequestIdentifier(apiInteractionPutRequestIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionIdentifiersPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionIdentifiersPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiInteractionPutRequestIdentifier** | [**ApiInteractionPutRequestIdentifier**](ApiInteractionPutRequestIdentifier.md) |  | 

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


## ApiInteractionMfaPut

> ApiInteractionMfaPut(ctx).ApiInteractionMfaPutRequest(apiInteractionMfaPutRequest).Execute()



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
	apiInteractionMfaPutRequest := openapiclient._api_interaction_mfa_put_request{ApiInteractionBindMfaPostRequestOneOf: openapiclient.NewApiInteractionBindMfaPostRequestOneOf("Type_example", "Code_example")} // ApiInteractionMfaPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InteractionAPI.ApiInteractionMfaPut(context.Background()).ApiInteractionMfaPutRequest(apiInteractionMfaPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionMfaPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionMfaPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiInteractionMfaPutRequest** | [**ApiInteractionMfaPutRequest**](ApiInteractionMfaPutRequest.md) |  | 

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


## ApiInteractionMfaSkippedPut

> ApiInteractionMfaSkippedPut(ctx).ApiInteractionMfaSkippedPutRequest(apiInteractionMfaSkippedPutRequest).Execute()



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
	apiInteractionMfaSkippedPutRequest := *openapiclient.NewApiInteractionMfaSkippedPutRequest(false) // ApiInteractionMfaSkippedPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InteractionAPI.ApiInteractionMfaSkippedPut(context.Background()).ApiInteractionMfaSkippedPutRequest(apiInteractionMfaSkippedPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionMfaSkippedPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionMfaSkippedPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiInteractionMfaSkippedPutRequest** | [**ApiInteractionMfaSkippedPutRequest**](ApiInteractionMfaSkippedPutRequest.md) |  | 

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


## ApiInteractionProfileDelete

> ApiInteractionProfileDelete(ctx).Execute()



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
	r, err := apiClient.InteractionAPI.ApiInteractionProfileDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionProfileDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionProfileDeleteRequest struct via the builder pattern


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


## ApiInteractionProfilePatch

> ApiInteractionProfilePatch(ctx).ApiInteractionPutRequestProfile(apiInteractionPutRequestProfile).Execute()



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
	apiInteractionPutRequestProfile := *openapiclient.NewApiInteractionPutRequestProfile() // ApiInteractionPutRequestProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InteractionAPI.ApiInteractionProfilePatch(context.Background()).ApiInteractionPutRequestProfile(apiInteractionPutRequestProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionProfilePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionProfilePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiInteractionPutRequestProfile** | [**ApiInteractionPutRequestProfile**](ApiInteractionPutRequestProfile.md) |  | 

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


## ApiInteractionProfilePut

> ApiInteractionProfilePut(ctx).ApiInteractionPutRequestProfile(apiInteractionPutRequestProfile).Execute()



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
	apiInteractionPutRequestProfile := *openapiclient.NewApiInteractionPutRequestProfile() // ApiInteractionPutRequestProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InteractionAPI.ApiInteractionProfilePut(context.Background()).ApiInteractionPutRequestProfile(apiInteractionPutRequestProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionProfilePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionProfilePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiInteractionPutRequestProfile** | [**ApiInteractionPutRequestProfile**](ApiInteractionPutRequestProfile.md) |  | 

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


## ApiInteractionPut

> ApiInteractionPut(ctx).ApiInteractionPutRequest(apiInteractionPutRequest).Execute()



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
	apiInteractionPutRequest := *openapiclient.NewApiInteractionPutRequest("Event_example") // ApiInteractionPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InteractionAPI.ApiInteractionPut(context.Background()).ApiInteractionPutRequest(apiInteractionPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiInteractionPutRequest** | [**ApiInteractionPutRequest**](ApiInteractionPutRequest.md) |  | 

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


## ApiInteractionSingleSignOnConnectorIdAuthenticationPost

> SubmitInteraction200Response ApiInteractionSingleSignOnConnectorIdAuthenticationPost(ctx, connectorId).RequestBody(requestBody).Execute()



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
	requestBody := map[string]interface{}{"key": interface{}({})} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InteractionAPI.ApiInteractionSingleSignOnConnectorIdAuthenticationPost(context.Background(), connectorId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionSingleSignOnConnectorIdAuthenticationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiInteractionSingleSignOnConnectorIdAuthenticationPost`: SubmitInteraction200Response
	fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.ApiInteractionSingleSignOnConnectorIdAuthenticationPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionSingleSignOnConnectorIdAuthenticationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**SubmitInteraction200Response**](SubmitInteraction200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPost

> SubmitInteraction200Response ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPost(ctx, connectorId).ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPostRequest(apiInteractionSingleSignOnConnectorIdAuthorizationUrlPostRequest).Execute()



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
	apiInteractionSingleSignOnConnectorIdAuthorizationUrlPostRequest := *openapiclient.NewApiInteractionSingleSignOnConnectorIdAuthorizationUrlPostRequest("State_example", map[string]interface{}(123)) // ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InteractionAPI.ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPost(context.Background(), connectorId).ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPostRequest(apiInteractionSingleSignOnConnectorIdAuthorizationUrlPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPost`: SubmitInteraction200Response
	fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionSingleSignOnConnectorIdAuthorizationUrlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiInteractionSingleSignOnConnectorIdAuthorizationUrlPostRequest** | [**ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPostRequest**](ApiInteractionSingleSignOnConnectorIdAuthorizationUrlPostRequest.md) |  | 

### Return type

[**SubmitInteraction200Response**](SubmitInteraction200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiInteractionSingleSignOnConnectorIdRegistrationPost

> SubmitInteraction200Response ApiInteractionSingleSignOnConnectorIdRegistrationPost(ctx, connectorId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InteractionAPI.ApiInteractionSingleSignOnConnectorIdRegistrationPost(context.Background(), connectorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionSingleSignOnConnectorIdRegistrationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiInteractionSingleSignOnConnectorIdRegistrationPost`: SubmitInteraction200Response
	fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.ApiInteractionSingleSignOnConnectorIdRegistrationPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorId** | **string** | The unique identifier of the connector. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionSingleSignOnConnectorIdRegistrationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ApiInteractionSingleSignOnConnectorsGet

> []string ApiInteractionSingleSignOnConnectorsGet(ctx).Email(email).Execute()



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
	email := "email_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InteractionAPI.ApiInteractionSingleSignOnConnectorsGet(context.Background()).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionSingleSignOnConnectorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiInteractionSingleSignOnConnectorsGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.ApiInteractionSingleSignOnConnectorsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionSingleSignOnConnectorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiInteractionSubmitPost

> SubmitInteraction200Response ApiInteractionSubmitPost(ctx).Execute()



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
	resp, r, err := apiClient.InteractionAPI.ApiInteractionSubmitPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiInteractionSubmitPost`: SubmitInteraction200Response
	fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.ApiInteractionSubmitPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionSubmitPostRequest struct via the builder pattern


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


## ApiInteractionVerificationSocialAuthorizationUriPost

> SubmitInteraction200Response ApiInteractionVerificationSocialAuthorizationUriPost(ctx).ApiInteractionVerificationSocialAuthorizationUriPostRequest(apiInteractionVerificationSocialAuthorizationUriPostRequest).Execute()



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
	apiInteractionVerificationSocialAuthorizationUriPostRequest := *openapiclient.NewApiInteractionVerificationSocialAuthorizationUriPostRequest("ConnectorId_example", "State_example", map[string]interface{}(123)) // ApiInteractionVerificationSocialAuthorizationUriPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InteractionAPI.ApiInteractionVerificationSocialAuthorizationUriPost(context.Background()).ApiInteractionVerificationSocialAuthorizationUriPostRequest(apiInteractionVerificationSocialAuthorizationUriPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionVerificationSocialAuthorizationUriPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiInteractionVerificationSocialAuthorizationUriPost`: SubmitInteraction200Response
	fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.ApiInteractionVerificationSocialAuthorizationUriPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionVerificationSocialAuthorizationUriPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiInteractionVerificationSocialAuthorizationUriPostRequest** | [**ApiInteractionVerificationSocialAuthorizationUriPostRequest**](ApiInteractionVerificationSocialAuthorizationUriPostRequest.md) |  | 

### Return type

[**SubmitInteraction200Response**](SubmitInteraction200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiInteractionVerificationTotpPost

> ApiInteractionVerificationTotpPost200Response ApiInteractionVerificationTotpPost(ctx).Execute()



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
	resp, r, err := apiClient.InteractionAPI.ApiInteractionVerificationTotpPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionVerificationTotpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiInteractionVerificationTotpPost`: ApiInteractionVerificationTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.ApiInteractionVerificationTotpPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionVerificationTotpPostRequest struct via the builder pattern


### Return type

[**ApiInteractionVerificationTotpPost200Response**](ApiInteractionVerificationTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiInteractionVerificationVerificationCodePost

> ApiInteractionVerificationVerificationCodePost(ctx).CreateVerificationCodeRequest(createVerificationCodeRequest).Execute()



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
	createVerificationCodeRequest := openapiclient.CreateVerificationCode_request{CreateVerificationCodeRequestOneOf: openapiclient.NewCreateVerificationCodeRequestOneOf("Email_example")} // CreateVerificationCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InteractionAPI.ApiInteractionVerificationVerificationCodePost(context.Background()).CreateVerificationCodeRequest(createVerificationCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionVerificationVerificationCodePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionVerificationVerificationCodePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVerificationCodeRequest** | [**CreateVerificationCodeRequest**](CreateVerificationCodeRequest.md) |  | 

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


## ApiInteractionVerificationWebauthnAuthenticationPost

> ApiInteractionVerificationWebauthnAuthenticationPost200Response ApiInteractionVerificationWebauthnAuthenticationPost(ctx).Execute()



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
	resp, r, err := apiClient.InteractionAPI.ApiInteractionVerificationWebauthnAuthenticationPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionVerificationWebauthnAuthenticationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiInteractionVerificationWebauthnAuthenticationPost`: ApiInteractionVerificationWebauthnAuthenticationPost200Response
	fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.ApiInteractionVerificationWebauthnAuthenticationPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionVerificationWebauthnAuthenticationPostRequest struct via the builder pattern


### Return type

[**ApiInteractionVerificationWebauthnAuthenticationPost200Response**](ApiInteractionVerificationWebauthnAuthenticationPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiInteractionVerificationWebauthnRegistrationPost

> ApiInteractionVerificationWebauthnRegistrationPost200Response ApiInteractionVerificationWebauthnRegistrationPost(ctx).Execute()



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
	resp, r, err := apiClient.InteractionAPI.ApiInteractionVerificationWebauthnRegistrationPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InteractionAPI.ApiInteractionVerificationWebauthnRegistrationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiInteractionVerificationWebauthnRegistrationPost`: ApiInteractionVerificationWebauthnRegistrationPost200Response
	fmt.Fprintf(os.Stdout, "Response from `InteractionAPI.ApiInteractionVerificationWebauthnRegistrationPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiInteractionVerificationWebauthnRegistrationPostRequest struct via the builder pattern


### Return type

[**ApiInteractionVerificationWebauthnRegistrationPost200Response**](ApiInteractionVerificationWebauthnRegistrationPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

