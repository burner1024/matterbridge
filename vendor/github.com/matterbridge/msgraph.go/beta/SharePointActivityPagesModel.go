// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// SharePointActivityPages undocumented
type SharePointActivityPages struct {
	// Entity is the base model of SharePointActivityPages
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// VisitedPageCount undocumented
	VisitedPageCount *int `json:"visitedPageCount,omitempty"`
	// ReportDate undocumented
	ReportDate *time.Time `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
