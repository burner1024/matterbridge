// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SetupStatus undocumented
type SetupStatus int

const (
	// SetupStatusVUnknown undocumented
	SetupStatusVUnknown SetupStatus = 0
	// SetupStatusVNotRegisteredYet undocumented
	SetupStatusVNotRegisteredYet SetupStatus = 1
	// SetupStatusVRegisteredSetupNotStarted undocumented
	SetupStatusVRegisteredSetupNotStarted SetupStatus = 2
	// SetupStatusVRegisteredSetupInProgress undocumented
	SetupStatusVRegisteredSetupInProgress SetupStatus = 3
	// SetupStatusVRegistrationAndSetupCompleted undocumented
	SetupStatusVRegistrationAndSetupCompleted SetupStatus = 4
	// SetupStatusVRegistrationFailed undocumented
	SetupStatusVRegistrationFailed SetupStatus = 5
	// SetupStatusVRegistrationTimedOut undocumented
	SetupStatusVRegistrationTimedOut SetupStatus = 6
	// SetupStatusVDisabled undocumented
	SetupStatusVDisabled SetupStatus = 7
)

// SetupStatusPUnknown returns a pointer to SetupStatusVUnknown
func SetupStatusPUnknown() *SetupStatus {
	v := SetupStatusVUnknown
	return &v
}

// SetupStatusPNotRegisteredYet returns a pointer to SetupStatusVNotRegisteredYet
func SetupStatusPNotRegisteredYet() *SetupStatus {
	v := SetupStatusVNotRegisteredYet
	return &v
}

// SetupStatusPRegisteredSetupNotStarted returns a pointer to SetupStatusVRegisteredSetupNotStarted
func SetupStatusPRegisteredSetupNotStarted() *SetupStatus {
	v := SetupStatusVRegisteredSetupNotStarted
	return &v
}

// SetupStatusPRegisteredSetupInProgress returns a pointer to SetupStatusVRegisteredSetupInProgress
func SetupStatusPRegisteredSetupInProgress() *SetupStatus {
	v := SetupStatusVRegisteredSetupInProgress
	return &v
}

// SetupStatusPRegistrationAndSetupCompleted returns a pointer to SetupStatusVRegistrationAndSetupCompleted
func SetupStatusPRegistrationAndSetupCompleted() *SetupStatus {
	v := SetupStatusVRegistrationAndSetupCompleted
	return &v
}

// SetupStatusPRegistrationFailed returns a pointer to SetupStatusVRegistrationFailed
func SetupStatusPRegistrationFailed() *SetupStatus {
	v := SetupStatusVRegistrationFailed
	return &v
}

// SetupStatusPRegistrationTimedOut returns a pointer to SetupStatusVRegistrationTimedOut
func SetupStatusPRegistrationTimedOut() *SetupStatus {
	v := SetupStatusVRegistrationTimedOut
	return &v
}

// SetupStatusPDisabled returns a pointer to SetupStatusVDisabled
func SetupStatusPDisabled() *SetupStatus {
	v := SetupStatusVDisabled
	return &v
}
