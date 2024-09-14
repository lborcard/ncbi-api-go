# V2reportsGenomicLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenomicAccessionVersion** | Pointer to **string** |  | [optional] 
**SequenceName** | Pointer to **string** |  | [optional] 
**GenomicRange** | Pointer to [**V2reportsRange**](V2reportsRange.md) |  | [optional] 
**Exons** | Pointer to [**[]V2reportsRange**](V2reportsRange.md) |  | [optional] 

## Methods

### NewV2reportsGenomicLocation

`func NewV2reportsGenomicLocation() *V2reportsGenomicLocation`

NewV2reportsGenomicLocation instantiates a new V2reportsGenomicLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsGenomicLocationWithDefaults

`func NewV2reportsGenomicLocationWithDefaults() *V2reportsGenomicLocation`

NewV2reportsGenomicLocationWithDefaults instantiates a new V2reportsGenomicLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenomicAccessionVersion

`func (o *V2reportsGenomicLocation) GetGenomicAccessionVersion() string`

GetGenomicAccessionVersion returns the GenomicAccessionVersion field if non-nil, zero value otherwise.

### GetGenomicAccessionVersionOk

`func (o *V2reportsGenomicLocation) GetGenomicAccessionVersionOk() (*string, bool)`

GetGenomicAccessionVersionOk returns a tuple with the GenomicAccessionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicAccessionVersion

`func (o *V2reportsGenomicLocation) SetGenomicAccessionVersion(v string)`

SetGenomicAccessionVersion sets GenomicAccessionVersion field to given value.

### HasGenomicAccessionVersion

`func (o *V2reportsGenomicLocation) HasGenomicAccessionVersion() bool`

HasGenomicAccessionVersion returns a boolean if a field has been set.

### GetSequenceName

`func (o *V2reportsGenomicLocation) GetSequenceName() string`

GetSequenceName returns the SequenceName field if non-nil, zero value otherwise.

### GetSequenceNameOk

`func (o *V2reportsGenomicLocation) GetSequenceNameOk() (*string, bool)`

GetSequenceNameOk returns a tuple with the SequenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceName

`func (o *V2reportsGenomicLocation) SetSequenceName(v string)`

SetSequenceName sets SequenceName field to given value.

### HasSequenceName

`func (o *V2reportsGenomicLocation) HasSequenceName() bool`

HasSequenceName returns a boolean if a field has been set.

### GetGenomicRange

`func (o *V2reportsGenomicLocation) GetGenomicRange() V2reportsRange`

GetGenomicRange returns the GenomicRange field if non-nil, zero value otherwise.

### GetGenomicRangeOk

`func (o *V2reportsGenomicLocation) GetGenomicRangeOk() (*V2reportsRange, bool)`

GetGenomicRangeOk returns a tuple with the GenomicRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicRange

`func (o *V2reportsGenomicLocation) SetGenomicRange(v V2reportsRange)`

SetGenomicRange sets GenomicRange field to given value.

### HasGenomicRange

`func (o *V2reportsGenomicLocation) HasGenomicRange() bool`

HasGenomicRange returns a boolean if a field has been set.

### GetExons

`func (o *V2reportsGenomicLocation) GetExons() []V2reportsRange`

GetExons returns the Exons field if non-nil, zero value otherwise.

### GetExonsOk

`func (o *V2reportsGenomicLocation) GetExonsOk() (*[]V2reportsRange, bool)`

GetExonsOk returns a tuple with the Exons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExons

`func (o *V2reportsGenomicLocation) SetExons(v []V2reportsRange)`

SetExons sets Exons field to given value.

### HasExons

`func (o *V2reportsGenomicLocation) HasExons() bool`

HasExons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


