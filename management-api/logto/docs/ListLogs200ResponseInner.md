# ListLogs200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**Key** | **string** |  | 
**Payload** | [**ListLogs200ResponseInnerPayload**](ListLogs200ResponseInnerPayload.md) |  | 
**CreatedAt** | **float32** |  | 

## Methods

### NewListLogs200ResponseInner

`func NewListLogs200ResponseInner(tenantId string, id string, key string, payload ListLogs200ResponseInnerPayload, createdAt float32, ) *ListLogs200ResponseInner`

NewListLogs200ResponseInner instantiates a new ListLogs200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogs200ResponseInnerWithDefaults

`func NewListLogs200ResponseInnerWithDefaults() *ListLogs200ResponseInner`

NewListLogs200ResponseInnerWithDefaults instantiates a new ListLogs200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListLogs200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListLogs200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListLogs200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListLogs200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLogs200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLogs200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *ListLogs200ResponseInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ListLogs200ResponseInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ListLogs200ResponseInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetPayload

`func (o *ListLogs200ResponseInner) GetPayload() ListLogs200ResponseInnerPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ListLogs200ResponseInner) GetPayloadOk() (*ListLogs200ResponseInnerPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ListLogs200ResponseInner) SetPayload(v ListLogs200ResponseInnerPayload)`

SetPayload sets Payload field to given value.


### GetCreatedAt

`func (o *ListLogs200ResponseInner) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListLogs200ResponseInner) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListLogs200ResponseInner) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


