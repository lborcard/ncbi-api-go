# V2reportsGeneDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TaxId** | Pointer to **string** |  | [optional] 
**Taxname** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**V2reportsGeneType**](V2reportsGeneType.md) |  | [optional] [default to V2REPORTSGENETYPE_UNKNOWN]
**RnaType** | Pointer to [**V2reportsRnaType**](V2reportsRnaType.md) |  | [optional] [default to V2REPORTSRNATYPE_RNA_UNKNOWN]
**Orientation** | Pointer to [**V2reportsOrientation**](V2reportsOrientation.md) |  | [optional] [default to V2REPORTSORIENTATION_NONE]
**ReferenceStandards** | Pointer to [**[]V2reportsGenomicRegion**](V2reportsGenomicRegion.md) |  | [optional] 
**GenomicRegions** | Pointer to [**[]V2reportsGenomicRegion**](V2reportsGenomicRegion.md) |  | [optional] 
**Chromosomes** | Pointer to **[]string** |  | [optional] 
**NomenclatureAuthority** | Pointer to [**V2reportsNomenclatureAuthority**](V2reportsNomenclatureAuthority.md) |  | [optional] 
**SwissProtAccessions** | Pointer to **[]string** |  | [optional] 
**EnsemblGeneIds** | Pointer to **[]string** |  | [optional] 
**OmimIds** | Pointer to **[]string** |  | [optional] 
**Synonyms** | Pointer to **[]string** |  | [optional] 
**ReplacedGeneId** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to [**[]V2reportsAnnotation**](V2reportsAnnotation.md) |  | [optional] 
**TranscriptCount** | Pointer to **int32** |  | [optional] 
**ProteinCount** | Pointer to **int32** |  | [optional] 
**TranscriptTypeCounts** | Pointer to [**[]V2reportsTranscriptTypeCount**](V2reportsTranscriptTypeCount.md) |  | [optional] 
**GeneGroups** | Pointer to [**[]V2reportsGeneGroup**](V2reportsGeneGroup.md) |  | [optional] 

## Methods

### NewV2reportsGeneDescriptor

`func NewV2reportsGeneDescriptor() *V2reportsGeneDescriptor`

NewV2reportsGeneDescriptor instantiates a new V2reportsGeneDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsGeneDescriptorWithDefaults

`func NewV2reportsGeneDescriptorWithDefaults() *V2reportsGeneDescriptor`

NewV2reportsGeneDescriptorWithDefaults instantiates a new V2reportsGeneDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneId

`func (o *V2reportsGeneDescriptor) GetGeneId() string`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *V2reportsGeneDescriptor) GetGeneIdOk() (*string, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *V2reportsGeneDescriptor) SetGeneId(v string)`

SetGeneId sets GeneId field to given value.

### HasGeneId

`func (o *V2reportsGeneDescriptor) HasGeneId() bool`

HasGeneId returns a boolean if a field has been set.

### GetSymbol

`func (o *V2reportsGeneDescriptor) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *V2reportsGeneDescriptor) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *V2reportsGeneDescriptor) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *V2reportsGeneDescriptor) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *V2reportsGeneDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2reportsGeneDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2reportsGeneDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V2reportsGeneDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTaxId

`func (o *V2reportsGeneDescriptor) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V2reportsGeneDescriptor) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V2reportsGeneDescriptor) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V2reportsGeneDescriptor) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxname

`func (o *V2reportsGeneDescriptor) GetTaxname() string`

GetTaxname returns the Taxname field if non-nil, zero value otherwise.

### GetTaxnameOk

`func (o *V2reportsGeneDescriptor) GetTaxnameOk() (*string, bool)`

GetTaxnameOk returns a tuple with the Taxname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxname

`func (o *V2reportsGeneDescriptor) SetTaxname(v string)`

SetTaxname sets Taxname field to given value.

### HasTaxname

`func (o *V2reportsGeneDescriptor) HasTaxname() bool`

HasTaxname returns a boolean if a field has been set.

### GetCommonName

`func (o *V2reportsGeneDescriptor) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V2reportsGeneDescriptor) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V2reportsGeneDescriptor) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V2reportsGeneDescriptor) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetType

