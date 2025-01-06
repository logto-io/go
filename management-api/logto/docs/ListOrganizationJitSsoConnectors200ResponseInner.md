# ListOrganizationJitSsoConnectors200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**ProviderName** | **string** |  | 
**ConnectorName** | **string** |  | 
**Config** | **map[string]interface{}** | arbitrary | 
**Domains** | **[]string** |  | 
**Branding** | [**ListOrganizationJitSsoConnectors200ResponseInnerBranding**](ListOrganizationJitSsoConnectors200ResponseInnerBranding.md) |  | 
**SyncProfile** | **bool** |  | 
**CreatedAt** | **float32** |  | 

## Methods

### NewListOrganizationJitSsoConnectors200ResponseInner

`func NewListOrganizationJitSsoConnectors200ResponseInner(tenantId string, id string, providerName string, connectorName string, config map[string]interface{}, domains []string, branding ListOrganizationJitSsoConnectors200ResponseInnerBranding, syncProfile bool, createdAt float32, ) *ListOrganizationJitSsoConnectors200ResponseInner`

NewListOrganizationJitSsoConnectors200ResponseInner instantiates a new ListOrganizationJitSsoConnectors200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizationJitSsoConnectors200ResponseInnerWithDefaults

`func NewListOrganizationJitSsoConnectors200ResponseInnerWithDefaults() *ListOrganizationJitSsoConnectors200ResponseInner`

NewListOrganizationJitSsoConnectors200ResponseInnerWithDefaults instantiates a new ListOrganizationJitSsoConnectors200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetProviderName

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetConnectorName

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetConnectorName() string`

GetConnectorName returns the ConnectorName field if non-nil, zero value otherwise.

### GetConnectorNameOk

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetConnectorNameOk() (*string, bool)`

GetConnectorNameOk returns a tuple with the ConnectorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorName

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) SetConnectorName(v string)`

SetConnectorName sets ConnectorName field to given value.


### GetConfig

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetDomains

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### GetBranding

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetBranding() ListOrganizationJitSsoConnectors200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetBrandingOk() (*ListOrganizationJitSsoConnectors200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) SetBranding(v ListOrganizationJitSsoConnectors200ResponseInnerBranding)`

SetBranding sets Branding field to given value.


### GetSyncProfile

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetSyncProfile() bool`

GetSyncProfile returns the SyncProfile field if non-nil, zero value otherwise.

### GetSyncProfileOk

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetSyncProfileOk() (*bool, bool)`

GetSyncProfileOk returns a tuple with the SyncProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProfile

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) SetSyncProfile(v bool)`

SetSyncProfile sets SyncProfile field to given value.


### GetCreatedAt

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListOrganizationJitSsoConnectors200ResponseInner) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


