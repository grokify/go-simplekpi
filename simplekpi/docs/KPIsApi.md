# \KPIsApi

All URIs are relative to *https://YOURDOMAIN.simplekpi.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllKPIs**](KPIsApi.md#GetAllKPIs) | **Get** /kpis | Get all KPIs
[**GetKPI**](KPIsApi.md#GetKPI) | **Get** /kpis/{kpiId} | Get a KPI



## GetAllKPIs

> []Kpi GetAllKPIs(ctx, )
Get all KPIs

Returns data on all KPIs. There are no parameters for this API.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Kpi**](KPI.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKPI

> Kpi GetKPI(ctx, kpiId)
Get a KPI

Returns data on a single KPIs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kpiId** | **int64**|  | 

### Return type

[**Kpi**](KPI.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

