# V2AssemblyDatasetDescriptorsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceOnly** | Pointer to **bool** |  | [optional] 
**AssemblySource** | Pointer to [**V2AssemblyDatasetDescriptorsFilterAssemblySource**](V2AssemblyDatasetDescriptorsFilterAssemblySource.md) |  | [optional] [default to V2ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYSOURCE_ALL]
**HasAnnotation** | Pointer to **bool** |  | [optional] 
**ExcludePairedReports** | Pointer to **bool** |  | [optional] 
**ExcludeAtypical** | Pointer to **bool** |  | [optional] 
**AssemblyVersion** | Pointer to [**V2AssemblyDatasetDescriptorsFilterAssemblyVersion**](V2AssemblyDatasetDescriptorsFilterAssemblyVersion.md) |  | [optional] [default to V2ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYVERSION_CURRENT]
**AssemblyLevel** | Pointer to [**[]V2reportsAssemblyLevel**](V2reportsAssemblyLevel.md) |  | [optional] 
**FirstReleaseDate** | Pointer to **time.Time** |  | [optional] 
**LastReleaseDate** | Pointer to **time.Time** |  | [optional] 
**SearchText** | Pointer to **[]string** |  | [optional] 
**IsMetagenomeDerived** | Pointer to [**V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter**](V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter.md) |  | [optional] [default to V2ASSEMBLYDATASETDESCRIPTORSFILTERMETAGENOMEDERIVEDFILTER_METAGENOME_DERIVED_UNSET]
**IsTypeMaterial** | Pointer to **bool** |  | [optional] 
**IsIctvExemplar** | Pointer to **bool** |  | [optional] 
**ExcludeMultiIsolate** | Pointer to **bool** |  | [optional] 
**TypeMaterialCategory** | Pointer to [**V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory**](V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory.md) |  | [optional] [default to V2ASSEMBLYDATASETDESCRIPTORSFILTERTYPEMATERIALCATEGORY_NONE]

## Methods

### NewV2AssemblyDatasetDescriptorsFilter

`func NewV2AssemblyDatasetDescriptorsFilter() *V2AssemblyDatasetDescriptorsFilter`

NewV2AssemblyDatasetDescriptorsFilter instantiates a new V2AssemblyDatasetDescriptorsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2AssemblyDatasetDescriptorsFilterWithDefaults

`func NewV2AssemblyDatasetDescriptorsFilterWithDefaults() *V2AssemblyDatasetDescriptorsFilter`

NewV2AssemblyDatasetDescriptorsFilterWithDefaults instantiates a new V2AssemblyDatasetDescriptorsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceOnly

`func (o *V2AssemblyDatasetDescriptorsFilter) GetReferenceOnly() bool`

GetReferenceOnly returns the ReferenceOnly field if non-nil, zero value otherwise.

### GetReferenceOnlyOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetReferenceOnlyOk() (*bool, bool)`

GetReferenceOnlyOk returns a tuple with the ReferenceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOnly

`func (o *V2AssemblyDatasetDescriptorsFilter) SetReferenceOnly(v bool)`

SetReferenceOnly sets ReferenceOnly field to given value.

### HasReferenceOnly

`func (o *V2AssemblyDatasetDescriptorsFilter) HasReferenceOnly() bool`

HasReferenceOnly returns a boolean if a field has been set.

### GetAssemblySource

`func (o *V2AssemblyDatasetDescriptorsFilter) GetAssemblySource() V2AssemblyDatasetDescriptorsFilterAssemblySource`

GetAssemblySource returns the AssemblySource field if non-nil, zero value otherwise.

### GetAssemblySourceOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetAssemblySourceOk() (*V2AssemblyDatasetDescriptorsFilterAssemblySource, bool)`

GetAssemblySourceOk returns a tuple with the AssemblySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblySource

`func (o *V2AssemblyDatasetDescriptorsFilter) SetAssemblySource(v V2AssemblyDatasetDescriptorsFilterAssemblySource)`

SetAssemblySource sets AssemblySource field to given value.

### HasAssemblySource

`func (o *V2AssemblyDatasetDescriptorsFilter) HasAssemblySource() bool`

HasAssemblySource returns a boolean if a field has been set.

### GetHasAnnotation

`func (o *V2AssemblyDatasetDescriptorsFilter) GetHasAnnotation() bool`

GetHasAnnotation returns the HasAnnotation field if non-nil, zero value otherwise.

### GetHasAnnotationOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetHasAnnotationOk() (*bool, bool)`

GetHasAnnotationOk returns a tuple with the HasAnnotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAnnotation

`func (o *V2AssemblyDatasetDescriptorsFilter) SetHasAnnotation(v bool)`

SetHasAnnotation sets HasAnnotation field to given value.

### HasHasAnnotation

`func (o *V2AssemblyDatasetDescriptorsFilter) HasHasAnnotation() bool`

