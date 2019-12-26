// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder is request builder for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequest
func (b *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) Request() *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequest {
	return &WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequest is request for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequest struct{ BaseRequest }

// Get performs GET request for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus
func (r *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequest) Get(ctx context.Context) (resObj *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus
func (r *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequest) Update(ctx context.Context, reqObj *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus
func (r *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Policy is navigation property
func (b *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) Policy() *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder {
	bb := &WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/policy"
	return bb
}
