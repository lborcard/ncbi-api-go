# V2reportsAssemblyStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNumberOfChromosomes** | Pointer to **int32** |  | [optional] 
**TotalSequenceLength** | Pointer to **string** |  | [optional] 
**TotalUngappedLength** | Pointer to **string** |  | [optional] 
**NumberOfContigs** | Pointer to **int32** |  | [optional] 
**ContigN50** | Pointer to **int32** |  | [optional] 
**ContigL50** | Pointer to **int32** |  | [optional] 
**NumberOfScaffolds** | Pointer to **int32** |  | [optional] 
**ScaffoldN50** | Pointer to **int32** |  | [optional] 
**ScaffoldL50** | Pointer to **int32** |  | [optional] 
**GapsBetweenScaffoldsCount** | Pointer to **int32** |  | [optional] 
**NumberOfComponentSequences** | Pointer to **int32** |  | [optional] 
**GcCount** | Pointer to **string** |  | [optional] 
**GcPercent** | Pointer to **float32** |  | [optional] 
**GenomeCoverage** | Pointer to **string** |  | [optional] 
**NumberOfOrganelles** | Pointer to **int32** |  | [optional] 

## Methods

### NewV2reportsAssemblyStats

`func NewV2reportsAssemblyStats() *V2reportsAssemblyStats`

NewV2reportsAssemblyStats instantiates a new V2reportsAssemblyStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsAssemblyStatsWithDefaults

`func NewV2reportsAssemblyStatsWithDefaults() *V2reportsAssemblyStats`

NewV2reportsAssemblyStatsWithDefaults instantiates a new V2reportsAssemblyStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNumberOfChromosomes

`func (o *V2reportsAssemblyStats) GetTotalNumberOfChromosomes() int32`

GetTotalNumberOfChromosomes returns the TotalNumberOfChromosomes field if non-nil, zero value otherwise.

### GetTotalNumberOfChromosomesOk

`func (o *V2reportsAssemblyStats) GetTotalNumberOfChromosomesOk() (*int32, bool)`

GetTotalNumberOfChromosomesOk returns a tuple with the TotalNumberOfChromosomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfChromosomes

`func (o *V2reportsAssemblyStats) SetTotalNumberOfChromosomes(v int32)`

SetTotalNumberOfChromosomes sets TotalNumberOfChromosomes field to given value.

### HasTotalNumberOfChromosomes

`func (o *V2reportsAssemblyStats) HasTotalNumberOfChromosomes() bool`

HasTotalNumberOfChromosomes returns a boolean if a field has been set.

### GetTotalSequenceLength

`func (o *V2reportsAssemblyStats) GetTotalSequenceLength() string`

GetTotalSequenceLength returns the TotalSequenceLength field if non-nil, zero value otherwise.

### GetTotalSequenceLengthOk

`func (o *V2reportsAssemblyStats) GetTotalSequenceLengthOk() (*string, bool)`

GetTotalSequenceLengthOk returns a tuple with the TotalSequenceLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSequenceLength

`func (o *V2reportsAssemblyStats) SetTotalSequenceLength(v string)`

SetTotalSequenceLength sets TotalSequenceLength field to given value.

### HasTotalSequenceLength

`func (o *V2reportsAssemblyStats) HasTotalSequenceLength() bool`

HasTotalSequenceLength returns a boolean if a field has been set.

### GetTotalUngappedLength

`func (o *V2reportsAssemblyStats) GetTotalUngappedLength() string`

GetTotalUngappedLength returns the TotalUngappedLength field if non-nil, zero value otherwise.

### GetTotalUngappedLengthOk

`func (o *V2reportsAssemblyStats) GetTotalUngappedLengthOk() (*string, bool)`

GetTotalUngappedLengthOk returns a tuple with the TotalUngappedLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUngappedLength

`func (o *V2reportsAssemblyStats) SetTotalUngappedLength(v string)`

SetTotalUngappedLength sets TotalUngappedLength field to given value.

### HasTotalUngappedLength

`func (o *V2reportsAssemblyStats) HasTotalUngappedLength() bool`

HasTotalUngappedLength returns a boolean if a field has been set.

### GetNumberOfContigs

`func (o *V2reportsAssemblyStats) GetNumberOfContigs() int32`

GetNumberOfContigs returns the NumberOfContigs field if non-nil, zero value otherwise.

### GetNumberOfContigsOk

