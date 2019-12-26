// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Place undocumented
type Place struct {
	// Entity is the base model of Place
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// GeoCoordinates undocumented
	GeoCoordinates *OutlookGeoCoordinates `json:"geoCoordinates,omitempty"`
	// Phone undocumented
	Phone *string `json:"phone,omitempty"`
	// Address undocumented
	Address *PhysicalAddress `json:"address,omitempty"`
}
