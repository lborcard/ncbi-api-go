# V2reportsAssemblyDataReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accession** | Pointer to **string** |  | [optional] 
**CurrentAccession** | Pointer to **string** |  | [optional] 
**PairedAccession** | Pointer to **string** |  | [optional] 
**SourceDatabase** | Pointer to [**V2reportsSourceDatabase**](V2reportsSourceDatabase.md) |  | [optional] [default to V2REPORTSSOURCEDATABASE_UNSPECIFIED]
**Organism** | Pointer to [**V2reportsOrganism**](V2reportsOrganism.md) |  | [optional] 
**AssemblyInfo** | Pointer to [**V2reportsAssemblyInfo**](V2reportsAssemblyInfo.md) |  | [optional] 
**AssemblyStats** | Pointer to [**V2reportsAssemblyStats**](V2reportsAssemblyStats.md) |  | [optional] 
**OrganelleInfo** | Pointer to [**[]V2reportsOrganelleInfo**](V2reportsOrganelleInfo.md) |  | [optional] 
**AnnotationInfo** | Pointer to [**V2reportsAnnotationInfo**](V2reportsAnnotationInfo.md) |  | [optional] 
**WgsInfo** | Pointer to [**V2reportsWGSInfo**](V2reportsWGSInfo.md) |  | [optional] 
**TypeMaterial** | Pointer to [**V2reportsTypeMaterial**](V2reportsTypeMaterial.md) |  | [optional] 
**CheckmInfo** | Pointer to [**V2reportsCheckM**](V2reportsCheckM.md) |  | [optional] 
**AverageNucleotideIdentity** | Pointer to [**V2reportsAverageNucleotideIdentity**](V2reportsAverageNucleotideIdentity.md) |  | [optional] 

## Methods

### NewV2reportsAssemblyDataReport

`func NewV2reportsAssemblyDataReport() *V2reportsAssemblyDataReport`

NewV2reportsAssemblyDataReport instantiates a new V2reportsAssemblyDataReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsAssemblyDataReportWithDefaults

`func NewV2reportsAssemblyDataReportWithDefaults() *V2reportsAssemblyDataReport`

NewV2reportsAssemblyDataReportWithDefaults instantiates a new V2reportsAssemblyDataReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccession

