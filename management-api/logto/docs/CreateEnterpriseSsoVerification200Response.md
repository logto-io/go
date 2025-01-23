# CreateEnterpriseSsoVerification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationUri** | **string** | The SSO authorization URI. | 
**VerificationId** | **string** | The unique verification ID of the newly created EnterpriseSSO verification record. The &#x60;verificationId&#x60; is required when verifying the SSO authorization response. | 

## Methods

### NewCreateEnterpriseSsoVerification200Response

`func NewCreateEnterpriseSsoVerification200Response(authorizationUri string, verificationId string, ) *CreateEnterpriseSsoVerification200Response`

NewCreateEnterpriseSsoVerification200Response instantiates a new CreateEnterpriseSsoVerification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnterpriseSsoVerification200ResponseWithDefaults

`func NewCreateEnterpriseSsoVerification200ResponseWithDefaults() *CreateEnterpriseSsoVerification200Response`

NewCreateEnterpriseSsoVerification200ResponseWithDefaults instantiates a new CreateEnterpriseSsoVerification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationUri

`func (o *CreateEnterpriseSsoVerification200Response) GetAuthorizationUri() string`

GetAuthorizationUri returns the AuthorizationUri field if non-nil, zero value otherwise.

### GetAuthorizationUriOk

`func (o *CreateEnterpriseSsoVerification200Response) GetAuthorizationUriOk() (*string, bool)`

GetAuthorizationUriOk returns a tuple with the AuthorizationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUri

`func (o *CreateEnterpriseSsoVerification200Response) SetAuthorizationUri(v string)`

SetAuthorizationUri sets AuthorizationUri field to given value.


### GetVerificationId

`func (o *CreateEnterpriseSsoVerification200Response) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *CreateEnterpriseSsoVerification200Response) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *CreateEnterpriseSsoVerification200Response) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


