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

// checks if the ApiInteractionConsentGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionConsentGet200Response{}

// ApiInteractionConsentGet200Response struct for ApiInteractionConsentGet200Response
type ApiInteractionConsentGet200Response struct {
	Application ApiInteractionConsentGet200ResponseApplication `json:"application"`
	User ApiInteractionConsentGet200ResponseUser `json:"user"`
	Organizations []ApiInteractionConsentGet200ResponseOrganizationsInner `json:"organizations,omitempty"`
	MissingOIDCScope []string `json:"missingOIDCScope,omitempty"`
	MissingResourceScopes []ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner `json:"missingResourceScopes,omitempty"`
	RedirectUri string `json:"redirectUri"`
}

type _ApiInteractionConsentGet200Response ApiInteractionConsentGet200Response

// NewApiInteractionConsentGet200Response instantiates a new ApiInteractionConsentGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionConsentGet200Response(application ApiInteractionConsentGet200ResponseApplication, user ApiInteractionConsentGet200ResponseUser, redirectUri string) *ApiInteractionConsentGet200Response {
	this := ApiInteractionConsentGet200Response{}
	this.Application = application
	this.User = user
	this.RedirectUri = redirectUri
	return &this
}

// NewApiInteractionConsentGet200ResponseWithDefaults instantiates a new ApiInteractionConsentGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionConsentGet200ResponseWithDefaults() *ApiInteractionConsentGet200Response {
	this := ApiInteractionConsentGet200Response{}
	return &this
}

// GetApplication returns the Application field value
func (o *ApiInteractionConsentGet200Response) GetApplication() ApiInteractionConsentGet200ResponseApplication {
	if o == nil {
		var ret ApiInteractionConsentGet200ResponseApplication
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionConsentGet200Response) GetApplicationOk() (*ApiInteractionConsentGet200ResponseApplication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *ApiInteractionConsentGet200Response) SetApplication(v ApiInteractionConsentGet200ResponseApplication) {
	o.Application = v
}

// GetUser returns the User field value
func (o *ApiInteractionConsentGet200Response) GetUser() ApiInteractionConsentGet200ResponseUser {
	if o == nil {
		var ret ApiInteractionConsentGet200ResponseUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionConsentGet200Response) GetUserOk() (*ApiInteractionConsentGet200ResponseUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ApiInteractionConsentGet200Response) SetUser(v ApiInteractionConsentGet200ResponseUser) {
	o.User = v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *ApiInteractionConsentGet200Response) GetOrganizations() []ApiInteractionConsentGet200ResponseOrganizationsInner {
	if o == nil || IsNil(o.Organizations) {
		var ret []ApiInteractionConsentGet200ResponseOrganizationsInner
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInteractionConsentGet200Response) GetOrganizationsOk() ([]ApiInteractionConsentGet200ResponseOrganizationsInner, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *ApiInteractionConsentGet200Response) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []ApiInteractionConsentGet200ResponseOrganizationsInner and assigns it to the Organizations field.
func (o *ApiInteractionConsentGet200Response) SetOrganizations(v []ApiInteractionConsentGet200ResponseOrganizationsInner) {
	o.Organizations = v
}

// GetMissingOIDCScope returns the MissingOIDCScope field value if set, zero value otherwise.
func (o *ApiInteractionConsentGet200Response) GetMissingOIDCScope() []string {
	if o == nil || IsNil(o.MissingOIDCScope) {
		var ret []string
		return ret
	}
	return o.MissingOIDCScope
}

// GetMissingOIDCScopeOk returns a tuple with the MissingOIDCScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInteractionConsentGet200Response) GetMissingOIDCScopeOk() ([]string, bool) {
	if o == nil || IsNil(o.MissingOIDCScope) {
		return nil, false
	}
	return o.MissingOIDCScope, true
}

// HasMissingOIDCScope returns a boolean if a field has been set.
func (o *ApiInteractionConsentGet200Response) HasMissingOIDCScope() bool {
	if o != nil && !IsNil(o.MissingOIDCScope) {
		return true
	}

	return false
}

// SetMissingOIDCScope gets a reference to the given []string and assigns it to the MissingOIDCScope field.
func (o *ApiInteractionConsentGet200Response) SetMissingOIDCScope(v []string) {
	o.MissingOIDCScope = v
}

// GetMissingResourceScopes returns the MissingResourceScopes field value if set, zero value otherwise.
func (o *ApiInteractionConsentGet200Response) GetMissingResourceScopes() []ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner {
	if o == nil || IsNil(o.MissingResourceScopes) {
		var ret []ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner
		return ret
	}
	return o.MissingResourceScopes
}

// GetMissingResourceScopesOk returns a tuple with the MissingResourceScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInteractionConsentGet200Response) GetMissingResourceScopesOk() ([]ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner, bool) {
	if o == nil || IsNil(o.MissingResourceScopes) {
		return nil, false
	}
	return o.MissingResourceScopes, true
}

// HasMissingResourceScopes returns a boolean if a field has been set.
func (o *ApiInteractionConsentGet200Response) HasMissingResourceScopes() bool {
	if o != nil && !IsNil(o.MissingResourceScopes) {
		return true
	}

	return false
}

// SetMissingResourceScopes gets a reference to the given []ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner and assigns it to the MissingResourceScopes field.
func (o *ApiInteractionConsentGet200Response) SetMissingResourceScopes(v []ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner) {
	o.MissingResourceScopes = v
}

// GetRedirectUri returns the RedirectUri field value
func (o *ApiInteractionConsentGet200Response) GetRedirectUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionConsentGet200Response) GetRedirectUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUri, true
}

// SetRedirectUri sets field value
func (o *ApiInteractionConsentGet200Response) SetRedirectUri(v string) {
	o.RedirectUri = v
}

func (o ApiInteractionConsentGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionConsentGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["application"] = o.Application
	toSerialize["user"] = o.User
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	if !IsNil(o.MissingOIDCScope) {
		toSerialize["missingOIDCScope"] = o.MissingOIDCScope
	}
	if !IsNil(o.MissingResourceScopes) {
		toSerialize["missingResourceScopes"] = o.MissingResourceScopes
	}
	toSerialize["redirectUri"] = o.RedirectUri
	return toSerialize, nil
}

func (o *ApiInteractionConsentGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"application",
		"user",
		"redirectUri",
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

	varApiInteractionConsentGet200Response := _ApiInteractionConsentGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiInteractionConsentGet200Response)

	if err != nil {
		return err
	}

	*o = ApiInteractionConsentGet200Response(varApiInteractionConsentGet200Response)

	return err
}

type NullableApiInteractionConsentGet200Response struct {
	value *ApiInteractionConsentGet200Response
	isSet bool
}

func (v NullableApiInteractionConsentGet200Response) Get() *ApiInteractionConsentGet200Response {
	return v.value
}

func (v *NullableApiInteractionConsentGet200Response) Set(val *ApiInteractionConsentGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionConsentGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionConsentGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionConsentGet200Response(val *ApiInteractionConsentGet200Response) *NullableApiInteractionConsentGet200Response {
	return &NullableApiInteractionConsentGet200Response{value: val, isSet: true}
}

func (v NullableApiInteractionConsentGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionConsentGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


