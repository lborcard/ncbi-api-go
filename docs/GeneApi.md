# \GeneApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v2alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadGenePackage**](GeneApi.md#DownloadGenePackage) | **Get** /gene/id/{gene_ids}/download | Get a gene dataset by gene ID
[**DownloadGenePackagePost**](GeneApi.md#DownloadGenePackagePost) | **Post** /gene/download | Get a gene dataset by POST
[**GeneChromosomeSummary**](GeneApi.md#GeneChromosomeSummary) | **Get** /gene/taxon/{taxon}/annotation/{annotation_name}/chromosome_summary | Get summary of chromosomes for a particular taxon&#39;s annotation
[**GeneCountsForTaxon**](GeneApi.md#GeneCountsForTaxon) | **Get** /gene/taxon/{taxon}/counts | Get gene counts by taxonomic identifier
[**GeneCountsForTaxonByPost**](GeneApi.md#GeneCountsForTaxonByPost) | **Post** /gene/taxon/counts | Get gene counts by taxonomic identifier
[**GeneDownloadSummaryById**](GeneApi.md#GeneDownloadSummaryById) | **Get** /gene/id/{gene_ids}/download_summary | Get gene download summary by GeneID
[**GeneDownloadSummaryByPost**](GeneApi.md#GeneDownloadSummaryByPost) | **Post** /gene/download_summary | Get gene download summary
[**GeneLinksById**](GeneApi.md#GeneLinksById) | **Get** /gene/id/{gene_ids}/links | Get gene links by gene ID
[**GeneLinksByIdByPost**](GeneApi.md#GeneLinksByIdByPost) | **Post** /gene/links | Get gene links by gene ID
[**GeneMetadataByAccession**](GeneApi.md#GeneMetadataByAccession) | **Get** /gene/accession/{accessions} | Get gene metadata by RefSeq Accession
[**GeneMetadataByPost**](GeneApi.md#GeneMetadataByPost) | **Post** /gene | Get gene metadata as JSON
[**GeneMetadataByTaxAndSymbol**](GeneApi.md#GeneMetadataByTaxAndSymbol) | **Get** /gene/symbol/{symbols}/taxon/{taxon} | Get gene metadata by gene symbol
[**GeneOrthologsById**](GeneApi.md#GeneOrthologsById) | **Get** /gene/id/{gene_id}/orthologs | Get gene orthologs by gene ID
[**GeneOrthologsByPost**](GeneApi.md#GeneOrthologsByPost) | **Post** /gene/orthologs | Get gene orthologs by gene ID
[**GeneReportsById**](GeneApi.md#GeneReportsById) | **Get** /gene/id/{gene_ids} | Get gene reports by GeneID
[**GeneReportsByTaxon**](GeneApi.md#GeneReportsByTaxon) | **Get** /gene/taxon/{taxon} | Get gene reports by taxonomic identifier



## DownloadGenePackage

> *os.File DownloadGenePackage(ctx, geneIds).IncludeAnnotationType(includeAnnotationType).FastaFilter(fastaFilter).AuxReport(auxReport).TableFields(tableFields).TableReportType(tableReportType).Filename(filename).Execute()

Get a gene dataset by gene ID



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
    geneIds := []int32{int32(123)} // []int32 | NCBI gene ids
    includeAnnotationType := []openapiclient.V2Fasta{openapiclient.v2Fasta("FASTA_UNSPECIFIED")} // []V2Fasta | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    fastaFilter := []string{"Inner_example"} // []string | Limit the FASTA sequences in the datasets package to these transcript and protein accessions (optional)
    auxReport := []openapiclient.V2GeneDatasetRequestGeneDatasetReportType{openapiclient.v2GeneDatasetRequestGeneDatasetReportType("DATASET_REPORT")} // []V2GeneDatasetRequestGeneDatasetReportType | list additional reports to include with download. Data report is included by default. (optional)
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    tableReportType := openapiclient.v2GeneDatasetRequestGeneDatasetReportType("DATASET_REPORT") // V2GeneDatasetRequestGeneDatasetReportType | Specify the report from which the table fields will be taken (optional) (default to "DATASET_REPORT")
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.DownloadGenePackage(context.Background(), geneIds).IncludeAnnotationType(includeAnnotationType).FastaFilter(fastaFilter).AuxReport(auxReport).TableFields(tableFields).TableReportType(tableReportType).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.DownloadGenePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadGenePackage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.DownloadGenePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneIds** | [**[]int32**](int32.md) | NCBI gene ids | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGenePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAnnotationType** | [**[]V2Fasta**](V2Fasta.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
 **fastaFilter** | **[]string** | Limit the FASTA sequences in the datasets package to these transcript and protein accessions | 
 **auxReport** | [**[]V2GeneDatasetRequestGeneDatasetReportType**](V2GeneDatasetRequestGeneDatasetReportType.md) | list additional reports to include with download. Data report is included by default. | 
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **tableReportType** | [**V2GeneDatasetRequestGeneDatasetReportType**](V2GeneDatasetRequestGeneDatasetReportType.md) | Specify the report from which the table fields will be taken | [default to &quot;DATASET_REPORT&quot;]
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


## DownloadGenePackagePost

> *os.File DownloadGenePackagePost(ctx).V2GeneDatasetRequest(v2GeneDatasetRequest).Filename(filename).Execute()

Get a gene dataset by POST



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
    v2GeneDatasetRequest := *openapiclient.NewV2GeneDatasetRequest() // V2GeneDatasetRequest | 
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.DownloadGenePackagePost(context.Background()).V2GeneDatasetRequest(v2GeneDatasetRequest).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.DownloadGenePackagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadGenePackagePost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.DownloadGenePackagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGenePackagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2GeneDatasetRequest** | [**V2GeneDatasetRequest**](V2GeneDatasetRequest.md) |  | 
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


## GeneChromosomeSummary

> V2GeneChromosomeSummaryReply GeneChromosomeSummary(ctx, taxon, annotationName).Execute()

Get summary of chromosomes for a particular taxon's annotation



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
    taxon := "9117" // string | 
    annotationName := "GCF_028858705.1-RS_2023_03" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneChromosomeSummary(context.Background(), taxon, annotationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneChromosomeSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneChromosomeSummary`: V2GeneChromosomeSummaryReply
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneChromosomeSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** |  | 
**annotationName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneChromosomeSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V2GeneChromosomeSummaryReply**](V2GeneChromosomeSummaryReply.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneCountsForTaxon

> V2GeneCountsByTaxonReply GeneCountsForTaxon(ctx, taxon).Execute()

Get gene counts by taxonomic identifier



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
    taxon := "9606" // string | Taxon for provided gene symbol

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneCountsForTaxon(context.Background(), taxon).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneCountsForTaxon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneCountsForTaxon`: V2GeneCountsByTaxonReply
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneCountsForTaxon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | Taxon for provided gene symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneCountsForTaxonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2GeneCountsByTaxonReply**](V2GeneCountsByTaxonReply.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneCountsForTaxonByPost

> V2GeneCountsByTaxonReply GeneCountsForTaxonByPost(ctx).V2GeneCountsByTaxonRequest(v2GeneCountsByTaxonRequest).Execute()

Get gene counts by taxonomic identifier



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
    v2GeneCountsByTaxonRequest := *openapiclient.NewV2GeneCountsByTaxonRequest() // V2GeneCountsByTaxonRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneCountsForTaxonByPost(context.Background()).V2GeneCountsByTaxonRequest(v2GeneCountsByTaxonRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneCountsForTaxonByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneCountsForTaxonByPost`: V2GeneCountsByTaxonReply
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneCountsForTaxonByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneCountsForTaxonByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2GeneCountsByTaxonRequest** | [**V2GeneCountsByTaxonRequest**](V2GeneCountsByTaxonRequest.md) |  | 

### Return type

[**V2GeneCountsByTaxonReply**](V2GeneCountsByTaxonReply.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneDownloadSummaryById

> V2DownloadSummary GeneDownloadSummaryById(ctx, geneIds).IncludeAnnotationType(includeAnnotationType).ReturnedContent(returnedContent).FastaFilter(fastaFilter).AuxReport(auxReport).TableFields(tableFields).TableReportType(tableReportType).Execute()

Get gene download summary by GeneID



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
    geneIds := []int32{int32(123)} // []int32 | NCBI gene ids
    includeAnnotationType := []openapiclient.V2Fasta{openapiclient.v2Fasta("FASTA_UNSPECIFIED")} // []V2Fasta | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    returnedContent := openapiclient.v2GeneDatasetRequestContentType("COMPLETE") // V2GeneDatasetRequestContentType | Return either gene-ids, or entire gene metadata (optional) (default to "COMPLETE")
    fastaFilter := []string{"Inner_example"} // []string | Limit the FASTA sequences in the datasets package to these transcript and protein accessions (optional)
    auxReport := []openapiclient.V2GeneDatasetRequestGeneDatasetReportType{openapiclient.v2GeneDatasetRequestGeneDatasetReportType("DATASET_REPORT")} // []V2GeneDatasetRequestGeneDatasetReportType | list additional reports to include with download. Data report is included by default. (optional)
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    tableReportType := openapiclient.v2GeneDatasetRequestGeneDatasetReportType("DATASET_REPORT") // V2GeneDatasetRequestGeneDatasetReportType | Specify the report from which the table fields will be taken (optional) (default to "DATASET_REPORT")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneDownloadSummaryById(context.Background(), geneIds).IncludeAnnotationType(includeAnnotationType).ReturnedContent(returnedContent).FastaFilter(fastaFilter).AuxReport(auxReport).TableFields(tableFields).TableReportType(tableReportType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneDownloadSummaryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneDownloadSummaryById`: V2DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneDownloadSummaryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneIds** | [**[]int32**](int32.md) | NCBI gene ids | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneDownloadSummaryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAnnotationType** | [**[]V2Fasta**](V2Fasta.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
 **returnedContent** | [**V2GeneDatasetRequestContentType**](V2GeneDatasetRequestContentType.md) | Return either gene-ids, or entire gene metadata | [default to &quot;COMPLETE&quot;]
 **fastaFilter** | **[]string** | Limit the FASTA sequences in the datasets package to these transcript and protein accessions | 
 **auxReport** | [**[]V2GeneDatasetRequestGeneDatasetReportType**](V2GeneDatasetRequestGeneDatasetReportType.md) | list additional reports to include with download. Data report is included by default. | 
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **tableReportType** | [**V2GeneDatasetRequestGeneDatasetReportType**](V2GeneDatasetRequestGeneDatasetReportType.md) | Specify the report from which the table fields will be taken | [default to &quot;DATASET_REPORT&quot;]

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


## GeneDownloadSummaryByPost

> V2DownloadSummary GeneDownloadSummaryByPost(ctx).V2GeneDatasetRequest(v2GeneDatasetRequest).Execute()

Get gene download summary



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
    v2GeneDatasetRequest := *openapiclient.NewV2GeneDatasetRequest() // V2GeneDatasetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneDownloadSummaryByPost(context.Background()).V2GeneDatasetRequest(v2GeneDatasetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneDownloadSummaryByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneDownloadSummaryByPost`: V2DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneDownloadSummaryByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneDownloadSummaryByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2GeneDatasetRequest** | [**V2GeneDatasetRequest**](V2GeneDatasetRequest.md) |  | 

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


## GeneLinksById

> V2GeneLinksReply GeneLinksById(ctx, geneIds).Execute()

Get gene links by gene ID



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
    geneIds := []int32{int32(123)} // []int32 | NCBI gene ids, limited to 1000 ids

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneLinksById(context.Background(), geneIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneLinksById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneLinksById`: V2GeneLinksReply
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneLinksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneIds** | [**[]int32**](int32.md) | NCBI gene ids, limited to 1000 ids | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneLinksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2GeneLinksReply**](V2GeneLinksReply.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneLinksByIdByPost

> V2GeneLinksReply GeneLinksByIdByPost(ctx).V2GeneLinksRequest(v2GeneLinksRequest).Execute()

Get gene links by gene ID



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
    v2GeneLinksRequest := *openapiclient.NewV2GeneLinksRequest() // V2GeneLinksRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneLinksByIdByPost(context.Background()).V2GeneLinksRequest(v2GeneLinksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneLinksByIdByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneLinksByIdByPost`: V2GeneLinksReply
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneLinksByIdByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneLinksByIdByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2GeneLinksRequest** | [**V2GeneLinksRequest**](V2GeneLinksRequest.md) |  | 

### Return type

[**V2GeneLinksReply**](V2GeneLinksReply.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneMetadataByAccession

> V2reportsGeneDataReportPage GeneMetadataByAccession(ctx, accessions).ReturnedContent(returnedContent).TableFields(tableFields).IncludeTabularHeader(includeTabularHeader).PageSize(pageSize).PageToken(pageToken).Execute()

Get gene metadata by RefSeq Accession



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
    accessions := []string{"Inner_example"} // []string | RNA or Protein accessions.
    returnedContent := openapiclient.v2GeneDatasetReportsRequestContentType("COMPLETE") // V2GeneDatasetReportsRequestContentType | Return either gene-ids, or entire gene metadata (optional) (default to "COMPLETE")
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")
    pageSize := int32(56) // int32 | The maximum number of gene reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyDatasetReportsRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyDatasetReportsRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneMetadataByAccession(context.Background(), accessions).ReturnedContent(returnedContent).TableFields(tableFields).IncludeTabularHeader(includeTabularHeader).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneMetadataByAccession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneMetadataByAccession`: V2reportsGeneDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneMetadataByAccession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | RNA or Protein accessions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneMetadataByAccessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | [**V2GeneDatasetReportsRequestContentType**](V2GeneDatasetReportsRequestContentType.md) | Return either gene-ids, or entire gene metadata | [default to &quot;COMPLETE&quot;]
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]
 **pageSize** | **int32** | The maximum number of gene reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyDatasetReportsRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyDatasetReportsRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V2reportsGeneDataReportPage**](V2reportsGeneDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneMetadataByPost

> V2reportsGeneDataReportPage GeneMetadataByPost(ctx).V2GeneDatasetReportsRequest(v2GeneDatasetReportsRequest).Execute()

Get gene metadata as JSON



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
    v2GeneDatasetReportsRequest := *openapiclient.NewV2GeneDatasetReportsRequest() // V2GeneDatasetReportsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneMetadataByPost(context.Background()).V2GeneDatasetReportsRequest(v2GeneDatasetReportsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneMetadataByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneMetadataByPost`: V2reportsGeneDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneMetadataByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneMetadataByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2GeneDatasetReportsRequest** | [**V2GeneDatasetReportsRequest**](V2GeneDatasetReportsRequest.md) |  | 

### Return type

[**V2reportsGeneDataReportPage**](V2reportsGeneDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneMetadataByTaxAndSymbol

> V2reportsGeneDataReportPage GeneMetadataByTaxAndSymbol(ctx, symbols, taxon).ReturnedContent(returnedContent).TableFields(tableFields).IncludeTabularHeader(includeTabularHeader).PageSize(pageSize).PageToken(pageToken).Execute()

Get gene metadata by gene symbol



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
    symbols := []string{"Inner_example"} // []string | Gene symbol
    taxon := "9606" // string | Taxon for provided gene symbol
    returnedContent := openapiclient.v2GeneDatasetReportsRequestContentType("COMPLETE") // V2GeneDatasetReportsRequestContentType | Return either gene-ids, or entire gene metadata (optional) (default to "COMPLETE")
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")
    pageSize := int32(56) // int32 | The maximum number of gene reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyDatasetReportsRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyDatasetReportsRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneMetadataByTaxAndSymbol(context.Background(), symbols, taxon).ReturnedContent(returnedContent).TableFields(tableFields).IncludeTabularHeader(includeTabularHeader).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneMetadataByTaxAndSymbol``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneMetadataByTaxAndSymbol`: V2reportsGeneDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneMetadataByTaxAndSymbol`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbols** | [**[]string**](string.md) | Gene symbol | 
**taxon** | **string** | Taxon for provided gene symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneMetadataByTaxAndSymbolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **returnedContent** | [**V2GeneDatasetReportsRequestContentType**](V2GeneDatasetReportsRequestContentType.md) | Return either gene-ids, or entire gene metadata | [default to &quot;COMPLETE&quot;]
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]
 **pageSize** | **int32** | The maximum number of gene reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyDatasetReportsRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyDatasetReportsRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V2reportsGeneDataReportPage**](V2reportsGeneDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneOrthologsById

> V2reportsGeneDataReportPage GeneOrthologsById(ctx, geneId).ReturnedContent(returnedContent).TaxonFilter(taxonFilter).PageSize(pageSize).PageToken(pageToken).Execute()

Get gene orthologs by gene ID



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
    geneId := int32(2778) // int32 | 
    returnedContent := openapiclient.v2OrthologRequestContentType("COMPLETE") // V2OrthologRequestContentType | Return either gene-ids, or entire gene metadata (optional) (default to "COMPLETE")
    taxonFilter := []string{"Inner_example"} // []string | Filter genes by taxa (optional)
    pageSize := int32(56) // int32 | The maximum number of gene reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `OrthologRequest` call with more than `page_size` results. Use this token, along with the previous `OrthologRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneOrthologsById(context.Background(), geneId).ReturnedContent(returnedContent).TaxonFilter(taxonFilter).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneOrthologsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneOrthologsById`: V2reportsGeneDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneOrthologsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneOrthologsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | [**V2OrthologRequestContentType**](V2OrthologRequestContentType.md) | Return either gene-ids, or entire gene metadata | [default to &quot;COMPLETE&quot;]
 **taxonFilter** | **[]string** | Filter genes by taxa | 
 **pageSize** | **int32** | The maximum number of gene reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;OrthologRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;OrthologRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V2reportsGeneDataReportPage**](V2reportsGeneDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneOrthologsByPost

> V2reportsGeneDataReportPage GeneOrthologsByPost(ctx).V2OrthologRequest(v2OrthologRequest).Execute()

Get gene orthologs by gene ID



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
    v2OrthologRequest := *openapiclient.NewV2OrthologRequest() // V2OrthologRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneOrthologsByPost(context.Background()).V2OrthologRequest(v2OrthologRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneOrthologsByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneOrthologsByPost`: V2reportsGeneDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneOrthologsByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneOrthologsByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2OrthologRequest** | [**V2OrthologRequest**](V2OrthologRequest.md) |  | 

### Return type

[**V2reportsGeneDataReportPage**](V2reportsGeneDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneReportsById

> V2reportsGeneDataReportPage GeneReportsById(ctx, geneIds).ReturnedContent(returnedContent).TableFields(tableFields).IncludeTabularHeader(includeTabularHeader).PageSize(pageSize).PageToken(pageToken).Execute()

Get gene reports by GeneID



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
    geneIds := []int32{int32(123)} // []int32 | NCBI gene ids
    returnedContent := openapiclient.v2GeneDatasetReportsRequestContentType("COMPLETE") // V2GeneDatasetReportsRequestContentType | Return either gene-ids, or entire gene metadata (optional) (default to "COMPLETE")
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")
    pageSize := int32(56) // int32 | The maximum number of gene reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyDatasetReportsRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyDatasetReportsRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneReportsById(context.Background(), geneIds).ReturnedContent(returnedContent).TableFields(tableFields).IncludeTabularHeader(includeTabularHeader).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneReportsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneReportsById`: V2reportsGeneDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneReportsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneIds** | [**[]int32**](int32.md) | NCBI gene ids | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneReportsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | [**V2GeneDatasetReportsRequestContentType**](V2GeneDatasetReportsRequestContentType.md) | Return either gene-ids, or entire gene metadata | [default to &quot;COMPLETE&quot;]
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]
 **pageSize** | **int32** | The maximum number of gene reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyDatasetReportsRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyDatasetReportsRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V2reportsGeneDataReportPage**](V2reportsGeneDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneReportsByTaxon

> V2reportsGeneDataReportPage GeneReportsByTaxon(ctx, taxon).ReturnedContent(returnedContent).TableFields(tableFields).IncludeTabularHeader(includeTabularHeader).PageSize(pageSize).PageToken(pageToken).Query(query).Types(types).Execute()

Get gene reports by taxonomic identifier



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
    taxon := "9606" // string | NCBI Taxonomy ID or name (common or scientific) that the genes are annotated at  end oneof }
    returnedContent := openapiclient.v2GeneDatasetReportsRequestContentType("COMPLETE") // V2GeneDatasetReportsRequestContentType | Return either gene-ids, or entire gene metadata (optional) (default to "COMPLETE")
    tableFields := []string{"Inner_example"} // []string | Specify which fields to include in the tabular report (optional)
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")
    pageSize := int32(56) // int32 | The maximum number of gene reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyDatasetReportsRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyDatasetReportsRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)
    query := "A2M immunoglobulin" // string | text search within gene symbol, aliases, name, and protein name (optional)
    types := []openapiclient.V2GeneType{openapiclient.v2GeneType("UNKNOWN")} // []V2GeneType | Gene types to filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneApi.GeneReportsByTaxon(context.Background(), taxon).ReturnedContent(returnedContent).TableFields(tableFields).IncludeTabularHeader(includeTabularHeader).PageSize(pageSize).PageToken(pageToken).Query(query).Types(types).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneReportsByTaxon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneReportsByTaxon`: V2reportsGeneDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneReportsByTaxon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) that the genes are annotated at  end oneof } | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneReportsByTaxonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | [**V2GeneDatasetReportsRequestContentType**](V2GeneDatasetReportsRequestContentType.md) | Return either gene-ids, or entire gene metadata | [default to &quot;COMPLETE&quot;]
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]
 **pageSize** | **int32** | The maximum number of gene reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyDatasetReportsRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyDatasetReportsRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **query** | **string** | text search within gene symbol, aliases, name, and protein name | 
 **types** | [**[]V2GeneType**](V2GeneType.md) | Gene types to filter | 

### Return type

[**V2reportsGeneDataReportPage**](V2reportsGeneDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

