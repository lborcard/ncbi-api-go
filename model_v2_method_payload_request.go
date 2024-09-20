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

// V2MethodPayloadRequest struct for V2MethodPayloadRequest
type V2MethodPayloadRequest struct {
	Method *string `json:"method,omitempty"`
	Payload *string `json:"payload,omitempty"`
}

// NewV2MethodPayloadRequest instantiates a new V2MethodPayloadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2MethodPayloadRequest() *V2MethodPayloadRequest {
	this := V2MethodPayloadRequest{}
	return &this
}

// NewV2MethodPayloadRequestWithDefaults instantiates a new V2MethodPayloadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2MethodPayloadRequestWithDefaults() *V2MethodPayloadRequest {
	this := V2MethodPayloadRequest{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *V2MethodPayloadRequest) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MethodPayloadRequest) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *V2MethodPayloadRequest) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *V2MethodPayloadRequest) SetMethod(v string) {
	o.Method = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *V2MethodPayloadRequest) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2MethodPayloadRequest) GetPayloadOk() (*string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *V2MethodPayloadRequest) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *V2MethodPayloadRequest) SetPayload(v string) {
	o.Payload = &v
}

func (o V2MethodPayloadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	return json.Marshal(toSerialize)
}

type NullableV2MethodPayloadRequest struct {
	value *V2MethodPayloadRequest
	isSet bool
}

func (v NullableV2MethodPayloadRequest) Get() *V2MethodPayloadRequest {
	return v.value
}

func (v *NullableV2MethodPayloadRequest) Set(val *V2MethodPayloadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2MethodPayloadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2MethodPayloadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2MethodPayloadRequest(val *V2MethodPayloadRequest) *NullableV2MethodPayloadRequest {
	return &NullableV2MethodPayloadRequest{value: val, isSet: true}
}

func (v NullableV2MethodPayloadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2MethodPayloadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


