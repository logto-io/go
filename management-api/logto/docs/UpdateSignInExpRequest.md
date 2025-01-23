# UpdateSignInExpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Color** | Pointer to [**UpdateSignInExpRequestColor**](UpdateSignInExpRequestColor.md) |  | [optional] 
**Branding** | Pointer to [**ListApplicationOrganizations200ResponseInnerBranding**](ListApplicationOrganizations200ResponseInnerBranding.md) |  | [optional] 
**LanguageInfo** | Pointer to [**UpdateSignInExpRequestLanguageInfo**](UpdateSignInExpRequestLanguageInfo.md) |  | [optional] 
**AgreeToTermsPolicy** | Pointer to **string** |  | [optional] 
**SignIn** | Pointer to [**UpdateSignInExpRequestSignIn**](UpdateSignInExpRequestSignIn.md) |  | [optional] 
**SignUp** | Pointer to [**UpdateSignInExpRequestSignUp**](UpdateSignInExpRequestSignUp.md) |  | [optional] 
**SocialSignIn** | Pointer to [**GetSignInExp200ResponseSocialSignIn**](GetSignInExp200ResponseSocialSignIn.md) |  | [optional] 
**SocialSignInConnectorTargets** | Pointer to **[]string** | Specify the social sign-in connectors to display on the sign-in page. | [optional] 
**SignInMode** | Pointer to **string** |  | [optional] 
**CustomCss** | Pointer to **NullableString** |  | [optional] 
**CustomContent** | Pointer to **map[string]string** | Custom content to display on experience flow pages. the page pathname will be the config key, the content will be the config value. | [optional] 
**CustomUiAssets** | Pointer to [**NullableGetSignInExp200ResponseCustomUiAssets**](GetSignInExp200ResponseCustomUiAssets.md) |  | [optional] 
**PasswordPolicy** | Pointer to [**GetSignInExp200ResponsePasswordPolicy**](GetSignInExp200ResponsePasswordPolicy.md) |  | [optional] 
**Mfa** | Pointer to [**GetSignInExp200ResponseMfa**](GetSignInExp200ResponseMfa.md) |  | [optional] 
**SingleSignOnEnabled** | Pointer to **bool** |  | [optional] 
**TermsOfUseUrl** | Pointer to [**UpdateSignInExpRequestTermsOfUseUrl**](UpdateSignInExpRequestTermsOfUseUrl.md) |  | [optional] 
**PrivacyPolicyUrl** | Pointer to [**UpdateSignInExpRequestTermsOfUseUrl**](UpdateSignInExpRequestTermsOfUseUrl.md) |  | [optional] 
**SupportEmail** | Pointer to [**UpdateSignInExpRequestSupportEmail**](UpdateSignInExpRequestSupportEmail.md) |  | [optional] 
**SupportWebsiteUrl** | Pointer to [**UpdateSignInExpRequestSupportWebsiteUrl**](UpdateSignInExpRequestSupportWebsiteUrl.md) |  | [optional] 
**UnknownSessionRedirectUrl** | Pointer to [**UpdateSignInExpRequestUnknownSessionRedirectUrl**](UpdateSignInExpRequestUnknownSessionRedirectUrl.md) |  | [optional] 

## Methods

### NewUpdateSignInExpRequest

`func NewUpdateSignInExpRequest() *UpdateSignInExpRequest`

NewUpdateSignInExpRequest instantiates a new UpdateSignInExpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSignInExpRequestWithDefaults

`func NewUpdateSignInExpRequestWithDefaults() *UpdateSignInExpRequest`

NewUpdateSignInExpRequestWithDefaults instantiates a new UpdateSignInExpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *UpdateSignInExpRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateSignInExpRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateSignInExpRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UpdateSignInExpRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetColor

`func (o *UpdateSignInExpRequest) GetColor() UpdateSignInExpRequestColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UpdateSignInExpRequest) GetColorOk() (*UpdateSignInExpRequestColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UpdateSignInExpRequest) SetColor(v UpdateSignInExpRequestColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *UpdateSignInExpRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetBranding

`func (o *UpdateSignInExpRequest) GetBranding() ListApplicationOrganizations200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *UpdateSignInExpRequest) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *UpdateSignInExpRequest) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *UpdateSignInExpRequest) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetLanguageInfo

`func (o *UpdateSignInExpRequest) GetLanguageInfo() UpdateSignInExpRequestLanguageInfo`

GetLanguageInfo returns the LanguageInfo field if non-nil, zero value otherwise.

### GetLanguageInfoOk

