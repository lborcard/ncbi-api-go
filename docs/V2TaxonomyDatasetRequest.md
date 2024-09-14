# V2TaxonomyDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxIds** | Pointer to **[]int32** |  | [optional] 
**AuxReports** | Pointer to [**[]V2TaxonomyDatasetRequestTaxonomyReportType**](V2TaxonomyDatasetRequestTaxonomyReportType.md) |  | [optional] 

## Methods

### NewV2TaxonomyDatasetRequest

`func NewV2TaxonomyDatasetRequest() *V2TaxonomyDatasetRequest`

NewV2TaxonomyDatasetRequest instantiates a new V2TaxonomyDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TaxonomyDatasetRequestWithDefaults

`func NewV2TaxonomyDatasetRequestWithDefaults() *V2TaxonomyDatasetRequest`

NewV2TaxonomyDatasetRequestWithDefaults instantiates a new V2TaxonomyDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxIds

`func (o *V2TaxonomyDatasetRequest) GetTaxIds() []int32`

GetTaxIds returns the TaxIds field if non-nil, zero value otherwise.

### GetTaxIdsOk

`func (o *V2TaxonomyDatasetRequest) GetTaxIdsOk() (*[]int32, bool)`

GetTaxIdsOk returns a tuple with the TaxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIds

`func (o *V2TaxonomyDatasetRequest) SetTaxIds(v []int32)`

SetTaxIds sets TaxIds field to given value.

### HasTaxIds

`func (o *V2TaxonomyDatasetRequest) HasTaxIds() bool`

HasTaxIds returns a boolean if a field has been set.

### GetAuxReports

`func (o *V2TaxonomyDatasetRequest) GetAuxReports() []V2TaxonomyDatasetRequestTaxonomyReportType`

GetAuxReports returns the AuxReports field if non-nil, zero value otherwise.

### GetAuxReportsOk

`func (o *V2TaxonomyDatasetRequest) GetAuxReportsOk() (*[]V2TaxonomyDatasetRequestTaxonomyReportType, bool)`

GetAuxReportsOk returns a tuple with the AuxReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxReports

`func (o *V2TaxonomyDatasetRequest) SetAuxReports(v []V2TaxonomyDatasetRequestTaxonomyReportType)`

SetAuxReports sets AuxReports field to given value.

### HasAuxReports

`func (o *V2TaxonomyDatasetRequest) HasAuxReports() bool`

HasAuxReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


