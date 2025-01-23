# ApiInteractionBindMfaPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Code** | **string** |  | 
**Id** | **string** |  | 
**RawId** | **string** |  | 
**Response** | [**VerifyWebAuthnRegistrationVerificationRequestPayloadResponse**](VerifyWebAuthnRegistrationVerificationRequestPayloadResponse.md) |  | 
**AuthenticatorAttachment** | Pointer to **string** |  | [optional] 
**ClientExtensionResults** | [**VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults**](VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults.md) |  | 

## Methods

### NewApiInteractionBindMfaPostRequest

`func NewApiInteractionBindMfaPostRequest(type_ string, code string, id string, rawId string, response VerifyWebAuthnRegistrationVerificationRequestPayloadResponse, clientExtensionResults VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults, ) *ApiInteractionBindMfaPostRequest`

NewApiInteractionBindMfaPostRequest instantiates a new ApiInteractionBindMfaPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInteractionBindMfaPostRequestWithDefaults

`func NewApiInteractionBindMfaPostRequestWithDefaults() *ApiInteractionBindMfaPostRequest`

NewApiInteractionBindMfaPostRequestWithDefaults instantiates a new ApiInteractionBindMfaPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiInteractionBindMfaPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiInteractionBindMfaPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiInteractionBindMfaPostRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *ApiInteractionBindMfaPostRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiInteractionBindMfaPostRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiInteractionBindMfaPostRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetId

`func (o *ApiInteractionBindMfaPostRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiInteractionBindMfaPostRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiInteractionBindMfaPostRequest) SetId(v string)`

SetId sets Id field to given value.


### GetRawId

`func (o *ApiInteractionBindMfaPostRequest) GetRawId() string`

GetRawId returns the RawId field if non-nil, zero value otherwise.

### GetRawIdOk

`func (o *ApiInteractionBindMfaPostRequest) GetRawIdOk() (*string, bool)`

GetRawIdOk returns a tuple with the RawId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawId

`func (o *ApiInteractionBindMfaPostRequest) SetRawId(v string)`

SetRawId sets RawId field to given value.


### GetResponse

`func (o *ApiInteractionBindMfaPostRequest) GetResponse() VerifyWebAuthnRegistrationVerificationRequestPayloadResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ApiInteractionBindMfaPostRequest) GetResponseOk() (*VerifyWebAuthnRegistrationVerificationRequestPayloadResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ApiInteractionBindMfaPostRequest) SetResponse(v VerifyWebAuthnRegistrationVerificationRequestPayloadResponse)`

SetResponse sets Response field to given value.


### GetAuthenticatorAttachment

`func (o *ApiInteractionBindMfaPostRequest) GetAuthenticatorAttachment() string`

GetAuthenticatorAttachment returns the AuthenticatorAttachment field if non-nil, zero value otherwise.

### GetAuthenticatorAttachmentOk

`func (o *ApiInteractionBindMfaPostRequest) GetAuthenticatorAttachmentOk() (*string, bool)`

GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorAttachment

`func (o *ApiInteractionBindMfaPostRequest) SetAuthenticatorAttachment(v string)`

SetAuthenticatorAttachment sets AuthenticatorAttachment field to given value.

### HasAuthenticatorAttachment

`func (o *ApiInteractionBindMfaPostRequest) HasAuthenticatorAttachment() bool`

HasAuthenticatorAttachment returns a boolean if a field has been set.

### GetClientExtensionResults

`func (o *ApiInteractionBindMfaPostRequest) GetClientExtensionResults() VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults`

GetClientExtensionResults returns the ClientExtensionResults field if non-nil, zero value otherwise.

### GetClientExtensionResultsOk

`func (o *ApiInteractionBindMfaPostRequest) GetClientExtensionResultsOk() (*VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults, bool)`

GetClientExtensionResultsOk returns a tuple with the ClientExtensionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientExtensionResults

`func (o *ApiInteractionBindMfaPostRequest) SetClientExtensionResults(v VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults)`

SetClientExtensionResults sets ClientExtensionResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


