# V2reportsANIMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assembly** | Pointer to **string** |  | [optional] 
**OrganismName** | Pointer to **string** |  | [optional] 
**Category** | Pointer to [**V2reportsANITypeCategory**](V2reportsANITypeCategory.md) |  | [optional] [default to V2REPORTSANITYPECATEGORY_ANI_CATEGORY_UNKNOWN]
**Ani** | Pointer to **float32** |  | [optional] 
**AssemblyCoverage** | Pointer to **float32** |  | [optional] 
**TypeAssemblyCoverage** | Pointer to **float32** |  | [optional] 

## Methods

### NewV2reportsANIMatch

`func NewV2reportsANIMatch() *V2reportsANIMatch`

NewV2reportsANIMatch instantiates a new V2reportsANIMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsANIMatchWithDefaults

`func NewV2reportsANIMatchWithDefaults() *V2reportsANIMatch`

NewV2reportsANIMatchWithDefaults instantiates a new V2reportsANIMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssembly

`func (o *V2reportsANIMatch) GetAssembly() string`

GetAssembly returns the Assembly field if non-nil, zero value otherwise.

### GetAssemblyOk

`func (o *V2reportsANIMatch) GetAssemblyOk() (*string, bool)`

GetAssemblyOk returns a tuple with the Assembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssembly

`func (o *V2reportsANIMatch) SetAssembly(v string)`

SetAssembly sets Assembly field to given value.

### HasAssembly

`func (o *V2reportsANIMatch) HasAssembly() bool`

HasAssembly returns a boolean if a field has been set.

### GetOrganismName

`func (o *V2reportsANIMatch) GetOrganismName() string`

GetOrganismName returns the OrganismName field if non-nil, zero value otherwise.

### GetOrganismNameOk

`func (o *V2reportsANIMatch) GetOrganismNameOk() (*string, bool)`

GetOrganismNameOk returns a tuple with the OrganismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganismName

`func (o *V2reportsANIMatch) SetOrganismName(v string)`

SetOrganismName sets OrganismName field to given value.

### HasOrganismName

`func (o *V2reportsANIMatch) HasOrganismName() bool`

HasOrganismName returns a boolean if a field has been set.

### GetCategory

`func (o *V2reportsANIMatch) GetCategory() V2reportsANITypeCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *V2reportsANIMatch) GetCategoryOk() (*V2reportsANITypeCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *V2reportsANIMatch) SetCategory(v V2reportsANITypeCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *V2reportsANIMatch) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetAni

`func (o *V2reportsANIMatch) GetAni() float32`

GetAni returns the Ani field if non-nil, zero value otherwise.

### GetAniOk

`func (o *V2reportsANIMatch) GetAniOk() (*float32, bool)`

GetAniOk returns a tuple with the Ani field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAni

`func (o *V2reportsANIMatch) SetAni(v float32)`

SetAni sets Ani field to given value.

### HasAni

`func (o *V2reportsANIMatch) HasAni() bool`

HasAni returns a boolean if a field has been set.

### GetAssemblyCoverage

`func (o *V2reportsANIMatch) GetAssemblyCoverage() float32`

GetAssemblyCoverage returns the AssemblyCoverage field if non-nil, zero value otherwise.

### GetAssemblyCoverageOk

`func (o *V2reportsANIMatch) GetAssemblyCoverageOk() (*float32, bool)`

GetAssemblyCoverageOk returns a tuple with the AssemblyCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyCoverage

`func (o *V2reportsANIMatch) SetAssemblyCoverage(v float32)`

SetAssemblyCoverage sets AssemblyCoverage field to given value.

### HasAssemblyCoverage

`func (o *V2reportsANIMatch) HasAssemblyCoverage() bool`

HasAssemblyCoverage returns a boolean if a field has been set.

### GetTypeAssemblyCoverage

`func (o *V2reportsANIMatch) GetTypeAssemblyCoverage() float32`

GetTypeAssemblyCoverage returns the TypeAssemblyCoverage field if non-nil, zero value otherwise.

### GetTypeAssemblyCoverageOk

`func (o *V2reportsANIMatch) GetTypeAssemblyCoverageOk() (*float32, bool)`

GetTypeAssemblyCoverageOk returns a tuple with the TypeAssemblyCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeAssemblyCoverage

`func (o *V2reportsANIMatch) SetTypeAssemblyCoverage(v float32)`

SetTypeAssemblyCoverage sets TypeAssemblyCoverage field to given value.

### HasTypeAssemblyCoverage

`func (o *V2reportsANIMatch) HasTypeAssemblyCoverage() bool`

HasTypeAssemblyCoverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


