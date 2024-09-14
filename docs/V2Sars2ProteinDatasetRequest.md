# V2Sars2ProteinDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Proteins** | Pointer to **[]string** |  | [optional] 
**RefseqOnly** | Pointer to **bool** |  | [optional] 
**AnnotatedOnly** | Pointer to **bool** |  | [optional] 
**ReleasedSince** | Pointer to **time.Time** |  | [optional] 
**UpdatedSince** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**GeoLocation** | Pointer to **string** |  | [optional] 
**CompleteOnly** | Pointer to **bool** |  | [optional] 
**TableFields** | Pointer to [**[]V2VirusTableField**](V2VirusTableField.md) |  | [optional] 
**IncludeSequence** | Pointer to [**[]V2ViralSequenceType**](V2ViralSequenceType.md) |  | [optional] 
**AuxReport** | Pointer to [**[]V2VirusDatasetReportType**](V2VirusDatasetReportType.md) |  | [optional] 
**Format** | Pointer to [**V2TableFormat**](V2TableFormat.md) |  | [optional] [default to V2TABLEFORMAT_TSV]

## Methods

### NewV2Sars2ProteinDatasetRequest

`func NewV2Sars2ProteinDatasetRequest() *V2Sars2ProteinDatasetRequest`

NewV2Sars2ProteinDatasetRequest instantiates a new V2Sars2ProteinDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2Sars2ProteinDatasetRequestWithDefaults

`func NewV2Sars2ProteinDatasetRequestWithDefaults() *V2Sars2ProteinDatasetRequest`

NewV2Sars2ProteinDatasetRequestWithDefaults instantiates a new V2Sars2ProteinDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProteins

`func (o *V2Sars2ProteinDatasetRequest) GetProteins() []string`

GetProteins returns the Proteins field if non-nil, zero value otherwise.

### GetProteinsOk

`func (o *V2Sars2ProteinDatasetRequest) GetProteinsOk() (*[]string, bool)`

GetProteinsOk returns a tuple with the Proteins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteins

`func (o *V2Sars2ProteinDatasetRequest) SetProteins(v []string)`

SetProteins sets Proteins field to given value.

### HasProteins

`func (o *V2Sars2ProteinDatasetRequest) HasProteins() bool`

HasProteins returns a boolean if a field has been set.

### GetRefseqOnly

`func (o *V2Sars2ProteinDatasetRequest) GetRefseqOnly() bool`

GetRefseqOnly returns the RefseqOnly field if non-nil, zero value otherwise.

### GetRefseqOnlyOk

`func (o *V2Sars2ProteinDatasetRequest) GetRefseqOnlyOk() (*bool, bool)`

GetRefseqOnlyOk returns a tuple with the RefseqOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqOnly

`func (o *V2Sars2ProteinDatasetRequest) SetRefseqOnly(v bool)`

SetRefseqOnly sets RefseqOnly field to given value.

### HasRefseqOnly

`func (o *V2Sars2ProteinDatasetRequest) HasRefseqOnly() bool`

HasRefseqOnly returns a boolean if a field has been set.

### GetAnnotatedOnly

`func (o *V2Sars2ProteinDatasetRequest) GetAnnotatedOnly() bool`

GetAnnotatedOnly returns the AnnotatedOnly field if non-nil, zero value otherwise.

### GetAnnotatedOnlyOk

`func (o *V2Sars2ProteinDatasetRequest) GetAnnotatedOnlyOk() (*bool, bool)`

GetAnnotatedOnlyOk returns a tuple with the AnnotatedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotatedOnly

`func (o *V2Sars2ProteinDatasetRequest) SetAnnotatedOnly(v bool)`

SetAnnotatedOnly sets AnnotatedOnly field to given value.

### HasAnnotatedOnly

`func (o *V2Sars2ProteinDatasetRequest) HasAnnotatedOnly() bool`

HasAnnotatedOnly returns a boolean if a field has been set.

### GetReleasedSince

`func (o *V2Sars2ProteinDatasetRequest) GetReleasedSince() time.Time`

GetReleasedSince returns the ReleasedSince field if non-nil, zero value otherwise.

### GetReleasedSinceOk

`func (o *V2Sars2ProteinDatasetRequest) GetReleasedSinceOk() (*time.Time, bool)`

GetReleasedSinceOk returns a tuple with the ReleasedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedSince

`func (o *V2Sars2ProteinDatasetRequest) SetReleasedSince(v time.Time)`

SetReleasedSince sets ReleasedSince field to given value.

### HasReleasedSince

`func (o *V2Sars2ProteinDatasetRequest) HasReleasedSince() bool`

HasReleasedSince returns a boolean if a field has been set.

### GetUpdatedSince

`func (o *V2Sars2ProteinDatasetRequest) GetUpdatedSince() time.Time`

GetUpdatedSince returns the UpdatedSince field if non-nil, zero value otherwise.

### GetUpdatedSinceOk