HasHasAnnotation returns a boolean if a field has been set.

### GetExcludePairedReports

`func (o *V2AssemblyDatasetDescriptorsFilter) GetExcludePairedReports() bool`

GetExcludePairedReports returns the ExcludePairedReports field if non-nil, zero value otherwise.

### GetExcludePairedReportsOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetExcludePairedReportsOk() (*bool, bool)`

GetExcludePairedReportsOk returns a tuple with the ExcludePairedReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePairedReports

`func (o *V2AssemblyDatasetDescriptorsFilter) SetExcludePairedReports(v bool)`

SetExcludePairedReports sets ExcludePairedReports field to given value.

### HasExcludePairedReports

`func (o *V2AssemblyDatasetDescriptorsFilter) HasExcludePairedReports() bool`

HasExcludePairedReports returns a boolean if a field has been set.

### GetExcludeAtypical

`func (o *V2AssemblyDatasetDescriptorsFilter) GetExcludeAtypical() bool`

GetExcludeAtypical returns the ExcludeAtypical field if non-nil, zero value otherwise.

### GetExcludeAtypicalOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetExcludeAtypicalOk() (*bool, bool)`

GetExcludeAtypicalOk returns a tuple with the ExcludeAtypical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAtypical

`func (o *V2AssemblyDatasetDescriptorsFilter) SetExcludeAtypical(v bool)`

SetExcludeAtypical sets ExcludeAtypical field to given value.

### HasExcludeAtypical

`func (o *V2AssemblyDatasetDescriptorsFilter) HasExcludeAtypical() bool`

HasExcludeAtypical returns a boolean if a field has been set.

### GetAssemblyVersion

`func (o *V2AssemblyDatasetDescriptorsFilter) GetAssemblyVersion() V2AssemblyDatasetDescriptorsFilterAssemblyVersion`

GetAssemblyVersion returns the AssemblyVersion field if non-nil, zero value otherwise.

### GetAssemblyVersionOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetAssemblyVersionOk() (*V2AssemblyDatasetDescriptorsFilterAssemblyVersion, bool)`

GetAssemblyVersionOk returns a tuple with the AssemblyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyVersion

`func (o *V2AssemblyDatasetDescriptorsFilter) SetAssemblyVersion(v V2AssemblyDatasetDescriptorsFilterAssemblyVersion)`

SetAssemblyVersion sets AssemblyVersion field to given value.

### HasAssemblyVersion

`func (o *V2AssemblyDatasetDescriptorsFilter) HasAssemblyVersion() bool`

HasAssemblyVersion returns a boolean if a field has been set.

### GetAssemblyLevel

`func (o *V2AssemblyDatasetDescriptorsFilter) GetAssemblyLevel() []V2reportsAssemblyLevel`

GetAssemblyLevel returns the AssemblyLevel field if non-nil, zero value otherwise.

### GetAssemblyLevelOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetAssemblyLevelOk() (*[]V2reportsAssemblyLevel, bool)`

GetAssemblyLevelOk returns a tuple with the AssemblyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyLevel

`func (o *V2AssemblyDatasetDescriptorsFilter) SetAssemblyLevel(v []V2reportsAssemblyLevel)`

SetAssemblyLevel sets AssemblyLevel field to given value.

### HasAssemblyLevel

`func (o *V2AssemblyDatasetDescriptorsFilter) HasAssemblyLevel() bool`

HasAssemblyLevel returns a boolean if a field has been set.

### GetFirstReleaseDate

`func (o *V2AssemblyDatasetDescriptorsFilter) GetFirstReleaseDate() time.Time`

GetFirstReleaseDate returns the FirstReleaseDate field if non-nil, zero value otherwise.

### GetFirstReleaseDateOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetFirstReleaseDateOk() (*time.Time, bool)`

GetFirstReleaseDateOk returns a tuple with the FirstReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstReleaseDate

`func (o *V2AssemblyDatasetDescriptorsFilter) SetFirstReleaseDate(v time.Time)`

SetFirstReleaseDate sets FirstReleaseDate field to given value.

### HasFirstReleaseDate

`func (o *V2AssemblyDatasetDescriptorsFilter) HasFirstReleaseDate() bool`

HasFirstReleaseDate returns a boolean if a field has been set.

### GetLastReleaseDate

`func (o *V2AssemblyDatasetDescriptorsFilter) GetLastReleaseDate() time.Time`

GetLastReleaseDate returns the LastReleaseDate field if non-nil, zero value otherwise.

### GetLastReleaseDateOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetLastReleaseDateOk() (*time.Time, bool)`

GetLastReleaseDateOk returns a tuple with the LastReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReleaseDate

`func (o *V2AssemblyDatasetDescriptorsFilter) SetLastReleaseDate(v time.Time)`

