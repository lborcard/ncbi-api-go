/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 2 API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated zip archive](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v2alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ncbi-openapi-v2_goland

import (
	"encoding/json"
	"fmt"
)

// V2TaxonomyMetadataRequestTableFormat the model 'V2TaxonomyMetadataRequestTableFormat'
type V2TaxonomyMetadataRequestTableFormat string

// List of v2TaxonomyMetadataRequestTableFormat
const (
	V2TAXONOMYMETADATAREQUESTTABLEFORMAT_SUMMARY V2TaxonomyMetadataRequestTableFormat = "SUMMARY"
)

// All allowed values of V2TaxonomyMetadataRequestTableFormat enum
var AllowedV2TaxonomyMetadataRequestTableFormatEnumValues = []V2TaxonomyMetadataRequestTableFormat{
	"SUMMARY",
}

func (v *V2TaxonomyMetadataRequestTableFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V2TaxonomyMetadataRequestTableFormat(value)
	for _, existing := range AllowedV2TaxonomyMetadataRequestTableFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V2TaxonomyMetadataRequestTableFormat", value)
}

// NewV2TaxonomyMetadataRequestTableFormatFromValue returns a pointer to a valid V2TaxonomyMetadataRequestTableFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV2TaxonomyMetadataRequestTableFormatFromValue(v string) (*V2TaxonomyMetadataRequestTableFormat, error) {
	ev := V2TaxonomyMetadataRequestTableFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V2TaxonomyMetadataRequestTableFormat: valid values are %v", v, AllowedV2TaxonomyMetadataRequestTableFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V2TaxonomyMetadataRequestTableFormat) IsValid() bool {
	for _, existing := range AllowedV2TaxonomyMetadataRequestTableFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v2TaxonomyMetadataRequestTableFormat value
func (v V2TaxonomyMetadataRequestTableFormat) Ptr() *V2TaxonomyMetadataRequestTableFormat {
	return &v
}

type NullableV2TaxonomyMetadataRequestTableFormat struct {
	value *V2TaxonomyMetadataRequestTableFormat
	isSet bool
}

func (v NullableV2TaxonomyMetadataRequestTableFormat) Get() *V2TaxonomyMetadataRequestTableFormat {
	return v.value
}

func (v *NullableV2TaxonomyMetadataRequestTableFormat) Set(val *V2TaxonomyMetadataRequestTableFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TaxonomyMetadataRequestTableFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TaxonomyMetadataRequestTableFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TaxonomyMetadataRequestTableFormat(val *V2TaxonomyMetadataRequestTableFormat) *NullableV2TaxonomyMetadataRequestTableFormat {
	return &NullableV2TaxonomyMetadataRequestTableFormat{value: val, isSet: true}
}

func (v NullableV2TaxonomyMetadataRequestTableFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TaxonomyMetadataRequestTableFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

