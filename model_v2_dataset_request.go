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

// V2DatasetRequest struct for V2DatasetRequest
type V2DatasetRequest struct {
	GenomeV2 *V2AssemblyDatasetRequest `json:"genome_v2,omitempty"`
	GeneV2 *V2GeneDatasetRequest `json:"gene_v2,omitempty"`
	VirusV2 *V2VirusDatasetRequest `json:"virus_v2,omitempty"`
}

// NewV2DatasetRequest instantiates a new V2DatasetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2DatasetRequest() *V2DatasetRequest {
	this := V2DatasetRequest{}
	return &this
}

// NewV2DatasetRequestWithDefaults instantiates a new V2DatasetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2DatasetRequestWithDefaults() *V2DatasetRequest {
	this := V2DatasetRequest{}
	return &this
}

// GetGenomeV2 returns the GenomeV2 field value if set, zero value otherwise.
func (o *V2DatasetRequest) GetGenomeV2() V2AssemblyDatasetRequest {
	if o == nil || o.GenomeV2 == nil {
		var ret V2AssemblyDatasetRequest
		return ret
	}
	return *o.GenomeV2
}

// GetGenomeV2Ok returns a tuple with the GenomeV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2DatasetRequest) GetGenomeV2Ok() (*V2AssemblyDatasetRequest, bool) {
	if o == nil || o.GenomeV2 == nil {
		return nil, false
	}
	return o.GenomeV2, true
}

// HasGenomeV2 returns a boolean if a field has been set.
func (o *V2DatasetRequest) HasGenomeV2() bool {
	if o != nil && o.GenomeV2 != nil {
		return true
	}

	return false
}

// SetGenomeV2 gets a reference to the given V2AssemblyDatasetRequest and assigns it to the GenomeV2 field.
func (o *V2DatasetRequest) SetGenomeV2(v V2AssemblyDatasetRequest) {
	o.GenomeV2 = &v
}

// GetGeneV2 returns the GeneV2 field value if set, zero value otherwise.
func (o *V2DatasetRequest) GetGeneV2() V2GeneDatasetRequest {
	if o == nil || o.GeneV2 == nil {
		var ret V2GeneDatasetRequest
		return ret
	}
	return *o.GeneV2
}

// GetGeneV2Ok returns a tuple with the GeneV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2DatasetRequest) GetGeneV2Ok() (*V2GeneDatasetRequest, bool) {
	if o == nil || o.GeneV2 == nil {
		return nil, false
	}
	return o.GeneV2, true
}

// HasGeneV2 returns a boolean if a field has been set.
func (o *V2DatasetRequest) HasGeneV2() bool {
	if o != nil && o.GeneV2 != nil {
		return true
	}

	return false
}

// SetGeneV2 gets a reference to the given V2GeneDatasetRequest and assigns it to the GeneV2 field.
func (o *V2DatasetRequest) SetGeneV2(v V2GeneDatasetRequest) {
	o.GeneV2 = &v
}

// GetVirusV2 returns the VirusV2 field value if set, zero value otherwise.
func (o *V2DatasetRequest) GetVirusV2() V2VirusDatasetRequest {
	if o == nil || o.VirusV2 == nil {
		var ret V2VirusDatasetRequest
		return ret
	}
	return *o.VirusV2
}

// GetVirusV2Ok returns a tuple with the VirusV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2DatasetRequest) GetVirusV2Ok() (*V2VirusDatasetRequest, bool) {
	if o == nil || o.VirusV2 == nil {
		return nil, false
	}
	return o.VirusV2, true
}

// HasVirusV2 returns a boolean if a field has been set.
func (o *V2DatasetRequest) HasVirusV2() bool {
	if o != nil && o.VirusV2 != nil {
		return true
	}

	return false
}

// SetVirusV2 gets a reference to the given V2VirusDatasetRequest and assigns it to the VirusV2 field.
func (o *V2DatasetRequest) SetVirusV2(v V2VirusDatasetRequest) {
	o.VirusV2 = &v
}

func (o V2DatasetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GenomeV2 != nil {
		toSerialize["genome_v2"] = o.GenomeV2
	}
	if o.GeneV2 != nil {
		toSerialize["gene_v2"] = o.GeneV2
	}
	if o.VirusV2 != nil {
		toSerialize["virus_v2"] = o.VirusV2
	}
	return json.Marshal(toSerialize)
}

type NullableV2DatasetRequest struct {
	value *V2DatasetRequest
	isSet bool
}

func (v NullableV2DatasetRequest) Get() *V2DatasetRequest {
	return v.value
}

func (v *NullableV2DatasetRequest) Set(val *V2DatasetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2DatasetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2DatasetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2DatasetRequest(val *V2DatasetRequest) *NullableV2DatasetRequest {
	return &NullableV2DatasetRequest{value: val, isSet: true}
}

func (v NullableV2DatasetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2DatasetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


