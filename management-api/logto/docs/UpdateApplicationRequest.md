# UpdateApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**OidcClientMetadata** | Pointer to [**ListApplications200ResponseInnerOidcClientMetadata**](ListApplications200ResponseInnerOidcClientMetadata.md) |  | [optional] 
**CustomClientMetadata** | Pointer to [**ListApplications200ResponseInnerCustomClientMetadata**](ListApplications200ResponseInnerCustomClientMetadata.md) |  | [optional] 
**CustomData** | Pointer to **map[string]interface{}** | arbitrary | [optional] 
**ProtectedAppMetadata** | Pointer to [**UpdateApplicationRequestProtectedAppMetadata**](UpdateApplicationRequestProtectedAppMetadata.md) |  | [optional] 
**IsAdmin** | Pointer to **bool** | Whether the application has admin access. User can enable the admin access for Machine-to-Machine apps. | [optional] 

## Methods

### NewUpdateApplicationRequest

`func NewUpdateApplicationRequest() *UpdateApplicationRequest`

NewUpdateApplicationRequest instantiates a new UpdateApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApplicationRequestWithDefaults

`func NewUpdateApplicationRequestWithDefaults() *UpdateApplicationRequest`

NewUpdateApplicationRequestWithDefaults instantiates a new UpdateApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateApplicationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateApplicationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateApplicationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateApplicationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateApplicationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateApplicationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateApplicationRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateApplicationRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOidcClientMetadata

`func (o *UpdateApplicationRequest) GetOidcClientMetadata() ListApplications200ResponseInnerOidcClientMetadata`

GetOidcClientMetadata returns the OidcClientMetadata field if non-nil, zero value otherwise.

### GetOidcClientMetadataOk

`func (o *UpdateApplicationRequest) GetOidcClientMetadataOk() (*ListApplications200ResponseInnerOidcClientMetadata, bool)`

GetOidcClientMetadataOk returns a tuple with the OidcClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientMetadata

`func (o *UpdateApplicationRequest) SetOidcClientMetadata(v ListApplications200ResponseInnerOidcClientMetadata)`

SetOidcClientMetadata sets OidcClientMetadata field to given value.

### HasOidcClientMetadata

`func (o *UpdateApplicationRequest) HasOidcClientMetadata() bool`

HasOidcClientMetadata returns a boolean if a field has been set.

### GetCustomClientMetadata

`func (o *UpdateApplicationRequest) GetCustomClientMetadata() ListApplications200ResponseInnerCustomClientMetadata`

GetCustomClientMetadata returns the CustomClientMetadata field if non-nil, zero value otherwise.

### GetCustomClientMetadataOk

`func (o *UpdateApplicationRequest) GetCustomClientMetadataOk() (*ListApplications200ResponseInnerCustomClientMetadata, bool)`

GetCustomClientMetadataOk returns a tuple with the CustomClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomClientMetadata

`func (o *UpdateApplicationRequest) SetCustomClientMetadata(v ListApplications200ResponseInnerCustomClientMetadata)`

SetCustomClientMetadata sets CustomClientMetadata field to given value.

### HasCustomClientMetadata

`func (o *UpdateApplicationRequest) HasCustomClientMetadata() bool`

HasCustomClientMetadata returns a boolean if a field has been set.

### GetCustomData

`func (o *UpdateApplicationRequest) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *UpdateApplicationRequest) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *UpdateApplicationRequest) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *UpdateApplicationRequest) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetProtectedAppMetadata

`func (o *UpdateApplicationRequest) GetProtectedAppMetadata() UpdateApplicationRequestProtectedAppMetadata`

GetProtectedAppMetadata returns the ProtectedAppMetadata field if non-nil, zero value otherwise.

### GetProtectedAppMetadataOk

`func (o *UpdateApplicationRequest) GetProtectedAppMetadataOk() (*UpdateApplicationRequestProtectedAppMetadata, bool)`

GetProtectedAppMetadataOk returns a tuple with the ProtectedAppMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedAppMetadata

`func (o *UpdateApplicationRequest) SetProtectedAppMetadata(v UpdateApplicationRequestProtectedAppMetadata)`

SetProtectedAppMetadata sets ProtectedAppMetadata field to given value.

### HasProtectedAppMetadata

`func (o *UpdateApplicationRequest) HasProtectedAppMetadata() bool`

HasProtectedAppMetadata returns a boolean if a field has been set.

### GetIsAdmin

`func (o *UpdateApplicationRequest) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UpdateApplicationRequest) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *UpdateApplicationRequest) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *UpdateApplicationRequest) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


