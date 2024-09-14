# V2reportsGenomeAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TaxId** | Pointer to **string** |  | [optional] 
**Taxname** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**V2reportsGeneType**](V2reportsGeneType.md) |  | [optional] [default to V2REPORTSGENETYPE_UNKNOWN]
**GeneType** | Pointer to **string** |  | [optional] 
**RnaType** | Pointer to [**V2reportsRnaType**](V2reportsRnaType.md) |  | [optional] [default to V2REPORTSRNATYPE_RNA_UNKNOWN]
**Orientation** | Pointer to [**V2reportsOrientation**](V2reportsOrientation.md) |  | [optional] [default to V2REPORTSORIENTATION_NONE]
**LocusTag** | Pointer to **string** |  | [optional] 
**ReferenceStandards** | Pointer to [**[]V2reportsGenomicRegion**](V2reportsGenomicRegion.md) |  | [optional] 
**GenomicRegions** | Pointer to [**[]V2reportsGenomicRegion**](V2reportsGenomicRegion.md) |  | [optional] 
**Transcripts** | Pointer to [**[]V2reportsTranscript**](V2reportsTranscript.md) |  | [optional] 
**Proteins** | Pointer to [**[]V2reportsProtein**](V2reportsProtein.md) |  | [optional] 
**Chromosomes** | Pointer to **[]string** |  | [optional] 
**SwissProtAccessions** | Pointer to **[]string** |  | [optional] 
**EnsemblGeneIds** | Pointer to **[]string** |  | [optional] 
**OmimIds** | Pointer to **[]string** |  | [optional] 
**Synonyms** | Pointer to **[]string** |  | [optional] 
**Annotations** | Pointer to [**[]V2reportsAnnotation**](V2reportsAnnotation.md) |  | [optional] 

## Methods

### NewV2reportsGenomeAnnotation

`func NewV2reportsGenomeAnnotation() *V2reportsGenomeAnnotation`

NewV2reportsGenomeAnnotation instantiates a new V2reportsGenomeAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsGenomeAnnotationWithDefaults

`func NewV2reportsGenomeAnnotationWithDefaults() *V2reportsGenomeAnnotation`

NewV2reportsGenomeAnnotationWithDefaults instantiates a new V2reportsGenomeAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneId

`func (o *V2reportsGenomeAnnotation) GetGeneId() string`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *V2reportsGenomeAnnotation) GetGeneIdOk() (*string, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *V2reportsGenomeAnnotation) SetGeneId(v string)`

SetGeneId sets GeneId field to given value.

### HasGeneId

`func (o *V2reportsGenomeAnnotation) HasGeneId() bool`

HasGeneId returns a boolean if a field has been set.

### GetSymbol

`func (o *V2reportsGenomeAnnotation) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *V2reportsGenomeAnnotation) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *V2reportsGenomeAnnotation) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *V2reportsGenomeAnnotation) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *V2reportsGenomeAnnotation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2reportsGenomeAnnotation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2reportsGenomeAnnotation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V2reportsGenomeAnnotation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *V2reportsGenomeAnnotation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2reportsGenomeAnnotation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2reportsGenomeAnnotation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2reportsGenomeAnnotation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxId

`func (o *V2reportsGenomeAnnotation) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V2reportsGenomeAnnotation) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V2reportsGenomeAnnotation) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V2reportsGenomeAnnotation) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxname

`func (o *V2reportsGenomeAnnotation) GetTaxname() string`

GetTaxname returns the Taxname field if non-nil, zero value otherwise.

### GetTaxnameOk

`func (o *V2reportsGenomeAnnotation) GetTaxnameOk() (*string, bool)`

GetTaxnameOk returns a tuple with the Taxname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxname

`func (o *V2reportsGenomeAnnotation) SetTaxname(v string)`

SetTaxname sets Taxname field to given value.

### HasTaxname

`func (o *V2reportsGenomeAnnotation) HasTaxname() bool`

HasTaxname returns a boolean if a field has been set.

### GetCommonName

`func (o *V2reportsGenomeAnnotation) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V2reportsGenomeAnnotation) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V2reportsGenomeAnnotation) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V2reportsGenomeAnnotation) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetType

`func (o *V2reportsGenomeAnnotation) GetType() V2reportsGeneType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2reportsGenomeAnnotation) GetTypeOk() (*V2reportsGeneType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2reportsGenomeAnnotation) SetType(v V2reportsGeneType)`

SetType sets Type field to given value.

### HasType

