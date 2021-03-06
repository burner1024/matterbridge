// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// ExternalConnectionRequestBuilder is request builder for ExternalConnection
type ExternalConnectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExternalConnectionRequest
func (b *ExternalConnectionRequestBuilder) Request() *ExternalConnectionRequest {
	return &ExternalConnectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExternalConnectionRequest is request for ExternalConnection
type ExternalConnectionRequest struct{ BaseRequest }

// Get performs GET request for ExternalConnection
func (r *ExternalConnectionRequest) Get(ctx context.Context) (resObj *ExternalConnection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExternalConnection
func (r *ExternalConnectionRequest) Update(ctx context.Context, reqObj *ExternalConnection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExternalConnection
func (r *ExternalConnectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Items returns request builder for ExternalItem collection
func (b *ExternalConnectionRequestBuilder) Items() *ExternalConnectionItemsCollectionRequestBuilder {
	bb := &ExternalConnectionItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/items"
	return bb
}

// ExternalConnectionItemsCollectionRequestBuilder is request builder for ExternalItem collection
type ExternalConnectionItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ExternalItem collection
func (b *ExternalConnectionItemsCollectionRequestBuilder) Request() *ExternalConnectionItemsCollectionRequest {
	return &ExternalConnectionItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ExternalItem item
func (b *ExternalConnectionItemsCollectionRequestBuilder) ID(id string) *ExternalItemRequestBuilder {
	bb := &ExternalItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ExternalConnectionItemsCollectionRequest is request for ExternalItem collection
type ExternalConnectionItemsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ExternalItem collection
func (r *ExternalConnectionItemsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ExternalItem, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ExternalItem
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ExternalItem
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ExternalItem collection
func (r *ExternalConnectionItemsCollectionRequest) Get(ctx context.Context) ([]ExternalItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ExternalItem collection
func (r *ExternalConnectionItemsCollectionRequest) Add(ctx context.Context, reqObj *ExternalItem) (resObj *ExternalItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Operations returns request builder for ConnectionOperation collection
func (b *ExternalConnectionRequestBuilder) Operations() *ExternalConnectionOperationsCollectionRequestBuilder {
	bb := &ExternalConnectionOperationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/operations"
	return bb
}

// ExternalConnectionOperationsCollectionRequestBuilder is request builder for ConnectionOperation collection
type ExternalConnectionOperationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ConnectionOperation collection
func (b *ExternalConnectionOperationsCollectionRequestBuilder) Request() *ExternalConnectionOperationsCollectionRequest {
	return &ExternalConnectionOperationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ConnectionOperation item
func (b *ExternalConnectionOperationsCollectionRequestBuilder) ID(id string) *ConnectionOperationRequestBuilder {
	bb := &ConnectionOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ExternalConnectionOperationsCollectionRequest is request for ConnectionOperation collection
type ExternalConnectionOperationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ConnectionOperation collection
func (r *ExternalConnectionOperationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ConnectionOperation, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ConnectionOperation
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ConnectionOperation
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ConnectionOperation collection
func (r *ExternalConnectionOperationsCollectionRequest) Get(ctx context.Context) ([]ConnectionOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ConnectionOperation collection
func (r *ExternalConnectionOperationsCollectionRequest) Add(ctx context.Context, reqObj *ConnectionOperation) (resObj *ConnectionOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Schema is navigation property
func (b *ExternalConnectionRequestBuilder) Schema() *SchemaRequestBuilder {
	bb := &SchemaRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/schema"
	return bb
}
