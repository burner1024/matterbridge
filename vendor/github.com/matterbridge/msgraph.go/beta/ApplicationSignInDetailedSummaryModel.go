// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ApplicationSignInDetailedSummary undocumented
type ApplicationSignInDetailedSummary struct {
	// Entity is the base model of ApplicationSignInDetailedSummary
	Entity
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// AppDisplayName undocumented
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// Status undocumented
	Status *SignInStatus `json:"status,omitempty"`
	// SignInCount undocumented
	SignInCount *int `json:"signInCount,omitempty"`
	// AggregatedEventDateTime undocumented
	AggregatedEventDateTime *time.Time `json:"aggregatedEventDateTime,omitempty"`
}
