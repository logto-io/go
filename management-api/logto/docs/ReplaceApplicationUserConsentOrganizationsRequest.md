# ReplaceApplicationUserConsentOrganizationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationIds** | **[]string** | A list of organization ids to be granted. &lt;br/&gt; All the existing organizations&#39; access will be revoked if not in the list. &lt;br/&gt; If the list is empty, all the organizations&#39; access will be revoked. | 

## Methods

### NewReplaceApplicationUserConsentOrganizationsRequest

`func NewReplaceApplicationUserConsentOrganizationsRequest(organizationIds []string, ) *ReplaceApplicationUserConsentOrganizationsRequest`

NewReplaceApplicationUserConsentOrganizationsRequest instantiates a new ReplaceApplicationUserConsentOrganizationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceApplicationUserConsentOrganizationsRequestWithDefaults

`func NewReplaceApplicationUserConsentOrganizationsRequestWithDefaults() *ReplaceApplicationUserConsentOrganizationsRequest`

NewReplaceApplicationUserConsentOrganizationsRequestWithDefaults instantiates a new ReplaceApplicationUserConsentOrganizationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationIds

`func (o *ReplaceApplicationUserConsentOrganizationsRequest) GetOrganizationIds() []string`

GetOrganizationIds returns the OrganizationIds field if non-nil, zero value otherwise.

### GetOrganizationIdsOk

`func (o *ReplaceApplicationUserConsentOrganizationsRequest) GetOrganizationIdsOk() (*[]string, bool)`

GetOrganizationIdsOk returns a tuple with the OrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationIds

`func (o *ReplaceApplicationUserConsentOrganizationsRequest) SetOrganizationIds(v []string)`

SetOrganizationIds sets OrganizationIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


