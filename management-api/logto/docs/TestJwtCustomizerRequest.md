# TestJwtCustomizerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenType** | Pointer to **interface{}** | The token type to test the JWT customizer for. | [optional] 
**Payload** | Pointer to [**TestJwtCustomizerRequestPayload**](TestJwtCustomizerRequestPayload.md) |  | [optional] 
**EnvironmentVariables** | Pointer to **map[string]string** |  | [optional] 
**Script** | **string** |  | 
**Token** | [**GetJwtCustomizer200ResponseOneOf1TokenSample**](GetJwtCustomizer200ResponseOneOf1TokenSample.md) |  | 
**Context** | [**GetJwtCustomizer200ResponseOneOfContextSample**](GetJwtCustomizer200ResponseOneOfContextSample.md) |  | 

## Methods

### NewTestJwtCustomizerRequest

`func NewTestJwtCustomizerRequest(script string, token GetJwtCustomizer200ResponseOneOf1TokenSample, context GetJwtCustomizer200ResponseOneOfContextSample, ) *TestJwtCustomizerRequest`

NewTestJwtCustomizerRequest instantiates a new TestJwtCustomizerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestJwtCustomizerRequestWithDefaults

`func NewTestJwtCustomizerRequestWithDefaults() *TestJwtCustomizerRequest`

NewTestJwtCustomizerRequestWithDefaults instantiates a new TestJwtCustomizerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenType

`func (o *TestJwtCustomizerRequest) GetTokenType() interface{}`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TestJwtCustomizerRequest) GetTokenTypeOk() (*interface{}, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TestJwtCustomizerRequest) SetTokenType(v interface{})`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *TestJwtCustomizerRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### SetTokenTypeNil

`func (o *TestJwtCustomizerRequest) SetTokenTypeNil(b bool)`

 SetTokenTypeNil sets the value for TokenType to be an explicit nil

### UnsetTokenType
`func (o *TestJwtCustomizerRequest) UnsetTokenType()`

UnsetTokenType ensures that no value is present for TokenType, not even an explicit nil
### GetPayload

`func (o *TestJwtCustomizerRequest) GetPayload() TestJwtCustomizerRequestPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *TestJwtCustomizerRequest) GetPayloadOk() (*TestJwtCustomizerRequestPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *TestJwtCustomizerRequest) SetPayload(v TestJwtCustomizerRequestPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *TestJwtCustomizerRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *TestJwtCustomizerRequest) GetEnvironmentVariables() map[string]string`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *TestJwtCustomizerRequest) GetEnvironmentVariablesOk() (*map[string]string, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *TestJwtCustomizerRequest) SetEnvironmentVariables(v map[string]string)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *TestJwtCustomizerRequest) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetScript

`func (o *TestJwtCustomizerRequest) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *TestJwtCustomizerRequest) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *TestJwtCustomizerRequest) SetScript(v string)`

SetScript sets Script field to given value.


### GetToken

`func (o *TestJwtCustomizerRequest) GetToken() GetJwtCustomizer200ResponseOneOf1TokenSample`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TestJwtCustomizerRequest) GetTokenOk() (*GetJwtCustomizer200ResponseOneOf1TokenSample, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TestJwtCustomizerRequest) SetToken(v GetJwtCustomizer200ResponseOneOf1TokenSample)`

SetToken sets Token field to given value.


### GetContext

`func (o *TestJwtCustomizerRequest) GetContext() GetJwtCustomizer200ResponseOneOfContextSample`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TestJwtCustomizerRequest) GetContextOk() (*GetJwtCustomizer200ResponseOneOfContextSample, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TestJwtCustomizerRequest) SetContext(v GetJwtCustomizer200ResponseOneOfContextSample)`

SetContext sets Context field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


