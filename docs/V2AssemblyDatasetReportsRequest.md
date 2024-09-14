# V2AssemblyDatasetReportsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taxons** | Pointer to **[]string** |  | [optional] 
**Bioprojects** | Pointer to **[]string** |  | [optional] 
**BiosampleIds** | Pointer to **[]string** |  | [optional] 
**AssemblyNames** | Pointer to **[]string** |  | [optional] 
**WgsAccessions** | Pointer to **[]string** |  | [optional] 
**Accessions** | Pointer to **[]string** |  | [optional] 
**Filters** | Pointer to [**V2AssemblyDatasetDescriptorsFilter**](V2AssemblyDatasetDescriptorsFilter.md) |  | [optional] 
**TaxExactMatch** | Pointer to **bool** |  | [optional] 
**Chromosomes** | Pointer to **[]string** |  | [optional] 
**TableFields** | Pointer to **[]string** |  | [optional] 
**ReturnedContent** | Pointer to [**V2AssemblyDatasetReportsRequestContentType**](V2AssemblyDatasetReportsRequestContentType.md) |  | [optional] [default to V2ASSEMBLYDATASETREPORTSREQUESTCONTENTTYPE_COMPLETE]
**PageSize** | Pointer to **int32** |  | [optional] 
**PageToken** | Pointer to **string** |  | [optional] 
**Sort** | Pointer to [**[]V2SortField**](V2SortField.md) |  | [optional] 
**IncludeTabularHeader** | Pointer to [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) |  | [optional] [default to V2INCLUDETABULARHEADER_FIRST_PAGE_ONLY]
**TableFormat** | Pointer to **string** |  | [optional] 

## Methods

### NewV2AssemblyDatasetReportsRequest

`func NewV2AssemblyDatasetReportsRequest() *V2AssemblyDatasetReportsRequest`

NewV2AssemblyDatasetReportsRequest instantiates a new V2AssemblyDatasetReportsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssemblyDatasetReportsRequestWithDefaults

`func NewV2AssemblyDatasetReportsRequestWithDefaults() *V2AssemblyDatasetReportsRequest`

NewV2AssemblyDatasetReportsRequestWithDefaults instantiates a new V2AssemblyDatasetReportsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxons

`func (o *V2AssemblyDatasetReportsRequest) GetTaxons() []string`

GetTaxons returns the Taxons field if non-nil, zero value otherwise.

### GetTaxonsOk

`func (o *V2AssemblyDatasetReportsRequest) GetTaxonsOk() (*[]string, bool)`

GetTaxonsOk returns a tuple with the Taxons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxons

`func (o *V2AssemblyDatasetReportsRequest) SetTaxons(v []string)`

SetTaxons sets Taxons field to given value.

### HasTaxons

`func (o *V2AssemblyDatasetReportsRequest) HasTaxons() bool`

HasTaxons returns a boolean if a field has been set.

### GetBioprojects

`func (o *V2AssemblyDatasetReportsRequest) GetBioprojects() []string`

GetBioprojects returns the Bioprojects field if non-nil, zero value otherwise.

### GetBioprojectsOk

`func (o *V2AssemblyDatasetReportsRequest) GetBioprojectsOk() (*[]string, bool)`

GetBioprojectsOk returns a tuple with the Bioprojects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioprojects

`func (o *V2AssemblyDatasetReportsRequest) SetBioprojects(v []string)`

SetBioprojects sets Bioprojects field to given value.

### HasBioprojects

`func (o *V2AssemblyDatasetReportsRequest) HasBioprojects() bool`

HasBioprojects returns a boolean if a field has been set.

### GetBiosampleIds

`func (o *V2AssemblyDatasetReportsRequest) GetBiosampleIds() []string`

GetBiosampleIds returns the BiosampleIds field if non-nil, zero value otherwise.

### GetBiosampleIdsOk

