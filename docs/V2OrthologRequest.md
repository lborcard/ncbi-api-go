# V2OrthologRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneId** | Pointer to **int32** |  | [optional] 
**ReturnedContent** | Pointer to [**V2OrthologRequestContentType**](V2OrthologRequestContentType.md) |  | [optional] [default to V2ORTHOLOGREQUESTCONTENTTYPE_COMPLETE]
**TaxonFilter** | Pointer to **[]string** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewV2OrthologRequest

`func NewV2OrthologRequest() *V2OrthologRequest`

NewV2OrthologRequest instantiates a new V2OrthologRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2OrthologRequestWithDefaults

`func NewV2OrthologRequestWithDefaults() *V2OrthologRequest`

NewV2OrthologRequestWithDefaults instantiates a new V2OrthologRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneId

`func (o *V2OrthologRequest) GetGeneId() int32`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *V2OrthologRequest) GetGeneIdOk() (*int32, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *V2OrthologRequest) SetGeneId(v int32)`

SetGeneId sets GeneId field to given value.

### HasGeneId

`func (o *V2OrthologRequest) HasGeneId() bool`

HasGeneId returns a boolean if a field has been set.

### GetReturnedContent

`func (o *V2OrthologRequest) GetReturnedContent() V2OrthologRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V2OrthologRequest) GetReturnedContentOk() (*V2OrthologRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V2OrthologRequest) SetReturnedContent(v V2OrthologRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V2OrthologRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.

### GetTaxonFilter

`func (o *V2OrthologRequest) GetTaxonFilter() []string`

GetTaxonFilter returns the TaxonFilter field if non-nil, zero value otherwise.

### GetTaxonFilterOk

`func (o *V2OrthologRequest) GetTaxonFilterOk() (*[]string, bool)`

GetTaxonFilterOk returns a tuple with the TaxonFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonFilter

`func (o *V2OrthologRequest) SetTaxonFilter(v []string)`

SetTaxonFilter sets TaxonFilter field to given value.

### HasTaxonFilter

`func (o *V2OrthologRequest) HasTaxonFilter() bool`

HasTaxonFilter returns a boolean if a field has been set.

### GetPageSize

`func (o *V2OrthologRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V2OrthologRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V2OrthologRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *V2OrthologRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *V2OrthologRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V2OrthologRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V2OrthologRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V2OrthologRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


