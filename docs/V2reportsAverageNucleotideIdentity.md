# V2reportsAverageNucleotideIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxonomyCheckStatus** | Pointer to [**V2reportsAverageNucleotideIdentityTaxonomyCheckStatus**](V2reportsAverageNucleotideIdentityTaxonomyCheckStatus.md) |  | [optional] [default to V2REPORTSAVERAGENUCLEOTIDEIDENTITYTAXONOMYCHECKSTATUS_TAXONOMY_CHECK_STATUS_UNKNOWN]
**MatchStatus** | Pointer to [**V2reportsAverageNucleotideIdentityMatchStatus**](V2reportsAverageNucleotideIdentityMatchStatus.md) |  | [optional] [default to V2REPORTSAVERAGENUCLEOTIDEIDENTITYMATCHSTATUS_BEST_MATCH_STATUS_UNKNOWN]
**SubmittedOrganism** | Pointer to **string** |  | [optional] 
**SubmittedSpecies** | Pointer to **string** |  | [optional] 
**Category** | Pointer to [**V2reportsANITypeCategory**](V2reportsANITypeCategory.md) |  | [optional] [default to V2REPORTSANITYPECATEGORY_ANI_CATEGORY_UNKNOWN]
**SubmittedAniMatch** | Pointer to [**V2reportsANIMatch**](V2reportsANIMatch.md) |  | [optional] 
**BestAniMatch** | Pointer to [**V2reportsANIMatch**](V2reportsANIMatch.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewV2reportsAverageNucleotideIdentity

`func NewV2reportsAverageNucleotideIdentity() *V2reportsAverageNucleotideIdentity`

NewV2reportsAverageNucleotideIdentity instantiates a new V2reportsAverageNucleotideIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsAverageNucleotideIdentityWithDefaults

`func NewV2reportsAverageNucleotideIdentityWithDefaults() *V2reportsAverageNucleotideIdentity`

NewV2reportsAverageNucleotideIdentityWithDefaults instantiates a new V2reportsAverageNucleotideIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxonomyCheckStatus

`func (o *V2reportsAverageNucleotideIdentity) GetTaxonomyCheckStatus() V2reportsAverageNucleotideIdentityTaxonomyCheckStatus`

GetTaxonomyCheckStatus returns the TaxonomyCheckStatus field if non-nil, zero value otherwise.

### GetTaxonomyCheckStatusOk

`func (o *V2reportsAverageNucleotideIdentity) GetTaxonomyCheckStatusOk() (*V2reportsAverageNucleotideIdentityTaxonomyCheckStatus, bool)`

GetTaxonomyCheckStatusOk returns a tuple with the TaxonomyCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyCheckStatus

`func (o *V2reportsAverageNucleotideIdentity) SetTaxonomyCheckStatus(v V2reportsAverageNucleotideIdentityTaxonomyCheckStatus)`

SetTaxonomyCheckStatus sets TaxonomyCheckStatus field to given value.

### HasTaxonomyCheckStatus

`func (o *V2reportsAverageNucleotideIdentity) HasTaxonomyCheckStatus() bool`

HasTaxonomyCheckStatus returns a boolean if a field has been set.

### GetMatchStatus

`func (o *V2reportsAverageNucleotideIdentity) GetMatchStatus() V2reportsAverageNucleotideIdentityMatchStatus`

GetMatchStatus returns the MatchStatus field if non-nil, zero value otherwise.

### GetMatchStatusOk

`func (o *V2reportsAverageNucleotideIdentity) GetMatchStatusOk() (*V2reportsAverageNucleotideIdentityMatchStatus, bool)`

GetMatchStatusOk returns a tuple with the MatchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchStatus

`func (o *V2reportsAverageNucleotideIdentity) SetMatchStatus(v V2reportsAverageNucleotideIdentityMatchStatus)`

SetMatchStatus sets MatchStatus field to given value.

### HasMatchStatus

`func (o *V2reportsAverageNucleotideIdentity) HasMatchStatus() bool`

HasMatchStatus returns a boolean if a field has been set.

### GetSubmittedOrganism

`func (o *V2reportsAverageNucleotideIdentity) GetSubmittedOrganism() string`

GetSubmittedOrganism returns the SubmittedOrganism field if non-nil, zero value otherwise.

### GetSubmittedOrganismOk

`func (o *V2reportsAverageNucleotideIdentity) GetSubmittedOrganismOk() (*string, bool)`

GetSubmittedOrganismOk returns a tuple with the SubmittedOrganism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedOrganism

`func (o *V2reportsAverageNucleotideIdentity) SetSubmittedOrganism(v string)`

SetSubmittedOrganism sets SubmittedOrganism field to given value.

### HasSubmittedOrganism

`func (o *V2reportsAverageNucleotideIdentity) HasSubmittedOrganism() bool`

HasSubmittedOrganism returns a boolean if a field has been set.

### GetSubmittedSpecies

`func (o *V2reportsAverageNucleotideIdentity) GetSubmittedSpecies() string`

GetSubmittedSpecies returns the SubmittedSpecies field if non-nil, zero value otherwise.

### GetSubmittedSpeciesOk

`func (o *V2reportsAverageNucleotideIdentity) GetSubmittedSpeciesOk() (*string, bool)`

GetSubmittedSpeciesOk returns a tuple with the SubmittedSpecies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedSpecies

`func (o *V2reportsAverageNucleotideIdentity) SetSubmittedSpecies(v string)`

SetSubmittedSpecies sets SubmittedSpecies field to given value.

### HasSubmittedSpecies

`func (o *V2reportsAverageNucleotideIdentity) HasSubmittedSpecies() bool`

HasSubmittedSpecies returns a boolean if a field has been set.

### GetCategory

`func (o *V2reportsAverageNucleotideIdentity) GetCategory() V2reportsANITypeCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *V2reportsAverageNucleotideIdentity) GetCategoryOk() (*V2reportsANITypeCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *V2reportsAverageNucleotideIdentity) SetCategory(v V2reportsANITypeCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *V2reportsAverageNucleotideIdentity) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSubmittedAniMatch

`func (o *V2reportsAverageNucleotideIdentity) GetSubmittedAniMatch() V2reportsANIMatch`

GetSubmittedAniMatch returns the SubmittedAniMatch field if non-nil, zero value otherwise.

### GetSubmittedAniMatchOk

`func (o *V2reportsAverageNucleotideIdentity) GetSubmittedAniMatchOk() (*V2reportsANIMatch, bool)`

GetSubmittedAniMatchOk returns a tuple with the SubmittedAniMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAniMatch

`func (o *V2reportsAverageNucleotideIdentity) SetSubmittedAniMatch(v V2reportsANIMatch)`

SetSubmittedAniMatch sets SubmittedAniMatch field to given value.

### HasSubmittedAniMatch

`func (o *V2reportsAverageNucleotideIdentity) HasSubmittedAniMatch() bool`

HasSubmittedAniMatch returns a boolean if a field has been set.

### GetBestAniMatch

`func (o *V2reportsAverageNucleotideIdentity) GetBestAniMatch() V2reportsANIMatch`

GetBestAniMatch returns the BestAniMatch field if non-nil, zero value otherwise.

### GetBestAniMatchOk

`func (o *V2reportsAverageNucleotideIdentity) GetBestAniMatchOk() (*V2reportsANIMatch, bool)`

GetBestAniMatchOk returns a tuple with the BestAniMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestAniMatch

`func (o *V2reportsAverageNucleotideIdentity) SetBestAniMatch(v V2reportsANIMatch)`

SetBestAniMatch sets BestAniMatch field to given value.

### HasBestAniMatch

`func (o *V2reportsAverageNucleotideIdentity) HasBestAniMatch() bool`

HasBestAniMatch returns a boolean if a field has been set.

### GetComment

`func (o *V2reportsAverageNucleotideIdentity) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V2reportsAverageNucleotideIdentity) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V2reportsAverageNucleotideIdentity) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V2reportsAverageNucleotideIdentity) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


