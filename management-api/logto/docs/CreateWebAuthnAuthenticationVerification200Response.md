# CreateWebAuthnAuthenticationVerification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | **string** | The unique ID for the WebAuthn authentication record, required to verify the WebAuthn authentication challenge. | 
**AuthenticationOptions** | [**CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions**](CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions.md) |  | 

## Methods

### NewCreateWebAuthnAuthenticationVerification200Response

`func NewCreateWebAuthnAuthenticationVerification200Response(verificationId string, authenticationOptions CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions, ) *CreateWebAuthnAuthenticationVerification200Response`

NewCreateWebAuthnAuthenticationVerification200Response instantiates a new CreateWebAuthnAuthenticationVerification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebAuthnAuthenticationVerification200ResponseWithDefaults

`func NewCreateWebAuthnAuthenticationVerification200ResponseWithDefaults() *CreateWebAuthnAuthenticationVerification200Response`

NewCreateWebAuthnAuthenticationVerification200ResponseWithDefaults instantiates a new CreateWebAuthnAuthenticationVerification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *CreateWebAuthnAuthenticationVerification200Response) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *CreateWebAuthnAuthenticationVerification200Response) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *CreateWebAuthnAuthenticationVerification200Response) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.


### GetAuthenticationOptions

`func (o *CreateWebAuthnAuthenticationVerification200Response) GetAuthenticationOptions() CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions`

GetAuthenticationOptions returns the AuthenticationOptions field if non-nil, zero value otherwise.

### GetAuthenticationOptionsOk

`func (o *CreateWebAuthnAuthenticationVerification200Response) GetAuthenticationOptionsOk() (*CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions, bool)`

GetAuthenticationOptionsOk returns a tuple with the AuthenticationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOptions

`func (o *CreateWebAuthnAuthenticationVerification200Response) SetAuthenticationOptions(v CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions)`

SetAuthenticationOptions sets AuthenticationOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


