# CreateApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**OidcClientMetadata** | Pointer to [**ListApplications200ResponseInnerOidcClientMetadata**](ListApplications200ResponseInnerOidcClientMetadata.md) |  | [optional] 
**CustomClientMetadata** | Pointer to [**ListApplications200ResponseInnerCustomClientMetadata**](ListApplications200ResponseInnerCustomClientMetadata.md) |  | [optional] 
**CustomData** | Pointer to **map[string]interface{}** | arbitrary | [optional] 
**IsThirdParty** | Pointer to **bool** |  | [optional] 
**ProtectedAppMetadata** | Pointer to [**CreateApplicationRequestProtectedAppMetadata**](CreateApplicationRequestProtectedAppMetadata.md) |  | [optional] 

## Methods

### NewCreateApplicationRequest

`func NewCreateApplicationRequest(name string, type_ string, ) *CreateApplicationRequest`

NewCreateApplicationRequest instantiates a new CreateApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationRequestWithDefaults

`func NewCreateApplicationRequestWithDefaults() *CreateApplicationRequest`

NewCreateApplicationRequestWithDefaults instantiates a new CreateApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApplicationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateApplicationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateApplicationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateApplicationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateApplicationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateApplicationRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateApplicationRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *CreateApplicationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateApplicationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateApplicationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetOidcClientMetadata

`func (o *CreateApplicationRequest) GetOidcClientMetadata() ListApplications200ResponseInnerOidcClientMetadata`

GetOidcClientMetadata returns the OidcClientMetadata field if non-nil, zero value otherwise.

### GetOidcClientMetadataOk

`func (o *CreateApplicationRequest) GetOidcClientMetadataOk() (*ListApplications200ResponseInnerOidcClientMetadata, bool)`

GetOidcClientMetadataOk returns a tuple with the OidcClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientMetadata

`func (o *CreateApplicationRequest) SetOidcClientMetadata(v ListApplications200ResponseInnerOidcClientMetadata)`

SetOidcClientMetadata sets OidcClientMetadata field to given value.

### HasOidcClientMetadata

`func (o *CreateApplicationRequest) HasOidcClientMetadata() bool`

HasOidcClientMetadata returns a boolean if a field has been set.

### GetCustomClientMetadata

`func (o *CreateApplicationRequest) GetCustomClientMetadata() ListApplications200ResponseInnerCustomClientMetadata`

GetCustomClientMetadata returns the CustomClientMetadata field if non-nil, zero value otherwise.

### GetCustomClientMetadataOk

`func (o *CreateApplicationRequest) GetCustomClientMetadataOk() (*ListApplications200ResponseInnerCustomClientMetadata, bool)`

GetCustomClientMetadataOk returns a tuple with the CustomClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomClientMetadata

`func (o *CreateApplicationRequest) SetCustomClientMetadata(v ListApplications200ResponseInnerCustomClientMetadata)`

SetCustomClientMetadata sets CustomClientMetadata field to given value.

### HasCustomClientMetadata

`func (o *CreateApplicationRequest) HasCustomClientMetadata() bool`

HasCustomClientMetadata returns a boolean if a field has been set.

### GetCustomData

`func (o *CreateApplicationRequest) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *CreateApplicationRequest) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *CreateApplicationRequest) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *CreateApplicationRequest) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetIsThirdParty

`func (o *CreateApplicationRequest) GetIsThirdParty() bool`

GetIsThirdParty returns the IsThirdParty field if non-nil, zero value otherwise.

### GetIsThirdPartyOk

`func (o *CreateApplicationRequest) GetIsThirdPartyOk() (*bool, bool)`

GetIsThirdPartyOk returns a tuple with the IsThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsThirdParty

`func (o *CreateApplicationRequest) SetIsThirdParty(v bool)`

SetIsThirdParty sets IsThirdParty field to given value.

### HasIsThirdParty

`func (o *CreateApplicationRequest) HasIsThirdParty() bool`

HasIsThirdParty returns a boolean if a field has been set.

### GetProtectedAppMetadata

`func (o *CreateApplicationRequest) GetProtectedAppMetadata() CreateApplicationRequestProtectedAppMetadata`

GetProtectedAppMetadata returns the ProtectedAppMetadata field if non-nil, zero value otherwise.

### GetProtectedAppMetadataOk

`func (o *CreateApplicationRequest) GetProtectedAppMetadataOk() (*CreateApplicationRequestProtectedAppMetadata, bool)`

GetProtectedAppMetadataOk returns a tuple with the ProtectedAppMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedAppMetadata

`func (o *CreateApplicationRequest) SetProtectedAppMetadata(v CreateApplicationRequestProtectedAppMetadata)`

SetProtectedAppMetadata sets ProtectedAppMetadata field to given value.

### HasProtectedAppMetadata

`func (o *CreateApplicationRequest) HasProtectedAppMetadata() bool`

HasProtectedAppMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


