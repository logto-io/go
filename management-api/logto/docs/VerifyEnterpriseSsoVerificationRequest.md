# VerifyEnterpriseSsoVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorData** | **map[string]interface{}** | Arbitrary data returned by the SSO provider to complete the verification process. | 
**VerificationId** | **string** | The ID of the EnterpriseSSO verification record. | 

## Methods

### NewVerifyEnterpriseSsoVerificationRequest

`func NewVerifyEnterpriseSsoVerificationRequest(connectorData map[string]interface{}, verificationId string, ) *VerifyEnterpriseSsoVerificationRequest`

NewVerifyEnterpriseSsoVerificationRequest instantiates a new VerifyEnterpriseSsoVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyEnterpriseSsoVerificationRequestWithDefaults

`func NewVerifyEnterpriseSsoVerificationRequestWithDefaults() *VerifyEnterpriseSsoVerificationRequest`

NewVerifyEnterpriseSsoVerificationRequestWithDefaults instantiates a new VerifyEnterpriseSsoVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorData

`func (o *VerifyEnterpriseSsoVerificationRequest) GetConnectorData() map[string]interface{}`

GetConnectorData returns the ConnectorData field if non-nil, zero value otherwise.

### GetConnectorDataOk

`func (o *VerifyEnterpriseSsoVerificationRequest) GetConnectorDataOk() (*map[string]interface{}, bool)`

GetConnectorDataOk returns a tuple with the ConnectorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorData

`func (o *VerifyEnterpriseSsoVerificationRequest) SetConnectorData(v map[string]interface{})`

SetConnectorData sets ConnectorData field to given value.


### GetVerificationId

`func (o *VerifyEnterpriseSsoVerificationRequest) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *VerifyEnterpriseSsoVerificationRequest) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *VerifyEnterpriseSsoVerificationRequest) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


