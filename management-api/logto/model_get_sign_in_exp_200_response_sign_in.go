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

// checks if the GetSignInExp200ResponseSignIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSignInExp200ResponseSignIn{}

// GetSignInExp200ResponseSignIn Sign-in method settings.
type GetSignInExp200ResponseSignIn struct {
	Methods []GetSignInExp200ResponseSignInMethodsInner `json:"methods"`
}

type _GetSignInExp200ResponseSignIn GetSignInExp200ResponseSignIn

// NewGetSignInExp200ResponseSignIn instantiates a new GetSignInExp200ResponseSignIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignInExp200ResponseSignIn(methods []GetSignInExp200ResponseSignInMethodsInner) *GetSignInExp200ResponseSignIn {
	this := GetSignInExp200ResponseSignIn{}
	this.Methods = methods
	return &this
}

// NewGetSignInExp200ResponseSignInWithDefaults instantiates a new GetSignInExp200ResponseSignIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignInExp200ResponseSignInWithDefaults() *GetSignInExp200ResponseSignIn {
	this := GetSignInExp200ResponseSignIn{}
	return &this
}

// GetMethods returns the Methods field value
func (o *GetSignInExp200ResponseSignIn) GetMethods() []GetSignInExp200ResponseSignInMethodsInner {
	if o == nil {
		var ret []GetSignInExp200ResponseSignInMethodsInner
		return ret
	}

	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200ResponseSignIn) GetMethodsOk() ([]GetSignInExp200ResponseSignInMethodsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Methods, true
}

// SetMethods sets field value
func (o *GetSignInExp200ResponseSignIn) SetMethods(v []GetSignInExp200ResponseSignInMethodsInner) {
	o.Methods = v
}

func (o GetSignInExp200ResponseSignIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSignInExp200ResponseSignIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["methods"] = o.Methods
	return toSerialize, nil
}

func (o *GetSignInExp200ResponseSignIn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"methods",
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

	varGetSignInExp200ResponseSignIn := _GetSignInExp200ResponseSignIn{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSignInExp200ResponseSignIn)

	if err != nil {
		return err
	}

	*o = GetSignInExp200ResponseSignIn(varGetSignInExp200ResponseSignIn)

	return err
}

type NullableGetSignInExp200ResponseSignIn struct {
	value *GetSignInExp200ResponseSignIn
	isSet bool
}

func (v NullableGetSignInExp200ResponseSignIn) Get() *GetSignInExp200ResponseSignIn {
	return v.value
}

func (v *NullableGetSignInExp200ResponseSignIn) Set(val *GetSignInExp200ResponseSignIn) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignInExp200ResponseSignIn) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignInExp200ResponseSignIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignInExp200ResponseSignIn(val *GetSignInExp200ResponseSignIn) *NullableGetSignInExp200ResponseSignIn {
	return &NullableGetSignInExp200ResponseSignIn{value: val, isSet: true}
}

func (v NullableGetSignInExp200ResponseSignIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignInExp200ResponseSignIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


