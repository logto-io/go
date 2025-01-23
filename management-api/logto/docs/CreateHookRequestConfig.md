# CreateHookRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Retries** | Pointer to **float32** | Now the retry times is fixed to 3. Keep for backward compatibility. | [optional] 

## Methods

### NewCreateHookRequestConfig

`func NewCreateHookRequestConfig(url string, ) *CreateHookRequestConfig`

NewCreateHookRequestConfig instantiates a new CreateHookRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHookRequestConfigWithDefaults

`func NewCreateHookRequestConfigWithDefaults() *CreateHookRequestConfig`

NewCreateHookRequestConfigWithDefaults instantiates a new CreateHookRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateHookRequestConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateHookRequestConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateHookRequestConfig) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *CreateHookRequestConfig) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CreateHookRequestConfig) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CreateHookRequestConfig) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CreateHookRequestConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRetries

`func (o *CreateHookRequestConfig) GetRetries() float32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *CreateHookRequestConfig) GetRetriesOk() (*float32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *CreateHookRequestConfig) SetRetries(v float32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *CreateHookRequestConfig) HasRetries() bool`

HasRetries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


