/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logto

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner - struct for ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner
type ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner struct {
	MapmapOfStringAny *map[string]interface{}
}

// map[string]interface{}AsListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner is a convenience function that returns map[string]interface{} wrapped in ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner
func MapmapOfStringAnyAsListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner(v *map[string]interface{}) ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner {
	return ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner{
		MapmapOfStringAny: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringAny); err != nil {
				dst.MapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner) MarshalJSON() ([]byte, error) {
	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner struct {
	value *ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner
	isSet bool
}

func (v NullableListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner) Get() *ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner {
	return v.value
}

func (v *NullableListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner) Set(val *ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner(val *ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner) *NullableListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner {
	return &NullableListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner{value: val, isSet: true}
}

func (v NullableListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


