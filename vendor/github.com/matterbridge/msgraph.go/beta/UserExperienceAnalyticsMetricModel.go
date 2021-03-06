// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserExperienceAnalyticsMetric The user experience analytics metric contains the score and units of a metric of a user experience anlaytics category.
type UserExperienceAnalyticsMetric struct {
	// Entity is the base model of UserExperienceAnalyticsMetric
	Entity
	// Value The value of the user experience analytics metric.
	Value *float64 `json:"value,omitempty"`
	// Unit The unit of the user experience analytics metric.
	Unit *string `json:"unit,omitempty"`
}
