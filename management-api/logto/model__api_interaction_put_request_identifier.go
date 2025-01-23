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

// ApiInteractionPutRequestIdentifier - struct for ApiInteractionPutRequestIdentifier
type ApiInteractionPutRequestIdentifier struct {
	ApiInteractionPutRequestIdentifierOneOf *ApiInteractionPutRequestIdentifierOneOf
	ApiInteractionPutRequestIdentifierOneOf1 *ApiInteractionPutRequestIdentifierOneOf1
	ApiInteractionPutRequestIdentifierOneOf2 *ApiInteractionPutRequestIdentifierOneOf2
	ApiInteractionPutRequestIdentifierOneOf3 *ApiInteractionPutRequestIdentifierOneOf3
	ApiInteractionPutRequestIdentifierOneOf4 *ApiInteractionPutRequestIdentifierOneOf4
	ApiInteractionPutRequestIdentifierOneOf5 *ApiInteractionPutRequestIdentifierOneOf5
	VerifyVerificationCodeRequestOneOf *VerifyVerificationCodeRequestOneOf
	VerifyVerificationCodeRequestOneOf1 *VerifyVerificationCodeRequestOneOf1
}

// ApiInteractionPutRequestIdentifierOneOfAsApiInteractionPutRequestIdentifier is a convenience function that returns ApiInteractionPutRequestIdentifierOneOf wrapped in ApiInteractionPutRequestIdentifier
func ApiInteractionPutRequestIdentifierOneOfAsApiInteractionPutRequestIdentifier(v *ApiInteractionPutRequestIdentifierOneOf) ApiInteractionPutRequestIdentifier {
	return ApiInteractionPutRequestIdentifier{
		ApiInteractionPutRequestIdentifierOneOf: v,
	}
}

// ApiInteractionPutRequestIdentifierOneOf1AsApiInteractionPutRequestIdentifier is a convenience function that returns ApiInteractionPutRequestIdentifierOneOf1 wrapped in ApiInteractionPutRequestIdentifier
func ApiInteractionPutRequestIdentifierOneOf1AsApiInteractionPutRequestIdentifier(v *ApiInteractionPutRequestIdentifierOneOf1) ApiInteractionPutRequestIdentifier {
	return ApiInteractionPutRequestIdentifier{
		ApiInteractionPutRequestIdentifierOneOf1: v,
	}
}

// ApiInteractionPutRequestIdentifierOneOf2AsApiInteractionPutRequestIdentifier is a convenience function that returns ApiInteractionPutRequestIdentifierOneOf2 wrapped in ApiInteractionPutRequestIdentifier
func ApiInteractionPutRequestIdentifierOneOf2AsApiInteractionPutRequestIdentifier(v *ApiInteractionPutRequestIdentifierOneOf2) ApiInteractionPutRequestIdentifier {
	return ApiInteractionPutRequestIdentifier{
		ApiInteractionPutRequestIdentifierOneOf2: v,
	}
}

// ApiInteractionPutRequestIdentifierOneOf3AsApiInteractionPutRequestIdentifier is a convenience function that returns ApiInteractionPutRequestIdentifierOneOf3 wrapped in ApiInteractionPutRequestIdentifier
func ApiInteractionPutRequestIdentifierOneOf3AsApiInteractionPutRequestIdentifier(v *ApiInteractionPutRequestIdentifierOneOf3) ApiInteractionPutRequestIdentifier {
	return ApiInteractionPutRequestIdentifier{
		ApiInteractionPutRequestIdentifierOneOf3: v,
	}
}

// ApiInteractionPutRequestIdentifierOneOf4AsApiInteractionPutRequestIdentifier is a convenience function that returns ApiInteractionPutRequestIdentifierOneOf4 wrapped in ApiInteractionPutRequestIdentifier
func ApiInteractionPutRequestIdentifierOneOf4AsApiInteractionPutRequestIdentifier(v *ApiInteractionPutRequestIdentifierOneOf4) ApiInteractionPutRequestIdentifier {
	return ApiInteractionPutRequestIdentifier{
		ApiInteractionPutRequestIdentifierOneOf4: v,
	}
}

