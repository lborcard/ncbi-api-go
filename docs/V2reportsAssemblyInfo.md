# V2reportsAssemblyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyLevel** | Pointer to **string** |  | [optional] 
**AssemblyStatus** | Pointer to [**V2reportsAssemblyStatus**](V2reportsAssemblyStatus.md) |  | [optional] [default to V2REPORTSASSEMBLYSTATUS_ASSEMBLY_STATUS_UNKNOWN]
**PairedAssembly** | Pointer to [**V2reportsPairedAssembly**](V2reportsPairedAssembly.md) |  | [optional] 
**AssemblyName** | Pointer to **string** |  | [optional] 
**AssemblyType** | Pointer to **string** |  | [optional] 
**BioprojectLineage** | Pointer to [**[]V2reportsBioProjectLineage**](V2reportsBioProjectLineage.md) |  | [optional] 
**BioprojectAccession** | Pointer to **string** |  | [optional] 
**SubmissionDate** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Submitter** | Pointer to **string** |  | [optional] 
**RefseqCategory** | Pointer to **string** |  | [optional] 
**Synonym** | Pointer to **string** |  | [optional] 
**LinkedAssembly** | Pointer to **string** |  | [optional] 
**LinkedAssemblies** | Pointer to [**[]V2reportsLinkedAssembly**](V2reportsLinkedAssembly.md) |  | [optional] 
**Atypical** | Pointer to [**V2reportsAtypicalInfo**](V2reportsAtypicalInfo.md) |  | [optional] 
**GenomeNotes** | Pointer to **[]string** |  | [optional] 
**SequencingTech** | Pointer to **string** |  | [optional] 
**AssemblyMethod** | Pointer to **string** |  | [optional] 
**Biosample** | Pointer to [**V2reportsBioSampleDescriptor**](V2reportsBioSampleDescriptor.md) |  | [optional] 
**BlastUrl** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**SuppressionReason** | Pointer to **string** |  | [optional] 
**DiploidRole** | Pointer to [**V2reportsLinkedAssemblyType**](V2reportsLinkedAssemblyType.md) |  | [optional] [default to V2REPORTSLINKEDASSEMBLYTYPE_LINKED_ASSEMBLY_TYPE_UNKNOWN]

## Methods

### NewV2reportsAssemblyInfo

`func NewV2reportsAssemblyInfo() *V2reportsAssemblyInfo`

NewV2reportsAssemblyInfo instantiates a new V2reportsAssemblyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsAssemblyInfoWithDefaults

`func NewV2reportsAssemblyInfoWithDefaults() *V2reportsAssemblyInfo`

NewV2reportsAssemblyInfoWithDefaults instantiates a new V2reportsAssemblyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssemblyLevel

`func (o *V2reportsAssemblyInfo) GetAssemblyLevel() string`

GetAssemblyLevel returns the AssemblyLevel field if non-nil, zero value otherwise.

### GetAssemblyLevelOk

`func (o *V2reportsAssemblyInfo) GetAssemblyLevelOk() (*string, bool)`

GetAssemblyLevelOk returns a tuple with the AssemblyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyLevel

`func (o *V2reportsAssemblyInfo) SetAssemblyLevel(v string)`

SetAssemblyLevel sets AssemblyLevel field to given value.

### HasAssemblyLevel

`func (o *V2reportsAssemblyInfo) HasAssemblyLevel() bool`

HasAssemblyLevel returns a boolean if a field has been set.

### GetAssemblyStatus

`func (o *V2reportsAssemblyInfo) GetAssemblyStatus() V2reportsAssemblyStatus`

GetAssemblyStatus returns the AssemblyStatus field if non-nil, zero value otherwise.

### GetAssemblyStatusOk

`func (o *V2reportsAssemblyInfo) GetAssemblyStatusOk() (*V2reportsAssemblyStatus, bool)`

GetAssemblyStatusOk returns a tuple with the AssemblyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyStatus

`func (o *V2reportsAssemblyInfo) SetAssemblyStatus(v V2reportsAssemblyStatus)`

SetAssemblyStatus sets AssemblyStatus field to given value.

### HasAssemblyStatus

`func (o *V2reportsAssemblyInfo) HasAssemblyStatus() bool`

HasAssemblyStatus returns a boolean if a field has been set.

### GetPairedAssembly

`func (o *V2reportsAssemblyInfo) GetPairedAssembly() V2reportsPairedAssembly`

GetPairedAssembly returns the PairedAssembly field if non-nil, zero value otherwise.

### GetPairedAssemblyOk

