# GetSignInExperienceConfig200Response

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
**SocialConnectors** | [**[]GetSignInExperienceConfig200ResponseSocialConnectorsInner**](GetSignInExperienceConfig200ResponseSocialConnectorsInner.md) |  | 
**SsoConnectors** | [**[]GetSignInExperienceConfig200ResponseSsoConnectorsInner**](GetSignInExperienceConfig200ResponseSsoConnectorsInner.md) |  | 
**ForgotPassword** | [**GetSignInExperienceConfig200ResponseForgotPassword**](GetSignInExperienceConfig200ResponseForgotPassword.md) |  | 
**IsDevelopmentTenant** | **bool** |  | 
**GoogleOneTap** | Pointer to [**GetSignInExperienceConfig200ResponseGoogleOneTap**](GetSignInExperienceConfig200ResponseGoogleOneTap.md) |  | [optional] 

## Methods

### NewGetSignInExperienceConfig200Response

`func NewGetSignInExperienceConfig200Response(tenantId string, id string, color UpdateSignInExp200ResponseColor, branding ListApplicationOrganizations200ResponseInnerBranding, languageInfo UpdateSignInExp200ResponseLanguageInfo, termsOfUseUrl NullableString, privacyPolicyUrl NullableString, agreeToTermsPolicy string, signIn UpdateSignInExp200ResponseSignIn, signUp UpdateSignInExp200ResponseSignUp, socialSignIn GetSignInExp200ResponseSocialSignIn, socialSignInConnectorTargets []string, signInMode string, customCss NullableString, customContent map[string]string, customUiAssets NullableGetSignInExp200ResponseCustomUiAssets, passwordPolicy UpdateSignInExp200ResponsePasswordPolicy, mfa UpdateSignInExp200ResponseMfa, singleSignOnEnabled bool, supportEmail NullableString, supportWebsiteUrl NullableString, unknownSessionRedirectUrl NullableString, socialConnectors []GetSignInExperienceConfig200ResponseSocialConnectorsInner, ssoConnectors []GetSignInExperienceConfig200ResponseSsoConnectorsInner, forgotPassword GetSignInExperienceConfig200ResponseForgotPassword, isDevelopmentTenant bool, ) *GetSignInExperienceConfig200Response`

NewGetSignInExperienceConfig200Response instantiates a new GetSignInExperienceConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignInExperienceConfig200ResponseWithDefaults

`func NewGetSignInExperienceConfig200ResponseWithDefaults() *GetSignInExperienceConfig200Response`

