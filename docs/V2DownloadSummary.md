# V2DownloadSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordCount** | Pointer to **int32** |  | [optional] 
**AssemblyCount** | Pointer to **int32** |  | [optional] 
**ResourceUpdatedOn** | Pointer to **time.Time** |  | [optional] 
**Hydrated** | Pointer to [**V2DownloadSummaryHydrated**](V2DownloadSummaryHydrated.md) |  | [optional] 
**Dehydrated** | Pointer to [**V2DownloadSummaryDehydrated**](V2DownloadSummaryDehydrated.md) |  | [optional] 
**Errors** | Pointer to [**[]V2reportsError**](V2reportsError.md) |  | [optional] 
**Messages** | Pointer to [**[]V2reportsMessage**](V2reportsMessage.md) |  | [optional] 
**AvailableFiles** | Pointer to [**V2DownloadSummaryAvailableFiles**](V2DownloadSummaryAvailableFiles.md) |  | [optional] 

## Methods

### NewV2DownloadSummary

`func NewV2DownloadSummary() *V2DownloadSummary`

NewV2DownloadSummary instantiates a new V2DownloadSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2DownloadSummaryWithDefaults

`func NewV2DownloadSummaryWithDefaults() *V2DownloadSummary`

NewV2DownloadSummaryWithDefaults instantiates a new V2DownloadSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordCount

`func (o *V2DownloadSummary) GetRecordCount() int32`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *V2DownloadSummary) GetRecordCountOk() (*int32, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *V2DownloadSummary) SetRecordCount(v int32)`

SetRecordCount sets RecordCount field to given value.

### HasRecordCount

`func (o *V2DownloadSummary) HasRecordCount() bool`

HasRecordCount returns a boolean if a field has been set.

### GetAssemblyCount

`func (o *V2DownloadSummary) GetAssemblyCount() int32`

GetAssemblyCount returns the AssemblyCount field if non-nil, zero value otherwise.

### GetAssemblyCountOk

`func (o *V2DownloadSummary) GetAssemblyCountOk() (*int32, bool)`

GetAssemblyCountOk returns a tuple with the AssemblyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyCount

`func (o *V2DownloadSummary) SetAssemblyCount(v int32)`

SetAssemblyCount sets AssemblyCount field to given value.

### HasAssemblyCount

`func (o *V2DownloadSummary) HasAssemblyCount() bool`

HasAssemblyCount returns a boolean if a field has been set.

### GetResourceUpdatedOn

`func (o *V2DownloadSummary) GetResourceUpdatedOn() time.Time`

GetResourceUpdatedOn returns the ResourceUpdatedOn field if non-nil, zero value otherwise.

### GetResourceUpdatedOnOk

`func (o *V2DownloadSummary) GetResourceUpdatedOnOk() (*time.Time, bool)`

GetResourceUpdatedOnOk returns a tuple with the ResourceUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUpdatedOn

`func (o *V2DownloadSummary) SetResourceUpdatedOn(v time.Time)`

SetResourceUpdatedOn sets ResourceUpdatedOn field to given value.

### HasResourceUpdatedOn

`func (o *V2DownloadSummary) HasResourceUpdatedOn() bool`

HasResourceUpdatedOn returns a boolean if a field has been set.

### GetHydrated

`func (o *V2DownloadSummary) GetHydrated() V2DownloadSummaryHydrated`

GetHydrated returns the Hydrated field if non-nil, zero value otherwise.

### GetHydratedOk

`func (o *V2DownloadSummary) GetHydratedOk() (*V2DownloadSummaryHydrated, bool)`

GetHydratedOk returns a tuple with the Hydrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHydrated

`func (o *V2DownloadSummary) SetHydrated(v V2DownloadSummaryHydrated)`

SetHydrated sets Hydrated field to given value.

### HasHydrated

`func (o *V2DownloadSummary) HasHydrated() bool`

HasHydrated returns a boolean if a field has been set.

### GetDehydrated

`func (o *V2DownloadSummary) GetDehydrated() V2DownloadSummaryDehydrated`

GetDehydrated returns the Dehydrated field if non-nil, zero value otherwise.

### GetDehydratedOk

`func (o *V2DownloadSummary) GetDehydratedOk() (*V2DownloadSummaryDehydrated, bool)`

GetDehydratedOk returns a tuple with the Dehydrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDehydrated

`func (o *V2DownloadSummary) SetDehydrated(v V2DownloadSummaryDehydrated)`

SetDehydrated sets Dehydrated field to given value.

### HasDehydrated

`func (o *V2DownloadSummary) HasDehydrated() bool`

HasDehydrated returns a boolean if a field has been set.

### GetErrors

`func (o *V2DownloadSummary) GetErrors() []V2reportsError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V2DownloadSummary) GetErrorsOk() (*[]V2reportsError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V2DownloadSummary) SetErrors(v []V2reportsError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V2DownloadSummary) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMessages

`func (o *V2DownloadSummary) GetMessages() []V2reportsMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *V2DownloadSummary) GetMessagesOk() (*[]V2reportsMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *V2DownloadSummary) SetMessages(v []V2reportsMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *V2DownloadSummary) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetAvailableFiles

`func (o *V2DownloadSummary) GetAvailableFiles() V2DownloadSummaryAvailableFiles`

GetAvailableFiles returns the AvailableFiles field if non-nil, zero value otherwise.

### GetAvailableFilesOk

`func (o *V2DownloadSummary) GetAvailableFilesOk() (*V2DownloadSummaryAvailableFiles, bool)`

GetAvailableFilesOk returns a tuple with the AvailableFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFiles

`func (o *V2DownloadSummary) SetAvailableFiles(v V2DownloadSummaryAvailableFiles)`

SetAvailableFiles sets AvailableFiles field to given value.

### HasAvailableFiles

`func (o *V2DownloadSummary) HasAvailableFiles() bool`

HasAvailableFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


