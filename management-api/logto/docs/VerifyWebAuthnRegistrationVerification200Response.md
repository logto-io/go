# VerifyWebAuthnRegistrationVerification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | **string** | The unique verification ID of the WebAuthn registration record. This &#x60;verificationId&#x60; is required to bind the WebAuthn credential to the user account via the &#x60;Profile&#x60; API. | 

## Methods

### NewVerifyWebAuthnRegistrationVerification200Response

`func NewVerifyWebAuthnRegistrationVerification200Response(verificationId string, ) *VerifyWebAuthnRegistrationVerification200Response`

NewVerifyWebAuthnRegistrationVerification200Response instantiates a new VerifyWebAuthnRegistrationVerification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyWebAuthnRegistrationVerification200ResponseWithDefaults

`func NewVerifyWebAuthnRegistrationVerification200ResponseWithDefaults() *VerifyWebAuthnRegistrationVerification200Response`

NewVerifyWebAuthnRegistrationVerification200ResponseWithDefaults instantiates a new VerifyWebAuthnRegistrationVerification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *VerifyWebAuthnRegistrationVerification200Response) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *VerifyWebAuthnRegistrationVerification200Response) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *VerifyWebAuthnRegistrationVerification200Response) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


