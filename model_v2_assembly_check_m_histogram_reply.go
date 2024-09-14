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

// V2AssemblyCheckMHistogramReply struct for V2AssemblyCheckMHistogramReply
type V2AssemblyCheckMHistogramReply struct {
	SpeciesTaxid *int32 `json:"species_taxid,omitempty"`
	HistogramIntervals []V2AssemblyCheckMHistogramReplyHistogramInterval `json:"histogram_intervals,omitempty"`
}

// NewV2AssemblyCheckMHistogramReply instantiates a new V2AssemblyCheckMHistogramReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssemblyCheckMHistogramReply() *V2AssemblyCheckMHistogramReply {
	this := V2AssemblyCheckMHistogramReply{}
	return &this
}

// NewV2AssemblyCheckMHistogramReplyWithDefaults instantiates a new V2AssemblyCheckMHistogramReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssemblyCheckMHistogramReplyWithDefaults() *V2AssemblyCheckMHistogramReply {
	this := V2AssemblyCheckMHistogramReply{}
	return &this
}

// GetSpeciesTaxid returns the SpeciesTaxid field value if set, zero value otherwise.
func (o *V2AssemblyCheckMHistogramReply) GetSpeciesTaxid() int32 {
	if o == nil || o.SpeciesTaxid == nil {
		var ret int32
		return ret
	}
	return *o.SpeciesTaxid
}

// GetSpeciesTaxidOk returns a tuple with the SpeciesTaxid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssemblyCheckMHistogramReply) GetSpeciesTaxidOk() (*int32, bool) {
	if o == nil || o.SpeciesTaxid == nil {
		return nil, false
	}
	return o.SpeciesTaxid, true
}

// HasSpeciesTaxid returns a boolean if a field has been set.
func (o *V2AssemblyCheckMHistogramReply) HasSpeciesTaxid() bool {
	if o != nil && o.SpeciesTaxid != nil {
		return true
	}

	return false
}

// SetSpeciesTaxid gets a reference to the given int32 and assigns it to the SpeciesTaxid field.
func (o *V2AssemblyCheckMHistogramReply) SetSpeciesTaxid(v int32) {
	o.SpeciesTaxid = &v
}

// GetHistogramIntervals returns the HistogramIntervals field value if set, zero value otherwise.
func (o *V2AssemblyCheckMHistogramReply) GetHistogramIntervals() []V2AssemblyCheckMHistogramReplyHistogramInterval {
	if o == nil || o.HistogramIntervals == nil {
		var ret []V2AssemblyCheckMHistogramReplyHistogramInterval
		return ret
	}
	return o.HistogramIntervals
}

// GetHistogramIntervalsOk returns a tuple with the HistogramIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssemblyCheckMHistogramReply) GetHistogramIntervalsOk() ([]V2AssemblyCheckMHistogramReplyHistogramInterval, bool) {
	if o == nil || o.HistogramIntervals == nil {
		return nil, false
	}
	return o.HistogramIntervals, true
}

// HasHistogramIntervals returns a boolean if a field has been set.
func (o *V2AssemblyCheckMHistogramReply) HasHistogramIntervals() bool {
	if o != nil && o.HistogramIntervals != nil {
		return true
	}

	return false
}

// SetHistogramIntervals gets a reference to the given []V2AssemblyCheckMHistogramReplyHistogramInterval and assigns it to the HistogramIntervals field.
func (o *V2AssemblyCheckMHistogramReply) SetHistogramIntervals(v []V2AssemblyCheckMHistogramReplyHistogramInterval) {
	o.HistogramIntervals = v
}

func (o V2AssemblyCheckMHistogramReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SpeciesTaxid != nil {
		toSerialize["species_taxid"] = o.SpeciesTaxid
	}
	if o.HistogramIntervals != nil {
		toSerialize["histogram_intervals"] = o.HistogramIntervals
	}
	return json.Marshal(toSerialize)
}

type NullableV2AssemblyCheckMHistogramReply struct {
	value *V2AssemblyCheckMHistogramReply
	isSet bool
}

func (v NullableV2AssemblyCheckMHistogramReply) Get() *V2AssemblyCheckMHistogramReply {
	return v.value
}

func (v *NullableV2AssemblyCheckMHistogramReply) Set(val *V2AssemblyCheckMHistogramReply) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssemblyCheckMHistogramReply) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssemblyCheckMHistogramReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssemblyCheckMHistogramReply(val *V2AssemblyCheckMHistogramReply) *NullableV2AssemblyCheckMHistogramReply {
	return &NullableV2AssemblyCheckMHistogramReply{value: val, isSet: true}
}

func (v NullableV2AssemblyCheckMHistogramReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssemblyCheckMHistogramReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


