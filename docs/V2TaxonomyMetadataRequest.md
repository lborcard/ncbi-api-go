# V2TaxonomyMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taxons** | Pointer to **[]string** |  | [optional] 
**ReturnedContent** | Pointer to [**V2TaxonomyMetadataRequestContentType**](V2TaxonomyMetadataRequestContentType.md) |  | [optional] [default to V2TAXONOMYMETADATAREQUESTCONTENTTYPE_COMPLETE]
**PageSize** | Pointer to **int32** |  | [optional] 
**IncludeTabularHeader** | Pointer to [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) |  | [optional] [default to V2INCLUDETABULARHEADER_FIRST_PAGE_ONLY]
**PageToken** | Pointer to **string** |  | [optional] 
**TableFormat** | Pointer to [**V2TaxonomyMetadataRequestTableFormat**](V2TaxonomyMetadataRequestTableFormat.md) |  | [optional] [default to V2TAXONOMYMETADATAREQUESTTABLEFORMAT_SUMMARY]
**Children** | Pointer to **bool** |  | [optional] 
**Ranks** | Pointer to [**[]V2reportsRankType**](V2reportsRankType.md) |  | [optional] 

## Methods

### NewV2TaxonomyMetadataRequest

`func NewV2TaxonomyMetadataRequest() *V2TaxonomyMetadataRequest`

NewV2TaxonomyMetadataRequest instantiates a new V2TaxonomyMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TaxonomyMetadataRequestWithDefaults

`func NewV2TaxonomyMetadataRequestWithDefaults() *V2TaxonomyMetadataRequest`

NewV2TaxonomyMetadataRequestWithDefaults instantiates a new V2TaxonomyMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxons

`func (o *V2TaxonomyMetadataRequest) GetTaxons() []string`

GetTaxons returns the Taxons field if non-nil, zero value otherwise.

### GetTaxonsOk

`func (o *V2TaxonomyMetadataRequest) GetTaxonsOk() (*[]string, bool)`

GetTaxonsOk returns a tuple with the Taxons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxons

`func (o *V2TaxonomyMetadataRequest) SetTaxons(v []string)`

SetTaxons sets Taxons field to given value.

### HasTaxons

`func (o *V2TaxonomyMetadataRequest) HasTaxons() bool`

HasTaxons returns a boolean if a field has been set.

### GetReturnedContent

`func (o *V2TaxonomyMetadataRequest) GetReturnedContent() V2TaxonomyMetadataRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V2TaxonomyMetadataRequest) GetReturnedContentOk() (*V2TaxonomyMetadataRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V2TaxonomyMetadataRequest) SetReturnedContent(v V2TaxonomyMetadataRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V2TaxonomyMetadataRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.

### GetPageSize

`func (o *V2TaxonomyMetadataRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V2TaxonomyMetadataRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V2TaxonomyMetadataRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *V2TaxonomyMetadataRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetIncludeTabularHeader

`func (o *V2TaxonomyMetadataRequest) GetIncludeTabularHeader() V2IncludeTabularHeader`

GetIncludeTabularHeader returns the IncludeTabularHeader field if non-nil, zero value otherwise.

### GetIncludeTabularHeaderOk

`func (o *V2TaxonomyMetadataRequest) GetIncludeTabularHeaderOk() (*V2IncludeTabularHeader, bool)`

GetIncludeTabularHeaderOk returns a tuple with the IncludeTabularHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTabularHeader

`func (o *V2TaxonomyMetadataRequest) SetIncludeTabularHeader(v V2IncludeTabularHeader)`

SetIncludeTabularHeader sets IncludeTabularHeader field to given value.

### HasIncludeTabularHeader

`func (o *V2TaxonomyMetadataRequest) HasIncludeTabularHeader() bool`

HasIncludeTabularHeader returns a boolean if a field has been set.

### GetPageToken

`func (o *V2TaxonomyMetadataRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V2TaxonomyMetadataRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V2TaxonomyMetadataRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V2TaxonomyMetadataRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetTableFormat

`func (o *V2TaxonomyMetadataRequest) GetTableFormat() V2TaxonomyMetadataRequestTableFormat`

GetTableFormat returns the TableFormat field if non-nil, zero value otherwise.

### GetTableFormatOk

`func (o *V2TaxonomyMetadataRequest) GetTableFormatOk() (*V2TaxonomyMetadataRequestTableFormat, bool)`

GetTableFormatOk returns a tuple with the TableFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFormat

`func (o *V2TaxonomyMetadataRequest) SetTableFormat(v V2TaxonomyMetadataRequestTableFormat)`

SetTableFormat sets TableFormat field to given value.

### HasTableFormat

`func (o *V2TaxonomyMetadataRequest) HasTableFormat() bool`

HasTableFormat returns a boolean if a field has been set.

### GetChildren

`func (o *V2TaxonomyMetadataRequest) GetChildren() bool`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *V2TaxonomyMetadataRequest) GetChildrenOk() (*bool, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *V2TaxonomyMetadataRequest) SetChildren(v bool)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *V2TaxonomyMetadataRequest) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetRanks

`func (o *V2TaxonomyMetadataRequest) GetRanks() []V2reportsRankType`

GetRanks returns the Ranks field if non-nil, zero value otherwise.

### GetRanksOk

`func (o *V2TaxonomyMetadataRequest) GetRanksOk() (*[]V2reportsRankType, bool)`

GetRanksOk returns a tuple with the Ranks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanks

`func (o *V2TaxonomyMetadataRequest) SetRanks(v []V2reportsRankType)`

SetRanks sets Ranks field to given value.

### HasRanks

`func (o *V2TaxonomyMetadataRequest) HasRanks() bool`

HasRanks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


