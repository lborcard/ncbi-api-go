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

// V2AssemblyRevisionHistoryRequest struct for V2AssemblyRevisionHistoryRequest
type V2AssemblyRevisionHistoryRequest struct {
	Accession *string `json:"accession,omitempty"`
}

// NewV2AssemblyRevisionHistoryRequest instantiates a new V2AssemblyRevisionHistoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2AssemblyRevisionHistoryRequest() *V2AssemblyRevisionHistoryRequest {
	this := V2AssemblyRevisionHistoryRequest{}
	return &this
}

// NewV2AssemblyRevisionHistoryRequestWithDefaults instantiates a new V2AssemblyRevisionHistoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2AssemblyRevisionHistoryRequestWithDefaults() *V2AssemblyRevisionHistoryRequest {
	this := V2AssemblyRevisionHistoryRequest{}
	return &this
}

// GetAccession returns the Accession field value if set, zero value otherwise.
func (o *V2AssemblyRevisionHistoryRequest) GetAccession() string {
	if o == nil || o.Accession == nil {
		var ret string
		return ret
	}
	return *o.Accession
}

// GetAccessionOk returns a tuple with the Accession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2AssemblyRevisionHistoryRequest) GetAccessionOk() (*string, bool) {
	if o == nil || o.Accession == nil {
		return nil, false
	}
	return o.Accession, true
}

// HasAccession returns a boolean if a field has been set.
func (o *V2AssemblyRevisionHistoryRequest) HasAccession() bool {
	if o != nil && o.Accession != nil {
		return true
	}

	return false
}

// SetAccession gets a reference to the given string and assigns it to the Accession field.
func (o *V2AssemblyRevisionHistoryRequest) SetAccession(v string) {
	o.Accession = &v
}

func (o V2AssemblyRevisionHistoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accession != nil {
		toSerialize["accession"] = o.Accession
	}
	return json.Marshal(toSerialize)
}

type NullableV2AssemblyRevisionHistoryRequest struct {
	value *V2AssemblyRevisionHistoryRequest
	isSet bool
}

func (v NullableV2AssemblyRevisionHistoryRequest) Get() *V2AssemblyRevisionHistoryRequest {
	return v.value
}

func (v *NullableV2AssemblyRevisionHistoryRequest) Set(val *V2AssemblyRevisionHistoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssemblyRevisionHistoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssemblyRevisionHistoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssemblyRevisionHistoryRequest(val *V2AssemblyRevisionHistoryRequest) *NullableV2AssemblyRevisionHistoryRequest {
	return &NullableV2AssemblyRevisionHistoryRequest{value: val, isSet: true}
}

func (v NullableV2AssemblyRevisionHistoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssemblyRevisionHistoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


