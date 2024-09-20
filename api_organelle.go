/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 2 API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated zip archive](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v2alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"os"
	"reflect"
	"time"
)


// OrganelleApiService OrganelleApi service
type OrganelleApiService service

type OrganelleApiDownloadOrganellePackageRequest struct {
	ctx context.Context
	ApiService *OrganelleApiService
	accessions []string
	excludeSequence *bool
	includeAnnotationType *[]V2AnnotationForOrganelleType
	filename *string
}

// Set to true to omit the genomic sequence.
func (r OrganelleApiDownloadOrganellePackageRequest) ExcludeSequence(excludeSequence bool) OrganelleApiDownloadOrganellePackageRequest {
	r.excludeSequence = &excludeSequence
	return r
}

// Select additional types of annotation to include in the data package.  If unset, no annotation is provided.
func (r OrganelleApiDownloadOrganellePackageRequest) IncludeAnnotationType(includeAnnotationType []V2AnnotationForOrganelleType) OrganelleApiDownloadOrganellePackageRequest {
	r.includeAnnotationType = &includeAnnotationType
	return r
}

// Output file name.
func (r OrganelleApiDownloadOrganellePackageRequest) Filename(filename string) OrganelleApiDownloadOrganellePackageRequest {
	r.filename = &filename
	return r
}

func (r OrganelleApiDownloadOrganellePackageRequest) Execute() (**os.File, *http.Response, error) {
	return r.ApiService.DownloadOrganellePackageExecute(r)
}

/*
DownloadOrganellePackage Get a organelle data package by accesions

Download a organelle data report and annotation data package.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accessions NCBI organelle assembly accessions
 @return OrganelleApiDownloadOrganellePackageRequest
*/
func (a *OrganelleApiService) DownloadOrganellePackage(ctx context.Context, accessions []string) OrganelleApiDownloadOrganellePackageRequest {
	return OrganelleApiDownloadOrganellePackageRequest{
		ApiService: a,
		ctx: ctx,
		accessions: accessions,
	}
}

