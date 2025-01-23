# CreateDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain name, e.g. &#x60;example.com&#x60;. | 

## Methods

### NewCreateDomainRequest

`func NewCreateDomainRequest(domain string, ) *CreateDomainRequest`

NewCreateDomainRequest instantiates a new CreateDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainRequestWithDefaults

`func NewCreateDomainRequestWithDefaults() *CreateDomainRequest`

NewCreateDomainRequestWithDefaults instantiates a new CreateDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CreateDomainRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateDomainRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateDomainRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


