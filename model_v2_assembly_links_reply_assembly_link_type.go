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

// V2AssemblyLinksReplyAssemblyLinkType the model 'V2AssemblyLinksReplyAssemblyLinkType'
type V2AssemblyLinksReplyAssemblyLinkType string

// List of v2AssemblyLinksReplyAssemblyLinkType
const (
	V2ASSEMBLYLINKSREPLYASSEMBLYLINKTYPE_GDV_LINK V2AssemblyLinksReplyAssemblyLinkType = "GDV_LINK"
	V2ASSEMBLYLINKSREPLYASSEMBLYLINKTYPE_FTP_LINK V2AssemblyLinksReplyAssemblyLinkType = "FTP_LINK"
	V2ASSEMBLYLINKSREPLYASSEMBLYLINKTYPE_ASSEMBLY_PUBMED V2AssemblyLinksReplyAssemblyLinkType = "ASSEMBLY_PUBMED"
	V2ASSEMBLYLINKSREPLYASSEMBLYLINKTYPE_BLAST_LINK V2AssemblyLinksReplyAssemblyLinkType = "BLAST_LINK"
	V2ASSEMBLYLINKSREPLYASSEMBLYLINKTYPE_ASSEMBLY_NUCCORE_REFSEQ V2AssemblyLinksReplyAssemblyLinkType = "ASSEMBLY_NUCCORE_REFSEQ"
	V2ASSEMBLYLINKSREPLYASSEMBLYLINKTYPE_ASSEMBLY_NUCCORE_GENBANK V2AssemblyLinksReplyAssemblyLinkType = "ASSEMBLY_NUCCORE_GENBANK"
)

// All allowed values of V2AssemblyLinksReplyAssemblyLinkType enum
var AllowedV2AssemblyLinksReplyAssemblyLinkTypeEnumValues = []V2AssemblyLinksReplyAssemblyLinkType{
	"GDV_LINK",
	"FTP_LINK",
	"ASSEMBLY_PUBMED",
	"BLAST_LINK",
	"ASSEMBLY_NUCCORE_REFSEQ",
	"ASSEMBLY_NUCCORE_GENBANK",
}

func (v *V2AssemblyLinksReplyAssemblyLinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V2AssemblyLinksReplyAssemblyLinkType(value)
	for _, existing := range AllowedV2AssemblyLinksReplyAssemblyLinkTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V2AssemblyLinksReplyAssemblyLinkType", value)
}

// NewV2AssemblyLinksReplyAssemblyLinkTypeFromValue returns a pointer to a valid V2AssemblyLinksReplyAssemblyLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV2AssemblyLinksReplyAssemblyLinkTypeFromValue(v string) (*V2AssemblyLinksReplyAssemblyLinkType, error) {
	ev := V2AssemblyLinksReplyAssemblyLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V2AssemblyLinksReplyAssemblyLinkType: valid values are %v", v, AllowedV2AssemblyLinksReplyAssemblyLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V2AssemblyLinksReplyAssemblyLinkType) IsValid() bool {
	for _, existing := range AllowedV2AssemblyLinksReplyAssemblyLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v2AssemblyLinksReplyAssemblyLinkType value
func (v V2AssemblyLinksReplyAssemblyLinkType) Ptr() *V2AssemblyLinksReplyAssemblyLinkType {
	return &v
}

type NullableV2AssemblyLinksReplyAssemblyLinkType struct {
	value *V2AssemblyLinksReplyAssemblyLinkType
	isSet bool
}

func (v NullableV2AssemblyLinksReplyAssemblyLinkType) Get() *V2AssemblyLinksReplyAssemblyLinkType {
	return v.value
}

func (v *NullableV2AssemblyLinksReplyAssemblyLinkType) Set(val *V2AssemblyLinksReplyAssemblyLinkType) {
	v.value = val
	v.isSet = true
}

func (v NullableV2AssemblyLinksReplyAssemblyLinkType) IsSet() bool {
	return v.isSet
}

func (v *NullableV2AssemblyLinksReplyAssemblyLinkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2AssemblyLinksReplyAssemblyLinkType(val *V2AssemblyLinksReplyAssemblyLinkType) *NullableV2AssemblyLinksReplyAssemblyLinkType {
	return &NullableV2AssemblyLinksReplyAssemblyLinkType{value: val, isSet: true}
}

func (v NullableV2AssemblyLinksReplyAssemblyLinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2AssemblyLinksReplyAssemblyLinkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

