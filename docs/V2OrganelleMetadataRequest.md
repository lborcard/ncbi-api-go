# V2OrganelleMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taxons** | Pointer to **[]string** |  | [optional] 
**Accessions** | Pointer to **[]string** |  | [optional] 
**OrganelleTypes** | Pointer to [**[]V2reportsOrganelleType**](V2reportsOrganelleType.md) |  | [optional] 
**FirstReleaseDate** | Pointer to **time.Time** |  | [optional] 
**LastReleaseDate** | Pointer to **time.Time** |  | [optional] 
**TaxExactMatch** | Pointer to **bool** |  | [optional] 
**Sort** | Pointer to [**[]V2OrganelleSort**](V2OrganelleSort.md) |  | [optional] 
**ReturnedContent** | Pointer to [**V2OrganelleMetadataRequestContentType**](V2OrganelleMetadataRequestContentType.md) |  | [optional] [default to V2ORGANELLEMETADATAREQUESTCONTENTTYPE_COMPLETE]
**PageSize** | Pointer to **int32** |  | [optional] 
**PageToken** | Pointer to **string** |  | [optional] 
**TableFormat** | Pointer to [**V2OrganelleMetadataRequestOrganelleTableFormat**](V2OrganelleMetadataRequestOrganelleTableFormat.md) |  | [optional] [default to V2ORGANELLEMETADATAREQUESTORGANELLETABLEFORMAT_ORGANELLE_TABLE_FORMAT_NO_TABLE]
**IncludeTabularHeader** | Pointer to [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) |  | [optional] [default to V2INCLUDETABULARHEADER_FIRST_PAGE_ONLY]

## Methods

### NewV2OrganelleMetadataRequest

`func NewV2OrganelleMetadataRequest() *V2OrganelleMetadataRequest`

NewV2OrganelleMetadataRequest instantiates a new V2OrganelleMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2OrganelleMetadataRequestWithDefaults

`func NewV2OrganelleMetadataRequestWithDefaults() *V2OrganelleMetadataRequest`

NewV2OrganelleMetadataRequestWithDefaults instantiates a new V2OrganelleMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxons

`func (o *V2OrganelleMetadataRequest) GetTaxons() []string`

GetTaxons returns the Taxons field if non-nil, zero value otherwise.

### GetTaxonsOk

`func (o *V2OrganelleMetadataRequest) GetTaxonsOk() (*[]string, bool)`

GetTaxonsOk returns a tuple with the Taxons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxons

`func (o *V2OrganelleMetadataRequest) SetTaxons(v []string)`

SetTaxons sets Taxons field to given value.

### HasTaxons

`func (o *V2OrganelleMetadataRequest) HasTaxons() bool`

HasTaxons returns a boolean if a field has been set.

### GetAccessions

`func (o *V2OrganelleMetadataRequest) GetAccessions() []string`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V2OrganelleMetadataRequest) GetAccessionsOk() (*[]string, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V2OrganelleMetadataRequest) SetAccessions(v []string)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V2OrganelleMetadataRequest) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetOrganelleTypes

`func (o *V2OrganelleMetadataRequest) GetOrganelleTypes() []V2reportsOrganelleType`

GetOrganelleTypes returns the OrganelleTypes field if non-nil, zero value otherwise.

### GetOrganelleTypesOk

`func (o *V2OrganelleMetadataRequest) GetOrganelleTypesOk() (*[]V2reportsOrganelleType, bool)`

GetOrganelleTypesOk returns a tuple with the OrganelleTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganelleTypes

`func (o *V2OrganelleMetadataRequest) SetOrganelleTypes(v []V2reportsOrganelleType)`

SetOrganelleTypes sets OrganelleTypes field to given value.

### HasOrganelleTypes

`func (o *V2OrganelleMetadataRequest) HasOrganelleTypes() bool`

HasOrganelleTypes returns a boolean if a field has been set.

### GetFirstReleaseDate

`func (o *V2OrganelleMetadataRequest) GetFirstReleaseDate() time.Time`

GetFirstReleaseDate returns the FirstReleaseDate field if non-nil, zero value otherwise.

### GetFirstReleaseDateOk

`func (o *V2OrganelleMetadataRequest) GetFirstReleaseDateOk() (*time.Time, bool)`

GetFirstReleaseDateOk returns a tuple with the FirstReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstReleaseDate

`func (o *V2OrganelleMetadataRequest) SetFirstReleaseDate(v time.Time)`

