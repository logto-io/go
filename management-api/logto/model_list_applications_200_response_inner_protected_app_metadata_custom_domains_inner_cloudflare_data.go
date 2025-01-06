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

// checks if the ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData{}

// ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData struct for ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData
type ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Ssl ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareDataSsl `json:"ssl"`
	VerificationErrors []string `json:"verification_errors,omitempty"`
}

type _ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData

// NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData instantiates a new ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData(id string, status string, ssl ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareDataSsl) *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData {
	this := ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData{}
	this.Id = id
	this.Status = status
	this.Ssl = ssl
	return &this
}

// NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareDataWithDefaults instantiates a new ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareDataWithDefaults() *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData {
	this := ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData{}
	return &this
}

// GetId returns the Id field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) SetStatus(v string) {
	o.Status = v
}

// GetSsl returns the Ssl field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) GetSsl() ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareDataSsl {
	if o == nil {
		var ret ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareDataSsl
		return ret
	}

	return o.Ssl
}

// GetSslOk returns a tuple with the Ssl field value
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) GetSslOk() (*ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareDataSsl, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssl, true
}

// SetSsl sets field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) SetSsl(v ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareDataSsl) {
	o.Ssl = v
}

// GetVerificationErrors returns the VerificationErrors field value if set, zero value otherwise.
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) GetVerificationErrors() []string {
	if o == nil || IsNil(o.VerificationErrors) {
		var ret []string
		return ret
	}
	return o.VerificationErrors
}

// GetVerificationErrorsOk returns a tuple with the VerificationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) GetVerificationErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.VerificationErrors) {
		return nil, false
	}
	return o.VerificationErrors, true
}

// HasVerificationErrors returns a boolean if a field has been set.
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) HasVerificationErrors() bool {
	if o != nil && !IsNil(o.VerificationErrors) {
		return true
	}

	return false
}

// SetVerificationErrors gets a reference to the given []string and assigns it to the VerificationErrors field.
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) SetVerificationErrors(v []string) {
	o.VerificationErrors = v
}

func (o ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["ssl"] = o.Ssl
	if !IsNil(o.VerificationErrors) {
		toSerialize["verification_errors"] = o.VerificationErrors
	}
	return toSerialize, nil
}

func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"ssl",
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

	varListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData := _ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData)

	if err != nil {
		return err
	}

	*o = ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData(varListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData)

	return err
}

type NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData struct {
	value *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData
	isSet bool
}

func (v NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) Get() *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData {
	return v.value
}

func (v *NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) Set(val *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) {
	v.value = val
	v.isSet = true
}

func (v NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) IsSet() bool {
	return v.isSet
}

func (v *NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData(val *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) *NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData {
	return &NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData{value: val, isSet: true}
}

func (v NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


