# UpdateResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The updated name of the resource. | [optional] 
**AccessTokenTtl** | Pointer to **float32** | The updated access token TTL in seconds. | [optional] 

## Methods

### NewUpdateResourceRequest

`func NewUpdateResourceRequest() *UpdateResourceRequest`

NewUpdateResourceRequest instantiates a new UpdateResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceRequestWithDefaults

`func NewUpdateResourceRequestWithDefaults() *UpdateResourceRequest`

NewUpdateResourceRequestWithDefaults instantiates a new UpdateResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *UpdateResourceRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateResourceRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateResourceRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UpdateResourceRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *UpdateResourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateResourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateResourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateResourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccessTokenTtl

`func (o *UpdateResourceRequest) GetAccessTokenTtl() float32`

GetAccessTokenTtl returns the AccessTokenTtl field if non-nil, zero value otherwise.

### GetAccessTokenTtlOk

`func (o *UpdateResourceRequest) GetAccessTokenTtlOk() (*float32, bool)`

GetAccessTokenTtlOk returns a tuple with the AccessTokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtl

`func (o *UpdateResourceRequest) SetAccessTokenTtl(v float32)`

SetAccessTokenTtl sets AccessTokenTtl field to given value.

### HasAccessTokenTtl

`func (o *UpdateResourceRequest) HasAccessTokenTtl() bool`

HasAccessTokenTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


