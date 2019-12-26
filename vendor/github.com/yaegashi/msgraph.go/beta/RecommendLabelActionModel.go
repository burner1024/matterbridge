// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RecommendLabelAction undocumented
type RecommendLabelAction struct {
	// InformationProtectionAction is the base model of RecommendLabelAction
	InformationProtectionAction
	// Label undocumented
	Label *LabelDetails `json:"label,omitempty"`
	// ResponsibleSensitiveTypeIDs undocumented
	ResponsibleSensitiveTypeIDs []UUID `json:"responsibleSensitiveTypeIds,omitempty"`
	// Actions undocumented
	Actions []InformationProtectionAction `json:"actions,omitempty"`
	// ActionSource undocumented
	ActionSource *ActionSource `json:"actionSource,omitempty"`
}
