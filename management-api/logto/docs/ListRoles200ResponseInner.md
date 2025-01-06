# ListRoles200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Type** | **string** |  | 
**IsDefault** | **bool** |  | 
**UsersCount** | **float32** |  | 
**FeaturedUsers** | [**[]ListRoles200ResponseInnerFeaturedUsersInner**](ListRoles200ResponseInnerFeaturedUsersInner.md) |  | 
**ApplicationsCount** | **float32** |  | 
**FeaturedApplications** | [**[]ListRoles200ResponseInnerFeaturedApplicationsInner**](ListRoles200ResponseInnerFeaturedApplicationsInner.md) |  | 

## Methods

### NewListRoles200ResponseInner

`func NewListRoles200ResponseInner(tenantId string, id string, name string, description string, type_ string, isDefault bool, usersCount float32, featuredUsers []ListRoles200ResponseInnerFeaturedUsersInner, applicationsCount float32, featuredApplications []ListRoles200ResponseInnerFeaturedApplicationsInner, ) *ListRoles200ResponseInner`

NewListRoles200ResponseInner instantiates a new ListRoles200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoles200ResponseInnerWithDefaults

`func NewListRoles200ResponseInnerWithDefaults() *ListRoles200ResponseInner`

NewListRoles200ResponseInnerWithDefaults instantiates a new ListRoles200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListRoles200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListRoles200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListRoles200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *ListRoles200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListRoles200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListRoles200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListRoles200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListRoles200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListRoles200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListRoles200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListRoles200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListRoles200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ListRoles200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListRoles200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListRoles200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetIsDefault

`func (o *ListRoles200ResponseInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ListRoles200ResponseInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ListRoles200ResponseInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetUsersCount

`func (o *ListRoles200ResponseInner) GetUsersCount() float32`

GetUsersCount returns the UsersCount field if non-nil, zero value otherwise.

### GetUsersCountOk

`func (o *ListRoles200ResponseInner) GetUsersCountOk() (*float32, bool)`

GetUsersCountOk returns a tuple with the UsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCount

`func (o *ListRoles200ResponseInner) SetUsersCount(v float32)`

SetUsersCount sets UsersCount field to given value.


### GetFeaturedUsers

`func (o *ListRoles200ResponseInner) GetFeaturedUsers() []ListRoles200ResponseInnerFeaturedUsersInner`

GetFeaturedUsers returns the FeaturedUsers field if non-nil, zero value otherwise.

### GetFeaturedUsersOk

`func (o *ListRoles200ResponseInner) GetFeaturedUsersOk() (*[]ListRoles200ResponseInnerFeaturedUsersInner, bool)`

GetFeaturedUsersOk returns a tuple with the FeaturedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedUsers

`func (o *ListRoles200ResponseInner) SetFeaturedUsers(v []ListRoles200ResponseInnerFeaturedUsersInner)`

SetFeaturedUsers sets FeaturedUsers field to given value.


### GetApplicationsCount

`func (o *ListRoles200ResponseInner) GetApplicationsCount() float32`

GetApplicationsCount returns the ApplicationsCount field if non-nil, zero value otherwise.

### GetApplicationsCountOk

`func (o *ListRoles200ResponseInner) GetApplicationsCountOk() (*float32, bool)`

GetApplicationsCountOk returns a tuple with the ApplicationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsCount

`func (o *ListRoles200ResponseInner) SetApplicationsCount(v float32)`

SetApplicationsCount sets ApplicationsCount field to given value.


### GetFeaturedApplications

`func (o *ListRoles200ResponseInner) GetFeaturedApplications() []ListRoles200ResponseInnerFeaturedApplicationsInner`

GetFeaturedApplications returns the FeaturedApplications field if non-nil, zero value otherwise.

### GetFeaturedApplicationsOk

`func (o *ListRoles200ResponseInner) GetFeaturedApplicationsOk() (*[]ListRoles200ResponseInnerFeaturedApplicationsInner, bool)`

GetFeaturedApplicationsOk returns a tuple with the FeaturedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedApplications

`func (o *ListRoles200ResponseInner) SetFeaturedApplications(v []ListRoles200ResponseInnerFeaturedApplicationsInner)`

SetFeaturedApplications sets FeaturedApplications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


