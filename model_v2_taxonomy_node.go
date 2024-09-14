/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 2 API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated zip archive](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v2alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ncbi-openapi-v2_goland

import (
	"encoding/json"
)

// V2TaxonomyNode struct for V2TaxonomyNode
type V2TaxonomyNode struct {
	TaxId *int32 `json:"tax_id,omitempty"`
	OrganismName *string `json:"organism_name,omitempty"`
	CommonName *string `json:"common_name,omitempty"`
	GenbankCommonName *string `json:"genbank_common_name,omitempty"`
	Acronyms []string `json:"acronyms,omitempty"`
	GenbankAcronym *string `json:"genbank_acronym,omitempty"`
	BlastName *string `json:"blast_name,omitempty"`
	Lineage []int32 `json:"lineage,omitempty"`
	Children []int32 `json:"children,omitempty"`
	DescendentWithDescribedSpeciesNamesCount *int32 `json:"descendent_with_described_species_names_count,omitempty"`
	Rank *V2reportsRankType `json:"rank,omitempty"`
	HasDescribedSpeciesName *bool `json:"has_described_species_name,omitempty"`
	Counts []V2TaxonomyNodeCountByType `json:"counts,omitempty"`
	MinOrd *int32 `json:"min_ord,omitempty"`
	MaxOrd *int32 `json:"max_ord,omitempty"`
	Extinct *bool `json:"extinct,omitempty"`
	GenomicMoltype *string `json:"genomic_moltype,omitempty"`
	IndexedNames []string `json:"_indexed_names,omitempty"`
}

// NewV2TaxonomyNode instantiates a new V2TaxonomyNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TaxonomyNode() *V2TaxonomyNode {
	this := V2TaxonomyNode{}
	var rank V2reportsRankType = V2REPORTSRANKTYPE_NO_RANK
	this.Rank = &rank
	return &this
}

// NewV2TaxonomyNodeWithDefaults instantiates a new V2TaxonomyNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TaxonomyNodeWithDefaults() *V2TaxonomyNode {
	this := V2TaxonomyNode{}
	var rank V2reportsRankType = V2REPORTSRANKTYPE_NO_RANK
	this.Rank = &rank
	return &this
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetTaxId() int32 {
	if o == nil || o.TaxId == nil {
		var ret int32
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetTaxIdOk() (*int32, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given int32 and assigns it to the TaxId field.
func (o *V2TaxonomyNode) SetTaxId(v int32) {
	o.TaxId = &v
}

// GetOrganismName returns the OrganismName field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetOrganismName() string {
	if o == nil || o.OrganismName == nil {
		var ret string
		return ret
	}
	return *o.OrganismName
}

// GetOrganismNameOk returns a tuple with the OrganismName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetOrganismNameOk() (*string, bool) {
	if o == nil || o.OrganismName == nil {
		return nil, false
	}
	return o.OrganismName, true
}

// HasOrganismName returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasOrganismName() bool {
	if o != nil && o.OrganismName != nil {
		return true
	}

	return false
}

// SetOrganismName gets a reference to the given string and assigns it to the OrganismName field.
func (o *V2TaxonomyNode) SetOrganismName(v string) {
	o.OrganismName = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *V2TaxonomyNode) SetCommonName(v string) {
	o.CommonName = &v
}

// GetGenbankCommonName returns the GenbankCommonName field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetGenbankCommonName() string {
	if o == nil || o.GenbankCommonName == nil {
		var ret string
		return ret
	}
	return *o.GenbankCommonName
}

// GetGenbankCommonNameOk returns a tuple with the GenbankCommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetGenbankCommonNameOk() (*string, bool) {
	if o == nil || o.GenbankCommonName == nil {
		return nil, false
	}
	return o.GenbankCommonName, true
}

// HasGenbankCommonName returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasGenbankCommonName() bool {
	if o != nil && o.GenbankCommonName != nil {
		return true
	}

	return false
}

// SetGenbankCommonName gets a reference to the given string and assigns it to the GenbankCommonName field.
func (o *V2TaxonomyNode) SetGenbankCommonName(v string) {
	o.GenbankCommonName = &v
}

// GetAcronyms returns the Acronyms field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetAcronyms() []string {
	if o == nil || o.Acronyms == nil {
		var ret []string
		return ret
	}
	return o.Acronyms
}

