# \BioSampleApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v2alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BioSampleDatasetReport**](BioSampleApi.md#BioSampleDatasetReport) | **Get** /biosample/accession/{accessions}/biosample_report | Get BioSample dataset reports by accession(s)



## BioSampleDatasetReport

> V2reportsBioSampleDataReportPage BioSampleDatasetReport(ctx, accessions).Execute()

Get BioSample dataset reports by accession(s)



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
    accessions := []string{"Inner_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BioSampleApi.BioSampleDatasetReport(context.Background(), accessions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BioSampleApi.BioSampleDatasetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BioSampleDatasetReport`: V2reportsBioSampleDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `BioSampleApi.BioSampleDatasetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBioSampleDatasetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2reportsBioSampleDataReportPage**](V2reportsBioSampleDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/x-ndjson, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

