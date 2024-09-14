# V2TaxonomyMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]V2reportsMessage**](V2reportsMessage.md) |  | [optional] 
**TaxonomyNodes** | Pointer to [**[]V2TaxonomyMatch**](V2TaxonomyMatch.md) |  | [optional] 

## Methods

### NewV2TaxonomyMetadataResponse

`func NewV2TaxonomyMetadataResponse() *V2TaxonomyMetadataResponse`

NewV2TaxonomyMetadataResponse instantiates a new V2TaxonomyMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TaxonomyMetadataResponseWithDefaults

`func NewV2TaxonomyMetadataResponseWithDefaults() *V2TaxonomyMetadataResponse`

NewV2TaxonomyMetadataResponseWithDefaults instantiates a new V2TaxonomyMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *V2TaxonomyMetadataResponse) GetMessages() []V2reportsMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *V2TaxonomyMetadataResponse) GetMessagesOk() (*[]V2reportsMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *V2TaxonomyMetadataResponse) SetMessages(v []V2reportsMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *V2TaxonomyMetadataResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetTaxonomyNodes

`func (o *V2TaxonomyMetadataResponse) GetTaxonomyNodes() []V2TaxonomyMatch`

GetTaxonomyNodes returns the TaxonomyNodes field if non-nil, zero value otherwise.

### GetTaxonomyNodesOk

`func (o *V2TaxonomyMetadataResponse) GetTaxonomyNodesOk() (*[]V2TaxonomyMatch, bool)`

GetTaxonomyNodesOk returns a tuple with the TaxonomyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyNodes

`func (o *V2TaxonomyMetadataResponse) SetTaxonomyNodes(v []V2TaxonomyMatch)`

SetTaxonomyNodes sets TaxonomyNodes field to given value.

### HasTaxonomyNodes

`func (o *V2TaxonomyMetadataResponse) HasTaxonomyNodes() bool`

HasTaxonomyNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