`func (o *V2reportsGeneDescriptor) GetType() V2reportsGeneType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2reportsGeneDescriptor) GetTypeOk() (*V2reportsGeneType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2reportsGeneDescriptor) SetType(v V2reportsGeneType)`

SetType sets Type field to given value.

### HasType

`func (o *V2reportsGeneDescriptor) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRnaType

`func (o *V2reportsGeneDescriptor) GetRnaType() V2reportsRnaType`

GetRnaType returns the RnaType field if non-nil, zero value otherwise.

### GetRnaTypeOk

`func (o *V2reportsGeneDescriptor) GetRnaTypeOk() (*V2reportsRnaType, bool)`

GetRnaTypeOk returns a tuple with the RnaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRnaType

`func (o *V2reportsGeneDescriptor) SetRnaType(v V2reportsRnaType)`

SetRnaType sets RnaType field to given value.

### HasRnaType

`func (o *V2reportsGeneDescriptor) HasRnaType() bool`

HasRnaType returns a boolean if a field has been set.

### GetOrientation

`func (o *V2reportsGeneDescriptor) GetOrientation() V2reportsOrientation`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *V2reportsGeneDescriptor) GetOrientationOk() (*V2reportsOrientation, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *V2reportsGeneDescriptor) SetOrientation(v V2reportsOrientation)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *V2reportsGeneDescriptor) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetReferenceStandards

`func (o *V2reportsGeneDescriptor) GetReferenceStandards() []V2reportsGenomicRegion`

GetReferenceStandards returns the ReferenceStandards field if non-nil, zero value otherwise.

### GetReferenceStandardsOk

`func (o *V2reportsGeneDescriptor) GetReferenceStandardsOk() (*[]V2reportsGenomicRegion, bool)`

GetReferenceStandardsOk returns a tuple with the ReferenceStandards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceStandards

`func (o *V2reportsGeneDescriptor) SetReferenceStandards(v []V2reportsGenomicRegion)`

SetReferenceStandards sets ReferenceStandards field to given value.

### HasReferenceStandards

`func (o *V2reportsGeneDescriptor) HasReferenceStandards() bool`

HasReferenceStandards returns a boolean if a field has been set.

### GetGenomicRegions

`func (o *V2reportsGeneDescriptor) GetGenomicRegions() []V2reportsGenomicRegion`

GetGenomicRegions returns the GenomicRegions field if non-nil, zero value otherwise.

### GetGenomicRegionsOk

`func (o *V2reportsGeneDescriptor) GetGenomicRegionsOk() (*[]V2reportsGenomicRegion, bool)`

GetGenomicRegionsOk returns a tuple with the GenomicRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicRegions

`func (o *V2reportsGeneDescriptor) SetGenomicRegions(v []V2reportsGenomicRegion)`

SetGenomicRegions sets GenomicRegions field to given value.

### HasGenomicRegions

`func (o *V2reportsGeneDescriptor) HasGenomicRegions() bool`

HasGenomicRegions returns a boolean if a field has been set.

### GetChromosomes

`func (o *V2reportsGeneDescriptor) GetChromosomes() []string`

GetChromosomes returns the Chromosomes field if non-nil, zero value otherwise.

### GetChromosomesOk

`func (o *V2reportsGeneDescriptor) GetChromosomesOk() (*[]string, bool)`

GetChromosomesOk returns a tuple with the Chromosomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosomes

`func (o *V2reportsGeneDescriptor) SetChromosomes(v []string)`

SetChromosomes sets Chromosomes field to given value.

### HasChromosomes

`func (o *V2reportsGeneDescriptor) HasChromosomes() bool`

HasChromosomes returns a boolean if a field has been set.

### GetNomenclatureAuthority

`func (o *V2reportsGeneDescriptor) GetNomenclatureAuthority() V2reportsNomenclatureAuthority`

GetNomenclatureAuthority returns the NomenclatureAuthority field if non-nil, zero value otherwise.

### GetNomenclatureAuthorityOk

`func (o *V2reportsGeneDescriptor) GetNomenclatureAuthorityOk() (*V2reportsNomenclatureAuthority, bool)`

GetNomenclatureAuthorityOk returns a tuple with the NomenclatureAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomenclatureAuthority

`func (o *V2reportsGeneDescriptor) SetNomenclatureAuthority(v V2reportsNomenclatureAuthority)`

SetNomenclatureAuthority sets NomenclatureAuthority field to given value.

### HasNomenclatureAuthority

`func (o *V2reportsGeneDescriptor) HasNomenclatureAuthority() bool`

HasNomenclatureAuthority returns a boolean if a field has been set.

### GetSwissProtAccessions

`func (o *V2reportsGeneDescriptor) GetSwissProtAccessions() []string`

GetSwissProtAccessions returns the SwissProtAccessions field if non-nil, zero value otherwise.

### GetSwissProtAccessionsOk

`func (o *V2reportsGeneDescriptor) GetSwissProtAccessionsOk() (*[]string, bool)`

GetSwissProtAccessionsOk returns a tuple with the SwissProtAccessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwissProtAccessions

`func (o *V2reportsGeneDescriptor) SetSwissProtAccessions(v []string)`

SetSwissProtAccessions sets SwissProtAccessions field to given value.

### HasSwissProtAccessions

`func (o *V2reportsGeneDescriptor) HasSwissProtAccessions() bool`

HasSwissProtAccessions returns a boolean if a field has been set.

### GetEnsemblGeneIds

`func (o *V2reportsGeneDescriptor) GetEnsemblGeneIds() []string`

GetEnsemblGeneIds returns the EnsemblGeneIds field if non-nil, zero value otherwise.

### GetEnsemblGeneIdsOk

`func (o *V2reportsGeneDescriptor) GetEnsemblGeneIdsOk() (*[]string, bool)`

GetEnsemblGeneIdsOk returns a tuple with the EnsemblGeneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsemblGeneIds

`func (o *V2reportsGeneDescriptor) SetEnsemblGeneIds(v []string)`

SetEnsemblGeneIds sets EnsemblGeneIds field to given value.

### HasEnsemblGeneIds

`func (o *V2reportsGeneDescriptor) HasEnsemblGeneIds() bool`

HasEnsemblGeneIds returns a boolean if a field has been set.

### GetOmimIds

`func (o *V2reportsGeneDescriptor) GetOmimIds() []string`

GetOmimIds returns the OmimIds field if non-nil, zero value otherwise.

### GetOmimIdsOk

`func (o *V2reportsGeneDescriptor) GetOmimIdsOk() (*[]string, bool)`

GetOmimIdsOk returns a tuple with the OmimIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmimIds

`func (o *V2reportsGeneDescriptor) SetOmimIds(v []string)`

SetOmimIds sets OmimIds field to given value.

### HasOmimIds

`func (o *V2reportsGeneDescriptor) HasOmimIds() bool`

HasOmimIds returns a boolean if a field has been set.

### GetSynonyms

`func (o *V2reportsGeneDescriptor) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *V2reportsGeneDescriptor) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *V2reportsGeneDescriptor) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *V2reportsGeneDescriptor) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetReplacedGeneId

