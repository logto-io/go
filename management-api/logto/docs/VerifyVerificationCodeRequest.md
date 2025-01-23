# VerifyVerificationCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**VerificationCode** | **string** |  | 
**Phone** | **string** |  | 

## Methods

### NewVerifyVerificationCodeRequest

`func NewVerifyVerificationCodeRequest(email string, verificationCode string, phone string, ) *VerifyVerificationCodeRequest`

NewVerifyVerificationCodeRequest instantiates a new VerifyVerificationCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyVerificationCodeRequestWithDefaults

`func NewVerifyVerificationCodeRequestWithDefaults() *VerifyVerificationCodeRequest`

NewVerifyVerificationCodeRequestWithDefaults instantiates a new VerifyVerificationCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *VerifyVerificationCodeRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VerifyVerificationCodeRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VerifyVerificationCodeRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetVerificationCode

`func (o *VerifyVerificationCodeRequest) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *VerifyVerificationCodeRequest) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *VerifyVerificationCodeRequest) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.


### GetPhone

`func (o *VerifyVerificationCodeRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *VerifyVerificationCodeRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *VerifyVerificationCodeRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


