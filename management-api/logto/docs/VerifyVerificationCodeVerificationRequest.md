# VerifyVerificationCodeVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [**VerifyVerificationCodeVerificationRequestIdentifier**](VerifyVerificationCodeVerificationRequestIdentifier.md) |  | 
**VerificationId** | **string** | The verification ID of the CodeVerification record. | 
**Code** | **string** | The verification code to be verified. | 

## Methods

### NewVerifyVerificationCodeVerificationRequest

`func NewVerifyVerificationCodeVerificationRequest(identifier VerifyVerificationCodeVerificationRequestIdentifier, verificationId string, code string, ) *VerifyVerificationCodeVerificationRequest`

NewVerifyVerificationCodeVerificationRequest instantiates a new VerifyVerificationCodeVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyVerificationCodeVerificationRequestWithDefaults

`func NewVerifyVerificationCodeVerificationRequestWithDefaults() *VerifyVerificationCodeVerificationRequest`

NewVerifyVerificationCodeVerificationRequestWithDefaults instantiates a new VerifyVerificationCodeVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *VerifyVerificationCodeVerificationRequest) GetIdentifier() VerifyVerificationCodeVerificationRequestIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VerifyVerificationCodeVerificationRequest) GetIdentifierOk() (*VerifyVerificationCodeVerificationRequestIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VerifyVerificationCodeVerificationRequest) SetIdentifier(v VerifyVerificationCodeVerificationRequestIdentifier)`

SetIdentifier sets Identifier field to given value.


### GetVerificationId

`func (o *VerifyVerificationCodeVerificationRequest) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *VerifyVerificationCodeVerificationRequest) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *VerifyVerificationCodeVerificationRequest) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.


### GetCode

`func (o *VerifyVerificationCodeVerificationRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VerifyVerificationCodeVerificationRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VerifyVerificationCodeVerificationRequest) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