SetLastReleaseDate sets LastReleaseDate field to given value.

### HasLastReleaseDate

`func (o *V2AssemblyDatasetDescriptorsFilter) HasLastReleaseDate() bool`

HasLastReleaseDate returns a boolean if a field has been set.

### GetSearchText

`func (o *V2AssemblyDatasetDescriptorsFilter) GetSearchText() []string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetSearchTextOk() (*[]string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *V2AssemblyDatasetDescriptorsFilter) SetSearchText(v []string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *V2AssemblyDatasetDescriptorsFilter) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetIsMetagenomeDerived

`func (o *V2AssemblyDatasetDescriptorsFilter) GetIsMetagenomeDerived() V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter`

GetIsMetagenomeDerived returns the IsMetagenomeDerived field if non-nil, zero value otherwise.

### GetIsMetagenomeDerivedOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetIsMetagenomeDerivedOk() (*V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter, bool)`

GetIsMetagenomeDerivedOk returns a tuple with the IsMetagenomeDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMetagenomeDerived

`func (o *V2AssemblyDatasetDescriptorsFilter) SetIsMetagenomeDerived(v V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter)`

SetIsMetagenomeDerived sets IsMetagenomeDerived field to given value.

### HasIsMetagenomeDerived

`func (o *V2AssemblyDatasetDescriptorsFilter) HasIsMetagenomeDerived() bool`

HasIsMetagenomeDerived returns a boolean if a field has been set.

### GetIsTypeMaterial

`func (o *V2AssemblyDatasetDescriptorsFilter) GetIsTypeMaterial() bool`

GetIsTypeMaterial returns the IsTypeMaterial field if non-nil, zero value otherwise.

### GetIsTypeMaterialOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetIsTypeMaterialOk() (*bool, bool)`

GetIsTypeMaterialOk returns a tuple with the IsTypeMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTypeMaterial

`func (o *V2AssemblyDatasetDescriptorsFilter) SetIsTypeMaterial(v bool)`

SetIsTypeMaterial sets IsTypeMaterial field to given value.

### HasIsTypeMaterial

`func (o *V2AssemblyDatasetDescriptorsFilter) HasIsTypeMaterial() bool`

HasIsTypeMaterial returns a boolean if a field has been set.

### GetIsIctvExemplar

`func (o *V2AssemblyDatasetDescriptorsFilter) GetIsIctvExemplar() bool`

GetIsIctvExemplar returns the IsIctvExemplar field if non-nil, zero value otherwise.

### GetIsIctvExemplarOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetIsIctvExemplarOk() (*bool, bool)`

GetIsIctvExemplarOk returns a tuple with the IsIctvExemplar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIctvExemplar

`func (o *V2AssemblyDatasetDescriptorsFilter) SetIsIctvExemplar(v bool)`

SetIsIctvExemplar sets IsIctvExemplar field to given value.

### HasIsIctvExemplar

`func (o *V2AssemblyDatasetDescriptorsFilter) HasIsIctvExemplar() bool`

HasIsIctvExemplar returns a boolean if a field has been set.

### GetExcludeMultiIsolate

`func (o *V2AssemblyDatasetDescriptorsFilter) GetExcludeMultiIsolate() bool`

GetExcludeMultiIsolate returns the ExcludeMultiIsolate field if non-nil, zero value otherwise.

### GetExcludeMultiIsolateOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetExcludeMultiIsolateOk() (*bool, bool)`

GetExcludeMultiIsolateOk returns a tuple with the ExcludeMultiIsolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeMultiIsolate

`func (o *V2AssemblyDatasetDescriptorsFilter) SetExcludeMultiIsolate(v bool)`

SetExcludeMultiIsolate sets ExcludeMultiIsolate field to given value.

### HasExcludeMultiIsolate

`func (o *V2AssemblyDatasetDescriptorsFilter) HasExcludeMultiIsolate() bool`

HasExcludeMultiIsolate returns a boolean if a field has been set.

### GetTypeMaterialCategory

`func (o *V2AssemblyDatasetDescriptorsFilter) GetTypeMaterialCategory() V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory`

GetTypeMaterialCategory returns the TypeMaterialCategory field if non-nil, zero value otherwise.

### GetTypeMaterialCategoryOk

`func (o *V2AssemblyDatasetDescriptorsFilter) GetTypeMaterialCategoryOk() (*V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory, bool)`

GetTypeMaterialCategoryOk returns a tuple with the TypeMaterialCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeMaterialCategory

`func (o *V2AssemblyDatasetDescriptorsFilter) SetTypeMaterialCategory(v V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory)`

SetTypeMaterialCategory sets TypeMaterialCategory field to given value.

### HasTypeMaterialCategory

`func (o *V2AssemblyDatasetDescriptorsFilter) HasTypeMaterialCategory() bool`

HasTypeMaterialCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


