// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// TaxArea undocumented
type TaxArea struct {
	// Entity is the base model of TaxArea
	Entity
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// TaxType undocumented
	TaxType *string `json:"taxType,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}
