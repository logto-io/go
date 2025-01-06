# CreateSsoConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | arbitrary | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Branding** | Pointer to [**ListOrganizationJitSsoConnectors200ResponseInnerBranding**](ListOrganizationJitSsoConnectors200ResponseInnerBranding.md) |  | [optional] 
**SyncProfile** | Pointer to **bool** |  | [optional] 
**ProviderName** | **string** |  | 
**ConnectorName** | **string** |  | 

## Methods

### NewCreateSsoConnectorRequest

`func NewCreateSsoConnectorRequest(providerName string, connectorName string, ) *CreateSsoConnectorRequest`

NewCreateSsoConnectorRequest instantiates a new CreateSsoConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSsoConnectorRequestWithDefaults

`func NewCreateSsoConnectorRequestWithDefaults() *CreateSsoConnectorRequest`

NewCreateSsoConnectorRequestWithDefaults instantiates a new CreateSsoConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CreateSsoConnectorRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateSsoConnectorRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateSsoConnectorRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateSsoConnectorRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDomains

`func (o *CreateSsoConnectorRequest) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CreateSsoConnectorRequest) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CreateSsoConnectorRequest) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CreateSsoConnectorRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetBranding

`func (o *CreateSsoConnectorRequest) GetBranding() ListOrganizationJitSsoConnectors200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *CreateSsoConnectorRequest) GetBrandingOk() (*ListOrganizationJitSsoConnectors200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *CreateSsoConnectorRequest) SetBranding(v ListOrganizationJitSsoConnectors200ResponseInnerBranding)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *CreateSsoConnectorRequest) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetSyncProfile

`func (o *CreateSsoConnectorRequest) GetSyncProfile() bool`

GetSyncProfile returns the SyncProfile field if non-nil, zero value otherwise.

### GetSyncProfileOk

`func (o *CreateSsoConnectorRequest) GetSyncProfileOk() (*bool, bool)`

GetSyncProfileOk returns a tuple with the SyncProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProfile

`func (o *CreateSsoConnectorRequest) SetSyncProfile(v bool)`

SetSyncProfile sets SyncProfile field to given value.

### HasSyncProfile

`func (o *CreateSsoConnectorRequest) HasSyncProfile() bool`

HasSyncProfile returns a boolean if a field has been set.

### GetProviderName

`func (o *CreateSsoConnectorRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *CreateSsoConnectorRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *CreateSsoConnectorRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetConnectorName

`func (o *CreateSsoConnectorRequest) GetConnectorName() string`

GetConnectorName returns the ConnectorName field if non-nil, zero value otherwise.

### GetConnectorNameOk

`func (o *CreateSsoConnectorRequest) GetConnectorNameOk() (*string, bool)`

GetConnectorNameOk returns a tuple with the ConnectorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorName

`func (o *CreateSsoConnectorRequest) SetConnectorName(v string)`

SetConnectorName sets ConnectorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


