# GetSignInExp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**Color** | [**GetSignInExp200ResponseColor**](GetSignInExp200ResponseColor.md) |  | 
**Branding** | [**ListApplicationOrganizations200ResponseInnerBranding**](ListApplicationOrganizations200ResponseInnerBranding.md) |  | 
**LanguageInfo** | [**GetSignInExp200ResponseLanguageInfo**](GetSignInExp200ResponseLanguageInfo.md) |  | 
**TermsOfUseUrl** | **NullableString** |  | 
**PrivacyPolicyUrl** | **NullableString** |  | 
**AgreeToTermsPolicy** | **string** |  | 
**SignIn** | [**GetSignInExp200ResponseSignIn**](GetSignInExp200ResponseSignIn.md) |  | 
**SignUp** | [**GetSignInExp200ResponseSignUp**](GetSignInExp200ResponseSignUp.md) |  | 
**SocialSignIn** | [**GetSignInExp200ResponseSocialSignIn**](GetSignInExp200ResponseSocialSignIn.md) |  | 
**SocialSignInConnectorTargets** | **[]string** | Enabled social sign-in connectors, will displayed on the sign-in page. | 
**SignInMode** | **string** |  | 
**CustomCss** | **NullableString** |  | 
**CustomContent** | **map[string]string** | Custom content to display on experience flow pages. the page pathname will be the config key, the content will be the config value. | 
**CustomUiAssets** | [**NullableGetSignInExp200ResponseCustomUiAssets**](GetSignInExp200ResponseCustomUiAssets.md) |  | 
**PasswordPolicy** | [**GetSignInExp200ResponsePasswordPolicy**](GetSignInExp200ResponsePasswordPolicy.md) |  | 
**Mfa** | [**GetSignInExp200ResponseMfa**](GetSignInExp200ResponseMfa.md) |  | 
**SingleSignOnEnabled** | **bool** |  | 
**SupportEmail** | **NullableString** | The support email address to display on the error pages. | 
**SupportWebsiteUrl** | **NullableString** | The support website URL to display on the error pages. | 
**UnknownSessionRedirectUrl** | **NullableString** | The fallback URL to redirect users when the sign-in session does not exist or unknown. Client should initiates a new authentication flow after the redirection. | 

## Methods

### NewGetSignInExp200Response

`func NewGetSignInExp200Response(tenantId string, id string, color GetSignInExp200ResponseColor, branding ListApplicationOrganizations200ResponseInnerBranding, languageInfo GetSignInExp200ResponseLanguageInfo, termsOfUseUrl NullableString, privacyPolicyUrl NullableString, agreeToTermsPolicy string, signIn GetSignInExp200ResponseSignIn, signUp GetSignInExp200ResponseSignUp, socialSignIn GetSignInExp200ResponseSocialSignIn, socialSignInConnectorTargets []string, signInMode string, customCss NullableString, customContent map[string]string, customUiAssets NullableGetSignInExp200ResponseCustomUiAssets, passwordPolicy GetSignInExp200ResponsePasswordPolicy, mfa GetSignInExp200ResponseMfa, singleSignOnEnabled bool, supportEmail NullableString, supportWebsiteUrl NullableString, unknownSessionRedirectUrl NullableString, ) *GetSignInExp200Response`

NewGetSignInExp200Response instantiates a new GetSignInExp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignInExp200ResponseWithDefaults

`func NewGetSignInExp200ResponseWithDefaults() *GetSignInExp200Response`

