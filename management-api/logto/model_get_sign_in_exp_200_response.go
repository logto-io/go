/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logto

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetSignInExp200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSignInExp200Response{}

// GetSignInExp200Response struct for GetSignInExp200Response
type GetSignInExp200Response struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	Color GetSignInExp200ResponseColor `json:"color"`
	Branding ListApplicationOrganizations200ResponseInnerBranding `json:"branding"`
	LanguageInfo GetSignInExp200ResponseLanguageInfo `json:"languageInfo"`
	TermsOfUseUrl NullableString `json:"termsOfUseUrl"`
	PrivacyPolicyUrl NullableString `json:"privacyPolicyUrl"`
	AgreeToTermsPolicy string `json:"agreeToTermsPolicy"`
	SignIn GetSignInExp200ResponseSignIn `json:"signIn"`
	SignUp GetSignInExp200ResponseSignUp `json:"signUp"`
	SocialSignIn GetSignInExp200ResponseSocialSignIn `json:"socialSignIn"`
	// Enabled social sign-in connectors, will displayed on the sign-in page.
	SocialSignInConnectorTargets []string `json:"socialSignInConnectorTargets"`
	SignInMode string `json:"signInMode"`
	CustomCss NullableString `json:"customCss"`
	// Custom content to display on experience flow pages. the page pathname will be the config key, the content will be the config value.
	CustomContent map[string]string `json:"customContent"`
	CustomUiAssets NullableGetSignInExp200ResponseCustomUiAssets `json:"customUiAssets"`
	PasswordPolicy GetSignInExp200ResponsePasswordPolicy `json:"passwordPolicy"`
	Mfa GetSignInExp200ResponseMfa `json:"mfa"`
	SingleSignOnEnabled bool `json:"singleSignOnEnabled"`
	// The support email address to display on the error pages.
	SupportEmail NullableString `json:"supportEmail"`
	// The support website URL to display on the error pages.
	SupportWebsiteUrl NullableString `json:"supportWebsiteUrl"`
	// The fallback URL to redirect users when the sign-in session does not exist or unknown. Client should initiates a new authentication flow after the redirection.
	UnknownSessionRedirectUrl NullableString `json:"unknownSessionRedirectUrl"`
}

type _GetSignInExp200Response GetSignInExp200Response

// NewGetSignInExp200Response instantiates a new GetSignInExp200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignInExp200Response(tenantId string, id string, color GetSignInExp200ResponseColor, branding ListApplicationOrganizations200ResponseInnerBranding, languageInfo GetSignInExp200ResponseLanguageInfo, termsOfUseUrl NullableString, privacyPolicyUrl NullableString, agreeToTermsPolicy string, signIn GetSignInExp200ResponseSignIn, signUp GetSignInExp200ResponseSignUp, socialSignIn GetSignInExp200ResponseSocialSignIn, socialSignInConnectorTargets []string, signInMode string, customCss NullableString, customContent map[string]string, customUiAssets NullableGetSignInExp200ResponseCustomUiAssets, passwordPolicy GetSignInExp200ResponsePasswordPolicy, mfa GetSignInExp200ResponseMfa, singleSignOnEnabled bool, supportEmail NullableString, supportWebsiteUrl NullableString, unknownSessionRedirectUrl NullableString) *GetSignInExp200Response {
	this := GetSignInExp200Response{}
	this.TenantId = tenantId
	this.Id = id
	this.Color = color
	this.Branding = branding
	this.LanguageInfo = languageInfo
	this.TermsOfUseUrl = termsOfUseUrl
	this.PrivacyPolicyUrl = privacyPolicyUrl
	this.AgreeToTermsPolicy = agreeToTermsPolicy
	this.SignIn = signIn
	this.SignUp = signUp
	this.SocialSignIn = socialSignIn
	this.SocialSignInConnectorTargets = socialSignInConnectorTargets
	this.SignInMode = signInMode
	this.CustomCss = customCss
	this.CustomContent = customContent
	this.CustomUiAssets = customUiAssets
	this.PasswordPolicy = passwordPolicy
	this.Mfa = mfa
	this.SingleSignOnEnabled = singleSignOnEnabled
	this.SupportEmail = supportEmail
	this.SupportWebsiteUrl = supportWebsiteUrl
	this.UnknownSessionRedirectUrl = unknownSessionRedirectUrl
	return &this
}

// NewGetSignInExp200ResponseWithDefaults instantiates a new GetSignInExp200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignInExp200ResponseWithDefaults() *GetSignInExp200Response {
	this := GetSignInExp200Response{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *GetSignInExp200Response) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *GetSignInExp200Response) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *GetSignInExp200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetSignInExp200Response) SetId(v string) {
	o.Id = v
}

