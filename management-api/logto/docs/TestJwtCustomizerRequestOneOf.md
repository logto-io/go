# TestJwtCustomizerRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenType** | **string** |  | 
**EnvironmentVariables** | Pointer to **map[string]string** |  | [optional] 
**Script** | **string** |  | 
**Token** | [**GetJwtCustomizer200ResponseOneOfTokenSample**](GetJwtCustomizer200ResponseOneOfTokenSample.md) |  | 
**Context** | [**GetJwtCustomizer200ResponseOneOfContextSample**](GetJwtCustomizer200ResponseOneOfContextSample.md) |  | 

## Methods

### NewTestJwtCustomizerRequestOneOf

`func NewTestJwtCustomizerRequestOneOf(tokenType string, script string, token GetJwtCustomizer200ResponseOneOfTokenSample, context GetJwtCustomizer200ResponseOneOfContextSample, ) *TestJwtCustomizerRequestOneOf`

NewTestJwtCustomizerRequestOneOf instantiates a new TestJwtCustomizerRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestJwtCustomizerRequestOneOfWithDefaults

`func NewTestJwtCustomizerRequestOneOfWithDefaults() *TestJwtCustomizerRequestOneOf`

NewTestJwtCustomizerRequestOneOfWithDefaults instantiates a new TestJwtCustomizerRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenType

`func (o *TestJwtCustomizerRequestOneOf) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TestJwtCustomizerRequestOneOf) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TestJwtCustomizerRequestOneOf) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetEnvironmentVariables

`func (o *TestJwtCustomizerRequestOneOf) GetEnvironmentVariables() map[string]string`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *TestJwtCustomizerRequestOneOf) GetEnvironmentVariablesOk() (*map[string]string, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *TestJwtCustomizerRequestOneOf) SetEnvironmentVariables(v map[string]string)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *TestJwtCustomizerRequestOneOf) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetScript

`func (o *TestJwtCustomizerRequestOneOf) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *TestJwtCustomizerRequestOneOf) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *TestJwtCustomizerRequestOneOf) SetScript(v string)`

SetScript sets Script field to given value.


### GetToken

`func (o *TestJwtCustomizerRequestOneOf) GetToken() GetJwtCustomizer200ResponseOneOfTokenSample`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TestJwtCustomizerRequestOneOf) GetTokenOk() (*GetJwtCustomizer200ResponseOneOfTokenSample, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TestJwtCustomizerRequestOneOf) SetToken(v GetJwtCustomizer200ResponseOneOfTokenSample)`

SetToken sets Token field to given value.


### GetContext

`func (o *TestJwtCustomizerRequestOneOf) GetContext() GetJwtCustomizer200ResponseOneOfContextSample`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TestJwtCustomizerRequestOneOf) GetContextOk() (*GetJwtCustomizer200ResponseOneOfContextSample, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TestJwtCustomizerRequestOneOf) SetContext(v GetJwtCustomizer200ResponseOneOfContextSample)`

SetContext sets Context field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


