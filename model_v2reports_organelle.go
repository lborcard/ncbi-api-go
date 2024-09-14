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

// V2reportsOrganelle struct for V2reportsOrganelle
type V2reportsOrganelle struct {
	Description *V2reportsOrganelleType `json:"description,omitempty"`
	Genbank *V2reportsSequenceInformation `json:"genbank,omitempty"`
	Refseq *V2reportsSequenceInformation `json:"refseq,omitempty"`
	Organism *V2reportsOrganism `json:"organism,omitempty"`
	Bioprojects []V2reportsBioProject `json:"bioprojects,omitempty"`
	Biosample *V2reportsOrganelleBiosample `json:"biosample,omitempty"`
	GeneCounts *V2reportsOrganelleGeneCounts `json:"gene_counts,omitempty"`
	Length *int32 `json:"length,omitempty"`
	Topology *V2reportsOrganelleTopology `json:"topology,omitempty"`
	GeneCount *int32 `json:"gene_count,omitempty"`
}

// NewV2reportsOrganelle instantiates a new V2reportsOrganelle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsOrganelle() *V2reportsOrganelle {
	this := V2reportsOrganelle{}
	var description V2reportsOrganelleType = V2REPORTSORGANELLETYPE_ORGANELLE_TYPE_UNKNOWN
	this.Description = &description
	var topology V2reportsOrganelleTopology = V2REPORTSORGANELLETOPOLOGY_TOPOLOGY_UNKNOWN
	this.Topology = &topology
	return &this
}

// NewV2reportsOrganelleWithDefaults instantiates a new V2reportsOrganelle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsOrganelleWithDefaults() *V2reportsOrganelle {
	this := V2reportsOrganelle{}
	var description V2reportsOrganelleType = V2REPORTSORGANELLETYPE_ORGANELLE_TYPE_UNKNOWN
	this.Description = &description
	var topology V2reportsOrganelleTopology = V2REPORTSORGANELLETOPOLOGY_TOPOLOGY_UNKNOWN
	this.Topology = &topology
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V2reportsOrganelle) GetDescription() V2reportsOrganelleType {
	if o == nil || o.Description == nil {
		var ret V2reportsOrganelleType
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsOrganelle) GetDescriptionOk() (*V2reportsOrganelleType, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V2reportsOrganelle) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given V2reportsOrganelleType and assigns it to the Description field.
func (o *V2reportsOrganelle) SetDescription(v V2reportsOrganelleType) {
	o.Description = &v
}

// GetGenbank returns the Genbank field value if set, zero value otherwise.
func (o *V2reportsOrganelle) GetGenbank() V2reportsSequenceInformation {
	if o == nil || o.Genbank == nil {
		var ret V2reportsSequenceInformation
		return ret
	}
	return *o.Genbank
}

// GetGenbankOk returns a tuple with the Genbank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsOrganelle) GetGenbankOk() (*V2reportsSequenceInformation, bool) {
	if o == nil || o.Genbank == nil {
		return nil, false
	}
	return o.Genbank, true
}

// HasGenbank returns a boolean if a field has been set.
func (o *V2reportsOrganelle) HasGenbank() bool {
	if o != nil && o.Genbank != nil {
		return true
	}

	return false
}

// SetGenbank gets a reference to the given V2reportsSequenceInformation and assigns it to the Genbank field.
func (o *V2reportsOrganelle) SetGenbank(v V2reportsSequenceInformation) {
	o.Genbank = &v
}

// GetRefseq returns the Refseq field value if set, zero value otherwise.
func (o *V2reportsOrganelle) GetRefseq() V2reportsSequenceInformation {
	if o == nil || o.Refseq == nil {
		var ret V2reportsSequenceInformation
		return ret
	}
	return *o.Refseq
}

// GetRefseqOk returns a tuple with the Refseq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsOrganelle) GetRefseqOk() (*V2reportsSequenceInformation, bool) {
	if o == nil || o.Refseq == nil {
		return nil, false
	}
	return o.Refseq, true
}