`func (o *V2Sars2ProteinDatasetRequest) GetUpdatedSinceOk() (*time.Time, bool)`

GetUpdatedSinceOk returns a tuple with the UpdatedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedSince

`func (o *V2Sars2ProteinDatasetRequest) SetUpdatedSince(v time.Time)`

SetUpdatedSince sets UpdatedSince field to given value.

### HasUpdatedSince

`func (o *V2Sars2ProteinDatasetRequest) HasUpdatedSince() bool`

HasUpdatedSince returns a boolean if a field has been set.

### GetHost

`func (o *V2Sars2ProteinDatasetRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V2Sars2ProteinDatasetRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V2Sars2ProteinDatasetRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V2Sars2ProteinDatasetRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetGeoLocation

`func (o *V2Sars2ProteinDatasetRequest) GetGeoLocation() string`

GetGeoLocation returns the GeoLocation field if non-nil, zero value otherwise.

### GetGeoLocationOk

`func (o *V2Sars2ProteinDatasetRequest) GetGeoLocationOk() (*string, bool)`

GetGeoLocationOk returns a tuple with the GeoLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoLocation

`func (o *V2Sars2ProteinDatasetRequest) SetGeoLocation(v string)`

SetGeoLocation sets GeoLocation field to given value.

### HasGeoLocation

`func (o *V2Sars2ProteinDatasetRequest) HasGeoLocation() bool`

HasGeoLocation returns a boolean if a field has been set.

### GetCompleteOnly

`func (o *V2Sars2ProteinDatasetRequest) GetCompleteOnly() bool`

GetCompleteOnly returns the CompleteOnly field if non-nil, zero value otherwise.

### GetCompleteOnlyOk

`func (o *V2Sars2ProteinDatasetRequest) GetCompleteOnlyOk() (*bool, bool)`

GetCompleteOnlyOk returns a tuple with the CompleteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteOnly

`func (o *V2Sars2ProteinDatasetRequest) SetCompleteOnly(v bool)`

SetCompleteOnly sets CompleteOnly field to given value.

### HasCompleteOnly

`func (o *V2Sars2ProteinDatasetRequest) HasCompleteOnly() bool`

HasCompleteOnly returns a boolean if a field has been set.

### GetTableFields

`func (o *V2Sars2ProteinDatasetRequest) GetTableFields() []V2VirusTableField`

GetTableFields returns the TableFields field if non-nil, zero value otherwise.

### GetTableFieldsOk

`func (o *V2Sars2ProteinDatasetRequest) GetTableFieldsOk() (*[]V2VirusTableField, bool)`

GetTableFieldsOk returns a tuple with the TableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFields

`func (o *V2Sars2ProteinDatasetRequest) SetTableFields(v []V2VirusTableField)`

SetTableFields sets TableFields field to given value.

### HasTableFields

`func (o *V2Sars2ProteinDatasetRequest) HasTableFields() bool`

HasTableFields returns a boolean if a field has been set.

### GetIncludeSequence

`func (o *V2Sars2ProteinDatasetRequest) GetIncludeSequence() []V2ViralSequenceType`

GetIncludeSequence returns the IncludeSequence field if non-nil, zero value otherwise.

### GetIncludeSequenceOk

`func (o *V2Sars2ProteinDatasetRequest) GetIncludeSequenceOk() (*[]V2ViralSequenceType, bool)`

GetIncludeSequenceOk returns a tuple with the IncludeSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSequence

`func (o *V2Sars2ProteinDatasetRequest) SetIncludeSequence(v []V2ViralSequenceType)`

SetIncludeSequence sets IncludeSequence field to given value.

### HasIncludeSequence

`func (o *V2Sars2ProteinDatasetRequest) HasIncludeSequence() bool`

HasIncludeSequence returns a boolean if a field has been set.

### GetAuxReport

`func (o *V2Sars2ProteinDatasetRequest) GetAuxReport() []V2VirusDatasetReportType`

GetAuxReport returns the AuxReport field if non-nil, zero value otherwise.

### GetAuxReportOk

`func (o *V2Sars2ProteinDatasetRequest) GetAuxReportOk() (*[]V2VirusDatasetReportType, bool)`

GetAuxReportOk returns a tuple with the AuxReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxReport

`func (o *V2Sars2ProteinDatasetRequest) SetAuxReport(v []V2VirusDatasetReportType)`

SetAuxReport sets AuxReport field to given value.

### HasAuxReport

`func (o *V2Sars2ProteinDatasetRequest) HasAuxReport() bool`

HasAuxReport returns a boolean if a field has been set.

### GetFormat

`func (o *V2Sars2ProteinDatasetRequest) GetFormat() V2TableFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V2Sars2ProteinDatasetRequest) GetFormatOk() (*V2TableFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V2Sars2ProteinDatasetRequest) SetFormat(v V2TableFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *V2Sars2ProteinDatasetRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


