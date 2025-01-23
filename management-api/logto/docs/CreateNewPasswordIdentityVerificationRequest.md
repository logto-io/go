# CreateNewPasswordIdentityVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [**CreateNewPasswordIdentityVerificationRequestIdentifier**](CreateNewPasswordIdentityVerificationRequestIdentifier.md) |  | 
**Password** | **string** | The new user password. (A password digest will be created and stored securely in the verification record.) | 

## Methods

### NewCreateNewPasswordIdentityVerificationRequest

`func NewCreateNewPasswordIdentityVerificationRequest(identifier CreateNewPasswordIdentityVerificationRequestIdentifier, password string, ) *CreateNewPasswordIdentityVerificationRequest`

NewCreateNewPasswordIdentityVerificationRequest instantiates a new CreateNewPasswordIdentityVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNewPasswordIdentityVerificationRequestWithDefaults

`func NewCreateNewPasswordIdentityVerificationRequestWithDefaults() *CreateNewPasswordIdentityVerificationRequest`

NewCreateNewPasswordIdentityVerificationRequestWithDefaults instantiates a new CreateNewPasswordIdentityVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CreateNewPasswordIdentityVerificationRequest) GetIdentifier() CreateNewPasswordIdentityVerificationRequestIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateNewPasswordIdentityVerificationRequest) GetIdentifierOk() (*CreateNewPasswordIdentityVerificationRequestIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateNewPasswordIdentityVerificationRequest) SetIdentifier(v CreateNewPasswordIdentityVerificationRequestIdentifier)`

SetIdentifier sets Identifier field to given value.


### GetPassword

`func (o *CreateNewPasswordIdentityVerificationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateNewPasswordIdentityVerificationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateNewPasswordIdentityVerificationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


