# GetJwtCustomizer200ResponseOneOfContextSampleUser

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
**SsoIdentities** | Pointer to [**[]GetJwtCustomizer200ResponseOneOfContextSampleUserSsoIdentitiesInner**](GetJwtCustomizer200ResponseOneOfContextSampleUserSsoIdentitiesInner.md) |  | [optional] 
**MfaVerificationFactors** | Pointer to **[]string** |  | [optional] 
**Roles** | Pointer to [**[]GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInner**](GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInner.md) |  | [optional] 
**Organizations** | Pointer to [**[]ListApplicationUserConsentScopes200ResponseOrganizationScopesInner**](ListApplicationUserConsentScopes200ResponseOrganizationScopesInner.md) |  | [optional] 
**OrganizationRoles** | Pointer to [**[]GetJwtCustomizer200ResponseOneOfContextSampleUserOrganizationRolesInner**](GetJwtCustomizer200ResponseOneOfContextSampleUserOrganizationRolesInner.md) |  | [optional] 

## Methods

### NewGetJwtCustomizer200ResponseOneOfContextSampleUser

`func NewGetJwtCustomizer200ResponseOneOfContextSampleUser() *GetJwtCustomizer200ResponseOneOfContextSampleUser`

NewGetJwtCustomizer200ResponseOneOfContextSampleUser instantiates a new GetJwtCustomizer200ResponseOneOfContextSampleUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJwtCustomizer200ResponseOneOfContextSampleUserWithDefaults

`func NewGetJwtCustomizer200ResponseOneOfContextSampleUserWithDefaults() *GetJwtCustomizer200ResponseOneOfContextSampleUser`

NewGetJwtCustomizer200ResponseOneOfContextSampleUserWithDefaults instantiates a new GetJwtCustomizer200ResponseOneOfContextSampleUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPrimaryEmail

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### SetPrimaryEmailNil

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetPrimaryEmailNil(b bool)`

 SetPrimaryEmailNil sets the value for PrimaryEmail to be an explicit nil

### UnsetPrimaryEmail
`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) UnsetPrimaryEmail()`

UnsetPrimaryEmail ensures that no value is present for PrimaryEmail, not even an explicit nil
### GetPrimaryPhone

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetPrimaryPhone() string`

GetPrimaryPhone returns the PrimaryPhone field if non-nil, zero value otherwise.

### GetPrimaryPhoneOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetPrimaryPhoneOk() (*string, bool)`

GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPhone

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetPrimaryPhone(v string)`

SetPrimaryPhone sets PrimaryPhone field to given value.

### HasPrimaryPhone

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasPrimaryPhone() bool`

HasPrimaryPhone returns a boolean if a field has been set.

### SetPrimaryPhoneNil

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetPrimaryPhoneNil(b bool)`

 SetPrimaryPhoneNil sets the value for PrimaryPhone to be an explicit nil

### UnsetPrimaryPhone
`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) UnsetPrimaryPhone()`

UnsetPrimaryPhone ensures that no value is present for PrimaryPhone, not even an explicit nil
### GetName

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAvatar

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetCustomData

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetIdentities

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetIdentities() map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetIdentitiesOk() (*map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetIdentities(v map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### GetLastSignInAt

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetLastSignInAt() float32`

GetLastSignInAt returns the LastSignInAt field if non-nil, zero value otherwise.

### GetLastSignInAtOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetLastSignInAtOk() (*float32, bool)`

GetLastSignInAtOk returns a tuple with the LastSignInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignInAt

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetLastSignInAt(v float32)`

SetLastSignInAt sets LastSignInAt field to given value.

### HasLastSignInAt

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasLastSignInAt() bool`

HasLastSignInAt returns a boolean if a field has been set.

### SetLastSignInAtNil

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetLastSignInAtNil(b bool)`

 SetLastSignInAtNil sets the value for LastSignInAt to be an explicit nil

### UnsetLastSignInAt
`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) UnsetLastSignInAt()`

UnsetLastSignInAt ensures that no value is present for LastSignInAt, not even an explicit nil
### GetCreatedAt

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProfile

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetProfile() GetJwtCustomizer200ResponseOneOfContextSampleUserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetProfileOk() (*GetJwtCustomizer200ResponseOneOfContextSampleUserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetProfile(v GetJwtCustomizer200ResponseOneOfContextSampleUserProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetApplicationId

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationIdNil

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetIsSuspended

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.

### HasIsSuspended

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasIsSuspended() bool`

HasIsSuspended returns a boolean if a field has been set.

### GetHasPassword

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetSsoIdentities

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetSsoIdentities() []GetJwtCustomizer200ResponseOneOfContextSampleUserSsoIdentitiesInner`

GetSsoIdentities returns the SsoIdentities field if non-nil, zero value otherwise.

### GetSsoIdentitiesOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetSsoIdentitiesOk() (*[]GetJwtCustomizer200ResponseOneOfContextSampleUserSsoIdentitiesInner, bool)`

GetSsoIdentitiesOk returns a tuple with the SsoIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoIdentities

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetSsoIdentities(v []GetJwtCustomizer200ResponseOneOfContextSampleUserSsoIdentitiesInner)`

SetSsoIdentities sets SsoIdentities field to given value.

### HasSsoIdentities

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasSsoIdentities() bool`

HasSsoIdentities returns a boolean if a field has been set.

### GetMfaVerificationFactors

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetMfaVerificationFactors() []string`

GetMfaVerificationFactors returns the MfaVerificationFactors field if non-nil, zero value otherwise.

### GetMfaVerificationFactorsOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetMfaVerificationFactorsOk() (*[]string, bool)`

GetMfaVerificationFactorsOk returns a tuple with the MfaVerificationFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaVerificationFactors

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetMfaVerificationFactors(v []string)`

SetMfaVerificationFactors sets MfaVerificationFactors field to given value.

### HasMfaVerificationFactors

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasMfaVerificationFactors() bool`

HasMfaVerificationFactors returns a boolean if a field has been set.

### GetRoles

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetRoles() []GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetRolesOk() (*[]GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetRoles(v []GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetOrganizations

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetOrganizations() []ListApplicationUserConsentScopes200ResponseOrganizationScopesInner`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetOrganizationsOk() (*[]ListApplicationUserConsentScopes200ResponseOrganizationScopesInner, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetOrganizations(v []ListApplicationUserConsentScopes200ResponseOrganizationScopesInner)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetOrganizationRoles

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetOrganizationRoles() []GetJwtCustomizer200ResponseOneOfContextSampleUserOrganizationRolesInner`

GetOrganizationRoles returns the OrganizationRoles field if non-nil, zero value otherwise.

### GetOrganizationRolesOk

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) GetOrganizationRolesOk() (*[]GetJwtCustomizer200ResponseOneOfContextSampleUserOrganizationRolesInner, bool)`

GetOrganizationRolesOk returns a tuple with the OrganizationRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoles

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) SetOrganizationRoles(v []GetJwtCustomizer200ResponseOneOfContextSampleUserOrganizationRolesInner)`

SetOrganizationRoles sets OrganizationRoles field to given value.

### HasOrganizationRoles

`func (o *GetJwtCustomizer200ResponseOneOfContextSampleUser) HasOrganizationRoles() bool`

HasOrganizationRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


