# GetApplicationSignInExperience200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**ApplicationId** | **string** |  | 
**Color** | [**GetApplicationSignInExperience200ResponseColor**](GetApplicationSignInExperience200ResponseColor.md) |  | 
**Branding** | [**ListApplicationOrganizations200ResponseInnerBranding**](ListApplicationOrganizations200ResponseInnerBranding.md) |  | 
**TermsOfUseUrl** | **NullableString** |  | 
**PrivacyPolicyUrl** | **NullableString** |  | 
**DisplayName** | **NullableString** |  | 

## Methods

### NewGetApplicationSignInExperience200Response

`func NewGetApplicationSignInExperience200Response(tenantId string, applicationId string, color GetApplicationSignInExperience200ResponseColor, branding ListApplicationOrganizations200ResponseInnerBranding, termsOfUseUrl NullableString, privacyPolicyUrl NullableString, displayName NullableString, ) *GetApplicationSignInExperience200Response`

NewGetApplicationSignInExperience200Response instantiates a new GetApplicationSignInExperience200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApplicationSignInExperience200ResponseWithDefaults

`func NewGetApplicationSignInExperience200ResponseWithDefaults() *GetApplicationSignInExperience200Response`

NewGetApplicationSignInExperience200ResponseWithDefaults instantiates a new GetApplicationSignInExperience200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *GetApplicationSignInExperience200Response) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GetApplicationSignInExperience200Response) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GetApplicationSignInExperience200Response) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetApplicationId

`func (o *GetApplicationSignInExperience200Response) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GetApplicationSignInExperience200Response) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GetApplicationSignInExperience200Response) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetColor

`func (o *GetApplicationSignInExperience200Response) GetColor() GetApplicationSignInExperience200ResponseColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *GetApplicationSignInExperience200Response) GetColorOk() (*GetApplicationSignInExperience200ResponseColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *GetApplicationSignInExperience200Response) SetColor(v GetApplicationSignInExperience200ResponseColor)`

SetColor sets Color field to given value.


### GetBranding

`func (o *GetApplicationSignInExperience200Response) GetBranding() ListApplicationOrganizations200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *GetApplicationSignInExperience200Response) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *GetApplicationSignInExperience200Response) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding)`

SetBranding sets Branding field to given value.


### GetTermsOfUseUrl

`func (o *GetApplicationSignInExperience200Response) GetTermsOfUseUrl() string`

GetTermsOfUseUrl returns the TermsOfUseUrl field if non-nil, zero value otherwise.

### GetTermsOfUseUrlOk

`func (o *GetApplicationSignInExperience200Response) GetTermsOfUseUrlOk() (*string, bool)`

GetTermsOfUseUrlOk returns a tuple with the TermsOfUseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUseUrl

`func (o *GetApplicationSignInExperience200Response) SetTermsOfUseUrl(v string)`

SetTermsOfUseUrl sets TermsOfUseUrl field to given value.


### SetTermsOfUseUrlNil

`func (o *GetApplicationSignInExperience200Response) SetTermsOfUseUrlNil(b bool)`

 SetTermsOfUseUrlNil sets the value for TermsOfUseUrl to be an explicit nil

### UnsetTermsOfUseUrl
`func (o *GetApplicationSignInExperience200Response) UnsetTermsOfUseUrl()`

UnsetTermsOfUseUrl ensures that no value is present for TermsOfUseUrl, not even an explicit nil
### GetPrivacyPolicyUrl

`func (o *GetApplicationSignInExperience200Response) GetPrivacyPolicyUrl() string`

GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetPrivacyPolicyUrlOk

`func (o *GetApplicationSignInExperience200Response) GetPrivacyPolicyUrlOk() (*string, bool)`

GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyUrl

`func (o *GetApplicationSignInExperience200Response) SetPrivacyPolicyUrl(v string)`

SetPrivacyPolicyUrl sets PrivacyPolicyUrl field to given value.


### SetPrivacyPolicyUrlNil

`func (o *GetApplicationSignInExperience200Response) SetPrivacyPolicyUrlNil(b bool)`

 SetPrivacyPolicyUrlNil sets the value for PrivacyPolicyUrl to be an explicit nil

### UnsetPrivacyPolicyUrl
`func (o *GetApplicationSignInExperience200Response) UnsetPrivacyPolicyUrl()`

UnsetPrivacyPolicyUrl ensures that no value is present for PrivacyPolicyUrl, not even an explicit nil
### GetDisplayName

`func (o *GetApplicationSignInExperience200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetApplicationSignInExperience200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetApplicationSignInExperience200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *GetApplicationSignInExperience200Response) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *GetApplicationSignInExperience200Response) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


