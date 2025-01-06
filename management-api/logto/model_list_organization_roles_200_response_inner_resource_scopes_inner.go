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

// checks if the ListOrganizationRoles200ResponseInnerResourceScopesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationRoles200ResponseInnerResourceScopesInner{}

// ListOrganizationRoles200ResponseInnerResourceScopesInner struct for ListOrganizationRoles200ResponseInnerResourceScopesInner
type ListOrganizationRoles200ResponseInnerResourceScopesInner struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Resource ListApplicationOrganizations200ResponseInnerOrganizationRolesInner `json:"resource"`
}

type _ListOrganizationRoles200ResponseInnerResourceScopesInner ListOrganizationRoles200ResponseInnerResourceScopesInner

// NewListOrganizationRoles200ResponseInnerResourceScopesInner instantiates a new ListOrganizationRoles200ResponseInnerResourceScopesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationRoles200ResponseInnerResourceScopesInner(id string, name string, resource ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) *ListOrganizationRoles200ResponseInnerResourceScopesInner {
	this := ListOrganizationRoles200ResponseInnerResourceScopesInner{}
	this.Id = id
	this.Name = name
	this.Resource = resource
	return &this
}

// NewListOrganizationRoles200ResponseInnerResourceScopesInnerWithDefaults instantiates a new ListOrganizationRoles200ResponseInnerResourceScopesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationRoles200ResponseInnerResourceScopesInnerWithDefaults() *ListOrganizationRoles200ResponseInnerResourceScopesInner {
	this := ListOrganizationRoles200ResponseInnerResourceScopesInner{}
	return &this
}

// GetId returns the Id field value
func (o *ListOrganizationRoles200ResponseInnerResourceScopesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationRoles200ResponseInnerResourceScopesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListOrganizationRoles200ResponseInnerResourceScopesInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ListOrganizationRoles200ResponseInnerResourceScopesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationRoles200ResponseInnerResourceScopesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListOrganizationRoles200ResponseInnerResourceScopesInner) SetName(v string) {
	o.Name = v
}

// GetResource returns the Resource field value
func (o *ListOrganizationRoles200ResponseInnerResourceScopesInner) GetResource() ListApplicationOrganizations200ResponseInnerOrganizationRolesInner {
	if o == nil {
		var ret ListApplicationOrganizations200ResponseInnerOrganizationRolesInner
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationRoles200ResponseInnerResourceScopesInner) GetResourceOk() (*ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ListOrganizationRoles200ResponseInnerResourceScopesInner) SetResource(v ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) {
	o.Resource = v
}

func (o ListOrganizationRoles200ResponseInnerResourceScopesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizationRoles200ResponseInnerResourceScopesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["resource"] = o.Resource
	return toSerialize, nil
}

func (o *ListOrganizationRoles200ResponseInnerResourceScopesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"resource",
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

	varListOrganizationRoles200ResponseInnerResourceScopesInner := _ListOrganizationRoles200ResponseInnerResourceScopesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListOrganizationRoles200ResponseInnerResourceScopesInner)

	if err != nil {
		return err
	}

	*o = ListOrganizationRoles200ResponseInnerResourceScopesInner(varListOrganizationRoles200ResponseInnerResourceScopesInner)

	return err
}

type NullableListOrganizationRoles200ResponseInnerResourceScopesInner struct {
	value *ListOrganizationRoles200ResponseInnerResourceScopesInner
	isSet bool
}

func (v NullableListOrganizationRoles200ResponseInnerResourceScopesInner) Get() *ListOrganizationRoles200ResponseInnerResourceScopesInner {
	return v.value
}

func (v *NullableListOrganizationRoles200ResponseInnerResourceScopesInner) Set(val *ListOrganizationRoles200ResponseInnerResourceScopesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationRoles200ResponseInnerResourceScopesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationRoles200ResponseInnerResourceScopesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationRoles200ResponseInnerResourceScopesInner(val *ListOrganizationRoles200ResponseInnerResourceScopesInner) *NullableListOrganizationRoles200ResponseInnerResourceScopesInner {
	return &NullableListOrganizationRoles200ResponseInnerResourceScopesInner{value: val, isSet: true}
}

func (v NullableListOrganizationRoles200ResponseInnerResourceScopesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationRoles200ResponseInnerResourceScopesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


