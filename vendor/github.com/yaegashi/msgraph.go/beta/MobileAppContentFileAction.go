// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MobileAppContentFileCommitRequestParameter undocumented
type MobileAppContentFileCommitRequestParameter struct {
	// FileEncryptionInfo undocumented
	FileEncryptionInfo *FileEncryptionInfo `json:"fileEncryptionInfo,omitempty"`
}

// MobileAppContentFileRenewUploadRequestParameter undocumented
type MobileAppContentFileRenewUploadRequestParameter struct {
}

//
type MobileAppContentFileCommitRequestBuilder struct{ BaseRequestBuilder }

// Commit action undocumented
func (b *MobileAppContentFileRequestBuilder) Commit(reqObj *MobileAppContentFileCommitRequestParameter) *MobileAppContentFileCommitRequestBuilder {
	bb := &MobileAppContentFileCommitRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/commit"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MobileAppContentFileCommitRequest struct{ BaseRequest }

//
func (b *MobileAppContentFileCommitRequestBuilder) Request() *MobileAppContentFileCommitRequest {
	return &MobileAppContentFileCommitRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MobileAppContentFileCommitRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type MobileAppContentFileRenewUploadRequestBuilder struct{ BaseRequestBuilder }

// RenewUpload action undocumented
func (b *MobileAppContentFileRequestBuilder) RenewUpload(reqObj *MobileAppContentFileRenewUploadRequestParameter) *MobileAppContentFileRenewUploadRequestBuilder {
	bb := &MobileAppContentFileRenewUploadRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/renewUpload"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MobileAppContentFileRenewUploadRequest struct{ BaseRequest }

//
func (b *MobileAppContentFileRenewUploadRequestBuilder) Request() *MobileAppContentFileRenewUploadRequest {
	return &MobileAppContentFileRenewUploadRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MobileAppContentFileRenewUploadRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
