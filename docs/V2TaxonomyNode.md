# V2TaxonomyNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxId** | Pointer to **int32** |  | [optional] 
**OrganismName** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**GenbankCommonName** | Pointer to **string** |  | [optional] 
**Acronyms** | Pointer to **[]string** |  | [optional] 
**GenbankAcronym** | Pointer to **string** |  | [optional] 
**BlastName** | Pointer to **string** |  | [optional] 
**Lineage** | Pointer to **[]int32** |  | [optional] 
**Children** | Pointer to **[]int32** |  | [optional] 
**DescendentWithDescribedSpeciesNamesCount** | Pointer to **int32** |  | [optional] 
**Rank** | Pointer to [**V2reportsRankType**](V2reportsRankType.md) |  | [optional] [default to V2REPORTSRANKTYPE_NO_RANK]
**HasDescribedSpeciesName** | Pointer to **bool** |  | [optional] 
**Counts** | Pointer to [**[]V2TaxonomyNodeCountByType**](V2TaxonomyNodeCountByType.md) |  | [optional] 
**MinOrd** | Pointer to **int32** |  | [optional] 
**MaxOrd** | Pointer to **int32** |  | [optional] 
**Extinct** | Pointer to **bool** |  | [optional] 
**GenomicMoltype** | Pointer to **string** |  | [optional] 
**IndexedNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV2TaxonomyNode

`func NewV2TaxonomyNode() *V2TaxonomyNode`

NewV2TaxonomyNode instantiates a new V2TaxonomyNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TaxonomyNodeWithDefaults

`func NewV2TaxonomyNodeWithDefaults() *V2TaxonomyNode`

NewV2TaxonomyNodeWithDefaults instantiates a new V2TaxonomyNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxId

`func (o *V2TaxonomyNode) GetTaxId() int32`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V2TaxonomyNode) GetTaxIdOk() (*int32, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V2TaxonomyNode) SetTaxId(v int32)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V2TaxonomyNode) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetOrganismName

`func (o *V2TaxonomyNode) GetOrganismName() string`

GetOrganismName returns the OrganismName field if non-nil, zero value otherwise.

### GetOrganismNameOk

`func (o *V2TaxonomyNode) GetOrganismNameOk() (*string, bool)`

GetOrganismNameOk returns a tuple with the OrganismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganismName

`func (o *V2TaxonomyNode) SetOrganismName(v string)`

SetOrganismName sets OrganismName field to given value.

### HasOrganismName

`func (o *V2TaxonomyNode) HasOrganismName() bool`

HasOrganismName returns a boolean if a field has been set.

### GetCommonName

