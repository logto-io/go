# RotateOidcKeysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningKeyAlgorithm** | Pointer to **string** | The signing key algorithm the new generated private key is using.  Only applicable when &#x60;keyType&#x60; is &#x60;private-keys&#x60;. | [optional] 

## Methods

### NewRotateOidcKeysRequest

`func NewRotateOidcKeysRequest() *RotateOidcKeysRequest`

NewRotateOidcKeysRequest instantiates a new RotateOidcKeysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateOidcKeysRequestWithDefaults

`func NewRotateOidcKeysRequestWithDefaults() *RotateOidcKeysRequest`

NewRotateOidcKeysRequestWithDefaults instantiates a new RotateOidcKeysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningKeyAlgorithm

`func (o *RotateOidcKeysRequest) GetSigningKeyAlgorithm() string`

GetSigningKeyAlgorithm returns the SigningKeyAlgorithm field if non-nil, zero value otherwise.

### GetSigningKeyAlgorithmOk

`func (o *RotateOidcKeysRequest) GetSigningKeyAlgorithmOk() (*string, bool)`

GetSigningKeyAlgorithmOk returns a tuple with the SigningKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeyAlgorithm

`func (o *RotateOidcKeysRequest) SetSigningKeyAlgorithm(v string)`

SetSigningKeyAlgorithm sets SigningKeyAlgorithm field to given value.

### HasSigningKeyAlgorithm

`func (o *RotateOidcKeysRequest) HasSigningKeyAlgorithm() bool`

HasSigningKeyAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