`func (o *V2reportsGenomeAnnotation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGeneType

`func (o *V2reportsGenomeAnnotation) GetGeneType() string`

GetGeneType returns the GeneType field if non-nil, zero value otherwise.

### GetGeneTypeOk

`func (o *V2reportsGenomeAnnotation) GetGeneTypeOk() (*string, bool)`

GetGeneTypeOk returns a tuple with the GeneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneType

`func (o *V2reportsGenomeAnnotation) SetGeneType(v string)`

SetGeneType sets GeneType field to given value.

### HasGeneType

`func (o *V2reportsGenomeAnnotation) HasGeneType() bool`

HasGeneType returns a boolean if a field has been set.

### GetRnaType

`func (o *V2reportsGenomeAnnotation) GetRnaType() V2reportsRnaType`

GetRnaType returns the RnaType field if non-nil, zero value otherwise.

### GetRnaTypeOk

`func (o *V2reportsGenomeAnnotation) GetRnaTypeOk() (*V2reportsRnaType, bool)`

GetRnaTypeOk returns a tuple with the RnaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRnaType

`func (o *V2reportsGenomeAnnotation) SetRnaType(v V2reportsRnaType)`

SetRnaType sets RnaType field to given value.

### HasRnaType

`func (o *V2reportsGenomeAnnotation) HasRnaType() bool`

HasRnaType returns a boolean if a field has been set.

### GetOrientation

`func (o *V2reportsGenomeAnnotation) GetOrientation() V2reportsOrientation`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *V2reportsGenomeAnnotation) GetOrientationOk() (*V2reportsOrientation, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *V2reportsGenomeAnnotation) SetOrientation(v V2reportsOrientation)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *V2reportsGenomeAnnotation) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetLocusTag

`func (o *V2reportsGenomeAnnotation) GetLocusTag() string`

GetLocusTag returns the LocusTag field if non-nil, zero value otherwise.

### GetLocusTagOk

`func (o *V2reportsGenomeAnnotation) GetLocusTagOk() (*string, bool)`

GetLocusTagOk returns a tuple with the LocusTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocusTag

`func (o *V2reportsGenomeAnnotation) SetLocusTag(v string)`

SetLocusTag sets LocusTag field to given value.

### HasLocusTag

`func (o *V2reportsGenomeAnnotation) HasLocusTag() bool`

HasLocusTag returns a boolean if a field has been set.

### GetReferenceStandards

`func (o *V2reportsGenomeAnnotation) GetReferenceStandards() []V2reportsGenomicRegion`

GetReferenceStandards returns the ReferenceStandards field if non-nil, zero value otherwise.

### GetReferenceStandardsOk

`func (o *V2reportsGenomeAnnotation) GetReferenceStandardsOk() (*[]V2reportsGenomicRegion, bool)`

GetReferenceStandardsOk returns a tuple with the ReferenceStandards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceStandards

`func (o *V2reportsGenomeAnnotation) SetReferenceStandards(v []V2reportsGenomicRegion)`

SetReferenceStandards sets ReferenceStandards field to given value.

### HasReferenceStandards

`func (o *V2reportsGenomeAnnotation) HasReferenceStandards() bool`

HasReferenceStandards returns a boolean if a field has been set.

### GetGenomicRegions

`func (o *V2reportsGenomeAnnotation) GetGenomicRegions() []V2reportsGenomicRegion`

GetGenomicRegions returns the GenomicRegions field if non-nil, zero value otherwise.

### GetGenomicRegionsOk

`func (o *V2reportsGenomeAnnotation) GetGenomicRegionsOk() (*[]V2reportsGenomicRegion, bool)`

GetGenomicRegionsOk returns a tuple with the GenomicRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicRegions

`func (o *V2reportsGenomeAnnotation) SetGenomicRegions(v []V2reportsGenomicRegion)`

SetGenomicRegions sets GenomicRegions field to given value.

### HasGenomicRegions

`func (o *V2reportsGenomeAnnotation) HasGenomicRegions() bool`

HasGenomicRegions returns a boolean if a field has been set.

### GetTranscripts

`func (o *V2reportsGenomeAnnotation) GetTranscripts() []V2reportsTranscript`

GetTranscripts returns the Transcripts field if non-nil, zero value otherwise.

### GetTranscriptsOk

`func (o *V2reportsGenomeAnnotation) GetTranscriptsOk() (*[]V2reportsTranscript, bool)`

GetTranscriptsOk returns a tuple with the Transcripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscripts

`func (o *V2reportsGenomeAnnotation) SetTranscripts(v []V2reportsTranscript)`

SetTranscripts sets Transcripts field to given value.

### HasTranscripts

`func (o *V2reportsGenomeAnnotation) HasTranscripts() bool`

HasTranscripts returns a boolean if a field has been set.

### GetProteins

`func (o *V2reportsGenomeAnnotation) GetProteins() []V2reportsProtein`

GetProteins returns the Proteins field if non-nil, zero value otherwise.

### GetProteinsOk

`func (o *V2reportsGenomeAnnotation) GetProteinsOk() (*[]V2reportsProtein, bool)`

GetProteinsOk returns a tuple with the Proteins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteins

`func (o *V2reportsGenomeAnnotation) SetProteins(v []V2reportsProtein)`

SetProteins sets Proteins field to given value.

### HasProteins

`func (o *V2reportsGenomeAnnotation) HasProteins() bool`

HasProteins returns a boolean if a field has been set.

### GetChromosomes

`func (o *V2reportsGenomeAnnotation) GetChromosomes() []string`

GetChromosomes returns the Chromosomes field if non-nil, zero value otherwise.

### GetChromosomesOk

`func (o *V2reportsGenomeAnnotation) GetChromosomesOk() (*[]string, bool)`

GetChromosomesOk returns a tuple with the Chromosomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosomes

`func (o *V2reportsGenomeAnnotation) SetChromosomes(v []string)`

SetChromosomes sets Chromosomes field to given value.

### HasChromosomes

`func (o *V2reportsGenomeAnnotation) HasChromosomes() bool`

HasChromosomes returns a boolean if a field has been set.

### GetSwissProtAccessions

`func (o *V2reportsGenomeAnnotation) GetSwissProtAccessions() []string`

GetSwissProtAccessions returns the SwissProtAccessions field if non-nil, zero value otherwise.

### GetSwissProtAccessionsOk

`func (o *V2reportsGenomeAnnotation) GetSwissProtAccessionsOk() (*[]string, bool)`

GetSwissProtAccessionsOk returns a tuple with the SwissProtAccessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwissProtAccessions

`func (o *V2reportsGenomeAnnotation) SetSwissProtAccessions(v []string)`

SetSwissProtAccessions sets SwissProtAccessions field to given value.

### HasSwissProtAccessions

`func (o *V2reportsGenomeAnnotation) HasSwissProtAccessions() bool`

HasSwissProtAccessions returns a boolean if a field has been set.

### GetEnsemblGeneIds

`func (o *V2reportsGenomeAnnotation) GetEnsemblGeneIds() []string`

GetEnsemblGeneIds returns the EnsemblGeneIds field if non-nil, zero value otherwise.

### GetEnsemblGeneIdsOk

`func (o *V2reportsGenomeAnnotation) GetEnsemblGeneIdsOk() (*[]string, bool)`

GetEnsemblGeneIdsOk returns a tuple with the EnsemblGeneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsemblGeneIds

`func (o *V2reportsGenomeAnnotation) SetEnsemblGeneIds(v []string)`

SetEnsemblGeneIds sets EnsemblGeneIds field to given value.

### HasEnsemblGeneIds

`func (o *V2reportsGenomeAnnotation) HasEnsemblGeneIds() bool`

HasEnsemblGeneIds returns a boolean if a field has been set.

### GetOmimIds

`func (o *V2reportsGenomeAnnotation) GetOmimIds() []string`

GetOmimIds returns the OmimIds field if non-nil, zero value otherwise.

### GetOmimIdsOk

`func (o *V2reportsGenomeAnnotation) GetOmimIdsOk() (*[]string, bool)`

GetOmimIdsOk returns a tuple with the OmimIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmimIds

`func (o *V2reportsGenomeAnnotation) SetOmimIds(v []string)`

SetOmimIds sets OmimIds field to given value.

### HasOmimIds

`func (o *V2reportsGenomeAnnotation) HasOmimIds() bool`

HasOmimIds returns a boolean if a field has been set.

### GetSynonyms

`func (o *V2reportsGenomeAnnotation) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *V2reportsGenomeAnnotation) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *V2reportsGenomeAnnotation) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *V2reportsGenomeAnnotation) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetAnnotations

`func (o *V2reportsGenomeAnnotation) GetAnnotations() []V2reportsAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *V2reportsGenomeAnnotation) GetAnnotationsOk() (*[]V2reportsAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *V2reportsGenomeAnnotation) SetAnnotations(v []V2reportsAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *V2reportsGenomeAnnotation) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


