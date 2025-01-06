# ListConnectorFactories200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**IsDemo** | Pointer to **bool** |  | [optional] 
**Id** | **string** |  | 
**Target** | **string** |  | 
**Name** | **map[string]interface{}** | Validator function | 
**Description** | **map[string]interface{}** | Validator function | 
**Logo** | **string** |  | 
**LogoDark** | **NullableString** |  | 
**Readme** | **string** |  | 
**ConfigTemplate** | Pointer to **string** |  | [optional] 
**FormItems** | Pointer to [**[]ListConnectors200ResponseInnerFormItemsInner**](ListConnectors200ResponseInnerFormItemsInner.md) |  | [optional] 
**CustomData** | Pointer to **map[string]interface{}** |  | [optional] 
**FromEmail** | Pointer to **string** |  | [optional] 
**Platform** | **NullableString** |  | 
**IsStandard** | Pointer to **bool** |  | [optional] 

## Methods

### NewListConnectorFactories200ResponseInner

`func NewListConnectorFactories200ResponseInner(type_ string, id string, target string, name map[string]interface{}, description map[string]interface{}, logo string, logoDark NullableString, readme string, platform NullableString, ) *ListConnectorFactories200ResponseInner`

NewListConnectorFactories200ResponseInner instantiates a new ListConnectorFactories200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectorFactories200ResponseInnerWithDefaults

`func NewListConnectorFactories200ResponseInnerWithDefaults() *ListConnectorFactories200ResponseInner`

NewListConnectorFactories200ResponseInnerWithDefaults instantiates a new ListConnectorFactories200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListConnectorFactories200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListConnectorFactories200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListConnectorFactories200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetIsDemo

`func (o *ListConnectorFactories200ResponseInner) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *ListConnectorFactories200ResponseInner) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *ListConnectorFactories200ResponseInner) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *ListConnectorFactories200ResponseInner) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### GetId

`func (o *ListConnectorFactories200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListConnectorFactories200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListConnectorFactories200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetTarget

`func (o *ListConnectorFactories200ResponseInner) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ListConnectorFactories200ResponseInner) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ListConnectorFactories200ResponseInner) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetName

`func (o *ListConnectorFactories200ResponseInner) GetName() map[string]interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListConnectorFactories200ResponseInner) GetNameOk() (*map[string]interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListConnectorFactories200ResponseInner) SetName(v map[string]interface{})`

SetName sets Name field to given value.


### GetDescription

`func (o *ListConnectorFactories200ResponseInner) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListConnectorFactories200ResponseInner) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListConnectorFactories200ResponseInner) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.


### GetLogo

`func (o *ListConnectorFactories200ResponseInner) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ListConnectorFactories200ResponseInner) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ListConnectorFactories200ResponseInner) SetLogo(v string)`

SetLogo sets Logo field to given value.


### GetLogoDark

`func (o *ListConnectorFactories200ResponseInner) GetLogoDark() string`

GetLogoDark returns the LogoDark field if non-nil, zero value otherwise.

### GetLogoDarkOk

`func (o *ListConnectorFactories200ResponseInner) GetLogoDarkOk() (*string, bool)`

GetLogoDarkOk returns a tuple with the LogoDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoDark

`func (o *ListConnectorFactories200ResponseInner) SetLogoDark(v string)`

SetLogoDark sets LogoDark field to given value.


### SetLogoDarkNil

`func (o *ListConnectorFactories200ResponseInner) SetLogoDarkNil(b bool)`

 SetLogoDarkNil sets the value for LogoDark to be an explicit nil

### UnsetLogoDark
`func (o *ListConnectorFactories200ResponseInner) UnsetLogoDark()`

UnsetLogoDark ensures that no value is present for LogoDark, not even an explicit nil
### GetReadme

`func (o *ListConnectorFactories200ResponseInner) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *ListConnectorFactories200ResponseInner) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *ListConnectorFactories200ResponseInner) SetReadme(v string)`

SetReadme sets Readme field to given value.


### GetConfigTemplate

`func (o *ListConnectorFactories200ResponseInner) GetConfigTemplate() string`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *ListConnectorFactories200ResponseInner) GetConfigTemplateOk() (*string, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *ListConnectorFactories200ResponseInner) SetConfigTemplate(v string)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *ListConnectorFactories200ResponseInner) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### GetFormItems

`func (o *ListConnectorFactories200ResponseInner) GetFormItems() []ListConnectors200ResponseInnerFormItemsInner`

GetFormItems returns the FormItems field if non-nil, zero value otherwise.

### GetFormItemsOk

`func (o *ListConnectorFactories200ResponseInner) GetFormItemsOk() (*[]ListConnectors200ResponseInnerFormItemsInner, bool)`

GetFormItemsOk returns a tuple with the FormItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormItems

`func (o *ListConnectorFactories200ResponseInner) SetFormItems(v []ListConnectors200ResponseInnerFormItemsInner)`

SetFormItems sets FormItems field to given value.

### HasFormItems

`func (o *ListConnectorFactories200ResponseInner) HasFormItems() bool`

HasFormItems returns a boolean if a field has been set.

### GetCustomData

`func (o *ListConnectorFactories200ResponseInner) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *ListConnectorFactories200ResponseInner) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *ListConnectorFactories200ResponseInner) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *ListConnectorFactories200ResponseInner) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetFromEmail

`func (o *ListConnectorFactories200ResponseInner) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *ListConnectorFactories200ResponseInner) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *ListConnectorFactories200ResponseInner) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *ListConnectorFactories200ResponseInner) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetPlatform

`func (o *ListConnectorFactories200ResponseInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListConnectorFactories200ResponseInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListConnectorFactories200ResponseInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### SetPlatformNil

`func (o *ListConnectorFactories200ResponseInner) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *ListConnectorFactories200ResponseInner) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetIsStandard

`func (o *ListConnectorFactories200ResponseInner) GetIsStandard() bool`

GetIsStandard returns the IsStandard field if non-nil, zero value otherwise.

### GetIsStandardOk

`func (o *ListConnectorFactories200ResponseInner) GetIsStandardOk() (*bool, bool)`

GetIsStandardOk returns a tuple with the IsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandard

`func (o *ListConnectorFactories200ResponseInner) SetIsStandard(v bool)`

SetIsStandard sets IsStandard field to given value.

### HasIsStandard

`func (o *ListConnectorFactories200ResponseInner) HasIsStandard() bool`

HasIsStandard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


