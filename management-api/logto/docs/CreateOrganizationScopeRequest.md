# CreateOrganizationScopeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Name** | **string** | The name of the organization scope. It must be unique within the organization template. | 
**Description** | Pointer to **NullableString** | The description of the organization scope. | [optional] 

## Methods

### NewCreateOrganizationScopeRequest

`func NewCreateOrganizationScopeRequest(name string, ) *CreateOrganizationScopeRequest`

NewCreateOrganizationScopeRequest instantiates a new CreateOrganizationScopeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationScopeRequestWithDefaults

`func NewCreateOrganizationScopeRequestWithDefaults() *CreateOrganizationScopeRequest`

NewCreateOrganizationScopeRequestWithDefaults instantiates a new CreateOrganizationScopeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateOrganizationScopeRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateOrganizationScopeRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateOrganizationScopeRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateOrganizationScopeRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *CreateOrganizationScopeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationScopeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationScopeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateOrganizationScopeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrganizationScopeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrganizationScopeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrganizationScopeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateOrganizationScopeRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateOrganizationScopeRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