`func (o *V2reportsAssemblyInfo) GetPairedAssemblyOk() (*V2reportsPairedAssembly, bool)`

GetPairedAssemblyOk returns a tuple with the PairedAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairedAssembly

`func (o *V2reportsAssemblyInfo) SetPairedAssembly(v V2reportsPairedAssembly)`

SetPairedAssembly sets PairedAssembly field to given value.

### HasPairedAssembly

`func (o *V2reportsAssemblyInfo) HasPairedAssembly() bool`

HasPairedAssembly returns a boolean if a field has been set.

### GetAssemblyName

`func (o *V2reportsAssemblyInfo) GetAssemblyName() string`

GetAssemblyName returns the AssemblyName field if non-nil, zero value otherwise.

### GetAssemblyNameOk

`func (o *V2reportsAssemblyInfo) GetAssemblyNameOk() (*string, bool)`

GetAssemblyNameOk returns a tuple with the AssemblyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyName

`func (o *V2reportsAssemblyInfo) SetAssemblyName(v string)`

SetAssemblyName sets AssemblyName field to given value.

### HasAssemblyName

`func (o *V2reportsAssemblyInfo) HasAssemblyName() bool`

HasAssemblyName returns a boolean if a field has been set.

### GetAssemblyType

`func (o *V2reportsAssemblyInfo) GetAssemblyType() string`

GetAssemblyType returns the AssemblyType field if non-nil, zero value otherwise.

### GetAssemblyTypeOk

`func (o *V2reportsAssemblyInfo) GetAssemblyTypeOk() (*string, bool)`

GetAssemblyTypeOk returns a tuple with the AssemblyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyType

`func (o *V2reportsAssemblyInfo) SetAssemblyType(v string)`

SetAssemblyType sets AssemblyType field to given value.

### HasAssemblyType

`func (o *V2reportsAssemblyInfo) HasAssemblyType() bool`

HasAssemblyType returns a boolean if a field has been set.

### GetBioprojectLineage

`func (o *V2reportsAssemblyInfo) GetBioprojectLineage() []V2reportsBioProjectLineage`

GetBioprojectLineage returns the BioprojectLineage field if non-nil, zero value otherwise.

### GetBioprojectLineageOk

`func (o *V2reportsAssemblyInfo) GetBioprojectLineageOk() (*[]V2reportsBioProjectLineage, bool)`

GetBioprojectLineageOk returns a tuple with the BioprojectLineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioprojectLineage

`func (o *V2reportsAssemblyInfo) SetBioprojectLineage(v []V2reportsBioProjectLineage)`

SetBioprojectLineage sets BioprojectLineage field to given value.

### HasBioprojectLineage

`func (o *V2reportsAssemblyInfo) HasBioprojectLineage() bool`

HasBioprojectLineage returns a boolean if a field has been set.

### GetBioprojectAccession

`func (o *V2reportsAssemblyInfo) GetBioprojectAccession() string`

GetBioprojectAccession returns the BioprojectAccession field if non-nil, zero value otherwise.

### GetBioprojectAccessionOk

`func (o *V2reportsAssemblyInfo) GetBioprojectAccessionOk() (*string, bool)`

GetBioprojectAccessionOk returns a tuple with the BioprojectAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioprojectAccession

`func (o *V2reportsAssemblyInfo) SetBioprojectAccession(v string)`

SetBioprojectAccession sets BioprojectAccession field to given value.

### HasBioprojectAccession

`func (o *V2reportsAssemblyInfo) HasBioprojectAccession() bool`

HasBioprojectAccession returns a boolean if a field has been set.

### GetSubmissionDate

`func (o *V2reportsAssemblyInfo) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *V2reportsAssemblyInfo) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *V2reportsAssemblyInfo) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *V2reportsAssemblyInfo) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.

### GetReleaseDate

`func (o *V2reportsAssemblyInfo) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *V2reportsAssemblyInfo) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *V2reportsAssemblyInfo) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *V2reportsAssemblyInfo) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetDescription

