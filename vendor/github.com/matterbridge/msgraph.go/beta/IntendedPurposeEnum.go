// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IntendedPurpose undocumented
type IntendedPurpose int

const (
	// IntendedPurposeVUnassigned undocumented
	IntendedPurposeVUnassigned IntendedPurpose = 1
	// IntendedPurposeVSmimeEncryption undocumented
	IntendedPurposeVSmimeEncryption IntendedPurpose = 2
	// IntendedPurposeVSmimeSigning undocumented
	IntendedPurposeVSmimeSigning IntendedPurpose = 3
	// IntendedPurposeVVpn undocumented
	IntendedPurposeVVpn IntendedPurpose = 4
	// IntendedPurposeVWifi undocumented
	IntendedPurposeVWifi IntendedPurpose = 5
)

// IntendedPurposePUnassigned returns a pointer to IntendedPurposeVUnassigned
func IntendedPurposePUnassigned() *IntendedPurpose {
	v := IntendedPurposeVUnassigned
	return &v
}

// IntendedPurposePSmimeEncryption returns a pointer to IntendedPurposeVSmimeEncryption
func IntendedPurposePSmimeEncryption() *IntendedPurpose {
	v := IntendedPurposeVSmimeEncryption
	return &v
}

// IntendedPurposePSmimeSigning returns a pointer to IntendedPurposeVSmimeSigning
func IntendedPurposePSmimeSigning() *IntendedPurpose {
	v := IntendedPurposeVSmimeSigning
	return &v
}

// IntendedPurposePVpn returns a pointer to IntendedPurposeVVpn
func IntendedPurposePVpn() *IntendedPurpose {
	v := IntendedPurposeVVpn
	return &v
}

// IntendedPurposePWifi returns a pointer to IntendedPurposeVWifi
func IntendedPurposePWifi() *IntendedPurpose {
	v := IntendedPurposeVWifi
	return &v
}