`func (o *V2reportsGeneDescriptor) GetReplacedGeneId() string`

GetReplacedGeneId returns the ReplacedGeneId field if non-nil, zero value otherwise.

### GetReplacedGeneIdOk

`func (o *V2reportsGeneDescriptor) GetReplacedGeneIdOk() (*string, bool)`

GetReplacedGeneIdOk returns a tuple with the ReplacedGeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedGeneId

`func (o *V2reportsGeneDescriptor) SetReplacedGeneId(v string)`

SetReplacedGeneId sets ReplacedGeneId field to given value.

### HasReplacedGeneId

`func (o *V2reportsGeneDescriptor) HasReplacedGeneId() bool`

HasReplacedGeneId returns a boolean if a field has been set.

### GetAnnotations

`func (o *V2reportsGeneDescriptor) GetAnnotations() []V2reportsAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *V2reportsGeneDescriptor) GetAnnotationsOk() (*[]V2reportsAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *V2reportsGeneDescriptor) SetAnnotations(v []V2reportsAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *V2reportsGeneDescriptor) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetTranscriptCount

`func (o *V2reportsGeneDescriptor) GetTranscriptCount() int32`

GetTranscriptCount returns the TranscriptCount field if non-nil, zero value otherwise.

### GetTranscriptCountOk

`func (o *V2reportsGeneDescriptor) GetTranscriptCountOk() (*int32, bool)`

GetTranscriptCountOk returns a tuple with the TranscriptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptCount

`func (o *V2reportsGeneDescriptor) SetTranscriptCount(v int32)`

SetTranscriptCount sets TranscriptCount field to given value.

### HasTranscriptCount

`func (o *V2reportsGeneDescriptor) HasTranscriptCount() bool`

HasTranscriptCount returns a boolean if a field has been set.

### GetProteinCount

`func (o *V2reportsGeneDescriptor) GetProteinCount() int32`

GetProteinCount returns the ProteinCount field if non-nil, zero value otherwise.

### GetProteinCountOk

`func (o *V2reportsGeneDescriptor) GetProteinCountOk() (*int32, bool)`

GetProteinCountOk returns a tuple with the ProteinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinCount

`func (o *V2reportsGeneDescriptor) SetProteinCount(v int32)`

SetProteinCount sets ProteinCount field to given value.

### HasProteinCount

`func (o *V2reportsGeneDescriptor) HasProteinCount() bool`

HasProteinCount returns a boolean if a field has been set.

### GetTranscriptTypeCounts

`func (o *V2reportsGeneDescriptor) GetTranscriptTypeCounts() []V2reportsTranscriptTypeCount`

GetTranscriptTypeCounts returns the TranscriptTypeCounts field if non-nil, zero value otherwise.

### GetTranscriptTypeCountsOk

`func (o *V2reportsGeneDescriptor) GetTranscriptTypeCountsOk() (*[]V2reportsTranscriptTypeCount, bool)`

GetTranscriptTypeCountsOk returns a tuple with the TranscriptTypeCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptTypeCounts

`func (o *V2reportsGeneDescriptor) SetTranscriptTypeCounts(v []V2reportsTranscriptTypeCount)`

SetTranscriptTypeCounts sets TranscriptTypeCounts field to given value.

### HasTranscriptTypeCounts

`func (o *V2reportsGeneDescriptor) HasTranscriptTypeCounts() bool`

HasTranscriptTypeCounts returns a boolean if a field has been set.

### GetGeneGroups

`func (o *V2reportsGeneDescriptor) GetGeneGroups() []V2reportsGeneGroup`

GetGeneGroups returns the GeneGroups field if non-nil, zero value otherwise.

### GetGeneGroupsOk

`func (o *V2reportsGeneDescriptor) GetGeneGroupsOk() (*[]V2reportsGeneGroup, bool)`

GetGeneGroupsOk returns a tuple with the GeneGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneGroups

`func (o *V2reportsGeneDescriptor) SetGeneGroups(v []V2reportsGeneGroup)`

SetGeneGroups sets GeneGroups field to given value.

### HasGeneGroups

`func (o *V2reportsGeneDescriptor) HasGeneGroups() bool`

HasGeneGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


