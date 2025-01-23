# ListApplicationOrganizations200ResponseInner

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
**OrganizationRoles** | [**[]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner**](ListApplicationOrganizations200ResponseInnerOrganizationRolesInner.md) |  | 

## Methods

### NewListApplicationOrganizations200ResponseInner

`func NewListApplicationOrganizations200ResponseInner(tenantId string, id string, name string, description NullableString, customData map[string]interface{}, isMfaRequired bool, branding ListApplicationOrganizations200ResponseInnerBranding, createdAt float32, organizationRoles []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, ) *ListApplicationOrganizations200ResponseInner`

NewListApplicationOrganizations200ResponseInner instantiates a new ListApplicationOrganizations200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApplicationOrganizations200ResponseInnerWithDefaults

`func NewListApplicationOrganizations200ResponseInnerWithDefaults() *ListApplicationOrganizations200ResponseInner`

NewListApplicationOrganizations200ResponseInnerWithDefaults instantiates a new ListApplicationOrganizations200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListApplicationOrganizations200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListApplicationOrganizations200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListApplicationOrganizations200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListApplicationOrganizations200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListApplicationOrganizations200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListApplicationOrganizations200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListApplicationOrganizations200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListApplicationOrganizations200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListApplicationOrganizations200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListApplicationOrganizations200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListApplicationOrganizations200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListApplicationOrganizations200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ListApplicationOrganizations200ResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListApplicationOrganizations200ResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCustomData

`func (o *ListApplicationOrganizations200ResponseInner) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *ListApplicationOrganizations200ResponseInner) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *ListApplicationOrganizations200ResponseInner) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.


### GetIsMfaRequired

`func (o *ListApplicationOrganizations200ResponseInner) GetIsMfaRequired() bool`

GetIsMfaRequired returns the IsMfaRequired field if non-nil, zero value otherwise.

### GetIsMfaRequiredOk

`func (o *ListApplicationOrganizations200ResponseInner) GetIsMfaRequiredOk() (*bool, bool)`

GetIsMfaRequiredOk returns a tuple with the IsMfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMfaRequired

`func (o *ListApplicationOrganizations200ResponseInner) SetIsMfaRequired(v bool)`

SetIsMfaRequired sets IsMfaRequired field to given value.


### GetBranding

`func (o *ListApplicationOrganizations200ResponseInner) GetBranding() ListApplicationOrganizations200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *ListApplicationOrganizations200ResponseInner) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *ListApplicationOrganizations200ResponseInner) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding)`

SetBranding sets Branding field to given value.


### GetCreatedAt

`func (o *ListApplicationOrganizations200ResponseInner) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListApplicationOrganizations200ResponseInner) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListApplicationOrganizations200ResponseInner) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetOrganizationRoles

`func (o *ListApplicationOrganizations200ResponseInner) GetOrganizationRoles() []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner`

GetOrganizationRoles returns the OrganizationRoles field if non-nil, zero value otherwise.

### GetOrganizationRolesOk

`func (o *ListApplicationOrganizations200ResponseInner) GetOrganizationRolesOk() (*[]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, bool)`

GetOrganizationRolesOk returns a tuple with the OrganizationRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoles

`func (o *ListApplicationOrganizations200ResponseInner) SetOrganizationRoles(v []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner)`

SetOrganizationRoles sets OrganizationRoles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


