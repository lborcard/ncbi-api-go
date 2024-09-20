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

// V2reportsAssemblyRevision struct for V2reportsAssemblyRevision
type V2reportsAssemblyRevision struct {
	GenbankAccession *string `json:"genbank_accession,omitempty"`
	RefseqAccession *string `json:"refseq_accession,omitempty"`
	AssemblyName *string `json:"assembly_name,omitempty"`
	AssemblyLevel *V2reportsAssemblyLevel `json:"assembly_level,omitempty"`
	ReleaseDate *string `json:"release_date,omitempty"`
	SubmissionDate *string `json:"submission_date,omitempty"`
	SequencingTechnology *string `json:"sequencing_technology,omitempty"`
}

// NewV2reportsAssemblyRevision instantiates a new V2reportsAssemblyRevision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsAssemblyRevision() *V2reportsAssemblyRevision {
	this := V2reportsAssemblyRevision{}
	var assemblyLevel V2reportsAssemblyLevel = V2REPORTSASSEMBLYLEVEL_CHROMOSOME
	this.AssemblyLevel = &assemblyLevel
	return &this
}

// NewV2reportsAssemblyRevisionWithDefaults instantiates a new V2reportsAssemblyRevision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsAssemblyRevisionWithDefaults() *V2reportsAssemblyRevision {
	this := V2reportsAssemblyRevision{}
	var assemblyLevel V2reportsAssemblyLevel = V2REPORTSASSEMBLYLEVEL_CHROMOSOME
	this.AssemblyLevel = &assemblyLevel
	return &this
}

// GetGenbankAccession returns the GenbankAccession field value if set, zero value otherwise.
func (o *V2reportsAssemblyRevision) GetGenbankAccession() string {
	if o == nil || o.GenbankAccession == nil {
		var ret string
		return ret
	}
	return *o.GenbankAccession
}

// GetGenbankAccessionOk returns a tuple with the GenbankAccession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyRevision) GetGenbankAccessionOk() (*string, bool) {
	if o == nil || o.GenbankAccession == nil {
		return nil, false
	}
	return o.GenbankAccession, true
}

// HasGenbankAccession returns a boolean if a field has been set.
func (o *V2reportsAssemblyRevision) HasGenbankAccession() bool {
	if o != nil && o.GenbankAccession != nil {
		return true
	}

	return false
}

// SetGenbankAccession gets a reference to the given string and assigns it to the GenbankAccession field.
func (o *V2reportsAssemblyRevision) SetGenbankAccession(v string) {
	o.GenbankAccession = &v
}

// GetRefseqAccession returns the RefseqAccession field value if set, zero value otherwise.
func (o *V2reportsAssemblyRevision) GetRefseqAccession() string {
	if o == nil || o.RefseqAccession == nil {
		var ret string
		return ret
	}
	return *o.RefseqAccession
}

// GetRefseqAccessionOk returns a tuple with the RefseqAccession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyRevision) GetRefseqAccessionOk() (*string, bool) {
	if o == nil || o.RefseqAccession == nil {
		return nil, false
	}
	return o.RefseqAccession, true
}

// HasRefseqAccession returns a boolean if a field has been set.
func (o *V2reportsAssemblyRevision) HasRefseqAccession() bool {
	if o != nil && o.RefseqAccession != nil {
		return true
	}

	return false
}

// SetRefseqAccession gets a reference to the given string and assigns it to the RefseqAccession field.
func (o *V2reportsAssemblyRevision) SetRefseqAccession(v string) {
	o.RefseqAccession = &v
}

// GetAssemblyName returns the AssemblyName field value if set, zero value otherwise.
func (o *V2reportsAssemblyRevision) GetAssemblyName() string {
	if o == nil || o.AssemblyName == nil {
		var ret string
		return ret
	}
	return *o.AssemblyName
}

// GetAssemblyNameOk returns a tuple with the AssemblyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyRevision) GetAssemblyNameOk() (*string, bool) {
	if o == nil || o.AssemblyName == nil {
		return nil, false
	}
	return o.AssemblyName, true
}

// HasAssemblyName returns a boolean if a field has been set.
func (o *V2reportsAssemblyRevision) HasAssemblyName() bool {
	if o != nil && o.AssemblyName != nil {
		return true
	}

	return false
}

// SetAssemblyName gets a reference to the given string and assigns it to the AssemblyName field.
func (o *V2reportsAssemblyRevision) SetAssemblyName(v string) {
	o.AssemblyName = &v
}

// GetAssemblyLevel returns the AssemblyLevel field value if set, zero value otherwise.
func (o *V2reportsAssemblyRevision) GetAssemblyLevel() V2reportsAssemblyLevel {
	if o == nil || o.AssemblyLevel == nil {
		var ret V2reportsAssemblyLevel
		return ret
	}
	return *o.AssemblyLevel
}

