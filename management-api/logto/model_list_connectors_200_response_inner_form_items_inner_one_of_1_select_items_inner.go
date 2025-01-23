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

// checks if the ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner{}

// ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner struct for ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner
type ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner struct {
	Value string `json:"value"`
}

type _ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner

// NewListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner instantiates a new ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner(value string) *ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner {
	this := ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner{}
	this.Value = value
	return &this
}

// NewListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInnerWithDefaults instantiates a new ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInnerWithDefaults() *ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner {
	this := ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner{}
	return &this
}

// GetValue returns the Value field value
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) SetValue(v string) {
	o.Value = v
}

func (o ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner := _ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner)

	if err != nil {
		return err
	}

	*o = ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner(varListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner)

	return err
}

type NullableListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner struct {
	value *ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner
	isSet bool
}

func (v NullableListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) Get() *ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner {
	return v.value
}

func (v *NullableListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) Set(val *ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner(val *ListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) *NullableListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner {
	return &NullableListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner{value: val, isSet: true}
}

func (v NullableListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectors200ResponseInnerFormItemsInnerOneOf1SelectItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


