// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsKioskProfile undocumented
type WindowsKioskProfile struct {
	// Object is the base model of WindowsKioskProfile
	Object
	// ProfileID Key of the entity.
	ProfileID *string `json:"profileId,omitempty"`
	// ProfileName This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
	ProfileName *string `json:"profileName,omitempty"`
	// AppConfiguration The App configuration that will be used for this kiosk configuration.
	AppConfiguration *WindowsKioskAppConfiguration `json:"appConfiguration,omitempty"`
	// UserAccountsConfiguration The user accounts that will be locked to this kiosk configuration. This collection can contain a maximum of 100 elements.
	UserAccountsConfiguration []WindowsKioskUser `json:"userAccountsConfiguration,omitempty"`
}
