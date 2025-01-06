# GetUserAssetServiceStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**GetUserAssetServiceStatus200ResponseStatus**](GetUserAssetServiceStatus200ResponseStatus.md) |  | 
**AllowUploadMimeTypes** | Pointer to **[]string** |  | [optional] 
**MaxUploadFileSize** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetUserAssetServiceStatus200Response

`func NewGetUserAssetServiceStatus200Response(status GetUserAssetServiceStatus200ResponseStatus, ) *GetUserAssetServiceStatus200Response`

NewGetUserAssetServiceStatus200Response instantiates a new GetUserAssetServiceStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserAssetServiceStatus200ResponseWithDefaults

`func NewGetUserAssetServiceStatus200ResponseWithDefaults() *GetUserAssetServiceStatus200Response`

NewGetUserAssetServiceStatus200ResponseWithDefaults instantiates a new GetUserAssetServiceStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetUserAssetServiceStatus200Response) GetStatus() GetUserAssetServiceStatus200ResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUserAssetServiceStatus200Response) GetStatusOk() (*GetUserAssetServiceStatus200ResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUserAssetServiceStatus200Response) SetStatus(v GetUserAssetServiceStatus200ResponseStatus)`

SetStatus sets Status field to given value.


### GetAllowUploadMimeTypes

`func (o *GetUserAssetServiceStatus200Response) GetAllowUploadMimeTypes() []string`

GetAllowUploadMimeTypes returns the AllowUploadMimeTypes field if non-nil, zero value otherwise.

### GetAllowUploadMimeTypesOk

`func (o *GetUserAssetServiceStatus200Response) GetAllowUploadMimeTypesOk() (*[]string, bool)`

GetAllowUploadMimeTypesOk returns a tuple with the AllowUploadMimeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUploadMimeTypes

`func (o *GetUserAssetServiceStatus200Response) SetAllowUploadMimeTypes(v []string)`

SetAllowUploadMimeTypes sets AllowUploadMimeTypes field to given value.

### HasAllowUploadMimeTypes

`func (o *GetUserAssetServiceStatus200Response) HasAllowUploadMimeTypes() bool`

HasAllowUploadMimeTypes returns a boolean if a field has been set.

### GetMaxUploadFileSize

`func (o *GetUserAssetServiceStatus200Response) GetMaxUploadFileSize() float32`

GetMaxUploadFileSize returns the MaxUploadFileSize field if non-nil, zero value otherwise.

### GetMaxUploadFileSizeOk

`func (o *GetUserAssetServiceStatus200Response) GetMaxUploadFileSizeOk() (*float32, bool)`

GetMaxUploadFileSizeOk returns a tuple with the MaxUploadFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUploadFileSize

`func (o *GetUserAssetServiceStatus200Response) SetMaxUploadFileSize(v float32)`

SetMaxUploadFileSize sets MaxUploadFileSize field to given value.

### HasMaxUploadFileSize

`func (o *GetUserAssetServiceStatus200Response) HasMaxUploadFileSize() bool`

HasMaxUploadFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


