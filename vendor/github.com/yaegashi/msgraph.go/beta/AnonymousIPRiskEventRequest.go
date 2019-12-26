// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AnonymousIPRiskEventRequestBuilder is request builder for AnonymousIPRiskEvent
type AnonymousIPRiskEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns AnonymousIPRiskEventRequest
func (b *AnonymousIPRiskEventRequestBuilder) Request() *AnonymousIPRiskEventRequest {
	return &AnonymousIPRiskEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AnonymousIPRiskEventRequest is request for AnonymousIPRiskEvent
type AnonymousIPRiskEventRequest struct{ BaseRequest }

// Get performs GET request for AnonymousIPRiskEvent
func (r *AnonymousIPRiskEventRequest) Get(ctx context.Context) (resObj *AnonymousIPRiskEvent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AnonymousIPRiskEvent
func (r *AnonymousIPRiskEventRequest) Update(ctx context.Context, reqObj *AnonymousIPRiskEvent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AnonymousIPRiskEvent
func (r *AnonymousIPRiskEventRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
