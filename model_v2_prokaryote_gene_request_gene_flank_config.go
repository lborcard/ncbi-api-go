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

// V2ProkaryoteGeneRequestGeneFlankConfig struct for V2ProkaryoteGeneRequestGeneFlankConfig
type V2ProkaryoteGeneRequestGeneFlankConfig struct {
	Length *int32 `json:"length,omitempty"`
}

// NewV2ProkaryoteGeneRequestGeneFlankConfig instantiates a new V2ProkaryoteGeneRequestGeneFlankConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ProkaryoteGeneRequestGeneFlankConfig() *V2ProkaryoteGeneRequestGeneFlankConfig {
	this := V2ProkaryoteGeneRequestGeneFlankConfig{}
	return &this
}

// NewV2ProkaryoteGeneRequestGeneFlankConfigWithDefaults instantiates a new V2ProkaryoteGeneRequestGeneFlankConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ProkaryoteGeneRequestGeneFlankConfigWithDefaults() *V2ProkaryoteGeneRequestGeneFlankConfig {
	this := V2ProkaryoteGeneRequestGeneFlankConfig{}
	return &this
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *V2ProkaryoteGeneRequestGeneFlankConfig) GetLength() int32 {
	if o == nil || o.Length == nil {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ProkaryoteGeneRequestGeneFlankConfig) GetLengthOk() (*int32, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *V2ProkaryoteGeneRequestGeneFlankConfig) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *V2ProkaryoteGeneRequestGeneFlankConfig) SetLength(v int32) {
	o.Length = &v
}

func (o V2ProkaryoteGeneRequestGeneFlankConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	return json.Marshal(toSerialize)
}

type NullableV2ProkaryoteGeneRequestGeneFlankConfig struct {
	value *V2ProkaryoteGeneRequestGeneFlankConfig
	isSet bool
}

func (v NullableV2ProkaryoteGeneRequestGeneFlankConfig) Get() *V2ProkaryoteGeneRequestGeneFlankConfig {
	return v.value
}

func (v *NullableV2ProkaryoteGeneRequestGeneFlankConfig) Set(val *V2ProkaryoteGeneRequestGeneFlankConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ProkaryoteGeneRequestGeneFlankConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ProkaryoteGeneRequestGeneFlankConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ProkaryoteGeneRequestGeneFlankConfig(val *V2ProkaryoteGeneRequestGeneFlankConfig) *NullableV2ProkaryoteGeneRequestGeneFlankConfig {
	return &NullableV2ProkaryoteGeneRequestGeneFlankConfig{value: val, isSet: true}
}

func (v NullableV2ProkaryoteGeneRequestGeneFlankConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ProkaryoteGeneRequestGeneFlankConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


