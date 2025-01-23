# UpdateUserPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | New password for the user. | 

## Methods

### NewUpdateUserPasswordRequest

`func NewUpdateUserPasswordRequest(password string, ) *UpdateUserPasswordRequest`

NewUpdateUserPasswordRequest instantiates a new UpdateUserPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserPasswordRequestWithDefaults

`func NewUpdateUserPasswordRequestWithDefaults() *UpdateUserPasswordRequest`

NewUpdateUserPasswordRequestWithDefaults instantiates a new UpdateUserPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UpdateUserPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateUserPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateUserPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


