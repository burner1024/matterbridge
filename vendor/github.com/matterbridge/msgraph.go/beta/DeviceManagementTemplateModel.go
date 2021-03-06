// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// DeviceManagementTemplate Entity that represents a defined collection of device settings
type DeviceManagementTemplate struct {
	// Entity is the base model of DeviceManagementTemplate
	Entity
	// DisplayName The template's display name
	DisplayName *string `json:"displayName,omitempty"`
	// Description The template's description
	Description *string `json:"description,omitempty"`
	// VersionInfo The template's version information
	VersionInfo *string `json:"versionInfo,omitempty"`
	// IsDeprecated The template is deprecated or not. Intents cannot be created from a deprecated template.
	IsDeprecated *bool `json:"isDeprecated,omitempty"`
	// IntentCount Number of Intents created from this template.
	IntentCount *int `json:"intentCount,omitempty"`
	// TemplateType The template's type.
	TemplateType *DeviceManagementTemplateType `json:"templateType,omitempty"`
	// PlatformType The template's platform.
	PlatformType *PolicyPlatformType `json:"platformType,omitempty"`
	// PublishedDateTime When the template was published
	PublishedDateTime *time.Time `json:"publishedDateTime,omitempty"`
	// Settings undocumented
	Settings []DeviceManagementSettingInstance `json:"settings,omitempty"`
	// Categories undocumented
	Categories []DeviceManagementTemplateSettingCategory `json:"categories,omitempty"`
	// MigratableTo undocumented
	MigratableTo []DeviceManagementTemplate `json:"migratableTo,omitempty"`
}
