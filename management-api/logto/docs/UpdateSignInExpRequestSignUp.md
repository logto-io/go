# UpdateSignInExpRequestSignUp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifiers** | **[]string** | Specify allowed identifiers when signing-up. | 
**Password** | **bool** | Whether the user is required to set a password when signing-up. | 
**Verify** | **bool** | Whether the user is required to verify their email/phone when signing-up. | 

## Methods

### NewUpdateSignInExpRequestSignUp

`func NewUpdateSignInExpRequestSignUp(identifiers []string, password bool, verify bool, ) *UpdateSignInExpRequestSignUp`

NewUpdateSignInExpRequestSignUp instantiates a new UpdateSignInExpRequestSignUp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSignInExpRequestSignUpWithDefaults

`func NewUpdateSignInExpRequestSignUpWithDefaults() *UpdateSignInExpRequestSignUp`

NewUpdateSignInExpRequestSignUpWithDefaults instantiates a new UpdateSignInExpRequestSignUp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifiers

`func (o *UpdateSignInExpRequestSignUp) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *UpdateSignInExpRequestSignUp) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *UpdateSignInExpRequestSignUp) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.


### GetPassword

`func (o *UpdateSignInExpRequestSignUp) GetPassword() bool`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateSignInExpRequestSignUp) GetPasswordOk() (*bool, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateSignInExpRequestSignUp) SetPassword(v bool)`

SetPassword sets Password field to given value.


### GetVerify

`func (o *UpdateSignInExpRequestSignUp) GetVerify() bool`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *UpdateSignInExpRequestSignUp) GetVerifyOk() (*bool, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *UpdateSignInExpRequestSignUp) SetVerify(v bool)`

SetVerify sets Verify field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


