# CheckPasswordWithDefaultSignInExperienceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | The password to check. | 
**UserId** | Pointer to **string** | The user ID to check the password for. It is required if rejects user info is enabled in the password policy. | [optional] 

## Methods

### NewCheckPasswordWithDefaultSignInExperienceRequest

`func NewCheckPasswordWithDefaultSignInExperienceRequest(password string, ) *CheckPasswordWithDefaultSignInExperienceRequest`

NewCheckPasswordWithDefaultSignInExperienceRequest instantiates a new CheckPasswordWithDefaultSignInExperienceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckPasswordWithDefaultSignInExperienceRequestWithDefaults

`func NewCheckPasswordWithDefaultSignInExperienceRequestWithDefaults() *CheckPasswordWithDefaultSignInExperienceRequest`

NewCheckPasswordWithDefaultSignInExperienceRequestWithDefaults instantiates a new CheckPasswordWithDefaultSignInExperienceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *CheckPasswordWithDefaultSignInExperienceRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CheckPasswordWithDefaultSignInExperienceRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CheckPasswordWithDefaultSignInExperienceRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUserId

`func (o *CheckPasswordWithDefaultSignInExperienceRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CheckPasswordWithDefaultSignInExperienceRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CheckPasswordWithDefaultSignInExperienceRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CheckPasswordWithDefaultSignInExperienceRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