`func (o *V2reportsAssemblyStats) GetNumberOfContigsOk() (*int32, bool)`

GetNumberOfContigsOk returns a tuple with the NumberOfContigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfContigs

`func (o *V2reportsAssemblyStats) SetNumberOfContigs(v int32)`

SetNumberOfContigs sets NumberOfContigs field to given value.

### HasNumberOfContigs

`func (o *V2reportsAssemblyStats) HasNumberOfContigs() bool`

HasNumberOfContigs returns a boolean if a field has been set.

### GetContigN50

`func (o *V2reportsAssemblyStats) GetContigN50() int32`

GetContigN50 returns the ContigN50 field if non-nil, zero value otherwise.

### GetContigN50Ok

`func (o *V2reportsAssemblyStats) GetContigN50Ok() (*int32, bool)`

GetContigN50Ok returns a tuple with the ContigN50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContigN50

`func (o *V2reportsAssemblyStats) SetContigN50(v int32)`

SetContigN50 sets ContigN50 field to given value.

### HasContigN50

`func (o *V2reportsAssemblyStats) HasContigN50() bool`

HasContigN50 returns a boolean if a field has been set.

### GetContigL50

`func (o *V2reportsAssemblyStats) GetContigL50() int32`

GetContigL50 returns the ContigL50 field if non-nil, zero value otherwise.

### GetContigL50Ok

`func (o *V2reportsAssemblyStats) GetContigL50Ok() (*int32, bool)`

GetContigL50Ok returns a tuple with the ContigL50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContigL50

`func (o *V2reportsAssemblyStats) SetContigL50(v int32)`

SetContigL50 sets ContigL50 field to given value.

### HasContigL50

`func (o *V2reportsAssemblyStats) HasContigL50() bool`

HasContigL50 returns a boolean if a field has been set.

### GetNumberOfScaffolds

`func (o *V2reportsAssemblyStats) GetNumberOfScaffolds() int32`

GetNumberOfScaffolds returns the NumberOfScaffolds field if non-nil, zero value otherwise.

### GetNumberOfScaffoldsOk

`func (o *V2reportsAssemblyStats) GetNumberOfScaffoldsOk() (*int32, bool)`

GetNumberOfScaffoldsOk returns a tuple with the NumberOfScaffolds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfScaffolds

`func (o *V2reportsAssemblyStats) SetNumberOfScaffolds(v int32)`

SetNumberOfScaffolds sets NumberOfScaffolds field to given value.

### HasNumberOfScaffolds

`func (o *V2reportsAssemblyStats) HasNumberOfScaffolds() bool`

HasNumberOfScaffolds returns a boolean if a field has been set.

### GetScaffoldN50

`func (o *V2reportsAssemblyStats) GetScaffoldN50() int32`

GetScaffoldN50 returns the ScaffoldN50 field if non-nil, zero value otherwise.

### GetScaffoldN50Ok

`func (o *V2reportsAssemblyStats) GetScaffoldN50Ok() (*int32, bool)`

GetScaffoldN50Ok returns a tuple with the ScaffoldN50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaffoldN50

`func (o *V2reportsAssemblyStats) SetScaffoldN50(v int32)`

SetScaffoldN50 sets ScaffoldN50 field to given value.

### HasScaffoldN50

`func (o *V2reportsAssemblyStats) HasScaffoldN50() bool`

HasScaffoldN50 returns a boolean if a field has been set.

### GetScaffoldL50

`func (o *V2reportsAssemblyStats) GetScaffoldL50() int32`

GetScaffoldL50 returns the ScaffoldL50 field if non-nil, zero value otherwise.

### GetScaffoldL50Ok

`func (o *V2reportsAssemblyStats) GetScaffoldL50Ok() (*int32, bool)`

GetScaffoldL50Ok returns a tuple with the ScaffoldL50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaffoldL50

`func (o *V2reportsAssemblyStats) SetScaffoldL50(v int32)`

SetScaffoldL50 sets ScaffoldL50 field to given value.

### HasScaffoldL50

`func (o *V2reportsAssemblyStats) HasScaffoldL50() bool`

HasScaffoldL50 returns a boolean if a field has been set.

### GetGapsBetweenScaffoldsCount

`func (o *V2reportsAssemblyStats) GetGapsBetweenScaffoldsCount() int32`

GetGapsBetweenScaffoldsCount returns the GapsBetweenScaffoldsCount field if non-nil, zero value otherwise.

