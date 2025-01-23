# AssignApplicationRolesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleIds** | **[]string** | An array of API resource role IDs to assign. | 

## Methods

### NewAssignApplicationRolesRequest

`func NewAssignApplicationRolesRequest(roleIds []string, ) *AssignApplicationRolesRequest`

NewAssignApplicationRolesRequest instantiates a new AssignApplicationRolesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignApplicationRolesRequestWithDefaults

`func NewAssignApplicationRolesRequestWithDefaults() *AssignApplicationRolesRequest`

NewAssignApplicationRolesRequestWithDefaults instantiates a new AssignApplicationRolesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleIds

`func (o *AssignApplicationRolesRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *AssignApplicationRolesRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *AssignApplicationRolesRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


