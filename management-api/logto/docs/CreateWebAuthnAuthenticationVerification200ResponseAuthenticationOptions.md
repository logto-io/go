# CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Challenge** | **string** |  | 
**Timeout** | Pointer to **float32** |  | [optional] 
**RpId** | Pointer to **string** |  | [optional] 
**AllowCredentials** | Pointer to [**[]CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner**](CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner.md) |  | [optional] 
**UserVerification** | Pointer to **string** |  | [optional] 
**Extensions** | Pointer to [**CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions**](CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions.md) |  | [optional] 

## Methods

### NewCreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions

`func NewCreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions(challenge string, ) *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions`

NewCreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions instantiates a new CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptionsWithDefaults

`func NewCreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptionsWithDefaults() *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions`

NewCreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptionsWithDefaults instantiates a new CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenge

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.


### GetTimeout

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetRpId

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) GetRpId() string`

GetRpId returns the RpId field if non-nil, zero value otherwise.

### GetRpIdOk

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) GetRpIdOk() (*string, bool)`

GetRpIdOk returns a tuple with the RpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpId

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) SetRpId(v string)`

SetRpId sets RpId field to given value.

### HasRpId

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) HasRpId() bool`

HasRpId returns a boolean if a field has been set.

### GetAllowCredentials

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) GetAllowCredentials() []CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner`

GetAllowCredentials returns the AllowCredentials field if non-nil, zero value otherwise.

### GetAllowCredentialsOk

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) GetAllowCredentialsOk() (*[]CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner, bool)`

GetAllowCredentialsOk returns a tuple with the AllowCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCredentials

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) SetAllowCredentials(v []CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner)`

SetAllowCredentials sets AllowCredentials field to given value.

### HasAllowCredentials

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) HasAllowCredentials() bool`

HasAllowCredentials returns a boolean if a field has been set.

### GetUserVerification

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) GetUserVerification() string`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) GetUserVerificationOk() (*string, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) SetUserVerification(v string)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.

### GetExtensions

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) GetExtensions() CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) GetExtensionsOk() (*CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) SetExtensions(v CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *CreateWebAuthnAuthenticationVerification200ResponseAuthenticationOptions) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


