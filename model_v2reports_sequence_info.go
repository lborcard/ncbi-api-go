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

// V2reportsSequenceInfo struct for V2reportsSequenceInfo
type V2reportsSequenceInfo struct {
	AssemblyAccession *string `json:"assembly_accession,omitempty"`
	ChrName *string `json:"chr_name,omitempty"`
	UcscStyleName *string `json:"ucsc_style_name,omitempty"`
	SortOrder *int32 `json:"sort_order,omitempty"`
	AssignedMoleculeLocationType *string `json:"assigned_molecule_location_type,omitempty"`
	RefseqAccession *string `json:"refseq_accession,omitempty"`
	AssemblyUnit *string `json:"assembly_unit,omitempty"`
	Length *int32 `json:"length,omitempty"`
	GenbankAccession *string `json:"genbank_accession,omitempty"`
	GcCount *string `json:"gc_count,omitempty"`
	GcPercent *float32 `json:"gc_percent,omitempty"`
	UnlocalizedCount *int32 `json:"unlocalized_count,omitempty"`
	AssemblyUnplacedCount *int32 `json:"assembly_unplaced_count,omitempty"`
	Role *string `json:"role,omitempty"`
	SequenceName *string `json:"sequence_name,omitempty"`
}

// NewV2reportsSequenceInfo instantiates a new V2reportsSequenceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsSequenceInfo() *V2reportsSequenceInfo {
	this := V2reportsSequenceInfo{}
	return &this
}

// NewV2reportsSequenceInfoWithDefaults instantiates a new V2reportsSequenceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsSequenceInfoWithDefaults() *V2reportsSequenceInfo {
	this := V2reportsSequenceInfo{}
	return &this
}

// GetAssemblyAccession returns the AssemblyAccession field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetAssemblyAccession() string {
	if o == nil || o.AssemblyAccession == nil {
		var ret string
		return ret
	}
	return *o.AssemblyAccession
}

// GetAssemblyAccessionOk returns a tuple with the AssemblyAccession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetAssemblyAccessionOk() (*string, bool) {
	if o == nil || o.AssemblyAccession == nil {
		return nil, false
	}
	return o.AssemblyAccession, true
}

// HasAssemblyAccession returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasAssemblyAccession() bool {
	if o != nil && o.AssemblyAccession != nil {
		return true
	}

	return false
}

// SetAssemblyAccession gets a reference to the given string and assigns it to the AssemblyAccession field.
func (o *V2reportsSequenceInfo) SetAssemblyAccession(v string) {
	o.AssemblyAccession = &v
}

// GetChrName returns the ChrName field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetChrName() string {
	if o == nil || o.ChrName == nil {
		var ret string
		return ret
	}
	return *o.ChrName
}

// GetChrNameOk returns a tuple with the ChrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetChrNameOk() (*string, bool) {
	if o == nil || o.ChrName == nil {
		return nil, false
	}
	return o.ChrName, true
}

// HasChrName returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasChrName() bool {
	if o != nil && o.ChrName != nil {
		return true
	}

	return false
}

// SetChrName gets a reference to the given string and assigns it to the ChrName field.
func (o *V2reportsSequenceInfo) SetChrName(v string) {
	o.ChrName = &v
}

// GetUcscStyleName returns the UcscStyleName field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetUcscStyleName() string {
	if o == nil || o.UcscStyleName == nil {
		var ret string
		return ret
	}
	return *o.UcscStyleName
}

// GetUcscStyleNameOk returns a tuple with the UcscStyleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetUcscStyleNameOk() (*string, bool) {
	if o == nil || o.UcscStyleName == nil {
		return nil, false
	}
	return o.UcscStyleName, true
}

// HasUcscStyleName returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasUcscStyleName() bool {
	if o != nil && o.UcscStyleName != nil {
		return true
	}

	return false
}

// SetUcscStyleName gets a reference to the given string and assigns it to the UcscStyleName field.
func (o *V2reportsSequenceInfo) SetUcscStyleName(v string) {
	o.UcscStyleName = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetSortOrder() int32 {
	if o == nil || o.SortOrder == nil {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetSortOrderOk() (*int32, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *V2reportsSequenceInfo) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetAssignedMoleculeLocationType returns the AssignedMoleculeLocationType field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetAssignedMoleculeLocationType() string {
	if o == nil || o.AssignedMoleculeLocationType == nil {
		var ret string
		return ret
	}
	return *o.AssignedMoleculeLocationType
}

// GetAssignedMoleculeLocationTypeOk returns a tuple with the AssignedMoleculeLocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetAssignedMoleculeLocationTypeOk() (*string, bool) {
	if o == nil || o.AssignedMoleculeLocationType == nil {
		return nil, false
	}
	return o.AssignedMoleculeLocationType, true
}

// HasAssignedMoleculeLocationType returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasAssignedMoleculeLocationType() bool {
	if o != nil && o.AssignedMoleculeLocationType != nil {
		return true
	}

	return false
}

// SetAssignedMoleculeLocationType gets a reference to the given string and assigns it to the AssignedMoleculeLocationType field.
func (o *V2reportsSequenceInfo) SetAssignedMoleculeLocationType(v string) {
	o.AssignedMoleculeLocationType = &v
}

// GetRefseqAccession returns the RefseqAccession field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetRefseqAccession() string {
	if o == nil || o.RefseqAccession == nil {
		var ret string
		return ret
	}
	return *o.RefseqAccession
}

