// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MobileAppPolicySetItem A class containing the properties used for mobile app PolicySetItem.
type MobileAppPolicySetItem struct {
	// PolicySetItem is the base model of MobileAppPolicySetItem
	PolicySetItem
	// Intent Install intent of the MobileAppPolicySetItem.
	Intent *InstallIntent `json:"intent,omitempty"`
	// Settings Settings of the MobileAppPolicySetItem.
	Settings *MobileAppAssignmentSettings `json:"settings,omitempty"`
}
