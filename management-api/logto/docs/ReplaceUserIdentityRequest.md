# ReplaceUserIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The user&#39;s social identity ID. | 
**Details** | Pointer to **map[string]interface{}** | The user&#39;s social identity details. | [optional] 

## Methods

### NewReplaceUserIdentityRequest

`func NewReplaceUserIdentityRequest(userId string, ) *ReplaceUserIdentityRequest`

NewReplaceUserIdentityRequest instantiates a new ReplaceUserIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceUserIdentityRequestWithDefaults

`func NewReplaceUserIdentityRequestWithDefaults() *ReplaceUserIdentityRequest`

NewReplaceUserIdentityRequestWithDefaults instantiates a new ReplaceUserIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ReplaceUserIdentityRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ReplaceUserIdentityRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ReplaceUserIdentityRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetDetails

`func (o *ReplaceUserIdentityRequest) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ReplaceUserIdentityRequest) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ReplaceUserIdentityRequest) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ReplaceUserIdentityRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


