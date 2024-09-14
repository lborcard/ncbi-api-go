# V2reportsTaxonomyNamesDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxId** | Pointer to **string** |  | [optional] 
**Rank** | Pointer to [**V2reportsRankType**](V2reportsRankType.md) |  | [optional] [default to V2REPORTSRANKTYPE_NO_RANK]
**CurrentScientificName** | Pointer to [**V2reportsNameAndAuthority**](V2reportsNameAndAuthority.md) |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**CuratorCommonName** | Pointer to **string** |  | [optional] 
**OtherCommonNames** | Pointer to **[]string** |  | [optional] 
**GeneralNotes** | Pointer to **[]string** |  | [optional] 
**LinksFromType** | Pointer to **string** |  | [optional] 
**Citations** | Pointer to [**[]V2reportsTaxonomyNamesDescriptorCitation**](V2reportsTaxonomyNamesDescriptorCitation.md) |  | [optional] 
**CurrentScientificNameIsFormal** | Pointer to **bool** |  | [optional] 

## Methods

### NewV2reportsTaxonomyNamesDescriptor

`func NewV2reportsTaxonomyNamesDescriptor() *V2reportsTaxonomyNamesDescriptor`

NewV2reportsTaxonomyNamesDescriptor instantiates a new V2reportsTaxonomyNamesDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsTaxonomyNamesDescriptorWithDefaults

`func NewV2reportsTaxonomyNamesDescriptorWithDefaults() *V2reportsTaxonomyNamesDescriptor`

NewV2reportsTaxonomyNamesDescriptorWithDefaults instantiates a new V2reportsTaxonomyNamesDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxId

`func (o *V2reportsTaxonomyNamesDescriptor) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V2reportsTaxonomyNamesDescriptor) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V2reportsTaxonomyNamesDescriptor) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V2reportsTaxonomyNamesDescriptor) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetRank

`func (o *V2reportsTaxonomyNamesDescriptor) GetRank() V2reportsRankType`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *V2reportsTaxonomyNamesDescriptor) GetRankOk() (*V2reportsRankType, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *V2reportsTaxonomyNamesDescriptor) SetRank(v V2reportsRankType)`

SetRank sets Rank field to given value.

### HasRank

`func (o *V2reportsTaxonomyNamesDescriptor) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetCurrentScientificName

`func (o *V2reportsTaxonomyNamesDescriptor) GetCurrentScientificName() V2reportsNameAndAuthority`

GetCurrentScientificName returns the CurrentScientificName field if non-nil, zero value otherwise.

### GetCurrentScientificNameOk

`func (o *V2reportsTaxonomyNamesDescriptor) GetCurrentScientificNameOk() (*V2reportsNameAndAuthority, bool)`

GetCurrentScientificNameOk returns a tuple with the CurrentScientificName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentScientificName

`func (o *V2reportsTaxonomyNamesDescriptor) SetCurrentScientificName(v V2reportsNameAndAuthority)`

SetCurrentScientificName sets CurrentScientificName field to given value.

### HasCurrentScientificName

`func (o *V2reportsTaxonomyNamesDescriptor) HasCurrentScientificName() bool`

HasCurrentScientificName returns a boolean if a field has been set.

### GetGroupName

`func (o *V2reportsTaxonomyNamesDescriptor) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *V2reportsTaxonomyNamesDescriptor) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *V2reportsTaxonomyNamesDescriptor) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *V2reportsTaxonomyNamesDescriptor) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetCuratorCommonName

`func (o *V2reportsTaxonomyNamesDescriptor) GetCuratorCommonName() string`

GetCuratorCommonName returns the CuratorCommonName field if non-nil, zero value otherwise.

### GetCuratorCommonNameOk

`func (o *V2reportsTaxonomyNamesDescriptor) GetCuratorCommonNameOk() (*string, bool)`

GetCuratorCommonNameOk returns a tuple with the CuratorCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuratorCommonName

`func (o *V2reportsTaxonomyNamesDescriptor) SetCuratorCommonName(v string)`

SetCuratorCommonName sets CuratorCommonName field to given value.

