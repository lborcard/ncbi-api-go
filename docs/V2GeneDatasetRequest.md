# V2GeneDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneIds** | Pointer to **[]int32** |  | [optional] 
**IncludeAnnotationType** | Pointer to [**[]V2Fasta**](V2Fasta.md) |  | [optional] 
**ReturnedContent** | Pointer to [**V2GeneDatasetRequestContentType**](V2GeneDatasetRequestContentType.md) |  | [optional] [default to V2GENEDATASETREQUESTCONTENTTYPE_COMPLETE]
**FastaFilter** | Pointer to **[]string** |  | [optional] 
**AuxReport** | Pointer to [**[]V2GeneDatasetRequestGeneDatasetReportType**](V2GeneDatasetRequestGeneDatasetReportType.md) |  | [optional] 
**TableFields** | Pointer to **[]string** |  | [optional] 
**TableReportType** | Pointer to [**V2GeneDatasetRequestGeneDatasetReportType**](V2GeneDatasetRequestGeneDatasetReportType.md) |  | [optional] [default to V2GENEDATASETREQUESTGENEDATASETREPORTTYPE_DATASET_REPORT]

## Methods

### NewV2GeneDatasetRequest

`func NewV2GeneDatasetRequest() *V2GeneDatasetRequest`

NewV2GeneDatasetRequest instantiates a new V2GeneDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2GeneDatasetRequestWithDefaults

`func NewV2GeneDatasetRequestWithDefaults() *V2GeneDatasetRequest`

NewV2GeneDatasetRequestWithDefaults instantiates a new V2GeneDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneIds

`func (o *V2GeneDatasetRequest) GetGeneIds() []int32`

GetGeneIds returns the GeneIds field if non-nil, zero value otherwise.

### GetGeneIdsOk

`func (o *V2GeneDatasetRequest) GetGeneIdsOk() (*[]int32, bool)`

GetGeneIdsOk returns a tuple with the GeneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneIds

`func (o *V2GeneDatasetRequest) SetGeneIds(v []int32)`

SetGeneIds sets GeneIds field to given value.

### HasGeneIds

`func (o *V2GeneDatasetRequest) HasGeneIds() bool`

HasGeneIds returns a boolean if a field has been set.

### GetIncludeAnnotationType

`func (o *V2GeneDatasetRequest) GetIncludeAnnotationType() []V2Fasta`

GetIncludeAnnotationType returns the IncludeAnnotationType field if non-nil, zero value otherwise.

### GetIncludeAnnotationTypeOk

`func (o *V2GeneDatasetRequest) GetIncludeAnnotationTypeOk() (*[]V2Fasta, bool)`

GetIncludeAnnotationTypeOk returns a tuple with the IncludeAnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnnotationType

`func (o *V2GeneDatasetRequest) SetIncludeAnnotationType(v []V2Fasta)`

SetIncludeAnnotationType sets IncludeAnnotationType field to given value.

### HasIncludeAnnotationType

`func (o *V2GeneDatasetRequest) HasIncludeAnnotationType() bool`

HasIncludeAnnotationType returns a boolean if a field has been set.

### GetReturnedContent

`func (o *V2GeneDatasetRequest) GetReturnedContent() V2GeneDatasetRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V2GeneDatasetRequest) GetReturnedContentOk() (*V2GeneDatasetRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V2GeneDatasetRequest) SetReturnedContent(v V2GeneDatasetRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V2GeneDatasetRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.

### GetFastaFilter

`func (o *V2GeneDatasetRequest) GetFastaFilter() []string`

GetFastaFilter returns the FastaFilter field if non-nil, zero value otherwise.

### GetFastaFilterOk

`func (o *V2GeneDatasetRequest) GetFastaFilterOk() (*[]string, bool)`

GetFastaFilterOk returns a tuple with the FastaFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastaFilter

`func (o *V2GeneDatasetRequest) SetFastaFilter(v []string)`

SetFastaFilter sets FastaFilter field to given value.

### HasFastaFilter

`func (o *V2GeneDatasetRequest) HasFastaFilter() bool`

HasFastaFilter returns a boolean if a field has been set.

### GetAuxReport

`func (o *V2GeneDatasetRequest) GetAuxReport() []V2GeneDatasetRequestGeneDatasetReportType`

GetAuxReport returns the AuxReport field if non-nil, zero value otherwise.

### GetAuxReportOk

`func (o *V2GeneDatasetRequest) GetAuxReportOk() (*[]V2GeneDatasetRequestGeneDatasetReportType, bool)`

GetAuxReportOk returns a tuple with the AuxReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxReport

`func (o *V2GeneDatasetRequest) SetAuxReport(v []V2GeneDatasetRequestGeneDatasetReportType)`

SetAuxReport sets AuxReport field to given value.

### HasAuxReport

`func (o *V2GeneDatasetRequest) HasAuxReport() bool`

HasAuxReport returns a boolean if a field has been set.

### GetTableFields

`func (o *V2GeneDatasetRequest) GetTableFields() []string`

GetTableFields returns the TableFields field if non-nil, zero value otherwise.

### GetTableFieldsOk

`func (o *V2GeneDatasetRequest) GetTableFieldsOk() (*[]string, bool)`

GetTableFieldsOk returns a tuple with the TableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFields

`func (o *V2GeneDatasetRequest) SetTableFields(v []string)`

SetTableFields sets TableFields field to given value.

### HasTableFields

`func (o *V2GeneDatasetRequest) HasTableFields() bool`

HasTableFields returns a boolean if a field has been set.

### GetTableReportType

`func (o *V2GeneDatasetRequest) GetTableReportType() V2GeneDatasetRequestGeneDatasetReportType`

GetTableReportType returns the TableReportType field if non-nil, zero value otherwise.

### GetTableReportTypeOk

`func (o *V2GeneDatasetRequest) GetTableReportTypeOk() (*V2GeneDatasetRequestGeneDatasetReportType, bool)`

GetTableReportTypeOk returns a tuple with the TableReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableReportType

`func (o *V2GeneDatasetRequest) SetTableReportType(v V2GeneDatasetRequestGeneDatasetReportType)`

SetTableReportType sets TableReportType field to given value.

### HasTableReportType

`func (o *V2GeneDatasetRequest) HasTableReportType() bool`

HasTableReportType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