// HasRefseq returns a boolean if a field has been set.
func (o *V2reportsOrganelle) HasRefseq() bool {
	if o != nil && o.Refseq != nil {
		return true
	}

	return false
}

// SetRefseq gets a reference to the given V2reportsSequenceInformation and assigns it to the Refseq field.
func (o *V2reportsOrganelle) SetRefseq(v V2reportsSequenceInformation) {
	o.Refseq = &v
}

// GetOrganism returns the Organism field value if set, zero value otherwise.
func (o *V2reportsOrganelle) GetOrganism() V2reportsOrganism {
	if o == nil || o.Organism == nil {
		var ret V2reportsOrganism
		return ret
	}
	return *o.Organism
}

// GetOrganismOk returns a tuple with the Organism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsOrganelle) GetOrganismOk() (*V2reportsOrganism, bool) {
	if o == nil || o.Organism == nil {
		return nil, false
	}
	return o.Organism, true
}

// HasOrganism returns a boolean if a field has been set.
func (o *V2reportsOrganelle) HasOrganism() bool {
	if o != nil && o.Organism != nil {
		return true
	}

	return false
}

// SetOrganism gets a reference to the given V2reportsOrganism and assigns it to the Organism field.
func (o *V2reportsOrganelle) SetOrganism(v V2reportsOrganism) {
	o.Organism = &v
}

// GetBioprojects returns the Bioprojects field value if set, zero value otherwise.
func (o *V2reportsOrganelle) GetBioprojects() []V2reportsBioProject {
	if o == nil || o.Bioprojects == nil {
		var ret []V2reportsBioProject
		return ret
	}
	return o.Bioprojects
}

// GetBioprojectsOk returns a tuple with the Bioprojects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsOrganelle) GetBioprojectsOk() ([]V2reportsBioProject, bool) {
	if o == nil || o.Bioprojects == nil {
		return nil, false
	}
	return o.Bioprojects, true
}

// HasBioprojects returns a boolean if a field has been set.
func (o *V2reportsOrganelle) HasBioprojects() bool {
	if o != nil && o.Bioprojects != nil {
		return true
	}

	return false
}

// SetBioprojects gets a reference to the given []V2reportsBioProject and assigns it to the Bioprojects field.
func (o *V2reportsOrganelle) SetBioprojects(v []V2reportsBioProject) {
	o.Bioprojects = v
}

// GetBiosample returns the Biosample field value if set, zero value otherwise.
func (o *V2reportsOrganelle) GetBiosample() V2reportsOrganelleBiosample {
	if o == nil || o.Biosample == nil {
		var ret V2reportsOrganelleBiosample
		return ret
	}
	return *o.Biosample
}

// GetBiosampleOk returns a tuple with the Biosample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsOrganelle) GetBiosampleOk() (*V2reportsOrganelleBiosample, bool) {
	if o == nil || o.Biosample == nil {
		return nil, false
	}
	return o.Biosample, true
}

// HasBiosample returns a boolean if a field has been set.
func (o *V2reportsOrganelle) HasBiosample() bool {
	if o != nil && o.Biosample != nil {
		return true
	}

	return false
}

// SetBiosample gets a reference to the given V2reportsOrganelleBiosample and assigns it to the Biosample field.
func (o *V2reportsOrganelle) SetBiosample(v V2reportsOrganelleBiosample) {
	o.Biosample = &v
}

// GetGeneCounts returns the GeneCounts field value if set, zero value otherwise.
func (o *V2reportsOrganelle) GetGeneCounts() V2reportsOrganelleGeneCounts {
	if o == nil || o.GeneCounts == nil {
		var ret V2reportsOrganelleGeneCounts
		return ret
	}
	return *o.GeneCounts
}

// GetGeneCountsOk returns a tuple with the GeneCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsOrganelle) GetGeneCountsOk() (*V2reportsOrganelleGeneCounts, bool) {
	if o == nil || o.GeneCounts == nil {
		return nil, false
	}
	return o.GeneCounts, true
}

