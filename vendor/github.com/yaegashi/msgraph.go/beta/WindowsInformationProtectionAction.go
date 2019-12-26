// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WindowsInformationProtectionAssignRequestParameter undocumented
type WindowsInformationProtectionAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

//
type WindowsInformationProtectionAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *WindowsInformationProtectionRequestBuilder) Assign(reqObj *WindowsInformationProtectionAssignRequestParameter) *WindowsInformationProtectionAssignRequestBuilder {
	bb := &WindowsInformationProtectionAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WindowsInformationProtectionAssignRequest struct{ BaseRequest }

//
func (b *WindowsInformationProtectionAssignRequestBuilder) Request() *WindowsInformationProtectionAssignRequest {
	return &WindowsInformationProtectionAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WindowsInformationProtectionAssignRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
