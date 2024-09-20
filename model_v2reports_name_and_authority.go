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

// V2reportsNameAndAuthority struct for V2reportsNameAndAuthority
type V2reportsNameAndAuthority struct {
	Name *string `json:"name,omitempty"`
	Authority *string `json:"authority,omitempty"`
	TypeStrains []V2reportsTaxonomyTypeMaterial `json:"type_strains,omitempty"`
	CuratorSynonym *string `json:"curator_synonym,omitempty"`
	HomotypicSynonyms []V2reportsNameAndAuthority `json:"homotypic_synonyms,omitempty"`
	HeterotypicSynonyms []V2reportsNameAndAuthority `json:"heterotypic_synonyms,omitempty"`
	OtherSynonyms []V2reportsNameAndAuthority `json:"other_synonyms,omitempty"`
	InformalNames []string `json:"informal_names,omitempty"`
	Basionym *V2reportsNameAndAuthority `json:"basionym,omitempty"`
	Publications []V2reportsNameAndAuthorityPublication `json:"publications,omitempty"`
	Notes []V2reportsNameAndAuthorityNote `json:"notes,omitempty"`
	Formal *bool `json:"formal,omitempty"`
}

// NewV2reportsNameAndAuthority instantiates a new V2reportsNameAndAuthority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsNameAndAuthority() *V2reportsNameAndAuthority {
	this := V2reportsNameAndAuthority{}
	return &this
}

// NewV2reportsNameAndAuthorityWithDefaults instantiates a new V2reportsNameAndAuthority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsNameAndAuthorityWithDefaults() *V2reportsNameAndAuthority {
	this := V2reportsNameAndAuthority{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthority) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthority) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthority) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V2reportsNameAndAuthority) SetName(v string) {
	o.Name = &v
}

// GetAuthority returns the Authority field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthority) GetAuthority() string {
	if o == nil || o.Authority == nil {
		var ret string
		return ret
	}
	return *o.Authority
}

// GetAuthorityOk returns a tuple with the Authority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthority) GetAuthorityOk() (*string, bool) {
	if o == nil || o.Authority == nil {
		return nil, false
	}
	return o.Authority, true
}

// HasAuthority returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthority) HasAuthority() bool {
	if o != nil && o.Authority != nil {
		return true
	}

	return false
}

// SetAuthority gets a reference to the given string and assigns it to the Authority field.
func (o *V2reportsNameAndAuthority) SetAuthority(v string) {
	o.Authority = &v
}

// GetTypeStrains returns the TypeStrains field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthority) GetTypeStrains() []V2reportsTaxonomyTypeMaterial {
	if o == nil || o.TypeStrains == nil {
		var ret []V2reportsTaxonomyTypeMaterial
		return ret
	}
	return o.TypeStrains
}

// GetTypeStrainsOk returns a tuple with the TypeStrains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthority) GetTypeStrainsOk() ([]V2reportsTaxonomyTypeMaterial, bool) {
	if o == nil || o.TypeStrains == nil {
		return nil, false
	}
	return o.TypeStrains, true
}

// HasTypeStrains returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthority) HasTypeStrains() bool {
	if o != nil && o.TypeStrains != nil {
		return true
	}

	return false
}

// SetTypeStrains gets a reference to the given []V2reportsTaxonomyTypeMaterial and assigns it to the TypeStrains field.
func (o *V2reportsNameAndAuthority) SetTypeStrains(v []V2reportsTaxonomyTypeMaterial) {
	o.TypeStrains = v
}

// GetCuratorSynonym returns the CuratorSynonym field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthority) GetCuratorSynonym() string {
	if o == nil || o.CuratorSynonym == nil {
		var ret string
		return ret
	}
	return *o.CuratorSynonym
}

// GetCuratorSynonymOk returns a tuple with the CuratorSynonym field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthority) GetCuratorSynonymOk() (*string, bool) {
	if o == nil || o.CuratorSynonym == nil {
		return nil, false
	}
	return o.CuratorSynonym, true
}

// HasCuratorSynonym returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthority) HasCuratorSynonym() bool {
	if o != nil && o.CuratorSynonym != nil {
		return true
	}

	return false
}

// SetCuratorSynonym gets a reference to the given string and assigns it to the CuratorSynonym field.
func (o *V2reportsNameAndAuthority) SetCuratorSynonym(v string) {
	o.CuratorSynonym = &v
}

