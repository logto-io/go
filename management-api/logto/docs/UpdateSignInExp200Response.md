# UpdateSignInExp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**Color** | [**UpdateSignInExp200ResponseColor**](UpdateSignInExp200ResponseColor.md) |  | 
**Branding** | [**ListApplicationOrganizations200ResponseInnerBranding**](ListApplicationOrganizations200ResponseInnerBranding.md) |  | 
**LanguageInfo** | [**UpdateSignInExp200ResponseLanguageInfo**](UpdateSignInExp200ResponseLanguageInfo.md) |  | 
**TermsOfUseUrl** | **NullableString** |  | 
**PrivacyPolicyUrl** | **NullableString** |  | 
**AgreeToTermsPolicy** | **string** |  | 
**SignIn** | [**UpdateSignInExp200ResponseSignIn**](UpdateSignInExp200ResponseSignIn.md) |  | 
**SignUp** | [**UpdateSignInExp200ResponseSignUp**](UpdateSignInExp200ResponseSignUp.md) |  | 
**SocialSignIn** | [**GetSignInExp200ResponseSocialSignIn**](GetSignInExp200ResponseSocialSignIn.md) |  | 
**SocialSignInConnectorTargets** | **[]string** |  | 
**SignInMode** | **string** |  | 
**CustomCss** | **NullableString** |  | 
**CustomContent** | **map[string]string** |  | 
**CustomUiAssets** | [**NullableGetSignInExp200ResponseCustomUiAssets**](GetSignInExp200ResponseCustomUiAssets.md) |  | 
**PasswordPolicy** | [**UpdateSignInExp200ResponsePasswordPolicy**](UpdateSignInExp200ResponsePasswordPolicy.md) |  | 
**Mfa** | [**UpdateSignInExp200ResponseMfa**](UpdateSignInExp200ResponseMfa.md) |  | 
**SingleSignOnEnabled** | **bool** |  | 
**SupportEmail** | **NullableString** |  | 
**SupportWebsiteUrl** | **NullableString** |  | 
**UnknownSessionRedirectUrl** | **NullableString** |  | 

## Methods

### NewUpdateSignInExp200Response

`func NewUpdateSignInExp200Response(tenantId string, id string, color UpdateSignInExp200ResponseColor, branding ListApplicationOrganizations200ResponseInnerBranding, languageInfo UpdateSignInExp200ResponseLanguageInfo, termsOfUseUrl NullableString, privacyPolicyUrl NullableString, agreeToTermsPolicy string, signIn UpdateSignInExp200ResponseSignIn, signUp UpdateSignInExp200ResponseSignUp, socialSignIn GetSignInExp200ResponseSocialSignIn, socialSignInConnectorTargets []string, signInMode string, customCss NullableString, customContent map[string]string, customUiAssets NullableGetSignInExp200ResponseCustomUiAssets, passwordPolicy UpdateSignInExp200ResponsePasswordPolicy, mfa UpdateSignInExp200ResponseMfa, singleSignOnEnabled bool, supportEmail NullableString, supportWebsiteUrl NullableString, unknownSessionRedirectUrl NullableString, ) *UpdateSignInExp200Response`

NewUpdateSignInExp200Response instantiates a new UpdateSignInExp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSignInExp200ResponseWithDefaults

`func NewUpdateSignInExp200ResponseWithDefaults() *UpdateSignInExp200Response`

