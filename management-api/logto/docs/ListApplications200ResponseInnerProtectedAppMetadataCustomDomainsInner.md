# ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | 
**Status** | **string** |  | 
**ErrorMessage** | **NullableString** |  | 
**DnsRecords** | [**[]ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner**](ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner.md) |  | 
**CloudflareData** | [**NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData**](ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData.md) |  | 

## Methods

### NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner

`func NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner(domain string, status string, errorMessage NullableString, dnsRecords []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner, cloudflareData NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData, ) *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner`

NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner instantiates a new ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerWithDefaults

`func NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerWithDefaults() *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner`

NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerWithDefaults instantiates a new ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetStatus

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetDnsRecords

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetDnsRecords() []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetDnsRecordsOk() (*[]ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) SetDnsRecords(v []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner)`

SetDnsRecords sets DnsRecords field to given value.


### GetCloudflareData

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetCloudflareData() ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData`

GetCloudflareData returns the CloudflareData field if non-nil, zero value otherwise.

### GetCloudflareDataOk

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetCloudflareDataOk() (*ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData, bool)`

GetCloudflareDataOk returns a tuple with the CloudflareData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflareData

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) SetCloudflareData(v ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData)`

SetCloudflareData sets CloudflareData field to given value.


### SetCloudflareDataNil

`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) SetCloudflareDataNil(b bool)`

 SetCloudflareDataNil sets the value for CloudflareData to be an explicit nil

### UnsetCloudflareData
`func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) UnsetCloudflareData()`

UnsetCloudflareData ensures that no value is present for CloudflareData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


