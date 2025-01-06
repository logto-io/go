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

// checks if the UpdateAccountCenterSettingsRequestFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAccountCenterSettingsRequestFields{}

// UpdateAccountCenterSettingsRequestFields The fields settings for the account API.
type UpdateAccountCenterSettingsRequestFields struct {
	Name *string `json:"name,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
	Profile *string `json:"profile,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
	Social *string `json:"social,omitempty"`
	CustomData *string `json:"customData,omitempty"`
}

// NewUpdateAccountCenterSettingsRequestFields instantiates a new UpdateAccountCenterSettingsRequestFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountCenterSettingsRequestFields() *UpdateAccountCenterSettingsRequestFields {
	this := UpdateAccountCenterSettingsRequestFields{}
	return &this
}

// NewUpdateAccountCenterSettingsRequestFieldsWithDefaults instantiates a new UpdateAccountCenterSettingsRequestFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountCenterSettingsRequestFieldsWithDefaults() *UpdateAccountCenterSettingsRequestFields {
	this := UpdateAccountCenterSettingsRequestFields{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAccountCenterSettingsRequestFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountCenterSettingsRequestFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAccountCenterSettingsRequestFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAccountCenterSettingsRequestFields) SetName(v string) {
	o.Name = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *UpdateAccountCenterSettingsRequestFields) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountCenterSettingsRequestFields) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *UpdateAccountCenterSettingsRequestFields) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *UpdateAccountCenterSettingsRequestFields) SetAvatar(v string) {
	o.Avatar = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UpdateAccountCenterSettingsRequestFields) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountCenterSettingsRequestFields) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UpdateAccountCenterSettingsRequestFields) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *UpdateAccountCenterSettingsRequestFields) SetProfile(v string) {
	o.Profile = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateAccountCenterSettingsRequestFields) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountCenterSettingsRequestFields) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateAccountCenterSettingsRequestFields) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateAccountCenterSettingsRequestFields) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UpdateAccountCenterSettingsRequestFields) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountCenterSettingsRequestFields) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UpdateAccountCenterSettingsRequestFields) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UpdateAccountCenterSettingsRequestFields) SetPhone(v string) {
	o.Phone = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateAccountCenterSettingsRequestFields) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountCenterSettingsRequestFields) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateAccountCenterSettingsRequestFields) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateAccountCenterSettingsRequestFields) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateAccountCenterSettingsRequestFields) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountCenterSettingsRequestFields) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateAccountCenterSettingsRequestFields) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateAccountCenterSettingsRequestFields) SetUsername(v string) {
	o.Username = &v
}

// GetSocial returns the Social field value if set, zero value otherwise.
func (o *UpdateAccountCenterSettingsRequestFields) GetSocial() string {
	if o == nil || IsNil(o.Social) {
		var ret string
		return ret
	}
	return *o.Social
}

// GetSocialOk returns a tuple with the Social field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountCenterSettingsRequestFields) GetSocialOk() (*string, bool) {
	if o == nil || IsNil(o.Social) {
		return nil, false
	}
	return o.Social, true
}

// HasSocial returns a boolean if a field has been set.
func (o *UpdateAccountCenterSettingsRequestFields) HasSocial() bool {
	if o != nil && !IsNil(o.Social) {
		return true
	}

	return false
}

// SetSocial gets a reference to the given string and assigns it to the Social field.
func (o *UpdateAccountCenterSettingsRequestFields) SetSocial(v string) {
	o.Social = &v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *UpdateAccountCenterSettingsRequestFields) GetCustomData() string {
	if o == nil || IsNil(o.CustomData) {
		var ret string
		return ret
	}
	return *o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountCenterSettingsRequestFields) GetCustomDataOk() (*string, bool) {
	if o == nil || IsNil(o.CustomData) {
		return nil, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *UpdateAccountCenterSettingsRequestFields) HasCustomData() bool {
	if o != nil && !IsNil(o.CustomData) {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given string and assigns it to the CustomData field.
func (o *UpdateAccountCenterSettingsRequestFields) SetCustomData(v string) {
	o.CustomData = &v
}

func (o UpdateAccountCenterSettingsRequestFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAccountCenterSettingsRequestFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Social) {
		toSerialize["social"] = o.Social
	}
	if !IsNil(o.CustomData) {
		toSerialize["customData"] = o.CustomData
	}
	return toSerialize, nil
}

type NullableUpdateAccountCenterSettingsRequestFields struct {
	value *UpdateAccountCenterSettingsRequestFields
	isSet bool
}

func (v NullableUpdateAccountCenterSettingsRequestFields) Get() *UpdateAccountCenterSettingsRequestFields {
	return v.value
}

func (v *NullableUpdateAccountCenterSettingsRequestFields) Set(val *UpdateAccountCenterSettingsRequestFields) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountCenterSettingsRequestFields) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountCenterSettingsRequestFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountCenterSettingsRequestFields(val *UpdateAccountCenterSettingsRequestFields) *NullableUpdateAccountCenterSettingsRequestFields {
	return &NullableUpdateAccountCenterSettingsRequestFields{value: val, isSet: true}
}

func (v NullableUpdateAccountCenterSettingsRequestFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountCenterSettingsRequestFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


