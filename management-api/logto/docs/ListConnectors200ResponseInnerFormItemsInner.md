# ListConnectors200ResponseInnerFormItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**SelectItems** | [**[]ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner**](ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner.md) |  | 
**Key** | **string** |  | 
**Label** | **string** |  | 
**Placeholder** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**DefaultValue** | Pointer to **interface{}** |  | [optional] 
**ShowConditions** | Pointer to [**[]ListConnectors200ResponseInnerFormItemsInnerOneOfShowConditionsInner**](ListConnectors200ResponseInnerFormItemsInnerOneOfShowConditionsInner.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tooltip** | Pointer to **string** |  | [optional] 
**IsConfidential** | Pointer to **bool** |  | [optional] 

## Methods

### NewListConnectors200ResponseInnerFormItemsInner

`func NewListConnectors200ResponseInnerFormItemsInner(type_ string, selectItems []ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner, key string, label string, ) *ListConnectors200ResponseInnerFormItemsInner`

NewListConnectors200ResponseInnerFormItemsInner instantiates a new ListConnectors200ResponseInnerFormItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectors200ResponseInnerFormItemsInnerWithDefaults

`func NewListConnectors200ResponseInnerFormItemsInnerWithDefaults() *ListConnectors200ResponseInnerFormItemsInner`

NewListConnectors200ResponseInnerFormItemsInnerWithDefaults instantiates a new ListConnectors200ResponseInnerFormItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListConnectors200ResponseInnerFormItemsInner) SetType(v string)`

SetType sets Type field to given value.


### GetSelectItems

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetSelectItems() []ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner`

GetSelectItems returns the SelectItems field if non-nil, zero value otherwise.

### GetSelectItemsOk

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetSelectItemsOk() (*[]ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner, bool)`

GetSelectItemsOk returns a tuple with the SelectItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectItems

`func (o *ListConnectors200ResponseInnerFormItemsInner) SetSelectItems(v []ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner)`

SetSelectItems sets SelectItems field to given value.


### GetKey

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ListConnectors200ResponseInnerFormItemsInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ListConnectors200ResponseInnerFormItemsInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPlaceholder

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *ListConnectors200ResponseInnerFormItemsInner) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *ListConnectors200ResponseInnerFormItemsInner) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetRequired

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ListConnectors200ResponseInnerFormItemsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ListConnectors200ResponseInnerFormItemsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ListConnectors200ResponseInnerFormItemsInner) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ListConnectors200ResponseInnerFormItemsInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ListConnectors200ResponseInnerFormItemsInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ListConnectors200ResponseInnerFormItemsInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetShowConditions

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetShowConditions() []ListConnectors200ResponseInnerFormItemsInnerOneOfShowConditionsInner`

GetShowConditions returns the ShowConditions field if non-nil, zero value otherwise.

### GetShowConditionsOk

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetShowConditionsOk() (*[]ListConnectors200ResponseInnerFormItemsInnerOneOfShowConditionsInner, bool)`

GetShowConditionsOk returns a tuple with the ShowConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowConditions

`func (o *ListConnectors200ResponseInnerFormItemsInner) SetShowConditions(v []ListConnectors200ResponseInnerFormItemsInnerOneOfShowConditionsInner)`

SetShowConditions sets ShowConditions field to given value.

### HasShowConditions

`func (o *ListConnectors200ResponseInnerFormItemsInner) HasShowConditions() bool`

HasShowConditions returns a boolean if a field has been set.

### GetDescription

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListConnectors200ResponseInnerFormItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListConnectors200ResponseInnerFormItemsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTooltip

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *ListConnectors200ResponseInnerFormItemsInner) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *ListConnectors200ResponseInnerFormItemsInner) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.

### GetIsConfidential

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetIsConfidential() bool`

GetIsConfidential returns the IsConfidential field if non-nil, zero value otherwise.

### GetIsConfidentialOk

`func (o *ListConnectors200ResponseInnerFormItemsInner) GetIsConfidentialOk() (*bool, bool)`

GetIsConfidentialOk returns a tuple with the IsConfidential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfidential

`func (o *ListConnectors200ResponseInnerFormItemsInner) SetIsConfidential(v bool)`

SetIsConfidential sets IsConfidential field to given value.

### HasIsConfidential

`func (o *ListConnectors200ResponseInnerFormItemsInner) HasIsConfidential() bool`

HasIsConfidential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


