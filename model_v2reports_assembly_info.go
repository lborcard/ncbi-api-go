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

// V2reportsAssemblyInfo struct for V2reportsAssemblyInfo
type V2reportsAssemblyInfo struct {
	AssemblyLevel *string `json:"assembly_level,omitempty"`
	AssemblyStatus *V2reportsAssemblyStatus `json:"assembly_status,omitempty"`
	PairedAssembly *V2reportsPairedAssembly `json:"paired_assembly,omitempty"`
	AssemblyName *string `json:"assembly_name,omitempty"`
	AssemblyType *string `json:"assembly_type,omitempty"`
	BioprojectLineage []V2reportsBioProjectLineage `json:"bioproject_lineage,omitempty"`
	BioprojectAccession *string `json:"bioproject_accession,omitempty"`
	SubmissionDate *string `json:"submission_date,omitempty"`
	ReleaseDate *string `json:"release_date,omitempty"`
	Description *string `json:"description,omitempty"`
	Submitter *string `json:"submitter,omitempty"`
	RefseqCategory *string `json:"refseq_category,omitempty"`
	Synonym *string `json:"synonym,omitempty"`
	LinkedAssembly *string `json:"linked_assembly,omitempty"`
	LinkedAssemblies []V2reportsLinkedAssembly `json:"linked_assemblies,omitempty"`
	Atypical *V2reportsAtypicalInfo `json:"atypical,omitempty"`
	GenomeNotes []string `json:"genome_notes,omitempty"`
	SequencingTech *string `json:"sequencing_tech,omitempty"`
	AssemblyMethod *string `json:"assembly_method,omitempty"`
	Biosample *V2reportsBioSampleDescriptor `json:"biosample,omitempty"`
	BlastUrl *string `json:"blast_url,omitempty"`
	Comments *string `json:"comments,omitempty"`
	SuppressionReason *string `json:"suppression_reason,omitempty"`
	DiploidRole *V2reportsLinkedAssemblyType `json:"diploid_role,omitempty"`
}

// NewV2reportsAssemblyInfo instantiates a new V2reportsAssemblyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsAssemblyInfo() *V2reportsAssemblyInfo {
	this := V2reportsAssemblyInfo{}
	var assemblyStatus V2reportsAssemblyStatus = V2REPORTSASSEMBLYSTATUS_ASSEMBLY_STATUS_UNKNOWN
	this.AssemblyStatus = &assemblyStatus
	var diploidRole V2reportsLinkedAssemblyType = V2REPORTSLINKEDASSEMBLYTYPE_LINKED_ASSEMBLY_TYPE_UNKNOWN
	this.DiploidRole = &diploidRole
	return &this
}

// NewV2reportsAssemblyInfoWithDefaults instantiates a new V2reportsAssemblyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsAssemblyInfoWithDefaults() *V2reportsAssemblyInfo {
	this := V2reportsAssemblyInfo{}
	var assemblyStatus V2reportsAssemblyStatus = V2REPORTSASSEMBLYSTATUS_ASSEMBLY_STATUS_UNKNOWN
	this.AssemblyStatus = &assemblyStatus
	var diploidRole V2reportsLinkedAssemblyType = V2REPORTSLINKEDASSEMBLYTYPE_LINKED_ASSEMBLY_TYPE_UNKNOWN
	this.DiploidRole = &diploidRole
	return &this
}

// GetAssemblyLevel returns the AssemblyLevel field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetAssemblyLevel() string {
	if o == nil || o.AssemblyLevel == nil {
		var ret string
		return ret
	}
	return *o.AssemblyLevel
}

// GetAssemblyLevelOk returns a tuple with the AssemblyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetAssemblyLevelOk() (*string, bool) {
	if o == nil || o.AssemblyLevel == nil {
		return nil, false
	}
	return o.AssemblyLevel, true
}

// HasAssemblyLevel returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasAssemblyLevel() bool {
	if o != nil && o.AssemblyLevel != nil {
		return true
	}

	return false
}

// SetAssemblyLevel gets a reference to the given string and assigns it to the AssemblyLevel field.
func (o *V2reportsAssemblyInfo) SetAssemblyLevel(v string) {
	o.AssemblyLevel = &v
}

