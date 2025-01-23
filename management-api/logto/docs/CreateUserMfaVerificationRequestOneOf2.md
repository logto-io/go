# CreateUserMfaVerificationRequestOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of MFA verification to create. | 
**Secret** | Pointer to **string** | The secret for the MFA verification, if not provided, a new secret will be generated. | [optional] 

## Methods

### NewCreateUserMfaVerificationRequestOneOf2

`func NewCreateUserMfaVerificationRequestOneOf2(type_ string, ) *CreateUserMfaVerificationRequestOneOf2`

NewCreateUserMfaVerificationRequestOneOf2 instantiates a new CreateUserMfaVerificationRequestOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserMfaVerificationRequestOneOf2WithDefaults

`func NewCreateUserMfaVerificationRequestOneOf2WithDefaults() *CreateUserMfaVerificationRequestOneOf2`

NewCreateUserMfaVerificationRequestOneOf2WithDefaults instantiates a new CreateUserMfaVerificationRequestOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateUserMfaVerificationRequestOneOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateUserMfaVerificationRequestOneOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateUserMfaVerificationRequestOneOf2) SetType(v string)`

SetType sets Type field to given value.


### GetSecret

`func (o *CreateUserMfaVerificationRequestOneOf2) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateUserMfaVerificationRequestOneOf2) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateUserMfaVerificationRequestOneOf2) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CreateUserMfaVerificationRequestOneOf2) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


