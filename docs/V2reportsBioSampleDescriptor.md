# V2reportsBioSampleDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accession** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**PublicationDate** | Pointer to **string** |  | [optional] 
**SubmissionDate** | Pointer to **string** |  | [optional] 
**SampleIds** | Pointer to [**[]V2reportsBioSampleId**](V2reportsBioSampleId.md) |  | [optional] 
**Description** | Pointer to [**V2reportsBioSampleDescription**](V2reportsBioSampleDescription.md) |  | [optional] 
**Owner** | Pointer to [**V2reportsBioSampleOwner**](V2reportsBioSampleOwner.md) |  | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 
**Bioprojects** | Pointer to [**[]V2reportsBioProject**](V2reportsBioProject.md) |  | [optional] 
**Package** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**[]V2reportsBioSampleAttribute**](V2reportsBioSampleAttribute.md) |  | [optional] 
**Status** | Pointer to [**V2reportsBioSampleStatus**](V2reportsBioSampleStatus.md) |  | [optional] 
**Age** | Pointer to **string** |  | [optional] 
**BiomaterialProvider** | Pointer to **string** |  | [optional] 
**Breed** | Pointer to **string** |  | [optional] 
**CollectedBy** | Pointer to **string** |  | [optional] 
**CollectionDate** | Pointer to **string** |  | [optional] 
**Cultivar** | Pointer to **string** |  | [optional] 
**DevStage** | Pointer to **string** |  | [optional] 
**Ecotype** | Pointer to **string** |  | [optional] 
**GeoLocName** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**HostDisease** | Pointer to **string** |  | [optional] 
**IdentifiedBy** | Pointer to **string** |  | [optional] 
**IfsacCategory** | Pointer to **string** |  | [optional] 
**Isolate** | Pointer to **string** |  | [optional] 
**IsolateNameAlias** | Pointer to **string** |  | [optional] 
**IsolationSource** | Pointer to **string** |  | [optional] 
**LatLon** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**SampleName** | Pointer to **string** |  | [optional] 
**Serovar** | Pointer to **string** |  | [optional] 
**Sex** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**Strain** | Pointer to **string** |  | [optional] 
**SubSpecies** | Pointer to **string** |  | [optional] 
**Tissue** | Pointer to **string** |  | [optional] 
**Serotype** | Pointer to **string** |  | [optional] 

## Methods

### NewV2reportsBioSampleDescriptor

`func NewV2reportsBioSampleDescriptor() *V2reportsBioSampleDescriptor`

NewV2reportsBioSampleDescriptor instantiates a new V2reportsBioSampleDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsBioSampleDescriptorWithDefaults

`func NewV2reportsBioSampleDescriptorWithDefaults() *V2reportsBioSampleDescriptor`

NewV2reportsBioSampleDescriptorWithDefaults instantiates a new V2reportsBioSampleDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccession

`func (o *V2reportsBioSampleDescriptor) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V2reportsBioSampleDescriptor) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V2reportsBioSampleDescriptor) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V2reportsBioSampleDescriptor) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetLastUpdated

`func (o *V2reportsBioSampleDescriptor) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *V2reportsBioSampleDescriptor) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *V2reportsBioSampleDescriptor) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *V2reportsBioSampleDescriptor) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPublicationDate

`func (o *V2reportsBioSampleDescriptor) GetPublicationDate() string`

GetPublicationDate returns the PublicationDate field if non-nil, zero value otherwise.

### GetPublicationDateOk

`func (o *V2reportsBioSampleDescriptor) GetPublicationDateOk() (*string, bool)`

GetPublicationDateOk returns a tuple with the PublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDate

`func (o *V2reportsBioSampleDescriptor) SetPublicationDate(v string)`

SetPublicationDate sets PublicationDate field to given value.

### HasPublicationDate

`func (o *V2reportsBioSampleDescriptor) HasPublicationDate() bool`

HasPublicationDate returns a boolean if a field has been set.

### GetSubmissionDate

