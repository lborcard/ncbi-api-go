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

// V2reportsBioSampleOwner struct for V2reportsBioSampleOwner
type V2reportsBioSampleOwner struct {
	Name *string `json:"name,omitempty"`
	Contacts []V2reportsBioSampleContact `json:"contacts,omitempty"`
}

// NewV2reportsBioSampleOwner instantiates a new V2reportsBioSampleOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2reportsBioSampleOwner() *V2reportsBioSampleOwner {
	this := V2reportsBioSampleOwner{}
	return &this
}

// NewV2reportsBioSampleOwnerWithDefaults instantiates a new V2reportsBioSampleOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2reportsBioSampleOwnerWithDefaults() *V2reportsBioSampleOwner {
	this := V2reportsBioSampleOwner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V2reportsBioSampleOwner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsBioSampleOwner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V2reportsBioSampleOwner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V2reportsBioSampleOwner) SetName(v string) {
	o.Name = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *V2reportsBioSampleOwner) GetContacts() []V2reportsBioSampleContact {
	if o == nil || o.Contacts == nil {
		var ret []V2reportsBioSampleContact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2reportsBioSampleOwner) GetContactsOk() ([]V2reportsBioSampleContact, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *V2reportsBioSampleOwner) HasContacts() bool {
	if o != nil && o.Contacts != nil {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []V2reportsBioSampleContact and assigns it to the Contacts field.
func (o *V2reportsBioSampleOwner) SetContacts(v []V2reportsBioSampleContact) {
	o.Contacts = v
}

func (o V2reportsBioSampleOwner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	return json.Marshal(toSerialize)
}

type NullableV2reportsBioSampleOwner struct {
	value *V2reportsBioSampleOwner
	isSet bool
}

func (v NullableV2reportsBioSampleOwner) Get() *V2reportsBioSampleOwner {
	return v.value
}

func (v *NullableV2reportsBioSampleOwner) Set(val *V2reportsBioSampleOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2reportsBioSampleOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2reportsBioSampleOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2reportsBioSampleOwner(val *V2reportsBioSampleOwner) *NullableV2reportsBioSampleOwner {
	return &NullableV2reportsBioSampleOwner{value: val, isSet: true}
}

func (v NullableV2reportsBioSampleOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2reportsBioSampleOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


