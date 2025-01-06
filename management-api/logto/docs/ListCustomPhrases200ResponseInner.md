# ListCustomPhrases200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**LanguageTag** | **string** |  | 
**Translation** | [**TranslationObject**](TranslationObject.md) |  | 

## Methods

### NewListCustomPhrases200ResponseInner

`func NewListCustomPhrases200ResponseInner(tenantId string, id string, languageTag string, translation TranslationObject, ) *ListCustomPhrases200ResponseInner`

NewListCustomPhrases200ResponseInner instantiates a new ListCustomPhrases200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomPhrases200ResponseInnerWithDefaults

`func NewListCustomPhrases200ResponseInnerWithDefaults() *ListCustomPhrases200ResponseInner`

NewListCustomPhrases200ResponseInnerWithDefaults instantiates a new ListCustomPhrases200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListCustomPhrases200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListCustomPhrases200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListCustomPhrases200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListCustomPhrases200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCustomPhrases200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCustomPhrases200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetLanguageTag

`func (o *ListCustomPhrases200ResponseInner) GetLanguageTag() string`

GetLanguageTag returns the LanguageTag field if non-nil, zero value otherwise.

### GetLanguageTagOk

`func (o *ListCustomPhrases200ResponseInner) GetLanguageTagOk() (*string, bool)`

GetLanguageTagOk returns a tuple with the LanguageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageTag

`func (o *ListCustomPhrases200ResponseInner) SetLanguageTag(v string)`

SetLanguageTag sets LanguageTag field to given value.


### GetTranslation

`func (o *ListCustomPhrases200ResponseInner) GetTranslation() TranslationObject`

GetTranslation returns the Translation field if non-nil, zero value otherwise.

### GetTranslationOk

`func (o *ListCustomPhrases200ResponseInner) GetTranslationOk() (*TranslationObject, bool)`

GetTranslationOk returns a tuple with the Translation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslation

`func (o *ListCustomPhrases200ResponseInner) SetTranslation(v TranslationObject)`

SetTranslation sets Translation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


