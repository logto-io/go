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

// checks if the ApiInteractionPutRequestIdentifierOneOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionPutRequestIdentifierOneOf3{}

// ApiInteractionPutRequestIdentifierOneOf3 struct for ApiInteractionPutRequestIdentifierOneOf3
type ApiInteractionPutRequestIdentifierOneOf3 struct {
	ConnectorId string `json:"connectorId"`
	// arbitrary
	ConnectorData map[string]interface{} `json:"connectorData"`
}

type _ApiInteractionPutRequestIdentifierOneOf3 ApiInteractionPutRequestIdentifierOneOf3

// NewApiInteractionPutRequestIdentifierOneOf3 instantiates a new ApiInteractionPutRequestIdentifierOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionPutRequestIdentifierOneOf3(connectorId string, connectorData map[string]interface{}) *ApiInteractionPutRequestIdentifierOneOf3 {
	this := ApiInteractionPutRequestIdentifierOneOf3{}
	this.ConnectorId = connectorId
	this.ConnectorData = connectorData
	return &this
}

// NewApiInteractionPutRequestIdentifierOneOf3WithDefaults instantiates a new ApiInteractionPutRequestIdentifierOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionPutRequestIdentifierOneOf3WithDefaults() *ApiInteractionPutRequestIdentifierOneOf3 {
	this := ApiInteractionPutRequestIdentifierOneOf3{}
	return &this
}

// GetConnectorId returns the ConnectorId field value
func (o *ApiInteractionPutRequestIdentifierOneOf3) GetConnectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionPutRequestIdentifierOneOf3) GetConnectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *ApiInteractionPutRequestIdentifierOneOf3) SetConnectorId(v string) {
	o.ConnectorId = v
}

// GetConnectorData returns the ConnectorData field value
func (o *ApiInteractionPutRequestIdentifierOneOf3) GetConnectorData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ConnectorData
}

// GetConnectorDataOk returns a tuple with the ConnectorData field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionPutRequestIdentifierOneOf3) GetConnectorDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ConnectorData, true
}

// SetConnectorData sets field value
func (o *ApiInteractionPutRequestIdentifierOneOf3) SetConnectorData(v map[string]interface{}) {
	o.ConnectorData = v
}

func (o ApiInteractionPutRequestIdentifierOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionPutRequestIdentifierOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectorId"] = o.ConnectorId
	toSerialize["connectorData"] = o.ConnectorData
	return toSerialize, nil
}

func (o *ApiInteractionPutRequestIdentifierOneOf3) UnmarshalJSON(data []byte) (err error) {
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

	varApiInteractionPutRequestIdentifierOneOf3 := _ApiInteractionPutRequestIdentifierOneOf3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiInteractionPutRequestIdentifierOneOf3)

	if err != nil {
		return err
	}

	*o = ApiInteractionPutRequestIdentifierOneOf3(varApiInteractionPutRequestIdentifierOneOf3)

	return err
}

type NullableApiInteractionPutRequestIdentifierOneOf3 struct {
	value *ApiInteractionPutRequestIdentifierOneOf3
	isSet bool
}

func (v NullableApiInteractionPutRequestIdentifierOneOf3) Get() *ApiInteractionPutRequestIdentifierOneOf3 {
	return v.value
}

func (v *NullableApiInteractionPutRequestIdentifierOneOf3) Set(val *ApiInteractionPutRequestIdentifierOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionPutRequestIdentifierOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionPutRequestIdentifierOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionPutRequestIdentifierOneOf3(val *ApiInteractionPutRequestIdentifierOneOf3) *NullableApiInteractionPutRequestIdentifierOneOf3 {
	return &NullableApiInteractionPutRequestIdentifierOneOf3{value: val, isSet: true}
}

func (v NullableApiInteractionPutRequestIdentifierOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionPutRequestIdentifierOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


