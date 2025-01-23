# CreateSocialVerification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationUri** | **string** | The social authorization URI. | 
**VerificationId** | **string** | The unique verification ID of the newly created SocialVerification record. The &#x60;verificationId&#x60; is required when verifying the social authorization response. | 

## Methods

### NewCreateSocialVerification200Response

`func NewCreateSocialVerification200Response(authorizationUri string, verificationId string, ) *CreateSocialVerification200Response`

NewCreateSocialVerification200Response instantiates a new CreateSocialVerification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSocialVerification200ResponseWithDefaults

`func NewCreateSocialVerification200ResponseWithDefaults() *CreateSocialVerification200Response`

NewCreateSocialVerification200ResponseWithDefaults instantiates a new CreateSocialVerification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationUri

`func (o *CreateSocialVerification200Response) GetAuthorizationUri() string`

GetAuthorizationUri returns the AuthorizationUri field if non-nil, zero value otherwise.

### GetAuthorizationUriOk

`func (o *CreateSocialVerification200Response) GetAuthorizationUriOk() (*string, bool)`

GetAuthorizationUriOk returns a tuple with the AuthorizationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUri

`func (o *CreateSocialVerification200Response) SetAuthorizationUri(v string)`

SetAuthorizationUri sets AuthorizationUri field to given value.


### GetVerificationId

`func (o *CreateSocialVerification200Response) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *CreateSocialVerification200Response) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *CreateSocialVerification200Response) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


