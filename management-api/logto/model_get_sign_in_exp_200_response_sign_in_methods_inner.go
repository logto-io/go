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

// checks if the GetSignInExp200ResponseSignInMethodsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSignInExp200ResponseSignInMethodsInner{}

// GetSignInExp200ResponseSignInMethodsInner struct for GetSignInExp200ResponseSignInMethodsInner
type GetSignInExp200ResponseSignInMethodsInner struct {
	Identifier string `json:"identifier"`
	Password bool `json:"password"`
	VerificationCode bool `json:"verificationCode"`
	IsPasswordPrimary bool `json:"isPasswordPrimary"`
}

type _GetSignInExp200ResponseSignInMethodsInner GetSignInExp200ResponseSignInMethodsInner

// NewGetSignInExp200ResponseSignInMethodsInner instantiates a new GetSignInExp200ResponseSignInMethodsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignInExp200ResponseSignInMethodsInner(identifier string, password bool, verificationCode bool, isPasswordPrimary bool) *GetSignInExp200ResponseSignInMethodsInner {
	this := GetSignInExp200ResponseSignInMethodsInner{}
	this.Identifier = identifier
	this.Password = password
	this.VerificationCode = verificationCode
	this.IsPasswordPrimary = isPasswordPrimary
	return &this
}

// NewGetSignInExp200ResponseSignInMethodsInnerWithDefaults instantiates a new GetSignInExp200ResponseSignInMethodsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignInExp200ResponseSignInMethodsInnerWithDefaults() *GetSignInExp200ResponseSignInMethodsInner {
	this := GetSignInExp200ResponseSignInMethodsInner{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *GetSignInExp200ResponseSignInMethodsInner) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200ResponseSignInMethodsInner) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *GetSignInExp200ResponseSignInMethodsInner) SetIdentifier(v string) {
	o.Identifier = v
}

// GetPassword returns the Password field value
func (o *GetSignInExp200ResponseSignInMethodsInner) GetPassword() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200ResponseSignInMethodsInner) GetPasswordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *GetSignInExp200ResponseSignInMethodsInner) SetPassword(v bool) {
	o.Password = v
}

// GetVerificationCode returns the VerificationCode field value
func (o *GetSignInExp200ResponseSignInMethodsInner) GetVerificationCode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VerificationCode
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200ResponseSignInMethodsInner) GetVerificationCodeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationCode, true
}

// SetVerificationCode sets field value
func (o *GetSignInExp200ResponseSignInMethodsInner) SetVerificationCode(v bool) {
	o.VerificationCode = v
}

// GetIsPasswordPrimary returns the IsPasswordPrimary field value
func (o *GetSignInExp200ResponseSignInMethodsInner) GetIsPasswordPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPasswordPrimary
}

// GetIsPasswordPrimaryOk returns a tuple with the IsPasswordPrimary field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200ResponseSignInMethodsInner) GetIsPasswordPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPasswordPrimary, true
}

// SetIsPasswordPrimary sets field value
func (o *GetSignInExp200ResponseSignInMethodsInner) SetIsPasswordPrimary(v bool) {
	o.IsPasswordPrimary = v
}

func (o GetSignInExp200ResponseSignInMethodsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSignInExp200ResponseSignInMethodsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifier"] = o.Identifier
	toSerialize["password"] = o.Password
	toSerialize["verificationCode"] = o.VerificationCode
	toSerialize["isPasswordPrimary"] = o.IsPasswordPrimary
	return toSerialize, nil
}

func (o *GetSignInExp200ResponseSignInMethodsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identifier",
		"password",
		"verificationCode",
		"isPasswordPrimary",
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

	varGetSignInExp200ResponseSignInMethodsInner := _GetSignInExp200ResponseSignInMethodsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSignInExp200ResponseSignInMethodsInner)

	if err != nil {
		return err
	}

	*o = GetSignInExp200ResponseSignInMethodsInner(varGetSignInExp200ResponseSignInMethodsInner)

	return err
}

type NullableGetSignInExp200ResponseSignInMethodsInner struct {
	value *GetSignInExp200ResponseSignInMethodsInner
	isSet bool
}

func (v NullableGetSignInExp200ResponseSignInMethodsInner) Get() *GetSignInExp200ResponseSignInMethodsInner {
	return v.value
}

func (v *NullableGetSignInExp200ResponseSignInMethodsInner) Set(val *GetSignInExp200ResponseSignInMethodsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignInExp200ResponseSignInMethodsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignInExp200ResponseSignInMethodsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignInExp200ResponseSignInMethodsInner(val *GetSignInExp200ResponseSignInMethodsInner) *NullableGetSignInExp200ResponseSignInMethodsInner {
	return &NullableGetSignInExp200ResponseSignInMethodsInner{value: val, isSet: true}
}

func (v NullableGetSignInExp200ResponseSignInMethodsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignInExp200ResponseSignInMethodsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


