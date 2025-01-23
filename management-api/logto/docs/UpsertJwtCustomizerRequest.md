# UpsertJwtCustomizerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **interface{}** | The script of the JWT customizer. | [optional] 
**EnvironmentVariables** | Pointer to **interface{}** | The environment variables for the JWT customizer. | [optional] 
**ContextSample** | Pointer to **interface{}** | The sample context for the JWT customizer script testing purpose. | [optional] 
**TokenSample** | Pointer to **interface{}** | The sample raw token payload for the JWT customizer script testing purpose. | [optional] 

## Methods

### NewUpsertJwtCustomizerRequest

`func NewUpsertJwtCustomizerRequest() *UpsertJwtCustomizerRequest`

NewUpsertJwtCustomizerRequest instantiates a new UpsertJwtCustomizerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertJwtCustomizerRequestWithDefaults

`func NewUpsertJwtCustomizerRequestWithDefaults() *UpsertJwtCustomizerRequest`

NewUpsertJwtCustomizerRequestWithDefaults instantiates a new UpsertJwtCustomizerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *UpsertJwtCustomizerRequest) GetScript() interface{}`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *UpsertJwtCustomizerRequest) GetScriptOk() (*interface{}, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *UpsertJwtCustomizerRequest) SetScript(v interface{})`

SetScript sets Script field to given value.

### HasScript

`func (o *UpsertJwtCustomizerRequest) HasScript() bool`

HasScript returns a boolean if a field has been set.

### SetScriptNil

`func (o *UpsertJwtCustomizerRequest) SetScriptNil(b bool)`

 SetScriptNil sets the value for Script to be an explicit nil

### UnsetScript
`func (o *UpsertJwtCustomizerRequest) UnsetScript()`

UnsetScript ensures that no value is present for Script, not even an explicit nil
### GetEnvironmentVariables

`func (o *UpsertJwtCustomizerRequest) GetEnvironmentVariables() interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *UpsertJwtCustomizerRequest) GetEnvironmentVariablesOk() (*interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *UpsertJwtCustomizerRequest) SetEnvironmentVariables(v interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *UpsertJwtCustomizerRequest) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *UpsertJwtCustomizerRequest) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *UpsertJwtCustomizerRequest) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil
### GetContextSample

`func (o *UpsertJwtCustomizerRequest) GetContextSample() interface{}`

GetContextSample returns the ContextSample field if non-nil, zero value otherwise.

### GetContextSampleOk

`func (o *UpsertJwtCustomizerRequest) GetContextSampleOk() (*interface{}, bool)`

GetContextSampleOk returns a tuple with the ContextSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextSample

`func (o *UpsertJwtCustomizerRequest) SetContextSample(v interface{})`

SetContextSample sets ContextSample field to given value.

### HasContextSample

`func (o *UpsertJwtCustomizerRequest) HasContextSample() bool`

HasContextSample returns a boolean if a field has been set.

### SetContextSampleNil

`func (o *UpsertJwtCustomizerRequest) SetContextSampleNil(b bool)`

 SetContextSampleNil sets the value for ContextSample to be an explicit nil

### UnsetContextSample
`func (o *UpsertJwtCustomizerRequest) UnsetContextSample()`

UnsetContextSample ensures that no value is present for ContextSample, not even an explicit nil
### GetTokenSample

`func (o *UpsertJwtCustomizerRequest) GetTokenSample() interface{}`

GetTokenSample returns the TokenSample field if non-nil, zero value otherwise.

### GetTokenSampleOk

`func (o *UpsertJwtCustomizerRequest) GetTokenSampleOk() (*interface{}, bool)`

GetTokenSampleOk returns a tuple with the TokenSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSample

`func (o *UpsertJwtCustomizerRequest) SetTokenSample(v interface{})`

SetTokenSample sets TokenSample field to given value.

### HasTokenSample

`func (o *UpsertJwtCustomizerRequest) HasTokenSample() bool`

HasTokenSample returns a boolean if a field has been set.

### SetTokenSampleNil

`func (o *UpsertJwtCustomizerRequest) SetTokenSampleNil(b bool)`

 SetTokenSampleNil sets the value for TokenSample to be an explicit nil

### UnsetTokenSample
`func (o *UpsertJwtCustomizerRequest) UnsetTokenSample()`

UnsetTokenSample ensures that no value is present for TokenSample, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


