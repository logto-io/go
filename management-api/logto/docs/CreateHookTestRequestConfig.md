# CreateHookTestRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Retries** | Pointer to **float32** | Now the retry times is fixed to 3. Keep for backward compatibility. | [optional] 

## Methods

### NewCreateHookTestRequestConfig

`func NewCreateHookTestRequestConfig(url string, ) *CreateHookTestRequestConfig`

NewCreateHookTestRequestConfig instantiates a new CreateHookTestRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHookTestRequestConfigWithDefaults

`func NewCreateHookTestRequestConfigWithDefaults() *CreateHookTestRequestConfig`

NewCreateHookTestRequestConfigWithDefaults instantiates a new CreateHookTestRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateHookTestRequestConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateHookTestRequestConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateHookTestRequestConfig) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *CreateHookTestRequestConfig) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CreateHookTestRequestConfig) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CreateHookTestRequestConfig) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CreateHookTestRequestConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRetries

`func (o *CreateHookTestRequestConfig) GetRetries() float32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *CreateHookTestRequestConfig) GetRetriesOk() (*float32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *CreateHookTestRequestConfig) SetRetries(v float32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *CreateHookTestRequestConfig) HasRetries() bool`

HasRetries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


