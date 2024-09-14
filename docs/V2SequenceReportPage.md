# V2SequenceReportPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reports** | Pointer to [**[]V2reportsSequenceInfo**](V2reportsSequenceInfo.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 
**ReportType** | Pointer to **string** |  | [optional] 
**ReportFields** | Pointer to **[]string** |  | [optional] 
**FirstPage** | Pointer to **bool** |  | [optional] 
**ReportFormat** | Pointer to **string** |  | [optional] 

## Methods

### NewV2SequenceReportPage

`func NewV2SequenceReportPage() *V2SequenceReportPage`

NewV2SequenceReportPage instantiates a new V2SequenceReportPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2SequenceReportPageWithDefaults

`func NewV2SequenceReportPageWithDefaults() *V2SequenceReportPage`

NewV2SequenceReportPageWithDefaults instantiates a new V2SequenceReportPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReports

`func (o *V2SequenceReportPage) GetReports() []V2reportsSequenceInfo`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *V2SequenceReportPage) GetReportsOk() (*[]V2reportsSequenceInfo, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *V2SequenceReportPage) SetReports(v []V2reportsSequenceInfo)`

SetReports sets Reports field to given value.

### HasReports

`func (o *V2SequenceReportPage) HasReports() bool`

HasReports returns a boolean if a field has been set.

### GetTotalCount

`func (o *V2SequenceReportPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *V2SequenceReportPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *V2SequenceReportPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *V2SequenceReportPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetNextPageToken

`func (o *V2SequenceReportPage) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *V2SequenceReportPage) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *V2SequenceReportPage) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *V2SequenceReportPage) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetReportType

`func (o *V2SequenceReportPage) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *V2SequenceReportPage) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *V2SequenceReportPage) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *V2SequenceReportPage) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetReportFields

`func (o *V2SequenceReportPage) GetReportFields() []string`

GetReportFields returns the ReportFields field if non-nil, zero value otherwise.

### GetReportFieldsOk

`func (o *V2SequenceReportPage) GetReportFieldsOk() (*[]string, bool)`

GetReportFieldsOk returns a tuple with the ReportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFields

`func (o *V2SequenceReportPage) SetReportFields(v []string)`

SetReportFields sets ReportFields field to given value.

### HasReportFields

`func (o *V2SequenceReportPage) HasReportFields() bool`

HasReportFields returns a boolean if a field has been set.

### GetFirstPage

`func (o *V2SequenceReportPage) GetFirstPage() bool`

GetFirstPage returns the FirstPage field if non-nil, zero value otherwise.

### GetFirstPageOk

`func (o *V2SequenceReportPage) GetFirstPageOk() (*bool, bool)`

GetFirstPageOk returns a tuple with the FirstPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPage

`func (o *V2SequenceReportPage) SetFirstPage(v bool)`

SetFirstPage sets FirstPage field to given value.

### HasFirstPage

`func (o *V2SequenceReportPage) HasFirstPage() bool`

HasFirstPage returns a boolean if a field has been set.

### GetReportFormat

`func (o *V2SequenceReportPage) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *V2SequenceReportPage) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *V2SequenceReportPage) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.

### HasReportFormat

`func (o *V2SequenceReportPage) HasReportFormat() bool`

HasReportFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


