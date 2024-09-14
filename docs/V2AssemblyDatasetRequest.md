# V2AssemblyDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessions** | Pointer to **[]string** |  | [optional] 
**Chromosomes** | Pointer to **[]string** |  | [optional] 
**IncludeAnnotationType** | Pointer to [**[]V2AnnotationForAssemblyType**](V2AnnotationForAssemblyType.md) |  | [optional] 
**Hydrated** | Pointer to [**V2AssemblyDatasetRequestResolution**](V2AssemblyDatasetRequestResolution.md) |  | [optional] [default to V2ASSEMBLYDATASETREQUESTRESOLUTION_FULLY_HYDRATED]
**IncludeTsv** | Pointer to **bool** |  | [optional] 
**ExpDebugValues** | Pointer to **string** |  | [optional] 

## Methods

### NewV2AssemblyDatasetRequest

`func NewV2AssemblyDatasetRequest() *V2AssemblyDatasetRequest`

NewV2AssemblyDatasetRequest instantiates a new V2AssemblyDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssemblyDatasetRequestWithDefaults

`func NewV2AssemblyDatasetRequestWithDefaults() *V2AssemblyDatasetRequest`

NewV2AssemblyDatasetRequestWithDefaults instantiates a new V2AssemblyDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessions

`func (o *V2AssemblyDatasetRequest) GetAccessions() []string`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V2AssemblyDatasetRequest) GetAccessionsOk() (*[]string, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V2AssemblyDatasetRequest) SetAccessions(v []string)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V2AssemblyDatasetRequest) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetChromosomes

`func (o *V2AssemblyDatasetRequest) GetChromosomes() []string`

GetChromosomes returns the Chromosomes field if non-nil, zero value otherwise.

### GetChromosomesOk

`func (o *V2AssemblyDatasetRequest) GetChromosomesOk() (*[]string, bool)`

GetChromosomesOk returns a tuple with the Chromosomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosomes

`func (o *V2AssemblyDatasetRequest) SetChromosomes(v []string)`

SetChromosomes sets Chromosomes field to given value.

### HasChromosomes

`func (o *V2AssemblyDatasetRequest) HasChromosomes() bool`

HasChromosomes returns a boolean if a field has been set.

### GetIncludeAnnotationType

`func (o *V2AssemblyDatasetRequest) GetIncludeAnnotationType() []V2AnnotationForAssemblyType`

GetIncludeAnnotationType returns the IncludeAnnotationType field if non-nil, zero value otherwise.

### GetIncludeAnnotationTypeOk

`func (o *V2AssemblyDatasetRequest) GetIncludeAnnotationTypeOk() (*[]V2AnnotationForAssemblyType, bool)`

GetIncludeAnnotationTypeOk returns a tuple with the IncludeAnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnnotationType

`func (o *V2AssemblyDatasetRequest) SetIncludeAnnotationType(v []V2AnnotationForAssemblyType)`

SetIncludeAnnotationType sets IncludeAnnotationType field to given value.

### HasIncludeAnnotationType

`func (o *V2AssemblyDatasetRequest) HasIncludeAnnotationType() bool`

HasIncludeAnnotationType returns a boolean if a field has been set.

### GetHydrated

`func (o *V2AssemblyDatasetRequest) GetHydrated() V2AssemblyDatasetRequestResolution`

GetHydrated returns the Hydrated field if non-nil, zero value otherwise.

### GetHydratedOk

`func (o *V2AssemblyDatasetRequest) GetHydratedOk() (*V2AssemblyDatasetRequestResolution, bool)`

GetHydratedOk returns a tuple with the Hydrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHydrated

`func (o *V2AssemblyDatasetRequest) SetHydrated(v V2AssemblyDatasetRequestResolution)`

SetHydrated sets Hydrated field to given value.

### HasHydrated

`func (o *V2AssemblyDatasetRequest) HasHydrated() bool`

HasHydrated returns a boolean if a field has been set.

### GetIncludeTsv

`func (o *V2AssemblyDatasetRequest) GetIncludeTsv() bool`

GetIncludeTsv returns the IncludeTsv field if non-nil, zero value otherwise.

### GetIncludeTsvOk

`func (o *V2AssemblyDatasetRequest) GetIncludeTsvOk() (*bool, bool)`

GetIncludeTsvOk returns a tuple with the IncludeTsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTsv

`func (o *V2AssemblyDatasetRequest) SetIncludeTsv(v bool)`

SetIncludeTsv sets IncludeTsv field to given value.

### HasIncludeTsv

`func (o *V2AssemblyDatasetRequest) HasIncludeTsv() bool`

HasIncludeTsv returns a boolean if a field has been set.

### GetExpDebugValues

`func (o *V2AssemblyDatasetRequest) GetExpDebugValues() string`

GetExpDebugValues returns the ExpDebugValues field if non-nil, zero value otherwise.

### GetExpDebugValuesOk

`func (o *V2AssemblyDatasetRequest) GetExpDebugValuesOk() (*string, bool)`

GetExpDebugValuesOk returns a tuple with the ExpDebugValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDebugValues

`func (o *V2AssemblyDatasetRequest) SetExpDebugValues(v string)`

SetExpDebugValues sets ExpDebugValues field to given value.

### HasExpDebugValues

`func (o *V2AssemblyDatasetRequest) HasExpDebugValues() bool`

HasExpDebugValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