// GetAcronymsOk returns a tuple with the Acronyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetAcronymsOk() ([]string, bool) {
	if o == nil || o.Acronyms == nil {
		return nil, false
	}
	return o.Acronyms, true
}

// HasAcronyms returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasAcronyms() bool {
	if o != nil && o.Acronyms != nil {
		return true
	}

	return false
}

// SetAcronyms gets a reference to the given []string and assigns it to the Acronyms field.
func (o *V2TaxonomyNode) SetAcronyms(v []string) {
	o.Acronyms = v
}

// GetGenbankAcronym returns the GenbankAcronym field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetGenbankAcronym() string {
	if o == nil || o.GenbankAcronym == nil {
		var ret string
		return ret
	}
	return *o.GenbankAcronym
}

// GetGenbankAcronymOk returns a tuple with the GenbankAcronym field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetGenbankAcronymOk() (*string, bool) {
	if o == nil || o.GenbankAcronym == nil {
		return nil, false
	}
	return o.GenbankAcronym, true
}

// HasGenbankAcronym returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasGenbankAcronym() bool {
	if o != nil && o.GenbankAcronym != nil {
		return true
	}

	return false
}

// SetGenbankAcronym gets a reference to the given string and assigns it to the GenbankAcronym field.
func (o *V2TaxonomyNode) SetGenbankAcronym(v string) {
	o.GenbankAcronym = &v
}

// GetBlastName returns the BlastName field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetBlastName() string {
	if o == nil || o.BlastName == nil {
		var ret string
		return ret
	}
	return *o.BlastName
}

// GetBlastNameOk returns a tuple with the BlastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetBlastNameOk() (*string, bool) {
	if o == nil || o.BlastName == nil {
		return nil, false
	}
	return o.BlastName, true
}

// HasBlastName returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasBlastName() bool {
	if o != nil && o.BlastName != nil {
		return true
	}

	return false
}

// SetBlastName gets a reference to the given string and assigns it to the BlastName field.
func (o *V2TaxonomyNode) SetBlastName(v string) {
	o.BlastName = &v
}

// GetLineage returns the Lineage field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetLineage() []int32 {
	if o == nil || o.Lineage == nil {
		var ret []int32
		return ret
	}
	return o.Lineage
}

// GetLineageOk returns a tuple with the Lineage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetLineageOk() ([]int32, bool) {
	if o == nil || o.Lineage == nil {
		return nil, false
	}
	return o.Lineage, true
}

// HasLineage returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasLineage() bool {
	if o != nil && o.Lineage != nil {
		return true
	}

	return false
}

