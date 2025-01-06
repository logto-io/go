# AssignOrganizationRolesToUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationRoleIds** | **[]string** | An array of organization role IDs to assign to the user. User existed roles assignment will be ignored. | 

## Methods

### NewAssignOrganizationRolesToUserRequest

`func NewAssignOrganizationRolesToUserRequest(organizationRoleIds []string, ) *AssignOrganizationRolesToUserRequest`

NewAssignOrganizationRolesToUserRequest instantiates a new AssignOrganizationRolesToUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignOrganizationRolesToUserRequestWithDefaults

`func NewAssignOrganizationRolesToUserRequestWithDefaults() *AssignOrganizationRolesToUserRequest`

NewAssignOrganizationRolesToUserRequestWithDefaults instantiates a new AssignOrganizationRolesToUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationRoleIds

`func (o *AssignOrganizationRolesToUserRequest) GetOrganizationRoleIds() []string`

GetOrganizationRoleIds returns the OrganizationRoleIds field if non-nil, zero value otherwise.

### GetOrganizationRoleIdsOk

`func (o *AssignOrganizationRolesToUserRequest) GetOrganizationRoleIdsOk() (*[]string, bool)`

GetOrganizationRoleIdsOk returns a tuple with the OrganizationRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoleIds

`func (o *AssignOrganizationRolesToUserRequest) SetOrganizationRoleIds(v []string)`

SetOrganizationRoleIds sets OrganizationRoleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


