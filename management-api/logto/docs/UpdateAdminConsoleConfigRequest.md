# UpdateAdminConsoleConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignInExperienceCustomized** | Pointer to **bool** |  | [optional] 
**OrganizationCreated** | Pointer to **bool** |  | [optional] 
**DevelopmentTenantMigrationNotification** | Pointer to [**GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification**](GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification.md) |  | [optional] 
**CheckedChargeNotification** | Pointer to [**GetAdminConsoleConfig200ResponseCheckedChargeNotification**](GetAdminConsoleConfig200ResponseCheckedChargeNotification.md) |  | [optional] 

## Methods

### NewUpdateAdminConsoleConfigRequest

`func NewUpdateAdminConsoleConfigRequest() *UpdateAdminConsoleConfigRequest`

NewUpdateAdminConsoleConfigRequest instantiates a new UpdateAdminConsoleConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAdminConsoleConfigRequestWithDefaults

`func NewUpdateAdminConsoleConfigRequestWithDefaults() *UpdateAdminConsoleConfigRequest`

NewUpdateAdminConsoleConfigRequestWithDefaults instantiates a new UpdateAdminConsoleConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignInExperienceCustomized

`func (o *UpdateAdminConsoleConfigRequest) GetSignInExperienceCustomized() bool`

GetSignInExperienceCustomized returns the SignInExperienceCustomized field if non-nil, zero value otherwise.

### GetSignInExperienceCustomizedOk

`func (o *UpdateAdminConsoleConfigRequest) GetSignInExperienceCustomizedOk() (*bool, bool)`

GetSignInExperienceCustomizedOk returns a tuple with the SignInExperienceCustomized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInExperienceCustomized

`func (o *UpdateAdminConsoleConfigRequest) SetSignInExperienceCustomized(v bool)`

SetSignInExperienceCustomized sets SignInExperienceCustomized field to given value.

### HasSignInExperienceCustomized

`func (o *UpdateAdminConsoleConfigRequest) HasSignInExperienceCustomized() bool`

HasSignInExperienceCustomized returns a boolean if a field has been set.

### GetOrganizationCreated

`func (o *UpdateAdminConsoleConfigRequest) GetOrganizationCreated() bool`

GetOrganizationCreated returns the OrganizationCreated field if non-nil, zero value otherwise.

### GetOrganizationCreatedOk

`func (o *UpdateAdminConsoleConfigRequest) GetOrganizationCreatedOk() (*bool, bool)`

GetOrganizationCreatedOk returns a tuple with the OrganizationCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCreated

`func (o *UpdateAdminConsoleConfigRequest) SetOrganizationCreated(v bool)`

SetOrganizationCreated sets OrganizationCreated field to given value.

### HasOrganizationCreated

`func (o *UpdateAdminConsoleConfigRequest) HasOrganizationCreated() bool`

HasOrganizationCreated returns a boolean if a field has been set.

### GetDevelopmentTenantMigrationNotification

`func (o *UpdateAdminConsoleConfigRequest) GetDevelopmentTenantMigrationNotification() GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification`

GetDevelopmentTenantMigrationNotification returns the DevelopmentTenantMigrationNotification field if non-nil, zero value otherwise.

### GetDevelopmentTenantMigrationNotificationOk

`func (o *UpdateAdminConsoleConfigRequest) GetDevelopmentTenantMigrationNotificationOk() (*GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification, bool)`

GetDevelopmentTenantMigrationNotificationOk returns a tuple with the DevelopmentTenantMigrationNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopmentTenantMigrationNotification

`func (o *UpdateAdminConsoleConfigRequest) SetDevelopmentTenantMigrationNotification(v GetAdminConsoleConfig200ResponseDevelopmentTenantMigrationNotification)`

SetDevelopmentTenantMigrationNotification sets DevelopmentTenantMigrationNotification field to given value.

### HasDevelopmentTenantMigrationNotification

`func (o *UpdateAdminConsoleConfigRequest) HasDevelopmentTenantMigrationNotification() bool`

HasDevelopmentTenantMigrationNotification returns a boolean if a field has been set.

### GetCheckedChargeNotification

`func (o *UpdateAdminConsoleConfigRequest) GetCheckedChargeNotification() GetAdminConsoleConfig200ResponseCheckedChargeNotification`

GetCheckedChargeNotification returns the CheckedChargeNotification field if non-nil, zero value otherwise.

### GetCheckedChargeNotificationOk

`func (o *UpdateAdminConsoleConfigRequest) GetCheckedChargeNotificationOk() (*GetAdminConsoleConfig200ResponseCheckedChargeNotification, bool)`

GetCheckedChargeNotificationOk returns a tuple with the CheckedChargeNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedChargeNotification

`func (o *UpdateAdminConsoleConfigRequest) SetCheckedChargeNotification(v GetAdminConsoleConfig200ResponseCheckedChargeNotification)`

SetCheckedChargeNotification sets CheckedChargeNotification field to given value.

### HasCheckedChargeNotification

`func (o *UpdateAdminConsoleConfigRequest) HasCheckedChargeNotification() bool`

HasCheckedChargeNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


