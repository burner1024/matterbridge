// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceKey undocumented
type DeviceKey struct {
	// Object is the base model of DeviceKey
	Object
	// KeyType undocumented
	KeyType *string `json:"keyType,omitempty"`
	// KeyMaterial undocumented
	KeyMaterial *Binary `json:"keyMaterial,omitempty"`
	// DeviceID undocumented
	DeviceID *UUID `json:"deviceId,omitempty"`
}