// GetAssemblyStatus returns the AssemblyStatus field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetAssemblyStatus() V2reportsAssemblyStatus {
	if o == nil || o.AssemblyStatus == nil {
		var ret V2reportsAssemblyStatus
		return ret
	}
	return *o.AssemblyStatus
}

// GetAssemblyStatusOk returns a tuple with the AssemblyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetAssemblyStatusOk() (*V2reportsAssemblyStatus, bool) {
	if o == nil || o.AssemblyStatus == nil {
		return nil, false
	}
	return o.AssemblyStatus, true
}

// HasAssemblyStatus returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasAssemblyStatus() bool {
	if o != nil && o.AssemblyStatus != nil {
		return true
	}

	return false
}

// SetAssemblyStatus gets a reference to the given V2reportsAssemblyStatus and assigns it to the AssemblyStatus field.
func (o *V2reportsAssemblyInfo) SetAssemblyStatus(v V2reportsAssemblyStatus) {
	o.AssemblyStatus = &v
}

// GetPairedAssembly returns the PairedAssembly field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetPairedAssembly() V2reportsPairedAssembly {
	if o == nil || o.PairedAssembly == nil {
		var ret V2reportsPairedAssembly
		return ret
	}
	return *o.PairedAssembly
}

// GetPairedAssemblyOk returns a tuple with the PairedAssembly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetPairedAssemblyOk() (*V2reportsPairedAssembly, bool) {
	if o == nil || o.PairedAssembly == nil {
		return nil, false
	}
	return o.PairedAssembly, true
}

// HasPairedAssembly returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasPairedAssembly() bool {
	if o != nil && o.PairedAssembly != nil {
		return true
	}

	return false
}

// SetPairedAssembly gets a reference to the given V2reportsPairedAssembly and assigns it to the PairedAssembly field.
func (o *V2reportsAssemblyInfo) SetPairedAssembly(v V2reportsPairedAssembly) {
	o.PairedAssembly = &v
}

// GetAssemblyName returns the AssemblyName field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetAssemblyName() string {
	if o == nil || o.AssemblyName == nil {
		var ret string
		return ret
	}
	return *o.AssemblyName
}

// GetAssemblyNameOk returns a tuple with the AssemblyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetAssemblyNameOk() (*string, bool) {
	if o == nil || o.AssemblyName == nil {
		return nil, false
	}
	return o.AssemblyName, true
}

// HasAssemblyName returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasAssemblyName() bool {
	if o != nil && o.AssemblyName != nil {
		return true
	}

	return false
}

// SetAssemblyName gets a reference to the given string and assigns it to the AssemblyName field.
func (o *V2reportsAssemblyInfo) SetAssemblyName(v string) {
	o.AssemblyName = &v
}

// GetAssemblyType returns the AssemblyType field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetAssemblyType() string {
	if o == nil || o.AssemblyType == nil {
		var ret string
		return ret
	}
	return *o.AssemblyType
}

// GetAssemblyTypeOk returns a tuple with the AssemblyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetAssemblyTypeOk() (*string, bool) {
	if o == nil || o.AssemblyType == nil {
		return nil, false
	}
	return o.AssemblyType, true
}

// HasAssemblyType returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasAssemblyType() bool {
	if o != nil && o.AssemblyType != nil {
		return true
	}

	return false
}

// SetAssemblyType gets a reference to the given string and assigns it to the AssemblyType field.
func (o *V2reportsAssemblyInfo) SetAssemblyType(v string) {
	o.AssemblyType = &v
}

// GetBioprojectLineage returns the BioprojectLineage field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetBioprojectLineage() []V2reportsBioProjectLineage {
	if o == nil || o.BioprojectLineage == nil {
		var ret []V2reportsBioProjectLineage
		return ret
	}
	return o.BioprojectLineage
}

// GetBioprojectLineageOk returns a tuple with the BioprojectLineage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetBioprojectLineageOk() ([]V2reportsBioProjectLineage, bool) {
	if o == nil || o.BioprojectLineage == nil {
		return nil, false
	}
	return o.BioprojectLineage, true
}

