// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementExchangeOnPremisesPolicy Singleton entity which represents the Exchange OnPremises policy configured for a tenant.
type DeviceManagementExchangeOnPremisesPolicy struct {
	// Entity is the base model of DeviceManagementExchangeOnPremisesPolicy
	Entity
	// NotificationContent Notification text that will be sent to users quarantined by this policy. This is UTF8 encoded byte array HTML.
	NotificationContent *Binary `json:"notificationContent,omitempty"`
	// DefaultAccessLevel Default access state in Exchange. This rule applies globally to the entire Exchange organization
	DefaultAccessLevel *DeviceManagementExchangeAccessLevel `json:"defaultAccessLevel,omitempty"`
	// AccessRules The list of device access rules in Exchange. The access rules apply globally to the entire Exchange organization
	AccessRules []DeviceManagementExchangeAccessRule `json:"accessRules,omitempty"`
	// KnownDeviceClasses The list of device classes known to Exchange
	KnownDeviceClasses []DeviceManagementExchangeDeviceClass `json:"knownDeviceClasses,omitempty"`
	// ConditionalAccessSettings undocumented
	ConditionalAccessSettings *OnPremisesConditionalAccessSettings `json:"conditionalAccessSettings,omitempty"`
}
