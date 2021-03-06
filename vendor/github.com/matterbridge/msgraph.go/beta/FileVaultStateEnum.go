// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// FileVaultState undocumented
type FileVaultState int

const (
	// FileVaultStateVSuccess undocumented
	FileVaultStateVSuccess FileVaultState = 0
	// FileVaultStateVDriveEncryptedByUser undocumented
	FileVaultStateVDriveEncryptedByUser FileVaultState = 1
	// FileVaultStateVUserDeferredEncryption undocumented
	FileVaultStateVUserDeferredEncryption FileVaultState = 2
	// FileVaultStateVEscrowNotEnabled undocumented
	FileVaultStateVEscrowNotEnabled FileVaultState = 4
)

// FileVaultStatePSuccess returns a pointer to FileVaultStateVSuccess
func FileVaultStatePSuccess() *FileVaultState {
	v := FileVaultStateVSuccess
	return &v
}

// FileVaultStatePDriveEncryptedByUser returns a pointer to FileVaultStateVDriveEncryptedByUser
func FileVaultStatePDriveEncryptedByUser() *FileVaultState {
	v := FileVaultStateVDriveEncryptedByUser
	return &v
}

// FileVaultStatePUserDeferredEncryption returns a pointer to FileVaultStateVUserDeferredEncryption
func FileVaultStatePUserDeferredEncryption() *FileVaultState {
	v := FileVaultStateVUserDeferredEncryption
	return &v
}

// FileVaultStatePEscrowNotEnabled returns a pointer to FileVaultStateVEscrowNotEnabled
func FileVaultStatePEscrowNotEnabled() *FileVaultState {
	v := FileVaultStateVEscrowNotEnabled
	return &v
}
