# IdentifyUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | Pointer to **string** | The ID of the verification record used to identify the user. &lt;br/&gt;- &#x60;SignIn&#x60; and &#x60;ForgotPassword&#x60; interactions: Required to verify the user&#39;s identity. &lt;br/&gt;- &#x60;Register&#x60; interaction: Optional. If provided, it updates the profile data with the verification record before account creation. If omitted, the account is created using existing profile data in the current interaction. | [optional] 
**LinkSocialIdentity** | Pointer to **bool** | Applies to the SignIn interaction only, and is used when a SocialVerification type verificationId is provided. &lt;br/&gt;- If &#x60;true&#x60;, the user is identified using the verified email or phone number from the social identity provider, and the social identity is linked to the user&#39;s account. &lt;br/&gt;- If &#x60;false&#x60; or not provided, the API identifies the user solely through the social identity. &lt;br/&gt; This parameters is used for linking a non-existing social identity to a related user account that can be identified through the verified email or phone number. | [optional] 

## Methods

### NewIdentifyUserRequest

`func NewIdentifyUserRequest() *IdentifyUserRequest`

NewIdentifyUserRequest instantiates a new IdentifyUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifyUserRequestWithDefaults

`func NewIdentifyUserRequestWithDefaults() *IdentifyUserRequest`

NewIdentifyUserRequestWithDefaults instantiates a new IdentifyUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *IdentifyUserRequest) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *IdentifyUserRequest) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *IdentifyUserRequest) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.

### HasVerificationId

`func (o *IdentifyUserRequest) HasVerificationId() bool`

HasVerificationId returns a boolean if a field has been set.

### GetLinkSocialIdentity

`func (o *IdentifyUserRequest) GetLinkSocialIdentity() bool`

GetLinkSocialIdentity returns the LinkSocialIdentity field if non-nil, zero value otherwise.

### GetLinkSocialIdentityOk

`func (o *IdentifyUserRequest) GetLinkSocialIdentityOk() (*bool, bool)`

GetLinkSocialIdentityOk returns a tuple with the LinkSocialIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSocialIdentity

`func (o *IdentifyUserRequest) SetLinkSocialIdentity(v bool)`

SetLinkSocialIdentity sets LinkSocialIdentity field to given value.

### HasLinkSocialIdentity

`func (o *IdentifyUserRequest) HasLinkSocialIdentity() bool`

HasLinkSocialIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


