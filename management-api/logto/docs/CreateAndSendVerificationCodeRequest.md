# CreateAndSendVerificationCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [**CreateAndSendVerificationCodeRequestIdentifier**](CreateAndSendVerificationCodeRequestIdentifier.md) |  | 
**InteractionEvent** | **string** | The interaction event for which the verification code will be used. Supported values are &#x60;SignIn&#x60;, &#x60;Register&#x60;, and &#x60;ForgotPassword&#x60;. This determines the template for the verification code. | 

## Methods

### NewCreateAndSendVerificationCodeRequest

`func NewCreateAndSendVerificationCodeRequest(identifier CreateAndSendVerificationCodeRequestIdentifier, interactionEvent string, ) *CreateAndSendVerificationCodeRequest`

NewCreateAndSendVerificationCodeRequest instantiates a new CreateAndSendVerificationCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAndSendVerificationCodeRequestWithDefaults

`func NewCreateAndSendVerificationCodeRequestWithDefaults() *CreateAndSendVerificationCodeRequest`

NewCreateAndSendVerificationCodeRequestWithDefaults instantiates a new CreateAndSendVerificationCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CreateAndSendVerificationCodeRequest) GetIdentifier() CreateAndSendVerificationCodeRequestIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateAndSendVerificationCodeRequest) GetIdentifierOk() (*CreateAndSendVerificationCodeRequestIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateAndSendVerificationCodeRequest) SetIdentifier(v CreateAndSendVerificationCodeRequestIdentifier)`

SetIdentifier sets Identifier field to given value.


### GetInteractionEvent

`func (o *CreateAndSendVerificationCodeRequest) GetInteractionEvent() string`

GetInteractionEvent returns the InteractionEvent field if non-nil, zero value otherwise.

### GetInteractionEventOk

`func (o *CreateAndSendVerificationCodeRequest) GetInteractionEventOk() (*string, bool)`

GetInteractionEventOk returns a tuple with the InteractionEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionEvent

`func (o *CreateAndSendVerificationCodeRequest) SetInteractionEvent(v string)`

SetInteractionEvent sets InteractionEvent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