`func (o *V2TaxonomyNode) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V2TaxonomyNode) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V2TaxonomyNode) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V2TaxonomyNode) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetGenbankCommonName

`func (o *V2TaxonomyNode) GetGenbankCommonName() string`

GetGenbankCommonName returns the GenbankCommonName field if non-nil, zero value otherwise.

### GetGenbankCommonNameOk

`func (o *V2TaxonomyNode) GetGenbankCommonNameOk() (*string, bool)`

GetGenbankCommonNameOk returns a tuple with the GenbankCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenbankCommonName

`func (o *V2TaxonomyNode) SetGenbankCommonName(v string)`

SetGenbankCommonName sets GenbankCommonName field to given value.

### HasGenbankCommonName

`func (o *V2TaxonomyNode) HasGenbankCommonName() bool`

HasGenbankCommonName returns a boolean if a field has been set.

### GetAcronyms

`func (o *V2TaxonomyNode) GetAcronyms() []string`

GetAcronyms returns the Acronyms field if non-nil, zero value otherwise.

### GetAcronymsOk

`func (o *V2TaxonomyNode) GetAcronymsOk() (*[]string, bool)`

GetAcronymsOk returns a tuple with the Acronyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcronyms

`func (o *V2TaxonomyNode) SetAcronyms(v []string)`

SetAcronyms sets Acronyms field to given value.

### HasAcronyms

`func (o *V2TaxonomyNode) HasAcronyms() bool`

HasAcronyms returns a boolean if a field has been set.

### GetGenbankAcronym

`func (o *V2TaxonomyNode) GetGenbankAcronym() string`

GetGenbankAcronym returns the GenbankAcronym field if non-nil, zero value otherwise.

### GetGenbankAcronymOk

`func (o *V2TaxonomyNode) GetGenbankAcronymOk() (*string, bool)`

GetGenbankAcronymOk returns a tuple with the GenbankAcronym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenbankAcronym

`func (o *V2TaxonomyNode) SetGenbankAcronym(v string)`

SetGenbankAcronym sets GenbankAcronym field to given value.

### HasGenbankAcronym

`func (o *V2TaxonomyNode) HasGenbankAcronym() bool`

HasGenbankAcronym returns a boolean if a field has been set.

### GetBlastName

`func (o *V2TaxonomyNode) GetBlastName() string`

GetBlastName returns the BlastName field if non-nil, zero value otherwise.

### GetBlastNameOk

`func (o *V2TaxonomyNode) GetBlastNameOk() (*string, bool)`

GetBlastNameOk returns a tuple with the BlastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlastName

`func (o *V2TaxonomyNode) SetBlastName(v string)`

SetBlastName sets BlastName field to given value.

### HasBlastName

`func (o *V2TaxonomyNode) HasBlastName() bool`

HasBlastName returns a boolean if a field has been set.

### GetLineage

`func (o *V2TaxonomyNode) GetLineage() []int32`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *V2TaxonomyNode) GetLineageOk() (*[]int32, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *V2TaxonomyNode) SetLineage(v []int32)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *V2TaxonomyNode) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetChildren

`func (o *V2TaxonomyNode) GetChildren() []int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *V2TaxonomyNode) GetChildrenOk() (*[]int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *V2TaxonomyNode) SetChildren(v []int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *V2TaxonomyNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetDescendentWithDescribedSpeciesNamesCount

`func (o *V2TaxonomyNode) GetDescendentWithDescribedSpeciesNamesCount() int32`

GetDescendentWithDescribedSpeciesNamesCount returns the DescendentWithDescribedSpeciesNamesCount field if non-nil, zero value otherwise.

### GetDescendentWithDescribedSpeciesNamesCountOk

`func (o *V2TaxonomyNode) GetDescendentWithDescribedSpeciesNamesCountOk() (*int32, bool)`

GetDescendentWithDescribedSpeciesNamesCountOk returns a tuple with the DescendentWithDescribedSpeciesNamesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescendentWithDescribedSpeciesNamesCount

`func (o *V2TaxonomyNode) SetDescendentWithDescribedSpeciesNamesCount(v int32)`

SetDescendentWithDescribedSpeciesNamesCount sets DescendentWithDescribedSpeciesNamesCount field to given value.

### HasDescendentWithDescribedSpeciesNamesCount

`func (o *V2TaxonomyNode) HasDescendentWithDescribedSpeciesNamesCount() bool`

HasDescendentWithDescribedSpeciesNamesCount returns a boolean if a field has been set.

### GetRank

`func (o *V2TaxonomyNode) GetRank() V2reportsRankType`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *V2TaxonomyNode) GetRankOk() (*V2reportsRankType, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *V2TaxonomyNode) SetRank(v V2reportsRankType)`

SetRank sets Rank field to given value.

### HasRank

`func (o *V2TaxonomyNode) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetHasDescribedSpeciesName

`func (o *V2TaxonomyNode) GetHasDescribedSpeciesName() bool`

GetHasDescribedSpeciesName returns the HasDescribedSpeciesName field if non-nil, zero value otherwise.

### GetHasDescribedSpeciesNameOk

`func (o *V2TaxonomyNode) GetHasDescribedSpeciesNameOk() (*bool, bool)`

GetHasDescribedSpeciesNameOk returns a tuple with the HasDescribedSpeciesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDescribedSpeciesName

