# TestJwtCustomizerRequestOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenType** | **string** |  | 
**EnvironmentVariables** | Pointer to **map[string]string** |  | [optional] 
**Script** | **string** |  | 
**Token** | [**GetJwtCustomizer200ResponseOneOf1TokenSample**](GetJwtCustomizer200ResponseOneOf1TokenSample.md) |  | 

## Methods

### NewTestJwtCustomizerRequestOneOf1

`func NewTestJwtCustomizerRequestOneOf1(tokenType string, script string, token GetJwtCustomizer200ResponseOneOf1TokenSample, ) *TestJwtCustomizerRequestOneOf1`

NewTestJwtCustomizerRequestOneOf1 instantiates a new TestJwtCustomizerRequestOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestJwtCustomizerRequestOneOf1WithDefaults

`func NewTestJwtCustomizerRequestOneOf1WithDefaults() *TestJwtCustomizerRequestOneOf1`

NewTestJwtCustomizerRequestOneOf1WithDefaults instantiates a new TestJwtCustomizerRequestOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenType

`func (o *TestJwtCustomizerRequestOneOf1) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TestJwtCustomizerRequestOneOf1) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TestJwtCustomizerRequestOneOf1) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetEnvironmentVariables

`func (o *TestJwtCustomizerRequestOneOf1) GetEnvironmentVariables() map[string]string`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *TestJwtCustomizerRequestOneOf1) GetEnvironmentVariablesOk() (*map[string]string, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *TestJwtCustomizerRequestOneOf1) SetEnvironmentVariables(v map[string]string)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *TestJwtCustomizerRequestOneOf1) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetScript

`func (o *TestJwtCustomizerRequestOneOf1) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *TestJwtCustomizerRequestOneOf1) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *TestJwtCustomizerRequestOneOf1) SetScript(v string)`

SetScript sets Script field to given value.


### GetToken

`func (o *TestJwtCustomizerRequestOneOf1) GetToken() GetJwtCustomizer200ResponseOneOf1TokenSample`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TestJwtCustomizerRequestOneOf1) GetTokenOk() (*GetJwtCustomizer200ResponseOneOf1TokenSample, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TestJwtCustomizerRequestOneOf1) SetToken(v GetJwtCustomizer200ResponseOneOf1TokenSample)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


