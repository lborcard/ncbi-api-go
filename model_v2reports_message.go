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

// V2reportsMessage struct for V2reportsMessage
type V2reportsMessage struct {
	Error *V2reportsError `json:"error,omitempty"`
	Warning *V2reportsWarning `json:"warning,omitempty"`
}

// NewV2reportsMessage instantiates a new V2reportsMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsMessage() *V2reportsMessage {
	this := V2reportsMessage{}
	return &this
}

// NewV2reportsMessageWithDefaults instantiates a new V2reportsMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsMessageWithDefaults() *V2reportsMessage {
	this := V2reportsMessage{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V2reportsMessage) GetError() V2reportsError {
	if o == nil || o.Error == nil {
		var ret V2reportsError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsMessage) GetErrorOk() (*V2reportsError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V2reportsMessage) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given V2reportsError and assigns it to the Error field.
func (o *V2reportsMessage) SetError(v V2reportsError) {
	o.Error = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *V2reportsMessage) GetWarning() V2reportsWarning {
	if o == nil || o.Warning == nil {
		var ret V2reportsWarning
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsMessage) GetWarningOk() (*V2reportsWarning, bool) {
	if o == nil || o.Warning == nil {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *V2reportsMessage) HasWarning() bool {
	if o != nil && o.Warning != nil {
		return true
	}

	return false
}

// SetWarning gets a reference to the given V2reportsWarning and assigns it to the Warning field.
func (o *V2reportsMessage) SetWarning(v V2reportsWarning) {
	o.Warning = &v
}

func (o V2reportsMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Warning != nil {
		toSerialize["warning"] = o.Warning
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsMessage struct {
	value *V2reportsMessage
	isSet bool
}

func (v NullableV2reportsMessage) Get() *V2reportsMessage {
	return v.value
}

func (v *NullableV2reportsMessage) Set(val *V2reportsMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsMessage(val *V2reportsMessage) *NullableV2reportsMessage {
	return &NullableV2reportsMessage{value: val, isSet: true}
}

func (v NullableV2reportsMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


