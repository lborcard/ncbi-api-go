# V2reportsGenomicRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneRange** | Pointer to [**V2reportsSeqRangeSet**](V2reportsSeqRangeSet.md) |  | [optional] 
**Type** | Pointer to [**V2reportsGenomicRegionGenomicRegionType**](V2reportsGenomicRegionGenomicRegionType.md) |  | [optional] [default to V2REPORTSGENOMICREGIONGENOMICREGIONTYPE_UNKNOWN]

## Methods

### NewV2reportsGenomicRegion

`func NewV2reportsGenomicRegion() *V2reportsGenomicRegion`

NewV2reportsGenomicRegion instantiates a new V2reportsGenomicRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsGenomicRegionWithDefaults

`func NewV2reportsGenomicRegionWithDefaults() *V2reportsGenomicRegion`

NewV2reportsGenomicRegionWithDefaults instantiates a new V2reportsGenomicRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneRange

`func (o *V2reportsGenomicRegion) GetGeneRange() V2reportsSeqRangeSet`

GetGeneRange returns the GeneRange field if non-nil, zero value otherwise.

### GetGeneRangeOk

`func (o *V2reportsGenomicRegion) GetGeneRangeOk() (*V2reportsSeqRangeSet, bool)`

GetGeneRangeOk returns a tuple with the GeneRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneRange

`func (o *V2reportsGenomicRegion) SetGeneRange(v V2reportsSeqRangeSet)`

SetGeneRange sets GeneRange field to given value.

### HasGeneRange

`func (o *V2reportsGenomicRegion) HasGeneRange() bool`

HasGeneRange returns a boolean if a field has been set.

### GetType

`func (o *V2reportsGenomicRegion) GetType() V2reportsGenomicRegionGenomicRegionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2reportsGenomicRegion) GetTypeOk() (*V2reportsGenomicRegionGenomicRegionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2reportsGenomicRegion) SetType(v V2reportsGenomicRegionGenomicRegionType)`

SetType sets Type field to given value.

### HasType

`func (o *V2reportsGenomicRegion) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


