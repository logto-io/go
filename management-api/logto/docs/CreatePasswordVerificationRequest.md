# CreatePasswordVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [**CreatePasswordVerificationRequestIdentifier**](CreatePasswordVerificationRequestIdentifier.md) |  | 
**Password** | **string** | The user password. | 

## Methods

### NewCreatePasswordVerificationRequest

`func NewCreatePasswordVerificationRequest(identifier CreatePasswordVerificationRequestIdentifier, password string, ) *CreatePasswordVerificationRequest`

NewCreatePasswordVerificationRequest instantiates a new CreatePasswordVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePasswordVerificationRequestWithDefaults

`func NewCreatePasswordVerificationRequestWithDefaults() *CreatePasswordVerificationRequest`

NewCreatePasswordVerificationRequestWithDefaults instantiates a new CreatePasswordVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CreatePasswordVerificationRequest) GetIdentifier() CreatePasswordVerificationRequestIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreatePasswordVerificationRequest) GetIdentifierOk() (*CreatePasswordVerificationRequestIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreatePasswordVerificationRequest) SetIdentifier(v CreatePasswordVerificationRequestIdentifier)`

SetIdentifier sets Identifier field to given value.


### GetPassword

`func (o *CreatePasswordVerificationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreatePasswordVerificationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreatePasswordVerificationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


