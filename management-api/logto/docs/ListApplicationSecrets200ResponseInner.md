# ListApplicationSecrets200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**ApplicationId** | **string** |  | 
**Name** | **string** |  | 
**Value** | **string** |  | 
**CreatedAt** | **float32** |  | 
**ExpiresAt** | **NullableFloat32** |  | 

## Methods

### NewListApplicationSecrets200ResponseInner

`func NewListApplicationSecrets200ResponseInner(tenantId string, applicationId string, name string, value string, createdAt float32, expiresAt NullableFloat32, ) *ListApplicationSecrets200ResponseInner`

NewListApplicationSecrets200ResponseInner instantiates a new ListApplicationSecrets200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApplicationSecrets200ResponseInnerWithDefaults

`func NewListApplicationSecrets200ResponseInnerWithDefaults() *ListApplicationSecrets200ResponseInner`

NewListApplicationSecrets200ResponseInnerWithDefaults instantiates a new ListApplicationSecrets200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListApplicationSecrets200ResponseInner) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListApplicationSecrets200ResponseInner) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListApplicationSecrets200ResponseInner) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetApplicationId

`func (o *ListApplicationSecrets200ResponseInner) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ListApplicationSecrets200ResponseInner) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ListApplicationSecrets200ResponseInner) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetName

`func (o *ListApplicationSecrets200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListApplicationSecrets200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListApplicationSecrets200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ListApplicationSecrets200ResponseInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListApplicationSecrets200ResponseInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListApplicationSecrets200ResponseInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetCreatedAt

`func (o *ListApplicationSecrets200ResponseInner) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListApplicationSecrets200ResponseInner) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListApplicationSecrets200ResponseInner) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *ListApplicationSecrets200ResponseInner) GetExpiresAt() float32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ListApplicationSecrets200ResponseInner) GetExpiresAtOk() (*float32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ListApplicationSecrets200ResponseInner) SetExpiresAt(v float32)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *ListApplicationSecrets200ResponseInner) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ListApplicationSecrets200ResponseInner) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


