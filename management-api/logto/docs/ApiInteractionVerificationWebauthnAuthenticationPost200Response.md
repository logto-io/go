# ApiInteractionVerificationWebauthnAuthenticationPost200Response

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

### NewApiInteractionVerificationWebauthnAuthenticationPost200Response

`func NewApiInteractionVerificationWebauthnAuthenticationPost200Response(challenge string, ) *ApiInteractionVerificationWebauthnAuthenticationPost200Response`

NewApiInteractionVerificationWebauthnAuthenticationPost200Response instantiates a new ApiInteractionVerificationWebauthnAuthenticationPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInteractionVerificationWebauthnAuthenticationPost200ResponseWithDefaults

`func NewApiInteractionVerificationWebauthnAuthenticationPost200ResponseWithDefaults() *ApiInteractionVerificationWebauthnAuthenticationPost200Response`

NewApiInteractionVerificationWebauthnAuthenticationPost200ResponseWithDefaults instantiates a new ApiInteractionVerificationWebauthnAuthenticationPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenge

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.


### GetTimeout

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetRpId

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) GetRpId() string`

GetRpId returns the RpId field if non-nil, zero value otherwise.

### GetRpIdOk

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) GetRpIdOk() (*string, bool)`

GetRpIdOk returns a tuple with the RpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpId

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) SetRpId(v string)`

SetRpId sets RpId field to given value.

### HasRpId

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) HasRpId() bool`

HasRpId returns a boolean if a field has been set.

### GetAllowCredentials

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) GetAllowCredentials() []CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner`

GetAllowCredentials returns the AllowCredentials field if non-nil, zero value otherwise.

### GetAllowCredentialsOk

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) GetAllowCredentialsOk() (*[]CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner, bool)`

GetAllowCredentialsOk returns a tuple with the AllowCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCredentials

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) SetAllowCredentials(v []CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner)`

SetAllowCredentials sets AllowCredentials field to given value.

### HasAllowCredentials

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) HasAllowCredentials() bool`

HasAllowCredentials returns a boolean if a field has been set.

### GetUserVerification

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) GetUserVerification() string`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) GetUserVerificationOk() (*string, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) SetUserVerification(v string)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.

### GetExtensions

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) GetExtensions() CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) GetExtensionsOk() (*CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) SetExtensions(v CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ApiInteractionVerificationWebauthnAuthenticationPost200Response) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


