// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// GroupPolicyConfigurationAssignment The group policy configuration assignment entity assigns one or more AAD groups to a specific group policy configuration.
type GroupPolicyConfigurationAssignment struct {
	// Entity is the base model of GroupPolicyConfigurationAssignment
	Entity
	// LastModifiedDateTime The date and time the entity was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Target The type of groups targeted the group policy configuration.
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}
