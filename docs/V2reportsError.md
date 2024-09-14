# V2reportsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyErrorCode** | Pointer to [**V2reportsErrorAssemblyErrorCode**](V2reportsErrorAssemblyErrorCode.md) |  | [optional] [default to V2REPORTSERRORASSEMBLYERRORCODE_UNKNOWN_ASSEMBLY_ERROR_CODE]
**GeneErrorCode** | Pointer to [**V2reportsErrorGeneErrorCode**](V2reportsErrorGeneErrorCode.md) |  | [optional] [default to V2REPORTSERRORGENEERRORCODE_UNKNOWN_GENE_ERROR_CODE]
**OrganelleErrorCode** | Pointer to [**V2reportsErrorOrganelleErrorCode**](V2reportsErrorOrganelleErrorCode.md) |  | [optional] [default to V2REPORTSERRORORGANELLEERRORCODE_UNKNOWN_ORGANELLE_ERROR_CODE]
**VirusErrorCode** | Pointer to [**V2reportsErrorVirusErrorCode**](V2reportsErrorVirusErrorCode.md) |  | [optional] [default to V2REPORTSERRORVIRUSERRORCODE_UNKNOWN_VIRUS_ERROR_CODE]
**TaxonomyErrorCode** | Pointer to [**V2reportsErrorTaxonomyErrorCode**](V2reportsErrorTaxonomyErrorCode.md) |  | [optional] [default to V2REPORTSERRORTAXONOMYERRORCODE_UNKNOWN_TAXONOMY_ERROR_CODE]
**Reason** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**InvalidIdentifiers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV2reportsError

`func NewV2reportsError() *V2reportsError`

NewV2reportsError instantiates a new V2reportsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsErrorWithDefaults

`func NewV2reportsErrorWithDefaults() *V2reportsError`

NewV2reportsErrorWithDefaults instantiates a new V2reportsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssemblyErrorCode

`func (o *V2reportsError) GetAssemblyErrorCode() V2reportsErrorAssemblyErrorCode`

GetAssemblyErrorCode returns the AssemblyErrorCode field if non-nil, zero value otherwise.

### GetAssemblyErrorCodeOk

`func (o *V2reportsError) GetAssemblyErrorCodeOk() (*V2reportsErrorAssemblyErrorCode, bool)`

GetAssemblyErrorCodeOk returns a tuple with the AssemblyErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyErrorCode

`func (o *V2reportsError) SetAssemblyErrorCode(v V2reportsErrorAssemblyErrorCode)`

SetAssemblyErrorCode sets AssemblyErrorCode field to given value.

### HasAssemblyErrorCode

`func (o *V2reportsError) HasAssemblyErrorCode() bool`

HasAssemblyErrorCode returns a boolean if a field has been set.

### GetGeneErrorCode

`func (o *V2reportsError) GetGeneErrorCode() V2reportsErrorGeneErrorCode`

GetGeneErrorCode returns the GeneErrorCode field if non-nil, zero value otherwise.

### GetGeneErrorCodeOk

`func (o *V2reportsError) GetGeneErrorCodeOk() (*V2reportsErrorGeneErrorCode, bool)`

GetGeneErrorCodeOk returns a tuple with the GeneErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneErrorCode

`func (o *V2reportsError) SetGeneErrorCode(v V2reportsErrorGeneErrorCode)`

SetGeneErrorCode sets GeneErrorCode field to given value.

### HasGeneErrorCode

`func (o *V2reportsError) HasGeneErrorCode() bool`

HasGeneErrorCode returns a boolean if a field has been set.

### GetOrganelleErrorCode

`func (o *V2reportsError) GetOrganelleErrorCode() V2reportsErrorOrganelleErrorCode`

GetOrganelleErrorCode returns the OrganelleErrorCode field if non-nil, zero value otherwise.

### GetOrganelleErrorCodeOk

`func (o *V2reportsError) GetOrganelleErrorCodeOk() (*V2reportsErrorOrganelleErrorCode, bool)`

GetOrganelleErrorCodeOk returns a tuple with the OrganelleErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganelleErrorCode

`func (o *V2reportsError) SetOrganelleErrorCode(v V2reportsErrorOrganelleErrorCode)`

SetOrganelleErrorCode sets OrganelleErrorCode field to given value.

### HasOrganelleErrorCode

`func (o *V2reportsError) HasOrganelleErrorCode() bool`

HasOrganelleErrorCode returns a boolean if a field has been set.

### GetVirusErrorCode

`func (o *V2reportsError) GetVirusErrorCode() V2reportsErrorVirusErrorCode`

GetVirusErrorCode returns the VirusErrorCode field if non-nil, zero value otherwise.

### GetVirusErrorCodeOk

`func (o *V2reportsError) GetVirusErrorCodeOk() (*V2reportsErrorVirusErrorCode, bool)`

GetVirusErrorCodeOk returns a tuple with the VirusErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirusErrorCode

`func (o *V2reportsError) SetVirusErrorCode(v V2reportsErrorVirusErrorCode)`

SetVirusErrorCode sets VirusErrorCode field to given value.

### HasVirusErrorCode

`func (o *V2reportsError) HasVirusErrorCode() bool`

HasVirusErrorCode returns a boolean if a field has been set.

### GetTaxonomyErrorCode

`func (o *V2reportsError) GetTaxonomyErrorCode() V2reportsErrorTaxonomyErrorCode`

GetTaxonomyErrorCode returns the TaxonomyErrorCode field if non-nil, zero value otherwise.

### GetTaxonomyErrorCodeOk

`func (o *V2reportsError) GetTaxonomyErrorCodeOk() (*V2reportsErrorTaxonomyErrorCode, bool)`

GetTaxonomyErrorCodeOk returns a tuple with the TaxonomyErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyErrorCode

`func (o *V2reportsError) SetTaxonomyErrorCode(v V2reportsErrorTaxonomyErrorCode)`

SetTaxonomyErrorCode sets TaxonomyErrorCode field to given value.

### HasTaxonomyErrorCode

`func (o *V2reportsError) HasTaxonomyErrorCode() bool`

HasTaxonomyErrorCode returns a boolean if a field has been set.

### GetReason

`func (o *V2reportsError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V2reportsError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V2reportsError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V2reportsError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *V2reportsError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2reportsError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2reportsError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V2reportsError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetInvalidIdentifiers

`func (o *V2reportsError) GetInvalidIdentifiers() []string`

GetInvalidIdentifiers returns the InvalidIdentifiers field if non-nil, zero value otherwise.

### GetInvalidIdentifiersOk

`func (o *V2reportsError) GetInvalidIdentifiersOk() (*[]string, bool)`

GetInvalidIdentifiersOk returns a tuple with the InvalidIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidIdentifiers

`func (o *V2reportsError) SetInvalidIdentifiers(v []string)`

SetInvalidIdentifiers sets InvalidIdentifiers field to given value.

### HasInvalidIdentifiers

`func (o *V2reportsError) HasInvalidIdentifiers() bool`

HasInvalidIdentifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


