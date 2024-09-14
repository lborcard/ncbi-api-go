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

// V2reportsNameAndAuthorityPublication struct for V2reportsNameAndAuthorityPublication
type V2reportsNameAndAuthorityPublication struct {
	Name *string `json:"name,omitempty"`
	Citation *string `json:"citation,omitempty"`
}

// NewV2reportsNameAndAuthorityPublication instantiates a new V2reportsNameAndAuthorityPublication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsNameAndAuthorityPublication() *V2reportsNameAndAuthorityPublication {
	this := V2reportsNameAndAuthorityPublication{}
	return &this
}

// NewV2reportsNameAndAuthorityPublicationWithDefaults instantiates a new V2reportsNameAndAuthorityPublication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsNameAndAuthorityPublicationWithDefaults() *V2reportsNameAndAuthorityPublication {
	this := V2reportsNameAndAuthorityPublication{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthorityPublication) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthorityPublication) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthorityPublication) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V2reportsNameAndAuthorityPublication) SetName(v string) {
	o.Name = &v
}

// GetCitation returns the Citation field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthorityPublication) GetCitation() string {
	if o == nil || o.Citation == nil {
		var ret string
		return ret
	}
	return *o.Citation
}

// GetCitationOk returns a tuple with the Citation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthorityPublication) GetCitationOk() (*string, bool) {
	if o == nil || o.Citation == nil {
		return nil, false
	}
	return o.Citation, true
}

// HasCitation returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthorityPublication) HasCitation() bool {
	if o != nil && o.Citation != nil {
		return true
	}

	return false
}

// SetCitation gets a reference to the given string and assigns it to the Citation field.
func (o *V2reportsNameAndAuthorityPublication) SetCitation(v string) {
	o.Citation = &v
}

func (o V2reportsNameAndAuthorityPublication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Citation != nil {
		toSerialize["citation"] = o.Citation
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsNameAndAuthorityPublication struct {
	value *V2reportsNameAndAuthorityPublication
	isSet bool
}

func (v NullableV2reportsNameAndAuthorityPublication) Get() *V2reportsNameAndAuthorityPublication {
	return v.value
}

func (v *NullableV2reportsNameAndAuthorityPublication) Set(val *V2reportsNameAndAuthorityPublication) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsNameAndAuthorityPublication) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsNameAndAuthorityPublication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsNameAndAuthorityPublication(val *V2reportsNameAndAuthorityPublication) *NullableV2reportsNameAndAuthorityPublication {
	return &NullableV2reportsNameAndAuthorityPublication{value: val, isSet: true}
}

func (v NullableV2reportsNameAndAuthorityPublication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsNameAndAuthorityPublication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


