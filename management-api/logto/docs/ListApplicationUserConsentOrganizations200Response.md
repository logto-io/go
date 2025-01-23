# ListApplicationUserConsentOrganizations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizations** | [**[]ListApplicationUserConsentOrganizations200ResponseOrganizationsInner**](ListApplicationUserConsentOrganizations200ResponseOrganizationsInner.md) | A list of organization entities granted by the user for the application. | 

## Methods

### NewListApplicationUserConsentOrganizations200Response

`func NewListApplicationUserConsentOrganizations200Response(organizations []ListApplicationUserConsentOrganizations200ResponseOrganizationsInner, ) *ListApplicationUserConsentOrganizations200Response`

NewListApplicationUserConsentOrganizations200Response instantiates a new ListApplicationUserConsentOrganizations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApplicationUserConsentOrganizations200ResponseWithDefaults

`func NewListApplicationUserConsentOrganizations200ResponseWithDefaults() *ListApplicationUserConsentOrganizations200Response`

NewListApplicationUserConsentOrganizations200ResponseWithDefaults instantiates a new ListApplicationUserConsentOrganizations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizations

`func (o *ListApplicationUserConsentOrganizations200Response) GetOrganizations() []ListApplicationUserConsentOrganizations200ResponseOrganizationsInner`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *ListApplicationUserConsentOrganizations200Response) GetOrganizationsOk() (*[]ListApplicationUserConsentOrganizations200ResponseOrganizationsInner, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *ListApplicationUserConsentOrganizations200Response) SetOrganizations(v []ListApplicationUserConsentOrganizations200ResponseOrganizationsInner)`

SetOrganizations sets Organizations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


