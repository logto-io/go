# CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryPhone** | Pointer to **string** | Primary phone number for the user. It should be unique across all users. | [optional] 
**PrimaryEmail** | Pointer to **string** | Primary email address for the user. It should be unique across all users. | [optional] 
**Username** | Pointer to **string** | Username for the user. It should be unique across all users. | [optional] 
**Password** | Pointer to **string** | Plain text password for the user. | [optional] 
**PasswordDigest** | Pointer to **string** | In case you already have the password digests and not the passwords, you can use them for the newly created user via this property. The value should be generated with one of the supported algorithms. The algorithm can be specified using the &#x60;passwordAlgorithm&#x60; property. | [optional] 
**PasswordAlgorithm** | Pointer to **string** | The hash algorithm used for the password. It should be one of the supported algorithms: argon2, md5, sha1, sha256. Should the encryption algorithm differ from argon2, it will automatically be upgraded to argon2 upon the user&#39;s next sign-in. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to [**NullableUpdateUserRequestAvatar**](UpdateUserRequestAvatar.md) |  | [optional] 
**CustomData** | Pointer to **map[string]interface{}** | arbitrary | [optional] 
**Profile** | Pointer to [**GetJwtCustomizer200ResponseOneOfContextSampleUserProfile**](GetJwtCustomizer200ResponseOneOfContextSampleUserProfile.md) |  | [optional] 

## Methods

### NewCreateUserRequest

`func NewCreateUserRequest() *CreateUserRequest`

NewCreateUserRequest instantiates a new CreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRequestWithDefaults

`func NewCreateUserRequestWithDefaults() *CreateUserRequest`

NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryPhone

`func (o *CreateUserRequest) GetPrimaryPhone() string`

GetPrimaryPhone returns the PrimaryPhone field if non-nil, zero value otherwise.

### GetPrimaryPhoneOk

`func (o *CreateUserRequest) GetPrimaryPhoneOk() (*string, bool)`

GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPhone

`func (o *CreateUserRequest) SetPrimaryPhone(v string)`

SetPrimaryPhone sets PrimaryPhone field to given value.

### HasPrimaryPhone

`func (o *CreateUserRequest) HasPrimaryPhone() bool`

HasPrimaryPhone returns a boolean if a field has been set.

### GetPrimaryEmail

`func (o *CreateUserRequest) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *CreateUserRequest) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *CreateUserRequest) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *CreateUserRequest) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### GetUsername

`func (o *CreateUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateUserRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *CreateUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordDigest

`func (o *CreateUserRequest) GetPasswordDigest() string`

GetPasswordDigest returns the PasswordDigest field if non-nil, zero value otherwise.

### GetPasswordDigestOk

`func (o *CreateUserRequest) GetPasswordDigestOk() (*string, bool)`

GetPasswordDigestOk returns a tuple with the PasswordDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordDigest

`func (o *CreateUserRequest) SetPasswordDigest(v string)`

SetPasswordDigest sets PasswordDigest field to given value.

### HasPasswordDigest

`func (o *CreateUserRequest) HasPasswordDigest() bool`

HasPasswordDigest returns a boolean if a field has been set.

### GetPasswordAlgorithm

`func (o *CreateUserRequest) GetPasswordAlgorithm() string`

GetPasswordAlgorithm returns the PasswordAlgorithm field if non-nil, zero value otherwise.

### GetPasswordAlgorithmOk

`func (o *CreateUserRequest) GetPasswordAlgorithmOk() (*string, bool)`

GetPasswordAlgorithmOk returns a tuple with the PasswordAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordAlgorithm

`func (o *CreateUserRequest) SetPasswordAlgorithm(v string)`

SetPasswordAlgorithm sets PasswordAlgorithm field to given value.

### HasPasswordAlgorithm

`func (o *CreateUserRequest) HasPasswordAlgorithm() bool`

HasPasswordAlgorithm returns a boolean if a field has been set.

### GetName

`func (o *CreateUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvatar

`func (o *CreateUserRequest) GetAvatar() UpdateUserRequestAvatar`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CreateUserRequest) GetAvatarOk() (*UpdateUserRequestAvatar, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CreateUserRequest) SetAvatar(v UpdateUserRequestAvatar)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CreateUserRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *CreateUserRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *CreateUserRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetCustomData

`func (o *CreateUserRequest) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *CreateUserRequest) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *CreateUserRequest) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *CreateUserRequest) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetProfile

`func (o *CreateUserRequest) GetProfile() GetJwtCustomizer200ResponseOneOfContextSampleUserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CreateUserRequest) GetProfileOk() (*GetJwtCustomizer200ResponseOneOfContextSampleUserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CreateUserRequest) SetProfile(v GetJwtCustomizer200ResponseOneOfContextSampleUserProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CreateUserRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


