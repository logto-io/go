# VerifySocialVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorData** | **map[string]interface{}** | Arbitrary data returned by the social provider to complete the verification process. | 
**VerificationId** | Pointer to **string** | The ID of the Social verification record. | [optional] 

## Methods

### NewVerifySocialVerificationRequest

`func NewVerifySocialVerificationRequest(connectorData map[string]interface{}, ) *VerifySocialVerificationRequest`

NewVerifySocialVerificationRequest instantiates a new VerifySocialVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifySocialVerificationRequestWithDefaults

`func NewVerifySocialVerificationRequestWithDefaults() *VerifySocialVerificationRequest`

NewVerifySocialVerificationRequestWithDefaults instantiates a new VerifySocialVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorData

`func (o *VerifySocialVerificationRequest) GetConnectorData() map[string]interface{}`

GetConnectorData returns the ConnectorData field if non-nil, zero value otherwise.

### GetConnectorDataOk

`func (o *VerifySocialVerificationRequest) GetConnectorDataOk() (*map[string]interface{}, bool)`

GetConnectorDataOk returns a tuple with the ConnectorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorData

`func (o *VerifySocialVerificationRequest) SetConnectorData(v map[string]interface{})`

SetConnectorData sets ConnectorData field to given value.


### GetVerificationId

`func (o *VerifySocialVerificationRequest) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *VerifySocialVerificationRequest) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *VerifySocialVerificationRequest) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.

### HasVerificationId

`func (o *VerifySocialVerificationRequest) HasVerificationId() bool`

HasVerificationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


