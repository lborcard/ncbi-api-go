# V2DatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenomeV2** | Pointer to [**V2AssemblyDatasetRequest**](V2AssemblyDatasetRequest.md) |  | [optional] 
**GeneV2** | Pointer to [**V2GeneDatasetRequest**](V2GeneDatasetRequest.md) |  | [optional] 
**VirusV2** | Pointer to [**V2VirusDatasetRequest**](V2VirusDatasetRequest.md) |  | [optional] 

## Methods

### NewV2DatasetRequest

`func NewV2DatasetRequest() *V2DatasetRequest`

NewV2DatasetRequest instantiates a new V2DatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2DatasetRequestWithDefaults

`func NewV2DatasetRequestWithDefaults() *V2DatasetRequest`

NewV2DatasetRequestWithDefaults instantiates a new V2DatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenomeV2

`func (o *V2DatasetRequest) GetGenomeV2() V2AssemblyDatasetRequest`

GetGenomeV2 returns the GenomeV2 field if non-nil, zero value otherwise.

### GetGenomeV2Ok

`func (o *V2DatasetRequest) GetGenomeV2Ok() (*V2AssemblyDatasetRequest, bool)`

GetGenomeV2Ok returns a tuple with the GenomeV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomeV2

`func (o *V2DatasetRequest) SetGenomeV2(v V2AssemblyDatasetRequest)`

SetGenomeV2 sets GenomeV2 field to given value.

### HasGenomeV2

`func (o *V2DatasetRequest) HasGenomeV2() bool`

HasGenomeV2 returns a boolean if a field has been set.

### GetGeneV2

`func (o *V2DatasetRequest) GetGeneV2() V2GeneDatasetRequest`

GetGeneV2 returns the GeneV2 field if non-nil, zero value otherwise.

### GetGeneV2Ok

`func (o *V2DatasetRequest) GetGeneV2Ok() (*V2GeneDatasetRequest, bool)`

GetGeneV2Ok returns a tuple with the GeneV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneV2

`func (o *V2DatasetRequest) SetGeneV2(v V2GeneDatasetRequest)`

SetGeneV2 sets GeneV2 field to given value.

### HasGeneV2

`func (o *V2DatasetRequest) HasGeneV2() bool`

HasGeneV2 returns a boolean if a field has been set.

### GetVirusV2

`func (o *V2DatasetRequest) GetVirusV2() V2VirusDatasetRequest`

GetVirusV2 returns the VirusV2 field if non-nil, zero value otherwise.

### GetVirusV2Ok

`func (o *V2DatasetRequest) GetVirusV2Ok() (*V2VirusDatasetRequest, bool)`

GetVirusV2Ok returns a tuple with the VirusV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirusV2

`func (o *V2DatasetRequest) SetVirusV2(v V2VirusDatasetRequest)`

SetVirusV2 sets VirusV2 field to given value.

### HasVirusV2

`func (o *V2DatasetRequest) HasVirusV2() bool`

HasVirusV2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


