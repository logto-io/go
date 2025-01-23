# VerifyVerificationBySocialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorData** | **map[string]interface{}** | A json object constructed from the url query params returned by the social platform. Typically it contains &#x60;code&#x60;, &#x60;state&#x60; and &#x60;redirectUri&#x60; fields. | 
**VerificationRecordId** | **string** |  | 
**VerificationId** | Pointer to **interface{}** | The verification ID of the SocialVerification record. | [optional] 

## Methods

### NewVerifyVerificationBySocialRequest

`func NewVerifyVerificationBySocialRequest(connectorData map[string]interface{}, verificationRecordId string, ) *VerifyVerificationBySocialRequest`

NewVerifyVerificationBySocialRequest instantiates a new VerifyVerificationBySocialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyVerificationBySocialRequestWithDefaults

`func NewVerifyVerificationBySocialRequestWithDefaults() *VerifyVerificationBySocialRequest`

NewVerifyVerificationBySocialRequestWithDefaults instantiates a new VerifyVerificationBySocialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorData

`func (o *VerifyVerificationBySocialRequest) GetConnectorData() map[string]interface{}`

GetConnectorData returns the ConnectorData field if non-nil, zero value otherwise.

### GetConnectorDataOk

`func (o *VerifyVerificationBySocialRequest) GetConnectorDataOk() (*map[string]interface{}, bool)`

GetConnectorDataOk returns a tuple with the ConnectorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorData

`func (o *VerifyVerificationBySocialRequest) SetConnectorData(v map[string]interface{})`

SetConnectorData sets ConnectorData field to given value.


### GetVerificationRecordId

`func (o *VerifyVerificationBySocialRequest) GetVerificationRecordId() string`

GetVerificationRecordId returns the VerificationRecordId field if non-nil, zero value otherwise.

### GetVerificationRecordIdOk

`func (o *VerifyVerificationBySocialRequest) GetVerificationRecordIdOk() (*string, bool)`

GetVerificationRecordIdOk returns a tuple with the VerificationRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRecordId

`func (o *VerifyVerificationBySocialRequest) SetVerificationRecordId(v string)`

SetVerificationRecordId sets VerificationRecordId field to given value.


### GetVerificationId

`func (o *VerifyVerificationBySocialRequest) GetVerificationId() interface{}`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *VerifyVerificationBySocialRequest) GetVerificationIdOk() (*interface{}, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *VerifyVerificationBySocialRequest) SetVerificationId(v interface{})`

SetVerificationId sets VerificationId field to given value.

### HasVerificationId

`func (o *VerifyVerificationBySocialRequest) HasVerificationId() bool`

HasVerificationId returns a boolean if a field has been set.

### SetVerificationIdNil

`func (o *VerifyVerificationBySocialRequest) SetVerificationIdNil(b bool)`

 SetVerificationIdNil sets the value for VerificationId to be an explicit nil

### UnsetVerificationId
`func (o *VerifyVerificationBySocialRequest) UnsetVerificationId()`

UnsetVerificationId ensures that no value is present for VerificationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


