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

// V2DownloadSummaryFileSummary struct for V2DownloadSummaryFileSummary
type V2DownloadSummaryFileSummary struct {
	FileCount *int32 `json:"file_count,omitempty"`
	SizeMb *float32 `json:"size_mb,omitempty"`
}

// NewV2DownloadSummaryFileSummary instantiates a new V2DownloadSummaryFileSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2DownloadSummaryFileSummary() *V2DownloadSummaryFileSummary {
	this := V2DownloadSummaryFileSummary{}
	return &this
}

// NewV2DownloadSummaryFileSummaryWithDefaults instantiates a new V2DownloadSummaryFileSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2DownloadSummaryFileSummaryWithDefaults() *V2DownloadSummaryFileSummary {
	this := V2DownloadSummaryFileSummary{}
	return &this
}

// GetFileCount returns the FileCount field value if set, zero value otherwise.
func (o *V2DownloadSummaryFileSummary) GetFileCount() int32 {
	if o == nil || o.FileCount == nil {
		var ret int32
		return ret
	}
	return *o.FileCount
}

// GetFileCountOk returns a tuple with the FileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2DownloadSummaryFileSummary) GetFileCountOk() (*int32, bool) {
	if o == nil || o.FileCount == nil {
		return nil, false
	}
	return o.FileCount, true
}

// HasFileCount returns a boolean if a field has been set.
func (o *V2DownloadSummaryFileSummary) HasFileCount() bool {
	if o != nil && o.FileCount != nil {
		return true
	}

	return false
}

// SetFileCount gets a reference to the given int32 and assigns it to the FileCount field.
func (o *V2DownloadSummaryFileSummary) SetFileCount(v int32) {
	o.FileCount = &v
}

// GetSizeMb returns the SizeMb field value if set, zero value otherwise.
func (o *V2DownloadSummaryFileSummary) GetSizeMb() float32 {
	if o == nil || o.SizeMb == nil {
		var ret float32
		return ret
	}
	return *o.SizeMb
}

// GetSizeMbOk returns a tuple with the SizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2DownloadSummaryFileSummary) GetSizeMbOk() (*float32, bool) {
	if o == nil || o.SizeMb == nil {
		return nil, false
	}
	return o.SizeMb, true
}

// HasSizeMb returns a boolean if a field has been set.
func (o *V2DownloadSummaryFileSummary) HasSizeMb() bool {
	if o != nil && o.SizeMb != nil {
		return true
	}

	return false
}

// SetSizeMb gets a reference to the given float32 and assigns it to the SizeMb field.
func (o *V2DownloadSummaryFileSummary) SetSizeMb(v float32) {
	o.SizeMb = &v
}

func (o V2DownloadSummaryFileSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FileCount != nil {
		toSerialize["file_count"] = o.FileCount
	}
	if o.SizeMb != nil {
		toSerialize["size_mb"] = o.SizeMb
	}
	return json.Marshal(toSerialize)
}

type NullableV2DownloadSummaryFileSummary struct {
	value *V2DownloadSummaryFileSummary
	isSet bool
}

func (v NullableV2DownloadSummaryFileSummary) Get() *V2DownloadSummaryFileSummary {
	return v.value
}

func (v *NullableV2DownloadSummaryFileSummary) Set(val *V2DownloadSummaryFileSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableV2DownloadSummaryFileSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableV2DownloadSummaryFileSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2DownloadSummaryFileSummary(val *V2DownloadSummaryFileSummary) *NullableV2DownloadSummaryFileSummary {
	return &NullableV2DownloadSummaryFileSummary{value: val, isSet: true}
}

func (v NullableV2DownloadSummaryFileSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2DownloadSummaryFileSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


