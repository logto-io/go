# DeleteApplicationLegacySecret200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Secret** | **string** |  | 
**Description** | **NullableString** |  | 
**Type** | **string** |  | 
**OidcClientMetadata** | [**ListApplications200ResponseInnerOidcClientMetadata**](ListApplications200ResponseInnerOidcClientMetadata.md) |  | 
**CustomClientMetadata** | [**ListApplications200ResponseInnerCustomClientMetadata**](ListApplications200ResponseInnerCustomClientMetadata.md) |  | 
**ProtectedAppMetadata** | [**NullableListApplications200ResponseInnerProtectedAppMetadata**](ListApplications200ResponseInnerProtectedAppMetadata.md) |  | 
**CustomData** | **map[string]interface{}** | arbitrary | 
**IsThirdParty** | **bool** |  | 
**CreatedAt** | **float32** |  | 

## Methods

### NewDeleteApplicationLegacySecret200Response

`func NewDeleteApplicationLegacySecret200Response(tenantId string, id string, name string, secret string, description NullableString, type_ string, oidcClientMetadata ListApplications200ResponseInnerOidcClientMetadata, customClientMetadata ListApplications200ResponseInnerCustomClientMetadata, protectedAppMetadata NullableListApplications200ResponseInnerProtectedAppMetadata, customData map[string]interface{}, isThirdParty bool, createdAt float32, ) *DeleteApplicationLegacySecret200Response`

NewDeleteApplicationLegacySecret200Response instantiates a new DeleteApplicationLegacySecret200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteApplicationLegacySecret200ResponseWithDefaults

`func NewDeleteApplicationLegacySecret200ResponseWithDefaults() *DeleteApplicationLegacySecret200Response`

NewDeleteApplicationLegacySecret200ResponseWithDefaults instantiates a new DeleteApplicationLegacySecret200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DeleteApplicationLegacySecret200Response) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DeleteApplicationLegacySecret200Response) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DeleteApplicationLegacySecret200Response) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *DeleteApplicationLegacySecret200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteApplicationLegacySecret200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteApplicationLegacySecret200Response) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeleteApplicationLegacySecret200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeleteApplicationLegacySecret200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeleteApplicationLegacySecret200Response) SetName(v string)`

SetName sets Name field to given value.


### GetSecret

`func (o *DeleteApplicationLegacySecret200Response) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DeleteApplicationLegacySecret200Response) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DeleteApplicationLegacySecret200Response) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetDescription

`func (o *DeleteApplicationLegacySecret200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeleteApplicationLegacySecret200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeleteApplicationLegacySecret200Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *DeleteApplicationLegacySecret200Response) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DeleteApplicationLegacySecret200Response) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *DeleteApplicationLegacySecret200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeleteApplicationLegacySecret200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeleteApplicationLegacySecret200Response) SetType(v string)`

SetType sets Type field to given value.


### GetOidcClientMetadata

`func (o *DeleteApplicationLegacySecret200Response) GetOidcClientMetadata() ListApplications200ResponseInnerOidcClientMetadata`

GetOidcClientMetadata returns the OidcClientMetadata field if non-nil, zero value otherwise.

### GetOidcClientMetadataOk

`func (o *DeleteApplicationLegacySecret200Response) GetOidcClientMetadataOk() (*ListApplications200ResponseInnerOidcClientMetadata, bool)`

GetOidcClientMetadataOk returns a tuple with the OidcClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientMetadata

`func (o *DeleteApplicationLegacySecret200Response) SetOidcClientMetadata(v ListApplications200ResponseInnerOidcClientMetadata)`

SetOidcClientMetadata sets OidcClientMetadata field to given value.


### GetCustomClientMetadata

`func (o *DeleteApplicationLegacySecret200Response) GetCustomClientMetadata() ListApplications200ResponseInnerCustomClientMetadata`

GetCustomClientMetadata returns the CustomClientMetadata field if non-nil, zero value otherwise.

### GetCustomClientMetadataOk

`func (o *DeleteApplicationLegacySecret200Response) GetCustomClientMetadataOk() (*ListApplications200ResponseInnerCustomClientMetadata, bool)`

GetCustomClientMetadataOk returns a tuple with the CustomClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomClientMetadata

`func (o *DeleteApplicationLegacySecret200Response) SetCustomClientMetadata(v ListApplications200ResponseInnerCustomClientMetadata)`

SetCustomClientMetadata sets CustomClientMetadata field to given value.


### GetProtectedAppMetadata

`func (o *DeleteApplicationLegacySecret200Response) GetProtectedAppMetadata() ListApplications200ResponseInnerProtectedAppMetadata`

GetProtectedAppMetadata returns the ProtectedAppMetadata field if non-nil, zero value otherwise.

### GetProtectedAppMetadataOk

`func (o *DeleteApplicationLegacySecret200Response) GetProtectedAppMetadataOk() (*ListApplications200ResponseInnerProtectedAppMetadata, bool)`

GetProtectedAppMetadataOk returns a tuple with the ProtectedAppMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedAppMetadata

`func (o *DeleteApplicationLegacySecret200Response) SetProtectedAppMetadata(v ListApplications200ResponseInnerProtectedAppMetadata)`

SetProtectedAppMetadata sets ProtectedAppMetadata field to given value.


### SetProtectedAppMetadataNil

`func (o *DeleteApplicationLegacySecret200Response) SetProtectedAppMetadataNil(b bool)`

 SetProtectedAppMetadataNil sets the value for ProtectedAppMetadata to be an explicit nil

### UnsetProtectedAppMetadata
`func (o *DeleteApplicationLegacySecret200Response) UnsetProtectedAppMetadata()`

UnsetProtectedAppMetadata ensures that no value is present for ProtectedAppMetadata, not even an explicit nil
### GetCustomData

`func (o *DeleteApplicationLegacySecret200Response) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *DeleteApplicationLegacySecret200Response) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *DeleteApplicationLegacySecret200Response) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.


### GetIsThirdParty

`func (o *DeleteApplicationLegacySecret200Response) GetIsThirdParty() bool`

GetIsThirdParty returns the IsThirdParty field if non-nil, zero value otherwise.

### GetIsThirdPartyOk

`func (o *DeleteApplicationLegacySecret200Response) GetIsThirdPartyOk() (*bool, bool)`

GetIsThirdPartyOk returns a tuple with the IsThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsThirdParty

`func (o *DeleteApplicationLegacySecret200Response) SetIsThirdParty(v bool)`

SetIsThirdParty sets IsThirdParty field to given value.


### GetCreatedAt

`func (o *DeleteApplicationLegacySecret200Response) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeleteApplicationLegacySecret200Response) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeleteApplicationLegacySecret200Response) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


