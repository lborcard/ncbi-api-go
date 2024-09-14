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

// V2reportsTranscriptTranscriptType the model 'V2reportsTranscriptTranscriptType'
type V2reportsTranscriptTranscriptType string

// List of v2reportsTranscriptTranscriptType
const (
	V2REPORTSTRANSCRIPTTRANSCRIPTTYPE_UNKNOWN V2reportsTranscriptTranscriptType = "UNKNOWN"
	V2REPORTSTRANSCRIPTTRANSCRIPTTYPE_PROTEIN_CODING V2reportsTranscriptTranscriptType = "PROTEIN_CODING"
	V2REPORTSTRANSCRIPTTRANSCRIPTTYPE_NON_CODING V2reportsTranscriptTranscriptType = "NON_CODING"
	V2REPORTSTRANSCRIPTTRANSCRIPTTYPE_PROTEIN_CODING_MODEL V2reportsTranscriptTranscriptType = "PROTEIN_CODING_MODEL"
	V2REPORTSTRANSCRIPTTRANSCRIPTTYPE_NON_CODING_MODEL V2reportsTranscriptTranscriptType = "NON_CODING_MODEL"
)

// All allowed values of V2reportsTranscriptTranscriptType enum
var AllowedV2reportsTranscriptTranscriptTypeEnumValues = []V2reportsTranscriptTranscriptType{
	"UNKNOWN",
	"PROTEIN_CODING",
	"NON_CODING",
	"PROTEIN_CODING_MODEL",
	"NON_CODING_MODEL",
}

func (v *V2reportsTranscriptTranscriptType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V2reportsTranscriptTranscriptType(value)
	for _, existing := range AllowedV2reportsTranscriptTranscriptTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V2reportsTranscriptTranscriptType", value)
}

// NewV2reportsTranscriptTranscriptTypeFromValue returns a pointer to a valid V2reportsTranscriptTranscriptType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV2reportsTranscriptTranscriptTypeFromValue(v string) (*V2reportsTranscriptTranscriptType, error) {
	ev := V2reportsTranscriptTranscriptType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V2reportsTranscriptTranscriptType: valid values are %v", v, AllowedV2reportsTranscriptTranscriptTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V2reportsTranscriptTranscriptType) IsValid() bool {
	for _, existing := range AllowedV2reportsTranscriptTranscriptTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v2reportsTranscriptTranscriptType value
func (v V2reportsTranscriptTranscriptType) Ptr() *V2reportsTranscriptTranscriptType {
	return &v
}

type NullableV2reportsTranscriptTranscriptType struct {
	value *V2reportsTranscriptTranscriptType
	isSet bool
}

func (v NullableV2reportsTranscriptTranscriptType) Get() *V2reportsTranscriptTranscriptType {
	return v.value
}

func (v *NullableV2reportsTranscriptTranscriptType) Set(val *V2reportsTranscriptTranscriptType) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsTranscriptTranscriptType) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsTranscriptTranscriptType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsTranscriptTranscriptType(val *V2reportsTranscriptTranscriptType) *NullableV2reportsTranscriptTranscriptType {
	return &NullableV2reportsTranscriptTranscriptType{value: val, isSet: true}
}

func (v NullableV2reportsTranscriptTranscriptType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsTranscriptTranscriptType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

