# GetProfile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**PrimaryEmail** | Pointer to **NullableString** |  | [optional] 
**PrimaryPhone** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**CustomData** | Pointer to **map[string]interface{}** | arbitrary | [optional] 
**Identities** | Pointer to [**map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue**](GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue.md) |  | [optional] 
**LastSignInAt** | Pointer to **NullableFloat32** |  | [optional] 
**CreatedAt** | Pointer to **float32** |  | [optional] 
**UpdatedAt** | Pointer to **float32** |  | [optional] 
**Profile** | Pointer to [**GetJwtCustomizer200ResponseOneOfContextSampleUserProfile**](GetJwtCustomizer200ResponseOneOfContextSampleUserProfile.md) |  | [optional] 
**ApplicationId** | Pointer to **NullableString** |  | [optional] 
**IsSuspended** | Pointer to **bool** |  | [optional] 
**HasPassword** | Pointer to **bool** |  | [optional] 
**SsoIdentities** | Pointer to [**[]GetUser200ResponseSsoIdentitiesInner**](GetUser200ResponseSsoIdentitiesInner.md) |  | [optional] 

## Methods

### NewGetProfile200Response

`func NewGetProfile200Response() *GetProfile200Response`

NewGetProfile200Response instantiates a new GetProfile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfile200ResponseWithDefaults

`func NewGetProfile200ResponseWithDefaults() *GetProfile200Response`

NewGetProfile200ResponseWithDefaults instantiates a new GetProfile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetProfile200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetProfile200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetProfile200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetProfile200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *GetProfile200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetProfile200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetProfile200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetProfile200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *GetProfile200Response) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *GetProfile200Response) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPrimaryEmail

`func (o *GetProfile200Response) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *GetProfile200Response) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *GetProfile200Response) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *GetProfile200Response) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### SetPrimaryEmailNil

`func (o *GetProfile200Response) SetPrimaryEmailNil(b bool)`

 SetPrimaryEmailNil sets the value for PrimaryEmail to be an explicit nil

### UnsetPrimaryEmail
`func (o *GetProfile200Response) UnsetPrimaryEmail()`

UnsetPrimaryEmail ensures that no value is present for PrimaryEmail, not even an explicit nil
### GetPrimaryPhone

`func (o *GetProfile200Response) GetPrimaryPhone() string`

GetPrimaryPhone returns the PrimaryPhone field if non-nil, zero value otherwise.

### GetPrimaryPhoneOk

`func (o *GetProfile200Response) GetPrimaryPhoneOk() (*string, bool)`

GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPhone

`func (o *GetProfile200Response) SetPrimaryPhone(v string)`

SetPrimaryPhone sets PrimaryPhone field to given value.

### HasPrimaryPhone

`func (o *GetProfile200Response) HasPrimaryPhone() bool`

HasPrimaryPhone returns a boolean if a field has been set.

### SetPrimaryPhoneNil

`func (o *GetProfile200Response) SetPrimaryPhoneNil(b bool)`

 SetPrimaryPhoneNil sets the value for PrimaryPhone to be an explicit nil

### UnsetPrimaryPhone
`func (o *GetProfile200Response) UnsetPrimaryPhone()`

UnsetPrimaryPhone ensures that no value is present for PrimaryPhone, not even an explicit nil
### GetName

`func (o *GetProfile200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetProfile200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetProfile200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetProfile200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GetProfile200Response) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetProfile200Response) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAvatar

`func (o *GetProfile200Response) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *GetProfile200Response) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *GetProfile200Response) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *GetProfile200Response) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *GetProfile200Response) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *GetProfile200Response) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetCustomData

`func (o *GetProfile200Response) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *GetProfile200Response) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *GetProfile200Response) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *GetProfile200Response) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetIdentities

`func (o *GetProfile200Response) GetIdentities() map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *GetProfile200Response) GetIdentitiesOk() (*map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *GetProfile200Response) SetIdentities(v map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *GetProfile200Response) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### GetLastSignInAt

`func (o *GetProfile200Response) GetLastSignInAt() float32`

GetLastSignInAt returns the LastSignInAt field if non-nil, zero value otherwise.

### GetLastSignInAtOk

`func (o *GetProfile200Response) GetLastSignInAtOk() (*float32, bool)`

GetLastSignInAtOk returns a tuple with the LastSignInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignInAt

`func (o *GetProfile200Response) SetLastSignInAt(v float32)`

SetLastSignInAt sets LastSignInAt field to given value.

### HasLastSignInAt

`func (o *GetProfile200Response) HasLastSignInAt() bool`

HasLastSignInAt returns a boolean if a field has been set.

### SetLastSignInAtNil

`func (o *GetProfile200Response) SetLastSignInAtNil(b bool)`

 SetLastSignInAtNil sets the value for LastSignInAt to be an explicit nil

### UnsetLastSignInAt
`func (o *GetProfile200Response) UnsetLastSignInAt()`

UnsetLastSignInAt ensures that no value is present for LastSignInAt, not even an explicit nil
### GetCreatedAt

`func (o *GetProfile200Response) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetProfile200Response) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetProfile200Response) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetProfile200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetProfile200Response) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetProfile200Response) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetProfile200Response) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetProfile200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProfile

`func (o *GetProfile200Response) GetProfile() GetJwtCustomizer200ResponseOneOfContextSampleUserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetProfile200Response) GetProfileOk() (*GetJwtCustomizer200ResponseOneOfContextSampleUserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetProfile200Response) SetProfile(v GetJwtCustomizer200ResponseOneOfContextSampleUserProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GetProfile200Response) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetApplicationId

`func (o *GetProfile200Response) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GetProfile200Response) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GetProfile200Response) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *GetProfile200Response) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationIdNil

`func (o *GetProfile200Response) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *GetProfile200Response) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetIsSuspended

`func (o *GetProfile200Response) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *GetProfile200Response) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *GetProfile200Response) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.

### HasIsSuspended

`func (o *GetProfile200Response) HasIsSuspended() bool`

HasIsSuspended returns a boolean if a field has been set.

### GetHasPassword

`func (o *GetProfile200Response) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *GetProfile200Response) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *GetProfile200Response) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *GetProfile200Response) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetSsoIdentities

`func (o *GetProfile200Response) GetSsoIdentities() []GetUser200ResponseSsoIdentitiesInner`

GetSsoIdentities returns the SsoIdentities field if non-nil, zero value otherwise.

### GetSsoIdentitiesOk

`func (o *GetProfile200Response) GetSsoIdentitiesOk() (*[]GetUser200ResponseSsoIdentitiesInner, bool)`

GetSsoIdentitiesOk returns a tuple with the SsoIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoIdentities

`func (o *GetProfile200Response) SetSsoIdentities(v []GetUser200ResponseSsoIdentitiesInner)`

SetSsoIdentities sets SsoIdentities field to given value.

### HasSsoIdentities

`func (o *GetProfile200Response) HasSsoIdentities() bool`

HasSsoIdentities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


