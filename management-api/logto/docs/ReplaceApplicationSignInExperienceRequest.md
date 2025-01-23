# ReplaceApplicationSignInExperienceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to [**GetApplicationSignInExperience200ResponseColor**](GetApplicationSignInExperience200ResponseColor.md) |  | [optional] 
**Branding** | Pointer to [**ListApplicationOrganizations200ResponseInnerBranding**](ListApplicationOrganizations200ResponseInnerBranding.md) |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**TermsOfUseUrl** | [**ReplaceApplicationSignInExperienceRequestTermsOfUseUrl**](ReplaceApplicationSignInExperienceRequestTermsOfUseUrl.md) |  | 
**PrivacyPolicyUrl** | [**ReplaceApplicationSignInExperienceRequestTermsOfUseUrl**](ReplaceApplicationSignInExperienceRequestTermsOfUseUrl.md) |  | 

## Methods

### NewReplaceApplicationSignInExperienceRequest

`func NewReplaceApplicationSignInExperienceRequest(termsOfUseUrl ReplaceApplicationSignInExperienceRequestTermsOfUseUrl, privacyPolicyUrl ReplaceApplicationSignInExperienceRequestTermsOfUseUrl, ) *ReplaceApplicationSignInExperienceRequest`

NewReplaceApplicationSignInExperienceRequest instantiates a new ReplaceApplicationSignInExperienceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceApplicationSignInExperienceRequestWithDefaults

`func NewReplaceApplicationSignInExperienceRequestWithDefaults() *ReplaceApplicationSignInExperienceRequest`

NewReplaceApplicationSignInExperienceRequestWithDefaults instantiates a new ReplaceApplicationSignInExperienceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *ReplaceApplicationSignInExperienceRequest) GetColor() GetApplicationSignInExperience200ResponseColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ReplaceApplicationSignInExperienceRequest) GetColorOk() (*GetApplicationSignInExperience200ResponseColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ReplaceApplicationSignInExperienceRequest) SetColor(v GetApplicationSignInExperience200ResponseColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *ReplaceApplicationSignInExperienceRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetBranding

`func (o *ReplaceApplicationSignInExperienceRequest) GetBranding() ListApplicationOrganizations200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *ReplaceApplicationSignInExperienceRequest) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *ReplaceApplicationSignInExperienceRequest) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *ReplaceApplicationSignInExperienceRequest) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetDisplayName

`func (o *ReplaceApplicationSignInExperienceRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ReplaceApplicationSignInExperienceRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ReplaceApplicationSignInExperienceRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ReplaceApplicationSignInExperienceRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ReplaceApplicationSignInExperienceRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ReplaceApplicationSignInExperienceRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetTermsOfUseUrl

`func (o *ReplaceApplicationSignInExperienceRequest) GetTermsOfUseUrl() ReplaceApplicationSignInExperienceRequestTermsOfUseUrl`

GetTermsOfUseUrl returns the TermsOfUseUrl field if non-nil, zero value otherwise.

### GetTermsOfUseUrlOk

`func (o *ReplaceApplicationSignInExperienceRequest) GetTermsOfUseUrlOk() (*ReplaceApplicationSignInExperienceRequestTermsOfUseUrl, bool)`

GetTermsOfUseUrlOk returns a tuple with the TermsOfUseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUseUrl

`func (o *ReplaceApplicationSignInExperienceRequest) SetTermsOfUseUrl(v ReplaceApplicationSignInExperienceRequestTermsOfUseUrl)`

SetTermsOfUseUrl sets TermsOfUseUrl field to given value.


### GetPrivacyPolicyUrl

`func (o *ReplaceApplicationSignInExperienceRequest) GetPrivacyPolicyUrl() ReplaceApplicationSignInExperienceRequestTermsOfUseUrl`

GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetPrivacyPolicyUrlOk

`func (o *ReplaceApplicationSignInExperienceRequest) GetPrivacyPolicyUrlOk() (*ReplaceApplicationSignInExperienceRequestTermsOfUseUrl, bool)`

GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyUrl

`func (o *ReplaceApplicationSignInExperienceRequest) SetPrivacyPolicyUrl(v ReplaceApplicationSignInExperienceRequestTermsOfUseUrl)`

SetPrivacyPolicyUrl sets PrivacyPolicyUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