### GetGapsBetweenScaffoldsCountOk

`func (o *V2reportsAssemblyStats) GetGapsBetweenScaffoldsCountOk() (*int32, bool)`

GetGapsBetweenScaffoldsCountOk returns a tuple with the GapsBetweenScaffoldsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGapsBetweenScaffoldsCount

`func (o *V2reportsAssemblyStats) SetGapsBetweenScaffoldsCount(v int32)`

SetGapsBetweenScaffoldsCount sets GapsBetweenScaffoldsCount field to given value.

### HasGapsBetweenScaffoldsCount

`func (o *V2reportsAssemblyStats) HasGapsBetweenScaffoldsCount() bool`

HasGapsBetweenScaffoldsCount returns a boolean if a field has been set.

### GetNumberOfComponentSequences

`func (o *V2reportsAssemblyStats) GetNumberOfComponentSequences() int32`

GetNumberOfComponentSequences returns the NumberOfComponentSequences field if non-nil, zero value otherwise.

### GetNumberOfComponentSequencesOk

`func (o *V2reportsAssemblyStats) GetNumberOfComponentSequencesOk() (*int32, bool)`

GetNumberOfComponentSequencesOk returns a tuple with the NumberOfComponentSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfComponentSequences

`func (o *V2reportsAssemblyStats) SetNumberOfComponentSequences(v int32)`

SetNumberOfComponentSequences sets NumberOfComponentSequences field to given value.

### HasNumberOfComponentSequences

`func (o *V2reportsAssemblyStats) HasNumberOfComponentSequences() bool`

HasNumberOfComponentSequences returns a boolean if a field has been set.

### GetGcCount

`func (o *V2reportsAssemblyStats) GetGcCount() string`

GetGcCount returns the GcCount field if non-nil, zero value otherwise.

### GetGcCountOk

`func (o *V2reportsAssemblyStats) GetGcCountOk() (*string, bool)`

GetGcCountOk returns a tuple with the GcCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcCount

`func (o *V2reportsAssemblyStats) SetGcCount(v string)`

SetGcCount sets GcCount field to given value.

### HasGcCount

`func (o *V2reportsAssemblyStats) HasGcCount() bool`

HasGcCount returns a boolean if a field has been set.

### GetGcPercent

`func (o *V2reportsAssemblyStats) GetGcPercent() float32`

GetGcPercent returns the GcPercent field if non-nil, zero value otherwise.

### GetGcPercentOk

`func (o *V2reportsAssemblyStats) GetGcPercentOk() (*float32, bool)`

GetGcPercentOk returns a tuple with the GcPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcPercent

`func (o *V2reportsAssemblyStats) SetGcPercent(v float32)`

SetGcPercent sets GcPercent field to given value.

### HasGcPercent

`func (o *V2reportsAssemblyStats) HasGcPercent() bool`

HasGcPercent returns a boolean if a field has been set.

### GetGenomeCoverage

`func (o *V2reportsAssemblyStats) GetGenomeCoverage() string`

GetGenomeCoverage returns the GenomeCoverage field if non-nil, zero value otherwise.

### GetGenomeCoverageOk

`func (o *V2reportsAssemblyStats) GetGenomeCoverageOk() (*string, bool)`

GetGenomeCoverageOk returns a tuple with the GenomeCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomeCoverage

`func (o *V2reportsAssemblyStats) SetGenomeCoverage(v string)`

SetGenomeCoverage sets GenomeCoverage field to given value.

### HasGenomeCoverage

`func (o *V2reportsAssemblyStats) HasGenomeCoverage() bool`

HasGenomeCoverage returns a boolean if a field has been set.

### GetNumberOfOrganelles

`func (o *V2reportsAssemblyStats) GetNumberOfOrganelles() int32`

GetNumberOfOrganelles returns the NumberOfOrganelles field if non-nil, zero value otherwise.

### GetNumberOfOrganellesOk

`func (o *V2reportsAssemblyStats) GetNumberOfOrganellesOk() (*int32, bool)`

GetNumberOfOrganellesOk returns a tuple with the NumberOfOrganelles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOrganelles

`func (o *V2reportsAssemblyStats) SetNumberOfOrganelles(v int32)`

SetNumberOfOrganelles sets NumberOfOrganelles field to given value.

### HasNumberOfOrganelles

`func (o *V2reportsAssemblyStats) HasNumberOfOrganelles() bool`

HasNumberOfOrganelles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