`func (o *V2reportsBioSampleDescriptor) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *V2reportsBioSampleDescriptor) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *V2reportsBioSampleDescriptor) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *V2reportsBioSampleDescriptor) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.

### GetSampleIds

`func (o *V2reportsBioSampleDescriptor) GetSampleIds() []V2reportsBioSampleId`

GetSampleIds returns the SampleIds field if non-nil, zero value otherwise.

### GetSampleIdsOk

`func (o *V2reportsBioSampleDescriptor) GetSampleIdsOk() (*[]V2reportsBioSampleId, bool)`

GetSampleIdsOk returns a tuple with the SampleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIds

`func (o *V2reportsBioSampleDescriptor) SetSampleIds(v []V2reportsBioSampleId)`

SetSampleIds sets SampleIds field to given value.

### HasSampleIds

`func (o *V2reportsBioSampleDescriptor) HasSampleIds() bool`

HasSampleIds returns a boolean if a field has been set.

### GetDescription

`func (o *V2reportsBioSampleDescriptor) GetDescription() V2reportsBioSampleDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2reportsBioSampleDescriptor) GetDescriptionOk() (*V2reportsBioSampleDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2reportsBioSampleDescriptor) SetDescription(v V2reportsBioSampleDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V2reportsBioSampleDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *V2reportsBioSampleDescriptor) GetOwner() V2reportsBioSampleOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V2reportsBioSampleDescriptor) GetOwnerOk() (*V2reportsBioSampleOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V2reportsBioSampleDescriptor) SetOwner(v V2reportsBioSampleOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V2reportsBioSampleDescriptor) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetModels

`func (o *V2reportsBioSampleDescriptor) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *V2reportsBioSampleDescriptor) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *V2reportsBioSampleDescriptor) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *V2reportsBioSampleDescriptor) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetBioprojects

`func (o *V2reportsBioSampleDescriptor) GetBioprojects() []V2reportsBioProject`

GetBioprojects returns the Bioprojects field if non-nil, zero value otherwise.

### GetBioprojectsOk

`func (o *V2reportsBioSampleDescriptor) GetBioprojectsOk() (*[]V2reportsBioProject, bool)`

GetBioprojectsOk returns a tuple with the Bioprojects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioprojects

`func (o *V2reportsBioSampleDescriptor) SetBioprojects(v []V2reportsBioProject)`

SetBioprojects sets Bioprojects field to given value.

### HasBioprojects

`func (o *V2reportsBioSampleDescriptor) HasBioprojects() bool`

HasBioprojects returns a boolean if a field has been set.

### GetPackage

`func (o *V2reportsBioSampleDescriptor) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *V2reportsBioSampleDescriptor) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *V2reportsBioSampleDescriptor) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *V2reportsBioSampleDescriptor) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetAttributes

