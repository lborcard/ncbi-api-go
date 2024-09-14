# V2reportsVirusAnnotationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accession** | Pointer to **string** |  | [optional] 
**IsolateName** | Pointer to **string** |  | [optional] 
**Genes** | Pointer to [**[]V2reportsVirusGene**](V2reportsVirusGene.md) |  | [optional] 

## Methods

### NewV2reportsVirusAnnotationReport

`func NewV2reportsVirusAnnotationReport() *V2reportsVirusAnnotationReport`

NewV2reportsVirusAnnotationReport instantiates a new V2reportsVirusAnnotationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsVirusAnnotationReportWithDefaults

`func NewV2reportsVirusAnnotationReportWithDefaults() *V2reportsVirusAnnotationReport`

NewV2reportsVirusAnnotationReportWithDefaults instantiates a new V2reportsVirusAnnotationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccession

`func (o *V2reportsVirusAnnotationReport) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V2reportsVirusAnnotationReport) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V2reportsVirusAnnotationReport) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V2reportsVirusAnnotationReport) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetIsolateName

`func (o *V2reportsVirusAnnotationReport) GetIsolateName() string`

GetIsolateName returns the IsolateName field if non-nil, zero value otherwise.

### GetIsolateNameOk

`func (o *V2reportsVirusAnnotationReport) GetIsolateNameOk() (*string, bool)`

GetIsolateNameOk returns a tuple with the IsolateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolateName

`func (o *V2reportsVirusAnnotationReport) SetIsolateName(v string)`

SetIsolateName sets IsolateName field to given value.

### HasIsolateName

`func (o *V2reportsVirusAnnotationReport) HasIsolateName() bool`

HasIsolateName returns a boolean if a field has been set.

### GetGenes

`func (o *V2reportsVirusAnnotationReport) GetGenes() []V2reportsVirusGene`

GetGenes returns the Genes field if non-nil, zero value otherwise.

### GetGenesOk

`func (o *V2reportsVirusAnnotationReport) GetGenesOk() (*[]V2reportsVirusGene, bool)`

GetGenesOk returns a tuple with the Genes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenes

`func (o *V2reportsVirusAnnotationReport) SetGenes(v []V2reportsVirusGene)`

SetGenes sets Genes field to given value.

### HasGenes

`func (o *V2reportsVirusAnnotationReport) HasGenes() bool`

HasGenes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


