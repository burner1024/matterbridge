// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ImportedDeviceIdentityCollectionImportDeviceIdentityListRequestParameter undocumented
type ImportedDeviceIdentityCollectionImportDeviceIdentityListRequestParameter struct {
	// ImportedDeviceIdentities undocumented
	ImportedDeviceIdentities []ImportedDeviceIdentity `json:"importedDeviceIdentities,omitempty"`
	// OverwriteImportedDeviceIdentities undocumented
	OverwriteImportedDeviceIdentities *bool `json:"overwriteImportedDeviceIdentities,omitempty"`
}

// ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequestParameter undocumented
type ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequestParameter struct {
	// ImportedDeviceIdentities undocumented
	ImportedDeviceIdentities []ImportedDeviceIdentity `json:"importedDeviceIdentities,omitempty"`
}

//
type ImportedDeviceIdentityCollectionImportDeviceIdentityListRequestBuilder struct{ BaseRequestBuilder }

// ImportDeviceIdentityList action undocumented
func (b *DeviceManagementImportedDeviceIdentitiesCollectionRequestBuilder) ImportDeviceIdentityList(reqObj *ImportedDeviceIdentityCollectionImportDeviceIdentityListRequestParameter) *ImportedDeviceIdentityCollectionImportDeviceIdentityListRequestBuilder {
	bb := &ImportedDeviceIdentityCollectionImportDeviceIdentityListRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/importDeviceIdentityList"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ImportedDeviceIdentityCollectionImportDeviceIdentityListRequest struct{ BaseRequest }

//
func (b *ImportedDeviceIdentityCollectionImportDeviceIdentityListRequestBuilder) Request() *ImportedDeviceIdentityCollectionImportDeviceIdentityListRequest {
	return &ImportedDeviceIdentityCollectionImportDeviceIdentityListRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ImportedDeviceIdentityCollectionImportDeviceIdentityListRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]ImportedDeviceIdentityResult, error) {
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
	var values [][]ImportedDeviceIdentityResult
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
			value  [][]ImportedDeviceIdentityResult
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

//
func (r *ImportedDeviceIdentityCollectionImportDeviceIdentityListRequest) Get(ctx context.Context) ([][]ImportedDeviceIdentityResult, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

//
type ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequestBuilder struct{ BaseRequestBuilder }

// SearchExistingIdentities action undocumented
func (b *DeviceManagementImportedDeviceIdentitiesCollectionRequestBuilder) SearchExistingIdentities(reqObj *ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequestParameter) *ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequestBuilder {
	bb := &ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/searchExistingIdentities"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequest struct{ BaseRequest }

//
func (b *ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequestBuilder) Request() *ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequest {
	return &ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]ImportedDeviceIdentity, error) {
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
	var values [][]ImportedDeviceIdentity
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
			value  [][]ImportedDeviceIdentity
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

//
func (r *ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequest) Get(ctx context.Context) ([][]ImportedDeviceIdentity, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}
