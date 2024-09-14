# V2GenomeAnnotationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accession** | Pointer to **string** |  | [optional] 
**AnnotationIds** | Pointer to **[]string** |  | [optional] 
**Symbols** | Pointer to **[]string** |  | [optional] 
**Locations** | Pointer to **[]string** |  | [optional] 
**GeneTypes** | Pointer to **[]string** |  | [optional] 
**SearchText** | Pointer to **[]string** |  | [optional] 
**Sort** | Pointer to [**[]V2SortField**](V2SortField.md) |  | [optional] 
**IncludeAnnotationType** | Pointer to [**[]V2GenomeAnnotationRequestAnnotationType**](V2GenomeAnnotationRequestAnnotationType.md) |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TableFields** | Pointer to **[]string** |  | [optional] 
**TableFormat** | Pointer to [**V2GenomeAnnotationRequestGenomeAnnotationTableFormat**](V2GenomeAnnotationRequestGenomeAnnotationTableFormat.md) |  | [optional] [default to V2GENOMEANNOTATIONREQUESTGENOMEANNOTATIONTABLEFORMAT_NO_TABLE]
**IncludeTabularHeader** | Pointer to [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) |  | [optional] [default to V2INCLUDETABULARHEADER_FIRST_PAGE_ONLY]
**PageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewV2GenomeAnnotationRequest

`func NewV2GenomeAnnotationRequest() *V2GenomeAnnotationRequest`

NewV2GenomeAnnotationRequest instantiates a new V2GenomeAnnotationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2GenomeAnnotationRequestWithDefaults

`func NewV2GenomeAnnotationRequestWithDefaults() *V2GenomeAnnotationRequest`

NewV2GenomeAnnotationRequestWithDefaults instantiates a new V2GenomeAnnotationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccession

`func (o *V2GenomeAnnotationRequest) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V2GenomeAnnotationRequest) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V2GenomeAnnotationRequest) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V2GenomeAnnotationRequest) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetAnnotationIds

`func (o *V2GenomeAnnotationRequest) GetAnnotationIds() []string`

GetAnnotationIds returns the AnnotationIds field if non-nil, zero value otherwise.

### GetAnnotationIdsOk

`func (o *V2GenomeAnnotationRequest) GetAnnotationIdsOk() (*[]string, bool)`

GetAnnotationIdsOk returns a tuple with the AnnotationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationIds

`func (o *V2GenomeAnnotationRequest) SetAnnotationIds(v []string)`

SetAnnotationIds sets AnnotationIds field to given value.

### HasAnnotationIds

`func (o *V2GenomeAnnotationRequest) HasAnnotationIds() bool`

HasAnnotationIds returns a boolean if a field has been set.

### GetSymbols

`func (o *V2GenomeAnnotationRequest) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *V2GenomeAnnotationRequest) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *V2GenomeAnnotationRequest) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *V2GenomeAnnotationRequest) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetLocations

`func (o *V2GenomeAnnotationRequest) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *V2GenomeAnnotationRequest) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *V2GenomeAnnotationRequest) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *V2GenomeAnnotationRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetGeneTypes

`func (o *V2GenomeAnnotationRequest) GetGeneTypes() []string`

GetGeneTypes returns the GeneTypes field if non-nil, zero value otherwise.

### GetGeneTypesOk

`func (o *V2GenomeAnnotationRequest) GetGeneTypesOk() (*[]string, bool)`

GetGeneTypesOk returns a tuple with the GeneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneTypes

`func (o *V2GenomeAnnotationRequest) SetGeneTypes(v []string)`

SetGeneTypes sets GeneTypes field to given value.

### HasGeneTypes

`func (o *V2GenomeAnnotationRequest) HasGeneTypes() bool`

HasGeneTypes returns a boolean if a field has been set.

### GetSearchText

`func (o *V2GenomeAnnotationRequest) GetSearchText() []string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *V2GenomeAnnotationRequest) GetSearchTextOk() (*[]string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *V2GenomeAnnotationRequest) SetSearchText(v []string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *V2GenomeAnnotationRequest) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetSort

