# ListOrganizationRoles200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**Type** | **string** |  | 
**Scopes** | [**[]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner**](ListApplicationOrganizations200ResponseInnerOrganizationRolesInner.md) |  | 
**ResourceScopes** | [**[]ListOrganizationRoles200ResponseInnerResourceScopesInner**](ListOrganizationRoles200ResponseInnerResourceScopesInner.md) |  | 

## Methods

### NewListOrganizationRoles200ResponseInner

`func NewListOrganizationRoles200ResponseInner(tenantId string, id string, name string, description NullableString, type_ string, scopes []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, resourceScopes []ListOrganizationRoles200ResponseInnerResourceScopesInner, ) *ListOrganizationRoles200ResponseInner`

NewListOrganizationRoles200ResponseInner instantiates a new ListOrganizationRoles200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizationRoles200ResponseInnerWithDefaults

`func NewListOrganizationRoles200ResponseInnerWithDefaults() *ListOrganizationRoles200ResponseInner`

NewListOrganizationRoles200ResponseInnerWithDefaults instantiates a new ListOrganizationRoles200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListOrganizationRoles200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListOrganizationRoles200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListOrganizationRoles200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListOrganizationRoles200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOrganizationRoles200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOrganizationRoles200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListOrganizationRoles200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListOrganizationRoles200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListOrganizationRoles200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListOrganizationRoles200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListOrganizationRoles200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListOrganizationRoles200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ListOrganizationRoles200ResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListOrganizationRoles200ResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *ListOrganizationRoles200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListOrganizationRoles200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListOrganizationRoles200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetScopes

`func (o *ListOrganizationRoles200ResponseInner) GetScopes() []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ListOrganizationRoles200ResponseInner) GetScopesOk() (*[]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ListOrganizationRoles200ResponseInner) SetScopes(v []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner)`

SetScopes sets Scopes field to given value.


### GetResourceScopes

`func (o *ListOrganizationRoles200ResponseInner) GetResourceScopes() []ListOrganizationRoles200ResponseInnerResourceScopesInner`

GetResourceScopes returns the ResourceScopes field if non-nil, zero value otherwise.

### GetResourceScopesOk

`func (o *ListOrganizationRoles200ResponseInner) GetResourceScopesOk() (*[]ListOrganizationRoles200ResponseInnerResourceScopesInner, bool)`

GetResourceScopesOk returns a tuple with the ResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceScopes

`func (o *ListOrganizationRoles200ResponseInner) SetResourceScopes(v []ListOrganizationRoles200ResponseInnerResourceScopesInner)`

SetResourceScopes sets ResourceScopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


