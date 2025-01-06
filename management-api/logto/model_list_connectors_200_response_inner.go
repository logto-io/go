/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logto

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListConnectors200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConnectors200ResponseInner{}

// ListConnectors200ResponseInner struct for ListConnectors200ResponseInner
type ListConnectors200ResponseInner struct {
	Id string `json:"id"`
	SyncProfile bool `json:"syncProfile"`
	// arbitrary
	Config map[string]interface{} `json:"config"`
	Metadata ListConnectors200ResponseInnerMetadata `json:"metadata"`
	ConnectorId string `json:"connectorId"`
	Target string `json:"target"`
	// Validator function
	Name map[string]interface{} `json:"name"`
	// Validator function
	Description map[string]interface{} `json:"description"`
	Logo string `json:"logo"`
	LogoDark NullableString `json:"logoDark"`
	Readme string `json:"readme"`
	ConfigTemplate *string `json:"configTemplate,omitempty"`
	FormItems []ListConnectors200ResponseInnerFormItemsInner `json:"formItems,omitempty"`
	CustomData map[string]interface{} `json:"customData,omitempty"`
	FromEmail *string `json:"fromEmail,omitempty"`
	Platform NullableString `json:"platform"`
	IsStandard *bool `json:"isStandard,omitempty"`
	Type string `json:"type"`
	IsDemo *bool `json:"isDemo,omitempty"`
	ExtraInfo map[string]interface{} `json:"extraInfo,omitempty"`
	Usage *float32 `json:"usage,omitempty"`
}

type _ListConnectors200ResponseInner ListConnectors200ResponseInner

// NewListConnectors200ResponseInner instantiates a new ListConnectors200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectors200ResponseInner(id string, syncProfile bool, config map[string]interface{}, metadata ListConnectors200ResponseInnerMetadata, connectorId string, target string, name map[string]interface{}, description map[string]interface{}, logo string, logoDark NullableString, readme string, platform NullableString, type_ string) *ListConnectors200ResponseInner {
	this := ListConnectors200ResponseInner{}
	this.Id = id
	this.SyncProfile = syncProfile
	this.Config = config
	this.Metadata = metadata
	this.ConnectorId = connectorId
	this.Target = target
	this.Name = name
	this.Description = description
	this.Logo = logo
	this.LogoDark = logoDark
	this.Readme = readme
	this.Platform = platform
	this.Type = type_
	return &this
}

// NewListConnectors200ResponseInnerWithDefaults instantiates a new ListConnectors200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectors200ResponseInnerWithDefaults() *ListConnectors200ResponseInner {
	this := ListConnectors200ResponseInner{}
	return &this
}

// GetId returns the Id field value
func (o *ListConnectors200ResponseInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListConnectors200ResponseInner) SetId(v string) {
	o.Id = v
}

// GetSyncProfile returns the SyncProfile field value
func (o *ListConnectors200ResponseInner) GetSyncProfile() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SyncProfile
}

// GetSyncProfileOk returns a tuple with the SyncProfile field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetSyncProfileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncProfile, true
}

// SetSyncProfile sets field value
func (o *ListConnectors200ResponseInner) SetSyncProfile(v bool) {
	o.SyncProfile = v
}

// GetConfig returns the Config field value
func (o *ListConnectors200ResponseInner) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// SetConfig sets field value
func (o *ListConnectors200ResponseInner) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetMetadata returns the Metadata field value
func (o *ListConnectors200ResponseInner) GetMetadata() ListConnectors200ResponseInnerMetadata {
	if o == nil {
		var ret ListConnectors200ResponseInnerMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetMetadataOk() (*ListConnectors200ResponseInnerMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ListConnectors200ResponseInner) SetMetadata(v ListConnectors200ResponseInnerMetadata) {
	o.Metadata = v
}

// GetConnectorId returns the ConnectorId field value
func (o *ListConnectors200ResponseInner) GetConnectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetConnectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *ListConnectors200ResponseInner) SetConnectorId(v string) {
	o.ConnectorId = v
}

// GetTarget returns the Target field value
func (o *ListConnectors200ResponseInner) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *ListConnectors200ResponseInner) SetTarget(v string) {
	o.Target = v
}

// GetName returns the Name field value
func (o *ListConnectors200ResponseInner) GetName() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetNameOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *ListConnectors200ResponseInner) SetName(v map[string]interface{}) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ListConnectors200ResponseInner) GetDescription() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetDescriptionOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Description, true
}

// SetDescription sets field value
func (o *ListConnectors200ResponseInner) SetDescription(v map[string]interface{}) {
	o.Description = v
}