// HasBioprojectLineage returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasBioprojectLineage() bool {
	if o != nil && o.BioprojectLineage != nil {
		return true
	}

	return false
}

// SetBioprojectLineage gets a reference to the given []V2reportsBioProjectLineage and assigns it to the BioprojectLineage field.
func (o *V2reportsAssemblyInfo) SetBioprojectLineage(v []V2reportsBioProjectLineage) {
	o.BioprojectLineage = v
}

// GetBioprojectAccession returns the BioprojectAccession field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetBioprojectAccession() string {
	if o == nil || o.BioprojectAccession == nil {
		var ret string
		return ret
	}
	return *o.BioprojectAccession
}

// GetBioprojectAccessionOk returns a tuple with the BioprojectAccession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetBioprojectAccessionOk() (*string, bool) {
	if o == nil || o.BioprojectAccession == nil {
		return nil, false
	}
	return o.BioprojectAccession, true
}

// HasBioprojectAccession returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasBioprojectAccession() bool {
	if o != nil && o.BioprojectAccession != nil {
		return true
	}

	return false
}

// SetBioprojectAccession gets a reference to the given string and assigns it to the BioprojectAccession field.
func (o *V2reportsAssemblyInfo) SetBioprojectAccession(v string) {
	o.BioprojectAccession = &v
}

// GetSubmissionDate returns the SubmissionDate field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetSubmissionDate() string {
	if o == nil || o.SubmissionDate == nil {
		var ret string
		return ret
	}
	return *o.SubmissionDate
}

// GetSubmissionDateOk returns a tuple with the SubmissionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetSubmissionDateOk() (*string, bool) {
	if o == nil || o.SubmissionDate == nil {
		return nil, false
	}
	return o.SubmissionDate, true
}

// HasSubmissionDate returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasSubmissionDate() bool {
	if o != nil && o.SubmissionDate != nil {
		return true
	}

	return false
}

// SetSubmissionDate gets a reference to the given string and assigns it to the SubmissionDate field.
func (o *V2reportsAssemblyInfo) SetSubmissionDate(v string) {
	o.SubmissionDate = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetReleaseDate() string {
	if o == nil || o.ReleaseDate == nil {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetReleaseDateOk() (*string, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate != nil {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *V2reportsAssemblyInfo) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V2reportsAssemblyInfo) SetDescription(v string) {
	o.Description = &v
}

// GetSubmitter returns the Submitter field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetSubmitter() string {
	if o == nil || o.Submitter == nil {
		var ret string
		return ret
	}
	return *o.Submitter
}

// GetSubmitterOk returns a tuple with the Submitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetSubmitterOk() (*string, bool) {
	if o == nil || o.Submitter == nil {
		return nil, false
	}
	return o.Submitter, true
}

// HasSubmitter returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasSubmitter() bool {
	if o != nil && o.Submitter != nil {
		return true
	}

	return false
}

// SetSubmitter gets a reference to the given string and assigns it to the Submitter field.
func (o *V2reportsAssemblyInfo) SetSubmitter(v string) {
	o.Submitter = &v
}

// GetRefseqCategory returns the RefseqCategory field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetRefseqCategory() string {
	if o == nil || o.RefseqCategory == nil {
		var ret string
		return ret
	}
	return *o.RefseqCategory
}

// GetRefseqCategoryOk returns a tuple with the RefseqCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetRefseqCategoryOk() (*string, bool) {
	if o == nil || o.RefseqCategory == nil {
		return nil, false
	}
	return o.RefseqCategory, true
}

// HasRefseqCategory returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasRefseqCategory() bool {
	if o != nil && o.RefseqCategory != nil {
		return true
	}

	return false
}

// SetRefseqCategory gets a reference to the given string and assigns it to the RefseqCategory field.
func (o *V2reportsAssemblyInfo) SetRefseqCategory(v string) {
	o.RefseqCategory = &v
}

// GetSynonym returns the Synonym field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetSynonym() string {
	if o == nil || o.Synonym == nil {
		var ret string
		return ret
	}
	return *o.Synonym
}

