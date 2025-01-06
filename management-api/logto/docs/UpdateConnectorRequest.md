# UpdateConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | The connector config object that will be passed to the connector. The config object should be compatible with the connector factory. | [optional] 
**Metadata** | Pointer to [**UpdateConnectorRequestMetadata**](UpdateConnectorRequestMetadata.md) |  | [optional] 
**SyncProfile** | Pointer to **bool** | Whether to sync user profile from the identity provider to Logto at each sign-in. If &#x60;false&#x60;, the user profile will only be synced when the user is created. | [optional] 

## Methods

### NewUpdateConnectorRequest

`func NewUpdateConnectorRequest() *UpdateConnectorRequest`

NewUpdateConnectorRequest instantiates a new UpdateConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConnectorRequestWithDefaults

`func NewUpdateConnectorRequestWithDefaults() *UpdateConnectorRequest`

NewUpdateConnectorRequestWithDefaults instantiates a new UpdateConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *UpdateConnectorRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateConnectorRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateConnectorRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateConnectorRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateConnectorRequest) GetMetadata() UpdateConnectorRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateConnectorRequest) GetMetadataOk() (*UpdateConnectorRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateConnectorRequest) SetMetadata(v UpdateConnectorRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateConnectorRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSyncProfile

`func (o *UpdateConnectorRequest) GetSyncProfile() bool`

GetSyncProfile returns the SyncProfile field if non-nil, zero value otherwise.

### GetSyncProfileOk

`func (o *UpdateConnectorRequest) GetSyncProfileOk() (*bool, bool)`

GetSyncProfileOk returns a tuple with the SyncProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProfile

`func (o *UpdateConnectorRequest) SetSyncProfile(v bool)`

SetSyncProfile sets SyncProfile field to given value.

### HasSyncProfile

`func (o *UpdateConnectorRequest) HasSyncProfile() bool`

HasSyncProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


