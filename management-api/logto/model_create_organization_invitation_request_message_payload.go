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

// CreateOrganizationInvitationRequestMessagePayload - The message payload for the \"OrganizationInvitation\" template to use when sending the invitation via email. If it is `false`, the invitation will not be sent via email.
type CreateOrganizationInvitationRequestMessagePayload struct {
	CreateOrganizationInvitationRequestMessagePayloadOneOf *CreateOrganizationInvitationRequestMessagePayloadOneOf
	Bool *bool
}

// CreateOrganizationInvitationRequestMessagePayloadOneOfAsCreateOrganizationInvitationRequestMessagePayload is a convenience function that returns CreateOrganizationInvitationRequestMessagePayloadOneOf wrapped in CreateOrganizationInvitationRequestMessagePayload
func CreateOrganizationInvitationRequestMessagePayloadOneOfAsCreateOrganizationInvitationRequestMessagePayload(v *CreateOrganizationInvitationRequestMessagePayloadOneOf) CreateOrganizationInvitationRequestMessagePayload {
	return CreateOrganizationInvitationRequestMessagePayload{
		CreateOrganizationInvitationRequestMessagePayloadOneOf: v,
	}
}

// boolAsCreateOrganizationInvitationRequestMessagePayload is a convenience function that returns bool wrapped in CreateOrganizationInvitationRequestMessagePayload
func BoolAsCreateOrganizationInvitationRequestMessagePayload(v *bool) CreateOrganizationInvitationRequestMessagePayload {
	return CreateOrganizationInvitationRequestMessagePayload{
		Bool: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateOrganizationInvitationRequestMessagePayload) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateOrganizationInvitationRequestMessagePayloadOneOf
	err = newStrictDecoder(data).Decode(&dst.CreateOrganizationInvitationRequestMessagePayloadOneOf)
	if err == nil {
		jsonCreateOrganizationInvitationRequestMessagePayloadOneOf, _ := json.Marshal(dst.CreateOrganizationInvitationRequestMessagePayloadOneOf)
		if string(jsonCreateOrganizationInvitationRequestMessagePayloadOneOf) == "{}" { // empty struct
			dst.CreateOrganizationInvitationRequestMessagePayloadOneOf = nil
		} else {
			if err = validator.Validate(dst.CreateOrganizationInvitationRequestMessagePayloadOneOf); err != nil {
				dst.CreateOrganizationInvitationRequestMessagePayloadOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.CreateOrganizationInvitationRequestMessagePayloadOneOf = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			if err = validator.Validate(dst.Bool); err != nil {
				dst.Bool = nil
			} else {
				match++
			}
		}
	} else {
		dst.Bool = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateOrganizationInvitationRequestMessagePayloadOneOf = nil
		dst.Bool = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateOrganizationInvitationRequestMessagePayload)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateOrganizationInvitationRequestMessagePayload)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateOrganizationInvitationRequestMessagePayload) MarshalJSON() ([]byte, error) {
	if src.CreateOrganizationInvitationRequestMessagePayloadOneOf != nil {
		return json.Marshal(&src.CreateOrganizationInvitationRequestMessagePayloadOneOf)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateOrganizationInvitationRequestMessagePayload) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CreateOrganizationInvitationRequestMessagePayloadOneOf != nil {
		return obj.CreateOrganizationInvitationRequestMessagePayloadOneOf
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	// all schemas are nil
	return nil
}

type NullableCreateOrganizationInvitationRequestMessagePayload struct {
	value *CreateOrganizationInvitationRequestMessagePayload
	isSet bool
}

func (v NullableCreateOrganizationInvitationRequestMessagePayload) Get() *CreateOrganizationInvitationRequestMessagePayload {
	return v.value
}

func (v *NullableCreateOrganizationInvitationRequestMessagePayload) Set(val *CreateOrganizationInvitationRequestMessagePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInvitationRequestMessagePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInvitationRequestMessagePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInvitationRequestMessagePayload(val *CreateOrganizationInvitationRequestMessagePayload) *NullableCreateOrganizationInvitationRequestMessagePayload {
	return &NullableCreateOrganizationInvitationRequestMessagePayload{value: val, isSet: true}
}

func (v NullableCreateOrganizationInvitationRequestMessagePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInvitationRequestMessagePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


