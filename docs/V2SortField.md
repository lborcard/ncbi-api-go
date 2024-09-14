# V2SortField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to [**V2SortDirection**](V2SortDirection.md) |  | [optional] [default to V2SORTDIRECTION_UNSPECIFIED]

## Methods

### NewV2SortField

`func NewV2SortField() *V2SortField`

NewV2SortField instantiates a new V2SortField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2SortFieldWithDefaults

`func NewV2SortFieldWithDefaults() *V2SortField`

NewV2SortFieldWithDefaults instantiates a new V2SortField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *V2SortField) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *V2SortField) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *V2SortField) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *V2SortField) HasField() bool`

HasField returns a boolean if a field has been set.

### GetDirection

`func (o *V2SortField) GetDirection() V2SortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *V2SortField) GetDirectionOk() (*V2SortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *V2SortField) SetDirection(v V2SortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *V2SortField) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


