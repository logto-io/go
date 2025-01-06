# GetNewUserCounts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Today** | [**GetNewUserCounts200ResponseToday**](GetNewUserCounts200ResponseToday.md) |  | 
**Last7Days** | [**GetNewUserCounts200ResponseToday**](GetNewUserCounts200ResponseToday.md) |  | 

## Methods

### NewGetNewUserCounts200Response

`func NewGetNewUserCounts200Response(today GetNewUserCounts200ResponseToday, last7Days GetNewUserCounts200ResponseToday, ) *GetNewUserCounts200Response`

NewGetNewUserCounts200Response instantiates a new GetNewUserCounts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNewUserCounts200ResponseWithDefaults

`func NewGetNewUserCounts200ResponseWithDefaults() *GetNewUserCounts200Response`

NewGetNewUserCounts200ResponseWithDefaults instantiates a new GetNewUserCounts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToday

`func (o *GetNewUserCounts200Response) GetToday() GetNewUserCounts200ResponseToday`

GetToday returns the Today field if non-nil, zero value otherwise.

### GetTodayOk

`func (o *GetNewUserCounts200Response) GetTodayOk() (*GetNewUserCounts200ResponseToday, bool)`

GetTodayOk returns a tuple with the Today field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToday

`func (o *GetNewUserCounts200Response) SetToday(v GetNewUserCounts200ResponseToday)`

SetToday sets Today field to given value.


### GetLast7Days

`func (o *GetNewUserCounts200Response) GetLast7Days() GetNewUserCounts200ResponseToday`

GetLast7Days returns the Last7Days field if non-nil, zero value otherwise.

### GetLast7DaysOk

`func (o *GetNewUserCounts200Response) GetLast7DaysOk() (*GetNewUserCounts200ResponseToday, bool)`

GetLast7DaysOk returns a tuple with the Last7Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast7Days

`func (o *GetNewUserCounts200Response) SetLast7Days(v GetNewUserCounts200ResponseToday)`

SetLast7Days sets Last7Days field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