NewGetSignInExperienceConfig200ResponseWithDefaults instantiates a new GetSignInExperienceConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *GetSignInExperienceConfig200Response) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GetSignInExperienceConfig200Response) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GetSignInExperienceConfig200Response) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *GetSignInExperienceConfig200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSignInExperienceConfig200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSignInExperienceConfig200Response) SetId(v string)`

SetId sets Id field to given value.


### GetColor

`func (o *GetSignInExperienceConfig200Response) GetColor() UpdateSignInExp200ResponseColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *GetSignInExperienceConfig200Response) GetColorOk() (*UpdateSignInExp200ResponseColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *GetSignInExperienceConfig200Response) SetColor(v UpdateSignInExp200ResponseColor)`

SetColor sets Color field to given value.


### GetBranding

`func (o *GetSignInExperienceConfig200Response) GetBranding() ListApplicationOrganizations200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *GetSignInExperienceConfig200Response) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *GetSignInExperienceConfig200Response) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding)`

SetBranding sets Branding field to given value.


### GetLanguageInfo

`func (o *GetSignInExperienceConfig200Response) GetLanguageInfo() UpdateSignInExp200ResponseLanguageInfo`

GetLanguageInfo returns the LanguageInfo field if non-nil, zero value otherwise.

### GetLanguageInfoOk

`func (o *GetSignInExperienceConfig200Response) GetLanguageInfoOk() (*UpdateSignInExp200ResponseLanguageInfo, bool)`

GetLanguageInfoOk returns a tuple with the LanguageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageInfo

`func (o *GetSignInExperienceConfig200Response) SetLanguageInfo(v UpdateSignInExp200ResponseLanguageInfo)`

SetLanguageInfo sets LanguageInfo field to given value.


### GetTermsOfUseUrl

`func (o *GetSignInExperienceConfig200Response) GetTermsOfUseUrl() string`

GetTermsOfUseUrl returns the TermsOfUseUrl field if non-nil, zero value otherwise.

### GetTermsOfUseUrlOk

`func (o *GetSignInExperienceConfig200Response) GetTermsOfUseUrlOk() (*string, bool)`

GetTermsOfUseUrlOk returns a tuple with the TermsOfUseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUseUrl

`func (o *GetSignInExperienceConfig200Response) SetTermsOfUseUrl(v string)`

SetTermsOfUseUrl sets TermsOfUseUrl field to given value.


### SetTermsOfUseUrlNil

`func (o *GetSignInExperienceConfig200Response) SetTermsOfUseUrlNil(b bool)`

 SetTermsOfUseUrlNil sets the value for TermsOfUseUrl to be an explicit nil

### UnsetTermsOfUseUrl
`func (o *GetSignInExperienceConfig200Response) UnsetTermsOfUseUrl()`

UnsetTermsOfUseUrl ensures that no value is present for TermsOfUseUrl, not even an explicit nil
### GetPrivacyPolicyUrl

`func (o *GetSignInExperienceConfig200Response) GetPrivacyPolicyUrl() string`

GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetPrivacyPolicyUrlOk

`func (o *GetSignInExperienceConfig200Response) GetPrivacyPolicyUrlOk() (*string, bool)`

GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyUrl

`func (o *GetSignInExperienceConfig200Response) SetPrivacyPolicyUrl(v string)`

SetPrivacyPolicyUrl sets PrivacyPolicyUrl field to given value.


### SetPrivacyPolicyUrlNil

`func (o *GetSignInExperienceConfig200Response) SetPrivacyPolicyUrlNil(b bool)`

 SetPrivacyPolicyUrlNil sets the value for PrivacyPolicyUrl to be an explicit nil

### UnsetPrivacyPolicyUrl
`func (o *GetSignInExperienceConfig200Response) UnsetPrivacyPolicyUrl()`

UnsetPrivacyPolicyUrl ensures that no value is present for PrivacyPolicyUrl, not even an explicit nil
### GetAgreeToTermsPolicy

`func (o *GetSignInExperienceConfig200Response) GetAgreeToTermsPolicy() string`

GetAgreeToTermsPolicy returns the AgreeToTermsPolicy field if non-nil, zero value otherwise.

### GetAgreeToTermsPolicyOk

`func (o *GetSignInExperienceConfig200Response) GetAgreeToTermsPolicyOk() (*string, bool)`

GetAgreeToTermsPolicyOk returns a tuple with the AgreeToTermsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreeToTermsPolicy

`func (o *GetSignInExperienceConfig200Response) SetAgreeToTermsPolicy(v string)`

SetAgreeToTermsPolicy sets AgreeToTermsPolicy field to given value.


### GetSignIn

`func (o *GetSignInExperienceConfig200Response) GetSignIn() UpdateSignInExp200ResponseSignIn`

GetSignIn returns the SignIn field if non-nil, zero value otherwise.

### GetSignInOk

`func (o *GetSignInExperienceConfig200Response) GetSignInOk() (*UpdateSignInExp200ResponseSignIn, bool)`

GetSignInOk returns a tuple with the SignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignIn

`func (o *GetSignInExperienceConfig200Response) SetSignIn(v UpdateSignInExp200ResponseSignIn)`

SetSignIn sets SignIn field to given value.


### GetSignUp

`func (o *GetSignInExperienceConfig200Response) GetSignUp() UpdateSignInExp200ResponseSignUp`

GetSignUp returns the SignUp field if non-nil, zero value otherwise.

### GetSignUpOk

`func (o *GetSignInExperienceConfig200Response) GetSignUpOk() (*UpdateSignInExp200ResponseSignUp, bool)`

GetSignUpOk returns a tuple with the SignUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUp

`func (o *GetSignInExperienceConfig200Response) SetSignUp(v UpdateSignInExp200ResponseSignUp)`

SetSignUp sets SignUp field to given value.


### GetSocialSignIn

`func (o *GetSignInExperienceConfig200Response) GetSocialSignIn() GetSignInExp200ResponseSocialSignIn`

GetSocialSignIn returns the SocialSignIn field if non-nil, zero value otherwise.

### GetSocialSignInOk

`func (o *GetSignInExperienceConfig200Response) GetSocialSignInOk() (*GetSignInExp200ResponseSocialSignIn, bool)`

GetSocialSignInOk returns a tuple with the SocialSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSignIn

`func (o *GetSignInExperienceConfig200Response) SetSocialSignIn(v GetSignInExp200ResponseSocialSignIn)`

SetSocialSignIn sets SocialSignIn field to given value.


### GetSocialSignInConnectorTargets

`func (o *GetSignInExperienceConfig200Response) GetSocialSignInConnectorTargets() []string`

GetSocialSignInConnectorTargets returns the SocialSignInConnectorTargets field if non-nil, zero value otherwise.

### GetSocialSignInConnectorTargetsOk

`func (o *GetSignInExperienceConfig200Response) GetSocialSignInConnectorTargetsOk() (*[]string, bool)`

GetSocialSignInConnectorTargetsOk returns a tuple with the SocialSignInConnectorTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSignInConnectorTargets

`func (o *GetSignInExperienceConfig200Response) SetSocialSignInConnectorTargets(v []string)`

SetSocialSignInConnectorTargets sets SocialSignInConnectorTargets field to given value.


### GetSignInMode

`func (o *GetSignInExperienceConfig200Response) GetSignInMode() string`

GetSignInMode returns the SignInMode field if non-nil, zero value otherwise.

### GetSignInModeOk

`func (o *GetSignInExperienceConfig200Response) GetSignInModeOk() (*string, bool)`

GetSignInModeOk returns a tuple with the SignInMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInMode

`func (o *GetSignInExperienceConfig200Response) SetSignInMode(v string)`

SetSignInMode sets SignInMode field to given value.


### GetCustomCss

`func (o *GetSignInExperienceConfig200Response) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *GetSignInExperienceConfig200Response) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *GetSignInExperienceConfig200Response) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.


