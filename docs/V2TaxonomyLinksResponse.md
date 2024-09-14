# V2TaxonomyLinksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxId** | Pointer to **string** |  | [optional] 
**EncyclopediaOfLife** | Pointer to **string** |  | [optional] 
**GlobalBiodiversityInformationFacility** | Pointer to **string** |  | [optional] 
**Inaturalist** | Pointer to **string** |  | [optional] 
**Viralzone** | Pointer to **string** |  | [optional] 
**Wikipedia** | Pointer to **string** |  | [optional] 
**GenericLinks** | Pointer to [**[]V2TaxonomyLinksResponseGenericLink**](V2TaxonomyLinksResponseGenericLink.md) |  | [optional] 

## Methods

### NewV2TaxonomyLinksResponse

`func NewV2TaxonomyLinksResponse() *V2TaxonomyLinksResponse`

NewV2TaxonomyLinksResponse instantiates a new V2TaxonomyLinksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TaxonomyLinksResponseWithDefaults

`func NewV2TaxonomyLinksResponseWithDefaults() *V2TaxonomyLinksResponse`

NewV2TaxonomyLinksResponseWithDefaults instantiates a new V2TaxonomyLinksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxId

`func (o *V2TaxonomyLinksResponse) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V2TaxonomyLinksResponse) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V2TaxonomyLinksResponse) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V2TaxonomyLinksResponse) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetEncyclopediaOfLife

`func (o *V2TaxonomyLinksResponse) GetEncyclopediaOfLife() string`

GetEncyclopediaOfLife returns the EncyclopediaOfLife field if non-nil, zero value otherwise.

### GetEncyclopediaOfLifeOk

`func (o *V2TaxonomyLinksResponse) GetEncyclopediaOfLifeOk() (*string, bool)`

GetEncyclopediaOfLifeOk returns a tuple with the EncyclopediaOfLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncyclopediaOfLife

`func (o *V2TaxonomyLinksResponse) SetEncyclopediaOfLife(v string)`

SetEncyclopediaOfLife sets EncyclopediaOfLife field to given value.

### HasEncyclopediaOfLife

`func (o *V2TaxonomyLinksResponse) HasEncyclopediaOfLife() bool`

HasEncyclopediaOfLife returns a boolean if a field has been set.

### GetGlobalBiodiversityInformationFacility

`func (o *V2TaxonomyLinksResponse) GetGlobalBiodiversityInformationFacility() string`

GetGlobalBiodiversityInformationFacility returns the GlobalBiodiversityInformationFacility field if non-nil, zero value otherwise.

### GetGlobalBiodiversityInformationFacilityOk

`func (o *V2TaxonomyLinksResponse) GetGlobalBiodiversityInformationFacilityOk() (*string, bool)`

GetGlobalBiodiversityInformationFacilityOk returns a tuple with the GlobalBiodiversityInformationFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalBiodiversityInformationFacility

`func (o *V2TaxonomyLinksResponse) SetGlobalBiodiversityInformationFacility(v string)`

SetGlobalBiodiversityInformationFacility sets GlobalBiodiversityInformationFacility field to given value.

### HasGlobalBiodiversityInformationFacility

`func (o *V2TaxonomyLinksResponse) HasGlobalBiodiversityInformationFacility() bool`

HasGlobalBiodiversityInformationFacility returns a boolean if a field has been set.

### GetInaturalist

`func (o *V2TaxonomyLinksResponse) GetInaturalist() string`

GetInaturalist returns the Inaturalist field if non-nil, zero value otherwise.

### GetInaturalistOk

`func (o *V2TaxonomyLinksResponse) GetInaturalistOk() (*string, bool)`

GetInaturalistOk returns a tuple with the Inaturalist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInaturalist

`func (o *V2TaxonomyLinksResponse) SetInaturalist(v string)`

SetInaturalist sets Inaturalist field to given value.

### HasInaturalist

`func (o *V2TaxonomyLinksResponse) HasInaturalist() bool`

HasInaturalist returns a boolean if a field has been set.

### GetViralzone

`func (o *V2TaxonomyLinksResponse) GetViralzone() string`

GetViralzone returns the Viralzone field if non-nil, zero value otherwise.

### GetViralzoneOk

`func (o *V2TaxonomyLinksResponse) GetViralzoneOk() (*string, bool)`

GetViralzoneOk returns a tuple with the Viralzone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViralzone

`func (o *V2TaxonomyLinksResponse) SetViralzone(v string)`

SetViralzone sets Viralzone field to given value.

### HasViralzone

`func (o *V2TaxonomyLinksResponse) HasViralzone() bool`

HasViralzone returns a boolean if a field has been set.

### GetWikipedia

`func (o *V2TaxonomyLinksResponse) GetWikipedia() string`

GetWikipedia returns the Wikipedia field if non-nil, zero value otherwise.

### GetWikipediaOk

`func (o *V2TaxonomyLinksResponse) GetWikipediaOk() (*string, bool)`

GetWikipediaOk returns a tuple with the Wikipedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikipedia

`func (o *V2TaxonomyLinksResponse) SetWikipedia(v string)`

SetWikipedia sets Wikipedia field to given value.

### HasWikipedia

`func (o *V2TaxonomyLinksResponse) HasWikipedia() bool`

HasWikipedia returns a boolean if a field has been set.

### GetGenericLinks

`func (o *V2TaxonomyLinksResponse) GetGenericLinks() []V2TaxonomyLinksResponseGenericLink`

GetGenericLinks returns the GenericLinks field if non-nil, zero value otherwise.

### GetGenericLinksOk

`func (o *V2TaxonomyLinksResponse) GetGenericLinksOk() (*[]V2TaxonomyLinksResponseGenericLink, bool)`

GetGenericLinksOk returns a tuple with the GenericLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericLinks

`func (o *V2TaxonomyLinksResponse) SetGenericLinks(v []V2TaxonomyLinksResponseGenericLink)`

SetGenericLinks sets GenericLinks field to given value.

### HasGenericLinks

`func (o *V2TaxonomyLinksResponse) HasGenericLinks() bool`

HasGenericLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


