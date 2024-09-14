# V2reportsRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Begin** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**Orientation** | Pointer to [**V2reportsOrientation**](V2reportsOrientation.md) |  | [optional] [default to V2REPORTSORIENTATION_NONE]
**Order** | Pointer to **int32** |  | [optional] 
**RibosomalSlippage** | Pointer to **int32** |  | [optional] 

## Methods

### NewV2reportsRange

`func NewV2reportsRange() *V2reportsRange`

NewV2reportsRange instantiates a new V2reportsRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsRangeWithDefaults

`func NewV2reportsRangeWithDefaults() *V2reportsRange`

NewV2reportsRangeWithDefaults instantiates a new V2reportsRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBegin

`func (o *V2reportsRange) GetBegin() string`

GetBegin returns the Begin field if non-nil, zero value otherwise.

### GetBeginOk

`func (o *V2reportsRange) GetBeginOk() (*string, bool)`

GetBeginOk returns a tuple with the Begin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegin

`func (o *V2reportsRange) SetBegin(v string)`

SetBegin sets Begin field to given value.

### HasBegin

`func (o *V2reportsRange) HasBegin() bool`

HasBegin returns a boolean if a field has been set.

### GetEnd

`func (o *V2reportsRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V2reportsRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V2reportsRange) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *V2reportsRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetOrientation

`func (o *V2reportsRange) GetOrientation() V2reportsOrientation`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *V2reportsRange) GetOrientationOk() (*V2reportsOrientation, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *V2reportsRange) SetOrientation(v V2reportsOrientation)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *V2reportsRange) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetOrder

`func (o *V2reportsRange) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *V2reportsRange) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *V2reportsRange) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *V2reportsRange) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetRibosomalSlippage

`func (o *V2reportsRange) GetRibosomalSlippage() int32`

GetRibosomalSlippage returns the RibosomalSlippage field if non-nil, zero value otherwise.

### GetRibosomalSlippageOk

`func (o *V2reportsRange) GetRibosomalSlippageOk() (*int32, bool)`

GetRibosomalSlippageOk returns a tuple with the RibosomalSlippage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRibosomalSlippage

`func (o *V2reportsRange) SetRibosomalSlippage(v int32)`

SetRibosomalSlippage sets RibosomalSlippage field to given value.

### HasRibosomalSlippage

`func (o *V2reportsRange) HasRibosomalSlippage() bool`

HasRibosomalSlippage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


