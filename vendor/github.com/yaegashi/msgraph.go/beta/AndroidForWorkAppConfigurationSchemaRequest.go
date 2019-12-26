// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidForWorkAppConfigurationSchemaRequestBuilder is request builder for AndroidForWorkAppConfigurationSchema
type AndroidForWorkAppConfigurationSchemaRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidForWorkAppConfigurationSchemaRequest
func (b *AndroidForWorkAppConfigurationSchemaRequestBuilder) Request() *AndroidForWorkAppConfigurationSchemaRequest {
	return &AndroidForWorkAppConfigurationSchemaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidForWorkAppConfigurationSchemaRequest is request for AndroidForWorkAppConfigurationSchema
type AndroidForWorkAppConfigurationSchemaRequest struct{ BaseRequest }

// Get performs GET request for AndroidForWorkAppConfigurationSchema
func (r *AndroidForWorkAppConfigurationSchemaRequest) Get(ctx context.Context) (resObj *AndroidForWorkAppConfigurationSchema, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidForWorkAppConfigurationSchema
func (r *AndroidForWorkAppConfigurationSchemaRequest) Update(ctx context.Context, reqObj *AndroidForWorkAppConfigurationSchema) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidForWorkAppConfigurationSchema
func (r *AndroidForWorkAppConfigurationSchemaRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
