# V2reportsTaxonomyReportMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taxonomy** | Pointer to [**V2reportsTaxonomyNode**](V2reportsTaxonomyNode.md) |  | [optional] 
**Query** | Pointer to **[]string** |  | [optional] 
**Warning** | Pointer to [**V2reportsWarning**](V2reportsWarning.md) |  | [optional] 
**Errors** | Pointer to [**[]V2reportsError**](V2reportsError.md) |  | [optional] 

## Methods

### NewV2reportsTaxonomyReportMatch

`func NewV2reportsTaxonomyReportMatch() *V2reportsTaxonomyReportMatch`

NewV2reportsTaxonomyReportMatch instantiates a new V2reportsTaxonomyReportMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsTaxonomyReportMatchWithDefaults

`func NewV2reportsTaxonomyReportMatchWithDefaults() *V2reportsTaxonomyReportMatch`

NewV2reportsTaxonomyReportMatchWithDefaults instantiates a new V2reportsTaxonomyReportMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxonomy

`func (o *V2reportsTaxonomyReportMatch) GetTaxonomy() V2reportsTaxonomyNode`

GetTaxonomy returns the Taxonomy field if non-nil, zero value otherwise.

### GetTaxonomyOk

`func (o *V2reportsTaxonomyReportMatch) GetTaxonomyOk() (*V2reportsTaxonomyNode, bool)`

GetTaxonomyOk returns a tuple with the Taxonomy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomy

`func (o *V2reportsTaxonomyReportMatch) SetTaxonomy(v V2reportsTaxonomyNode)`

SetTaxonomy sets Taxonomy field to given value.

### HasTaxonomy

`func (o *V2reportsTaxonomyReportMatch) HasTaxonomy() bool`

HasTaxonomy returns a boolean if a field has been set.

### GetQuery

`func (o *V2reportsTaxonomyReportMatch) GetQuery() []string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *V2reportsTaxonomyReportMatch) GetQueryOk() (*[]string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *V2reportsTaxonomyReportMatch) SetQuery(v []string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *V2reportsTaxonomyReportMatch) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetWarning

`func (o *V2reportsTaxonomyReportMatch) GetWarning() V2reportsWarning`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *V2reportsTaxonomyReportMatch) GetWarningOk() (*V2reportsWarning, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *V2reportsTaxonomyReportMatch) SetWarning(v V2reportsWarning)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *V2reportsTaxonomyReportMatch) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetErrors

`func (o *V2reportsTaxonomyReportMatch) GetErrors() []V2reportsError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V2reportsTaxonomyReportMatch) GetErrorsOk() (*[]V2reportsError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V2reportsTaxonomyReportMatch) SetErrors(v []V2reportsError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V2reportsTaxonomyReportMatch) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


