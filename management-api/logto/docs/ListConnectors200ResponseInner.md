# ListConnectors200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SyncProfile** | **bool** |  | 
**Config** | **map[string]interface{}** | arbitrary | 
**Metadata** | [**ListConnectors200ResponseInnerMetadata**](ListConnectors200ResponseInnerMetadata.md) |  | 
**ConnectorId** | **string** |  | 
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
**Type** | **string** |  | 
**IsDemo** | Pointer to **bool** |  | [optional] 
**ExtraInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**Usage** | Pointer to **float32** |  | [optional] 

## Methods

### NewListConnectors200ResponseInner

`func NewListConnectors200ResponseInner(id string, syncProfile bool, config map[string]interface{}, metadata ListConnectors200ResponseInnerMetadata, connectorId string, target string, name map[string]interface{}, description map[string]interface{}, logo string, logoDark NullableString, readme string, platform NullableString, type_ string, ) *ListConnectors200ResponseInner`

NewListConnectors200ResponseInner instantiates a new ListConnectors200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectors200ResponseInnerWithDefaults

`func NewListConnectors200ResponseInnerWithDefaults() *ListConnectors200ResponseInner`

NewListConnectors200ResponseInnerWithDefaults instantiates a new ListConnectors200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListConnectors200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListConnectors200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListConnectors200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetSyncProfile

`func (o *ListConnectors200ResponseInner) GetSyncProfile() bool`

GetSyncProfile returns the SyncProfile field if non-nil, zero value otherwise.

### GetSyncProfileOk

`func (o *ListConnectors200ResponseInner) GetSyncProfileOk() (*bool, bool)`

GetSyncProfileOk returns a tuple with the SyncProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProfile

`func (o *ListConnectors200ResponseInner) SetSyncProfile(v bool)`

SetSyncProfile sets SyncProfile field to given value.


### GetConfig

`func (o *ListConnectors200ResponseInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListConnectors200ResponseInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListConnectors200ResponseInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetMetadata

`func (o *ListConnectors200ResponseInner) GetMetadata() ListConnectors200ResponseInnerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListConnectors200ResponseInner) GetMetadataOk() (*ListConnectors200ResponseInnerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListConnectors200ResponseInner) SetMetadata(v ListConnectors200ResponseInnerMetadata)`

SetMetadata sets Metadata field to given value.


### GetConnectorId

`func (o *ListConnectors200ResponseInner) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *ListConnectors200ResponseInner) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *ListConnectors200ResponseInner) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.


### GetTarget

`func (o *ListConnectors200ResponseInner) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ListConnectors200ResponseInner) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ListConnectors200ResponseInner) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetName

`func (o *ListConnectors200ResponseInner) GetName() map[string]interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListConnectors200ResponseInner) GetNameOk() (*map[string]interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListConnectors200ResponseInner) SetName(v map[string]interface{})`

SetName sets Name field to given value.


### GetDescription

`func (o *ListConnectors200ResponseInner) GetDescription() map[string]interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListConnectors200ResponseInner) GetDescriptionOk() (*map[string]interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListConnectors200ResponseInner) SetDescription(v map[string]interface{})`

SetDescription sets Description field to given value.


### GetLogo

`func (o *ListConnectors200ResponseInner) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ListConnectors200ResponseInner) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ListConnectors200ResponseInner) SetLogo(v string)`

SetLogo sets Logo field to given value.


### GetLogoDark

`func (o *ListConnectors200ResponseInner) GetLogoDark() string`

GetLogoDark returns the LogoDark field if non-nil, zero value otherwise.

### GetLogoDarkOk

`func (o *ListConnectors200ResponseInner) GetLogoDarkOk() (*string, bool)`

GetLogoDarkOk returns a tuple with the LogoDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoDark

`func (o *ListConnectors200ResponseInner) SetLogoDark(v string)`

SetLogoDark sets LogoDark field to given value.


### SetLogoDarkNil

`func (o *ListConnectors200ResponseInner) SetLogoDarkNil(b bool)`

 SetLogoDarkNil sets the value for LogoDark to be an explicit nil

### UnsetLogoDark
`func (o *ListConnectors200ResponseInner) UnsetLogoDark()`

UnsetLogoDark ensures that no value is present for LogoDark, not even an explicit nil
### GetReadme

`func (o *ListConnectors200ResponseInner) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *ListConnectors200ResponseInner) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *ListConnectors200ResponseInner) SetReadme(v string)`

SetReadme sets Readme field to given value.


### GetConfigTemplate

`func (o *ListConnectors200ResponseInner) GetConfigTemplate() string`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *ListConnectors200ResponseInner) GetConfigTemplateOk() (*string, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *ListConnectors200ResponseInner) SetConfigTemplate(v string)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *ListConnectors200ResponseInner) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### GetFormItems

`func (o *ListConnectors200ResponseInner) GetFormItems() []ListConnectors200ResponseInnerFormItemsInner`

GetFormItems returns the FormItems field if non-nil, zero value otherwise.

### GetFormItemsOk

`func (o *ListConnectors200ResponseInner) GetFormItemsOk() (*[]ListConnectors200ResponseInnerFormItemsInner, bool)`

GetFormItemsOk returns a tuple with the FormItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormItems

`func (o *ListConnectors200ResponseInner) SetFormItems(v []ListConnectors200ResponseInnerFormItemsInner)`

SetFormItems sets FormItems field to given value.

### HasFormItems

`func (o *ListConnectors200ResponseInner) HasFormItems() bool`

HasFormItems returns a boolean if a field has been set.

### GetCustomData

`func (o *ListConnectors200ResponseInner) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *ListConnectors200ResponseInner) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *ListConnectors200ResponseInner) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *ListConnectors200ResponseInner) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetFromEmail

`func (o *ListConnectors200ResponseInner) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *ListConnectors200ResponseInner) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *ListConnectors200ResponseInner) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *ListConnectors200ResponseInner) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetPlatform

`func (o *ListConnectors200ResponseInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListConnectors200ResponseInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListConnectors200ResponseInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### SetPlatformNil

`func (o *ListConnectors200ResponseInner) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *ListConnectors200ResponseInner) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetIsStandard

`func (o *ListConnectors200ResponseInner) GetIsStandard() bool`

GetIsStandard returns the IsStandard field if non-nil, zero value otherwise.

### GetIsStandardOk

`func (o *ListConnectors200ResponseInner) GetIsStandardOk() (*bool, bool)`

GetIsStandardOk returns a tuple with the IsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandard

`func (o *ListConnectors200ResponseInner) SetIsStandard(v bool)`

SetIsStandard sets IsStandard field to given value.

### HasIsStandard

`func (o *ListConnectors200ResponseInner) HasIsStandard() bool`

HasIsStandard returns a boolean if a field has been set.

### GetType

`func (o *ListConnectors200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListConnectors200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListConnectors200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetIsDemo

`func (o *ListConnectors200ResponseInner) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *ListConnectors200ResponseInner) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *ListConnectors200ResponseInner) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *ListConnectors200ResponseInner) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### GetExtraInfo

`func (o *ListConnectors200ResponseInner) GetExtraInfo() map[string]interface{}`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *ListConnectors200ResponseInner) GetExtraInfoOk() (*map[string]interface{}, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *ListConnectors200ResponseInner) SetExtraInfo(v map[string]interface{})`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *ListConnectors200ResponseInner) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetUsage

`func (o *ListConnectors200ResponseInner) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ListConnectors200ResponseInner) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ListConnectors200ResponseInner) SetUsage(v float32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ListConnectors200ResponseInner) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


