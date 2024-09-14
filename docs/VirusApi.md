# \VirusApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v2alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Sars2ProteinDownload**](VirusApi.md#Sars2ProteinDownload) | **Get** /virus/taxon/sars2/protein/{proteins}/download | Download SARS-CoV-2 protein and CDS datasets by protein name
[**Sars2ProteinDownloadPost**](VirusApi.md#Sars2ProteinDownloadPost) | **Post** /virus/taxon/sars2/protein/download | Download SARS-CoV-2 protein and CDS datasets by protein name by POST request
[**Sars2ProteinSummary**](VirusApi.md#Sars2ProteinSummary) | **Get** /virus/taxon/sars2/protein/{proteins} | Summary of SARS-CoV-2 protein and CDS datasets by protein name
[**Sars2ProteinSummaryByPost**](VirusApi.md#Sars2ProteinSummaryByPost) | **Post** /virus/taxon/sars2/protein | Summary of SARS-CoV-2 protein and CDS datasets by protein name
[**Sars2ProteinTable**](VirusApi.md#Sars2ProteinTable) | **Get** /virus/taxon/sars2/protein/{proteins}/table | Get SARS-CoV-2 protein metadata in a tabular format.
[**VirusAccessionAvailability**](VirusApi.md#VirusAccessionAvailability) | **Get** /virus/accession/{accessions}/check | Check available viruses by accession
[**VirusAccessionAvailabilityPost**](VirusApi.md#VirusAccessionAvailabilityPost) | **Post** /virus/check | Check available viruses by accession
[**VirusAnnotationReportsByAcessions**](VirusApi.md#VirusAnnotationReportsByAcessions) | **Get** /virus/accession/{accessions}/annotation_report | Get virus annotation report by accession
[**VirusAnnotationReportsByPost**](VirusApi.md#VirusAnnotationReportsByPost) | **Post** /virus/annotation_report | Get virus annotation report by POST
[**VirusAnnotationReportsByTaxon**](VirusApi.md#VirusAnnotationReportsByTaxon) | **Get** /virus/taxon/{taxon}/annotation_report | Get virus annotation report by taxon
[**VirusGenomeDownload**](VirusApi.md#VirusGenomeDownload) | **Get** /virus/taxon/{taxon}/genome/download | Download a virus genome dataset by taxon
[**VirusGenomeDownloadAccession**](VirusApi.md#VirusGenomeDownloadAccession) | **Get** /virus/accession/{accessions}/genome/download | Download a virus genome dataset by accession
[**VirusGenomeDownloadPost**](VirusApi.md#VirusGenomeDownloadPost) | **Post** /virus/genome/download | Get a virus genome dataset by post
[**VirusGenomeSummary**](VirusApi.md#VirusGenomeSummary) | **Get** /virus/taxon/{taxon}/genome | Get summary data for virus genomes by taxon
[**VirusGenomeSummaryByPost**](VirusApi.md#VirusGenomeSummaryByPost) | **Post** /virus/genome | Get summary data for virus genomes by post
[**VirusGenomeTable**](VirusApi.md#VirusGenomeTable) | **Get** /virus/taxon/{taxon}/genome/table | Get virus genome metadata in a tabular format.
[**VirusReportsByAcessions**](VirusApi.md#VirusReportsByAcessions) | **Get** /virus/accession/{accessions}/dataset_report | Get virus metadata by accession
[**VirusReportsByPost**](VirusApi.md#VirusReportsByPost) | **Post** /virus | Get virus metadata by POST
[**VirusReportsByTaxon**](VirusApi.md#VirusReportsByTaxon) | **Get** /virus/taxon/{taxon}/dataset_report | Get virus metadata by taxon



## Sars2ProteinDownload

> *os.File Sars2ProteinDownload(ctx, proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeSequence(includeSequence).AuxReport(auxReport).Filename(filename).Execute()

Download SARS-CoV-2 protein and CDS datasets by protein name



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
    proteins := []string{"Inner_example"} // []string | Which proteins to retrieve in the data package
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    includeSequence := []openapiclient.V2ViralSequenceType{openapiclient.v2ViralSequenceType("GENOME")} // []V2ViralSequenceType | Specify which sequence files to include in the download (optional)
    auxReport := []openapiclient.V2VirusDatasetReportType{openapiclient.v2VirusDatasetReportType("DATASET_REPORT")} // []V2VirusDatasetReportType | List additional reports to include with download. Data report is included by default. (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.Sars2ProteinDownload(context.Background(), proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeSequence(includeSequence).AuxReport(auxReport).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.Sars2ProteinDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Sars2ProteinDownload`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.Sars2ProteinDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proteins** | [**[]string**](string.md) | Which proteins to retrieve in the data package | 

### Other Parameters

Other parameters are passed through a pointer to a apiSars2ProteinDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **updatedSince** | **time.Time** |  | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **includeSequence** | [**[]V2ViralSequenceType**](V2ViralSequenceType.md) | Specify which sequence files to include in the download | 
 **auxReport** | [**[]V2VirusDatasetReportType**](V2VirusDatasetReportType.md) | List additional reports to include with download. Data report is included by default. | 
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


## Sars2ProteinDownloadPost

> *os.File Sars2ProteinDownloadPost(ctx).V2Sars2ProteinDatasetRequest(v2Sars2ProteinDatasetRequest).Filename(filename).Execute()

Download SARS-CoV-2 protein and CDS datasets by protein name by POST request



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
    v2Sars2ProteinDatasetRequest := *openapiclient.NewV2Sars2ProteinDatasetRequest() // V2Sars2ProteinDatasetRequest | 
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.Sars2ProteinDownloadPost(context.Background()).V2Sars2ProteinDatasetRequest(v2Sars2ProteinDatasetRequest).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.Sars2ProteinDownloadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Sars2ProteinDownloadPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.Sars2ProteinDownloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSars2ProteinDownloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2Sars2ProteinDatasetRequest** | [**V2Sars2ProteinDatasetRequest**](V2Sars2ProteinDatasetRequest.md) |  | 
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


## Sars2ProteinSummary

> V2DownloadSummary Sars2ProteinSummary(ctx, proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeSequence(includeSequence).AuxReport(auxReport).Execute()

Summary of SARS-CoV-2 protein and CDS datasets by protein name



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
    proteins := []string{"Inner_example"} // []string | Which proteins to retrieve in the data package
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    includeSequence := []openapiclient.V2ViralSequenceType{openapiclient.v2ViralSequenceType("GENOME")} // []V2ViralSequenceType | Specify which sequence files to include in the download (optional)
    auxReport := []openapiclient.V2VirusDatasetReportType{openapiclient.v2VirusDatasetReportType("DATASET_REPORT")} // []V2VirusDatasetReportType | List additional reports to include with download. Data report is included by default. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.Sars2ProteinSummary(context.Background(), proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeSequence(includeSequence).AuxReport(auxReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.Sars2ProteinSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Sars2ProteinSummary`: V2DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.Sars2ProteinSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proteins** | [**[]string**](string.md) | Which proteins to retrieve in the data package | 

### Other Parameters

Other parameters are passed through a pointer to a apiSars2ProteinSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **updatedSince** | **time.Time** |  | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **includeSequence** | [**[]V2ViralSequenceType**](V2ViralSequenceType.md) | Specify which sequence files to include in the download | 
 **auxReport** | [**[]V2VirusDatasetReportType**](V2VirusDatasetReportType.md) | List additional reports to include with download. Data report is included by default. | 

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


## Sars2ProteinSummaryByPost

> V2DownloadSummary Sars2ProteinSummaryByPost(ctx).V2Sars2ProteinDatasetRequest(v2Sars2ProteinDatasetRequest).Execute()

Summary of SARS-CoV-2 protein and CDS datasets by protein name



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
    v2Sars2ProteinDatasetRequest := *openapiclient.NewV2Sars2ProteinDatasetRequest() // V2Sars2ProteinDatasetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.Sars2ProteinSummaryByPost(context.Background()).V2Sars2ProteinDatasetRequest(v2Sars2ProteinDatasetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.Sars2ProteinSummaryByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Sars2ProteinSummaryByPost`: V2DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.Sars2ProteinSummaryByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSars2ProteinSummaryByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2Sars2ProteinDatasetRequest** | [**V2Sars2ProteinDatasetRequest**](V2Sars2ProteinDatasetRequest.md) |  | 

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


## Sars2ProteinTable

> V2TabularOutput Sars2ProteinTable(ctx, proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).TableFields(tableFields).IncludeSequence(includeSequence).AuxReport(auxReport).Format(format).Execute()

Get SARS-CoV-2 protein metadata in a tabular format.



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
    proteins := []string{"Inner_example"} // []string | Which proteins to retrieve in the data package
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    tableFields := []openapiclient.V2VirusTableField{openapiclient.v2VirusTableField("unspecified")} // []V2VirusTableField | Specify which fields to include in the tabular report (optional)
    includeSequence := []openapiclient.V2ViralSequenceType{openapiclient.v2ViralSequenceType("GENOME")} // []V2ViralSequenceType | Specify which sequence files to include in the download (optional)
    auxReport := []openapiclient.V2VirusDatasetReportType{openapiclient.v2VirusDatasetReportType("DATASET_REPORT")} // []V2VirusDatasetReportType | List additional reports to include with download. Data report is included by default. (optional)
    format := openapiclient.v2TableFormat("tsv") // V2TableFormat | Choose download format (tsv, csv or jsonl) (optional) (default to "tsv")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.Sars2ProteinTable(context.Background(), proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).TableFields(tableFields).IncludeSequence(includeSequence).AuxReport(auxReport).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.Sars2ProteinTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Sars2ProteinTable`: V2TabularOutput
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.Sars2ProteinTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proteins** | [**[]string**](string.md) | Which proteins to retrieve in the data package | 

### Other Parameters

Other parameters are passed through a pointer to a apiSars2ProteinTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **updatedSince** | **time.Time** |  | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **tableFields** | [**[]V2VirusTableField**](V2VirusTableField.md) | Specify which fields to include in the tabular report | 
 **includeSequence** | [**[]V2ViralSequenceType**](V2ViralSequenceType.md) | Specify which sequence files to include in the download | 
 **auxReport** | [**[]V2VirusDatasetReportType**](V2VirusDatasetReportType.md) | List additional reports to include with download. Data report is included by default. | 
 **format** | [**V2TableFormat**](V2TableFormat.md) | Choose download format (tsv, csv or jsonl) | [default to &quot;tsv&quot;]

### Return type

[**V2TabularOutput**](V2TabularOutput.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusAccessionAvailability

> V2VirusAvailability VirusAccessionAvailability(ctx, accessions).Execute()

Check available viruses by accession



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
    accessions := []string{"Inner_example"} // []string | virus accessions

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusAccessionAvailability(context.Background(), accessions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusAccessionAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusAccessionAvailability`: V2VirusAvailability
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusAccessionAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | virus accessions | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusAccessionAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2VirusAvailability**](V2VirusAvailability.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusAccessionAvailabilityPost

> V2VirusAvailability VirusAccessionAvailabilityPost(ctx).V2VirusAvailabilityRequest(v2VirusAvailabilityRequest).Execute()

Check available viruses by accession



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
    v2VirusAvailabilityRequest := *openapiclient.NewV2VirusAvailabilityRequest() // V2VirusAvailabilityRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusAccessionAvailabilityPost(context.Background()).V2VirusAvailabilityRequest(v2VirusAvailabilityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusAccessionAvailabilityPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusAccessionAvailabilityPost`: V2VirusAvailability
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusAccessionAvailabilityPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirusAccessionAvailabilityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2VirusAvailabilityRequest** | [**V2VirusAvailabilityRequest**](V2VirusAvailabilityRequest.md) |  | 

### Return type

[**V2VirusAvailability**](V2VirusAvailability.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusAnnotationReportsByAcessions

> V2reportsVirusAnnotationReportPage VirusAnnotationReportsByAcessions(ctx, accessions).FilterRefseqOnly(filterRefseqOnly).FilterReleasedSince(filterReleasedSince).FilterUpdatedSince(filterUpdatedSince).FilterHost(filterHost).FilterPangolinClassification(filterPangolinClassification).FilterGeoLocation(filterGeoLocation).FilterCompleteOnly(filterCompleteOnly).TableFields(tableFields).PageSize(pageSize).PageToken(pageToken).Execute()

Get virus annotation report by accession



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
    accessions := []string{"Inner_example"} // []string | genome sequence accessions
    filterRefseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    filterReleasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    filterUpdatedSince := time.Now() // time.Time |  (optional)
    filterHost := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    filterPangolinClassification := "filterPangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    filterGeoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    filterCompleteOnly := true // bool | only include complete genomes. (optional) (default to false)
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    pageSize := int32(56) // int32 | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from a `GetVirusDataReports` call with more than `page_size` results. Use this token, along with the previous `VirusDataReportRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusAnnotationReportsByAcessions(context.Background(), accessions).FilterRefseqOnly(filterRefseqOnly).FilterReleasedSince(filterReleasedSince).FilterUpdatedSince(filterUpdatedSince).FilterHost(filterHost).FilterPangolinClassification(filterPangolinClassification).FilterGeoLocation(filterGeoLocation).FilterCompleteOnly(filterCompleteOnly).TableFields(tableFields).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusAnnotationReportsByAcessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusAnnotationReportsByAcessions`: V2reportsVirusAnnotationReportPage
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusAnnotationReportsByAcessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | genome sequence accessions | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusAnnotationReportsByAcessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterRefseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **filterReleasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **filterUpdatedSince** | **time.Time** |  | 
 **filterHost** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **filterPangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **filterGeoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **filterCompleteOnly** | **bool** | only include complete genomes. | [default to false]
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **pageSize** | **int32** | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from a &#x60;GetVirusDataReports&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;VirusDataReportRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V2reportsVirusAnnotationReportPage**](V2reportsVirusAnnotationReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusAnnotationReportsByPost

> V2reportsVirusAnnotationReportPage VirusAnnotationReportsByPost(ctx).V2VirusAnnotationReportRequest(v2VirusAnnotationReportRequest).Execute()

Get virus annotation report by POST



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
    v2VirusAnnotationReportRequest := *openapiclient.NewV2VirusAnnotationReportRequest() // V2VirusAnnotationReportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusAnnotationReportsByPost(context.Background()).V2VirusAnnotationReportRequest(v2VirusAnnotationReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusAnnotationReportsByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusAnnotationReportsByPost`: V2reportsVirusAnnotationReportPage
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusAnnotationReportsByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirusAnnotationReportsByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2VirusAnnotationReportRequest** | [**V2VirusAnnotationReportRequest**](V2VirusAnnotationReportRequest.md) |  | 

### Return type

[**V2reportsVirusAnnotationReportPage**](V2reportsVirusAnnotationReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusAnnotationReportsByTaxon

> V2reportsVirusAnnotationReportPage VirusAnnotationReportsByTaxon(ctx, taxon).FilterRefseqOnly(filterRefseqOnly).FilterReleasedSince(filterReleasedSince).FilterUpdatedSince(filterUpdatedSince).FilterHost(filterHost).FilterPangolinClassification(filterPangolinClassification).FilterGeoLocation(filterGeoLocation).FilterCompleteOnly(filterCompleteOnly).TableFields(tableFields).PageSize(pageSize).PageToken(pageToken).Execute()

Get virus annotation report by taxon



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
    taxon := "1335626" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    filterRefseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    filterReleasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    filterUpdatedSince := time.Now() // time.Time |  (optional)
    filterHost := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    filterPangolinClassification := "filterPangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    filterGeoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    filterCompleteOnly := true // bool | only include complete genomes. (optional) (default to false)
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    pageSize := int32(56) // int32 | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from a `GetVirusDataReports` call with more than `page_size` results. Use this token, along with the previous `VirusDataReportRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusAnnotationReportsByTaxon(context.Background(), taxon).FilterRefseqOnly(filterRefseqOnly).FilterReleasedSince(filterReleasedSince).FilterUpdatedSince(filterUpdatedSince).FilterHost(filterHost).FilterPangolinClassification(filterPangolinClassification).FilterGeoLocation(filterGeoLocation).FilterCompleteOnly(filterCompleteOnly).TableFields(tableFields).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusAnnotationReportsByTaxon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusAnnotationReportsByTaxon`: V2reportsVirusAnnotationReportPage
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusAnnotationReportsByTaxon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusAnnotationReportsByTaxonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterRefseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **filterReleasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **filterUpdatedSince** | **time.Time** |  | 
 **filterHost** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **filterPangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **filterGeoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **filterCompleteOnly** | **bool** | only include complete genomes. | [default to false]
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **pageSize** | **int32** | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from a &#x60;GetVirusDataReports&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;VirusDataReportRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V2reportsVirusAnnotationReportPage**](V2reportsVirusAnnotationReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusGenomeDownload

> *os.File VirusGenomeDownload(ctx, taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeSequence(includeSequence).AuxReport(auxReport).UsePsg(usePsg).Filename(filename).Execute()

Download a virus genome dataset by taxon



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
    taxon := "1335626" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolinClassification := "pangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    includeSequence := []openapiclient.V2ViralSequenceType{openapiclient.v2ViralSequenceType("GENOME")} // []V2ViralSequenceType | specify which sequence files to include in the download (optional)
    auxReport := []openapiclient.V2VirusDatasetReportType{openapiclient.v2VirusDatasetReportType("DATASET_REPORT")} // []V2VirusDatasetReportType | list additional reports to include with download. Data report is included by default. (optional)
    usePsg := true // bool | Experimental approach to retrieving sequence data. (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusGenomeDownload(context.Background(), taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeSequence(includeSequence).AuxReport(auxReport).UsePsg(usePsg).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeDownload`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **updatedSince** | **time.Time** |  | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **pangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **includeSequence** | [**[]V2ViralSequenceType**](V2ViralSequenceType.md) | specify which sequence files to include in the download | 
 **auxReport** | [**[]V2VirusDatasetReportType**](V2VirusDatasetReportType.md) | list additional reports to include with download. Data report is included by default. | 
 **usePsg** | **bool** | Experimental approach to retrieving sequence data. | 
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


## VirusGenomeDownloadAccession

> *os.File VirusGenomeDownloadAccession(ctx, accessions).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeSequence(includeSequence).AuxReport(auxReport).UsePsg(usePsg).Filename(filename).Execute()

Download a virus genome dataset by accession



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
    accessions := []string{"Inner_example"} // []string | genome sequence accessions
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolinClassification := "pangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    includeSequence := []openapiclient.V2ViralSequenceType{openapiclient.v2ViralSequenceType("GENOME")} // []V2ViralSequenceType | specify which sequence files to include in the download (optional)
    auxReport := []openapiclient.V2VirusDatasetReportType{openapiclient.v2VirusDatasetReportType("DATASET_REPORT")} // []V2VirusDatasetReportType | list additional reports to include with download. Data report is included by default. (optional)
    usePsg := true // bool | Experimental approach to retrieving sequence data. (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusGenomeDownloadAccession(context.Background(), accessions).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeSequence(includeSequence).AuxReport(auxReport).UsePsg(usePsg).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeDownloadAccession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeDownloadAccession`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeDownloadAccession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | genome sequence accessions | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeDownloadAccessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **updatedSince** | **time.Time** |  | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **pangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **includeSequence** | [**[]V2ViralSequenceType**](V2ViralSequenceType.md) | specify which sequence files to include in the download | 
 **auxReport** | [**[]V2VirusDatasetReportType**](V2VirusDatasetReportType.md) | list additional reports to include with download. Data report is included by default. | 
 **usePsg** | **bool** | Experimental approach to retrieving sequence data. | 
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


## VirusGenomeDownloadPost

> *os.File VirusGenomeDownloadPost(ctx).V2VirusDatasetRequest(v2VirusDatasetRequest).Filename(filename).Execute()

Get a virus genome dataset by post



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
    v2VirusDatasetRequest := *openapiclient.NewV2VirusDatasetRequest() // V2VirusDatasetRequest | 
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusGenomeDownloadPost(context.Background()).V2VirusDatasetRequest(v2VirusDatasetRequest).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeDownloadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeDownloadPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeDownloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeDownloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2VirusDatasetRequest** | [**V2VirusDatasetRequest**](V2VirusDatasetRequest.md) |  | 
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


## VirusGenomeSummary

> V2DownloadSummary VirusGenomeSummary(ctx, taxon).Accessions(accessions).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeSequence(includeSequence).AuxReport(auxReport).Execute()

Get summary data for virus genomes by taxon



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
    taxon := "1335626" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    accessions := []string{"Inner_example"} // []string | genome sequence accessions (optional)
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolinClassification := "pangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    includeSequence := []openapiclient.V2ViralSequenceType{openapiclient.v2ViralSequenceType("GENOME")} // []V2ViralSequenceType | specify which sequence files to include in the download (optional)
    auxReport := []openapiclient.V2VirusDatasetReportType{openapiclient.v2VirusDatasetReportType("DATASET_REPORT")} // []V2VirusDatasetReportType | list additional reports to include with download. Data report is included by default. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusGenomeSummary(context.Background(), taxon).Accessions(accessions).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeSequence(includeSequence).AuxReport(auxReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeSummary`: V2DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessions** | **[]string** | genome sequence accessions | 
 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **updatedSince** | **time.Time** |  | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **pangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **includeSequence** | [**[]V2ViralSequenceType**](V2ViralSequenceType.md) | specify which sequence files to include in the download | 
 **auxReport** | [**[]V2VirusDatasetReportType**](V2VirusDatasetReportType.md) | list additional reports to include with download. Data report is included by default. | 

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


## VirusGenomeSummaryByPost

> V2DownloadSummary VirusGenomeSummaryByPost(ctx).V2VirusDatasetRequest(v2VirusDatasetRequest).Execute()

Get summary data for virus genomes by post



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
    v2VirusDatasetRequest := *openapiclient.NewV2VirusDatasetRequest() // V2VirusDatasetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusGenomeSummaryByPost(context.Background()).V2VirusDatasetRequest(v2VirusDatasetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeSummaryByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeSummaryByPost`: V2DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeSummaryByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeSummaryByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2VirusDatasetRequest** | [**V2VirusDatasetRequest**](V2VirusDatasetRequest.md) |  | 

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


## VirusGenomeTable

> V2TabularOutput VirusGenomeTable(ctx, taxon).Accessions(accessions).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).TableFields(tableFields).IncludeSequence(includeSequence).AuxReport(auxReport).Format(format).Execute()

Get virus genome metadata in a tabular format.



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
    taxon := "1335626" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    accessions := []string{"Inner_example"} // []string | genome sequence accessions (optional)
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolinClassification := "pangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    tableFields := []openapiclient.V2VirusTableField{openapiclient.v2VirusTableField("unspecified")} // []V2VirusTableField | Specify which fields to include in the tabular report (optional)
    includeSequence := []openapiclient.V2ViralSequenceType{openapiclient.v2ViralSequenceType("GENOME")} // []V2ViralSequenceType | specify which sequence files to include in the download (optional)
    auxReport := []openapiclient.V2VirusDatasetReportType{openapiclient.v2VirusDatasetReportType("DATASET_REPORT")} // []V2VirusDatasetReportType | list additional reports to include with download. Data report is included by default. (optional)
    format := openapiclient.v2TableFormat("tsv") // V2TableFormat | Choose download format (tsv, csv or jsonl) (optional) (default to "tsv")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusGenomeTable(context.Background(), taxon).Accessions(accessions).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).TableFields(tableFields).IncludeSequence(includeSequence).AuxReport(auxReport).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeTable`: V2TabularOutput
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessions** | **[]string** | genome sequence accessions | 
 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **updatedSince** | **time.Time** |  | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **pangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **tableFields** | [**[]V2VirusTableField**](V2VirusTableField.md) | Specify which fields to include in the tabular report | 
 **includeSequence** | [**[]V2ViralSequenceType**](V2ViralSequenceType.md) | specify which sequence files to include in the download | 
 **auxReport** | [**[]V2VirusDatasetReportType**](V2VirusDatasetReportType.md) | list additional reports to include with download. Data report is included by default. | 
 **format** | [**V2TableFormat**](V2TableFormat.md) | Choose download format (tsv, csv or jsonl) | [default to &quot;tsv&quot;]

### Return type

[**V2TabularOutput**](V2TabularOutput.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusReportsByAcessions

> V2reportsVirusDataReportPage VirusReportsByAcessions(ctx, accessions).FilterRefseqOnly(filterRefseqOnly).FilterAnnotatedOnly(filterAnnotatedOnly).FilterReleasedSince(filterReleasedSince).FilterUpdatedSince(filterUpdatedSince).FilterHost(filterHost).FilterPangolinClassification(filterPangolinClassification).FilterGeoLocation(filterGeoLocation).FilterCompleteOnly(filterCompleteOnly).ReturnedContent(returnedContent).TableFields(tableFields).PageSize(pageSize).PageToken(pageToken).Execute()

Get virus metadata by accession



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
    accessions := []string{"Inner_example"} // []string | genome sequence accessions
    filterRefseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    filterAnnotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    filterReleasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    filterUpdatedSince := time.Now() // time.Time |  (optional)
    filterHost := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    filterPangolinClassification := "filterPangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    filterGeoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    filterCompleteOnly := true // bool | only include complete genomes. (optional) (default to false)
    returnedContent := openapiclient.v2VirusDataReportRequestContentType("COMPLETE") // V2VirusDataReportRequestContentType | Return either virus genome accessions, or complete virus metadata (optional) (default to "COMPLETE")
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    pageSize := int32(56) // int32 | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from a `GetVirusDataReports` call with more than `page_size` results. Use this token, along with the previous `VirusDataReportRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusReportsByAcessions(context.Background(), accessions).FilterRefseqOnly(filterRefseqOnly).FilterAnnotatedOnly(filterAnnotatedOnly).FilterReleasedSince(filterReleasedSince).FilterUpdatedSince(filterUpdatedSince).FilterHost(filterHost).FilterPangolinClassification(filterPangolinClassification).FilterGeoLocation(filterGeoLocation).FilterCompleteOnly(filterCompleteOnly).ReturnedContent(returnedContent).TableFields(tableFields).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusReportsByAcessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusReportsByAcessions`: V2reportsVirusDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusReportsByAcessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | genome sequence accessions | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusReportsByAcessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterRefseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **filterAnnotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **filterReleasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **filterUpdatedSince** | **time.Time** |  | 
 **filterHost** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **filterPangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **filterGeoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **filterCompleteOnly** | **bool** | only include complete genomes. | [default to false]
 **returnedContent** | [**V2VirusDataReportRequestContentType**](V2VirusDataReportRequestContentType.md) | Return either virus genome accessions, or complete virus metadata | [default to &quot;COMPLETE&quot;]
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **pageSize** | **int32** | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from a &#x60;GetVirusDataReports&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;VirusDataReportRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V2reportsVirusDataReportPage**](V2reportsVirusDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusReportsByPost

> V2reportsVirusDataReportPage VirusReportsByPost(ctx).V2VirusDataReportRequest(v2VirusDataReportRequest).Execute()

Get virus metadata by POST



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
    v2VirusDataReportRequest := *openapiclient.NewV2VirusDataReportRequest() // V2VirusDataReportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusReportsByPost(context.Background()).V2VirusDataReportRequest(v2VirusDataReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusReportsByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusReportsByPost`: V2reportsVirusDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusReportsByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirusReportsByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2VirusDataReportRequest** | [**V2VirusDataReportRequest**](V2VirusDataReportRequest.md) |  | 

### Return type

[**V2reportsVirusDataReportPage**](V2reportsVirusDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusReportsByTaxon

> V2reportsVirusDataReportPage VirusReportsByTaxon(ctx, taxon).FilterRefseqOnly(filterRefseqOnly).FilterAnnotatedOnly(filterAnnotatedOnly).FilterReleasedSince(filterReleasedSince).FilterUpdatedSince(filterUpdatedSince).FilterHost(filterHost).FilterPangolinClassification(filterPangolinClassification).FilterGeoLocation(filterGeoLocation).FilterCompleteOnly(filterCompleteOnly).ReturnedContent(returnedContent).TableFields(tableFields).PageSize(pageSize).PageToken(pageToken).Execute()

Get virus metadata by taxon



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
    taxon := "1335626" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    filterRefseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    filterAnnotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    filterReleasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    filterUpdatedSince := time.Now() // time.Time |  (optional)
    filterHost := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    filterPangolinClassification := "filterPangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    filterGeoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    filterCompleteOnly := true // bool | only include complete genomes. (optional) (default to false)
    returnedContent := openapiclient.v2VirusDataReportRequestContentType("COMPLETE") // V2VirusDataReportRequestContentType | Return either virus genome accessions, or complete virus metadata (optional) (default to "COMPLETE")
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    pageSize := int32(56) // int32 | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from a `GetVirusDataReports` call with more than `page_size` results. Use this token, along with the previous `VirusDataReportRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirusApi.VirusReportsByTaxon(context.Background(), taxon).FilterRefseqOnly(filterRefseqOnly).FilterAnnotatedOnly(filterAnnotatedOnly).FilterReleasedSince(filterReleasedSince).FilterUpdatedSince(filterUpdatedSince).FilterHost(filterHost).FilterPangolinClassification(filterPangolinClassification).FilterGeoLocation(filterGeoLocation).FilterCompleteOnly(filterCompleteOnly).ReturnedContent(returnedContent).TableFields(tableFields).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusReportsByTaxon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusReportsByTaxon`: V2reportsVirusDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusReportsByTaxon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusReportsByTaxonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterRefseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **filterAnnotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **filterReleasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **filterUpdatedSince** | **time.Time** |  | 
 **filterHost** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **filterPangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **filterGeoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **filterCompleteOnly** | **bool** | only include complete genomes. | [default to false]
 **returnedContent** | [**V2VirusDataReportRequestContentType**](V2VirusDataReportRequestContentType.md) | Return either virus genome accessions, or complete virus metadata | [default to &quot;COMPLETE&quot;]
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **pageSize** | **int32** | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from a &#x60;GetVirusDataReports&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;VirusDataReportRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V2reportsVirusDataReportPage**](V2reportsVirusDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

