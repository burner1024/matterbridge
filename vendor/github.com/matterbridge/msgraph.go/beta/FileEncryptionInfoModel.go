// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// FileEncryptionInfo undocumented
type FileEncryptionInfo struct {
	// Object is the base model of FileEncryptionInfo
	Object
	// EncryptionKey The key used to encrypt the file content.
	EncryptionKey *Binary `json:"encryptionKey,omitempty"`
	// InitializationVector The initialization vector used for the encryption algorithm.
	InitializationVector *Binary `json:"initializationVector,omitempty"`
	// Mac The hash of the encrypted file content + IV (content hash).
	Mac *Binary `json:"mac,omitempty"`
	// MacKey The key used to get mac.
	MacKey *Binary `json:"macKey,omitempty"`
	// ProfileIdentifier The the profile identifier.
	ProfileIdentifier *string `json:"profileIdentifier,omitempty"`
	// FileDigest The file digest prior to encryption.
	FileDigest *Binary `json:"fileDigest,omitempty"`
	// FileDigestAlgorithm The file digest algorithm.
	FileDigestAlgorithm *string `json:"fileDigestAlgorithm,omitempty"`
}
