# VerifyTotpVerification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | **string** | The unique verification ID of the TOTP verification record. For newly created TOTP secret verification record, this ID is required to bind the TOTP secret to the user account through &#x60;Profile&#x60; API. | 

## Methods

### NewVerifyTotpVerification200Response

`func NewVerifyTotpVerification200Response(verificationId string, ) *VerifyTotpVerification200Response`

NewVerifyTotpVerification200Response instantiates a new VerifyTotpVerification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyTotpVerification200ResponseWithDefaults

`func NewVerifyTotpVerification200ResponseWithDefaults() *VerifyTotpVerification200Response`

NewVerifyTotpVerification200ResponseWithDefaults instantiates a new VerifyTotpVerification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *VerifyTotpVerification200Response) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *VerifyTotpVerification200Response) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *VerifyTotpVerification200Response) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


