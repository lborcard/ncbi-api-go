# V2TaxonomyFilteredSubtreeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootNodes** | Pointer to **[]int32** |  | [optional] 
**Edges** | Pointer to [**V2TaxonomyFilteredSubtreeResponseEdgesEntry**](V2TaxonomyFilteredSubtreeResponseEdgesEntry.md) |  | [optional] 
**Warnings** | Pointer to [**[]V2reportsWarning**](V2reportsWarning.md) |  | [optional] 
**Errors** | Pointer to [**[]V2reportsError**](V2reportsError.md) |  | [optional] 

## Methods

### NewV2TaxonomyFilteredSubtreeResponse

`func NewV2TaxonomyFilteredSubtreeResponse() *V2TaxonomyFilteredSubtreeResponse`

NewV2TaxonomyFilteredSubtreeResponse instantiates a new V2TaxonomyFilteredSubtreeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TaxonomyFilteredSubtreeResponseWithDefaults

`func NewV2TaxonomyFilteredSubtreeResponseWithDefaults() *V2TaxonomyFilteredSubtreeResponse`

NewV2TaxonomyFilteredSubtreeResponseWithDefaults instantiates a new V2TaxonomyFilteredSubtreeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootNodes

`func (o *V2TaxonomyFilteredSubtreeResponse) GetRootNodes() []int32`

GetRootNodes returns the RootNodes field if non-nil, zero value otherwise.

### GetRootNodesOk

`func (o *V2TaxonomyFilteredSubtreeResponse) GetRootNodesOk() (*[]int32, bool)`

GetRootNodesOk returns a tuple with the RootNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootNodes

`func (o *V2TaxonomyFilteredSubtreeResponse) SetRootNodes(v []int32)`

SetRootNodes sets RootNodes field to given value.

### HasRootNodes

`func (o *V2TaxonomyFilteredSubtreeResponse) HasRootNodes() bool`

HasRootNodes returns a boolean if a field has been set.

### GetEdges

`func (o *V2TaxonomyFilteredSubtreeResponse) GetEdges() V2TaxonomyFilteredSubtreeResponseEdgesEntry`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *V2TaxonomyFilteredSubtreeResponse) GetEdgesOk() (*V2TaxonomyFilteredSubtreeResponseEdgesEntry, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *V2TaxonomyFilteredSubtreeResponse) SetEdges(v V2TaxonomyFilteredSubtreeResponseEdgesEntry)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *V2TaxonomyFilteredSubtreeResponse) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetWarnings

`func (o *V2TaxonomyFilteredSubtreeResponse) GetWarnings() []V2reportsWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V2TaxonomyFilteredSubtreeResponse) GetWarningsOk() (*[]V2reportsWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V2TaxonomyFilteredSubtreeResponse) SetWarnings(v []V2reportsWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V2TaxonomyFilteredSubtreeResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetErrors

`func (o *V2TaxonomyFilteredSubtreeResponse) GetErrors() []V2reportsError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V2TaxonomyFilteredSubtreeResponse) GetErrorsOk() (*[]V2reportsError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V2TaxonomyFilteredSubtreeResponse) SetErrors(v []V2reportsError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V2TaxonomyFilteredSubtreeResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


