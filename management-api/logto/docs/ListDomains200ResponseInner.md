# ListDomains200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Domain** | **string** |  | 
**Status** | **string** |  | 
**ErrorMessage** | **NullableString** |  | 
**DnsRecords** | [**[]ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner**](ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner.md) |  | 

## Methods

### NewListDomains200ResponseInner

`func NewListDomains200ResponseInner(id string, domain string, status string, errorMessage NullableString, dnsRecords []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner, ) *ListDomains200ResponseInner`

NewListDomains200ResponseInner instantiates a new ListDomains200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDomains200ResponseInnerWithDefaults

`func NewListDomains200ResponseInnerWithDefaults() *ListDomains200ResponseInner`

NewListDomains200ResponseInnerWithDefaults instantiates a new ListDomains200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListDomains200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListDomains200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListDomains200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetDomain

`func (o *ListDomains200ResponseInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ListDomains200ResponseInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ListDomains200ResponseInner) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetStatus

`func (o *ListDomains200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListDomains200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListDomains200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *ListDomains200ResponseInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListDomains200ResponseInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListDomains200ResponseInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *ListDomains200ResponseInner) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ListDomains200ResponseInner) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetDnsRecords

`func (o *ListDomains200ResponseInner) GetDnsRecords() []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *ListDomains200ResponseInner) GetDnsRecordsOk() (*[]ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *ListDomains200ResponseInner) SetDnsRecords(v []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner)`

SetDnsRecords sets DnsRecords field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


