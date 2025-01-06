# ListHooks200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Event** | **NullableString** |  | 
**Events** | **[]string** |  | 
**Config** | [**ListHooks200ResponseInnerConfig**](ListHooks200ResponseInnerConfig.md) |  | 
**SigningKey** | **string** |  | 
**Enabled** | **bool** |  | 
**CreatedAt** | **float32** |  | 
**ExecutionStats** | Pointer to [**ListHooks200ResponseInnerExecutionStats**](ListHooks200ResponseInnerExecutionStats.md) |  | [optional] 

## Methods

### NewListHooks200ResponseInner

`func NewListHooks200ResponseInner(tenantId string, id string, name string, event NullableString, events []string, config ListHooks200ResponseInnerConfig, signingKey string, enabled bool, createdAt float32, ) *ListHooks200ResponseInner`

NewListHooks200ResponseInner instantiates a new ListHooks200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHooks200ResponseInnerWithDefaults

`func NewListHooks200ResponseInnerWithDefaults() *ListHooks200ResponseInner`

NewListHooks200ResponseInnerWithDefaults instantiates a new ListHooks200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListHooks200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListHooks200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListHooks200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListHooks200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListHooks200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListHooks200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListHooks200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListHooks200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListHooks200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetEvent

`func (o *ListHooks200ResponseInner) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ListHooks200ResponseInner) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ListHooks200ResponseInner) SetEvent(v string)`

SetEvent sets Event field to given value.


### SetEventNil

`func (o *ListHooks200ResponseInner) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *ListHooks200ResponseInner) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetEvents

`func (o *ListHooks200ResponseInner) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListHooks200ResponseInner) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListHooks200ResponseInner) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetConfig

`func (o *ListHooks200ResponseInner) GetConfig() ListHooks200ResponseInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListHooks200ResponseInner) GetConfigOk() (*ListHooks200ResponseInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListHooks200ResponseInner) SetConfig(v ListHooks200ResponseInnerConfig)`

SetConfig sets Config field to given value.


### GetSigningKey

`func (o *ListHooks200ResponseInner) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *ListHooks200ResponseInner) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *ListHooks200ResponseInner) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.


### GetEnabled

`func (o *ListHooks200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListHooks200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListHooks200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedAt

`func (o *ListHooks200ResponseInner) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListHooks200ResponseInner) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListHooks200ResponseInner) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetExecutionStats

`func (o *ListHooks200ResponseInner) GetExecutionStats() ListHooks200ResponseInnerExecutionStats`

GetExecutionStats returns the ExecutionStats field if non-nil, zero value otherwise.

### GetExecutionStatsOk

`func (o *ListHooks200ResponseInner) GetExecutionStatsOk() (*ListHooks200ResponseInnerExecutionStats, bool)`

GetExecutionStatsOk returns a tuple with the ExecutionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStats

`func (o *ListHooks200ResponseInner) SetExecutionStats(v ListHooks200ResponseInnerExecutionStats)`

SetExecutionStats sets ExecutionStats field to given value.

### HasExecutionStats

`func (o *ListHooks200ResponseInner) HasExecutionStats() bool`

HasExecutionStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


