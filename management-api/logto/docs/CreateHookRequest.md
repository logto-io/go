# CreateHookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of the hook. | [optional] 
**Event** | Pointer to **string** | Use &#x60;events&#x60; instead. | [optional] 
**Events** | Pointer to **[]string** | An array of hook events. | [optional] 
**Config** | [**CreateHookRequestConfig**](CreateHookRequestConfig.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **float32** |  | [optional] 

## Methods

### NewCreateHookRequest

`func NewCreateHookRequest(config CreateHookRequestConfig, ) *CreateHookRequest`

NewCreateHookRequest instantiates a new CreateHookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHookRequestWithDefaults

`func NewCreateHookRequestWithDefaults() *CreateHookRequest`

NewCreateHookRequestWithDefaults instantiates a new CreateHookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateHookRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateHookRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateHookRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateHookRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *CreateHookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateHookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateHookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateHookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEvent

`func (o *CreateHookRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CreateHookRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CreateHookRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CreateHookRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEvents

`func (o *CreateHookRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateHookRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateHookRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CreateHookRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetConfig

`func (o *CreateHookRequest) GetConfig() CreateHookRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateHookRequest) GetConfigOk() (*CreateHookRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateHookRequest) SetConfig(v CreateHookRequestConfig)`

SetConfig sets Config field to given value.


### GetEnabled

`func (o *CreateHookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateHookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateHookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateHookRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateHookRequest) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateHookRequest) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateHookRequest) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateHookRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


