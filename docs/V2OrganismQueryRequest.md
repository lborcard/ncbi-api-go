# V2OrganismQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganismQuery** | Pointer to **string** |  | [optional] 
**TaxonQuery** | Pointer to **string** |  | [optional] 
**TaxRankFilter** | Pointer to [**V2OrganismQueryRequestTaxRankFilter**](V2OrganismQueryRequestTaxRankFilter.md) |  | [optional] [default to V2ORGANISMQUERYREQUESTTAXRANKFILTER_SPECIES]
**TaxonResourceFilter** | Pointer to [**V2OrganismQueryRequestTaxonResourceFilter**](V2OrganismQueryRequestTaxonResourceFilter.md) |  | [optional] [default to V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_ALL]
**ExactMatch** | Pointer to **bool** |  | [optional] 

## Methods

### NewV2OrganismQueryRequest

`func NewV2OrganismQueryRequest() *V2OrganismQueryRequest`

NewV2OrganismQueryRequest instantiates a new V2OrganismQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2OrganismQueryRequestWithDefaults

`func NewV2OrganismQueryRequestWithDefaults() *V2OrganismQueryRequest`

NewV2OrganismQueryRequestWithDefaults instantiates a new V2OrganismQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganismQuery

`func (o *V2OrganismQueryRequest) GetOrganismQuery() string`

GetOrganismQuery returns the OrganismQuery field if non-nil, zero value otherwise.

### GetOrganismQueryOk

`func (o *V2OrganismQueryRequest) GetOrganismQueryOk() (*string, bool)`

GetOrganismQueryOk returns a tuple with the OrganismQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganismQuery

`func (o *V2OrganismQueryRequest) SetOrganismQuery(v string)`

SetOrganismQuery sets OrganismQuery field to given value.

### HasOrganismQuery

`func (o *V2OrganismQueryRequest) HasOrganismQuery() bool`

HasOrganismQuery returns a boolean if a field has been set.

### GetTaxonQuery

`func (o *V2OrganismQueryRequest) GetTaxonQuery() string`

GetTaxonQuery returns the TaxonQuery field if non-nil, zero value otherwise.

### GetTaxonQueryOk

`func (o *V2OrganismQueryRequest) GetTaxonQueryOk() (*string, bool)`

GetTaxonQueryOk returns a tuple with the TaxonQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonQuery

`func (o *V2OrganismQueryRequest) SetTaxonQuery(v string)`

SetTaxonQuery sets TaxonQuery field to given value.

### HasTaxonQuery

`func (o *V2OrganismQueryRequest) HasTaxonQuery() bool`

HasTaxonQuery returns a boolean if a field has been set.

### GetTaxRankFilter

`func (o *V2OrganismQueryRequest) GetTaxRankFilter() V2OrganismQueryRequestTaxRankFilter`

GetTaxRankFilter returns the TaxRankFilter field if non-nil, zero value otherwise.

### GetTaxRankFilterOk

`func (o *V2OrganismQueryRequest) GetTaxRankFilterOk() (*V2OrganismQueryRequestTaxRankFilter, bool)`

GetTaxRankFilterOk returns a tuple with the TaxRankFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRankFilter

`func (o *V2OrganismQueryRequest) SetTaxRankFilter(v V2OrganismQueryRequestTaxRankFilter)`

SetTaxRankFilter sets TaxRankFilter field to given value.

### HasTaxRankFilter

`func (o *V2OrganismQueryRequest) HasTaxRankFilter() bool`

HasTaxRankFilter returns a boolean if a field has been set.

### GetTaxonResourceFilter

`func (o *V2OrganismQueryRequest) GetTaxonResourceFilter() V2OrganismQueryRequestTaxonResourceFilter`

GetTaxonResourceFilter returns the TaxonResourceFilter field if non-nil, zero value otherwise.

### GetTaxonResourceFilterOk

`func (o *V2OrganismQueryRequest) GetTaxonResourceFilterOk() (*V2OrganismQueryRequestTaxonResourceFilter, bool)`

GetTaxonResourceFilterOk returns a tuple with the TaxonResourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonResourceFilter

`func (o *V2OrganismQueryRequest) SetTaxonResourceFilter(v V2OrganismQueryRequestTaxonResourceFilter)`

SetTaxonResourceFilter sets TaxonResourceFilter field to given value.

### HasTaxonResourceFilter

`func (o *V2OrganismQueryRequest) HasTaxonResourceFilter() bool`

HasTaxonResourceFilter returns a boolean if a field has been set.

### GetExactMatch

`func (o *V2OrganismQueryRequest) GetExactMatch() bool`

GetExactMatch returns the ExactMatch field if non-nil, zero value otherwise.

### GetExactMatchOk

`func (o *V2OrganismQueryRequest) GetExactMatchOk() (*bool, bool)`

GetExactMatchOk returns a tuple with the ExactMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactMatch

`func (o *V2OrganismQueryRequest) SetExactMatch(v bool)`

SetExactMatch sets ExactMatch field to given value.

### HasExactMatch

`func (o *V2OrganismQueryRequest) HasExactMatch() bool`

HasExactMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


