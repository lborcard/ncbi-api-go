# V2reportsVirusPeptide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accession** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OtherNames** | Pointer to **[]string** |  | [optional] 
**Nucleotide** | Pointer to [**V2reportsSeqRangeSetFasta**](V2reportsSeqRangeSetFasta.md) |  | [optional] 
**Protein** | Pointer to [**V2reportsSeqRangeSetFasta**](V2reportsSeqRangeSetFasta.md) |  | [optional] 
**PdbIds** | Pointer to **[]string** |  | [optional] 
**Cdd** | Pointer to [**[]V2reportsConservedDomain**](V2reportsConservedDomain.md) |  | [optional] 
**UniProtKb** | Pointer to [**V2reportsVirusPeptideUniProtId**](V2reportsVirusPeptideUniProtId.md) |  | [optional] 
**MaturePeptide** | Pointer to [**[]V2reportsVirusPeptide**](V2reportsVirusPeptide.md) |  | [optional] 
**ProteinCompleteness** | Pointer to [**V2reportsVirusPeptideViralPeptideCompleteness**](V2reportsVirusPeptideViralPeptideCompleteness.md) |  | [optional] [default to V2REPORTSVIRUSPEPTIDEVIRALPEPTIDECOMPLETENESS_UNKNOWN]

## Methods

### NewV2reportsVirusPeptide

`func NewV2reportsVirusPeptide() *V2reportsVirusPeptide`

NewV2reportsVirusPeptide instantiates a new V2reportsVirusPeptide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsVirusPeptideWithDefaults

`func NewV2reportsVirusPeptideWithDefaults() *V2reportsVirusPeptide`

NewV2reportsVirusPeptideWithDefaults instantiates a new V2reportsVirusPeptide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccession

`func (o *V2reportsVirusPeptide) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V2reportsVirusPeptide) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V2reportsVirusPeptide) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V2reportsVirusPeptide) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetName

`func (o *V2reportsVirusPeptide) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2reportsVirusPeptide) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2reportsVirusPeptide) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2reportsVirusPeptide) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOtherNames

`func (o *V2reportsVirusPeptide) GetOtherNames() []string`

GetOtherNames returns the OtherNames field if non-nil, zero value otherwise.

### GetOtherNamesOk

`func (o *V2reportsVirusPeptide) GetOtherNamesOk() (*[]string, bool)`

GetOtherNamesOk returns a tuple with the OtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNames

`func (o *V2reportsVirusPeptide) SetOtherNames(v []string)`

SetOtherNames sets OtherNames field to given value.

### HasOtherNames

`func (o *V2reportsVirusPeptide) HasOtherNames() bool`

HasOtherNames returns a boolean if a field has been set.

### GetNucleotide

`func (o *V2reportsVirusPeptide) GetNucleotide() V2reportsSeqRangeSetFasta`

GetNucleotide returns the Nucleotide field if non-nil, zero value otherwise.

### GetNucleotideOk

`func (o *V2reportsVirusPeptide) GetNucleotideOk() (*V2reportsSeqRangeSetFasta, bool)`

GetNucleotideOk returns a tuple with the Nucleotide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleotide

`func (o *V2reportsVirusPeptide) SetNucleotide(v V2reportsSeqRangeSetFasta)`

SetNucleotide sets Nucleotide field to given value.

### HasNucleotide

`func (o *V2reportsVirusPeptide) HasNucleotide() bool`

HasNucleotide returns a boolean if a field has been set.

### GetProtein

`func (o *V2reportsVirusPeptide) GetProtein() V2reportsSeqRangeSetFasta`

GetProtein returns the Protein field if non-nil, zero value otherwise.

### GetProteinOk

`func (o *V2reportsVirusPeptide) GetProteinOk() (*V2reportsSeqRangeSetFasta, bool)`

GetProteinOk returns a tuple with the Protein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtein

`func (o *V2reportsVirusPeptide) SetProtein(v V2reportsSeqRangeSetFasta)`

