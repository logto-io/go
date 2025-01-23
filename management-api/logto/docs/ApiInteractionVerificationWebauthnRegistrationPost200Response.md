# ApiInteractionVerificationWebauthnRegistrationPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rp** | [**CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsRp**](CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsRp.md) |  | 
**User** | [**CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsUser**](CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsUser.md) |  | 
**Challenge** | **string** |  | 
**PubKeyCredParams** | [**[]CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsPubKeyCredParamsInner**](CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsPubKeyCredParamsInner.md) |  | 
**Timeout** | Pointer to **float32** |  | [optional] 
**ExcludeCredentials** | Pointer to [**[]CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner**](CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner.md) |  | [optional] 
**AuthenticatorSelection** | Pointer to [**CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection**](CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection.md) |  | [optional] 
**Attestation** | Pointer to **string** |  | [optional] 
**Extensions** | Pointer to [**CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions**](CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions.md) |  | [optional] 

## Methods

### NewApiInteractionVerificationWebauthnRegistrationPost200Response

`func NewApiInteractionVerificationWebauthnRegistrationPost200Response(rp CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsRp, user CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsUser, challenge string, pubKeyCredParams []CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsPubKeyCredParamsInner, ) *ApiInteractionVerificationWebauthnRegistrationPost200Response`

NewApiInteractionVerificationWebauthnRegistrationPost200Response instantiates a new ApiInteractionVerificationWebauthnRegistrationPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInteractionVerificationWebauthnRegistrationPost200ResponseWithDefaults

`func NewApiInteractionVerificationWebauthnRegistrationPost200ResponseWithDefaults() *ApiInteractionVerificationWebauthnRegistrationPost200Response`

NewApiInteractionVerificationWebauthnRegistrationPost200ResponseWithDefaults instantiates a new ApiInteractionVerificationWebauthnRegistrationPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRp

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetRp() CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsRp`

GetRp returns the Rp field if non-nil, zero value otherwise.

### GetRpOk

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetRpOk() (*CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsRp, bool)`

GetRpOk returns a tuple with the Rp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRp

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) SetRp(v CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsRp)`

SetRp sets Rp field to given value.


### GetUser

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetUser() CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetUserOk() (*CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) SetUser(v CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsUser)`

SetUser sets User field to given value.


### GetChallenge

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.


### GetPubKeyCredParams

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetPubKeyCredParams() []CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsPubKeyCredParamsInner`

GetPubKeyCredParams returns the PubKeyCredParams field if non-nil, zero value otherwise.

### GetPubKeyCredParamsOk

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetPubKeyCredParamsOk() (*[]CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsPubKeyCredParamsInner, bool)`

GetPubKeyCredParamsOk returns a tuple with the PubKeyCredParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKeyCredParams

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) SetPubKeyCredParams(v []CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsPubKeyCredParamsInner)`

SetPubKeyCredParams sets PubKeyCredParams field to given value.


### GetTimeout

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetExcludeCredentials

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetExcludeCredentials() []CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner`

GetExcludeCredentials returns the ExcludeCredentials field if non-nil, zero value otherwise.

### GetExcludeCredentialsOk

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetExcludeCredentialsOk() (*[]CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner, bool)`

GetExcludeCredentialsOk returns a tuple with the ExcludeCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCredentials

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) SetExcludeCredentials(v []CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExcludeCredentialsInner)`

SetExcludeCredentials sets ExcludeCredentials field to given value.

### HasExcludeCredentials

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) HasExcludeCredentials() bool`

HasExcludeCredentials returns a boolean if a field has been set.

### GetAuthenticatorSelection

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetAuthenticatorSelection() CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection`

GetAuthenticatorSelection returns the AuthenticatorSelection field if non-nil, zero value otherwise.

### GetAuthenticatorSelectionOk

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetAuthenticatorSelectionOk() (*CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection, bool)`

GetAuthenticatorSelectionOk returns a tuple with the AuthenticatorSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorSelection

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) SetAuthenticatorSelection(v CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsAuthenticatorSelection)`

SetAuthenticatorSelection sets AuthenticatorSelection field to given value.

### HasAuthenticatorSelection

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) HasAuthenticatorSelection() bool`

HasAuthenticatorSelection returns a boolean if a field has been set.

### GetAttestation

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetAttestation() string`

GetAttestation returns the Attestation field if non-nil, zero value otherwise.

### GetAttestationOk

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetAttestationOk() (*string, bool)`

GetAttestationOk returns a tuple with the Attestation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestation

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) SetAttestation(v string)`

SetAttestation sets Attestation field to given value.

### HasAttestation

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) HasAttestation() bool`

HasAttestation returns a boolean if a field has been set.

### GetExtensions

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetExtensions() CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) GetExtensionsOk() (*CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) SetExtensions(v CreateWebAuthnRegistrationVerification200ResponseRegistrationOptionsExtensions)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ApiInteractionVerificationWebauthnRegistrationPost200Response) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


