/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 2 API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated zip archive](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v2alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ncbi-openapi-v2_goland

import (
	"encoding/json"
)

// V2TaxonomyImageRequest struct for V2TaxonomyImageRequest
type V2TaxonomyImageRequest struct {
	Taxon *string `json:"taxon,omitempty"`
	ImageSize *V2ImageSize `json:"image_size,omitempty"`
}

// NewV2TaxonomyImageRequest instantiates a new V2TaxonomyImageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TaxonomyImageRequest() *V2TaxonomyImageRequest {
	this := V2TaxonomyImageRequest{}
	var imageSize V2ImageSize = V2IMAGESIZE_UNSPECIFIED
	this.ImageSize = &imageSize
	return &this
}

// NewV2TaxonomyImageRequestWithDefaults instantiates a new V2TaxonomyImageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TaxonomyImageRequestWithDefaults() *V2TaxonomyImageRequest {
	this := V2TaxonomyImageRequest{}
	var imageSize V2ImageSize = V2IMAGESIZE_UNSPECIFIED
	this.ImageSize = &imageSize
	return &this
}

// GetTaxon returns the Taxon field value if set, zero value otherwise.
func (o *V2TaxonomyImageRequest) GetTaxon() string {
	if o == nil || o.Taxon == nil {
		var ret string
		return ret
	}
	return *o.Taxon
}

// GetTaxonOk returns a tuple with the Taxon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyImageRequest) GetTaxonOk() (*string, bool) {
	if o == nil || o.Taxon == nil {
		return nil, false
	}
	return o.Taxon, true
}

// HasTaxon returns a boolean if a field has been set.
func (o *V2TaxonomyImageRequest) HasTaxon() bool {
	if o != nil && o.Taxon != nil {
		return true
	}

	return false
}

// SetTaxon gets a reference to the given string and assigns it to the Taxon field.
func (o *V2TaxonomyImageRequest) SetTaxon(v string) {
	o.Taxon = &v
}

// GetImageSize returns the ImageSize field value if set, zero value otherwise.
func (o *V2TaxonomyImageRequest) GetImageSize() V2ImageSize {
	if o == nil || o.ImageSize == nil {
		var ret V2ImageSize
		return ret
	}
	return *o.ImageSize
}

// GetImageSizeOk returns a tuple with the ImageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyImageRequest) GetImageSizeOk() (*V2ImageSize, bool) {
	if o == nil || o.ImageSize == nil {
		return nil, false
	}
	return o.ImageSize, true
}

// HasImageSize returns a boolean if a field has been set.
func (o *V2TaxonomyImageRequest) HasImageSize() bool {
	if o != nil && o.ImageSize != nil {
		return true
	}

	return false
}

// SetImageSize gets a reference to the given V2ImageSize and assigns it to the ImageSize field.
func (o *V2TaxonomyImageRequest) SetImageSize(v V2ImageSize) {
	o.ImageSize = &v
}

func (o V2TaxonomyImageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Taxon != nil {
		toSerialize["taxon"] = o.Taxon
	}
	if o.ImageSize != nil {
		toSerialize["image_size"] = o.ImageSize
	}
	return json.Marshal(toSerialize)
}

type NullableV2TaxonomyImageRequest struct {
	value *V2TaxonomyImageRequest
	isSet bool
}

func (v NullableV2TaxonomyImageRequest) Get() *V2TaxonomyImageRequest {
	return v.value
}

func (v *NullableV2TaxonomyImageRequest) Set(val *V2TaxonomyImageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TaxonomyImageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TaxonomyImageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TaxonomyImageRequest(val *V2TaxonomyImageRequest) *NullableV2TaxonomyImageRequest {
	return &NullableV2TaxonomyImageRequest{value: val, isSet: true}
}

func (v NullableV2TaxonomyImageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TaxonomyImageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


