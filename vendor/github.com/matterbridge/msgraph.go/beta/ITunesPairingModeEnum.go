// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ITunesPairingMode undocumented
type ITunesPairingMode int

const (
	// ITunesPairingModeVDisallow undocumented
	ITunesPairingModeVDisallow ITunesPairingMode = 0
	// ITunesPairingModeVAllow undocumented
	ITunesPairingModeVAllow ITunesPairingMode = 1
	// ITunesPairingModeVRequiresCertificate undocumented
	ITunesPairingModeVRequiresCertificate ITunesPairingMode = 2
)

// ITunesPairingModePDisallow returns a pointer to ITunesPairingModeVDisallow
func ITunesPairingModePDisallow() *ITunesPairingMode {
	v := ITunesPairingModeVDisallow
	return &v
}

// ITunesPairingModePAllow returns a pointer to ITunesPairingModeVAllow
func ITunesPairingModePAllow() *ITunesPairingMode {
	v := ITunesPairingModeVAllow
	return &v
}

// ITunesPairingModePRequiresCertificate returns a pointer to ITunesPairingModeVRequiresCertificate
func ITunesPairingModePRequiresCertificate() *ITunesPairingMode {
	v := ITunesPairingModeVRequiresCertificate
	return &v
}