`func (o *V2TaxonomyNode) SetHasDescribedSpeciesName(v bool)`

SetHasDescribedSpeciesName sets HasDescribedSpeciesName field to given value.

### HasHasDescribedSpeciesName

`func (o *V2TaxonomyNode) HasHasDescribedSpeciesName() bool`

HasHasDescribedSpeciesName returns a boolean if a field has been set.

### GetCounts

`func (o *V2TaxonomyNode) GetCounts() []V2TaxonomyNodeCountByType`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *V2TaxonomyNode) GetCountsOk() (*[]V2TaxonomyNodeCountByType, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *V2TaxonomyNode) SetCounts(v []V2TaxonomyNodeCountByType)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *V2TaxonomyNode) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetMinOrd

`func (o *V2TaxonomyNode) GetMinOrd() int32`

GetMinOrd returns the MinOrd field if non-nil, zero value otherwise.

### GetMinOrdOk

`func (o *V2TaxonomyNode) GetMinOrdOk() (*int32, bool)`

GetMinOrdOk returns a tuple with the MinOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrd

`func (o *V2TaxonomyNode) SetMinOrd(v int32)`

SetMinOrd sets MinOrd field to given value.

### HasMinOrd

`func (o *V2TaxonomyNode) HasMinOrd() bool`

HasMinOrd returns a boolean if a field has been set.

### GetMaxOrd

`func (o *V2TaxonomyNode) GetMaxOrd() int32`

GetMaxOrd returns the MaxOrd field if non-nil, zero value otherwise.

### GetMaxOrdOk

`func (o *V2TaxonomyNode) GetMaxOrdOk() (*int32, bool)`

GetMaxOrdOk returns a tuple with the MaxOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrd

`func (o *V2TaxonomyNode) SetMaxOrd(v int32)`

SetMaxOrd sets MaxOrd field to given value.

### HasMaxOrd

`func (o *V2TaxonomyNode) HasMaxOrd() bool`

HasMaxOrd returns a boolean if a field has been set.

### GetExtinct

`func (o *V2TaxonomyNode) GetExtinct() bool`

GetExtinct returns the Extinct field if non-nil, zero value otherwise.

### GetExtinctOk

`func (o *V2TaxonomyNode) GetExtinctOk() (*bool, bool)`

GetExtinctOk returns a tuple with the Extinct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtinct

`func (o *V2TaxonomyNode) SetExtinct(v bool)`

SetExtinct sets Extinct field to given value.

### HasExtinct

`func (o *V2TaxonomyNode) HasExtinct() bool`

HasExtinct returns a boolean if a field has been set.

### GetGenomicMoltype

`func (o *V2TaxonomyNode) GetGenomicMoltype() string`

GetGenomicMoltype returns the GenomicMoltype field if non-nil, zero value otherwise.

### GetGenomicMoltypeOk

`func (o *V2TaxonomyNode) GetGenomicMoltypeOk() (*string, bool)`

GetGenomicMoltypeOk returns a tuple with the GenomicMoltype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicMoltype

`func (o *V2TaxonomyNode) SetGenomicMoltype(v string)`

SetGenomicMoltype sets GenomicMoltype field to given value.

### HasGenomicMoltype

`func (o *V2TaxonomyNode) HasGenomicMoltype() bool`

HasGenomicMoltype returns a boolean if a field has been set.

### GetIndexedNames

`func (o *V2TaxonomyNode) GetIndexedNames() []string`

GetIndexedNames returns the IndexedNames field if non-nil, zero value otherwise.

### GetIndexedNamesOk

`func (o *V2TaxonomyNode) GetIndexedNamesOk() (*[]string, bool)`

GetIndexedNamesOk returns a tuple with the IndexedNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedNames

`func (o *V2TaxonomyNode) SetIndexedNames(v []string)`

SetIndexedNames sets IndexedNames field to given value.

### HasIndexedNames

`func (o *V2TaxonomyNode) HasIndexedNames() bool`

HasIndexedNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