SetProtein sets Protein field to given value.

### HasProtein

`func (o *V2reportsVirusPeptide) HasProtein() bool`

HasProtein returns a boolean if a field has been set.

### GetPdbIds

`func (o *V2reportsVirusPeptide) GetPdbIds() []string`

GetPdbIds returns the PdbIds field if non-nil, zero value otherwise.

### GetPdbIdsOk

`func (o *V2reportsVirusPeptide) GetPdbIdsOk() (*[]string, bool)`

GetPdbIdsOk returns a tuple with the PdbIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdbIds

`func (o *V2reportsVirusPeptide) SetPdbIds(v []string)`

SetPdbIds sets PdbIds field to given value.

### HasPdbIds

`func (o *V2reportsVirusPeptide) HasPdbIds() bool`

HasPdbIds returns a boolean if a field has been set.

### GetCdd

`func (o *V2reportsVirusPeptide) GetCdd() []V2reportsConservedDomain`

GetCdd returns the Cdd field if non-nil, zero value otherwise.

### GetCddOk

`func (o *V2reportsVirusPeptide) GetCddOk() (*[]V2reportsConservedDomain, bool)`

GetCddOk returns a tuple with the Cdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdd

`func (o *V2reportsVirusPeptide) SetCdd(v []V2reportsConservedDomain)`

SetCdd sets Cdd field to given value.

### HasCdd

`func (o *V2reportsVirusPeptide) HasCdd() bool`

HasCdd returns a boolean if a field has been set.

### GetUniProtKb

`func (o *V2reportsVirusPeptide) GetUniProtKb() V2reportsVirusPeptideUniProtId`

GetUniProtKb returns the UniProtKb field if non-nil, zero value otherwise.

### GetUniProtKbOk

`func (o *V2reportsVirusPeptide) GetUniProtKbOk() (*V2reportsVirusPeptideUniProtId, bool)`

GetUniProtKbOk returns a tuple with the UniProtKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniProtKb

`func (o *V2reportsVirusPeptide) SetUniProtKb(v V2reportsVirusPeptideUniProtId)`

SetUniProtKb sets UniProtKb field to given value.

### HasUniProtKb

`func (o *V2reportsVirusPeptide) HasUniProtKb() bool`

HasUniProtKb returns a boolean if a field has been set.

### GetMaturePeptide

`func (o *V2reportsVirusPeptide) GetMaturePeptide() []V2reportsVirusPeptide`

GetMaturePeptide returns the MaturePeptide field if non-nil, zero value otherwise.

### GetMaturePeptideOk

`func (o *V2reportsVirusPeptide) GetMaturePeptideOk() (*[]V2reportsVirusPeptide, bool)`

GetMaturePeptideOk returns a tuple with the MaturePeptide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturePeptide

`func (o *V2reportsVirusPeptide) SetMaturePeptide(v []V2reportsVirusPeptide)`

SetMaturePeptide sets MaturePeptide field to given value.

### HasMaturePeptide

`func (o *V2reportsVirusPeptide) HasMaturePeptide() bool`

HasMaturePeptide returns a boolean if a field has been set.

### GetProteinCompleteness

`func (o *V2reportsVirusPeptide) GetProteinCompleteness() V2reportsVirusPeptideViralPeptideCompleteness`

GetProteinCompleteness returns the ProteinCompleteness field if non-nil, zero value otherwise.

### GetProteinCompletenessOk

`func (o *V2reportsVirusPeptide) GetProteinCompletenessOk() (*V2reportsVirusPeptideViralPeptideCompleteness, bool)`

GetProteinCompletenessOk returns a tuple with the ProteinCompleteness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinCompleteness

`func (o *V2reportsVirusPeptide) SetProteinCompleteness(v V2reportsVirusPeptideViralPeptideCompleteness)`

SetProteinCompleteness sets ProteinCompleteness field to given value.

### HasProteinCompleteness

`func (o *V2reportsVirusPeptide) HasProteinCompleteness() bool`

HasProteinCompleteness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