`func (o *V2reportsBioSampleDescriptor) GetAttributes() []V2reportsBioSampleAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *V2reportsBioSampleDescriptor) GetAttributesOk() (*[]V2reportsBioSampleAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *V2reportsBioSampleDescriptor) SetAttributes(v []V2reportsBioSampleAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *V2reportsBioSampleDescriptor) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetStatus

`func (o *V2reportsBioSampleDescriptor) GetStatus() V2reportsBioSampleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2reportsBioSampleDescriptor) GetStatusOk() (*V2reportsBioSampleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2reportsBioSampleDescriptor) SetStatus(v V2reportsBioSampleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V2reportsBioSampleDescriptor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAge

`func (o *V2reportsBioSampleDescriptor) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *V2reportsBioSampleDescriptor) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *V2reportsBioSampleDescriptor) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *V2reportsBioSampleDescriptor) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetBiomaterialProvider

`func (o *V2reportsBioSampleDescriptor) GetBiomaterialProvider() string`

GetBiomaterialProvider returns the BiomaterialProvider field if non-nil, zero value otherwise.

### GetBiomaterialProviderOk

`func (o *V2reportsBioSampleDescriptor) GetBiomaterialProviderOk() (*string, bool)`

GetBiomaterialProviderOk returns a tuple with the BiomaterialProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomaterialProvider

`func (o *V2reportsBioSampleDescriptor) SetBiomaterialProvider(v string)`

SetBiomaterialProvider sets BiomaterialProvider field to given value.

### HasBiomaterialProvider

`func (o *V2reportsBioSampleDescriptor) HasBiomaterialProvider() bool`

HasBiomaterialProvider returns a boolean if a field has been set.

### GetBreed

`func (o *V2reportsBioSampleDescriptor) GetBreed() string`

GetBreed returns the Breed field if non-nil, zero value otherwise.

### GetBreedOk

`func (o *V2reportsBioSampleDescriptor) GetBreedOk() (*string, bool)`

GetBreedOk returns a tuple with the Breed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreed

`func (o *V2reportsBioSampleDescriptor) SetBreed(v string)`

SetBreed sets Breed field to given value.

### HasBreed

`func (o *V2reportsBioSampleDescriptor) HasBreed() bool`

HasBreed returns a boolean if a field has been set.

### GetCollectedBy

`func (o *V2reportsBioSampleDescriptor) GetCollectedBy() string`

GetCollectedBy returns the CollectedBy field if non-nil, zero value otherwise.

### GetCollectedByOk

`func (o *V2reportsBioSampleDescriptor) GetCollectedByOk() (*string, bool)`

GetCollectedByOk returns a tuple with the CollectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectedBy

`func (o *V2reportsBioSampleDescriptor) SetCollectedBy(v string)`

SetCollectedBy sets CollectedBy field to given value.

### HasCollectedBy

`func (o *V2reportsBioSampleDescriptor) HasCollectedBy() bool`

HasCollectedBy returns a boolean if a field has been set.

### GetCollectionDate

`func (o *V2reportsBioSampleDescriptor) GetCollectionDate() string`

GetCollectionDate returns the CollectionDate field if non-nil, zero value otherwise.

### GetCollectionDateOk

`func (o *V2reportsBioSampleDescriptor) GetCollectionDateOk() (*string, bool)`

GetCollectionDateOk returns a tuple with the CollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionDate

`func (o *V2reportsBioSampleDescriptor) SetCollectionDate(v string)`

SetCollectionDate sets CollectionDate field to given value.

### HasCollectionDate

`func (o *V2reportsBioSampleDescriptor) HasCollectionDate() bool`

HasCollectionDate returns a boolean if a field has been set.

### GetCultivar

`func (o *V2reportsBioSampleDescriptor) GetCultivar() string`

GetCultivar returns the Cultivar field if non-nil, zero value otherwise.

### GetCultivarOk

`func (o *V2reportsBioSampleDescriptor) GetCultivarOk() (*string, bool)`

GetCultivarOk returns a tuple with the Cultivar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultivar

`func (o *V2reportsBioSampleDescriptor) SetCultivar(v string)`

SetCultivar sets Cultivar field to given value.

### HasCultivar

`func (o *V2reportsBioSampleDescriptor) HasCultivar() bool`

HasCultivar returns a boolean if a field has been set.

### GetDevStage

`func (o *V2reportsBioSampleDescriptor) GetDevStage() string`

GetDevStage returns the DevStage field if non-nil, zero value otherwise.

### GetDevStageOk

`func (o *V2reportsBioSampleDescriptor) GetDevStageOk() (*string, bool)`

GetDevStageOk returns a tuple with the DevStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevStage

`func (o *V2reportsBioSampleDescriptor) SetDevStage(v string)`

SetDevStage sets DevStage field to given value.

### HasDevStage

`func (o *V2reportsBioSampleDescriptor) HasDevStage() bool`

HasDevStage returns a boolean if a field has been set.

### GetEcotype

`func (o *V2reportsBioSampleDescriptor) GetEcotype() string`

GetEcotype returns the Ecotype field if non-nil, zero value otherwise.

### GetEcotypeOk

`func (o *V2reportsBioSampleDescriptor) GetEcotypeOk() (*string, bool)`

GetEcotypeOk returns a tuple with the Ecotype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcotype

`func (o *V2reportsBioSampleDescriptor) SetEcotype(v string)`

SetEcotype sets Ecotype field to given value.

### HasEcotype

`func (o *V2reportsBioSampleDescriptor) HasEcotype() bool`

HasEcotype returns a boolean if a field has been set.

### GetGeoLocName

`func (o *V2reportsBioSampleDescriptor) GetGeoLocName() string`

GetGeoLocName returns the GeoLocName field if non-nil, zero value otherwise.

### GetGeoLocNameOk

`func (o *V2reportsBioSampleDescriptor) GetGeoLocNameOk() (*string, bool)`

GetGeoLocNameOk returns a tuple with the GeoLocName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoLocName

`func (o *V2reportsBioSampleDescriptor) SetGeoLocName(v string)`

SetGeoLocName sets GeoLocName field to given value.

### HasGeoLocName

`func (o *V2reportsBioSampleDescriptor) HasGeoLocName() bool`

HasGeoLocName returns a boolean if a field has been set.

### GetHost

`func (o *V2reportsBioSampleDescriptor) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V2reportsBioSampleDescriptor) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V2reportsBioSampleDescriptor) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V2reportsBioSampleDescriptor) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHostDisease

`func (o *V2reportsBioSampleDescriptor) GetHostDisease() string`

GetHostDisease returns the HostDisease field if non-nil, zero value otherwise.

### GetHostDiseaseOk

`func (o *V2reportsBioSampleDescriptor) GetHostDiseaseOk() (*string, bool)`

GetHostDiseaseOk returns a tuple with the HostDisease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDisease

