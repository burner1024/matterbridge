// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RoleAssignmentScopeType undocumented
type RoleAssignmentScopeType int

const (
	// RoleAssignmentScopeTypeVResourceScope undocumented
	RoleAssignmentScopeTypeVResourceScope RoleAssignmentScopeType = 0
	// RoleAssignmentScopeTypeVAllDevices undocumented
	RoleAssignmentScopeTypeVAllDevices RoleAssignmentScopeType = 1
	// RoleAssignmentScopeTypeVAllLicensedUsers undocumented
	RoleAssignmentScopeTypeVAllLicensedUsers RoleAssignmentScopeType = 2
	// RoleAssignmentScopeTypeVAllDevicesAndLicensedUsers undocumented
	RoleAssignmentScopeTypeVAllDevicesAndLicensedUsers RoleAssignmentScopeType = 3
)

// RoleAssignmentScopeTypePResourceScope returns a pointer to RoleAssignmentScopeTypeVResourceScope
func RoleAssignmentScopeTypePResourceScope() *RoleAssignmentScopeType {
	v := RoleAssignmentScopeTypeVResourceScope
	return &v
}

// RoleAssignmentScopeTypePAllDevices returns a pointer to RoleAssignmentScopeTypeVAllDevices
func RoleAssignmentScopeTypePAllDevices() *RoleAssignmentScopeType {
	v := RoleAssignmentScopeTypeVAllDevices
	return &v
}

// RoleAssignmentScopeTypePAllLicensedUsers returns a pointer to RoleAssignmentScopeTypeVAllLicensedUsers
func RoleAssignmentScopeTypePAllLicensedUsers() *RoleAssignmentScopeType {
	v := RoleAssignmentScopeTypeVAllLicensedUsers
	return &v
}

// RoleAssignmentScopeTypePAllDevicesAndLicensedUsers returns a pointer to RoleAssignmentScopeTypeVAllDevicesAndLicensedUsers
func RoleAssignmentScopeTypePAllDevicesAndLicensedUsers() *RoleAssignmentScopeType {
	v := RoleAssignmentScopeTypeVAllDevicesAndLicensedUsers
	return &v
}
