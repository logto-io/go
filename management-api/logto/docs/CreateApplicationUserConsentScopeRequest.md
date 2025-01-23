# CreateApplicationUserConsentScopeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationScopes** | Pointer to **[]string** | A list of organization scope id to assign to the application. Throws error if any given organization scope is not found. | [optional] 
**ResourceScopes** | Pointer to **[]string** | A list of resource scope id to assign to the application. Throws error if any given resource scope is not found. | [optional] 
**OrganizationResourceScopes** | Pointer to **[]string** | A list of organization resource scope id to assign to the application. Throws error if any given resource scope is not found. | [optional] 
**UserScopes** | Pointer to **[]string** | A list of user scope enum value to assign to the application. | [optional] 

## Methods

### NewCreateApplicationUserConsentScopeRequest

`func NewCreateApplicationUserConsentScopeRequest() *CreateApplicationUserConsentScopeRequest`

NewCreateApplicationUserConsentScopeRequest instantiates a new CreateApplicationUserConsentScopeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationUserConsentScopeRequestWithDefaults

`func NewCreateApplicationUserConsentScopeRequestWithDefaults() *CreateApplicationUserConsentScopeRequest`

NewCreateApplicationUserConsentScopeRequestWithDefaults instantiates a new CreateApplicationUserConsentScopeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationScopes

`func (o *CreateApplicationUserConsentScopeRequest) GetOrganizationScopes() []string`

GetOrganizationScopes returns the OrganizationScopes field if non-nil, zero value otherwise.

### GetOrganizationScopesOk

`func (o *CreateApplicationUserConsentScopeRequest) GetOrganizationScopesOk() (*[]string, bool)`

GetOrganizationScopesOk returns a tuple with the OrganizationScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationScopes

`func (o *CreateApplicationUserConsentScopeRequest) SetOrganizationScopes(v []string)`

SetOrganizationScopes sets OrganizationScopes field to given value.

### HasOrganizationScopes

`func (o *CreateApplicationUserConsentScopeRequest) HasOrganizationScopes() bool`

HasOrganizationScopes returns a boolean if a field has been set.

### GetResourceScopes

`func (o *CreateApplicationUserConsentScopeRequest) GetResourceScopes() []string`

GetResourceScopes returns the ResourceScopes field if non-nil, zero value otherwise.

### GetResourceScopesOk

`func (o *CreateApplicationUserConsentScopeRequest) GetResourceScopesOk() (*[]string, bool)`

GetResourceScopesOk returns a tuple with the ResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceScopes

`func (o *CreateApplicationUserConsentScopeRequest) SetResourceScopes(v []string)`

SetResourceScopes sets ResourceScopes field to given value.

### HasResourceScopes

`func (o *CreateApplicationUserConsentScopeRequest) HasResourceScopes() bool`

HasResourceScopes returns a boolean if a field has been set.

### GetOrganizationResourceScopes

`func (o *CreateApplicationUserConsentScopeRequest) GetOrganizationResourceScopes() []string`

GetOrganizationResourceScopes returns the OrganizationResourceScopes field if non-nil, zero value otherwise.

### GetOrganizationResourceScopesOk

`func (o *CreateApplicationUserConsentScopeRequest) GetOrganizationResourceScopesOk() (*[]string, bool)`

GetOrganizationResourceScopesOk returns a tuple with the OrganizationResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationResourceScopes

`func (o *CreateApplicationUserConsentScopeRequest) SetOrganizationResourceScopes(v []string)`

SetOrganizationResourceScopes sets OrganizationResourceScopes field to given value.

### HasOrganizationResourceScopes

`func (o *CreateApplicationUserConsentScopeRequest) HasOrganizationResourceScopes() bool`

HasOrganizationResourceScopes returns a boolean if a field has been set.

### GetUserScopes

`func (o *CreateApplicationUserConsentScopeRequest) GetUserScopes() []string`

GetUserScopes returns the UserScopes field if non-nil, zero value otherwise.

### GetUserScopesOk

`func (o *CreateApplicationUserConsentScopeRequest) GetUserScopesOk() (*[]string, bool)`

GetUserScopesOk returns a tuple with the UserScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScopes

`func (o *CreateApplicationUserConsentScopeRequest) SetUserScopes(v []string)`

SetUserScopes sets UserScopes field to given value.

### HasUserScopes

`func (o *CreateApplicationUserConsentScopeRequest) HasUserScopes() bool`

HasUserScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