`func (o *UpdateSignInExpRequest) GetLanguageInfoOk() (*UpdateSignInExpRequestLanguageInfo, bool)`

GetLanguageInfoOk returns a tuple with the LanguageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageInfo

`func (o *UpdateSignInExpRequest) SetLanguageInfo(v UpdateSignInExpRequestLanguageInfo)`

SetLanguageInfo sets LanguageInfo field to given value.

### HasLanguageInfo

`func (o *UpdateSignInExpRequest) HasLanguageInfo() bool`

HasLanguageInfo returns a boolean if a field has been set.

### GetAgreeToTermsPolicy

`func (o *UpdateSignInExpRequest) GetAgreeToTermsPolicy() string`

GetAgreeToTermsPolicy returns the AgreeToTermsPolicy field if non-nil, zero value otherwise.

### GetAgreeToTermsPolicyOk

`func (o *UpdateSignInExpRequest) GetAgreeToTermsPolicyOk() (*string, bool)`

GetAgreeToTermsPolicyOk returns a tuple with the AgreeToTermsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreeToTermsPolicy

`func (o *UpdateSignInExpRequest) SetAgreeToTermsPolicy(v string)`

SetAgreeToTermsPolicy sets AgreeToTermsPolicy field to given value.

### HasAgreeToTermsPolicy

`func (o *UpdateSignInExpRequest) HasAgreeToTermsPolicy() bool`

HasAgreeToTermsPolicy returns a boolean if a field has been set.

### GetSignIn

`func (o *UpdateSignInExpRequest) GetSignIn() UpdateSignInExpRequestSignIn`

GetSignIn returns the SignIn field if non-nil, zero value otherwise.

### GetSignInOk

`func (o *UpdateSignInExpRequest) GetSignInOk() (*UpdateSignInExpRequestSignIn, bool)`

GetSignInOk returns a tuple with the SignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignIn

`func (o *UpdateSignInExpRequest) SetSignIn(v UpdateSignInExpRequestSignIn)`

SetSignIn sets SignIn field to given value.

### HasSignIn

`func (o *UpdateSignInExpRequest) HasSignIn() bool`

HasSignIn returns a boolean if a field has been set.

### GetSignUp

`func (o *UpdateSignInExpRequest) GetSignUp() UpdateSignInExpRequestSignUp`

GetSignUp returns the SignUp field if non-nil, zero value otherwise.

### GetSignUpOk

`func (o *UpdateSignInExpRequest) GetSignUpOk() (*UpdateSignInExpRequestSignUp, bool)`

GetSignUpOk returns a tuple with the SignUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUp

`func (o *UpdateSignInExpRequest) SetSignUp(v UpdateSignInExpRequestSignUp)`

SetSignUp sets SignUp field to given value.

### HasSignUp

`func (o *UpdateSignInExpRequest) HasSignUp() bool`

HasSignUp returns a boolean if a field has been set.

### GetSocialSignIn

`func (o *UpdateSignInExpRequest) GetSocialSignIn() GetSignInExp200ResponseSocialSignIn`

GetSocialSignIn returns the SocialSignIn field if non-nil, zero value otherwise.

### GetSocialSignInOk

`func (o *UpdateSignInExpRequest) GetSocialSignInOk() (*GetSignInExp200ResponseSocialSignIn, bool)`

GetSocialSignInOk returns a tuple with the SocialSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSignIn

`func (o *UpdateSignInExpRequest) SetSocialSignIn(v GetSignInExp200ResponseSocialSignIn)`

SetSocialSignIn sets SocialSignIn field to given value.

### HasSocialSignIn

`func (o *UpdateSignInExpRequest) HasSocialSignIn() bool`

HasSocialSignIn returns a boolean if a field has been set.

### GetSocialSignInConnectorTargets

`func (o *UpdateSignInExpRequest) GetSocialSignInConnectorTargets() []string`

GetSocialSignInConnectorTargets returns the SocialSignInConnectorTargets field if non-nil, zero value otherwise.

### GetSocialSignInConnectorTargetsOk

`func (o *UpdateSignInExpRequest) GetSocialSignInConnectorTargetsOk() (*[]string, bool)`

GetSocialSignInConnectorTargetsOk returns a tuple with the SocialSignInConnectorTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSignInConnectorTargets

`func (o *UpdateSignInExpRequest) SetSocialSignInConnectorTargets(v []string)`

SetSocialSignInConnectorTargets sets SocialSignInConnectorTargets field to given value.

### HasSocialSignInConnectorTargets

`func (o *UpdateSignInExpRequest) HasSocialSignInConnectorTargets() bool`