// GetColor returns the Color field value
func (o *GetSignInExp200Response) GetColor() GetSignInExp200ResponseColor {
	if o == nil {
		var ret GetSignInExp200ResponseColor
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetColorOk() (*GetSignInExp200ResponseColor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *GetSignInExp200Response) SetColor(v GetSignInExp200ResponseColor) {
	o.Color = v
}

// GetBranding returns the Branding field value
func (o *GetSignInExp200Response) GetBranding() ListApplicationOrganizations200ResponseInnerBranding {
	if o == nil {
		var ret ListApplicationOrganizations200ResponseInnerBranding
		return ret
	}

	return o.Branding
}

// GetBrandingOk returns a tuple with the Branding field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branding, true
}

// SetBranding sets field value
func (o *GetSignInExp200Response) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding) {
	o.Branding = v
}

// GetLanguageInfo returns the LanguageInfo field value
func (o *GetSignInExp200Response) GetLanguageInfo() GetSignInExp200ResponseLanguageInfo {
	if o == nil {
		var ret GetSignInExp200ResponseLanguageInfo
		return ret
	}

	return o.LanguageInfo
}

// GetLanguageInfoOk returns a tuple with the LanguageInfo field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetLanguageInfoOk() (*GetSignInExp200ResponseLanguageInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LanguageInfo, true
}

// SetLanguageInfo sets field value
func (o *GetSignInExp200Response) SetLanguageInfo(v GetSignInExp200ResponseLanguageInfo) {
	o.LanguageInfo = v
}

// GetTermsOfUseUrl returns the TermsOfUseUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetSignInExp200Response) GetTermsOfUseUrl() string {
	if o == nil || o.TermsOfUseUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.TermsOfUseUrl.Get()
}

// GetTermsOfUseUrlOk returns a tuple with the TermsOfUseUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSignInExp200Response) GetTermsOfUseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TermsOfUseUrl.Get(), o.TermsOfUseUrl.IsSet()
}

// SetTermsOfUseUrl sets field value
func (o *GetSignInExp200Response) SetTermsOfUseUrl(v string) {
	o.TermsOfUseUrl.Set(&v)
}

// GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetSignInExp200Response) GetPrivacyPolicyUrl() string {
	if o == nil || o.PrivacyPolicyUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.PrivacyPolicyUrl.Get()
}

// GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSignInExp200Response) GetPrivacyPolicyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivacyPolicyUrl.Get(), o.PrivacyPolicyUrl.IsSet()
}

// SetPrivacyPolicyUrl sets field value
func (o *GetSignInExp200Response) SetPrivacyPolicyUrl(v string) {
	o.PrivacyPolicyUrl.Set(&v)
}

// GetAgreeToTermsPolicy returns the AgreeToTermsPolicy field value
func (o *GetSignInExp200Response) GetAgreeToTermsPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgreeToTermsPolicy
}

// GetAgreeToTermsPolicyOk returns a tuple with the AgreeToTermsPolicy field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetAgreeToTermsPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgreeToTermsPolicy, true
}

// SetAgreeToTermsPolicy sets field value
func (o *GetSignInExp200Response) SetAgreeToTermsPolicy(v string) {
	o.AgreeToTermsPolicy = v
}

// GetSignIn returns the SignIn field value
func (o *GetSignInExp200Response) GetSignIn() GetSignInExp200ResponseSignIn {
	if o == nil {
		var ret GetSignInExp200ResponseSignIn
		return ret
	}

	return o.SignIn
}

// GetSignInOk returns a tuple with the SignIn field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetSignInOk() (*GetSignInExp200ResponseSignIn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignIn, true
}

// SetSignIn sets field value
func (o *GetSignInExp200Response) SetSignIn(v GetSignInExp200ResponseSignIn) {
	o.SignIn = v
}

// GetSignUp returns the SignUp field value
func (o *GetSignInExp200Response) GetSignUp() GetSignInExp200ResponseSignUp {
	if o == nil {
		var ret GetSignInExp200ResponseSignUp
		return ret
	}

	return o.SignUp
}

// GetSignUpOk returns a tuple with the SignUp field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetSignUpOk() (*GetSignInExp200ResponseSignUp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignUp, true
}

// SetSignUp sets field value
func (o *GetSignInExp200Response) SetSignUp(v GetSignInExp200ResponseSignUp) {
	o.SignUp = v
}

// GetSocialSignIn returns the SocialSignIn field value
func (o *GetSignInExp200Response) GetSocialSignIn() GetSignInExp200ResponseSocialSignIn {
	if o == nil {
		var ret GetSignInExp200ResponseSocialSignIn
		return ret
	}

	return o.SocialSignIn
}

// GetSocialSignInOk returns a tuple with the SocialSignIn field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetSocialSignInOk() (*GetSignInExp200ResponseSocialSignIn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SocialSignIn, true
}