// GetRefseqAccessionOk returns a tuple with the RefseqAccession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetRefseqAccessionOk() (*string, bool) {
	if o == nil || o.RefseqAccession == nil {
		return nil, false
	}
	return o.RefseqAccession, true
}

// HasRefseqAccession returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasRefseqAccession() bool {
	if o != nil && o.RefseqAccession != nil {
		return true
	}

	return false
}

// SetRefseqAccession gets a reference to the given string and assigns it to the RefseqAccession field.
func (o *V2reportsSequenceInfo) SetRefseqAccession(v string) {
	o.RefseqAccession = &v
}

// GetAssemblyUnit returns the AssemblyUnit field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetAssemblyUnit() string {
	if o == nil || o.AssemblyUnit == nil {
		var ret string
		return ret
	}
	return *o.AssemblyUnit
}

// GetAssemblyUnitOk returns a tuple with the AssemblyUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetAssemblyUnitOk() (*string, bool) {
	if o == nil || o.AssemblyUnit == nil {
		return nil, false
	}
	return o.AssemblyUnit, true
}

// HasAssemblyUnit returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasAssemblyUnit() bool {
	if o != nil && o.AssemblyUnit != nil {
		return true
	}

	return false
}

// SetAssemblyUnit gets a reference to the given string and assigns it to the AssemblyUnit field.
func (o *V2reportsSequenceInfo) SetAssemblyUnit(v string) {
	o.AssemblyUnit = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetLength() int32 {
	if o == nil || o.Length == nil {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetLengthOk() (*int32, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *V2reportsSequenceInfo) SetLength(v int32) {
	o.Length = &v
}

// GetGenbankAccession returns the GenbankAccession field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetGenbankAccession() string {
	if o == nil || o.GenbankAccession == nil {
		var ret string
		return ret
	}
	return *o.GenbankAccession
}

// GetGenbankAccessionOk returns a tuple with the GenbankAccession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetGenbankAccessionOk() (*string, bool) {
	if o == nil || o.GenbankAccession == nil {
		return nil, false
	}
	return o.GenbankAccession, true
}

// HasGenbankAccession returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasGenbankAccession() bool {
	if o != nil && o.GenbankAccession != nil {
		return true
	}

	return false
}

// SetGenbankAccession gets a reference to the given string and assigns it to the GenbankAccession field.
func (o *V2reportsSequenceInfo) SetGenbankAccession(v string) {
	o.GenbankAccession = &v
}

// GetGcCount returns the GcCount field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetGcCount() string {
	if o == nil || o.GcCount == nil {
		var ret string
		return ret
	}
	return *o.GcCount
}

// GetGcCountOk returns a tuple with the GcCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetGcCountOk() (*string, bool) {
	if o == nil || o.GcCount == nil {
		return nil, false
	}
	return o.GcCount, true
}

// HasGcCount returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasGcCount() bool {
	if o != nil && o.GcCount != nil {
		return true
	}

	return false
}

// SetGcCount gets a reference to the given string and assigns it to the GcCount field.
func (o *V2reportsSequenceInfo) SetGcCount(v string) {
	o.GcCount = &v
}

// GetGcPercent returns the GcPercent field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetGcPercent() float32 {
	if o == nil || o.GcPercent == nil {
		var ret float32
		return ret
	}
	return *o.GcPercent
}

// GetGcPercentOk returns a tuple with the GcPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetGcPercentOk() (*float32, bool) {
	if o == nil || o.GcPercent == nil {
		return nil, false
	}
	return o.GcPercent, true
}

// HasGcPercent returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasGcPercent() bool {
	if o != nil && o.GcPercent != nil {
		return true
	}

	return false
}

// SetGcPercent gets a reference to the given float32 and assigns it to the GcPercent field.
func (o *V2reportsSequenceInfo) SetGcPercent(v float32) {
	o.GcPercent = &v
}

// GetUnlocalizedCount returns the UnlocalizedCount field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetUnlocalizedCount() int32 {
	if o == nil || o.UnlocalizedCount == nil {
		var ret int32
		return ret
	}
	return *o.UnlocalizedCount
}

