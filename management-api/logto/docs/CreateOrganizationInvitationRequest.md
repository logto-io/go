# CreateOrganizationInvitationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InviterId** | Pointer to **NullableString** | The ID of the user who is inviting the user to join the organization. | [optional] 
**Invitee** | **string** | The email address of the user to invite to join the organization. | 
**OrganizationId** | **string** | The ID of the organization to invite the user to join. | 
**ExpiresAt** | **float32** | The epoch time in milliseconds when the invitation expires. | 
**OrganizationRoleIds** | Pointer to **[]string** | The IDs of the organization roles to assign to the user when they accept the invitation. | [optional] 
**MessagePayload** | [**CreateOrganizationInvitationRequestMessagePayload**](CreateOrganizationInvitationRequestMessagePayload.md) |  | [default to false]

## Methods

### NewCreateOrganizationInvitationRequest

`func NewCreateOrganizationInvitationRequest(invitee string, organizationId string, expiresAt float32, messagePayload CreateOrganizationInvitationRequestMessagePayload, ) *CreateOrganizationInvitationRequest`

NewCreateOrganizationInvitationRequest instantiates a new CreateOrganizationInvitationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationInvitationRequestWithDefaults

`func NewCreateOrganizationInvitationRequestWithDefaults() *CreateOrganizationInvitationRequest`

NewCreateOrganizationInvitationRequestWithDefaults instantiates a new CreateOrganizationInvitationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInviterId

`func (o *CreateOrganizationInvitationRequest) GetInviterId() string`

GetInviterId returns the InviterId field if non-nil, zero value otherwise.

### GetInviterIdOk

`func (o *CreateOrganizationInvitationRequest) GetInviterIdOk() (*string, bool)`

GetInviterIdOk returns a tuple with the InviterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviterId

`func (o *CreateOrganizationInvitationRequest) SetInviterId(v string)`

SetInviterId sets InviterId field to given value.

### HasInviterId

`func (o *CreateOrganizationInvitationRequest) HasInviterId() bool`

HasInviterId returns a boolean if a field has been set.

### SetInviterIdNil

`func (o *CreateOrganizationInvitationRequest) SetInviterIdNil(b bool)`

 SetInviterIdNil sets the value for InviterId to be an explicit nil

### UnsetInviterId
`func (o *CreateOrganizationInvitationRequest) UnsetInviterId()`

UnsetInviterId ensures that no value is present for InviterId, not even an explicit nil
### GetInvitee

`func (o *CreateOrganizationInvitationRequest) GetInvitee() string`

GetInvitee returns the Invitee field if non-nil, zero value otherwise.

### GetInviteeOk

`func (o *CreateOrganizationInvitationRequest) GetInviteeOk() (*string, bool)`

GetInviteeOk returns a tuple with the Invitee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitee

`func (o *CreateOrganizationInvitationRequest) SetInvitee(v string)`

SetInvitee sets Invitee field to given value.


### GetOrganizationId

`func (o *CreateOrganizationInvitationRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateOrganizationInvitationRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateOrganizationInvitationRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetExpiresAt

`func (o *CreateOrganizationInvitationRequest) GetExpiresAt() float32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateOrganizationInvitationRequest) GetExpiresAtOk() (*float32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateOrganizationInvitationRequest) SetExpiresAt(v float32)`

SetExpiresAt sets ExpiresAt field to given value.


### GetOrganizationRoleIds

`func (o *CreateOrganizationInvitationRequest) GetOrganizationRoleIds() []string`

GetOrganizationRoleIds returns the OrganizationRoleIds field if non-nil, zero value otherwise.

### GetOrganizationRoleIdsOk

`func (o *CreateOrganizationInvitationRequest) GetOrganizationRoleIdsOk() (*[]string, bool)`

GetOrganizationRoleIdsOk returns a tuple with the OrganizationRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoleIds

`func (o *CreateOrganizationInvitationRequest) SetOrganizationRoleIds(v []string)`

SetOrganizationRoleIds sets OrganizationRoleIds field to given value.

### HasOrganizationRoleIds

`func (o *CreateOrganizationInvitationRequest) HasOrganizationRoleIds() bool`

HasOrganizationRoleIds returns a boolean if a field has been set.

### GetMessagePayload

`func (o *CreateOrganizationInvitationRequest) GetMessagePayload() CreateOrganizationInvitationRequestMessagePayload`

GetMessagePayload returns the MessagePayload field if non-nil, zero value otherwise.

### GetMessagePayloadOk

`func (o *CreateOrganizationInvitationRequest) GetMessagePayloadOk() (*CreateOrganizationInvitationRequestMessagePayload, bool)`

GetMessagePayloadOk returns a tuple with the MessagePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagePayload

`func (o *CreateOrganizationInvitationRequest) SetMessagePayload(v CreateOrganizationInvitationRequestMessagePayload)`

SetMessagePayload sets MessagePayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