`func (o *V2AssemblyDatasetReportsRequest) GetBiosampleIdsOk() (*[]string, bool)`

GetBiosampleIdsOk returns a tuple with the BiosampleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosampleIds

`func (o *V2AssemblyDatasetReportsRequest) SetBiosampleIds(v []string)`

SetBiosampleIds sets BiosampleIds field to given value.

### HasBiosampleIds

`func (o *V2AssemblyDatasetReportsRequest) HasBiosampleIds() bool`

HasBiosampleIds returns a boolean if a field has been set.

### GetAssemblyNames

`func (o *V2AssemblyDatasetReportsRequest) GetAssemblyNames() []string`

GetAssemblyNames returns the AssemblyNames field if non-nil, zero value otherwise.

### GetAssemblyNamesOk

`func (o *V2AssemblyDatasetReportsRequest) GetAssemblyNamesOk() (*[]string, bool)`

GetAssemblyNamesOk returns a tuple with the AssemblyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyNames

`func (o *V2AssemblyDatasetReportsRequest) SetAssemblyNames(v []string)`

SetAssemblyNames sets AssemblyNames field to given value.

### HasAssemblyNames

`func (o *V2AssemblyDatasetReportsRequest) HasAssemblyNames() bool`

HasAssemblyNames returns a boolean if a field has been set.

### GetWgsAccessions

`func (o *V2AssemblyDatasetReportsRequest) GetWgsAccessions() []string`

GetWgsAccessions returns the WgsAccessions field if non-nil, zero value otherwise.

### GetWgsAccessionsOk

`func (o *V2AssemblyDatasetReportsRequest) GetWgsAccessionsOk() (*[]string, bool)`

GetWgsAccessionsOk returns a tuple with the WgsAccessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWgsAccessions

`func (o *V2AssemblyDatasetReportsRequest) SetWgsAccessions(v []string)`

SetWgsAccessions sets WgsAccessions field to given value.

### HasWgsAccessions

`func (o *V2AssemblyDatasetReportsRequest) HasWgsAccessions() bool`

HasWgsAccessions returns a boolean if a field has been set.

### GetAccessions

`func (o *V2AssemblyDatasetReportsRequest) GetAccessions() []string`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V2AssemblyDatasetReportsRequest) GetAccessionsOk() (*[]string, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V2AssemblyDatasetReportsRequest) SetAccessions(v []string)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V2AssemblyDatasetReportsRequest) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetFilters

