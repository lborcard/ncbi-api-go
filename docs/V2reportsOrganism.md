# V2reportsOrganism

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxId** | Pointer to **int32** |  | [optional] 
**SciName** | Pointer to **string** |  | [optional] 
**OrganismName** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Lineage** | Pointer to [**[]V2reportsLineageOrganism**](V2reportsLineageOrganism.md) |  | [optional] 
**Strain** | Pointer to **string** |  | [optional] 
**PangolinClassification** | Pointer to **string** |  | [optional] 
**InfraspecificNames** | Pointer to [**V2reportsInfraspecificNames**](V2reportsInfraspecificNames.md) |  | [optional] 

## Methods

### NewV2reportsOrganism

`func NewV2reportsOrganism() *V2reportsOrganism`

NewV2reportsOrganism instantiates a new V2reportsOrganism object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2reportsOrganismWithDefaults

`func NewV2reportsOrganismWithDefaults() *V2reportsOrganism`

NewV2reportsOrganismWithDefaults instantiates a new V2reportsOrganism object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxId

`func (o *V2reportsOrganism) GetTaxId() int32`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V2reportsOrganism) GetTaxIdOk() (*int32, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V2reportsOrganism) SetTaxId(v int32)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V2reportsOrganism) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetSciName

`func (o *V2reportsOrganism) GetSciName() string`

GetSciName returns the SciName field if non-nil, zero value otherwise.

### GetSciNameOk

`func (o *V2reportsOrganism) GetSciNameOk() (*string, bool)`

GetSciNameOk returns a tuple with the SciName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSciName

`func (o *V2reportsOrganism) SetSciName(v string)`

SetSciName sets SciName field to given value.

### HasSciName

`func (o *V2reportsOrganism) HasSciName() bool`

HasSciName returns a boolean if a field has been set.

### GetOrganismName

`func (o *V2reportsOrganism) GetOrganismName() string`

GetOrganismName returns the OrganismName field if non-nil, zero value otherwise.

### GetOrganismNameOk

`func (o *V2reportsOrganism) GetOrganismNameOk() (*string, bool)`

GetOrganismNameOk returns a tuple with the OrganismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganismName

`func (o *V2reportsOrganism) SetOrganismName(v string)`

SetOrganismName sets OrganismName field to given value.

### HasOrganismName

`func (o *V2reportsOrganism) HasOrganismName() bool`

HasOrganismName returns a boolean if a field has been set.

### GetCommonName

`func (o *V2reportsOrganism) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V2reportsOrganism) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V2reportsOrganism) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V2reportsOrganism) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetLineage

`func (o *V2reportsOrganism) GetLineage() []V2reportsLineageOrganism`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *V2reportsOrganism) GetLineageOk() (*[]V2reportsLineageOrganism, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *V2reportsOrganism) SetLineage(v []V2reportsLineageOrganism)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *V2reportsOrganism) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetStrain

`func (o *V2reportsOrganism) GetStrain() string`

GetStrain returns the Strain field if non-nil, zero value otherwise.

### GetStrainOk

`func (o *V2reportsOrganism) GetStrainOk() (*string, bool)`

GetStrainOk returns a tuple with the Strain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrain

`func (o *V2reportsOrganism) SetStrain(v string)`

SetStrain sets Strain field to given value.

### HasStrain

`func (o *V2reportsOrganism) HasStrain() bool`

HasStrain returns a boolean if a field has been set.

### GetPangolinClassification

`func (o *V2reportsOrganism) GetPangolinClassification() string`

GetPangolinClassification returns the PangolinClassification field if non-nil, zero value otherwise.

### GetPangolinClassificationOk

`func (o *V2reportsOrganism) GetPangolinClassificationOk() (*string, bool)`

GetPangolinClassificationOk returns a tuple with the PangolinClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPangolinClassification

`func (o *V2reportsOrganism) SetPangolinClassification(v string)`

SetPangolinClassification sets PangolinClassification field to given value.

### HasPangolinClassification

`func (o *V2reportsOrganism) HasPangolinClassification() bool`

HasPangolinClassification returns a boolean if a field has been set.

### GetInfraspecificNames

`func (o *V2reportsOrganism) GetInfraspecificNames() V2reportsInfraspecificNames`

GetInfraspecificNames returns the InfraspecificNames field if non-nil, zero value otherwise.

### GetInfraspecificNamesOk

`func (o *V2reportsOrganism) GetInfraspecificNamesOk() (*V2reportsInfraspecificNames, bool)`

GetInfraspecificNamesOk returns a tuple with the InfraspecificNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraspecificNames

`func (o *V2reportsOrganism) SetInfraspecificNames(v V2reportsInfraspecificNames)`

SetInfraspecificNames sets InfraspecificNames field to given value.

### HasInfraspecificNames

`func (o *V2reportsOrganism) HasInfraspecificNames() bool`

HasInfraspecificNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


