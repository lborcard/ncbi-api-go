# V2reportsNameAndAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Authority** | Pointer to **string** |  | [optional] 
**TypeStrains** | Pointer to [**[]V2reportsTaxonomyTypeMaterial**](V2reportsTaxonomyTypeMaterial.md) |  | [optional] 
**CuratorSynonym** | Pointer to **string** |  | [optional] 
**HomotypicSynonyms** | Pointer to [**[]V2reportsNameAndAuthority**](V2reportsNameAndAuthority.md) |  | [optional] 
**HeterotypicSynonyms** | Pointer to [**[]V2reportsNameAndAuthority**](V2reportsNameAndAuthority.md) |  | [optional] 
**OtherSynonyms** | Pointer to [**[]V2reportsNameAndAuthority**](V2reportsNameAndAuthority.md) |  | [optional] 
**InformalNames** | Pointer to **[]string** |  | [optional] 
**Basionym** | Pointer to [**V2reportsNameAndAuthority**](V2reportsNameAndAuthority.md) |  | [optional] 
**Publications** | Pointer to [**[]V2reportsNameAndAuthorityPublication**](V2reportsNameAndAuthorityPublication.md) |  | [optional] 
**Notes** | Pointer to [**[]V2reportsNameAndAuthorityNote**](V2reportsNameAndAuthorityNote.md) |  | [optional] 
**Formal** | Pointer to **bool** |  | [optional] 

## Methods

### NewV2reportsNameAndAuthority

`func NewV2reportsNameAndAuthority() *V2reportsNameAndAuthority`

NewV2reportsNameAndAuthority instantiates a new V2reportsNameAndAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsNameAndAuthorityWithDefaults

`func NewV2reportsNameAndAuthorityWithDefaults() *V2reportsNameAndAuthority`

NewV2reportsNameAndAuthorityWithDefaults instantiates a new V2reportsNameAndAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2reportsNameAndAuthority) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2reportsNameAndAuthority) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2reportsNameAndAuthority) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2reportsNameAndAuthority) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthority

