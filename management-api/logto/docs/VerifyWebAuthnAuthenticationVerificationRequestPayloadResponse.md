# VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientDataJSON** | **string** |  | 
**AuthenticatorData** | **string** |  | 
**Signature** | **string** |  | 
**UserHandle** | Pointer to **string** |  | [optional] 

## Methods

### NewVerifyWebAuthnAuthenticationVerificationRequestPayloadResponse

`func NewVerifyWebAuthnAuthenticationVerificationRequestPayloadResponse(clientDataJSON string, authenticatorData string, signature string, ) *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse`

NewVerifyWebAuthnAuthenticationVerificationRequestPayloadResponse instantiates a new VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyWebAuthnAuthenticationVerificationRequestPayloadResponseWithDefaults

`func NewVerifyWebAuthnAuthenticationVerificationRequestPayloadResponseWithDefaults() *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse`

NewVerifyWebAuthnAuthenticationVerificationRequestPayloadResponseWithDefaults instantiates a new VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientDataJSON

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) GetClientDataJSON() string`

GetClientDataJSON returns the ClientDataJSON field if non-nil, zero value otherwise.

### GetClientDataJSONOk

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) GetClientDataJSONOk() (*string, bool)`

GetClientDataJSONOk returns a tuple with the ClientDataJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDataJSON

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) SetClientDataJSON(v string)`

SetClientDataJSON sets ClientDataJSON field to given value.


### GetAuthenticatorData

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) GetAuthenticatorData() string`

GetAuthenticatorData returns the AuthenticatorData field if non-nil, zero value otherwise.

### GetAuthenticatorDataOk

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) GetAuthenticatorDataOk() (*string, bool)`

GetAuthenticatorDataOk returns a tuple with the AuthenticatorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorData

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) SetAuthenticatorData(v string)`

SetAuthenticatorData sets AuthenticatorData field to given value.


### GetSignature

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetUserHandle

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) GetUserHandle() string`

GetUserHandle returns the UserHandle field if non-nil, zero value otherwise.

### GetUserHandleOk

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) GetUserHandleOk() (*string, bool)`

GetUserHandleOk returns a tuple with the UserHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHandle

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) SetUserHandle(v string)`

SetUserHandle sets UserHandle field to given value.

### HasUserHandle

`func (o *VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse) HasUserHandle() bool`

HasUserHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


