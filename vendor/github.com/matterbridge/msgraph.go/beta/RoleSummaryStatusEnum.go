// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RoleSummaryStatus undocumented
type RoleSummaryStatus int

const (
	// RoleSummaryStatusVOk undocumented
	RoleSummaryStatusVOk RoleSummaryStatus = 0
	// RoleSummaryStatusVBad undocumented
	RoleSummaryStatusVBad RoleSummaryStatus = 1
)

// RoleSummaryStatusPOk returns a pointer to RoleSummaryStatusVOk
func RoleSummaryStatusPOk() *RoleSummaryStatus {
	v := RoleSummaryStatusVOk
	return &v
}

// RoleSummaryStatusPBad returns a pointer to RoleSummaryStatusVBad
func RoleSummaryStatusPBad() *RoleSummaryStatus {
	v := RoleSummaryStatusVBad
	return &v
}
