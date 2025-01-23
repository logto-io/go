# UpdateOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The updated name of the organization. | [optional] 
**Description** | Pointer to **NullableString** | The updated description of the organization. | [optional] 
**CustomData** | Pointer to **map[string]interface{}** | arbitrary | [optional] 
**IsMfaRequired** | Pointer to **bool** |  | [optional] 
**Branding** | Pointer to [**ListApplicationOrganizations200ResponseInnerBranding**](ListApplicationOrganizations200ResponseInnerBranding.md) |  | [optional] 
**CreatedAt** | Pointer to **float32** |  | [optional] 

## Methods

### NewUpdateOrganizationRequest

`func NewUpdateOrganizationRequest() *UpdateOrganizationRequest`

NewUpdateOrganizationRequest instantiates a new UpdateOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationRequestWithDefaults

`func NewUpdateOrganizationRequestWithDefaults() *UpdateOrganizationRequest`

NewUpdateOrganizationRequestWithDefaults instantiates a new UpdateOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *UpdateOrganizationRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateOrganizationRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateOrganizationRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UpdateOrganizationRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetId

`func (o *UpdateOrganizationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOrganizationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOrganizationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateOrganizationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateOrganizationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOrganizationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateOrganizationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOrganizationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOrganizationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOrganizationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateOrganizationRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateOrganizationRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCustomData

`func (o *UpdateOrganizationRequest) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *UpdateOrganizationRequest) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *UpdateOrganizationRequest) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *UpdateOrganizationRequest) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetIsMfaRequired

`func (o *UpdateOrganizationRequest) GetIsMfaRequired() bool`

GetIsMfaRequired returns the IsMfaRequired field if non-nil, zero value otherwise.

### GetIsMfaRequiredOk

`func (o *UpdateOrganizationRequest) GetIsMfaRequiredOk() (*bool, bool)`

GetIsMfaRequiredOk returns a tuple with the IsMfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMfaRequired

`func (o *UpdateOrganizationRequest) SetIsMfaRequired(v bool)`

SetIsMfaRequired sets IsMfaRequired field to given value.

### HasIsMfaRequired

`func (o *UpdateOrganizationRequest) HasIsMfaRequired() bool`

HasIsMfaRequired returns a boolean if a field has been set.

### GetBranding

`func (o *UpdateOrganizationRequest) GetBranding() ListApplicationOrganizations200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *UpdateOrganizationRequest) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *UpdateOrganizationRequest) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *UpdateOrganizationRequest) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UpdateOrganizationRequest) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateOrganizationRequest) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateOrganizationRequest) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UpdateOrganizationRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


