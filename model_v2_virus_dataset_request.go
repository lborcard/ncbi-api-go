/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 2 API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated zip archive](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v2alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// V2VirusDatasetRequest struct for V2VirusDatasetRequest
type V2VirusDatasetRequest struct {
	Accessions []string `json:"accessions,omitempty"`
	Taxon *string `json:"taxon,omitempty"`
	RefseqOnly *bool `json:"refseq_only,omitempty"`
	AnnotatedOnly *bool `json:"annotated_only,omitempty"`
	ReleasedSince *time.Time `json:"released_since,omitempty"`
	UpdatedSince *time.Time `json:"updated_since,omitempty"`
	Host *string `json:"host,omitempty"`
	PangolinClassification *string `json:"pangolin_classification,omitempty"`
	GeoLocation *string `json:"geo_location,omitempty"`
	CompleteOnly *bool `json:"complete_only,omitempty"`
	TableFields []V2VirusTableField `json:"table_fields,omitempty"`
	IncludeSequence []V2ViralSequenceType `json:"include_sequence,omitempty"`
	AuxReport []V2VirusDatasetReportType `json:"aux_report,omitempty"`
	Format *V2TableFormat `json:"format,omitempty"`
	UsePsg *bool `json:"use_psg,omitempty"`
}

// NewV2VirusDatasetRequest instantiates a new V2VirusDatasetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2VirusDatasetRequest() *V2VirusDatasetRequest {
	this := V2VirusDatasetRequest{}
	var format V2TableFormat = V2TABLEFORMAT_TSV
	this.Format = &format
	return &this
}

// NewV2VirusDatasetRequestWithDefaults instantiates a new V2VirusDatasetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2VirusDatasetRequestWithDefaults() *V2VirusDatasetRequest {
	this := V2VirusDatasetRequest{}
	var format V2TableFormat = V2TABLEFORMAT_TSV
	this.Format = &format
	return &this
}

// GetAccessions returns the Accessions field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetAccessions() []string {
	if o == nil || o.Accessions == nil {
		var ret []string
		return ret
	}
	return o.Accessions
}

// GetAccessionsOk returns a tuple with the Accessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetAccessionsOk() ([]string, bool) {
	if o == nil || o.Accessions == nil {
		return nil, false
	}
	return o.Accessions, true
}

// HasAccessions returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasAccessions() bool {
	if o != nil && o.Accessions != nil {
		return true
	}

	return false
}

// SetAccessions gets a reference to the given []string and assigns it to the Accessions field.
func (o *V2VirusDatasetRequest) SetAccessions(v []string) {
	o.Accessions = v
}

// GetTaxon returns the Taxon field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetTaxon() string {
	if o == nil || o.Taxon == nil {
		var ret string
		return ret
	}
	return *o.Taxon
}

// GetTaxonOk returns a tuple with the Taxon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetTaxonOk() (*string, bool) {
	if o == nil || o.Taxon == nil {
		return nil, false
	}
	return o.Taxon, true
}

// HasTaxon returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasTaxon() bool {
	if o != nil && o.Taxon != nil {
		return true
	}

	return false
}

// SetTaxon gets a reference to the given string and assigns it to the Taxon field.
func (o *V2VirusDatasetRequest) SetTaxon(v string) {
	o.Taxon = &v
}

// GetRefseqOnly returns the RefseqOnly field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetRefseqOnly() bool {
	if o == nil || o.RefseqOnly == nil {
		var ret bool
		return ret
	}
	return *o.RefseqOnly
}

// GetRefseqOnlyOk returns a tuple with the RefseqOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetRefseqOnlyOk() (*bool, bool) {
	if o == nil || o.RefseqOnly == nil {
		return nil, false
	}
	return o.RefseqOnly, true
}

// HasRefseqOnly returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasRefseqOnly() bool {
	if o != nil && o.RefseqOnly != nil {
		return true
	}

	return false
}

// SetRefseqOnly gets a reference to the given bool and assigns it to the RefseqOnly field.
func (o *V2VirusDatasetRequest) SetRefseqOnly(v bool) {
	o.RefseqOnly = &v
}

// GetAnnotatedOnly returns the AnnotatedOnly field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetAnnotatedOnly() bool {
	if o == nil || o.AnnotatedOnly == nil {
		var ret bool
		return ret
	}
	return *o.AnnotatedOnly
}

// GetAnnotatedOnlyOk returns a tuple with the AnnotatedOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetAnnotatedOnlyOk() (*bool, bool) {
	if o == nil || o.AnnotatedOnly == nil {
		return nil, false
	}
	return o.AnnotatedOnly, true
}

// HasAnnotatedOnly returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasAnnotatedOnly() bool {
	if o != nil && o.AnnotatedOnly != nil {
		return true
	}

	return false
}

// SetAnnotatedOnly gets a reference to the given bool and assigns it to the AnnotatedOnly field.
func (o *V2VirusDatasetRequest) SetAnnotatedOnly(v bool) {
	o.AnnotatedOnly = &v
}

