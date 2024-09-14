# \TaxonomyApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v2alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadTaxonomyPackage**](TaxonomyApi.md#DownloadTaxonomyPackage) | **Get** /taxonomy/taxon/{tax_ids}/download | Get a taxonomy data package by tax ID
[**DownloadTaxonomyPackageByPost**](TaxonomyApi.md#DownloadTaxonomyPackageByPost) | **Post** /taxonomy/download | Get a taxonomy data package by tax_id
[**TaxNameQuery**](TaxonomyApi.md#TaxNameQuery) | **Get** /taxonomy/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs given a partial taxonomic name
[**TaxNameQueryByPost**](TaxonomyApi.md#TaxNameQueryByPost) | **Post** /taxonomy/taxon_suggest | Get a list of taxonomy names and IDs given a partial taxonomic name
[**TaxonomyDataReport**](TaxonomyApi.md#TaxonomyDataReport) | **Get** /taxonomy/taxon/{taxons}/dataset_report | Use taxonomic identifiers to get taxonomic data report
[**TaxonomyDataReportPost**](TaxonomyApi.md#TaxonomyDataReportPost) | **Post** /taxonomy/dataset_report | Use taxonomic identifiers to get taxonomic names data report by post
[**TaxonomyFilteredSubtree**](TaxonomyApi.md#TaxonomyFilteredSubtree) | **Get** /taxonomy/taxon/{taxons}/filtered_subtree | Use taxonomic identifiers to get a filtered taxonomic subtree
[**TaxonomyFilteredSubtreePost**](TaxonomyApi.md#TaxonomyFilteredSubtreePost) | **Post** /taxonomy/filtered_subtree | Use taxonomic identifiers to get a filtered taxonomic subtree by post
[**TaxonomyImage**](TaxonomyApi.md#TaxonomyImage) | **Get** /taxonomy/taxon/{taxon}/image | Retrieve image associated with a taxonomic identifier
[**TaxonomyImageMetadata**](TaxonomyApi.md#TaxonomyImageMetadata) | **Get** /taxonomy/taxon/{taxon}/image/metadata | Retrieve image metadata associated with a taxonomic identifier
[**TaxonomyImageMetadataPost**](TaxonomyApi.md#TaxonomyImageMetadataPost) | **Post** /taxonomy/image/metadata | Retrieve image metadata associated with a taxonomic identifier by post
[**TaxonomyImagePost**](TaxonomyApi.md#TaxonomyImagePost) | **Post** /taxonomy/image | Retrieve image associated with a taxonomic identifier by post
[**TaxonomyLinks**](TaxonomyApi.md#TaxonomyLinks) | **Get** /taxonomy/taxon/{taxon}/links | Retrieve external links associated with a taxonomic identifier.
[**TaxonomyLinksByPost**](TaxonomyApi.md#TaxonomyLinksByPost) | **Post** /taxonomy/links | Retrieve external links associated with a taxonomic identifier.
[**TaxonomyMetadata**](TaxonomyApi.md#TaxonomyMetadata) | **Get** /taxonomy/taxon/{taxons} | Use taxonomic identifiers to get taxonomic metadata
[**TaxonomyMetadataPost**](TaxonomyApi.md#TaxonomyMetadataPost) | **Post** /taxonomy | Use taxonomic identifiers to get taxonomic metadata by post
[**TaxonomyNames**](TaxonomyApi.md#TaxonomyNames) | **Get** /taxonomy/taxon/{taxons}/name_report | Use taxonomic identifiers to get taxonomic names data report
[**TaxonomyNamesPost**](TaxonomyApi.md#TaxonomyNamesPost) | **Post** /taxonomy/name_report | Use taxonomic identifiers to get taxonomic names data report by post
[**TaxonomyRelatedIds**](TaxonomyApi.md#TaxonomyRelatedIds) | **Get** /taxonomy/taxon/{tax_id}/related_ids | Use taxonomic identifier to get related taxonomic identifiers, such as children
[**TaxonomyRelatedIdsPost**](TaxonomyApi.md#TaxonomyRelatedIdsPost) | **Post** /taxonomy/related_ids | Use taxonomic identifier to get related taxonomic identifiers, such as children



## DownloadTaxonomyPackage

> *os.File DownloadTaxonomyPackage(ctx, taxIds).AuxReports(auxReports).Filename(filename).Execute()

Get a taxonomy data package by tax ID



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
    taxIds := []int32{int32(123)} // []int32 | 
    auxReports := []openapiclient.V2TaxonomyDatasetRequestTaxonomyReportType{openapiclient.v2TaxonomyDatasetRequestTaxonomyReportType("TAXONOMY_SUMMARY")} // []V2TaxonomyDatasetRequestTaxonomyReportType | list additional reports to include with download. TAXONOMY_REPORT is included by default. (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.DownloadTaxonomyPackage(context.Background(), taxIds).AuxReports(auxReports).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.DownloadTaxonomyPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadTaxonomyPackage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.DownloadTaxonomyPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxIds** | [**[]int32**](int32.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadTaxonomyPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auxReports** | [**[]V2TaxonomyDatasetRequestTaxonomyReportType**](V2TaxonomyDatasetRequestTaxonomyReportType.md) | list additional reports to include with download. TAXONOMY_REPORT is included by default. | 
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


## DownloadTaxonomyPackageByPost

> *os.File DownloadTaxonomyPackageByPost(ctx).V2TaxonomyDatasetRequest(v2TaxonomyDatasetRequest).Filename(filename).Execute()

Get a taxonomy data package by tax_id



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
    v2TaxonomyDatasetRequest := *openapiclient.NewV2TaxonomyDatasetRequest() // V2TaxonomyDatasetRequest | 
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.DownloadTaxonomyPackageByPost(context.Background()).V2TaxonomyDatasetRequest(v2TaxonomyDatasetRequest).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.DownloadTaxonomyPackageByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadTaxonomyPackageByPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.DownloadTaxonomyPackageByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadTaxonomyPackageByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2TaxonomyDatasetRequest** | [**V2TaxonomyDatasetRequest**](V2TaxonomyDatasetRequest.md) |  | 
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


## TaxNameQuery

> V2SciNameAndIds TaxNameQuery(ctx, taxonQuery).TaxRankFilter(taxRankFilter).TaxonResourceFilter(taxonResourceFilter).ExactMatch(exactMatch).Execute()

Get a list of taxonomy names and IDs given a partial taxonomic name



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
    taxonQuery := "hum" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    taxRankFilter := openapiclient.v2OrganismQueryRequestTaxRankFilter("species") // V2OrganismQueryRequestTaxRankFilter | Set the scope of searched tax ranks when filtering by gene or genome.  Not used for 'all' (optional) (default to "species")
    taxonResourceFilter := openapiclient.v2OrganismQueryRequestTaxonResourceFilter("TAXON_RESOURCE_FILTER_ALL") // V2OrganismQueryRequestTaxonResourceFilter | Limit results to those with gene or genome counts (no filter by default) (optional) (default to "TAXON_RESOURCE_FILTER_ALL")
    exactMatch := true // bool | If true, only return results that exactly match the provided name or tax-id (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxNameQuery(context.Background(), taxonQuery).TaxRankFilter(taxRankFilter).TaxonResourceFilter(taxonResourceFilter).ExactMatch(exactMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxNameQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxNameQuery`: V2SciNameAndIds
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxNameQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxonQuery** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxNameQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxRankFilter** | [**V2OrganismQueryRequestTaxRankFilter**](V2OrganismQueryRequestTaxRankFilter.md) | Set the scope of searched tax ranks when filtering by gene or genome.  Not used for &#39;all&#39; | [default to &quot;species&quot;]
 **taxonResourceFilter** | [**V2OrganismQueryRequestTaxonResourceFilter**](V2OrganismQueryRequestTaxonResourceFilter.md) | Limit results to those with gene or genome counts (no filter by default) | [default to &quot;TAXON_RESOURCE_FILTER_ALL&quot;]
 **exactMatch** | **bool** | If true, only return results that exactly match the provided name or tax-id | [default to false]

### Return type

[**V2SciNameAndIds**](V2SciNameAndIds.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxNameQueryByPost

> V2SciNameAndIds TaxNameQueryByPost(ctx).V2OrganismQueryRequest(v2OrganismQueryRequest).Execute()

Get a list of taxonomy names and IDs given a partial taxonomic name



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
    v2OrganismQueryRequest := *openapiclient.NewV2OrganismQueryRequest() // V2OrganismQueryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxNameQueryByPost(context.Background()).V2OrganismQueryRequest(v2OrganismQueryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxNameQueryByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxNameQueryByPost`: V2SciNameAndIds
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxNameQueryByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxNameQueryByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2OrganismQueryRequest** | [**V2OrganismQueryRequest**](V2OrganismQueryRequest.md) |  | 

### Return type

[**V2SciNameAndIds**](V2SciNameAndIds.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyDataReport

> V2reportsTaxonomyDataReportPage TaxonomyDataReport(ctx, taxons).ReturnedContent(returnedContent).PageSize(pageSize).IncludeTabularHeader(includeTabularHeader).PageToken(pageToken).TableFormat(tableFormat).Children(children).Ranks(ranks).Execute()

Use taxonomic identifiers to get taxonomic data report



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
    taxons := []string{"Inner_example"} // []string | 
    returnedContent := openapiclient.v2TaxonomyMetadataRequestContentType("COMPLETE") // V2TaxonomyMetadataRequestContentType | Return either tax-ids alone, or entire taxononmy-metadata records (optional) (default to "COMPLETE")
    pageSize := int32(56) // int32 | The maximum number of taxons to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")
    pageToken := "pageToken_example" // string | A page token is returned from `GetTaxonomyDataReportFor` and `GetTaxonomyNamesDataReportFor` calls with more than `page_size` results. When `page_token` is empty, all results have been retrieved. (optional)
    tableFormat := openapiclient.v2TaxonomyMetadataRequestTableFormat("SUMMARY") // V2TaxonomyMetadataRequestTableFormat |  (optional) (default to "SUMMARY")
    children := true // bool | Flag for tax explosion. (optional)
    ranks := []openapiclient.V2reportsRankType{openapiclient.v2reportsRankType("NO_RANK")} // []V2reportsRankType | Only include taxons of the provided ranks. If empty, return all ranks. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyDataReport(context.Background(), taxons).ReturnedContent(returnedContent).PageSize(pageSize).IncludeTabularHeader(includeTabularHeader).PageToken(pageToken).TableFormat(tableFormat).Children(children).Ranks(ranks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyDataReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyDataReport`: V2reportsTaxonomyDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyDataReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxons** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyDataReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | [**V2TaxonomyMetadataRequestContentType**](V2TaxonomyMetadataRequestContentType.md) | Return either tax-ids alone, or entire taxononmy-metadata records | [default to &quot;COMPLETE&quot;]
 **pageSize** | **int32** | The maximum number of taxons to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]
 **pageToken** | **string** | A page token is returned from &#x60;GetTaxonomyDataReportFor&#x60; and &#x60;GetTaxonomyNamesDataReportFor&#x60; calls with more than &#x60;page_size&#x60; results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **tableFormat** | [**V2TaxonomyMetadataRequestTableFormat**](V2TaxonomyMetadataRequestTableFormat.md) |  | [default to &quot;SUMMARY&quot;]
 **children** | **bool** | Flag for tax explosion. | 
 **ranks** | [**[]V2reportsRankType**](V2reportsRankType.md) | Only include taxons of the provided ranks. If empty, return all ranks. | 

### Return type

[**V2reportsTaxonomyDataReportPage**](V2reportsTaxonomyDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyDataReportPost

> V2reportsTaxonomyDataReportPage TaxonomyDataReportPost(ctx).V2TaxonomyMetadataRequest(v2TaxonomyMetadataRequest).Execute()

Use taxonomic identifiers to get taxonomic names data report by post



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
    v2TaxonomyMetadataRequest := *openapiclient.NewV2TaxonomyMetadataRequest() // V2TaxonomyMetadataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyDataReportPost(context.Background()).V2TaxonomyMetadataRequest(v2TaxonomyMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyDataReportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyDataReportPost`: V2reportsTaxonomyDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyDataReportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyDataReportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2TaxonomyMetadataRequest** | [**V2TaxonomyMetadataRequest**](V2TaxonomyMetadataRequest.md) |  | 

### Return type

[**V2reportsTaxonomyDataReportPage**](V2reportsTaxonomyDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyFilteredSubtree

> V2TaxonomyFilteredSubtreeResponse TaxonomyFilteredSubtree(ctx, taxons).SpecifiedLimit(specifiedLimit).RankLimits(rankLimits).Execute()

Use taxonomic identifiers to get a filtered taxonomic subtree



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
    taxons := []string{"Inner_example"} // []string | 
    specifiedLimit := true // bool | Limit to specified species (optional)
    rankLimits := []openapiclient.V2reportsRankType{openapiclient.v2reportsRankType("NO_RANK")} // []V2reportsRankType | Limit to the provided ranks.  If empty, accept any rank. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyFilteredSubtree(context.Background(), taxons).SpecifiedLimit(specifiedLimit).RankLimits(rankLimits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyFilteredSubtree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyFilteredSubtree`: V2TaxonomyFilteredSubtreeResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyFilteredSubtree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxons** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyFilteredSubtreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **specifiedLimit** | **bool** | Limit to specified species | 
 **rankLimits** | [**[]V2reportsRankType**](V2reportsRankType.md) | Limit to the provided ranks.  If empty, accept any rank. | 

### Return type

[**V2TaxonomyFilteredSubtreeResponse**](V2TaxonomyFilteredSubtreeResponse.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyFilteredSubtreePost

> V2TaxonomyFilteredSubtreeResponse TaxonomyFilteredSubtreePost(ctx).V2TaxonomyFilteredSubtreeRequest(v2TaxonomyFilteredSubtreeRequest).Execute()

Use taxonomic identifiers to get a filtered taxonomic subtree by post



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
    v2TaxonomyFilteredSubtreeRequest := *openapiclient.NewV2TaxonomyFilteredSubtreeRequest() // V2TaxonomyFilteredSubtreeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyFilteredSubtreePost(context.Background()).V2TaxonomyFilteredSubtreeRequest(v2TaxonomyFilteredSubtreeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyFilteredSubtreePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyFilteredSubtreePost`: V2TaxonomyFilteredSubtreeResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyFilteredSubtreePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyFilteredSubtreePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2TaxonomyFilteredSubtreeRequest** | [**V2TaxonomyFilteredSubtreeRequest**](V2TaxonomyFilteredSubtreeRequest.md) |  | 

### Return type

[**V2TaxonomyFilteredSubtreeResponse**](V2TaxonomyFilteredSubtreeResponse.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyImage

> *os.File TaxonomyImage(ctx, taxon).ImageSize(imageSize).Execute()

Retrieve image associated with a taxonomic identifier



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
    taxon := "9606" // string | 
    imageSize := openapiclient.v2ImageSize("UNSPECIFIED") // V2ImageSize |  (optional) (default to "UNSPECIFIED")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyImage(context.Background(), taxon).ImageSize(imageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyImage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imageSize** | [**V2ImageSize**](V2ImageSize.md) |  | [default to &quot;UNSPECIFIED&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, image/jpeg, image/png, image/gif, image/tiff, image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyImageMetadata

> V2TaxonomyImageMetadataResponse TaxonomyImageMetadata(ctx, taxon).Execute()

Retrieve image metadata associated with a taxonomic identifier



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
    taxon := "9606" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyImageMetadata(context.Background(), taxon).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyImageMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyImageMetadata`: V2TaxonomyImageMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyImageMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyImageMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2TaxonomyImageMetadataResponse**](V2TaxonomyImageMetadataResponse.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyImageMetadataPost

> V2TaxonomyImageMetadataResponse TaxonomyImageMetadataPost(ctx).V2TaxonomyImageMetadataRequest(v2TaxonomyImageMetadataRequest).Execute()

Retrieve image metadata associated with a taxonomic identifier by post



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
    v2TaxonomyImageMetadataRequest := *openapiclient.NewV2TaxonomyImageMetadataRequest() // V2TaxonomyImageMetadataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyImageMetadataPost(context.Background()).V2TaxonomyImageMetadataRequest(v2TaxonomyImageMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyImageMetadataPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyImageMetadataPost`: V2TaxonomyImageMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyImageMetadataPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyImageMetadataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2TaxonomyImageMetadataRequest** | [**V2TaxonomyImageMetadataRequest**](V2TaxonomyImageMetadataRequest.md) |  | 

### Return type

[**V2TaxonomyImageMetadataResponse**](V2TaxonomyImageMetadataResponse.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyImagePost

> *os.File TaxonomyImagePost(ctx).V2TaxonomyImageRequest(v2TaxonomyImageRequest).Execute()

Retrieve image associated with a taxonomic identifier by post



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
    v2TaxonomyImageRequest := *openapiclient.NewV2TaxonomyImageRequest() // V2TaxonomyImageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyImagePost(context.Background()).V2TaxonomyImageRequest(v2TaxonomyImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyImagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyImagePost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyImagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyImagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2TaxonomyImageRequest** | [**V2TaxonomyImageRequest**](V2TaxonomyImageRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, image/jpeg, image/png, image/tiff, image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyLinks

> V2TaxonomyLinksResponse TaxonomyLinks(ctx, taxon).Execute()

Retrieve external links associated with a taxonomic identifier.



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
    taxon := "9606" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyLinks(context.Background(), taxon).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyLinks`: V2TaxonomyLinksResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2TaxonomyLinksResponse**](V2TaxonomyLinksResponse.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyLinksByPost

> V2TaxonomyLinksResponse TaxonomyLinksByPost(ctx).V2TaxonomyLinksRequest(v2TaxonomyLinksRequest).Execute()

Retrieve external links associated with a taxonomic identifier.



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
    v2TaxonomyLinksRequest := *openapiclient.NewV2TaxonomyLinksRequest() // V2TaxonomyLinksRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyLinksByPost(context.Background()).V2TaxonomyLinksRequest(v2TaxonomyLinksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyLinksByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyLinksByPost`: V2TaxonomyLinksResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyLinksByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyLinksByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2TaxonomyLinksRequest** | [**V2TaxonomyLinksRequest**](V2TaxonomyLinksRequest.md) |  | 

### Return type

[**V2TaxonomyLinksResponse**](V2TaxonomyLinksResponse.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyMetadata

> V2TaxonomyMetadataResponse TaxonomyMetadata(ctx, taxons).ReturnedContent(returnedContent).PageSize(pageSize).IncludeTabularHeader(includeTabularHeader).PageToken(pageToken).TableFormat(tableFormat).Children(children).Ranks(ranks).Execute()

Use taxonomic identifiers to get taxonomic metadata



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
    taxons := []string{"Inner_example"} // []string | 
    returnedContent := openapiclient.v2TaxonomyMetadataRequestContentType("COMPLETE") // V2TaxonomyMetadataRequestContentType | Return either tax-ids alone, or entire taxononmy-metadata records (optional) (default to "COMPLETE")
    pageSize := int32(56) // int32 | The maximum number of taxons to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")
    pageToken := "pageToken_example" // string | A page token is returned from `GetTaxonomyDataReportFor` and `GetTaxonomyNamesDataReportFor` calls with more than `page_size` results. When `page_token` is empty, all results have been retrieved. (optional)
    tableFormat := openapiclient.v2TaxonomyMetadataRequestTableFormat("SUMMARY") // V2TaxonomyMetadataRequestTableFormat |  (optional) (default to "SUMMARY")
    children := true // bool | Flag for tax explosion. (optional)
    ranks := []openapiclient.V2reportsRankType{openapiclient.v2reportsRankType("NO_RANK")} // []V2reportsRankType | Only include taxons of the provided ranks. If empty, return all ranks. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyMetadata(context.Background(), taxons).ReturnedContent(returnedContent).PageSize(pageSize).IncludeTabularHeader(includeTabularHeader).PageToken(pageToken).TableFormat(tableFormat).Children(children).Ranks(ranks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyMetadata`: V2TaxonomyMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxons** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | [**V2TaxonomyMetadataRequestContentType**](V2TaxonomyMetadataRequestContentType.md) | Return either tax-ids alone, or entire taxononmy-metadata records | [default to &quot;COMPLETE&quot;]
 **pageSize** | **int32** | The maximum number of taxons to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]
 **pageToken** | **string** | A page token is returned from &#x60;GetTaxonomyDataReportFor&#x60; and &#x60;GetTaxonomyNamesDataReportFor&#x60; calls with more than &#x60;page_size&#x60; results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **tableFormat** | [**V2TaxonomyMetadataRequestTableFormat**](V2TaxonomyMetadataRequestTableFormat.md) |  | [default to &quot;SUMMARY&quot;]
 **children** | **bool** | Flag for tax explosion. | 
 **ranks** | [**[]V2reportsRankType**](V2reportsRankType.md) | Only include taxons of the provided ranks. If empty, return all ranks. | 

### Return type

[**V2TaxonomyMetadataResponse**](V2TaxonomyMetadataResponse.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyMetadataPost

> V2TaxonomyMetadataResponse TaxonomyMetadataPost(ctx).V2TaxonomyMetadataRequest(v2TaxonomyMetadataRequest).Execute()

Use taxonomic identifiers to get taxonomic metadata by post



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
    v2TaxonomyMetadataRequest := *openapiclient.NewV2TaxonomyMetadataRequest() // V2TaxonomyMetadataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyMetadataPost(context.Background()).V2TaxonomyMetadataRequest(v2TaxonomyMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyMetadataPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyMetadataPost`: V2TaxonomyMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyMetadataPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyMetadataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2TaxonomyMetadataRequest** | [**V2TaxonomyMetadataRequest**](V2TaxonomyMetadataRequest.md) |  | 

### Return type

[**V2TaxonomyMetadataResponse**](V2TaxonomyMetadataResponse.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyNames

> V2reportsTaxonomyNamesDataReportPage TaxonomyNames(ctx, taxons).ReturnedContent(returnedContent).PageSize(pageSize).IncludeTabularHeader(includeTabularHeader).PageToken(pageToken).TableFormat(tableFormat).Children(children).Ranks(ranks).Execute()

Use taxonomic identifiers to get taxonomic names data report



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
    taxons := []string{"Inner_example"} // []string | 
    returnedContent := openapiclient.v2TaxonomyMetadataRequestContentType("COMPLETE") // V2TaxonomyMetadataRequestContentType | Return either tax-ids alone, or entire taxononmy-metadata records (optional) (default to "COMPLETE")
    pageSize := int32(56) // int32 | The maximum number of taxons to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    includeTabularHeader := openapiclient.v2IncludeTabularHeader("INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY") // V2IncludeTabularHeader | Whether this request for tabular data should include the header row (optional) (default to "INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY")
    pageToken := "pageToken_example" // string | A page token is returned from `GetTaxonomyDataReportFor` and `GetTaxonomyNamesDataReportFor` calls with more than `page_size` results. When `page_token` is empty, all results have been retrieved. (optional)
    tableFormat := openapiclient.v2TaxonomyMetadataRequestTableFormat("SUMMARY") // V2TaxonomyMetadataRequestTableFormat |  (optional) (default to "SUMMARY")
    children := true // bool | Flag for tax explosion. (optional)
    ranks := []openapiclient.V2reportsRankType{openapiclient.v2reportsRankType("NO_RANK")} // []V2reportsRankType | Only include taxons of the provided ranks. If empty, return all ranks. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyNames(context.Background(), taxons).ReturnedContent(returnedContent).PageSize(pageSize).IncludeTabularHeader(includeTabularHeader).PageToken(pageToken).TableFormat(tableFormat).Children(children).Ranks(ranks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyNames`: V2reportsTaxonomyNamesDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxons** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | [**V2TaxonomyMetadataRequestContentType**](V2TaxonomyMetadataRequestContentType.md) | Return either tax-ids alone, or entire taxononmy-metadata records | [default to &quot;COMPLETE&quot;]
 **pageSize** | **int32** | The maximum number of taxons to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **includeTabularHeader** | [**V2IncludeTabularHeader**](V2IncludeTabularHeader.md) | Whether this request for tabular data should include the header row | [default to &quot;INCLUDE_TABULAR_HEADER_FIRST_PAGE_ONLY&quot;]
 **pageToken** | **string** | A page token is returned from &#x60;GetTaxonomyDataReportFor&#x60; and &#x60;GetTaxonomyNamesDataReportFor&#x60; calls with more than &#x60;page_size&#x60; results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 
 **tableFormat** | [**V2TaxonomyMetadataRequestTableFormat**](V2TaxonomyMetadataRequestTableFormat.md) |  | [default to &quot;SUMMARY&quot;]
 **children** | **bool** | Flag for tax explosion. | 
 **ranks** | [**[]V2reportsRankType**](V2reportsRankType.md) | Only include taxons of the provided ranks. If empty, return all ranks. | 

### Return type

[**V2reportsTaxonomyNamesDataReportPage**](V2reportsTaxonomyNamesDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyNamesPost

> V2reportsTaxonomyNamesDataReportPage TaxonomyNamesPost(ctx).V2TaxonomyMetadataRequest(v2TaxonomyMetadataRequest).Execute()

Use taxonomic identifiers to get taxonomic names data report by post



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
    v2TaxonomyMetadataRequest := *openapiclient.NewV2TaxonomyMetadataRequest() // V2TaxonomyMetadataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyNamesPost(context.Background()).V2TaxonomyMetadataRequest(v2TaxonomyMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyNamesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyNamesPost`: V2reportsTaxonomyNamesDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyNamesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyNamesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2TaxonomyMetadataRequest** | [**V2TaxonomyMetadataRequest**](V2TaxonomyMetadataRequest.md) |  | 

### Return type

[**V2reportsTaxonomyNamesDataReportPage**](V2reportsTaxonomyNamesDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyRelatedIds

> V2TaxonomyTaxIdsPage TaxonomyRelatedIds(ctx, taxId).IncludeLineage(includeLineage).IncludeSubtree(includeSubtree).Ranks(ranks).PageSize(pageSize).PageToken(pageToken).Execute()

Use taxonomic identifier to get related taxonomic identifiers, such as children



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
    taxId := int32(9606) // int32 | 
    includeLineage := true // bool | If true, return reports for all taxonomy nodes in the lineages of the requested tax_id (optional) (default to false)
    includeSubtree := true // bool | If true, return reports for all taxonomy nodes that are children of the requested tax_ids (optional) (default to false)
    ranks := []openapiclient.V2reportsRankType{openapiclient.v2reportsRankType("NO_RANK")} // []V2reportsRankType | Only include taxons of the provided ranks. If empty, return all ranks. (optional)
    pageSize := int32(56) // int32 | The maximum number of taxids to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from a `GetRelatedTaxids` call with more than `page_size` results. Use this token, along with the previous `TaxonomyRelatedIdRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyRelatedIds(context.Background(), taxId).IncludeLineage(includeLineage).IncludeSubtree(includeSubtree).Ranks(ranks).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyRelatedIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyRelatedIds`: V2TaxonomyTaxIdsPage
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyRelatedIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyRelatedIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeLineage** | **bool** | If true, return reports for all taxonomy nodes in the lineages of the requested tax_id | [default to false]
 **includeSubtree** | **bool** | If true, return reports for all taxonomy nodes that are children of the requested tax_ids | [default to false]
 **ranks** | [**[]V2reportsRankType**](V2reportsRankType.md) | Only include taxons of the provided ranks. If empty, return all ranks. | 
 **pageSize** | **int32** | The maximum number of taxids to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from a &#x60;GetRelatedTaxids&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;TaxonomyRelatedIdRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V2TaxonomyTaxIdsPage**](V2TaxonomyTaxIdsPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyRelatedIdsPost

> V2TaxonomyTaxIdsPage TaxonomyRelatedIdsPost(ctx).V2TaxonomyRelatedIdRequest(v2TaxonomyRelatedIdRequest).Execute()

Use taxonomic identifier to get related taxonomic identifiers, such as children



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
    v2TaxonomyRelatedIdRequest := *openapiclient.NewV2TaxonomyRelatedIdRequest() // V2TaxonomyRelatedIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.TaxonomyRelatedIdsPost(context.Background()).V2TaxonomyRelatedIdRequest(v2TaxonomyRelatedIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyRelatedIdsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyRelatedIdsPost`: V2TaxonomyTaxIdsPage
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyRelatedIdsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyRelatedIdsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2TaxonomyRelatedIdRequest** | [**V2TaxonomyRelatedIdRequest**](V2TaxonomyRelatedIdRequest.md) |  | 

### Return type

[**V2TaxonomyTaxIdsPage**](V2TaxonomyTaxIdsPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

