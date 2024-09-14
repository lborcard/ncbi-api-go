# V2reportsVirusAssembly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accession** | Pointer to **string** |  | [optional] 
**IsComplete** | Pointer to **bool** |  | [optional] 
**IsAnnotated** | Pointer to **bool** |  | [optional] 
**Isolate** | Pointer to [**V2reportsIsolate**](V2reportsIsolate.md) |  | [optional] 
**SourceDatabase** | Pointer to **string** |  | [optional] 
**ProteinCount** | Pointer to **int32** |  | [optional] 
**Host** | Pointer to [**V2reportsOrganism**](V2reportsOrganism.md) |  | [optional] 
**Virus** | Pointer to [**V2reportsOrganism**](V2reportsOrganism.md) |  | [optional] 
**Bioprojects** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to [**V2reportsVirusAssemblyCollectionLocation**](V2reportsVirusAssemblyCollectionLocation.md) |  | [optional] 
**UpdateDate** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**NucleotideCompleteness** | Pointer to **string** |  | [optional] 
**Completeness** | Pointer to [**V2reportsVirusAssemblyCompleteness**](V2reportsVirusAssemblyCompleteness.md) |  | [optional] [default to V2REPORTSVIRUSASSEMBLYCOMPLETENESS_UNKNOWN]
**Length** | Pointer to **int32** |  | [optional] 
**GeneCount** | Pointer to **int32** |  | [optional] 
**MaturePeptideCount** | Pointer to **int32** |  | [optional] 
**Biosample** | Pointer to **string** |  | [optional] 
**MolType** | Pointer to **string** |  | [optional] 
**Nucleotide** | Pointer to [**V2reportsSeqRangeSetFasta**](V2reportsSeqRangeSetFasta.md) |  | [optional] 
**PurposeOfSampling** | Pointer to [**V2reportsPurposeOfSampling**](V2reportsPurposeOfSampling.md) |  | [optional] [default to V2REPORTSPURPOSEOFSAMPLING_UNKNOWN]
**SraAccessions** | Pointer to **[]string** |  | [optional] 
**Submitter** | Pointer to [**V2reportsVirusAssemblySubmitterInfo**](V2reportsVirusAssemblySubmitterInfo.md) |  | [optional] 
**LabHost** | Pointer to **string** |  | [optional] 
**IsLabHost** | Pointer to **bool** |  | [optional] 
**IsVaccineStrain** | Pointer to **bool** |  | [optional] 

## Methods

### NewV2reportsVirusAssembly

`func NewV2reportsVirusAssembly() *V2reportsVirusAssembly`

NewV2reportsVirusAssembly instantiates a new V2reportsVirusAssembly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsVirusAssemblyWithDefaults

`func NewV2reportsVirusAssemblyWithDefaults() *V2reportsVirusAssembly`

NewV2reportsVirusAssemblyWithDefaults instantiates a new V2reportsVirusAssembly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccession

`func (o *V2reportsVirusAssembly) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V2reportsVirusAssembly) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V2reportsVirusAssembly) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V2reportsVirusAssembly) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetIsComplete

`func (o *V2reportsVirusAssembly) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *V2reportsVirusAssembly) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *V2reportsVirusAssembly) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.

### HasIsComplete

`func (o *V2reportsVirusAssembly) HasIsComplete() bool`

HasIsComplete returns a boolean if a field has been set.

### GetIsAnnotated

`func (o *V2reportsVirusAssembly) GetIsAnnotated() bool`

GetIsAnnotated returns the IsAnnotated field if non-nil, zero value otherwise.

### GetIsAnnotatedOk

`func (o *V2reportsVirusAssembly) GetIsAnnotatedOk() (*bool, bool)`

GetIsAnnotatedOk returns a tuple with the IsAnnotated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnnotated

`func (o *V2reportsVirusAssembly) SetIsAnnotated(v bool)`

SetIsAnnotated sets IsAnnotated field to given value.

### HasIsAnnotated

`func (o *V2reportsVirusAssembly) HasIsAnnotated() bool`

HasIsAnnotated returns a boolean if a field has been set.

### GetIsolate

`func (o *V2reportsVirusAssembly) GetIsolate() V2reportsIsolate`

GetIsolate returns the Isolate field if non-nil, zero value otherwise.

### GetIsolateOk

`func (o *V2reportsVirusAssembly) GetIsolateOk() (*V2reportsIsolate, bool)`

GetIsolateOk returns a tuple with the Isolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolate

`func (o *V2reportsVirusAssembly) SetIsolate(v V2reportsIsolate)`