// GetReleasedSince returns the ReleasedSince field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetReleasedSince() time.Time {
	if o == nil || o.ReleasedSince == nil {
		var ret time.Time
		return ret
	}
	return *o.ReleasedSince
}

// GetReleasedSinceOk returns a tuple with the ReleasedSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetReleasedSinceOk() (*time.Time, bool) {
	if o == nil || o.ReleasedSince == nil {
		return nil, false
	}
	return o.ReleasedSince, true
}

// HasReleasedSince returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasReleasedSince() bool {
	if o != nil && o.ReleasedSince != nil {
		return true
	}

	return false
}

// SetReleasedSince gets a reference to the given time.Time and assigns it to the ReleasedSince field.
func (o *V2VirusDatasetRequest) SetReleasedSince(v time.Time) {
	o.ReleasedSince = &v
}

// GetUpdatedSince returns the UpdatedSince field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetUpdatedSince() time.Time {
	if o == nil || o.UpdatedSince == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedSince
}

// GetUpdatedSinceOk returns a tuple with the UpdatedSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetUpdatedSinceOk() (*time.Time, bool) {
	if o == nil || o.UpdatedSince == nil {
		return nil, false
	}
	return o.UpdatedSince, true
}

// HasUpdatedSince returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasUpdatedSince() bool {
	if o != nil && o.UpdatedSince != nil {
		return true
	}

	return false
}

// SetUpdatedSince gets a reference to the given time.Time and assigns it to the UpdatedSince field.
func (o *V2VirusDatasetRequest) SetUpdatedSince(v time.Time) {
	o.UpdatedSince = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *V2VirusDatasetRequest) SetHost(v string) {
	o.Host = &v
}

// GetPangolinClassification returns the PangolinClassification field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetPangolinClassification() string {
	if o == nil || o.PangolinClassification == nil {
		var ret string
		return ret
	}
	return *o.PangolinClassification
}

// GetPangolinClassificationOk returns a tuple with the PangolinClassification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetPangolinClassificationOk() (*string, bool) {
	if o == nil || o.PangolinClassification == nil {
		return nil, false
	}
	return o.PangolinClassification, true
}

// HasPangolinClassification returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasPangolinClassification() bool {
	if o != nil && o.PangolinClassification != nil {
		return true
	}

	return false
}

// SetPangolinClassification gets a reference to the given string and assigns it to the PangolinClassification field.
func (o *V2VirusDatasetRequest) SetPangolinClassification(v string) {
	o.PangolinClassification = &v
}

// GetGeoLocation returns the GeoLocation field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetGeoLocation() string {
	if o == nil || o.GeoLocation == nil {
		var ret string
		return ret
	}
	return *o.GeoLocation
}

// GetGeoLocationOk returns a tuple with the GeoLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetGeoLocationOk() (*string, bool) {
	if o == nil || o.GeoLocation == nil {
		return nil, false
	}
	return o.GeoLocation, true
}

// HasGeoLocation returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasGeoLocation() bool {
	if o != nil && o.GeoLocation != nil {
		return true
	}

	return false
}

// SetGeoLocation gets a reference to the given string and assigns it to the GeoLocation field.
func (o *V2VirusDatasetRequest) SetGeoLocation(v string) {
	o.GeoLocation = &v
}

// GetCompleteOnly returns the CompleteOnly field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetCompleteOnly() bool {
	if o == nil || o.CompleteOnly == nil {
		var ret bool
		return ret
	}
	return *o.CompleteOnly
}

// GetCompleteOnlyOk returns a tuple with the CompleteOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetCompleteOnlyOk() (*bool, bool) {
	if o == nil || o.CompleteOnly == nil {
		return nil, false
	}
	return o.CompleteOnly, true
}

// HasCompleteOnly returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasCompleteOnly() bool {
	if o != nil && o.CompleteOnly != nil {
		return true
	}

	return false
}

// SetCompleteOnly gets a reference to the given bool and assigns it to the CompleteOnly field.
func (o *V2VirusDatasetRequest) SetCompleteOnly(v bool) {
	o.CompleteOnly = &v
}

// GetTableFields returns the TableFields field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetTableFields() []V2VirusTableField {
	if o == nil || o.TableFields == nil {
		var ret []V2VirusTableField
		return ret
	}
	return o.TableFields
}

// GetTableFieldsOk returns a tuple with the TableFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetTableFieldsOk() ([]V2VirusTableField, bool) {
	if o == nil || o.TableFields == nil {
		return nil, false
	}
	return o.TableFields, true
}

// HasTableFields returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasTableFields() bool {
	if o != nil && o.TableFields != nil {
		return true
	}

	return false
}

// SetTableFields gets a reference to the given []V2VirusTableField and assigns it to the TableFields field.
func (o *V2VirusDatasetRequest) SetTableFields(v []V2VirusTableField) {
	o.TableFields = v
}

