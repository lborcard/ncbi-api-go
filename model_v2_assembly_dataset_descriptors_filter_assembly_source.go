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

// V2AssemblyDatasetDescriptorsFilterAssemblySource the model 'V2AssemblyDatasetDescriptorsFilterAssemblySource'
type V2AssemblyDatasetDescriptorsFilterAssemblySource string

// List of v2AssemblyDatasetDescriptorsFilterAssemblySource
const (
	V2ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYSOURCE_ALL V2AssemblyDatasetDescriptorsFilterAssemblySource = "all"
	V2ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYSOURCE_REFSEQ V2AssemblyDatasetDescriptorsFilterAssemblySource = "refseq"
	V2ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYSOURCE_GENBANK V2AssemblyDatasetDescriptorsFilterAssemblySource = "genbank"
)

// All allowed values of V2AssemblyDatasetDescriptorsFilterAssemblySource enum
var AllowedV2AssemblyDatasetDescriptorsFilterAssemblySourceEnumValues = []V2AssemblyDatasetDescriptorsFilterAssemblySource{
	"all",
	"refseq",
	"genbank",
}

func (v *V2AssemblyDatasetDescriptorsFilterAssemblySource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V2AssemblyDatasetDescriptorsFilterAssemblySource(value)
	for _, existing := range AllowedV2AssemblyDatasetDescriptorsFilterAssemblySourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V2AssemblyDatasetDescriptorsFilterAssemblySource", value)
}

// NewV2AssemblyDatasetDescriptorsFilterAssemblySourceFromValue returns a pointer to a valid V2AssemblyDatasetDescriptorsFilterAssemblySource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV2AssemblyDatasetDescriptorsFilterAssemblySourceFromValue(v string) (*V2AssemblyDatasetDescriptorsFilterAssemblySource, error) {
	ev := V2AssemblyDatasetDescriptorsFilterAssemblySource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V2AssemblyDatasetDescriptorsFilterAssemblySource: valid values are %v", v, AllowedV2AssemblyDatasetDescriptorsFilterAssemblySourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V2AssemblyDatasetDescriptorsFilterAssemblySource) IsValid() bool {
	for _, existing := range AllowedV2AssemblyDatasetDescriptorsFilterAssemblySourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v2AssemblyDatasetDescriptorsFilterAssemblySource value
func (v V2AssemblyDatasetDescriptorsFilterAssemblySource) Ptr() *V2AssemblyDatasetDescriptorsFilterAssemblySource {
	return &v
}

type NullableV2AssemblyDatasetDescriptorsFilterAssemblySource struct {
	value *V2AssemblyDatasetDescriptorsFilterAssemblySource
	isSet bool
}

func (v NullableV2AssemblyDatasetDescriptorsFilterAssemblySource) Get() *V2AssemblyDatasetDescriptorsFilterAssemblySource {
	return v.value
}

func (v *NullableV2AssemblyDatasetDescriptorsFilterAssemblySource) Set(val *V2AssemblyDatasetDescriptorsFilterAssemblySource) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssemblyDatasetDescriptorsFilterAssemblySource) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssemblyDatasetDescriptorsFilterAssemblySource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssemblyDatasetDescriptorsFilterAssemblySource(val *V2AssemblyDatasetDescriptorsFilterAssemblySource) *NullableV2AssemblyDatasetDescriptorsFilterAssemblySource {
	return &NullableV2AssemblyDatasetDescriptorsFilterAssemblySource{value: val, isSet: true}
}

func (v NullableV2AssemblyDatasetDescriptorsFilterAssemblySource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssemblyDatasetDescriptorsFilterAssemblySource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

