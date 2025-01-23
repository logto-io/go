# ListRoleScopes200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**ResourceId** | **string** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**CreatedAt** | **float32** |  | 
**Resource** | [**GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource**](GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource.md) |  | 

## Methods

### NewListRoleScopes200ResponseInner

`func NewListRoleScopes200ResponseInner(tenantId string, id string, resourceId string, name string, description NullableString, createdAt float32, resource GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource, ) *ListRoleScopes200ResponseInner`

NewListRoleScopes200ResponseInner instantiates a new ListRoleScopes200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoleScopes200ResponseInnerWithDefaults

`func NewListRoleScopes200ResponseInnerWithDefaults() *ListRoleScopes200ResponseInner`

NewListRoleScopes200ResponseInnerWithDefaults instantiates a new ListRoleScopes200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListRoleScopes200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListRoleScopes200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListRoleScopes200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListRoleScopes200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListRoleScopes200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListRoleScopes200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetResourceId

`func (o *ListRoleScopes200ResponseInner) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ListRoleScopes200ResponseInner) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ListRoleScopes200ResponseInner) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetName

`func (o *ListRoleScopes200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListRoleScopes200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListRoleScopes200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListRoleScopes200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListRoleScopes200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListRoleScopes200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ListRoleScopes200ResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListRoleScopes200ResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatedAt

`func (o *ListRoleScopes200ResponseInner) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListRoleScopes200ResponseInner) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListRoleScopes200ResponseInner) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetResource

`func (o *ListRoleScopes200ResponseInner) GetResource() GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ListRoleScopes200ResponseInner) GetResourceOk() (*GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ListRoleScopes200ResponseInner) SetResource(v GetJwtCustomizer200ResponseOneOfContextSampleUserRolesInnerScopesInnerResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


