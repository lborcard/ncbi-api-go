# V2VirusDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessions** | Pointer to **[]string** |  | [optional] 
**Taxon** | Pointer to **string** |  | [optional] 
**RefseqOnly** | Pointer to **bool** |  | [optional] 
**AnnotatedOnly** | Pointer to **bool** |  | [optional] 
**ReleasedSince** | Pointer to **time.Time** |  | [optional] 
**UpdatedSince** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**PangolinClassification** | Pointer to **string** |  | [optional] 
**GeoLocation** | Pointer to **string** |  | [optional] 
**CompleteOnly** | Pointer to **bool** |  | [optional] 
**TableFields** | Pointer to [**[]V2VirusTableField**](V2VirusTableField.md) |  | [optional] 
**IncludeSequence** | Pointer to [**[]V2ViralSequenceType**](V2ViralSequenceType.md) |  | [optional] 
**AuxReport** | Pointer to [**[]V2VirusDatasetReportType**](V2VirusDatasetReportType.md) |  | [optional] 
**Format** | Pointer to [**V2TableFormat**](V2TableFormat.md) |  | [optional] [default to V2TABLEFORMAT_TSV]
**UsePsg** | Pointer to **bool** |  | [optional] 

## Methods

### NewV2VirusDatasetRequest

`func NewV2VirusDatasetRequest() *V2VirusDatasetRequest`

NewV2VirusDatasetRequest instantiates a new V2VirusDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2VirusDatasetRequestWithDefaults

`func NewV2VirusDatasetRequestWithDefaults() *V2VirusDatasetRequest`

NewV2VirusDatasetRequestWithDefaults instantiates a new V2VirusDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessions

`func (o *V2VirusDatasetRequest) GetAccessions() []string`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V2VirusDatasetRequest) GetAccessionsOk() (*[]string, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V2VirusDatasetRequest) SetAccessions(v []string)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V2VirusDatasetRequest) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetTaxon

`func (o *V2VirusDatasetRequest) GetTaxon() string`

GetTaxon returns the Taxon field if non-nil, zero value otherwise.

### GetTaxonOk

`func (o *V2VirusDatasetRequest) GetTaxonOk() (*string, bool)`

GetTaxonOk returns a tuple with the Taxon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxon

`func (o *V2VirusDatasetRequest) SetTaxon(v string)`

SetTaxon sets Taxon field to given value.

### HasTaxon

`func (o *V2VirusDatasetRequest) HasTaxon() bool`

HasTaxon returns a boolean if a field has been set.

### GetRefseqOnly

`func (o *V2VirusDatasetRequest) GetRefseqOnly() bool`

GetRefseqOnly returns the RefseqOnly field if non-nil, zero value otherwise.

### GetRefseqOnlyOk

`func (o *V2VirusDatasetRequest) GetRefseqOnlyOk() (*bool, bool)`

GetRefseqOnlyOk returns a tuple with the RefseqOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqOnly

`func (o *V2VirusDatasetRequest) SetRefseqOnly(v bool)`

SetRefseqOnly sets RefseqOnly field to given value.

### HasRefseqOnly

`func (o *V2VirusDatasetRequest) HasRefseqOnly() bool`

HasRefseqOnly returns a boolean if a field has been set.

### GetAnnotatedOnly

`func (o *V2VirusDatasetRequest) GetAnnotatedOnly() bool`

GetAnnotatedOnly returns the AnnotatedOnly field if non-nil, zero value otherwise.

### GetAnnotatedOnlyOk

`func (o *V2VirusDatasetRequest) GetAnnotatedOnlyOk() (*bool, bool)`

GetAnnotatedOnlyOk returns a tuple with the AnnotatedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotatedOnly

`func (o *V2VirusDatasetRequest) SetAnnotatedOnly(v bool)`

SetAnnotatedOnly sets AnnotatedOnly field to given value.

### HasAnnotatedOnly

`func (o *V2VirusDatasetRequest) HasAnnotatedOnly() bool`

HasAnnotatedOnly returns a boolean if a field has been set.

### GetReleasedSince

`func (o *V2VirusDatasetRequest) GetReleasedSince() time.Time`

GetReleasedSince returns the ReleasedSince field if non-nil, zero value otherwise.

### GetReleasedSinceOk

`func (o *V2VirusDatasetRequest) GetReleasedSinceOk() (*time.Time, bool)`

GetReleasedSinceOk returns a tuple with the ReleasedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedSince

`func (o *V2VirusDatasetRequest) SetReleasedSince(v time.Time)`

SetReleasedSince sets ReleasedSince field to given value.

### HasReleasedSince

`func (o *V2VirusDatasetRequest) HasReleasedSince() bool`

HasReleasedSince returns a boolean if a field has been set.

### GetUpdatedSince

`func (o *V2VirusDatasetRequest) GetUpdatedSince() time.Time`

GetUpdatedSince returns the UpdatedSince field if non-nil, zero value otherwise.

### GetUpdatedSinceOk

`func (o *V2VirusDatasetRequest) GetUpdatedSinceOk() (*time.Time, bool)`

GetUpdatedSinceOk returns a tuple with the UpdatedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedSince

`func (o *V2VirusDatasetRequest) SetUpdatedSince(v time.Time)`

SetUpdatedSince sets UpdatedSince field to given value.

### HasUpdatedSince

`func (o *V2VirusDatasetRequest) HasUpdatedSince() bool`

