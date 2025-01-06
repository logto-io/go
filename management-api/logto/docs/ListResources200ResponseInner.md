# ListResources200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Indicator** | **string** |  | 
**IsDefault** | **bool** |  | 
**AccessTokenTtl** | **float32** |  | 
**Scopes** | Pointer to [**[]ListResources200ResponseInnerScopesInner**](ListResources200ResponseInnerScopesInner.md) |  | [optional] 

## Methods

### NewListResources200ResponseInner

`func NewListResources200ResponseInner(tenantId string, id string, name string, indicator string, isDefault bool, accessTokenTtl float32, ) *ListResources200ResponseInner`

NewListResources200ResponseInner instantiates a new ListResources200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResources200ResponseInnerWithDefaults

`func NewListResources200ResponseInnerWithDefaults() *ListResources200ResponseInner`

NewListResources200ResponseInnerWithDefaults instantiates a new ListResources200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListResources200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListResources200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListResources200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListResources200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListResources200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListResources200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListResources200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListResources200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListResources200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetIndicator

`func (o *ListResources200ResponseInner) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *ListResources200ResponseInner) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *ListResources200ResponseInner) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.


### GetIsDefault

`func (o *ListResources200ResponseInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ListResources200ResponseInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ListResources200ResponseInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetAccessTokenTtl

`func (o *ListResources200ResponseInner) GetAccessTokenTtl() float32`

GetAccessTokenTtl returns the AccessTokenTtl field if non-nil, zero value otherwise.

### GetAccessTokenTtlOk

`func (o *ListResources200ResponseInner) GetAccessTokenTtlOk() (*float32, bool)`

GetAccessTokenTtlOk returns a tuple with the AccessTokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtl

`func (o *ListResources200ResponseInner) SetAccessTokenTtl(v float32)`

SetAccessTokenTtl sets AccessTokenTtl field to given value.


### GetScopes

`func (o *ListResources200ResponseInner) GetScopes() []ListResources200ResponseInnerScopesInner`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ListResources200ResponseInner) GetScopesOk() (*[]ListResources200ResponseInnerScopesInner, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ListResources200ResponseInner) SetScopes(v []ListResources200ResponseInnerScopesInner)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ListResources200ResponseInner) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