### HasCuratorCommonName

`func (o *V2reportsTaxonomyNamesDescriptor) HasCuratorCommonName() bool`

HasCuratorCommonName returns a boolean if a field has been set.

### GetOtherCommonNames

`func (o *V2reportsTaxonomyNamesDescriptor) GetOtherCommonNames() []string`

GetOtherCommonNames returns the OtherCommonNames field if non-nil, zero value otherwise.

### GetOtherCommonNamesOk

`func (o *V2reportsTaxonomyNamesDescriptor) GetOtherCommonNamesOk() (*[]string, bool)`

GetOtherCommonNamesOk returns a tuple with the OtherCommonNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCommonNames

`func (o *V2reportsTaxonomyNamesDescriptor) SetOtherCommonNames(v []string)`

SetOtherCommonNames sets OtherCommonNames field to given value.

### HasOtherCommonNames

`func (o *V2reportsTaxonomyNamesDescriptor) HasOtherCommonNames() bool`

HasOtherCommonNames returns a boolean if a field has been set.

### GetGeneralNotes

`func (o *V2reportsTaxonomyNamesDescriptor) GetGeneralNotes() []string`

GetGeneralNotes returns the GeneralNotes field if non-nil, zero value otherwise.

### GetGeneralNotesOk

`func (o *V2reportsTaxonomyNamesDescriptor) GetGeneralNotesOk() (*[]string, bool)`

GetGeneralNotesOk returns a tuple with the GeneralNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralNotes

`func (o *V2reportsTaxonomyNamesDescriptor) SetGeneralNotes(v []string)`

SetGeneralNotes sets GeneralNotes field to given value.

### HasGeneralNotes

`func (o *V2reportsTaxonomyNamesDescriptor) HasGeneralNotes() bool`

HasGeneralNotes returns a boolean if a field has been set.

### GetLinksFromType

`func (o *V2reportsTaxonomyNamesDescriptor) GetLinksFromType() string`

GetLinksFromType returns the LinksFromType field if non-nil, zero value otherwise.

### GetLinksFromTypeOk

`func (o *V2reportsTaxonomyNamesDescriptor) GetLinksFromTypeOk() (*string, bool)`

GetLinksFromTypeOk returns a tuple with the LinksFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinksFromType

`func (o *V2reportsTaxonomyNamesDescriptor) SetLinksFromType(v string)`

SetLinksFromType sets LinksFromType field to given value.

### HasLinksFromType

`func (o *V2reportsTaxonomyNamesDescriptor) HasLinksFromType() bool`

HasLinksFromType returns a boolean if a field has been set.

### GetCitations

`func (o *V2reportsTaxonomyNamesDescriptor) GetCitations() []V2reportsTaxonomyNamesDescriptorCitation`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *V2reportsTaxonomyNamesDescriptor) GetCitationsOk() (*[]V2reportsTaxonomyNamesDescriptorCitation, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *V2reportsTaxonomyNamesDescriptor) SetCitations(v []V2reportsTaxonomyNamesDescriptorCitation)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *V2reportsTaxonomyNamesDescriptor) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### GetCurrentScientificNameIsFormal

`func (o *V2reportsTaxonomyNamesDescriptor) GetCurrentScientificNameIsFormal() bool`

GetCurrentScientificNameIsFormal returns the CurrentScientificNameIsFormal field if non-nil, zero value otherwise.

### GetCurrentScientificNameIsFormalOk

`func (o *V2reportsTaxonomyNamesDescriptor) GetCurrentScientificNameIsFormalOk() (*bool, bool)`

GetCurrentScientificNameIsFormalOk returns a tuple with the CurrentScientificNameIsFormal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentScientificNameIsFormal

`func (o *V2reportsTaxonomyNamesDescriptor) SetCurrentScientificNameIsFormal(v bool)`

SetCurrentScientificNameIsFormal sets CurrentScientificNameIsFormal field to given value.

### HasCurrentScientificNameIsFormal

`func (o *V2reportsTaxonomyNamesDescriptor) HasCurrentScientificNameIsFormal() bool`

HasCurrentScientificNameIsFormal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


