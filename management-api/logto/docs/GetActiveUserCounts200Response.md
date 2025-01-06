# GetActiveUserCounts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DauCurve** | [**[]GetActiveUserCounts200ResponseDauCurveInner**](GetActiveUserCounts200ResponseDauCurveInner.md) |  | 
**Dau** | [**GetNewUserCounts200ResponseToday**](GetNewUserCounts200ResponseToday.md) |  | 
**Wau** | [**GetNewUserCounts200ResponseToday**](GetNewUserCounts200ResponseToday.md) |  | 
**Mau** | [**GetNewUserCounts200ResponseToday**](GetNewUserCounts200ResponseToday.md) |  | 

## Methods

### NewGetActiveUserCounts200Response

`func NewGetActiveUserCounts200Response(dauCurve []GetActiveUserCounts200ResponseDauCurveInner, dau GetNewUserCounts200ResponseToday, wau GetNewUserCounts200ResponseToday, mau GetNewUserCounts200ResponseToday, ) *GetActiveUserCounts200Response`

NewGetActiveUserCounts200Response instantiates a new GetActiveUserCounts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActiveUserCounts200ResponseWithDefaults

`func NewGetActiveUserCounts200ResponseWithDefaults() *GetActiveUserCounts200Response`

NewGetActiveUserCounts200ResponseWithDefaults instantiates a new GetActiveUserCounts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDauCurve

`func (o *GetActiveUserCounts200Response) GetDauCurve() []GetActiveUserCounts200ResponseDauCurveInner`

GetDauCurve returns the DauCurve field if non-nil, zero value otherwise.

### GetDauCurveOk

`func (o *GetActiveUserCounts200Response) GetDauCurveOk() (*[]GetActiveUserCounts200ResponseDauCurveInner, bool)`

GetDauCurveOk returns a tuple with the DauCurve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDauCurve

`func (o *GetActiveUserCounts200Response) SetDauCurve(v []GetActiveUserCounts200ResponseDauCurveInner)`

SetDauCurve sets DauCurve field to given value.


### GetDau

`func (o *GetActiveUserCounts200Response) GetDau() GetNewUserCounts200ResponseToday`

GetDau returns the Dau field if non-nil, zero value otherwise.

### GetDauOk

`func (o *GetActiveUserCounts200Response) GetDauOk() (*GetNewUserCounts200ResponseToday, bool)`

GetDauOk returns a tuple with the Dau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDau

`func (o *GetActiveUserCounts200Response) SetDau(v GetNewUserCounts200ResponseToday)`

SetDau sets Dau field to given value.


### GetWau

`func (o *GetActiveUserCounts200Response) GetWau() GetNewUserCounts200ResponseToday`

GetWau returns the Wau field if non-nil, zero value otherwise.

### GetWauOk

`func (o *GetActiveUserCounts200Response) GetWauOk() (*GetNewUserCounts200ResponseToday, bool)`

GetWauOk returns a tuple with the Wau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWau

`func (o *GetActiveUserCounts200Response) SetWau(v GetNewUserCounts200ResponseToday)`

SetWau sets Wau field to given value.


### GetMau

`func (o *GetActiveUserCounts200Response) GetMau() GetNewUserCounts200ResponseToday`

GetMau returns the Mau field if non-nil, zero value otherwise.

### GetMauOk

`func (o *GetActiveUserCounts200Response) GetMauOk() (*GetNewUserCounts200ResponseToday, bool)`

GetMauOk returns a tuple with the Mau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMau

`func (o *GetActiveUserCounts200Response) SetMau(v GetNewUserCounts200ResponseToday)`

SetMau sets Mau field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


