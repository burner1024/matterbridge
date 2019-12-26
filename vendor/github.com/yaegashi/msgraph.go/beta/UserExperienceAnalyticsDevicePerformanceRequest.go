// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// UserExperienceAnalyticsDevicePerformanceRequestBuilder is request builder for UserExperienceAnalyticsDevicePerformance
type UserExperienceAnalyticsDevicePerformanceRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsDevicePerformanceRequest
func (b *UserExperienceAnalyticsDevicePerformanceRequestBuilder) Request() *UserExperienceAnalyticsDevicePerformanceRequest {
	return &UserExperienceAnalyticsDevicePerformanceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsDevicePerformanceRequest is request for UserExperienceAnalyticsDevicePerformance
type UserExperienceAnalyticsDevicePerformanceRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsDevicePerformance
func (r *UserExperienceAnalyticsDevicePerformanceRequest) Get(ctx context.Context) (resObj *UserExperienceAnalyticsDevicePerformance, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsDevicePerformance
func (r *UserExperienceAnalyticsDevicePerformanceRequest) Update(ctx context.Context, reqObj *UserExperienceAnalyticsDevicePerformance) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsDevicePerformance
func (r *UserExperienceAnalyticsDevicePerformanceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
