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

// checks if the ListRoles200ResponseInnerFeaturedApplicationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRoles200ResponseInnerFeaturedApplicationsInner{}

// ListRoles200ResponseInnerFeaturedApplicationsInner struct for ListRoles200ResponseInnerFeaturedApplicationsInner
type ListRoles200ResponseInnerFeaturedApplicationsInner struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type _ListRoles200ResponseInnerFeaturedApplicationsInner ListRoles200ResponseInnerFeaturedApplicationsInner

// NewListRoles200ResponseInnerFeaturedApplicationsInner instantiates a new ListRoles200ResponseInnerFeaturedApplicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRoles200ResponseInnerFeaturedApplicationsInner(id string, name string, type_ string) *ListRoles200ResponseInnerFeaturedApplicationsInner {
	this := ListRoles200ResponseInnerFeaturedApplicationsInner{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewListRoles200ResponseInnerFeaturedApplicationsInnerWithDefaults instantiates a new ListRoles200ResponseInnerFeaturedApplicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRoles200ResponseInnerFeaturedApplicationsInnerWithDefaults() *ListRoles200ResponseInnerFeaturedApplicationsInner {
	this := ListRoles200ResponseInnerFeaturedApplicationsInner{}
	return &this
}

// GetId returns the Id field value
func (o *ListRoles200ResponseInnerFeaturedApplicationsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListRoles200ResponseInnerFeaturedApplicationsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListRoles200ResponseInnerFeaturedApplicationsInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ListRoles200ResponseInnerFeaturedApplicationsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListRoles200ResponseInnerFeaturedApplicationsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListRoles200ResponseInnerFeaturedApplicationsInner) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ListRoles200ResponseInnerFeaturedApplicationsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListRoles200ResponseInnerFeaturedApplicationsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListRoles200ResponseInnerFeaturedApplicationsInner) SetType(v string) {
	o.Type = v
}

func (o ListRoles200ResponseInnerFeaturedApplicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRoles200ResponseInnerFeaturedApplicationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ListRoles200ResponseInnerFeaturedApplicationsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
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

	varListRoles200ResponseInnerFeaturedApplicationsInner := _ListRoles200ResponseInnerFeaturedApplicationsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListRoles200ResponseInnerFeaturedApplicationsInner)

	if err != nil {
		return err
	}

	*o = ListRoles200ResponseInnerFeaturedApplicationsInner(varListRoles200ResponseInnerFeaturedApplicationsInner)

	return err
}

type NullableListRoles200ResponseInnerFeaturedApplicationsInner struct {
	value *ListRoles200ResponseInnerFeaturedApplicationsInner
	isSet bool
}

func (v NullableListRoles200ResponseInnerFeaturedApplicationsInner) Get() *ListRoles200ResponseInnerFeaturedApplicationsInner {
	return v.value
}

func (v *NullableListRoles200ResponseInnerFeaturedApplicationsInner) Set(val *ListRoles200ResponseInnerFeaturedApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListRoles200ResponseInnerFeaturedApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListRoles200ResponseInnerFeaturedApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRoles200ResponseInnerFeaturedApplicationsInner(val *ListRoles200ResponseInnerFeaturedApplicationsInner) *NullableListRoles200ResponseInnerFeaturedApplicationsInner {
	return &NullableListRoles200ResponseInnerFeaturedApplicationsInner{value: val, isSet: true}
}

func (v NullableListRoles200ResponseInnerFeaturedApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRoles200ResponseInnerFeaturedApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


