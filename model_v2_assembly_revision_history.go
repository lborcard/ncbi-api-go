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

// V2AssemblyRevisionHistory struct for V2AssemblyRevisionHistory
type V2AssemblyRevisionHistory struct {
	AssemblyRevisions []V2reportsAssemblyRevision `json:"assembly_revisions,omitempty"`
	TotalCount *int32 `json:"total_count,omitempty"`
}

// NewV2AssemblyRevisionHistory instantiates a new V2AssemblyRevisionHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssemblyRevisionHistory() *V2AssemblyRevisionHistory {
	this := V2AssemblyRevisionHistory{}
	return &this
}

// NewV2AssemblyRevisionHistoryWithDefaults instantiates a new V2AssemblyRevisionHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssemblyRevisionHistoryWithDefaults() *V2AssemblyRevisionHistory {
	this := V2AssemblyRevisionHistory{}
	return &this
}

// GetAssemblyRevisions returns the AssemblyRevisions field value if set, zero value otherwise.
func (o *V2AssemblyRevisionHistory) GetAssemblyRevisions() []V2reportsAssemblyRevision {
	if o == nil || o.AssemblyRevisions == nil {
		var ret []V2reportsAssemblyRevision
		return ret
	}
	return o.AssemblyRevisions
}

// GetAssemblyRevisionsOk returns a tuple with the AssemblyRevisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssemblyRevisionHistory) GetAssemblyRevisionsOk() ([]V2reportsAssemblyRevision, bool) {
	if o == nil || o.AssemblyRevisions == nil {
		return nil, false
	}
	return o.AssemblyRevisions, true
}

// HasAssemblyRevisions returns a boolean if a field has been set.
func (o *V2AssemblyRevisionHistory) HasAssemblyRevisions() bool {
	if o != nil && o.AssemblyRevisions != nil {
		return true
	}

	return false
}

// SetAssemblyRevisions gets a reference to the given []V2reportsAssemblyRevision and assigns it to the AssemblyRevisions field.
func (o *V2AssemblyRevisionHistory) SetAssemblyRevisions(v []V2reportsAssemblyRevision) {
	o.AssemblyRevisions = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *V2AssemblyRevisionHistory) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssemblyRevisionHistory) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *V2AssemblyRevisionHistory) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *V2AssemblyRevisionHistory) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o V2AssemblyRevisionHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssemblyRevisions != nil {
		toSerialize["assembly_revisions"] = o.AssemblyRevisions
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableV2AssemblyRevisionHistory struct {
	value *V2AssemblyRevisionHistory
	isSet bool
}

func (v NullableV2AssemblyRevisionHistory) Get() *V2AssemblyRevisionHistory {
	return v.value
}

func (v *NullableV2AssemblyRevisionHistory) Set(val *V2AssemblyRevisionHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssemblyRevisionHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssemblyRevisionHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssemblyRevisionHistory(val *V2AssemblyRevisionHistory) *NullableV2AssemblyRevisionHistory {
	return &NullableV2AssemblyRevisionHistory{value: val, isSet: true}
}

func (v NullableV2AssemblyRevisionHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssemblyRevisionHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