// HasGeneCounts returns a boolean if a field has been set.
func (o *V2reportsOrganelle) HasGeneCounts() bool {
	if o != nil && o.GeneCounts != nil {
		return true
	}

	return false
}

// SetGeneCounts gets a reference to the given V2reportsOrganelleGeneCounts and assigns it to the GeneCounts field.
func (o *V2reportsOrganelle) SetGeneCounts(v V2reportsOrganelleGeneCounts) {
	o.GeneCounts = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *V2reportsOrganelle) GetLength() int32 {
	if o == nil || o.Length == nil {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsOrganelle) GetLengthOk() (*int32, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *V2reportsOrganelle) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *V2reportsOrganelle) SetLength(v int32) {
	o.Length = &v
}

// GetTopology returns the Topology field value if set, zero value otherwise.
func (o *V2reportsOrganelle) GetTopology() V2reportsOrganelleTopology {
	if o == nil || o.Topology == nil {
		var ret V2reportsOrganelleTopology
		return ret
	}
	return *o.Topology
}

// GetTopologyOk returns a tuple with the Topology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsOrganelle) GetTopologyOk() (*V2reportsOrganelleTopology, bool) {
	if o == nil || o.Topology == nil {
		return nil, false
	}
	return o.Topology, true
}

// HasTopology returns a boolean if a field has been set.
func (o *V2reportsOrganelle) HasTopology() bool {
	if o != nil && o.Topology != nil {
		return true
	}

	return false
}

// SetTopology gets a reference to the given V2reportsOrganelleTopology and assigns it to the Topology field.
func (o *V2reportsOrganelle) SetTopology(v V2reportsOrganelleTopology) {
	o.Topology = &v
}

// GetGeneCount returns the GeneCount field value if set, zero value otherwise.
func (o *V2reportsOrganelle) GetGeneCount() int32 {
	if o == nil || o.GeneCount == nil {
		var ret int32
		return ret
	}
	return *o.GeneCount
}

// GetGeneCountOk returns a tuple with the GeneCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsOrganelle) GetGeneCountOk() (*int32, bool) {
	if o == nil || o.GeneCount == nil {
		return nil, false
	}
	return o.GeneCount, true
}

// HasGeneCount returns a boolean if a field has been set.
func (o *V2reportsOrganelle) HasGeneCount() bool {
	if o != nil && o.GeneCount != nil {
		return true
	}

	return false
}

// SetGeneCount gets a reference to the given int32 and assigns it to the GeneCount field.
func (o *V2reportsOrganelle) SetGeneCount(v int32) {
	o.GeneCount = &v
}

func (o V2reportsOrganelle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Genbank != nil {
		toSerialize["genbank"] = o.Genbank
	}
	if o.Refseq != nil {
		toSerialize["refseq"] = o.Refseq
	}
	if o.Organism != nil {
		toSerialize["organism"] = o.Organism
	}
	if o.Bioprojects != nil {
		toSerialize["bioprojects"] = o.Bioprojects
	}
	if o.Biosample != nil {
		toSerialize["biosample"] = o.Biosample
	}
	if o.GeneCounts != nil {
		toSerialize["gene_counts"] = o.GeneCounts
	}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.Topology != nil {
		toSerialize["topology"] = o.Topology
	}
	if o.GeneCount != nil {
		toSerialize["gene_count"] = o.GeneCount
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsOrganelle struct {
	value *V2reportsOrganelle
	isSet bool
}

func (v NullableV2reportsOrganelle) Get() *V2reportsOrganelle {
	return v.value
}

func (v *NullableV2reportsOrganelle) Set(val *V2reportsOrganelle) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsOrganelle) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsOrganelle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsOrganelle(val *V2reportsOrganelle) *NullableV2reportsOrganelle {
	return &NullableV2reportsOrganelle{value: val, isSet: true}
}

func (v NullableV2reportsOrganelle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsOrganelle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


