# CreateOrganizationRoleScopeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationScopeIds** | **[]string** | An array of organization scope IDs to be assigned. Existed scope IDs assignments will be ignored. | 

## Methods

### NewCreateOrganizationRoleScopeRequest

`func NewCreateOrganizationRoleScopeRequest(organizationScopeIds []string, ) *CreateOrganizationRoleScopeRequest`

NewCreateOrganizationRoleScopeRequest instantiates a new CreateOrganizationRoleScopeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationRoleScopeRequestWithDefaults

`func NewCreateOrganizationRoleScopeRequestWithDefaults() *CreateOrganizationRoleScopeRequest`

NewCreateOrganizationRoleScopeRequestWithDefaults instantiates a new CreateOrganizationRoleScopeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationScopeIds

`func (o *CreateOrganizationRoleScopeRequest) GetOrganizationScopeIds() []string`

GetOrganizationScopeIds returns the OrganizationScopeIds field if non-nil, zero value otherwise.

### GetOrganizationScopeIdsOk

`func (o *CreateOrganizationRoleScopeRequest) GetOrganizationScopeIdsOk() (*[]string, bool)`

GetOrganizationScopeIdsOk returns a tuple with the OrganizationScopeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationScopeIds

`func (o *CreateOrganizationRoleScopeRequest) SetOrganizationScopeIds(v []string)`

SetOrganizationScopeIds sets OrganizationScopeIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


