# V2GeneDatasetReportsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnedContent** | Pointer to [**V2GeneDatasetReportsRequestContentType**](V2GeneDatasetReportsRequestContentType.md) |  | [optional] [default to V2GENEDATASETREPORTSREQUESTCONTENTTYPE_COMPLETE]
**GeneIds** | Pointer to **[]int32** |  | [optional] 
**Accessions** | Pointer to **[]string** |  | [optional] 
**SymbolsForTaxon** | Pointer to [**V2GeneDatasetReportsRequestSymbolsForTaxon**](V2GeneDatasetReportsRequestSymbolsForTaxon.md) |  | [optional] 
**Taxon** | Pointer to **string** |  | [optional] 
**TableFields** | Pointer to **[]string** |  | [optional] 
**TableFormat** | Pointer to **string** |  | [optional] 
**IncludeTabularHeader** | Pointer to [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) |  | [optional] [default to V2INCLUDETABULARHEADER_FIRST_PAGE_ONLY]
**PageSize** | Pointer to **int32** |  | [optional] 
**PageToken** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Types** | Pointer to [**[]V2GeneType**](V2GeneType.md) |  | [optional] 

## Methods

### NewV2GeneDatasetReportsRequest

`func NewV2GeneDatasetReportsRequest() *V2GeneDatasetReportsRequest`

NewV2GeneDatasetReportsRequest instantiates a new V2GeneDatasetReportsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2GeneDatasetReportsRequestWithDefaults

`func NewV2GeneDatasetReportsRequestWithDefaults() *V2GeneDatasetReportsRequest`

NewV2GeneDatasetReportsRequestWithDefaults instantiates a new V2GeneDatasetReportsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnedContent

