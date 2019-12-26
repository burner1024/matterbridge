// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidForWorkScepCertificateProfile Android For Work SCEP certificate profile
type AndroidForWorkScepCertificateProfile struct {
	// AndroidForWorkCertificateProfileBase is the base model of AndroidForWorkScepCertificateProfile
	AndroidForWorkCertificateProfileBase
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
	// CertificateStore Target store certificate
	CertificateStore *CertificateStore `json:"certificateStore,omitempty"`
	// CustomSubjectAlternativeNames Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
	CustomSubjectAlternativeNames []CustomSubjectAlternativeName `json:"customSubjectAlternativeNames,omitempty"`
	// ManagedDeviceCertificateStates undocumented
	ManagedDeviceCertificateStates []ManagedDeviceCertificateState `json:"managedDeviceCertificateStates,omitempty"`
}