// SetSocialSignIn sets field value
func (o *GetSignInExp200Response) SetSocialSignIn(v GetSignInExp200ResponseSocialSignIn) {
	o.SocialSignIn = v
}

// GetSocialSignInConnectorTargets returns the SocialSignInConnectorTargets field value
func (o *GetSignInExp200Response) GetSocialSignInConnectorTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SocialSignInConnectorTargets
}

// GetSocialSignInConnectorTargetsOk returns a tuple with the SocialSignInConnectorTargets field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetSocialSignInConnectorTargetsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SocialSignInConnectorTargets, true
}

// SetSocialSignInConnectorTargets sets field value
func (o *GetSignInExp200Response) SetSocialSignInConnectorTargets(v []string) {
	o.SocialSignInConnectorTargets = v
}

// GetSignInMode returns the SignInMode field value
func (o *GetSignInExp200Response) GetSignInMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignInMode
}

// GetSignInModeOk returns a tuple with the SignInMode field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetSignInModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignInMode, true
}

// SetSignInMode sets field value
func (o *GetSignInExp200Response) SetSignInMode(v string) {
	o.SignInMode = v
}

// GetCustomCss returns the CustomCss field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetSignInExp200Response) GetCustomCss() string {
	if o == nil || o.CustomCss.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustomCss.Get()
}

// GetCustomCssOk returns a tuple with the CustomCss field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSignInExp200Response) GetCustomCssOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomCss.Get(), o.CustomCss.IsSet()
}

// SetCustomCss sets field value
func (o *GetSignInExp200Response) SetCustomCss(v string) {
	o.CustomCss.Set(&v)
}

// GetCustomContent returns the CustomContent field value
func (o *GetSignInExp200Response) GetCustomContent() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.CustomContent
}

// GetCustomContentOk returns a tuple with the CustomContent field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetCustomContentOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomContent, true
}

// SetCustomContent sets field value
func (o *GetSignInExp200Response) SetCustomContent(v map[string]string) {
	o.CustomContent = v
}

// GetCustomUiAssets returns the CustomUiAssets field value
// If the value is explicit nil, the zero value for GetSignInExp200ResponseCustomUiAssets will be returned
func (o *GetSignInExp200Response) GetCustomUiAssets() GetSignInExp200ResponseCustomUiAssets {
	if o == nil || o.CustomUiAssets.Get() == nil {
		var ret GetSignInExp200ResponseCustomUiAssets
		return ret
	}

	return *o.CustomUiAssets.Get()
}

// GetCustomUiAssetsOk returns a tuple with the CustomUiAssets field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSignInExp200Response) GetCustomUiAssetsOk() (*GetSignInExp200ResponseCustomUiAssets, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomUiAssets.Get(), o.CustomUiAssets.IsSet()
}

// SetCustomUiAssets sets field value
func (o *GetSignInExp200Response) SetCustomUiAssets(v GetSignInExp200ResponseCustomUiAssets) {
	o.CustomUiAssets.Set(&v)
}

// GetPasswordPolicy returns the PasswordPolicy field value
func (o *GetSignInExp200Response) GetPasswordPolicy() GetSignInExp200ResponsePasswordPolicy {
	if o == nil {
		var ret GetSignInExp200ResponsePasswordPolicy
		return ret
	}

	return o.PasswordPolicy
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetPasswordPolicyOk() (*GetSignInExp200ResponsePasswordPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordPolicy, true
}

// SetPasswordPolicy sets field value
func (o *GetSignInExp200Response) SetPasswordPolicy(v GetSignInExp200ResponsePasswordPolicy) {
	o.PasswordPolicy = v
}

// GetMfa returns the Mfa field value
func (o *GetSignInExp200Response) GetMfa() GetSignInExp200ResponseMfa {
	if o == nil {
		var ret GetSignInExp200ResponseMfa
		return ret
	}

	return o.Mfa
}

// GetMfaOk returns a tuple with the Mfa field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetMfaOk() (*GetSignInExp200ResponseMfa, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mfa, true
}

// SetMfa sets field value
func (o *GetSignInExp200Response) SetMfa(v GetSignInExp200ResponseMfa) {
	o.Mfa = v
}

// GetSingleSignOnEnabled returns the SingleSignOnEnabled field value
func (o *GetSignInExp200Response) GetSingleSignOnEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SingleSignOnEnabled
}

// GetSingleSignOnEnabledOk returns a tuple with the SingleSignOnEnabled field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200Response) GetSingleSignOnEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SingleSignOnEnabled, true
}

// SetSingleSignOnEnabled sets field value
func (o *GetSignInExp200Response) SetSingleSignOnEnabled(v bool) {
	o.SingleSignOnEnabled = v
}

