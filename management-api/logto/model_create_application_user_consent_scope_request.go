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

// checks if the CreateApplicationUserConsentScopeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApplicationUserConsentScopeRequest{}

// CreateApplicationUserConsentScopeRequest struct for CreateApplicationUserConsentScopeRequest
type CreateApplicationUserConsentScopeRequest struct {
	// A list of organization scope id to assign to the application. Throws error if any given organization scope is not found.
	OrganizationScopes []string `json:"organizationScopes,omitempty"`
	// A list of resource scope id to assign to the application. Throws error if any given resource scope is not found.
	ResourceScopes []string `json:"resourceScopes,omitempty"`
	// A list of organization resource scope id to assign to the application. Throws error if any given resource scope is not found.
	OrganizationResourceScopes []string `json:"organizationResourceScopes,omitempty"`
	// A list of user scope enum value to assign to the application.
	UserScopes []string `json:"userScopes,omitempty"`
}

// NewCreateApplicationUserConsentScopeRequest instantiates a new CreateApplicationUserConsentScopeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApplicationUserConsentScopeRequest() *CreateApplicationUserConsentScopeRequest {
	this := CreateApplicationUserConsentScopeRequest{}
	return &this
}

// NewCreateApplicationUserConsentScopeRequestWithDefaults instantiates a new CreateApplicationUserConsentScopeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApplicationUserConsentScopeRequestWithDefaults() *CreateApplicationUserConsentScopeRequest {
	this := CreateApplicationUserConsentScopeRequest{}
	return &this
}

// GetOrganizationScopes returns the OrganizationScopes field value if set, zero value otherwise.
func (o *CreateApplicationUserConsentScopeRequest) GetOrganizationScopes() []string {
	if o == nil || IsNil(o.OrganizationScopes) {
		var ret []string
		return ret
	}
	return o.OrganizationScopes
}

// GetOrganizationScopesOk returns a tuple with the OrganizationScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationUserConsentScopeRequest) GetOrganizationScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationScopes) {
		return nil, false
	}
	return o.OrganizationScopes, true
}

// HasOrganizationScopes returns a boolean if a field has been set.
func (o *CreateApplicationUserConsentScopeRequest) HasOrganizationScopes() bool {
	if o != nil && !IsNil(o.OrganizationScopes) {
		return true
	}

	return false
}

// SetOrganizationScopes gets a reference to the given []string and assigns it to the OrganizationScopes field.
func (o *CreateApplicationUserConsentScopeRequest) SetOrganizationScopes(v []string) {
	o.OrganizationScopes = v
}

// GetResourceScopes returns the ResourceScopes field value if set, zero value otherwise.
func (o *CreateApplicationUserConsentScopeRequest) GetResourceScopes() []string {
	if o == nil || IsNil(o.ResourceScopes) {
		var ret []string
		return ret
	}
	return o.ResourceScopes
}

// GetResourceScopesOk returns a tuple with the ResourceScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationUserConsentScopeRequest) GetResourceScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.ResourceScopes) {
		return nil, false
	}
	return o.ResourceScopes, true
}

// HasResourceScopes returns a boolean if a field has been set.
func (o *CreateApplicationUserConsentScopeRequest) HasResourceScopes() bool {
	if o != nil && !IsNil(o.ResourceScopes) {
		return true
	}

	return false
}

// SetResourceScopes gets a reference to the given []string and assigns it to the ResourceScopes field.
func (o *CreateApplicationUserConsentScopeRequest) SetResourceScopes(v []string) {
	o.ResourceScopes = v
}

// GetOrganizationResourceScopes returns the OrganizationResourceScopes field value if set, zero value otherwise.
func (o *CreateApplicationUserConsentScopeRequest) GetOrganizationResourceScopes() []string {
	if o == nil || IsNil(o.OrganizationResourceScopes) {
		var ret []string
		return ret
	}
	return o.OrganizationResourceScopes
}

// GetOrganizationResourceScopesOk returns a tuple with the OrganizationResourceScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationUserConsentScopeRequest) GetOrganizationResourceScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationResourceScopes) {
		return nil, false
	}
	return o.OrganizationResourceScopes, true
}

// HasOrganizationResourceScopes returns a boolean if a field has been set.
func (o *CreateApplicationUserConsentScopeRequest) HasOrganizationResourceScopes() bool {
	if o != nil && !IsNil(o.OrganizationResourceScopes) {
		return true
	}

	return false
}

// SetOrganizationResourceScopes gets a reference to the given []string and assigns it to the OrganizationResourceScopes field.
func (o *CreateApplicationUserConsentScopeRequest) SetOrganizationResourceScopes(v []string) {
	o.OrganizationResourceScopes = v
}

// GetUserScopes returns the UserScopes field value if set, zero value otherwise.
func (o *CreateApplicationUserConsentScopeRequest) GetUserScopes() []string {
	if o == nil || IsNil(o.UserScopes) {
		var ret []string
		return ret
	}
	return o.UserScopes
}

// GetUserScopesOk returns a tuple with the UserScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationUserConsentScopeRequest) GetUserScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.UserScopes) {
		return nil, false
	}
	return o.UserScopes, true
}

// HasUserScopes returns a boolean if a field has been set.
func (o *CreateApplicationUserConsentScopeRequest) HasUserScopes() bool {
	if o != nil && !IsNil(o.UserScopes) {
		return true
	}

	return false
}

// SetUserScopes gets a reference to the given []string and assigns it to the UserScopes field.
func (o *CreateApplicationUserConsentScopeRequest) SetUserScopes(v []string) {
	o.UserScopes = v
}

func (o CreateApplicationUserConsentScopeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateApplicationUserConsentScopeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrganizationScopes) {
		toSerialize["organizationScopes"] = o.OrganizationScopes
	}
	if !IsNil(o.ResourceScopes) {
		toSerialize["resourceScopes"] = o.ResourceScopes
	}
	if !IsNil(o.OrganizationResourceScopes) {
		toSerialize["organizationResourceScopes"] = o.OrganizationResourceScopes
	}
	if !IsNil(o.UserScopes) {
		toSerialize["userScopes"] = o.UserScopes
	}
	return toSerialize, nil
}

type NullableCreateApplicationUserConsentScopeRequest struct {
	value *CreateApplicationUserConsentScopeRequest
	isSet bool
}

func (v NullableCreateApplicationUserConsentScopeRequest) Get() *CreateApplicationUserConsentScopeRequest {
	return v.value
}

func (v *NullableCreateApplicationUserConsentScopeRequest) Set(val *CreateApplicationUserConsentScopeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApplicationUserConsentScopeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApplicationUserConsentScopeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApplicationUserConsentScopeRequest(val *CreateApplicationUserConsentScopeRequest) *NullableCreateApplicationUserConsentScopeRequest {
	return &NullableCreateApplicationUserConsentScopeRequest{value: val, isSet: true}
}

func (v NullableCreateApplicationUserConsentScopeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApplicationUserConsentScopeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


