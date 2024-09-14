# V2VirusDataReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**V2VirusDatasetFilter**](V2VirusDatasetFilter.md) |  | [optional] 
**ReturnedContent** | Pointer to [**V2VirusDataReportRequestContentType**](V2VirusDataReportRequestContentType.md) |  | [optional] [default to V2VIRUSDATAREPORTREQUESTCONTENTTYPE_COMPLETE]
**TableFields** | Pointer to **[]string** |  | [optional] 
**TableFormat** | Pointer to **string** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewV2VirusDataReportRequest

`func NewV2VirusDataReportRequest() *V2VirusDataReportRequest`

NewV2VirusDataReportRequest instantiates a new V2VirusDataReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2VirusDataReportRequestWithDefaults

`func NewV2VirusDataReportRequestWithDefaults() *V2VirusDataReportRequest`

NewV2VirusDataReportRequestWithDefaults instantiates a new V2VirusDataReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *V2VirusDataReportRequest) GetFilter() V2VirusDatasetFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *V2VirusDataReportRequest) GetFilterOk() (*V2VirusDatasetFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *V2VirusDataReportRequest) SetFilter(v V2VirusDatasetFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *V2VirusDataReportRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetReturnedContent

`func (o *V2VirusDataReportRequest) GetReturnedContent() V2VirusDataReportRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V2VirusDataReportRequest) GetReturnedContentOk() (*V2VirusDataReportRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V2VirusDataReportRequest) SetReturnedContent(v V2VirusDataReportRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V2VirusDataReportRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.

### GetTableFields

`func (o *V2VirusDataReportRequest) GetTableFields() []string`

GetTableFields returns the TableFields field if non-nil, zero value otherwise.

### GetTableFieldsOk

`func (o *V2VirusDataReportRequest) GetTableFieldsOk() (*[]string, bool)`

GetTableFieldsOk returns a tuple with the TableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFields

`func (o *V2VirusDataReportRequest) SetTableFields(v []string)`

SetTableFields sets TableFields field to given value.

### HasTableFields

`func (o *V2VirusDataReportRequest) HasTableFields() bool`

HasTableFields returns a boolean if a field has been set.

### GetTableFormat

`func (o *V2VirusDataReportRequest) GetTableFormat() string`

GetTableFormat returns the TableFormat field if non-nil, zero value otherwise.

### GetTableFormatOk

`func (o *V2VirusDataReportRequest) GetTableFormatOk() (*string, bool)`

GetTableFormatOk returns a tuple with the TableFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFormat

`func (o *V2VirusDataReportRequest) SetTableFormat(v string)`

SetTableFormat sets TableFormat field to given value.

### HasTableFormat

`func (o *V2VirusDataReportRequest) HasTableFormat() bool`

HasTableFormat returns a boolean if a field has been set.

### GetPageSize

`func (o *V2VirusDataReportRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V2VirusDataReportRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V2VirusDataReportRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *V2VirusDataReportRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *V2VirusDataReportRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V2VirusDataReportRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V2VirusDataReportRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V2VirusDataReportRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


