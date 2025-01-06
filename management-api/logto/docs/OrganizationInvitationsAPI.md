# \OrganizationInvitationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationInvitation**](OrganizationInvitationsAPI.md#CreateOrganizationInvitation) | **Post** /api/organization-invitations | Create organization invitation
[**CreateOrganizationInvitationMessage**](OrganizationInvitationsAPI.md#CreateOrganizationInvitationMessage) | **Post** /api/organization-invitations/{id}/message | Resend invitation message
[**DeleteOrganizationInvitation**](OrganizationInvitationsAPI.md#DeleteOrganizationInvitation) | **Delete** /api/organization-invitations/{id} | Delete organization invitation
[**GetOrganizationInvitation**](OrganizationInvitationsAPI.md#GetOrganizationInvitation) | **Get** /api/organization-invitations/{id} | Get organization invitation
[**ListOrganizationInvitations**](OrganizationInvitationsAPI.md#ListOrganizationInvitations) | **Get** /api/organization-invitations | Get organization invitations
[**ReplaceOrganizationInvitationStatus**](OrganizationInvitationsAPI.md#ReplaceOrganizationInvitationStatus) | **Put** /api/organization-invitations/{id}/status | Update organization invitation status



## CreateOrganizationInvitation

> GetOrganizationInvitation200Response CreateOrganizationInvitation(ctx).CreateOrganizationInvitationRequest(createOrganizationInvitationRequest).Execute()

Create organization invitation



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
	createOrganizationInvitationRequest := *openapiclient.NewCreateOrganizationInvitationRequest("Invitee_example", "OrganizationId_example", float32(123), openapiclient.CreateOrganizationInvitation_request_messagePayload{CreateOrganizationInvitationRequestMessagePayloadOneOf: openapiclient.NewCreateOrganizationInvitationRequestMessagePayloadOneOf()}) // CreateOrganizationInvitationRequest | The organization invitation to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationInvitationsAPI.CreateOrganizationInvitation(context.Background()).CreateOrganizationInvitationRequest(createOrganizationInvitationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsAPI.CreateOrganizationInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationInvitation`: GetOrganizationInvitation200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationInvitationsAPI.CreateOrganizationInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganizationInvitationRequest** | [**CreateOrganizationInvitationRequest**](CreateOrganizationInvitationRequest.md) | The organization invitation to create. | 

### Return type

[**GetOrganizationInvitation200Response**](GetOrganizationInvitation200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInvitationMessage

> CreateOrganizationInvitationMessage(ctx, id).CreateOrganizationInvitationRequestMessagePayloadOneOf(createOrganizationInvitationRequestMessagePayloadOneOf).Execute()

Resend invitation message



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
	id := "id_example" // string | The unique identifier of the organization invitation.
	createOrganizationInvitationRequestMessagePayloadOneOf := *openapiclient.NewCreateOrganizationInvitationRequestMessagePayloadOneOf() // CreateOrganizationInvitationRequestMessagePayloadOneOf | The message payload for the \"OrganizationInvitation\" template to use when sending the invitation via email.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationInvitationsAPI.CreateOrganizationInvitationMessage(context.Background(), id).CreateOrganizationInvitationRequestMessagePayloadOneOf(createOrganizationInvitationRequestMessagePayloadOneOf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsAPI.CreateOrganizationInvitationMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInvitationMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInvitationRequestMessagePayloadOneOf** | [**CreateOrganizationInvitationRequestMessagePayloadOneOf**](CreateOrganizationInvitationRequestMessagePayloadOneOf.md) | The message payload for the \&quot;OrganizationInvitation\&quot; template to use when sending the invitation via email. | 

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


## DeleteOrganizationInvitation

> DeleteOrganizationInvitation(ctx, id).Execute()

Delete organization invitation



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
	id := "id_example" // string | The unique identifier of the organization invitation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationInvitationsAPI.DeleteOrganizationInvitation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsAPI.DeleteOrganizationInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationInvitationRequest struct via the builder pattern


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


## GetOrganizationInvitation

> GetOrganizationInvitation200Response GetOrganizationInvitation(ctx, id).Execute()

Get organization invitation



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
	id := "id_example" // string | The unique identifier of the organization invitation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationInvitationsAPI.GetOrganizationInvitation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsAPI.GetOrganizationInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInvitation`: GetOrganizationInvitation200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationInvitationsAPI.GetOrganizationInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationInvitation200Response**](GetOrganizationInvitation200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationInvitations

> []GetOrganizationInvitation200Response ListOrganizationInvitations(ctx).OrganizationId(organizationId).InviterId(inviterId).Invitee(invitee).Execute()

Get organization invitations



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
	organizationId := "organizationId_example" // string |  (optional)
	inviterId := "inviterId_example" // string |  (optional)
	invitee := "invitee_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationInvitationsAPI.ListOrganizationInvitations(context.Background()).OrganizationId(organizationId).InviterId(inviterId).Invitee(invitee).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsAPI.ListOrganizationInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationInvitations`: []GetOrganizationInvitation200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationInvitationsAPI.ListOrganizationInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** |  | 
 **inviterId** | **string** |  | 
 **invitee** | **string** |  | 

### Return type

[**[]GetOrganizationInvitation200Response**](GetOrganizationInvitation200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOrganizationInvitationStatus

> GetOrganizationInvitation200Response ReplaceOrganizationInvitationStatus(ctx, id).ReplaceOrganizationInvitationStatusRequest(replaceOrganizationInvitationStatusRequest).Execute()

Update organization invitation status



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
	id := "id_example" // string | The unique identifier of the organization invitation.
	replaceOrganizationInvitationStatusRequest := *openapiclient.NewReplaceOrganizationInvitationStatusRequest("Status_example") // ReplaceOrganizationInvitationStatusRequest | The organization invitation status to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationInvitationsAPI.ReplaceOrganizationInvitationStatus(context.Background(), id).ReplaceOrganizationInvitationStatusRequest(replaceOrganizationInvitationStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationInvitationsAPI.ReplaceOrganizationInvitationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceOrganizationInvitationStatus`: GetOrganizationInvitation200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationInvitationsAPI.ReplaceOrganizationInvitationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the organization invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrganizationInvitationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceOrganizationInvitationStatusRequest** | [**ReplaceOrganizationInvitationStatusRequest**](ReplaceOrganizationInvitationStatusRequest.md) | The organization invitation status to update. | 

### Return type

[**GetOrganizationInvitation200Response**](GetOrganizationInvitation200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

