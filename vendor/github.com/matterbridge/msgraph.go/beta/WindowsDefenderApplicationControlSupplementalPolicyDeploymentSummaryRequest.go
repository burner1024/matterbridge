// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequestBuilder is request builder for WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequest
func (b *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequestBuilder) Request() *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequest {
	return &WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequest is request for WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequest struct{ BaseRequest }

// Get performs GET request for WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary
func (r *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequest) Get(ctx context.Context) (resObj *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary
func (r *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequest) Update(ctx context.Context, reqObj *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary
func (r *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}