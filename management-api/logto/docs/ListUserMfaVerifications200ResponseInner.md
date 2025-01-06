# ListUserMfaVerifications200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **string** |  | 
**Type** | **string** |  | 
**Agent** | Pointer to **string** |  | [optional] 
**RemainCodes** | Pointer to **float32** |  | [optional] 

## Methods

### NewListUserMfaVerifications200ResponseInner

`func NewListUserMfaVerifications200ResponseInner(id string, createdAt string, type_ string, ) *ListUserMfaVerifications200ResponseInner`

NewListUserMfaVerifications200ResponseInner instantiates a new ListUserMfaVerifications200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserMfaVerifications200ResponseInnerWithDefaults

`func NewListUserMfaVerifications200ResponseInnerWithDefaults() *ListUserMfaVerifications200ResponseInner`

NewListUserMfaVerifications200ResponseInnerWithDefaults instantiates a new ListUserMfaVerifications200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListUserMfaVerifications200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListUserMfaVerifications200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListUserMfaVerifications200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListUserMfaVerifications200ResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListUserMfaVerifications200ResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListUserMfaVerifications200ResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetType

`func (o *ListUserMfaVerifications200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListUserMfaVerifications200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListUserMfaVerifications200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetAgent

`func (o *ListUserMfaVerifications200ResponseInner) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *ListUserMfaVerifications200ResponseInner) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *ListUserMfaVerifications200ResponseInner) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *ListUserMfaVerifications200ResponseInner) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetRemainCodes

`func (o *ListUserMfaVerifications200ResponseInner) GetRemainCodes() float32`

GetRemainCodes returns the RemainCodes field if non-nil, zero value otherwise.

### GetRemainCodesOk

`func (o *ListUserMfaVerifications200ResponseInner) GetRemainCodesOk() (*float32, bool)`

GetRemainCodesOk returns a tuple with the RemainCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainCodes

`func (o *ListUserMfaVerifications200ResponseInner) SetRemainCodes(v float32)`

SetRemainCodes sets RemainCodes field to given value.

### HasRemainCodes

`func (o *ListUserMfaVerifications200ResponseInner) HasRemainCodes() bool`

HasRemainCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