// GetUnlocalizedCountOk returns a tuple with the UnlocalizedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetUnlocalizedCountOk() (*int32, bool) {
	if o == nil || o.UnlocalizedCount == nil {
		return nil, false
	}
	return o.UnlocalizedCount, true
}

// HasUnlocalizedCount returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasUnlocalizedCount() bool {
	if o != nil && o.UnlocalizedCount != nil {
		return true
	}

	return false
}

// SetUnlocalizedCount gets a reference to the given int32 and assigns it to the UnlocalizedCount field.
func (o *V2reportsSequenceInfo) SetUnlocalizedCount(v int32) {
	o.UnlocalizedCount = &v
}

// GetAssemblyUnplacedCount returns the AssemblyUnplacedCount field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetAssemblyUnplacedCount() int32 {
	if o == nil || o.AssemblyUnplacedCount == nil {
		var ret int32
		return ret
	}
	return *o.AssemblyUnplacedCount
}

// GetAssemblyUnplacedCountOk returns a tuple with the AssemblyUnplacedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetAssemblyUnplacedCountOk() (*int32, bool) {
	if o == nil || o.AssemblyUnplacedCount == nil {
		return nil, false
	}
	return o.AssemblyUnplacedCount, true
}

// HasAssemblyUnplacedCount returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasAssemblyUnplacedCount() bool {
	if o != nil && o.AssemblyUnplacedCount != nil {
		return true
	}

	return false
}

// SetAssemblyUnplacedCount gets a reference to the given int32 and assigns it to the AssemblyUnplacedCount field.
func (o *V2reportsSequenceInfo) SetAssemblyUnplacedCount(v int32) {
	o.AssemblyUnplacedCount = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *V2reportsSequenceInfo) SetRole(v string) {
	o.Role = &v
}

// GetSequenceName returns the SequenceName field value if set, zero value otherwise.
func (o *V2reportsSequenceInfo) GetSequenceName() string {
	if o == nil || o.SequenceName == nil {
		var ret string
		return ret
	}
	return *o.SequenceName
}

// GetSequenceNameOk returns a tuple with the SequenceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsSequenceInfo) GetSequenceNameOk() (*string, bool) {
	if o == nil || o.SequenceName == nil {
		return nil, false
	}
	return o.SequenceName, true
}

// HasSequenceName returns a boolean if a field has been set.
func (o *V2reportsSequenceInfo) HasSequenceName() bool {
	if o != nil && o.SequenceName != nil {
		return true
	}

	return false
}

// SetSequenceName gets a reference to the given string and assigns it to the SequenceName field.
func (o *V2reportsSequenceInfo) SetSequenceName(v string) {
	o.SequenceName = &v
}

func (o V2reportsSequenceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssemblyAccession != nil {
		toSerialize["assembly_accession"] = o.AssemblyAccession
	}
	if o.ChrName != nil {
		toSerialize["chr_name"] = o.ChrName
	}
	if o.UcscStyleName != nil {
		toSerialize["ucsc_style_name"] = o.UcscStyleName
	}
	if o.SortOrder != nil {
		toSerialize["sort_order"] = o.SortOrder
	}
	if o.AssignedMoleculeLocationType != nil {
		toSerialize["assigned_molecule_location_type"] = o.AssignedMoleculeLocationType
	}
	if o.RefseqAccession != nil {
		toSerialize["refseq_accession"] = o.RefseqAccession
	}
	if o.AssemblyUnit != nil {
		toSerialize["assembly_unit"] = o.AssemblyUnit
	}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.GenbankAccession != nil {
		toSerialize["genbank_accession"] = o.GenbankAccession
	}
	if o.GcCount != nil {
		toSerialize["gc_count"] = o.GcCount
	}
	if o.GcPercent != nil {
		toSerialize["gc_percent"] = o.GcPercent
	}
	if o.UnlocalizedCount != nil {
		toSerialize["unlocalized_count"] = o.UnlocalizedCount
	}
	if o.AssemblyUnplacedCount != nil {
		toSerialize["assembly_unplaced_count"] = o.AssemblyUnplacedCount
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.SequenceName != nil {
		toSerialize["sequence_name"] = o.SequenceName
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsSequenceInfo struct {
	value *V2reportsSequenceInfo
	isSet bool
}

func (v NullableV2reportsSequenceInfo) Get() *V2reportsSequenceInfo {
	return v.value
}

func (v *NullableV2reportsSequenceInfo) Set(val *V2reportsSequenceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsSequenceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsSequenceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsSequenceInfo(val *V2reportsSequenceInfo) *NullableV2reportsSequenceInfo {
	return &NullableV2reportsSequenceInfo{value: val, isSet: true}
}

func (v NullableV2reportsSequenceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsSequenceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