// SetLineage gets a reference to the given []int32 and assigns it to the Lineage field.
func (o *V2TaxonomyNode) SetLineage(v []int32) {
	o.Lineage = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetChildren() []int32 {
	if o == nil || o.Children == nil {
		var ret []int32
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetChildrenOk() ([]int32, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []int32 and assigns it to the Children field.
func (o *V2TaxonomyNode) SetChildren(v []int32) {
	o.Children = v
}

// GetDescendentWithDescribedSpeciesNamesCount returns the DescendentWithDescribedSpeciesNamesCount field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetDescendentWithDescribedSpeciesNamesCount() int32 {
	if o == nil || o.DescendentWithDescribedSpeciesNamesCount == nil {
		var ret int32
		return ret
	}
	return *o.DescendentWithDescribedSpeciesNamesCount
}

// GetDescendentWithDescribedSpeciesNamesCountOk returns a tuple with the DescendentWithDescribedSpeciesNamesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetDescendentWithDescribedSpeciesNamesCountOk() (*int32, bool) {
	if o == nil || o.DescendentWithDescribedSpeciesNamesCount == nil {
		return nil, false
	}
	return o.DescendentWithDescribedSpeciesNamesCount, true
}

// HasDescendentWithDescribedSpeciesNamesCount returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasDescendentWithDescribedSpeciesNamesCount() bool {
	if o != nil && o.DescendentWithDescribedSpeciesNamesCount != nil {
		return true
	}

	return false
}

// SetDescendentWithDescribedSpeciesNamesCount gets a reference to the given int32 and assigns it to the DescendentWithDescribedSpeciesNamesCount field.
func (o *V2TaxonomyNode) SetDescendentWithDescribedSpeciesNamesCount(v int32) {
	o.DescendentWithDescribedSpeciesNamesCount = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetRank() V2reportsRankType {
	if o == nil || o.Rank == nil {
		var ret V2reportsRankType
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetRankOk() (*V2reportsRankType, bool) {
	if o == nil || o.Rank == nil {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasRank() bool {
	if o != nil && o.Rank != nil {
		return true
	}

	return false
}

// SetRank gets a reference to the given V2reportsRankType and assigns it to the Rank field.
func (o *V2TaxonomyNode) SetRank(v V2reportsRankType) {
	o.Rank = &v
}

// GetHasDescribedSpeciesName returns the HasDescribedSpeciesName field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetHasDescribedSpeciesName() bool {
	if o == nil || o.HasDescribedSpeciesName == nil {
		var ret bool
		return ret
	}
	return *o.HasDescribedSpeciesName
}

// GetHasDescribedSpeciesNameOk returns a tuple with the HasDescribedSpeciesName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetHasDescribedSpeciesNameOk() (*bool, bool) {
	if o == nil || o.HasDescribedSpeciesName == nil {
		return nil, false
	}
	return o.HasDescribedSpeciesName, true
}

// HasHasDescribedSpeciesName returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasHasDescribedSpeciesName() bool {
	if o != nil && o.HasDescribedSpeciesName != nil {
		return true
	}

	return false
}

// SetHasDescribedSpeciesName gets a reference to the given bool and assigns it to the HasDescribedSpeciesName field.
func (o *V2TaxonomyNode) SetHasDescribedSpeciesName(v bool) {
	o.HasDescribedSpeciesName = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetCounts() []V2TaxonomyNodeCountByType {
	if o == nil || o.Counts == nil {
		var ret []V2TaxonomyNodeCountByType
		return ret
	}
	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetCountsOk() ([]V2TaxonomyNodeCountByType, bool) {
	if o == nil || o.Counts == nil {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasCounts() bool {
	if o != nil && o.Counts != nil {
		return true
	}

	return false
}

// SetCounts gets a reference to the given []V2TaxonomyNodeCountByType and assigns it to the Counts field.
func (o *V2TaxonomyNode) SetCounts(v []V2TaxonomyNodeCountByType) {
	o.Counts = v
}

// GetMinOrd returns the MinOrd field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetMinOrd() int32 {
	if o == nil || o.MinOrd == nil {
		var ret int32
		return ret
	}
	return *o.MinOrd
}

// GetMinOrdOk returns a tuple with the MinOrd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetMinOrdOk() (*int32, bool) {
	if o == nil || o.MinOrd == nil {
		return nil, false
	}
	return o.MinOrd, true
}

// HasMinOrd returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasMinOrd() bool {
	if o != nil && o.MinOrd != nil {
		return true
	}

	return false
}

// SetMinOrd gets a reference to the given int32 and assigns it to the MinOrd field.
func (o *V2TaxonomyNode) SetMinOrd(v int32) {
	o.MinOrd = &v
}

// GetMaxOrd returns the MaxOrd field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetMaxOrd() int32 {
	if o == nil || o.MaxOrd == nil {
		var ret int32
		return ret
	}
	return *o.MaxOrd
}

// GetMaxOrdOk returns a tuple with the MaxOrd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetMaxOrdOk() (*int32, bool) {
	if o == nil || o.MaxOrd == nil {
		return nil, false
	}
	return o.MaxOrd, true
}

// HasMaxOrd returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasMaxOrd() bool {
	if o != nil && o.MaxOrd != nil {
		return true
	}

	return false
}

// SetMaxOrd gets a reference to the given int32 and assigns it to the MaxOrd field.
func (o *V2TaxonomyNode) SetMaxOrd(v int32) {
	o.MaxOrd = &v
}

// GetExtinct returns the Extinct field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetExtinct() bool {
	if o == nil || o.Extinct == nil {
		var ret bool
		return ret
	}
	return *o.Extinct
}

// GetExtinctOk returns a tuple with the Extinct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetExtinctOk() (*bool, bool) {
	if o == nil || o.Extinct == nil {
		return nil, false
	}
	return o.Extinct, true
}

// HasExtinct returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasExtinct() bool {
	if o != nil && o.Extinct != nil {
		return true
	}

	return false
}

