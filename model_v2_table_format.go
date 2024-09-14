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

// V2TableFormat the model 'V2TableFormat'
type V2TableFormat string

// List of v2TableFormat
const (
	V2TABLEFORMAT_TSV V2TableFormat = "tsv"
	V2TABLEFORMAT_CSV V2TableFormat = "csv"
	V2TABLEFORMAT_JSONL V2TableFormat = "jsonl"
)

// All allowed values of V2TableFormat enum
var AllowedV2TableFormatEnumValues = []V2TableFormat{
	"tsv",
	"csv",
	"jsonl",
}

func (v *V2TableFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V2TableFormat(value)
	for _, existing := range AllowedV2TableFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V2TableFormat", value)
}

// NewV2TableFormatFromValue returns a pointer to a valid V2TableFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV2TableFormatFromValue(v string) (*V2TableFormat, error) {
	ev := V2TableFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V2TableFormat: valid values are %v", v, AllowedV2TableFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V2TableFormat) IsValid() bool {
	for _, existing := range AllowedV2TableFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v2TableFormat value
func (v V2TableFormat) Ptr() *V2TableFormat {
	return &v
}

type NullableV2TableFormat struct {
	value *V2TableFormat
	isSet bool
}

func (v NullableV2TableFormat) Get() *V2TableFormat {
	return v.value
}

func (v *NullableV2TableFormat) Set(val *V2TableFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TableFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TableFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TableFormat(val *V2TableFormat) *NullableV2TableFormat {
	return &NullableV2TableFormat{value: val, isSet: true}
}

func (v NullableV2TableFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TableFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

