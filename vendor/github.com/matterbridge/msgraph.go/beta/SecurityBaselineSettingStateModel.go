// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SecurityBaselineSettingState The security baseline compliance state of a setting for a device
type SecurityBaselineSettingState struct {
	// Entity is the base model of SecurityBaselineSettingState
	Entity
	// SettingName The setting name that is being reported
	SettingName *string `json:"settingName,omitempty"`
	// State The compliance state of the security baseline setting
	State *SecurityBaselineComplianceState `json:"state,omitempty"`
	// SettingCategoryID The setting category id which this setting belongs to
	SettingCategoryID *string `json:"settingCategoryId,omitempty"`
}