// GetSupportEmail returns the SupportEmail field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetSignInExp200Response) GetSupportEmail() string {
	if o == nil || o.SupportEmail.Get() == nil {
		var ret string
		return ret
	}

	return *o.SupportEmail.Get()
}

// GetSupportEmailOk returns a tuple with the SupportEmail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSignInExp200Response) GetSupportEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportEmail.Get(), o.SupportEmail.IsSet()
}

// SetSupportEmail sets field value
func (o *GetSignInExp200Response) SetSupportEmail(v string) {
	o.SupportEmail.Set(&v)
}

// GetSupportWebsiteUrl returns the SupportWebsiteUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetSignInExp200Response) GetSupportWebsiteUrl() string {
	if o == nil || o.SupportWebsiteUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.SupportWebsiteUrl.Get()
}

// GetSupportWebsiteUrlOk returns a tuple with the SupportWebsiteUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSignInExp200Response) GetSupportWebsiteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportWebsiteUrl.Get(), o.SupportWebsiteUrl.IsSet()
}

// SetSupportWebsiteUrl sets field value
func (o *GetSignInExp200Response) SetSupportWebsiteUrl(v string) {
	o.SupportWebsiteUrl.Set(&v)
}

// GetUnknownSessionRedirectUrl returns the UnknownSessionRedirectUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetSignInExp200Response) GetUnknownSessionRedirectUrl() string {
	if o == nil || o.UnknownSessionRedirectUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnknownSessionRedirectUrl.Get()
}

// GetUnknownSessionRedirectUrlOk returns a tuple with the UnknownSessionRedirectUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSignInExp200Response) GetUnknownSessionRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnknownSessionRedirectUrl.Get(), o.UnknownSessionRedirectUrl.IsSet()
}

// SetUnknownSessionRedirectUrl sets field value
func (o *GetSignInExp200Response) SetUnknownSessionRedirectUrl(v string) {
	o.UnknownSessionRedirectUrl.Set(&v)
}

func (o GetSignInExp200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSignInExp200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["color"] = o.Color
	toSerialize["branding"] = o.Branding
	toSerialize["languageInfo"] = o.LanguageInfo
	toSerialize["termsOfUseUrl"] = o.TermsOfUseUrl.Get()
	toSerialize["privacyPolicyUrl"] = o.PrivacyPolicyUrl.Get()
	toSerialize["agreeToTermsPolicy"] = o.AgreeToTermsPolicy
	toSerialize["signIn"] = o.SignIn
	toSerialize["signUp"] = o.SignUp
	toSerialize["socialSignIn"] = o.SocialSignIn
	toSerialize["socialSignInConnectorTargets"] = o.SocialSignInConnectorTargets
	toSerialize["signInMode"] = o.SignInMode
	toSerialize["customCss"] = o.CustomCss.Get()
	toSerialize["customContent"] = o.CustomContent
	toSerialize["customUiAssets"] = o.CustomUiAssets.Get()
	toSerialize["passwordPolicy"] = o.PasswordPolicy
	toSerialize["mfa"] = o.Mfa
	toSerialize["singleSignOnEnabled"] = o.SingleSignOnEnabled
	toSerialize["supportEmail"] = o.SupportEmail.Get()
	toSerialize["supportWebsiteUrl"] = o.SupportWebsiteUrl.Get()
	toSerialize["unknownSessionRedirectUrl"] = o.UnknownSessionRedirectUrl.Get()
	return toSerialize, nil
}

func (o *GetSignInExp200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"color",
		"branding",
		"languageInfo",
		"termsOfUseUrl",
		"privacyPolicyUrl",
		"agreeToTermsPolicy",
		"signIn",
		"signUp",
		"socialSignIn",
		"socialSignInConnectorTargets",
		"signInMode",
		"customCss",
		"customContent",
		"customUiAssets",
		"passwordPolicy",
		"mfa",
		"singleSignOnEnabled",
		"supportEmail",
		"supportWebsiteUrl",
		"unknownSessionRedirectUrl",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetSignInExp200Response := _GetSignInExp200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSignInExp200Response)

	if err != nil {
		return err
	}

	*o = GetSignInExp200Response(varGetSignInExp200Response)

	return err
}

type NullableGetSignInExp200Response struct {
	value *GetSignInExp200Response
	isSet bool
}

func (v NullableGetSignInExp200Response) Get() *GetSignInExp200Response {
	return v.value
}

func (v *NullableGetSignInExp200Response) Set(val *GetSignInExp200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignInExp200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignInExp200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignInExp200Response(val *GetSignInExp200Response) *NullableGetSignInExp200Response {
	return &NullableGetSignInExp200Response{value: val, isSet: true}
}

func (v NullableGetSignInExp200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignInExp200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


