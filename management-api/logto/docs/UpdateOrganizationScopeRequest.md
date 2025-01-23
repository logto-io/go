# UpdateOrganizationScopeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The updated name of the organization scope. It must be unique within the organization template. | [optional] 
**Description** | Pointer to **NullableString** | The updated description of the organization scope. | [optional] 

## Methods

### NewUpdateOrganizationScopeRequest

`func NewUpdateOrganizationScopeRequest() *UpdateOrganizationScopeRequest`

NewUpdateOrganizationScopeRequest instantiates a new UpdateOrganizationScopeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationScopeRequestWithDefaults

`func NewUpdateOrganizationScopeRequestWithDefaults() *UpdateOrganizationScopeRequest`

NewUpdateOrganizationScopeRequestWithDefaults instantiates a new UpdateOrganizationScopeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *UpdateOrganizationScopeRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateOrganizationScopeRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateOrganizationScopeRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UpdateOrganizationScopeRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetId

`func (o *UpdateOrganizationScopeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOrganizationScopeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOrganizationScopeRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateOrganizationScopeRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateOrganizationScopeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationScopeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationScopeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOrganizationScopeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateOrganizationScopeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOrganizationScopeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOrganizationScopeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOrganizationScopeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateOrganizationScopeRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateOrganizationScopeRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


