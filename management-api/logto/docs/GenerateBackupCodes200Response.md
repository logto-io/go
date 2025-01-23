# GenerateBackupCodes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationId** | **string** | The unique verification ID of the newly created BackupCode verification record. This ID is required when adding the backup codes to the user profile via the Profile API. | 
**Codes** | **[]string** | The generated backup codes. | 

## Methods

### NewGenerateBackupCodes200Response

`func NewGenerateBackupCodes200Response(verificationId string, codes []string, ) *GenerateBackupCodes200Response`

NewGenerateBackupCodes200Response instantiates a new GenerateBackupCodes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateBackupCodes200ResponseWithDefaults

`func NewGenerateBackupCodes200ResponseWithDefaults() *GenerateBackupCodes200Response`

NewGenerateBackupCodes200ResponseWithDefaults instantiates a new GenerateBackupCodes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationId

`func (o *GenerateBackupCodes200Response) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *GenerateBackupCodes200Response) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *GenerateBackupCodes200Response) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.


### GetCodes

`func (o *GenerateBackupCodes200Response) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *GenerateBackupCodes200Response) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *GenerateBackupCodes200Response) SetCodes(v []string)`

SetCodes sets Codes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


