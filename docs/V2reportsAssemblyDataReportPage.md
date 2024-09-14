# V2reportsAssemblyDataReportPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reports** | Pointer to [**[]V2reportsAssemblyDataReport**](V2reportsAssemblyDataReport.md) |  | [optional] 
**ContentType** | Pointer to [**V2reportsContentType**](V2reportsContentType.md) |  | [optional] [default to V2REPORTSCONTENTTYPE_COMPLETE]
**TotalCount** | Pointer to **int32** |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to [**[]V2reportsMessage**](V2reportsMessage.md) |  | [optional] 
**ReportType** | Pointer to **string** |  | [optional] 
**ReportFields** | Pointer to **[]string** |  | [optional] 
**FirstPage** | Pointer to **bool** |  | [optional] 
**ReportFormat** | Pointer to **string** |  | [optional] 

## Methods

### NewV2reportsAssemblyDataReportPage

`func NewV2reportsAssemblyDataReportPage() *V2reportsAssemblyDataReportPage`

NewV2reportsAssemblyDataReportPage instantiates a new V2reportsAssemblyDataReportPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsAssemblyDataReportPageWithDefaults

`func NewV2reportsAssemblyDataReportPageWithDefaults() *V2reportsAssemblyDataReportPage`

NewV2reportsAssemblyDataReportPageWithDefaults instantiates a new V2reportsAssemblyDataReportPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReports

`func (o *V2reportsAssemblyDataReportPage) GetReports() []V2reportsAssemblyDataReport`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *V2reportsAssemblyDataReportPage) GetReportsOk() (*[]V2reportsAssemblyDataReport, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *V2reportsAssemblyDataReportPage) SetReports(v []V2reportsAssemblyDataReport)`

SetReports sets Reports field to given value.

### HasReports

`func (o *V2reportsAssemblyDataReportPage) HasReports() bool`

HasReports returns a boolean if a field has been set.

### GetContentType

`func (o *V2reportsAssemblyDataReportPage) GetContentType() V2reportsContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *V2reportsAssemblyDataReportPage) GetContentTypeOk() (*V2reportsContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *V2reportsAssemblyDataReportPage) SetContentType(v V2reportsContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *V2reportsAssemblyDataReportPage) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetTotalCount

`func (o *V2reportsAssemblyDataReportPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *V2reportsAssemblyDataReportPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *V2reportsAssemblyDataReportPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *V2reportsAssemblyDataReportPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetNextPageToken

`func (o *V2reportsAssemblyDataReportPage) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *V2reportsAssemblyDataReportPage) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *V2reportsAssemblyDataReportPage) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *V2reportsAssemblyDataReportPage) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetMessages

`func (o *V2reportsAssemblyDataReportPage) GetMessages() []V2reportsMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *V2reportsAssemblyDataReportPage) GetMessagesOk() (*[]V2reportsMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *V2reportsAssemblyDataReportPage) SetMessages(v []V2reportsMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *V2reportsAssemblyDataReportPage) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetReportType

`func (o *V2reportsAssemblyDataReportPage) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *V2reportsAssemblyDataReportPage) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *V2reportsAssemblyDataReportPage) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *V2reportsAssemblyDataReportPage) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetReportFields

`func (o *V2reportsAssemblyDataReportPage) GetReportFields() []string`

GetReportFields returns the ReportFields field if non-nil, zero value otherwise.

### GetReportFieldsOk

`func (o *V2reportsAssemblyDataReportPage) GetReportFieldsOk() (*[]string, bool)`

GetReportFieldsOk returns a tuple with the ReportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFields

`func (o *V2reportsAssemblyDataReportPage) SetReportFields(v []string)`

SetReportFields sets ReportFields field to given value.

### HasReportFields

`func (o *V2reportsAssemblyDataReportPage) HasReportFields() bool`

HasReportFields returns a boolean if a field has been set.

### GetFirstPage

`func (o *V2reportsAssemblyDataReportPage) GetFirstPage() bool`

GetFirstPage returns the FirstPage field if non-nil, zero value otherwise.

### GetFirstPageOk

`func (o *V2reportsAssemblyDataReportPage) GetFirstPageOk() (*bool, bool)`

GetFirstPageOk returns a tuple with the FirstPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPage

`func (o *V2reportsAssemblyDataReportPage) SetFirstPage(v bool)`

SetFirstPage sets FirstPage field to given value.

### HasFirstPage

`func (o *V2reportsAssemblyDataReportPage) HasFirstPage() bool`

HasFirstPage returns a boolean if a field has been set.

### GetReportFormat

`func (o *V2reportsAssemblyDataReportPage) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *V2reportsAssemblyDataReportPage) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *V2reportsAssemblyDataReportPage) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.

### HasReportFormat

`func (o *V2reportsAssemblyDataReportPage) HasReportFormat() bool`

HasReportFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