NewUpdateSignInExp200ResponseWithDefaults instantiates a new UpdateSignInExp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *UpdateSignInExp200Response) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateSignInExp200Response) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateSignInExp200Response) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *UpdateSignInExp200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSignInExp200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSignInExp200Response) SetId(v string)`

SetId sets Id field to given value.


### GetColor

`func (o *UpdateSignInExp200Response) GetColor() UpdateSignInExp200ResponseColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UpdateSignInExp200Response) GetColorOk() (*UpdateSignInExp200ResponseColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UpdateSignInExp200Response) SetColor(v UpdateSignInExp200ResponseColor)`

SetColor sets Color field to given value.


### GetBranding

`func (o *UpdateSignInExp200Response) GetBranding() ListApplicationOrganizations200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *UpdateSignInExp200Response) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *UpdateSignInExp200Response) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding)`

SetBranding sets Branding field to given value.


### GetLanguageInfo

`func (o *UpdateSignInExp200Response) GetLanguageInfo() UpdateSignInExp200ResponseLanguageInfo`

GetLanguageInfo returns the LanguageInfo field if non-nil, zero value otherwise.

### GetLanguageInfoOk

`func (o *UpdateSignInExp200Response) GetLanguageInfoOk() (*UpdateSignInExp200ResponseLanguageInfo, bool)`

GetLanguageInfoOk returns a tuple with the LanguageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageInfo

`func (o *UpdateSignInExp200Response) SetLanguageInfo(v UpdateSignInExp200ResponseLanguageInfo)`

SetLanguageInfo sets LanguageInfo field to given value.


### GetTermsOfUseUrl

`func (o *UpdateSignInExp200Response) GetTermsOfUseUrl() string`

GetTermsOfUseUrl returns the TermsOfUseUrl field if non-nil, zero value otherwise.

### GetTermsOfUseUrlOk

`func (o *UpdateSignInExp200Response) GetTermsOfUseUrlOk() (*string, bool)`

GetTermsOfUseUrlOk returns a tuple with the TermsOfUseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUseUrl

`func (o *UpdateSignInExp200Response) SetTermsOfUseUrl(v string)`

SetTermsOfUseUrl sets TermsOfUseUrl field to given value.


### SetTermsOfUseUrlNil

`func (o *UpdateSignInExp200Response) SetTermsOfUseUrlNil(b bool)`

 SetTermsOfUseUrlNil sets the value for TermsOfUseUrl to be an explicit nil

### UnsetTermsOfUseUrl
`func (o *UpdateSignInExp200Response) UnsetTermsOfUseUrl()`

UnsetTermsOfUseUrl ensures that no value is present for TermsOfUseUrl, not even an explicit nil
### GetPrivacyPolicyUrl

`func (o *UpdateSignInExp200Response) GetPrivacyPolicyUrl() string`

GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetPrivacyPolicyUrlOk

`func (o *UpdateSignInExp200Response) GetPrivacyPolicyUrlOk() (*string, bool)`

GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyUrl

`func (o *UpdateSignInExp200Response) SetPrivacyPolicyUrl(v string)`

SetPrivacyPolicyUrl sets PrivacyPolicyUrl field to given value.


### SetPrivacyPolicyUrlNil

`func (o *UpdateSignInExp200Response) SetPrivacyPolicyUrlNil(b bool)`

 SetPrivacyPolicyUrlNil sets the value for PrivacyPolicyUrl to be an explicit nil

### UnsetPrivacyPolicyUrl
`func (o *UpdateSignInExp200Response) UnsetPrivacyPolicyUrl()`

UnsetPrivacyPolicyUrl ensures that no value is present for PrivacyPolicyUrl, not even an explicit nil
### GetAgreeToTermsPolicy

`func (o *UpdateSignInExp200Response) GetAgreeToTermsPolicy() string`

GetAgreeToTermsPolicy returns the AgreeToTermsPolicy field if non-nil, zero value otherwise.

### GetAgreeToTermsPolicyOk

`func (o *UpdateSignInExp200Response) GetAgreeToTermsPolicyOk() (*string, bool)`

GetAgreeToTermsPolicyOk returns a tuple with the AgreeToTermsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreeToTermsPolicy

`func (o *UpdateSignInExp200Response) SetAgreeToTermsPolicy(v string)`

SetAgreeToTermsPolicy sets AgreeToTermsPolicy field to given value.


### GetSignIn

`func (o *UpdateSignInExp200Response) GetSignIn() UpdateSignInExp200ResponseSignIn`

GetSignIn returns the SignIn field if non-nil, zero value otherwise.

### GetSignInOk

`func (o *UpdateSignInExp200Response) GetSignInOk() (*UpdateSignInExp200ResponseSignIn, bool)`

GetSignInOk returns a tuple with the SignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignIn

`func (o *UpdateSignInExp200Response) SetSignIn(v UpdateSignInExp200ResponseSignIn)`

SetSignIn sets SignIn field to given value.


### GetSignUp

`func (o *UpdateSignInExp200Response) GetSignUp() UpdateSignInExp200ResponseSignUp`

GetSignUp returns the SignUp field if non-nil, zero value otherwise.

### GetSignUpOk

`func (o *UpdateSignInExp200Response) GetSignUpOk() (*UpdateSignInExp200ResponseSignUp, bool)`

GetSignUpOk returns a tuple with the SignUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUp

`func (o *UpdateSignInExp200Response) SetSignUp(v UpdateSignInExp200ResponseSignUp)`

SetSignUp sets SignUp field to given value.


### GetSocialSignIn

`func (o *UpdateSignInExp200Response) GetSocialSignIn() GetSignInExp200ResponseSocialSignIn`

GetSocialSignIn returns the SocialSignIn field if non-nil, zero value otherwise.

### GetSocialSignInOk

`func (o *UpdateSignInExp200Response) GetSocialSignInOk() (*GetSignInExp200ResponseSocialSignIn, bool)`

GetSocialSignInOk returns a tuple with the SocialSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSignIn

`func (o *UpdateSignInExp200Response) SetSocialSignIn(v GetSignInExp200ResponseSocialSignIn)`

SetSocialSignIn sets SocialSignIn field to given value.


### GetSocialSignInConnectorTargets

`func (o *UpdateSignInExp200Response) GetSocialSignInConnectorTargets() []string`

GetSocialSignInConnectorTargets returns the SocialSignInConnectorTargets field if non-nil, zero value otherwise.

### GetSocialSignInConnectorTargetsOk

`func (o *UpdateSignInExp200Response) GetSocialSignInConnectorTargetsOk() (*[]string, bool)`

GetSocialSignInConnectorTargetsOk returns a tuple with the SocialSignInConnectorTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSignInConnectorTargets

`func (o *UpdateSignInExp200Response) SetSocialSignInConnectorTargets(v []string)`

SetSocialSignInConnectorTargets sets SocialSignInConnectorTargets field to given value.


### GetSignInMode

`func (o *UpdateSignInExp200Response) GetSignInMode() string`

GetSignInMode returns the SignInMode field if non-nil, zero value otherwise.

### GetSignInModeOk

`func (o *UpdateSignInExp200Response) GetSignInModeOk() (*string, bool)`

GetSignInModeOk returns a tuple with the SignInMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInMode

`func (o *UpdateSignInExp200Response) SetSignInMode(v string)`

SetSignInMode sets SignInMode field to given value.


### GetCustomCss

`func (o *UpdateSignInExp200Response) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *UpdateSignInExp200Response) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *UpdateSignInExp200Response) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.


