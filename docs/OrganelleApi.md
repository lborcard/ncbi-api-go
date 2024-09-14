# \OrganelleApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v2alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadOrganellePackage**](OrganelleApi.md#DownloadOrganellePackage) | **Get** /organelle/accession/{accessions}/download | Get a organelle data package by accesions
[**DownloadOrganellePackageByPost**](OrganelleApi.md#DownloadOrganellePackageByPost) | **Post** /organelle/download | Get a organelle data package by post
[**OrganelleDatareportByAccession**](OrganelleApi.md#OrganelleDatareportByAccession) | **Get** /organelle/accessions/{accessions}/dataset_report | Get Organelle dataset report by accession
[**OrganelleDatareportByPost**](OrganelleApi.md#OrganelleDatareportByPost) | **Post** /organelle/dataset_report | Get Organelle dataset report by http post
[**OrganelleDatareportByTaxon**](OrganelleApi.md#OrganelleDatareportByTaxon) | **Get** /organelle/taxon/{taxons}/dataset_report | Get Organelle dataset report by taxons



## DownloadOrganellePackage

> *os.File DownloadOrganellePackage(ctx, accessions).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Filename(filename).Execute()

Get a organelle data package by accesions



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
    accessions := []string{"Inner_example"} // []string | NCBI organelle assembly accessions
    excludeSequence := true // bool | Set to true to omit the genomic sequence. (optional)
    includeAnnotationType := []openapiclient.V2AnnotationForOrganelleType{openapiclient.v2AnnotationForOrganelleType("GENOME_FASTA")} // []V2AnnotationForOrganelleType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganelleApi.DownloadOrganellePackage(context.Background(), accessions).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganelleApi.DownloadOrganellePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadOrganellePackage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OrganelleApi.DownloadOrganellePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | NCBI organelle assembly accessions | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadOrganellePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeSequence** | **bool** | Set to true to omit the genomic sequence. | 
 **includeAnnotationType** | [**[]V2AnnotationForOrganelleType**](V2AnnotationForOrganelleType.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
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


## DownloadOrganellePackageByPost

> *os.File DownloadOrganellePackageByPost(ctx).V2OrganelleDownloadRequest(v2OrganelleDownloadRequest).Filename(filename).Execute()

Get a organelle data package by post



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
    v2OrganelleDownloadRequest := *openapiclient.NewV2OrganelleDownloadRequest() // V2OrganelleDownloadRequest | 
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganelleApi.DownloadOrganellePackageByPost(context.Background()).V2OrganelleDownloadRequest(v2OrganelleDownloadRequest).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganelleApi.DownloadOrganellePackageByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadOrganellePackageByPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OrganelleApi.DownloadOrganellePackageByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadOrganellePackageByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2OrganelleDownloadRequest** | [**V2OrganelleDownloadRequest**](V2OrganelleDownloadRequest.md) |  | 
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


## OrganelleDatareportByAccession

> V2reportsOrganelleDataReports OrganelleDatareportByAccession(ctx, accessions).Taxons(taxons).OrganelleTypes(organelleTypes).FirstReleaseDate(firstReleaseDate).LastReleaseDate(lastReleaseDate).TaxExactMatch(taxExactMatch).SortField(sortField).SortDirection(sortDirection).ReturnedContent(returnedContent).TableFormat(tableFormat).IncludeTabularHeader(includeTabularHeader).Execute()

Get Organelle dataset report by accession



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
    accessions := []string{"Inner_example"} // []string | NCBI assembly accession
    taxons := []string{"Inner_example"} // []string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank (optional)
    organelleTypes := []openapiclient.V2reportsOrganelleType{openapiclient.v2reportsOrganelleType("ORGANELLE_TYPE_UNKNOWN")} // []V2reportsOrganelleType |  (optional)
    firstReleaseDate := time.Now() // time.Time | Only return organelle assemblies that were released on or after the specified date By default, do not filter. (optional)
    lastReleaseDate := time.Now() // time.Time | Only return organelle assemblies that were released on or before to the specified date By default, do not filter. (optional)
    taxExactMatch := true // bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. (optional) (default to false)
    sortField := "sortField_example" // string |  (optional)
    sortDirection := openapiclient.v2SortDirection("SORT_DIRECTION_UNSPECIFIED") // V2SortDirection |  (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    returnedContent := openapiclient.v2OrganelleMetadataRequestContentType("COMPLETE") // V2OrganelleMetadataRequestContentType | Return either assembly accessions, or entire assembly-metadata records (optional) (default to "COMPLETE")
    tableFormat := openapiclient.v2OrganelleMetadataRequestOrganelleTableFormat("ORGANELLE_TABLE_FORMAT_NO_TABLE") // V2OrganelleMetadataRequestOrganelleTableFormat | Optional pre-defined template for processing a tabular data request (optional) (default to "ORGANELLE_TABLE_FORMAT_NO_TABLE")
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganelleApi.OrganelleDatareportByAccession(context.Background(), accessions).Taxons(taxons).OrganelleTypes(organelleTypes).FirstReleaseDate(firstReleaseDate).LastReleaseDate(lastReleaseDate).TaxExactMatch(taxExactMatch).SortField(sortField).SortDirection(sortDirection).ReturnedContent(returnedContent).TableFormat(tableFormat).IncludeTabularHeader(includeTabularHeader).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganelleApi.OrganelleDatareportByAccession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganelleDatareportByAccession`: V2reportsOrganelleDataReports
    fmt.Fprintf(os.Stdout, "Response from `OrganelleApi.OrganelleDatareportByAccession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | NCBI assembly accession | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganelleDatareportByAccessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxons** | **[]string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 
 **organelleTypes** | [**[]V2reportsOrganelleType**](V2reportsOrganelleType.md) |  | 
 **firstReleaseDate** | **time.Time** | Only return organelle assemblies that were released on or after the specified date By default, do not filter. | 
 **lastReleaseDate** | **time.Time** | Only return organelle assemblies that were released on or before to the specified date By default, do not filter. | 
 **taxExactMatch** | **bool** | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. | [default to false]
 **sortField** | **string** |  | 
 **sortDirection** | [**V2SortDirection**](V2SortDirection.md) |  | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **returnedContent** | [**V2OrganelleMetadataRequestContentType**](V2OrganelleMetadataRequestContentType.md) | Return either assembly accessions, or entire assembly-metadata records | [default to &quot;COMPLETE&quot;]
 **tableFormat** | [**V2OrganelleMetadataRequestOrganelleTableFormat**](V2OrganelleMetadataRequestOrganelleTableFormat.md) | Optional pre-defined template for processing a tabular data request | [default to &quot;ORGANELLE_TABLE_FORMAT_NO_TABLE&quot;]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]

### Return type

[**V2reportsOrganelleDataReports**](V2reportsOrganelleDataReports.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganelleDatareportByPost

> V2reportsOrganelleDataReports OrganelleDatareportByPost(ctx).V2OrganelleMetadataRequest(v2OrganelleMetadataRequest).Execute()

Get Organelle dataset report by http post



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
    v2OrganelleMetadataRequest := *openapiclient.NewV2OrganelleMetadataRequest() // V2OrganelleMetadataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganelleApi.OrganelleDatareportByPost(context.Background()).V2OrganelleMetadataRequest(v2OrganelleMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganelleApi.OrganelleDatareportByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganelleDatareportByPost`: V2reportsOrganelleDataReports
    fmt.Fprintf(os.Stdout, "Response from `OrganelleApi.OrganelleDatareportByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganelleDatareportByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2OrganelleMetadataRequest** | [**V2OrganelleMetadataRequest**](V2OrganelleMetadataRequest.md) |  | 

### Return type

[**V2reportsOrganelleDataReports**](V2reportsOrganelleDataReports.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganelleDatareportByTaxon

> V2reportsOrganelleDataReports OrganelleDatareportByTaxon(ctx, taxons).OrganelleTypes(organelleTypes).FirstReleaseDate(firstReleaseDate).LastReleaseDate(lastReleaseDate).TaxExactMatch(taxExactMatch).SortField(sortField).SortDirection(sortDirection).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).TableFormat(tableFormat).IncludeTabularHeader(includeTabularHeader).Execute()

Get Organelle dataset report by taxons



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
    organelleTypes := []openapiclient.V2reportsOrganelleType{openapiclient.v2reportsOrganelleType("ORGANELLE_TYPE_UNKNOWN")} // []V2reportsOrganelleType |  (optional)
    firstReleaseDate := time.Now() // time.Time | Only return organelle assemblies that were released on or after the specified date By default, do not filter. (optional)
    lastReleaseDate := time.Now() // time.Time | Only return organelle assemblies that were released on or before to the specified date By default, do not filter. (optional)
    taxExactMatch := true // bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. (optional) (default to false)
    sortField := "sortField_example" // string |  (optional)
    sortDirection := openapiclient.v2SortDirection("SORT_DIRECTION_UNSPECIFIED") // V2SortDirection |  (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    returnedContent := openapiclient.v2OrganelleMetadataRequestContentType("COMPLETE") // V2OrganelleMetadataRequestContentType | Return either assembly accessions, or entire assembly-metadata records (optional) (default to "COMPLETE")
    pageSize := int32(56) // int32 | The maximum number of organelle assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional)
    pageToken := "pageToken_example" // string | A page token is returned from an `OrganelleMetadata` call with more than `page_size` results. Use this token, along with the previous `OrganelleMetadata` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)
    tableFormat := openapiclient.v2OrganelleMetadataRequestOrganelleTableFormat("ORGANELLE_TABLE_FORMAT_NO_TABLE") // V2OrganelleMetadataRequestOrganelleTableFormat | Optional pre-defined template for processing a tabular data request (optional) (default to "ORGANELLE_TABLE_FORMAT_NO_TABLE")
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganelleApi.OrganelleDatareportByTaxon(context.Background(), taxons).OrganelleTypes(organelleTypes).FirstReleaseDate(firstReleaseDate).LastReleaseDate(lastReleaseDate).TaxExactMatch(taxExactMatch).SortField(sortField).SortDirection(sortDirection).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).TableFormat(tableFormat).IncludeTabularHeader(includeTabularHeader).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganelleApi.OrganelleDatareportByTaxon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganelleDatareportByTaxon`: V2reportsOrganelleDataReports
    fmt.Fprintf(os.Stdout, "Response from `OrganelleApi.OrganelleDatareportByTaxon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxons** | [**[]string**](string.md) | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganelleDatareportByTaxonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organelleTypes** | [**[]V2reportsOrganelleType**](V2reportsOrganelleType.md) |  | 
 **firstReleaseDate** | **time.Time** | Only return organelle assemblies that were released on or after the specified date By default, do not filter. | 
 **lastReleaseDate** | **time.Time** | Only return organelle assemblies that were released on or before to the specified date By default, do not filter. | 
 **taxExactMatch** | **bool** | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. | [default to false]
 **sortField** | **string** |  | 
 **sortDirection** | [**V2SortDirection**](V2SortDirection.md) |  | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **returnedContent** | [**V2OrganelleMetadataRequestContentType**](V2OrganelleMetadataRequestContentType.md) | Return either assembly accessions, or entire assembly-metadata records | [default to &quot;COMPLETE&quot;]
 **pageSize** | **int32** | The maximum number of organelle assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | 
 **pageToken** | **string** | A page token is returned from an &#x60;OrganelleMetadata&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;OrganelleMetadata&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **tableFormat** | [**V2OrganelleMetadataRequestOrganelleTableFormat**](V2OrganelleMetadataRequestOrganelleTableFormat.md) | Optional pre-defined template for processing a tabular data request | [default to &quot;ORGANELLE_TABLE_FORMAT_NO_TABLE&quot;]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]

### Return type

[**V2reportsOrganelleDataReports**](V2reportsOrganelleDataReports.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

