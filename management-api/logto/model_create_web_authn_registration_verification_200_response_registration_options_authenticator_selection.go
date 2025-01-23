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

// checks if the CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection{}

// CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection struct for CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection
type CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection struct {
	AuthenticatorAttachment *string `json:"authenticatorAttachment,omitempty"`
	RequireResidentKey *bool `json:"requireResidentKey,omitempty"`
	ResidentKey *string `json:"residentKey,omitempty"`
	UserVerification *string `json:"userVerification,omitempty"`
}

// NewCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection instantiates a new CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection() *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection {
	this := CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection{}
	return &this
}

// NewCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelectionWithDefaults instantiates a new CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelectionWithDefaults() *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection {
	this := CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection{}
	return &this
}

// GetAuthenticatorAttachment returns the AuthenticatorAttachment field value if set, zero value otherwise.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) GetAuthenticatorAttachment() string {
	if o == nil || IsNil(o.AuthenticatorAttachment) {
		var ret string
		return ret
	}
	return *o.AuthenticatorAttachment
}

// GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) GetAuthenticatorAttachmentOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticatorAttachment) {
		return nil, false
	}
	return o.AuthenticatorAttachment, true
}

// HasAuthenticatorAttachment returns a boolean if a field has been set.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) HasAuthenticatorAttachment() bool {
	if o != nil && !IsNil(o.AuthenticatorAttachment) {
		return true
	}

	return false
}

// SetAuthenticatorAttachment gets a reference to the given string and assigns it to the AuthenticatorAttachment field.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) SetAuthenticatorAttachment(v string) {
	o.AuthenticatorAttachment = &v
}

// GetRequireResidentKey returns the RequireResidentKey field value if set, zero value otherwise.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) GetRequireResidentKey() bool {
	if o == nil || IsNil(o.RequireResidentKey) {
		var ret bool
		return ret
	}
	return *o.RequireResidentKey
}

// GetRequireResidentKeyOk returns a tuple with the RequireResidentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) GetRequireResidentKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireResidentKey) {
		return nil, false
	}
	return o.RequireResidentKey, true
}

// HasRequireResidentKey returns a boolean if a field has been set.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) HasRequireResidentKey() bool {
	if o != nil && !IsNil(o.RequireResidentKey) {
		return true
	}

	return false
}

// SetRequireResidentKey gets a reference to the given bool and assigns it to the RequireResidentKey field.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) SetRequireResidentKey(v bool) {
	o.RequireResidentKey = &v
}

// GetResidentKey returns the ResidentKey field value if set, zero value otherwise.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) GetResidentKey() string {
	if o == nil || IsNil(o.ResidentKey) {
		var ret string
		return ret
	}
	return *o.ResidentKey
}

// GetResidentKeyOk returns a tuple with the ResidentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) GetResidentKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ResidentKey) {
		return nil, false
	}
	return o.ResidentKey, true
}

// HasResidentKey returns a boolean if a field has been set.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) HasResidentKey() bool {
	if o != nil && !IsNil(o.ResidentKey) {
		return true
	}

	return false
}

// SetResidentKey gets a reference to the given string and assigns it to the ResidentKey field.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) SetResidentKey(v string) {
	o.ResidentKey = &v
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) GetUserVerification() string {
	if o == nil || IsNil(o.UserVerification) {
		var ret string
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) GetUserVerificationOk() (*string, bool) {
	if o == nil || IsNil(o.UserVerification) {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) HasUserVerification() bool {
	if o != nil && !IsNil(o.UserVerification) {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given string and assigns it to the UserVerification field.
func (o *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) SetUserVerification(v string) {
	o.UserVerification = &v
}

func (o CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticatorAttachment) {
		toSerialize["authenticatorAttachment"] = o.AuthenticatorAttachment
	}
	if !IsNil(o.RequireResidentKey) {
		toSerialize["requireResidentKey"] = o.RequireResidentKey
	}
	if !IsNil(o.ResidentKey) {
		toSerialize["residentKey"] = o.ResidentKey
	}
	if !IsNil(o.UserVerification) {
		toSerialize["userVerification"] = o.UserVerification
	}
	return toSerialize, nil
}

type NullableCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection struct {
	value *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection
	isSet bool
}

func (v NullableCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) Get() *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection {
	return v.value
}

func (v *NullableCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) Set(val *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection(val *CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) *NullableCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection {
	return &NullableCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection{value: val, isSet: true}
}

func (v NullableCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