### SetCustomCssNil

`func (o *UpdateSignInExp200Response) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *UpdateSignInExp200Response) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetCustomContent

`func (o *UpdateSignInExp200Response) GetCustomContent() map[string]string`

GetCustomContent returns the CustomContent field if non-nil, zero value otherwise.

### GetCustomContentOk

`func (o *UpdateSignInExp200Response) GetCustomContentOk() (*map[string]string, bool)`

GetCustomContentOk returns a tuple with the CustomContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContent

`func (o *UpdateSignInExp200Response) SetCustomContent(v map[string]string)`

SetCustomContent sets CustomContent field to given value.


### GetCustomUiAssets

`func (o *UpdateSignInExp200Response) GetCustomUiAssets() GetSignInExp200ResponseCustomUiAssets`

GetCustomUiAssets returns the CustomUiAssets field if non-nil, zero value otherwise.

### GetCustomUiAssetsOk

`func (o *UpdateSignInExp200Response) GetCustomUiAssetsOk() (*GetSignInExp200ResponseCustomUiAssets, bool)`

GetCustomUiAssetsOk returns a tuple with the CustomUiAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUiAssets

`func (o *UpdateSignInExp200Response) SetCustomUiAssets(v GetSignInExp200ResponseCustomUiAssets)`

SetCustomUiAssets sets CustomUiAssets field to given value.


### SetCustomUiAssetsNil

`func (o *UpdateSignInExp200Response) SetCustomUiAssetsNil(b bool)`

 SetCustomUiAssetsNil sets the value for CustomUiAssets to be an explicit nil

### UnsetCustomUiAssets
`func (o *UpdateSignInExp200Response) UnsetCustomUiAssets()`

UnsetCustomUiAssets ensures that no value is present for CustomUiAssets, not even an explicit nil
### GetPasswordPolicy

`func (o *UpdateSignInExp200Response) GetPasswordPolicy() UpdateSignInExp200ResponsePasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *UpdateSignInExp200Response) GetPasswordPolicyOk() (*UpdateSignInExp200ResponsePasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *UpdateSignInExp200Response) SetPasswordPolicy(v UpdateSignInExp200ResponsePasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.


### GetMfa

`func (o *UpdateSignInExp200Response) GetMfa() UpdateSignInExp200ResponseMfa`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *UpdateSignInExp200Response) GetMfaOk() (*UpdateSignInExp200ResponseMfa, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *UpdateSignInExp200Response) SetMfa(v UpdateSignInExp200ResponseMfa)`

SetMfa sets Mfa field to given value.


### GetSingleSignOnEnabled

`func (o *UpdateSignInExp200Response) GetSingleSignOnEnabled() bool`

GetSingleSignOnEnabled returns the SingleSignOnEnabled field if non-nil, zero value otherwise.

### GetSingleSignOnEnabledOk

`func (o *UpdateSignInExp200Response) GetSingleSignOnEnabledOk() (*bool, bool)`

GetSingleSignOnEnabledOk returns a tuple with the SingleSignOnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSignOnEnabled

`func (o *UpdateSignInExp200Response) SetSingleSignOnEnabled(v bool)`

SetSingleSignOnEnabled sets SingleSignOnEnabled field to given value.


### GetSupportEmail

`func (o *UpdateSignInExp200Response) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *UpdateSignInExp200Response) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *UpdateSignInExp200Response) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.


