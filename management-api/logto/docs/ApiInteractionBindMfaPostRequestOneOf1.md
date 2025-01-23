# ApiInteractionBindMfaPostRequestOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**RawId** | **string** |  | 
**Response** | [**VerifyWebAuthnRegistrationVerificationRequestPayloadResponse**](VerifyWebAuthnRegistrationVerificationRequestPayloadResponse.md) |  | 
**AuthenticatorAttachment** | Pointer to **string** |  | [optional] 
**ClientExtensionResults** | [**VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults**](VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults.md) |  | 

## Methods

### NewApiInteractionBindMfaPostRequestOneOf1

`func NewApiInteractionBindMfaPostRequestOneOf1(type_ string, id string, rawId string, response VerifyWebAuthnRegistrationVerificationRequestPayloadResponse, clientExtensionResults VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults, ) *ApiInteractionBindMfaPostRequestOneOf1`

NewApiInteractionBindMfaPostRequestOneOf1 instantiates a new ApiInteractionBindMfaPostRequestOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInteractionBindMfaPostRequestOneOf1WithDefaults

`func NewApiInteractionBindMfaPostRequestOneOf1WithDefaults() *ApiInteractionBindMfaPostRequestOneOf1`

NewApiInteractionBindMfaPostRequestOneOf1WithDefaults instantiates a new ApiInteractionBindMfaPostRequestOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiInteractionBindMfaPostRequestOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiInteractionBindMfaPostRequestOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiInteractionBindMfaPostRequestOneOf1) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ApiInteractionBindMfaPostRequestOneOf1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiInteractionBindMfaPostRequestOneOf1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiInteractionBindMfaPostRequestOneOf1) SetId(v string)`

SetId sets Id field to given value.


### GetRawId

`func (o *ApiInteractionBindMfaPostRequestOneOf1) GetRawId() string`

GetRawId returns the RawId field if non-nil, zero value otherwise.

### GetRawIdOk

`func (o *ApiInteractionBindMfaPostRequestOneOf1) GetRawIdOk() (*string, bool)`

GetRawIdOk returns a tuple with the RawId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawId

`func (o *ApiInteractionBindMfaPostRequestOneOf1) SetRawId(v string)`

SetRawId sets RawId field to given value.


### GetResponse

`func (o *ApiInteractionBindMfaPostRequestOneOf1) GetResponse() VerifyWebAuthnRegistrationVerificationRequestPayloadResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ApiInteractionBindMfaPostRequestOneOf1) GetResponseOk() (*VerifyWebAuthnRegistrationVerificationRequestPayloadResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ApiInteractionBindMfaPostRequestOneOf1) SetResponse(v VerifyWebAuthnRegistrationVerificationRequestPayloadResponse)`

SetResponse sets Response field to given value.


### GetAuthenticatorAttachment

`func (o *ApiInteractionBindMfaPostRequestOneOf1) GetAuthenticatorAttachment() string`

GetAuthenticatorAttachment returns the AuthenticatorAttachment field if non-nil, zero value otherwise.

### GetAuthenticatorAttachmentOk

`func (o *ApiInteractionBindMfaPostRequestOneOf1) GetAuthenticatorAttachmentOk() (*string, bool)`

GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorAttachment

`func (o *ApiInteractionBindMfaPostRequestOneOf1) SetAuthenticatorAttachment(v string)`

SetAuthenticatorAttachment sets AuthenticatorAttachment field to given value.

### HasAuthenticatorAttachment

`func (o *ApiInteractionBindMfaPostRequestOneOf1) HasAuthenticatorAttachment() bool`

HasAuthenticatorAttachment returns a boolean if a field has been set.

### GetClientExtensionResults

`func (o *ApiInteractionBindMfaPostRequestOneOf1) GetClientExtensionResults() VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults`

GetClientExtensionResults returns the ClientExtensionResults field if non-nil, zero value otherwise.

### GetClientExtensionResultsOk

`func (o *ApiInteractionBindMfaPostRequestOneOf1) GetClientExtensionResultsOk() (*VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults, bool)`

GetClientExtensionResultsOk returns a tuple with the ClientExtensionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientExtensionResults

`func (o *ApiInteractionBindMfaPostRequestOneOf1) SetClientExtensionResults(v VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults)`

SetClientExtensionResults sets ClientExtensionResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