### SetCustomCssNil

`func (o *GetSignInExperienceConfig200Response) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *GetSignInExperienceConfig200Response) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetCustomContent

`func (o *GetSignInExperienceConfig200Response) GetCustomContent() map[string]string`

GetCustomContent returns the CustomContent field if non-nil, zero value otherwise.

### GetCustomContentOk

`func (o *GetSignInExperienceConfig200Response) GetCustomContentOk() (*map[string]string, bool)`

GetCustomContentOk returns a tuple with the CustomContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContent

`func (o *GetSignInExperienceConfig200Response) SetCustomContent(v map[string]string)`

SetCustomContent sets CustomContent field to given value.


### GetCustomUiAssets

`func (o *GetSignInExperienceConfig200Response) GetCustomUiAssets() GetSignInExp200ResponseCustomUiAssets`

GetCustomUiAssets returns the CustomUiAssets field if non-nil, zero value otherwise.

### GetCustomUiAssetsOk

`func (o *GetSignInExperienceConfig200Response) GetCustomUiAssetsOk() (*GetSignInExp200ResponseCustomUiAssets, bool)`

GetCustomUiAssetsOk returns a tuple with the CustomUiAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUiAssets

`func (o *GetSignInExperienceConfig200Response) SetCustomUiAssets(v GetSignInExp200ResponseCustomUiAssets)`

SetCustomUiAssets sets CustomUiAssets field to given value.


### SetCustomUiAssetsNil

`func (o *GetSignInExperienceConfig200Response) SetCustomUiAssetsNil(b bool)`

 SetCustomUiAssetsNil sets the value for CustomUiAssets to be an explicit nil

### UnsetCustomUiAssets
`func (o *GetSignInExperienceConfig200Response) UnsetCustomUiAssets()`

UnsetCustomUiAssets ensures that no value is present for CustomUiAssets, not even an explicit nil
### GetPasswordPolicy

`func (o *GetSignInExperienceConfig200Response) GetPasswordPolicy() UpdateSignInExp200ResponsePasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *GetSignInExperienceConfig200Response) GetPasswordPolicyOk() (*UpdateSignInExp200ResponsePasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *GetSignInExperienceConfig200Response) SetPasswordPolicy(v UpdateSignInExp200ResponsePasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.


### GetMfa

`func (o *GetSignInExperienceConfig200Response) GetMfa() UpdateSignInExp200ResponseMfa`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *GetSignInExperienceConfig200Response) GetMfaOk() (*UpdateSignInExp200ResponseMfa, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *GetSignInExperienceConfig200Response) SetMfa(v UpdateSignInExp200ResponseMfa)`

SetMfa sets Mfa field to given value.


### GetSingleSignOnEnabled

`func (o *GetSignInExperienceConfig200Response) GetSingleSignOnEnabled() bool`

GetSingleSignOnEnabled returns the SingleSignOnEnabled field if non-nil, zero value otherwise.

### GetSingleSignOnEnabledOk

`func (o *GetSignInExperienceConfig200Response) GetSingleSignOnEnabledOk() (*bool, bool)`

GetSingleSignOnEnabledOk returns a tuple with the SingleSignOnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSignOnEnabled

`func (o *GetSignInExperienceConfig200Response) SetSingleSignOnEnabled(v bool)`

SetSingleSignOnEnabled sets SingleSignOnEnabled field to given value.


### GetSupportEmail

`func (o *GetSignInExperienceConfig200Response) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *GetSignInExperienceConfig200Response) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *GetSignInExperienceConfig200Response) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.


