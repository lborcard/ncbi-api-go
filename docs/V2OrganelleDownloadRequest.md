# V2OrganelleDownloadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessions** | Pointer to **[]string** |  | [optional] 
**ExcludeSequence** | Pointer to **bool** |  | [optional] 
**IncludeAnnotationType** | Pointer to [**[]V2AnnotationForOrganelleType**](V2AnnotationForOrganelleType.md) |  | [optional] 

## Methods

### NewV2OrganelleDownloadRequest

`func NewV2OrganelleDownloadRequest() *V2OrganelleDownloadRequest`

NewV2OrganelleDownloadRequest instantiates a new V2OrganelleDownloadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2OrganelleDownloadRequestWithDefaults

`func NewV2OrganelleDownloadRequestWithDefaults() *V2OrganelleDownloadRequest`

NewV2OrganelleDownloadRequestWithDefaults instantiates a new V2OrganelleDownloadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessions

`func (o *V2OrganelleDownloadRequest) GetAccessions() []string`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V2OrganelleDownloadRequest) GetAccessionsOk() (*[]string, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V2OrganelleDownloadRequest) SetAccessions(v []string)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V2OrganelleDownloadRequest) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetExcludeSequence

`func (o *V2OrganelleDownloadRequest) GetExcludeSequence() bool`

GetExcludeSequence returns the ExcludeSequence field if non-nil, zero value otherwise.

### GetExcludeSequenceOk

`func (o *V2OrganelleDownloadRequest) GetExcludeSequenceOk() (*bool, bool)`

GetExcludeSequenceOk returns a tuple with the ExcludeSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSequence

`func (o *V2OrganelleDownloadRequest) SetExcludeSequence(v bool)`

SetExcludeSequence sets ExcludeSequence field to given value.

### HasExcludeSequence

`func (o *V2OrganelleDownloadRequest) HasExcludeSequence() bool`

HasExcludeSequence returns a boolean if a field has been set.

### GetIncludeAnnotationType

`func (o *V2OrganelleDownloadRequest) GetIncludeAnnotationType() []V2AnnotationForOrganelleType`

GetIncludeAnnotationType returns the IncludeAnnotationType field if non-nil, zero value otherwise.

### GetIncludeAnnotationTypeOk

`func (o *V2OrganelleDownloadRequest) GetIncludeAnnotationTypeOk() (*[]V2AnnotationForOrganelleType, bool)`

GetIncludeAnnotationTypeOk returns a tuple with the IncludeAnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnnotationType

`func (o *V2OrganelleDownloadRequest) SetIncludeAnnotationType(v []V2AnnotationForOrganelleType)`

SetIncludeAnnotationType sets IncludeAnnotationType field to given value.

### HasIncludeAnnotationType

`func (o *V2OrganelleDownloadRequest) HasIncludeAnnotationType() bool`

HasIncludeAnnotationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


