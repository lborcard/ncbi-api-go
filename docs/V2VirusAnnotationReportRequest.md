# V2VirusAnnotationReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**V2VirusAnnotationFilter**](V2VirusAnnotationFilter.md) |  | [optional] 
**TableFields** | Pointer to **[]string** |  | [optional] 
**TableFormat** | Pointer to **string** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewV2VirusAnnotationReportRequest

`func NewV2VirusAnnotationReportRequest() *V2VirusAnnotationReportRequest`

NewV2VirusAnnotationReportRequest instantiates a new V2VirusAnnotationReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2VirusAnnotationReportRequestWithDefaults

`func NewV2VirusAnnotationReportRequestWithDefaults() *V2VirusAnnotationReportRequest`

NewV2VirusAnnotationReportRequestWithDefaults instantiates a new V2VirusAnnotationReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *V2VirusAnnotationReportRequest) GetFilter() V2VirusAnnotationFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *V2VirusAnnotationReportRequest) GetFilterOk() (*V2VirusAnnotationFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *V2VirusAnnotationReportRequest) SetFilter(v V2VirusAnnotationFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *V2VirusAnnotationReportRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetTableFields

`func (o *V2VirusAnnotationReportRequest) GetTableFields() []string`

GetTableFields returns the TableFields field if non-nil, zero value otherwise.

### GetTableFieldsOk

`func (o *V2VirusAnnotationReportRequest) GetTableFieldsOk() (*[]string, bool)`

GetTableFieldsOk returns a tuple with the TableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFields

`func (o *V2VirusAnnotationReportRequest) SetTableFields(v []string)`

SetTableFields sets TableFields field to given value.

### HasTableFields

`func (o *V2VirusAnnotationReportRequest) HasTableFields() bool`

HasTableFields returns a boolean if a field has been set.

### GetTableFormat

`func (o *V2VirusAnnotationReportRequest) GetTableFormat() string`

GetTableFormat returns the TableFormat field if non-nil, zero value otherwise.

### GetTableFormatOk

`func (o *V2VirusAnnotationReportRequest) GetTableFormatOk() (*string, bool)`

GetTableFormatOk returns a tuple with the TableFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFormat

`func (o *V2VirusAnnotationReportRequest) SetTableFormat(v string)`

SetTableFormat sets TableFormat field to given value.

### HasTableFormat

`func (o *V2VirusAnnotationReportRequest) HasTableFormat() bool`

HasTableFormat returns a boolean if a field has been set.

### GetPageSize

`func (o *V2VirusAnnotationReportRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V2VirusAnnotationReportRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V2VirusAnnotationReportRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *V2VirusAnnotationReportRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *V2VirusAnnotationReportRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V2VirusAnnotationReportRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V2VirusAnnotationReportRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V2VirusAnnotationReportRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


