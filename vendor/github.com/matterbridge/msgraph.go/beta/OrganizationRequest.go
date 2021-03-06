// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// OrganizationRequestBuilder is request builder for Organization
type OrganizationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OrganizationRequest
func (b *OrganizationRequestBuilder) Request() *OrganizationRequest {
	return &OrganizationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OrganizationRequest is request for Organization
type OrganizationRequest struct{ BaseRequest }

// Get performs GET request for Organization
func (r *OrganizationRequest) Get(ctx context.Context) (resObj *Organization, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Organization
func (r *OrganizationRequest) Update(ctx context.Context, reqObj *Organization) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Organization
func (r *OrganizationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Brandings returns request builder for OrganizationalBranding collection
func (b *OrganizationRequestBuilder) Brandings() *OrganizationBrandingsCollectionRequestBuilder {
	bb := &OrganizationBrandingsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/brandings"
	return bb
}

// OrganizationBrandingsCollectionRequestBuilder is request builder for OrganizationalBranding collection
type OrganizationBrandingsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OrganizationalBranding collection
func (b *OrganizationBrandingsCollectionRequestBuilder) Request() *OrganizationBrandingsCollectionRequest {
	return &OrganizationBrandingsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OrganizationalBranding item
func (b *OrganizationBrandingsCollectionRequestBuilder) ID(id string) *OrganizationalBrandingRequestBuilder {
	bb := &OrganizationalBrandingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OrganizationBrandingsCollectionRequest is request for OrganizationalBranding collection
type OrganizationBrandingsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OrganizationalBranding collection
func (r *OrganizationBrandingsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OrganizationalBranding, error) {
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
	var values []OrganizationalBranding
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
			value  []OrganizationalBranding
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

// Get performs GET request for OrganizationalBranding collection
func (r *OrganizationBrandingsCollectionRequest) Get(ctx context.Context) ([]OrganizationalBranding, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OrganizationalBranding collection
func (r *OrganizationBrandingsCollectionRequest) Add(ctx context.Context, reqObj *OrganizationalBranding) (resObj *OrganizationalBranding, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// CertificateBasedAuthConfiguration returns request builder for CertificateBasedAuthConfiguration collection
func (b *OrganizationRequestBuilder) CertificateBasedAuthConfiguration() *OrganizationCertificateBasedAuthConfigurationCollectionRequestBuilder {
	bb := &OrganizationCertificateBasedAuthConfigurationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/certificateBasedAuthConfiguration"
	return bb
}

// OrganizationCertificateBasedAuthConfigurationCollectionRequestBuilder is request builder for CertificateBasedAuthConfiguration collection
type OrganizationCertificateBasedAuthConfigurationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CertificateBasedAuthConfiguration collection
func (b *OrganizationCertificateBasedAuthConfigurationCollectionRequestBuilder) Request() *OrganizationCertificateBasedAuthConfigurationCollectionRequest {
	return &OrganizationCertificateBasedAuthConfigurationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CertificateBasedAuthConfiguration item
func (b *OrganizationCertificateBasedAuthConfigurationCollectionRequestBuilder) ID(id string) *CertificateBasedAuthConfigurationRequestBuilder {
	bb := &CertificateBasedAuthConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OrganizationCertificateBasedAuthConfigurationCollectionRequest is request for CertificateBasedAuthConfiguration collection
type OrganizationCertificateBasedAuthConfigurationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CertificateBasedAuthConfiguration collection
func (r *OrganizationCertificateBasedAuthConfigurationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]CertificateBasedAuthConfiguration, error) {
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
	var values []CertificateBasedAuthConfiguration
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
			value  []CertificateBasedAuthConfiguration
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

// Get performs GET request for CertificateBasedAuthConfiguration collection
func (r *OrganizationCertificateBasedAuthConfigurationCollectionRequest) Get(ctx context.Context) ([]CertificateBasedAuthConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for CertificateBasedAuthConfiguration collection
func (r *OrganizationCertificateBasedAuthConfigurationCollectionRequest) Add(ctx context.Context, reqObj *CertificateBasedAuthConfiguration) (resObj *CertificateBasedAuthConfiguration, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Extensions returns request builder for Extension collection
func (b *OrganizationRequestBuilder) Extensions() *OrganizationExtensionsCollectionRequestBuilder {
	bb := &OrganizationExtensionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensions"
	return bb
}

// OrganizationExtensionsCollectionRequestBuilder is request builder for Extension collection
type OrganizationExtensionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Extension collection
func (b *OrganizationExtensionsCollectionRequestBuilder) Request() *OrganizationExtensionsCollectionRequest {
	return &OrganizationExtensionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Extension item
func (b *OrganizationExtensionsCollectionRequestBuilder) ID(id string) *ExtensionRequestBuilder {
	bb := &ExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OrganizationExtensionsCollectionRequest is request for Extension collection
type OrganizationExtensionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Extension collection
func (r *OrganizationExtensionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Extension, error) {
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
	var values []Extension
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
			value  []Extension
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

// Get performs GET request for Extension collection
func (r *OrganizationExtensionsCollectionRequest) Get(ctx context.Context) ([]Extension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Extension collection
func (r *OrganizationExtensionsCollectionRequest) Add(ctx context.Context, reqObj *Extension) (resObj *Extension, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
