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

// V2reportsVirusAnnotationReport struct for V2reportsVirusAnnotationReport
type V2reportsVirusAnnotationReport struct {
	Accession *string `json:"accession,omitempty"`
	IsolateName *string `json:"isolate_name,omitempty"`
	Genes []V2reportsVirusGene `json:"genes,omitempty"`
}

// NewV2reportsVirusAnnotationReport instantiates a new V2reportsVirusAnnotationReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsVirusAnnotationReport() *V2reportsVirusAnnotationReport {
	this := V2reportsVirusAnnotationReport{}
	return &this
}

// NewV2reportsVirusAnnotationReportWithDefaults instantiates a new V2reportsVirusAnnotationReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsVirusAnnotationReportWithDefaults() *V2reportsVirusAnnotationReport {
	this := V2reportsVirusAnnotationReport{}
	return &this
}

// GetAccession returns the Accession field value if set, zero value otherwise.
func (o *V2reportsVirusAnnotationReport) GetAccession() string {
	if o == nil || o.Accession == nil {
		var ret string
		return ret
	}
	return *o.Accession
}

// GetAccessionOk returns a tuple with the Accession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsVirusAnnotationReport) GetAccessionOk() (*string, bool) {
	if o == nil || o.Accession == nil {
		return nil, false
	}
	return o.Accession, true
}

// HasAccession returns a boolean if a field has been set.
func (o *V2reportsVirusAnnotationReport) HasAccession() bool {
	if o != nil && o.Accession != nil {
		return true
	}

	return false
}

// SetAccession gets a reference to the given string and assigns it to the Accession field.
func (o *V2reportsVirusAnnotationReport) SetAccession(v string) {
	o.Accession = &v
}

// GetIsolateName returns the IsolateName field value if set, zero value otherwise.
func (o *V2reportsVirusAnnotationReport) GetIsolateName() string {
	if o == nil || o.IsolateName == nil {
		var ret string
		return ret
	}
	return *o.IsolateName
}

// GetIsolateNameOk returns a tuple with the IsolateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsVirusAnnotationReport) GetIsolateNameOk() (*string, bool) {
	if o == nil || o.IsolateName == nil {
		return nil, false
	}
	return o.IsolateName, true
}

// HasIsolateName returns a boolean if a field has been set.
func (o *V2reportsVirusAnnotationReport) HasIsolateName() bool {
	if o != nil && o.IsolateName != nil {
		return true
	}

	return false
}

// SetIsolateName gets a reference to the given string and assigns it to the IsolateName field.
func (o *V2reportsVirusAnnotationReport) SetIsolateName(v string) {
	o.IsolateName = &v
}

// GetGenes returns the Genes field value if set, zero value otherwise.
func (o *V2reportsVirusAnnotationReport) GetGenes() []V2reportsVirusGene {
	if o == nil || o.Genes == nil {
		var ret []V2reportsVirusGene
		return ret
	}
	return o.Genes
}

// GetGenesOk returns a tuple with the Genes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsVirusAnnotationReport) GetGenesOk() ([]V2reportsVirusGene, bool) {
	if o == nil || o.Genes == nil {
		return nil, false
	}
	return o.Genes, true
}

// HasGenes returns a boolean if a field has been set.
func (o *V2reportsVirusAnnotationReport) HasGenes() bool {
	if o != nil && o.Genes != nil {
		return true
	}

	return false
}

// SetGenes gets a reference to the given []V2reportsVirusGene and assigns it to the Genes field.
func (o *V2reportsVirusAnnotationReport) SetGenes(v []V2reportsVirusGene) {
	o.Genes = v
}

func (o V2reportsVirusAnnotationReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accession != nil {
		toSerialize["accession"] = o.Accession
	}
	if o.IsolateName != nil {
		toSerialize["isolate_name"] = o.IsolateName
	}
	if o.Genes != nil {
		toSerialize["genes"] = o.Genes
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsVirusAnnotationReport struct {
	value *V2reportsVirusAnnotationReport
	isSet bool
}

func (v NullableV2reportsVirusAnnotationReport) Get() *V2reportsVirusAnnotationReport {
	return v.value
}

func (v *NullableV2reportsVirusAnnotationReport) Set(val *V2reportsVirusAnnotationReport) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsVirusAnnotationReport) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsVirusAnnotationReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsVirusAnnotationReport(val *V2reportsVirusAnnotationReport) *NullableV2reportsVirusAnnotationReport {
	return &NullableV2reportsVirusAnnotationReport{value: val, isSet: true}
}

func (v NullableV2reportsVirusAnnotationReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsVirusAnnotationReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


