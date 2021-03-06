// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PlannerTaskRequestBuilder is request builder for PlannerTask
type PlannerTaskRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlannerTaskRequest
func (b *PlannerTaskRequestBuilder) Request() *PlannerTaskRequest {
	return &PlannerTaskRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlannerTaskRequest is request for PlannerTask
type PlannerTaskRequest struct{ BaseRequest }

// Get performs GET request for PlannerTask
func (r *PlannerTaskRequest) Get(ctx context.Context) (resObj *PlannerTask, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PlannerTask
func (r *PlannerTaskRequest) Update(ctx context.Context, reqObj *PlannerTask) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PlannerTask
func (r *PlannerTaskRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AssignedToTaskBoardFormat is navigation property
func (b *PlannerTaskRequestBuilder) AssignedToTaskBoardFormat() *PlannerAssignedToTaskBoardTaskFormatRequestBuilder {
	bb := &PlannerAssignedToTaskBoardTaskFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignedToTaskBoardFormat"
	return bb
}

// BucketTaskBoardFormat is navigation property
func (b *PlannerTaskRequestBuilder) BucketTaskBoardFormat() *PlannerBucketTaskBoardTaskFormatRequestBuilder {
	bb := &PlannerBucketTaskBoardTaskFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/bucketTaskBoardFormat"
	return bb
}

// Details is navigation property
func (b *PlannerTaskRequestBuilder) Details() *PlannerTaskDetailsRequestBuilder {
	bb := &PlannerTaskDetailsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/details"
	return bb
}

// ProgressTaskBoardFormat is navigation property
func (b *PlannerTaskRequestBuilder) ProgressTaskBoardFormat() *PlannerProgressTaskBoardTaskFormatRequestBuilder {
	bb := &PlannerProgressTaskBoardTaskFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/progressTaskBoardFormat"
	return bb
}
