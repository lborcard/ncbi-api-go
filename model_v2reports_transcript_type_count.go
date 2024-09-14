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

// V2reportsTranscriptTypeCount struct for V2reportsTranscriptTypeCount
type V2reportsTranscriptTypeCount struct {
	Type *V2reportsTranscriptTranscriptType `json:"type,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// NewV2reportsTranscriptTypeCount instantiates a new V2reportsTranscriptTypeCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsTranscriptTypeCount() *V2reportsTranscriptTypeCount {
	this := V2reportsTranscriptTypeCount{}
	var type_ V2reportsTranscriptTranscriptType = V2REPORTSTRANSCRIPTTRANSCRIPTTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// NewV2reportsTranscriptTypeCountWithDefaults instantiates a new V2reportsTranscriptTypeCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsTranscriptTypeCountWithDefaults() *V2reportsTranscriptTypeCount {
	this := V2reportsTranscriptTypeCount{}
	var type_ V2reportsTranscriptTranscriptType = V2REPORTSTRANSCRIPTTRANSCRIPTTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V2reportsTranscriptTypeCount) GetType() V2reportsTranscriptTranscriptType {
	if o == nil || o.Type == nil {
		var ret V2reportsTranscriptTranscriptType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsTranscriptTypeCount) GetTypeOk() (*V2reportsTranscriptTranscriptType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V2reportsTranscriptTypeCount) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given V2reportsTranscriptTranscriptType and assigns it to the Type field.
func (o *V2reportsTranscriptTypeCount) SetType(v V2reportsTranscriptTranscriptType) {
	o.Type = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V2reportsTranscriptTypeCount) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsTranscriptTypeCount) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V2reportsTranscriptTypeCount) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V2reportsTranscriptTypeCount) SetCount(v int32) {
	o.Count = &v
}

func (o V2reportsTranscriptTypeCount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsTranscriptTypeCount struct {
	value *V2reportsTranscriptTypeCount
	isSet bool
}

func (v NullableV2reportsTranscriptTypeCount) Get() *V2reportsTranscriptTypeCount {
	return v.value
}

func (v *NullableV2reportsTranscriptTypeCount) Set(val *V2reportsTranscriptTypeCount) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsTranscriptTypeCount) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsTranscriptTypeCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsTranscriptTypeCount(val *V2reportsTranscriptTypeCount) *NullableV2reportsTranscriptTypeCount {
	return &NullableV2reportsTranscriptTypeCount{value: val, isSet: true}
}

func (v NullableV2reportsTranscriptTypeCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsTranscriptTypeCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