// GetLogo returns the Logo field value
func (o *ListConnectors200ResponseInner) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value
func (o *ListConnectors200ResponseInner) SetLogo(v string) {
	o.Logo = v
}

// GetLogoDark returns the LogoDark field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListConnectors200ResponseInner) GetLogoDark() string {
	if o == nil || o.LogoDark.Get() == nil {
		var ret string
		return ret
	}

	return *o.LogoDark.Get()
}

// GetLogoDarkOk returns a tuple with the LogoDark field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListConnectors200ResponseInner) GetLogoDarkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoDark.Get(), o.LogoDark.IsSet()
}

// SetLogoDark sets field value
func (o *ListConnectors200ResponseInner) SetLogoDark(v string) {
	o.LogoDark.Set(&v)
}

// GetReadme returns the Readme field value
func (o *ListConnectors200ResponseInner) GetReadme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Readme
}

// GetReadmeOk returns a tuple with the Readme field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetReadmeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Readme, true
}

// SetReadme sets field value
func (o *ListConnectors200ResponseInner) SetReadme(v string) {
	o.Readme = v
}

// GetConfigTemplate returns the ConfigTemplate field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInner) GetConfigTemplate() string {
	if o == nil || IsNil(o.ConfigTemplate) {
		var ret string
		return ret
	}
	return *o.ConfigTemplate
}

// GetConfigTemplateOk returns a tuple with the ConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetConfigTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigTemplate) {
		return nil, false
	}
	return o.ConfigTemplate, true
}

// HasConfigTemplate returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInner) HasConfigTemplate() bool {
	if o != nil && !IsNil(o.ConfigTemplate) {
		return true
	}

	return false
}

// SetConfigTemplate gets a reference to the given string and assigns it to the ConfigTemplate field.
func (o *ListConnectors200ResponseInner) SetConfigTemplate(v string) {
	o.ConfigTemplate = &v
}

// GetFormItems returns the FormItems field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInner) GetFormItems() []ListConnectors200ResponseInnerFormItemsInner {
	if o == nil || IsNil(o.FormItems) {
		var ret []ListConnectors200ResponseInnerFormItemsInner
		return ret
	}
	return o.FormItems
}

// GetFormItemsOk returns a tuple with the FormItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetFormItemsOk() ([]ListConnectors200ResponseInnerFormItemsInner, bool) {
	if o == nil || IsNil(o.FormItems) {
		return nil, false
	}
	return o.FormItems, true
}

// HasFormItems returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInner) HasFormItems() bool {
	if o != nil && !IsNil(o.FormItems) {
		return true
	}

	return false
}

// SetFormItems gets a reference to the given []ListConnectors200ResponseInnerFormItemsInner and assigns it to the FormItems field.
func (o *ListConnectors200ResponseInner) SetFormItems(v []ListConnectors200ResponseInnerFormItemsInner) {
	o.FormItems = v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInner) GetCustomData() map[string]interface{} {
	if o == nil || IsNil(o.CustomData) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetCustomDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomData) {
		return map[string]interface{}{}, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInner) HasCustomData() bool {
	if o != nil && !IsNil(o.CustomData) {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given map[string]interface{} and assigns it to the CustomData field.
func (o *ListConnectors200ResponseInner) SetCustomData(v map[string]interface{}) {
	o.CustomData = v
}

// GetFromEmail returns the FromEmail field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInner) GetFromEmail() string {
	if o == nil || IsNil(o.FromEmail) {
		var ret string
		return ret
	}
	return *o.FromEmail
}

// GetFromEmailOk returns a tuple with the FromEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetFromEmailOk() (*string, bool) {
	if o == nil || IsNil(o.FromEmail) {
		return nil, false
	}
	return o.FromEmail, true
}

// HasFromEmail returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInner) HasFromEmail() bool {
	if o != nil && !IsNil(o.FromEmail) {
		return true
	}

	return false
}

// SetFromEmail gets a reference to the given string and assigns it to the FromEmail field.
func (o *ListConnectors200ResponseInner) SetFromEmail(v string) {
	o.FromEmail = &v
}

// GetPlatform returns the Platform field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListConnectors200ResponseInner) GetPlatform() string {
	if o == nil || o.Platform.Get() == nil {
		var ret string
		return ret
	}

	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListConnectors200ResponseInner) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// SetPlatform sets field value
func (o *ListConnectors200ResponseInner) SetPlatform(v string) {
	o.Platform.Set(&v)
}

// GetIsStandard returns the IsStandard field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInner) GetIsStandard() bool {
	if o == nil || IsNil(o.IsStandard) {
		var ret bool
		return ret
	}
	return *o.IsStandard
}

