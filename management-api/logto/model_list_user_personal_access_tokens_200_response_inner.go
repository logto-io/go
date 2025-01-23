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

// checks if the ListUserPersonalAccessTokens200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUserPersonalAccessTokens200ResponseInner{}

// ListUserPersonalAccessTokens200ResponseInner struct for ListUserPersonalAccessTokens200ResponseInner
type ListUserPersonalAccessTokens200ResponseInner struct {
	TenantId string `json:"tenantId"`
	UserId string `json:"userId"`
	Name string `json:"name"`
	Value string `json:"value"`
	CreatedAt float32 `json:"createdAt"`
	ExpiresAt NullableFloat32 `json:"expiresAt"`
}

type _ListUserPersonalAccessTokens200ResponseInner ListUserPersonalAccessTokens200ResponseInner

// NewListUserPersonalAccessTokens200ResponseInner instantiates a new ListUserPersonalAccessTokens200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUserPersonalAccessTokens200ResponseInner(tenantId string, userId string, name string, value string, createdAt float32, expiresAt NullableFloat32) *ListUserPersonalAccessTokens200ResponseInner {
	this := ListUserPersonalAccessTokens200ResponseInner{}
	this.TenantId = tenantId
	this.UserId = userId
	this.Name = name
	this.Value = value
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	return &this
}

// NewListUserPersonalAccessTokens200ResponseInnerWithDefaults instantiates a new ListUserPersonalAccessTokens200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUserPersonalAccessTokens200ResponseInnerWithDefaults() *ListUserPersonalAccessTokens200ResponseInner {
	this := ListUserPersonalAccessTokens200ResponseInner{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ListUserPersonalAccessTokens200ResponseInner) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListUserPersonalAccessTokens200ResponseInner) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListUserPersonalAccessTokens200ResponseInner) SetTenantId(v string) {
	o.TenantId = v
}

// GetUserId returns the UserId field value
func (o *ListUserPersonalAccessTokens200ResponseInner) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ListUserPersonalAccessTokens200ResponseInner) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ListUserPersonalAccessTokens200ResponseInner) SetUserId(v string) {
	o.UserId = v
}

// GetName returns the Name field value
func (o *ListUserPersonalAccessTokens200ResponseInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListUserPersonalAccessTokens200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListUserPersonalAccessTokens200ResponseInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *ListUserPersonalAccessTokens200ResponseInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListUserPersonalAccessTokens200ResponseInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListUserPersonalAccessTokens200ResponseInner) SetValue(v string) {
	o.Value = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ListUserPersonalAccessTokens200ResponseInner) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ListUserPersonalAccessTokens200ResponseInner) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ListUserPersonalAccessTokens200ResponseInner) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ListUserPersonalAccessTokens200ResponseInner) GetExpiresAt() float32 {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListUserPersonalAccessTokens200ResponseInner) GetExpiresAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *ListUserPersonalAccessTokens200ResponseInner) SetExpiresAt(v float32) {
	o.ExpiresAt.Set(&v)
}

func (o ListUserPersonalAccessTokens200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUserPersonalAccessTokens200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["userId"] = o.UserId
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["expiresAt"] = o.ExpiresAt.Get()
	return toSerialize, nil
}

func (o *ListUserPersonalAccessTokens200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"userId",
		"name",
		"value",
		"createdAt",
		"expiresAt",
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

	varListUserPersonalAccessTokens200ResponseInner := _ListUserPersonalAccessTokens200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListUserPersonalAccessTokens200ResponseInner)

	if err != nil {
		return err
	}

	*o = ListUserPersonalAccessTokens200ResponseInner(varListUserPersonalAccessTokens200ResponseInner)

	return err
}

type NullableListUserPersonalAccessTokens200ResponseInner struct {
	value *ListUserPersonalAccessTokens200ResponseInner
	isSet bool
}

func (v NullableListUserPersonalAccessTokens200ResponseInner) Get() *ListUserPersonalAccessTokens200ResponseInner {
	return v.value
}

func (v *NullableListUserPersonalAccessTokens200ResponseInner) Set(val *ListUserPersonalAccessTokens200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListUserPersonalAccessTokens200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListUserPersonalAccessTokens200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUserPersonalAccessTokens200ResponseInner(val *ListUserPersonalAccessTokens200ResponseInner) *NullableListUserPersonalAccessTokens200ResponseInner {
	return &NullableListUserPersonalAccessTokens200ResponseInner{value: val, isSet: true}
}

func (v NullableListUserPersonalAccessTokens200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUserPersonalAccessTokens200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