HasSocialSignInConnectorTargets returns a boolean if a field has been set.

### GetSignInMode

`func (o *UpdateSignInExpRequest) GetSignInMode() string`

GetSignInMode returns the SignInMode field if non-nil, zero value otherwise.

### GetSignInModeOk

`func (o *UpdateSignInExpRequest) GetSignInModeOk() (*string, bool)`

GetSignInModeOk returns a tuple with the SignInMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInMode

`func (o *UpdateSignInExpRequest) SetSignInMode(v string)`

SetSignInMode sets SignInMode field to given value.

### HasSignInMode

`func (o *UpdateSignInExpRequest) HasSignInMode() bool`

HasSignInMode returns a boolean if a field has been set.

### GetCustomCss

`func (o *UpdateSignInExpRequest) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *UpdateSignInExpRequest) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *UpdateSignInExpRequest) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.

### HasCustomCss

`func (o *UpdateSignInExpRequest) HasCustomCss() bool`

HasCustomCss returns a boolean if a field has been set.

### SetCustomCssNil

`func (o *UpdateSignInExpRequest) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *UpdateSignInExpRequest) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetCustomContent

`func (o *UpdateSignInExpRequest) GetCustomContent() map[string]string`

GetCustomContent returns the CustomContent field if non-nil, zero value otherwise.

### GetCustomContentOk

`func (o *UpdateSignInExpRequest) GetCustomContentOk() (*map[string]string, bool)`

GetCustomContentOk returns a tuple with the CustomContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContent

`func (o *UpdateSignInExpRequest) SetCustomContent(v map[string]string)`

SetCustomContent sets CustomContent field to given value.

### HasCustomContent

`func (o *UpdateSignInExpRequest) HasCustomContent() bool`

HasCustomContent returns a boolean if a field has been set.

### GetCustomUiAssets

`func (o *UpdateSignInExpRequest) GetCustomUiAssets() GetSignInExp200ResponseCustomUiAssets`

GetCustomUiAssets returns the CustomUiAssets field if non-nil, zero value otherwise.

### GetCustomUiAssetsOk

`func (o *UpdateSignInExpRequest) GetCustomUiAssetsOk() (*GetSignInExp200ResponseCustomUiAssets, bool)`

GetCustomUiAssetsOk returns a tuple with the CustomUiAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUiAssets

`func (o *UpdateSignInExpRequest) SetCustomUiAssets(v GetSignInExp200ResponseCustomUiAssets)`

SetCustomUiAssets sets CustomUiAssets field to given value.

### HasCustomUiAssets

`func (o *UpdateSignInExpRequest) HasCustomUiAssets() bool`

HasCustomUiAssets returns a boolean if a field has been set.

### SetCustomUiAssetsNil

`func (o *UpdateSignInExpRequest) SetCustomUiAssetsNil(b bool)`

 SetCustomUiAssetsNil sets the value for CustomUiAssets to be an explicit nil

### UnsetCustomUiAssets
`func (o *UpdateSignInExpRequest) UnsetCustomUiAssets()`

UnsetCustomUiAssets ensures that no value is present for CustomUiAssets, not even an explicit nil
### GetPasswordPolicy

`func (o *UpdateSignInExpRequest) GetPasswordPolicy() GetSignInExp200ResponsePasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *UpdateSignInExpRequest) GetPasswordPolicyOk() (*GetSignInExp200ResponsePasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *UpdateSignInExpRequest) SetPasswordPolicy(v GetSignInExp200ResponsePasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *UpdateSignInExpRequest) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetMfa

`func (o *UpdateSignInExpRequest) GetMfa() GetSignInExp200ResponseMfa`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *UpdateSignInExpRequest) GetMfaOk() (*GetSignInExp200ResponseMfa, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *UpdateSignInExpRequest) SetMfa(v GetSignInExp200ResponseMfa)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *UpdateSignInExpRequest) HasMfa() bool`

HasMfa returns a boolean if a field has been set.

### GetSingleSignOnEnabled

`func (o *UpdateSignInExpRequest) GetSingleSignOnEnabled() bool`

GetSingleSignOnEnabled returns the SingleSignOnEnabled field if non-nil, zero value otherwise.

### GetSingleSignOnEnabledOk

`func (o *UpdateSignInExpRequest) GetSingleSignOnEnabledOk() (*bool, bool)`

GetSingleSignOnEnabledOk returns a tuple with the SingleSignOnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSignOnEnabled

`func (o *UpdateSignInExpRequest) SetSingleSignOnEnabled(v bool)`

SetSingleSignOnEnabled sets SingleSignOnEnabled field to given value.

