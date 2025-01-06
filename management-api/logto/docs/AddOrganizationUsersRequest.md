# AddOrganizationUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | **[]string** | An array of user IDs to be added to the organization. Organization existed users assignment will be ignored. | 

## Methods

### NewAddOrganizationUsersRequest

`func NewAddOrganizationUsersRequest(userIds []string, ) *AddOrganizationUsersRequest`

NewAddOrganizationUsersRequest instantiates a new AddOrganizationUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOrganizationUsersRequestWithDefaults

`func NewAddOrganizationUsersRequestWithDefaults() *AddOrganizationUsersRequest`

NewAddOrganizationUsersRequestWithDefaults instantiates a new AddOrganizationUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *AddOrganizationUsersRequest) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *AddOrganizationUsersRequest) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *AddOrganizationUsersRequest) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


