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

// V2reportsTaxonomyReportMatch struct for V2reportsTaxonomyReportMatch
type V2reportsTaxonomyReportMatch struct {
	Taxonomy *V2reportsTaxonomyNode `json:"taxonomy,omitempty"`
	Query []string `json:"query,omitempty"`
	Warning *V2reportsWarning `json:"warning,omitempty"`
	Errors []V2reportsError `json:"errors,omitempty"`
}

// NewV2reportsTaxonomyReportMatch instantiates a new V2reportsTaxonomyReportMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsTaxonomyReportMatch() *V2reportsTaxonomyReportMatch {
	this := V2reportsTaxonomyReportMatch{}
	return &this
}

// NewV2reportsTaxonomyReportMatchWithDefaults instantiates a new V2reportsTaxonomyReportMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsTaxonomyReportMatchWithDefaults() *V2reportsTaxonomyReportMatch {
	this := V2reportsTaxonomyReportMatch{}
	return &this
}

// GetTaxonomy returns the Taxonomy field value if set, zero value otherwise.
func (o *V2reportsTaxonomyReportMatch) GetTaxonomy() V2reportsTaxonomyNode {
	if o == nil || o.Taxonomy == nil {
		var ret V2reportsTaxonomyNode
		return ret
	}
	return *o.Taxonomy
}

// GetTaxonomyOk returns a tuple with the Taxonomy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsTaxonomyReportMatch) GetTaxonomyOk() (*V2reportsTaxonomyNode, bool) {
	if o == nil || o.Taxonomy == nil {
		return nil, false
	}
	return o.Taxonomy, true
}

// HasTaxonomy returns a boolean if a field has been set.
func (o *V2reportsTaxonomyReportMatch) HasTaxonomy() bool {
	if o != nil && o.Taxonomy != nil {
		return true
	}

	return false
}

// SetTaxonomy gets a reference to the given V2reportsTaxonomyNode and assigns it to the Taxonomy field.
func (o *V2reportsTaxonomyReportMatch) SetTaxonomy(v V2reportsTaxonomyNode) {
	o.Taxonomy = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *V2reportsTaxonomyReportMatch) GetQuery() []string {
	if o == nil || o.Query == nil {
		var ret []string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsTaxonomyReportMatch) GetQueryOk() ([]string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *V2reportsTaxonomyReportMatch) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given []string and assigns it to the Query field.
func (o *V2reportsTaxonomyReportMatch) SetQuery(v []string) {
	o.Query = v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *V2reportsTaxonomyReportMatch) GetWarning() V2reportsWarning {
	if o == nil || o.Warning == nil {
		var ret V2reportsWarning
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsTaxonomyReportMatch) GetWarningOk() (*V2reportsWarning, bool) {
	if o == nil || o.Warning == nil {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *V2reportsTaxonomyReportMatch) HasWarning() bool {
	if o != nil && o.Warning != nil {
		return true
	}

	return false
}

// SetWarning gets a reference to the given V2reportsWarning and assigns it to the Warning field.
func (o *V2reportsTaxonomyReportMatch) SetWarning(v V2reportsWarning) {
	o.Warning = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V2reportsTaxonomyReportMatch) GetErrors() []V2reportsError {
	if o == nil || o.Errors == nil {
		var ret []V2reportsError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsTaxonomyReportMatch) GetErrorsOk() ([]V2reportsError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V2reportsTaxonomyReportMatch) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V2reportsError and assigns it to the Errors field.
func (o *V2reportsTaxonomyReportMatch) SetErrors(v []V2reportsError) {
	o.Errors = v
}

func (o V2reportsTaxonomyReportMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Taxonomy != nil {
		toSerialize["taxonomy"] = o.Taxonomy
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Warning != nil {
		toSerialize["warning"] = o.Warning
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsTaxonomyReportMatch struct {
	value *V2reportsTaxonomyReportMatch
	isSet bool
}

func (v NullableV2reportsTaxonomyReportMatch) Get() *V2reportsTaxonomyReportMatch {
	return v.value
}

func (v *NullableV2reportsTaxonomyReportMatch) Set(val *V2reportsTaxonomyReportMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsTaxonomyReportMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsTaxonomyReportMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsTaxonomyReportMatch(val *V2reportsTaxonomyReportMatch) *NullableV2reportsTaxonomyReportMatch {
	return &NullableV2reportsTaxonomyReportMatch{value: val, isSet: true}
}

func (v NullableV2reportsTaxonomyReportMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsTaxonomyReportMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


