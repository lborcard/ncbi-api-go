# V2TaxonomyRelatedIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxId** | Pointer to **int32** |  | [optional] 
**IncludeLineage** | Pointer to **bool** |  | [optional] 
**IncludeSubtree** | Pointer to **bool** |  | [optional] 
**Ranks** | Pointer to [**[]V2reportsRankType**](V2reportsRankType.md) |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewV2TaxonomyRelatedIdRequest

`func NewV2TaxonomyRelatedIdRequest() *V2TaxonomyRelatedIdRequest`

NewV2TaxonomyRelatedIdRequest instantiates a new V2TaxonomyRelatedIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TaxonomyRelatedIdRequestWithDefaults

`func NewV2TaxonomyRelatedIdRequestWithDefaults() *V2TaxonomyRelatedIdRequest`

NewV2TaxonomyRelatedIdRequestWithDefaults instantiates a new V2TaxonomyRelatedIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxId

`func (o *V2TaxonomyRelatedIdRequest) GetTaxId() int32`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V2TaxonomyRelatedIdRequest) GetTaxIdOk() (*int32, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V2TaxonomyRelatedIdRequest) SetTaxId(v int32)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V2TaxonomyRelatedIdRequest) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetIncludeLineage

`func (o *V2TaxonomyRelatedIdRequest) GetIncludeLineage() bool`

GetIncludeLineage returns the IncludeLineage field if non-nil, zero value otherwise.

### GetIncludeLineageOk

`func (o *V2TaxonomyRelatedIdRequest) GetIncludeLineageOk() (*bool, bool)`

GetIncludeLineageOk returns a tuple with the IncludeLineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLineage

`func (o *V2TaxonomyRelatedIdRequest) SetIncludeLineage(v bool)`

SetIncludeLineage sets IncludeLineage field to given value.

### HasIncludeLineage

`func (o *V2TaxonomyRelatedIdRequest) HasIncludeLineage() bool`

HasIncludeLineage returns a boolean if a field has been set.

### GetIncludeSubtree

`func (o *V2TaxonomyRelatedIdRequest) GetIncludeSubtree() bool`

GetIncludeSubtree returns the IncludeSubtree field if non-nil, zero value otherwise.

### GetIncludeSubtreeOk

`func (o *V2TaxonomyRelatedIdRequest) GetIncludeSubtreeOk() (*bool, bool)`

GetIncludeSubtreeOk returns a tuple with the IncludeSubtree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubtree

`func (o *V2TaxonomyRelatedIdRequest) SetIncludeSubtree(v bool)`

SetIncludeSubtree sets IncludeSubtree field to given value.

### HasIncludeSubtree

`func (o *V2TaxonomyRelatedIdRequest) HasIncludeSubtree() bool`

HasIncludeSubtree returns a boolean if a field has been set.

### GetRanks

`func (o *V2TaxonomyRelatedIdRequest) GetRanks() []V2reportsRankType`

GetRanks returns the Ranks field if non-nil, zero value otherwise.

### GetRanksOk

`func (o *V2TaxonomyRelatedIdRequest) GetRanksOk() (*[]V2reportsRankType, bool)`

GetRanksOk returns a tuple with the Ranks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanks

`func (o *V2TaxonomyRelatedIdRequest) SetRanks(v []V2reportsRankType)`

SetRanks sets Ranks field to given value.

### HasRanks

`func (o *V2TaxonomyRelatedIdRequest) HasRanks() bool`

HasRanks returns a boolean if a field has been set.

### GetPageSize

`func (o *V2TaxonomyRelatedIdRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V2TaxonomyRelatedIdRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V2TaxonomyRelatedIdRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *V2TaxonomyRelatedIdRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *V2TaxonomyRelatedIdRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V2TaxonomyRelatedIdRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V2TaxonomyRelatedIdRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V2TaxonomyRelatedIdRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


