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

// checks if the CreateTotpSecret200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTotpSecret200Response{}

// CreateTotpSecret200Response struct for CreateTotpSecret200Response
type CreateTotpSecret200Response struct {
	// The unique verification ID for the TOTP record. This ID is required to verify the TOTP code.
	VerificationId string `json:"verificationId"`
	// The newly generated TOTP secret.
	Secret string `json:"secret"`
	// A QR code image data URL for the TOTP secret. The user can scan this QR code with their TOTP authenticator app.
	SecretQrCode string `json:"secretQrCode"`
}

type _CreateTotpSecret200Response CreateTotpSecret200Response

// NewCreateTotpSecret200Response instantiates a new CreateTotpSecret200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTotpSecret200Response(verificationId string, secret string, secretQrCode string) *CreateTotpSecret200Response {
	this := CreateTotpSecret200Response{}
	this.VerificationId = verificationId
	this.Secret = secret
	this.SecretQrCode = secretQrCode
	return &this
}

// NewCreateTotpSecret200ResponseWithDefaults instantiates a new CreateTotpSecret200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTotpSecret200ResponseWithDefaults() *CreateTotpSecret200Response {
	this := CreateTotpSecret200Response{}
	return &this
}

// GetVerificationId returns the VerificationId field value
func (o *CreateTotpSecret200Response) GetVerificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationId
}

// GetVerificationIdOk returns a tuple with the VerificationId field value
// and a boolean to check if the value has been set.
func (o *CreateTotpSecret200Response) GetVerificationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationId, true
}

// SetVerificationId sets field value
func (o *CreateTotpSecret200Response) SetVerificationId(v string) {
	o.VerificationId = v
}

// GetSecret returns the Secret field value
func (o *CreateTotpSecret200Response) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CreateTotpSecret200Response) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CreateTotpSecret200Response) SetSecret(v string) {
	o.Secret = v
}

// GetSecretQrCode returns the SecretQrCode field value
func (o *CreateTotpSecret200Response) GetSecretQrCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretQrCode
}

// GetSecretQrCodeOk returns a tuple with the SecretQrCode field value
// and a boolean to check if the value has been set.
func (o *CreateTotpSecret200Response) GetSecretQrCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretQrCode, true
}

// SetSecretQrCode sets field value
func (o *CreateTotpSecret200Response) SetSecretQrCode(v string) {
	o.SecretQrCode = v
}

func (o CreateTotpSecret200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTotpSecret200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["verificationId"] = o.VerificationId
	toSerialize["secret"] = o.Secret
	toSerialize["secretQrCode"] = o.SecretQrCode
	return toSerialize, nil
}

func (o *CreateTotpSecret200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"verificationId",
		"secret",
		"secretQrCode",
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

	varCreateTotpSecret200Response := _CreateTotpSecret200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTotpSecret200Response)

	if err != nil {
		return err
	}

	*o = CreateTotpSecret200Response(varCreateTotpSecret200Response)

	return err
}

type NullableCreateTotpSecret200Response struct {
	value *CreateTotpSecret200Response
	isSet bool
}

func (v NullableCreateTotpSecret200Response) Get() *CreateTotpSecret200Response {
	return v.value
}

func (v *NullableCreateTotpSecret200Response) Set(val *CreateTotpSecret200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTotpSecret200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTotpSecret200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTotpSecret200Response(val *CreateTotpSecret200Response) *NullableCreateTotpSecret200Response {
	return &NullableCreateTotpSecret200Response{value: val, isSet: true}
}

func (v NullableCreateTotpSecret200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTotpSecret200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


