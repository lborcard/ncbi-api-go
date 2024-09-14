# V2ProkaryoteGeneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessions** | Pointer to **[]string** |  | [optional] 
**IncludeAnnotationType** | Pointer to [**[]V2Fasta**](V2Fasta.md) |  | [optional] 
**GeneFlankConfig** | Pointer to [**V2ProkaryoteGeneRequestGeneFlankConfig**](V2ProkaryoteGeneRequestGeneFlankConfig.md) |  | [optional] 
**Taxon** | Pointer to **string** |  | [optional] 

## Methods

### NewV2ProkaryoteGeneRequest

`func NewV2ProkaryoteGeneRequest() *V2ProkaryoteGeneRequest`

NewV2ProkaryoteGeneRequest instantiates a new V2ProkaryoteGeneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ProkaryoteGeneRequestWithDefaults

`func NewV2ProkaryoteGeneRequestWithDefaults() *V2ProkaryoteGeneRequest`

NewV2ProkaryoteGeneRequestWithDefaults instantiates a new V2ProkaryoteGeneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessions

`func (o *V2ProkaryoteGeneRequest) GetAccessions() []string`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V2ProkaryoteGeneRequest) GetAccessionsOk() (*[]string, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V2ProkaryoteGeneRequest) SetAccessions(v []string)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V2ProkaryoteGeneRequest) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetIncludeAnnotationType

`func (o *V2ProkaryoteGeneRequest) GetIncludeAnnotationType() []V2Fasta`

GetIncludeAnnotationType returns the IncludeAnnotationType field if non-nil, zero value otherwise.

### GetIncludeAnnotationTypeOk

`func (o *V2ProkaryoteGeneRequest) GetIncludeAnnotationTypeOk() (*[]V2Fasta, bool)`

GetIncludeAnnotationTypeOk returns a tuple with the IncludeAnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnnotationType

`func (o *V2ProkaryoteGeneRequest) SetIncludeAnnotationType(v []V2Fasta)`

SetIncludeAnnotationType sets IncludeAnnotationType field to given value.

### HasIncludeAnnotationType

`func (o *V2ProkaryoteGeneRequest) HasIncludeAnnotationType() bool`

HasIncludeAnnotationType returns a boolean if a field has been set.

### GetGeneFlankConfig

`func (o *V2ProkaryoteGeneRequest) GetGeneFlankConfig() V2ProkaryoteGeneRequestGeneFlankConfig`

GetGeneFlankConfig returns the GeneFlankConfig field if non-nil, zero value otherwise.

### GetGeneFlankConfigOk

`func (o *V2ProkaryoteGeneRequest) GetGeneFlankConfigOk() (*V2ProkaryoteGeneRequestGeneFlankConfig, bool)`

GetGeneFlankConfigOk returns a tuple with the GeneFlankConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneFlankConfig

`func (o *V2ProkaryoteGeneRequest) SetGeneFlankConfig(v V2ProkaryoteGeneRequestGeneFlankConfig)`

SetGeneFlankConfig sets GeneFlankConfig field to given value.

### HasGeneFlankConfig

`func (o *V2ProkaryoteGeneRequest) HasGeneFlankConfig() bool`

HasGeneFlankConfig returns a boolean if a field has been set.

### GetTaxon

`func (o *V2ProkaryoteGeneRequest) GetTaxon() string`

GetTaxon returns the Taxon field if non-nil, zero value otherwise.

### GetTaxonOk

`func (o *V2ProkaryoteGeneRequest) GetTaxonOk() (*string, bool)`

GetTaxonOk returns a tuple with the Taxon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxon

`func (o *V2ProkaryoteGeneRequest) SetTaxon(v string)`

SetTaxon sets Taxon field to given value.

### HasTaxon

`func (o *V2ProkaryoteGeneRequest) HasTaxon() bool`

HasTaxon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