// GetAssemblyLevelOk returns a tuple with the AssemblyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyRevision) GetAssemblyLevelOk() (*V2reportsAssemblyLevel, bool) {
	if o == nil || o.AssemblyLevel == nil {
		return nil, false
	}
	return o.AssemblyLevel, true
}

// HasAssemblyLevel returns a boolean if a field has been set.
func (o *V2reportsAssemblyRevision) HasAssemblyLevel() bool {
	if o != nil && o.AssemblyLevel != nil {
		return true
	}

	return false
}

// SetAssemblyLevel gets a reference to the given V2reportsAssemblyLevel and assigns it to the AssemblyLevel field.
func (o *V2reportsAssemblyRevision) SetAssemblyLevel(v V2reportsAssemblyLevel) {
	o.AssemblyLevel = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *V2reportsAssemblyRevision) GetReleaseDate() string {
	if o == nil || o.ReleaseDate == nil {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyRevision) GetReleaseDateOk() (*string, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *V2reportsAssemblyRevision) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate != nil {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *V2reportsAssemblyRevision) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

// GetSubmissionDate returns the SubmissionDate field value if set, zero value otherwise.
func (o *V2reportsAssemblyRevision) GetSubmissionDate() string {
	if o == nil || o.SubmissionDate == nil {
		var ret string
		return ret
	}
	return *o.SubmissionDate
}

// GetSubmissionDateOk returns a tuple with the SubmissionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyRevision) GetSubmissionDateOk() (*string, bool) {
	if o == nil || o.SubmissionDate == nil {
		return nil, false
	}
	return o.SubmissionDate, true
}

// HasSubmissionDate returns a boolean if a field has been set.
func (o *V2reportsAssemblyRevision) HasSubmissionDate() bool {
	if o != nil && o.SubmissionDate != nil {
		return true
	}

	return false
}

// SetSubmissionDate gets a reference to the given string and assigns it to the SubmissionDate field.
func (o *V2reportsAssemblyRevision) SetSubmissionDate(v string) {
	o.SubmissionDate = &v
}

// GetSequencingTechnology returns the SequencingTechnology field value if set, zero value otherwise.
func (o *V2reportsAssemblyRevision) GetSequencingTechnology() string {
	if o == nil || o.SequencingTechnology == nil {
		var ret string
		return ret
	}
	return *o.SequencingTechnology
}

// GetSequencingTechnologyOk returns a tuple with the SequencingTechnology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyRevision) GetSequencingTechnologyOk() (*string, bool) {
	if o == nil || o.SequencingTechnology == nil {
		return nil, false
	}
	return o.SequencingTechnology, true
}

// HasSequencingTechnology returns a boolean if a field has been set.
func (o *V2reportsAssemblyRevision) HasSequencingTechnology() bool {
	if o != nil && o.SequencingTechnology != nil {
		return true
	}

	return false
}

// SetSequencingTechnology gets a reference to the given string and assigns it to the SequencingTechnology field.
func (o *V2reportsAssemblyRevision) SetSequencingTechnology(v string) {
	o.SequencingTechnology = &v
}

func (o V2reportsAssemblyRevision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GenbankAccession != nil {
		toSerialize["genbank_accession"] = o.GenbankAccession
	}
	if o.RefseqAccession != nil {
		toSerialize["refseq_accession"] = o.RefseqAccession
	}
	if o.AssemblyName != nil {
		toSerialize["assembly_name"] = o.AssemblyName
	}
	if o.AssemblyLevel != nil {
		toSerialize["assembly_level"] = o.AssemblyLevel
	}
	if o.ReleaseDate != nil {
		toSerialize["release_date"] = o.ReleaseDate
	}
	if o.SubmissionDate != nil {
		toSerialize["submission_date"] = o.SubmissionDate
	}
	if o.SequencingTechnology != nil {
		toSerialize["sequencing_technology"] = o.SequencingTechnology
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsAssemblyRevision struct {
	value *V2reportsAssemblyRevision
	isSet bool
}

func (v NullableV2reportsAssemblyRevision) Get() *V2reportsAssemblyRevision {
	return v.value
}

func (v *NullableV2reportsAssemblyRevision) Set(val *V2reportsAssemblyRevision) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsAssemblyRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsAssemblyRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsAssemblyRevision(val *V2reportsAssemblyRevision) *NullableV2reportsAssemblyRevision {
	return &NullableV2reportsAssemblyRevision{value: val, isSet: true}
}

func (v NullableV2reportsAssemblyRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsAssemblyRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


