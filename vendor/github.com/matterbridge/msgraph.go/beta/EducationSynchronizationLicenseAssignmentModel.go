// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationSynchronizationLicenseAssignment undocumented
type EducationSynchronizationLicenseAssignment struct {
	// Object is the base model of EducationSynchronizationLicenseAssignment
	Object
	// AppliesTo undocumented
	AppliesTo *EducationUserRole `json:"appliesTo,omitempty"`
	// SKUIDs undocumented
	SKUIDs []string `json:"skuIds,omitempty"`
}
