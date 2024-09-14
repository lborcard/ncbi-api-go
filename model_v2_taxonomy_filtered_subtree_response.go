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

// V2TaxonomyFilteredSubtreeResponse struct for V2TaxonomyFilteredSubtreeResponse
type V2TaxonomyFilteredSubtreeResponse struct {
	RootNodes []int32 `json:"root_nodes,omitempty"`
	Edges *V2TaxonomyFilteredSubtreeResponseEdgesEntry `json:"edges,omitempty"`
	Warnings []V2reportsWarning `json:"warnings,omitempty"`
	Errors []V2reportsError `json:"errors,omitempty"`
}

// NewV2TaxonomyFilteredSubtreeResponse instantiates a new V2TaxonomyFilteredSubtreeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TaxonomyFilteredSubtreeResponse() *V2TaxonomyFilteredSubtreeResponse {
	this := V2TaxonomyFilteredSubtreeResponse{}
	return &this
}

// NewV2TaxonomyFilteredSubtreeResponseWithDefaults instantiates a new V2TaxonomyFilteredSubtreeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TaxonomyFilteredSubtreeResponseWithDefaults() *V2TaxonomyFilteredSubtreeResponse {
	this := V2TaxonomyFilteredSubtreeResponse{}
	return &this
}

// GetRootNodes returns the RootNodes field value if set, zero value otherwise.
func (o *V2TaxonomyFilteredSubtreeResponse) GetRootNodes() []int32 {
	if o == nil || o.RootNodes == nil {
		var ret []int32
		return ret
	}
	return o.RootNodes
}

// GetRootNodesOk returns a tuple with the RootNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyFilteredSubtreeResponse) GetRootNodesOk() ([]int32, bool) {
	if o == nil || o.RootNodes == nil {
		return nil, false
	}
	return o.RootNodes, true
}

// HasRootNodes returns a boolean if a field has been set.
func (o *V2TaxonomyFilteredSubtreeResponse) HasRootNodes() bool {
	if o != nil && o.RootNodes != nil {
		return true
	}

	return false
}

// SetRootNodes gets a reference to the given []int32 and assigns it to the RootNodes field.
func (o *V2TaxonomyFilteredSubtreeResponse) SetRootNodes(v []int32) {
	o.RootNodes = v
}

// GetEdges returns the Edges field value if set, zero value otherwise.
func (o *V2TaxonomyFilteredSubtreeResponse) GetEdges() V2TaxonomyFilteredSubtreeResponseEdgesEntry {
	if o == nil || o.Edges == nil {
		var ret V2TaxonomyFilteredSubtreeResponseEdgesEntry
		return ret
	}
	return *o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyFilteredSubtreeResponse) GetEdgesOk() (*V2TaxonomyFilteredSubtreeResponseEdgesEntry, bool) {
	if o == nil || o.Edges == nil {
		return nil, false
	}
	return o.Edges, true
}

// HasEdges returns a boolean if a field has been set.
func (o *V2TaxonomyFilteredSubtreeResponse) HasEdges() bool {
	if o != nil && o.Edges != nil {
		return true
	}

	return false
}

// SetEdges gets a reference to the given V2TaxonomyFilteredSubtreeResponseEdgesEntry and assigns it to the Edges field.
func (o *V2TaxonomyFilteredSubtreeResponse) SetEdges(v V2TaxonomyFilteredSubtreeResponseEdgesEntry) {
	o.Edges = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V2TaxonomyFilteredSubtreeResponse) GetWarnings() []V2reportsWarning {
	if o == nil || o.Warnings == nil {
		var ret []V2reportsWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyFilteredSubtreeResponse) GetWarningsOk() ([]V2reportsWarning, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V2TaxonomyFilteredSubtreeResponse) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V2reportsWarning and assigns it to the Warnings field.
func (o *V2TaxonomyFilteredSubtreeResponse) SetWarnings(v []V2reportsWarning) {
	o.Warnings = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V2TaxonomyFilteredSubtreeResponse) GetErrors() []V2reportsError {
	if o == nil || o.Errors == nil {
		var ret []V2reportsError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyFilteredSubtreeResponse) GetErrorsOk() ([]V2reportsError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V2TaxonomyFilteredSubtreeResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V2reportsError and assigns it to the Errors field.
func (o *V2TaxonomyFilteredSubtreeResponse) SetErrors(v []V2reportsError) {
	o.Errors = v
}

func (o V2TaxonomyFilteredSubtreeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RootNodes != nil {
		toSerialize["root_nodes"] = o.RootNodes
	}
	if o.Edges != nil {
		toSerialize["edges"] = o.Edges
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableV2TaxonomyFilteredSubtreeResponse struct {
	value *V2TaxonomyFilteredSubtreeResponse
	isSet bool
}

func (v NullableV2TaxonomyFilteredSubtreeResponse) Get() *V2TaxonomyFilteredSubtreeResponse {
	return v.value
}

func (v *NullableV2TaxonomyFilteredSubtreeResponse) Set(val *V2TaxonomyFilteredSubtreeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TaxonomyFilteredSubtreeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TaxonomyFilteredSubtreeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TaxonomyFilteredSubtreeResponse(val *V2TaxonomyFilteredSubtreeResponse) *NullableV2TaxonomyFilteredSubtreeResponse {
	return &NullableV2TaxonomyFilteredSubtreeResponse{value: val, isSet: true}
}

func (v NullableV2TaxonomyFilteredSubtreeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TaxonomyFilteredSubtreeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


