# CreateNewPasswordIdentityVerification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | **string** | The unique verification ID of the newly created NewPasswordIdentity verification record. The &#x60;verificationId&#x60; is required when creating a new user account via the &#x60;Identification&#x60; API. | 

## Methods

### NewCreateNewPasswordIdentityVerification200Response

`func NewCreateNewPasswordIdentityVerification200Response(verificationId string, ) *CreateNewPasswordIdentityVerification200Response`

NewCreateNewPasswordIdentityVerification200Response instantiates a new CreateNewPasswordIdentityVerification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNewPasswordIdentityVerification200ResponseWithDefaults

`func NewCreateNewPasswordIdentityVerification200ResponseWithDefaults() *CreateNewPasswordIdentityVerification200Response`

NewCreateNewPasswordIdentityVerification200ResponseWithDefaults instantiates a new CreateNewPasswordIdentityVerification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *CreateNewPasswordIdentityVerification200Response) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *CreateNewPasswordIdentityVerification200Response) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *CreateNewPasswordIdentityVerification200Response) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


