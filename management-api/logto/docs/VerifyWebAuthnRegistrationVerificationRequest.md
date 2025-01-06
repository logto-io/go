# VerifyWebAuthnRegistrationVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | **string** | The verification ID of the WebAuthn registration record. | 
**Payload** | [**VerifyWebAuthnRegistrationVerificationRequestPayload**](VerifyWebAuthnRegistrationVerificationRequestPayload.md) |  | 

## Methods

### NewVerifyWebAuthnRegistrationVerificationRequest

`func NewVerifyWebAuthnRegistrationVerificationRequest(verificationId string, payload VerifyWebAuthnRegistrationVerificationRequestPayload, ) *VerifyWebAuthnRegistrationVerificationRequest`

NewVerifyWebAuthnRegistrationVerificationRequest instantiates a new VerifyWebAuthnRegistrationVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyWebAuthnRegistrationVerificationRequestWithDefaults

`func NewVerifyWebAuthnRegistrationVerificationRequestWithDefaults() *VerifyWebAuthnRegistrationVerificationRequest`

NewVerifyWebAuthnRegistrationVerificationRequestWithDefaults instantiates a new VerifyWebAuthnRegistrationVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *VerifyWebAuthnRegistrationVerificationRequest) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *VerifyWebAuthnRegistrationVerificationRequest) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *VerifyWebAuthnRegistrationVerificationRequest) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.


### GetPayload

`func (o *VerifyWebAuthnRegistrationVerificationRequest) GetPayload() VerifyWebAuthnRegistrationVerificationRequestPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *VerifyWebAuthnRegistrationVerificationRequest) GetPayloadOk() (*VerifyWebAuthnRegistrationVerificationRequestPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *VerifyWebAuthnRegistrationVerificationRequest) SetPayload(v VerifyWebAuthnRegistrationVerificationRequestPayload)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


