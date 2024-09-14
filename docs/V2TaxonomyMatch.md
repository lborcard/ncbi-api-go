# V2TaxonomyMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warnings** | Pointer to [**[]V2reportsWarning**](V2reportsWarning.md) |  | [optional] 
**Errors** | Pointer to [**[]V2reportsError**](V2reportsError.md) |  | [optional] 
**Query** | Pointer to **[]string** |  | [optional] 
**Taxonomy** | Pointer to [**V2TaxonomyNode**](V2TaxonomyNode.md) |  | [optional] 

## Methods

### NewV2TaxonomyMatch

`func NewV2TaxonomyMatch() *V2TaxonomyMatch`

NewV2TaxonomyMatch instantiates a new V2TaxonomyMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TaxonomyMatchWithDefaults

`func NewV2TaxonomyMatchWithDefaults() *V2TaxonomyMatch`

NewV2TaxonomyMatchWithDefaults instantiates a new V2TaxonomyMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarnings

`func (o *V2TaxonomyMatch) GetWarnings() []V2reportsWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V2TaxonomyMatch) GetWarningsOk() (*[]V2reportsWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V2TaxonomyMatch) SetWarnings(v []V2reportsWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V2TaxonomyMatch) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetErrors

`func (o *V2TaxonomyMatch) GetErrors() []V2reportsError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V2TaxonomyMatch) GetErrorsOk() (*[]V2reportsError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V2TaxonomyMatch) SetErrors(v []V2reportsError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V2TaxonomyMatch) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetQuery

`func (o *V2TaxonomyMatch) GetQuery() []string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *V2TaxonomyMatch) GetQueryOk() (*[]string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *V2TaxonomyMatch) SetQuery(v []string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *V2TaxonomyMatch) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTaxonomy

`func (o *V2TaxonomyMatch) GetTaxonomy() V2TaxonomyNode`

GetTaxonomy returns the Taxonomy field if non-nil, zero value otherwise.

### GetTaxonomyOk

`func (o *V2TaxonomyMatch) GetTaxonomyOk() (*V2TaxonomyNode, bool)`

GetTaxonomyOk returns a tuple with the Taxonomy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomy

`func (o *V2TaxonomyMatch) SetTaxonomy(v V2TaxonomyNode)`

SetTaxonomy sets Taxonomy field to given value.

### HasTaxonomy

`func (o *V2TaxonomyMatch) HasTaxonomy() bool`

HasTaxonomy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


