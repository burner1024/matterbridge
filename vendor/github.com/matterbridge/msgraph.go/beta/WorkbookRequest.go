// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// WorkbookRequestBuilder is request builder for Workbook
type WorkbookRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookRequest
func (b *WorkbookRequestBuilder) Request() *WorkbookRequest {
	return &WorkbookRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookRequest is request for Workbook
type WorkbookRequest struct{ BaseRequest }

// Get performs GET request for Workbook
func (r *WorkbookRequest) Get(ctx context.Context) (resObj *Workbook, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Workbook
func (r *WorkbookRequest) Update(ctx context.Context, reqObj *Workbook) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Workbook
func (r *WorkbookRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Application is navigation property
func (b *WorkbookRequestBuilder) Application() *WorkbookApplicationRequestBuilder {
	bb := &WorkbookApplicationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/application"
	return bb
}

// Comments returns request builder for WorkbookComment collection
func (b *WorkbookRequestBuilder) Comments() *WorkbookCommentsCollectionRequestBuilder {
	bb := &WorkbookCommentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/comments"
	return bb
}

// WorkbookCommentsCollectionRequestBuilder is request builder for WorkbookComment collection
type WorkbookCommentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookComment collection
func (b *WorkbookCommentsCollectionRequestBuilder) Request() *WorkbookCommentsCollectionRequest {
	return &WorkbookCommentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookComment item
func (b *WorkbookCommentsCollectionRequestBuilder) ID(id string) *WorkbookCommentRequestBuilder {
	bb := &WorkbookCommentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookCommentsCollectionRequest is request for WorkbookComment collection
type WorkbookCommentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WorkbookComment collection
func (r *WorkbookCommentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]WorkbookComment, error) {
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
	var values []WorkbookComment
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
			value  []WorkbookComment
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

// Get performs GET request for WorkbookComment collection
func (r *WorkbookCommentsCollectionRequest) Get(ctx context.Context) ([]WorkbookComment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for WorkbookComment collection
func (r *WorkbookCommentsCollectionRequest) Add(ctx context.Context, reqObj *WorkbookComment) (resObj *WorkbookComment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Functions is navigation property
func (b *WorkbookRequestBuilder) Functions() *WorkbookFunctionsRequestBuilder {
	bb := &WorkbookFunctionsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/functions"
	return bb
}

// Names returns request builder for WorkbookNamedItem collection
func (b *WorkbookRequestBuilder) Names() *WorkbookNamesCollectionRequestBuilder {
	bb := &WorkbookNamesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/names"
	return bb
}

// WorkbookNamesCollectionRequestBuilder is request builder for WorkbookNamedItem collection
type WorkbookNamesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookNamedItem collection
func (b *WorkbookNamesCollectionRequestBuilder) Request() *WorkbookNamesCollectionRequest {
	return &WorkbookNamesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookNamedItem item
func (b *WorkbookNamesCollectionRequestBuilder) ID(id string) *WorkbookNamedItemRequestBuilder {
	bb := &WorkbookNamedItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookNamesCollectionRequest is request for WorkbookNamedItem collection
type WorkbookNamesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WorkbookNamedItem collection
func (r *WorkbookNamesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]WorkbookNamedItem, error) {
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
	var values []WorkbookNamedItem
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
			value  []WorkbookNamedItem
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

// Get performs GET request for WorkbookNamedItem collection
func (r *WorkbookNamesCollectionRequest) Get(ctx context.Context) ([]WorkbookNamedItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for WorkbookNamedItem collection
func (r *WorkbookNamesCollectionRequest) Add(ctx context.Context, reqObj *WorkbookNamedItem) (resObj *WorkbookNamedItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Tables returns request builder for WorkbookTable collection
func (b *WorkbookRequestBuilder) Tables() *WorkbookTablesCollectionRequestBuilder {
	bb := &WorkbookTablesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tables"
	return bb
}

// WorkbookTablesCollectionRequestBuilder is request builder for WorkbookTable collection
type WorkbookTablesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookTable collection
func (b *WorkbookTablesCollectionRequestBuilder) Request() *WorkbookTablesCollectionRequest {
	return &WorkbookTablesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookTable item
func (b *WorkbookTablesCollectionRequestBuilder) ID(id string) *WorkbookTableRequestBuilder {
	bb := &WorkbookTableRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookTablesCollectionRequest is request for WorkbookTable collection
type WorkbookTablesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WorkbookTable collection
func (r *WorkbookTablesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]WorkbookTable, error) {
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
	var values []WorkbookTable
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
			value  []WorkbookTable
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

// Get performs GET request for WorkbookTable collection
func (r *WorkbookTablesCollectionRequest) Get(ctx context.Context) ([]WorkbookTable, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for WorkbookTable collection
func (r *WorkbookTablesCollectionRequest) Add(ctx context.Context, reqObj *WorkbookTable) (resObj *WorkbookTable, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Worksheets returns request builder for WorkbookWorksheet collection
func (b *WorkbookRequestBuilder) Worksheets() *WorkbookWorksheetsCollectionRequestBuilder {
	bb := &WorkbookWorksheetsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/worksheets"
	return bb
}

// WorkbookWorksheetsCollectionRequestBuilder is request builder for WorkbookWorksheet collection
type WorkbookWorksheetsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookWorksheet collection
func (b *WorkbookWorksheetsCollectionRequestBuilder) Request() *WorkbookWorksheetsCollectionRequest {
	return &WorkbookWorksheetsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookWorksheet item
func (b *WorkbookWorksheetsCollectionRequestBuilder) ID(id string) *WorkbookWorksheetRequestBuilder {
	bb := &WorkbookWorksheetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookWorksheetsCollectionRequest is request for WorkbookWorksheet collection
type WorkbookWorksheetsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WorkbookWorksheet collection
func (r *WorkbookWorksheetsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]WorkbookWorksheet, error) {
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
	var values []WorkbookWorksheet
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
			value  []WorkbookWorksheet
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

// Get performs GET request for WorkbookWorksheet collection
func (r *WorkbookWorksheetsCollectionRequest) Get(ctx context.Context) ([]WorkbookWorksheet, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for WorkbookWorksheet collection
func (r *WorkbookWorksheetsCollectionRequest) Add(ctx context.Context, reqObj *WorkbookWorksheet) (resObj *WorkbookWorksheet, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
