// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ServicePrincipal undocumented
type ServicePrincipal struct {
	// DirectoryObject is the base model of ServicePrincipal
	DirectoryObject
	// AccountEnabled undocumented
	AccountEnabled *bool `json:"accountEnabled,omitempty"`
	// AddIns undocumented
	AddIns []AddIn `json:"addIns,omitempty"`
	// AppDisplayName undocumented
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// ApplicationTemplateID undocumented
	ApplicationTemplateID *string `json:"applicationTemplateId,omitempty"`
	// AppOwnerOrganizationID undocumented
	AppOwnerOrganizationID *UUID `json:"appOwnerOrganizationId,omitempty"`
	// AppRoleAssignmentRequired undocumented
	AppRoleAssignmentRequired *bool `json:"appRoleAssignmentRequired,omitempty"`
	// AppRoles undocumented
	AppRoles []AppRole `json:"appRoles,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Homepage undocumented
	Homepage *string `json:"homepage,omitempty"`
	// KeyCredentials undocumented
	KeyCredentials []KeyCredential `json:"keyCredentials,omitempty"`
	// Info undocumented
	Info *InformationalURL `json:"info,omitempty"`
	// LogoutURL undocumented
	LogoutURL *string `json:"logoutUrl,omitempty"`
	// NotificationEmailAddresses undocumented
	NotificationEmailAddresses []string `json:"notificationEmailAddresses,omitempty"`
	// PublishedPermissionScopes undocumented
	PublishedPermissionScopes []PermissionScope `json:"publishedPermissionScopes,omitempty"`
	// PasswordCredentials undocumented
	PasswordCredentials []PasswordCredential `json:"passwordCredentials,omitempty"`
	// PreferredTokenSigningKeyThumbprint undocumented
	PreferredTokenSigningKeyThumbprint *string `json:"preferredTokenSigningKeyThumbprint,omitempty"`
	// PublisherName undocumented
	PublisherName *string `json:"publisherName,omitempty"`
	// ReplyUrls undocumented
	ReplyUrls []string `json:"replyUrls,omitempty"`
	// SamlMetadataURL undocumented
	SamlMetadataURL *string `json:"samlMetadataUrl,omitempty"`
	// ServicePrincipalNames undocumented
	ServicePrincipalNames []string `json:"servicePrincipalNames,omitempty"`
	// Tags undocumented
	Tags []string `json:"tags,omitempty"`
	// AppRoleAssignedTo undocumented
	AppRoleAssignedTo []AppRoleAssignment `json:"appRoleAssignedTo,omitempty"`
	// AppRoleAssignments undocumented
	AppRoleAssignments []AppRoleAssignment `json:"appRoleAssignments,omitempty"`
	// Oauth2PermissionGrants undocumented
	Oauth2PermissionGrants []OAuth2PermissionGrant `json:"oauth2PermissionGrants,omitempty"`
	// MemberOf undocumented
	MemberOf []DirectoryObject `json:"memberOf,omitempty"`
	// TransitiveMemberOf undocumented
	TransitiveMemberOf []DirectoryObject `json:"transitiveMemberOf,omitempty"`
	// CreatedObjects undocumented
	CreatedObjects []DirectoryObject `json:"createdObjects,omitempty"`
	// LicenseDetails undocumented
	LicenseDetails []LicenseDetails `json:"licenseDetails,omitempty"`
	// Owners undocumented
	Owners []DirectoryObject `json:"owners,omitempty"`
	// OwnedObjects undocumented
	OwnedObjects []DirectoryObject `json:"ownedObjects,omitempty"`
	// Policies undocumented
	Policies []DirectoryObject `json:"policies,omitempty"`
	// Synchronization undocumented
	Synchronization *Synchronization `json:"synchronization,omitempty"`
}
