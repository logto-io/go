# ListJwtCustomizers200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | [**GetJwtCustomizer200ResponseOneOf1**](GetJwtCustomizer200ResponseOneOf1.md) |  | 

## Methods

### NewListJwtCustomizers200ResponseInner

`func NewListJwtCustomizers200ResponseInner(key string, value GetJwtCustomizer200ResponseOneOf1, ) *ListJwtCustomizers200ResponseInner`

NewListJwtCustomizers200ResponseInner instantiates a new ListJwtCustomizers200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListJwtCustomizers200ResponseInnerWithDefaults

`func NewListJwtCustomizers200ResponseInnerWithDefaults() *ListJwtCustomizers200ResponseInner`

NewListJwtCustomizers200ResponseInnerWithDefaults instantiates a new ListJwtCustomizers200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ListJwtCustomizers200ResponseInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ListJwtCustomizers200ResponseInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ListJwtCustomizers200ResponseInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *ListJwtCustomizers200ResponseInner) GetValue() GetJwtCustomizer200ResponseOneOf1`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListJwtCustomizers200ResponseInner) GetValueOk() (*GetJwtCustomizer200ResponseOneOf1, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListJwtCustomizers200ResponseInner) SetValue(v GetJwtCustomizer200ResponseOneOf1)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


