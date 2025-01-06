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

// checks if the GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress{}

// GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress struct for GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress
type GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress struct {
	Formatted *string `json:"formatted,omitempty"`
	StreetAddress *string `json:"streetAddress,omitempty"`
	Locality *string `json:"locality,omitempty"`
	Region *string `json:"region,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	Country *string `json:"country,omitempty"`
}

// NewGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress instantiates a new GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress() *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress {
	this := GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress{}
	return &this
}

// NewGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddressWithDefaults instantiates a new GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddressWithDefaults() *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress {
	this := GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress{}
	return &this
}

// GetFormatted returns the Formatted field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) GetFormatted() string {
	if o == nil || IsNil(o.Formatted) {
		var ret string
		return ret
	}
	return *o.Formatted
}

// GetFormattedOk returns a tuple with the Formatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) GetFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.Formatted) {
		return nil, false
	}
	return o.Formatted, true
}

// HasFormatted returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) HasFormatted() bool {
	if o != nil && !IsNil(o.Formatted) {
		return true
	}

	return false
}

// SetFormatted gets a reference to the given string and assigns it to the Formatted field.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) SetFormatted(v string) {
	o.Formatted = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) GetStreetAddress() string {
	if o == nil || IsNil(o.StreetAddress) {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) GetStreetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StreetAddress) {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) HasStreetAddress() bool {
	if o != nil && !IsNil(o.StreetAddress) {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) GetLocality() string {
	if o == nil || IsNil(o.Locality) {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) GetLocalityOk() (*string, bool) {
	if o == nil || IsNil(o.Locality) {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) HasLocality() bool {
	if o != nil && !IsNil(o.Locality) {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) SetLocality(v string) {
	o.Locality = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) SetRegion(v string) {
	o.Region = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) SetCountry(v string) {
	o.Country = &v
}

func (o GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Formatted) {
		toSerialize["formatted"] = o.Formatted
	}
	if !IsNil(o.StreetAddress) {
		toSerialize["streetAddress"] = o.StreetAddress
	}
	if !IsNil(o.Locality) {
		toSerialize["locality"] = o.Locality
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	return toSerialize, nil
}

type NullableGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress struct {
	value *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress
	isSet bool
}

func (v NullableGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) Get() *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress {
	return v.value
}

func (v *NullableGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) Set(val *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress(val *GetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) *NullableGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress {
	return &NullableGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress{value: val, isSet: true}
}

func (v NullableGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetJwtCustomizer200ResponseOneOfContextSampleUserProfileAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


