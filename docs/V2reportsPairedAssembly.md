# V2reportsPairedAssembly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accession** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**V2reportsAssemblyStatus**](V2reportsAssemblyStatus.md) |  | [optional] [default to V2REPORTSASSEMBLYSTATUS_ASSEMBLY_STATUS_UNKNOWN]
**AnnotationName** | Pointer to **string** |  | [optional] 
**OnlyGenbank** | Pointer to **string** |  | [optional] 
**OnlyRefseq** | Pointer to **string** |  | [optional] 
**Changed** | Pointer to **string** |  | [optional] 
**ManualDiff** | Pointer to **string** |  | [optional] 

## Methods

### NewV2reportsPairedAssembly

`func NewV2reportsPairedAssembly() *V2reportsPairedAssembly`

NewV2reportsPairedAssembly instantiates a new V2reportsPairedAssembly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsPairedAssemblyWithDefaults

`func NewV2reportsPairedAssemblyWithDefaults() *V2reportsPairedAssembly`

NewV2reportsPairedAssemblyWithDefaults instantiates a new V2reportsPairedAssembly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccession

`func (o *V2reportsPairedAssembly) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V2reportsPairedAssembly) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V2reportsPairedAssembly) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V2reportsPairedAssembly) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetStatus

`func (o *V2reportsPairedAssembly) GetStatus() V2reportsAssemblyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2reportsPairedAssembly) GetStatusOk() (*V2reportsAssemblyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2reportsPairedAssembly) SetStatus(v V2reportsAssemblyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V2reportsPairedAssembly) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAnnotationName

`func (o *V2reportsPairedAssembly) GetAnnotationName() string`

GetAnnotationName returns the AnnotationName field if non-nil, zero value otherwise.

### GetAnnotationNameOk

`func (o *V2reportsPairedAssembly) GetAnnotationNameOk() (*string, bool)`

GetAnnotationNameOk returns a tuple with the AnnotationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationName

`func (o *V2reportsPairedAssembly) SetAnnotationName(v string)`

SetAnnotationName sets AnnotationName field to given value.

### HasAnnotationName

`func (o *V2reportsPairedAssembly) HasAnnotationName() bool`

HasAnnotationName returns a boolean if a field has been set.

### GetOnlyGenbank

`func (o *V2reportsPairedAssembly) GetOnlyGenbank() string`

GetOnlyGenbank returns the OnlyGenbank field if non-nil, zero value otherwise.

### GetOnlyGenbankOk

`func (o *V2reportsPairedAssembly) GetOnlyGenbankOk() (*string, bool)`

GetOnlyGenbankOk returns a tuple with the OnlyGenbank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyGenbank

`func (o *V2reportsPairedAssembly) SetOnlyGenbank(v string)`

SetOnlyGenbank sets OnlyGenbank field to given value.

### HasOnlyGenbank

`func (o *V2reportsPairedAssembly) HasOnlyGenbank() bool`

HasOnlyGenbank returns a boolean if a field has been set.

### GetOnlyRefseq

`func (o *V2reportsPairedAssembly) GetOnlyRefseq() string`

GetOnlyRefseq returns the OnlyRefseq field if non-nil, zero value otherwise.

### GetOnlyRefseqOk

`func (o *V2reportsPairedAssembly) GetOnlyRefseqOk() (*string, bool)`

GetOnlyRefseqOk returns a tuple with the OnlyRefseq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyRefseq

`func (o *V2reportsPairedAssembly) SetOnlyRefseq(v string)`

SetOnlyRefseq sets OnlyRefseq field to given value.

### HasOnlyRefseq

`func (o *V2reportsPairedAssembly) HasOnlyRefseq() bool`

HasOnlyRefseq returns a boolean if a field has been set.

### GetChanged

`func (o *V2reportsPairedAssembly) GetChanged() string`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *V2reportsPairedAssembly) GetChangedOk() (*string, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *V2reportsPairedAssembly) SetChanged(v string)`

SetChanged sets Changed field to given value.

### HasChanged

`func (o *V2reportsPairedAssembly) HasChanged() bool`

HasChanged returns a boolean if a field has been set.

### GetManualDiff

`func (o *V2reportsPairedAssembly) GetManualDiff() string`

GetManualDiff returns the ManualDiff field if non-nil, zero value otherwise.

### GetManualDiffOk

`func (o *V2reportsPairedAssembly) GetManualDiffOk() (*string, bool)`

GetManualDiffOk returns a tuple with the ManualDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualDiff

`func (o *V2reportsPairedAssembly) SetManualDiff(v string)`

SetManualDiff sets ManualDiff field to given value.

### HasManualDiff

`func (o *V2reportsPairedAssembly) HasManualDiff() bool`

HasManualDiff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


