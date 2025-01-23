# CreateOrganizationRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Name** | **string** | The name of the organization role. It must be unique within the organization template. | 
**Description** | Pointer to **NullableString** | The description of the organization role. | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**OrganizationScopeIds** | **[]string** | An array of organization scope IDs to be assigned to the organization role. | [default to []]
**ResourceScopeIds** | **[]string** | An array of resource scope IDs to be assigned to the organization role. | [default to []]

## Methods

### NewCreateOrganizationRoleRequest

`func NewCreateOrganizationRoleRequest(name string, organizationScopeIds []string, resourceScopeIds []string, ) *CreateOrganizationRoleRequest`

NewCreateOrganizationRoleRequest instantiates a new CreateOrganizationRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationRoleRequestWithDefaults

`func NewCreateOrganizationRoleRequestWithDefaults() *CreateOrganizationRoleRequest`

NewCreateOrganizationRoleRequestWithDefaults instantiates a new CreateOrganizationRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateOrganizationRoleRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateOrganizationRoleRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateOrganizationRoleRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateOrganizationRoleRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *CreateOrganizationRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateOrganizationRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrganizationRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrganizationRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrganizationRoleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateOrganizationRoleRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateOrganizationRoleRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *CreateOrganizationRoleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrganizationRoleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrganizationRoleRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateOrganizationRoleRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrganizationScopeIds

`func (o *CreateOrganizationRoleRequest) GetOrganizationScopeIds() []string`

GetOrganizationScopeIds returns the OrganizationScopeIds field if non-nil, zero value otherwise.

### GetOrganizationScopeIdsOk

`func (o *CreateOrganizationRoleRequest) GetOrganizationScopeIdsOk() (*[]string, bool)`

GetOrganizationScopeIdsOk returns a tuple with the OrganizationScopeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationScopeIds

`func (o *CreateOrganizationRoleRequest) SetOrganizationScopeIds(v []string)`

SetOrganizationScopeIds sets OrganizationScopeIds field to given value.


### GetResourceScopeIds

`func (o *CreateOrganizationRoleRequest) GetResourceScopeIds() []string`

GetResourceScopeIds returns the ResourceScopeIds field if non-nil, zero value otherwise.

### GetResourceScopeIdsOk

`func (o *CreateOrganizationRoleRequest) GetResourceScopeIdsOk() (*[]string, bool)`

GetResourceScopeIdsOk returns a tuple with the ResourceScopeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceScopeIds

`func (o *CreateOrganizationRoleRequest) SetResourceScopeIds(v []string)`

SetResourceScopeIds sets ResourceScopeIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


