# ListSsoConnectors200ResponseInner

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
**Name** | **string** |  | 
**ProviderType** | **string** |  | 
**ProviderLogo** | **string** |  | 
**ProviderLogoDark** | **string** |  | 
**ProviderConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewListSsoConnectors200ResponseInner

`func NewListSsoConnectors200ResponseInner(tenantId string, id string, providerName string, connectorName string, config map[string]interface{}, domains []string, branding ListOrganizationJitSsoConnectors200ResponseInnerBranding, syncProfile bool, createdAt float32, name string, providerType string, providerLogo string, providerLogoDark string, ) *ListSsoConnectors200ResponseInner`

NewListSsoConnectors200ResponseInner instantiates a new ListSsoConnectors200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSsoConnectors200ResponseInnerWithDefaults

`func NewListSsoConnectors200ResponseInnerWithDefaults() *ListSsoConnectors200ResponseInner`

NewListSsoConnectors200ResponseInnerWithDefaults instantiates a new ListSsoConnectors200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListSsoConnectors200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListSsoConnectors200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListSsoConnectors200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListSsoConnectors200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSsoConnectors200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSsoConnectors200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetProviderName

`func (o *ListSsoConnectors200ResponseInner) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ListSsoConnectors200ResponseInner) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ListSsoConnectors200ResponseInner) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetConnectorName

`func (o *ListSsoConnectors200ResponseInner) GetConnectorName() string`

GetConnectorName returns the ConnectorName field if non-nil, zero value otherwise.

### GetConnectorNameOk

`func (o *ListSsoConnectors200ResponseInner) GetConnectorNameOk() (*string, bool)`

GetConnectorNameOk returns a tuple with the ConnectorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorName

`func (o *ListSsoConnectors200ResponseInner) SetConnectorName(v string)`

SetConnectorName sets ConnectorName field to given value.


### GetConfig

`func (o *ListSsoConnectors200ResponseInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListSsoConnectors200ResponseInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListSsoConnectors200ResponseInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetDomains

`func (o *ListSsoConnectors200ResponseInner) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ListSsoConnectors200ResponseInner) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ListSsoConnectors200ResponseInner) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### GetBranding

`func (o *ListSsoConnectors200ResponseInner) GetBranding() ListOrganizationJitSsoConnectors200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *ListSsoConnectors200ResponseInner) GetBrandingOk() (*ListOrganizationJitSsoConnectors200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *ListSsoConnectors200ResponseInner) SetBranding(v ListOrganizationJitSsoConnectors200ResponseInnerBranding)`

SetBranding sets Branding field to given value.


### GetSyncProfile

`func (o *ListSsoConnectors200ResponseInner) GetSyncProfile() bool`

GetSyncProfile returns the SyncProfile field if non-nil, zero value otherwise.

### GetSyncProfileOk

`func (o *ListSsoConnectors200ResponseInner) GetSyncProfileOk() (*bool, bool)`

GetSyncProfileOk returns a tuple with the SyncProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProfile

`func (o *ListSsoConnectors200ResponseInner) SetSyncProfile(v bool)`

SetSyncProfile sets SyncProfile field to given value.


### GetCreatedAt

`func (o *ListSsoConnectors200ResponseInner) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListSsoConnectors200ResponseInner) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListSsoConnectors200ResponseInner) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *ListSsoConnectors200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSsoConnectors200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSsoConnectors200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetProviderType

`func (o *ListSsoConnectors200ResponseInner) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *ListSsoConnectors200ResponseInner) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *ListSsoConnectors200ResponseInner) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetProviderLogo

`func (o *ListSsoConnectors200ResponseInner) GetProviderLogo() string`

GetProviderLogo returns the ProviderLogo field if non-nil, zero value otherwise.

### GetProviderLogoOk

`func (o *ListSsoConnectors200ResponseInner) GetProviderLogoOk() (*string, bool)`

GetProviderLogoOk returns a tuple with the ProviderLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderLogo

`func (o *ListSsoConnectors200ResponseInner) SetProviderLogo(v string)`

SetProviderLogo sets ProviderLogo field to given value.


### GetProviderLogoDark

`func (o *ListSsoConnectors200ResponseInner) GetProviderLogoDark() string`

GetProviderLogoDark returns the ProviderLogoDark field if non-nil, zero value otherwise.

### GetProviderLogoDarkOk

`func (o *ListSsoConnectors200ResponseInner) GetProviderLogoDarkOk() (*string, bool)`

GetProviderLogoDarkOk returns a tuple with the ProviderLogoDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderLogoDark

`func (o *ListSsoConnectors200ResponseInner) SetProviderLogoDark(v string)`

SetProviderLogoDark sets ProviderLogoDark field to given value.


### GetProviderConfig

`func (o *ListSsoConnectors200ResponseInner) GetProviderConfig() map[string]interface{}`

GetProviderConfig returns the ProviderConfig field if non-nil, zero value otherwise.

### GetProviderConfigOk

`func (o *ListSsoConnectors200ResponseInner) GetProviderConfigOk() (*map[string]interface{}, bool)`

GetProviderConfigOk returns a tuple with the ProviderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfig

`func (o *ListSsoConnectors200ResponseInner) SetProviderConfig(v map[string]interface{})`

SetProviderConfig sets ProviderConfig field to given value.

### HasProviderConfig

`func (o *ListSsoConnectors200ResponseInner) HasProviderConfig() bool`

HasProviderConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


