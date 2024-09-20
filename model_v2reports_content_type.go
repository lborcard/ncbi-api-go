/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 2 API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated zip archive](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v2alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// V2reportsContentType the model 'V2reportsContentType'
type V2reportsContentType string

// List of v2reportsContentType
const (
	V2REPORTSCONTENTTYPE_COMPLETE V2reportsContentType = "COMPLETE"
	V2REPORTSCONTENTTYPE_ASSM_ACC V2reportsContentType = "ASSM_ACC"
	V2REPORTSCONTENTTYPE_PAIRED_ACC V2reportsContentType = "PAIRED_ACC"
)

// All allowed values of V2reportsContentType enum
var AllowedV2reportsContentTypeEnumValues = []V2reportsContentType{
	"COMPLETE",
	"ASSM_ACC",
	"PAIRED_ACC",
}

func (v *V2reportsContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V2reportsContentType(value)
	for _, existing := range AllowedV2reportsContentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V2reportsContentType", value)
}

// NewV2reportsContentTypeFromValue returns a pointer to a valid V2reportsContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV2reportsContentTypeFromValue(v string) (*V2reportsContentType, error) {
	ev := V2reportsContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V2reportsContentType: valid values are %v", v, AllowedV2reportsContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V2reportsContentType) IsValid() bool {
	for _, existing := range AllowedV2reportsContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v2reportsContentType value
func (v V2reportsContentType) Ptr() *V2reportsContentType {
	return &v
}

type NullableV2reportsContentType struct {
	value *V2reportsContentType
	isSet bool
}

func (v NullableV2reportsContentType) Get() *V2reportsContentType {
	return v.value
}

func (v *NullableV2reportsContentType) Set(val *V2reportsContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsContentType(val *V2reportsContentType) *NullableV2reportsContentType {
	return &NullableV2reportsContentType{value: val, isSet: true}
}

func (v NullableV2reportsContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

