# CreateSocialVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The state parameter to pass to the social connector. | 
**RedirectUri** | **string** | The URI to redirect the user after the social authorization is completed. | 

## Methods

### NewCreateSocialVerificationRequest

`func NewCreateSocialVerificationRequest(state string, redirectUri string, ) *CreateSocialVerificationRequest`

NewCreateSocialVerificationRequest instantiates a new CreateSocialVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSocialVerificationRequestWithDefaults

`func NewCreateSocialVerificationRequestWithDefaults() *CreateSocialVerificationRequest`

NewCreateSocialVerificationRequestWithDefaults instantiates a new CreateSocialVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CreateSocialVerificationRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateSocialVerificationRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateSocialVerificationRequest) SetState(v string)`

SetState sets State field to given value.


### GetRedirectUri

`func (o *CreateSocialVerificationRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *CreateSocialVerificationRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *CreateSocialVerificationRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


