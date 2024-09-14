# V2reportsWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneWarningCode** | Pointer to [**V2reportsWarningGeneWarningCode**](V2reportsWarningGeneWarningCode.md) |  | [optional] [default to V2REPORTSWARNINGGENEWARNINGCODE_UNKNOWN_GENE_WARNING_CODE]
**Reason** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ReplacedId** | Pointer to [**V2reportsWarningReplacedId**](V2reportsWarningReplacedId.md) |  | [optional] 
**UnrecognizedIdentifier** | Pointer to **string** |  | [optional] 

## Methods

### NewV2reportsWarning

`func NewV2reportsWarning() *V2reportsWarning`

NewV2reportsWarning instantiates a new V2reportsWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsWarningWithDefaults

`func NewV2reportsWarningWithDefaults() *V2reportsWarning`

NewV2reportsWarningWithDefaults instantiates a new V2reportsWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneWarningCode

`func (o *V2reportsWarning) GetGeneWarningCode() V2reportsWarningGeneWarningCode`

GetGeneWarningCode returns the GeneWarningCode field if non-nil, zero value otherwise.

### GetGeneWarningCodeOk

`func (o *V2reportsWarning) GetGeneWarningCodeOk() (*V2reportsWarningGeneWarningCode, bool)`

GetGeneWarningCodeOk returns a tuple with the GeneWarningCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneWarningCode

`func (o *V2reportsWarning) SetGeneWarningCode(v V2reportsWarningGeneWarningCode)`

SetGeneWarningCode sets GeneWarningCode field to given value.

### HasGeneWarningCode

`func (o *V2reportsWarning) HasGeneWarningCode() bool`

HasGeneWarningCode returns a boolean if a field has been set.

### GetReason

`func (o *V2reportsWarning) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V2reportsWarning) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V2reportsWarning) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V2reportsWarning) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *V2reportsWarning) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2reportsWarning) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2reportsWarning) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V2reportsWarning) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReplacedId

`func (o *V2reportsWarning) GetReplacedId() V2reportsWarningReplacedId`

GetReplacedId returns the ReplacedId field if non-nil, zero value otherwise.

### GetReplacedIdOk

`func (o *V2reportsWarning) GetReplacedIdOk() (*V2reportsWarningReplacedId, bool)`

GetReplacedIdOk returns a tuple with the ReplacedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedId

`func (o *V2reportsWarning) SetReplacedId(v V2reportsWarningReplacedId)`

SetReplacedId sets ReplacedId field to given value.

### HasReplacedId

`func (o *V2reportsWarning) HasReplacedId() bool`

HasReplacedId returns a boolean if a field has been set.

### GetUnrecognizedIdentifier

`func (o *V2reportsWarning) GetUnrecognizedIdentifier() string`

GetUnrecognizedIdentifier returns the UnrecognizedIdentifier field if non-nil, zero value otherwise.

### GetUnrecognizedIdentifierOk

`func (o *V2reportsWarning) GetUnrecognizedIdentifierOk() (*string, bool)`

GetUnrecognizedIdentifierOk returns a tuple with the UnrecognizedIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrecognizedIdentifier

`func (o *V2reportsWarning) SetUnrecognizedIdentifier(v string)`

SetUnrecognizedIdentifier sets UnrecognizedIdentifier field to given value.

### HasUnrecognizedIdentifier

`func (o *V2reportsWarning) HasUnrecognizedIdentifier() bool`

HasUnrecognizedIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


