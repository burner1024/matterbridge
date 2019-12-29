// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// GroupPolicyDefinitionRequestBuilder is request builder for GroupPolicyDefinition
type GroupPolicyDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupPolicyDefinitionRequest
func (b *GroupPolicyDefinitionRequestBuilder) Request() *GroupPolicyDefinitionRequest {
	return &GroupPolicyDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupPolicyDefinitionRequest is request for GroupPolicyDefinition
type GroupPolicyDefinitionRequest struct{ BaseRequest }

// Get performs GET request for GroupPolicyDefinition
func (r *GroupPolicyDefinitionRequest) Get(ctx context.Context) (resObj *GroupPolicyDefinition, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GroupPolicyDefinition
func (r *GroupPolicyDefinitionRequest) Update(ctx context.Context, reqObj *GroupPolicyDefinition) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GroupPolicyDefinition
func (r *GroupPolicyDefinitionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DefinitionFile is navigation property
func (b *GroupPolicyDefinitionRequestBuilder) DefinitionFile() *GroupPolicyDefinitionFileRequestBuilder {
	bb := &GroupPolicyDefinitionFileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/definitionFile"
	return bb
}

// Presentations returns request builder for GroupPolicyPresentation collection
func (b *GroupPolicyDefinitionRequestBuilder) Presentations() *GroupPolicyDefinitionPresentationsCollectionRequestBuilder {
	bb := &GroupPolicyDefinitionPresentationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/presentations"
	return bb
}

// GroupPolicyDefinitionPresentationsCollectionRequestBuilder is request builder for GroupPolicyPresentation collection
type GroupPolicyDefinitionPresentationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GroupPolicyPresentation collection
func (b *GroupPolicyDefinitionPresentationsCollectionRequestBuilder) Request() *GroupPolicyDefinitionPresentationsCollectionRequest {
	return &GroupPolicyDefinitionPresentationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GroupPolicyPresentation item
func (b *GroupPolicyDefinitionPresentationsCollectionRequestBuilder) ID(id string) *GroupPolicyPresentationRequestBuilder {
	bb := &GroupPolicyPresentationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GroupPolicyDefinitionPresentationsCollectionRequest is request for GroupPolicyPresentation collection
type GroupPolicyDefinitionPresentationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GroupPolicyPresentation collection
func (r *GroupPolicyDefinitionPresentationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]GroupPolicyPresentation, error) {
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
	var values []GroupPolicyPresentation
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
			value  []GroupPolicyPresentation
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

// Get performs GET request for GroupPolicyPresentation collection
func (r *GroupPolicyDefinitionPresentationsCollectionRequest) Get(ctx context.Context) ([]GroupPolicyPresentation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for GroupPolicyPresentation collection
func (r *GroupPolicyDefinitionPresentationsCollectionRequest) Add(ctx context.Context, reqObj *GroupPolicyPresentation) (resObj *GroupPolicyPresentation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}