`func (o *V2GenomeAnnotationRequest) GetSort() []V2SortField`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *V2GenomeAnnotationRequest) GetSortOk() (*[]V2SortField, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *V2GenomeAnnotationRequest) SetSort(v []V2SortField)`

SetSort sets Sort field to given value.

### HasSort

`func (o *V2GenomeAnnotationRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetIncludeAnnotationType

`func (o *V2GenomeAnnotationRequest) GetIncludeAnnotationType() []V2GenomeAnnotationRequestAnnotationType`

GetIncludeAnnotationType returns the IncludeAnnotationType field if non-nil, zero value otherwise.

### GetIncludeAnnotationTypeOk

`func (o *V2GenomeAnnotationRequest) GetIncludeAnnotationTypeOk() (*[]V2GenomeAnnotationRequestAnnotationType, bool)`

GetIncludeAnnotationTypeOk returns a tuple with the IncludeAnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnnotationType

`func (o *V2GenomeAnnotationRequest) SetIncludeAnnotationType(v []V2GenomeAnnotationRequestAnnotationType)`

SetIncludeAnnotationType sets IncludeAnnotationType field to given value.

### HasIncludeAnnotationType

`func (o *V2GenomeAnnotationRequest) HasIncludeAnnotationType() bool`

HasIncludeAnnotationType returns a boolean if a field has been set.

### GetPageSize

`func (o *V2GenomeAnnotationRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V2GenomeAnnotationRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V2GenomeAnnotationRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *V2GenomeAnnotationRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTableFields

`func (o *V2GenomeAnnotationRequest) GetTableFields() []string`

GetTableFields returns the TableFields field if non-nil, zero value otherwise.

### GetTableFieldsOk

`func (o *V2GenomeAnnotationRequest) GetTableFieldsOk() (*[]string, bool)`

GetTableFieldsOk returns a tuple with the TableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFields

`func (o *V2GenomeAnnotationRequest) SetTableFields(v []string)`

SetTableFields sets TableFields field to given value.

### HasTableFields

`func (o *V2GenomeAnnotationRequest) HasTableFields() bool`

HasTableFields returns a boolean if a field has been set.

### GetTableFormat

`func (o *V2GenomeAnnotationRequest) GetTableFormat() V2GenomeAnnotationRequestGenomeAnnotationTableFormat`

GetTableFormat returns the TableFormat field if non-nil, zero value otherwise.

### GetTableFormatOk

`func (o *V2GenomeAnnotationRequest) GetTableFormatOk() (*V2GenomeAnnotationRequestGenomeAnnotationTableFormat, bool)`

GetTableFormatOk returns a tuple with the TableFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFormat

`func (o *V2GenomeAnnotationRequest) SetTableFormat(v V2GenomeAnnotationRequestGenomeAnnotationTableFormat)`

SetTableFormat sets TableFormat field to given value.

### HasTableFormat

`func (o *V2GenomeAnnotationRequest) HasTableFormat() bool`

HasTableFormat returns a boolean if a field has been set.

### GetIncludeTabularHeader

`func (o *V2GenomeAnnotationRequest) GetIncludeTabularHeader() V2IncludeTabularHeader`

GetIncludeTabularHeader returns the IncludeTabularHeader field if non-nil, zero value otherwise.

### GetIncludeTabularHeaderOk

`func (o *V2GenomeAnnotationRequest) GetIncludeTabularHeaderOk() (*V2IncludeTabularHeader, bool)`

GetIncludeTabularHeaderOk returns a tuple with the IncludeTabularHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTabularHeader

`func (o *V2GenomeAnnotationRequest) SetIncludeTabularHeader(v V2IncludeTabularHeader)`

SetIncludeTabularHeader sets IncludeTabularHeader field to given value.

### HasIncludeTabularHeader

`func (o *V2GenomeAnnotationRequest) HasIncludeTabularHeader() bool`

HasIncludeTabularHeader returns a boolean if a field has been set.

### GetPageToken

`func (o *V2GenomeAnnotationRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V2GenomeAnnotationRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V2GenomeAnnotationRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V2GenomeAnnotationRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


