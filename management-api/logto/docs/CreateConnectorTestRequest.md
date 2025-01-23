# CreateConnectorTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Phone number to send test message to. If this is set, email will be ignored. | [optional] 
**Email** | Pointer to **string** | Email address to send test message to. If phone is set, this will be ignored. | [optional] 
**Config** | **map[string]interface{}** | Connector configuration object for testing. | 

## Methods

### NewCreateConnectorTestRequest

`func NewCreateConnectorTestRequest(config map[string]interface{}, ) *CreateConnectorTestRequest`

NewCreateConnectorTestRequest instantiates a new CreateConnectorTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectorTestRequestWithDefaults

`func NewCreateConnectorTestRequestWithDefaults() *CreateConnectorTestRequest`

NewCreateConnectorTestRequestWithDefaults instantiates a new CreateConnectorTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *CreateConnectorTestRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CreateConnectorTestRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CreateConnectorTestRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CreateConnectorTestRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *CreateConnectorTestRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateConnectorTestRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateConnectorTestRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateConnectorTestRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetConfig

`func (o *CreateConnectorTestRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateConnectorTestRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateConnectorTestRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


