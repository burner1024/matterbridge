// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// EducationAssignmentGrade undocumented
type EducationAssignmentGrade struct {
	// Object is the base model of EducationAssignmentGrade
	Object
	// GradedBy undocumented
	GradedBy *IdentitySet `json:"gradedBy,omitempty"`
	// GradedDateTime undocumented
	GradedDateTime *time.Time `json:"gradedDateTime,omitempty"`
}
