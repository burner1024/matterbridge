// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TeamsAsyncOperationType undocumented
type TeamsAsyncOperationType int

const (
	// TeamsAsyncOperationTypeVInvalid undocumented
	TeamsAsyncOperationTypeVInvalid TeamsAsyncOperationType = 0
	// TeamsAsyncOperationTypeVCloneTeam undocumented
	TeamsAsyncOperationTypeVCloneTeam TeamsAsyncOperationType = 1
	// TeamsAsyncOperationTypeVArchiveTeam undocumented
	TeamsAsyncOperationTypeVArchiveTeam TeamsAsyncOperationType = 2
	// TeamsAsyncOperationTypeVUnarchiveTeam undocumented
	TeamsAsyncOperationTypeVUnarchiveTeam TeamsAsyncOperationType = 3
	// TeamsAsyncOperationTypeVCreateTeam undocumented
	TeamsAsyncOperationTypeVCreateTeam TeamsAsyncOperationType = 4
	// TeamsAsyncOperationTypeVUnknownFutureValue undocumented
	TeamsAsyncOperationTypeVUnknownFutureValue TeamsAsyncOperationType = 5
)

// TeamsAsyncOperationTypePInvalid returns a pointer to TeamsAsyncOperationTypeVInvalid
func TeamsAsyncOperationTypePInvalid() *TeamsAsyncOperationType {
	v := TeamsAsyncOperationTypeVInvalid
	return &v
}

// TeamsAsyncOperationTypePCloneTeam returns a pointer to TeamsAsyncOperationTypeVCloneTeam
func TeamsAsyncOperationTypePCloneTeam() *TeamsAsyncOperationType {
	v := TeamsAsyncOperationTypeVCloneTeam
	return &v
}

// TeamsAsyncOperationTypePArchiveTeam returns a pointer to TeamsAsyncOperationTypeVArchiveTeam
func TeamsAsyncOperationTypePArchiveTeam() *TeamsAsyncOperationType {
	v := TeamsAsyncOperationTypeVArchiveTeam
	return &v
}

// TeamsAsyncOperationTypePUnarchiveTeam returns a pointer to TeamsAsyncOperationTypeVUnarchiveTeam
func TeamsAsyncOperationTypePUnarchiveTeam() *TeamsAsyncOperationType {
	v := TeamsAsyncOperationTypeVUnarchiveTeam
	return &v
}

// TeamsAsyncOperationTypePCreateTeam returns a pointer to TeamsAsyncOperationTypeVCreateTeam
func TeamsAsyncOperationTypePCreateTeam() *TeamsAsyncOperationType {
	v := TeamsAsyncOperationTypeVCreateTeam
	return &v
}

// TeamsAsyncOperationTypePUnknownFutureValue returns a pointer to TeamsAsyncOperationTypeVUnknownFutureValue
func TeamsAsyncOperationTypePUnknownFutureValue() *TeamsAsyncOperationType {
	v := TeamsAsyncOperationTypeVUnknownFutureValue
	return &v
}
