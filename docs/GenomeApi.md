# \GenomeApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v2alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnnotationReportFacets**](GenomeApi.md#AnnotationReportFacets) | **Get** /genome/accession/{accession}/id/{annotation_ids}/annotation_summary | Get genome annotation report summary information
[**AnnotationReportFacetsByPost**](GenomeApi.md#AnnotationReportFacetsByPost) | **Post** /genome/annotation_summary | Get genome annotation report summary information
[**AssemblyAccessionsForSequenceAccession**](GenomeApi.md#AssemblyAccessionsForSequenceAccession) | **Get** /genome/sequence_accession/{accession}/sequence_assemblies | Get assembly accessions for a sequence accession
[**AssemblyAccessionsForSequenceAccessionByPost**](GenomeApi.md#AssemblyAccessionsForSequenceAccessionByPost) | **Post** /genome/sequence_assemblies | Get assembly accessions for a sequence accession
[**AssemblyRevisionHistoryByGet**](GenomeApi.md#AssemblyRevisionHistoryByGet) | **Get** /genome/accession/{accession}/revision_history | Get revision history for assembly by accession
[**AssemblyRevisionHistoryByPost**](GenomeApi.md#AssemblyRevisionHistoryByPost) | **Post** /genome/revision_history | Get revision history for assembly by accession
[**CheckAssemblyAvailability**](GenomeApi.md#CheckAssemblyAvailability) | **Get** /genome/accession/{accessions}/check | Check the validity of genome accessions
[**CheckAssemblyAvailabilityPost**](GenomeApi.md#CheckAssemblyAvailabilityPost) | **Post** /genome/check | Check the validity of many genome accessions in a single request
[**CheckmHistogramByTaxon**](GenomeApi.md#CheckmHistogramByTaxon) | **Get** /genome/taxon/{species_taxon}/checkm_histogram | Get CheckM histogram by species taxon
[**CheckmHistogramByTaxonByPost**](GenomeApi.md#CheckmHistogramByTaxonByPost) | **Post** /genome/checkm_histogram | Get CheckM histogram by species taxon
[**DownloadAssemblyPackage**](GenomeApi.md#DownloadAssemblyPackage) | **Get** /genome/accession/{accessions}/download | Get a genome dataset by accession
[**DownloadAssemblyPackagePost**](GenomeApi.md#DownloadAssemblyPackagePost) | **Post** /genome/download | Get a genome dataset by post
[**DownloadGenomeAnnotationPackage**](GenomeApi.md#DownloadGenomeAnnotationPackage) | **Get** /genome/accession/{accession}/annotation_report/download | Get an annotation report dataset by accession
[**DownloadGenomeAnnotationPackageByPost**](GenomeApi.md#DownloadGenomeAnnotationPackageByPost) | **Post** /genome/annotation_report/download | Get an annotation report dataset by accession
[**GenomeAnnotationDownloadSummary**](GenomeApi.md#GenomeAnnotationDownloadSummary) | **Get** /genome/accession/{accession}/annotation_report/download_summary | Preview feature dataset download
[**GenomeAnnotationDownloadSummaryByPost**](GenomeApi.md#GenomeAnnotationDownloadSummaryByPost) | **Post** /genome/annotation_report/download_summary | Preview feature download by POST
[**GenomeAnnotationReport**](GenomeApi.md#GenomeAnnotationReport) | **Get** /genome/accession/{accession}/annotation_report | Get genome annotation reports by genome accession
[**GenomeAnnotationReportByPost**](GenomeApi.md#GenomeAnnotationReportByPost) | **Post** /genome/annotation_report | Get genome annotation reports by genome accession
[**GenomeDatasetReport**](GenomeApi.md#GenomeDatasetReport) | **Get** /genome/accession/{accessions}/dataset_report | Get dataset reports by accessions
[**GenomeDatasetReportByPost**](GenomeApi.md#GenomeDatasetReportByPost) | **Post** /genome/dataset_report | Get dataset reports by accessions
[**GenomeDatasetReportsByAssemblyName**](GenomeApi.md#GenomeDatasetReportsByAssemblyName) | **Get** /genome/assembly_name/{assembly_names}/dataset_report | Get dataset reports by assembly name (exact)
[**GenomeDatasetReportsByBioproject**](GenomeApi.md#GenomeDatasetReportsByBioproject) | **Get** /genome/bioproject/{bioprojects}/dataset_report | Get dataset reports by bioproject
[**GenomeDatasetReportsByBiosampleId**](GenomeApi.md#GenomeDatasetReportsByBiosampleId) | **Get** /genome/biosample/{biosample_ids}/dataset_report | Get dataset reports by biosample id
[**GenomeDatasetReportsByTaxon**](GenomeApi.md#GenomeDatasetReportsByTaxon) | **Get** /genome/taxon/{taxons}/dataset_report | Get dataset reports by taxons
[**GenomeDatasetReportsByWgs**](GenomeApi.md#GenomeDatasetReportsByWgs) | **Get** /genome/wgs/{wgs_accessions}/dataset_report | Get dataset reports by wgs accession
[**GenomeDownloadSummary**](GenomeApi.md#GenomeDownloadSummary) | **Get** /genome/accession/{accessions}/download_summary | Preview genome dataset download
[**GenomeDownloadSummaryByPost**](GenomeApi.md#GenomeDownloadSummaryByPost) | **Post** /genome/download_summary | Preview genome dataset download by POST
[**GenomeLinksByAccession**](GenomeApi.md#GenomeLinksByAccession) | **Get** /genome/accession/{accessions}/links | Get assembly links by accessions
[**GenomeLinksByAccessionByPost**](GenomeApi.md#GenomeLinksByAccessionByPost) | **Post** /genome/links | Get assembly links by accessions
[**GenomeSequenceReport**](GenomeApi.md#GenomeSequenceReport) | **Get** /genome/accession/{accession}/sequence_reports | Get sequence reports by accessions
[**GenomeSequenceReportByPost**](GenomeApi.md#GenomeSequenceReportByPost) | **Post** /genome/sequence_reports | Get sequence reports by accessions



## AnnotationReportFacets

> V2GenomeAnnotationTableSummaryReply AnnotationReportFacets(ctx, accession, annotationIds).SortField(sortField).SortDirection(sortDirection).Execute()

Get genome annotation report summary information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accession := "GCF_000001635.27" // string | 
    annotationIds := []string{"Inner_example"} // []string | Limit the reports by internal, unstable annotation ids.
    sortField := "sortField_example" // string |  (optional)
    sortDirection := openapiclient.v2SortDirection("SORT_DIRECTION_UNSPECIFIED") // V2SortDirection |  (optional) (default to "SORT_DIRECTION_UNSPECIFIED")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.AnnotationReportFacets(context.Background(), accession, annotationIds).SortField(sortField).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.AnnotationReportFacets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AnnotationReportFacets`: V2GenomeAnnotationTableSummaryReply
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.AnnotationReportFacets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accession** | **string** |  | 
**annotationIds** | [**[]string**](string.md) | Limit the reports by internal, unstable annotation ids. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnnotationReportFacetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortField** | **string** |  | 
 **sortDirection** | [**V2SortDirection**](V2SortDirection.md) |  | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]

### Return type

[**V2GenomeAnnotationTableSummaryReply**](V2GenomeAnnotationTableSummaryReply.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnotationReportFacetsByPost

> V2GenomeAnnotationTableSummaryReply AnnotationReportFacetsByPost(ctx).V2GenomeAnnotationRequest(v2GenomeAnnotationRequest).Execute()

Get genome annotation report summary information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2GenomeAnnotationRequest := *openapiclient.NewV2GenomeAnnotationRequest() // V2GenomeAnnotationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.AnnotationReportFacetsByPost(context.Background()).V2GenomeAnnotationRequest(v2GenomeAnnotationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.AnnotationReportFacetsByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AnnotationReportFacetsByPost`: V2GenomeAnnotationTableSummaryReply
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.AnnotationReportFacetsByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnnotationReportFacetsByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2GenomeAnnotationRequest** | [**V2GenomeAnnotationRequest**](V2GenomeAnnotationRequest.md) |  | 

### Return type

[**V2GenomeAnnotationTableSummaryReply**](V2GenomeAnnotationTableSummaryReply.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssemblyAccessionsForSequenceAccession

> V2AssemblyAccessions AssemblyAccessionsForSequenceAccession(ctx, accession).Execute()

Get assembly accessions for a sequence accession



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accession := "NC_000001.11" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.AssemblyAccessionsForSequenceAccession(context.Background(), accession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.AssemblyAccessionsForSequenceAccession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssemblyAccessionsForSequenceAccession`: V2AssemblyAccessions
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.AssemblyAccessionsForSequenceAccession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accession** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssemblyAccessionsForSequenceAccessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2AssemblyAccessions**](V2AssemblyAccessions.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssemblyAccessionsForSequenceAccessionByPost

> V2AssemblyAccessions AssemblyAccessionsForSequenceAccessionByPost(ctx).V2SequenceAccessionRequest(v2SequenceAccessionRequest).Execute()

Get assembly accessions for a sequence accession



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2SequenceAccessionRequest := *openapiclient.NewV2SequenceAccessionRequest() // V2SequenceAccessionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.AssemblyAccessionsForSequenceAccessionByPost(context.Background()).V2SequenceAccessionRequest(v2SequenceAccessionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.AssemblyAccessionsForSequenceAccessionByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssemblyAccessionsForSequenceAccessionByPost`: V2AssemblyAccessions
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.AssemblyAccessionsForSequenceAccessionByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssemblyAccessionsForSequenceAccessionByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2SequenceAccessionRequest** | [**V2SequenceAccessionRequest**](V2SequenceAccessionRequest.md) |  | 

### Return type

[**V2AssemblyAccessions**](V2AssemblyAccessions.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssemblyRevisionHistoryByGet

> V2AssemblyRevisionHistory AssemblyRevisionHistoryByGet(ctx, accession).Execute()

Get revision history for assembly by accession



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accession := "GCF_000001405.40" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.AssemblyRevisionHistoryByGet(context.Background(), accession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.AssemblyRevisionHistoryByGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssemblyRevisionHistoryByGet`: V2AssemblyRevisionHistory
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.AssemblyRevisionHistoryByGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accession** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssemblyRevisionHistoryByGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2AssemblyRevisionHistory**](V2AssemblyRevisionHistory.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssemblyRevisionHistoryByPost

> V2AssemblyRevisionHistory AssemblyRevisionHistoryByPost(ctx).V2AssemblyRevisionHistoryRequest(v2AssemblyRevisionHistoryRequest).Execute()

Get revision history for assembly by accession



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2AssemblyRevisionHistoryRequest := *openapiclient.NewV2AssemblyRevisionHistoryRequest() // V2AssemblyRevisionHistoryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.AssemblyRevisionHistoryByPost(context.Background()).V2AssemblyRevisionHistoryRequest(v2AssemblyRevisionHistoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.AssemblyRevisionHistoryByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssemblyRevisionHistoryByPost`: V2AssemblyRevisionHistory
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.AssemblyRevisionHistoryByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssemblyRevisionHistoryByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2AssemblyRevisionHistoryRequest** | [**V2AssemblyRevisionHistoryRequest**](V2AssemblyRevisionHistoryRequest.md) |  | 

### Return type

[**V2AssemblyRevisionHistory**](V2AssemblyRevisionHistory.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckAssemblyAvailability

> V2AssemblyDatasetAvailability CheckAssemblyAvailability(ctx, accessions).ExpDebugValues(expDebugValues).Execute()

Check the validity of genome accessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accessions := []string{"Inner_example"} // []string | NCBI genome assembly accessions
    expDebugValues := "expDebugValues_example" // string | Supports debugging, e.g. by controlling data download speeds (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.CheckAssemblyAvailability(context.Background(), accessions).ExpDebugValues(expDebugValues).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.CheckAssemblyAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAssemblyAvailability`: V2AssemblyDatasetAvailability
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.CheckAssemblyAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | NCBI genome assembly accessions | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckAssemblyAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expDebugValues** | **string** | Supports debugging, e.g. by controlling data download speeds | 

### Return type

[**V2AssemblyDatasetAvailability**](V2AssemblyDatasetAvailability.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckAssemblyAvailabilityPost

> V2AssemblyDatasetAvailability CheckAssemblyAvailabilityPost(ctx).V2AssemblyDatasetRequest(v2AssemblyDatasetRequest).Execute()

Check the validity of many genome accessions in a single request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2AssemblyDatasetRequest := *openapiclient.NewV2AssemblyDatasetRequest() // V2AssemblyDatasetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.CheckAssemblyAvailabilityPost(context.Background()).V2AssemblyDatasetRequest(v2AssemblyDatasetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.CheckAssemblyAvailabilityPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAssemblyAvailabilityPost`: V2AssemblyDatasetAvailability
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.CheckAssemblyAvailabilityPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAssemblyAvailabilityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2AssemblyDatasetRequest** | [**V2AssemblyDatasetRequest**](V2AssemblyDatasetRequest.md) |  | 

### Return type

[**V2AssemblyDatasetAvailability**](V2AssemblyDatasetAvailability.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckmHistogramByTaxon

> V2AssemblyCheckMHistogramReply CheckmHistogramByTaxon(ctx, speciesTaxon).Execute()

Get CheckM histogram by species taxon



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    speciesTaxon := "202956" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.CheckmHistogramByTaxon(context.Background(), speciesTaxon).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.CheckmHistogramByTaxon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckmHistogramByTaxon`: V2AssemblyCheckMHistogramReply
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.CheckmHistogramByTaxon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**speciesTaxon** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckmHistogramByTaxonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2AssemblyCheckMHistogramReply**](V2AssemblyCheckMHistogramReply.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckmHistogramByTaxonByPost

> V2AssemblyCheckMHistogramReply CheckmHistogramByTaxonByPost(ctx).V2AssemblyCheckMHistogramRequest(v2AssemblyCheckMHistogramRequest).Execute()

Get CheckM histogram by species taxon



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2AssemblyCheckMHistogramRequest := *openapiclient.NewV2AssemblyCheckMHistogramRequest() // V2AssemblyCheckMHistogramRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.CheckmHistogramByTaxonByPost(context.Background()).V2AssemblyCheckMHistogramRequest(v2AssemblyCheckMHistogramRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.CheckmHistogramByTaxonByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckmHistogramByTaxonByPost`: V2AssemblyCheckMHistogramReply
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.CheckmHistogramByTaxonByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckmHistogramByTaxonByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2AssemblyCheckMHistogramRequest** | [**V2AssemblyCheckMHistogramRequest**](V2AssemblyCheckMHistogramRequest.md) |  | 

### Return type

[**V2AssemblyCheckMHistogramReply**](V2AssemblyCheckMHistogramReply.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadAssemblyPackage

> *os.File DownloadAssemblyPackage(ctx, accessions).Chromosomes(chromosomes).IncludeAnnotationType(includeAnnotationType).Hydrated(hydrated).ExpDebugValues(expDebugValues).Filename(filename).Execute()

Get a genome dataset by accession



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accessions := []string{"Inner_example"} // []string | NCBI genome assembly accessions
    chromosomes := []string{"Inner_example"} // []string | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome ('Un'). The filter only applies to fasta sequence. (optional)
    includeAnnotationType := []openapiclient.V2AnnotationForAssemblyType{openapiclient.v2AnnotationForAssemblyType("GENOME_GFF")} // []V2AnnotationForAssemblyType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    hydrated := openapiclient.v2AssemblyDatasetRequestResolution("FULLY_HYDRATED") // V2AssemblyDatasetRequestResolution | Set to DATA_REPORT_ONLY, to only retrieve data-reports. (optional) (default to "FULLY_HYDRATED")
    expDebugValues := "expDebugValues_example" // string | Supports debugging, e.g. by controlling data download speeds (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.DownloadAssemblyPackage(context.Background(), accessions).Chromosomes(chromosomes).IncludeAnnotationType(includeAnnotationType).Hydrated(hydrated).ExpDebugValues(expDebugValues).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.DownloadAssemblyPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadAssemblyPackage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.DownloadAssemblyPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | NCBI genome assembly accessions | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAssemblyPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chromosomes** | **[]string** | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | 
 **includeAnnotationType** | [**[]V2AnnotationForAssemblyType**](V2AnnotationForAssemblyType.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
 **hydrated** | [**V2AssemblyDatasetRequestResolution**](V2AssemblyDatasetRequestResolution.md) | Set to DATA_REPORT_ONLY, to only retrieve data-reports. | [default to &quot;FULLY_HYDRATED&quot;]
 **expDebugValues** | **string** | Supports debugging, e.g. by controlling data download speeds | 
 **filename** | **string** | Output file name. | [default to &quot;ncbi_dataset.zip&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadAssemblyPackagePost

> *os.File DownloadAssemblyPackagePost(ctx).V2AssemblyDatasetRequest(v2AssemblyDatasetRequest).Filename(filename).Execute()

Get a genome dataset by post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2AssemblyDatasetRequest := *openapiclient.NewV2AssemblyDatasetRequest() // V2AssemblyDatasetRequest | 
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.DownloadAssemblyPackagePost(context.Background()).V2AssemblyDatasetRequest(v2AssemblyDatasetRequest).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.DownloadAssemblyPackagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadAssemblyPackagePost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.DownloadAssemblyPackagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAssemblyPackagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2AssemblyDatasetRequest** | [**V2AssemblyDatasetRequest**](V2AssemblyDatasetRequest.md) |  | 
 **filename** | **string** | Output file name. | [default to &quot;ncbi_dataset.zip&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadGenomeAnnotationPackage

> *os.File DownloadGenomeAnnotationPackage(ctx, accession).AnnotationIds(annotationIds).Symbols(symbols).Locations(locations).GeneTypes(geneTypes).SearchText(searchText).SortField(sortField).SortDirection(sortDirection).IncludeAnnotationType(includeAnnotationType).PageSize(pageSize).TableFields(tableFields).TableFormat(tableFormat).IncludeTabularHeader(includeTabularHeader).PageToken(pageToken).Filename(filename).Execute()

Get an annotation report dataset by accession



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accession := "GCF_000001635.27" // string | 
    annotationIds := []string{"Inner_example"} // []string | Limit the reports by internal, unstable annotation ids. (optional)
    symbols := []string{"Inner_example"} // []string | Filter parameters (optional)
    locations := []string{"Inner_example"} // []string | Locations with a chromosome or accession and optional start-stop range: chromosome|accession[:start-end] (optional)
    geneTypes := []string{"Inner_example"} // []string | granular gene_types (optional)
    searchText := []string{"Inner_example"} // []string |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortDirection := openapiclient.v2SortDirection("SORT_DIRECTION_UNSPECIFIED") // V2SortDirection |  (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    includeAnnotationType := []openapiclient.V2GenomeAnnotationRequestAnnotationType{openapiclient.v2GenomeAnnotationRequestAnnotationType("GENOME_FASTA")} // []V2GenomeAnnotationRequestAnnotationType |  (optional)
    pageSize := int32(56) // int32 | The maximum number of features to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    tableFormat := openapiclient.v2GenomeAnnotationRequestGenomeAnnotationTableFormat("NO_TABLE") // V2GenomeAnnotationRequestGenomeAnnotationTableFormat | Optional pre-defined template for processing a tabular data request (optional) (default to "NO_TABLE")
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")
    pageToken := "pageToken_example" // string | A page token is returned from a `GetFeatures` call with more than `page_size` results. Use this token, along with the previous `FeatureRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.DownloadGenomeAnnotationPackage(context.Background(), accession).AnnotationIds(annotationIds).Symbols(symbols).Locations(locations).GeneTypes(geneTypes).SearchText(searchText).SortField(sortField).SortDirection(sortDirection).IncludeAnnotationType(includeAnnotationType).PageSize(pageSize).TableFields(tableFields).TableFormat(tableFormat).IncludeTabularHeader(includeTabularHeader).PageToken(pageToken).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.DownloadGenomeAnnotationPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadGenomeAnnotationPackage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.DownloadGenomeAnnotationPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accession** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGenomeAnnotationPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **annotationIds** | **[]string** | Limit the reports by internal, unstable annotation ids. | 
 **symbols** | **[]string** | Filter parameters | 
 **locations** | **[]string** | Locations with a chromosome or accession and optional start-stop range: chromosome|accession[:start-end] | 
 **geneTypes** | **[]string** | granular gene_types | 
 **searchText** | **[]string** |  | 
 **sortField** | **string** |  | 
 **sortDirection** | [**V2SortDirection**](V2SortDirection.md) |  | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **includeAnnotationType** | [**[]V2GenomeAnnotationRequestAnnotationType**](V2GenomeAnnotationRequestAnnotationType.md) |  | 
 **pageSize** | **int32** | The maximum number of features to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **tableFormat** | [**V2GenomeAnnotationRequestGenomeAnnotationTableFormat**](V2GenomeAnnotationRequestGenomeAnnotationTableFormat.md) | Optional pre-defined template for processing a tabular data request | [default to &quot;NO_TABLE&quot;]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]
 **pageToken** | **string** | A page token is returned from a &#x60;GetFeatures&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;FeatureRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **filename** | **string** | Output file name. | [default to &quot;ncbi_dataset.zip&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadGenomeAnnotationPackageByPost

> *os.File DownloadGenomeAnnotationPackageByPost(ctx).V2GenomeAnnotationRequest(v2GenomeAnnotationRequest).Filename(filename).Execute()

Get an annotation report dataset by accession



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2GenomeAnnotationRequest := *openapiclient.NewV2GenomeAnnotationRequest() // V2GenomeAnnotationRequest | 
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.DownloadGenomeAnnotationPackageByPost(context.Background()).V2GenomeAnnotationRequest(v2GenomeAnnotationRequest).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.DownloadGenomeAnnotationPackageByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadGenomeAnnotationPackageByPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.DownloadGenomeAnnotationPackageByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGenomeAnnotationPackageByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2GenomeAnnotationRequest** | [**V2GenomeAnnotationRequest**](V2GenomeAnnotationRequest.md) |  | 
 **filename** | **string** | Output file name. | [default to &quot;ncbi_dataset.zip&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeAnnotationDownloadSummary

> V2DownloadSummary GenomeAnnotationDownloadSummary(ctx, accession).AnnotationIds(annotationIds).Symbols(symbols).Locations(locations).GeneTypes(geneTypes).SearchText(searchText).SortField(sortField).SortDirection(sortDirection).IncludeAnnotationType(includeAnnotationType).TableFields(tableFields).TableFormat(tableFormat).IncludeTabularHeader(includeTabularHeader).Execute()

Preview feature dataset download



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accession := "GCF_000001635.27" // string | 
    annotationIds := []string{"Inner_example"} // []string | Limit the reports by internal, unstable annotation ids. (optional)
    symbols := []string{"Inner_example"} // []string | Filter parameters (optional)
    locations := []string{"Inner_example"} // []string | Locations with a chromosome or accession and optional start-stop range: chromosome|accession[:start-end] (optional)
    geneTypes := []string{"Inner_example"} // []string | granular gene_types (optional)
    searchText := []string{"Inner_example"} // []string |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortDirection := openapiclient.v2SortDirection("SORT_DIRECTION_UNSPECIFIED") // V2SortDirection |  (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    includeAnnotationType := []openapiclient.V2GenomeAnnotationRequestAnnotationType{openapiclient.v2GenomeAnnotationRequestAnnotationType("GENOME_FASTA")} // []V2GenomeAnnotationRequestAnnotationType |  (optional)
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    tableFormat := openapiclient.v2GenomeAnnotationRequestGenomeAnnotationTableFormat("NO_TABLE") // V2GenomeAnnotationRequestGenomeAnnotationTableFormat | Optional pre-defined template for processing a tabular data request (optional) (default to "NO_TABLE")
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeAnnotationDownloadSummary(context.Background(), accession).AnnotationIds(annotationIds).Symbols(symbols).Locations(locations).GeneTypes(geneTypes).SearchText(searchText).SortField(sortField).SortDirection(sortDirection).IncludeAnnotationType(includeAnnotationType).TableFields(tableFields).TableFormat(tableFormat).IncludeTabularHeader(includeTabularHeader).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeAnnotationDownloadSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeAnnotationDownloadSummary`: V2DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeAnnotationDownloadSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accession** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeAnnotationDownloadSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **annotationIds** | **[]string** | Limit the reports by internal, unstable annotation ids. | 
 **symbols** | **[]string** | Filter parameters | 
 **locations** | **[]string** | Locations with a chromosome or accession and optional start-stop range: chromosome|accession[:start-end] | 
 **geneTypes** | **[]string** | granular gene_types | 
 **searchText** | **[]string** |  | 
 **sortField** | **string** |  | 
 **sortDirection** | [**V2SortDirection**](V2SortDirection.md) |  | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **includeAnnotationType** | [**[]V2GenomeAnnotationRequestAnnotationType**](V2GenomeAnnotationRequestAnnotationType.md) |  | 
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **tableFormat** | [**V2GenomeAnnotationRequestGenomeAnnotationTableFormat**](V2GenomeAnnotationRequestGenomeAnnotationTableFormat.md) | Optional pre-defined template for processing a tabular data request | [default to &quot;NO_TABLE&quot;]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]

### Return type

[**V2DownloadSummary**](V2DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeAnnotationDownloadSummaryByPost

> V2DownloadSummary GenomeAnnotationDownloadSummaryByPost(ctx).V2GenomeAnnotationRequest(v2GenomeAnnotationRequest).Execute()

Preview feature download by POST



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2GenomeAnnotationRequest := *openapiclient.NewV2GenomeAnnotationRequest() // V2GenomeAnnotationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeAnnotationDownloadSummaryByPost(context.Background()).V2GenomeAnnotationRequest(v2GenomeAnnotationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeAnnotationDownloadSummaryByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeAnnotationDownloadSummaryByPost`: V2DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeAnnotationDownloadSummaryByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenomeAnnotationDownloadSummaryByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2GenomeAnnotationRequest** | [**V2GenomeAnnotationRequest**](V2GenomeAnnotationRequest.md) |  | 

### Return type

[**V2DownloadSummary**](V2DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeAnnotationReport

> V2reportsGenomeAnnotationReportPage GenomeAnnotationReport(ctx, accession).AnnotationIds(annotationIds).Symbols(symbols).Locations(locations).GeneTypes(geneTypes).SearchText(searchText).SortField(sortField).SortDirection(sortDirection).PageSize(pageSize).TableFields(tableFields).TableFormat(tableFormat).IncludeTabularHeader(includeTabularHeader).PageToken(pageToken).Execute()

Get genome annotation reports by genome accession



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accession := "GCF_000001635.27" // string | 
    annotationIds := []string{"Inner_example"} // []string | Limit the reports by internal, unstable annotation ids. (optional)
    symbols := []string{"Inner_example"} // []string | Filter parameters (optional)
    locations := []string{"Inner_example"} // []string | Locations with a chromosome or accession and optional start-stop range: chromosome|accession[:start-end] (optional)
    geneTypes := []string{"Inner_example"} // []string | granular gene_types (optional)
    searchText := []string{"Inner_example"} // []string |  (optional)
    sortField := "sortField_example" // string |  (optional)
    sortDirection := openapiclient.v2SortDirection("SORT_DIRECTION_UNSPECIFIED") // V2SortDirection |  (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    pageSize := int32(56) // int32 | The maximum number of features to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    tableFormat := openapiclient.v2GenomeAnnotationRequestGenomeAnnotationTableFormat("NO_TABLE") // V2GenomeAnnotationRequestGenomeAnnotationTableFormat | Optional pre-defined template for processing a tabular data request (optional) (default to "NO_TABLE")
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")
    pageToken := "pageToken_example" // string | A page token is returned from a `GetFeatures` call with more than `page_size` results. Use this token, along with the previous `FeatureRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeAnnotationReport(context.Background(), accession).AnnotationIds(annotationIds).Symbols(symbols).Locations(locations).GeneTypes(geneTypes).SearchText(searchText).SortField(sortField).SortDirection(sortDirection).PageSize(pageSize).TableFields(tableFields).TableFormat(tableFormat).IncludeTabularHeader(includeTabularHeader).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeAnnotationReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeAnnotationReport`: V2reportsGenomeAnnotationReportPage
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeAnnotationReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accession** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeAnnotationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **annotationIds** | **[]string** | Limit the reports by internal, unstable annotation ids. | 
 **symbols** | **[]string** | Filter parameters | 
 **locations** | **[]string** | Locations with a chromosome or accession and optional start-stop range: chromosome|accession[:start-end] | 
 **geneTypes** | **[]string** | granular gene_types | 
 **searchText** | **[]string** |  | 
 **sortField** | **string** |  | 
 **sortDirection** | [**V2SortDirection**](V2SortDirection.md) |  | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **pageSize** | **int32** | The maximum number of features to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **tableFormat** | [**V2GenomeAnnotationRequestGenomeAnnotationTableFormat**](V2GenomeAnnotationRequestGenomeAnnotationTableFormat.md) | Optional pre-defined template for processing a tabular data request | [default to &quot;NO_TABLE&quot;]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]
 **pageToken** | **string** | A page token is returned from a &#x60;GetFeatures&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;FeatureRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V2reportsGenomeAnnotationReportPage**](V2reportsGenomeAnnotationReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeAnnotationReportByPost

> V2reportsGenomeAnnotationReportPage GenomeAnnotationReportByPost(ctx).V2GenomeAnnotationRequest(v2GenomeAnnotationRequest).Execute()

Get genome annotation reports by genome accession



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2GenomeAnnotationRequest := *openapiclient.NewV2GenomeAnnotationRequest() // V2GenomeAnnotationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeAnnotationReportByPost(context.Background()).V2GenomeAnnotationRequest(v2GenomeAnnotationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeAnnotationReportByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeAnnotationReportByPost`: V2reportsGenomeAnnotationReportPage
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeAnnotationReportByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenomeAnnotationReportByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2GenomeAnnotationRequest** | [**V2GenomeAnnotationRequest**](V2GenomeAnnotationRequest.md) |  | 

### Return type

[**V2reportsGenomeAnnotationReportPage**](V2reportsGenomeAnnotationReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeDatasetReport

> V2reportsAssemblyDataReportPage GenomeDatasetReport(ctx, accessions).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).FiltersIsMetagenomeDerived(filtersIsMetagenomeDerived).FiltersIsTypeMaterial(filtersIsTypeMaterial).FiltersIsIctvExemplar(filtersIsIctvExemplar).FiltersExcludeMultiIsolate(filtersExcludeMultiIsolate).FiltersTypeMaterialCategory(filtersTypeMaterialCategory).TaxExactMatch(taxExactMatch).TableFields(tableFields).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).SortField(sortField).SortDirection(sortDirection).IncludeTabularHeader(includeTabularHeader).Execute()

Get dataset reports by accessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accessions := []string{"Inner_example"} // []string | 
    filtersReferenceOnly := true // bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional) (default to false)
    filtersAssemblySource := openapiclient.v2AssemblyDatasetDescriptorsFilterAssemblySource("all") // V2AssemblyDatasetDescriptorsFilterAssemblySource | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies (optional) (default to "all")
    filtersHasAnnotation := true // bool | Return only annotated genome assemblies (optional) (default to false)
    filtersExcludePairedReports := false // bool | For paired (GCA/GCF) records, only return the primary record (optional) (default to false)
    filtersExcludeAtypical := false // bool | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical (optional) (default to false)
    filtersAssemblyVersion := openapiclient.v2AssemblyDatasetDescriptorsFilterAssemblyVersion("current") // V2AssemblyDatasetDescriptorsFilterAssemblyVersion | Return all assemblies, including replaced and suppressed, or only current assemblies (optional) (default to "current")
    filtersAssemblyLevel := []openapiclient.V2reportsAssemblyLevel{openapiclient.v2reportsAssemblyLevel("chromosome")} // []V2reportsAssemblyLevel | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. (optional)
    filtersFirstReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or after the specified date By default, do not filter. (optional)
    filtersLastReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or before to the specified date By default, do not filter. (optional)
    filtersSearchText := []string{"Inner_example"} // []string | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter (optional)
    filtersIsMetagenomeDerived := openapiclient.v2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter("METAGENOME_DERIVED_UNSET") // V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter |  (optional) (default to "METAGENOME_DERIVED_UNSET")
    filtersIsTypeMaterial := false // bool | If true, include only type materials (optional) (default to false)
    filtersIsIctvExemplar := false // bool | If true, include only ICTV Exemplars (optional) (default to false)
    filtersExcludeMultiIsolate := false // bool | If true, exclude large multi-isolate projects (optional) (default to false)
    filtersTypeMaterialCategory := openapiclient.v2AssemblyDatasetDescriptorsFilterTypeMaterialCategory("NONE") // V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory |  (optional) (default to "NONE")
    taxExactMatch := true // bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. (optional) (default to false)
    tableFields := []string{"Inner_example"} // []string |  (optional)
    returnedContent := openapiclient.v2AssemblyDatasetReportsRequestContentType("COMPLETE") // V2AssemblyDatasetReportsRequestContentType | Return either assembly accessions, or complete assembly reports (optional) (default to "COMPLETE")
    pageSize := int32(56) // int32 | The maximum number of genome assembly reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyDatasetReportsRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyDatasetReportsRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)
    sortField := "sortField_example" // string |  (optional)
    sortDirection := openapiclient.v2SortDirection("SORT_DIRECTION_UNSPECIFIED") // V2SortDirection |  (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeDatasetReport(context.Background(), accessions).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).FiltersIsMetagenomeDerived(filtersIsMetagenomeDerived).FiltersIsTypeMaterial(filtersIsTypeMaterial).FiltersIsIctvExemplar(filtersIsIctvExemplar).FiltersExcludeMultiIsolate(filtersExcludeMultiIsolate).FiltersTypeMaterialCategory(filtersTypeMaterialCategory).TaxExactMatch(taxExactMatch).TableFields(tableFields).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).SortField(sortField).SortDirection(sortDirection).IncludeTabularHeader(includeTabularHeader).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeDatasetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeDatasetReport`: V2reportsAssemblyDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeDatasetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeDatasetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtersReferenceOnly** | **bool** | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [default to false]
 **filtersAssemblySource** | [**V2AssemblyDatasetDescriptorsFilterAssemblySource**](V2AssemblyDatasetDescriptorsFilterAssemblySource.md) | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies | [default to &quot;all&quot;]
 **filtersHasAnnotation** | **bool** | Return only annotated genome assemblies | [default to false]
 **filtersExcludePairedReports** | **bool** | For paired (GCA/GCF) records, only return the primary record | [default to false]
 **filtersExcludeAtypical** | **bool** | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical | [default to false]
 **filtersAssemblyVersion** | [**V2AssemblyDatasetDescriptorsFilterAssemblyVersion**](V2AssemblyDatasetDescriptorsFilterAssemblyVersion.md) | Return all assemblies, including replaced and suppressed, or only current assemblies | [default to &quot;current&quot;]
 **filtersAssemblyLevel** | [**[]V2reportsAssemblyLevel**](V2reportsAssemblyLevel.md) | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. | 
 **filtersFirstReleaseDate** | **time.Time** | Only return genome assemblies that were released on or after the specified date By default, do not filter. | 
 **filtersLastReleaseDate** | **time.Time** | Only return genome assemblies that were released on or before to the specified date By default, do not filter. | 
 **filtersSearchText** | **[]string** | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter | 
 **filtersIsMetagenomeDerived** | [**V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter**](V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter.md) |  | [default to &quot;METAGENOME_DERIVED_UNSET&quot;]
 **filtersIsTypeMaterial** | **bool** | If true, include only type materials | [default to false]
 **filtersIsIctvExemplar** | **bool** | If true, include only ICTV Exemplars | [default to false]
 **filtersExcludeMultiIsolate** | **bool** | If true, exclude large multi-isolate projects | [default to false]
 **filtersTypeMaterialCategory** | [**V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory**](V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory.md) |  | [default to &quot;NONE&quot;]
 **taxExactMatch** | **bool** | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. | [default to false]
 **tableFields** | **[]string** |  | 
 **returnedContent** | [**V2AssemblyDatasetReportsRequestContentType**](V2AssemblyDatasetReportsRequestContentType.md) | Return either assembly accessions, or complete assembly reports | [default to &quot;COMPLETE&quot;]
 **pageSize** | **int32** | The maximum number of genome assembly reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyDatasetReportsRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyDatasetReportsRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **sortField** | **string** |  | 
 **sortDirection** | [**V2SortDirection**](V2SortDirection.md) |  | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]

### Return type

[**V2reportsAssemblyDataReportPage**](V2reportsAssemblyDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeDatasetReportByPost

> V2reportsAssemblyDataReportPage GenomeDatasetReportByPost(ctx).V2AssemblyDatasetReportsRequest(v2AssemblyDatasetReportsRequest).Execute()

Get dataset reports by accessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2AssemblyDatasetReportsRequest := *openapiclient.NewV2AssemblyDatasetReportsRequest() // V2AssemblyDatasetReportsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeDatasetReportByPost(context.Background()).V2AssemblyDatasetReportsRequest(v2AssemblyDatasetReportsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeDatasetReportByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeDatasetReportByPost`: V2reportsAssemblyDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeDatasetReportByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenomeDatasetReportByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2AssemblyDatasetReportsRequest** | [**V2AssemblyDatasetReportsRequest**](V2AssemblyDatasetReportsRequest.md) |  | 

### Return type

[**V2reportsAssemblyDataReportPage**](V2reportsAssemblyDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeDatasetReportsByAssemblyName

> V2reportsAssemblyDataReportPage GenomeDatasetReportsByAssemblyName(ctx, assemblyNames).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).FiltersIsMetagenomeDerived(filtersIsMetagenomeDerived).FiltersIsTypeMaterial(filtersIsTypeMaterial).FiltersIsIctvExemplar(filtersIsIctvExemplar).FiltersExcludeMultiIsolate(filtersExcludeMultiIsolate).FiltersTypeMaterialCategory(filtersTypeMaterialCategory).TaxExactMatch(taxExactMatch).TableFields(tableFields).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).SortField(sortField).SortDirection(sortDirection).IncludeTabularHeader(includeTabularHeader).Execute()

Get dataset reports by assembly name (exact)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    assemblyNames := []string{"Inner_example"} // []string | 
    filtersReferenceOnly := true // bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional) (default to false)
    filtersAssemblySource := openapiclient.v2AssemblyDatasetDescriptorsFilterAssemblySource("all") // V2AssemblyDatasetDescriptorsFilterAssemblySource | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies (optional) (default to "all")
    filtersHasAnnotation := true // bool | Return only annotated genome assemblies (optional) (default to false)
    filtersExcludePairedReports := false // bool | For paired (GCA/GCF) records, only return the primary record (optional) (default to false)
    filtersExcludeAtypical := false // bool | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical (optional) (default to false)
    filtersAssemblyVersion := openapiclient.v2AssemblyDatasetDescriptorsFilterAssemblyVersion("current") // V2AssemblyDatasetDescriptorsFilterAssemblyVersion | Return all assemblies, including replaced and suppressed, or only current assemblies (optional) (default to "current")
    filtersAssemblyLevel := []openapiclient.V2reportsAssemblyLevel{openapiclient.v2reportsAssemblyLevel("chromosome")} // []V2reportsAssemblyLevel | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. (optional)
    filtersFirstReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or after the specified date By default, do not filter. (optional)
    filtersLastReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or before to the specified date By default, do not filter. (optional)
    filtersSearchText := []string{"Inner_example"} // []string | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter (optional)
    filtersIsMetagenomeDerived := openapiclient.v2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter("METAGENOME_DERIVED_UNSET") // V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter |  (optional) (default to "METAGENOME_DERIVED_UNSET")
    filtersIsTypeMaterial := false // bool | If true, include only type materials (optional) (default to false)
    filtersIsIctvExemplar := false // bool | If true, include only ICTV Exemplars (optional) (default to false)
    filtersExcludeMultiIsolate := false // bool | If true, exclude large multi-isolate projects (optional) (default to false)
    filtersTypeMaterialCategory := openapiclient.v2AssemblyDatasetDescriptorsFilterTypeMaterialCategory("NONE") // V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory |  (optional) (default to "NONE")
    taxExactMatch := true // bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. (optional) (default to false)
    tableFields := []string{"Inner_example"} // []string |  (optional)
    returnedContent := openapiclient.v2AssemblyDatasetReportsRequestContentType("COMPLETE") // V2AssemblyDatasetReportsRequestContentType | Return either assembly accessions, or complete assembly reports (optional) (default to "COMPLETE")
    pageSize := int32(56) // int32 | The maximum number of genome assembly reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyDatasetReportsRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyDatasetReportsRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)
    sortField := "sortField_example" // string |  (optional)
    sortDirection := openapiclient.v2SortDirection("SORT_DIRECTION_UNSPECIFIED") // V2SortDirection |  (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeDatasetReportsByAssemblyName(context.Background(), assemblyNames).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).FiltersIsMetagenomeDerived(filtersIsMetagenomeDerived).FiltersIsTypeMaterial(filtersIsTypeMaterial).FiltersIsIctvExemplar(filtersIsIctvExemplar).FiltersExcludeMultiIsolate(filtersExcludeMultiIsolate).FiltersTypeMaterialCategory(filtersTypeMaterialCategory).TaxExactMatch(taxExactMatch).TableFields(tableFields).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).SortField(sortField).SortDirection(sortDirection).IncludeTabularHeader(includeTabularHeader).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeDatasetReportsByAssemblyName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeDatasetReportsByAssemblyName`: V2reportsAssemblyDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeDatasetReportsByAssemblyName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assemblyNames** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeDatasetReportsByAssemblyNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtersReferenceOnly** | **bool** | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [default to false]
 **filtersAssemblySource** | [**V2AssemblyDatasetDescriptorsFilterAssemblySource**](V2AssemblyDatasetDescriptorsFilterAssemblySource.md) | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies | [default to &quot;all&quot;]
 **filtersHasAnnotation** | **bool** | Return only annotated genome assemblies | [default to false]
 **filtersExcludePairedReports** | **bool** | For paired (GCA/GCF) records, only return the primary record | [default to false]
 **filtersExcludeAtypical** | **bool** | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical | [default to false]
 **filtersAssemblyVersion** | [**V2AssemblyDatasetDescriptorsFilterAssemblyVersion**](V2AssemblyDatasetDescriptorsFilterAssemblyVersion.md) | Return all assemblies, including replaced and suppressed, or only current assemblies | [default to &quot;current&quot;]
 **filtersAssemblyLevel** | [**[]V2reportsAssemblyLevel**](V2reportsAssemblyLevel.md) | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. | 
 **filtersFirstReleaseDate** | **time.Time** | Only return genome assemblies that were released on or after the specified date By default, do not filter. | 
 **filtersLastReleaseDate** | **time.Time** | Only return genome assemblies that were released on or before to the specified date By default, do not filter. | 
 **filtersSearchText** | **[]string** | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter | 
 **filtersIsMetagenomeDerived** | [**V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter**](V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter.md) |  | [default to &quot;METAGENOME_DERIVED_UNSET&quot;]
 **filtersIsTypeMaterial** | **bool** | If true, include only type materials | [default to false]
 **filtersIsIctvExemplar** | **bool** | If true, include only ICTV Exemplars | [default to false]
 **filtersExcludeMultiIsolate** | **bool** | If true, exclude large multi-isolate projects | [default to false]
 **filtersTypeMaterialCategory** | [**V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory**](V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory.md) |  | [default to &quot;NONE&quot;]
 **taxExactMatch** | **bool** | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. | [default to false]
 **tableFields** | **[]string** |  | 
 **returnedContent** | [**V2AssemblyDatasetReportsRequestContentType**](V2AssemblyDatasetReportsRequestContentType.md) | Return either assembly accessions, or complete assembly reports | [default to &quot;COMPLETE&quot;]
 **pageSize** | **int32** | The maximum number of genome assembly reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyDatasetReportsRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyDatasetReportsRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **sortField** | **string** |  | 
 **sortDirection** | [**V2SortDirection**](V2SortDirection.md) |  | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]

### Return type

[**V2reportsAssemblyDataReportPage**](V2reportsAssemblyDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeDatasetReportsByBioproject

> V2reportsAssemblyDataReportPage GenomeDatasetReportsByBioproject(ctx, bioprojects).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).FiltersIsMetagenomeDerived(filtersIsMetagenomeDerived).FiltersIsTypeMaterial(filtersIsTypeMaterial).FiltersIsIctvExemplar(filtersIsIctvExemplar).FiltersExcludeMultiIsolate(filtersExcludeMultiIsolate).FiltersTypeMaterialCategory(filtersTypeMaterialCategory).TaxExactMatch(taxExactMatch).TableFields(tableFields).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).SortField(sortField).SortDirection(sortDirection).IncludeTabularHeader(includeTabularHeader).Execute()

Get dataset reports by bioproject



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    bioprojects := []string{"Inner_example"} // []string | 
    filtersReferenceOnly := true // bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional) (default to false)
    filtersAssemblySource := openapiclient.v2AssemblyDatasetDescriptorsFilterAssemblySource("all") // V2AssemblyDatasetDescriptorsFilterAssemblySource | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies (optional) (default to "all")
    filtersHasAnnotation := true // bool | Return only annotated genome assemblies (optional) (default to false)
    filtersExcludePairedReports := false // bool | For paired (GCA/GCF) records, only return the primary record (optional) (default to false)
    filtersExcludeAtypical := false // bool | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical (optional) (default to false)
    filtersAssemblyVersion := openapiclient.v2AssemblyDatasetDescriptorsFilterAssemblyVersion("current") // V2AssemblyDatasetDescriptorsFilterAssemblyVersion | Return all assemblies, including replaced and suppressed, or only current assemblies (optional) (default to "current")
    filtersAssemblyLevel := []openapiclient.V2reportsAssemblyLevel{openapiclient.v2reportsAssemblyLevel("chromosome")} // []V2reportsAssemblyLevel | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. (optional)
    filtersFirstReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or after the specified date By default, do not filter. (optional)
    filtersLastReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or before to the specified date By default, do not filter. (optional)
    filtersSearchText := []string{"Inner_example"} // []string | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter (optional)
    filtersIsMetagenomeDerived := openapiclient.v2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter("METAGENOME_DERIVED_UNSET") // V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter |  (optional) (default to "METAGENOME_DERIVED_UNSET")
    filtersIsTypeMaterial := false // bool | If true, include only type materials (optional) (default to false)
    filtersIsIctvExemplar := false // bool | If true, include only ICTV Exemplars (optional) (default to false)
    filtersExcludeMultiIsolate := false // bool | If true, exclude large multi-isolate projects (optional) (default to false)
    filtersTypeMaterialCategory := openapiclient.v2AssemblyDatasetDescriptorsFilterTypeMaterialCategory("NONE") // V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory |  (optional) (default to "NONE")
    taxExactMatch := true // bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. (optional) (default to false)
    tableFields := []string{"Inner_example"} // []string |  (optional)
    returnedContent := openapiclient.v2AssemblyDatasetReportsRequestContentType("COMPLETE") // V2AssemblyDatasetReportsRequestContentType | Return either assembly accessions, or complete assembly reports (optional) (default to "COMPLETE")
    pageSize := int32(56) // int32 | The maximum number of genome assembly reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyDatasetReportsRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyDatasetReportsRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)
    sortField := "sortField_example" // string |  (optional)
    sortDirection := openapiclient.v2SortDirection("SORT_DIRECTION_UNSPECIFIED") // V2SortDirection |  (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeDatasetReportsByBioproject(context.Background(), bioprojects).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).FiltersIsMetagenomeDerived(filtersIsMetagenomeDerived).FiltersIsTypeMaterial(filtersIsTypeMaterial).FiltersIsIctvExemplar(filtersIsIctvExemplar).FiltersExcludeMultiIsolate(filtersExcludeMultiIsolate).FiltersTypeMaterialCategory(filtersTypeMaterialCategory).TaxExactMatch(taxExactMatch).TableFields(tableFields).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).SortField(sortField).SortDirection(sortDirection).IncludeTabularHeader(includeTabularHeader).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeDatasetReportsByBioproject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeDatasetReportsByBioproject`: V2reportsAssemblyDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeDatasetReportsByBioproject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bioprojects** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeDatasetReportsByBioprojectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtersReferenceOnly** | **bool** | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [default to false]
 **filtersAssemblySource** | [**V2AssemblyDatasetDescriptorsFilterAssemblySource**](V2AssemblyDatasetDescriptorsFilterAssemblySource.md) | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies | [default to &quot;all&quot;]
 **filtersHasAnnotation** | **bool** | Return only annotated genome assemblies | [default to false]
 **filtersExcludePairedReports** | **bool** | For paired (GCA/GCF) records, only return the primary record | [default to false]
 **filtersExcludeAtypical** | **bool** | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical | [default to false]
 **filtersAssemblyVersion** | [**V2AssemblyDatasetDescriptorsFilterAssemblyVersion**](V2AssemblyDatasetDescriptorsFilterAssemblyVersion.md) | Return all assemblies, including replaced and suppressed, or only current assemblies | [default to &quot;current&quot;]
 **filtersAssemblyLevel** | [**[]V2reportsAssemblyLevel**](V2reportsAssemblyLevel.md) | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. | 
 **filtersFirstReleaseDate** | **time.Time** | Only return genome assemblies that were released on or after the specified date By default, do not filter. | 
 **filtersLastReleaseDate** | **time.Time** | Only return genome assemblies that were released on or before to the specified date By default, do not filter. | 
 **filtersSearchText** | **[]string** | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter | 
 **filtersIsMetagenomeDerived** | [**V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter**](V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter.md) |  | [default to &quot;METAGENOME_DERIVED_UNSET&quot;]
 **filtersIsTypeMaterial** | **bool** | If true, include only type materials | [default to false]
 **filtersIsIctvExemplar** | **bool** | If true, include only ICTV Exemplars | [default to false]
 **filtersExcludeMultiIsolate** | **bool** | If true, exclude large multi-isolate projects | [default to false]
 **filtersTypeMaterialCategory** | [**V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory**](V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory.md) |  | [default to &quot;NONE&quot;]
 **taxExactMatch** | **bool** | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. | [default to false]
 **tableFields** | **[]string** |  | 
 **returnedContent** | [**V2AssemblyDatasetReportsRequestContentType**](V2AssemblyDatasetReportsRequestContentType.md) | Return either assembly accessions, or complete assembly reports | [default to &quot;COMPLETE&quot;]
 **pageSize** | **int32** | The maximum number of genome assembly reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyDatasetReportsRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyDatasetReportsRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **sortField** | **string** |  | 
 **sortDirection** | [**V2SortDirection**](V2SortDirection.md) |  | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]

### Return type

[**V2reportsAssemblyDataReportPage**](V2reportsAssemblyDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeDatasetReportsByBiosampleId

> V2reportsAssemblyDataReportPage GenomeDatasetReportsByBiosampleId(ctx, biosampleIds).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).FiltersIsMetagenomeDerived(filtersIsMetagenomeDerived).FiltersIsTypeMaterial(filtersIsTypeMaterial).FiltersIsIctvExemplar(filtersIsIctvExemplar).FiltersExcludeMultiIsolate(filtersExcludeMultiIsolate).FiltersTypeMaterialCategory(filtersTypeMaterialCategory).TaxExactMatch(taxExactMatch).TableFields(tableFields).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).SortField(sortField).SortDirection(sortDirection).IncludeTabularHeader(includeTabularHeader).Execute()

Get dataset reports by biosample id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    biosampleIds := []string{"Inner_example"} // []string | 
    filtersReferenceOnly := true // bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional) (default to false)
    filtersAssemblySource := openapiclient.v2AssemblyDatasetDescriptorsFilterAssemblySource("all") // V2AssemblyDatasetDescriptorsFilterAssemblySource | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies (optional) (default to "all")
    filtersHasAnnotation := true // bool | Return only annotated genome assemblies (optional) (default to false)
    filtersExcludePairedReports := false // bool | For paired (GCA/GCF) records, only return the primary record (optional) (default to false)
    filtersExcludeAtypical := false // bool | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical (optional) (default to false)
    filtersAssemblyVersion := openapiclient.v2AssemblyDatasetDescriptorsFilterAssemblyVersion("current") // V2AssemblyDatasetDescriptorsFilterAssemblyVersion | Return all assemblies, including replaced and suppressed, or only current assemblies (optional) (default to "current")
    filtersAssemblyLevel := []openapiclient.V2reportsAssemblyLevel{openapiclient.v2reportsAssemblyLevel("chromosome")} // []V2reportsAssemblyLevel | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. (optional)
    filtersFirstReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or after the specified date By default, do not filter. (optional)
    filtersLastReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or before to the specified date By default, do not filter. (optional)
    filtersSearchText := []string{"Inner_example"} // []string | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter (optional)
    filtersIsMetagenomeDerived := openapiclient.v2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter("METAGENOME_DERIVED_UNSET") // V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter |  (optional) (default to "METAGENOME_DERIVED_UNSET")
    filtersIsTypeMaterial := false // bool | If true, include only type materials (optional) (default to false)
    filtersIsIctvExemplar := false // bool | If true, include only ICTV Exemplars (optional) (default to false)
    filtersExcludeMultiIsolate := false // bool | If true, exclude large multi-isolate projects (optional) (default to false)
    filtersTypeMaterialCategory := openapiclient.v2AssemblyDatasetDescriptorsFilterTypeMaterialCategory("NONE") // V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory |  (optional) (default to "NONE")
    taxExactMatch := true // bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. (optional) (default to false)
    tableFields := []string{"Inner_example"} // []string |  (optional)
    returnedContent := openapiclient.v2AssemblyDatasetReportsRequestContentType("COMPLETE") // V2AssemblyDatasetReportsRequestContentType | Return either assembly accessions, or complete assembly reports (optional) (default to "COMPLETE")
    pageSize := int32(56) // int32 | The maximum number of genome assembly reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyDatasetReportsRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyDatasetReportsRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)
    sortField := "sortField_example" // string |  (optional)
    sortDirection := openapiclient.v2SortDirection("SORT_DIRECTION_UNSPECIFIED") // V2SortDirection |  (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeDatasetReportsByBiosampleId(context.Background(), biosampleIds).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).FiltersIsMetagenomeDerived(filtersIsMetagenomeDerived).FiltersIsTypeMaterial(filtersIsTypeMaterial).FiltersIsIctvExemplar(filtersIsIctvExemplar).FiltersExcludeMultiIsolate(filtersExcludeMultiIsolate).FiltersTypeMaterialCategory(filtersTypeMaterialCategory).TaxExactMatch(taxExactMatch).TableFields(tableFields).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).SortField(sortField).SortDirection(sortDirection).IncludeTabularHeader(includeTabularHeader).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeDatasetReportsByBiosampleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeDatasetReportsByBiosampleId`: V2reportsAssemblyDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeDatasetReportsByBiosampleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**biosampleIds** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeDatasetReportsByBiosampleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtersReferenceOnly** | **bool** | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [default to false]
 **filtersAssemblySource** | [**V2AssemblyDatasetDescriptorsFilterAssemblySource**](V2AssemblyDatasetDescriptorsFilterAssemblySource.md) | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies | [default to &quot;all&quot;]
 **filtersHasAnnotation** | **bool** | Return only annotated genome assemblies | [default to false]
 **filtersExcludePairedReports** | **bool** | For paired (GCA/GCF) records, only return the primary record | [default to false]
 **filtersExcludeAtypical** | **bool** | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical | [default to false]
 **filtersAssemblyVersion** | [**V2AssemblyDatasetDescriptorsFilterAssemblyVersion**](V2AssemblyDatasetDescriptorsFilterAssemblyVersion.md) | Return all assemblies, including replaced and suppressed, or only current assemblies | [default to &quot;current&quot;]
 **filtersAssemblyLevel** | [**[]V2reportsAssemblyLevel**](V2reportsAssemblyLevel.md) | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. | 
 **filtersFirstReleaseDate** | **time.Time** | Only return genome assemblies that were released on or after the specified date By default, do not filter. | 
 **filtersLastReleaseDate** | **time.Time** | Only return genome assemblies that were released on or before to the specified date By default, do not filter. | 
 **filtersSearchText** | **[]string** | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter | 
 **filtersIsMetagenomeDerived** | [**V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter**](V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter.md) |  | [default to &quot;METAGENOME_DERIVED_UNSET&quot;]
 **filtersIsTypeMaterial** | **bool** | If true, include only type materials | [default to false]
 **filtersIsIctvExemplar** | **bool** | If true, include only ICTV Exemplars | [default to false]
 **filtersExcludeMultiIsolate** | **bool** | If true, exclude large multi-isolate projects | [default to false]
 **filtersTypeMaterialCategory** | [**V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory**](V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory.md) |  | [default to &quot;NONE&quot;]
 **taxExactMatch** | **bool** | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. | [default to false]
 **tableFields** | **[]string** |  | 
 **returnedContent** | [**V2AssemblyDatasetReportsRequestContentType**](V2AssemblyDatasetReportsRequestContentType.md) | Return either assembly accessions, or complete assembly reports | [default to &quot;COMPLETE&quot;]
 **pageSize** | **int32** | The maximum number of genome assembly reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyDatasetReportsRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyDatasetReportsRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **sortField** | **string** |  | 
 **sortDirection** | [**V2SortDirection**](V2SortDirection.md) |  | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]

### Return type

[**V2reportsAssemblyDataReportPage**](V2reportsAssemblyDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeDatasetReportsByTaxon

> V2reportsAssemblyDataReportPage GenomeDatasetReportsByTaxon(ctx, taxons).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).FiltersIsMetagenomeDerived(filtersIsMetagenomeDerived).FiltersIsTypeMaterial(filtersIsTypeMaterial).FiltersIsIctvExemplar(filtersIsIctvExemplar).FiltersExcludeMultiIsolate(filtersExcludeMultiIsolate).FiltersTypeMaterialCategory(filtersTypeMaterialCategory).TaxExactMatch(taxExactMatch).TableFields(tableFields).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).SortField(sortField).SortDirection(sortDirection).IncludeTabularHeader(includeTabularHeader).Execute()

Get dataset reports by taxons



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    taxons := []string{"Inner_example"} // []string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    filtersReferenceOnly := true // bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional) (default to false)
    filtersAssemblySource := openapiclient.v2AssemblyDatasetDescriptorsFilterAssemblySource("all") // V2AssemblyDatasetDescriptorsFilterAssemblySource | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies (optional) (default to "all")
    filtersHasAnnotation := true // bool | Return only annotated genome assemblies (optional) (default to false)
    filtersExcludePairedReports := false // bool | For paired (GCA/GCF) records, only return the primary record (optional) (default to false)
    filtersExcludeAtypical := false // bool | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical (optional) (default to false)
    filtersAssemblyVersion := openapiclient.v2AssemblyDatasetDescriptorsFilterAssemblyVersion("current") // V2AssemblyDatasetDescriptorsFilterAssemblyVersion | Return all assemblies, including replaced and suppressed, or only current assemblies (optional) (default to "current")
    filtersAssemblyLevel := []openapiclient.V2reportsAssemblyLevel{openapiclient.v2reportsAssemblyLevel("chromosome")} // []V2reportsAssemblyLevel | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. (optional)
    filtersFirstReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or after the specified date By default, do not filter. (optional)
    filtersLastReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or before to the specified date By default, do not filter. (optional)
    filtersSearchText := []string{"Inner_example"} // []string | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter (optional)
    filtersIsMetagenomeDerived := openapiclient.v2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter("METAGENOME_DERIVED_UNSET") // V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter |  (optional) (default to "METAGENOME_DERIVED_UNSET")
    filtersIsTypeMaterial := false // bool | If true, include only type materials (optional) (default to false)
    filtersIsIctvExemplar := false // bool | If true, include only ICTV Exemplars (optional) (default to false)
    filtersExcludeMultiIsolate := false // bool | If true, exclude large multi-isolate projects (optional) (default to false)
    filtersTypeMaterialCategory := openapiclient.v2AssemblyDatasetDescriptorsFilterTypeMaterialCategory("NONE") // V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory |  (optional) (default to "NONE")
    taxExactMatch := true // bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. (optional) (default to false)
    tableFields := []string{"Inner_example"} // []string |  (optional)
    returnedContent := openapiclient.v2AssemblyDatasetReportsRequestContentType("COMPLETE") // V2AssemblyDatasetReportsRequestContentType | Return either assembly accessions, or complete assembly reports (optional) (default to "COMPLETE")
    pageSize := int32(56) // int32 | The maximum number of genome assembly reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyDatasetReportsRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyDatasetReportsRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)
    sortField := "sortField_example" // string |  (optional)
    sortDirection := openapiclient.v2SortDirection("SORT_DIRECTION_UNSPECIFIED") // V2SortDirection |  (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeDatasetReportsByTaxon(context.Background(), taxons).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).FiltersIsMetagenomeDerived(filtersIsMetagenomeDerived).FiltersIsTypeMaterial(filtersIsTypeMaterial).FiltersIsIctvExemplar(filtersIsIctvExemplar).FiltersExcludeMultiIsolate(filtersExcludeMultiIsolate).FiltersTypeMaterialCategory(filtersTypeMaterialCategory).TaxExactMatch(taxExactMatch).TableFields(tableFields).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).SortField(sortField).SortDirection(sortDirection).IncludeTabularHeader(includeTabularHeader).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeDatasetReportsByTaxon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeDatasetReportsByTaxon`: V2reportsAssemblyDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeDatasetReportsByTaxon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxons** | [**[]string**](string.md) | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeDatasetReportsByTaxonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtersReferenceOnly** | **bool** | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [default to false]
 **filtersAssemblySource** | [**V2AssemblyDatasetDescriptorsFilterAssemblySource**](V2AssemblyDatasetDescriptorsFilterAssemblySource.md) | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies | [default to &quot;all&quot;]
 **filtersHasAnnotation** | **bool** | Return only annotated genome assemblies | [default to false]
 **filtersExcludePairedReports** | **bool** | For paired (GCA/GCF) records, only return the primary record | [default to false]
 **filtersExcludeAtypical** | **bool** | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical | [default to false]
 **filtersAssemblyVersion** | [**V2AssemblyDatasetDescriptorsFilterAssemblyVersion**](V2AssemblyDatasetDescriptorsFilterAssemblyVersion.md) | Return all assemblies, including replaced and suppressed, or only current assemblies | [default to &quot;current&quot;]
 **filtersAssemblyLevel** | [**[]V2reportsAssemblyLevel**](V2reportsAssemblyLevel.md) | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. | 
 **filtersFirstReleaseDate** | **time.Time** | Only return genome assemblies that were released on or after the specified date By default, do not filter. | 
 **filtersLastReleaseDate** | **time.Time** | Only return genome assemblies that were released on or before to the specified date By default, do not filter. | 
 **filtersSearchText** | **[]string** | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter | 
 **filtersIsMetagenomeDerived** | [**V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter**](V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter.md) |  | [default to &quot;METAGENOME_DERIVED_UNSET&quot;]
 **filtersIsTypeMaterial** | **bool** | If true, include only type materials | [default to false]
 **filtersIsIctvExemplar** | **bool** | If true, include only ICTV Exemplars | [default to false]
 **filtersExcludeMultiIsolate** | **bool** | If true, exclude large multi-isolate projects | [default to false]
 **filtersTypeMaterialCategory** | [**V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory**](V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory.md) |  | [default to &quot;NONE&quot;]
 **taxExactMatch** | **bool** | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. | [default to false]
 **tableFields** | **[]string** |  | 
 **returnedContent** | [**V2AssemblyDatasetReportsRequestContentType**](V2AssemblyDatasetReportsRequestContentType.md) | Return either assembly accessions, or complete assembly reports | [default to &quot;COMPLETE&quot;]
 **pageSize** | **int32** | The maximum number of genome assembly reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyDatasetReportsRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyDatasetReportsRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **sortField** | **string** |  | 
 **sortDirection** | [**V2SortDirection**](V2SortDirection.md) |  | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]

### Return type

[**V2reportsAssemblyDataReportPage**](V2reportsAssemblyDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeDatasetReportsByWgs

> V2reportsAssemblyDataReportPage GenomeDatasetReportsByWgs(ctx, wgsAccessions).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).FiltersIsMetagenomeDerived(filtersIsMetagenomeDerived).FiltersIsTypeMaterial(filtersIsTypeMaterial).FiltersIsIctvExemplar(filtersIsIctvExemplar).FiltersExcludeMultiIsolate(filtersExcludeMultiIsolate).FiltersTypeMaterialCategory(filtersTypeMaterialCategory).TaxExactMatch(taxExactMatch).TableFields(tableFields).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).SortField(sortField).SortDirection(sortDirection).IncludeTabularHeader(includeTabularHeader).Execute()

Get dataset reports by wgs accession



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    wgsAccessions := []string{"Inner_example"} // []string | 
    filtersReferenceOnly := true // bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional) (default to false)
    filtersAssemblySource := openapiclient.v2AssemblyDatasetDescriptorsFilterAssemblySource("all") // V2AssemblyDatasetDescriptorsFilterAssemblySource | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies (optional) (default to "all")
    filtersHasAnnotation := true // bool | Return only annotated genome assemblies (optional) (default to false)
    filtersExcludePairedReports := false // bool | For paired (GCA/GCF) records, only return the primary record (optional) (default to false)
    filtersExcludeAtypical := false // bool | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical (optional) (default to false)
    filtersAssemblyVersion := openapiclient.v2AssemblyDatasetDescriptorsFilterAssemblyVersion("current") // V2AssemblyDatasetDescriptorsFilterAssemblyVersion | Return all assemblies, including replaced and suppressed, or only current assemblies (optional) (default to "current")
    filtersAssemblyLevel := []openapiclient.V2reportsAssemblyLevel{openapiclient.v2reportsAssemblyLevel("chromosome")} // []V2reportsAssemblyLevel | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. (optional)
    filtersFirstReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or after the specified date By default, do not filter. (optional)
    filtersLastReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or before to the specified date By default, do not filter. (optional)
    filtersSearchText := []string{"Inner_example"} // []string | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter (optional)
    filtersIsMetagenomeDerived := openapiclient.v2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter("METAGENOME_DERIVED_UNSET") // V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter |  (optional) (default to "METAGENOME_DERIVED_UNSET")
    filtersIsTypeMaterial := false // bool | If true, include only type materials (optional) (default to false)
    filtersIsIctvExemplar := false // bool | If true, include only ICTV Exemplars (optional) (default to false)
    filtersExcludeMultiIsolate := false // bool | If true, exclude large multi-isolate projects (optional) (default to false)
    filtersTypeMaterialCategory := openapiclient.v2AssemblyDatasetDescriptorsFilterTypeMaterialCategory("NONE") // V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory |  (optional) (default to "NONE")
    taxExactMatch := true // bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. (optional) (default to false)
    tableFields := []string{"Inner_example"} // []string |  (optional)
    returnedContent := openapiclient.v2AssemblyDatasetReportsRequestContentType("COMPLETE") // V2AssemblyDatasetReportsRequestContentType | Return either assembly accessions, or complete assembly reports (optional) (default to "COMPLETE")
    pageSize := int32(56) // int32 | The maximum number of genome assembly reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyDatasetReportsRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyDatasetReportsRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)
    sortField := "sortField_example" // string |  (optional)
    sortDirection := openapiclient.v2SortDirection("SORT_DIRECTION_UNSPECIFIED") // V2SortDirection |  (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeDatasetReportsByWgs(context.Background(), wgsAccessions).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).FiltersIsMetagenomeDerived(filtersIsMetagenomeDerived).FiltersIsTypeMaterial(filtersIsTypeMaterial).FiltersIsIctvExemplar(filtersIsIctvExemplar).FiltersExcludeMultiIsolate(filtersExcludeMultiIsolate).FiltersTypeMaterialCategory(filtersTypeMaterialCategory).TaxExactMatch(taxExactMatch).TableFields(tableFields).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).SortField(sortField).SortDirection(sortDirection).IncludeTabularHeader(includeTabularHeader).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeDatasetReportsByWgs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeDatasetReportsByWgs`: V2reportsAssemblyDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeDatasetReportsByWgs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wgsAccessions** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeDatasetReportsByWgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtersReferenceOnly** | **bool** | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [default to false]
 **filtersAssemblySource** | [**V2AssemblyDatasetDescriptorsFilterAssemblySource**](V2AssemblyDatasetDescriptorsFilterAssemblySource.md) | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies | [default to &quot;all&quot;]
 **filtersHasAnnotation** | **bool** | Return only annotated genome assemblies | [default to false]
 **filtersExcludePairedReports** | **bool** | For paired (GCA/GCF) records, only return the primary record | [default to false]
 **filtersExcludeAtypical** | **bool** | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical | [default to false]
 **filtersAssemblyVersion** | [**V2AssemblyDatasetDescriptorsFilterAssemblyVersion**](V2AssemblyDatasetDescriptorsFilterAssemblyVersion.md) | Return all assemblies, including replaced and suppressed, or only current assemblies | [default to &quot;current&quot;]
 **filtersAssemblyLevel** | [**[]V2reportsAssemblyLevel**](V2reportsAssemblyLevel.md) | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. | 
 **filtersFirstReleaseDate** | **time.Time** | Only return genome assemblies that were released on or after the specified date By default, do not filter. | 
 **filtersLastReleaseDate** | **time.Time** | Only return genome assemblies that were released on or before to the specified date By default, do not filter. | 
 **filtersSearchText** | **[]string** | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter | 
 **filtersIsMetagenomeDerived** | [**V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter**](V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter.md) |  | [default to &quot;METAGENOME_DERIVED_UNSET&quot;]
 **filtersIsTypeMaterial** | **bool** | If true, include only type materials | [default to false]
 **filtersIsIctvExemplar** | **bool** | If true, include only ICTV Exemplars | [default to false]
 **filtersExcludeMultiIsolate** | **bool** | If true, exclude large multi-isolate projects | [default to false]
 **filtersTypeMaterialCategory** | [**V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory**](V2AssemblyDatasetDescriptorsFilterTypeMaterialCategory.md) |  | [default to &quot;NONE&quot;]
 **taxExactMatch** | **bool** | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. | [default to false]
 **tableFields** | **[]string** |  | 
 **returnedContent** | [**V2AssemblyDatasetReportsRequestContentType**](V2AssemblyDatasetReportsRequestContentType.md) | Return either assembly accessions, or complete assembly reports | [default to &quot;COMPLETE&quot;]
 **pageSize** | **int32** | The maximum number of genome assembly reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyDatasetReportsRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyDatasetReportsRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **sortField** | **string** |  | 
 **sortDirection** | [**V2SortDirection**](V2SortDirection.md) |  | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]

### Return type

[**V2reportsAssemblyDataReportPage**](V2reportsAssemblyDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeDownloadSummary

> V2DownloadSummary GenomeDownloadSummary(ctx, accessions).Chromosomes(chromosomes).IncludeAnnotationType(includeAnnotationType).ExpDebugValues(expDebugValues).Execute()

Preview genome dataset download



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accessions := []string{"Inner_example"} // []string | NCBI genome assembly accessions
    chromosomes := []string{"Inner_example"} // []string | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome ('Un'). The filter only applies to fasta sequence. (optional)
    includeAnnotationType := []openapiclient.V2AnnotationForAssemblyType{openapiclient.v2AnnotationForAssemblyType("GENOME_GFF")} // []V2AnnotationForAssemblyType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    expDebugValues := "expDebugValues_example" // string | Supports debugging, e.g. by controlling data download speeds (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeDownloadSummary(context.Background(), accessions).Chromosomes(chromosomes).IncludeAnnotationType(includeAnnotationType).ExpDebugValues(expDebugValues).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeDownloadSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeDownloadSummary`: V2DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeDownloadSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | NCBI genome assembly accessions | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeDownloadSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chromosomes** | **[]string** | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | 
 **includeAnnotationType** | [**[]V2AnnotationForAssemblyType**](V2AnnotationForAssemblyType.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
 **expDebugValues** | **string** | Supports debugging, e.g. by controlling data download speeds | 

### Return type

[**V2DownloadSummary**](V2DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeDownloadSummaryByPost

> V2DownloadSummary GenomeDownloadSummaryByPost(ctx).V2AssemblyDatasetRequest(v2AssemblyDatasetRequest).Execute()

Preview genome dataset download by POST



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2AssemblyDatasetRequest := *openapiclient.NewV2AssemblyDatasetRequest() // V2AssemblyDatasetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeDownloadSummaryByPost(context.Background()).V2AssemblyDatasetRequest(v2AssemblyDatasetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeDownloadSummaryByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeDownloadSummaryByPost`: V2DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeDownloadSummaryByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenomeDownloadSummaryByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2AssemblyDatasetRequest** | [**V2AssemblyDatasetRequest**](V2AssemblyDatasetRequest.md) |  | 

### Return type

[**V2DownloadSummary**](V2DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeLinksByAccession

> V2AssemblyLinksReply GenomeLinksByAccession(ctx, accessions).Execute()

Get assembly links by accessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accessions := []string{"Inner_example"} // []string | NCBI genome assembly accessions, limited to 1000

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeLinksByAccession(context.Background(), accessions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeLinksByAccession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeLinksByAccession`: V2AssemblyLinksReply
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeLinksByAccession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | NCBI genome assembly accessions, limited to 1000 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeLinksByAccessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2AssemblyLinksReply**](V2AssemblyLinksReply.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeLinksByAccessionByPost

> V2AssemblyLinksReply GenomeLinksByAccessionByPost(ctx).V2AssemblyLinksRequest(v2AssemblyLinksRequest).Execute()

Get assembly links by accessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2AssemblyLinksRequest := *openapiclient.NewV2AssemblyLinksRequest() // V2AssemblyLinksRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeLinksByAccessionByPost(context.Background()).V2AssemblyLinksRequest(v2AssemblyLinksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeLinksByAccessionByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeLinksByAccessionByPost`: V2AssemblyLinksReply
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeLinksByAccessionByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenomeLinksByAccessionByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2AssemblyLinksRequest** | [**V2AssemblyLinksRequest**](V2AssemblyLinksRequest.md) |  | 

### Return type

[**V2AssemblyLinksReply**](V2AssemblyLinksReply.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeSequenceReport

> V2SequenceReportPage GenomeSequenceReport(ctx, accession).Chromosomes(chromosomes).RoleFilters(roleFilters).TableFields(tableFields).CountAssemblyUnplaced(countAssemblyUnplaced).PageSize(pageSize).PageToken(pageToken).IncludeTabularHeader(includeTabularHeader).Execute()

Get sequence reports by accessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accession := "GCF_000001635.27" // string | 
    chromosomes := []string{"Inner_example"} // []string |  (optional)
    roleFilters := []string{"Inner_example"} // []string |  (optional)
    tableFields := []string{"Inner_example"} // []string |  (optional)
    countAssemblyUnplaced := true // bool |  (optional) (default to false)
    pageSize := int32(56) // int32 | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `GetSequenceReports` call with more than `page_size` results. Use this token, along with the previous `AssemblyMetadataRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader |  (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeSequenceReport(context.Background(), accession).Chromosomes(chromosomes).RoleFilters(roleFilters).TableFields(tableFields).CountAssemblyUnplaced(countAssemblyUnplaced).PageSize(pageSize).PageToken(pageToken).IncludeTabularHeader(includeTabularHeader).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeSequenceReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeSequenceReport`: V2SequenceReportPage
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeSequenceReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accession** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeSequenceReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chromosomes** | **[]string** |  | 
 **roleFilters** | **[]string** |  | 
 **tableFields** | **[]string** |  | 
 **countAssemblyUnplaced** | **bool** |  | [default to false]
 **pageSize** | **int32** | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;GetSequenceReports&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) |  | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]

### Return type

[**V2SequenceReportPage**](V2SequenceReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, text/tab-separated-values, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeSequenceReportByPost

> V2SequenceReportPage GenomeSequenceReportByPost(ctx).V2AssemblySequenceReportsRequest(v2AssemblySequenceReportsRequest).Execute()

Get sequence reports by accessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v2AssemblySequenceReportsRequest := *openapiclient.NewV2AssemblySequenceReportsRequest() // V2AssemblySequenceReportsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenomeApi.GenomeSequenceReportByPost(context.Background()).V2AssemblySequenceReportsRequest(v2AssemblySequenceReportsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeSequenceReportByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeSequenceReportByPost`: V2SequenceReportPage
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeSequenceReportByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenomeSequenceReportByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2AssemblySequenceReportsRequest** | [**V2AssemblySequenceReportsRequest**](V2AssemblySequenceReportsRequest.md) |  | 

### Return type

[**V2SequenceReportPage**](V2SequenceReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, text/tab-separated-values, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

