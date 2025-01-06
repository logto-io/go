# ListOrganizations200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**CustomData** | **map[string]interface{}** | arbitrary | 
**IsMfaRequired** | **bool** |  | 
**Branding** | [**ListApplicationOrganizations200ResponseInnerBranding**](ListApplicationOrganizations200ResponseInnerBranding.md) |  | 
**CreatedAt** | **float32** |  | 
**UsersCount** | Pointer to **float32** |  | [optional] 
**FeaturedUsers** | Pointer to [**[]ListRoles200ResponseInnerFeaturedUsersInner**](ListRoles200ResponseInnerFeaturedUsersInner.md) |  | [optional] 

## Methods

### NewListOrganizations200ResponseInner

`func NewListOrganizations200ResponseInner(tenantId string, id string, name string, description NullableString, customData map[string]interface{}, isMfaRequired bool, branding ListApplicationOrganizations200ResponseInnerBranding, createdAt float32, ) *ListOrganizations200ResponseInner`

NewListOrganizations200ResponseInner instantiates a new ListOrganizations200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizations200ResponseInnerWithDefaults

`func NewListOrganizations200ResponseInnerWithDefaults() *ListOrganizations200ResponseInner`

NewListOrganizations200ResponseInnerWithDefaults instantiates a new ListOrganizations200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListOrganizations200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListOrganizations200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListOrganizations200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListOrganizations200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOrganizations200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOrganizations200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListOrganizations200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListOrganizations200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListOrganizations200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListOrganizations200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListOrganizations200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListOrganizations200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ListOrganizations200ResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListOrganizations200ResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCustomData

`func (o *ListOrganizations200ResponseInner) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *ListOrganizations200ResponseInner) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *ListOrganizations200ResponseInner) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.


### GetIsMfaRequired

`func (o *ListOrganizations200ResponseInner) GetIsMfaRequired() bool`

GetIsMfaRequired returns the IsMfaRequired field if non-nil, zero value otherwise.

### GetIsMfaRequiredOk

`func (o *ListOrganizations200ResponseInner) GetIsMfaRequiredOk() (*bool, bool)`

GetIsMfaRequiredOk returns a tuple with the IsMfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMfaRequired

`func (o *ListOrganizations200ResponseInner) SetIsMfaRequired(v bool)`

SetIsMfaRequired sets IsMfaRequired field to given value.


### GetBranding

`func (o *ListOrganizations200ResponseInner) GetBranding() ListApplicationOrganizations200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *ListOrganizations200ResponseInner) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *ListOrganizations200ResponseInner) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding)`

SetBranding sets Branding field to given value.


### GetCreatedAt

`func (o *ListOrganizations200ResponseInner) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListOrganizations200ResponseInner) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListOrganizations200ResponseInner) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUsersCount

`func (o *ListOrganizations200ResponseInner) GetUsersCount() float32`

GetUsersCount returns the UsersCount field if non-nil, zero value otherwise.

### GetUsersCountOk

`func (o *ListOrganizations200ResponseInner) GetUsersCountOk() (*float32, bool)`

GetUsersCountOk returns a tuple with the UsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCount

`func (o *ListOrganizations200ResponseInner) SetUsersCount(v float32)`

SetUsersCount sets UsersCount field to given value.

### HasUsersCount

`func (o *ListOrganizations200ResponseInner) HasUsersCount() bool`

HasUsersCount returns a boolean if a field has been set.

### GetFeaturedUsers

`func (o *ListOrganizations200ResponseInner) GetFeaturedUsers() []ListRoles200ResponseInnerFeaturedUsersInner`

GetFeaturedUsers returns the FeaturedUsers field if non-nil, zero value otherwise.

### GetFeaturedUsersOk

`func (o *ListOrganizations200ResponseInner) GetFeaturedUsersOk() (*[]ListRoles200ResponseInnerFeaturedUsersInner, bool)`

GetFeaturedUsersOk returns a tuple with the FeaturedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedUsers

`func (o *ListOrganizations200ResponseInner) SetFeaturedUsers(v []ListRoles200ResponseInnerFeaturedUsersInner)`

SetFeaturedUsers sets FeaturedUsers field to given value.

### HasFeaturedUsers

`func (o *ListOrganizations200ResponseInner) HasFeaturedUsers() bool`

HasFeaturedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


