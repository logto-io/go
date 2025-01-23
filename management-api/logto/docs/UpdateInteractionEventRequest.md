# UpdateInteractionEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InteractionEvent** | **string** | The type of the interaction event. Only &#x60;SignIn&#x60; and &#x60;Register&#x60; are supported. | 

## Methods

### NewUpdateInteractionEventRequest

`func NewUpdateInteractionEventRequest(interactionEvent string, ) *UpdateInteractionEventRequest`

NewUpdateInteractionEventRequest instantiates a new UpdateInteractionEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInteractionEventRequestWithDefaults

`func NewUpdateInteractionEventRequestWithDefaults() *UpdateInteractionEventRequest`

NewUpdateInteractionEventRequestWithDefaults instantiates a new UpdateInteractionEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteractionEvent

`func (o *UpdateInteractionEventRequest) GetInteractionEvent() string`

GetInteractionEvent returns the InteractionEvent field if non-nil, zero value otherwise.

### GetInteractionEventOk

`func (o *UpdateInteractionEventRequest) GetInteractionEventOk() (*string, bool)`

GetInteractionEventOk returns a tuple with the InteractionEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionEvent

`func (o *UpdateInteractionEventRequest) SetInteractionEvent(v string)`

SetInteractionEvent sets InteractionEvent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


