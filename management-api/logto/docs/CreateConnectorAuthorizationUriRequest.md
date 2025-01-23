# CreateConnectorAuthorizationUriRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | A random string generated on the client side to prevent CSRF (Cross-Site Request Forgery) attacks. | 
**RedirectUri** | **string** | The URI to navigate back to after the user is authenticated by the connected social identity provider and has granted access to the connector. | 

## Methods

### NewCreateConnectorAuthorizationUriRequest

`func NewCreateConnectorAuthorizationUriRequest(state string, redirectUri string, ) *CreateConnectorAuthorizationUriRequest`

NewCreateConnectorAuthorizationUriRequest instantiates a new CreateConnectorAuthorizationUriRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectorAuthorizationUriRequestWithDefaults

`func NewCreateConnectorAuthorizationUriRequestWithDefaults() *CreateConnectorAuthorizationUriRequest`

NewCreateConnectorAuthorizationUriRequestWithDefaults instantiates a new CreateConnectorAuthorizationUriRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CreateConnectorAuthorizationUriRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateConnectorAuthorizationUriRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateConnectorAuthorizationUriRequest) SetState(v string)`

SetState sets State field to given value.


### GetRedirectUri

`func (o *CreateConnectorAuthorizationUriRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *CreateConnectorAuthorizationUriRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *CreateConnectorAuthorizationUriRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


