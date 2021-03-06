// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationAssignmentStatus undocumented
type EducationAssignmentStatus int

const (
	// EducationAssignmentStatusVDraft undocumented
	EducationAssignmentStatusVDraft EducationAssignmentStatus = 0
	// EducationAssignmentStatusVPublished undocumented
	EducationAssignmentStatusVPublished EducationAssignmentStatus = 1
	// EducationAssignmentStatusVAssigned undocumented
	EducationAssignmentStatusVAssigned EducationAssignmentStatus = 2
	// EducationAssignmentStatusVUnknownFutureValue undocumented
	EducationAssignmentStatusVUnknownFutureValue EducationAssignmentStatus = 3
)

// EducationAssignmentStatusPDraft returns a pointer to EducationAssignmentStatusVDraft
func EducationAssignmentStatusPDraft() *EducationAssignmentStatus {
	v := EducationAssignmentStatusVDraft
	return &v
}

// EducationAssignmentStatusPPublished returns a pointer to EducationAssignmentStatusVPublished
func EducationAssignmentStatusPPublished() *EducationAssignmentStatus {
	v := EducationAssignmentStatusVPublished
	return &v
}

// EducationAssignmentStatusPAssigned returns a pointer to EducationAssignmentStatusVAssigned
func EducationAssignmentStatusPAssigned() *EducationAssignmentStatus {
	v := EducationAssignmentStatusVAssigned
	return &v
}

// EducationAssignmentStatusPUnknownFutureValue returns a pointer to EducationAssignmentStatusVUnknownFutureValue
func EducationAssignmentStatusPUnknownFutureValue() *EducationAssignmentStatus {
	v := EducationAssignmentStatusVUnknownFutureValue
	return &v
}
