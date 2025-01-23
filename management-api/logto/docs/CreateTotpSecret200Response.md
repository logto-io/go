# CreateTotpSecret200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | **string** | The unique verification ID for the TOTP record. This ID is required to verify the TOTP code. | 
**Secret** | **string** | The newly generated TOTP secret. | 
**SecretQrCode** | **string** | A QR code image data URL for the TOTP secret. The user can scan this QR code with their TOTP authenticator app. | 

## Methods

### NewCreateTotpSecret200Response

`func NewCreateTotpSecret200Response(verificationId string, secret string, secretQrCode string, ) *CreateTotpSecret200Response`

NewCreateTotpSecret200Response instantiates a new CreateTotpSecret200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTotpSecret200ResponseWithDefaults

`func NewCreateTotpSecret200ResponseWithDefaults() *CreateTotpSecret200Response`

NewCreateTotpSecret200ResponseWithDefaults instantiates a new CreateTotpSecret200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *CreateTotpSecret200Response) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *CreateTotpSecret200Response) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *CreateTotpSecret200Response) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.


### GetSecret

`func (o *CreateTotpSecret200Response) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateTotpSecret200Response) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateTotpSecret200Response) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetSecretQrCode

`func (o *CreateTotpSecret200Response) GetSecretQrCode() string`

GetSecretQrCode returns the SecretQrCode field if non-nil, zero value otherwise.

### GetSecretQrCodeOk

`func (o *CreateTotpSecret200Response) GetSecretQrCodeOk() (*string, bool)`

GetSecretQrCodeOk returns a tuple with the SecretQrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretQrCode

`func (o *CreateTotpSecret200Response) SetSecretQrCode(v string)`

SetSecretQrCode sets SecretQrCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