// GetIsStandardOk returns a tuple with the IsStandard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetIsStandardOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStandard) {
		return nil, false
	}
	return o.IsStandard, true
}

// HasIsStandard returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInner) HasIsStandard() bool {
	if o != nil && !IsNil(o.IsStandard) {
		return true
	}

	return false
}

// SetIsStandard gets a reference to the given bool and assigns it to the IsStandard field.
func (o *ListConnectors200ResponseInner) SetIsStandard(v bool) {
	o.IsStandard = &v
}

// GetType returns the Type field value
func (o *ListConnectors200ResponseInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListConnectors200ResponseInner) SetType(v string) {
	o.Type = v
}

// GetIsDemo returns the IsDemo field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInner) GetIsDemo() bool {
	if o == nil || IsNil(o.IsDemo) {
		var ret bool
		return ret
	}
	return *o.IsDemo
}

// GetIsDemoOk returns a tuple with the IsDemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetIsDemoOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDemo) {
		return nil, false
	}
	return o.IsDemo, true
}

// HasIsDemo returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInner) HasIsDemo() bool {
	if o != nil && !IsNil(o.IsDemo) {
		return true
	}

	return false
}

// SetIsDemo gets a reference to the given bool and assigns it to the IsDemo field.
func (o *ListConnectors200ResponseInner) SetIsDemo(v bool) {
	o.IsDemo = &v
}

// GetExtraInfo returns the ExtraInfo field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInner) GetExtraInfo() map[string]interface{} {
	if o == nil || IsNil(o.ExtraInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraInfo
}

// GetExtraInfoOk returns a tuple with the ExtraInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetExtraInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtraInfo) {
		return map[string]interface{}{}, false
	}
	return o.ExtraInfo, true
}

// HasExtraInfo returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInner) HasExtraInfo() bool {
	if o != nil && !IsNil(o.ExtraInfo) {
		return true
	}

	return false
}

// SetExtraInfo gets a reference to the given map[string]interface{} and assigns it to the ExtraInfo field.
func (o *ListConnectors200ResponseInner) SetExtraInfo(v map[string]interface{}) {
	o.ExtraInfo = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInner) GetUsage() float32 {
	if o == nil || IsNil(o.Usage) {
		var ret float32
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInner) GetUsageOk() (*float32, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInner) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given float32 and assigns it to the Usage field.
func (o *ListConnectors200ResponseInner) SetUsage(v float32) {
	o.Usage = &v
}

func (o ListConnectors200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConnectors200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["syncProfile"] = o.SyncProfile
	toSerialize["config"] = o.Config
	toSerialize["metadata"] = o.Metadata
	toSerialize["connectorId"] = o.ConnectorId
	toSerialize["target"] = o.Target
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["logo"] = o.Logo
	toSerialize["logoDark"] = o.LogoDark.Get()
	toSerialize["readme"] = o.Readme
	if !IsNil(o.ConfigTemplate) {
		toSerialize["configTemplate"] = o.ConfigTemplate
	}
	if !IsNil(o.FormItems) {
		toSerialize["formItems"] = o.FormItems
	}
	if !IsNil(o.CustomData) {
		toSerialize["customData"] = o.CustomData
	}
	if !IsNil(o.FromEmail) {
		toSerialize["fromEmail"] = o.FromEmail
	}
	toSerialize["platform"] = o.Platform.Get()
	if !IsNil(o.IsStandard) {
		toSerialize["isStandard"] = o.IsStandard
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.IsDemo) {
		toSerialize["isDemo"] = o.IsDemo
	}
	if !IsNil(o.ExtraInfo) {
		toSerialize["extraInfo"] = o.ExtraInfo
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

func (o *ListConnectors200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"syncProfile",
		"config",
		"metadata",
		"connectorId",
		"target",
		"name",
		"description",
		"logo",
		"logoDark",
		"readme",
		"platform",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListConnectors200ResponseInner := _ListConnectors200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListConnectors200ResponseInner)

	if err != nil {
		return err
	}

	*o = ListConnectors200ResponseInner(varListConnectors200ResponseInner)

	return err
}

type NullableListConnectors200ResponseInner struct {
	value *ListConnectors200ResponseInner
	isSet bool
}

func (v NullableListConnectors200ResponseInner) Get() *ListConnectors200ResponseInner {
	return v.value
}

func (v *NullableListConnectors200ResponseInner) Set(val *ListConnectors200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectors200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectors200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectors200ResponseInner(val *ListConnectors200ResponseInner) *NullableListConnectors200ResponseInner {
	return &NullableListConnectors200ResponseInner{value: val, isSet: true}
}

func (v NullableListConnectors200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectors200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