### HasSingleSignOnEnabled

`func (o *UpdateSignInExpRequest) HasSingleSignOnEnabled() bool`

HasSingleSignOnEnabled returns a boolean if a field has been set.

### GetTermsOfUseUrl

`func (o *UpdateSignInExpRequest) GetTermsOfUseUrl() UpdateSignInExpRequestTermsOfUseUrl`

GetTermsOfUseUrl returns the TermsOfUseUrl field if non-nil, zero value otherwise.

### GetTermsOfUseUrlOk

`func (o *UpdateSignInExpRequest) GetTermsOfUseUrlOk() (*UpdateSignInExpRequestTermsOfUseUrl, bool)`

GetTermsOfUseUrlOk returns a tuple with the TermsOfUseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUseUrl

`func (o *UpdateSignInExpRequest) SetTermsOfUseUrl(v UpdateSignInExpRequestTermsOfUseUrl)`

SetTermsOfUseUrl sets TermsOfUseUrl field to given value.

### HasTermsOfUseUrl

`func (o *UpdateSignInExpRequest) HasTermsOfUseUrl() bool`

HasTermsOfUseUrl returns a boolean if a field has been set.

### GetPrivacyPolicyUrl

`func (o *UpdateSignInExpRequest) GetPrivacyPolicyUrl() UpdateSignInExpRequestTermsOfUseUrl`

GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetPrivacyPolicyUrlOk

`func (o *UpdateSignInExpRequest) GetPrivacyPolicyUrlOk() (*UpdateSignInExpRequestTermsOfUseUrl, bool)`

GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyUrl

`func (o *UpdateSignInExpRequest) SetPrivacyPolicyUrl(v UpdateSignInExpRequestTermsOfUseUrl)`

SetPrivacyPolicyUrl sets PrivacyPolicyUrl field to given value.

### HasPrivacyPolicyUrl

`func (o *UpdateSignInExpRequest) HasPrivacyPolicyUrl() bool`

HasPrivacyPolicyUrl returns a boolean if a field has been set.

### GetSupportEmail

`func (o *UpdateSignInExpRequest) GetSupportEmail() UpdateSignInExpRequestSupportEmail`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *UpdateSignInExpRequest) GetSupportEmailOk() (*UpdateSignInExpRequestSupportEmail, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *UpdateSignInExpRequest) SetSupportEmail(v UpdateSignInExpRequestSupportEmail)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *UpdateSignInExpRequest) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### GetSupportWebsiteUrl

`func (o *UpdateSignInExpRequest) GetSupportWebsiteUrl() UpdateSignInExpRequestSupportWebsiteUrl`

GetSupportWebsiteUrl returns the SupportWebsiteUrl field if non-nil, zero value otherwise.

### GetSupportWebsiteUrlOk

`func (o *UpdateSignInExpRequest) GetSupportWebsiteUrlOk() (*UpdateSignInExpRequestSupportWebsiteUrl, bool)`

GetSupportWebsiteUrlOk returns a tuple with the SupportWebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportWebsiteUrl

`func (o *UpdateSignInExpRequest) SetSupportWebsiteUrl(v UpdateSignInExpRequestSupportWebsiteUrl)`

SetSupportWebsiteUrl sets SupportWebsiteUrl field to given value.

### HasSupportWebsiteUrl

`func (o *UpdateSignInExpRequest) HasSupportWebsiteUrl() bool`

HasSupportWebsiteUrl returns a boolean if a field has been set.

### GetUnknownSessionRedirectUrl

`func (o *UpdateSignInExpRequest) GetUnknownSessionRedirectUrl() UpdateSignInExpRequestUnknownSessionRedirectUrl`

GetUnknownSessionRedirectUrl returns the UnknownSessionRedirectUrl field if non-nil, zero value otherwise.

### GetUnknownSessionRedirectUrlOk

`func (o *UpdateSignInExpRequest) GetUnknownSessionRedirectUrlOk() (*UpdateSignInExpRequestUnknownSessionRedirectUrl, bool)`

GetUnknownSessionRedirectUrlOk returns a tuple with the UnknownSessionRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownSessionRedirectUrl

`func (o *UpdateSignInExpRequest) SetUnknownSessionRedirectUrl(v UpdateSignInExpRequestUnknownSessionRedirectUrl)`

SetUnknownSessionRedirectUrl sets UnknownSessionRedirectUrl field to given value.

### HasUnknownSessionRedirectUrl

`func (o *UpdateSignInExpRequest) HasUnknownSessionRedirectUrl() bool`

HasUnknownSessionRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


