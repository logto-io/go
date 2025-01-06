# UpdateOtherProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FamilyName** | Pointer to **string** | The new family name for the user. | [optional] 
**GivenName** | Pointer to **string** | The new given name for the user. | [optional] 
**MiddleName** | Pointer to **string** | The new middle name for the user. | [optional] 
**Nickname** | Pointer to **string** | The new nickname for the user. | [optional] 
**PreferredUsername** | Pointer to **string** | The new preferred username for the user. | [optional] 
**Profile** | Pointer to **string** | The new profile for the user. | [optional] 
**Website** | Pointer to **string** | The new website for the user. | [optional] 
**Gender** | Pointer to **string** | The new gender for the user. | [optional] 
**Birthdate** | Pointer to **string** | The new birthdate for the user. | [optional] 
**Zoneinfo** | Pointer to **string** | The new zoneinfo for the user. | [optional] 
**Locale** | Pointer to **string** | The new locale for the user. | [optional] 
**Address** | Pointer to [**UpdateOtherProfileRequestAddress**](UpdateOtherProfileRequestAddress.md) |  | [optional] 

## Methods

### NewUpdateOtherProfileRequest

`func NewUpdateOtherProfileRequest() *UpdateOtherProfileRequest`

NewUpdateOtherProfileRequest instantiates a new UpdateOtherProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOtherProfileRequestWithDefaults

`func NewUpdateOtherProfileRequestWithDefaults() *UpdateOtherProfileRequest`

NewUpdateOtherProfileRequestWithDefaults instantiates a new UpdateOtherProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamilyName

`func (o *UpdateOtherProfileRequest) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *UpdateOtherProfileRequest) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *UpdateOtherProfileRequest) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *UpdateOtherProfileRequest) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetGivenName

`func (o *UpdateOtherProfileRequest) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *UpdateOtherProfileRequest) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *UpdateOtherProfileRequest) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *UpdateOtherProfileRequest) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetMiddleName

`func (o *UpdateOtherProfileRequest) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *UpdateOtherProfileRequest) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *UpdateOtherProfileRequest) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *UpdateOtherProfileRequest) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetNickname

`func (o *UpdateOtherProfileRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *UpdateOtherProfileRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *UpdateOtherProfileRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *UpdateOtherProfileRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPreferredUsername

`func (o *UpdateOtherProfileRequest) GetPreferredUsername() string`

GetPreferredUsername returns the PreferredUsername field if non-nil, zero value otherwise.

### GetPreferredUsernameOk

`func (o *UpdateOtherProfileRequest) GetPreferredUsernameOk() (*string, bool)`

GetPreferredUsernameOk returns a tuple with the PreferredUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUsername

`func (o *UpdateOtherProfileRequest) SetPreferredUsername(v string)`

SetPreferredUsername sets PreferredUsername field to given value.

### HasPreferredUsername

`func (o *UpdateOtherProfileRequest) HasPreferredUsername() bool`

HasPreferredUsername returns a boolean if a field has been set.

### GetProfile

`func (o *UpdateOtherProfileRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UpdateOtherProfileRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UpdateOtherProfileRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UpdateOtherProfileRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetWebsite

`func (o *UpdateOtherProfileRequest) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *UpdateOtherProfileRequest) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *UpdateOtherProfileRequest) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *UpdateOtherProfileRequest) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetGender

`func (o *UpdateOtherProfileRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UpdateOtherProfileRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UpdateOtherProfileRequest) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UpdateOtherProfileRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetBirthdate

`func (o *UpdateOtherProfileRequest) GetBirthdate() string`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *UpdateOtherProfileRequest) GetBirthdateOk() (*string, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *UpdateOtherProfileRequest) SetBirthdate(v string)`

SetBirthdate sets Birthdate field to given value.

### HasBirthdate

`func (o *UpdateOtherProfileRequest) HasBirthdate() bool`

HasBirthdate returns a boolean if a field has been set.

### GetZoneinfo

`func (o *UpdateOtherProfileRequest) GetZoneinfo() string`

GetZoneinfo returns the Zoneinfo field if non-nil, zero value otherwise.

### GetZoneinfoOk

`func (o *UpdateOtherProfileRequest) GetZoneinfoOk() (*string, bool)`

GetZoneinfoOk returns a tuple with the Zoneinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneinfo

`func (o *UpdateOtherProfileRequest) SetZoneinfo(v string)`

SetZoneinfo sets Zoneinfo field to given value.

### HasZoneinfo

`func (o *UpdateOtherProfileRequest) HasZoneinfo() bool`

HasZoneinfo returns a boolean if a field has been set.

### GetLocale

`func (o *UpdateOtherProfileRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UpdateOtherProfileRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UpdateOtherProfileRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UpdateOtherProfileRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetAddress

`func (o *UpdateOtherProfileRequest) GetAddress() UpdateOtherProfileRequestAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateOtherProfileRequest) GetAddressOk() (*UpdateOtherProfileRequestAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateOtherProfileRequest) SetAddress(v UpdateOtherProfileRequestAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateOtherProfileRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


