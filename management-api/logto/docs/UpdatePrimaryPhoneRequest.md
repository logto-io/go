# UpdatePrimaryPhoneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** | The new phone for the user. | 
**NewIdentifierVerificationRecordId** | **string** | The identifier verification record ID for the new phone ownership verification. | 

## Methods

### NewUpdatePrimaryPhoneRequest

`func NewUpdatePrimaryPhoneRequest(phone string, newIdentifierVerificationRecordId string, ) *UpdatePrimaryPhoneRequest`

NewUpdatePrimaryPhoneRequest instantiates a new UpdatePrimaryPhoneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePrimaryPhoneRequestWithDefaults

`func NewUpdatePrimaryPhoneRequestWithDefaults() *UpdatePrimaryPhoneRequest`

NewUpdatePrimaryPhoneRequestWithDefaults instantiates a new UpdatePrimaryPhoneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *UpdatePrimaryPhoneRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdatePrimaryPhoneRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdatePrimaryPhoneRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetNewIdentifierVerificationRecordId

`func (o *UpdatePrimaryPhoneRequest) GetNewIdentifierVerificationRecordId() string`

GetNewIdentifierVerificationRecordId returns the NewIdentifierVerificationRecordId field if non-nil, zero value otherwise.

### GetNewIdentifierVerificationRecordIdOk

`func (o *UpdatePrimaryPhoneRequest) GetNewIdentifierVerificationRecordIdOk() (*string, bool)`

GetNewIdentifierVerificationRecordIdOk returns a tuple with the NewIdentifierVerificationRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIdentifierVerificationRecordId

`func (o *UpdatePrimaryPhoneRequest) SetNewIdentifierVerificationRecordId(v string)`

SetNewIdentifierVerificationRecordId sets NewIdentifierVerificationRecordId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