// GetHomotypicSynonyms returns the HomotypicSynonyms field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthority) GetHomotypicSynonyms() []V2reportsNameAndAuthority {
	if o == nil || o.HomotypicSynonyms == nil {
		var ret []V2reportsNameAndAuthority
		return ret
	}
	return o.HomotypicSynonyms
}

// GetHomotypicSynonymsOk returns a tuple with the HomotypicSynonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthority) GetHomotypicSynonymsOk() ([]V2reportsNameAndAuthority, bool) {
	if o == nil || o.HomotypicSynonyms == nil {
		return nil, false
	}
	return o.HomotypicSynonyms, true
}

// HasHomotypicSynonyms returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthority) HasHomotypicSynonyms() bool {
	if o != nil && o.HomotypicSynonyms != nil {
		return true
	}

	return false
}

// SetHomotypicSynonyms gets a reference to the given []V2reportsNameAndAuthority and assigns it to the HomotypicSynonyms field.
func (o *V2reportsNameAndAuthority) SetHomotypicSynonyms(v []V2reportsNameAndAuthority) {
	o.HomotypicSynonyms = v
}

// GetHeterotypicSynonyms returns the HeterotypicSynonyms field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthority) GetHeterotypicSynonyms() []V2reportsNameAndAuthority {
	if o == nil || o.HeterotypicSynonyms == nil {
		var ret []V2reportsNameAndAuthority
		return ret
	}
	return o.HeterotypicSynonyms
}

// GetHeterotypicSynonymsOk returns a tuple with the HeterotypicSynonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthority) GetHeterotypicSynonymsOk() ([]V2reportsNameAndAuthority, bool) {
	if o == nil || o.HeterotypicSynonyms == nil {
		return nil, false
	}
	return o.HeterotypicSynonyms, true
}

// HasHeterotypicSynonyms returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthority) HasHeterotypicSynonyms() bool {
	if o != nil && o.HeterotypicSynonyms != nil {
		return true
	}

	return false
}

// SetHeterotypicSynonyms gets a reference to the given []V2reportsNameAndAuthority and assigns it to the HeterotypicSynonyms field.
func (o *V2reportsNameAndAuthority) SetHeterotypicSynonyms(v []V2reportsNameAndAuthority) {
	o.HeterotypicSynonyms = v
}

// GetOtherSynonyms returns the OtherSynonyms field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthority) GetOtherSynonyms() []V2reportsNameAndAuthority {
	if o == nil || o.OtherSynonyms == nil {
		var ret []V2reportsNameAndAuthority
		return ret
	}
	return o.OtherSynonyms
}

// GetOtherSynonymsOk returns a tuple with the OtherSynonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthority) GetOtherSynonymsOk() ([]V2reportsNameAndAuthority, bool) {
	if o == nil || o.OtherSynonyms == nil {
		return nil, false
	}
	return o.OtherSynonyms, true
}

// HasOtherSynonyms returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthority) HasOtherSynonyms() bool {
	if o != nil && o.OtherSynonyms != nil {
		return true
	}

	return false
}

// SetOtherSynonyms gets a reference to the given []V2reportsNameAndAuthority and assigns it to the OtherSynonyms field.
func (o *V2reportsNameAndAuthority) SetOtherSynonyms(v []V2reportsNameAndAuthority) {
	o.OtherSynonyms = v
}

// GetInformalNames returns the InformalNames field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthority) GetInformalNames() []string {
	if o == nil || o.InformalNames == nil {
		var ret []string
		return ret
	}
	return o.InformalNames
}

// GetInformalNamesOk returns a tuple with the InformalNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthority) GetInformalNamesOk() ([]string, bool) {
	if o == nil || o.InformalNames == nil {
		return nil, false
	}
	return o.InformalNames, true
}

// HasInformalNames returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthority) HasInformalNames() bool {
	if o != nil && o.InformalNames != nil {
		return true
	}

	return false
}

// SetInformalNames gets a reference to the given []string and assigns it to the InformalNames field.
func (o *V2reportsNameAndAuthority) SetInformalNames(v []string) {
	o.InformalNames = v
}

// GetBasionym returns the Basionym field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthority) GetBasionym() V2reportsNameAndAuthority {
	if o == nil || o.Basionym == nil {
		var ret V2reportsNameAndAuthority
		return ret
	}
	return *o.Basionym
}