// GetIncludeSequence returns the IncludeSequence field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetIncludeSequence() []V2ViralSequenceType {
	if o == nil || o.IncludeSequence == nil {
		var ret []V2ViralSequenceType
		return ret
	}
	return o.IncludeSequence
}

// GetIncludeSequenceOk returns a tuple with the IncludeSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetIncludeSequenceOk() ([]V2ViralSequenceType, bool) {
	if o == nil || o.IncludeSequence == nil {
		return nil, false
	}
	return o.IncludeSequence, true
}

// HasIncludeSequence returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasIncludeSequence() bool {
	if o != nil && o.IncludeSequence != nil {
		return true
	}

	return false
}

// SetIncludeSequence gets a reference to the given []V2ViralSequenceType and assigns it to the IncludeSequence field.
func (o *V2VirusDatasetRequest) SetIncludeSequence(v []V2ViralSequenceType) {
	o.IncludeSequence = v
}

// GetAuxReport returns the AuxReport field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetAuxReport() []V2VirusDatasetReportType {
	if o == nil || o.AuxReport == nil {
		var ret []V2VirusDatasetReportType
		return ret
	}
	return o.AuxReport
}

// GetAuxReportOk returns a tuple with the AuxReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetAuxReportOk() ([]V2VirusDatasetReportType, bool) {
	if o == nil || o.AuxReport == nil {
		return nil, false
	}
	return o.AuxReport, true
}

// HasAuxReport returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasAuxReport() bool {
	if o != nil && o.AuxReport != nil {
		return true
	}

	return false
}

// SetAuxReport gets a reference to the given []V2VirusDatasetReportType and assigns it to the AuxReport field.
func (o *V2VirusDatasetRequest) SetAuxReport(v []V2VirusDatasetReportType) {
	o.AuxReport = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetFormat() V2TableFormat {
	if o == nil || o.Format == nil {
		var ret V2TableFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetFormatOk() (*V2TableFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given V2TableFormat and assigns it to the Format field.
func (o *V2VirusDatasetRequest) SetFormat(v V2TableFormat) {
	o.Format = &v
}

// GetUsePsg returns the UsePsg field value if set, zero value otherwise.
func (o *V2VirusDatasetRequest) GetUsePsg() bool {
	if o == nil || o.UsePsg == nil {
		var ret bool
		return ret
	}
	return *o.UsePsg
}

// GetUsePsgOk returns a tuple with the UsePsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDatasetRequest) GetUsePsgOk() (*bool, bool) {
	if o == nil || o.UsePsg == nil {
		return nil, false
	}
	return o.UsePsg, true
}

// HasUsePsg returns a boolean if a field has been set.
func (o *V2VirusDatasetRequest) HasUsePsg() bool {
	if o != nil && o.UsePsg != nil {
		return true
	}

	return false
}

// SetUsePsg gets a reference to the given bool and assigns it to the UsePsg field.
func (o *V2VirusDatasetRequest) SetUsePsg(v bool) {
	o.UsePsg = &v
}

func (o V2VirusDatasetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessions != nil {
		toSerialize["accessions"] = o.Accessions
	}
	if o.Taxon != nil {
		toSerialize["taxon"] = o.Taxon
	}
	if o.RefseqOnly != nil {
		toSerialize["refseq_only"] = o.RefseqOnly
	}
	if o.AnnotatedOnly != nil {
		toSerialize["annotated_only"] = o.AnnotatedOnly
	}
	if o.ReleasedSince != nil {
		toSerialize["released_since"] = o.ReleasedSince
	}
	if o.UpdatedSince != nil {
		toSerialize["updated_since"] = o.UpdatedSince
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.PangolinClassification != nil {
		toSerialize["pangolin_classification"] = o.PangolinClassification
	}
	if o.GeoLocation != nil {
		toSerialize["geo_location"] = o.GeoLocation
	}
	if o.CompleteOnly != nil {
		toSerialize["complete_only"] = o.CompleteOnly
	}
	if o.TableFields != nil {
		toSerialize["table_fields"] = o.TableFields
	}
	if o.IncludeSequence != nil {
		toSerialize["include_sequence"] = o.IncludeSequence
	}
	if o.AuxReport != nil {
		toSerialize["aux_report"] = o.AuxReport
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.UsePsg != nil {
		toSerialize["use_psg"] = o.UsePsg
	}
	return json.Marshal(toSerialize)
}

type NullableV2VirusDatasetRequest struct {
	value *V2VirusDatasetRequest
	isSet bool
}

func (v NullableV2VirusDatasetRequest) Get() *V2VirusDatasetRequest {
	return v.value
}

func (v *NullableV2VirusDatasetRequest) Set(val *V2VirusDatasetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2VirusDatasetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2VirusDatasetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2VirusDatasetRequest(val *V2VirusDatasetRequest) *NullableV2VirusDatasetRequest {
	return &NullableV2VirusDatasetRequest{value: val, isSet: true}
}

func (v NullableV2VirusDatasetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2VirusDatasetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