`func (o *V2AssemblyDatasetReportsRequest) GetFilters() V2AssemblyDatasetDescriptorsFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *V2AssemblyDatasetReportsRequest) GetFiltersOk() (*V2AssemblyDatasetDescriptorsFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *V2AssemblyDatasetReportsRequest) SetFilters(v V2AssemblyDatasetDescriptorsFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *V2AssemblyDatasetReportsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTaxExactMatch

`func (o *V2AssemblyDatasetReportsRequest) GetTaxExactMatch() bool`

GetTaxExactMatch returns the TaxExactMatch field if non-nil, zero value otherwise.

### GetTaxExactMatchOk

`func (o *V2AssemblyDatasetReportsRequest) GetTaxExactMatchOk() (*bool, bool)`

GetTaxExactMatchOk returns a tuple with the TaxExactMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExactMatch

`func (o *V2AssemblyDatasetReportsRequest) SetTaxExactMatch(v bool)`

SetTaxExactMatch sets TaxExactMatch field to given value.

### HasTaxExactMatch

`func (o *V2AssemblyDatasetReportsRequest) HasTaxExactMatch() bool`

HasTaxExactMatch returns a boolean if a field has been set.

### GetChromosomes

`func (o *V2AssemblyDatasetReportsRequest) GetChromosomes() []string`

GetChromosomes returns the Chromosomes field if non-nil, zero value otherwise.

### GetChromosomesOk

`func (o *V2AssemblyDatasetReportsRequest) GetChromosomesOk() (*[]string, bool)`

GetChromosomesOk returns a tuple with the Chromosomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosomes

`func (o *V2AssemblyDatasetReportsRequest) SetChromosomes(v []string)`

SetChromosomes sets Chromosomes field to given value.

### HasChromosomes

`func (o *V2AssemblyDatasetReportsRequest) HasChromosomes() bool`

HasChromosomes returns a boolean if a field has been set.

### GetTableFields

`func (o *V2AssemblyDatasetReportsRequest) GetTableFields() []string`

GetTableFields returns the TableFields field if non-nil, zero value otherwise.

### GetTableFieldsOk

`func (o *V2AssemblyDatasetReportsRequest) GetTableFieldsOk() (*[]string, bool)`

GetTableFieldsOk returns a tuple with the TableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFields

`func (o *V2AssemblyDatasetReportsRequest) SetTableFields(v []string)`

SetTableFields sets TableFields field to given value.

### HasTableFields

`func (o *V2AssemblyDatasetReportsRequest) HasTableFields() bool`

HasTableFields returns a boolean if a field has been set.

### GetReturnedContent

`func (o *V2AssemblyDatasetReportsRequest) GetReturnedContent() V2AssemblyDatasetReportsRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V2AssemblyDatasetReportsRequest) GetReturnedContentOk() (*V2AssemblyDatasetReportsRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V2AssemblyDatasetReportsRequest) SetReturnedContent(v V2AssemblyDatasetReportsRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V2AssemblyDatasetReportsRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.

### GetPageSize

`func (o *V2AssemblyDatasetReportsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V2AssemblyDatasetReportsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V2AssemblyDatasetReportsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *V2AssemblyDatasetReportsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *V2AssemblyDatasetReportsRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V2AssemblyDatasetReportsRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V2AssemblyDatasetReportsRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V2AssemblyDatasetReportsRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetSort

`func (o *V2AssemblyDatasetReportsRequest) GetSort() []V2SortField`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *V2AssemblyDatasetReportsRequest) GetSortOk() (*[]V2SortField, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *V2AssemblyDatasetReportsRequest) SetSort(v []V2SortField)`

SetSort sets Sort field to given value.

### HasSort

`func (o *V2AssemblyDatasetReportsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetIncludeTabularHeader

`func (o *V2AssemblyDatasetReportsRequest) GetIncludeTabularHeader() V2IncludeTabularHeader`

GetIncludeTabularHeader returns the IncludeTabularHeader field if non-nil, zero value otherwise.

### GetIncludeTabularHeaderOk

`func (o *V2AssemblyDatasetReportsRequest) GetIncludeTabularHeaderOk() (*V2IncludeTabularHeader, bool)`

GetIncludeTabularHeaderOk returns a tuple with the IncludeTabularHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTabularHeader

`func (o *V2AssemblyDatasetReportsRequest) SetIncludeTabularHeader(v V2IncludeTabularHeader)`

SetIncludeTabularHeader sets IncludeTabularHeader field to given value.

### HasIncludeTabularHeader

`func (o *V2AssemblyDatasetReportsRequest) HasIncludeTabularHeader() bool`

HasIncludeTabularHeader returns a boolean if a field has been set.

### GetTableFormat

`func (o *V2AssemblyDatasetReportsRequest) GetTableFormat() string`

GetTableFormat returns the TableFormat field if non-nil, zero value otherwise.

### GetTableFormatOk

`func (o *V2AssemblyDatasetReportsRequest) GetTableFormatOk() (*string, bool)`

GetTableFormatOk returns a tuple with the TableFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFormat

`func (o *V2AssemblyDatasetReportsRequest) SetTableFormat(v string)`

SetTableFormat sets TableFormat field to given value.

### HasTableFormat

`func (o *V2AssemblyDatasetReportsRequest) HasTableFormat() bool`

HasTableFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