// SetExtinct gets a reference to the given bool and assigns it to the Extinct field.
func (o *V2TaxonomyNode) SetExtinct(v bool) {
	o.Extinct = &v
}

// GetGenomicMoltype returns the GenomicMoltype field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetGenomicMoltype() string {
	if o == nil || o.GenomicMoltype == nil {
		var ret string
		return ret
	}
	return *o.GenomicMoltype
}

// GetGenomicMoltypeOk returns a tuple with the GenomicMoltype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetGenomicMoltypeOk() (*string, bool) {
	if o == nil || o.GenomicMoltype == nil {
		return nil, false
	}
	return o.GenomicMoltype, true
}

// HasGenomicMoltype returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasGenomicMoltype() bool {
	if o != nil && o.GenomicMoltype != nil {
		return true
	}

	return false
}

// SetGenomicMoltype gets a reference to the given string and assigns it to the GenomicMoltype field.
func (o *V2TaxonomyNode) SetGenomicMoltype(v string) {
	o.GenomicMoltype = &v
}

// GetIndexedNames returns the IndexedNames field value if set, zero value otherwise.
func (o *V2TaxonomyNode) GetIndexedNames() []string {
	if o == nil || o.IndexedNames == nil {
		var ret []string
		return ret
	}
	return o.IndexedNames
}

// GetIndexedNamesOk returns a tuple with the IndexedNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyNode) GetIndexedNamesOk() ([]string, bool) {
	if o == nil || o.IndexedNames == nil {
		return nil, false
	}
	return o.IndexedNames, true
}

// HasIndexedNames returns a boolean if a field has been set.
func (o *V2TaxonomyNode) HasIndexedNames() bool {
	if o != nil && o.IndexedNames != nil {
		return true
	}

	return false
}

// SetIndexedNames gets a reference to the given []string and assigns it to the IndexedNames field.
func (o *V2TaxonomyNode) SetIndexedNames(v []string) {
	o.IndexedNames = v
}

func (o V2TaxonomyNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxId != nil {
		toSerialize["tax_id"] = o.TaxId
	}
	if o.OrganismName != nil {
		toSerialize["organism_name"] = o.OrganismName
	}
	if o.CommonName != nil {
		toSerialize["common_name"] = o.CommonName
	}
	if o.GenbankCommonName != nil {
		toSerialize["genbank_common_name"] = o.GenbankCommonName
	}
	if o.Acronyms != nil {
		toSerialize["acronyms"] = o.Acronyms
	}
	if o.GenbankAcronym != nil {
		toSerialize["genbank_acronym"] = o.GenbankAcronym
	}
	if o.BlastName != nil {
		toSerialize["blast_name"] = o.BlastName
	}
	if o.Lineage != nil {
		toSerialize["lineage"] = o.Lineage
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.DescendentWithDescribedSpeciesNamesCount != nil {
		toSerialize["descendent_with_described_species_names_count"] = o.DescendentWithDescribedSpeciesNamesCount
	}
	if o.Rank != nil {
		toSerialize["rank"] = o.Rank
	}
	if o.HasDescribedSpeciesName != nil {
		toSerialize["has_described_species_name"] = o.HasDescribedSpeciesName
	}
	if o.Counts != nil {
		toSerialize["counts"] = o.Counts
	}
	if o.MinOrd != nil {
		toSerialize["min_ord"] = o.MinOrd
	}
	if o.MaxOrd != nil {
		toSerialize["max_ord"] = o.MaxOrd
	}
	if o.Extinct != nil {
		toSerialize["extinct"] = o.Extinct
	}
	if o.GenomicMoltype != nil {
		toSerialize["genomic_moltype"] = o.GenomicMoltype
	}
	if o.IndexedNames != nil {
		toSerialize["_indexed_names"] = o.IndexedNames
	}
	return json.Marshal(toSerialize)
}

type NullableV2TaxonomyNode struct {
	value *V2TaxonomyNode
	isSet bool
}

func (v NullableV2TaxonomyNode) Get() *V2TaxonomyNode {
	return v.value
}

func (v *NullableV2TaxonomyNode) Set(val *V2TaxonomyNode) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TaxonomyNode) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TaxonomyNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TaxonomyNode(val *V2TaxonomyNode) *NullableV2TaxonomyNode {
	return &NullableV2TaxonomyNode{value: val, isSet: true}
}

func (v NullableV2TaxonomyNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TaxonomyNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


