// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ApplicationTemplate undocumented
type ApplicationTemplate struct {
	// Entity is the base model of ApplicationTemplate
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// HomePageURL undocumented
	HomePageURL *string `json:"homePageUrl,omitempty"`
	// SupportedSingleSignOnModes undocumented
	SupportedSingleSignOnModes []string `json:"supportedSingleSignOnModes,omitempty"`
	// SupportedProvisioningTypes undocumented
	SupportedProvisioningTypes []string `json:"supportedProvisioningTypes,omitempty"`
	// LogoURL undocumented
	LogoURL *string `json:"logoUrl,omitempty"`
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
	// Publisher undocumented
	Publisher *string `json:"publisher,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
}
