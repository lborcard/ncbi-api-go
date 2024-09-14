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

// V2TaxonomyFilteredSubtreeResponseEdge struct for V2TaxonomyFilteredSubtreeResponseEdge
type V2TaxonomyFilteredSubtreeResponseEdge struct {
	VisibleChildren []int32 `json:"visible_children,omitempty"`
	ChildrenStatus *V2TaxonomyFilteredSubtreeResponseEdgeChildStatus `json:"children_status,omitempty"`
}

// NewV2TaxonomyFilteredSubtreeResponseEdge instantiates a new V2TaxonomyFilteredSubtreeResponseEdge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TaxonomyFilteredSubtreeResponseEdge() *V2TaxonomyFilteredSubtreeResponseEdge {
	this := V2TaxonomyFilteredSubtreeResponseEdge{}
	var childrenStatus V2TaxonomyFilteredSubtreeResponseEdgeChildStatus = V2TAXONOMYFILTEREDSUBTREERESPONSEEDGECHILDSTATUS_UNSPECIFIED
	this.ChildrenStatus = &childrenStatus
	return &this
}

// NewV2TaxonomyFilteredSubtreeResponseEdgeWithDefaults instantiates a new V2TaxonomyFilteredSubtreeResponseEdge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TaxonomyFilteredSubtreeResponseEdgeWithDefaults() *V2TaxonomyFilteredSubtreeResponseEdge {
	this := V2TaxonomyFilteredSubtreeResponseEdge{}
	var childrenStatus V2TaxonomyFilteredSubtreeResponseEdgeChildStatus = V2TAXONOMYFILTEREDSUBTREERESPONSEEDGECHILDSTATUS_UNSPECIFIED
	this.ChildrenStatus = &childrenStatus
	return &this
}

// GetVisibleChildren returns the VisibleChildren field value if set, zero value otherwise.
func (o *V2TaxonomyFilteredSubtreeResponseEdge) GetVisibleChildren() []int32 {
	if o == nil || o.VisibleChildren == nil {
		var ret []int32
		return ret
	}
	return o.VisibleChildren
}

// GetVisibleChildrenOk returns a tuple with the VisibleChildren field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyFilteredSubtreeResponseEdge) GetVisibleChildrenOk() ([]int32, bool) {
	if o == nil || o.VisibleChildren == nil {
		return nil, false
	}
	return o.VisibleChildren, true
}

// HasVisibleChildren returns a boolean if a field has been set.
func (o *V2TaxonomyFilteredSubtreeResponseEdge) HasVisibleChildren() bool {
	if o != nil && o.VisibleChildren != nil {
		return true
	}

	return false
}

// SetVisibleChildren gets a reference to the given []int32 and assigns it to the VisibleChildren field.
func (o *V2TaxonomyFilteredSubtreeResponseEdge) SetVisibleChildren(v []int32) {
	o.VisibleChildren = v
}

// GetChildrenStatus returns the ChildrenStatus field value if set, zero value otherwise.
func (o *V2TaxonomyFilteredSubtreeResponseEdge) GetChildrenStatus() V2TaxonomyFilteredSubtreeResponseEdgeChildStatus {
	if o == nil || o.ChildrenStatus == nil {
		var ret V2TaxonomyFilteredSubtreeResponseEdgeChildStatus
		return ret
	}
	return *o.ChildrenStatus
}

// GetChildrenStatusOk returns a tuple with the ChildrenStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TaxonomyFilteredSubtreeResponseEdge) GetChildrenStatusOk() (*V2TaxonomyFilteredSubtreeResponseEdgeChildStatus, bool) {
	if o == nil || o.ChildrenStatus == nil {
		return nil, false
	}
	return o.ChildrenStatus, true
}

// HasChildrenStatus returns a boolean if a field has been set.
func (o *V2TaxonomyFilteredSubtreeResponseEdge) HasChildrenStatus() bool {
	if o != nil && o.ChildrenStatus != nil {
		return true
	}

	return false
}

// SetChildrenStatus gets a reference to the given V2TaxonomyFilteredSubtreeResponseEdgeChildStatus and assigns it to the ChildrenStatus field.
func (o *V2TaxonomyFilteredSubtreeResponseEdge) SetChildrenStatus(v V2TaxonomyFilteredSubtreeResponseEdgeChildStatus) {
	o.ChildrenStatus = &v
}

func (o V2TaxonomyFilteredSubtreeResponseEdge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VisibleChildren != nil {
		toSerialize["visible_children"] = o.VisibleChildren
	}
	if o.ChildrenStatus != nil {
		toSerialize["children_status"] = o.ChildrenStatus
	}
	return json.Marshal(toSerialize)
}

type NullableV2TaxonomyFilteredSubtreeResponseEdge struct {
	value *V2TaxonomyFilteredSubtreeResponseEdge
	isSet bool
}

func (v NullableV2TaxonomyFilteredSubtreeResponseEdge) Get() *V2TaxonomyFilteredSubtreeResponseEdge {
	return v.value
}

func (v *NullableV2TaxonomyFilteredSubtreeResponseEdge) Set(val *V2TaxonomyFilteredSubtreeResponseEdge) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TaxonomyFilteredSubtreeResponseEdge) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TaxonomyFilteredSubtreeResponseEdge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TaxonomyFilteredSubtreeResponseEdge(val *V2TaxonomyFilteredSubtreeResponseEdge) *NullableV2TaxonomyFilteredSubtreeResponseEdge {
	return &NullableV2TaxonomyFilteredSubtreeResponseEdge{value: val, isSet: true}
}

func (v NullableV2TaxonomyFilteredSubtreeResponseEdge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TaxonomyFilteredSubtreeResponseEdge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


