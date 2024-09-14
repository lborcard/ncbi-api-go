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

// V2reportsGenomeAnnotation struct for V2reportsGenomeAnnotation
type V2reportsGenomeAnnotation struct {
	GeneId *string `json:"gene_id,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	TaxId *string `json:"tax_id,omitempty"`
	Taxname *string `json:"taxname,omitempty"`
	CommonName *string `json:"common_name,omitempty"`
	Type *V2reportsGeneType `json:"type,omitempty"`
	GeneType *string `json:"gene_type,omitempty"`
	RnaType *V2reportsRnaType `json:"rna_type,omitempty"`
	Orientation *V2reportsOrientation `json:"orientation,omitempty"`
	LocusTag *string `json:"locus_tag,omitempty"`
	ReferenceStandards []V2reportsGenomicRegion `json:"reference_standards,omitempty"`
	GenomicRegions []V2reportsGenomicRegion `json:"genomic_regions,omitempty"`
	Transcripts []V2reportsTranscript `json:"transcripts,omitempty"`
	Proteins []V2reportsProtein `json:"proteins,omitempty"`
	Chromosomes []string `json:"chromosomes,omitempty"`
	SwissProtAccessions []string `json:"swiss_prot_accessions,omitempty"`
	EnsemblGeneIds []string `json:"ensembl_gene_ids,omitempty"`
	OmimIds []string `json:"omim_ids,omitempty"`
	Synonyms []string `json:"synonyms,omitempty"`
	Annotations []V2reportsAnnotation `json:"annotations,omitempty"`
}

// NewV2reportsGenomeAnnotation instantiates a new V2reportsGenomeAnnotation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsGenomeAnnotation() *V2reportsGenomeAnnotation {
	this := V2reportsGenomeAnnotation{}
	var type_ V2reportsGeneType = V2REPORTSGENETYPE_UNKNOWN
	this.Type = &type_
	var rnaType V2reportsRnaType = V2REPORTSRNATYPE_RNA_UNKNOWN
	this.RnaType = &rnaType
	var orientation V2reportsOrientation = V2REPORTSORIENTATION_NONE
	this.Orientation = &orientation
	return &this
}

// NewV2reportsGenomeAnnotationWithDefaults instantiates a new V2reportsGenomeAnnotation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsGenomeAnnotationWithDefaults() *V2reportsGenomeAnnotation {
	this := V2reportsGenomeAnnotation{}
	var type_ V2reportsGeneType = V2REPORTSGENETYPE_UNKNOWN
	this.Type = &type_
	var rnaType V2reportsRnaType = V2REPORTSRNATYPE_RNA_UNKNOWN
	this.RnaType = &rnaType
	var orientation V2reportsOrientation = V2REPORTSORIENTATION_NONE
	this.Orientation = &orientation
	return &this
}

// GetGeneId returns the GeneId field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetGeneId() string {
	if o == nil || o.GeneId == nil {
		var ret string
		return ret
	}
	return *o.GeneId
}

// GetGeneIdOk returns a tuple with the GeneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetGeneIdOk() (*string, bool) {
	if o == nil || o.GeneId == nil {
		return nil, false
	}
	return o.GeneId, true
}

// HasGeneId returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasGeneId() bool {
	if o != nil && o.GeneId != nil {
		return true
	}

	return false
}

// SetGeneId gets a reference to the given string and assigns it to the GeneId field.
func (o *V2reportsGenomeAnnotation) SetGeneId(v string) {
	o.GeneId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *V2reportsGenomeAnnotation) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V2reportsGenomeAnnotation) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V2reportsGenomeAnnotation) SetName(v string) {
	o.Name = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetTaxId() string {
	if o == nil || o.TaxId == nil {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetTaxIdOk() (*string, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *V2reportsGenomeAnnotation) SetTaxId(v string) {
	o.TaxId = &v
}

// GetTaxname returns the Taxname field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetTaxname() string {
	if o == nil || o.Taxname == nil {
		var ret string
		return ret
	}
	return *o.Taxname
}

// GetTaxnameOk returns a tuple with the Taxname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetTaxnameOk() (*string, bool) {
	if o == nil || o.Taxname == nil {
		return nil, false
	}
	return o.Taxname, true
}

// HasTaxname returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasTaxname() bool {
	if o != nil && o.Taxname != nil {
		return true
	}

	return false
}

// SetTaxname gets a reference to the given string and assigns it to the Taxname field.
func (o *V2reportsGenomeAnnotation) SetTaxname(v string) {
	o.Taxname = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *V2reportsGenomeAnnotation) SetCommonName(v string) {
	o.CommonName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetType() V2reportsGeneType {
	if o == nil || o.Type == nil {
		var ret V2reportsGeneType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetTypeOk() (*V2reportsGeneType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given V2reportsGeneType and assigns it to the Type field.
func (o *V2reportsGenomeAnnotation) SetType(v V2reportsGeneType) {
	o.Type = &v
}

// GetGeneType returns the GeneType field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetGeneType() string {
	if o == nil || o.GeneType == nil {
		var ret string
		return ret
	}
	return *o.GeneType
}

// GetGeneTypeOk returns a tuple with the GeneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetGeneTypeOk() (*string, bool) {
	if o == nil || o.GeneType == nil {
		return nil, false
	}
	return o.GeneType, true
}

// HasGeneType returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasGeneType() bool {
	if o != nil && o.GeneType != nil {
		return true
	}

	return false
}

// SetGeneType gets a reference to the given string and assigns it to the GeneType field.
func (o *V2reportsGenomeAnnotation) SetGeneType(v string) {
	o.GeneType = &v
}

// GetRnaType returns the RnaType field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetRnaType() V2reportsRnaType {
	if o == nil || o.RnaType == nil {
		var ret V2reportsRnaType
		return ret
	}
	return *o.RnaType
}

// GetRnaTypeOk returns a tuple with the RnaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetRnaTypeOk() (*V2reportsRnaType, bool) {
	if o == nil || o.RnaType == nil {
		return nil, false
	}
	return o.RnaType, true
}

// HasRnaType returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasRnaType() bool {
	if o != nil && o.RnaType != nil {
		return true
	}

	return false
}

// SetRnaType gets a reference to the given V2reportsRnaType and assigns it to the RnaType field.
func (o *V2reportsGenomeAnnotation) SetRnaType(v V2reportsRnaType) {
	o.RnaType = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetOrientation() V2reportsOrientation {
	if o == nil || o.Orientation == nil {
		var ret V2reportsOrientation
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetOrientationOk() (*V2reportsOrientation, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given V2reportsOrientation and assigns it to the Orientation field.
func (o *V2reportsGenomeAnnotation) SetOrientation(v V2reportsOrientation) {
	o.Orientation = &v
}

// GetLocusTag returns the LocusTag field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetLocusTag() string {
	if o == nil || o.LocusTag == nil {
		var ret string
		return ret
	}
	return *o.LocusTag
}

// GetLocusTagOk returns a tuple with the LocusTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetLocusTagOk() (*string, bool) {
	if o == nil || o.LocusTag == nil {
		return nil, false
	}
	return o.LocusTag, true
}

// HasLocusTag returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasLocusTag() bool {
	if o != nil && o.LocusTag != nil {
		return true
	}

	return false
}

// SetLocusTag gets a reference to the given string and assigns it to the LocusTag field.
func (o *V2reportsGenomeAnnotation) SetLocusTag(v string) {
	o.LocusTag = &v
}

// GetReferenceStandards returns the ReferenceStandards field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetReferenceStandards() []V2reportsGenomicRegion {
	if o == nil || o.ReferenceStandards == nil {
		var ret []V2reportsGenomicRegion
		return ret
	}
	return o.ReferenceStandards
}

// GetReferenceStandardsOk returns a tuple with the ReferenceStandards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetReferenceStandardsOk() ([]V2reportsGenomicRegion, bool) {
	if o == nil || o.ReferenceStandards == nil {
		return nil, false
	}
	return o.ReferenceStandards, true
}

// HasReferenceStandards returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasReferenceStandards() bool {
	if o != nil && o.ReferenceStandards != nil {
		return true
	}

	return false
}

// SetReferenceStandards gets a reference to the given []V2reportsGenomicRegion and assigns it to the ReferenceStandards field.
func (o *V2reportsGenomeAnnotation) SetReferenceStandards(v []V2reportsGenomicRegion) {
	o.ReferenceStandards = v
}

// GetGenomicRegions returns the GenomicRegions field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetGenomicRegions() []V2reportsGenomicRegion {
	if o == nil || o.GenomicRegions == nil {
		var ret []V2reportsGenomicRegion
		return ret
	}
	return o.GenomicRegions
}

// GetGenomicRegionsOk returns a tuple with the GenomicRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetGenomicRegionsOk() ([]V2reportsGenomicRegion, bool) {
	if o == nil || o.GenomicRegions == nil {
		return nil, false
	}
	return o.GenomicRegions, true
}

// HasGenomicRegions returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasGenomicRegions() bool {
	if o != nil && o.GenomicRegions != nil {
		return true
	}

	return false
}

// SetGenomicRegions gets a reference to the given []V2reportsGenomicRegion and assigns it to the GenomicRegions field.
func (o *V2reportsGenomeAnnotation) SetGenomicRegions(v []V2reportsGenomicRegion) {
	o.GenomicRegions = v
}

// GetTranscripts returns the Transcripts field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetTranscripts() []V2reportsTranscript {
	if o == nil || o.Transcripts == nil {
		var ret []V2reportsTranscript
		return ret
	}
	return o.Transcripts
}

// GetTranscriptsOk returns a tuple with the Transcripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetTranscriptsOk() ([]V2reportsTranscript, bool) {
	if o == nil || o.Transcripts == nil {
		return nil, false
	}
	return o.Transcripts, true
}

// HasTranscripts returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasTranscripts() bool {
	if o != nil && o.Transcripts != nil {
		return true
	}

	return false
}

// SetTranscripts gets a reference to the given []V2reportsTranscript and assigns it to the Transcripts field.
func (o *V2reportsGenomeAnnotation) SetTranscripts(v []V2reportsTranscript) {
	o.Transcripts = v
}

// GetProteins returns the Proteins field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetProteins() []V2reportsProtein {
	if o == nil || o.Proteins == nil {
		var ret []V2reportsProtein
		return ret
	}
	return o.Proteins
}

// GetProteinsOk returns a tuple with the Proteins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetProteinsOk() ([]V2reportsProtein, bool) {
	if o == nil || o.Proteins == nil {
		return nil, false
	}
	return o.Proteins, true
}

// HasProteins returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasProteins() bool {
	if o != nil && o.Proteins != nil {
		return true
	}

	return false
}

// SetProteins gets a reference to the given []V2reportsProtein and assigns it to the Proteins field.
func (o *V2reportsGenomeAnnotation) SetProteins(v []V2reportsProtein) {
	o.Proteins = v
}

// GetChromosomes returns the Chromosomes field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetChromosomes() []string {
	if o == nil || o.Chromosomes == nil {
		var ret []string
		return ret
	}
	return o.Chromosomes
}

// GetChromosomesOk returns a tuple with the Chromosomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetChromosomesOk() ([]string, bool) {
	if o == nil || o.Chromosomes == nil {
		return nil, false
	}
	return o.Chromosomes, true
}

// HasChromosomes returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasChromosomes() bool {
	if o != nil && o.Chromosomes != nil {
		return true
	}

	return false
}

// SetChromosomes gets a reference to the given []string and assigns it to the Chromosomes field.
func (o *V2reportsGenomeAnnotation) SetChromosomes(v []string) {
	o.Chromosomes = v
}

// GetSwissProtAccessions returns the SwissProtAccessions field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetSwissProtAccessions() []string {
	if o == nil || o.SwissProtAccessions == nil {
		var ret []string
		return ret
	}
	return o.SwissProtAccessions
}

// GetSwissProtAccessionsOk returns a tuple with the SwissProtAccessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetSwissProtAccessionsOk() ([]string, bool) {
	if o == nil || o.SwissProtAccessions == nil {
		return nil, false
	}
	return o.SwissProtAccessions, true
}

// HasSwissProtAccessions returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasSwissProtAccessions() bool {
	if o != nil && o.SwissProtAccessions != nil {
		return true
	}

	return false
}

// SetSwissProtAccessions gets a reference to the given []string and assigns it to the SwissProtAccessions field.
func (o *V2reportsGenomeAnnotation) SetSwissProtAccessions(v []string) {
	o.SwissProtAccessions = v
}

// GetEnsemblGeneIds returns the EnsemblGeneIds field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetEnsemblGeneIds() []string {
	if o == nil || o.EnsemblGeneIds == nil {
		var ret []string
		return ret
	}
	return o.EnsemblGeneIds
}

// GetEnsemblGeneIdsOk returns a tuple with the EnsemblGeneIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetEnsemblGeneIdsOk() ([]string, bool) {
	if o == nil || o.EnsemblGeneIds == nil {
		return nil, false
	}
	return o.EnsemblGeneIds, true
}

// HasEnsemblGeneIds returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasEnsemblGeneIds() bool {
	if o != nil && o.EnsemblGeneIds != nil {
		return true
	}

	return false
}

// SetEnsemblGeneIds gets a reference to the given []string and assigns it to the EnsemblGeneIds field.
func (o *V2reportsGenomeAnnotation) SetEnsemblGeneIds(v []string) {
	o.EnsemblGeneIds = v
}

// GetOmimIds returns the OmimIds field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetOmimIds() []string {
	if o == nil || o.OmimIds == nil {
		var ret []string
		return ret
	}
	return o.OmimIds
}

// GetOmimIdsOk returns a tuple with the OmimIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetOmimIdsOk() ([]string, bool) {
	if o == nil || o.OmimIds == nil {
		return nil, false
	}
	return o.OmimIds, true
}

// HasOmimIds returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasOmimIds() bool {
	if o != nil && o.OmimIds != nil {
		return true
	}

	return false
}

// SetOmimIds gets a reference to the given []string and assigns it to the OmimIds field.
func (o *V2reportsGenomeAnnotation) SetOmimIds(v []string) {
	o.OmimIds = v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetSynonyms() []string {
	if o == nil || o.Synonyms == nil {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetSynonymsOk() ([]string, bool) {
	if o == nil || o.Synonyms == nil {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasSynonyms() bool {
	if o != nil && o.Synonyms != nil {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *V2reportsGenomeAnnotation) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotation) GetAnnotations() []V2reportsAnnotation {
	if o == nil || o.Annotations == nil {
		var ret []V2reportsAnnotation
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotation) GetAnnotationsOk() ([]V2reportsAnnotation, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotation) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []V2reportsAnnotation and assigns it to the Annotations field.
func (o *V2reportsGenomeAnnotation) SetAnnotations(v []V2reportsAnnotation) {
	o.Annotations = v
}

func (o V2reportsGenomeAnnotation) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.GeneType != nil {
		toSerialize["gene_type"] = o.GeneType
	}
	if o.RnaType != nil {
		toSerialize["rna_type"] = o.RnaType
	}
	if o.Orientation != nil {
		toSerialize["orientation"] = o.Orientation
	}
	if o.LocusTag != nil {
		toSerialize["locus_tag"] = o.LocusTag
	}
	if o.ReferenceStandards != nil {
		toSerialize["reference_standards"] = o.ReferenceStandards
	}
	if o.GenomicRegions != nil {
		toSerialize["genomic_regions"] = o.GenomicRegions
	}
	if o.Transcripts != nil {
		toSerialize["transcripts"] = o.Transcripts
	}
	if o.Proteins != nil {
		toSerialize["proteins"] = o.Proteins
	}
	if o.Chromosomes != nil {
		toSerialize["chromosomes"] = o.Chromosomes
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
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsGenomeAnnotation struct {
	value *V2reportsGenomeAnnotation
	isSet bool
}

func (v NullableV2reportsGenomeAnnotation) Get() *V2reportsGenomeAnnotation {
	return v.value
}

func (v *NullableV2reportsGenomeAnnotation) Set(val *V2reportsGenomeAnnotation) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsGenomeAnnotation) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsGenomeAnnotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsGenomeAnnotation(val *V2reportsGenomeAnnotation) *NullableV2reportsGenomeAnnotation {
	return &NullableV2reportsGenomeAnnotation{value: val, isSet: true}
}

func (v NullableV2reportsGenomeAnnotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsGenomeAnnotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


