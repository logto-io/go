# CreateUserMfaVerificationRequestOneOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of MFA verification to create. | 
**Codes** | Pointer to **[]string** | The backup codes for the MFA verification, if not provided, a new group of backup codes will be generated. | [optional] 

## Methods

### NewCreateUserMfaVerificationRequestOneOf3

`func NewCreateUserMfaVerificationRequestOneOf3(type_ string, ) *CreateUserMfaVerificationRequestOneOf3`

NewCreateUserMfaVerificationRequestOneOf3 instantiates a new CreateUserMfaVerificationRequestOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserMfaVerificationRequestOneOf3WithDefaults

`func NewCreateUserMfaVerificationRequestOneOf3WithDefaults() *CreateUserMfaVerificationRequestOneOf3`

NewCreateUserMfaVerificationRequestOneOf3WithDefaults instantiates a new CreateUserMfaVerificationRequestOneOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateUserMfaVerificationRequestOneOf3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateUserMfaVerificationRequestOneOf3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateUserMfaVerificationRequestOneOf3) SetType(v string)`

SetType sets Type field to given value.


### GetCodes

`func (o *CreateUserMfaVerificationRequestOneOf3) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *CreateUserMfaVerificationRequestOneOf3) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *CreateUserMfaVerificationRequestOneOf3) SetCodes(v []string)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *CreateUserMfaVerificationRequestOneOf3) HasCodes() bool`

HasCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


