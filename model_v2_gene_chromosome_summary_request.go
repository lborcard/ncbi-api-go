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

// V2GeneChromosomeSummaryRequest struct for V2GeneChromosomeSummaryRequest
type V2GeneChromosomeSummaryRequest struct {
	Taxon *string `json:"taxon,omitempty"`
	AnnotationName *string `json:"annotation_name,omitempty"`
}

// NewV2GeneChromosomeSummaryRequest instantiates a new V2GeneChromosomeSummaryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2GeneChromosomeSummaryRequest() *V2GeneChromosomeSummaryRequest {
	this := V2GeneChromosomeSummaryRequest{}
	return &this
}

// NewV2GeneChromosomeSummaryRequestWithDefaults instantiates a new V2GeneChromosomeSummaryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2GeneChromosomeSummaryRequestWithDefaults() *V2GeneChromosomeSummaryRequest {
	this := V2GeneChromosomeSummaryRequest{}
	return &this
}

// GetTaxon returns the Taxon field value if set, zero value otherwise.
func (o *V2GeneChromosomeSummaryRequest) GetTaxon() string {
	if o == nil || o.Taxon == nil {
		var ret string
		return ret
	}
	return *o.Taxon
}

// GetTaxonOk returns a tuple with the Taxon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2GeneChromosomeSummaryRequest) GetTaxonOk() (*string, bool) {
	if o == nil || o.Taxon == nil {
		return nil, false
	}
	return o.Taxon, true
}

// HasTaxon returns a boolean if a field has been set.
func (o *V2GeneChromosomeSummaryRequest) HasTaxon() bool {
	if o != nil && o.Taxon != nil {
		return true
	}

	return false
}

// SetTaxon gets a reference to the given string and assigns it to the Taxon field.
func (o *V2GeneChromosomeSummaryRequest) SetTaxon(v string) {
	o.Taxon = &v
}

// GetAnnotationName returns the AnnotationName field value if set, zero value otherwise.
func (o *V2GeneChromosomeSummaryRequest) GetAnnotationName() string {
	if o == nil || o.AnnotationName == nil {
		var ret string
		return ret
	}
	return *o.AnnotationName
}

// GetAnnotationNameOk returns a tuple with the AnnotationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2GeneChromosomeSummaryRequest) GetAnnotationNameOk() (*string, bool) {
	if o == nil || o.AnnotationName == nil {
		return nil, false
	}
	return o.AnnotationName, true
}

// HasAnnotationName returns a boolean if a field has been set.
func (o *V2GeneChromosomeSummaryRequest) HasAnnotationName() bool {
	if o != nil && o.AnnotationName != nil {
		return true
	}

	return false
}

// SetAnnotationName gets a reference to the given string and assigns it to the AnnotationName field.
func (o *V2GeneChromosomeSummaryRequest) SetAnnotationName(v string) {
	o.AnnotationName = &v
}

func (o V2GeneChromosomeSummaryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Taxon != nil {
		toSerialize["taxon"] = o.Taxon
	}
	if o.AnnotationName != nil {
		toSerialize["annotation_name"] = o.AnnotationName
	}
	return json.Marshal(toSerialize)
}

type NullableV2GeneChromosomeSummaryRequest struct {
	value *V2GeneChromosomeSummaryRequest
	isSet bool
}

func (v NullableV2GeneChromosomeSummaryRequest) Get() *V2GeneChromosomeSummaryRequest {
	return v.value
}

func (v *NullableV2GeneChromosomeSummaryRequest) Set(val *V2GeneChromosomeSummaryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2GeneChromosomeSummaryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2GeneChromosomeSummaryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2GeneChromosomeSummaryRequest(val *V2GeneChromosomeSummaryRequest) *NullableV2GeneChromosomeSummaryRequest {
	return &NullableV2GeneChromosomeSummaryRequest{value: val, isSet: true}
}

func (v NullableV2GeneChromosomeSummaryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2GeneChromosomeSummaryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


