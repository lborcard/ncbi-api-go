# V2AssemblySequenceReportsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accession** | Pointer to **string** |  | [optional] 
**Chromosomes** | Pointer to **[]string** |  | [optional] 
**RoleFilters** | Pointer to **[]string** |  | [optional] 
**TableFields** | Pointer to **[]string** |  | [optional] 
**CountAssemblyUnplaced** | Pointer to **bool** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PageToken** | Pointer to **string** |  | [optional] 
**IncludeTabularHeader** | Pointer to [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) |  | [optional] [default to V2INCLUDETABULARHEADER_FIRST_PAGE_ONLY]
**TableFormat** | Pointer to **string** |  | [optional] 

## Methods

### NewV2AssemblySequenceReportsRequest

`func NewV2AssemblySequenceReportsRequest() *V2AssemblySequenceReportsRequest`

NewV2AssemblySequenceReportsRequest instantiates a new V2AssemblySequenceReportsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssemblySequenceReportsRequestWithDefaults

`func NewV2AssemblySequenceReportsRequestWithDefaults() *V2AssemblySequenceReportsRequest`

NewV2AssemblySequenceReportsRequestWithDefaults instantiates a new V2AssemblySequenceReportsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccession

`func (o *V2AssemblySequenceReportsRequest) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V2AssemblySequenceReportsRequest) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V2AssemblySequenceReportsRequest) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V2AssemblySequenceReportsRequest) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetChromosomes

`func (o *V2AssemblySequenceReportsRequest) GetChromosomes() []string`

GetChromosomes returns the Chromosomes field if non-nil, zero value otherwise.

### GetChromosomesOk

`func (o *V2AssemblySequenceReportsRequest) GetChromosomesOk() (*[]string, bool)`

GetChromosomesOk returns a tuple with the Chromosomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosomes

`func (o *V2AssemblySequenceReportsRequest) SetChromosomes(v []string)`

SetChromosomes sets Chromosomes field to given value.

### HasChromosomes

`func (o *V2AssemblySequenceReportsRequest) HasChromosomes() bool`

HasChromosomes returns a boolean if a field has been set.

### GetRoleFilters

`func (o *V2AssemblySequenceReportsRequest) GetRoleFilters() []string`

GetRoleFilters returns the RoleFilters field if non-nil, zero value otherwise.

### GetRoleFiltersOk

`func (o *V2AssemblySequenceReportsRequest) GetRoleFiltersOk() (*[]string, bool)`

GetRoleFiltersOk returns a tuple with the RoleFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleFilters

`func (o *V2AssemblySequenceReportsRequest) SetRoleFilters(v []string)`

SetRoleFilters sets RoleFilters field to given value.

### HasRoleFilters

`func (o *V2AssemblySequenceReportsRequest) HasRoleFilters() bool`

HasRoleFilters returns a boolean if a field has been set.

### GetTableFields

`func (o *V2AssemblySequenceReportsRequest) GetTableFields() []string`

GetTableFields returns the TableFields field if non-nil, zero value otherwise.

### GetTableFieldsOk

`func (o *V2AssemblySequenceReportsRequest) GetTableFieldsOk() (*[]string, bool)`

GetTableFieldsOk returns a tuple with the TableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFields

`func (o *V2AssemblySequenceReportsRequest) SetTableFields(v []string)`

SetTableFields sets TableFields field to given value.

### HasTableFields

`func (o *V2AssemblySequenceReportsRequest) HasTableFields() bool`

HasTableFields returns a boolean if a field has been set.

### GetCountAssemblyUnplaced

`func (o *V2AssemblySequenceReportsRequest) GetCountAssemblyUnplaced() bool`

GetCountAssemblyUnplaced returns the CountAssemblyUnplaced field if non-nil, zero value otherwise.

### GetCountAssemblyUnplacedOk

`func (o *V2AssemblySequenceReportsRequest) GetCountAssemblyUnplacedOk() (*bool, bool)`

GetCountAssemblyUnplacedOk returns a tuple with the CountAssemblyUnplaced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAssemblyUnplaced

`func (o *V2AssemblySequenceReportsRequest) SetCountAssemblyUnplaced(v bool)`

SetCountAssemblyUnplaced sets CountAssemblyUnplaced field to given value.

### HasCountAssemblyUnplaced

`func (o *V2AssemblySequenceReportsRequest) HasCountAssemblyUnplaced() bool`

HasCountAssemblyUnplaced returns a boolean if a field has been set.

### GetPageSize

`func (o *V2AssemblySequenceReportsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V2AssemblySequenceReportsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V2AssemblySequenceReportsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *V2AssemblySequenceReportsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *V2AssemblySequenceReportsRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V2AssemblySequenceReportsRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V2AssemblySequenceReportsRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V2AssemblySequenceReportsRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetIncludeTabularHeader

`func (o *V2AssemblySequenceReportsRequest) GetIncludeTabularHeader() V2IncludeTabularHeader`

GetIncludeTabularHeader returns the IncludeTabularHeader field if non-nil, zero value otherwise.

### GetIncludeTabularHeaderOk

`func (o *V2AssemblySequenceReportsRequest) GetIncludeTabularHeaderOk() (*V2IncludeTabularHeader, bool)`

GetIncludeTabularHeaderOk returns a tuple with the IncludeTabularHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTabularHeader

`func (o *V2AssemblySequenceReportsRequest) SetIncludeTabularHeader(v V2IncludeTabularHeader)`

SetIncludeTabularHeader sets IncludeTabularHeader field to given value.

### HasIncludeTabularHeader

`func (o *V2AssemblySequenceReportsRequest) HasIncludeTabularHeader() bool`

HasIncludeTabularHeader returns a boolean if a field has been set.

### GetTableFormat

`func (o *V2AssemblySequenceReportsRequest) GetTableFormat() string`

GetTableFormat returns the TableFormat field if non-nil, zero value otherwise.

### GetTableFormatOk

`func (o *V2AssemblySequenceReportsRequest) GetTableFormatOk() (*string, bool)`

GetTableFormatOk returns a tuple with the TableFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFormat

`func (o *V2AssemblySequenceReportsRequest) SetTableFormat(v string)`

SetTableFormat sets TableFormat field to given value.

### HasTableFormat

`func (o *V2AssemblySequenceReportsRequest) HasTableFormat() bool`

HasTableFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


