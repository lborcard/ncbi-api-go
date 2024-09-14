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

// V2reportsBioSampleId struct for V2reportsBioSampleId
type V2reportsBioSampleId struct {
	Db *string `json:"db,omitempty"`
	Label *string `json:"label,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewV2reportsBioSampleId instantiates a new V2reportsBioSampleId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsBioSampleId() *V2reportsBioSampleId {
	this := V2reportsBioSampleId{}
	return &this
}

// NewV2reportsBioSampleIdWithDefaults instantiates a new V2reportsBioSampleId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsBioSampleIdWithDefaults() *V2reportsBioSampleId {
	this := V2reportsBioSampleId{}
	return &this
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *V2reportsBioSampleId) GetDb() string {
	if o == nil || o.Db == nil {
		var ret string
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsBioSampleId) GetDbOk() (*string, bool) {
	if o == nil || o.Db == nil {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *V2reportsBioSampleId) HasDb() bool {
	if o != nil && o.Db != nil {
		return true
	}

	return false
}

// SetDb gets a reference to the given string and assigns it to the Db field.
func (o *V2reportsBioSampleId) SetDb(v string) {
	o.Db = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *V2reportsBioSampleId) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsBioSampleId) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *V2reportsBioSampleId) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *V2reportsBioSampleId) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *V2reportsBioSampleId) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsBioSampleId) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *V2reportsBioSampleId) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *V2reportsBioSampleId) SetValue(v string) {
	o.Value = &v
}

func (o V2reportsBioSampleId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Db != nil {
		toSerialize["db"] = o.Db
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsBioSampleId struct {
	value *V2reportsBioSampleId
	isSet bool
}

func (v NullableV2reportsBioSampleId) Get() *V2reportsBioSampleId {
	return v.value
}

func (v *NullableV2reportsBioSampleId) Set(val *V2reportsBioSampleId) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsBioSampleId) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsBioSampleId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsBioSampleId(val *V2reportsBioSampleId) *NullableV2reportsBioSampleId {
	return &NullableV2reportsBioSampleId{value: val, isSet: true}
}

func (v NullableV2reportsBioSampleId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsBioSampleId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