### SetSupportEmailNil

`func (o *UpdateSignInExp200Response) SetSupportEmailNil(b bool)`

 SetSupportEmailNil sets the value for SupportEmail to be an explicit nil

### UnsetSupportEmail
`func (o *UpdateSignInExp200Response) UnsetSupportEmail()`

UnsetSupportEmail ensures that no value is present for SupportEmail, not even an explicit nil
### GetSupportWebsiteUrl

`func (o *UpdateSignInExp200Response) GetSupportWebsiteUrl() string`

GetSupportWebsiteUrl returns the SupportWebsiteUrl field if non-nil, zero value otherwise.

### GetSupportWebsiteUrlOk

`func (o *UpdateSignInExp200Response) GetSupportWebsiteUrlOk() (*string, bool)`

GetSupportWebsiteUrlOk returns a tuple with the SupportWebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportWebsiteUrl

`func (o *UpdateSignInExp200Response) SetSupportWebsiteUrl(v string)`

SetSupportWebsiteUrl sets SupportWebsiteUrl field to given value.


### SetSupportWebsiteUrlNil

`func (o *UpdateSignInExp200Response) SetSupportWebsiteUrlNil(b bool)`

 SetSupportWebsiteUrlNil sets the value for SupportWebsiteUrl to be an explicit nil

### UnsetSupportWebsiteUrl
`func (o *UpdateSignInExp200Response) UnsetSupportWebsiteUrl()`

UnsetSupportWebsiteUrl ensures that no value is present for SupportWebsiteUrl, not even an explicit nil
### GetUnknownSessionRedirectUrl

`func (o *UpdateSignInExp200Response) GetUnknownSessionRedirectUrl() string`

GetUnknownSessionRedirectUrl returns the UnknownSessionRedirectUrl field if non-nil, zero value otherwise.

### GetUnknownSessionRedirectUrlOk

`func (o *UpdateSignInExp200Response) GetUnknownSessionRedirectUrlOk() (*string, bool)`

GetUnknownSessionRedirectUrlOk returns a tuple with the UnknownSessionRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownSessionRedirectUrl

`func (o *UpdateSignInExp200Response) SetUnknownSessionRedirectUrl(v string)`

SetUnknownSessionRedirectUrl sets UnknownSessionRedirectUrl field to given value.


### SetUnknownSessionRedirectUrlNil

`func (o *UpdateSignInExp200Response) SetUnknownSessionRedirectUrlNil(b bool)`

 SetUnknownSessionRedirectUrlNil sets the value for UnknownSessionRedirectUrl to be an explicit nil

### UnsetUnknownSessionRedirectUrl
`func (o *UpdateSignInExp200Response) UnsetUnknownSessionRedirectUrl()`

UnsetUnknownSessionRedirectUrl ensures that no value is present for UnknownSessionRedirectUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


