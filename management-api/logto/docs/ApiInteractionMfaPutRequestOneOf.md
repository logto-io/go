# ApiInteractionMfaPutRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**RawId** | **string** |  | 
**AuthenticatorAttachment** | Pointer to **string** |  | [optional] 
**ClientExtensionResults** | [**VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults**](VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults.md) |  | 
**Response** | [**VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse**](VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse.md) |  | 

## Methods

### NewApiInteractionMfaPutRequestOneOf

`func NewApiInteractionMfaPutRequestOneOf(type_ string, id string, rawId string, clientExtensionResults VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults, response VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse, ) *ApiInteractionMfaPutRequestOneOf`

NewApiInteractionMfaPutRequestOneOf instantiates a new ApiInteractionMfaPutRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInteractionMfaPutRequestOneOfWithDefaults

`func NewApiInteractionMfaPutRequestOneOfWithDefaults() *ApiInteractionMfaPutRequestOneOf`

NewApiInteractionMfaPutRequestOneOfWithDefaults instantiates a new ApiInteractionMfaPutRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiInteractionMfaPutRequestOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiInteractionMfaPutRequestOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiInteractionMfaPutRequestOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ApiInteractionMfaPutRequestOneOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiInteractionMfaPutRequestOneOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiInteractionMfaPutRequestOneOf) SetId(v string)`

SetId sets Id field to given value.


### GetRawId

`func (o *ApiInteractionMfaPutRequestOneOf) GetRawId() string`

GetRawId returns the RawId field if non-nil, zero value otherwise.

### GetRawIdOk

`func (o *ApiInteractionMfaPutRequestOneOf) GetRawIdOk() (*string, bool)`

GetRawIdOk returns a tuple with the RawId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawId

`func (o *ApiInteractionMfaPutRequestOneOf) SetRawId(v string)`

SetRawId sets RawId field to given value.


### GetAuthenticatorAttachment

`func (o *ApiInteractionMfaPutRequestOneOf) GetAuthenticatorAttachment() string`

GetAuthenticatorAttachment returns the AuthenticatorAttachment field if non-nil, zero value otherwise.

### GetAuthenticatorAttachmentOk

`func (o *ApiInteractionMfaPutRequestOneOf) GetAuthenticatorAttachmentOk() (*string, bool)`

GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorAttachment

`func (o *ApiInteractionMfaPutRequestOneOf) SetAuthenticatorAttachment(v string)`

SetAuthenticatorAttachment sets AuthenticatorAttachment field to given value.

### HasAuthenticatorAttachment

`func (o *ApiInteractionMfaPutRequestOneOf) HasAuthenticatorAttachment() bool`

HasAuthenticatorAttachment returns a boolean if a field has been set.

### GetClientExtensionResults

`func (o *ApiInteractionMfaPutRequestOneOf) GetClientExtensionResults() VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults`

GetClientExtensionResults returns the ClientExtensionResults field if non-nil, zero value otherwise.

### GetClientExtensionResultsOk

`func (o *ApiInteractionMfaPutRequestOneOf) GetClientExtensionResultsOk() (*VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults, bool)`

GetClientExtensionResultsOk returns a tuple with the ClientExtensionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientExtensionResults

`func (o *ApiInteractionMfaPutRequestOneOf) SetClientExtensionResults(v VerifyWebAuthnRegistrationVerificationRequestPayloadClientExtensionResults)`

SetClientExtensionResults sets ClientExtensionResults field to given value.


### GetResponse

`func (o *ApiInteractionMfaPutRequestOneOf) GetResponse() VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ApiInteractionMfaPutRequestOneOf) GetResponseOk() (*VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ApiInteractionMfaPutRequestOneOf) SetResponse(v VerifyWebAuthnAuthenticationVerificationRequestPayloadResponse)`

SetResponse sets Response field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