`func (o *V2reportsNameAndAuthority) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *V2reportsNameAndAuthority) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *V2reportsNameAndAuthority) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *V2reportsNameAndAuthority) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetTypeStrains

`func (o *V2reportsNameAndAuthority) GetTypeStrains() []V2reportsTaxonomyTypeMaterial`

GetTypeStrains returns the TypeStrains field if non-nil, zero value otherwise.

### GetTypeStrainsOk

`func (o *V2reportsNameAndAuthority) GetTypeStrainsOk() (*[]V2reportsTaxonomyTypeMaterial, bool)`

GetTypeStrainsOk returns a tuple with the TypeStrains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeStrains

`func (o *V2reportsNameAndAuthority) SetTypeStrains(v []V2reportsTaxonomyTypeMaterial)`

SetTypeStrains sets TypeStrains field to given value.

### HasTypeStrains

`func (o *V2reportsNameAndAuthority) HasTypeStrains() bool`

HasTypeStrains returns a boolean if a field has been set.

### GetCuratorSynonym

`func (o *V2reportsNameAndAuthority) GetCuratorSynonym() string`

GetCuratorSynonym returns the CuratorSynonym field if non-nil, zero value otherwise.

### GetCuratorSynonymOk

`func (o *V2reportsNameAndAuthority) GetCuratorSynonymOk() (*string, bool)`

GetCuratorSynonymOk returns a tuple with the CuratorSynonym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuratorSynonym

`func (o *V2reportsNameAndAuthority) SetCuratorSynonym(v string)`

SetCuratorSynonym sets CuratorSynonym field to given value.

### HasCuratorSynonym

`func (o *V2reportsNameAndAuthority) HasCuratorSynonym() bool`

HasCuratorSynonym returns a boolean if a field has been set.

### GetHomotypicSynonyms

`func (o *V2reportsNameAndAuthority) GetHomotypicSynonyms() []V2reportsNameAndAuthority`

GetHomotypicSynonyms returns the HomotypicSynonyms field if non-nil, zero value otherwise.

### GetHomotypicSynonymsOk

`func (o *V2reportsNameAndAuthority) GetHomotypicSynonymsOk() (*[]V2reportsNameAndAuthority, bool)`

GetHomotypicSynonymsOk returns a tuple with the HomotypicSynonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomotypicSynonyms

`func (o *V2reportsNameAndAuthority) SetHomotypicSynonyms(v []V2reportsNameAndAuthority)`

SetHomotypicSynonyms sets HomotypicSynonyms field to given value.

### HasHomotypicSynonyms

`func (o *V2reportsNameAndAuthority) HasHomotypicSynonyms() bool`

HasHomotypicSynonyms returns a boolean if a field has been set.

### GetHeterotypicSynonyms

`func (o *V2reportsNameAndAuthority) GetHeterotypicSynonyms() []V2reportsNameAndAuthority`

GetHeterotypicSynonyms returns the HeterotypicSynonyms field if non-nil, zero value otherwise.

### GetHeterotypicSynonymsOk

`func (o *V2reportsNameAndAuthority) GetHeterotypicSynonymsOk() (*[]V2reportsNameAndAuthority, bool)`

GetHeterotypicSynonymsOk returns a tuple with the HeterotypicSynonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeterotypicSynonyms

`func (o *V2reportsNameAndAuthority) SetHeterotypicSynonyms(v []V2reportsNameAndAuthority)`

SetHeterotypicSynonyms sets HeterotypicSynonyms field to given value.

### HasHeterotypicSynonyms

`func (o *V2reportsNameAndAuthority) HasHeterotypicSynonyms() bool`

HasHeterotypicSynonyms returns a boolean if a field has been set.

### GetOtherSynonyms

`func (o *V2reportsNameAndAuthority) GetOtherSynonyms() []V2reportsNameAndAuthority`

GetOtherSynonyms returns the OtherSynonyms field if non-nil, zero value otherwise.

### GetOtherSynonymsOk

`func (o *V2reportsNameAndAuthority) GetOtherSynonymsOk() (*[]V2reportsNameAndAuthority, bool)`

GetOtherSynonymsOk returns a tuple with the OtherSynonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSynonyms

`func (o *V2reportsNameAndAuthority) SetOtherSynonyms(v []V2reportsNameAndAuthority)`

SetOtherSynonyms sets OtherSynonyms field to given value.

### HasOtherSynonyms

`func (o *V2reportsNameAndAuthority) HasOtherSynonyms() bool`

HasOtherSynonyms returns a boolean if a field has been set.

### GetInformalNames

`func (o *V2reportsNameAndAuthority) GetInformalNames() []string`

GetInformalNames returns the InformalNames field if non-nil, zero value otherwise.

### GetInformalNamesOk

`func (o *V2reportsNameAndAuthority) GetInformalNamesOk() (*[]string, bool)`

GetInformalNamesOk returns a tuple with the InformalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformalNames

`func (o *V2reportsNameAndAuthority) SetInformalNames(v []string)`

SetInformalNames sets InformalNames field to given value.

### HasInformalNames

`func (o *V2reportsNameAndAuthority) HasInformalNames() bool`

HasInformalNames returns a boolean if a field has been set.

### GetBasionym

`func (o *V2reportsNameAndAuthority) GetBasionym() V2reportsNameAndAuthority`

GetBasionym returns the Basionym field if non-nil, zero value otherwise.

### GetBasionymOk

`func (o *V2reportsNameAndAuthority) GetBasionymOk() (*V2reportsNameAndAuthority, bool)`

GetBasionymOk returns a tuple with the Basionym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasionym

`func (o *V2reportsNameAndAuthority) SetBasionym(v V2reportsNameAndAuthority)`

SetBasionym sets Basionym field to given value.

### HasBasionym

`func (o *V2reportsNameAndAuthority) HasBasionym() bool`

HasBasionym returns a boolean if a field has been set.

### GetPublications

`func (o *V2reportsNameAndAuthority) GetPublications() []V2reportsNameAndAuthorityPublication`

GetPublications returns the Publications field if non-nil, zero value otherwise.

### GetPublicationsOk

`func (o *V2reportsNameAndAuthority) GetPublicationsOk() (*[]V2reportsNameAndAuthorityPublication, bool)`

GetPublicationsOk returns a tuple with the Publications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublications

`func (o *V2reportsNameAndAuthority) SetPublications(v []V2reportsNameAndAuthorityPublication)`

SetPublications sets Publications field to given value.

### HasPublications

`func (o *V2reportsNameAndAuthority) HasPublications() bool`

HasPublications returns a boolean if a field has been set.

### GetNotes

`func (o *V2reportsNameAndAuthority) GetNotes() []V2reportsNameAndAuthorityNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *V2reportsNameAndAuthority) GetNotesOk() (*[]V2reportsNameAndAuthorityNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *V2reportsNameAndAuthority) SetNotes(v []V2reportsNameAndAuthorityNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *V2reportsNameAndAuthority) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetFormal

`func (o *V2reportsNameAndAuthority) GetFormal() bool`

GetFormal returns the Formal field if non-nil, zero value otherwise.

### GetFormalOk

`func (o *V2reportsNameAndAuthority) GetFormalOk() (*bool, bool)`

GetFormalOk returns a tuple with the Formal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormal

`func (o *V2reportsNameAndAuthority) SetFormal(v bool)`

SetFormal sets Formal field to given value.

### HasFormal

`func (o *V2reportsNameAndAuthority) HasFormal() bool`

HasFormal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