SetFirstReleaseDate sets FirstReleaseDate field to given value.

### HasFirstReleaseDate

`func (o *V2OrganelleMetadataRequest) HasFirstReleaseDate() bool`

HasFirstReleaseDate returns a boolean if a field has been set.

### GetLastReleaseDate

`func (o *V2OrganelleMetadataRequest) GetLastReleaseDate() time.Time`

GetLastReleaseDate returns the LastReleaseDate field if non-nil, zero value otherwise.

### GetLastReleaseDateOk

`func (o *V2OrganelleMetadataRequest) GetLastReleaseDateOk() (*time.Time, bool)`

GetLastReleaseDateOk returns a tuple with the LastReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReleaseDate

`func (o *V2OrganelleMetadataRequest) SetLastReleaseDate(v time.Time)`

SetLastReleaseDate sets LastReleaseDate field to given value.

### HasLastReleaseDate

`func (o *V2OrganelleMetadataRequest) HasLastReleaseDate() bool`

HasLastReleaseDate returns a boolean if a field has been set.

### GetTaxExactMatch

`func (o *V2OrganelleMetadataRequest) GetTaxExactMatch() bool`

GetTaxExactMatch returns the TaxExactMatch field if non-nil, zero value otherwise.

### GetTaxExactMatchOk

`func (o *V2OrganelleMetadataRequest) GetTaxExactMatchOk() (*bool, bool)`

GetTaxExactMatchOk returns a tuple with the TaxExactMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExactMatch

`func (o *V2OrganelleMetadataRequest) SetTaxExactMatch(v bool)`

SetTaxExactMatch sets TaxExactMatch field to given value.

### HasTaxExactMatch

`func (o *V2OrganelleMetadataRequest) HasTaxExactMatch() bool`

HasTaxExactMatch returns a boolean if a field has been set.

### GetSort

`func (o *V2OrganelleMetadataRequest) GetSort() []V2OrganelleSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *V2OrganelleMetadataRequest) GetSortOk() (*[]V2OrganelleSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *V2OrganelleMetadataRequest) SetSort(v []V2OrganelleSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *V2OrganelleMetadataRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetReturnedContent

`func (o *V2OrganelleMetadataRequest) GetReturnedContent() V2OrganelleMetadataRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V2OrganelleMetadataRequest) GetReturnedContentOk() (*V2OrganelleMetadataRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V2OrganelleMetadataRequest) SetReturnedContent(v V2OrganelleMetadataRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V2OrganelleMetadataRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.

### GetPageSize

`func (o *V2OrganelleMetadataRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V2OrganelleMetadataRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V2OrganelleMetadataRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *V2OrganelleMetadataRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *V2OrganelleMetadataRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V2OrganelleMetadataRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V2OrganelleMetadataRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V2OrganelleMetadataRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetTableFormat

`func (o *V2OrganelleMetadataRequest) GetTableFormat() V2OrganelleMetadataRequestOrganelleTableFormat`

GetTableFormat returns the TableFormat field if non-nil, zero value otherwise.

### GetTableFormatOk

`func (o *V2OrganelleMetadataRequest) GetTableFormatOk() (*V2OrganelleMetadataRequestOrganelleTableFormat, bool)`

GetTableFormatOk returns a tuple with the TableFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFormat

`func (o *V2OrganelleMetadataRequest) SetTableFormat(v V2OrganelleMetadataRequestOrganelleTableFormat)`

SetTableFormat sets TableFormat field to given value.

### HasTableFormat

`func (o *V2OrganelleMetadataRequest) HasTableFormat() bool`

HasTableFormat returns a boolean if a field has been set.

### GetIncludeTabularHeader

`func (o *V2OrganelleMetadataRequest) GetIncludeTabularHeader() V2IncludeTabularHeader`

GetIncludeTabularHeader returns the IncludeTabularHeader field if non-nil, zero value otherwise.

### GetIncludeTabularHeaderOk

`func (o *V2OrganelleMetadataRequest) GetIncludeTabularHeaderOk() (*V2IncludeTabularHeader, bool)`

GetIncludeTabularHeaderOk returns a tuple with the IncludeTabularHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTabularHeader

`func (o *V2OrganelleMetadataRequest) SetIncludeTabularHeader(v V2IncludeTabularHeader)`

SetIncludeTabularHeader sets IncludeTabularHeader field to given value.

### HasIncludeTabularHeader

`func (o *V2OrganelleMetadataRequest) HasIncludeTabularHeader() bool`

HasIncludeTabularHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


