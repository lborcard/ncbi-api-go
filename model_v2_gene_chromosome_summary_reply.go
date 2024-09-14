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

// V2GeneChromosomeSummaryReply struct for V2GeneChromosomeSummaryReply
type V2GeneChromosomeSummaryReply struct {
	GeneChromosomeSummaries []V2GeneChromosomeSummaryReplyGeneChromosomeSummary `json:"gene_chromosome_summaries,omitempty"`
}

// NewV2GeneChromosomeSummaryReply instantiates a new V2GeneChromosomeSummaryReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2GeneChromosomeSummaryReply() *V2GeneChromosomeSummaryReply {
	this := V2GeneChromosomeSummaryReply{}
	return &this
}

// NewV2GeneChromosomeSummaryReplyWithDefaults instantiates a new V2GeneChromosomeSummaryReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2GeneChromosomeSummaryReplyWithDefaults() *V2GeneChromosomeSummaryReply {
	this := V2GeneChromosomeSummaryReply{}
	return &this
}

// GetGeneChromosomeSummaries returns the GeneChromosomeSummaries field value if set, zero value otherwise.
func (o *V2GeneChromosomeSummaryReply) GetGeneChromosomeSummaries() []V2GeneChromosomeSummaryReplyGeneChromosomeSummary {
	if o == nil || o.GeneChromosomeSummaries == nil {
		var ret []V2GeneChromosomeSummaryReplyGeneChromosomeSummary
		return ret
	}
	return o.GeneChromosomeSummaries
}

// GetGeneChromosomeSummariesOk returns a tuple with the GeneChromosomeSummaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2GeneChromosomeSummaryReply) GetGeneChromosomeSummariesOk() ([]V2GeneChromosomeSummaryReplyGeneChromosomeSummary, bool) {
	if o == nil || o.GeneChromosomeSummaries == nil {
		return nil, false
	}
	return o.GeneChromosomeSummaries, true
}

// HasGeneChromosomeSummaries returns a boolean if a field has been set.
func (o *V2GeneChromosomeSummaryReply) HasGeneChromosomeSummaries() bool {
	if o != nil && o.GeneChromosomeSummaries != nil {
		return true
	}

	return false
}

// SetGeneChromosomeSummaries gets a reference to the given []V2GeneChromosomeSummaryReplyGeneChromosomeSummary and assigns it to the GeneChromosomeSummaries field.
func (o *V2GeneChromosomeSummaryReply) SetGeneChromosomeSummaries(v []V2GeneChromosomeSummaryReplyGeneChromosomeSummary) {
	o.GeneChromosomeSummaries = v
}

func (o V2GeneChromosomeSummaryReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GeneChromosomeSummaries != nil {
		toSerialize["gene_chromosome_summaries"] = o.GeneChromosomeSummaries
	}
	return json.Marshal(toSerialize)
}

type NullableV2GeneChromosomeSummaryReply struct {
	value *V2GeneChromosomeSummaryReply
	isSet bool
}

func (v NullableV2GeneChromosomeSummaryReply) Get() *V2GeneChromosomeSummaryReply {
	return v.value
}

func (v *NullableV2GeneChromosomeSummaryReply) Set(val *V2GeneChromosomeSummaryReply) {
	v.value = val
	v.isSet = true
}

func (v NullableV2GeneChromosomeSummaryReply) IsSet() bool {
	return v.isSet
}

func (v *NullableV2GeneChromosomeSummaryReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2GeneChromosomeSummaryReply(val *V2GeneChromosomeSummaryReply) *NullableV2GeneChromosomeSummaryReply {
	return &NullableV2GeneChromosomeSummaryReply{value: val, isSet: true}
}

func (v NullableV2GeneChromosomeSummaryReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2GeneChromosomeSummaryReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


