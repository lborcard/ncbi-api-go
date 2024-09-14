# V2reportsVirusGene

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**GeneId** | Pointer to **int32** |  | [optional] 
**Nucleotide** | Pointer to [**V2reportsSeqRangeSetFasta**](V2reportsSeqRangeSetFasta.md) |  | [optional] 
**Cds** | Pointer to [**[]V2reportsVirusPeptide**](V2reportsVirusPeptide.md) |  | [optional] 

## Methods

### NewV2reportsVirusGene

`func NewV2reportsVirusGene() *V2reportsVirusGene`

NewV2reportsVirusGene instantiates a new V2reportsVirusGene object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsVirusGeneWithDefaults

`func NewV2reportsVirusGeneWithDefaults() *V2reportsVirusGene`

NewV2reportsVirusGeneWithDefaults instantiates a new V2reportsVirusGene object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2reportsVirusGene) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2reportsVirusGene) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2reportsVirusGene) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2reportsVirusGene) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGeneId

`func (o *V2reportsVirusGene) GetGeneId() int32`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *V2reportsVirusGene) GetGeneIdOk() (*int32, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *V2reportsVirusGene) SetGeneId(v int32)`

SetGeneId sets GeneId field to given value.

### HasGeneId

`func (o *V2reportsVirusGene) HasGeneId() bool`

HasGeneId returns a boolean if a field has been set.

### GetNucleotide

`func (o *V2reportsVirusGene) GetNucleotide() V2reportsSeqRangeSetFasta`

GetNucleotide returns the Nucleotide field if non-nil, zero value otherwise.

### GetNucleotideOk

`func (o *V2reportsVirusGene) GetNucleotideOk() (*V2reportsSeqRangeSetFasta, bool)`

GetNucleotideOk returns a tuple with the Nucleotide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleotide

`func (o *V2reportsVirusGene) SetNucleotide(v V2reportsSeqRangeSetFasta)`

SetNucleotide sets Nucleotide field to given value.

### HasNucleotide

`func (o *V2reportsVirusGene) HasNucleotide() bool`

HasNucleotide returns a boolean if a field has been set.

### GetCds

`func (o *V2reportsVirusGene) GetCds() []V2reportsVirusPeptide`

GetCds returns the Cds field if non-nil, zero value otherwise.

### GetCdsOk

`func (o *V2reportsVirusGene) GetCdsOk() (*[]V2reportsVirusPeptide, bool)`

GetCdsOk returns a tuple with the Cds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCds

`func (o *V2reportsVirusGene) SetCds(v []V2reportsVirusPeptide)`

SetCds sets Cds field to given value.

### HasCds

`func (o *V2reportsVirusGene) HasCds() bool`

HasCds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


