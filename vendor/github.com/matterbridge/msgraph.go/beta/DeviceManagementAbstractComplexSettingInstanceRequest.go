// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// DeviceManagementAbstractComplexSettingInstanceRequestBuilder is request builder for DeviceManagementAbstractComplexSettingInstance
type DeviceManagementAbstractComplexSettingInstanceRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementAbstractComplexSettingInstanceRequest
func (b *DeviceManagementAbstractComplexSettingInstanceRequestBuilder) Request() *DeviceManagementAbstractComplexSettingInstanceRequest {
	return &DeviceManagementAbstractComplexSettingInstanceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementAbstractComplexSettingInstanceRequest is request for DeviceManagementAbstractComplexSettingInstance
type DeviceManagementAbstractComplexSettingInstanceRequest struct{ BaseRequest }

// Get performs GET request for DeviceManagementAbstractComplexSettingInstance
func (r *DeviceManagementAbstractComplexSettingInstanceRequest) Get(ctx context.Context) (resObj *DeviceManagementAbstractComplexSettingInstance, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceManagementAbstractComplexSettingInstance
func (r *DeviceManagementAbstractComplexSettingInstanceRequest) Update(ctx context.Context, reqObj *DeviceManagementAbstractComplexSettingInstance) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceManagementAbstractComplexSettingInstance
func (r *DeviceManagementAbstractComplexSettingInstanceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Value returns request builder for DeviceManagementSettingInstance collection
func (b *DeviceManagementAbstractComplexSettingInstanceRequestBuilder) Value() *DeviceManagementAbstractComplexSettingInstanceValueCollectionRequestBuilder {
	bb := &DeviceManagementAbstractComplexSettingInstanceValueCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/value"
	return bb
}

// DeviceManagementAbstractComplexSettingInstanceValueCollectionRequestBuilder is request builder for DeviceManagementSettingInstance collection
type DeviceManagementAbstractComplexSettingInstanceValueCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementSettingInstance collection
func (b *DeviceManagementAbstractComplexSettingInstanceValueCollectionRequestBuilder) Request() *DeviceManagementAbstractComplexSettingInstanceValueCollectionRequest {
	return &DeviceManagementAbstractComplexSettingInstanceValueCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementSettingInstance item
func (b *DeviceManagementAbstractComplexSettingInstanceValueCollectionRequestBuilder) ID(id string) *DeviceManagementSettingInstanceRequestBuilder {
	bb := &DeviceManagementSettingInstanceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementAbstractComplexSettingInstanceValueCollectionRequest is request for DeviceManagementSettingInstance collection
type DeviceManagementAbstractComplexSettingInstanceValueCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceManagementSettingInstance collection
func (r *DeviceManagementAbstractComplexSettingInstanceValueCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceManagementSettingInstance, error) {
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
	var values []DeviceManagementSettingInstance
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
			value  []DeviceManagementSettingInstance
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

// Get performs GET request for DeviceManagementSettingInstance collection
func (r *DeviceManagementAbstractComplexSettingInstanceValueCollectionRequest) Get(ctx context.Context) ([]DeviceManagementSettingInstance, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceManagementSettingInstance collection
func (r *DeviceManagementAbstractComplexSettingInstanceValueCollectionRequest) Add(ctx context.Context, reqObj *DeviceManagementSettingInstance) (resObj *DeviceManagementSettingInstance, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
