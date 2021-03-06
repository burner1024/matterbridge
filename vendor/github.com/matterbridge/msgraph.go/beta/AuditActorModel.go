// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AuditActor undocumented
type AuditActor struct {
	// Object is the base model of AuditActor
	Object
	// Type Actor Type.
	Type *string `json:"type,omitempty"`
	// UserPermissions List of user permissions when the audit was performed.
	UserPermissions []string `json:"userPermissions,omitempty"`
	// ApplicationID AAD Application Id.
	ApplicationID *string `json:"applicationId,omitempty"`
	// ApplicationDisplayName Name of the Application.
	ApplicationDisplayName *string `json:"applicationDisplayName,omitempty"`
	// UserPrincipalName User Principal Name (UPN).
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// ServicePrincipalName Service Principal Name (SPN).
	ServicePrincipalName *string `json:"servicePrincipalName,omitempty"`
	// IPAddress IPAddress.
	IPAddress *string `json:"ipAddress,omitempty"`
	// UserID User Id.
	UserID *string `json:"userId,omitempty"`
	// UserRoleScopeTags List of user scope tags when the audit was performed.
	UserRoleScopeTags []RoleScopeTagInfo `json:"userRoleScopeTags,omitempty"`
}
