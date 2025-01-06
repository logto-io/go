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

// checks if the CreateRoleScopeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoleScopeRequest{}

// CreateRoleScopeRequest struct for CreateRoleScopeRequest
type CreateRoleScopeRequest struct {
	// An array of API resource scope IDs to be linked.
	ScopeIds []string `json:"scopeIds"`
}

type _CreateRoleScopeRequest CreateRoleScopeRequest

// NewCreateRoleScopeRequest instantiates a new CreateRoleScopeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoleScopeRequest(scopeIds []string) *CreateRoleScopeRequest {
	this := CreateRoleScopeRequest{}
	this.ScopeIds = scopeIds
	return &this
}

// NewCreateRoleScopeRequestWithDefaults instantiates a new CreateRoleScopeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleScopeRequestWithDefaults() *CreateRoleScopeRequest {
	this := CreateRoleScopeRequest{}
	return &this
}

// GetScopeIds returns the ScopeIds field value
func (o *CreateRoleScopeRequest) GetScopeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ScopeIds
}

// GetScopeIdsOk returns a tuple with the ScopeIds field value
// and a boolean to check if the value has been set.
func (o *CreateRoleScopeRequest) GetScopeIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeIds, true
}

// SetScopeIds sets field value
func (o *CreateRoleScopeRequest) SetScopeIds(v []string) {
	o.ScopeIds = v
}

func (o CreateRoleScopeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoleScopeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scopeIds"] = o.ScopeIds
	return toSerialize, nil
}

func (o *CreateRoleScopeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"scopeIds",
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

	varCreateRoleScopeRequest := _CreateRoleScopeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRoleScopeRequest)

	if err != nil {
		return err
	}

	*o = CreateRoleScopeRequest(varCreateRoleScopeRequest)

	return err
}

type NullableCreateRoleScopeRequest struct {
	value *CreateRoleScopeRequest
	isSet bool
}

func (v NullableCreateRoleScopeRequest) Get() *CreateRoleScopeRequest {
	return v.value
}

func (v *NullableCreateRoleScopeRequest) Set(val *CreateRoleScopeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoleScopeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoleScopeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoleScopeRequest(val *CreateRoleScopeRequest) *NullableCreateRoleScopeRequest {
	return &NullableCreateRoleScopeRequest{value: val, isSet: true}
}

func (v NullableCreateRoleScopeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoleScopeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


