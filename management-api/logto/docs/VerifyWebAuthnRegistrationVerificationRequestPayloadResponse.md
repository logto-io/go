# VerifyWebAuthnRegistrationVerificationRequestPayloadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientDataJSON** | **string** |  | 
**AttestationObject** | **string** |  | 
**AuthenticatorData** | Pointer to **string** |  | [optional] 
**Transports** | Pointer to **[]string** |  | [optional] 
**PublicKeyAlgorithm** | Pointer to **float32** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 

## Methods

### NewVerifyWebAuthnRegistrationVerificationRequestPayloadResponse

`func NewVerifyWebAuthnRegistrationVerificationRequestPayloadResponse(clientDataJSON string, attestationObject string, ) *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse`

NewVerifyWebAuthnRegistrationVerificationRequestPayloadResponse instantiates a new VerifyWebAuthnRegistrationVerificationRequestPayloadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyWebAuthnRegistrationVerificationRequestPayloadResponseWithDefaults

`func NewVerifyWebAuthnRegistrationVerificationRequestPayloadResponseWithDefaults() *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse`

NewVerifyWebAuthnRegistrationVerificationRequestPayloadResponseWithDefaults instantiates a new VerifyWebAuthnRegistrationVerificationRequestPayloadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientDataJSON

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) GetClientDataJSON() string`

GetClientDataJSON returns the ClientDataJSON field if non-nil, zero value otherwise.

### GetClientDataJSONOk

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) GetClientDataJSONOk() (*string, bool)`

GetClientDataJSONOk returns a tuple with the ClientDataJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDataJSON

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) SetClientDataJSON(v string)`

SetClientDataJSON sets ClientDataJSON field to given value.


### GetAttestationObject

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) GetAttestationObject() string`

GetAttestationObject returns the AttestationObject field if non-nil, zero value otherwise.

### GetAttestationObjectOk

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) GetAttestationObjectOk() (*string, bool)`

GetAttestationObjectOk returns a tuple with the AttestationObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationObject

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) SetAttestationObject(v string)`

SetAttestationObject sets AttestationObject field to given value.


### GetAuthenticatorData

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) GetAuthenticatorData() string`

GetAuthenticatorData returns the AuthenticatorData field if non-nil, zero value otherwise.

### GetAuthenticatorDataOk

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) GetAuthenticatorDataOk() (*string, bool)`

GetAuthenticatorDataOk returns a tuple with the AuthenticatorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorData

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) SetAuthenticatorData(v string)`

SetAuthenticatorData sets AuthenticatorData field to given value.

### HasAuthenticatorData

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) HasAuthenticatorData() bool`

HasAuthenticatorData returns a boolean if a field has been set.

### GetTransports

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) GetTransports() []string`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) GetTransportsOk() (*[]string, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) SetTransports(v []string)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### GetPublicKeyAlgorithm

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) GetPublicKeyAlgorithm() float32`

GetPublicKeyAlgorithm returns the PublicKeyAlgorithm field if non-nil, zero value otherwise.

### GetPublicKeyAlgorithmOk

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) GetPublicKeyAlgorithmOk() (*float32, bool)`

GetPublicKeyAlgorithmOk returns a tuple with the PublicKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyAlgorithm

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) SetPublicKeyAlgorithm(v float32)`

SetPublicKeyAlgorithm sets PublicKeyAlgorithm field to given value.

### HasPublicKeyAlgorithm

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) HasPublicKeyAlgorithm() bool`

HasPublicKeyAlgorithm returns a boolean if a field has been set.

### GetPublicKey

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *VerifyWebAuthnRegistrationVerificationRequestPayloadResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


