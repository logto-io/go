# CreateUserMfaVerification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Secret** | **string** |  | 
**SecretQrCode** | **string** |  | 
**Codes** | **[]string** |  | 

## Methods

### NewCreateUserMfaVerification200Response

`func NewCreateUserMfaVerification200Response(type_ string, secret string, secretQrCode string, codes []string, ) *CreateUserMfaVerification200Response`

NewCreateUserMfaVerification200Response instantiates a new CreateUserMfaVerification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserMfaVerification200ResponseWithDefaults

`func NewCreateUserMfaVerification200ResponseWithDefaults() *CreateUserMfaVerification200Response`

NewCreateUserMfaVerification200ResponseWithDefaults instantiates a new CreateUserMfaVerification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateUserMfaVerification200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateUserMfaVerification200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateUserMfaVerification200Response) SetType(v string)`

SetType sets Type field to given value.


### GetSecret

`func (o *CreateUserMfaVerification200Response) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateUserMfaVerification200Response) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateUserMfaVerification200Response) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetSecretQrCode

`func (o *CreateUserMfaVerification200Response) GetSecretQrCode() string`

GetSecretQrCode returns the SecretQrCode field if non-nil, zero value otherwise.

### GetSecretQrCodeOk

`func (o *CreateUserMfaVerification200Response) GetSecretQrCodeOk() (*string, bool)`

GetSecretQrCodeOk returns a tuple with the SecretQrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretQrCode

`func (o *CreateUserMfaVerification200Response) SetSecretQrCode(v string)`

SetSecretQrCode sets SecretQrCode field to given value.


### GetCodes

`func (o *CreateUserMfaVerification200Response) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *CreateUserMfaVerification200Response) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *CreateUserMfaVerification200Response) SetCodes(v []string)`

SetCodes sets Codes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


