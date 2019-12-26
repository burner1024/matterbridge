// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// EducationalActivity undocumented
type EducationalActivity struct {
	// ItemFacet is the base model of EducationalActivity
	ItemFacet
	// CompletionMonthYear undocumented
	CompletionMonthYear *time.Time `json:"completionMonthYear,omitempty"`
	// EndMonthYear undocumented
	EndMonthYear *time.Time `json:"endMonthYear,omitempty"`
	// Institution undocumented
	Institution *InstitutionData `json:"institution,omitempty"`
	// Program undocumented
	Program *EducationalActivityDetail `json:"program,omitempty"`
	// StartMonthYear undocumented
	StartMonthYear *time.Time `json:"startMonthYear,omitempty"`
}
