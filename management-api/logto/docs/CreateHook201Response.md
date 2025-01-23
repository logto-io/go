# CreateHook201Response

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

## Methods

### NewCreateHook201Response

`func NewCreateHook201Response(tenantId string, id string, name string, event NullableString, events []string, config ListHooks200ResponseInnerConfig, signingKey string, enabled bool, createdAt float32, ) *CreateHook201Response`

NewCreateHook201Response instantiates a new CreateHook201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHook201ResponseWithDefaults

`func NewCreateHook201ResponseWithDefaults() *CreateHook201Response`

NewCreateHook201ResponseWithDefaults instantiates a new CreateHook201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateHook201Response) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateHook201Response) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateHook201Response) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *CreateHook201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateHook201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateHook201Response) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateHook201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateHook201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateHook201Response) SetName(v string)`

SetName sets Name field to given value.


### GetEvent

`func (o *CreateHook201Response) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CreateHook201Response) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CreateHook201Response) SetEvent(v string)`

SetEvent sets Event field to given value.


### SetEventNil

`func (o *CreateHook201Response) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *CreateHook201Response) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetEvents

`func (o *CreateHook201Response) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateHook201Response) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateHook201Response) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetConfig

`func (o *CreateHook201Response) GetConfig() ListHooks200ResponseInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateHook201Response) GetConfigOk() (*ListHooks200ResponseInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateHook201Response) SetConfig(v ListHooks200ResponseInnerConfig)`

SetConfig sets Config field to given value.


### GetSigningKey

`func (o *CreateHook201Response) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *CreateHook201Response) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *CreateHook201Response) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.


### GetEnabled

`func (o *CreateHook201Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateHook201Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateHook201Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedAt

`func (o *CreateHook201Response) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateHook201Response) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateHook201Response) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


