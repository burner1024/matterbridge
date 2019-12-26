// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidCustomConfiguration This topic provides descriptions of the declared methods, properties and relationships exposed by the androidCustomConfiguration resource.
type AndroidCustomConfiguration struct {
	// DeviceConfiguration is the base model of AndroidCustomConfiguration
	DeviceConfiguration
	// OMASettings OMA settings. This collection can contain a maximum of 1000 elements.
	OMASettings []OMASetting `json:"omaSettings,omitempty"`
}
