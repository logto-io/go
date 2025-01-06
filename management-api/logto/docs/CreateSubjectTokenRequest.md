# CreateSubjectTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The ID of the user to impersonate. | 
**Context** | Pointer to **map[string]interface{}** | The additional context to be included in the token, this can be used in custom JWT. | [optional] 

## Methods

### NewCreateSubjectTokenRequest

`func NewCreateSubjectTokenRequest(userId string, ) *CreateSubjectTokenRequest`

NewCreateSubjectTokenRequest instantiates a new CreateSubjectTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubjectTokenRequestWithDefaults

`func NewCreateSubjectTokenRequestWithDefaults() *CreateSubjectTokenRequest`

NewCreateSubjectTokenRequestWithDefaults instantiates a new CreateSubjectTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateSubjectTokenRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateSubjectTokenRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateSubjectTokenRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetContext

`func (o *CreateSubjectTokenRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CreateSubjectTokenRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CreateSubjectTokenRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *CreateSubjectTokenRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