// GetSynonymOk returns a tuple with the Synonym field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetSynonymOk() (*string, bool) {
	if o == nil || o.Synonym == nil {
		return nil, false
	}
	return o.Synonym, true
}

// HasSynonym returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasSynonym() bool {
	if o != nil && o.Synonym != nil {
		return true
	}

	return false
}

// SetSynonym gets a reference to the given string and assigns it to the Synonym field.
func (o *V2reportsAssemblyInfo) SetSynonym(v string) {
	o.Synonym = &v
}

// GetLinkedAssembly returns the LinkedAssembly field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetLinkedAssembly() string {
	if o == nil || o.LinkedAssembly == nil {
		var ret string
		return ret
	}
	return *o.LinkedAssembly
}

// GetLinkedAssemblyOk returns a tuple with the LinkedAssembly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetLinkedAssemblyOk() (*string, bool) {
	if o == nil || o.LinkedAssembly == nil {
		return nil, false
	}
	return o.LinkedAssembly, true
}

// HasLinkedAssembly returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasLinkedAssembly() bool {
	if o != nil && o.LinkedAssembly != nil {
		return true
	}

	return false
}

// SetLinkedAssembly gets a reference to the given string and assigns it to the LinkedAssembly field.
func (o *V2reportsAssemblyInfo) SetLinkedAssembly(v string) {
	o.LinkedAssembly = &v
}

// GetLinkedAssemblies returns the LinkedAssemblies field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetLinkedAssemblies() []V2reportsLinkedAssembly {
	if o == nil || o.LinkedAssemblies == nil {
		var ret []V2reportsLinkedAssembly
		return ret
	}
	return o.LinkedAssemblies
}

// GetLinkedAssembliesOk returns a tuple with the LinkedAssemblies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetLinkedAssembliesOk() ([]V2reportsLinkedAssembly, bool) {
	if o == nil || o.LinkedAssemblies == nil {
		return nil, false
	}
	return o.LinkedAssemblies, true
}

// HasLinkedAssemblies returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasLinkedAssemblies() bool {
	if o != nil && o.LinkedAssemblies != nil {
		return true
	}

	return false
}

// SetLinkedAssemblies gets a reference to the given []V2reportsLinkedAssembly and assigns it to the LinkedAssemblies field.
func (o *V2reportsAssemblyInfo) SetLinkedAssemblies(v []V2reportsLinkedAssembly) {
	o.LinkedAssemblies = v
}

// GetAtypical returns the Atypical field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetAtypical() V2reportsAtypicalInfo {
	if o == nil || o.Atypical == nil {
		var ret V2reportsAtypicalInfo
		return ret
	}
	return *o.Atypical
}

// GetAtypicalOk returns a tuple with the Atypical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetAtypicalOk() (*V2reportsAtypicalInfo, bool) {
	if o == nil || o.Atypical == nil {
		return nil, false
	}
	return o.Atypical, true
}

// HasAtypical returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasAtypical() bool {
	if o != nil && o.Atypical != nil {
		return true
	}

	return false
}

// SetAtypical gets a reference to the given V2reportsAtypicalInfo and assigns it to the Atypical field.
func (o *V2reportsAssemblyInfo) SetAtypical(v V2reportsAtypicalInfo) {
	o.Atypical = &v
}

// GetGenomeNotes returns the GenomeNotes field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetGenomeNotes() []string {
	if o == nil || o.GenomeNotes == nil {
		var ret []string
		return ret
	}
	return o.GenomeNotes
}

// GetGenomeNotesOk returns a tuple with the GenomeNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetGenomeNotesOk() ([]string, bool) {
	if o == nil || o.GenomeNotes == nil {
		return nil, false
	}
	return o.GenomeNotes, true
}

// HasGenomeNotes returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasGenomeNotes() bool {
	if o != nil && o.GenomeNotes != nil {
		return true
	}

	return false
}

// SetGenomeNotes gets a reference to the given []string and assigns it to the GenomeNotes field.
func (o *V2reportsAssemblyInfo) SetGenomeNotes(v []string) {
	o.GenomeNotes = v
}