NewGetSignInExp200ResponseWithDefaults instantiates a new GetSignInExp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *GetSignInExp200Response) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GetSignInExp200Response) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GetSignInExp200Response) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *GetSignInExp200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSignInExp200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSignInExp200Response) SetId(v string)`

SetId sets Id field to given value.


### GetColor

`func (o *GetSignInExp200Response) GetColor() GetSignInExp200ResponseColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *GetSignInExp200Response) GetColorOk() (*GetSignInExp200ResponseColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *GetSignInExp200Response) SetColor(v GetSignInExp200ResponseColor)`

SetColor sets Color field to given value.


### GetBranding

`func (o *GetSignInExp200Response) GetBranding() ListApplicationOrganizations200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *GetSignInExp200Response) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *GetSignInExp200Response) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding)`

SetBranding sets Branding field to given value.


### GetLanguageInfo

`func (o *GetSignInExp200Response) GetLanguageInfo() GetSignInExp200ResponseLanguageInfo`

GetLanguageInfo returns the LanguageInfo field if non-nil, zero value otherwise.

### GetLanguageInfoOk

`func (o *GetSignInExp200Response) GetLanguageInfoOk() (*GetSignInExp200ResponseLanguageInfo, bool)`

GetLanguageInfoOk returns a tuple with the LanguageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageInfo

`func (o *GetSignInExp200Response) SetLanguageInfo(v GetSignInExp200ResponseLanguageInfo)`

SetLanguageInfo sets LanguageInfo field to given value.


### GetTermsOfUseUrl

`func (o *GetSignInExp200Response) GetTermsOfUseUrl() string`

GetTermsOfUseUrl returns the TermsOfUseUrl field if non-nil, zero value otherwise.

### GetTermsOfUseUrlOk

`func (o *GetSignInExp200Response) GetTermsOfUseUrlOk() (*string, bool)`

GetTermsOfUseUrlOk returns a tuple with the TermsOfUseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUseUrl

`func (o *GetSignInExp200Response) SetTermsOfUseUrl(v string)`

SetTermsOfUseUrl sets TermsOfUseUrl field to given value.


### SetTermsOfUseUrlNil

`func (o *GetSignInExp200Response) SetTermsOfUseUrlNil(b bool)`

 SetTermsOfUseUrlNil sets the value for TermsOfUseUrl to be an explicit nil

### UnsetTermsOfUseUrl
`func (o *GetSignInExp200Response) UnsetTermsOfUseUrl()`

UnsetTermsOfUseUrl ensures that no value is present for TermsOfUseUrl, not even an explicit nil
### GetPrivacyPolicyUrl

`func (o *GetSignInExp200Response) GetPrivacyPolicyUrl() string`

GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetPrivacyPolicyUrlOk

`func (o *GetSignInExp200Response) GetPrivacyPolicyUrlOk() (*string, bool)`

GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyUrl

`func (o *GetSignInExp200Response) SetPrivacyPolicyUrl(v string)`

SetPrivacyPolicyUrl sets PrivacyPolicyUrl field to given value.


### SetPrivacyPolicyUrlNil

`func (o *GetSignInExp200Response) SetPrivacyPolicyUrlNil(b bool)`

 SetPrivacyPolicyUrlNil sets the value for PrivacyPolicyUrl to be an explicit nil

### UnsetPrivacyPolicyUrl
`func (o *GetSignInExp200Response) UnsetPrivacyPolicyUrl()`

UnsetPrivacyPolicyUrl ensures that no value is present for PrivacyPolicyUrl, not even an explicit nil
### GetAgreeToTermsPolicy

`func (o *GetSignInExp200Response) GetAgreeToTermsPolicy() string`

GetAgreeToTermsPolicy returns the AgreeToTermsPolicy field if non-nil, zero value otherwise.

### GetAgreeToTermsPolicyOk

`func (o *GetSignInExp200Response) GetAgreeToTermsPolicyOk() (*string, bool)`

GetAgreeToTermsPolicyOk returns a tuple with the AgreeToTermsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreeToTermsPolicy

`func (o *GetSignInExp200Response) SetAgreeToTermsPolicy(v string)`

SetAgreeToTermsPolicy sets AgreeToTermsPolicy field to given value.


### GetSignIn

`func (o *GetSignInExp200Response) GetSignIn() GetSignInExp200ResponseSignIn`

GetSignIn returns the SignIn field if non-nil, zero value otherwise.

### GetSignInOk

`func (o *GetSignInExp200Response) GetSignInOk() (*GetSignInExp200ResponseSignIn, bool)`

GetSignInOk returns a tuple with the SignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignIn

`func (o *GetSignInExp200Response) SetSignIn(v GetSignInExp200ResponseSignIn)`

SetSignIn sets SignIn field to given value.


### GetSignUp

`func (o *GetSignInExp200Response) GetSignUp() GetSignInExp200ResponseSignUp`

GetSignUp returns the SignUp field if non-nil, zero value otherwise.

### GetSignUpOk

`func (o *GetSignInExp200Response) GetSignUpOk() (*GetSignInExp200ResponseSignUp, bool)`

GetSignUpOk returns a tuple with the SignUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUp

`func (o *GetSignInExp200Response) SetSignUp(v GetSignInExp200ResponseSignUp)`

SetSignUp sets SignUp field to given value.


### GetSocialSignIn

`func (o *GetSignInExp200Response) GetSocialSignIn() GetSignInExp200ResponseSocialSignIn`

GetSocialSignIn returns the SocialSignIn field if non-nil, zero value otherwise.

### GetSocialSignInOk

`func (o *GetSignInExp200Response) GetSocialSignInOk() (*GetSignInExp200ResponseSocialSignIn, bool)`

GetSocialSignInOk returns a tuple with the SocialSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSignIn

`func (o *GetSignInExp200Response) SetSocialSignIn(v GetSignInExp200ResponseSocialSignIn)`

SetSocialSignIn sets SocialSignIn field to given value.


### GetSocialSignInConnectorTargets

`func (o *GetSignInExp200Response) GetSocialSignInConnectorTargets() []string`

GetSocialSignInConnectorTargets returns the SocialSignInConnectorTargets field if non-nil, zero value otherwise.

### GetSocialSignInConnectorTargetsOk

`func (o *GetSignInExp200Response) GetSocialSignInConnectorTargetsOk() (*[]string, bool)`

GetSocialSignInConnectorTargetsOk returns a tuple with the SocialSignInConnectorTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSignInConnectorTargets

`func (o *GetSignInExp200Response) SetSocialSignInConnectorTargets(v []string)`

SetSocialSignInConnectorTargets sets SocialSignInConnectorTargets field to given value.


### GetSignInMode

`func (o *GetSignInExp200Response) GetSignInMode() string`

GetSignInMode returns the SignInMode field if non-nil, zero value otherwise.

### GetSignInModeOk

`func (o *GetSignInExp200Response) GetSignInModeOk() (*string, bool)`

GetSignInModeOk returns a tuple with the SignInMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInMode

`func (o *GetSignInExp200Response) SetSignInMode(v string)`

SetSignInMode sets SignInMode field to given value.


### GetCustomCss

`func (o *GetSignInExp200Response) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *GetSignInExp200Response) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *GetSignInExp200Response) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.


