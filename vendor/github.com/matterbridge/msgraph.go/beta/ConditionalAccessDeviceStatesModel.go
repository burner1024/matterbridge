// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConditionalAccessDeviceStates undocumented
type ConditionalAccessDeviceStates struct {
	// Object is the base model of ConditionalAccessDeviceStates
	Object
	// IncludeStates undocumented
	IncludeStates []string `json:"includeStates,omitempty"`
	// ExcludeStates undocumented
	ExcludeStates []string `json:"excludeStates,omitempty"`
}