// GetSequencingTech returns the SequencingTech field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetSequencingTech() string {
	if o == nil || o.SequencingTech == nil {
		var ret string
		return ret
	}
	return *o.SequencingTech
}

// GetSequencingTechOk returns a tuple with the SequencingTech field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetSequencingTechOk() (*string, bool) {
	if o == nil || o.SequencingTech == nil {
		return nil, false
	}
	return o.SequencingTech, true
}

// HasSequencingTech returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasSequencingTech() bool {
	if o != nil && o.SequencingTech != nil {
		return true
	}

	return false
}

// SetSequencingTech gets a reference to the given string and assigns it to the SequencingTech field.
func (o *V2reportsAssemblyInfo) SetSequencingTech(v string) {
	o.SequencingTech = &v
}

// GetAssemblyMethod returns the AssemblyMethod field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetAssemblyMethod() string {
	if o == nil || o.AssemblyMethod == nil {
		var ret string
		return ret
	}
	return *o.AssemblyMethod
}

// GetAssemblyMethodOk returns a tuple with the AssemblyMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetAssemblyMethodOk() (*string, bool) {
	if o == nil || o.AssemblyMethod == nil {
		return nil, false
	}
	return o.AssemblyMethod, true
}

// HasAssemblyMethod returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasAssemblyMethod() bool {
	if o != nil && o.AssemblyMethod != nil {
		return true
	}

	return false
}

// SetAssemblyMethod gets a reference to the given string and assigns it to the AssemblyMethod field.
func (o *V2reportsAssemblyInfo) SetAssemblyMethod(v string) {
	o.AssemblyMethod = &v
}

// GetBiosample returns the Biosample field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetBiosample() V2reportsBioSampleDescriptor {
	if o == nil || o.Biosample == nil {
		var ret V2reportsBioSampleDescriptor
		return ret
	}
	return *o.Biosample
}

// GetBiosampleOk returns a tuple with the Biosample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetBiosampleOk() (*V2reportsBioSampleDescriptor, bool) {
	if o == nil || o.Biosample == nil {
		return nil, false
	}
	return o.Biosample, true
}

// HasBiosample returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasBiosample() bool {
	if o != nil && o.Biosample != nil {
		return true
	}

	return false
}

// SetBiosample gets a reference to the given V2reportsBioSampleDescriptor and assigns it to the Biosample field.
func (o *V2reportsAssemblyInfo) SetBiosample(v V2reportsBioSampleDescriptor) {
	o.Biosample = &v
}

// GetBlastUrl returns the BlastUrl field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetBlastUrl() string {
	if o == nil || o.BlastUrl == nil {
		var ret string
		return ret
	}
	return *o.BlastUrl
}

// GetBlastUrlOk returns a tuple with the BlastUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetBlastUrlOk() (*string, bool) {
	if o == nil || o.BlastUrl == nil {
		return nil, false
	}
	return o.BlastUrl, true
}

// HasBlastUrl returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasBlastUrl() bool {
	if o != nil && o.BlastUrl != nil {
		return true
	}

	return false
}

// SetBlastUrl gets a reference to the given string and assigns it to the BlastUrl field.
func (o *V2reportsAssemblyInfo) SetBlastUrl(v string) {
	o.BlastUrl = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *V2reportsAssemblyInfo) SetComments(v string) {
	o.Comments = &v
}

// GetSuppressionReason returns the SuppressionReason field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetSuppressionReason() string {
	if o == nil || o.SuppressionReason == nil {
		var ret string
		return ret
	}
	return *o.SuppressionReason
}

// GetSuppressionReasonOk returns a tuple with the SuppressionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetSuppressionReasonOk() (*string, bool) {
	if o == nil || o.SuppressionReason == nil {
		return nil, false
	}
	return o.SuppressionReason, true
}

// HasSuppressionReason returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasSuppressionReason() bool {
	if o != nil && o.SuppressionReason != nil {
		return true
	}

	return false
}

// SetSuppressionReason gets a reference to the given string and assigns it to the SuppressionReason field.
func (o *V2reportsAssemblyInfo) SetSuppressionReason(v string) {
	o.SuppressionReason = &v
}

