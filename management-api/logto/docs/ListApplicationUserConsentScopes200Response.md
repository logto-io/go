# ListApplicationUserConsentScopes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationScopes** | [**[]ListApplicationUserConsentScopes200ResponseOrganizationScopesInner**](ListApplicationUserConsentScopes200ResponseOrganizationScopesInner.md) | A list of organization scope details assigned to the application. | 
**ResourceScopes** | [**[]ListApplicationUserConsentScopes200ResponseResourceScopesInner**](ListApplicationUserConsentScopes200ResponseResourceScopesInner.md) | A list of resource scope details grouped by resource id assigned to the application. | 
**OrganizationResourceScopes** | [**[]ListApplicationUserConsentScopes200ResponseResourceScopesInner**](ListApplicationUserConsentScopes200ResponseResourceScopesInner.md) | A list of organization resource scope details grouped by resource id assigned to the application. | 
**UserScopes** | **[]string** | A list of user scope enum value assigned to the application. | 

## Methods

### NewListApplicationUserConsentScopes200Response

`func NewListApplicationUserConsentScopes200Response(organizationScopes []ListApplicationUserConsentScopes200ResponseOrganizationScopesInner, resourceScopes []ListApplicationUserConsentScopes200ResponseResourceScopesInner, organizationResourceScopes []ListApplicationUserConsentScopes200ResponseResourceScopesInner, userScopes []string, ) *ListApplicationUserConsentScopes200Response`

NewListApplicationUserConsentScopes200Response instantiates a new ListApplicationUserConsentScopes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApplicationUserConsentScopes200ResponseWithDefaults

`func NewListApplicationUserConsentScopes200ResponseWithDefaults() *ListApplicationUserConsentScopes200Response`

NewListApplicationUserConsentScopes200ResponseWithDefaults instantiates a new ListApplicationUserConsentScopes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationScopes

`func (o *ListApplicationUserConsentScopes200Response) GetOrganizationScopes() []ListApplicationUserConsentScopes200ResponseOrganizationScopesInner`

GetOrganizationScopes returns the OrganizationScopes field if non-nil, zero value otherwise.

### GetOrganizationScopesOk

`func (o *ListApplicationUserConsentScopes200Response) GetOrganizationScopesOk() (*[]ListApplicationUserConsentScopes200ResponseOrganizationScopesInner, bool)`

GetOrganizationScopesOk returns a tuple with the OrganizationScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationScopes

`func (o *ListApplicationUserConsentScopes200Response) SetOrganizationScopes(v []ListApplicationUserConsentScopes200ResponseOrganizationScopesInner)`

SetOrganizationScopes sets OrganizationScopes field to given value.


### GetResourceScopes

`func (o *ListApplicationUserConsentScopes200Response) GetResourceScopes() []ListApplicationUserConsentScopes200ResponseResourceScopesInner`

GetResourceScopes returns the ResourceScopes field if non-nil, zero value otherwise.

### GetResourceScopesOk

`func (o *ListApplicationUserConsentScopes200Response) GetResourceScopesOk() (*[]ListApplicationUserConsentScopes200ResponseResourceScopesInner, bool)`

GetResourceScopesOk returns a tuple with the ResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceScopes

`func (o *ListApplicationUserConsentScopes200Response) SetResourceScopes(v []ListApplicationUserConsentScopes200ResponseResourceScopesInner)`

SetResourceScopes sets ResourceScopes field to given value.


### GetOrganizationResourceScopes

`func (o *ListApplicationUserConsentScopes200Response) GetOrganizationResourceScopes() []ListApplicationUserConsentScopes200ResponseResourceScopesInner`

GetOrganizationResourceScopes returns the OrganizationResourceScopes field if non-nil, zero value otherwise.

### GetOrganizationResourceScopesOk

`func (o *ListApplicationUserConsentScopes200Response) GetOrganizationResourceScopesOk() (*[]ListApplicationUserConsentScopes200ResponseResourceScopesInner, bool)`

GetOrganizationResourceScopesOk returns a tuple with the OrganizationResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationResourceScopes

`func (o *ListApplicationUserConsentScopes200Response) SetOrganizationResourceScopes(v []ListApplicationUserConsentScopes200ResponseResourceScopesInner)`

SetOrganizationResourceScopes sets OrganizationResourceScopes field to given value.


### GetUserScopes

`func (o *ListApplicationUserConsentScopes200Response) GetUserScopes() []string`

GetUserScopes returns the UserScopes field if non-nil, zero value otherwise.

### GetUserScopesOk

`func (o *ListApplicationUserConsentScopes200Response) GetUserScopesOk() (*[]string, bool)`

GetUserScopesOk returns a tuple with the UserScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScopes

`func (o *ListApplicationUserConsentScopes200Response) SetUserScopes(v []string)`

SetUserScopes sets UserScopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


