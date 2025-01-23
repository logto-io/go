# CreateWebAuthnRegistrationVerification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | **string** | The unique verification ID for the WebAuthn registration record. This ID is required to verify the WebAuthn registration challenge. | 
**RegistrationOptions** | [**CreateWebAuthnRegistrationVerification200ResponseRegistrationOptions**](CreateWebAuthnRegistrationVerification200ResponseRegistrationOptions.md) |  | 

## Methods

### NewCreateWebAuthnRegistrationVerification200Response

`func NewCreateWebAuthnRegistrationVerification200Response(verificationId string, registrationOptions CreateWebAuthnRegistrationVerification200ResponseRegistrationOptions, ) *CreateWebAuthnRegistrationVerification200Response`

NewCreateWebAuthnRegistrationVerification200Response instantiates a new CreateWebAuthnRegistrationVerification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebAuthnRegistrationVerification200ResponseWithDefaults

`func NewCreateWebAuthnRegistrationVerification200ResponseWithDefaults() *CreateWebAuthnRegistrationVerification200Response`

NewCreateWebAuthnRegistrationVerification200ResponseWithDefaults instantiates a new CreateWebAuthnRegistrationVerification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *CreateWebAuthnRegistrationVerification200Response) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *CreateWebAuthnRegistrationVerification200Response) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *CreateWebAuthnRegistrationVerification200Response) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.


### GetRegistrationOptions

`func (o *CreateWebAuthnRegistrationVerification200Response) GetRegistrationOptions() CreateWebAuthnRegistrationVerification200ResponseRegistrationOptions`

GetRegistrationOptions returns the RegistrationOptions field if non-nil, zero value otherwise.

### GetRegistrationOptionsOk

`func (o *CreateWebAuthnRegistrationVerification200Response) GetRegistrationOptionsOk() (*CreateWebAuthnRegistrationVerification200ResponseRegistrationOptions, bool)`

GetRegistrationOptionsOk returns a tuple with the RegistrationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationOptions

`func (o *CreateWebAuthnRegistrationVerification200Response) SetRegistrationOptions(v CreateWebAuthnRegistrationVerification200ResponseRegistrationOptions)`

SetRegistrationOptions sets RegistrationOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


