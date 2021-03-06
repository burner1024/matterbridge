// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AppleDeviceFeaturesConfigurationBase Apple device features configuration profile.
type AppleDeviceFeaturesConfigurationBase struct {
	// DeviceConfiguration is the base model of AppleDeviceFeaturesConfigurationBase
	DeviceConfiguration
	// AirPrintDestinations An array of AirPrint printers that should always be shown. This collection can contain a maximum of 500 elements.
	AirPrintDestinations []AirPrintDestination `json:"airPrintDestinations,omitempty"`
}
