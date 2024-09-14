# V2reportsNameAndAuthorityNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**NoteClassifier** | Pointer to [**V2reportsNameAndAuthorityNoteClassifier**](V2reportsNameAndAuthorityNoteClassifier.md) |  | [optional] [default to V2REPORTSNAMEANDAUTHORITYNOTECLASSIFIER_NO_AUTHORITY_CLASSIFIER]

## Methods

### NewV2reportsNameAndAuthorityNote

`func NewV2reportsNameAndAuthorityNote() *V2reportsNameAndAuthorityNote`

NewV2reportsNameAndAuthorityNote instantiates a new V2reportsNameAndAuthorityNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsNameAndAuthorityNoteWithDefaults

`func NewV2reportsNameAndAuthorityNoteWithDefaults() *V2reportsNameAndAuthorityNote`

NewV2reportsNameAndAuthorityNoteWithDefaults instantiates a new V2reportsNameAndAuthorityNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2reportsNameAndAuthorityNote) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2reportsNameAndAuthorityNote) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2reportsNameAndAuthorityNote) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2reportsNameAndAuthorityNote) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *V2reportsNameAndAuthorityNote) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *V2reportsNameAndAuthorityNote) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *V2reportsNameAndAuthorityNote) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *V2reportsNameAndAuthorityNote) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetNoteClassifier

`func (o *V2reportsNameAndAuthorityNote) GetNoteClassifier() V2reportsNameAndAuthorityNoteClassifier`

GetNoteClassifier returns the NoteClassifier field if non-nil, zero value otherwise.

### GetNoteClassifierOk

`func (o *V2reportsNameAndAuthorityNote) GetNoteClassifierOk() (*V2reportsNameAndAuthorityNoteClassifier, bool)`

GetNoteClassifierOk returns a tuple with the NoteClassifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteClassifier

`func (o *V2reportsNameAndAuthorityNote) SetNoteClassifier(v V2reportsNameAndAuthorityNoteClassifier)`

SetNoteClassifier sets NoteClassifier field to given value.

### HasNoteClassifier

`func (o *V2reportsNameAndAuthorityNote) HasNoteClassifier() bool`

HasNoteClassifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


