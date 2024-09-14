# V2reportsMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**V2reportsError**](V2reportsError.md) |  | [optional] 
**Warning** | Pointer to [**V2reportsWarning**](V2reportsWarning.md) |  | [optional] 

## Methods

### NewV2reportsMessage

`func NewV2reportsMessage() *V2reportsMessage`

NewV2reportsMessage instantiates a new V2reportsMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsMessageWithDefaults

`func NewV2reportsMessageWithDefaults() *V2reportsMessage`

NewV2reportsMessageWithDefaults instantiates a new V2reportsMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V2reportsMessage) GetError() V2reportsError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2reportsMessage) GetErrorOk() (*V2reportsError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2reportsMessage) SetError(v V2reportsError)`

SetError sets Error field to given value.

### HasError

`func (o *V2reportsMessage) HasError() bool`

HasError returns a boolean if a field has been set.

### GetWarning

`func (o *V2reportsMessage) GetWarning() V2reportsWarning`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *V2reportsMessage) GetWarningOk() (*V2reportsWarning, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *V2reportsMessage) SetWarning(v V2reportsWarning)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *V2reportsMessage) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