// Execute executes the request
//  @return *os.File
func (a *OrganelleApiService) DownloadOrganellePackageExecute(r OrganelleApiDownloadOrganellePackageRequest) (**os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  **os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganelleApiService.DownloadOrganellePackage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organelle/accession/{accessions}/download"
	localVarPath = strings.Replace(localVarPath, "{"+"accessions"+"}", url.PathEscape(parameterToString(r.accessions, "csv")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.excludeSequence != nil {
		localVarQueryParams.Add("exclude_sequence", parameterToString(*r.excludeSequence, ""))
	}
	if r.includeAnnotationType != nil {
		t := *r.includeAnnotationType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("include_annotation_type", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("include_annotation_type", parameterToString(t, "multi"))
		}
	}
	if r.filename != nil {
		localVarQueryParams.Add("filename", parameterToString(*r.filename, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/zip"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OrganelleApiDownloadOrganellePackageByPostRequest struct {
	ctx context.Context
	ApiService *OrganelleApiService
	v2OrganelleDownloadRequest *V2OrganelleDownloadRequest
	filename *string
}

func (r OrganelleApiDownloadOrganellePackageByPostRequest) V2OrganelleDownloadRequest(v2OrganelleDownloadRequest V2OrganelleDownloadRequest) OrganelleApiDownloadOrganellePackageByPostRequest {
	r.v2OrganelleDownloadRequest = &v2OrganelleDownloadRequest
	return r
}

// Output file name.
func (r OrganelleApiDownloadOrganellePackageByPostRequest) Filename(filename string) OrganelleApiDownloadOrganellePackageByPostRequest {
	r.filename = &filename
	return r
}

func (r OrganelleApiDownloadOrganellePackageByPostRequest) Execute() (**os.File, *http.Response, error) {
	return r.ApiService.DownloadOrganellePackageByPostExecute(r)
}

/*
DownloadOrganellePackageByPost Get a organelle data package by post

Download a organelle report and annotation data package by post.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OrganelleApiDownloadOrganellePackageByPostRequest
*/
func (a *OrganelleApiService) DownloadOrganellePackageByPost(ctx context.Context) OrganelleApiDownloadOrganellePackageByPostRequest {
	return OrganelleApiDownloadOrganellePackageByPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *OrganelleApiService) DownloadOrganellePackageByPostExecute(r OrganelleApiDownloadOrganellePackageByPostRequest) (**os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  **os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganelleApiService.DownloadOrganellePackageByPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organelle/download"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.v2OrganelleDownloadRequest == nil {
		return localVarReturnValue, nil, reportError("v2OrganelleDownloadRequest is required and must be specified")
	}

	if r.filename != nil {
		localVarQueryParams.Add("filename", parameterToString(*r.filename, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/zip"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.v2OrganelleDownloadRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OrganelleApiOrganelleDatareportByAccessionRequest struct {
	ctx context.Context
	ApiService *OrganelleApiService
	accessions []string
	taxons *[]string
	organelleTypes *[]V2reportsOrganelleType
	firstReleaseDate *time.Time
	lastReleaseDate *time.Time
	taxExactMatch *bool
	sortField *string
	sortDirection *V2SortDirection
	returnedContent *V2OrganelleMetadataRequestContentType
	tableFormat *V2OrganelleMetadataRequestOrganelleTableFormat
	includeTabularHeader *V2IncludeTabularHeader
}

// NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
func (r OrganelleApiOrganelleDatareportByAccessionRequest) Taxons(taxons []string) OrganelleApiOrganelleDatareportByAccessionRequest {
	r.taxons = &taxons
	return r
}

func (r OrganelleApiOrganelleDatareportByAccessionRequest) OrganelleTypes(organelleTypes []V2reportsOrganelleType) OrganelleApiOrganelleDatareportByAccessionRequest {
	r.organelleTypes = &organelleTypes
	return r
}

// Only return organelle assemblies that were released on or after the specified date By default, do not filter.
func (r OrganelleApiOrganelleDatareportByAccessionRequest) FirstReleaseDate(firstReleaseDate time.Time) OrganelleApiOrganelleDatareportByAccessionRequest {
	r.firstReleaseDate = &firstReleaseDate
	return r
}

// Only return organelle assemblies that were released on or before to the specified date By default, do not filter.
func (r OrganelleApiOrganelleDatareportByAccessionRequest) LastReleaseDate(lastReleaseDate time.Time) OrganelleApiOrganelleDatareportByAccessionRequest {
	r.lastReleaseDate = &lastReleaseDate
	return r
}

// If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too.
func (r OrganelleApiOrganelleDatareportByAccessionRequest) TaxExactMatch(taxExactMatch bool) OrganelleApiOrganelleDatareportByAccessionRequest {
	r.taxExactMatch = &taxExactMatch
	return r
}

func (r OrganelleApiOrganelleDatareportByAccessionRequest) SortField(sortField string) OrganelleApiOrganelleDatareportByAccessionRequest {
	r.sortField = &sortField
	return r
}

func (r OrganelleApiOrganelleDatareportByAccessionRequest) SortDirection(sortDirection V2SortDirection) OrganelleApiOrganelleDatareportByAccessionRequest {
	r.sortDirection = &sortDirection
	return r
}

// Return either assembly accessions, or entire assembly-metadata records
func (r OrganelleApiOrganelleDatareportByAccessionRequest) ReturnedContent(returnedContent V2OrganelleMetadataRequestContentType) OrganelleApiOrganelleDatareportByAccessionRequest {
	r.returnedContent = &returnedContent
	return r
}

// Optional pre-defined template for processing a tabular data request
func (r OrganelleApiOrganelleDatareportByAccessionRequest) TableFormat(tableFormat V2OrganelleMetadataRequestOrganelleTableFormat) OrganelleApiOrganelleDatareportByAccessionRequest {
	r.tableFormat = &tableFormat
	return r
}

// Whether this request for tabular data should include the header row
func (r OrganelleApiOrganelleDatareportByAccessionRequest) IncludeTabularHeader(includeTabularHeader V2IncludeTabularHeader) OrganelleApiOrganelleDatareportByAccessionRequest {
	r.includeTabularHeader = &includeTabularHeader
	return r
}

func (r OrganelleApiOrganelleDatareportByAccessionRequest) Execute() (*V2reportsOrganelleDataReports, *http.Response, error) {
	return r.ApiService.OrganelleDatareportByAccessionExecute(r)
}

/*
OrganelleDatareportByAccession Get Organelle dataset report by accession

Get Organelle dataset report by accession.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accessions NCBI assembly accession
 @return OrganelleApiOrganelleDatareportByAccessionRequest
*/
func (a *OrganelleApiService) OrganelleDatareportByAccession(ctx context.Context, accessions []string) OrganelleApiOrganelleDatareportByAccessionRequest {
	return OrganelleApiOrganelleDatareportByAccessionRequest{
		ApiService: a,
		ctx: ctx,
		accessions: accessions,
	}
}

// Execute executes the request
//  @return V2reportsOrganelleDataReports
func (a *OrganelleApiService) OrganelleDatareportByAccessionExecute(r OrganelleApiOrganelleDatareportByAccessionRequest) (*V2reportsOrganelleDataReports, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V2reportsOrganelleDataReports
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganelleApiService.OrganelleDatareportByAccession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organelle/accessions/{accessions}/dataset_report"
	localVarPath = strings.Replace(localVarPath, "{"+"accessions"+"}", url.PathEscape(parameterToString(r.accessions, "csv")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.taxons != nil {
		t := *r.taxons
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("taxons", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("taxons", parameterToString(t, "multi"))
		}
	}
	if r.organelleTypes != nil {
		t := *r.organelleTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("organelle_types", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("organelle_types", parameterToString(t, "multi"))
		}
	}
	if r.firstReleaseDate != nil {
		localVarQueryParams.Add("first_release_date", parameterToString(*r.firstReleaseDate, ""))
	}
	if r.lastReleaseDate != nil {
		localVarQueryParams.Add("last_release_date", parameterToString(*r.lastReleaseDate, ""))
	}
	if r.taxExactMatch != nil {
		localVarQueryParams.Add("tax_exact_match", parameterToString(*r.taxExactMatch, ""))
	}
	if r.sortField != nil {
		localVarQueryParams.Add("sort.field", parameterToString(*r.sortField, ""))
	}
	if r.sortDirection != nil {
		localVarQueryParams.Add("sort.direction", parameterToString(*r.sortDirection, ""))
	}
	if r.returnedContent != nil {
		localVarQueryParams.Add("returned_content", parameterToString(*r.returnedContent, ""))
	}
	if r.tableFormat != nil {
		localVarQueryParams.Add("table_format", parameterToString(*r.tableFormat, ""))
	}
	if r.includeTabularHeader != nil {
		localVarQueryParams.Add("include_tabular_header", parameterToString(*r.includeTabularHeader, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OrganelleApiOrganelleDatareportByPostRequest struct {
	ctx context.Context
	ApiService *OrganelleApiService
	v2OrganelleMetadataRequest *V2OrganelleMetadataRequest
}

func (r OrganelleApiOrganelleDatareportByPostRequest) V2OrganelleMetadataRequest(v2OrganelleMetadataRequest V2OrganelleMetadataRequest) OrganelleApiOrganelleDatareportByPostRequest {
	r.v2OrganelleMetadataRequest = &v2OrganelleMetadataRequest
	return r
}

func (r OrganelleApiOrganelleDatareportByPostRequest) Execute() (*V2reportsOrganelleDataReports, *http.Response, error) {
	return r.ApiService.OrganelleDatareportByPostExecute(r)
}

/*
OrganelleDatareportByPost Get Organelle dataset report by http post

Get Organelle dataset report by http post.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OrganelleApiOrganelleDatareportByPostRequest
*/
func (a *OrganelleApiService) OrganelleDatareportByPost(ctx context.Context) OrganelleApiOrganelleDatareportByPostRequest {
	return OrganelleApiOrganelleDatareportByPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V2reportsOrganelleDataReports
func (a *OrganelleApiService) OrganelleDatareportByPostExecute(r OrganelleApiOrganelleDatareportByPostRequest) (*V2reportsOrganelleDataReports, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V2reportsOrganelleDataReports
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganelleApiService.OrganelleDatareportByPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organelle/dataset_report"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.v2OrganelleMetadataRequest == nil {
		return localVarReturnValue, nil, reportError("v2OrganelleMetadataRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.v2OrganelleMetadataRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OrganelleApiOrganelleDatareportByTaxonRequest struct {
	ctx context.Context
	ApiService *OrganelleApiService
	taxons []string
	organelleTypes *[]V2reportsOrganelleType
	firstReleaseDate *time.Time
	lastReleaseDate *time.Time
	taxExactMatch *bool
	sortField *string
	sortDirection *V2SortDirection
	returnedContent *V2OrganelleMetadataRequestContentType
	pageSize *int32
	pageToken *string
	tableFormat *V2OrganelleMetadataRequestOrganelleTableFormat
	includeTabularHeader *V2IncludeTabularHeader
}

func (r OrganelleApiOrganelleDatareportByTaxonRequest) OrganelleTypes(organelleTypes []V2reportsOrganelleType) OrganelleApiOrganelleDatareportByTaxonRequest {
	r.organelleTypes = &organelleTypes
	return r
}

// Only return organelle assemblies that were released on or after the specified date By default, do not filter.
func (r OrganelleApiOrganelleDatareportByTaxonRequest) FirstReleaseDate(firstReleaseDate time.Time) OrganelleApiOrganelleDatareportByTaxonRequest {
	r.firstReleaseDate = &firstReleaseDate
	return r
}

// Only return organelle assemblies that were released on or before to the specified date By default, do not filter.
func (r OrganelleApiOrganelleDatareportByTaxonRequest) LastReleaseDate(lastReleaseDate time.Time) OrganelleApiOrganelleDatareportByTaxonRequest {
	r.lastReleaseDate = &lastReleaseDate
	return r
}

// If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too.
func (r OrganelleApiOrganelleDatareportByTaxonRequest) TaxExactMatch(taxExactMatch bool) OrganelleApiOrganelleDatareportByTaxonRequest {
	r.taxExactMatch = &taxExactMatch
	return r
}

func (r OrganelleApiOrganelleDatareportByTaxonRequest) SortField(sortField string) OrganelleApiOrganelleDatareportByTaxonRequest {
	r.sortField = &sortField
	return r
}

func (r OrganelleApiOrganelleDatareportByTaxonRequest) SortDirection(sortDirection V2SortDirection) OrganelleApiOrganelleDatareportByTaxonRequest {
	r.sortDirection = &sortDirection
	return r
}

// Return either assembly accessions, or entire assembly-metadata records
func (r OrganelleApiOrganelleDatareportByTaxonRequest) ReturnedContent(returnedContent V2OrganelleMetadataRequestContentType) OrganelleApiOrganelleDatareportByTaxonRequest {
	r.returnedContent = &returnedContent
	return r
}

// The maximum number of organelle assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results.
func (r OrganelleApiOrganelleDatareportByTaxonRequest) PageSize(pageSize int32) OrganelleApiOrganelleDatareportByTaxonRequest {
	r.pageSize = &pageSize
	return r
}

// A page token is returned from an &#x60;OrganelleMetadata&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;OrganelleMetadata&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved.
func (r OrganelleApiOrganelleDatareportByTaxonRequest) PageToken(pageToken string) OrganelleApiOrganelleDatareportByTaxonRequest {
	r.pageToken = &pageToken
	return r
}

// Optional pre-defined template for processing a tabular data request
func (r OrganelleApiOrganelleDatareportByTaxonRequest) TableFormat(tableFormat V2OrganelleMetadataRequestOrganelleTableFormat) OrganelleApiOrganelleDatareportByTaxonRequest {
	r.tableFormat = &tableFormat
	return r
}

// Whether this request for tabular data should include the header row
func (r OrganelleApiOrganelleDatareportByTaxonRequest) IncludeTabularHeader(includeTabularHeader V2IncludeTabularHeader) OrganelleApiOrganelleDatareportByTaxonRequest {
	r.includeTabularHeader = &includeTabularHeader
	return r
}

func (r OrganelleApiOrganelleDatareportByTaxonRequest) Execute() (*V2reportsOrganelleDataReports, *http.Response, error) {
	return r.ApiService.OrganelleDatareportByTaxonExecute(r)
}

/*
OrganelleDatareportByTaxon Get Organelle dataset report by taxons

Get Organelle dataset report by taxons.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taxons NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
 @return OrganelleApiOrganelleDatareportByTaxonRequest
*/
func (a *OrganelleApiService) OrganelleDatareportByTaxon(ctx context.Context, taxons []string) OrganelleApiOrganelleDatareportByTaxonRequest {
	return OrganelleApiOrganelleDatareportByTaxonRequest{
		ApiService: a,
		ctx: ctx,
		taxons: taxons,
	}
}

// Execute executes the request
//  @return V2reportsOrganelleDataReports
func (a *OrganelleApiService) OrganelleDatareportByTaxonExecute(r OrganelleApiOrganelleDatareportByTaxonRequest) (*V2reportsOrganelleDataReports, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V2reportsOrganelleDataReports
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganelleApiService.OrganelleDatareportByTaxon")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organelle/taxon/{taxons}/dataset_report"
	localVarPath = strings.Replace(localVarPath, "{"+"taxons"+"}", url.PathEscape(parameterToString(r.taxons, "csv")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.organelleTypes != nil {
		t := *r.organelleTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("organelle_types", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("organelle_types", parameterToString(t, "multi"))
		}
	}
	if r.firstReleaseDate != nil {
		localVarQueryParams.Add("first_release_date", parameterToString(*r.firstReleaseDate, ""))
	}
	if r.lastReleaseDate != nil {
		localVarQueryParams.Add("last_release_date", parameterToString(*r.lastReleaseDate, ""))
	}
	if r.taxExactMatch != nil {
		localVarQueryParams.Add("tax_exact_match", parameterToString(*r.taxExactMatch, ""))
	}
	if r.sortField != nil {
		localVarQueryParams.Add("sort.field", parameterToString(*r.sortField, ""))
	}
	if r.sortDirection != nil {
		localVarQueryParams.Add("sort.direction", parameterToString(*r.sortDirection, ""))
	}
	if r.returnedContent != nil {
		localVarQueryParams.Add("returned_content", parameterToString(*r.returnedContent, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.pageToken != nil {
		localVarQueryParams.Add("page_token", parameterToString(*r.pageToken, ""))
	}
	if r.tableFormat != nil {
		localVarQueryParams.Add("table_format", parameterToString(*r.tableFormat, ""))
	}
	if r.includeTabularHeader != nil {
		localVarQueryParams.Add("include_tabular_header", parameterToString(*r.includeTabularHeader, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
