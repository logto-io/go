# UpdateAccountCenterSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable or disable the account API. | [optional] 
**Fields** | Pointer to [**UpdateAccountCenterSettingsRequestFields**](UpdateAccountCenterSettingsRequestFields.md) |  | [optional] 

## Methods

### NewUpdateAccountCenterSettingsRequest

`func NewUpdateAccountCenterSettingsRequest() *UpdateAccountCenterSettingsRequest`

NewUpdateAccountCenterSettingsRequest instantiates a new UpdateAccountCenterSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountCenterSettingsRequestWithDefaults

`func NewUpdateAccountCenterSettingsRequestWithDefaults() *UpdateAccountCenterSettingsRequest`

NewUpdateAccountCenterSettingsRequestWithDefaults instantiates a new UpdateAccountCenterSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateAccountCenterSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateAccountCenterSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateAccountCenterSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateAccountCenterSettingsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFields

`func (o *UpdateAccountCenterSettingsRequest) GetFields() UpdateAccountCenterSettingsRequestFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *UpdateAccountCenterSettingsRequest) GetFieldsOk() (*UpdateAccountCenterSettingsRequestFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *UpdateAccountCenterSettingsRequest) SetFields(v UpdateAccountCenterSettingsRequestFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *UpdateAccountCenterSettingsRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


