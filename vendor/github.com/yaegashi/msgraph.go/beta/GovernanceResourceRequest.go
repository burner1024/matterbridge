// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// GovernanceResourceRequestBuilder is request builder for GovernanceResource
type GovernanceResourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns GovernanceResourceRequest
func (b *GovernanceResourceRequestBuilder) Request() *GovernanceResourceRequest {
	return &GovernanceResourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GovernanceResourceRequest is request for GovernanceResource
type GovernanceResourceRequest struct{ BaseRequest }

// Get performs GET request for GovernanceResource
func (r *GovernanceResourceRequest) Get(ctx context.Context) (resObj *GovernanceResource, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GovernanceResource
func (r *GovernanceResourceRequest) Update(ctx context.Context, reqObj *GovernanceResource) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GovernanceResource
func (r *GovernanceResourceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Parent is navigation property
func (b *GovernanceResourceRequestBuilder) Parent() *GovernanceResourceRequestBuilder {
	bb := &GovernanceResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parent"
	return bb
}

// RoleAssignmentRequests returns request builder for GovernanceRoleAssignmentRequestObject collection
func (b *GovernanceResourceRequestBuilder) RoleAssignmentRequests() *GovernanceResourceRoleAssignmentRequestsCollectionRequestBuilder {
	bb := &GovernanceResourceRoleAssignmentRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleAssignmentRequests"
	return bb
}

// GovernanceResourceRoleAssignmentRequestsCollectionRequestBuilder is request builder for GovernanceRoleAssignmentRequestObject collection
type GovernanceResourceRoleAssignmentRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleAssignmentRequestObject collection
func (b *GovernanceResourceRoleAssignmentRequestsCollectionRequestBuilder) Request() *GovernanceResourceRoleAssignmentRequestsCollectionRequest {
	return &GovernanceResourceRoleAssignmentRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleAssignmentRequestObject item
func (b *GovernanceResourceRoleAssignmentRequestsCollectionRequestBuilder) ID(id string) *GovernanceRoleAssignmentRequestObjectRequestBuilder {
	bb := &GovernanceRoleAssignmentRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GovernanceResourceRoleAssignmentRequestsCollectionRequest is request for GovernanceRoleAssignmentRequestObject collection
type GovernanceResourceRoleAssignmentRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleAssignmentRequestObject collection
func (r *GovernanceResourceRoleAssignmentRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]GovernanceRoleAssignmentRequestObject, error) {
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
	var values []GovernanceRoleAssignmentRequestObject
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
			value  []GovernanceRoleAssignmentRequestObject
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

// Get performs GET request for GovernanceRoleAssignmentRequestObject collection
func (r *GovernanceResourceRoleAssignmentRequestsCollectionRequest) Get(ctx context.Context) ([]GovernanceRoleAssignmentRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for GovernanceRoleAssignmentRequestObject collection
func (r *GovernanceResourceRoleAssignmentRequestsCollectionRequest) Add(ctx context.Context, reqObj *GovernanceRoleAssignmentRequestObject) (resObj *GovernanceRoleAssignmentRequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RoleAssignments returns request builder for GovernanceRoleAssignment collection
func (b *GovernanceResourceRequestBuilder) RoleAssignments() *GovernanceResourceRoleAssignmentsCollectionRequestBuilder {
	bb := &GovernanceResourceRoleAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleAssignments"
	return bb
}

// GovernanceResourceRoleAssignmentsCollectionRequestBuilder is request builder for GovernanceRoleAssignment collection
type GovernanceResourceRoleAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleAssignment collection
func (b *GovernanceResourceRoleAssignmentsCollectionRequestBuilder) Request() *GovernanceResourceRoleAssignmentsCollectionRequest {
	return &GovernanceResourceRoleAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleAssignment item
func (b *GovernanceResourceRoleAssignmentsCollectionRequestBuilder) ID(id string) *GovernanceRoleAssignmentRequestBuilder {
	bb := &GovernanceRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GovernanceResourceRoleAssignmentsCollectionRequest is request for GovernanceRoleAssignment collection
type GovernanceResourceRoleAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleAssignment collection
func (r *GovernanceResourceRoleAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]GovernanceRoleAssignment, error) {
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
	var values []GovernanceRoleAssignment
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
			value  []GovernanceRoleAssignment
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

// Get performs GET request for GovernanceRoleAssignment collection
func (r *GovernanceResourceRoleAssignmentsCollectionRequest) Get(ctx context.Context) ([]GovernanceRoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for GovernanceRoleAssignment collection
func (r *GovernanceResourceRoleAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *GovernanceRoleAssignment) (resObj *GovernanceRoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RoleDefinitions returns request builder for GovernanceRoleDefinition collection
func (b *GovernanceResourceRequestBuilder) RoleDefinitions() *GovernanceResourceRoleDefinitionsCollectionRequestBuilder {
	bb := &GovernanceResourceRoleDefinitionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinitions"
	return bb
}

// GovernanceResourceRoleDefinitionsCollectionRequestBuilder is request builder for GovernanceRoleDefinition collection
type GovernanceResourceRoleDefinitionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleDefinition collection
func (b *GovernanceResourceRoleDefinitionsCollectionRequestBuilder) Request() *GovernanceResourceRoleDefinitionsCollectionRequest {
	return &GovernanceResourceRoleDefinitionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleDefinition item
func (b *GovernanceResourceRoleDefinitionsCollectionRequestBuilder) ID(id string) *GovernanceRoleDefinitionRequestBuilder {
	bb := &GovernanceRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GovernanceResourceRoleDefinitionsCollectionRequest is request for GovernanceRoleDefinition collection
type GovernanceResourceRoleDefinitionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleDefinition collection
func (r *GovernanceResourceRoleDefinitionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]GovernanceRoleDefinition, error) {
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
	var values []GovernanceRoleDefinition
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
			value  []GovernanceRoleDefinition
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

// Get performs GET request for GovernanceRoleDefinition collection
func (r *GovernanceResourceRoleDefinitionsCollectionRequest) Get(ctx context.Context) ([]GovernanceRoleDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for GovernanceRoleDefinition collection
func (r *GovernanceResourceRoleDefinitionsCollectionRequest) Add(ctx context.Context, reqObj *GovernanceRoleDefinition) (resObj *GovernanceRoleDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RoleSettings returns request builder for GovernanceRoleSetting collection
func (b *GovernanceResourceRequestBuilder) RoleSettings() *GovernanceResourceRoleSettingsCollectionRequestBuilder {
	bb := &GovernanceResourceRoleSettingsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleSettings"
	return bb
}

// GovernanceResourceRoleSettingsCollectionRequestBuilder is request builder for GovernanceRoleSetting collection
type GovernanceResourceRoleSettingsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleSetting collection
func (b *GovernanceResourceRoleSettingsCollectionRequestBuilder) Request() *GovernanceResourceRoleSettingsCollectionRequest {
	return &GovernanceResourceRoleSettingsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleSetting item
func (b *GovernanceResourceRoleSettingsCollectionRequestBuilder) ID(id string) *GovernanceRoleSettingRequestBuilder {
	bb := &GovernanceRoleSettingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GovernanceResourceRoleSettingsCollectionRequest is request for GovernanceRoleSetting collection
type GovernanceResourceRoleSettingsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleSetting collection
func (r *GovernanceResourceRoleSettingsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]GovernanceRoleSetting, error) {
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
	var values []GovernanceRoleSetting
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
			value  []GovernanceRoleSetting
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

// Get performs GET request for GovernanceRoleSetting collection
func (r *GovernanceResourceRoleSettingsCollectionRequest) Get(ctx context.Context) ([]GovernanceRoleSetting, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for GovernanceRoleSetting collection
func (r *GovernanceResourceRoleSettingsCollectionRequest) Add(ctx context.Context, reqObj *GovernanceRoleSetting) (resObj *GovernanceRoleSetting, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
