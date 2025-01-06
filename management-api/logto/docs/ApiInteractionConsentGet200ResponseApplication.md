# ApiInteractionConsentGet200ResponseApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Branding** | Pointer to [**ListApplicationOrganizations200ResponseInnerBranding**](ListApplicationOrganizations200ResponseInnerBranding.md) |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**PrivacyPolicyUrl** | Pointer to **NullableString** |  | [optional] 
**TermsOfUseUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApiInteractionConsentGet200ResponseApplication

`func NewApiInteractionConsentGet200ResponseApplication(id string, name string, ) *ApiInteractionConsentGet200ResponseApplication`

NewApiInteractionConsentGet200ResponseApplication instantiates a new ApiInteractionConsentGet200ResponseApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInteractionConsentGet200ResponseApplicationWithDefaults

`func NewApiInteractionConsentGet200ResponseApplicationWithDefaults() *ApiInteractionConsentGet200ResponseApplication`

NewApiInteractionConsentGet200ResponseApplicationWithDefaults instantiates a new ApiInteractionConsentGet200ResponseApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiInteractionConsentGet200ResponseApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiInteractionConsentGet200ResponseApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiInteractionConsentGet200ResponseApplication) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiInteractionConsentGet200ResponseApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiInteractionConsentGet200ResponseApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiInteractionConsentGet200ResponseApplication) SetName(v string)`

SetName sets Name field to given value.


### GetBranding

`func (o *ApiInteractionConsentGet200ResponseApplication) GetBranding() ListApplicationOrganizations200ResponseInnerBranding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *ApiInteractionConsentGet200ResponseApplication) GetBrandingOk() (*ListApplicationOrganizations200ResponseInnerBranding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *ApiInteractionConsentGet200ResponseApplication) SetBranding(v ListApplicationOrganizations200ResponseInnerBranding)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *ApiInteractionConsentGet200ResponseApplication) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiInteractionConsentGet200ResponseApplication) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiInteractionConsentGet200ResponseApplication) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiInteractionConsentGet200ResponseApplication) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiInteractionConsentGet200ResponseApplication) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ApiInteractionConsentGet200ResponseApplication) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ApiInteractionConsentGet200ResponseApplication) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetPrivacyPolicyUrl

`func (o *ApiInteractionConsentGet200ResponseApplication) GetPrivacyPolicyUrl() string`

GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetPrivacyPolicyUrlOk

`func (o *ApiInteractionConsentGet200ResponseApplication) GetPrivacyPolicyUrlOk() (*string, bool)`

GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyUrl

`func (o *ApiInteractionConsentGet200ResponseApplication) SetPrivacyPolicyUrl(v string)`

SetPrivacyPolicyUrl sets PrivacyPolicyUrl field to given value.

### HasPrivacyPolicyUrl

`func (o *ApiInteractionConsentGet200ResponseApplication) HasPrivacyPolicyUrl() bool`

HasPrivacyPolicyUrl returns a boolean if a field has been set.

### SetPrivacyPolicyUrlNil

`func (o *ApiInteractionConsentGet200ResponseApplication) SetPrivacyPolicyUrlNil(b bool)`

 SetPrivacyPolicyUrlNil sets the value for PrivacyPolicyUrl to be an explicit nil

### UnsetPrivacyPolicyUrl
`func (o *ApiInteractionConsentGet200ResponseApplication) UnsetPrivacyPolicyUrl()`

UnsetPrivacyPolicyUrl ensures that no value is present for PrivacyPolicyUrl, not even an explicit nil
### GetTermsOfUseUrl

`func (o *ApiInteractionConsentGet200ResponseApplication) GetTermsOfUseUrl() string`

GetTermsOfUseUrl returns the TermsOfUseUrl field if non-nil, zero value otherwise.

### GetTermsOfUseUrlOk

`func (o *ApiInteractionConsentGet200ResponseApplication) GetTermsOfUseUrlOk() (*string, bool)`

GetTermsOfUseUrlOk returns a tuple with the TermsOfUseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUseUrl

`func (o *ApiInteractionConsentGet200ResponseApplication) SetTermsOfUseUrl(v string)`

SetTermsOfUseUrl sets TermsOfUseUrl field to given value.

### HasTermsOfUseUrl

`func (o *ApiInteractionConsentGet200ResponseApplication) HasTermsOfUseUrl() bool`

HasTermsOfUseUrl returns a boolean if a field has been set.

### SetTermsOfUseUrlNil

`func (o *ApiInteractionConsentGet200ResponseApplication) SetTermsOfUseUrlNil(b bool)`

 SetTermsOfUseUrlNil sets the value for TermsOfUseUrl to be an explicit nil

### UnsetTermsOfUseUrl
`func (o *ApiInteractionConsentGet200ResponseApplication) UnsetTermsOfUseUrl()`

UnsetTermsOfUseUrl ensures that no value is present for TermsOfUseUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


