# V2reportsProductDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TaxId** | Pointer to **string** |  | [optional] 
**Taxname** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**V2reportsGeneType**](V2reportsGeneType.md) |  | [optional] [default to V2REPORTSGENETYPE_UNKNOWN]
**RnaType** | Pointer to [**V2reportsRnaType**](V2reportsRnaType.md) |  | [optional] [default to V2REPORTSRNATYPE_RNA_UNKNOWN]
**Transcripts** | Pointer to [**[]V2reportsTranscript**](V2reportsTranscript.md) |  | [optional] 
**TranscriptCount** | Pointer to **int32** |  | [optional] 
**ProteinCount** | Pointer to **int32** |  | [optional] 
**TranscriptTypeCounts** | Pointer to [**[]V2reportsTranscriptTypeCount**](V2reportsTranscriptTypeCount.md) |  | [optional] 

## Methods

### NewV2reportsProductDescriptor

`func NewV2reportsProductDescriptor() *V2reportsProductDescriptor`

NewV2reportsProductDescriptor instantiates a new V2reportsProductDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsProductDescriptorWithDefaults

`func NewV2reportsProductDescriptorWithDefaults() *V2reportsProductDescriptor`

NewV2reportsProductDescriptorWithDefaults instantiates a new V2reportsProductDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneId

`func (o *V2reportsProductDescriptor) GetGeneId() string`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *V2reportsProductDescriptor) GetGeneIdOk() (*string, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *V2reportsProductDescriptor) SetGeneId(v string)`

SetGeneId sets GeneId field to given value.

### HasGeneId

`func (o *V2reportsProductDescriptor) HasGeneId() bool`

HasGeneId returns a boolean if a field has been set.

### GetSymbol

`func (o *V2reportsProductDescriptor) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *V2reportsProductDescriptor) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *V2reportsProductDescriptor) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *V2reportsProductDescriptor) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *V2reportsProductDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2reportsProductDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2reportsProductDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V2reportsProductDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTaxId

`func (o *V2reportsProductDescriptor) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V2reportsProductDescriptor) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V2reportsProductDescriptor) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V2reportsProductDescriptor) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxname

`func (o *V2reportsProductDescriptor) GetTaxname() string`

GetTaxname returns the Taxname field if non-nil, zero value otherwise.

### GetTaxnameOk

`func (o *V2reportsProductDescriptor) GetTaxnameOk() (*string, bool)`

GetTaxnameOk returns a tuple with the Taxname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxname

`func (o *V2reportsProductDescriptor) SetTaxname(v string)`

SetTaxname sets Taxname field to given value.

### HasTaxname

`func (o *V2reportsProductDescriptor) HasTaxname() bool`

HasTaxname returns a boolean if a field has been set.

### GetCommonName

`func (o *V2reportsProductDescriptor) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V2reportsProductDescriptor) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V2reportsProductDescriptor) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V2reportsProductDescriptor) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetType

`func (o *V2reportsProductDescriptor) GetType() V2reportsGeneType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2reportsProductDescriptor) GetTypeOk() (*V2reportsGeneType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2reportsProductDescriptor) SetType(v V2reportsGeneType)`

SetType sets Type field to given value.

### HasType

`func (o *V2reportsProductDescriptor) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRnaType

`func (o *V2reportsProductDescriptor) GetRnaType() V2reportsRnaType`

GetRnaType returns the RnaType field if non-nil, zero value otherwise.

### GetRnaTypeOk

`func (o *V2reportsProductDescriptor) GetRnaTypeOk() (*V2reportsRnaType, bool)`

GetRnaTypeOk returns a tuple with the RnaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRnaType

`func (o *V2reportsProductDescriptor) SetRnaType(v V2reportsRnaType)`

SetRnaType sets RnaType field to given value.

### HasRnaType

`func (o *V2reportsProductDescriptor) HasRnaType() bool`

HasRnaType returns a boolean if a field has been set.

### GetTranscripts

`func (o *V2reportsProductDescriptor) GetTranscripts() []V2reportsTranscript`

GetTranscripts returns the Transcripts field if non-nil, zero value otherwise.

### GetTranscriptsOk

`func (o *V2reportsProductDescriptor) GetTranscriptsOk() (*[]V2reportsTranscript, bool)`

GetTranscriptsOk returns a tuple with the Transcripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscripts

`func (o *V2reportsProductDescriptor) SetTranscripts(v []V2reportsTranscript)`

SetTranscripts sets Transcripts field to given value.

### HasTranscripts

`func (o *V2reportsProductDescriptor) HasTranscripts() bool`

HasTranscripts returns a boolean if a field has been set.

### GetTranscriptCount

`func (o *V2reportsProductDescriptor) GetTranscriptCount() int32`

GetTranscriptCount returns the TranscriptCount field if non-nil, zero value otherwise.

### GetTranscriptCountOk

`func (o *V2reportsProductDescriptor) GetTranscriptCountOk() (*int32, bool)`

GetTranscriptCountOk returns a tuple with the TranscriptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptCount

`func (o *V2reportsProductDescriptor) SetTranscriptCount(v int32)`

SetTranscriptCount sets TranscriptCount field to given value.

### HasTranscriptCount

`func (o *V2reportsProductDescriptor) HasTranscriptCount() bool`

HasTranscriptCount returns a boolean if a field has been set.

### GetProteinCount

`func (o *V2reportsProductDescriptor) GetProteinCount() int32`

GetProteinCount returns the ProteinCount field if non-nil, zero value otherwise.

### GetProteinCountOk

`func (o *V2reportsProductDescriptor) GetProteinCountOk() (*int32, bool)`

GetProteinCountOk returns a tuple with the ProteinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinCount

`func (o *V2reportsProductDescriptor) SetProteinCount(v int32)`

SetProteinCount sets ProteinCount field to given value.

### HasProteinCount

`func (o *V2reportsProductDescriptor) HasProteinCount() bool`

HasProteinCount returns a boolean if a field has been set.

### GetTranscriptTypeCounts

`func (o *V2reportsProductDescriptor) GetTranscriptTypeCounts() []V2reportsTranscriptTypeCount`

GetTranscriptTypeCounts returns the TranscriptTypeCounts field if non-nil, zero value otherwise.

### GetTranscriptTypeCountsOk

`func (o *V2reportsProductDescriptor) GetTranscriptTypeCountsOk() (*[]V2reportsTranscriptTypeCount, bool)`

GetTranscriptTypeCountsOk returns a tuple with the TranscriptTypeCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptTypeCounts

`func (o *V2reportsProductDescriptor) SetTranscriptTypeCounts(v []V2reportsTranscriptTypeCount)`

SetTranscriptTypeCounts sets TranscriptTypeCounts field to given value.

### HasTranscriptTypeCounts

`func (o *V2reportsProductDescriptor) HasTranscriptTypeCounts() bool`

HasTranscriptTypeCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


