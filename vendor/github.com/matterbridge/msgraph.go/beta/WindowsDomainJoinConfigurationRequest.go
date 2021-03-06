// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// WindowsDomainJoinConfigurationRequestBuilder is request builder for WindowsDomainJoinConfiguration
type WindowsDomainJoinConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsDomainJoinConfigurationRequest
func (b *WindowsDomainJoinConfigurationRequestBuilder) Request() *WindowsDomainJoinConfigurationRequest {
	return &WindowsDomainJoinConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsDomainJoinConfigurationRequest is request for WindowsDomainJoinConfiguration
type WindowsDomainJoinConfigurationRequest struct{ BaseRequest }

// Get performs GET request for WindowsDomainJoinConfiguration
func (r *WindowsDomainJoinConfigurationRequest) Get(ctx context.Context) (resObj *WindowsDomainJoinConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsDomainJoinConfiguration
func (r *WindowsDomainJoinConfigurationRequest) Update(ctx context.Context, reqObj *WindowsDomainJoinConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsDomainJoinConfiguration
func (r *WindowsDomainJoinConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// NetworkAccessConfigurations returns request builder for DeviceConfiguration collection
func (b *WindowsDomainJoinConfigurationRequestBuilder) NetworkAccessConfigurations() *WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequestBuilder {
	bb := &WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/networkAccessConfigurations"
	return bb
}

// WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequestBuilder is request builder for DeviceConfiguration collection
type WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceConfiguration collection
func (b *WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequestBuilder) Request() *WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequest {
	return &WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceConfiguration item
func (b *WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequestBuilder) ID(id string) *DeviceConfigurationRequestBuilder {
	bb := &DeviceConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequest is request for DeviceConfiguration collection
type WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceConfiguration collection
func (r *WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceConfiguration, error) {
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
	var values []DeviceConfiguration
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
			value  []DeviceConfiguration
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

// Get performs GET request for DeviceConfiguration collection
func (r *WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequest) Get(ctx context.Context) ([]DeviceConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceConfiguration collection
func (r *WindowsDomainJoinConfigurationNetworkAccessConfigurationsCollectionRequest) Add(ctx context.Context, reqObj *DeviceConfiguration) (resObj *DeviceConfiguration, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
