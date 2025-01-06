# ListLogs200ResponseInnerPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Result** | **string** |  | 
**Error** | Pointer to [**ListLogs200ResponseInnerPayloadError**](ListLogs200ResponseInnerPayloadError.md) |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewListLogs200ResponseInnerPayload

`func NewListLogs200ResponseInnerPayload(key string, result string, ) *ListLogs200ResponseInnerPayload`

NewListLogs200ResponseInnerPayload instantiates a new ListLogs200ResponseInnerPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogs200ResponseInnerPayloadWithDefaults

`func NewListLogs200ResponseInnerPayloadWithDefaults() *ListLogs200ResponseInnerPayload`

NewListLogs200ResponseInnerPayloadWithDefaults instantiates a new ListLogs200ResponseInnerPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ListLogs200ResponseInnerPayload) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ListLogs200ResponseInnerPayload) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ListLogs200ResponseInnerPayload) SetKey(v string)`

SetKey sets Key field to given value.


### GetResult

`func (o *ListLogs200ResponseInnerPayload) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListLogs200ResponseInnerPayload) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListLogs200ResponseInnerPayload) SetResult(v string)`

SetResult sets Result field to given value.


### GetError

`func (o *ListLogs200ResponseInnerPayload) GetError() ListLogs200ResponseInnerPayloadError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListLogs200ResponseInnerPayload) GetErrorOk() (*ListLogs200ResponseInnerPayloadError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListLogs200ResponseInnerPayload) SetError(v ListLogs200ResponseInnerPayloadError)`

SetError sets Error field to given value.

### HasError

`func (o *ListLogs200ResponseInnerPayload) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIp

`func (o *ListLogs200ResponseInnerPayload) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ListLogs200ResponseInnerPayload) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ListLogs200ResponseInnerPayload) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ListLogs200ResponseInnerPayload) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUserAgent

`func (o *ListLogs200ResponseInnerPayload) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ListLogs200ResponseInnerPayload) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ListLogs200ResponseInnerPayload) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ListLogs200ResponseInnerPayload) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserId

`func (o *ListLogs200ResponseInnerPayload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListLogs200ResponseInnerPayload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListLogs200ResponseInnerPayload) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ListLogs200ResponseInnerPayload) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetApplicationId

`func (o *ListLogs200ResponseInnerPayload) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ListLogs200ResponseInnerPayload) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ListLogs200ResponseInnerPayload) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ListLogs200ResponseInnerPayload) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetSessionId

`func (o *ListLogs200ResponseInnerPayload) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ListLogs200ResponseInnerPayload) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ListLogs200ResponseInnerPayload) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ListLogs200ResponseInnerPayload) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetParams

`func (o *ListLogs200ResponseInnerPayload) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListLogs200ResponseInnerPayload) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListLogs200ResponseInnerPayload) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *ListLogs200ResponseInnerPayload) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


