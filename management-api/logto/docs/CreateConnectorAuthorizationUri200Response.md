# CreateConnectorAuthorizationUri200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectTo** | **string** |  | 
**RedirectUri** | Pointer to **interface{}** | The URI to navigate for authentication and authorization in the connected social identity provider. | [optional] 

## Methods

### NewCreateConnectorAuthorizationUri200Response

`func NewCreateConnectorAuthorizationUri200Response(redirectTo string, ) *CreateConnectorAuthorizationUri200Response`

NewCreateConnectorAuthorizationUri200Response instantiates a new CreateConnectorAuthorizationUri200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectorAuthorizationUri200ResponseWithDefaults

`func NewCreateConnectorAuthorizationUri200ResponseWithDefaults() *CreateConnectorAuthorizationUri200Response`

NewCreateConnectorAuthorizationUri200ResponseWithDefaults instantiates a new CreateConnectorAuthorizationUri200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectTo

`func (o *CreateConnectorAuthorizationUri200Response) GetRedirectTo() string`

GetRedirectTo returns the RedirectTo field if non-nil, zero value otherwise.

### GetRedirectToOk

`func (o *CreateConnectorAuthorizationUri200Response) GetRedirectToOk() (*string, bool)`

GetRedirectToOk returns a tuple with the RedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTo

`func (o *CreateConnectorAuthorizationUri200Response) SetRedirectTo(v string)`

SetRedirectTo sets RedirectTo field to given value.


### GetRedirectUri

`func (o *CreateConnectorAuthorizationUri200Response) GetRedirectUri() interface{}`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *CreateConnectorAuthorizationUri200Response) GetRedirectUriOk() (*interface{}, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *CreateConnectorAuthorizationUri200Response) SetRedirectUri(v interface{})`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *CreateConnectorAuthorizationUri200Response) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### SetRedirectUriNil

`func (o *CreateConnectorAuthorizationUri200Response) SetRedirectUriNil(b bool)`

 SetRedirectUriNil sets the value for RedirectUri to be an explicit nil

### UnsetRedirectUri
`func (o *CreateConnectorAuthorizationUri200Response) UnsetRedirectUri()`

UnsetRedirectUri ensures that no value is present for RedirectUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


