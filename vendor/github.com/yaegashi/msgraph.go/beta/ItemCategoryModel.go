// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ItemCategory undocumented
type ItemCategory struct {
	// Entity is the base model of ItemCategory
	Entity
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}
