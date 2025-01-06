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

// ListApplicationsTypesParameter - struct for ListApplicationsTypesParameter
type ListApplicationsTypesParameter struct {
	ArrayOfString *[]string
	String *string
}

// []stringAsListApplicationsTypesParameter is a convenience function that returns []string wrapped in ListApplicationsTypesParameter
func ArrayOfStringAsListApplicationsTypesParameter(v *[]string) ListApplicationsTypesParameter {
	return ListApplicationsTypesParameter{
		ArrayOfString: v,
	}
}

// stringAsListApplicationsTypesParameter is a convenience function that returns string wrapped in ListApplicationsTypesParameter
func StringAsListApplicationsTypesParameter(v *string) ListApplicationsTypesParameter {
	return ListApplicationsTypesParameter{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListApplicationsTypesParameter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			if err = validator.Validate(dst.ArrayOfString); err != nil {
				dst.ArrayOfString = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfString = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListApplicationsTypesParameter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListApplicationsTypesParameter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListApplicationsTypesParameter) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListApplicationsTypesParameter) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableListApplicationsTypesParameter struct {
	value *ListApplicationsTypesParameter
	isSet bool
}

func (v NullableListApplicationsTypesParameter) Get() *ListApplicationsTypesParameter {
	return v.value
}

func (v *NullableListApplicationsTypesParameter) Set(val *ListApplicationsTypesParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableListApplicationsTypesParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableListApplicationsTypesParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApplicationsTypesParameter(val *ListApplicationsTypesParameter) *NullableListApplicationsTypesParameter {
	return &NullableListApplicationsTypesParameter{value: val, isSet: true}
}

func (v NullableListApplicationsTypesParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApplicationsTypesParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


