# V2OrganelleSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to [**V2SortDirection**](V2SortDirection.md) |  | [optional] [default to V2SORTDIRECTION_UNSPECIFIED]

## Methods

### NewV2OrganelleSort

`func NewV2OrganelleSort() *V2OrganelleSort`

NewV2OrganelleSort instantiates a new V2OrganelleSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2OrganelleSortWithDefaults

`func NewV2OrganelleSortWithDefaults() *V2OrganelleSort`

NewV2OrganelleSortWithDefaults instantiates a new V2OrganelleSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *V2OrganelleSort) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *V2OrganelleSort) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *V2OrganelleSort) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *V2OrganelleSort) HasField() bool`

HasField returns a boolean if a field has been set.

### GetDirection

`func (o *V2OrganelleSort) GetDirection() V2SortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *V2OrganelleSort) GetDirectionOk() (*V2SortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *V2OrganelleSort) SetDirection(v V2SortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *V2OrganelleSort) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