### SetSupportEmailNil

`func (o *GetSignInExperienceConfig200Response) SetSupportEmailNil(b bool)`

 SetSupportEmailNil sets the value for SupportEmail to be an explicit nil

### UnsetSupportEmail
`func (o *GetSignInExperienceConfig200Response) UnsetSupportEmail()`

UnsetSupportEmail ensures that no value is present for SupportEmail, not even an explicit nil
### GetSupportWebsiteUrl

`func (o *GetSignInExperienceConfig200Response) GetSupportWebsiteUrl() string`

GetSupportWebsiteUrl returns the SupportWebsiteUrl field if non-nil, zero value otherwise.

### GetSupportWebsiteUrlOk

`func (o *GetSignInExperienceConfig200Response) GetSupportWebsiteUrlOk() (*string, bool)`

GetSupportWebsiteUrlOk returns a tuple with the SupportWebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportWebsiteUrl

`func (o *GetSignInExperienceConfig200Response) SetSupportWebsiteUrl(v string)`

SetSupportWebsiteUrl sets SupportWebsiteUrl field to given value.


### SetSupportWebsiteUrlNil

`func (o *GetSignInExperienceConfig200Response) SetSupportWebsiteUrlNil(b bool)`

 SetSupportWebsiteUrlNil sets the value for SupportWebsiteUrl to be an explicit nil

### UnsetSupportWebsiteUrl
`func (o *GetSignInExperienceConfig200Response) UnsetSupportWebsiteUrl()`

UnsetSupportWebsiteUrl ensures that no value is present for SupportWebsiteUrl, not even an explicit nil
### GetUnknownSessionRedirectUrl

`func (o *GetSignInExperienceConfig200Response) GetUnknownSessionRedirectUrl() string`

GetUnknownSessionRedirectUrl returns the UnknownSessionRedirectUrl field if non-nil, zero value otherwise.

### GetUnknownSessionRedirectUrlOk

`func (o *GetSignInExperienceConfig200Response) GetUnknownSessionRedirectUrlOk() (*string, bool)`

GetUnknownSessionRedirectUrlOk returns a tuple with the UnknownSessionRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownSessionRedirectUrl

`func (o *GetSignInExperienceConfig200Response) SetUnknownSessionRedirectUrl(v string)`

SetUnknownSessionRedirectUrl sets UnknownSessionRedirectUrl field to given value.


### SetUnknownSessionRedirectUrlNil

`func (o *GetSignInExperienceConfig200Response) SetUnknownSessionRedirectUrlNil(b bool)`

 SetUnknownSessionRedirectUrlNil sets the value for UnknownSessionRedirectUrl to be an explicit nil

