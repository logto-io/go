# ApiInteractionPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**Identifier** | Pointer to [**ApiInteractionPutRequestIdentifier**](ApiInteractionPutRequestIdentifier.md) |  | [optional] 
**Profile** | Pointer to [**ApiInteractionPutRequestProfile**](ApiInteractionPutRequestProfile.md) |  | [optional] 

## Methods

### NewApiInteractionPutRequest

`func NewApiInteractionPutRequest(event string, ) *ApiInteractionPutRequest`

NewApiInteractionPutRequest instantiates a new ApiInteractionPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInteractionPutRequestWithDefaults

`func NewApiInteractionPutRequestWithDefaults() *ApiInteractionPutRequest`

NewApiInteractionPutRequestWithDefaults instantiates a new ApiInteractionPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *ApiInteractionPutRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ApiInteractionPutRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ApiInteractionPutRequest) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetIdentifier

`func (o *ApiInteractionPutRequest) GetIdentifier() ApiInteractionPutRequestIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ApiInteractionPutRequest) GetIdentifierOk() (*ApiInteractionPutRequestIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ApiInteractionPutRequest) SetIdentifier(v ApiInteractionPutRequestIdentifier)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ApiInteractionPutRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetProfile

`func (o *ApiInteractionPutRequest) GetProfile() ApiInteractionPutRequestProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ApiInteractionPutRequest) GetProfileOk() (*ApiInteractionPutRequestProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ApiInteractionPutRequest) SetProfile(v ApiInteractionPutRequestProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ApiInteractionPutRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


