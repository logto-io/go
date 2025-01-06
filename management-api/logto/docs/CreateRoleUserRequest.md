# CreateRoleUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | **[]string** | An array of user IDs to be assigned. | 

## Methods

### NewCreateRoleUserRequest

`func NewCreateRoleUserRequest(userIds []string, ) *CreateRoleUserRequest`

NewCreateRoleUserRequest instantiates a new CreateRoleUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleUserRequestWithDefaults

`func NewCreateRoleUserRequestWithDefaults() *CreateRoleUserRequest`

NewCreateRoleUserRequestWithDefaults instantiates a new CreateRoleUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *CreateRoleUserRequest) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *CreateRoleUserRequest) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *CreateRoleUserRequest) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


