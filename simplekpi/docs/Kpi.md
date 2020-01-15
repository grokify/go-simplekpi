# Kpi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregateFunction** | **string** | The aggregate function determines how the KPI is calculated and can be either AVG (Average) or SUM (Total Sum) | 
**CategoryId** | **int64** | The id of the category the KPI is in | 
**CreatedAt** | **string** | The UTC date and time the KPI was created. Date time format without timezone, e.g. &#x60;2019-01-01T00:00:00&#x60; | [optional] 
**Description** | **string** | The description of the KPI | [optional] 
**FrequencyId** | **string** |  | 
**IconId** | **int64** | The id of the icon to assign to the KPI | 
**Id** | **int64** | Automatically generated for the KPI | [optional] 
**IsActive** | **bool** | Active KPIs can have date entered against them otherwise they are display only KPIs | 
**IsCalculated** | **bool** | Calculated KPIs cannot be amended via the API and must be added / amended in the interface | [optional] 
**Name** | **string** | The name of the KPI | 
**SortOrder** | **int32** | The display order of the KPI | 
**TargetDefault** | **float32** | The default target value for the KPI. If left blank or null the KPI will not have a target | [optional] 
**UnitId** | **int64** | The id of the unit of measure to assign to the KPI | 
**UpdatedAt** | **string** | The UTC date and time the KPI was updated. Date time format without timezone, e.g. &#x60;2019-01-01T00:00:00&#x60; | [optional] 
**ValueDirection** | **string** | The value direction is case sensitive and can only be U(p), D(own) and N(one) | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


