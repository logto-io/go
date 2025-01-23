# CreateResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Name** | **string** | The name of the resource. | 
**Indicator** | **string** | The unique resource indicator. Should be a valid URI. | 
**AccessTokenTtl** | Pointer to **float32** | The access token TTL in seconds. It affects the &#x60;exp&#x60; claim of the access token granted for this resource. | [optional] [default to 3600]

## Methods

### NewCreateResourceRequest

`func NewCreateResourceRequest(name string, indicator string, ) *CreateResourceRequest`

NewCreateResourceRequest instantiates a new CreateResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourceRequestWithDefaults

`func NewCreateResourceRequestWithDefaults() *CreateResourceRequest`

NewCreateResourceRequestWithDefaults instantiates a new CreateResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateResourceRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateResourceRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateResourceRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateResourceRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *CreateResourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateResourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateResourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIndicator

`func (o *CreateResourceRequest) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *CreateResourceRequest) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *CreateResourceRequest) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.


### GetAccessTokenTtl

`func (o *CreateResourceRequest) GetAccessTokenTtl() float32`

GetAccessTokenTtl returns the AccessTokenTtl field if non-nil, zero value otherwise.

### GetAccessTokenTtlOk

`func (o *CreateResourceRequest) GetAccessTokenTtlOk() (*float32, bool)`

GetAccessTokenTtlOk returns a tuple with the AccessTokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtl

`func (o *CreateResourceRequest) SetAccessTokenTtl(v float32)`

SetAccessTokenTtl sets AccessTokenTtl field to given value.

### HasAccessTokenTtl

`func (o *CreateResourceRequest) HasAccessTokenTtl() bool`

HasAccessTokenTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


