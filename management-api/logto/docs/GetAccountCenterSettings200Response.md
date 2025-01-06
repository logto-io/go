# GetAccountCenterSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Id** | **string** |  | 
**Enabled** | **bool** |  | 
**Fields** | [**GetAccountCenterSettings200ResponseFields**](GetAccountCenterSettings200ResponseFields.md) |  | 

## Methods

### NewGetAccountCenterSettings200Response

`func NewGetAccountCenterSettings200Response(tenantId string, id string, enabled bool, fields GetAccountCenterSettings200ResponseFields, ) *GetAccountCenterSettings200Response`

NewGetAccountCenterSettings200Response instantiates a new GetAccountCenterSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountCenterSettings200ResponseWithDefaults

`func NewGetAccountCenterSettings200ResponseWithDefaults() *GetAccountCenterSettings200Response`

NewGetAccountCenterSettings200ResponseWithDefaults instantiates a new GetAccountCenterSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *GetAccountCenterSettings200Response) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GetAccountCenterSettings200Response) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GetAccountCenterSettings200Response) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *GetAccountCenterSettings200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccountCenterSettings200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccountCenterSettings200Response) SetId(v string)`

SetId sets Id field to given value.


### GetEnabled

`func (o *GetAccountCenterSettings200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetAccountCenterSettings200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetAccountCenterSettings200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFields

`func (o *GetAccountCenterSettings200Response) GetFields() GetAccountCenterSettings200ResponseFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GetAccountCenterSettings200Response) GetFieldsOk() (*GetAccountCenterSettings200ResponseFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GetAccountCenterSettings200Response) SetFields(v GetAccountCenterSettings200ResponseFields)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


