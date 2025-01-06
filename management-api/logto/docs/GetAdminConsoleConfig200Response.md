# GetAdminConsoleConfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignInExperienceCustomized** | **bool** |  | 
**OrganizationCreated** | **bool** |  | 
**DevelopmentTenantMigrationNotification** | Pointer to [**GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification**](GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification.md) |  | [optional] 
**CheckedChargeNotification** | Pointer to [**GetAdminConsoleConfig200ResponseCheckedChargeNotification**](GetAdminConsoleConfig200ResponseCheckedChargeNotification.md) |  | [optional] 

## Methods

### NewGetAdminConsoleConfig200Response

`func NewGetAdminConsoleConfig200Response(signInExperienceCustomized bool, organizationCreated bool, ) *GetAdminConsoleConfig200Response`

NewGetAdminConsoleConfig200Response instantiates a new GetAdminConsoleConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdminConsoleConfig200ResponseWithDefaults

`func NewGetAdminConsoleConfig200ResponseWithDefaults() *GetAdminConsoleConfig200Response`

NewGetAdminConsoleConfig200ResponseWithDefaults instantiates a new GetAdminConsoleConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignInExperienceCustomized

`func (o *GetAdminConsoleConfig200Response) GetSignInExperienceCustomized() bool`

GetSignInExperienceCustomized returns the SignInExperienceCustomized field if non-nil, zero value otherwise.

### GetSignInExperienceCustomizedOk

`func (o *GetAdminConsoleConfig200Response) GetSignInExperienceCustomizedOk() (*bool, bool)`

GetSignInExperienceCustomizedOk returns a tuple with the SignInExperienceCustomized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInExperienceCustomized

`func (o *GetAdminConsoleConfig200Response) SetSignInExperienceCustomized(v bool)`

SetSignInExperienceCustomized sets SignInExperienceCustomized field to given value.


### GetOrganizationCreated

`func (o *GetAdminConsoleConfig200Response) GetOrganizationCreated() bool`

GetOrganizationCreated returns the OrganizationCreated field if non-nil, zero value otherwise.

### GetOrganizationCreatedOk

`func (o *GetAdminConsoleConfig200Response) GetOrganizationCreatedOk() (*bool, bool)`

GetOrganizationCreatedOk returns a tuple with the OrganizationCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCreated

`func (o *GetAdminConsoleConfig200Response) SetOrganizationCreated(v bool)`

SetOrganizationCreated sets OrganizationCreated field to given value.


### GetDevelopmentTenantMigrationNotification

`func (o *GetAdminConsoleConfig200Response) GetDevelopmentTenantMigrationNotification() GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification`

GetDevelopmentTenantMigrationNotification returns the DevelopmentTenantMigrationNotification field if non-nil, zero value otherwise.

### GetDevelopmentTenantMigrationNotificationOk

`func (o *GetAdminConsoleConfig200Response) GetDevelopmentTenantMigrationNotificationOk() (*GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification, bool)`

GetDevelopmentTenantMigrationNotificationOk returns a tuple with the DevelopmentTenantMigrationNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopmentTenantMigrationNotification

`func (o *GetAdminConsoleConfig200Response) SetDevelopmentTenantMigrationNotification(v GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification)`

SetDevelopmentTenantMigrationNotification sets DevelopmentTenantMigrationNotification field to given value.

### HasDevelopmentTenantMigrationNotification

`func (o *GetAdminConsoleConfig200Response) HasDevelopmentTenantMigrationNotification() bool`

HasDevelopmentTenantMigrationNotification returns a boolean if a field has been set.

### GetCheckedChargeNotification

`func (o *GetAdminConsoleConfig200Response) GetCheckedChargeNotification() GetAdminConsoleConfig200ResponseCheckedChargeNotification`

GetCheckedChargeNotification returns the CheckedChargeNotification field if non-nil, zero value otherwise.

### GetCheckedChargeNotificationOk

`func (o *GetAdminConsoleConfig200Response) GetCheckedChargeNotificationOk() (*GetAdminConsoleConfig200ResponseCheckedChargeNotification, bool)`

GetCheckedChargeNotificationOk returns a tuple with the CheckedChargeNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedChargeNotification

`func (o *GetAdminConsoleConfig200Response) SetCheckedChargeNotification(v GetAdminConsoleConfig200ResponseCheckedChargeNotification)`

SetCheckedChargeNotification sets CheckedChargeNotification field to given value.

### HasCheckedChargeNotification

`func (o *GetAdminConsoleConfig200Response) HasCheckedChargeNotification() bool`

HasCheckedChargeNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


