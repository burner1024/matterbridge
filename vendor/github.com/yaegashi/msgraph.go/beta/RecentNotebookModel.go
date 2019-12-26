// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// RecentNotebook undocumented
type RecentNotebook struct {
	// Object is the base model of RecentNotebook
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastAccessedTime undocumented
	LastAccessedTime *time.Time `json:"lastAccessedTime,omitempty"`
	// Links undocumented
	Links *RecentNotebookLinks `json:"links,omitempty"`
	// SourceService undocumented
	SourceService *OnenoteSourceService `json:"sourceService,omitempty"`
}
