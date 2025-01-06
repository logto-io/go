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

// checks if the VerifyBackupCode200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyBackupCode200Response{}

// VerifyBackupCode200Response struct for VerifyBackupCode200Response
type VerifyBackupCode200Response struct {
	// The unique verification ID of the BackupCode verification record.
	VerificationId string `json:"verificationId"`
}

type _VerifyBackupCode200Response VerifyBackupCode200Response

// NewVerifyBackupCode200Response instantiates a new VerifyBackupCode200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyBackupCode200Response(verificationId string) *VerifyBackupCode200Response {
	this := VerifyBackupCode200Response{}
	this.VerificationId = verificationId
	return &this
}

// NewVerifyBackupCode200ResponseWithDefaults instantiates a new VerifyBackupCode200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyBackupCode200ResponseWithDefaults() *VerifyBackupCode200Response {
	this := VerifyBackupCode200Response{}
	return &this
}

// GetVerificationId returns the VerificationId field value
func (o *VerifyBackupCode200Response) GetVerificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationId
}

// GetVerificationIdOk returns a tuple with the VerificationId field value
// and a boolean to check if the value has been set.
func (o *VerifyBackupCode200Response) GetVerificationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationId, true
}

// SetVerificationId sets field value
func (o *VerifyBackupCode200Response) SetVerificationId(v string) {
	o.VerificationId = v
}

func (o VerifyBackupCode200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyBackupCode200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["verificationId"] = o.VerificationId
	return toSerialize, nil
}

func (o *VerifyBackupCode200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"verificationId",
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

	varVerifyBackupCode200Response := _VerifyBackupCode200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVerifyBackupCode200Response)

	if err != nil {
		return err
	}

	*o = VerifyBackupCode200Response(varVerifyBackupCode200Response)

	return err
}

type NullableVerifyBackupCode200Response struct {
	value *VerifyBackupCode200Response
	isSet bool
}

func (v NullableVerifyBackupCode200Response) Get() *VerifyBackupCode200Response {
	return v.value
}

func (v *NullableVerifyBackupCode200Response) Set(val *VerifyBackupCode200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyBackupCode200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyBackupCode200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyBackupCode200Response(val *VerifyBackupCode200Response) *NullableVerifyBackupCode200Response {
	return &NullableVerifyBackupCode200Response{value: val, isSet: true}
}

func (v NullableVerifyBackupCode200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyBackupCode200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


