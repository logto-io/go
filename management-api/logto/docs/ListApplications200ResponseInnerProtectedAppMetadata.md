# ListApplications200ResponseInnerProtectedAppMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Origin** | **string** |  | 
**SessionDuration** | **float32** |  | 
**PageRules** | [**[]ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner**](ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner.md) |  | 
**CustomDomains** | Pointer to [**[]ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner**](ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner.md) |  | [optional] 

## Methods

### NewListApplications200ResponseInnerProtectedAppMetadata

`func NewListApplications200ResponseInnerProtectedAppMetadata(host string, origin string, sessionDuration float32, pageRules []ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner, ) *ListApplications200ResponseInnerProtectedAppMetadata`

NewListApplications200ResponseInnerProtectedAppMetadata instantiates a new ListApplications200ResponseInnerProtectedAppMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApplications200ResponseInnerProtectedAppMetadataWithDefaults

`func NewListApplications200ResponseInnerProtectedAppMetadataWithDefaults() *ListApplications200ResponseInnerProtectedAppMetadata`

NewListApplications200ResponseInnerProtectedAppMetadataWithDefaults instantiates a new ListApplications200ResponseInnerProtectedAppMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) SetHost(v string)`

SetHost sets Host field to given value.


### GetOrigin

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetSessionDuration

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) GetSessionDuration() float32`

GetSessionDuration returns the SessionDuration field if non-nil, zero value otherwise.

### GetSessionDurationOk

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) GetSessionDurationOk() (*float32, bool)`

GetSessionDurationOk returns a tuple with the SessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDuration

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) SetSessionDuration(v float32)`

SetSessionDuration sets SessionDuration field to given value.


### GetPageRules

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) GetPageRules() []ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner`

GetPageRules returns the PageRules field if non-nil, zero value otherwise.

### GetPageRulesOk

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) GetPageRulesOk() (*[]ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner, bool)`

GetPageRulesOk returns a tuple with the PageRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageRules

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) SetPageRules(v []ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner)`

SetPageRules sets PageRules field to given value.


### GetCustomDomains

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) GetCustomDomains() []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner`

GetCustomDomains returns the CustomDomains field if non-nil, zero value otherwise.

### GetCustomDomainsOk

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) GetCustomDomainsOk() (*[]ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner, bool)`

GetCustomDomainsOk returns a tuple with the CustomDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomains

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) SetCustomDomains(v []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner)`

SetCustomDomains sets CustomDomains field to given value.

### HasCustomDomains

`func (o *ListApplications200ResponseInnerProtectedAppMetadata) HasCustomDomains() bool`

HasCustomDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


