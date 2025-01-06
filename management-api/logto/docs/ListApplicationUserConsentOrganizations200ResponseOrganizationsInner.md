# ListApplicationUserConsentOrganizations200ResponseOrganizationsInner

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

## Methods

### NewListApplicationUserConsentOrganizations200ResponseOrganizationsInner

`func NewListApplicationUserConsentOrganizations200ResponseOrganizationsInner(tenantId string, id string, name string, description NullableString, customData map[string]interface{}, isMfaRequired bool, branding ListApplicationOrganizations200ResponseInnerBranding, createdAt float32, ) *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner`

NewListApplicationUserConsentOrganizations200ResponseOrganizationsInner instantiates a new ListApplicationUserConsentOrganizations200ResponseOrganizationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApplicationUserConsentOrganizations200ResponseOrganizationsInnerWithDefaults

`func NewListApplicationUserConsentOrganizations200ResponseOrganizationsInnerWithDefaults() *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner`

NewListApplicationUserConsentOrganizations200ResponseOrganizationsInnerWithDefaults instantiates a new ListApplicationUserConsentOrganizations200ResponseOrganizationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCustomData

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.


### GetIsMfaRequired

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetIsMfaRequired() bool`

GetIsMfaRequired returns the IsMfaRequired field if non-nil, zero value otherwise.

### GetIsMfaRequiredOk

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetIsMfaRequiredOk() (*bool, bool)`

GetIsMfaRequiredOk returns a tuple with the IsMfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMfaRequired

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetIsMfaRequired(v bool)`

SetIsMfaRequired sets IsMfaRequired field to given value.


### GetBranding

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetBranding() ListApplicationOrganizations200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding)`

SetBranding sets Branding field to given value.


### GetCreatedAt

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListApplicationUserConsentOrganizations200ResponseOrganizationsInner) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


