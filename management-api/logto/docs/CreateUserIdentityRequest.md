# CreateUserIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorId** | **string** | The Logto connector ID. | 
**ConnectorData** | **map[string]interface{}** | A json object constructed from the url query params returned by the social platform. Typically it contains &#x60;code&#x60;, &#x60;state&#x60; and &#x60;redirectUri&#x60; fields. | 

## Methods

### NewCreateUserIdentityRequest

`func NewCreateUserIdentityRequest(connectorId string, connectorData map[string]interface{}, ) *CreateUserIdentityRequest`

NewCreateUserIdentityRequest instantiates a new CreateUserIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserIdentityRequestWithDefaults

`func NewCreateUserIdentityRequestWithDefaults() *CreateUserIdentityRequest`

NewCreateUserIdentityRequestWithDefaults instantiates a new CreateUserIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorId

`func (o *CreateUserIdentityRequest) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *CreateUserIdentityRequest) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *CreateUserIdentityRequest) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.


### GetConnectorData

`func (o *CreateUserIdentityRequest) GetConnectorData() map[string]interface{}`

GetConnectorData returns the ConnectorData field if non-nil, zero value otherwise.

### GetConnectorDataOk

`func (o *CreateUserIdentityRequest) GetConnectorDataOk() (*map[string]interface{}, bool)`

GetConnectorDataOk returns a tuple with the ConnectorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorData

`func (o *CreateUserIdentityRequest) SetConnectorData(v map[string]interface{})`

SetConnectorData sets ConnectorData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


