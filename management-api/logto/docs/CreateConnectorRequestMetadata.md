# CreateConnectorRequestMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **map[string]interface{}** | Validator function | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**LogoDark** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateConnectorRequestMetadata

`func NewCreateConnectorRequestMetadata() *CreateConnectorRequestMetadata`

NewCreateConnectorRequestMetadata instantiates a new CreateConnectorRequestMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectorRequestMetadataWithDefaults

`func NewCreateConnectorRequestMetadataWithDefaults() *CreateConnectorRequestMetadata`

NewCreateConnectorRequestMetadataWithDefaults instantiates a new CreateConnectorRequestMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *CreateConnectorRequestMetadata) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CreateConnectorRequestMetadata) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CreateConnectorRequestMetadata) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CreateConnectorRequestMetadata) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetName

`func (o *CreateConnectorRequestMetadata) GetName() map[string]interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConnectorRequestMetadata) GetNameOk() (*map[string]interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConnectorRequestMetadata) SetName(v map[string]interface{})`

SetName sets Name field to given value.

### HasName

`func (o *CreateConnectorRequestMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogo

`func (o *CreateConnectorRequestMetadata) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CreateConnectorRequestMetadata) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CreateConnectorRequestMetadata) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CreateConnectorRequestMetadata) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoDark

`func (o *CreateConnectorRequestMetadata) GetLogoDark() string`

GetLogoDark returns the LogoDark field if non-nil, zero value otherwise.

### GetLogoDarkOk

`func (o *CreateConnectorRequestMetadata) GetLogoDarkOk() (*string, bool)`

GetLogoDarkOk returns a tuple with the LogoDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoDark

`func (o *CreateConnectorRequestMetadata) SetLogoDark(v string)`

SetLogoDark sets LogoDark field to given value.

### HasLogoDark

`func (o *CreateConnectorRequestMetadata) HasLogoDark() bool`

HasLogoDark returns a boolean if a field has been set.

### SetLogoDarkNil

`func (o *CreateConnectorRequestMetadata) SetLogoDarkNil(b bool)`

 SetLogoDarkNil sets the value for LogoDark to be an explicit nil

### UnsetLogoDark
`func (o *CreateConnectorRequestMetadata) UnsetLogoDark()`

UnsetLogoDark ensures that no value is present for LogoDark, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


