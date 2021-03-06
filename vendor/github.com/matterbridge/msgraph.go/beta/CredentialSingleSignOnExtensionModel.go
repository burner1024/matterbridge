// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CredentialSingleSignOnExtension undocumented
type CredentialSingleSignOnExtension struct {
	// SingleSignOnExtension is the base model of CredentialSingleSignOnExtension
	SingleSignOnExtension
	// ExtensionIdentifier Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
	ExtensionIdentifier *string `json:"extensionIdentifier,omitempty"`
	// TeamIdentifier Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
	TeamIdentifier *string `json:"teamIdentifier,omitempty"`
	// Domains Gets or sets a list of hosts or domain names for which the app extension performs SSO.
	Domains []string `json:"domains,omitempty"`
	// Realm Gets or sets the case-sensitive realm name for this profile.
	Realm *string `json:"realm,omitempty"`
	// Configurations Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
	Configurations []KeyTypedValuePair `json:"configurations,omitempty"`
}
