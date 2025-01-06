# CreatePasswordVerification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | **string** | The unique verification ID of the newly created Password verification record. The &#x60;verificationId&#x60; is required when verifying the user&#39;s identity via the &#x60;Identification&#x60; API. | 

## Methods

### NewCreatePasswordVerification200Response

`func NewCreatePasswordVerification200Response(verificationId string, ) *CreatePasswordVerification200Response`

NewCreatePasswordVerification200Response instantiates a new CreatePasswordVerification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePasswordVerification200ResponseWithDefaults

`func NewCreatePasswordVerification200ResponseWithDefaults() *CreatePasswordVerification200Response`

NewCreatePasswordVerification200ResponseWithDefaults instantiates a new CreatePasswordVerification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *CreatePasswordVerification200Response) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *CreatePasswordVerification200Response) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *CreatePasswordVerification200Response) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


