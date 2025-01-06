# AssertSingleSignOnSamlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelayState** | Pointer to **string** | SAML standard parameter that will be transmitted between the identity provider and the service provider. It will be used as the session ID (jti) of the user&#39;s Logto authentication session. This API will use this session ID to retrieve the SSO connector authentication session from the database. | [optional] 
**SAMLResponse** | **string** | The SAML assertion response from the identity provider (IdP). | 

## Methods

### NewAssertSingleSignOnSamlRequest

`func NewAssertSingleSignOnSamlRequest(sAMLResponse string, ) *AssertSingleSignOnSamlRequest`

NewAssertSingleSignOnSamlRequest instantiates a new AssertSingleSignOnSamlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssertSingleSignOnSamlRequestWithDefaults

`func NewAssertSingleSignOnSamlRequestWithDefaults() *AssertSingleSignOnSamlRequest`

NewAssertSingleSignOnSamlRequestWithDefaults instantiates a new AssertSingleSignOnSamlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelayState

`func (o *AssertSingleSignOnSamlRequest) GetRelayState() string`

GetRelayState returns the RelayState field if non-nil, zero value otherwise.

### GetRelayStateOk

`func (o *AssertSingleSignOnSamlRequest) GetRelayStateOk() (*string, bool)`

GetRelayStateOk returns a tuple with the RelayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayState

`func (o *AssertSingleSignOnSamlRequest) SetRelayState(v string)`

SetRelayState sets RelayState field to given value.

### HasRelayState

`func (o *AssertSingleSignOnSamlRequest) HasRelayState() bool`

HasRelayState returns a boolean if a field has been set.

### GetSAMLResponse

`func (o *AssertSingleSignOnSamlRequest) GetSAMLResponse() string`

GetSAMLResponse returns the SAMLResponse field if non-nil, zero value otherwise.

### GetSAMLResponseOk

`func (o *AssertSingleSignOnSamlRequest) GetSAMLResponseOk() (*string, bool)`

GetSAMLResponseOk returns a tuple with the SAMLResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAMLResponse

`func (o *AssertSingleSignOnSamlRequest) SetSAMLResponse(v string)`

SetSAMLResponse sets SAMLResponse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


