# GetJwtCustomizer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | **string** |  | 
**EnvironmentVariables** | Pointer to **map[string]string** |  | [optional] 
**ContextSample** | Pointer to **map[string]interface{}** | arbitrary | [optional] 
**TokenSample** | Pointer to [**GetJwtCustomizer200ResponseOneOf1TokenSample**](GetJwtCustomizer200ResponseOneOf1TokenSample.md) |  | [optional] 

## Methods

### NewGetJwtCustomizer200Response

`func NewGetJwtCustomizer200Response(script string, ) *GetJwtCustomizer200Response`

NewGetJwtCustomizer200Response instantiates a new GetJwtCustomizer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJwtCustomizer200ResponseWithDefaults

`func NewGetJwtCustomizer200ResponseWithDefaults() *GetJwtCustomizer200Response`

NewGetJwtCustomizer200ResponseWithDefaults instantiates a new GetJwtCustomizer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *GetJwtCustomizer200Response) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *GetJwtCustomizer200Response) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *GetJwtCustomizer200Response) SetScript(v string)`

SetScript sets Script field to given value.


### GetEnvironmentVariables

`func (o *GetJwtCustomizer200Response) GetEnvironmentVariables() map[string]string`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GetJwtCustomizer200Response) GetEnvironmentVariablesOk() (*map[string]string, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GetJwtCustomizer200Response) SetEnvironmentVariables(v map[string]string)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GetJwtCustomizer200Response) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetContextSample

`func (o *GetJwtCustomizer200Response) GetContextSample() map[string]interface{}`

GetContextSample returns the ContextSample field if non-nil, zero value otherwise.

### GetContextSampleOk

`func (o *GetJwtCustomizer200Response) GetContextSampleOk() (*map[string]interface{}, bool)`

GetContextSampleOk returns a tuple with the ContextSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextSample

`func (o *GetJwtCustomizer200Response) SetContextSample(v map[string]interface{})`

SetContextSample sets ContextSample field to given value.

### HasContextSample

`func (o *GetJwtCustomizer200Response) HasContextSample() bool`

HasContextSample returns a boolean if a field has been set.

### GetTokenSample

`func (o *GetJwtCustomizer200Response) GetTokenSample() GetJwtCustomizer200ResponseOneOf1TokenSample`

GetTokenSample returns the TokenSample field if non-nil, zero value otherwise.

### GetTokenSampleOk

`func (o *GetJwtCustomizer200Response) GetTokenSampleOk() (*GetJwtCustomizer200ResponseOneOf1TokenSample, bool)`

GetTokenSampleOk returns a tuple with the TokenSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSample

`func (o *GetJwtCustomizer200Response) SetTokenSample(v GetJwtCustomizer200ResponseOneOf1TokenSample)`

SetTokenSample sets TokenSample field to given value.

### HasTokenSample

`func (o *GetJwtCustomizer200Response) HasTokenSample() bool`

HasTokenSample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


