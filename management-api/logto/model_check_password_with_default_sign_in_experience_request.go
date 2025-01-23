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

// checks if the CheckPasswordWithDefaultSignInExperienceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckPasswordWithDefaultSignInExperienceRequest{}

// CheckPasswordWithDefaultSignInExperienceRequest struct for CheckPasswordWithDefaultSignInExperienceRequest
type CheckPasswordWithDefaultSignInExperienceRequest struct {
	// The password to check.
	Password string `json:"password"`
	// The user ID to check the password for. It is required if rejects user info is enabled in the password policy.
	UserId *string `json:"userId,omitempty"`
}

type _CheckPasswordWithDefaultSignInExperienceRequest CheckPasswordWithDefaultSignInExperienceRequest

// NewCheckPasswordWithDefaultSignInExperienceRequest instantiates a new CheckPasswordWithDefaultSignInExperienceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckPasswordWithDefaultSignInExperienceRequest(password string) *CheckPasswordWithDefaultSignInExperienceRequest {
	this := CheckPasswordWithDefaultSignInExperienceRequest{}
	this.Password = password
	return &this
}

// NewCheckPasswordWithDefaultSignInExperienceRequestWithDefaults instantiates a new CheckPasswordWithDefaultSignInExperienceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckPasswordWithDefaultSignInExperienceRequestWithDefaults() *CheckPasswordWithDefaultSignInExperienceRequest {
	this := CheckPasswordWithDefaultSignInExperienceRequest{}
	return &this
}

// GetPassword returns the Password field value
func (o *CheckPasswordWithDefaultSignInExperienceRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CheckPasswordWithDefaultSignInExperienceRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CheckPasswordWithDefaultSignInExperienceRequest) SetPassword(v string) {
	o.Password = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CheckPasswordWithDefaultSignInExperienceRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckPasswordWithDefaultSignInExperienceRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CheckPasswordWithDefaultSignInExperienceRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *CheckPasswordWithDefaultSignInExperienceRequest) SetUserId(v string) {
	o.UserId = &v
}

func (o CheckPasswordWithDefaultSignInExperienceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckPasswordWithDefaultSignInExperienceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

func (o *CheckPasswordWithDefaultSignInExperienceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
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

	varCheckPasswordWithDefaultSignInExperienceRequest := _CheckPasswordWithDefaultSignInExperienceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCheckPasswordWithDefaultSignInExperienceRequest)

	if err != nil {
		return err
	}

	*o = CheckPasswordWithDefaultSignInExperienceRequest(varCheckPasswordWithDefaultSignInExperienceRequest)

	return err
}

type NullableCheckPasswordWithDefaultSignInExperienceRequest struct {
	value *CheckPasswordWithDefaultSignInExperienceRequest
	isSet bool
}

func (v NullableCheckPasswordWithDefaultSignInExperienceRequest) Get() *CheckPasswordWithDefaultSignInExperienceRequest {
	return v.value
}

func (v *NullableCheckPasswordWithDefaultSignInExperienceRequest) Set(val *CheckPasswordWithDefaultSignInExperienceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckPasswordWithDefaultSignInExperienceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckPasswordWithDefaultSignInExperienceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckPasswordWithDefaultSignInExperienceRequest(val *CheckPasswordWithDefaultSignInExperienceRequest) *NullableCheckPasswordWithDefaultSignInExperienceRequest {
	return &NullableCheckPasswordWithDefaultSignInExperienceRequest{value: val, isSet: true}
}

func (v NullableCheckPasswordWithDefaultSignInExperienceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckPasswordWithDefaultSignInExperienceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


