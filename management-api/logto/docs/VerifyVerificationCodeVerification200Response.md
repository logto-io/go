# VerifyVerificationCodeVerification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | **string** | he unique ID of the verification record. Required for user identification via the &#x60;Identification&#x60; API or to bind the identifier to the user&#39;s account via the &#x60;Profile&#x60; API. | 

## Methods

### NewVerifyVerificationCodeVerification200Response

`func NewVerifyVerificationCodeVerification200Response(verificationId string, ) *VerifyVerificationCodeVerification200Response`

NewVerifyVerificationCodeVerification200Response instantiates a new VerifyVerificationCodeVerification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyVerificationCodeVerification200ResponseWithDefaults

`func NewVerifyVerificationCodeVerification200ResponseWithDefaults() *VerifyVerificationCodeVerification200Response`

NewVerifyVerificationCodeVerification200ResponseWithDefaults instantiates a new VerifyVerificationCodeVerification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *VerifyVerificationCodeVerification200Response) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *VerifyVerificationCodeVerification200Response) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *VerifyVerificationCodeVerification200Response) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


