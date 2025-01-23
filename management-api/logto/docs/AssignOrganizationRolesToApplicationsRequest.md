# AssignOrganizationRolesToApplicationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationIds** | **[]string** | An array of application IDs to assign roles to. | 
**OrganizationRoleIds** | **[]string** | An array of organization role IDs to assign to the applications. | 

## Methods

### NewAssignOrganizationRolesToApplicationsRequest

`func NewAssignOrganizationRolesToApplicationsRequest(applicationIds []string, organizationRoleIds []string, ) *AssignOrganizationRolesToApplicationsRequest`

NewAssignOrganizationRolesToApplicationsRequest instantiates a new AssignOrganizationRolesToApplicationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignOrganizationRolesToApplicationsRequestWithDefaults

`func NewAssignOrganizationRolesToApplicationsRequestWithDefaults() *AssignOrganizationRolesToApplicationsRequest`

NewAssignOrganizationRolesToApplicationsRequestWithDefaults instantiates a new AssignOrganizationRolesToApplicationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationIds

`func (o *AssignOrganizationRolesToApplicationsRequest) GetApplicationIds() []string`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *AssignOrganizationRolesToApplicationsRequest) GetApplicationIdsOk() (*[]string, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *AssignOrganizationRolesToApplicationsRequest) SetApplicationIds(v []string)`

SetApplicationIds sets ApplicationIds field to given value.


### GetOrganizationRoleIds

`func (o *AssignOrganizationRolesToApplicationsRequest) GetOrganizationRoleIds() []string`

GetOrganizationRoleIds returns the OrganizationRoleIds field if non-nil, zero value otherwise.

### GetOrganizationRoleIdsOk

`func (o *AssignOrganizationRolesToApplicationsRequest) GetOrganizationRoleIdsOk() (*[]string, bool)`

GetOrganizationRoleIdsOk returns a tuple with the OrganizationRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoleIds

`func (o *AssignOrganizationRolesToApplicationsRequest) SetOrganizationRoleIds(v []string)`

SetOrganizationRoleIds sets OrganizationRoleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


