# \KPIEntriesApi

All URIs are relative to *https://YOURDOMAIN.simplekpi.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKPIEntry**](KPIEntriesApi.md#AddKPIEntry) | **Post** /kpientries | Add KPI Entry
[**DeleteKPIEntry**](KPIEntriesApi.md#DeleteKPIEntry) | **Delete** /kpientries/{kpientryid} | Delete KPI Entry
[**GetAllKPIEntries**](KPIEntriesApi.md#GetAllKPIEntries) | **Get** /kpientries | Get all KPI Entries
[**GetKPIEntry**](KPIEntriesApi.md#GetKPIEntry) | **Get** /kpientries/{kpientryid} | Get KPI Entry
[**UpdateKPIEntry**](KPIEntriesApi.md#UpdateKPIEntry) | **Put** /kpientries/{kpientryid} | Update KPI Entry



## AddKPIEntry

> KpiEntry AddKPIEntry(ctx, kpiEntry)
Add KPI Entry

The KPI entries are filtered based on the search query string. All the search criteria is optional and we will return a maximum of 500 entries per page. If the result set has the amount of rows you set `&rows=100`, it's your responsibility to check the next page to see if there are any more -- you do this by adding &page=2 to the query, then &page=3 and so on.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kpiEntry** | [**KpiEntry**](KpiEntry.md)| KPI Entry Object | 

### Return type

[**KpiEntry**](KPIEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKPIEntry

> DeleteKPIEntry(ctx, kpientryid)
Delete KPI Entry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kpientryid** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllKPIEntries

> []KpiEntry GetAllKPIEntries(ctx, dateFrom, dateTo, optional)
Get all KPI Entries

The KPI entries are filtered based on the search query string. All the search criteria is optional and we will return a maximum of 500 entries per page. If the result set has the amount of rows you set `&rows=100`, it's your responsibility to check the next page to see if there are any more -- you do this by adding &page=2 to the query, then &page=3 and so on.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dateFrom** | **string**|  | 
**dateTo** | **string**|  | 
 **optional** | ***GetAllKPIEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllKPIEntriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userid** | **optional.Int32**|  | 
 **kpiid** | **optional.Int32**|  | 
 **rows** | **optional.Int32**|  | 
 **page** | **optional.Int32**|  | 

### Return type

[**[]KpiEntry**](KPIEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKPIEntry

> KpiEntry GetKPIEntry(ctx, kpientryid)
Get KPI Entry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kpientryid** | **int64**|  | 

### Return type

[**KpiEntry**](KPIEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKPIEntry

> KpiEntry UpdateKPIEntry(ctx, kpientryid, kpiEntry)
Update KPI Entry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kpientryid** | **int64**|  | 
**kpiEntry** | [**KpiEntry**](KpiEntry.md)| KPI Entry Object | 

### Return type

[**KpiEntry**](KPIEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

