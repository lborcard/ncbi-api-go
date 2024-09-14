# V2TaxonomyImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taxon** | Pointer to **string** |  | [optional] 
**ImageSize** | Pointer to [**V2ImageSize**](V2ImageSize.md) |  | [optional] [default to V2IMAGESIZE_UNSPECIFIED]

## Methods

### NewV2TaxonomyImageRequest

`func NewV2TaxonomyImageRequest() *V2TaxonomyImageRequest`

NewV2TaxonomyImageRequest instantiates a new V2TaxonomyImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TaxonomyImageRequestWithDefaults

`func NewV2TaxonomyImageRequestWithDefaults() *V2TaxonomyImageRequest`

NewV2TaxonomyImageRequestWithDefaults instantiates a new V2TaxonomyImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxon

`func (o *V2TaxonomyImageRequest) GetTaxon() string`

GetTaxon returns the Taxon field if non-nil, zero value otherwise.

### GetTaxonOk

`func (o *V2TaxonomyImageRequest) GetTaxonOk() (*string, bool)`

GetTaxonOk returns a tuple with the Taxon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxon

`func (o *V2TaxonomyImageRequest) SetTaxon(v string)`

SetTaxon sets Taxon field to given value.

### HasTaxon

`func (o *V2TaxonomyImageRequest) HasTaxon() bool`

HasTaxon returns a boolean if a field has been set.

### GetImageSize

`func (o *V2TaxonomyImageRequest) GetImageSize() V2ImageSize`

GetImageSize returns the ImageSize field if non-nil, zero value otherwise.

### GetImageSizeOk

`func (o *V2TaxonomyImageRequest) GetImageSizeOk() (*V2ImageSize, bool)`

GetImageSizeOk returns a tuple with the ImageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSize

`func (o *V2TaxonomyImageRequest) SetImageSize(v V2ImageSize)`

SetImageSize sets ImageSize field to given value.

### HasImageSize

`func (o *V2TaxonomyImageRequest) HasImageSize() bool`

HasImageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


