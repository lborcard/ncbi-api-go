# V2RefGeneCatalogDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpaqueSolrQuery** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]V2RefGeneCatalogDatasetRequestFileType**](V2RefGeneCatalogDatasetRequestFileType.md) |  | [optional] 
**ElementFlankConfig** | Pointer to [**V2ElementFlankConfig**](V2ElementFlankConfig.md) |  | [optional] 

## Methods

### NewV2RefGeneCatalogDatasetRequest

`func NewV2RefGeneCatalogDatasetRequest() *V2RefGeneCatalogDatasetRequest`

NewV2RefGeneCatalogDatasetRequest instantiates a new V2RefGeneCatalogDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RefGeneCatalogDatasetRequestWithDefaults

`func NewV2RefGeneCatalogDatasetRequestWithDefaults() *V2RefGeneCatalogDatasetRequest`

NewV2RefGeneCatalogDatasetRequestWithDefaults instantiates a new V2RefGeneCatalogDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpaqueSolrQuery

`func (o *V2RefGeneCatalogDatasetRequest) GetOpaqueSolrQuery() string`

GetOpaqueSolrQuery returns the OpaqueSolrQuery field if non-nil, zero value otherwise.

### GetOpaqueSolrQueryOk

`func (o *V2RefGeneCatalogDatasetRequest) GetOpaqueSolrQueryOk() (*string, bool)`

GetOpaqueSolrQueryOk returns a tuple with the OpaqueSolrQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaqueSolrQuery

`func (o *V2RefGeneCatalogDatasetRequest) SetOpaqueSolrQuery(v string)`

SetOpaqueSolrQuery sets OpaqueSolrQuery field to given value.

### HasOpaqueSolrQuery

`func (o *V2RefGeneCatalogDatasetRequest) HasOpaqueSolrQuery() bool`

HasOpaqueSolrQuery returns a boolean if a field has been set.

### GetFiles

`func (o *V2RefGeneCatalogDatasetRequest) GetFiles() []V2RefGeneCatalogDatasetRequestFileType`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *V2RefGeneCatalogDatasetRequest) GetFilesOk() (*[]V2RefGeneCatalogDatasetRequestFileType, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *V2RefGeneCatalogDatasetRequest) SetFiles(v []V2RefGeneCatalogDatasetRequestFileType)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *V2RefGeneCatalogDatasetRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetElementFlankConfig

`func (o *V2RefGeneCatalogDatasetRequest) GetElementFlankConfig() V2ElementFlankConfig`

GetElementFlankConfig returns the ElementFlankConfig field if non-nil, zero value otherwise.

### GetElementFlankConfigOk

`func (o *V2RefGeneCatalogDatasetRequest) GetElementFlankConfigOk() (*V2ElementFlankConfig, bool)`

GetElementFlankConfigOk returns a tuple with the ElementFlankConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementFlankConfig

`func (o *V2RefGeneCatalogDatasetRequest) SetElementFlankConfig(v V2ElementFlankConfig)`

SetElementFlankConfig sets ElementFlankConfig field to given value.

### HasElementFlankConfig

`func (o *V2RefGeneCatalogDatasetRequest) HasElementFlankConfig() bool`

HasElementFlankConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


