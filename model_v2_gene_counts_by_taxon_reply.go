/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 2 API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated zip archive](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v2alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// V2GeneCountsByTaxonReply struct for V2GeneCountsByTaxonReply
type V2GeneCountsByTaxonReply struct {
	Report []V2GeneCountsByTaxonReplyGeneTypeAndCount `json:"report,omitempty"`
}

// NewV2GeneCountsByTaxonReply instantiates a new V2GeneCountsByTaxonReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2GeneCountsByTaxonReply() *V2GeneCountsByTaxonReply {
	this := V2GeneCountsByTaxonReply{}
	return &this
}

// NewV2GeneCountsByTaxonReplyWithDefaults instantiates a new V2GeneCountsByTaxonReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2GeneCountsByTaxonReplyWithDefaults() *V2GeneCountsByTaxonReply {
	this := V2GeneCountsByTaxonReply{}
	return &this
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *V2GeneCountsByTaxonReply) GetReport() []V2GeneCountsByTaxonReplyGeneTypeAndCount {
	if o == nil || o.Report == nil {
		var ret []V2GeneCountsByTaxonReplyGeneTypeAndCount
		return ret
	}
	return o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2GeneCountsByTaxonReply) GetReportOk() ([]V2GeneCountsByTaxonReplyGeneTypeAndCount, bool) {
	if o == nil || o.Report == nil {
		return nil, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *V2GeneCountsByTaxonReply) HasReport() bool {
	if o != nil && o.Report != nil {
		return true
	}

	return false
}

// SetReport gets a reference to the given []V2GeneCountsByTaxonReplyGeneTypeAndCount and assigns it to the Report field.
func (o *V2GeneCountsByTaxonReply) SetReport(v []V2GeneCountsByTaxonReplyGeneTypeAndCount) {
	o.Report = v
}

func (o V2GeneCountsByTaxonReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Report != nil {
		toSerialize["report"] = o.Report
	}
	return json.Marshal(toSerialize)
}

type NullableV2GeneCountsByTaxonReply struct {
	value *V2GeneCountsByTaxonReply
	isSet bool
}

func (v NullableV2GeneCountsByTaxonReply) Get() *V2GeneCountsByTaxonReply {
	return v.value
}

func (v *NullableV2GeneCountsByTaxonReply) Set(val *V2GeneCountsByTaxonReply) {
	v.value = val
	v.isSet = true
}

func (v NullableV2GeneCountsByTaxonReply) IsSet() bool {
	return v.isSet
}

func (v *NullableV2GeneCountsByTaxonReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2GeneCountsByTaxonReply(val *V2GeneCountsByTaxonReply) *NullableV2GeneCountsByTaxonReply {
	return &NullableV2GeneCountsByTaxonReply{value: val, isSet: true}
}

func (v NullableV2GeneCountsByTaxonReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2GeneCountsByTaxonReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


