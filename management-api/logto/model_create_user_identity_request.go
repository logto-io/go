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

// checks if the CreateUserIdentityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserIdentityRequest{}

// CreateUserIdentityRequest struct for CreateUserIdentityRequest
type CreateUserIdentityRequest struct {
	// The Logto connector ID.
	ConnectorId string `json:"connectorId"`
	// A json object constructed from the url query params returned by the social platform. Typically it contains `code`, `state` and `redirectUri` fields.
	ConnectorData map[string]interface{} `json:"connectorData"`
}

type _CreateUserIdentityRequest CreateUserIdentityRequest

// NewCreateUserIdentityRequest instantiates a new CreateUserIdentityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserIdentityRequest(connectorId string, connectorData map[string]interface{}) *CreateUserIdentityRequest {
	this := CreateUserIdentityRequest{}
	this.ConnectorId = connectorId
	this.ConnectorData = connectorData
	return &this
}

// NewCreateUserIdentityRequestWithDefaults instantiates a new CreateUserIdentityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserIdentityRequestWithDefaults() *CreateUserIdentityRequest {
	this := CreateUserIdentityRequest{}
	return &this
}

// GetConnectorId returns the ConnectorId field value
func (o *CreateUserIdentityRequest) GetConnectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *CreateUserIdentityRequest) GetConnectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *CreateUserIdentityRequest) SetConnectorId(v string) {
	o.ConnectorId = v
}

// GetConnectorData returns the ConnectorData field value
func (o *CreateUserIdentityRequest) GetConnectorData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ConnectorData
}

// GetConnectorDataOk returns a tuple with the ConnectorData field value
// and a boolean to check if the value has been set.
func (o *CreateUserIdentityRequest) GetConnectorDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ConnectorData, true
}

// SetConnectorData sets field value
func (o *CreateUserIdentityRequest) SetConnectorData(v map[string]interface{}) {
	o.ConnectorData = v
}

func (o CreateUserIdentityRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserIdentityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectorId"] = o.ConnectorId
	toSerialize["connectorData"] = o.ConnectorData
	return toSerialize, nil
}

func (o *CreateUserIdentityRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectorId",
		"connectorData",
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

	varCreateUserIdentityRequest := _CreateUserIdentityRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateUserIdentityRequest)

	if err != nil {
		return err
	}

	*o = CreateUserIdentityRequest(varCreateUserIdentityRequest)

	return err
}

type NullableCreateUserIdentityRequest struct {
	value *CreateUserIdentityRequest
	isSet bool
}

func (v NullableCreateUserIdentityRequest) Get() *CreateUserIdentityRequest {
	return v.value
}

func (v *NullableCreateUserIdentityRequest) Set(val *CreateUserIdentityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserIdentityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserIdentityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserIdentityRequest(val *CreateUserIdentityRequest) *NullableCreateUserIdentityRequest {
	return &NullableCreateUserIdentityRequest{value: val, isSet: true}
}

func (v NullableCreateUserIdentityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserIdentityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


