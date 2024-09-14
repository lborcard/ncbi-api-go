# V2reportsOrganelle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**V2reportsOrganelleType**](V2reportsOrganelleType.md) |  | [optional] [default to V2REPORTSORGANELLETYPE_ORGANELLE_TYPE_UNKNOWN]
**Genbank** | Pointer to [**V2reportsSequenceInformation**](V2reportsSequenceInformation.md) |  | [optional] 
**Refseq** | Pointer to [**V2reportsSequenceInformation**](V2reportsSequenceInformation.md) |  | [optional] 
**Organism** | Pointer to [**V2reportsOrganism**](V2reportsOrganism.md) |  | [optional] 
**Bioprojects** | Pointer to [**[]V2reportsBioProject**](V2reportsBioProject.md) |  | [optional] 
**Biosample** | Pointer to [**V2reportsOrganelleBiosample**](V2reportsOrganelleBiosample.md) |  | [optional] 
**GeneCounts** | Pointer to [**V2reportsOrganelleGeneCounts**](V2reportsOrganelleGeneCounts.md) |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**Topology** | Pointer to [**V2reportsOrganelleTopology**](V2reportsOrganelleTopology.md) |  | [optional] [default to V2REPORTSORGANELLETOPOLOGY_TOPOLOGY_UNKNOWN]
**GeneCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewV2reportsOrganelle

`func NewV2reportsOrganelle() *V2reportsOrganelle`

NewV2reportsOrganelle instantiates a new V2reportsOrganelle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsOrganelleWithDefaults

`func NewV2reportsOrganelleWithDefaults() *V2reportsOrganelle`

NewV2reportsOrganelleWithDefaults instantiates a new V2reportsOrganelle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V2reportsOrganelle) GetDescription() V2reportsOrganelleType`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2reportsOrganelle) GetDescriptionOk() (*V2reportsOrganelleType, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2reportsOrganelle) SetDescription(v V2reportsOrganelleType)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V2reportsOrganelle) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGenbank

`func (o *V2reportsOrganelle) GetGenbank() V2reportsSequenceInformation`

GetGenbank returns the Genbank field if non-nil, zero value otherwise.

### GetGenbankOk

`func (o *V2reportsOrganelle) GetGenbankOk() (*V2reportsSequenceInformation, bool)`

GetGenbankOk returns a tuple with the Genbank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenbank

`func (o *V2reportsOrganelle) SetGenbank(v V2reportsSequenceInformation)`

SetGenbank sets Genbank field to given value.

### HasGenbank

`func (o *V2reportsOrganelle) HasGenbank() bool`

HasGenbank returns a boolean if a field has been set.

### GetRefseq

`func (o *V2reportsOrganelle) GetRefseq() V2reportsSequenceInformation`

GetRefseq returns the Refseq field if non-nil, zero value otherwise.

### GetRefseqOk

`func (o *V2reportsOrganelle) GetRefseqOk() (*V2reportsSequenceInformation, bool)`

GetRefseqOk returns a tuple with the Refseq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseq

`func (o *V2reportsOrganelle) SetRefseq(v V2reportsSequenceInformation)`

SetRefseq sets Refseq field to given value.

### HasRefseq

`func (o *V2reportsOrganelle) HasRefseq() bool`

HasRefseq returns a boolean if a field has been set.

### GetOrganism

`func (o *V2reportsOrganelle) GetOrganism() V2reportsOrganism`

GetOrganism returns the Organism field if non-nil, zero value otherwise.

### GetOrganismOk

`func (o *V2reportsOrganelle) GetOrganismOk() (*V2reportsOrganism, bool)`

GetOrganismOk returns a tuple with the Organism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganism

`func (o *V2reportsOrganelle) SetOrganism(v V2reportsOrganism)`

SetOrganism sets Organism field to given value.

### HasOrganism

`func (o *V2reportsOrganelle) HasOrganism() bool`

HasOrganism returns a boolean if a field has been set.

### GetBioprojects

`func (o *V2reportsOrganelle) GetBioprojects() []V2reportsBioProject`

GetBioprojects returns the Bioprojects field if non-nil, zero value otherwise.

### GetBioprojectsOk

`func (o *V2reportsOrganelle) GetBioprojectsOk() (*[]V2reportsBioProject, bool)`

GetBioprojectsOk returns a tuple with the Bioprojects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioprojects

`func (o *V2reportsOrganelle) SetBioprojects(v []V2reportsBioProject)`

SetBioprojects sets Bioprojects field to given value.

### HasBioprojects

`func (o *V2reportsOrganelle) HasBioprojects() bool`

HasBioprojects returns a boolean if a field has been set.

### GetBiosample

`func (o *V2reportsOrganelle) GetBiosample() V2reportsOrganelleBiosample`

GetBiosample returns the Biosample field if non-nil, zero value otherwise.

### GetBiosampleOk

`func (o *V2reportsOrganelle) GetBiosampleOk() (*V2reportsOrganelleBiosample, bool)`

GetBiosampleOk returns a tuple with the Biosample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosample

`func (o *V2reportsOrganelle) SetBiosample(v V2reportsOrganelleBiosample)`

SetBiosample sets Biosample field to given value.

### HasBiosample

`func (o *V2reportsOrganelle) HasBiosample() bool`

HasBiosample returns a boolean if a field has been set.

### GetGeneCounts

`func (o *V2reportsOrganelle) GetGeneCounts() V2reportsOrganelleGeneCounts`

GetGeneCounts returns the GeneCounts field if non-nil, zero value otherwise.

### GetGeneCountsOk

`func (o *V2reportsOrganelle) GetGeneCountsOk() (*V2reportsOrganelleGeneCounts, bool)`

GetGeneCountsOk returns a tuple with the GeneCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneCounts

`func (o *V2reportsOrganelle) SetGeneCounts(v V2reportsOrganelleGeneCounts)`

SetGeneCounts sets GeneCounts field to given value.

### HasGeneCounts

`func (o *V2reportsOrganelle) HasGeneCounts() bool`

HasGeneCounts returns a boolean if a field has been set.

### GetLength

`func (o *V2reportsOrganelle) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *V2reportsOrganelle) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *V2reportsOrganelle) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *V2reportsOrganelle) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetTopology

`func (o *V2reportsOrganelle) GetTopology() V2reportsOrganelleTopology`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *V2reportsOrganelle) GetTopologyOk() (*V2reportsOrganelleTopology, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *V2reportsOrganelle) SetTopology(v V2reportsOrganelleTopology)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *V2reportsOrganelle) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetGeneCount

`func (o *V2reportsOrganelle) GetGeneCount() int32`

GetGeneCount returns the GeneCount field if non-nil, zero value otherwise.

### GetGeneCountOk

`func (o *V2reportsOrganelle) GetGeneCountOk() (*int32, bool)`

GetGeneCountOk returns a tuple with the GeneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneCount

`func (o *V2reportsOrganelle) SetGeneCount(v int32)`

SetGeneCount sets GeneCount field to given value.

### HasGeneCount

`func (o *V2reportsOrganelle) HasGeneCount() bool`

HasGeneCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


