# ListApplications200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Secret** | **string** | The internal client secret. Note it is only used for internal validation, and the actual secrets should be retrieved from &#x60;/api/applications/{id}/secrets&#x60; endpoints. | 
**Description** | **NullableString** |  | 
**Type** | **string** |  | 
**OidcClientMetadata** | [**ListApplications200ResponseInnerOidcClientMetadata**](ListApplications200ResponseInnerOidcClientMetadata.md) |  | 
**CustomClientMetadata** | [**ListApplications200ResponseInnerCustomClientMetadata**](ListApplications200ResponseInnerCustomClientMetadata.md) |  | 
**ProtectedAppMetadata** | [**NullableListApplications200ResponseInnerProtectedAppMetadata**](ListApplications200ResponseInnerProtectedAppMetadata.md) |  | 
**CustomData** | **map[string]interface{}** | arbitrary | 
**IsThirdParty** | **bool** |  | 
**CreatedAt** | **float32** |  | 

## Methods

### NewListApplications200ResponseInner

`func NewListApplications200ResponseInner(tenantId string, id string, name string, secret string, description NullableString, type_ string, oidcClientMetadata ListApplications200ResponseInnerOidcClientMetadata, customClientMetadata ListApplications200ResponseInnerCustomClientMetadata, protectedAppMetadata NullableListApplications200ResponseInnerProtectedAppMetadata, customData map[string]interface{}, isThirdParty bool, createdAt float32, ) *ListApplications200ResponseInner`

NewListApplications200ResponseInner instantiates a new ListApplications200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApplications200ResponseInnerWithDefaults

`func NewListApplications200ResponseInnerWithDefaults() *ListApplications200ResponseInner`

NewListApplications200ResponseInnerWithDefaults instantiates a new ListApplications200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListApplications200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListApplications200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListApplications200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListApplications200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListApplications200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListApplications200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListApplications200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListApplications200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListApplications200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetSecret

`func (o *ListApplications200ResponseInner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ListApplications200ResponseInner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ListApplications200ResponseInner) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetDescription

`func (o *ListApplications200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListApplications200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListApplications200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ListApplications200ResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListApplications200ResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *ListApplications200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListApplications200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListApplications200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetOidcClientMetadata

`func (o *ListApplications200ResponseInner) GetOidcClientMetadata() ListApplications200ResponseInnerOidcClientMetadata`

GetOidcClientMetadata returns the OidcClientMetadata field if non-nil, zero value otherwise.

### GetOidcClientMetadataOk

`func (o *ListApplications200ResponseInner) GetOidcClientMetadataOk() (*ListApplications200ResponseInnerOidcClientMetadata, bool)`

GetOidcClientMetadataOk returns a tuple with the OidcClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientMetadata

`func (o *ListApplications200ResponseInner) SetOidcClientMetadata(v ListApplications200ResponseInnerOidcClientMetadata)`

SetOidcClientMetadata sets OidcClientMetadata field to given value.


### GetCustomClientMetadata

`func (o *ListApplications200ResponseInner) GetCustomClientMetadata() ListApplications200ResponseInnerCustomClientMetadata`

GetCustomClientMetadata returns the CustomClientMetadata field if non-nil, zero value otherwise.

### GetCustomClientMetadataOk

`func (o *ListApplications200ResponseInner) GetCustomClientMetadataOk() (*ListApplications200ResponseInnerCustomClientMetadata, bool)`

GetCustomClientMetadataOk returns a tuple with the CustomClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomClientMetadata

`func (o *ListApplications200ResponseInner) SetCustomClientMetadata(v ListApplications200ResponseInnerCustomClientMetadata)`

SetCustomClientMetadata sets CustomClientMetadata field to given value.


### GetProtectedAppMetadata

`func (o *ListApplications200ResponseInner) GetProtectedAppMetadata() ListApplications200ResponseInnerProtectedAppMetadata`

GetProtectedAppMetadata returns the ProtectedAppMetadata field if non-nil, zero value otherwise.

### GetProtectedAppMetadataOk

`func (o *ListApplications200ResponseInner) GetProtectedAppMetadataOk() (*ListApplications200ResponseInnerProtectedAppMetadata, bool)`

GetProtectedAppMetadataOk returns a tuple with the ProtectedAppMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedAppMetadata

`func (o *ListApplications200ResponseInner) SetProtectedAppMetadata(v ListApplications200ResponseInnerProtectedAppMetadata)`

SetProtectedAppMetadata sets ProtectedAppMetadata field to given value.


### SetProtectedAppMetadataNil

`func (o *ListApplications200ResponseInner) SetProtectedAppMetadataNil(b bool)`

 SetProtectedAppMetadataNil sets the value for ProtectedAppMetadata to be an explicit nil

### UnsetProtectedAppMetadata
`func (o *ListApplications200ResponseInner) UnsetProtectedAppMetadata()`

UnsetProtectedAppMetadata ensures that no value is present for ProtectedAppMetadata, not even an explicit nil
### GetCustomData

`func (o *ListApplications200ResponseInner) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *ListApplications200ResponseInner) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *ListApplications200ResponseInner) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.


### GetIsThirdParty

`func (o *ListApplications200ResponseInner) GetIsThirdParty() bool`

GetIsThirdParty returns the IsThirdParty field if non-nil, zero value otherwise.

### GetIsThirdPartyOk

`func (o *ListApplications200ResponseInner) GetIsThirdPartyOk() (*bool, bool)`

GetIsThirdPartyOk returns a tuple with the IsThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsThirdParty

`func (o *ListApplications200ResponseInner) SetIsThirdParty(v bool)`

SetIsThirdParty sets IsThirdParty field to given value.


### GetCreatedAt

`func (o *ListApplications200ResponseInner) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListApplications200ResponseInner) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListApplications200ResponseInner) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


