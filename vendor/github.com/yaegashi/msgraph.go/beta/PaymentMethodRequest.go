// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PaymentMethodRequestBuilder is request builder for PaymentMethod
type PaymentMethodRequestBuilder struct{ BaseRequestBuilder }

// Request returns PaymentMethodRequest
func (b *PaymentMethodRequestBuilder) Request() *PaymentMethodRequest {
	return &PaymentMethodRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PaymentMethodRequest is request for PaymentMethod
type PaymentMethodRequest struct{ BaseRequest }

// Get performs GET request for PaymentMethod
func (r *PaymentMethodRequest) Get(ctx context.Context) (resObj *PaymentMethod, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PaymentMethod
func (r *PaymentMethodRequest) Update(ctx context.Context, reqObj *PaymentMethod) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PaymentMethod
func (r *PaymentMethodRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
