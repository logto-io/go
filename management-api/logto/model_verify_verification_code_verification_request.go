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

// checks if the VerifyVerificationCodeVerificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyVerificationCodeVerificationRequest{}

// VerifyVerificationCodeVerificationRequest struct for VerifyVerificationCodeVerificationRequest
type VerifyVerificationCodeVerificationRequest struct {
	Identifier VerifyVerificationCodeVerificationRequestIdentifier `json:"identifier"`
	// The verification ID of the CodeVerification record.
	VerificationId string `json:"verificationId"`
	// The verification code to be verified.
	Code string `json:"code"`
}

type _VerifyVerificationCodeVerificationRequest VerifyVerificationCodeVerificationRequest

// NewVerifyVerificationCodeVerificationRequest instantiates a new VerifyVerificationCodeVerificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyVerificationCodeVerificationRequest(identifier VerifyVerificationCodeVerificationRequestIdentifier, verificationId string, code string) *VerifyVerificationCodeVerificationRequest {
	this := VerifyVerificationCodeVerificationRequest{}
	this.Identifier = identifier
	this.VerificationId = verificationId
	this.Code = code
	return &this
}

// NewVerifyVerificationCodeVerificationRequestWithDefaults instantiates a new VerifyVerificationCodeVerificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyVerificationCodeVerificationRequestWithDefaults() *VerifyVerificationCodeVerificationRequest {
	this := VerifyVerificationCodeVerificationRequest{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *VerifyVerificationCodeVerificationRequest) GetIdentifier() VerifyVerificationCodeVerificationRequestIdentifier {
	if o == nil {
		var ret VerifyVerificationCodeVerificationRequestIdentifier
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *VerifyVerificationCodeVerificationRequest) GetIdentifierOk() (*VerifyVerificationCodeVerificationRequestIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *VerifyVerificationCodeVerificationRequest) SetIdentifier(v VerifyVerificationCodeVerificationRequestIdentifier) {
	o.Identifier = v
}

// GetVerificationId returns the VerificationId field value
func (o *VerifyVerificationCodeVerificationRequest) GetVerificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationId
}

// GetVerificationIdOk returns a tuple with the VerificationId field value
// and a boolean to check if the value has been set.
func (o *VerifyVerificationCodeVerificationRequest) GetVerificationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationId, true
}

// SetVerificationId sets field value
func (o *VerifyVerificationCodeVerificationRequest) SetVerificationId(v string) {
	o.VerificationId = v
}

// GetCode returns the Code field value
func (o *VerifyVerificationCodeVerificationRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *VerifyVerificationCodeVerificationRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *VerifyVerificationCodeVerificationRequest) SetCode(v string) {
	o.Code = v
}

func (o VerifyVerificationCodeVerificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyVerificationCodeVerificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifier"] = o.Identifier
	toSerialize["verificationId"] = o.VerificationId
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

func (o *VerifyVerificationCodeVerificationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identifier",
		"verificationId",
		"code",
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

	varVerifyVerificationCodeVerificationRequest := _VerifyVerificationCodeVerificationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVerifyVerificationCodeVerificationRequest)

	if err != nil {
		return err
	}

	*o = VerifyVerificationCodeVerificationRequest(varVerifyVerificationCodeVerificationRequest)

	return err
}

type NullableVerifyVerificationCodeVerificationRequest struct {
	value *VerifyVerificationCodeVerificationRequest
	isSet bool
}

func (v NullableVerifyVerificationCodeVerificationRequest) Get() *VerifyVerificationCodeVerificationRequest {
	return v.value
}

func (v *NullableVerifyVerificationCodeVerificationRequest) Set(val *VerifyVerificationCodeVerificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyVerificationCodeVerificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyVerificationCodeVerificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyVerificationCodeVerificationRequest(val *VerifyVerificationCodeVerificationRequest) *NullableVerifyVerificationCodeVerificationRequest {
	return &NullableVerifyVerificationCodeVerificationRequest{value: val, isSet: true}
}

func (v NullableVerifyVerificationCodeVerificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyVerificationCodeVerificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