// GetDiploidRole returns the DiploidRole field value if set, zero value otherwise.
func (o *V2reportsAssemblyInfo) GetDiploidRole() V2reportsLinkedAssemblyType {
	if o == nil || o.DiploidRole == nil {
		var ret V2reportsLinkedAssemblyType
		return ret
	}
	return *o.DiploidRole
}

// GetDiploidRoleOk returns a tuple with the DiploidRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsAssemblyInfo) GetDiploidRoleOk() (*V2reportsLinkedAssemblyType, bool) {
	if o == nil || o.DiploidRole == nil {
		return nil, false
	}
	return o.DiploidRole, true
}

// HasDiploidRole returns a boolean if a field has been set.
func (o *V2reportsAssemblyInfo) HasDiploidRole() bool {
	if o != nil && o.DiploidRole != nil {
		return true
	}

	return false
}

// SetDiploidRole gets a reference to the given V2reportsLinkedAssemblyType and assigns it to the DiploidRole field.
func (o *V2reportsAssemblyInfo) SetDiploidRole(v V2reportsLinkedAssemblyType) {
	o.DiploidRole = &v
}

func (o V2reportsAssemblyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssemblyLevel != nil {
		toSerialize["assembly_level"] = o.AssemblyLevel
	}
	if o.AssemblyStatus != nil {
		toSerialize["assembly_status"] = o.AssemblyStatus
	}
	if o.PairedAssembly != nil {
		toSerialize["paired_assembly"] = o.PairedAssembly
	}
	if o.AssemblyName != nil {
		toSerialize["assembly_name"] = o.AssemblyName
	}
	if o.AssemblyType != nil {
		toSerialize["assembly_type"] = o.AssemblyType
	}
	if o.BioprojectLineage != nil {
		toSerialize["bioproject_lineage"] = o.BioprojectLineage
	}
	if o.BioprojectAccession != nil {
		toSerialize["bioproject_accession"] = o.BioprojectAccession
	}
	if o.SubmissionDate != nil {
		toSerialize["submission_date"] = o.SubmissionDate
	}
	if o.ReleaseDate != nil {
		toSerialize["release_date"] = o.ReleaseDate
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Submitter != nil {
		toSerialize["submitter"] = o.Submitter
	}
	if o.RefseqCategory != nil {
		toSerialize["refseq_category"] = o.RefseqCategory
	}
	if o.Synonym != nil {
		toSerialize["synonym"] = o.Synonym
	}
	if o.LinkedAssembly != nil {
		toSerialize["linked_assembly"] = o.LinkedAssembly
	}
	if o.LinkedAssemblies != nil {
		toSerialize["linked_assemblies"] = o.LinkedAssemblies
	}
	if o.Atypical != nil {
		toSerialize["atypical"] = o.Atypical
	}
	if o.GenomeNotes != nil {
		toSerialize["genome_notes"] = o.GenomeNotes
	}
	if o.SequencingTech != nil {
		toSerialize["sequencing_tech"] = o.SequencingTech
	}
	if o.AssemblyMethod != nil {
		toSerialize["assembly_method"] = o.AssemblyMethod
	}
	if o.Biosample != nil {
		toSerialize["biosample"] = o.Biosample
	}
	if o.BlastUrl != nil {
		toSerialize["blast_url"] = o.BlastUrl
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.SuppressionReason != nil {
		toSerialize["suppression_reason"] = o.SuppressionReason
	}
	if o.DiploidRole != nil {
		toSerialize["diploid_role"] = o.DiploidRole
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsAssemblyInfo struct {
	value *V2reportsAssemblyInfo
	isSet bool
}

func (v NullableV2reportsAssemblyInfo) Get() *V2reportsAssemblyInfo {
	return v.value
}

func (v *NullableV2reportsAssemblyInfo) Set(val *V2reportsAssemblyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsAssemblyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsAssemblyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsAssemblyInfo(val *V2reportsAssemblyInfo) *NullableV2reportsAssemblyInfo {
	return &NullableV2reportsAssemblyInfo{value: val, isSet: true}
}

func (v NullableV2reportsAssemblyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsAssemblyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


