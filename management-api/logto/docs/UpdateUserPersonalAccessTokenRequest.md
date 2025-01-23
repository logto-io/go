# UpdateUserPersonalAccessTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The token name to update. Must be unique within the user. | 

## Methods

### NewUpdateUserPersonalAccessTokenRequest

`func NewUpdateUserPersonalAccessTokenRequest(name string, ) *UpdateUserPersonalAccessTokenRequest`

NewUpdateUserPersonalAccessTokenRequest instantiates a new UpdateUserPersonalAccessTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserPersonalAccessTokenRequestWithDefaults

`func NewUpdateUserPersonalAccessTokenRequestWithDefaults() *UpdateUserPersonalAccessTokenRequest`

NewUpdateUserPersonalAccessTokenRequestWithDefaults instantiates a new UpdateUserPersonalAccessTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateUserPersonalAccessTokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserPersonalAccessTokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserPersonalAccessTokenRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