### SetCustomCssNil

`func (o *GetSignInExp200Response) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *GetSignInExp200Response) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetCustomContent

`func (o *GetSignInExp200Response) GetCustomContent() map[string]string`

GetCustomContent returns the CustomContent field if non-nil, zero value otherwise.

### GetCustomContentOk

`func (o *GetSignInExp200Response) GetCustomContentOk() (*map[string]string, bool)`

GetCustomContentOk returns a tuple with the CustomContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContent

`func (o *GetSignInExp200Response) SetCustomContent(v map[string]string)`

SetCustomContent sets CustomContent field to given value.


### GetCustomUiAssets

`func (o *GetSignInExp200Response) GetCustomUiAssets() GetSignInExp200ResponseCustomUiAssets`

GetCustomUiAssets returns the CustomUiAssets field if non-nil, zero value otherwise.

### GetCustomUiAssetsOk

`func (o *GetSignInExp200Response) GetCustomUiAssetsOk() (*GetSignInExp200ResponseCustomUiAssets, bool)`

GetCustomUiAssetsOk returns a tuple with the CustomUiAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUiAssets

`func (o *GetSignInExp200Response) SetCustomUiAssets(v GetSignInExp200ResponseCustomUiAssets)`

SetCustomUiAssets sets CustomUiAssets field to given value.


### SetCustomUiAssetsNil

`func (o *GetSignInExp200Response) SetCustomUiAssetsNil(b bool)`

 SetCustomUiAssetsNil sets the value for CustomUiAssets to be an explicit nil

### UnsetCustomUiAssets
`func (o *GetSignInExp200Response) UnsetCustomUiAssets()`

UnsetCustomUiAssets ensures that no value is present for CustomUiAssets, not even an explicit nil
### GetPasswordPolicy

`func (o *GetSignInExp200Response) GetPasswordPolicy() GetSignInExp200ResponsePasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *GetSignInExp200Response) GetPasswordPolicyOk() (*GetSignInExp200ResponsePasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *GetSignInExp200Response) SetPasswordPolicy(v GetSignInExp200ResponsePasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.


### GetMfa

`func (o *GetSignInExp200Response) GetMfa() GetSignInExp200ResponseMfa`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *GetSignInExp200Response) GetMfaOk() (*GetSignInExp200ResponseMfa, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *GetSignInExp200Response) SetMfa(v GetSignInExp200ResponseMfa)`

SetMfa sets Mfa field to given value.


### GetSingleSignOnEnabled

`func (o *GetSignInExp200Response) GetSingleSignOnEnabled() bool`

GetSingleSignOnEnabled returns the SingleSignOnEnabled field if non-nil, zero value otherwise.

### GetSingleSignOnEnabledOk

`func (o *GetSignInExp200Response) GetSingleSignOnEnabledOk() (*bool, bool)`

GetSingleSignOnEnabledOk returns a tuple with the SingleSignOnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSignOnEnabled

`func (o *GetSignInExp200Response) SetSingleSignOnEnabled(v bool)`

SetSingleSignOnEnabled sets SingleSignOnEnabled field to given value.


### GetSupportEmail

`func (o *GetSignInExp200Response) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *GetSignInExp200Response) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *GetSignInExp200Response) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.


