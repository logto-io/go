# UpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to [**NullableUpdateUserRequestUsername**](UpdateUserRequestUsername.md) |  | [optional] 
**PrimaryEmail** | Pointer to [**NullableUpdateUserRequestPrimaryEmail**](UpdateUserRequestPrimaryEmail.md) |  | [optional] 
**PrimaryPhone** | Pointer to [**NullableUpdateUserRequestPrimaryPhone**](UpdateUserRequestPrimaryPhone.md) |  | [optional] 
**Name** | Pointer to [**NullableUpdateUserRequestName**](UpdateUserRequestName.md) |  | [optional] 
**Avatar** | Pointer to [**NullableUpdateUserRequestAvatar**](UpdateUserRequestAvatar.md) |  | [optional] 
**CustomData** | Pointer to **map[string]interface{}** | Custom data object to update for the given user ID. Note this will replace the entire custom data object.  If you want to perform a partial update, use the &#x60;PATCH /api/users/{userId}/custom-data&#x60; endpoint instead. | [optional] 
**Profile** | Pointer to [**GetJwtCustomizer200ResponseOneOfContextSampleUserProfile**](GetJwtCustomizer200ResponseOneOfContextSampleUserProfile.md) |  | [optional] 

## Methods

### NewUpdateUserRequest

`func NewUpdateUserRequest() *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestWithDefaults

`func NewUpdateUserRequestWithDefaults() *UpdateUserRequest`

NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UpdateUserRequest) GetUsername() UpdateUserRequestUsername`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateUserRequest) GetUsernameOk() (*UpdateUserRequestUsername, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateUserRequest) SetUsername(v UpdateUserRequestUsername)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateUserRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UpdateUserRequest) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateUserRequest) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPrimaryEmail

`func (o *UpdateUserRequest) GetPrimaryEmail() UpdateUserRequestPrimaryEmail`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *UpdateUserRequest) GetPrimaryEmailOk() (*UpdateUserRequestPrimaryEmail, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *UpdateUserRequest) SetPrimaryEmail(v UpdateUserRequestPrimaryEmail)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *UpdateUserRequest) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### SetPrimaryEmailNil

`func (o *UpdateUserRequest) SetPrimaryEmailNil(b bool)`

 SetPrimaryEmailNil sets the value for PrimaryEmail to be an explicit nil

### UnsetPrimaryEmail
`func (o *UpdateUserRequest) UnsetPrimaryEmail()`

UnsetPrimaryEmail ensures that no value is present for PrimaryEmail, not even an explicit nil
### GetPrimaryPhone

`func (o *UpdateUserRequest) GetPrimaryPhone() UpdateUserRequestPrimaryPhone`

GetPrimaryPhone returns the PrimaryPhone field if non-nil, zero value otherwise.

### GetPrimaryPhoneOk

`func (o *UpdateUserRequest) GetPrimaryPhoneOk() (*UpdateUserRequestPrimaryPhone, bool)`

GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPhone

`func (o *UpdateUserRequest) SetPrimaryPhone(v UpdateUserRequestPrimaryPhone)`

SetPrimaryPhone sets PrimaryPhone field to given value.

### HasPrimaryPhone

`func (o *UpdateUserRequest) HasPrimaryPhone() bool`

HasPrimaryPhone returns a boolean if a field has been set.

### SetPrimaryPhoneNil

`func (o *UpdateUserRequest) SetPrimaryPhoneNil(b bool)`

 SetPrimaryPhoneNil sets the value for PrimaryPhone to be an explicit nil

### UnsetPrimaryPhone
`func (o *UpdateUserRequest) UnsetPrimaryPhone()`

UnsetPrimaryPhone ensures that no value is present for PrimaryPhone, not even an explicit nil
### GetName

`func (o *UpdateUserRequest) GetName() UpdateUserRequestName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserRequest) GetNameOk() (*UpdateUserRequestName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserRequest) SetName(v UpdateUserRequestName)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateUserRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateUserRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAvatar

`func (o *UpdateUserRequest) GetAvatar() UpdateUserRequestAvatar`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UpdateUserRequest) GetAvatarOk() (*UpdateUserRequestAvatar, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UpdateUserRequest) SetAvatar(v UpdateUserRequestAvatar)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UpdateUserRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *UpdateUserRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *UpdateUserRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetCustomData

`func (o *UpdateUserRequest) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *UpdateUserRequest) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *UpdateUserRequest) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *UpdateUserRequest) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetProfile

`func (o *UpdateUserRequest) GetProfile() GetJwtCustomizer200ResponseOneOfContextSampleUserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UpdateUserRequest) GetProfileOk() (*GetJwtCustomizer200ResponseOneOfContextSampleUserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UpdateUserRequest) SetProfile(v GetJwtCustomizer200ResponseOneOfContextSampleUserProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UpdateUserRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


