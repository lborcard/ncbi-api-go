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

// V2reportsVirusAssemblySubmitterInfo struct for V2reportsVirusAssemblySubmitterInfo
type V2reportsVirusAssemblySubmitterInfo struct {
	Names []string `json:"names,omitempty"`
	Affiliation *string `json:"affiliation,omitempty"`
	Country *string `json:"country,omitempty"`
}

// NewV2reportsVirusAssemblySubmitterInfo instantiates a new V2reportsVirusAssemblySubmitterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsVirusAssemblySubmitterInfo() *V2reportsVirusAssemblySubmitterInfo {
	this := V2reportsVirusAssemblySubmitterInfo{}
	return &this
}

// NewV2reportsVirusAssemblySubmitterInfoWithDefaults instantiates a new V2reportsVirusAssemblySubmitterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsVirusAssemblySubmitterInfoWithDefaults() *V2reportsVirusAssemblySubmitterInfo {
	this := V2reportsVirusAssemblySubmitterInfo{}
	return &this
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *V2reportsVirusAssemblySubmitterInfo) GetNames() []string {
	if o == nil || o.Names == nil {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsVirusAssemblySubmitterInfo) GetNamesOk() ([]string, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *V2reportsVirusAssemblySubmitterInfo) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *V2reportsVirusAssemblySubmitterInfo) SetNames(v []string) {
	o.Names = v
}

// GetAffiliation returns the Affiliation field value if set, zero value otherwise.
func (o *V2reportsVirusAssemblySubmitterInfo) GetAffiliation() string {
	if o == nil || o.Affiliation == nil {
		var ret string
		return ret
	}
	return *o.Affiliation
}

// GetAffiliationOk returns a tuple with the Affiliation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsVirusAssemblySubmitterInfo) GetAffiliationOk() (*string, bool) {
	if o == nil || o.Affiliation == nil {
		return nil, false
	}
	return o.Affiliation, true
}

// HasAffiliation returns a boolean if a field has been set.
func (o *V2reportsVirusAssemblySubmitterInfo) HasAffiliation() bool {
	if o != nil && o.Affiliation != nil {
		return true
	}

	return false
}

// SetAffiliation gets a reference to the given string and assigns it to the Affiliation field.
func (o *V2reportsVirusAssemblySubmitterInfo) SetAffiliation(v string) {
	o.Affiliation = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *V2reportsVirusAssemblySubmitterInfo) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsVirusAssemblySubmitterInfo) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *V2reportsVirusAssemblySubmitterInfo) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *V2reportsVirusAssemblySubmitterInfo) SetCountry(v string) {
	o.Country = &v
}

func (o V2reportsVirusAssemblySubmitterInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	if o.Affiliation != nil {
		toSerialize["affiliation"] = o.Affiliation
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsVirusAssemblySubmitterInfo struct {
	value *V2reportsVirusAssemblySubmitterInfo
	isSet bool
}

func (v NullableV2reportsVirusAssemblySubmitterInfo) Get() *V2reportsVirusAssemblySubmitterInfo {
	return v.value
}

func (v *NullableV2reportsVirusAssemblySubmitterInfo) Set(val *V2reportsVirusAssemblySubmitterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsVirusAssemblySubmitterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsVirusAssemblySubmitterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsVirusAssemblySubmitterInfo(val *V2reportsVirusAssemblySubmitterInfo) *NullableV2reportsVirusAssemblySubmitterInfo {
	return &NullableV2reportsVirusAssemblySubmitterInfo{value: val, isSet: true}
}

func (v NullableV2reportsVirusAssemblySubmitterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsVirusAssemblySubmitterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