### SetSupportEmailNil

`func (o *GetSignInExp200Response) SetSupportEmailNil(b bool)`

 SetSupportEmailNil sets the value for SupportEmail to be an explicit nil

### UnsetSupportEmail
`func (o *GetSignInExp200Response) UnsetSupportEmail()`

UnsetSupportEmail ensures that no value is present for SupportEmail, not even an explicit nil
### GetSupportWebsiteUrl

`func (o *GetSignInExp200Response) GetSupportWebsiteUrl() string`

GetSupportWebsiteUrl returns the SupportWebsiteUrl field if non-nil, zero value otherwise.

### GetSupportWebsiteUrlOk

`func (o *GetSignInExp200Response) GetSupportWebsiteUrlOk() (*string, bool)`

GetSupportWebsiteUrlOk returns a tuple with the SupportWebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportWebsiteUrl

`func (o *GetSignInExp200Response) SetSupportWebsiteUrl(v string)`

SetSupportWebsiteUrl sets SupportWebsiteUrl field to given value.


### SetSupportWebsiteUrlNil

`func (o *GetSignInExp200Response) SetSupportWebsiteUrlNil(b bool)`

 SetSupportWebsiteUrlNil sets the value for SupportWebsiteUrl to be an explicit nil

### UnsetSupportWebsiteUrl
`func (o *GetSignInExp200Response) UnsetSupportWebsiteUrl()`

UnsetSupportWebsiteUrl ensures that no value is present for SupportWebsiteUrl, not even an explicit nil
### GetUnknownSessionRedirectUrl

`func (o *GetSignInExp200Response) GetUnknownSessionRedirectUrl() string`

GetUnknownSessionRedirectUrl returns the UnknownSessionRedirectUrl field if non-nil, zero value otherwise.

### GetUnknownSessionRedirectUrlOk

`func (o *GetSignInExp200Response) GetUnknownSessionRedirectUrlOk() (*string, bool)`

GetUnknownSessionRedirectUrlOk returns a tuple with the UnknownSessionRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownSessionRedirectUrl

`func (o *GetSignInExp200Response) SetUnknownSessionRedirectUrl(v string)`

SetUnknownSessionRedirectUrl sets UnknownSessionRedirectUrl field to given value.


### SetUnknownSessionRedirectUrlNil

`func (o *GetSignInExp200Response) SetUnknownSessionRedirectUrlNil(b bool)`

 SetUnknownSessionRedirectUrlNil sets the value for UnknownSessionRedirectUrl to be an explicit nil

### UnsetUnknownSessionRedirectUrl
`func (o *GetSignInExp200Response) UnsetUnknownSessionRedirectUrl()`

UnsetUnknownSessionRedirectUrl ensures that no value is present for UnknownSessionRedirectUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


