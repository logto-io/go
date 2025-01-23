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

// checks if the GetSignInExp200ResponseMfa type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSignInExp200ResponseMfa{}

// GetSignInExp200ResponseMfa MFA settings
type GetSignInExp200ResponseMfa struct {
	Factors []string `json:"factors"`
	Policy string `json:"policy"`
}

type _GetSignInExp200ResponseMfa GetSignInExp200ResponseMfa

// NewGetSignInExp200ResponseMfa instantiates a new GetSignInExp200ResponseMfa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignInExp200ResponseMfa(factors []string, policy string) *GetSignInExp200ResponseMfa {
	this := GetSignInExp200ResponseMfa{}
	this.Factors = factors
	this.Policy = policy
	return &this
}

// NewGetSignInExp200ResponseMfaWithDefaults instantiates a new GetSignInExp200ResponseMfa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignInExp200ResponseMfaWithDefaults() *GetSignInExp200ResponseMfa {
	this := GetSignInExp200ResponseMfa{}
	return &this
}

// GetFactors returns the Factors field value
func (o *GetSignInExp200ResponseMfa) GetFactors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Factors
}

// GetFactorsOk returns a tuple with the Factors field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200ResponseMfa) GetFactorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Factors, true
}

// SetFactors sets field value
func (o *GetSignInExp200ResponseMfa) SetFactors(v []string) {
	o.Factors = v
}

// GetPolicy returns the Policy field value
func (o *GetSignInExp200ResponseMfa) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200ResponseMfa) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *GetSignInExp200ResponseMfa) SetPolicy(v string) {
	o.Policy = v
}

func (o GetSignInExp200ResponseMfa) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSignInExp200ResponseMfa) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["factors"] = o.Factors
	toSerialize["policy"] = o.Policy
	return toSerialize, nil
}

func (o *GetSignInExp200ResponseMfa) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"factors",
		"policy",
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

	varGetSignInExp200ResponseMfa := _GetSignInExp200ResponseMfa{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSignInExp200ResponseMfa)

	if err != nil {
		return err
	}

	*o = GetSignInExp200ResponseMfa(varGetSignInExp200ResponseMfa)

	return err
}

type NullableGetSignInExp200ResponseMfa struct {
	value *GetSignInExp200ResponseMfa
	isSet bool
}

func (v NullableGetSignInExp200ResponseMfa) Get() *GetSignInExp200ResponseMfa {
	return v.value
}

func (v *NullableGetSignInExp200ResponseMfa) Set(val *GetSignInExp200ResponseMfa) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignInExp200ResponseMfa) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignInExp200ResponseMfa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignInExp200ResponseMfa(val *GetSignInExp200ResponseMfa) *NullableGetSignInExp200ResponseMfa {
	return &NullableGetSignInExp200ResponseMfa{value: val, isSet: true}
}

func (v NullableGetSignInExp200ResponseMfa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignInExp200ResponseMfa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