`func (o *V2reportsAssemblyInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2reportsAssemblyInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2reportsAssemblyInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V2reportsAssemblyInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubmitter

`func (o *V2reportsAssemblyInfo) GetSubmitter() string`

GetSubmitter returns the Submitter field if non-nil, zero value otherwise.

### GetSubmitterOk

`func (o *V2reportsAssemblyInfo) GetSubmitterOk() (*string, bool)`

GetSubmitterOk returns a tuple with the Submitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitter

`func (o *V2reportsAssemblyInfo) SetSubmitter(v string)`

SetSubmitter sets Submitter field to given value.

### HasSubmitter

`func (o *V2reportsAssemblyInfo) HasSubmitter() bool`

HasSubmitter returns a boolean if a field has been set.

### GetRefseqCategory

`func (o *V2reportsAssemblyInfo) GetRefseqCategory() string`

GetRefseqCategory returns the RefseqCategory field if non-nil, zero value otherwise.

### GetRefseqCategoryOk

`func (o *V2reportsAssemblyInfo) GetRefseqCategoryOk() (*string, bool)`

GetRefseqCategoryOk returns a tuple with the RefseqCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqCategory

`func (o *V2reportsAssemblyInfo) SetRefseqCategory(v string)`

SetRefseqCategory sets RefseqCategory field to given value.

### HasRefseqCategory

`func (o *V2reportsAssemblyInfo) HasRefseqCategory() bool`

HasRefseqCategory returns a boolean if a field has been set.

### GetSynonym

`func (o *V2reportsAssemblyInfo) GetSynonym() string`

GetSynonym returns the Synonym field if non-nil, zero value otherwise.

### GetSynonymOk

`func (o *V2reportsAssemblyInfo) GetSynonymOk() (*string, bool)`

GetSynonymOk returns a tuple with the Synonym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonym

`func (o *V2reportsAssemblyInfo) SetSynonym(v string)`

SetSynonym sets Synonym field to given value.

### HasSynonym

`func (o *V2reportsAssemblyInfo) HasSynonym() bool`

HasSynonym returns a boolean if a field has been set.

### GetLinkedAssembly

`func (o *V2reportsAssemblyInfo) GetLinkedAssembly() string`

GetLinkedAssembly returns the LinkedAssembly field if non-nil, zero value otherwise.

### GetLinkedAssemblyOk

`func (o *V2reportsAssemblyInfo) GetLinkedAssemblyOk() (*string, bool)`

GetLinkedAssemblyOk returns a tuple with the LinkedAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAssembly

`func (o *V2reportsAssemblyInfo) SetLinkedAssembly(v string)`

SetLinkedAssembly sets LinkedAssembly field to given value.

### HasLinkedAssembly

`func (o *V2reportsAssemblyInfo) HasLinkedAssembly() bool`

HasLinkedAssembly returns a boolean if a field has been set.

### GetLinkedAssemblies

`func (o *V2reportsAssemblyInfo) GetLinkedAssemblies() []V2reportsLinkedAssembly`

GetLinkedAssemblies returns the LinkedAssemblies field if non-nil, zero value otherwise.

### GetLinkedAssembliesOk

`func (o *V2reportsAssemblyInfo) GetLinkedAssembliesOk() (*[]V2reportsLinkedAssembly, bool)`

GetLinkedAssembliesOk returns a tuple with the LinkedAssemblies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAssemblies

`func (o *V2reportsAssemblyInfo) SetLinkedAssemblies(v []V2reportsLinkedAssembly)`

SetLinkedAssemblies sets LinkedAssemblies field to given value.

### HasLinkedAssemblies

`func (o *V2reportsAssemblyInfo) HasLinkedAssemblies() bool`

HasLinkedAssemblies returns a boolean if a field has been set.

### GetAtypical

`func (o *V2reportsAssemblyInfo) GetAtypical() V2reportsAtypicalInfo`

GetAtypical returns the Atypical field if non-nil, zero value otherwise.

### GetAtypicalOk

`func (o *V2reportsAssemblyInfo) GetAtypicalOk() (*V2reportsAtypicalInfo, bool)`

GetAtypicalOk returns a tuple with the Atypical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtypical

`func (o *V2reportsAssemblyInfo) SetAtypical(v V2reportsAtypicalInfo)`

SetAtypical sets Atypical field to given value.

### HasAtypical

`func (o *V2reportsAssemblyInfo) HasAtypical() bool`

HasAtypical returns a boolean if a field has been set.

### GetGenomeNotes

`func (o *V2reportsAssemblyInfo) GetGenomeNotes() []string`

GetGenomeNotes returns the GenomeNotes field if non-nil, zero value otherwise.

### GetGenomeNotesOk

`func (o *V2reportsAssemblyInfo) GetGenomeNotesOk() (*[]string, bool)`

GetGenomeNotesOk returns a tuple with the GenomeNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomeNotes

`func (o *V2reportsAssemblyInfo) SetGenomeNotes(v []string)`

SetGenomeNotes sets GenomeNotes field to given value.

### HasGenomeNotes

`func (o *V2reportsAssemblyInfo) HasGenomeNotes() bool`

HasGenomeNotes returns a boolean if a field has been set.

### GetSequencingTech

`func (o *V2reportsAssemblyInfo) GetSequencingTech() string`

GetSequencingTech returns the SequencingTech field if non-nil, zero value otherwise.

### GetSequencingTechOk

`func (o *V2reportsAssemblyInfo) GetSequencingTechOk() (*string, bool)`

GetSequencingTechOk returns a tuple with the SequencingTech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequencingTech

`func (o *V2reportsAssemblyInfo) SetSequencingTech(v string)`

SetSequencingTech sets SequencingTech field to given value.

### HasSequencingTech

`func (o *V2reportsAssemblyInfo) HasSequencingTech() bool`

HasSequencingTech returns a boolean if a field has been set.

### GetAssemblyMethod

`func (o *V2reportsAssemblyInfo) GetAssemblyMethod() string`

GetAssemblyMethod returns the AssemblyMethod field if non-nil, zero value otherwise.

### GetAssemblyMethodOk

`func (o *V2reportsAssemblyInfo) GetAssemblyMethodOk() (*string, bool)`

GetAssemblyMethodOk returns a tuple with the AssemblyMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyMethod

`func (o *V2reportsAssemblyInfo) SetAssemblyMethod(v string)`

SetAssemblyMethod sets AssemblyMethod field to given value.

### HasAssemblyMethod

`func (o *V2reportsAssemblyInfo) HasAssemblyMethod() bool`

HasAssemblyMethod returns a boolean if a field has been set.

### GetBiosample

`func (o *V2reportsAssemblyInfo) GetBiosample() V2reportsBioSampleDescriptor`

GetBiosample returns the Biosample field if non-nil, zero value otherwise.

### GetBiosampleOk

`func (o *V2reportsAssemblyInfo) GetBiosampleOk() (*V2reportsBioSampleDescriptor, bool)`

GetBiosampleOk returns a tuple with the Biosample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosample

`func (o *V2reportsAssemblyInfo) SetBiosample(v V2reportsBioSampleDescriptor)`

SetBiosample sets Biosample field to given value.

### HasBiosample

`func (o *V2reportsAssemblyInfo) HasBiosample() bool`

HasBiosample returns a boolean if a field has been set.

### GetBlastUrl

`func (o *V2reportsAssemblyInfo) GetBlastUrl() string`

GetBlastUrl returns the BlastUrl field if non-nil, zero value otherwise.

### GetBlastUrlOk

`func (o *V2reportsAssemblyInfo) GetBlastUrlOk() (*string, bool)`

GetBlastUrlOk returns a tuple with the BlastUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlastUrl

`func (o *V2reportsAssemblyInfo) SetBlastUrl(v string)`

SetBlastUrl sets BlastUrl field to given value.

### HasBlastUrl

`func (o *V2reportsAssemblyInfo) HasBlastUrl() bool`

HasBlastUrl returns a boolean if a field has been set.

### GetComments

`func (o *V2reportsAssemblyInfo) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *V2reportsAssemblyInfo) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *V2reportsAssemblyInfo) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *V2reportsAssemblyInfo) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetSuppressionReason

