# ListOrganizationApplications200ResponseInner

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
**OrganizationRoles** | [**[]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner**](ListApplicationOrganizations200ResponseInnerOrganizationRolesInner.md) |  | 

## Methods

### NewListOrganizationApplications200ResponseInner

`func NewListOrganizationApplications200ResponseInner(tenantId string, id string, name string, secret string, description NullableString, type_ string, oidcClientMetadata ListApplications200ResponseInnerOidcClientMetadata, customClientMetadata ListApplications200ResponseInnerCustomClientMetadata, protectedAppMetadata NullableListApplications200ResponseInnerProtectedAppMetadata, customData map[string]interface{}, isThirdParty bool, createdAt float32, organizationRoles []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, ) *ListOrganizationApplications200ResponseInner`

NewListOrganizationApplications200ResponseInner instantiates a new ListOrganizationApplications200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizationApplications200ResponseInnerWithDefaults

`func NewListOrganizationApplications200ResponseInnerWithDefaults() *ListOrganizationApplications200ResponseInner`

NewListOrganizationApplications200ResponseInnerWithDefaults instantiates a new ListOrganizationApplications200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListOrganizationApplications200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListOrganizationApplications200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListOrganizationApplications200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListOrganizationApplications200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOrganizationApplications200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOrganizationApplications200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListOrganizationApplications200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListOrganizationApplications200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListOrganizationApplications200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetSecret

`func (o *ListOrganizationApplications200ResponseInner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ListOrganizationApplications200ResponseInner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ListOrganizationApplications200ResponseInner) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetDescription

`func (o *ListOrganizationApplications200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListOrganizationApplications200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListOrganizationApplications200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ListOrganizationApplications200ResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListOrganizationApplications200ResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *ListOrganizationApplications200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListOrganizationApplications200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListOrganizationApplications200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetOidcClientMetadata

`func (o *ListOrganizationApplications200ResponseInner) GetOidcClientMetadata() ListApplications200ResponseInnerOidcClientMetadata`

GetOidcClientMetadata returns the OidcClientMetadata field if non-nil, zero value otherwise.

### GetOidcClientMetadataOk

`func (o *ListOrganizationApplications200ResponseInner) GetOidcClientMetadataOk() (*ListApplications200ResponseInnerOidcClientMetadata, bool)`

GetOidcClientMetadataOk returns a tuple with the OidcClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientMetadata

`func (o *ListOrganizationApplications200ResponseInner) SetOidcClientMetadata(v ListApplications200ResponseInnerOidcClientMetadata)`

SetOidcClientMetadata sets OidcClientMetadata field to given value.


### GetCustomClientMetadata

`func (o *ListOrganizationApplications200ResponseInner) GetCustomClientMetadata() ListApplications200ResponseInnerCustomClientMetadata`

GetCustomClientMetadata returns the CustomClientMetadata field if non-nil, zero value otherwise.

### GetCustomClientMetadataOk

`func (o *ListOrganizationApplications200ResponseInner) GetCustomClientMetadataOk() (*ListApplications200ResponseInnerCustomClientMetadata, bool)`

GetCustomClientMetadataOk returns a tuple with the CustomClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomClientMetadata

`func (o *ListOrganizationApplications200ResponseInner) SetCustomClientMetadata(v ListApplications200ResponseInnerCustomClientMetadata)`

SetCustomClientMetadata sets CustomClientMetadata field to given value.


### GetProtectedAppMetadata

`func (o *ListOrganizationApplications200ResponseInner) GetProtectedAppMetadata() ListApplications200ResponseInnerProtectedAppMetadata`

GetProtectedAppMetadata returns the ProtectedAppMetadata field if non-nil, zero value otherwise.

### GetProtectedAppMetadataOk

`func (o *ListOrganizationApplications200ResponseInner) GetProtectedAppMetadataOk() (*ListApplications200ResponseInnerProtectedAppMetadata, bool)`

GetProtectedAppMetadataOk returns a tuple with the ProtectedAppMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedAppMetadata

`func (o *ListOrganizationApplications200ResponseInner) SetProtectedAppMetadata(v ListApplications200ResponseInnerProtectedAppMetadata)`

SetProtectedAppMetadata sets ProtectedAppMetadata field to given value.


### SetProtectedAppMetadataNil

`func (o *ListOrganizationApplications200ResponseInner) SetProtectedAppMetadataNil(b bool)`

 SetProtectedAppMetadataNil sets the value for ProtectedAppMetadata to be an explicit nil

### UnsetProtectedAppMetadata
`func (o *ListOrganizationApplications200ResponseInner) UnsetProtectedAppMetadata()`

UnsetProtectedAppMetadata ensures that no value is present for ProtectedAppMetadata, not even an explicit nil
### GetCustomData

`func (o *ListOrganizationApplications200ResponseInner) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *ListOrganizationApplications200ResponseInner) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *ListOrganizationApplications200ResponseInner) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.


### GetIsThirdParty

`func (o *ListOrganizationApplications200ResponseInner) GetIsThirdParty() bool`

GetIsThirdParty returns the IsThirdParty field if non-nil, zero value otherwise.

### GetIsThirdPartyOk

`func (o *ListOrganizationApplications200ResponseInner) GetIsThirdPartyOk() (*bool, bool)`

GetIsThirdPartyOk returns a tuple with the IsThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsThirdParty

`func (o *ListOrganizationApplications200ResponseInner) SetIsThirdParty(v bool)`

SetIsThirdParty sets IsThirdParty field to given value.


### GetCreatedAt

`func (o *ListOrganizationApplications200ResponseInner) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListOrganizationApplications200ResponseInner) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListOrganizationApplications200ResponseInner) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetOrganizationRoles

`func (o *ListOrganizationApplications200ResponseInner) GetOrganizationRoles() []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner`

GetOrganizationRoles returns the OrganizationRoles field if non-nil, zero value otherwise.

### GetOrganizationRolesOk

`func (o *ListOrganizationApplications200ResponseInner) GetOrganizationRolesOk() (*[]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, bool)`

GetOrganizationRolesOk returns a tuple with the OrganizationRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoles

`func (o *ListOrganizationApplications200ResponseInner) SetOrganizationRoles(v []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner)`

SetOrganizationRoles sets OrganizationRoles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


