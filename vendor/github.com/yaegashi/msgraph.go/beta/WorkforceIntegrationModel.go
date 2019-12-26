// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkforceIntegration undocumented
type WorkforceIntegration struct {
	// ChangeTrackedEntity is the base model of WorkforceIntegration
	ChangeTrackedEntity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// APIVersion undocumented
	APIVersion *int `json:"apiVersion,omitempty"`
	// Encryption undocumented
	Encryption *WorkforceIntegrationEncryption `json:"encryption,omitempty"`
	// IsActive undocumented
	IsActive *bool `json:"isActive,omitempty"`
	// URL undocumented
	URL *string `json:"url,omitempty"`
	// Supports undocumented
	Supports *WorkforceIntegrationSupportedEntities `json:"supports,omitempty"`
}
