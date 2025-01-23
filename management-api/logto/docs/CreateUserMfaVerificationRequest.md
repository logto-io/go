# CreateUserMfaVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of MFA verification to create. | 
**Secret** | Pointer to **string** | The secret for the MFA verification, if not provided, a new secret will be generated. | [optional] 
**Codes** | Pointer to **[]string** | The backup codes for the MFA verification, if not provided, a new group of backup codes will be generated. | [optional] 

## Methods

### NewCreateUserMfaVerificationRequest

`func NewCreateUserMfaVerificationRequest(type_ string, ) *CreateUserMfaVerificationRequest`

NewCreateUserMfaVerificationRequest instantiates a new CreateUserMfaVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserMfaVerificationRequestWithDefaults

`func NewCreateUserMfaVerificationRequestWithDefaults() *CreateUserMfaVerificationRequest`

NewCreateUserMfaVerificationRequestWithDefaults instantiates a new CreateUserMfaVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateUserMfaVerificationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateUserMfaVerificationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateUserMfaVerificationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetSecret

`func (o *CreateUserMfaVerificationRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateUserMfaVerificationRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateUserMfaVerificationRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CreateUserMfaVerificationRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetCodes

`func (o *CreateUserMfaVerificationRequest) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *CreateUserMfaVerificationRequest) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *CreateUserMfaVerificationRequest) SetCodes(v []string)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *CreateUserMfaVerificationRequest) HasCodes() bool`

HasCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


