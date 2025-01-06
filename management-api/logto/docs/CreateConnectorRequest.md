# CreateConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | The connector config object that will be passed to the connector. The config object should be compatible with the connector factory. | [optional] 
**ConnectorId** | **string** | The connector factory ID for creating the connector. | 
**Metadata** | Pointer to [**CreateConnectorRequestMetadata**](CreateConnectorRequestMetadata.md) |  | [optional] 
**SyncProfile** | Pointer to **bool** | Whether to sync user profile from the identity provider to Logto at each sign-in. If &#x60;false&#x60;, the user profile will only be synced when the user is created. | [optional] 
**Id** | Pointer to **string** | The unique ID for the connector. If not provided, a random ID will be generated. | [optional] 

## Methods

### NewCreateConnectorRequest

`func NewCreateConnectorRequest(connectorId string, ) *CreateConnectorRequest`

NewCreateConnectorRequest instantiates a new CreateConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectorRequestWithDefaults

`func NewCreateConnectorRequestWithDefaults() *CreateConnectorRequest`

NewCreateConnectorRequestWithDefaults instantiates a new CreateConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CreateConnectorRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateConnectorRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateConnectorRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateConnectorRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConnectorId

`func (o *CreateConnectorRequest) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *CreateConnectorRequest) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *CreateConnectorRequest) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.


### GetMetadata

`func (o *CreateConnectorRequest) GetMetadata() CreateConnectorRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateConnectorRequest) GetMetadataOk() (*CreateConnectorRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateConnectorRequest) SetMetadata(v CreateConnectorRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateConnectorRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSyncProfile

`func (o *CreateConnectorRequest) GetSyncProfile() bool`

GetSyncProfile returns the SyncProfile field if non-nil, zero value otherwise.

### GetSyncProfileOk

`func (o *CreateConnectorRequest) GetSyncProfileOk() (*bool, bool)`

GetSyncProfileOk returns a tuple with the SyncProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProfile

`func (o *CreateConnectorRequest) SetSyncProfile(v bool)`

SetSyncProfile sets SyncProfile field to given value.

### HasSyncProfile

`func (o *CreateConnectorRequest) HasSyncProfile() bool`

HasSyncProfile returns a boolean if a field has been set.

### GetId

`func (o *CreateConnectorRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateConnectorRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateConnectorRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateConnectorRequest) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