`func (o *V2reportsBioSampleDescriptor) SetHostDisease(v string)`

SetHostDisease sets HostDisease field to given value.

### HasHostDisease

`func (o *V2reportsBioSampleDescriptor) HasHostDisease() bool`

HasHostDisease returns a boolean if a field has been set.

### GetIdentifiedBy

`func (o *V2reportsBioSampleDescriptor) GetIdentifiedBy() string`

GetIdentifiedBy returns the IdentifiedBy field if non-nil, zero value otherwise.

### GetIdentifiedByOk

`func (o *V2reportsBioSampleDescriptor) GetIdentifiedByOk() (*string, bool)`

GetIdentifiedByOk returns a tuple with the IdentifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedBy

`func (o *V2reportsBioSampleDescriptor) SetIdentifiedBy(v string)`

SetIdentifiedBy sets IdentifiedBy field to given value.

### HasIdentifiedBy

`func (o *V2reportsBioSampleDescriptor) HasIdentifiedBy() bool`

HasIdentifiedBy returns a boolean if a field has been set.

### GetIfsacCategory

`func (o *V2reportsBioSampleDescriptor) GetIfsacCategory() string`

GetIfsacCategory returns the IfsacCategory field if non-nil, zero value otherwise.

### GetIfsacCategoryOk

`func (o *V2reportsBioSampleDescriptor) GetIfsacCategoryOk() (*string, bool)`

GetIfsacCategoryOk returns a tuple with the IfsacCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfsacCategory

`func (o *V2reportsBioSampleDescriptor) SetIfsacCategory(v string)`

SetIfsacCategory sets IfsacCategory field to given value.

### HasIfsacCategory

`func (o *V2reportsBioSampleDescriptor) HasIfsacCategory() bool`

HasIfsacCategory returns a boolean if a field has been set.

### GetIsolate

`func (o *V2reportsBioSampleDescriptor) GetIsolate() string`

GetIsolate returns the Isolate field if non-nil, zero value otherwise.

### GetIsolateOk

`func (o *V2reportsBioSampleDescriptor) GetIsolateOk() (*string, bool)`

GetIsolateOk returns a tuple with the Isolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolate

`func (o *V2reportsBioSampleDescriptor) SetIsolate(v string)`

SetIsolate sets Isolate field to given value.

### HasIsolate

`func (o *V2reportsBioSampleDescriptor) HasIsolate() bool`

HasIsolate returns a boolean if a field has been set.

### GetIsolateNameAlias

`func (o *V2reportsBioSampleDescriptor) GetIsolateNameAlias() string`

GetIsolateNameAlias returns the IsolateNameAlias field if non-nil, zero value otherwise.

### GetIsolateNameAliasOk

`func (o *V2reportsBioSampleDescriptor) GetIsolateNameAliasOk() (*string, bool)`

GetIsolateNameAliasOk returns a tuple with the IsolateNameAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolateNameAlias

`func (o *V2reportsBioSampleDescriptor) SetIsolateNameAlias(v string)`

SetIsolateNameAlias sets IsolateNameAlias field to given value.

### HasIsolateNameAlias

`func (o *V2reportsBioSampleDescriptor) HasIsolateNameAlias() bool`

HasIsolateNameAlias returns a boolean if a field has been set.

### GetIsolationSource

`func (o *V2reportsBioSampleDescriptor) GetIsolationSource() string`

GetIsolationSource returns the IsolationSource field if non-nil, zero value otherwise.

### GetIsolationSourceOk

`func (o *V2reportsBioSampleDescriptor) GetIsolationSourceOk() (*string, bool)`

GetIsolationSourceOk returns a tuple with the IsolationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationSource

`func (o *V2reportsBioSampleDescriptor) SetIsolationSource(v string)`

SetIsolationSource sets IsolationSource field to given value.

### HasIsolationSource

`func (o *V2reportsBioSampleDescriptor) HasIsolationSource() bool`

HasIsolationSource returns a boolean if a field has been set.

### GetLatLon

`func (o *V2reportsBioSampleDescriptor) GetLatLon() string`

GetLatLon returns the LatLon field if non-nil, zero value otherwise.

### GetLatLonOk

`func (o *V2reportsBioSampleDescriptor) GetLatLonOk() (*string, bool)`

GetLatLonOk returns a tuple with the LatLon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatLon

`func (o *V2reportsBioSampleDescriptor) SetLatLon(v string)`

SetLatLon sets LatLon field to given value.

### HasLatLon

`func (o *V2reportsBioSampleDescriptor) HasLatLon() bool`

