# ApiInteractionConsentGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | [**ApiInteractionConsentGet200ResponseApplication**](ApiInteractionConsentGet200ResponseApplication.md) |  | 
**User** | [**ApiInteractionConsentGet200ResponseUser**](ApiInteractionConsentGet200ResponseUser.md) |  | 
**Organizations** | Pointer to [**[]ApiInteractionConsentGet200ResponseOrganizationsInner**](ApiInteractionConsentGet200ResponseOrganizationsInner.md) |  | [optional] 
**MissingOIDCScope** | Pointer to **[]string** |  | [optional] 
**MissingResourceScopes** | Pointer to [**[]ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner**](ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner.md) |  | [optional] 
**RedirectUri** | **string** |  | 

## Methods

### NewApiInteractionConsentGet200Response

`func NewApiInteractionConsentGet200Response(application ApiInteractionConsentGet200ResponseApplication, user ApiInteractionConsentGet200ResponseUser, redirectUri string, ) *ApiInteractionConsentGet200Response`

NewApiInteractionConsentGet200Response instantiates a new ApiInteractionConsentGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInteractionConsentGet200ResponseWithDefaults

`func NewApiInteractionConsentGet200ResponseWithDefaults() *ApiInteractionConsentGet200Response`

NewApiInteractionConsentGet200ResponseWithDefaults instantiates a new ApiInteractionConsentGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ApiInteractionConsentGet200Response) GetApplication() ApiInteractionConsentGet200ResponseApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApiInteractionConsentGet200Response) GetApplicationOk() (*ApiInteractionConsentGet200ResponseApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApiInteractionConsentGet200Response) SetApplication(v ApiInteractionConsentGet200ResponseApplication)`

SetApplication sets Application field to given value.


### GetUser

`func (o *ApiInteractionConsentGet200Response) GetUser() ApiInteractionConsentGet200ResponseUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApiInteractionConsentGet200Response) GetUserOk() (*ApiInteractionConsentGet200ResponseUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApiInteractionConsentGet200Response) SetUser(v ApiInteractionConsentGet200ResponseUser)`

SetUser sets User field to given value.


### GetOrganizations

`func (o *ApiInteractionConsentGet200Response) GetOrganizations() []ApiInteractionConsentGet200ResponseOrganizationsInner`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *ApiInteractionConsentGet200Response) GetOrganizationsOk() (*[]ApiInteractionConsentGet200ResponseOrganizationsInner, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *ApiInteractionConsentGet200Response) SetOrganizations(v []ApiInteractionConsentGet200ResponseOrganizationsInner)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *ApiInteractionConsentGet200Response) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetMissingOIDCScope

`func (o *ApiInteractionConsentGet200Response) GetMissingOIDCScope() []string`

GetMissingOIDCScope returns the MissingOIDCScope field if non-nil, zero value otherwise.

### GetMissingOIDCScopeOk

`func (o *ApiInteractionConsentGet200Response) GetMissingOIDCScopeOk() (*[]string, bool)`

GetMissingOIDCScopeOk returns a tuple with the MissingOIDCScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingOIDCScope

`func (o *ApiInteractionConsentGet200Response) SetMissingOIDCScope(v []string)`

SetMissingOIDCScope sets MissingOIDCScope field to given value.

### HasMissingOIDCScope

`func (o *ApiInteractionConsentGet200Response) HasMissingOIDCScope() bool`

HasMissingOIDCScope returns a boolean if a field has been set.

### GetMissingResourceScopes

`func (o *ApiInteractionConsentGet200Response) GetMissingResourceScopes() []ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner`

GetMissingResourceScopes returns the MissingResourceScopes field if non-nil, zero value otherwise.

### GetMissingResourceScopesOk

`func (o *ApiInteractionConsentGet200Response) GetMissingResourceScopesOk() (*[]ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner, bool)`

GetMissingResourceScopesOk returns a tuple with the MissingResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingResourceScopes

`func (o *ApiInteractionConsentGet200Response) SetMissingResourceScopes(v []ApiInteractionConsentGet200ResponseOrganizationsInnerMissingResourceScopesInner)`

SetMissingResourceScopes sets MissingResourceScopes field to given value.

### HasMissingResourceScopes

`func (o *ApiInteractionConsentGet200Response) HasMissingResourceScopes() bool`

HasMissingResourceScopes returns a boolean if a field has been set.

### GetRedirectUri

`func (o *ApiInteractionConsentGet200Response) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *ApiInteractionConsentGet200Response) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *ApiInteractionConsentGet200Response) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


