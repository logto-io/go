# GetUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | **NullableString** |  | 
**PrimaryEmail** | **NullableString** |  | 
**PrimaryPhone** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**Avatar** | **NullableString** |  | 
**CustomData** | **map[string]interface{}** | arbitrary | 
**Identities** | [**map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue**](GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue.md) |  | 
**LastSignInAt** | **NullableFloat32** |  | 
**CreatedAt** | **float32** |  | 
**UpdatedAt** | **float32** |  | 
**Profile** | [**GetJwtCustomizer200ResponseOneOfContextSampleUserProfile**](GetJwtCustomizer200ResponseOneOfContextSampleUserProfile.md) |  | 
**ApplicationId** | **NullableString** |  | 
**IsSuspended** | **bool** |  | 
**HasPassword** | Pointer to **bool** |  | [optional] 
**SsoIdentities** | Pointer to [**[]GetUser200ResponseSsoIdentitiesInner**](GetUser200ResponseSsoIdentitiesInner.md) | List of SSO identities associated with the user. Only available when the &#x60;includeSsoIdentities&#x60; query parameter is provided with a truthy value. | [optional] 

## Methods

### NewGetUser200Response

`func NewGetUser200Response(id string, username NullableString, primaryEmail NullableString, primaryPhone NullableString, name NullableString, avatar NullableString, customData map[string]interface{}, identities map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue, lastSignInAt NullableFloat32, createdAt float32, updatedAt float32, profile GetJwtCustomizer200ResponseOneOfContextSampleUserProfile, applicationId NullableString, isSuspended bool, ) *GetUser200Response`

NewGetUser200Response instantiates a new GetUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUser200ResponseWithDefaults

`func NewGetUser200ResponseWithDefaults() *GetUser200Response`

NewGetUser200ResponseWithDefaults instantiates a new GetUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetUser200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUser200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUser200Response) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *GetUser200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetUser200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetUser200Response) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *GetUser200Response) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *GetUser200Response) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPrimaryEmail

`func (o *GetUser200Response) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *GetUser200Response) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *GetUser200Response) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.


### SetPrimaryEmailNil

`func (o *GetUser200Response) SetPrimaryEmailNil(b bool)`

 SetPrimaryEmailNil sets the value for PrimaryEmail to be an explicit nil

### UnsetPrimaryEmail
`func (o *GetUser200Response) UnsetPrimaryEmail()`

UnsetPrimaryEmail ensures that no value is present for PrimaryEmail, not even an explicit nil
### GetPrimaryPhone

`func (o *GetUser200Response) GetPrimaryPhone() string`

GetPrimaryPhone returns the PrimaryPhone field if non-nil, zero value otherwise.

### GetPrimaryPhoneOk

`func (o *GetUser200Response) GetPrimaryPhoneOk() (*string, bool)`

GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPhone

`func (o *GetUser200Response) SetPrimaryPhone(v string)`

SetPrimaryPhone sets PrimaryPhone field to given value.


### SetPrimaryPhoneNil

`func (o *GetUser200Response) SetPrimaryPhoneNil(b bool)`

 SetPrimaryPhoneNil sets the value for PrimaryPhone to be an explicit nil

### UnsetPrimaryPhone
`func (o *GetUser200Response) UnsetPrimaryPhone()`

UnsetPrimaryPhone ensures that no value is present for PrimaryPhone, not even an explicit nil
### GetName

`func (o *GetUser200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUser200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUser200Response) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GetUser200Response) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetUser200Response) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAvatar

`func (o *GetUser200Response) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *GetUser200Response) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *GetUser200Response) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *GetUser200Response) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *GetUser200Response) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetCustomData

`func (o *GetUser200Response) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *GetUser200Response) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *GetUser200Response) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.


### GetIdentities

`func (o *GetUser200Response) GetIdentities() map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *GetUser200Response) GetIdentitiesOk() (*map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *GetUser200Response) SetIdentities(v map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue)`

SetIdentities sets Identities field to given value.


### GetLastSignInAt

`func (o *GetUser200Response) GetLastSignInAt() float32`

GetLastSignInAt returns the LastSignInAt field if non-nil, zero value otherwise.

### GetLastSignInAtOk

`func (o *GetUser200Response) GetLastSignInAtOk() (*float32, bool)`

GetLastSignInAtOk returns a tuple with the LastSignInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignInAt

`func (o *GetUser200Response) SetLastSignInAt(v float32)`

SetLastSignInAt sets LastSignInAt field to given value.


### SetLastSignInAtNil

`func (o *GetUser200Response) SetLastSignInAtNil(b bool)`

 SetLastSignInAtNil sets the value for LastSignInAt to be an explicit nil

### UnsetLastSignInAt
`func (o *GetUser200Response) UnsetLastSignInAt()`

UnsetLastSignInAt ensures that no value is present for LastSignInAt, not even an explicit nil
### GetCreatedAt

`func (o *GetUser200Response) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetUser200Response) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetUser200Response) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetUser200Response) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetUser200Response) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetUser200Response) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetProfile

`func (o *GetUser200Response) GetProfile() GetJwtCustomizer200ResponseOneOfContextSampleUserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetUser200Response) GetProfileOk() (*GetJwtCustomizer200ResponseOneOfContextSampleUserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetUser200Response) SetProfile(v GetJwtCustomizer200ResponseOneOfContextSampleUserProfile)`

SetProfile sets Profile field to given value.


### GetApplicationId

`func (o *GetUser200Response) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GetUser200Response) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GetUser200Response) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### SetApplicationIdNil

`func (o *GetUser200Response) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *GetUser200Response) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetIsSuspended

`func (o *GetUser200Response) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *GetUser200Response) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *GetUser200Response) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.


### GetHasPassword

`func (o *GetUser200Response) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *GetUser200Response) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *GetUser200Response) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *GetUser200Response) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetSsoIdentities

`func (o *GetUser200Response) GetSsoIdentities() []GetUser200ResponseSsoIdentitiesInner`

GetSsoIdentities returns the SsoIdentities field if non-nil, zero value otherwise.

### GetSsoIdentitiesOk

`func (o *GetUser200Response) GetSsoIdentitiesOk() (*[]GetUser200ResponseSsoIdentitiesInner, bool)`

GetSsoIdentitiesOk returns a tuple with the SsoIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoIdentities

`func (o *GetUser200Response) SetSsoIdentities(v []GetUser200ResponseSsoIdentitiesInner)`

SetSsoIdentities sets SsoIdentities field to given value.

### HasSsoIdentities

`func (o *GetUser200Response) HasSsoIdentities() bool`

HasSsoIdentities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