SetIsolate sets Isolate field to given value.

### HasIsolate

`func (o *V2reportsVirusAssembly) HasIsolate() bool`

HasIsolate returns a boolean if a field has been set.

### GetSourceDatabase

`func (o *V2reportsVirusAssembly) GetSourceDatabase() string`

GetSourceDatabase returns the SourceDatabase field if non-nil, zero value otherwise.

### GetSourceDatabaseOk

`func (o *V2reportsVirusAssembly) GetSourceDatabaseOk() (*string, bool)`

GetSourceDatabaseOk returns a tuple with the SourceDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatabase

`func (o *V2reportsVirusAssembly) SetSourceDatabase(v string)`

SetSourceDatabase sets SourceDatabase field to given value.

### HasSourceDatabase

`func (o *V2reportsVirusAssembly) HasSourceDatabase() bool`

HasSourceDatabase returns a boolean if a field has been set.

### GetProteinCount

`func (o *V2reportsVirusAssembly) GetProteinCount() int32`

GetProteinCount returns the ProteinCount field if non-nil, zero value otherwise.

### GetProteinCountOk

`func (o *V2reportsVirusAssembly) GetProteinCountOk() (*int32, bool)`

GetProteinCountOk returns a tuple with the ProteinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinCount

`func (o *V2reportsVirusAssembly) SetProteinCount(v int32)`

SetProteinCount sets ProteinCount field to given value.

### HasProteinCount

`func (o *V2reportsVirusAssembly) HasProteinCount() bool`

HasProteinCount returns a boolean if a field has been set.

### GetHost

`func (o *V2reportsVirusAssembly) GetHost() V2reportsOrganism`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V2reportsVirusAssembly) GetHostOk() (*V2reportsOrganism, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V2reportsVirusAssembly) SetHost(v V2reportsOrganism)`

SetHost sets Host field to given value.

### HasHost

`func (o *V2reportsVirusAssembly) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetVirus

`func (o *V2reportsVirusAssembly) GetVirus() V2reportsOrganism`

GetVirus returns the Virus field if non-nil, zero value otherwise.

### GetVirusOk

`func (o *V2reportsVirusAssembly) GetVirusOk() (*V2reportsOrganism, bool)`

GetVirusOk returns a tuple with the Virus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirus

`func (o *V2reportsVirusAssembly) SetVirus(v V2reportsOrganism)`

SetVirus sets Virus field to given value.

### HasVirus

`func (o *V2reportsVirusAssembly) HasVirus() bool`

HasVirus returns a boolean if a field has been set.

### GetBioprojects

`func (o *V2reportsVirusAssembly) GetBioprojects() []string`

GetBioprojects returns the Bioprojects field if non-nil, zero value otherwise.

### GetBioprojectsOk

`func (o *V2reportsVirusAssembly) GetBioprojectsOk() (*[]string, bool)`

GetBioprojectsOk returns a tuple with the Bioprojects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioprojects

`func (o *V2reportsVirusAssembly) SetBioprojects(v []string)`

SetBioprojects sets Bioprojects field to given value.

### HasBioprojects

`func (o *V2reportsVirusAssembly) HasBioprojects() bool`

HasBioprojects returns a boolean if a field has been set.

### GetLocation

`func (o *V2reportsVirusAssembly) GetLocation() V2reportsVirusAssemblyCollectionLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V2reportsVirusAssembly) GetLocationOk() (*V2reportsVirusAssemblyCollectionLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V2reportsVirusAssembly) SetLocation(v V2reportsVirusAssemblyCollectionLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V2reportsVirusAssembly) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetUpdateDate

`func (o *V2reportsVirusAssembly) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *V2reportsVirusAssembly) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *V2reportsVirusAssembly) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *V2reportsVirusAssembly) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetReleaseDate

