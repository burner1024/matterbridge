// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PolicySetStatus undocumented
type PolicySetStatus int

const (
	// PolicySetStatusVUnknown undocumented
	PolicySetStatusVUnknown PolicySetStatus = 0
	// PolicySetStatusVValidating undocumented
	PolicySetStatusVValidating PolicySetStatus = 1
	// PolicySetStatusVPartialSuccess undocumented
	PolicySetStatusVPartialSuccess PolicySetStatus = 2
	// PolicySetStatusVSuccess undocumented
	PolicySetStatusVSuccess PolicySetStatus = 3
	// PolicySetStatusVError undocumented
	PolicySetStatusVError PolicySetStatus = 4
	// PolicySetStatusVNotAssigned undocumented
	PolicySetStatusVNotAssigned PolicySetStatus = 5
)

// PolicySetStatusPUnknown returns a pointer to PolicySetStatusVUnknown
func PolicySetStatusPUnknown() *PolicySetStatus {
	v := PolicySetStatusVUnknown
	return &v
}

// PolicySetStatusPValidating returns a pointer to PolicySetStatusVValidating
func PolicySetStatusPValidating() *PolicySetStatus {
	v := PolicySetStatusVValidating
	return &v
}

// PolicySetStatusPPartialSuccess returns a pointer to PolicySetStatusVPartialSuccess
func PolicySetStatusPPartialSuccess() *PolicySetStatus {
	v := PolicySetStatusVPartialSuccess
	return &v
}

// PolicySetStatusPSuccess returns a pointer to PolicySetStatusVSuccess
func PolicySetStatusPSuccess() *PolicySetStatus {
	v := PolicySetStatusVSuccess
	return &v
}

// PolicySetStatusPError returns a pointer to PolicySetStatusVError
func PolicySetStatusPError() *PolicySetStatus {
	v := PolicySetStatusVError
	return &v
}

// PolicySetStatusPNotAssigned returns a pointer to PolicySetStatusVNotAssigned
func PolicySetStatusPNotAssigned() *PolicySetStatus {
	v := PolicySetStatusVNotAssigned
	return &v
}
