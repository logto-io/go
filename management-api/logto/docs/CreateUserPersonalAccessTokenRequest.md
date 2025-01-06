# CreateUserPersonalAccessTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The personal access token name. Must be unique within the user. | 
**ExpiresAt** | Pointer to **NullableFloat32** | The epoch time in milliseconds when the token will expire. If not provided, the token will never expire. | [optional] 

## Methods

### NewCreateUserPersonalAccessTokenRequest

`func NewCreateUserPersonalAccessTokenRequest(name string, ) *CreateUserPersonalAccessTokenRequest`

NewCreateUserPersonalAccessTokenRequest instantiates a new CreateUserPersonalAccessTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserPersonalAccessTokenRequestWithDefaults

`func NewCreateUserPersonalAccessTokenRequestWithDefaults() *CreateUserPersonalAccessTokenRequest`

NewCreateUserPersonalAccessTokenRequestWithDefaults instantiates a new CreateUserPersonalAccessTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUserPersonalAccessTokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserPersonalAccessTokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserPersonalAccessTokenRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExpiresAt

`func (o *CreateUserPersonalAccessTokenRequest) GetExpiresAt() float32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateUserPersonalAccessTokenRequest) GetExpiresAtOk() (*float32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateUserPersonalAccessTokenRequest) SetExpiresAt(v float32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateUserPersonalAccessTokenRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CreateUserPersonalAccessTokenRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreateUserPersonalAccessTokenRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


