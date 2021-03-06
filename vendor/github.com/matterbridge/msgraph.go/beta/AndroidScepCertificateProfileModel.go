// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidScepCertificateProfile Android SCEP certificate profile
type AndroidScepCertificateProfile struct {
	// AndroidCertificateProfileBase is the base model of AndroidScepCertificateProfile
	AndroidCertificateProfileBase
	// ScepServerUrls SCEP Server Url(s)
	ScepServerUrls []string `json:"scepServerUrls,omitempty"`
	// SubjectNameFormatString Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
	SubjectNameFormatString *string `json:"subjectNameFormatString,omitempty"`
	// KeyUsage SCEP Key Usage
	KeyUsage *KeyUsages `json:"keyUsage,omitempty"`
	// KeySize SCEP Key Size
	KeySize *KeySize `json:"keySize,omitempty"`
	// HashAlgorithm SCEP Hash Algorithm
	HashAlgorithm *HashAlgorithms `json:"hashAlgorithm,omitempty"`
	// SubjectAlternativeNameFormatString Custom String that defines the AAD Attribute.
	SubjectAlternativeNameFormatString *string `json:"subjectAlternativeNameFormatString,omitempty"`
	// ManagedDeviceCertificateStates undocumented
	ManagedDeviceCertificateStates []ManagedDeviceCertificateState `json:"managedDeviceCertificateStates,omitempty"`
}
