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

// V2reportsBioSampleDataReportPage struct for V2reportsBioSampleDataReportPage
type V2reportsBioSampleDataReportPage struct {
	Reports []V2reportsBioSampleDataReport `json:"reports,omitempty"`
	TotalCount *int32 `json:"total_count,omitempty"`
	NextPageToken *string `json:"next_page_token,omitempty"`
	Messages []V2reportsMessage `json:"messages,omitempty"`
	ReportType *string `json:"_report_type,omitempty"`
	ReportFields []string `json:"_report_fields,omitempty"`
	ReportFormat *string `json:"_report_format,omitempty"`
}

// NewV2reportsBioSampleDataReportPage instantiates a new V2reportsBioSampleDataReportPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsBioSampleDataReportPage() *V2reportsBioSampleDataReportPage {
	this := V2reportsBioSampleDataReportPage{}
	return &this
}

// NewV2reportsBioSampleDataReportPageWithDefaults instantiates a new V2reportsBioSampleDataReportPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsBioSampleDataReportPageWithDefaults() *V2reportsBioSampleDataReportPage {
	this := V2reportsBioSampleDataReportPage{}
	return &this
}

// GetReports returns the Reports field value if set, zero value otherwise.
func (o *V2reportsBioSampleDataReportPage) GetReports() []V2reportsBioSampleDataReport {
	if o == nil || o.Reports == nil {
		var ret []V2reportsBioSampleDataReport
		return ret
	}
	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsBioSampleDataReportPage) GetReportsOk() ([]V2reportsBioSampleDataReport, bool) {
	if o == nil || o.Reports == nil {
		return nil, false
	}
	return o.Reports, true
}

// HasReports returns a boolean if a field has been set.
func (o *V2reportsBioSampleDataReportPage) HasReports() bool {
	if o != nil && o.Reports != nil {
		return true
	}

	return false
}

// SetReports gets a reference to the given []V2reportsBioSampleDataReport and assigns it to the Reports field.
func (o *V2reportsBioSampleDataReportPage) SetReports(v []V2reportsBioSampleDataReport) {
	o.Reports = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *V2reportsBioSampleDataReportPage) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsBioSampleDataReportPage) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *V2reportsBioSampleDataReportPage) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *V2reportsBioSampleDataReportPage) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *V2reportsBioSampleDataReportPage) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsBioSampleDataReportPage) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *V2reportsBioSampleDataReportPage) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *V2reportsBioSampleDataReportPage) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *V2reportsBioSampleDataReportPage) GetMessages() []V2reportsMessage {
	if o == nil || o.Messages == nil {
		var ret []V2reportsMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsBioSampleDataReportPage) GetMessagesOk() ([]V2reportsMessage, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *V2reportsBioSampleDataReportPage) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []V2reportsMessage and assigns it to the Messages field.
func (o *V2reportsBioSampleDataReportPage) SetMessages(v []V2reportsMessage) {
	o.Messages = v
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *V2reportsBioSampleDataReportPage) GetReportType() string {
	if o == nil || o.ReportType == nil {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsBioSampleDataReportPage) GetReportTypeOk() (*string, bool) {
	if o == nil || o.ReportType == nil {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *V2reportsBioSampleDataReportPage) HasReportType() bool {
	if o != nil && o.ReportType != nil {
		return true
	}

	return false
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *V2reportsBioSampleDataReportPage) SetReportType(v string) {
	o.ReportType = &v
}

// GetReportFields returns the ReportFields field value if set, zero value otherwise.
func (o *V2reportsBioSampleDataReportPage) GetReportFields() []string {
	if o == nil || o.ReportFields == nil {
		var ret []string
		return ret
	}
	return o.ReportFields
}

// GetReportFieldsOk returns a tuple with the ReportFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsBioSampleDataReportPage) GetReportFieldsOk() ([]string, bool) {
	if o == nil || o.ReportFields == nil {
		return nil, false
	}
	return o.ReportFields, true
}

// HasReportFields returns a boolean if a field has been set.
func (o *V2reportsBioSampleDataReportPage) HasReportFields() bool {
	if o != nil && o.ReportFields != nil {
		return true
	}

	return false
}

// SetReportFields gets a reference to the given []string and assigns it to the ReportFields field.
func (o *V2reportsBioSampleDataReportPage) SetReportFields(v []string) {
	o.ReportFields = v
}

// GetReportFormat returns the ReportFormat field value if set, zero value otherwise.
func (o *V2reportsBioSampleDataReportPage) GetReportFormat() string {
	if o == nil || o.ReportFormat == nil {
		var ret string
		return ret
	}
	return *o.ReportFormat
}

// GetReportFormatOk returns a tuple with the ReportFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsBioSampleDataReportPage) GetReportFormatOk() (*string, bool) {
	if o == nil || o.ReportFormat == nil {
		return nil, false
	}
	return o.ReportFormat, true
}

// HasReportFormat returns a boolean if a field has been set.
func (o *V2reportsBioSampleDataReportPage) HasReportFormat() bool {
	if o != nil && o.ReportFormat != nil {
		return true
	}

	return false
}

// SetReportFormat gets a reference to the given string and assigns it to the ReportFormat field.
func (o *V2reportsBioSampleDataReportPage) SetReportFormat(v string) {
	o.ReportFormat = &v
}

func (o V2reportsBioSampleDataReportPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reports != nil {
		toSerialize["reports"] = o.Reports
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.ReportType != nil {
		toSerialize["_report_type"] = o.ReportType
	}
	if o.ReportFields != nil {
		toSerialize["_report_fields"] = o.ReportFields
	}
	if o.ReportFormat != nil {
		toSerialize["_report_format"] = o.ReportFormat
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsBioSampleDataReportPage struct {
	value *V2reportsBioSampleDataReportPage
	isSet bool
}

func (v NullableV2reportsBioSampleDataReportPage) Get() *V2reportsBioSampleDataReportPage {
	return v.value
}

func (v *NullableV2reportsBioSampleDataReportPage) Set(val *V2reportsBioSampleDataReportPage) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsBioSampleDataReportPage) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsBioSampleDataReportPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsBioSampleDataReportPage(val *V2reportsBioSampleDataReportPage) *NullableV2reportsBioSampleDataReportPage {
	return &NullableV2reportsBioSampleDataReportPage{value: val, isSet: true}
}

func (v NullableV2reportsBioSampleDataReportPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsBioSampleDataReportPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


