# VerifyWebAuthnAuthenticationVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | **string** | The verification ID of the WebAuthn authentication verification record. | 
**Payload** | [**VerifyWebAuthnAuthenticationVerificationRequestPayload**](VerifyWebAuthnAuthenticationVerificationRequestPayload.md) |  | 

## Methods

### NewVerifyWebAuthnAuthenticationVerificationRequest

`func NewVerifyWebAuthnAuthenticationVerificationRequest(verificationId string, payload VerifyWebAuthnAuthenticationVerificationRequestPayload, ) *VerifyWebAuthnAuthenticationVerificationRequest`

NewVerifyWebAuthnAuthenticationVerificationRequest instantiates a new VerifyWebAuthnAuthenticationVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyWebAuthnAuthenticationVerificationRequestWithDefaults

`func NewVerifyWebAuthnAuthenticationVerificationRequestWithDefaults() *VerifyWebAuthnAuthenticationVerificationRequest`

NewVerifyWebAuthnAuthenticationVerificationRequestWithDefaults instantiates a new VerifyWebAuthnAuthenticationVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *VerifyWebAuthnAuthenticationVerificationRequest) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *VerifyWebAuthnAuthenticationVerificationRequest) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *VerifyWebAuthnAuthenticationVerificationRequest) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.


### GetPayload

`func (o *VerifyWebAuthnAuthenticationVerificationRequest) GetPayload() VerifyWebAuthnAuthenticationVerificationRequestPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *VerifyWebAuthnAuthenticationVerificationRequest) GetPayloadOk() (*VerifyWebAuthnAuthenticationVerificationRequestPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *VerifyWebAuthnAuthenticationVerificationRequest) SetPayload(v VerifyWebAuthnAuthenticationVerificationRequestPayload)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