`func (o *V2reportsVirusAssembly) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *V2reportsVirusAssembly) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *V2reportsVirusAssembly) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *V2reportsVirusAssembly) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetNucleotideCompleteness

`func (o *V2reportsVirusAssembly) GetNucleotideCompleteness() string`

GetNucleotideCompleteness returns the NucleotideCompleteness field if non-nil, zero value otherwise.

### GetNucleotideCompletenessOk

`func (o *V2reportsVirusAssembly) GetNucleotideCompletenessOk() (*string, bool)`

GetNucleotideCompletenessOk returns a tuple with the NucleotideCompleteness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleotideCompleteness

`func (o *V2reportsVirusAssembly) SetNucleotideCompleteness(v string)`

SetNucleotideCompleteness sets NucleotideCompleteness field to given value.

### HasNucleotideCompleteness

`func (o *V2reportsVirusAssembly) HasNucleotideCompleteness() bool`

HasNucleotideCompleteness returns a boolean if a field has been set.

### GetCompleteness

`func (o *V2reportsVirusAssembly) GetCompleteness() V2reportsVirusAssemblyCompleteness`

GetCompleteness returns the Completeness field if non-nil, zero value otherwise.

### GetCompletenessOk

`func (o *V2reportsVirusAssembly) GetCompletenessOk() (*V2reportsVirusAssemblyCompleteness, bool)`

GetCompletenessOk returns a tuple with the Completeness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteness

`func (o *V2reportsVirusAssembly) SetCompleteness(v V2reportsVirusAssemblyCompleteness)`

SetCompleteness sets Completeness field to given value.

### HasCompleteness

`func (o *V2reportsVirusAssembly) HasCompleteness() bool`

HasCompleteness returns a boolean if a field has been set.

### GetLength

`func (o *V2reportsVirusAssembly) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *V2reportsVirusAssembly) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *V2reportsVirusAssembly) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *V2reportsVirusAssembly) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetGeneCount

`func (o *V2reportsVirusAssembly) GetGeneCount() int32`

GetGeneCount returns the GeneCount field if non-nil, zero value otherwise.

### GetGeneCountOk

`func (o *V2reportsVirusAssembly) GetGeneCountOk() (*int32, bool)`

GetGeneCountOk returns a tuple with the GeneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneCount

`func (o *V2reportsVirusAssembly) SetGeneCount(v int32)`

SetGeneCount sets GeneCount field to given value.

### HasGeneCount

`func (o *V2reportsVirusAssembly) HasGeneCount() bool`

HasGeneCount returns a boolean if a field has been set.

### GetMaturePeptideCount

`func (o *V2reportsVirusAssembly) GetMaturePeptideCount() int32`

GetMaturePeptideCount returns the MaturePeptideCount field if non-nil, zero value otherwise.

### GetMaturePeptideCountOk

`func (o *V2reportsVirusAssembly) GetMaturePeptideCountOk() (*int32, bool)`

GetMaturePeptideCountOk returns a tuple with the MaturePeptideCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturePeptideCount

`func (o *V2reportsVirusAssembly) SetMaturePeptideCount(v int32)`

SetMaturePeptideCount sets MaturePeptideCount field to given value.

### HasMaturePeptideCount

`func (o *V2reportsVirusAssembly) HasMaturePeptideCount() bool`

HasMaturePeptideCount returns a boolean if a field has been set.

### GetBiosample

`func (o *V2reportsVirusAssembly) GetBiosample() string`

GetBiosample returns the Biosample field if non-nil, zero value otherwise.

### GetBiosampleOk

`func (o *V2reportsVirusAssembly) GetBiosampleOk() (*string, bool)`

GetBiosampleOk returns a tuple with the Biosample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosample

`func (o *V2reportsVirusAssembly) SetBiosample(v string)`

SetBiosample sets Biosample field to given value.

### HasBiosample

`func (o *V2reportsVirusAssembly) HasBiosample() bool`

HasBiosample returns a boolean if a field has been set.

### GetMolType

`func (o *V2reportsVirusAssembly) GetMolType() string`

GetMolType returns the MolType field if non-nil, zero value otherwise.

### GetMolTypeOk

`func (o *V2reportsVirusAssembly) GetMolTypeOk() (*string, bool)`

GetMolTypeOk returns a tuple with the MolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMolType

`func (o *V2reportsVirusAssembly) SetMolType(v string)`

SetMolType sets MolType field to given value.

### HasMolType

`func (o *V2reportsVirusAssembly) HasMolType() bool`

HasMolType returns a boolean if a field has been set.

### GetNucleotide

`func (o *V2reportsVirusAssembly) GetNucleotide() V2reportsSeqRangeSetFasta`

GetNucleotide returns the Nucleotide field if non-nil, zero value otherwise.

### GetNucleotideOk

`func (o *V2reportsVirusAssembly) GetNucleotideOk() (*V2reportsSeqRangeSetFasta, bool)`

GetNucleotideOk returns a tuple with the Nucleotide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleotide

`func (o *V2reportsVirusAssembly) SetNucleotide(v V2reportsSeqRangeSetFasta)`

SetNucleotide sets Nucleotide field to given value.

### HasNucleotide

