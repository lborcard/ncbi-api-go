/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 2 API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated zip archive](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v2alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ncbi-openapi-v2_goland

import (
	"encoding/json"
)

// V2DownloadSummaryHydrated struct for V2DownloadSummaryHydrated
type V2DownloadSummaryHydrated struct {
	EstimatedFileSizeMb *int32 `json:"estimated_file_size_mb,omitempty"`
	Url *string `json:"url,omitempty"`
	CliDownloadCommandLine *string `json:"cli_download_command_line,omitempty"`
}

// NewV2DownloadSummaryHydrated instantiates a new V2DownloadSummaryHydrated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2DownloadSummaryHydrated() *V2DownloadSummaryHydrated {
	this := V2DownloadSummaryHydrated{}
	return &this
}

// NewV2DownloadSummaryHydratedWithDefaults instantiates a new V2DownloadSummaryHydrated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2DownloadSummaryHydratedWithDefaults() *V2DownloadSummaryHydrated {
	this := V2DownloadSummaryHydrated{}
	return &this
}

// GetEstimatedFileSizeMb returns the EstimatedFileSizeMb field value if set, zero value otherwise.
func (o *V2DownloadSummaryHydrated) GetEstimatedFileSizeMb() int32 {
	if o == nil || o.EstimatedFileSizeMb == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedFileSizeMb
}

// GetEstimatedFileSizeMbOk returns a tuple with the EstimatedFileSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2DownloadSummaryHydrated) GetEstimatedFileSizeMbOk() (*int32, bool) {
	if o == nil || o.EstimatedFileSizeMb == nil {
		return nil, false
	}
	return o.EstimatedFileSizeMb, true
}

// HasEstimatedFileSizeMb returns a boolean if a field has been set.
func (o *V2DownloadSummaryHydrated) HasEstimatedFileSizeMb() bool {
	if o != nil && o.EstimatedFileSizeMb != nil {
		return true
	}

	return false
}

// SetEstimatedFileSizeMb gets a reference to the given int32 and assigns it to the EstimatedFileSizeMb field.
func (o *V2DownloadSummaryHydrated) SetEstimatedFileSizeMb(v int32) {
	o.EstimatedFileSizeMb = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *V2DownloadSummaryHydrated) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2DownloadSummaryHydrated) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *V2DownloadSummaryHydrated) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *V2DownloadSummaryHydrated) SetUrl(v string) {
	o.Url = &v
}

// GetCliDownloadCommandLine returns the CliDownloadCommandLine field value if set, zero value otherwise.
func (o *V2DownloadSummaryHydrated) GetCliDownloadCommandLine() string {
	if o == nil || o.CliDownloadCommandLine == nil {
		var ret string
		return ret
	}
	return *o.CliDownloadCommandLine
}

// GetCliDownloadCommandLineOk returns a tuple with the CliDownloadCommandLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2DownloadSummaryHydrated) GetCliDownloadCommandLineOk() (*string, bool) {
	if o == nil || o.CliDownloadCommandLine == nil {
		return nil, false
	}
	return o.CliDownloadCommandLine, true
}

// HasCliDownloadCommandLine returns a boolean if a field has been set.
func (o *V2DownloadSummaryHydrated) HasCliDownloadCommandLine() bool {
	if o != nil && o.CliDownloadCommandLine != nil {
		return true
	}

	return false
}

// SetCliDownloadCommandLine gets a reference to the given string and assigns it to the CliDownloadCommandLine field.
func (o *V2DownloadSummaryHydrated) SetCliDownloadCommandLine(v string) {
	o.CliDownloadCommandLine = &v
}

func (o V2DownloadSummaryHydrated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EstimatedFileSizeMb != nil {
		toSerialize["estimated_file_size_mb"] = o.EstimatedFileSizeMb
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.CliDownloadCommandLine != nil {
		toSerialize["cli_download_command_line"] = o.CliDownloadCommandLine
	}
	return json.Marshal(toSerialize)
}

type NullableV2DownloadSummaryHydrated struct {
	value *V2DownloadSummaryHydrated
	isSet bool
}

func (v NullableV2DownloadSummaryHydrated) Get() *V2DownloadSummaryHydrated {
	return v.value
}

func (v *NullableV2DownloadSummaryHydrated) Set(val *V2DownloadSummaryHydrated) {
	v.value = val
	v.isSet = true
}

func (v NullableV2DownloadSummaryHydrated) IsSet() bool {
	return v.isSet
}

func (v *NullableV2DownloadSummaryHydrated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2DownloadSummaryHydrated(val *V2DownloadSummaryHydrated) *NullableV2DownloadSummaryHydrated {
	return &NullableV2DownloadSummaryHydrated{value: val, isSet: true}
}

func (v NullableV2DownloadSummaryHydrated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2DownloadSummaryHydrated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


