// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AttributeMappingSourceType undocumented
type AttributeMappingSourceType int

const (
	// AttributeMappingSourceTypeVAttribute undocumented
	AttributeMappingSourceTypeVAttribute AttributeMappingSourceType = 0
	// AttributeMappingSourceTypeVConstant undocumented
	AttributeMappingSourceTypeVConstant AttributeMappingSourceType = 1
	// AttributeMappingSourceTypeVFunction undocumented
	AttributeMappingSourceTypeVFunction AttributeMappingSourceType = 2
)

// AttributeMappingSourceTypePAttribute returns a pointer to AttributeMappingSourceTypeVAttribute
func AttributeMappingSourceTypePAttribute() *AttributeMappingSourceType {
	v := AttributeMappingSourceTypeVAttribute
	return &v
}

// AttributeMappingSourceTypePConstant returns a pointer to AttributeMappingSourceTypeVConstant
func AttributeMappingSourceTypePConstant() *AttributeMappingSourceType {
	v := AttributeMappingSourceTypeVConstant
	return &v
}

// AttributeMappingSourceTypePFunction returns a pointer to AttributeMappingSourceTypeVFunction
func AttributeMappingSourceTypePFunction() *AttributeMappingSourceType {
	v := AttributeMappingSourceTypeVFunction
	return &v
}
