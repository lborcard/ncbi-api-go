# \ProkaryoteApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v2alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadProkaryoteGenePackage**](ProkaryoteApi.md#DownloadProkaryoteGenePackage) | **Get** /protein/accession/{accessions}/download | Get a prokaryote gene dataset by RefSeq protein accession
[**DownloadProkaryoteGenePackagePost**](ProkaryoteApi.md#DownloadProkaryoteGenePackagePost) | **Post** /protein/accession/download | Get a prokaryote gene dataset by RefSeq protein accession by POST



## DownloadProkaryoteGenePackage

> *os.File DownloadProkaryoteGenePackage(ctx, accessions).IncludeAnnotationType(includeAnnotationType).GeneFlankConfigLength(geneFlankConfigLength).Taxon(taxon).Filename(filename).Execute()

Get a prokaryote gene dataset by RefSeq protein accession



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
    accessions := []string{"Inner_example"} // []string | WP prokaryote protein accession
    includeAnnotationType := []openapiclient.V2Fasta{openapiclient.v2Fasta("FASTA_UNSPECIFIED")} // []V2Fasta | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    geneFlankConfigLength := int32(56) // int32 |  (optional)
    taxon := "taxon_example" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank When specified, return data from this taxon and its subtree (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProkaryoteApi.DownloadProkaryoteGenePackage(context.Background(), accessions).IncludeAnnotationType(includeAnnotationType).GeneFlankConfigLength(geneFlankConfigLength).Taxon(taxon).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProkaryoteApi.DownloadProkaryoteGenePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadProkaryoteGenePackage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ProkaryoteApi.DownloadProkaryoteGenePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | WP prokaryote protein accession | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadProkaryoteGenePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAnnotationType** | [**[]V2Fasta**](V2Fasta.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
 **geneFlankConfigLength** | **int32** |  | 
 **taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank When specified, return data from this taxon and its subtree | 
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


## DownloadProkaryoteGenePackagePost

> *os.File DownloadProkaryoteGenePackagePost(ctx).V2ProkaryoteGeneRequest(v2ProkaryoteGeneRequest).Filename(filename).Execute()

Get a prokaryote gene dataset by RefSeq protein accession by POST



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
    v2ProkaryoteGeneRequest := *openapiclient.NewV2ProkaryoteGeneRequest() // V2ProkaryoteGeneRequest | 
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProkaryoteApi.DownloadProkaryoteGenePackagePost(context.Background()).V2ProkaryoteGeneRequest(v2ProkaryoteGeneRequest).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProkaryoteApi.DownloadProkaryoteGenePackagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadProkaryoteGenePackagePost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ProkaryoteApi.DownloadProkaryoteGenePackagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadProkaryoteGenePackagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2ProkaryoteGeneRequest** | [**V2ProkaryoteGeneRequest**](V2ProkaryoteGeneRequest.md) |  | 
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

