// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// PendingContentUpdate undocumented
type PendingContentUpdate struct {
	// Object is the base model of PendingContentUpdate
	Object
	// QueuedDateTime undocumented
	QueuedDateTime *time.Time `json:"queuedDateTime,omitempty"`
}
