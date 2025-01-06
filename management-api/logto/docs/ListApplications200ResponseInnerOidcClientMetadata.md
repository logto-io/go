# ListApplications200ResponseInnerOidcClientMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUris** | [**[]ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner**](ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner.md) |  | 
**PostLogoutRedirectUris** | **[]string** |  | 
**BackchannelLogoutUri** | Pointer to **string** |  | [optional] 
**BackchannelLogoutSessionRequired** | Pointer to **bool** |  | [optional] 
**LogoUri** | Pointer to **string** |  | [optional] 

## Methods

### NewListApplications200ResponseInnerOidcClientMetadata

`func NewListApplications200ResponseInnerOidcClientMetadata(redirectUris []ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner, postLogoutRedirectUris []string, ) *ListApplications200ResponseInnerOidcClientMetadata`

NewListApplications200ResponseInnerOidcClientMetadata instantiates a new ListApplications200ResponseInnerOidcClientMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApplications200ResponseInnerOidcClientMetadataWithDefaults

`func NewListApplications200ResponseInnerOidcClientMetadataWithDefaults() *ListApplications200ResponseInnerOidcClientMetadata`

NewListApplications200ResponseInnerOidcClientMetadataWithDefaults instantiates a new ListApplications200ResponseInnerOidcClientMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUris

`func (o *ListApplications200ResponseInnerOidcClientMetadata) GetRedirectUris() []ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ListApplications200ResponseInnerOidcClientMetadata) GetRedirectUrisOk() (*[]ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ListApplications200ResponseInnerOidcClientMetadata) SetRedirectUris(v []ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner)`

SetRedirectUris sets RedirectUris field to given value.


### GetPostLogoutRedirectUris

`func (o *ListApplications200ResponseInnerOidcClientMetadata) GetPostLogoutRedirectUris() []string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *ListApplications200ResponseInnerOidcClientMetadata) GetPostLogoutRedirectUrisOk() (*[]string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *ListApplications200ResponseInnerOidcClientMetadata) SetPostLogoutRedirectUris(v []string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.


### GetBackchannelLogoutUri

`func (o *ListApplications200ResponseInnerOidcClientMetadata) GetBackchannelLogoutUri() string`

GetBackchannelLogoutUri returns the BackchannelLogoutUri field if non-nil, zero value otherwise.

### GetBackchannelLogoutUriOk

`func (o *ListApplications200ResponseInnerOidcClientMetadata) GetBackchannelLogoutUriOk() (*string, bool)`

GetBackchannelLogoutUriOk returns a tuple with the BackchannelLogoutUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelLogoutUri

`func (o *ListApplications200ResponseInnerOidcClientMetadata) SetBackchannelLogoutUri(v string)`

SetBackchannelLogoutUri sets BackchannelLogoutUri field to given value.

### HasBackchannelLogoutUri

`func (o *ListApplications200ResponseInnerOidcClientMetadata) HasBackchannelLogoutUri() bool`

HasBackchannelLogoutUri returns a boolean if a field has been set.

### GetBackchannelLogoutSessionRequired

`func (o *ListApplications200ResponseInnerOidcClientMetadata) GetBackchannelLogoutSessionRequired() bool`

GetBackchannelLogoutSessionRequired returns the BackchannelLogoutSessionRequired field if non-nil, zero value otherwise.

### GetBackchannelLogoutSessionRequiredOk

`func (o *ListApplications200ResponseInnerOidcClientMetadata) GetBackchannelLogoutSessionRequiredOk() (*bool, bool)`

GetBackchannelLogoutSessionRequiredOk returns a tuple with the BackchannelLogoutSessionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelLogoutSessionRequired

`func (o *ListApplications200ResponseInnerOidcClientMetadata) SetBackchannelLogoutSessionRequired(v bool)`

SetBackchannelLogoutSessionRequired sets BackchannelLogoutSessionRequired field to given value.

### HasBackchannelLogoutSessionRequired

`func (o *ListApplications200ResponseInnerOidcClientMetadata) HasBackchannelLogoutSessionRequired() bool`

HasBackchannelLogoutSessionRequired returns a boolean if a field has been set.

### GetLogoUri

`func (o *ListApplications200ResponseInnerOidcClientMetadata) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *ListApplications200ResponseInnerOidcClientMetadata) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *ListApplications200ResponseInnerOidcClientMetadata) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.

### HasLogoUri

`func (o *ListApplications200ResponseInnerOidcClientMetadata) HasLogoUri() bool`

HasLogoUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


