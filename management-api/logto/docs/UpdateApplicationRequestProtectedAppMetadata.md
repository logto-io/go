# UpdateApplicationRequestProtectedAppMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | Pointer to **string** |  | [optional] 
**SessionDuration** | Pointer to **float32** |  | [optional] 
**PageRules** | Pointer to [**[]ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner**](ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner.md) |  | [optional] 

## Methods

### NewUpdateApplicationRequestProtectedAppMetadata

`func NewUpdateApplicationRequestProtectedAppMetadata() *UpdateApplicationRequestProtectedAppMetadata`

NewUpdateApplicationRequestProtectedAppMetadata instantiates a new UpdateApplicationRequestProtectedAppMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApplicationRequestProtectedAppMetadataWithDefaults

`func NewUpdateApplicationRequestProtectedAppMetadataWithDefaults() *UpdateApplicationRequestProtectedAppMetadata`

NewUpdateApplicationRequestProtectedAppMetadataWithDefaults instantiates a new UpdateApplicationRequestProtectedAppMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *UpdateApplicationRequestProtectedAppMetadata) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *UpdateApplicationRequestProtectedAppMetadata) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *UpdateApplicationRequestProtectedAppMetadata) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *UpdateApplicationRequestProtectedAppMetadata) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetSessionDuration

`func (o *UpdateApplicationRequestProtectedAppMetadata) GetSessionDuration() float32`

GetSessionDuration returns the SessionDuration field if non-nil, zero value otherwise.

### GetSessionDurationOk

`func (o *UpdateApplicationRequestProtectedAppMetadata) GetSessionDurationOk() (*float32, bool)`

GetSessionDurationOk returns a tuple with the SessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDuration

`func (o *UpdateApplicationRequestProtectedAppMetadata) SetSessionDuration(v float32)`

SetSessionDuration sets SessionDuration field to given value.

### HasSessionDuration

`func (o *UpdateApplicationRequestProtectedAppMetadata) HasSessionDuration() bool`

HasSessionDuration returns a boolean if a field has been set.

### GetPageRules

`func (o *UpdateApplicationRequestProtectedAppMetadata) GetPageRules() []ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner`

GetPageRules returns the PageRules field if non-nil, zero value otherwise.

### GetPageRulesOk

`func (o *UpdateApplicationRequestProtectedAppMetadata) GetPageRulesOk() (*[]ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner, bool)`

GetPageRulesOk returns a tuple with the PageRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageRules

`func (o *UpdateApplicationRequestProtectedAppMetadata) SetPageRules(v []ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner)`

SetPageRules sets PageRules field to given value.

### HasPageRules

`func (o *UpdateApplicationRequestProtectedAppMetadata) HasPageRules() bool`

HasPageRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


