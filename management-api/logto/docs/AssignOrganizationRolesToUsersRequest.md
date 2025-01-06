# AssignOrganizationRolesToUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | **[]string** | An array of user IDs to assign roles. | 
**OrganizationRoleIds** | **[]string** | An array of organization role IDs to assign. User existed roles assignment will be ignored. | 

## Methods

### NewAssignOrganizationRolesToUsersRequest

`func NewAssignOrganizationRolesToUsersRequest(userIds []string, organizationRoleIds []string, ) *AssignOrganizationRolesToUsersRequest`

NewAssignOrganizationRolesToUsersRequest instantiates a new AssignOrganizationRolesToUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignOrganizationRolesToUsersRequestWithDefaults

`func NewAssignOrganizationRolesToUsersRequestWithDefaults() *AssignOrganizationRolesToUsersRequest`

NewAssignOrganizationRolesToUsersRequestWithDefaults instantiates a new AssignOrganizationRolesToUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *AssignOrganizationRolesToUsersRequest) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *AssignOrganizationRolesToUsersRequest) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *AssignOrganizationRolesToUsersRequest) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.


### GetOrganizationRoleIds

`func (o *AssignOrganizationRolesToUsersRequest) GetOrganizationRoleIds() []string`

GetOrganizationRoleIds returns the OrganizationRoleIds field if non-nil, zero value otherwise.

### GetOrganizationRoleIdsOk

`func (o *AssignOrganizationRolesToUsersRequest) GetOrganizationRoleIdsOk() (*[]string, bool)`

GetOrganizationRoleIdsOk returns a tuple with the OrganizationRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoleIds

`func (o *AssignOrganizationRolesToUsersRequest) SetOrganizationRoleIds(v []string)`

SetOrganizationRoleIds sets OrganizationRoleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