`func (o *V2reportsVirusAssembly) HasNucleotide() bool`

HasNucleotide returns a boolean if a field has been set.

### GetPurposeOfSampling

`func (o *V2reportsVirusAssembly) GetPurposeOfSampling() V2reportsPurposeOfSampling`

GetPurposeOfSampling returns the PurposeOfSampling field if non-nil, zero value otherwise.

### GetPurposeOfSamplingOk

`func (o *V2reportsVirusAssembly) GetPurposeOfSamplingOk() (*V2reportsPurposeOfSampling, bool)`

GetPurposeOfSamplingOk returns a tuple with the PurposeOfSampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeOfSampling

`func (o *V2reportsVirusAssembly) SetPurposeOfSampling(v V2reportsPurposeOfSampling)`

SetPurposeOfSampling sets PurposeOfSampling field to given value.

### HasPurposeOfSampling

`func (o *V2reportsVirusAssembly) HasPurposeOfSampling() bool`

HasPurposeOfSampling returns a boolean if a field has been set.

### GetSraAccessions

`func (o *V2reportsVirusAssembly) GetSraAccessions() []string`

GetSraAccessions returns the SraAccessions field if non-nil, zero value otherwise.

### GetSraAccessionsOk

`func (o *V2reportsVirusAssembly) GetSraAccessionsOk() (*[]string, bool)`

GetSraAccessionsOk returns a tuple with the SraAccessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSraAccessions

`func (o *V2reportsVirusAssembly) SetSraAccessions(v []string)`

SetSraAccessions sets SraAccessions field to given value.

### HasSraAccessions

`func (o *V2reportsVirusAssembly) HasSraAccessions() bool`

HasSraAccessions returns a boolean if a field has been set.

### GetSubmitter

`func (o *V2reportsVirusAssembly) GetSubmitter() V2reportsVirusAssemblySubmitterInfo`

GetSubmitter returns the Submitter field if non-nil, zero value otherwise.

### GetSubmitterOk

`func (o *V2reportsVirusAssembly) GetSubmitterOk() (*V2reportsVirusAssemblySubmitterInfo, bool)`

GetSubmitterOk returns a tuple with the Submitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitter

`func (o *V2reportsVirusAssembly) SetSubmitter(v V2reportsVirusAssemblySubmitterInfo)`

SetSubmitter sets Submitter field to given value.

### HasSubmitter

`func (o *V2reportsVirusAssembly) HasSubmitter() bool`

HasSubmitter returns a boolean if a field has been set.

### GetLabHost

`func (o *V2reportsVirusAssembly) GetLabHost() string`

GetLabHost returns the LabHost field if non-nil, zero value otherwise.

### GetLabHostOk

`func (o *V2reportsVirusAssembly) GetLabHostOk() (*string, bool)`

GetLabHostOk returns a tuple with the LabHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabHost

`func (o *V2reportsVirusAssembly) SetLabHost(v string)`

SetLabHost sets LabHost field to given value.

### HasLabHost

`func (o *V2reportsVirusAssembly) HasLabHost() bool`

HasLabHost returns a boolean if a field has been set.

### GetIsLabHost

`func (o *V2reportsVirusAssembly) GetIsLabHost() bool`

GetIsLabHost returns the IsLabHost field if non-nil, zero value otherwise.

### GetIsLabHostOk

`func (o *V2reportsVirusAssembly) GetIsLabHostOk() (*bool, bool)`

GetIsLabHostOk returns a tuple with the IsLabHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLabHost

`func (o *V2reportsVirusAssembly) SetIsLabHost(v bool)`

SetIsLabHost sets IsLabHost field to given value.

### HasIsLabHost

`func (o *V2reportsVirusAssembly) HasIsLabHost() bool`

HasIsLabHost returns a boolean if a field has been set.

### GetIsVaccineStrain

`func (o *V2reportsVirusAssembly) GetIsVaccineStrain() bool`

GetIsVaccineStrain returns the IsVaccineStrain field if non-nil, zero value otherwise.

### GetIsVaccineStrainOk

`func (o *V2reportsVirusAssembly) GetIsVaccineStrainOk() (*bool, bool)`

GetIsVaccineStrainOk returns a tuple with the IsVaccineStrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVaccineStrain

`func (o *V2reportsVirusAssembly) SetIsVaccineStrain(v bool)`

SetIsVaccineStrain sets IsVaccineStrain field to given value.

### HasIsVaccineStrain

`func (o *V2reportsVirusAssembly) HasIsVaccineStrain() bool`

HasIsVaccineStrain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