`func (o *V2reportsAssemblyDataReport) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V2reportsAssemblyDataReport) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V2reportsAssemblyDataReport) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V2reportsAssemblyDataReport) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetCurrentAccession

`func (o *V2reportsAssemblyDataReport) GetCurrentAccession() string`

GetCurrentAccession returns the CurrentAccession field if non-nil, zero value otherwise.

### GetCurrentAccessionOk

`func (o *V2reportsAssemblyDataReport) GetCurrentAccessionOk() (*string, bool)`

GetCurrentAccessionOk returns a tuple with the CurrentAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAccession

`func (o *V2reportsAssemblyDataReport) SetCurrentAccession(v string)`

SetCurrentAccession sets CurrentAccession field to given value.

### HasCurrentAccession

`func (o *V2reportsAssemblyDataReport) HasCurrentAccession() bool`

HasCurrentAccession returns a boolean if a field has been set.

### GetPairedAccession

`func (o *V2reportsAssemblyDataReport) GetPairedAccession() string`

GetPairedAccession returns the PairedAccession field if non-nil, zero value otherwise.

### GetPairedAccessionOk

`func (o *V2reportsAssemblyDataReport) GetPairedAccessionOk() (*string, bool)`

GetPairedAccessionOk returns a tuple with the PairedAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairedAccession

`func (o *V2reportsAssemblyDataReport) SetPairedAccession(v string)`

SetPairedAccession sets PairedAccession field to given value.

### HasPairedAccession

`func (o *V2reportsAssemblyDataReport) HasPairedAccession() bool`

HasPairedAccession returns a boolean if a field has been set.

### GetSourceDatabase

`func (o *V2reportsAssemblyDataReport) GetSourceDatabase() V2reportsSourceDatabase`

GetSourceDatabase returns the SourceDatabase field if non-nil, zero value otherwise.

### GetSourceDatabaseOk

`func (o *V2reportsAssemblyDataReport) GetSourceDatabaseOk() (*V2reportsSourceDatabase, bool)`

GetSourceDatabaseOk returns a tuple with the SourceDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatabase

`func (o *V2reportsAssemblyDataReport) SetSourceDatabase(v V2reportsSourceDatabase)`

SetSourceDatabase sets SourceDatabase field to given value.

### HasSourceDatabase

`func (o *V2reportsAssemblyDataReport) HasSourceDatabase() bool`

HasSourceDatabase returns a boolean if a field has been set.

### GetOrganism

`func (o *V2reportsAssemblyDataReport) GetOrganism() V2reportsOrganism`

GetOrganism returns the Organism field if non-nil, zero value otherwise.

### GetOrganismOk

`func (o *V2reportsAssemblyDataReport) GetOrganismOk() (*V2reportsOrganism, bool)`

GetOrganismOk returns a tuple with the Organism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganism

`func (o *V2reportsAssemblyDataReport) SetOrganism(v V2reportsOrganism)`

SetOrganism sets Organism field to given value.

### HasOrganism

`func (o *V2reportsAssemblyDataReport) HasOrganism() bool`

HasOrganism returns a boolean if a field has been set.

### GetAssemblyInfo

`func (o *V2reportsAssemblyDataReport) GetAssemblyInfo() V2reportsAssemblyInfo`

GetAssemblyInfo returns the AssemblyInfo field if non-nil, zero value otherwise.

### GetAssemblyInfoOk

`func (o *V2reportsAssemblyDataReport) GetAssemblyInfoOk() (*V2reportsAssemblyInfo, bool)`

GetAssemblyInfoOk returns a tuple with the AssemblyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyInfo

`func (o *V2reportsAssemblyDataReport) SetAssemblyInfo(v V2reportsAssemblyInfo)`

SetAssemblyInfo sets AssemblyInfo field to given value.

### HasAssemblyInfo

`func (o *V2reportsAssemblyDataReport) HasAssemblyInfo() bool`

HasAssemblyInfo returns a boolean if a field has been set.

### GetAssemblyStats

`func (o *V2reportsAssemblyDataReport) GetAssemblyStats() V2reportsAssemblyStats`

GetAssemblyStats returns the AssemblyStats field if non-nil, zero value otherwise.

### GetAssemblyStatsOk

`func (o *V2reportsAssemblyDataReport) GetAssemblyStatsOk() (*V2reportsAssemblyStats, bool)`

GetAssemblyStatsOk returns a tuple with the AssemblyStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyStats

`func (o *V2reportsAssemblyDataReport) SetAssemblyStats(v V2reportsAssemblyStats)`

SetAssemblyStats sets AssemblyStats field to given value.

### HasAssemblyStats

`func (o *V2reportsAssemblyDataReport) HasAssemblyStats() bool`

HasAssemblyStats returns a boolean if a field has been set.

### GetOrganelleInfo

`func (o *V2reportsAssemblyDataReport) GetOrganelleInfo() []V2reportsOrganelleInfo`

GetOrganelleInfo returns the OrganelleInfo field if non-nil, zero value otherwise.

### GetOrganelleInfoOk

`func (o *V2reportsAssemblyDataReport) GetOrganelleInfoOk() (*[]V2reportsOrganelleInfo, bool)`

GetOrganelleInfoOk returns a tuple with the OrganelleInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganelleInfo

`func (o *V2reportsAssemblyDataReport) SetOrganelleInfo(v []V2reportsOrganelleInfo)`

SetOrganelleInfo sets OrganelleInfo field to given value.

### HasOrganelleInfo

`func (o *V2reportsAssemblyDataReport) HasOrganelleInfo() bool`

HasOrganelleInfo returns a boolean if a field has been set.

### GetAnnotationInfo

`func (o *V2reportsAssemblyDataReport) GetAnnotationInfo() V2reportsAnnotationInfo`

GetAnnotationInfo returns the AnnotationInfo field if non-nil, zero value otherwise.

### GetAnnotationInfoOk

`func (o *V2reportsAssemblyDataReport) GetAnnotationInfoOk() (*V2reportsAnnotationInfo, bool)`

GetAnnotationInfoOk returns a tuple with the AnnotationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationInfo

`func (o *V2reportsAssemblyDataReport) SetAnnotationInfo(v V2reportsAnnotationInfo)`

SetAnnotationInfo sets AnnotationInfo field to given value.

### HasAnnotationInfo

`func (o *V2reportsAssemblyDataReport) HasAnnotationInfo() bool`

HasAnnotationInfo returns a boolean if a field has been set.

### GetWgsInfo

`func (o *V2reportsAssemblyDataReport) GetWgsInfo() V2reportsWGSInfo`

GetWgsInfo returns the WgsInfo field if non-nil, zero value otherwise.

### GetWgsInfoOk

`func (o *V2reportsAssemblyDataReport) GetWgsInfoOk() (*V2reportsWGSInfo, bool)`

GetWgsInfoOk returns a tuple with the WgsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWgsInfo

`func (o *V2reportsAssemblyDataReport) SetWgsInfo(v V2reportsWGSInfo)`

SetWgsInfo sets WgsInfo field to given value.

### HasWgsInfo

`func (o *V2reportsAssemblyDataReport) HasWgsInfo() bool`

HasWgsInfo returns a boolean if a field has been set.

### GetTypeMaterial

`func (o *V2reportsAssemblyDataReport) GetTypeMaterial() V2reportsTypeMaterial`

GetTypeMaterial returns the TypeMaterial field if non-nil, zero value otherwise.

### GetTypeMaterialOk

`func (o *V2reportsAssemblyDataReport) GetTypeMaterialOk() (*V2reportsTypeMaterial, bool)`

GetTypeMaterialOk returns a tuple with the TypeMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeMaterial

`func (o *V2reportsAssemblyDataReport) SetTypeMaterial(v V2reportsTypeMaterial)`

SetTypeMaterial sets TypeMaterial field to given value.

### HasTypeMaterial

`func (o *V2reportsAssemblyDataReport) HasTypeMaterial() bool`

HasTypeMaterial returns a boolean if a field has been set.

### GetCheckmInfo

`func (o *V2reportsAssemblyDataReport) GetCheckmInfo() V2reportsCheckM`

GetCheckmInfo returns the CheckmInfo field if non-nil, zero value otherwise.

### GetCheckmInfoOk

`func (o *V2reportsAssemblyDataReport) GetCheckmInfoOk() (*V2reportsCheckM, bool)`

GetCheckmInfoOk returns a tuple with the CheckmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckmInfo

`func (o *V2reportsAssemblyDataReport) SetCheckmInfo(v V2reportsCheckM)`

SetCheckmInfo sets CheckmInfo field to given value.

### HasCheckmInfo

`func (o *V2reportsAssemblyDataReport) HasCheckmInfo() bool`

HasCheckmInfo returns a boolean if a field has been set.

### GetAverageNucleotideIdentity

`func (o *V2reportsAssemblyDataReport) GetAverageNucleotideIdentity() V2reportsAverageNucleotideIdentity`

GetAverageNucleotideIdentity returns the AverageNucleotideIdentity field if non-nil, zero value otherwise.

### GetAverageNucleotideIdentityOk

`func (o *V2reportsAssemblyDataReport) GetAverageNucleotideIdentityOk() (*V2reportsAverageNucleotideIdentity, bool)`

GetAverageNucleotideIdentityOk returns a tuple with the AverageNucleotideIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageNucleotideIdentity

`func (o *V2reportsAssemblyDataReport) SetAverageNucleotideIdentity(v V2reportsAverageNucleotideIdentity)`

SetAverageNucleotideIdentity sets AverageNucleotideIdentity field to given value.

### HasAverageNucleotideIdentity

`func (o *V2reportsAssemblyDataReport) HasAverageNucleotideIdentity() bool`

HasAverageNucleotideIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


