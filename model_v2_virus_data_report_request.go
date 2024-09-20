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

// V2VirusDataReportRequest struct for V2VirusDataReportRequest
type V2VirusDataReportRequest struct {
	Filter *V2VirusDatasetFilter `json:"filter,omitempty"`
	ReturnedContent *V2VirusDataReportRequestContentType `json:"returned_content,omitempty"`
	TableFields []string `json:"table_fields,omitempty"`
	TableFormat *string `json:"table_format,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	PageToken *string `json:"page_token,omitempty"`
}

// NewV2VirusDataReportRequest instantiates a new V2VirusDataReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2VirusDataReportRequest() *V2VirusDataReportRequest {
	this := V2VirusDataReportRequest{}
	var returnedContent V2VirusDataReportRequestContentType = V2VIRUSDATAREPORTREQUESTCONTENTTYPE_COMPLETE
	this.ReturnedContent = &returnedContent
	return &this
}

// NewV2VirusDataReportRequestWithDefaults instantiates a new V2VirusDataReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2VirusDataReportRequestWithDefaults() *V2VirusDataReportRequest {
	this := V2VirusDataReportRequest{}
	var returnedContent V2VirusDataReportRequestContentType = V2VIRUSDATAREPORTREQUESTCONTENTTYPE_COMPLETE
	this.ReturnedContent = &returnedContent
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *V2VirusDataReportRequest) GetFilter() V2VirusDatasetFilter {
	if o == nil || o.Filter == nil {
		var ret V2VirusDatasetFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDataReportRequest) GetFilterOk() (*V2VirusDatasetFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *V2VirusDataReportRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given V2VirusDatasetFilter and assigns it to the Filter field.
func (o *V2VirusDataReportRequest) SetFilter(v V2VirusDatasetFilter) {
	o.Filter = &v
}

// GetReturnedContent returns the ReturnedContent field value if set, zero value otherwise.
func (o *V2VirusDataReportRequest) GetReturnedContent() V2VirusDataReportRequestContentType {
	if o == nil || o.ReturnedContent == nil {
		var ret V2VirusDataReportRequestContentType
		return ret
	}
	return *o.ReturnedContent
}

// GetReturnedContentOk returns a tuple with the ReturnedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDataReportRequest) GetReturnedContentOk() (*V2VirusDataReportRequestContentType, bool) {
	if o == nil || o.ReturnedContent == nil {
		return nil, false
	}
	return o.ReturnedContent, true
}

// HasReturnedContent returns a boolean if a field has been set.
func (o *V2VirusDataReportRequest) HasReturnedContent() bool {
	if o != nil && o.ReturnedContent != nil {
		return true
	}

	return false
}

// SetReturnedContent gets a reference to the given V2VirusDataReportRequestContentType and assigns it to the ReturnedContent field.
func (o *V2VirusDataReportRequest) SetReturnedContent(v V2VirusDataReportRequestContentType) {
	o.ReturnedContent = &v
}

// GetTableFields returns the TableFields field value if set, zero value otherwise.
func (o *V2VirusDataReportRequest) GetTableFields() []string {
	if o == nil || o.TableFields == nil {
		var ret []string
		return ret
	}
	return o.TableFields
}

// GetTableFieldsOk returns a tuple with the TableFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDataReportRequest) GetTableFieldsOk() ([]string, bool) {
	if o == nil || o.TableFields == nil {
		return nil, false
	}
	return o.TableFields, true
}

// HasTableFields returns a boolean if a field has been set.
func (o *V2VirusDataReportRequest) HasTableFields() bool {
	if o != nil && o.TableFields != nil {
		return true
	}

	return false
}

// SetTableFields gets a reference to the given []string and assigns it to the TableFields field.
func (o *V2VirusDataReportRequest) SetTableFields(v []string) {
	o.TableFields = v
}

// GetTableFormat returns the TableFormat field value if set, zero value otherwise.
func (o *V2VirusDataReportRequest) GetTableFormat() string {
	if o == nil || o.TableFormat == nil {
		var ret string
		return ret
	}
	return *o.TableFormat
}

// GetTableFormatOk returns a tuple with the TableFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDataReportRequest) GetTableFormatOk() (*string, bool) {
	if o == nil || o.TableFormat == nil {
		return nil, false
	}
	return o.TableFormat, true
}

// HasTableFormat returns a boolean if a field has been set.
func (o *V2VirusDataReportRequest) HasTableFormat() bool {
	if o != nil && o.TableFormat != nil {
		return true
	}

	return false
}

// SetTableFormat gets a reference to the given string and assigns it to the TableFormat field.
func (o *V2VirusDataReportRequest) SetTableFormat(v string) {
	o.TableFormat = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *V2VirusDataReportRequest) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDataReportRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *V2VirusDataReportRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *V2VirusDataReportRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *V2VirusDataReportRequest) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2VirusDataReportRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || o.PageToken == nil {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *V2VirusDataReportRequest) HasPageToken() bool {
	if o != nil && o.PageToken != nil {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *V2VirusDataReportRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o V2VirusDataReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.ReturnedContent != nil {
		toSerialize["returned_content"] = o.ReturnedContent
	}
	if o.TableFields != nil {
		toSerialize["table_fields"] = o.TableFields
	}
	if o.TableFormat != nil {
		toSerialize["table_format"] = o.TableFormat
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PageToken != nil {
		toSerialize["page_token"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

type NullableV2VirusDataReportRequest struct {
	value *V2VirusDataReportRequest
	isSet bool
}

func (v NullableV2VirusDataReportRequest) Get() *V2VirusDataReportRequest {
	return v.value
}

func (v *NullableV2VirusDataReportRequest) Set(val *V2VirusDataReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2VirusDataReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2VirusDataReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2VirusDataReportRequest(val *V2VirusDataReportRequest) *NullableV2VirusDataReportRequest {
	return &NullableV2VirusDataReportRequest{value: val, isSet: true}
}

func (v NullableV2VirusDataReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2VirusDataReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


