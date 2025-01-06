/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logto

import (
	"encoding/json"
)

// checks if the CreateConnectorRequestMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConnectorRequestMetadata{}

// CreateConnectorRequestMetadata Custom connector metadata, will be used to overwrite the default connector factory metadata.
type CreateConnectorRequestMetadata struct {
	Target *string `json:"target,omitempty"`
	// Validator function
	Name map[string]interface{} `json:"name,omitempty"`
	Logo *string `json:"logo,omitempty"`
	LogoDark NullableString `json:"logoDark,omitempty"`
}

// NewCreateConnectorRequestMetadata instantiates a new CreateConnectorRequestMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConnectorRequestMetadata() *CreateConnectorRequestMetadata {
	this := CreateConnectorRequestMetadata{}
	return &this
}

// NewCreateConnectorRequestMetadataWithDefaults instantiates a new CreateConnectorRequestMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConnectorRequestMetadataWithDefaults() *CreateConnectorRequestMetadata {
	this := CreateConnectorRequestMetadata{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *CreateConnectorRequestMetadata) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequestMetadata) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *CreateConnectorRequestMetadata) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *CreateConnectorRequestMetadata) SetTarget(v string) {
	o.Target = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateConnectorRequestMetadata) GetName() map[string]interface{} {
	if o == nil || IsNil(o.Name) {
		var ret map[string]interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequestMetadata) GetNameOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return map[string]interface{}{}, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateConnectorRequestMetadata) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given map[string]interface{} and assigns it to the Name field.
func (o *CreateConnectorRequestMetadata) SetName(v map[string]interface{}) {
	o.Name = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CreateConnectorRequestMetadata) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequestMetadata) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CreateConnectorRequestMetadata) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *CreateConnectorRequestMetadata) SetLogo(v string) {
	o.Logo = &v
}

// GetLogoDark returns the LogoDark field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateConnectorRequestMetadata) GetLogoDark() string {
	if o == nil || IsNil(o.LogoDark.Get()) {
		var ret string
		return ret
	}
	return *o.LogoDark.Get()
}

// GetLogoDarkOk returns a tuple with the LogoDark field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateConnectorRequestMetadata) GetLogoDarkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoDark.Get(), o.LogoDark.IsSet()
}

// HasLogoDark returns a boolean if a field has been set.
func (o *CreateConnectorRequestMetadata) HasLogoDark() bool {
	if o != nil && o.LogoDark.IsSet() {
		return true
	}

	return false
}

// SetLogoDark gets a reference to the given NullableString and assigns it to the LogoDark field.
func (o *CreateConnectorRequestMetadata) SetLogoDark(v string) {
	o.LogoDark.Set(&v)
}
// SetLogoDarkNil sets the value for LogoDark to be an explicit nil
func (o *CreateConnectorRequestMetadata) SetLogoDarkNil() {
	o.LogoDark.Set(nil)
}

// UnsetLogoDark ensures that no value is present for LogoDark, not even an explicit nil
func (o *CreateConnectorRequestMetadata) UnsetLogoDark() {
	o.LogoDark.Unset()
}

func (o CreateConnectorRequestMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateConnectorRequestMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if o.LogoDark.IsSet() {
		toSerialize["logoDark"] = o.LogoDark.Get()
	}
	return toSerialize, nil
}

type NullableCreateConnectorRequestMetadata struct {
	value *CreateConnectorRequestMetadata
	isSet bool
}

func (v NullableCreateConnectorRequestMetadata) Get() *CreateConnectorRequestMetadata {
	return v.value
}

func (v *NullableCreateConnectorRequestMetadata) Set(val *CreateConnectorRequestMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConnectorRequestMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConnectorRequestMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConnectorRequestMetadata(val *CreateConnectorRequestMetadata) *NullableCreateConnectorRequestMetadata {
	return &NullableCreateConnectorRequestMetadata{value: val, isSet: true}
}

func (v NullableCreateConnectorRequestMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConnectorRequestMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