### UnsetUnknownSessionRedirectUrl
`func (o *GetSignInExperienceConfig200Response) UnsetUnknownSessionRedirectUrl()`

UnsetUnknownSessionRedirectUrl ensures that no value is present for UnknownSessionRedirectUrl, not even an explicit nil
### GetSocialConnectors

`func (o *GetSignInExperienceConfig200Response) GetSocialConnectors() []GetSignInExperienceConfig200ResponseSocialConnectorsInner`

GetSocialConnectors returns the SocialConnectors field if non-nil, zero value otherwise.

### GetSocialConnectorsOk

`func (o *GetSignInExperienceConfig200Response) GetSocialConnectorsOk() (*[]GetSignInExperienceConfig200ResponseSocialConnectorsInner, bool)`

GetSocialConnectorsOk returns a tuple with the SocialConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialConnectors

`func (o *GetSignInExperienceConfig200Response) SetSocialConnectors(v []GetSignInExperienceConfig200ResponseSocialConnectorsInner)`

SetSocialConnectors sets SocialConnectors field to given value.


### GetSsoConnectors

`func (o *GetSignInExperienceConfig200Response) GetSsoConnectors() []GetSignInExperienceConfig200ResponseSsoConnectorsInner`

GetSsoConnectors returns the SsoConnectors field if non-nil, zero value otherwise.

### GetSsoConnectorsOk

`func (o *GetSignInExperienceConfig200Response) GetSsoConnectorsOk() (*[]GetSignInExperienceConfig200ResponseSsoConnectorsInner, bool)`

GetSsoConnectorsOk returns a tuple with the SsoConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoConnectors

`func (o *GetSignInExperienceConfig200Response) SetSsoConnectors(v []GetSignInExperienceConfig200ResponseSsoConnectorsInner)`

SetSsoConnectors sets SsoConnectors field to given value.


### GetForgotPassword

`func (o *GetSignInExperienceConfig200Response) GetForgotPassword() GetSignInExperienceConfig200ResponseForgotPassword`

GetForgotPassword returns the ForgotPassword field if non-nil, zero value otherwise.

### GetForgotPasswordOk

`func (o *GetSignInExperienceConfig200Response) GetForgotPasswordOk() (*GetSignInExperienceConfig200ResponseForgotPassword, bool)`

GetForgotPasswordOk returns a tuple with the ForgotPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgotPassword

`func (o *GetSignInExperienceConfig200Response) SetForgotPassword(v GetSignInExperienceConfig200ResponseForgotPassword)`

SetForgotPassword sets ForgotPassword field to given value.


### GetIsDevelopmentTenant

`func (o *GetSignInExperienceConfig200Response) GetIsDevelopmentTenant() bool`

GetIsDevelopmentTenant returns the IsDevelopmentTenant field if non-nil, zero value otherwise.

### GetIsDevelopmentTenantOk

`func (o *GetSignInExperienceConfig200Response) GetIsDevelopmentTenantOk() (*bool, bool)`

GetIsDevelopmentTenantOk returns a tuple with the IsDevelopmentTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDevelopmentTenant

`func (o *GetSignInExperienceConfig200Response) SetIsDevelopmentTenant(v bool)`

SetIsDevelopmentTenant sets IsDevelopmentTenant field to given value.


### GetGoogleOneTap

`func (o *GetSignInExperienceConfig200Response) GetGoogleOneTap() GetSignInExperienceConfig200ResponseGoogleOneTap`

GetGoogleOneTap returns the GoogleOneTap field if non-nil, zero value otherwise.

### GetGoogleOneTapOk

`func (o *GetSignInExperienceConfig200Response) GetGoogleOneTapOk() (*GetSignInExperienceConfig200ResponseGoogleOneTap, bool)`

GetGoogleOneTapOk returns a tuple with the GoogleOneTap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleOneTap

`func (o *GetSignInExperienceConfig200Response) SetGoogleOneTap(v GetSignInExperienceConfig200ResponseGoogleOneTap)`

SetGoogleOneTap sets GoogleOneTap field to given value.

### HasGoogleOneTap

`func (o *GetSignInExperienceConfig200Response) HasGoogleOneTap() bool`

HasGoogleOneTap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


