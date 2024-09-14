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

// V2GeneDatasetReportsRequestSymbolsForTaxon struct for V2GeneDatasetReportsRequestSymbolsForTaxon
type V2GeneDatasetReportsRequestSymbolsForTaxon struct {
	Symbols []string `json:"symbols,omitempty"`
	Taxon *string `json:"taxon,omitempty"`
}

// NewV2GeneDatasetReportsRequestSymbolsForTaxon instantiates a new V2GeneDatasetReportsRequestSymbolsForTaxon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2GeneDatasetReportsRequestSymbolsForTaxon() *V2GeneDatasetReportsRequestSymbolsForTaxon {
	this := V2GeneDatasetReportsRequestSymbolsForTaxon{}
	return &this
}

// NewV2GeneDatasetReportsRequestSymbolsForTaxonWithDefaults instantiates a new V2GeneDatasetReportsRequestSymbolsForTaxon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2GeneDatasetReportsRequestSymbolsForTaxonWithDefaults() *V2GeneDatasetReportsRequestSymbolsForTaxon {
	this := V2GeneDatasetReportsRequestSymbolsForTaxon{}
	return &this
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *V2GeneDatasetReportsRequestSymbolsForTaxon) GetSymbols() []string {
	if o == nil || o.Symbols == nil {
		var ret []string
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2GeneDatasetReportsRequestSymbolsForTaxon) GetSymbolsOk() ([]string, bool) {
	if o == nil || o.Symbols == nil {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *V2GeneDatasetReportsRequestSymbolsForTaxon) HasSymbols() bool {
	if o != nil && o.Symbols != nil {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []string and assigns it to the Symbols field.
func (o *V2GeneDatasetReportsRequestSymbolsForTaxon) SetSymbols(v []string) {
	o.Symbols = v
}

// GetTaxon returns the Taxon field value if set, zero value otherwise.
func (o *V2GeneDatasetReportsRequestSymbolsForTaxon) GetTaxon() string {
	if o == nil || o.Taxon == nil {
		var ret string
		return ret
	}
	return *o.Taxon
}

// GetTaxonOk returns a tuple with the Taxon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2GeneDatasetReportsRequestSymbolsForTaxon) GetTaxonOk() (*string, bool) {
	if o == nil || o.Taxon == nil {
		return nil, false
	}
	return o.Taxon, true
}

// HasTaxon returns a boolean if a field has been set.
func (o *V2GeneDatasetReportsRequestSymbolsForTaxon) HasTaxon() bool {
	if o != nil && o.Taxon != nil {
		return true
	}

	return false
}

// SetTaxon gets a reference to the given string and assigns it to the Taxon field.
func (o *V2GeneDatasetReportsRequestSymbolsForTaxon) SetTaxon(v string) {
	o.Taxon = &v
}

func (o V2GeneDatasetReportsRequestSymbolsForTaxon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbols != nil {
		toSerialize["symbols"] = o.Symbols
	}
	if o.Taxon != nil {
		toSerialize["taxon"] = o.Taxon
	}
	return json.Marshal(toSerialize)
}

type NullableV2GeneDatasetReportsRequestSymbolsForTaxon struct {
	value *V2GeneDatasetReportsRequestSymbolsForTaxon
	isSet bool
}

func (v NullableV2GeneDatasetReportsRequestSymbolsForTaxon) Get() *V2GeneDatasetReportsRequestSymbolsForTaxon {
	return v.value
}

func (v *NullableV2GeneDatasetReportsRequestSymbolsForTaxon) Set(val *V2GeneDatasetReportsRequestSymbolsForTaxon) {
	v.value = val
	v.isSet = true
}

func (v NullableV2GeneDatasetReportsRequestSymbolsForTaxon) IsSet() bool {
	return v.isSet
}

func (v *NullableV2GeneDatasetReportsRequestSymbolsForTaxon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2GeneDatasetReportsRequestSymbolsForTaxon(val *V2GeneDatasetReportsRequestSymbolsForTaxon) *NullableV2GeneDatasetReportsRequestSymbolsForTaxon {
	return &NullableV2GeneDatasetReportsRequestSymbolsForTaxon{value: val, isSet: true}
}

func (v NullableV2GeneDatasetReportsRequestSymbolsForTaxon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2GeneDatasetReportsRequestSymbolsForTaxon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


