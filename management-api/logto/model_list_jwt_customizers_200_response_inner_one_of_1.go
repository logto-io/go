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

// checks if the ListJwtCustomizers200ResponseInnerOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListJwtCustomizers200ResponseInnerOneOf1{}

// ListJwtCustomizers200ResponseInnerOneOf1 struct for ListJwtCustomizers200ResponseInnerOneOf1
type ListJwtCustomizers200ResponseInnerOneOf1 struct {
	Key string `json:"key"`
	Value GetJwtCustomizer200ResponseOneOf1 `json:"value"`
}

type _ListJwtCustomizers200ResponseInnerOneOf1 ListJwtCustomizers200ResponseInnerOneOf1

// NewListJwtCustomizers200ResponseInnerOneOf1 instantiates a new ListJwtCustomizers200ResponseInnerOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListJwtCustomizers200ResponseInnerOneOf1(key string, value GetJwtCustomizer200ResponseOneOf1) *ListJwtCustomizers200ResponseInnerOneOf1 {
	this := ListJwtCustomizers200ResponseInnerOneOf1{}
	this.Key = key
	this.Value = value
	return &this
}

// NewListJwtCustomizers200ResponseInnerOneOf1WithDefaults instantiates a new ListJwtCustomizers200ResponseInnerOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListJwtCustomizers200ResponseInnerOneOf1WithDefaults() *ListJwtCustomizers200ResponseInnerOneOf1 {
	this := ListJwtCustomizers200ResponseInnerOneOf1{}
	return &this
}

// GetKey returns the Key field value
func (o *ListJwtCustomizers200ResponseInnerOneOf1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ListJwtCustomizers200ResponseInnerOneOf1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ListJwtCustomizers200ResponseInnerOneOf1) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *ListJwtCustomizers200ResponseInnerOneOf1) GetValue() GetJwtCustomizer200ResponseOneOf1 {
	if o == nil {
		var ret GetJwtCustomizer200ResponseOneOf1
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListJwtCustomizers200ResponseInnerOneOf1) GetValueOk() (*GetJwtCustomizer200ResponseOneOf1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListJwtCustomizers200ResponseInnerOneOf1) SetValue(v GetJwtCustomizer200ResponseOneOf1) {
	o.Value = v
}

func (o ListJwtCustomizers200ResponseInnerOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListJwtCustomizers200ResponseInnerOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ListJwtCustomizers200ResponseInnerOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
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

	varListJwtCustomizers200ResponseInnerOneOf1 := _ListJwtCustomizers200ResponseInnerOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListJwtCustomizers200ResponseInnerOneOf1)

	if err != nil {
		return err
	}

	*o = ListJwtCustomizers200ResponseInnerOneOf1(varListJwtCustomizers200ResponseInnerOneOf1)

	return err
}

type NullableListJwtCustomizers200ResponseInnerOneOf1 struct {
	value *ListJwtCustomizers200ResponseInnerOneOf1
	isSet bool
}

func (v NullableListJwtCustomizers200ResponseInnerOneOf1) Get() *ListJwtCustomizers200ResponseInnerOneOf1 {
	return v.value
}

func (v *NullableListJwtCustomizers200ResponseInnerOneOf1) Set(val *ListJwtCustomizers200ResponseInnerOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableListJwtCustomizers200ResponseInnerOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableListJwtCustomizers200ResponseInnerOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListJwtCustomizers200ResponseInnerOneOf1(val *ListJwtCustomizers200ResponseInnerOneOf1) *NullableListJwtCustomizers200ResponseInnerOneOf1 {
	return &NullableListJwtCustomizers200ResponseInnerOneOf1{value: val, isSet: true}
}

func (v NullableListJwtCustomizers200ResponseInnerOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListJwtCustomizers200ResponseInnerOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


