# V2reportsAnnotationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**ReportUrl** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**V2reportsFeatureCounts**](V2reportsFeatureCounts.md) |  | [optional] 
**Busco** | Pointer to [**V2reportsBuscoStat**](V2reportsBuscoStat.md) |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Pipeline** | Pointer to **string** |  | [optional] 
**SoftwareVersion** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ReleaseVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewV2reportsAnnotationInfo

`func NewV2reportsAnnotationInfo() *V2reportsAnnotationInfo`

NewV2reportsAnnotationInfo instantiates a new V2reportsAnnotationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsAnnotationInfoWithDefaults

`func NewV2reportsAnnotationInfoWithDefaults() *V2reportsAnnotationInfo`

NewV2reportsAnnotationInfoWithDefaults instantiates a new V2reportsAnnotationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2reportsAnnotationInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2reportsAnnotationInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2reportsAnnotationInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2reportsAnnotationInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *V2reportsAnnotationInfo) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *V2reportsAnnotationInfo) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *V2reportsAnnotationInfo) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *V2reportsAnnotationInfo) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReleaseDate

`func (o *V2reportsAnnotationInfo) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *V2reportsAnnotationInfo) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *V2reportsAnnotationInfo) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *V2reportsAnnotationInfo) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetReportUrl

`func (o *V2reportsAnnotationInfo) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *V2reportsAnnotationInfo) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *V2reportsAnnotationInfo) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *V2reportsAnnotationInfo) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.

### GetStats

`func (o *V2reportsAnnotationInfo) GetStats() V2reportsFeatureCounts`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *V2reportsAnnotationInfo) GetStatsOk() (*V2reportsFeatureCounts, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *V2reportsAnnotationInfo) SetStats(v V2reportsFeatureCounts)`

SetStats sets Stats field to given value.

### HasStats

`func (o *V2reportsAnnotationInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetBusco

`func (o *V2reportsAnnotationInfo) GetBusco() V2reportsBuscoStat`

GetBusco returns the Busco field if non-nil, zero value otherwise.

### GetBuscoOk

`func (o *V2reportsAnnotationInfo) GetBuscoOk() (*V2reportsBuscoStat, bool)`

GetBuscoOk returns a tuple with the Busco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusco

`func (o *V2reportsAnnotationInfo) SetBusco(v V2reportsBuscoStat)`

SetBusco sets Busco field to given value.

### HasBusco

`func (o *V2reportsAnnotationInfo) HasBusco() bool`

HasBusco returns a boolean if a field has been set.

### GetMethod

`func (o *V2reportsAnnotationInfo) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2reportsAnnotationInfo) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2reportsAnnotationInfo) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2reportsAnnotationInfo) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPipeline

`func (o *V2reportsAnnotationInfo) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *V2reportsAnnotationInfo) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *V2reportsAnnotationInfo) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *V2reportsAnnotationInfo) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *V2reportsAnnotationInfo) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *V2reportsAnnotationInfo) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *V2reportsAnnotationInfo) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *V2reportsAnnotationInfo) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetStatus

`func (o *V2reportsAnnotationInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2reportsAnnotationInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2reportsAnnotationInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V2reportsAnnotationInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReleaseVersion

`func (o *V2reportsAnnotationInfo) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *V2reportsAnnotationInfo) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *V2reportsAnnotationInfo) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.

### HasReleaseVersion

`func (o *V2reportsAnnotationInfo) HasReleaseVersion() bool`

HasReleaseVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


