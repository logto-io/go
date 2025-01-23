# AddUserProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **interface{}** | The type of profile data to add. &#x60;email&#x60;, &#x60;phone&#x60;, &#x60;username&#x60;, &#x60;password&#x60;, etc. | [optional] 
**Value** | Pointer to **interface{}** | The plain text value of the profile data. Only supported for profile data types that does not require verification, such as &#x60;username&#x60; and &#x60;password&#x60;. | [optional] 
**VerificationId** | Pointer to **interface{}** | The ID of the verification record used to verify the profile data. Required for profile data types that require verification, such as &#x60;email&#x60;, &#x60;phone&#x60; and &#x60;social&#x60;. | [optional] 

## Methods

### NewAddUserProfileRequest

`func NewAddUserProfileRequest() *AddUserProfileRequest`

NewAddUserProfileRequest instantiates a new AddUserProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserProfileRequestWithDefaults

`func NewAddUserProfileRequestWithDefaults() *AddUserProfileRequest`

NewAddUserProfileRequestWithDefaults instantiates a new AddUserProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddUserProfileRequest) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddUserProfileRequest) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddUserProfileRequest) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *AddUserProfileRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AddUserProfileRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AddUserProfileRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValue

`func (o *AddUserProfileRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddUserProfileRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddUserProfileRequest) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *AddUserProfileRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *AddUserProfileRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AddUserProfileRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetVerificationId

`func (o *AddUserProfileRequest) GetVerificationId() interface{}`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *AddUserProfileRequest) GetVerificationIdOk() (*interface{}, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *AddUserProfileRequest) SetVerificationId(v interface{})`

SetVerificationId sets VerificationId field to given value.

### HasVerificationId

`func (o *AddUserProfileRequest) HasVerificationId() bool`

HasVerificationId returns a boolean if a field has been set.

### SetVerificationIdNil

`func (o *AddUserProfileRequest) SetVerificationIdNil(b bool)`

 SetVerificationIdNil sets the value for VerificationId to be an explicit nil

### UnsetVerificationId
`func (o *AddUserProfileRequest) UnsetVerificationId()`

UnsetVerificationId ensures that no value is present for VerificationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


