# UpdateUserCustomDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomData** | **map[string]interface{}** | Partial custom data object to update for the given user ID. | 

## Methods

### NewUpdateUserCustomDataRequest

`func NewUpdateUserCustomDataRequest(customData map[string]interface{}, ) *UpdateUserCustomDataRequest`

NewUpdateUserCustomDataRequest instantiates a new UpdateUserCustomDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserCustomDataRequestWithDefaults

`func NewUpdateUserCustomDataRequestWithDefaults() *UpdateUserCustomDataRequest`

NewUpdateUserCustomDataRequestWithDefaults instantiates a new UpdateUserCustomDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomData

`func (o *UpdateUserCustomDataRequest) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *UpdateUserCustomDataRequest) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *UpdateUserCustomDataRequest) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


