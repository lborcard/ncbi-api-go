# V2reportsTaxonomyNamesDataReportPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reports** | Pointer to [**[]V2reportsTaxonomyNamesReportMatch**](V2reportsTaxonomyNamesReportMatch.md) |  | [optional] 
**Messages** | Pointer to [**[]V2reportsMessage**](V2reportsMessage.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 
**ReportType** | Pointer to **string** |  | [optional] 
**ReportFields** | Pointer to **[]string** |  | [optional] 
**FirstPage** | Pointer to **bool** |  | [optional] 
**ReportFormat** | Pointer to **string** |  | [optional] 

## Methods

### NewV2reportsTaxonomyNamesDataReportPage

`func NewV2reportsTaxonomyNamesDataReportPage() *V2reportsTaxonomyNamesDataReportPage`

NewV2reportsTaxonomyNamesDataReportPage instantiates a new V2reportsTaxonomyNamesDataReportPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsTaxonomyNamesDataReportPageWithDefaults

`func NewV2reportsTaxonomyNamesDataReportPageWithDefaults() *V2reportsTaxonomyNamesDataReportPage`

NewV2reportsTaxonomyNamesDataReportPageWithDefaults instantiates a new V2reportsTaxonomyNamesDataReportPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReports

`func (o *V2reportsTaxonomyNamesDataReportPage) GetReports() []V2reportsTaxonomyNamesReportMatch`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *V2reportsTaxonomyNamesDataReportPage) GetReportsOk() (*[]V2reportsTaxonomyNamesReportMatch, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *V2reportsTaxonomyNamesDataReportPage) SetReports(v []V2reportsTaxonomyNamesReportMatch)`

SetReports sets Reports field to given value.

### HasReports

`func (o *V2reportsTaxonomyNamesDataReportPage) HasReports() bool`

HasReports returns a boolean if a field has been set.

### GetMessages

`func (o *V2reportsTaxonomyNamesDataReportPage) GetMessages() []V2reportsMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *V2reportsTaxonomyNamesDataReportPage) GetMessagesOk() (*[]V2reportsMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *V2reportsTaxonomyNamesDataReportPage) SetMessages(v []V2reportsMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *V2reportsTaxonomyNamesDataReportPage) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetTotalCount

`func (o *V2reportsTaxonomyNamesDataReportPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *V2reportsTaxonomyNamesDataReportPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *V2reportsTaxonomyNamesDataReportPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *V2reportsTaxonomyNamesDataReportPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetNextPageToken

`func (o *V2reportsTaxonomyNamesDataReportPage) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *V2reportsTaxonomyNamesDataReportPage) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *V2reportsTaxonomyNamesDataReportPage) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *V2reportsTaxonomyNamesDataReportPage) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetReportType

`func (o *V2reportsTaxonomyNamesDataReportPage) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *V2reportsTaxonomyNamesDataReportPage) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *V2reportsTaxonomyNamesDataReportPage) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *V2reportsTaxonomyNamesDataReportPage) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetReportFields

`func (o *V2reportsTaxonomyNamesDataReportPage) GetReportFields() []string`

GetReportFields returns the ReportFields field if non-nil, zero value otherwise.

### GetReportFieldsOk

`func (o *V2reportsTaxonomyNamesDataReportPage) GetReportFieldsOk() (*[]string, bool)`

GetReportFieldsOk returns a tuple with the ReportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFields

`func (o *V2reportsTaxonomyNamesDataReportPage) SetReportFields(v []string)`

SetReportFields sets ReportFields field to given value.

### HasReportFields

`func (o *V2reportsTaxonomyNamesDataReportPage) HasReportFields() bool`

HasReportFields returns a boolean if a field has been set.

### GetFirstPage

`func (o *V2reportsTaxonomyNamesDataReportPage) GetFirstPage() bool`

GetFirstPage returns the FirstPage field if non-nil, zero value otherwise.

### GetFirstPageOk

`func (o *V2reportsTaxonomyNamesDataReportPage) GetFirstPageOk() (*bool, bool)`

GetFirstPageOk returns a tuple with the FirstPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPage

`func (o *V2reportsTaxonomyNamesDataReportPage) SetFirstPage(v bool)`

SetFirstPage sets FirstPage field to given value.

### HasFirstPage

`func (o *V2reportsTaxonomyNamesDataReportPage) HasFirstPage() bool`

HasFirstPage returns a boolean if a field has been set.

### GetReportFormat

`func (o *V2reportsTaxonomyNamesDataReportPage) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *V2reportsTaxonomyNamesDataReportPage) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *V2reportsTaxonomyNamesDataReportPage) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.

### HasReportFormat

`func (o *V2reportsTaxonomyNamesDataReportPage) HasReportFormat() bool`

HasReportFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


