# V2reportsTaxonomyNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxId** | Pointer to **int32** |  | [optional] 
**Rank** | Pointer to [**V2reportsRankType**](V2reportsRankType.md) |  | [optional] [default to V2REPORTSRANKTYPE_NO_RANK]
**CurrentScientificName** | Pointer to [**V2reportsNameAndAuthority**](V2reportsNameAndAuthority.md) |  | [optional] 
**Basionym** | Pointer to [**V2reportsNameAndAuthority**](V2reportsNameAndAuthority.md) |  | [optional] 
**CuratorCommonName** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**HasTypeMaterial** | Pointer to **bool** |  | [optional] 
**Classification** | Pointer to [**V2reportsClassification**](V2reportsClassification.md) |  | [optional] 
**Parents** | Pointer to **[]int32** |  | [optional] 
**Children** | Pointer to **[]int32** |  | [optional] 
**Counts** | Pointer to [**[]V2reportsTaxonomyNodeCountByType**](V2reportsTaxonomyNodeCountByType.md) |  | [optional] 
**GenomicMoltype** | Pointer to **string** |  | [optional] 

## Methods

### NewV2reportsTaxonomyNode

`func NewV2reportsTaxonomyNode() *V2reportsTaxonomyNode`

NewV2reportsTaxonomyNode instantiates a new V2reportsTaxonomyNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsTaxonomyNodeWithDefaults

`func NewV2reportsTaxonomyNodeWithDefaults() *V2reportsTaxonomyNode`

NewV2reportsTaxonomyNodeWithDefaults instantiates a new V2reportsTaxonomyNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxId

`func (o *V2reportsTaxonomyNode) GetTaxId() int32`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V2reportsTaxonomyNode) GetTaxIdOk() (*int32, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V2reportsTaxonomyNode) SetTaxId(v int32)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V2reportsTaxonomyNode) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetRank

`func (o *V2reportsTaxonomyNode) GetRank() V2reportsRankType`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *V2reportsTaxonomyNode) GetRankOk() (*V2reportsRankType, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *V2reportsTaxonomyNode) SetRank(v V2reportsRankType)`

SetRank sets Rank field to given value.

### HasRank

`func (o *V2reportsTaxonomyNode) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetCurrentScientificName

`func (o *V2reportsTaxonomyNode) GetCurrentScientificName() V2reportsNameAndAuthority`

GetCurrentScientificName returns the CurrentScientificName field if non-nil, zero value otherwise.

### GetCurrentScientificNameOk

`func (o *V2reportsTaxonomyNode) GetCurrentScientificNameOk() (*V2reportsNameAndAuthority, bool)`

GetCurrentScientificNameOk returns a tuple with the CurrentScientificName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentScientificName

`func (o *V2reportsTaxonomyNode) SetCurrentScientificName(v V2reportsNameAndAuthority)`

SetCurrentScientificName sets CurrentScientificName field to given value.

### HasCurrentScientificName

`func (o *V2reportsTaxonomyNode) HasCurrentScientificName() bool`

HasCurrentScientificName returns a boolean if a field has been set.

### GetBasionym

`func (o *V2reportsTaxonomyNode) GetBasionym() V2reportsNameAndAuthority`

GetBasionym returns the Basionym field if non-nil, zero value otherwise.

### GetBasionymOk

`func (o *V2reportsTaxonomyNode) GetBasionymOk() (*V2reportsNameAndAuthority, bool)`

GetBasionymOk returns a tuple with the Basionym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasionym

`func (o *V2reportsTaxonomyNode) SetBasionym(v V2reportsNameAndAuthority)`

SetBasionym sets Basionym field to given value.

### HasBasionym

`func (o *V2reportsTaxonomyNode) HasBasionym() bool`

HasBasionym returns a boolean if a field has been set.

### GetCuratorCommonName

`func (o *V2reportsTaxonomyNode) GetCuratorCommonName() string`

GetCuratorCommonName returns the CuratorCommonName field if non-nil, zero value otherwise.

### GetCuratorCommonNameOk

`func (o *V2reportsTaxonomyNode) GetCuratorCommonNameOk() (*string, bool)`

GetCuratorCommonNameOk returns a tuple with the CuratorCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuratorCommonName

`func (o *V2reportsTaxonomyNode) SetCuratorCommonName(v string)`

SetCuratorCommonName sets CuratorCommonName field to given value.

### HasCuratorCommonName

`func (o *V2reportsTaxonomyNode) HasCuratorCommonName() bool`

HasCuratorCommonName returns a boolean if a field has been set.

### GetGroupName

`func (o *V2reportsTaxonomyNode) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *V2reportsTaxonomyNode) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *V2reportsTaxonomyNode) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *V2reportsTaxonomyNode) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetHasTypeMaterial

`func (o *V2reportsTaxonomyNode) GetHasTypeMaterial() bool`

GetHasTypeMaterial returns the HasTypeMaterial field if non-nil, zero value otherwise.

### GetHasTypeMaterialOk

`func (o *V2reportsTaxonomyNode) GetHasTypeMaterialOk() (*bool, bool)`

GetHasTypeMaterialOk returns a tuple with the HasTypeMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTypeMaterial

`func (o *V2reportsTaxonomyNode) SetHasTypeMaterial(v bool)`

SetHasTypeMaterial sets HasTypeMaterial field to given value.

### HasHasTypeMaterial

`func (o *V2reportsTaxonomyNode) HasHasTypeMaterial() bool`

HasHasTypeMaterial returns a boolean if a field has been set.

### GetClassification

`func (o *V2reportsTaxonomyNode) GetClassification() V2reportsClassification`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *V2reportsTaxonomyNode) GetClassificationOk() (*V2reportsClassification, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *V2reportsTaxonomyNode) SetClassification(v V2reportsClassification)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *V2reportsTaxonomyNode) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetParents

`func (o *V2reportsTaxonomyNode) GetParents() []int32`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *V2reportsTaxonomyNode) GetParentsOk() (*[]int32, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *V2reportsTaxonomyNode) SetParents(v []int32)`

SetParents sets Parents field to given value.

### HasParents

`func (o *V2reportsTaxonomyNode) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetChildren

`func (o *V2reportsTaxonomyNode) GetChildren() []int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *V2reportsTaxonomyNode) GetChildrenOk() (*[]int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *V2reportsTaxonomyNode) SetChildren(v []int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *V2reportsTaxonomyNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCounts

`func (o *V2reportsTaxonomyNode) GetCounts() []V2reportsTaxonomyNodeCountByType`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *V2reportsTaxonomyNode) GetCountsOk() (*[]V2reportsTaxonomyNodeCountByType, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *V2reportsTaxonomyNode) SetCounts(v []V2reportsTaxonomyNodeCountByType)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *V2reportsTaxonomyNode) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetGenomicMoltype

`func (o *V2reportsTaxonomyNode) GetGenomicMoltype() string`

GetGenomicMoltype returns the GenomicMoltype field if non-nil, zero value otherwise.

### GetGenomicMoltypeOk

`func (o *V2reportsTaxonomyNode) GetGenomicMoltypeOk() (*string, bool)`

GetGenomicMoltypeOk returns a tuple with the GenomicMoltype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicMoltype

`func (o *V2reportsTaxonomyNode) SetGenomicMoltype(v string)`

SetGenomicMoltype sets GenomicMoltype field to given value.

### HasGenomicMoltype

`func (o *V2reportsTaxonomyNode) HasGenomicMoltype() bool`

HasGenomicMoltype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


