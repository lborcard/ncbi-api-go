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

// V2reportsRankType the model 'V2reportsRankType'
type V2reportsRankType string

// List of v2reportsRankType
const (
	V2REPORTSRANKTYPE_NO_RANK V2reportsRankType = "NO_RANK"
	V2REPORTSRANKTYPE_SUPERKINGDOM V2reportsRankType = "SUPERKINGDOM"
	V2REPORTSRANKTYPE_KINGDOM V2reportsRankType = "KINGDOM"
	V2REPORTSRANKTYPE_SUBKINGDOM V2reportsRankType = "SUBKINGDOM"
	V2REPORTSRANKTYPE_SUPERPHYLUM V2reportsRankType = "SUPERPHYLUM"
	V2REPORTSRANKTYPE_SUBPHYLUM V2reportsRankType = "SUBPHYLUM"
	V2REPORTSRANKTYPE_PHYLUM V2reportsRankType = "PHYLUM"
	V2REPORTSRANKTYPE_CLADE V2reportsRankType = "CLADE"
	V2REPORTSRANKTYPE_SUPERCLASS V2reportsRankType = "SUPERCLASS"
	V2REPORTSRANKTYPE_CLASS V2reportsRankType = "CLASS"
	V2REPORTSRANKTYPE_SUBCLASS V2reportsRankType = "SUBCLASS"
	V2REPORTSRANKTYPE_INFRACLASS V2reportsRankType = "INFRACLASS"
	V2REPORTSRANKTYPE_COHORT V2reportsRankType = "COHORT"
	V2REPORTSRANKTYPE_SUBCOHORT V2reportsRankType = "SUBCOHORT"
	V2REPORTSRANKTYPE_SUPERORDER V2reportsRankType = "SUPERORDER"
	V2REPORTSRANKTYPE_ORDER V2reportsRankType = "ORDER"
	V2REPORTSRANKTYPE_SUBORDER V2reportsRankType = "SUBORDER"
	V2REPORTSRANKTYPE_INFRAORDER V2reportsRankType = "INFRAORDER"
	V2REPORTSRANKTYPE_PARVORDER V2reportsRankType = "PARVORDER"
	V2REPORTSRANKTYPE_SUPERFAMILY V2reportsRankType = "SUPERFAMILY"
	V2REPORTSRANKTYPE_FAMILY V2reportsRankType = "FAMILY"
	V2REPORTSRANKTYPE_SUBFAMILY V2reportsRankType = "SUBFAMILY"
	V2REPORTSRANKTYPE_GENUS V2reportsRankType = "GENUS"
	V2REPORTSRANKTYPE_SUBGENUS V2reportsRankType = "SUBGENUS"
	V2REPORTSRANKTYPE_SPECIES_GROUP V2reportsRankType = "SPECIES_GROUP"
	V2REPORTSRANKTYPE_SPECIES_SUBGROUP V2reportsRankType = "SPECIES_SUBGROUP"
	V2REPORTSRANKTYPE_SPECIES V2reportsRankType = "SPECIES"
	V2REPORTSRANKTYPE_SUBSPECIES V2reportsRankType = "SUBSPECIES"
	V2REPORTSRANKTYPE_TRIBE V2reportsRankType = "TRIBE"
	V2REPORTSRANKTYPE_SUBTRIBE V2reportsRankType = "SUBTRIBE"
	V2REPORTSRANKTYPE_FORMA V2reportsRankType = "FORMA"
	V2REPORTSRANKTYPE_VARIETAS V2reportsRankType = "VARIETAS"
	V2REPORTSRANKTYPE_STRAIN V2reportsRankType = "STRAIN"
	V2REPORTSRANKTYPE_SECTION V2reportsRankType = "SECTION"
	V2REPORTSRANKTYPE_SUBSECTION V2reportsRankType = "SUBSECTION"
	V2REPORTSRANKTYPE_PATHOGROUP V2reportsRankType = "PATHOGROUP"
	V2REPORTSRANKTYPE_SUBVARIETY V2reportsRankType = "SUBVARIETY"
	V2REPORTSRANKTYPE_GENOTYPE V2reportsRankType = "GENOTYPE"
	V2REPORTSRANKTYPE_SEROTYPE V2reportsRankType = "SEROTYPE"
	V2REPORTSRANKTYPE_ISOLATE V2reportsRankType = "ISOLATE"
	V2REPORTSRANKTYPE_MORPH V2reportsRankType = "MORPH"
	V2REPORTSRANKTYPE_SERIES V2reportsRankType = "SERIES"
	V2REPORTSRANKTYPE_FORMA_SPECIALIS V2reportsRankType = "FORMA_SPECIALIS"
	V2REPORTSRANKTYPE_SEROGROUP V2reportsRankType = "SEROGROUP"
	V2REPORTSRANKTYPE_BIOTYPE V2reportsRankType = "BIOTYPE"
)

// All allowed values of V2reportsRankType enum
var AllowedV2reportsRankTypeEnumValues = []V2reportsRankType{
	"NO_RANK",
	"SUPERKINGDOM",
	"KINGDOM",
	"SUBKINGDOM",
	"SUPERPHYLUM",
	"SUBPHYLUM",
	"PHYLUM",
	"CLADE",
	"SUPERCLASS",
	"CLASS",
	"SUBCLASS",
	"INFRACLASS",
	"COHORT",
	"SUBCOHORT",
	"SUPERORDER",
	"ORDER",
	"SUBORDER",
	"INFRAORDER",
	"PARVORDER",
	"SUPERFAMILY",
	"FAMILY",
	"SUBFAMILY",
	"GENUS",
	"SUBGENUS",
	"SPECIES_GROUP",
	"SPECIES_SUBGROUP",
	"SPECIES",
	"SUBSPECIES",
	"TRIBE",
	"SUBTRIBE",
	"FORMA",
	"VARIETAS",
	"STRAIN",
	"SECTION",
	"SUBSECTION",
	"PATHOGROUP",
	"SUBVARIETY",
	"GENOTYPE",
	"SEROTYPE",
	"ISOLATE",
	"MORPH",
	"SERIES",
	"FORMA_SPECIALIS",
	"SEROGROUP",
	"BIOTYPE",
}

func (v *V2reportsRankType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V2reportsRankType(value)
	for _, existing := range AllowedV2reportsRankTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V2reportsRankType", value)
}

// NewV2reportsRankTypeFromValue returns a pointer to a valid V2reportsRankType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV2reportsRankTypeFromValue(v string) (*V2reportsRankType, error) {
	ev := V2reportsRankType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V2reportsRankType: valid values are %v", v, AllowedV2reportsRankTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V2reportsRankType) IsValid() bool {
	for _, existing := range AllowedV2reportsRankTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v2reportsRankType value
func (v V2reportsRankType) Ptr() *V2reportsRankType {
	return &v
}

type NullableV2reportsRankType struct {
	value *V2reportsRankType
	isSet bool
}

func (v NullableV2reportsRankType) Get() *V2reportsRankType {
	return v.value
}

func (v *NullableV2reportsRankType) Set(val *V2reportsRankType) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsRankType) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsRankType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsRankType(val *V2reportsRankType) *NullableV2reportsRankType {
	return &NullableV2reportsRankType{value: val, isSet: true}
}

func (v NullableV2reportsRankType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsRankType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

