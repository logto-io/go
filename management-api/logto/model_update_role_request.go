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

// checks if the UpdateRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoleRequest{}

// UpdateRoleRequest struct for UpdateRoleRequest
type UpdateRoleRequest struct {
	// The name of the role. It should be unique within the tenant.
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
}

// NewUpdateRoleRequest instantiates a new UpdateRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleRequest() *UpdateRoleRequest {
	this := UpdateRoleRequest{}
	return &this
}

// NewUpdateRoleRequestWithDefaults instantiates a new UpdateRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleRequestWithDefaults() *UpdateRoleRequest {
	this := UpdateRoleRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateRoleRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateRoleRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateRoleRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateRoleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateRoleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateRoleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *UpdateRoleRequest) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleRequest) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *UpdateRoleRequest) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *UpdateRoleRequest) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o UpdateRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	return toSerialize, nil
}

type NullableUpdateRoleRequest struct {
	value *UpdateRoleRequest
	isSet bool
}

func (v NullableUpdateRoleRequest) Get() *UpdateRoleRequest {
	return v.value
}

func (v *NullableUpdateRoleRequest) Set(val *UpdateRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoleRequest(val *UpdateRoleRequest) *NullableUpdateRoleRequest {
	return &NullableUpdateRoleRequest{value: val, isSet: true}
}

func (v NullableUpdateRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


