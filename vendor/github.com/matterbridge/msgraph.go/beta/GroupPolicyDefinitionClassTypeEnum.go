// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupPolicyDefinitionClassType undocumented
type GroupPolicyDefinitionClassType int

const (
	// GroupPolicyDefinitionClassTypeVUser undocumented
	GroupPolicyDefinitionClassTypeVUser GroupPolicyDefinitionClassType = 0
	// GroupPolicyDefinitionClassTypeVMachine undocumented
	GroupPolicyDefinitionClassTypeVMachine GroupPolicyDefinitionClassType = 1
)

// GroupPolicyDefinitionClassTypePUser returns a pointer to GroupPolicyDefinitionClassTypeVUser
func GroupPolicyDefinitionClassTypePUser() *GroupPolicyDefinitionClassType {
	v := GroupPolicyDefinitionClassTypeVUser
	return &v
}

// GroupPolicyDefinitionClassTypePMachine returns a pointer to GroupPolicyDefinitionClassTypeVMachine
func GroupPolicyDefinitionClassTypePMachine() *GroupPolicyDefinitionClassType {
	v := GroupPolicyDefinitionClassTypeVMachine
	return &v
}