HasLatLon returns a boolean if a field has been set.

### GetProjectName

`func (o *V2reportsBioSampleDescriptor) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *V2reportsBioSampleDescriptor) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *V2reportsBioSampleDescriptor) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *V2reportsBioSampleDescriptor) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetSampleName

`func (o *V2reportsBioSampleDescriptor) GetSampleName() string`

GetSampleName returns the SampleName field if non-nil, zero value otherwise.

### GetSampleNameOk

`func (o *V2reportsBioSampleDescriptor) GetSampleNameOk() (*string, bool)`

GetSampleNameOk returns a tuple with the SampleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleName

`func (o *V2reportsBioSampleDescriptor) SetSampleName(v string)`

SetSampleName sets SampleName field to given value.

### HasSampleName

`func (o *V2reportsBioSampleDescriptor) HasSampleName() bool`

HasSampleName returns a boolean if a field has been set.

### GetSerovar

`func (o *V2reportsBioSampleDescriptor) GetSerovar() string`

GetSerovar returns the Serovar field if non-nil, zero value otherwise.

### GetSerovarOk

`func (o *V2reportsBioSampleDescriptor) GetSerovarOk() (*string, bool)`

GetSerovarOk returns a tuple with the Serovar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerovar

`func (o *V2reportsBioSampleDescriptor) SetSerovar(v string)`

SetSerovar sets Serovar field to given value.

### HasSerovar

`func (o *V2reportsBioSampleDescriptor) HasSerovar() bool`

HasSerovar returns a boolean if a field has been set.

### GetSex

`func (o *V2reportsBioSampleDescriptor) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *V2reportsBioSampleDescriptor) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *V2reportsBioSampleDescriptor) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *V2reportsBioSampleDescriptor) HasSex() bool`

HasSex returns a boolean if a field has been set.

### GetSourceType

`func (o *V2reportsBioSampleDescriptor) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *V2reportsBioSampleDescriptor) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *V2reportsBioSampleDescriptor) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *V2reportsBioSampleDescriptor) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetStrain

`func (o *V2reportsBioSampleDescriptor) GetStrain() string`

GetStrain returns the Strain field if non-nil, zero value otherwise.

### GetStrainOk

`func (o *V2reportsBioSampleDescriptor) GetStrainOk() (*string, bool)`

GetStrainOk returns a tuple with the Strain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrain

`func (o *V2reportsBioSampleDescriptor) SetStrain(v string)`

SetStrain sets Strain field to given value.

### HasStrain

`func (o *V2reportsBioSampleDescriptor) HasStrain() bool`

HasStrain returns a boolean if a field has been set.

### GetSubSpecies

`func (o *V2reportsBioSampleDescriptor) GetSubSpecies() string`

GetSubSpecies returns the SubSpecies field if non-nil, zero value otherwise.

### GetSubSpeciesOk

`func (o *V2reportsBioSampleDescriptor) GetSubSpeciesOk() (*string, bool)`

GetSubSpeciesOk returns a tuple with the SubSpecies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubSpecies

`func (o *V2reportsBioSampleDescriptor) SetSubSpecies(v string)`

SetSubSpecies sets SubSpecies field to given value.

### HasSubSpecies

`func (o *V2reportsBioSampleDescriptor) HasSubSpecies() bool`

HasSubSpecies returns a boolean if a field has been set.

### GetTissue

`func (o *V2reportsBioSampleDescriptor) GetTissue() string`

GetTissue returns the Tissue field if non-nil, zero value otherwise.

### GetTissueOk

`func (o *V2reportsBioSampleDescriptor) GetTissueOk() (*string, bool)`

GetTissueOk returns a tuple with the Tissue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTissue

`func (o *V2reportsBioSampleDescriptor) SetTissue(v string)`

SetTissue sets Tissue field to given value.

### HasTissue

`func (o *V2reportsBioSampleDescriptor) HasTissue() bool`

HasTissue returns a boolean if a field has been set.

### GetSerotype

`func (o *V2reportsBioSampleDescriptor) GetSerotype() string`

GetSerotype returns the Serotype field if non-nil, zero value otherwise.

### GetSerotypeOk

`func (o *V2reportsBioSampleDescriptor) GetSerotypeOk() (*string, bool)`

GetSerotypeOk returns a tuple with the Serotype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerotype

`func (o *V2reportsBioSampleDescriptor) SetSerotype(v string)`

SetSerotype sets Serotype field to given value.

### HasSerotype

`func (o *V2reportsBioSampleDescriptor) HasSerotype() bool`

HasSerotype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


