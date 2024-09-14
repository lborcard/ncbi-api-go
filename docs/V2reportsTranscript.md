# V2reportsTranscript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessionVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**Cds** | Pointer to [**V2reportsSeqRangeSet**](V2reportsSeqRangeSet.md) |  | [optional] 
**GenomicLocations** | Pointer to [**[]V2reportsGenomicLocation**](V2reportsGenomicLocation.md) |  | [optional] 
**EnsemblTranscript** | Pointer to **string** |  | [optional] 
**Protein** | Pointer to [**V2reportsProtein**](V2reportsProtein.md) |  | [optional] 
**Type** | Pointer to [**V2reportsTranscriptTranscriptType**](V2reportsTranscriptTranscriptType.md) |  | [optional] [default to V2REPORTSTRANSCRIPTTRANSCRIPTTYPE_UNKNOWN]

## Methods

### NewV2reportsTranscript

`func NewV2reportsTranscript() *V2reportsTranscript`

NewV2reportsTranscript instantiates a new V2reportsTranscript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsTranscriptWithDefaults

`func NewV2reportsTranscriptWithDefaults() *V2reportsTranscript`

NewV2reportsTranscriptWithDefaults instantiates a new V2reportsTranscript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessionVersion

`func (o *V2reportsTranscript) GetAccessionVersion() string`

GetAccessionVersion returns the AccessionVersion field if non-nil, zero value otherwise.

### GetAccessionVersionOk

`func (o *V2reportsTranscript) GetAccessionVersionOk() (*string, bool)`

GetAccessionVersionOk returns a tuple with the AccessionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessionVersion

`func (o *V2reportsTranscript) SetAccessionVersion(v string)`

SetAccessionVersion sets AccessionVersion field to given value.

### HasAccessionVersion

`func (o *V2reportsTranscript) HasAccessionVersion() bool`

HasAccessionVersion returns a boolean if a field has been set.

### GetName

`func (o *V2reportsTranscript) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2reportsTranscript) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2reportsTranscript) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2reportsTranscript) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLength

`func (o *V2reportsTranscript) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *V2reportsTranscript) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *V2reportsTranscript) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *V2reportsTranscript) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetCds

`func (o *V2reportsTranscript) GetCds() V2reportsSeqRangeSet`

GetCds returns the Cds field if non-nil, zero value otherwise.

### GetCdsOk

`func (o *V2reportsTranscript) GetCdsOk() (*V2reportsSeqRangeSet, bool)`

GetCdsOk returns a tuple with the Cds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCds

`func (o *V2reportsTranscript) SetCds(v V2reportsSeqRangeSet)`

SetCds sets Cds field to given value.

### HasCds

`func (o *V2reportsTranscript) HasCds() bool`

HasCds returns a boolean if a field has been set.

### GetGenomicLocations

`func (o *V2reportsTranscript) GetGenomicLocations() []V2reportsGenomicLocation`

GetGenomicLocations returns the GenomicLocations field if non-nil, zero value otherwise.

### GetGenomicLocationsOk

`func (o *V2reportsTranscript) GetGenomicLocationsOk() (*[]V2reportsGenomicLocation, bool)`

GetGenomicLocationsOk returns a tuple with the GenomicLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicLocations

`func (o *V2reportsTranscript) SetGenomicLocations(v []V2reportsGenomicLocation)`

SetGenomicLocations sets GenomicLocations field to given value.

### HasGenomicLocations

`func (o *V2reportsTranscript) HasGenomicLocations() bool`

HasGenomicLocations returns a boolean if a field has been set.

### GetEnsemblTranscript

`func (o *V2reportsTranscript) GetEnsemblTranscript() string`

GetEnsemblTranscript returns the EnsemblTranscript field if non-nil, zero value otherwise.

### GetEnsemblTranscriptOk

`func (o *V2reportsTranscript) GetEnsemblTranscriptOk() (*string, bool)`

GetEnsemblTranscriptOk returns a tuple with the EnsemblTranscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsemblTranscript

`func (o *V2reportsTranscript) SetEnsemblTranscript(v string)`

SetEnsemblTranscript sets EnsemblTranscript field to given value.

### HasEnsemblTranscript

`func (o *V2reportsTranscript) HasEnsemblTranscript() bool`

HasEnsemblTranscript returns a boolean if a field has been set.

### GetProtein

`func (o *V2reportsTranscript) GetProtein() V2reportsProtein`

GetProtein returns the Protein field if non-nil, zero value otherwise.

### GetProteinOk

`func (o *V2reportsTranscript) GetProteinOk() (*V2reportsProtein, bool)`

GetProteinOk returns a tuple with the Protein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtein

`func (o *V2reportsTranscript) SetProtein(v V2reportsProtein)`

SetProtein sets Protein field to given value.

### HasProtein

`func (o *V2reportsTranscript) HasProtein() bool`

HasProtein returns a boolean if a field has been set.

### GetType

`func (o *V2reportsTranscript) GetType() V2reportsTranscriptTranscriptType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2reportsTranscript) GetTypeOk() (*V2reportsTranscriptTranscriptType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2reportsTranscript) SetType(v V2reportsTranscriptTranscriptType)`

SetType sets Type field to given value.

### HasType

`func (o *V2reportsTranscript) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


