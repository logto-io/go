# UpdatePrimaryEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The new email for the user. | 
**NewIdentifierVerificationRecordId** | **string** | The identifier verification record ID for the new email ownership verification. | 

## Methods

### NewUpdatePrimaryEmailRequest

`func NewUpdatePrimaryEmailRequest(email string, newIdentifierVerificationRecordId string, ) *UpdatePrimaryEmailRequest`

NewUpdatePrimaryEmailRequest instantiates a new UpdatePrimaryEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePrimaryEmailRequestWithDefaults

`func NewUpdatePrimaryEmailRequestWithDefaults() *UpdatePrimaryEmailRequest`

NewUpdatePrimaryEmailRequestWithDefaults instantiates a new UpdatePrimaryEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdatePrimaryEmailRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdatePrimaryEmailRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdatePrimaryEmailRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetNewIdentifierVerificationRecordId

`func (o *UpdatePrimaryEmailRequest) GetNewIdentifierVerificationRecordId() string`

GetNewIdentifierVerificationRecordId returns the NewIdentifierVerificationRecordId field if non-nil, zero value otherwise.

### GetNewIdentifierVerificationRecordIdOk

`func (o *UpdatePrimaryEmailRequest) GetNewIdentifierVerificationRecordIdOk() (*string, bool)`

GetNewIdentifierVerificationRecordIdOk returns a tuple with the NewIdentifierVerificationRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIdentifierVerificationRecordId

`func (o *UpdatePrimaryEmailRequest) SetNewIdentifierVerificationRecordId(v string)`

SetNewIdentifierVerificationRecordId sets NewIdentifierVerificationRecordId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


