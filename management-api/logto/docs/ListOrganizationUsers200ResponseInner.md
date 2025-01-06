# ListOrganizationUsers200ResponseInner

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
**OrganizationRoles** | [**[]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner**](ListApplicationOrganizations200ResponseInnerOrganizationRolesInner.md) |  | 

## Methods

### NewListOrganizationUsers200ResponseInner

`func NewListOrganizationUsers200ResponseInner(id string, username NullableString, primaryEmail NullableString, primaryPhone NullableString, name NullableString, avatar NullableString, customData map[string]interface{}, identities map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue, lastSignInAt NullableFloat32, createdAt float32, updatedAt float32, profile GetJwtCustomizer200ResponseOneOfContextSampleUserProfile, applicationId NullableString, isSuspended bool, organizationRoles []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, ) *ListOrganizationUsers200ResponseInner`

NewListOrganizationUsers200ResponseInner instantiates a new ListOrganizationUsers200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizationUsers200ResponseInnerWithDefaults

`func NewListOrganizationUsers200ResponseInnerWithDefaults() *ListOrganizationUsers200ResponseInner`

NewListOrganizationUsers200ResponseInnerWithDefaults instantiates a new ListOrganizationUsers200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListOrganizationUsers200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOrganizationUsers200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOrganizationUsers200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *ListOrganizationUsers200ResponseInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListOrganizationUsers200ResponseInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListOrganizationUsers200ResponseInner) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *ListOrganizationUsers200ResponseInner) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ListOrganizationUsers200ResponseInner) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPrimaryEmail

`func (o *ListOrganizationUsers200ResponseInner) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *ListOrganizationUsers200ResponseInner) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *ListOrganizationUsers200ResponseInner) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.


### SetPrimaryEmailNil

`func (o *ListOrganizationUsers200ResponseInner) SetPrimaryEmailNil(b bool)`

 SetPrimaryEmailNil sets the value for PrimaryEmail to be an explicit nil

### UnsetPrimaryEmail
`func (o *ListOrganizationUsers200ResponseInner) UnsetPrimaryEmail()`

UnsetPrimaryEmail ensures that no value is present for PrimaryEmail, not even an explicit nil
### GetPrimaryPhone

`func (o *ListOrganizationUsers200ResponseInner) GetPrimaryPhone() string`

GetPrimaryPhone returns the PrimaryPhone field if non-nil, zero value otherwise.

### GetPrimaryPhoneOk

`func (o *ListOrganizationUsers200ResponseInner) GetPrimaryPhoneOk() (*string, bool)`

GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPhone

`func (o *ListOrganizationUsers200ResponseInner) SetPrimaryPhone(v string)`

SetPrimaryPhone sets PrimaryPhone field to given value.


### SetPrimaryPhoneNil

`func (o *ListOrganizationUsers200ResponseInner) SetPrimaryPhoneNil(b bool)`

 SetPrimaryPhoneNil sets the value for PrimaryPhone to be an explicit nil

### UnsetPrimaryPhone
`func (o *ListOrganizationUsers200ResponseInner) UnsetPrimaryPhone()`

UnsetPrimaryPhone ensures that no value is present for PrimaryPhone, not even an explicit nil
### GetName

`func (o *ListOrganizationUsers200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListOrganizationUsers200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListOrganizationUsers200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ListOrganizationUsers200ResponseInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ListOrganizationUsers200ResponseInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAvatar

`func (o *ListOrganizationUsers200ResponseInner) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ListOrganizationUsers200ResponseInner) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ListOrganizationUsers200ResponseInner) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *ListOrganizationUsers200ResponseInner) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *ListOrganizationUsers200ResponseInner) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetCustomData

`func (o *ListOrganizationUsers200ResponseInner) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *ListOrganizationUsers200ResponseInner) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *ListOrganizationUsers200ResponseInner) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.


### GetIdentities

`func (o *ListOrganizationUsers200ResponseInner) GetIdentities() map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ListOrganizationUsers200ResponseInner) GetIdentitiesOk() (*map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ListOrganizationUsers200ResponseInner) SetIdentities(v map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue)`

SetIdentities sets Identities field to given value.


### GetLastSignInAt

`func (o *ListOrganizationUsers200ResponseInner) GetLastSignInAt() float32`

GetLastSignInAt returns the LastSignInAt field if non-nil, zero value otherwise.

### GetLastSignInAtOk

`func (o *ListOrganizationUsers200ResponseInner) GetLastSignInAtOk() (*float32, bool)`

GetLastSignInAtOk returns a tuple with the LastSignInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignInAt

`func (o *ListOrganizationUsers200ResponseInner) SetLastSignInAt(v float32)`

SetLastSignInAt sets LastSignInAt field to given value.


### SetLastSignInAtNil

`func (o *ListOrganizationUsers200ResponseInner) SetLastSignInAtNil(b bool)`

 SetLastSignInAtNil sets the value for LastSignInAt to be an explicit nil

### UnsetLastSignInAt
`func (o *ListOrganizationUsers200ResponseInner) UnsetLastSignInAt()`

UnsetLastSignInAt ensures that no value is present for LastSignInAt, not even an explicit nil
### GetCreatedAt

`func (o *ListOrganizationUsers200ResponseInner) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListOrganizationUsers200ResponseInner) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListOrganizationUsers200ResponseInner) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListOrganizationUsers200ResponseInner) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListOrganizationUsers200ResponseInner) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListOrganizationUsers200ResponseInner) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetProfile

`func (o *ListOrganizationUsers200ResponseInner) GetProfile() GetJwtCustomizer200ResponseOneOfContextSampleUserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ListOrganizationUsers200ResponseInner) GetProfileOk() (*GetJwtCustomizer200ResponseOneOfContextSampleUserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ListOrganizationUsers200ResponseInner) SetProfile(v GetJwtCustomizer200ResponseOneOfContextSampleUserProfile)`

SetProfile sets Profile field to given value.


### GetApplicationId

`func (o *ListOrganizationUsers200ResponseInner) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ListOrganizationUsers200ResponseInner) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ListOrganizationUsers200ResponseInner) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### SetApplicationIdNil

`func (o *ListOrganizationUsers200ResponseInner) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *ListOrganizationUsers200ResponseInner) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetIsSuspended

`func (o *ListOrganizationUsers200ResponseInner) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *ListOrganizationUsers200ResponseInner) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *ListOrganizationUsers200ResponseInner) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.


### GetOrganizationRoles

`func (o *ListOrganizationUsers200ResponseInner) GetOrganizationRoles() []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner`

GetOrganizationRoles returns the OrganizationRoles field if non-nil, zero value otherwise.

### GetOrganizationRolesOk

`func (o *ListOrganizationUsers200ResponseInner) GetOrganizationRolesOk() (*[]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, bool)`

GetOrganizationRolesOk returns a tuple with the OrganizationRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoles

`func (o *ListOrganizationUsers200ResponseInner) SetOrganizationRoles(v []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner)`

SetOrganizationRoles sets OrganizationRoles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


