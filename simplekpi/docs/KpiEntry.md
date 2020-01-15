# KpiEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actual** | **float64** | The actual value cannot be null if the target and notes are both null | [optional] 
**CreatedAt** | **string** | The UTC date and time the KPI entry was created. Date time format without timezone, e.g. &#x60;2019-01-01T00:00:00&#x60; | [optional] 
**EntryDate** | **string** | The date of the entry. Date time format without timezone, e.g. &#x60;2019-01-01T00:00:00&#x60; | 
**Id** | **int64** | Automatically generated for the KPI entry | [optional] 
**KpiId** | **int64** | The kpi must be active and cannot be a calculated KPI. The KPI must also be assigned to the user | 
**Notes** | **string** | The note associated with the KPI entry | [optional] 
**Target** | **float64** | The target value of the entry. This value will be ignored if the KPI has a null target | [optional] 
**UpdatedAt** | **string** | The UTC date and time the KPI entry was updated. Date time format without timezone, e.g. &#x60;2019-01-01T00:00:00&#x60; | [optional] 
**UserId** | **int64** | An id of an active user to assign to the KPI entry | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


