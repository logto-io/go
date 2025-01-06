# GetOrganizationInvitation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**InviterId** | **NullableString** |  | 
**Invitee** | **string** |  | 
**AcceptedUserId** | **NullableString** |  | 
**OrganizationId** | **string** |  | 
**Status** | **string** |  | 
**CreatedAt** | **float32** |  | 
**UpdatedAt** | **float32** |  | 
**ExpiresAt** | **float32** |  | 
**OrganizationRoles** | [**[]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner**](ListApplicationOrganizations200ResponseInnerOrganizationRolesInner.md) |  | 

## Methods

### NewGetOrganizationInvitation200Response

`func NewGetOrganizationInvitation200Response(tenantId string, id string, inviterId NullableString, invitee string, acceptedUserId NullableString, organizationId string, status string, createdAt float32, updatedAt float32, expiresAt float32, organizationRoles []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, ) *GetOrganizationInvitation200Response`

NewGetOrganizationInvitation200Response instantiates a new GetOrganizationInvitation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationInvitation200ResponseWithDefaults

`func NewGetOrganizationInvitation200ResponseWithDefaults() *GetOrganizationInvitation200Response`

NewGetOrganizationInvitation200ResponseWithDefaults instantiates a new GetOrganizationInvitation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *GetOrganizationInvitation200Response) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GetOrganizationInvitation200Response) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GetOrganizationInvitation200Response) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *GetOrganizationInvitation200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationInvitation200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationInvitation200Response) SetId(v string)`

SetId sets Id field to given value.


### GetInviterId

`func (o *GetOrganizationInvitation200Response) GetInviterId() string`

GetInviterId returns the InviterId field if non-nil, zero value otherwise.

### GetInviterIdOk

`func (o *GetOrganizationInvitation200Response) GetInviterIdOk() (*string, bool)`

GetInviterIdOk returns a tuple with the InviterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviterId

`func (o *GetOrganizationInvitation200Response) SetInviterId(v string)`

SetInviterId sets InviterId field to given value.


### SetInviterIdNil

`func (o *GetOrganizationInvitation200Response) SetInviterIdNil(b bool)`

 SetInviterIdNil sets the value for InviterId to be an explicit nil

### UnsetInviterId
`func (o *GetOrganizationInvitation200Response) UnsetInviterId()`

UnsetInviterId ensures that no value is present for InviterId, not even an explicit nil
### GetInvitee

`func (o *GetOrganizationInvitation200Response) GetInvitee() string`

GetInvitee returns the Invitee field if non-nil, zero value otherwise.

### GetInviteeOk

`func (o *GetOrganizationInvitation200Response) GetInviteeOk() (*string, bool)`

GetInviteeOk returns a tuple with the Invitee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitee

`func (o *GetOrganizationInvitation200Response) SetInvitee(v string)`

SetInvitee sets Invitee field to given value.


### GetAcceptedUserId

`func (o *GetOrganizationInvitation200Response) GetAcceptedUserId() string`

GetAcceptedUserId returns the AcceptedUserId field if non-nil, zero value otherwise.

### GetAcceptedUserIdOk

`func (o *GetOrganizationInvitation200Response) GetAcceptedUserIdOk() (*string, bool)`

GetAcceptedUserIdOk returns a tuple with the AcceptedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedUserId

`func (o *GetOrganizationInvitation200Response) SetAcceptedUserId(v string)`

SetAcceptedUserId sets AcceptedUserId field to given value.


### SetAcceptedUserIdNil

`func (o *GetOrganizationInvitation200Response) SetAcceptedUserIdNil(b bool)`

 SetAcceptedUserIdNil sets the value for AcceptedUserId to be an explicit nil

### UnsetAcceptedUserId
`func (o *GetOrganizationInvitation200Response) UnsetAcceptedUserId()`

UnsetAcceptedUserId ensures that no value is present for AcceptedUserId, not even an explicit nil
### GetOrganizationId

`func (o *GetOrganizationInvitation200Response) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetOrganizationInvitation200Response) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetOrganizationInvitation200Response) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetStatus

`func (o *GetOrganizationInvitation200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationInvitation200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationInvitation200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *GetOrganizationInvitation200Response) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrganizationInvitation200Response) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrganizationInvitation200Response) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetOrganizationInvitation200Response) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetOrganizationInvitation200Response) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetOrganizationInvitation200Response) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetExpiresAt

`func (o *GetOrganizationInvitation200Response) GetExpiresAt() float32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetOrganizationInvitation200Response) GetExpiresAtOk() (*float32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetOrganizationInvitation200Response) SetExpiresAt(v float32)`

SetExpiresAt sets ExpiresAt field to given value.


### GetOrganizationRoles

`func (o *GetOrganizationInvitation200Response) GetOrganizationRoles() []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner`

GetOrganizationRoles returns the OrganizationRoles field if non-nil, zero value otherwise.

### GetOrganizationRolesOk

`func (o *GetOrganizationInvitation200Response) GetOrganizationRolesOk() (*[]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, bool)`

GetOrganizationRolesOk returns a tuple with the OrganizationRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoles

`func (o *GetOrganizationInvitation200Response) SetOrganizationRoles(v []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner)`

SetOrganizationRoles sets OrganizationRoles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


