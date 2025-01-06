# ReplaceOrganizationInvitationStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedUserId** | Pointer to **NullableString** | The ID of the user who accepted the organization invitation. Required if the status is \&quot;Accepted\&quot;. | [optional] 
**Status** | **string** | The status of the organization invitation. | 

## Methods

### NewReplaceOrganizationInvitationStatusRequest

`func NewReplaceOrganizationInvitationStatusRequest(status string, ) *ReplaceOrganizationInvitationStatusRequest`

NewReplaceOrganizationInvitationStatusRequest instantiates a new ReplaceOrganizationInvitationStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceOrganizationInvitationStatusRequestWithDefaults

`func NewReplaceOrganizationInvitationStatusRequestWithDefaults() *ReplaceOrganizationInvitationStatusRequest`

NewReplaceOrganizationInvitationStatusRequestWithDefaults instantiates a new ReplaceOrganizationInvitationStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedUserId

`func (o *ReplaceOrganizationInvitationStatusRequest) GetAcceptedUserId() string`

GetAcceptedUserId returns the AcceptedUserId field if non-nil, zero value otherwise.

### GetAcceptedUserIdOk

`func (o *ReplaceOrganizationInvitationStatusRequest) GetAcceptedUserIdOk() (*string, bool)`

GetAcceptedUserIdOk returns a tuple with the AcceptedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedUserId

`func (o *ReplaceOrganizationInvitationStatusRequest) SetAcceptedUserId(v string)`

SetAcceptedUserId sets AcceptedUserId field to given value.

### HasAcceptedUserId

`func (o *ReplaceOrganizationInvitationStatusRequest) HasAcceptedUserId() bool`

HasAcceptedUserId returns a boolean if a field has been set.

### SetAcceptedUserIdNil

`func (o *ReplaceOrganizationInvitationStatusRequest) SetAcceptedUserIdNil(b bool)`

 SetAcceptedUserIdNil sets the value for AcceptedUserId to be an explicit nil

### UnsetAcceptedUserId
`func (o *ReplaceOrganizationInvitationStatusRequest) UnsetAcceptedUserId()`

UnsetAcceptedUserId ensures that no value is present for AcceptedUserId, not even an explicit nil
### GetStatus

`func (o *ReplaceOrganizationInvitationStatusRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReplaceOrganizationInvitationStatusRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReplaceOrganizationInvitationStatusRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


