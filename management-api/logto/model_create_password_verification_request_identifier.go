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

// checks if the CreatePasswordVerificationRequestIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePasswordVerificationRequestIdentifier{}

// CreatePasswordVerificationRequestIdentifier The unique identifier of the user that will be used to identify the user along with the provided password.
type CreatePasswordVerificationRequestIdentifier struct {
	Type string `json:"type"`
	Value string `json:"value"`
}

type _CreatePasswordVerificationRequestIdentifier CreatePasswordVerificationRequestIdentifier

// NewCreatePasswordVerificationRequestIdentifier instantiates a new CreatePasswordVerificationRequestIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePasswordVerificationRequestIdentifier(type_ string, value string) *CreatePasswordVerificationRequestIdentifier {
	this := CreatePasswordVerificationRequestIdentifier{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewCreatePasswordVerificationRequestIdentifierWithDefaults instantiates a new CreatePasswordVerificationRequestIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePasswordVerificationRequestIdentifierWithDefaults() *CreatePasswordVerificationRequestIdentifier {
	this := CreatePasswordVerificationRequestIdentifier{}
	return &this
}

// GetType returns the Type field value
func (o *CreatePasswordVerificationRequestIdentifier) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreatePasswordVerificationRequestIdentifier) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreatePasswordVerificationRequestIdentifier) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *CreatePasswordVerificationRequestIdentifier) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CreatePasswordVerificationRequestIdentifier) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CreatePasswordVerificationRequestIdentifier) SetValue(v string) {
	o.Value = v
}

func (o CreatePasswordVerificationRequestIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePasswordVerificationRequestIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *CreatePasswordVerificationRequestIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varCreatePasswordVerificationRequestIdentifier := _CreatePasswordVerificationRequestIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreatePasswordVerificationRequestIdentifier)

	if err != nil {
		return err
	}

	*o = CreatePasswordVerificationRequestIdentifier(varCreatePasswordVerificationRequestIdentifier)

	return err
}

type NullableCreatePasswordVerificationRequestIdentifier struct {
	value *CreatePasswordVerificationRequestIdentifier
	isSet bool
}

func (v NullableCreatePasswordVerificationRequestIdentifier) Get() *CreatePasswordVerificationRequestIdentifier {
	return v.value
}

func (v *NullableCreatePasswordVerificationRequestIdentifier) Set(val *CreatePasswordVerificationRequestIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePasswordVerificationRequestIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePasswordVerificationRequestIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePasswordVerificationRequestIdentifier(val *CreatePasswordVerificationRequestIdentifier) *NullableCreatePasswordVerificationRequestIdentifier {
	return &NullableCreatePasswordVerificationRequestIdentifier{value: val, isSet: true}
}

func (v NullableCreatePasswordVerificationRequestIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePasswordVerificationRequestIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


