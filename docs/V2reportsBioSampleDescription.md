# V2reportsBioSampleDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Organism** | Pointer to [**V2reportsOrganism**](V2reportsOrganism.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewV2reportsBioSampleDescription

`func NewV2reportsBioSampleDescription() *V2reportsBioSampleDescription`

NewV2reportsBioSampleDescription instantiates a new V2reportsBioSampleDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsBioSampleDescriptionWithDefaults

`func NewV2reportsBioSampleDescriptionWithDefaults() *V2reportsBioSampleDescription`

NewV2reportsBioSampleDescriptionWithDefaults instantiates a new V2reportsBioSampleDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *V2reportsBioSampleDescription) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V2reportsBioSampleDescription) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V2reportsBioSampleDescription) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *V2reportsBioSampleDescription) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOrganism

`func (o *V2reportsBioSampleDescription) GetOrganism() V2reportsOrganism`

GetOrganism returns the Organism field if non-nil, zero value otherwise.

### GetOrganismOk

`func (o *V2reportsBioSampleDescription) GetOrganismOk() (*V2reportsOrganism, bool)`

GetOrganismOk returns a tuple with the Organism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganism

`func (o *V2reportsBioSampleDescription) SetOrganism(v V2reportsOrganism)`

SetOrganism sets Organism field to given value.

### HasOrganism

`func (o *V2reportsBioSampleDescription) HasOrganism() bool`

HasOrganism returns a boolean if a field has been set.

### GetComment

`func (o *V2reportsBioSampleDescription) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V2reportsBioSampleDescription) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V2reportsBioSampleDescription) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V2reportsBioSampleDescription) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


