# UpdateSsoConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | arbitrary | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Branding** | Pointer to [**ListOrganizationJitSsoConnectors200ResponseInnerBranding**](ListOrganizationJitSsoConnectors200ResponseInnerBranding.md) |  | [optional] 
**SyncProfile** | Pointer to **bool** |  | [optional] 
**ConnectorName** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateSsoConnectorRequest

`func NewUpdateSsoConnectorRequest() *UpdateSsoConnectorRequest`

NewUpdateSsoConnectorRequest instantiates a new UpdateSsoConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSsoConnectorRequestWithDefaults

`func NewUpdateSsoConnectorRequestWithDefaults() *UpdateSsoConnectorRequest`

NewUpdateSsoConnectorRequestWithDefaults instantiates a new UpdateSsoConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *UpdateSsoConnectorRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateSsoConnectorRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateSsoConnectorRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateSsoConnectorRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDomains

`func (o *UpdateSsoConnectorRequest) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *UpdateSsoConnectorRequest) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *UpdateSsoConnectorRequest) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *UpdateSsoConnectorRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetBranding

`func (o *UpdateSsoConnectorRequest) GetBranding() ListOrganizationJitSsoConnectors200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *UpdateSsoConnectorRequest) GetBrandingOk() (*ListOrganizationJitSsoConnectors200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *UpdateSsoConnectorRequest) SetBranding(v ListOrganizationJitSsoConnectors200ResponseInnerBranding)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *UpdateSsoConnectorRequest) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetSyncProfile

`func (o *UpdateSsoConnectorRequest) GetSyncProfile() bool`

GetSyncProfile returns the SyncProfile field if non-nil, zero value otherwise.

### GetSyncProfileOk

`func (o *UpdateSsoConnectorRequest) GetSyncProfileOk() (*bool, bool)`

GetSyncProfileOk returns a tuple with the SyncProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProfile

`func (o *UpdateSsoConnectorRequest) SetSyncProfile(v bool)`

SetSyncProfile sets SyncProfile field to given value.

### HasSyncProfile

`func (o *UpdateSsoConnectorRequest) HasSyncProfile() bool`

HasSyncProfile returns a boolean if a field has been set.

### GetConnectorName

`func (o *UpdateSsoConnectorRequest) GetConnectorName() string`

GetConnectorName returns the ConnectorName field if non-nil, zero value otherwise.

### GetConnectorNameOk

`func (o *UpdateSsoConnectorRequest) GetConnectorNameOk() (*string, bool)`

GetConnectorNameOk returns a tuple with the ConnectorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorName

`func (o *UpdateSsoConnectorRequest) SetConnectorName(v string)`

SetConnectorName sets ConnectorName field to given value.

### HasConnectorName

`func (o *UpdateSsoConnectorRequest) HasConnectorName() bool`

HasConnectorName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


