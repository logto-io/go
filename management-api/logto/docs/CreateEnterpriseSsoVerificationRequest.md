# CreateEnterpriseSsoVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The state parameter to pass to the SSO connector. | 
**RedirectUri** | **string** | The URI to redirect the user after the SSO authorization is completed. | 

## Methods

### NewCreateEnterpriseSsoVerificationRequest

`func NewCreateEnterpriseSsoVerificationRequest(state string, redirectUri string, ) *CreateEnterpriseSsoVerificationRequest`

NewCreateEnterpriseSsoVerificationRequest instantiates a new CreateEnterpriseSsoVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnterpriseSsoVerificationRequestWithDefaults

`func NewCreateEnterpriseSsoVerificationRequestWithDefaults() *CreateEnterpriseSsoVerificationRequest`

NewCreateEnterpriseSsoVerificationRequestWithDefaults instantiates a new CreateEnterpriseSsoVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CreateEnterpriseSsoVerificationRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateEnterpriseSsoVerificationRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateEnterpriseSsoVerificationRequest) SetState(v string)`

SetState sets State field to given value.


### GetRedirectUri

`func (o *CreateEnterpriseSsoVerificationRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *CreateEnterpriseSsoVerificationRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *CreateEnterpriseSsoVerificationRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