`func (o *V2reportsAssemblyInfo) GetSuppressionReason() string`

GetSuppressionReason returns the SuppressionReason field if non-nil, zero value otherwise.

### GetSuppressionReasonOk

`func (o *V2reportsAssemblyInfo) GetSuppressionReasonOk() (*string, bool)`

GetSuppressionReasonOk returns a tuple with the SuppressionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionReason

`func (o *V2reportsAssemblyInfo) SetSuppressionReason(v string)`

SetSuppressionReason sets SuppressionReason field to given value.

### HasSuppressionReason

`func (o *V2reportsAssemblyInfo) HasSuppressionReason() bool`

HasSuppressionReason returns a boolean if a field has been set.

### GetDiploidRole

`func (o *V2reportsAssemblyInfo) GetDiploidRole() V2reportsLinkedAssemblyType`

GetDiploidRole returns the DiploidRole field if non-nil, zero value otherwise.

### GetDiploidRoleOk

`func (o *V2reportsAssemblyInfo) GetDiploidRoleOk() (*V2reportsLinkedAssemblyType, bool)`

GetDiploidRoleOk returns a tuple with the DiploidRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiploidRole

`func (o *V2reportsAssemblyInfo) SetDiploidRole(v V2reportsLinkedAssemblyType)`

SetDiploidRole sets DiploidRole field to given value.

### HasDiploidRole

`func (o *V2reportsAssemblyInfo) HasDiploidRole() bool`

HasDiploidRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


