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

// V2reportsLinkedAssembly struct for V2reportsLinkedAssembly
type V2reportsLinkedAssembly struct {
	LinkedAssembly *string `json:"linked_assembly,omitempty"`
	AssemblyType *V2reportsLinkedAssemblyType `json:"assembly_type,omitempty"`
}

// NewV2reportsLinkedAssembly instantiates a new V2reportsLinkedAssembly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsLinkedAssembly() *V2reportsLinkedAssembly {
	this := V2reportsLinkedAssembly{}
	var assemblyType V2reportsLinkedAssemblyType = V2REPORTSLINKEDASSEMBLYTYPE_LINKED_ASSEMBLY_TYPE_UNKNOWN
	this.AssemblyType = &assemblyType
	return &this
}

// NewV2reportsLinkedAssemblyWithDefaults instantiates a new V2reportsLinkedAssembly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsLinkedAssemblyWithDefaults() *V2reportsLinkedAssembly {
	this := V2reportsLinkedAssembly{}
	var assemblyType V2reportsLinkedAssemblyType = V2REPORTSLINKEDASSEMBLYTYPE_LINKED_ASSEMBLY_TYPE_UNKNOWN
	this.AssemblyType = &assemblyType
	return &this
}

// GetLinkedAssembly returns the LinkedAssembly field value if set, zero value otherwise.
func (o *V2reportsLinkedAssembly) GetLinkedAssembly() string {
	if o == nil || o.LinkedAssembly == nil {
		var ret string
		return ret
	}
	return *o.LinkedAssembly
}

// GetLinkedAssemblyOk returns a tuple with the LinkedAssembly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsLinkedAssembly) GetLinkedAssemblyOk() (*string, bool) {
	if o == nil || o.LinkedAssembly == nil {
		return nil, false
	}
	return o.LinkedAssembly, true
}

// HasLinkedAssembly returns a boolean if a field has been set.
func (o *V2reportsLinkedAssembly) HasLinkedAssembly() bool {
	if o != nil && o.LinkedAssembly != nil {
		return true
	}

	return false
}

// SetLinkedAssembly gets a reference to the given string and assigns it to the LinkedAssembly field.
func (o *V2reportsLinkedAssembly) SetLinkedAssembly(v string) {
	o.LinkedAssembly = &v
}

// GetAssemblyType returns the AssemblyType field value if set, zero value otherwise.
func (o *V2reportsLinkedAssembly) GetAssemblyType() V2reportsLinkedAssemblyType {
	if o == nil || o.AssemblyType == nil {
		var ret V2reportsLinkedAssemblyType
		return ret
	}
	return *o.AssemblyType
}

// GetAssemblyTypeOk returns a tuple with the AssemblyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsLinkedAssembly) GetAssemblyTypeOk() (*V2reportsLinkedAssemblyType, bool) {
	if o == nil || o.AssemblyType == nil {
		return nil, false
	}
	return o.AssemblyType, true
}

// HasAssemblyType returns a boolean if a field has been set.
func (o *V2reportsLinkedAssembly) HasAssemblyType() bool {
	if o != nil && o.AssemblyType != nil {
		return true
	}

	return false
}

// SetAssemblyType gets a reference to the given V2reportsLinkedAssemblyType and assigns it to the AssemblyType field.
func (o *V2reportsLinkedAssembly) SetAssemblyType(v V2reportsLinkedAssemblyType) {
	o.AssemblyType = &v
}

func (o V2reportsLinkedAssembly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinkedAssembly != nil {
		toSerialize["linked_assembly"] = o.LinkedAssembly
	}
	if o.AssemblyType != nil {
		toSerialize["assembly_type"] = o.AssemblyType
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsLinkedAssembly struct {
	value *V2reportsLinkedAssembly
	isSet bool
}

func (v NullableV2reportsLinkedAssembly) Get() *V2reportsLinkedAssembly {
	return v.value
}

func (v *NullableV2reportsLinkedAssembly) Set(val *V2reportsLinkedAssembly) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsLinkedAssembly) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsLinkedAssembly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsLinkedAssembly(val *V2reportsLinkedAssembly) *NullableV2reportsLinkedAssembly {
	return &NullableV2reportsLinkedAssembly{value: val, isSet: true}
}

func (v NullableV2reportsLinkedAssembly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsLinkedAssembly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


