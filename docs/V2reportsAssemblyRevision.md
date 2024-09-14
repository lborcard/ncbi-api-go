# V2reportsAssemblyRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenbankAccession** | Pointer to **string** |  | [optional] 
**RefseqAccession** | Pointer to **string** |  | [optional] 
**AssemblyName** | Pointer to **string** |  | [optional] 
**AssemblyLevel** | Pointer to [**V2reportsAssemblyLevel**](V2reportsAssemblyLevel.md) |  | [optional] [default to V2REPORTSASSEMBLYLEVEL_CHROMOSOME]
**ReleaseDate** | Pointer to **string** |  | [optional] 
**SubmissionDate** | Pointer to **string** |  | [optional] 
**SequencingTechnology** | Pointer to **string** |  | [optional] 

## Methods

### NewV2reportsAssemblyRevision

`func NewV2reportsAssemblyRevision() *V2reportsAssemblyRevision`

NewV2reportsAssemblyRevision instantiates a new V2reportsAssemblyRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsAssemblyRevisionWithDefaults

`func NewV2reportsAssemblyRevisionWithDefaults() *V2reportsAssemblyRevision`

NewV2reportsAssemblyRevisionWithDefaults instantiates a new V2reportsAssemblyRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenbankAccession

`func (o *V2reportsAssemblyRevision) GetGenbankAccession() string`

GetGenbankAccession returns the GenbankAccession field if non-nil, zero value otherwise.

### GetGenbankAccessionOk

`func (o *V2reportsAssemblyRevision) GetGenbankAccessionOk() (*string, bool)`

GetGenbankAccessionOk returns a tuple with the GenbankAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenbankAccession

`func (o *V2reportsAssemblyRevision) SetGenbankAccession(v string)`

SetGenbankAccession sets GenbankAccession field to given value.

### HasGenbankAccession

`func (o *V2reportsAssemblyRevision) HasGenbankAccession() bool`

HasGenbankAccession returns a boolean if a field has been set.

### GetRefseqAccession

`func (o *V2reportsAssemblyRevision) GetRefseqAccession() string`

GetRefseqAccession returns the RefseqAccession field if non-nil, zero value otherwise.

### GetRefseqAccessionOk

`func (o *V2reportsAssemblyRevision) GetRefseqAccessionOk() (*string, bool)`

GetRefseqAccessionOk returns a tuple with the RefseqAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqAccession

`func (o *V2reportsAssemblyRevision) SetRefseqAccession(v string)`

SetRefseqAccession sets RefseqAccession field to given value.

### HasRefseqAccession

`func (o *V2reportsAssemblyRevision) HasRefseqAccession() bool`

HasRefseqAccession returns a boolean if a field has been set.

### GetAssemblyName

`func (o *V2reportsAssemblyRevision) GetAssemblyName() string`

GetAssemblyName returns the AssemblyName field if non-nil, zero value otherwise.

### GetAssemblyNameOk

`func (o *V2reportsAssemblyRevision) GetAssemblyNameOk() (*string, bool)`

GetAssemblyNameOk returns a tuple with the AssemblyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyName

`func (o *V2reportsAssemblyRevision) SetAssemblyName(v string)`

SetAssemblyName sets AssemblyName field to given value.

### HasAssemblyName

`func (o *V2reportsAssemblyRevision) HasAssemblyName() bool`

HasAssemblyName returns a boolean if a field has been set.

### GetAssemblyLevel

`func (o *V2reportsAssemblyRevision) GetAssemblyLevel() V2reportsAssemblyLevel`

GetAssemblyLevel returns the AssemblyLevel field if non-nil, zero value otherwise.

### GetAssemblyLevelOk

`func (o *V2reportsAssemblyRevision) GetAssemblyLevelOk() (*V2reportsAssemblyLevel, bool)`

GetAssemblyLevelOk returns a tuple with the AssemblyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyLevel

`func (o *V2reportsAssemblyRevision) SetAssemblyLevel(v V2reportsAssemblyLevel)`

SetAssemblyLevel sets AssemblyLevel field to given value.

### HasAssemblyLevel

`func (o *V2reportsAssemblyRevision) HasAssemblyLevel() bool`

HasAssemblyLevel returns a boolean if a field has been set.

### GetReleaseDate

`func (o *V2reportsAssemblyRevision) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *V2reportsAssemblyRevision) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *V2reportsAssemblyRevision) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *V2reportsAssemblyRevision) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetSubmissionDate

`func (o *V2reportsAssemblyRevision) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *V2reportsAssemblyRevision) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *V2reportsAssemblyRevision) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *V2reportsAssemblyRevision) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.

### GetSequencingTechnology

`func (o *V2reportsAssemblyRevision) GetSequencingTechnology() string`

GetSequencingTechnology returns the SequencingTechnology field if non-nil, zero value otherwise.

### GetSequencingTechnologyOk

`func (o *V2reportsAssemblyRevision) GetSequencingTechnologyOk() (*string, bool)`

GetSequencingTechnologyOk returns a tuple with the SequencingTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequencingTechnology

`func (o *V2reportsAssemblyRevision) SetSequencingTechnology(v string)`

SetSequencingTechnology sets SequencingTechnology field to given value.

### HasSequencingTechnology

`func (o *V2reportsAssemblyRevision) HasSequencingTechnology() bool`

HasSequencingTechnology returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


