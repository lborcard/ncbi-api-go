/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 2 API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated zip archive](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v2alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// V2reportsGenomeAnnotationReportPage struct for V2reportsGenomeAnnotationReportPage
type V2reportsGenomeAnnotationReportPage struct {
	Reports []V2reportsGenomeAnnotationReportMatch `json:"reports,omitempty"`
	Messages []V2reportsMessage `json:"messages,omitempty"`
	TotalCount *int32 `json:"total_count,omitempty"`
	NextPageToken *string `json:"next_page_token,omitempty"`
	ReportType *string `json:"_report_type,omitempty"`
	ReportFields []string `json:"_report_fields,omitempty"`
	FirstPage *bool `json:"_first_page,omitempty"`
	ReportFormat *string `json:"_report_format,omitempty"`
}

// NewV2reportsGenomeAnnotationReportPage instantiates a new V2reportsGenomeAnnotationReportPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsGenomeAnnotationReportPage() *V2reportsGenomeAnnotationReportPage {
	this := V2reportsGenomeAnnotationReportPage{}
	return &this
}

// NewV2reportsGenomeAnnotationReportPageWithDefaults instantiates a new V2reportsGenomeAnnotationReportPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsGenomeAnnotationReportPageWithDefaults() *V2reportsGenomeAnnotationReportPage {
	this := V2reportsGenomeAnnotationReportPage{}
	return &this
}

// GetReports returns the Reports field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotationReportPage) GetReports() []V2reportsGenomeAnnotationReportMatch {
	if o == nil || o.Reports == nil {
		var ret []V2reportsGenomeAnnotationReportMatch
		return ret
	}
	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotationReportPage) GetReportsOk() ([]V2reportsGenomeAnnotationReportMatch, bool) {
	if o == nil || o.Reports == nil {
		return nil, false
	}
	return o.Reports, true
}

// HasReports returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotationReportPage) HasReports() bool {
	if o != nil && o.Reports != nil {
		return true
	}

	return false
}

// SetReports gets a reference to the given []V2reportsGenomeAnnotationReportMatch and assigns it to the Reports field.
func (o *V2reportsGenomeAnnotationReportPage) SetReports(v []V2reportsGenomeAnnotationReportMatch) {
	o.Reports = v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotationReportPage) GetMessages() []V2reportsMessage {
	if o == nil || o.Messages == nil {
		var ret []V2reportsMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotationReportPage) GetMessagesOk() ([]V2reportsMessage, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotationReportPage) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []V2reportsMessage and assigns it to the Messages field.
func (o *V2reportsGenomeAnnotationReportPage) SetMessages(v []V2reportsMessage) {
	o.Messages = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotationReportPage) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotationReportPage) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotationReportPage) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *V2reportsGenomeAnnotationReportPage) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotationReportPage) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotationReportPage) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotationReportPage) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *V2reportsGenomeAnnotationReportPage) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotationReportPage) GetReportType() string {
	if o == nil || o.ReportType == nil {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotationReportPage) GetReportTypeOk() (*string, bool) {
	if o == nil || o.ReportType == nil {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotationReportPage) HasReportType() bool {
	if o != nil && o.ReportType != nil {
		return true
	}

	return false
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *V2reportsGenomeAnnotationReportPage) SetReportType(v string) {
	o.ReportType = &v
}

// GetReportFields returns the ReportFields field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotationReportPage) GetReportFields() []string {
	if o == nil || o.ReportFields == nil {
		var ret []string
		return ret
	}
	return o.ReportFields
}

// GetReportFieldsOk returns a tuple with the ReportFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotationReportPage) GetReportFieldsOk() ([]string, bool) {
	if o == nil || o.ReportFields == nil {
		return nil, false
	}
	return o.ReportFields, true
}

// HasReportFields returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotationReportPage) HasReportFields() bool {
	if o != nil && o.ReportFields != nil {
		return true
	}

	return false
}

// SetReportFields gets a reference to the given []string and assigns it to the ReportFields field.
func (o *V2reportsGenomeAnnotationReportPage) SetReportFields(v []string) {
	o.ReportFields = v
}

// GetFirstPage returns the FirstPage field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotationReportPage) GetFirstPage() bool {
	if o == nil || o.FirstPage == nil {
		var ret bool
		return ret
	}
	return *o.FirstPage
}

// GetFirstPageOk returns a tuple with the FirstPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotationReportPage) GetFirstPageOk() (*bool, bool) {
	if o == nil || o.FirstPage == nil {
		return nil, false
	}
	return o.FirstPage, true
}

// HasFirstPage returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotationReportPage) HasFirstPage() bool {
	if o != nil && o.FirstPage != nil {
		return true
	}

	return false
}

// SetFirstPage gets a reference to the given bool and assigns it to the FirstPage field.
func (o *V2reportsGenomeAnnotationReportPage) SetFirstPage(v bool) {
	o.FirstPage = &v
}

// GetReportFormat returns the ReportFormat field value if set, zero value otherwise.
func (o *V2reportsGenomeAnnotationReportPage) GetReportFormat() string {
	if o == nil || o.ReportFormat == nil {
		var ret string
		return ret
	}
	return *o.ReportFormat
}

// GetReportFormatOk returns a tuple with the ReportFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsGenomeAnnotationReportPage) GetReportFormatOk() (*string, bool) {
	if o == nil || o.ReportFormat == nil {
		return nil, false
	}
	return o.ReportFormat, true
}

// HasReportFormat returns a boolean if a field has been set.
func (o *V2reportsGenomeAnnotationReportPage) HasReportFormat() bool {
	if o != nil && o.ReportFormat != nil {
		return true
	}

	return false
}

// SetReportFormat gets a reference to the given string and assigns it to the ReportFormat field.
func (o *V2reportsGenomeAnnotationReportPage) SetReportFormat(v string) {
	o.ReportFormat = &v
}

func (o V2reportsGenomeAnnotationReportPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reports != nil {
		toSerialize["reports"] = o.Reports
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	if o.ReportType != nil {
		toSerialize["_report_type"] = o.ReportType
	}
	if o.ReportFields != nil {
		toSerialize["_report_fields"] = o.ReportFields
	}
	if o.FirstPage != nil {
		toSerialize["_first_page"] = o.FirstPage
	}
	if o.ReportFormat != nil {
		toSerialize["_report_format"] = o.ReportFormat
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsGenomeAnnotationReportPage struct {
	value *V2reportsGenomeAnnotationReportPage
	isSet bool
}

func (v NullableV2reportsGenomeAnnotationReportPage) Get() *V2reportsGenomeAnnotationReportPage {
	return v.value
}

func (v *NullableV2reportsGenomeAnnotationReportPage) Set(val *V2reportsGenomeAnnotationReportPage) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsGenomeAnnotationReportPage) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsGenomeAnnotationReportPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsGenomeAnnotationReportPage(val *V2reportsGenomeAnnotationReportPage) *NullableV2reportsGenomeAnnotationReportPage {
	return &NullableV2reportsGenomeAnnotationReportPage{value: val, isSet: true}
}

func (v NullableV2reportsGenomeAnnotationReportPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsGenomeAnnotationReportPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


