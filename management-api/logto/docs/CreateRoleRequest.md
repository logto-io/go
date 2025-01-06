# CreateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Name** | **string** | The name of the role. It should be unique within the tenant. | 
**Description** | **string** |  | 
**Type** | Pointer to **string** | The type of the role. It cannot be changed after creation. | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**ScopeIds** | Pointer to **[]string** | The initial API resource scopes assigned to the role. | [optional] 

## Methods

### NewCreateRoleRequest

`func NewCreateRoleRequest(name string, description string, ) *CreateRoleRequest`

NewCreateRoleRequest instantiates a new CreateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleRequestWithDefaults

`func NewCreateRoleRequestWithDefaults() *CreateRoleRequest`

NewCreateRoleRequestWithDefaults instantiates a new CreateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateRoleRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateRoleRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateRoleRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateRoleRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *CreateRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *CreateRoleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRoleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRoleRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateRoleRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsDefault

`func (o *CreateRoleRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreateRoleRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreateRoleRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CreateRoleRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetScopeIds

`func (o *CreateRoleRequest) GetScopeIds() []string`

GetScopeIds returns the ScopeIds field if non-nil, zero value otherwise.

### GetScopeIdsOk

`func (o *CreateRoleRequest) GetScopeIdsOk() (*[]string, bool)`

GetScopeIdsOk returns a tuple with the ScopeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeIds

`func (o *CreateRoleRequest) SetScopeIds(v []string)`

SetScopeIds sets ScopeIds field to given value.

### HasScopeIds

`func (o *CreateRoleRequest) HasScopeIds() bool`

HasScopeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


