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

// checks if the UpdateUserIsSuspendedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserIsSuspendedRequest{}

// UpdateUserIsSuspendedRequest struct for UpdateUserIsSuspendedRequest
type UpdateUserIsSuspendedRequest struct {
	// New suspension status for the user.
	IsSuspended bool `json:"isSuspended"`
}

type _UpdateUserIsSuspendedRequest UpdateUserIsSuspendedRequest

// NewUpdateUserIsSuspendedRequest instantiates a new UpdateUserIsSuspendedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserIsSuspendedRequest(isSuspended bool) *UpdateUserIsSuspendedRequest {
	this := UpdateUserIsSuspendedRequest{}
	this.IsSuspended = isSuspended
	return &this
}

// NewUpdateUserIsSuspendedRequestWithDefaults instantiates a new UpdateUserIsSuspendedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserIsSuspendedRequestWithDefaults() *UpdateUserIsSuspendedRequest {
	this := UpdateUserIsSuspendedRequest{}
	return &this
}

// GetIsSuspended returns the IsSuspended field value
func (o *UpdateUserIsSuspendedRequest) GetIsSuspended() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSuspended
}

// GetIsSuspendedOk returns a tuple with the IsSuspended field value
// and a boolean to check if the value has been set.
func (o *UpdateUserIsSuspendedRequest) GetIsSuspendedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSuspended, true
}

// SetIsSuspended sets field value
func (o *UpdateUserIsSuspendedRequest) SetIsSuspended(v bool) {
	o.IsSuspended = v
}

func (o UpdateUserIsSuspendedRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserIsSuspendedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isSuspended"] = o.IsSuspended
	return toSerialize, nil
}

func (o *UpdateUserIsSuspendedRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isSuspended",
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

	varUpdateUserIsSuspendedRequest := _UpdateUserIsSuspendedRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateUserIsSuspendedRequest)

	if err != nil {
		return err
	}

	*o = UpdateUserIsSuspendedRequest(varUpdateUserIsSuspendedRequest)

	return err
}

type NullableUpdateUserIsSuspendedRequest struct {
	value *UpdateUserIsSuspendedRequest
	isSet bool
}

func (v NullableUpdateUserIsSuspendedRequest) Get() *UpdateUserIsSuspendedRequest {
	return v.value
}

func (v *NullableUpdateUserIsSuspendedRequest) Set(val *UpdateUserIsSuspendedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserIsSuspendedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserIsSuspendedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserIsSuspendedRequest(val *UpdateUserIsSuspendedRequest) *NullableUpdateUserIsSuspendedRequest {
	return &NullableUpdateUserIsSuspendedRequest{value: val, isSet: true}
}

func (v NullableUpdateUserIsSuspendedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserIsSuspendedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


