# CreateOrganizationRoleResourceScopeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScopeIds** | **[]string** | An array of resource scope IDs to be assigned. Existed scope IDs assignments will be ignored. | 

## Methods

### NewCreateOrganizationRoleResourceScopeRequest

`func NewCreateOrganizationRoleResourceScopeRequest(scopeIds []string, ) *CreateOrganizationRoleResourceScopeRequest`

NewCreateOrganizationRoleResourceScopeRequest instantiates a new CreateOrganizationRoleResourceScopeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationRoleResourceScopeRequestWithDefaults

`func NewCreateOrganizationRoleResourceScopeRequestWithDefaults() *CreateOrganizationRoleResourceScopeRequest`

NewCreateOrganizationRoleResourceScopeRequestWithDefaults instantiates a new CreateOrganizationRoleResourceScopeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopeIds

`func (o *CreateOrganizationRoleResourceScopeRequest) GetScopeIds() []string`

GetScopeIds returns the ScopeIds field if non-nil, zero value otherwise.

### GetScopeIdsOk

`func (o *CreateOrganizationRoleResourceScopeRequest) GetScopeIdsOk() (*[]string, bool)`

GetScopeIdsOk returns a tuple with the ScopeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeIds

`func (o *CreateOrganizationRoleResourceScopeRequest) SetScopeIds(v []string)`

SetScopeIds sets ScopeIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


