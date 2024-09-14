# V2MicroBiggeDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpaqueSolrQuery** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]V2MicroBiggeDatasetRequestFileType**](V2MicroBiggeDatasetRequestFileType.md) |  | [optional] 
**ElementFlankConfig** | Pointer to [**V2ElementFlankConfig**](V2ElementFlankConfig.md) |  | [optional] 

## Methods

### NewV2MicroBiggeDatasetRequest

`func NewV2MicroBiggeDatasetRequest() *V2MicroBiggeDatasetRequest`

NewV2MicroBiggeDatasetRequest instantiates a new V2MicroBiggeDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MicroBiggeDatasetRequestWithDefaults

`func NewV2MicroBiggeDatasetRequestWithDefaults() *V2MicroBiggeDatasetRequest`

NewV2MicroBiggeDatasetRequestWithDefaults instantiates a new V2MicroBiggeDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpaqueSolrQuery

`func (o *V2MicroBiggeDatasetRequest) GetOpaqueSolrQuery() string`

GetOpaqueSolrQuery returns the OpaqueSolrQuery field if non-nil, zero value otherwise.

### GetOpaqueSolrQueryOk

`func (o *V2MicroBiggeDatasetRequest) GetOpaqueSolrQueryOk() (*string, bool)`

GetOpaqueSolrQueryOk returns a tuple with the OpaqueSolrQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaqueSolrQuery

`func (o *V2MicroBiggeDatasetRequest) SetOpaqueSolrQuery(v string)`

SetOpaqueSolrQuery sets OpaqueSolrQuery field to given value.

### HasOpaqueSolrQuery

`func (o *V2MicroBiggeDatasetRequest) HasOpaqueSolrQuery() bool`

HasOpaqueSolrQuery returns a boolean if a field has been set.

### GetFiles

`func (o *V2MicroBiggeDatasetRequest) GetFiles() []V2MicroBiggeDatasetRequestFileType`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *V2MicroBiggeDatasetRequest) GetFilesOk() (*[]V2MicroBiggeDatasetRequestFileType, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *V2MicroBiggeDatasetRequest) SetFiles(v []V2MicroBiggeDatasetRequestFileType)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *V2MicroBiggeDatasetRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetElementFlankConfig

`func (o *V2MicroBiggeDatasetRequest) GetElementFlankConfig() V2ElementFlankConfig`

GetElementFlankConfig returns the ElementFlankConfig field if non-nil, zero value otherwise.

### GetElementFlankConfigOk

`func (o *V2MicroBiggeDatasetRequest) GetElementFlankConfigOk() (*V2ElementFlankConfig, bool)`

GetElementFlankConfigOk returns a tuple with the ElementFlankConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementFlankConfig

`func (o *V2MicroBiggeDatasetRequest) SetElementFlankConfig(v V2ElementFlankConfig)`

SetElementFlankConfig sets ElementFlankConfig field to given value.

### HasElementFlankConfig

`func (o *V2MicroBiggeDatasetRequest) HasElementFlankConfig() bool`

HasElementFlankConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


