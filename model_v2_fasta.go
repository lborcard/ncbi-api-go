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

// V2Fasta the model 'V2Fasta'
type V2Fasta string

// List of v2Fasta
const (
	V2FASTA_UNSPECIFIED V2Fasta = "FASTA_UNSPECIFIED"
	V2FASTA_GENE V2Fasta = "FASTA_GENE"
	V2FASTA_RNA V2Fasta = "FASTA_RNA"
	V2FASTA_PROTEIN V2Fasta = "FASTA_PROTEIN"
	V2FASTA_GENE_FLANK V2Fasta = "FASTA_GENE_FLANK"
	V2FASTA_CDS V2Fasta = "FASTA_CDS"
	V2FASTA__5_P_UTR V2Fasta = "FASTA_5P_UTR"
	V2FASTA__3_P_UTR V2Fasta = "FASTA_3P_UTR"
)

// All allowed values of V2Fasta enum
var AllowedV2FastaEnumValues = []V2Fasta{
	"FASTA_UNSPECIFIED",
	"FASTA_GENE",
	"FASTA_RNA",
	"FASTA_PROTEIN",
	"FASTA_GENE_FLANK",
	"FASTA_CDS",
	"FASTA_5P_UTR",
	"FASTA_3P_UTR",
}

func (v *V2Fasta) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V2Fasta(value)
	for _, existing := range AllowedV2FastaEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V2Fasta", value)
}

// NewV2FastaFromValue returns a pointer to a valid V2Fasta
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV2FastaFromValue(v string) (*V2Fasta, error) {
	ev := V2Fasta(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V2Fasta: valid values are %v", v, AllowedV2FastaEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V2Fasta) IsValid() bool {
	for _, existing := range AllowedV2FastaEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v2Fasta value
func (v V2Fasta) Ptr() *V2Fasta {
	return &v
}

type NullableV2Fasta struct {
	value *V2Fasta
	isSet bool
}

func (v NullableV2Fasta) Get() *V2Fasta {
	return v.value
}

func (v *NullableV2Fasta) Set(val *V2Fasta) {
	v.value = val
	v.isSet = true
}

func (v NullableV2Fasta) IsSet() bool {
	return v.isSet
}

func (v *NullableV2Fasta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2Fasta(val *V2Fasta) *NullableV2Fasta {
	return &NullableV2Fasta{value: val, isSet: true}
}

func (v NullableV2Fasta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2Fasta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