`func (o *V2GeneDatasetReportsRequest) GetReturnedContent() V2GeneDatasetReportsRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V2GeneDatasetReportsRequest) GetReturnedContentOk() (*V2GeneDatasetReportsRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V2GeneDatasetReportsRequest) SetReturnedContent(v V2GeneDatasetReportsRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V2GeneDatasetReportsRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.

### GetGeneIds

`func (o *V2GeneDatasetReportsRequest) GetGeneIds() []int32`

GetGeneIds returns the GeneIds field if non-nil, zero value otherwise.

### GetGeneIdsOk

`func (o *V2GeneDatasetReportsRequest) GetGeneIdsOk() (*[]int32, bool)`

GetGeneIdsOk returns a tuple with the GeneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneIds

`func (o *V2GeneDatasetReportsRequest) SetGeneIds(v []int32)`

SetGeneIds sets GeneIds field to given value.

### HasGeneIds

`func (o *V2GeneDatasetReportsRequest) HasGeneIds() bool`

HasGeneIds returns a boolean if a field has been set.

### GetAccessions

`func (o *V2GeneDatasetReportsRequest) GetAccessions() []string`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V2GeneDatasetReportsRequest) GetAccessionsOk() (*[]string, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V2GeneDatasetReportsRequest) SetAccessions(v []string)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V2GeneDatasetReportsRequest) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetSymbolsForTaxon

`func (o *V2GeneDatasetReportsRequest) GetSymbolsForTaxon() V2GeneDatasetReportsRequestSymbolsForTaxon`

GetSymbolsForTaxon returns the SymbolsForTaxon field if non-nil, zero value otherwise.

### GetSymbolsForTaxonOk

`func (o *V2GeneDatasetReportsRequest) GetSymbolsForTaxonOk() (*V2GeneDatasetReportsRequestSymbolsForTaxon, bool)`

GetSymbolsForTaxonOk returns a tuple with the SymbolsForTaxon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolsForTaxon

`func (o *V2GeneDatasetReportsRequest) SetSymbolsForTaxon(v V2GeneDatasetReportsRequestSymbolsForTaxon)`

SetSymbolsForTaxon sets SymbolsForTaxon field to given value.

### HasSymbolsForTaxon

`func (o *V2GeneDatasetReportsRequest) HasSymbolsForTaxon() bool`

HasSymbolsForTaxon returns a boolean if a field has been set.

### GetTaxon

`func (o *V2GeneDatasetReportsRequest) GetTaxon() string`

GetTaxon returns the Taxon field if non-nil, zero value otherwise.

### GetTaxonOk

`func (o *V2GeneDatasetReportsRequest) GetTaxonOk() (*string, bool)`

GetTaxonOk returns a tuple with the Taxon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxon

`func (o *V2GeneDatasetReportsRequest) SetTaxon(v string)`

SetTaxon sets Taxon field to given value.

### HasTaxon

`func (o *V2GeneDatasetReportsRequest) HasTaxon() bool`

HasTaxon returns a boolean if a field has been set.

### GetTableFields

`func (o *V2GeneDatasetReportsRequest) GetTableFields() []string`

GetTableFields returns the TableFields field if non-nil, zero value otherwise.

### GetTableFieldsOk

`func (o *V2GeneDatasetReportsRequest) GetTableFieldsOk() (*[]string, bool)`

GetTableFieldsOk returns a tuple with the TableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFields

`func (o *V2GeneDatasetReportsRequest) SetTableFields(v []string)`

SetTableFields sets TableFields field to given value.

### HasTableFields

`func (o *V2GeneDatasetReportsRequest) HasTableFields() bool`

HasTableFields returns a boolean if a field has been set.

### GetTableFormat

`func (o *V2GeneDatasetReportsRequest) GetTableFormat() string`

GetTableFormat returns the TableFormat field if non-nil, zero value otherwise.

### GetTableFormatOk

`func (o *V2GeneDatasetReportsRequest) GetTableFormatOk() (*string, bool)`

GetTableFormatOk returns a tuple with the TableFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFormat

`func (o *V2GeneDatasetReportsRequest) SetTableFormat(v string)`

SetTableFormat sets TableFormat field to given value.

### HasTableFormat

`func (o *V2GeneDatasetReportsRequest) HasTableFormat() bool`

HasTableFormat returns a boolean if a field has been set.

### GetIncludeTabularHeader

`func (o *V2GeneDatasetReportsRequest) GetIncludeTabularHeader() V2IncludeTabularHeader`

GetIncludeTabularHeader returns the IncludeTabularHeader field if non-nil, zero value otherwise.

### GetIncludeTabularHeaderOk

`func (o *V2GeneDatasetReportsRequest) GetIncludeTabularHeaderOk() (*V2IncludeTabularHeader, bool)`

GetIncludeTabularHeaderOk returns a tuple with the IncludeTabularHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTabularHeader

`func (o *V2GeneDatasetReportsRequest) SetIncludeTabularHeader(v V2IncludeTabularHeader)`

SetIncludeTabularHeader sets IncludeTabularHeader field to given value.

### HasIncludeTabularHeader

`func (o *V2GeneDatasetReportsRequest) HasIncludeTabularHeader() bool`

HasIncludeTabularHeader returns a boolean if a field has been set.

### GetPageSize

`func (o *V2GeneDatasetReportsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V2GeneDatasetReportsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V2GeneDatasetReportsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *V2GeneDatasetReportsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *V2GeneDatasetReportsRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V2GeneDatasetReportsRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V2GeneDatasetReportsRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V2GeneDatasetReportsRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetQuery

`func (o *V2GeneDatasetReportsRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *V2GeneDatasetReportsRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *V2GeneDatasetReportsRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *V2GeneDatasetReportsRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTypes

`func (o *V2GeneDatasetReportsRequest) GetTypes() []V2GeneType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *V2GeneDatasetReportsRequest) GetTypesOk() (*[]V2GeneType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *V2GeneDatasetReportsRequest) SetTypes(v []V2GeneType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *V2GeneDatasetReportsRequest) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