// ApiInteractionPutRequestIdentifierOneOf5AsApiInteractionPutRequestIdentifier is a convenience function that returns ApiInteractionPutRequestIdentifierOneOf5 wrapped in ApiInteractionPutRequestIdentifier
func ApiInteractionPutRequestIdentifierOneOf5AsApiInteractionPutRequestIdentifier(v *ApiInteractionPutRequestIdentifierOneOf5) ApiInteractionPutRequestIdentifier {
	return ApiInteractionPutRequestIdentifier{
		ApiInteractionPutRequestIdentifierOneOf5: v,
	}
}

// VerifyVerificationCodeRequestOneOfAsApiInteractionPutRequestIdentifier is a convenience function that returns VerifyVerificationCodeRequestOneOf wrapped in ApiInteractionPutRequestIdentifier
func VerifyVerificationCodeRequestOneOfAsApiInteractionPutRequestIdentifier(v *VerifyVerificationCodeRequestOneOf) ApiInteractionPutRequestIdentifier {
	return ApiInteractionPutRequestIdentifier{
		VerifyVerificationCodeRequestOneOf: v,
	}
}

// VerifyVerificationCodeRequestOneOf1AsApiInteractionPutRequestIdentifier is a convenience function that returns VerifyVerificationCodeRequestOneOf1 wrapped in ApiInteractionPutRequestIdentifier
func VerifyVerificationCodeRequestOneOf1AsApiInteractionPutRequestIdentifier(v *VerifyVerificationCodeRequestOneOf1) ApiInteractionPutRequestIdentifier {
	return ApiInteractionPutRequestIdentifier{
		VerifyVerificationCodeRequestOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApiInteractionPutRequestIdentifier) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApiInteractionPutRequestIdentifierOneOf
	err = newStrictDecoder(data).Decode(&dst.ApiInteractionPutRequestIdentifierOneOf)
	if err == nil {
		jsonApiInteractionPutRequestIdentifierOneOf, _ := json.Marshal(dst.ApiInteractionPutRequestIdentifierOneOf)
		if string(jsonApiInteractionPutRequestIdentifierOneOf) == "{}" { // empty struct
			dst.ApiInteractionPutRequestIdentifierOneOf = nil
		} else {
			if err = validator.Validate(dst.ApiInteractionPutRequestIdentifierOneOf); err != nil {
				dst.ApiInteractionPutRequestIdentifierOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiInteractionPutRequestIdentifierOneOf = nil
	}

	// try to unmarshal data into ApiInteractionPutRequestIdentifierOneOf1
	err = newStrictDecoder(data).Decode(&dst.ApiInteractionPutRequestIdentifierOneOf1)
	if err == nil {
		jsonApiInteractionPutRequestIdentifierOneOf1, _ := json.Marshal(dst.ApiInteractionPutRequestIdentifierOneOf1)
		if string(jsonApiInteractionPutRequestIdentifierOneOf1) == "{}" { // empty struct
			dst.ApiInteractionPutRequestIdentifierOneOf1 = nil
		} else {
			if err = validator.Validate(dst.ApiInteractionPutRequestIdentifierOneOf1); err != nil {
				dst.ApiInteractionPutRequestIdentifierOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiInteractionPutRequestIdentifierOneOf1 = nil
	}

	// try to unmarshal data into ApiInteractionPutRequestIdentifierOneOf2
	err = newStrictDecoder(data).Decode(&dst.ApiInteractionPutRequestIdentifierOneOf2)
	if err == nil {
		jsonApiInteractionPutRequestIdentifierOneOf2, _ := json.Marshal(dst.ApiInteractionPutRequestIdentifierOneOf2)
		if string(jsonApiInteractionPutRequestIdentifierOneOf2) == "{}" { // empty struct
			dst.ApiInteractionPutRequestIdentifierOneOf2 = nil
		} else {
			if err = validator.Validate(dst.ApiInteractionPutRequestIdentifierOneOf2); err != nil {
				dst.ApiInteractionPutRequestIdentifierOneOf2 = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiInteractionPutRequestIdentifierOneOf2 = nil
	}

	// try to unmarshal data into ApiInteractionPutRequestIdentifierOneOf3
	err = newStrictDecoder(data).Decode(&dst.ApiInteractionPutRequestIdentifierOneOf3)
	if err == nil {
		jsonApiInteractionPutRequestIdentifierOneOf3, _ := json.Marshal(dst.ApiInteractionPutRequestIdentifierOneOf3)
		if string(jsonApiInteractionPutRequestIdentifierOneOf3) == "{}" { // empty struct
			dst.ApiInteractionPutRequestIdentifierOneOf3 = nil
		} else {
			if err = validator.Validate(dst.ApiInteractionPutRequestIdentifierOneOf3); err != nil {
				dst.ApiInteractionPutRequestIdentifierOneOf3 = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiInteractionPutRequestIdentifierOneOf3 = nil
	}

	// try to unmarshal data into ApiInteractionPutRequestIdentifierOneOf4
	err = newStrictDecoder(data).Decode(&dst.ApiInteractionPutRequestIdentifierOneOf4)
	if err == nil {
		jsonApiInteractionPutRequestIdentifierOneOf4, _ := json.Marshal(dst.ApiInteractionPutRequestIdentifierOneOf4)
		if string(jsonApiInteractionPutRequestIdentifierOneOf4) == "{}" { // empty struct
			dst.ApiInteractionPutRequestIdentifierOneOf4 = nil
		} else {
			if err = validator.Validate(dst.ApiInteractionPutRequestIdentifierOneOf4); err != nil {
				dst.ApiInteractionPutRequestIdentifierOneOf4 = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiInteractionPutRequestIdentifierOneOf4 = nil
	}

	// try to unmarshal data into ApiInteractionPutRequestIdentifierOneOf5
	err = newStrictDecoder(data).Decode(&dst.ApiInteractionPutRequestIdentifierOneOf5)
	if err == nil {
		jsonApiInteractionPutRequestIdentifierOneOf5, _ := json.Marshal(dst.ApiInteractionPutRequestIdentifierOneOf5)
		if string(jsonApiInteractionPutRequestIdentifierOneOf5) == "{}" { // empty struct
			dst.ApiInteractionPutRequestIdentifierOneOf5 = nil
		} else {
			if err = validator.Validate(dst.ApiInteractionPutRequestIdentifierOneOf5); err != nil {
				dst.ApiInteractionPutRequestIdentifierOneOf5 = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApiInteractionPutRequestIdentifierOneOf5 = nil
	}

	// try to unmarshal data into VerifyVerificationCodeRequestOneOf
	err = newStrictDecoder(data).Decode(&dst.VerifyVerificationCodeRequestOneOf)
	if err == nil {
		jsonVerifyVerificationCodeRequestOneOf, _ := json.Marshal(dst.VerifyVerificationCodeRequestOneOf)
		if string(jsonVerifyVerificationCodeRequestOneOf) == "{}" { // empty struct
			dst.VerifyVerificationCodeRequestOneOf = nil
		} else {
			if err = validator.Validate(dst.VerifyVerificationCodeRequestOneOf); err != nil {
				dst.VerifyVerificationCodeRequestOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.VerifyVerificationCodeRequestOneOf = nil
	}

	// try to unmarshal data into VerifyVerificationCodeRequestOneOf1
	err = newStrictDecoder(data).Decode(&dst.VerifyVerificationCodeRequestOneOf1)
	if err == nil {
		jsonVerifyVerificationCodeRequestOneOf1, _ := json.Marshal(dst.VerifyVerificationCodeRequestOneOf1)
		if string(jsonVerifyVerificationCodeRequestOneOf1) == "{}" { // empty struct
			dst.VerifyVerificationCodeRequestOneOf1 = nil
		} else {
			if err = validator.Validate(dst.VerifyVerificationCodeRequestOneOf1); err != nil {
				dst.VerifyVerificationCodeRequestOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.VerifyVerificationCodeRequestOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApiInteractionPutRequestIdentifierOneOf = nil
		dst.ApiInteractionPutRequestIdentifierOneOf1 = nil
		dst.ApiInteractionPutRequestIdentifierOneOf2 = nil
		dst.ApiInteractionPutRequestIdentifierOneOf3 = nil
		dst.ApiInteractionPutRequestIdentifierOneOf4 = nil
		dst.ApiInteractionPutRequestIdentifierOneOf5 = nil
		dst.VerifyVerificationCodeRequestOneOf = nil
		dst.VerifyVerificationCodeRequestOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApiInteractionPutRequestIdentifier)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApiInteractionPutRequestIdentifier)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApiInteractionPutRequestIdentifier) MarshalJSON() ([]byte, error) {
	if src.ApiInteractionPutRequestIdentifierOneOf != nil {
		return json.Marshal(&src.ApiInteractionPutRequestIdentifierOneOf)
	}

	if src.ApiInteractionPutRequestIdentifierOneOf1 != nil {
		return json.Marshal(&src.ApiInteractionPutRequestIdentifierOneOf1)
	}

	if src.ApiInteractionPutRequestIdentifierOneOf2 != nil {
		return json.Marshal(&src.ApiInteractionPutRequestIdentifierOneOf2)
	}

	if src.ApiInteractionPutRequestIdentifierOneOf3 != nil {
		return json.Marshal(&src.ApiInteractionPutRequestIdentifierOneOf3)
	}

	if src.ApiInteractionPutRequestIdentifierOneOf4 != nil {
		return json.Marshal(&src.ApiInteractionPutRequestIdentifierOneOf4)
	}

	if src.ApiInteractionPutRequestIdentifierOneOf5 != nil {
		return json.Marshal(&src.ApiInteractionPutRequestIdentifierOneOf5)
	}

	if src.VerifyVerificationCodeRequestOneOf != nil {
		return json.Marshal(&src.VerifyVerificationCodeRequestOneOf)
	}

	if src.VerifyVerificationCodeRequestOneOf1 != nil {
		return json.Marshal(&src.VerifyVerificationCodeRequestOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApiInteractionPutRequestIdentifier) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApiInteractionPutRequestIdentifierOneOf != nil {
		return obj.ApiInteractionPutRequestIdentifierOneOf
	}

	if obj.ApiInteractionPutRequestIdentifierOneOf1 != nil {
		return obj.ApiInteractionPutRequestIdentifierOneOf1
	}

	if obj.ApiInteractionPutRequestIdentifierOneOf2 != nil {
		return obj.ApiInteractionPutRequestIdentifierOneOf2
	}

	if obj.ApiInteractionPutRequestIdentifierOneOf3 != nil {
		return obj.ApiInteractionPutRequestIdentifierOneOf3
	}

	if obj.ApiInteractionPutRequestIdentifierOneOf4 != nil {
		return obj.ApiInteractionPutRequestIdentifierOneOf4
	}

	if obj.ApiInteractionPutRequestIdentifierOneOf5 != nil {
		return obj.ApiInteractionPutRequestIdentifierOneOf5
	}

	if obj.VerifyVerificationCodeRequestOneOf != nil {
		return obj.VerifyVerificationCodeRequestOneOf
	}

	if obj.VerifyVerificationCodeRequestOneOf1 != nil {
		return obj.VerifyVerificationCodeRequestOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableApiInteractionPutRequestIdentifier struct {
	value *ApiInteractionPutRequestIdentifier
	isSet bool
}

func (v NullableApiInteractionPutRequestIdentifier) Get() *ApiInteractionPutRequestIdentifier {
	return v.value
}

func (v *NullableApiInteractionPutRequestIdentifier) Set(val *ApiInteractionPutRequestIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionPutRequestIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionPutRequestIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionPutRequestIdentifier(val *ApiInteractionPutRequestIdentifier) *NullableApiInteractionPutRequestIdentifier {
	return &NullableApiInteractionPutRequestIdentifier{value: val, isSet: true}
}

func (v NullableApiInteractionPutRequestIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionPutRequestIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