// GetBasionymOk returns a tuple with the Basionym field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthority) GetBasionymOk() (*V2reportsNameAndAuthority, bool) {
	if o == nil || o.Basionym == nil {
		return nil, false
	}
	return o.Basionym, true
}

// HasBasionym returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthority) HasBasionym() bool {
	if o != nil && o.Basionym != nil {
		return true
	}

	return false
}

// SetBasionym gets a reference to the given V2reportsNameAndAuthority and assigns it to the Basionym field.
func (o *V2reportsNameAndAuthority) SetBasionym(v V2reportsNameAndAuthority) {
	o.Basionym = &v
}

// GetPublications returns the Publications field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthority) GetPublications() []V2reportsNameAndAuthorityPublication {
	if o == nil || o.Publications == nil {
		var ret []V2reportsNameAndAuthorityPublication
		return ret
	}
	return o.Publications
}

// GetPublicationsOk returns a tuple with the Publications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthority) GetPublicationsOk() ([]V2reportsNameAndAuthorityPublication, bool) {
	if o == nil || o.Publications == nil {
		return nil, false
	}
	return o.Publications, true
}

// HasPublications returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthority) HasPublications() bool {
	if o != nil && o.Publications != nil {
		return true
	}

	return false
}

// SetPublications gets a reference to the given []V2reportsNameAndAuthorityPublication and assigns it to the Publications field.
func (o *V2reportsNameAndAuthority) SetPublications(v []V2reportsNameAndAuthorityPublication) {
	o.Publications = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthority) GetNotes() []V2reportsNameAndAuthorityNote {
	if o == nil || o.Notes == nil {
		var ret []V2reportsNameAndAuthorityNote
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthority) GetNotesOk() ([]V2reportsNameAndAuthorityNote, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthority) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []V2reportsNameAndAuthorityNote and assigns it to the Notes field.
func (o *V2reportsNameAndAuthority) SetNotes(v []V2reportsNameAndAuthorityNote) {
	o.Notes = v
}

// GetFormal returns the Formal field value if set, zero value otherwise.
func (o *V2reportsNameAndAuthority) GetFormal() bool {
	if o == nil || o.Formal == nil {
		var ret bool
		return ret
	}
	return *o.Formal
}

// GetFormalOk returns a tuple with the Formal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsNameAndAuthority) GetFormalOk() (*bool, bool) {
	if o == nil || o.Formal == nil {
		return nil, false
	}
	return o.Formal, true
}

// HasFormal returns a boolean if a field has been set.
func (o *V2reportsNameAndAuthority) HasFormal() bool {
	if o != nil && o.Formal != nil {
		return true
	}

	return false
}

// SetFormal gets a reference to the given bool and assigns it to the Formal field.
func (o *V2reportsNameAndAuthority) SetFormal(v bool) {
	o.Formal = &v
}

func (o V2reportsNameAndAuthority) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Authority != nil {
		toSerialize["authority"] = o.Authority
	}
	if o.TypeStrains != nil {
		toSerialize["type_strains"] = o.TypeStrains
	}
	if o.CuratorSynonym != nil {
		toSerialize["curator_synonym"] = o.CuratorSynonym
	}
	if o.HomotypicSynonyms != nil {
		toSerialize["homotypic_synonyms"] = o.HomotypicSynonyms
	}
	if o.HeterotypicSynonyms != nil {
		toSerialize["heterotypic_synonyms"] = o.HeterotypicSynonyms
	}
	if o.OtherSynonyms != nil {
		toSerialize["other_synonyms"] = o.OtherSynonyms
	}
	if o.InformalNames != nil {
		toSerialize["informal_names"] = o.InformalNames
	}
	if o.Basionym != nil {
		toSerialize["basionym"] = o.Basionym
	}
	if o.Publications != nil {
		toSerialize["publications"] = o.Publications
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Formal != nil {
		toSerialize["formal"] = o.Formal
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsNameAndAuthority struct {
	value *V2reportsNameAndAuthority
	isSet bool
}

func (v NullableV2reportsNameAndAuthority) Get() *V2reportsNameAndAuthority {
	return v.value
}

func (v *NullableV2reportsNameAndAuthority) Set(val *V2reportsNameAndAuthority) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsNameAndAuthority) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsNameAndAuthority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsNameAndAuthority(val *V2reportsNameAndAuthority) *NullableV2reportsNameAndAuthority {
	return &NullableV2reportsNameAndAuthority{value: val, isSet: true}
}

func (v NullableV2reportsNameAndAuthority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsNameAndAuthority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


