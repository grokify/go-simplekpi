/*
 * SimpleKPI API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package simplekpi

import (
	"time"
)

type Entry struct {
	Actual    float64   `json:"actual,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	EntryDate time.Time `json:"entry_date,omitempty"`
	Id        int64     `json:"id,omitempty"`
	KpiId     int64     `json:"kpi_id,omitempty"`
	Notes     string    `json:"notes,omitempty"`
	Target    float64   `json:"target,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UserId    int64     `json:"user_id,omitempty"`
}
