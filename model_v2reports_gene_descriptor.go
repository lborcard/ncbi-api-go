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

// V2reportsGeneDescriptor struct for V2reportsGeneDescriptor
type V2reportsGeneDescriptor struct {
	GeneId *string `json:"gene_id,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Description *string `json:"description,omitempty"`
	TaxId *string `json:"tax_id,omitempty"`
	Taxname *string `json:"taxname,omitempty"`
	CommonName *string `json:"common_name,omitempty"`
	Type *V2reportsGeneType `json:"type,omitempty"`
	RnaType *V2reportsRnaType `json:"rna_type,omitempty"`
	Orientation *V2reportsOrientation `json:"orientation,omitempty"`
	ReferenceStandards []V2reportsGenomicRegion `json:"reference_standards,omitempty"`
	GenomicRegions []V2reportsGenomicRegion `json:"genomic_regions,omitempty"`
	Chromosomes []string `json:"chromosomes,omitempty"`
	NomenclatureAuthority *V2reportsNomenclatureAuthority `json:"nomenclature_authority,omitempty"`
	SwissProtAccessions []string `json:"swiss_prot_accessions,omitempty"`
	EnsemblGeneIds []string `json:"ensembl_gene_ids,omitempty"`
	OmimIds []string `json:"omim_ids,omitempty"`
	Synonyms []string `json:"synonyms,omitempty"`
	ReplacedGeneId *string `json:"replaced_gene_id,omitempty"`
	Annotations []V2reportsAnnotation `json:"annotations,omitempty"`
	TranscriptCount *int32 `json:"transcript_count,omitempty"`
	ProteinCount *int32 `json:"protein_count,omitempty"`
	TranscriptTypeCounts []V2reportsTranscriptTypeCount `json:"transcript_type_counts,omitempty"`
	GeneGroups []V2reportsGeneGroup `json:"gene_groups,omitempty"`
}

// NewV2reportsGeneDescriptor instantiates a new V2reportsGeneDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsGeneDescriptor() *V2reportsGeneDescriptor {
	this := V2reportsGeneDescriptor{}
	var type_ V2reportsGeneType = V2REPORTSGENETYPE_UNKNOWN
	this.Type = &type_
	var rnaType V2reportsRnaType = V2REPORTSRNATYPE_RNA_UNKNOWN
	this.RnaType = &rnaType
	var orientation V2reportsOrientation = V2REPORTSORIENTATION_NONE
	this.Orientation = &orientation
	return &this
}

// NewV2reportsGeneDescriptorWithDefaults instantiates a new V2reportsGeneDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsGeneDescriptorWithDefaults() *V2reportsGeneDescriptor {
	this := V2reportsGeneDescriptor{}
	var type_ V2reportsGeneType = V2REPORTSGENETYPE_UNKNOWN
	this.Type = &type_
	var rnaType V2reportsRnaType = V2REPORTSRNATYPE_RNA_UNKNOWN
	this.RnaType = &rnaType
	var orientation V2reportsOrientation = V2REPORTSORIENTATION_NONE
	this.Orientation = &orientation
	return &this
}

// GetGeneId returns the GeneId field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetGeneId() string {
	if o == nil || o.GeneId == nil {
		var ret string
		return ret
	}
	return *o.GeneId
}

// GetGeneIdOk returns a tuple with the GeneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetGeneIdOk() (*string, bool) {
	if o == nil || o.GeneId == nil {
		return nil, false
	}
	return o.GeneId, true
}

// HasGeneId returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasGeneId() bool {
	if o != nil && o.GeneId != nil {
		return true
	}

	return false
}

// SetGeneId gets a reference to the given string and assigns it to the GeneId field.
func (o *V2reportsGeneDescriptor) SetGeneId(v string) {
	o.GeneId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *V2reportsGeneDescriptor) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V2reportsGeneDescriptor) SetDescription(v string) {
	o.Description = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetTaxId() string {
	if o == nil || o.TaxId == nil {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetTaxIdOk() (*string, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *V2reportsGeneDescriptor) SetTaxId(v string) {
	o.TaxId = &v
}

// GetTaxname returns the Taxname field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetTaxname() string {
	if o == nil || o.Taxname == nil {
		var ret string
		return ret
	}
	return *o.Taxname
}

// GetTaxnameOk returns a tuple with the Taxname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetTaxnameOk() (*string, bool) {
	if o == nil || o.Taxname == nil {
		return nil, false
	}
	return o.Taxname, true
}

// HasTaxname returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasTaxname() bool {
	if o != nil && o.Taxname != nil {
		return true
	}

	return false
}

// SetTaxname gets a reference to the given string and assigns it to the Taxname field.
func (o *V2reportsGeneDescriptor) SetTaxname(v string) {
	o.Taxname = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *V2reportsGeneDescriptor) SetCommonName(v string) {
	o.CommonName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetType() V2reportsGeneType {
	if o == nil || o.Type == nil {
		var ret V2reportsGeneType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetTypeOk() (*V2reportsGeneType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given V2reportsGeneType and assigns it to the Type field.
func (o *V2reportsGeneDescriptor) SetType(v V2reportsGeneType) {
	o.Type = &v
}

// GetRnaType returns the RnaType field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetRnaType() V2reportsRnaType {
	if o == nil || o.RnaType == nil {
		var ret V2reportsRnaType
		return ret
	}
	return *o.RnaType
}

// GetRnaTypeOk returns a tuple with the RnaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetRnaTypeOk() (*V2reportsRnaType, bool) {
	if o == nil || o.RnaType == nil {
		return nil, false
	}
	return o.RnaType, true
}

// HasRnaType returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasRnaType() bool {
	if o != nil && o.RnaType != nil {
		return true
	}

	return false
}

// SetRnaType gets a reference to the given V2reportsRnaType and assigns it to the RnaType field.
func (o *V2reportsGeneDescriptor) SetRnaType(v V2reportsRnaType) {
	o.RnaType = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetOrientation() V2reportsOrientation {
	if o == nil || o.Orientation == nil {
		var ret V2reportsOrientation
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetOrientationOk() (*V2reportsOrientation, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given V2reportsOrientation and assigns it to the Orientation field.
func (o *V2reportsGeneDescriptor) SetOrientation(v V2reportsOrientation) {
	o.Orientation = &v
}

// GetReferenceStandards returns the ReferenceStandards field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetReferenceStandards() []V2reportsGenomicRegion {
	if o == nil || o.ReferenceStandards == nil {
		var ret []V2reportsGenomicRegion
		return ret
	}
	return o.ReferenceStandards
}

// GetReferenceStandardsOk returns a tuple with the ReferenceStandards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetReferenceStandardsOk() ([]V2reportsGenomicRegion, bool) {
	if o == nil || o.ReferenceStandards == nil {
		return nil, false
	}
	return o.ReferenceStandards, true
}

// HasReferenceStandards returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasReferenceStandards() bool {
	if o != nil && o.ReferenceStandards != nil {
		return true
	}

	return false
}

// SetReferenceStandards gets a reference to the given []V2reportsGenomicRegion and assigns it to the ReferenceStandards field.
func (o *V2reportsGeneDescriptor) SetReferenceStandards(v []V2reportsGenomicRegion) {
	o.ReferenceStandards = v
}

// GetGenomicRegions returns the GenomicRegions field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetGenomicRegions() []V2reportsGenomicRegion {
	if o == nil || o.GenomicRegions == nil {
		var ret []V2reportsGenomicRegion
		return ret
	}
	return o.GenomicRegions
}

// GetGenomicRegionsOk returns a tuple with the GenomicRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetGenomicRegionsOk() ([]V2reportsGenomicRegion, bool) {
	if o == nil || o.GenomicRegions == nil {
		return nil, false
	}
	return o.GenomicRegions, true
}

// HasGenomicRegions returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasGenomicRegions() bool {
	if o != nil && o.GenomicRegions != nil {
		return true
	}

	return false
}

// SetGenomicRegions gets a reference to the given []V2reportsGenomicRegion and assigns it to the GenomicRegions field.
func (o *V2reportsGeneDescriptor) SetGenomicRegions(v []V2reportsGenomicRegion) {
	o.GenomicRegions = v
}

// GetChromosomes returns the Chromosomes field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetChromosomes() []string {
	if o == nil || o.Chromosomes == nil {
		var ret []string
		return ret
	}
	return o.Chromosomes
}

// GetChromosomesOk returns a tuple with the Chromosomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetChromosomesOk() ([]string, bool) {
	if o == nil || o.Chromosomes == nil {
		return nil, false
	}
	return o.Chromosomes, true
}

// HasChromosomes returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasChromosomes() bool {
	if o != nil && o.Chromosomes != nil {
		return true
	}

	return false
}

// SetChromosomes gets a reference to the given []string and assigns it to the Chromosomes field.
func (o *V2reportsGeneDescriptor) SetChromosomes(v []string) {
	o.Chromosomes = v
}

// GetNomenclatureAuthority returns the NomenclatureAuthority field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetNomenclatureAuthority() V2reportsNomenclatureAuthority {
	if o == nil || o.NomenclatureAuthority == nil {
		var ret V2reportsNomenclatureAuthority
		return ret
	}
	return *o.NomenclatureAuthority
}

// GetNomenclatureAuthorityOk returns a tuple with the NomenclatureAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetNomenclatureAuthorityOk() (*V2reportsNomenclatureAuthority, bool) {
	if o == nil || o.NomenclatureAuthority == nil {
		return nil, false
	}
	return o.NomenclatureAuthority, true
}

// HasNomenclatureAuthority returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasNomenclatureAuthority() bool {
	if o != nil && o.NomenclatureAuthority != nil {
		return true
	}

	return false
}

// SetNomenclatureAuthority gets a reference to the given V2reportsNomenclatureAuthority and assigns it to the NomenclatureAuthority field.
func (o *V2reportsGeneDescriptor) SetNomenclatureAuthority(v V2reportsNomenclatureAuthority) {
	o.NomenclatureAuthority = &v
}

// GetSwissProtAccessions returns the SwissProtAccessions field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetSwissProtAccessions() []string {
	if o == nil || o.SwissProtAccessions == nil {
		var ret []string
		return ret
	}
	return o.SwissProtAccessions
}

// GetSwissProtAccessionsOk returns a tuple with the SwissProtAccessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetSwissProtAccessionsOk() ([]string, bool) {
	if o == nil || o.SwissProtAccessions == nil {
		return nil, false
	}
	return o.SwissProtAccessions, true
}

// HasSwissProtAccessions returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasSwissProtAccessions() bool {
	if o != nil && o.SwissProtAccessions != nil {
		return true
	}

	return false
}

// SetSwissProtAccessions gets a reference to the given []string and assigns it to the SwissProtAccessions field.
func (o *V2reportsGeneDescriptor) SetSwissProtAccessions(v []string) {
	o.SwissProtAccessions = v
}

// GetEnsemblGeneIds returns the EnsemblGeneIds field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetEnsemblGeneIds() []string {
	if o == nil || o.EnsemblGeneIds == nil {
		var ret []string
		return ret
	}
	return o.EnsemblGeneIds
}

// GetEnsemblGeneIdsOk returns a tuple with the EnsemblGeneIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetEnsemblGeneIdsOk() ([]string, bool) {
	if o == nil || o.EnsemblGeneIds == nil {
		return nil, false
	}
	return o.EnsemblGeneIds, true
}

// HasEnsemblGeneIds returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasEnsemblGeneIds() bool {
	if o != nil && o.EnsemblGeneIds != nil {
		return true
	}

	return false
}

// SetEnsemblGeneIds gets a reference to the given []string and assigns it to the EnsemblGeneIds field.
func (o *V2reportsGeneDescriptor) SetEnsemblGeneIds(v []string) {
	o.EnsemblGeneIds = v
}

// GetOmimIds returns the OmimIds field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetOmimIds() []string {
	if o == nil || o.OmimIds == nil {
		var ret []string
		return ret
	}
	return o.OmimIds
}

// GetOmimIdsOk returns a tuple with the OmimIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetOmimIdsOk() ([]string, bool) {
	if o == nil || o.OmimIds == nil {
		return nil, false
	}
	return o.OmimIds, true
}

// HasOmimIds returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasOmimIds() bool {
	if o != nil && o.OmimIds != nil {
		return true
	}

	return false
}

// SetOmimIds gets a reference to the given []string and assigns it to the OmimIds field.
func (o *V2reportsGeneDescriptor) SetOmimIds(v []string) {
	o.OmimIds = v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetSynonyms() []string {
	if o == nil || o.Synonyms == nil {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetSynonymsOk() ([]string, bool) {
	if o == nil || o.Synonyms == nil {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasSynonyms() bool {
	if o != nil && o.Synonyms != nil {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *V2reportsGeneDescriptor) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetReplacedGeneId returns the ReplacedGeneId field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetReplacedGeneId() string {
	if o == nil || o.ReplacedGeneId == nil {
		var ret string
		return ret
	}
	return *o.ReplacedGeneId
}

// GetReplacedGeneIdOk returns a tuple with the ReplacedGeneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetReplacedGeneIdOk() (*string, bool) {
	if o == nil || o.ReplacedGeneId == nil {
		return nil, false
	}
	return o.ReplacedGeneId, true
}

// HasReplacedGeneId returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasReplacedGeneId() bool {
	if o != nil && o.ReplacedGeneId != nil {
		return true
	}

	return false
}

// SetReplacedGeneId gets a reference to the given string and assigns it to the ReplacedGeneId field.
func (o *V2reportsGeneDescriptor) SetReplacedGeneId(v string) {
	o.ReplacedGeneId = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetAnnotations() []V2reportsAnnotation {
	if o == nil || o.Annotations == nil {
		var ret []V2reportsAnnotation
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetAnnotationsOk() ([]V2reportsAnnotation, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []V2reportsAnnotation and assigns it to the Annotations field.
func (o *V2reportsGeneDescriptor) SetAnnotations(v []V2reportsAnnotation) {
	o.Annotations = v
}

// GetTranscriptCount returns the TranscriptCount field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetTranscriptCount() int32 {
	if o == nil || o.TranscriptCount == nil {
		var ret int32
		return ret
	}
	return *o.TranscriptCount
}

// GetTranscriptCountOk returns a tuple with the TranscriptCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetTranscriptCountOk() (*int32, bool) {
	if o == nil || o.TranscriptCount == nil {
		return nil, false
	}
	return o.TranscriptCount, true
}

// HasTranscriptCount returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasTranscriptCount() bool {
	if o != nil && o.TranscriptCount != nil {
		return true
	}

	return false
}

// SetTranscriptCount gets a reference to the given int32 and assigns it to the TranscriptCount field.
func (o *V2reportsGeneDescriptor) SetTranscriptCount(v int32) {
	o.TranscriptCount = &v
}

// GetProteinCount returns the ProteinCount field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetProteinCount() int32 {
	if o == nil || o.ProteinCount == nil {
		var ret int32
		return ret
	}
	return *o.ProteinCount
}

// GetProteinCountOk returns a tuple with the ProteinCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetProteinCountOk() (*int32, bool) {
	if o == nil || o.ProteinCount == nil {
		return nil, false
	}
	return o.ProteinCount, true
}

// HasProteinCount returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasProteinCount() bool {
	if o != nil && o.ProteinCount != nil {
		return true
	}

	return false
}

// SetProteinCount gets a reference to the given int32 and assigns it to the ProteinCount field.
func (o *V2reportsGeneDescriptor) SetProteinCount(v int32) {
	o.ProteinCount = &v
}

// GetTranscriptTypeCounts returns the TranscriptTypeCounts field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetTranscriptTypeCounts() []V2reportsTranscriptTypeCount {
	if o == nil || o.TranscriptTypeCounts == nil {
		var ret []V2reportsTranscriptTypeCount
		return ret
	}
	return o.TranscriptTypeCounts
}

// GetTranscriptTypeCountsOk returns a tuple with the TranscriptTypeCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetTranscriptTypeCountsOk() ([]V2reportsTranscriptTypeCount, bool) {
	if o == nil || o.TranscriptTypeCounts == nil {
		return nil, false
	}
	return o.TranscriptTypeCounts, true
}

// HasTranscriptTypeCounts returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasTranscriptTypeCounts() bool {
	if o != nil && o.TranscriptTypeCounts != nil {
		return true
	}

	return false
}

// SetTranscriptTypeCounts gets a reference to the given []V2reportsTranscriptTypeCount and assigns it to the TranscriptTypeCounts field.
func (o *V2reportsGeneDescriptor) SetTranscriptTypeCounts(v []V2reportsTranscriptTypeCount) {
	o.TranscriptTypeCounts = v
}

// GetGeneGroups returns the GeneGroups field value if set, zero value otherwise.
func (o *V2reportsGeneDescriptor) GetGeneGroups() []V2reportsGeneGroup {
	if o == nil || o.GeneGroups == nil {
		var ret []V2reportsGeneGroup
		return ret
	}
	return o.GeneGroups
}

// GetGeneGroupsOk returns a tuple with the GeneGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGeneDescriptor) GetGeneGroupsOk() ([]V2reportsGeneGroup, bool) {
	if o == nil || o.GeneGroups == nil {
		return nil, false
	}
	return o.GeneGroups, true
}

// HasGeneGroups returns a boolean if a field has been set.
func (o *V2reportsGeneDescriptor) HasGeneGroups() bool {
	if o != nil && o.GeneGroups != nil {
		return true
	}

	return false
}

// SetGeneGroups gets a reference to the given []V2reportsGeneGroup and assigns it to the GeneGroups field.
func (o *V2reportsGeneDescriptor) SetGeneGroups(v []V2reportsGeneGroup) {
	o.GeneGroups = v
}

func (o V2reportsGeneDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GeneId != nil {
		toSerialize["gene_id"] = o.GeneId
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.TaxId != nil {
		toSerialize["tax_id"] = o.TaxId
	}
	if o.Taxname != nil {
		toSerialize["taxname"] = o.Taxname
	}
	if o.CommonName != nil {
		toSerialize["common_name"] = o.CommonName
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.RnaType != nil {
		toSerialize["rna_type"] = o.RnaType
	}
	if o.Orientation != nil {
		toSerialize["orientation"] = o.Orientation
	}
	if o.ReferenceStandards != nil {
		toSerialize["reference_standards"] = o.ReferenceStandards
	}
	if o.GenomicRegions != nil {
		toSerialize["genomic_regions"] = o.GenomicRegions
	}
	if o.Chromosomes != nil {
		toSerialize["chromosomes"] = o.Chromosomes
	}
	if o.NomenclatureAuthority != nil {
		toSerialize["nomenclature_authority"] = o.NomenclatureAuthority
	}
	if o.SwissProtAccessions != nil {
		toSerialize["swiss_prot_accessions"] = o.SwissProtAccessions
	}
	if o.EnsemblGeneIds != nil {
		toSerialize["ensembl_gene_ids"] = o.EnsemblGeneIds
	}
	if o.OmimIds != nil {
		toSerialize["omim_ids"] = o.OmimIds
	}
	if o.Synonyms != nil {
		toSerialize["synonyms"] = o.Synonyms
	}
	if o.ReplacedGeneId != nil {
		toSerialize["replaced_gene_id"] = o.ReplacedGeneId
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	if o.TranscriptCount != nil {
		toSerialize["transcript_count"] = o.TranscriptCount
	}
	if o.ProteinCount != nil {
		toSerialize["protein_count"] = o.ProteinCount
	}
	if o.TranscriptTypeCounts != nil {
		toSerialize["transcript_type_counts"] = o.TranscriptTypeCounts
	}
	if o.GeneGroups != nil {
		toSerialize["gene_groups"] = o.GeneGroups
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsGeneDescriptor struct {
	value *V2reportsGeneDescriptor
	isSet bool
}

func (v NullableV2reportsGeneDescriptor) Get() *V2reportsGeneDescriptor {
	return v.value
}

func (v *NullableV2reportsGeneDescriptor) Set(val *V2reportsGeneDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsGeneDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsGeneDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsGeneDescriptor(val *V2reportsGeneDescriptor) *NullableV2reportsGeneDescriptor {
	return &NullableV2reportsGeneDescriptor{value: val, isSet: true}
}

func (v NullableV2reportsGeneDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsGeneDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


