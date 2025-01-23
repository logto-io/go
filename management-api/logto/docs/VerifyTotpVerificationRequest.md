# VerifyTotpVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The TOTP code to be verified. | 
**VerificationId** | Pointer to **string** | The verification ID of the newly created TOTP secret. This ID is required to verify a newly created TOTP secret that needs to be bound to the user account. If not provided, the API will create a new TOTP verification record and verify the code against the user&#39;s existing TOTP secret. | [optional] 

## Methods

### NewVerifyTotpVerificationRequest

`func NewVerifyTotpVerificationRequest(code string, ) *VerifyTotpVerificationRequest`

NewVerifyTotpVerificationRequest instantiates a new VerifyTotpVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyTotpVerificationRequestWithDefaults

`func NewVerifyTotpVerificationRequestWithDefaults() *VerifyTotpVerificationRequest`

NewVerifyTotpVerificationRequestWithDefaults instantiates a new VerifyTotpVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *VerifyTotpVerificationRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VerifyTotpVerificationRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VerifyTotpVerificationRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetVerificationId

`func (o *VerifyTotpVerificationRequest) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *VerifyTotpVerificationRequest) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *VerifyTotpVerificationRequest) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.

### HasVerificationId

`func (o *VerifyTotpVerificationRequest) HasVerificationId() bool`

HasVerificationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


