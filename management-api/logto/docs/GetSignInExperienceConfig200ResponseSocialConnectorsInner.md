# GetSignInExperienceConfig200ResponseSocialConnectorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Target** | **string** |  | 
**Name** | **map[string]interface{}** | Validator function | 
**Logo** | **string** |  | 
**LogoDark** | **NullableString** |  | 
**FromEmail** | Pointer to **string** |  | [optional] 
**Platform** | **NullableString** |  | 
**IsStandard** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetSignInExperienceConfig200ResponseSocialConnectorsInner

`func NewGetSignInExperienceConfig200ResponseSocialConnectorsInner(id string, target string, name map[string]interface{}, logo string, logoDark NullableString, platform NullableString, ) *GetSignInExperienceConfig200ResponseSocialConnectorsInner`

NewGetSignInExperienceConfig200ResponseSocialConnectorsInner instantiates a new GetSignInExperienceConfig200ResponseSocialConnectorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignInExperienceConfig200ResponseSocialConnectorsInnerWithDefaults

`func NewGetSignInExperienceConfig200ResponseSocialConnectorsInnerWithDefaults() *GetSignInExperienceConfig200ResponseSocialConnectorsInner`

NewGetSignInExperienceConfig200ResponseSocialConnectorsInnerWithDefaults instantiates a new GetSignInExperienceConfig200ResponseSocialConnectorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) SetId(v string)`

SetId sets Id field to given value.


### GetTarget

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetName

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetName() map[string]interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetNameOk() (*map[string]interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) SetName(v map[string]interface{})`

SetName sets Name field to given value.


### GetLogo

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) SetLogo(v string)`

SetLogo sets Logo field to given value.


### GetLogoDark

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetLogoDark() string`

GetLogoDark returns the LogoDark field if non-nil, zero value otherwise.

### GetLogoDarkOk

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetLogoDarkOk() (*string, bool)`

GetLogoDarkOk returns a tuple with the LogoDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoDark

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) SetLogoDark(v string)`

SetLogoDark sets LogoDark field to given value.


### SetLogoDarkNil

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) SetLogoDarkNil(b bool)`

 SetLogoDarkNil sets the value for LogoDark to be an explicit nil

### UnsetLogoDark
`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) UnsetLogoDark()`

UnsetLogoDark ensures that no value is present for LogoDark, not even an explicit nil
### GetFromEmail

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetPlatform

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### SetPlatformNil

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetIsStandard

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetIsStandard() bool`

GetIsStandard returns the IsStandard field if non-nil, zero value otherwise.

### GetIsStandardOk

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) GetIsStandardOk() (*bool, bool)`

GetIsStandardOk returns a tuple with the IsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandard

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) SetIsStandard(v bool)`

SetIsStandard sets IsStandard field to given value.

### HasIsStandard

`func (o *GetSignInExperienceConfig200ResponseSocialConnectorsInner) HasIsStandard() bool`

HasIsStandard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


