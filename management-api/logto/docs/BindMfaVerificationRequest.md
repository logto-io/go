# BindMfaVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of MFA. | 
**VerificationId** | **string** | The ID of the MFA verification record. | 

## Methods

### NewBindMfaVerificationRequest

`func NewBindMfaVerificationRequest(type_ string, verificationId string, ) *BindMfaVerificationRequest`

NewBindMfaVerificationRequest instantiates a new BindMfaVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindMfaVerificationRequestWithDefaults

`func NewBindMfaVerificationRequestWithDefaults() *BindMfaVerificationRequest`

NewBindMfaVerificationRequestWithDefaults instantiates a new BindMfaVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BindMfaVerificationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BindMfaVerificationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BindMfaVerificationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetVerificationId

`func (o *BindMfaVerificationRequest) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *BindMfaVerificationRequest) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *BindMfaVerificationRequest) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


