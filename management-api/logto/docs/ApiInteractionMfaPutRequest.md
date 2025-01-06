# ApiInteractionMfaPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Code** | **string** |  | 
**Id** | **string** |  | 
**RawId** | **string** |  | 
**AuthenticatorAttachment** | Pointer to **string** |  | [optional] 
**ClientExtensionResults** | [**VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults**](VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults.md) |  | 
**Response** | [**VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse**](VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse.md) |  | 

## Methods

### NewApiInteractionMfaPutRequest

`func NewApiInteractionMfaPutRequest(type_ string, code string, id string, rawId string, clientExtensionResults VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults, response VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse, ) *ApiInteractionMfaPutRequest`

NewApiInteractionMfaPutRequest instantiates a new ApiInteractionMfaPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInteractionMfaPutRequestWithDefaults

`func NewApiInteractionMfaPutRequestWithDefaults() *ApiInteractionMfaPutRequest`

NewApiInteractionMfaPutRequestWithDefaults instantiates a new ApiInteractionMfaPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiInteractionMfaPutRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiInteractionMfaPutRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiInteractionMfaPutRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *ApiInteractionMfaPutRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiInteractionMfaPutRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiInteractionMfaPutRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetId

`func (o *ApiInteractionMfaPutRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiInteractionMfaPutRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiInteractionMfaPutRequest) SetId(v string)`

SetId sets Id field to given value.


### GetRawId

`func (o *ApiInteractionMfaPutRequest) GetRawId() string`

GetRawId returns the RawId field if non-nil, zero value otherwise.

### GetRawIdOk

`func (o *ApiInteractionMfaPutRequest) GetRawIdOk() (*string, bool)`

GetRawIdOk returns a tuple with the RawId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawId

`func (o *ApiInteractionMfaPutRequest) SetRawId(v string)`

SetRawId sets RawId field to given value.


### GetAuthenticatorAttachment

`func (o *ApiInteractionMfaPutRequest) GetAuthenticatorAttachment() string`

GetAuthenticatorAttachment returns the AuthenticatorAttachment field if non-nil, zero value otherwise.

### GetAuthenticatorAttachmentOk

`func (o *ApiInteractionMfaPutRequest) GetAuthenticatorAttachmentOk() (*string, bool)`

GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorAttachment

`func (o *ApiInteractionMfaPutRequest) SetAuthenticatorAttachment(v string)`

SetAuthenticatorAttachment sets AuthenticatorAttachment field to given value.

### HasAuthenticatorAttachment

`func (o *ApiInteractionMfaPutRequest) HasAuthenticatorAttachment() bool`

HasAuthenticatorAttachment returns a boolean if a field has been set.

### GetClientExtensionResults

`func (o *ApiInteractionMfaPutRequest) GetClientExtensionResults() VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults`

GetClientExtensionResults returns the ClientExtensionResults field if non-nil, zero value otherwise.

### GetClientExtensionResultsOk

`func (o *ApiInteractionMfaPutRequest) GetClientExtensionResultsOk() (*VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults, bool)`

GetClientExtensionResultsOk returns a tuple with the ClientExtensionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientExtensionResults

`func (o *ApiInteractionMfaPutRequest) SetClientExtensionResults(v VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults)`

SetClientExtensionResults sets ClientExtensionResults field to given value.


### GetResponse

`func (o *ApiInteractionMfaPutRequest) GetResponse() VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ApiInteractionMfaPutRequest) GetResponseOk() (*VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ApiInteractionMfaPutRequest) SetResponse(v VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse)`

SetResponse sets Response field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