HasUpdatedSince returns a boolean if a field has been set.

### GetHost

`func (o *V2VirusDatasetRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V2VirusDatasetRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V2VirusDatasetRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V2VirusDatasetRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPangolinClassification

`func (o *V2VirusDatasetRequest) GetPangolinClassification() string`

GetPangolinClassification returns the PangolinClassification field if non-nil, zero value otherwise.

### GetPangolinClassificationOk

`func (o *V2VirusDatasetRequest) GetPangolinClassificationOk() (*string, bool)`

GetPangolinClassificationOk returns a tuple with the PangolinClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPangolinClassification

`func (o *V2VirusDatasetRequest) SetPangolinClassification(v string)`

SetPangolinClassification sets PangolinClassification field to given value.

### HasPangolinClassification

`func (o *V2VirusDatasetRequest) HasPangolinClassification() bool`

HasPangolinClassification returns a boolean if a field has been set.

### GetGeoLocation

`func (o *V2VirusDatasetRequest) GetGeoLocation() string`

GetGeoLocation returns the GeoLocation field if non-nil, zero value otherwise.

### GetGeoLocationOk

`func (o *V2VirusDatasetRequest) GetGeoLocationOk() (*string, bool)`

GetGeoLocationOk returns a tuple with the GeoLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoLocation

`func (o *V2VirusDatasetRequest) SetGeoLocation(v string)`

SetGeoLocation sets GeoLocation field to given value.

### HasGeoLocation

`func (o *V2VirusDatasetRequest) HasGeoLocation() bool`

HasGeoLocation returns a boolean if a field has been set.

### GetCompleteOnly

`func (o *V2VirusDatasetRequest) GetCompleteOnly() bool`

GetCompleteOnly returns the CompleteOnly field if non-nil, zero value otherwise.

### GetCompleteOnlyOk

`func (o *V2VirusDatasetRequest) GetCompleteOnlyOk() (*bool, bool)`

GetCompleteOnlyOk returns a tuple with the CompleteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteOnly

`func (o *V2VirusDatasetRequest) SetCompleteOnly(v bool)`

SetCompleteOnly sets CompleteOnly field to given value.

### HasCompleteOnly

`func (o *V2VirusDatasetRequest) HasCompleteOnly() bool`

HasCompleteOnly returns a boolean if a field has been set.

### GetTableFields

`func (o *V2VirusDatasetRequest) GetTableFields() []V2VirusTableField`

GetTableFields returns the TableFields field if non-nil, zero value otherwise.

### GetTableFieldsOk

`func (o *V2VirusDatasetRequest) GetTableFieldsOk() (*[]V2VirusTableField, bool)`

GetTableFieldsOk returns a tuple with the TableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFields

`func (o *V2VirusDatasetRequest) SetTableFields(v []V2VirusTableField)`

SetTableFields sets TableFields field to given value.

### HasTableFields

`func (o *V2VirusDatasetRequest) HasTableFields() bool`

HasTableFields returns a boolean if a field has been set.

### GetIncludeSequence

`func (o *V2VirusDatasetRequest) GetIncludeSequence() []V2ViralSequenceType`

GetIncludeSequence returns the IncludeSequence field if non-nil, zero value otherwise.

### GetIncludeSequenceOk

`func (o *V2VirusDatasetRequest) GetIncludeSequenceOk() (*[]V2ViralSequenceType, bool)`

GetIncludeSequenceOk returns a tuple with the IncludeSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSequence

`func (o *V2VirusDatasetRequest) SetIncludeSequence(v []V2ViralSequenceType)`

SetIncludeSequence sets IncludeSequence field to given value.

### HasIncludeSequence

`func (o *V2VirusDatasetRequest) HasIncludeSequence() bool`

HasIncludeSequence returns a boolean if a field has been set.

### GetAuxReport

`func (o *V2VirusDatasetRequest) GetAuxReport() []V2VirusDatasetReportType`

GetAuxReport returns the AuxReport field if non-nil, zero value otherwise.

### GetAuxReportOk

`func (o *V2VirusDatasetRequest) GetAuxReportOk() (*[]V2VirusDatasetReportType, bool)`

GetAuxReportOk returns a tuple with the AuxReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxReport

`func (o *V2VirusDatasetRequest) SetAuxReport(v []V2VirusDatasetReportType)`

SetAuxReport sets AuxReport field to given value.

### HasAuxReport

`func (o *V2VirusDatasetRequest) HasAuxReport() bool`

HasAuxReport returns a boolean if a field has been set.

### GetFormat

`func (o *V2VirusDatasetRequest) GetFormat() V2TableFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V2VirusDatasetRequest) GetFormatOk() (*V2TableFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V2VirusDatasetRequest) SetFormat(v V2TableFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *V2VirusDatasetRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetUsePsg

`func (o *V2VirusDatasetRequest) GetUsePsg() bool`

GetUsePsg returns the UsePsg field if non-nil, zero value otherwise.

### GetUsePsgOk

`func (o *V2VirusDatasetRequest) GetUsePsgOk() (*bool, bool)`

GetUsePsgOk returns a tuple with the UsePsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePsg

`func (o *V2VirusDatasetRequest) SetUsePsg(v bool)`

SetUsePsg sets UsePsg field to given value.

### HasUsePsg

`func (o *V2VirusDatasetRequest) HasUsePsg() bool`

HasUsePsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


