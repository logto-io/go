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

// checks if the ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner{}

// ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner struct for ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner
type ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner struct {
	Resource ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInnerResource `json:"resource"`
	Scopes []ListApplicationUserConsentScopes200ResponseResourceScopesInnerScopesInner `json:"scopes"`
}

type _ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner

// NewApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner instantiates a new ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner(resource ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInnerResource, scopes []ListApplicationUserConsentScopes200ResponseResourceScopesInnerScopesInner) *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner {
	this := ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner{}
	this.Resource = resource
	this.Scopes = scopes
	return &this
}

// NewApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInnerWithDefaults instantiates a new ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInnerWithDefaults() *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner {
	this := ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner{}
	return &this
}

// GetResource returns the Resource field value
func (o *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) GetResource() ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInnerResource {
	if o == nil {
		var ret ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInnerResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) GetResourceOk() (*ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInnerResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) SetResource(v ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInnerResource) {
	o.Resource = v
}

// GetScopes returns the Scopes field value
func (o *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) GetScopes() []ListApplicationUserConsentScopes200ResponseResourceScopesInnerScopesInner {
	if o == nil {
		var ret []ListApplicationUserConsentScopes200ResponseResourceScopesInnerScopesInner
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) GetScopesOk() ([]ListApplicationUserConsentScopes200ResponseResourceScopesInnerScopesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) SetScopes(v []ListApplicationUserConsentScopes200ResponseResourceScopesInnerScopesInner) {
	o.Scopes = v
}

func (o ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource"] = o.Resource
	toSerialize["scopes"] = o.Scopes
	return toSerialize, nil
}

func (o *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource",
		"scopes",
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

	varApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner := _ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner)

	if err != nil {
		return err
	}

	*o = ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner(varApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner)

	return err
}

type NullableApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner struct {
	value *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner
	isSet bool
}

func (v NullableApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) Get() *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner {
	return v.value
}

func (v *NullableApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) Set(val *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner(val *ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) *NullableApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner {
	return &NullableApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner{value: val, isSet: true}
}

func (v NullableApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


