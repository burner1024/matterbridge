// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CalendarGroup undocumented
type CalendarGroup struct {
	// Entity is the base model of CalendarGroup
	Entity
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// ClassID undocumented
	ClassID *UUID `json:"classId,omitempty"`
	// ChangeKey undocumented
	ChangeKey *string `json:"changeKey,omitempty"`
	// Calendars undocumented
	Calendars []Calendar `json:"calendars,omitempty"`
}
