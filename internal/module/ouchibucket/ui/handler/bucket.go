package handler

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/google/uuid"

	bucketv1 "github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/bucket/v1"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/bucket/v1/bucketv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/usecase"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/usecase/dto"
)

type bucketHandler struct {
	bucket usecase.Bucket
}

func NewBucketHandler(b usecase.Bucket) bucketv1connect.BucketServiceHandler {
	return &bucketHandler{
		bucket: b,
	}
}

func (h *bucketHandler) CreateBucket(ctx context.Context, req *connect.Request[bucketv1.CreateBucketRequest]) (*connect.Response[bucketv1.CreateBucketResponse], error) {
	storageID, err := uuid.Parse(req.Msg.GetStorageId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	id, err := h.bucket.CreateBucket(ctx, &dto.CreateBucketParams{
		Name:      req.Msg.Name,
		StorageID: storageID,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&bucketv1.CreateBucketResponse{Id: id.String()}), nil
}

func (h *bucketHandler) ListBuckets(ctx context.Context, req *connect.Request[bucketv1.ListBucketsRequest]) (*connect.Response[bucketv1.ListBucketsResponse], error) {
	storageID, err := uuid.Parse(req.Msg.GetStorageId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	buckets, err := h.bucket.ListBuckets(ctx, storageID)
	if err != nil {
		// TODO:
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&bucketv1.ListBucketsResponse{Buckets: bucketv1.ToMessages(buckets)}), nil
}

func (h *bucketHandler) SyncBucket(ctx context.Context, req *connect.Request[bucketv1.SyncBucketRequest]) (*connect.Response[bucketv1.SyncBucketResponse], error) {
	bucketID, err := uuid.Parse(req.Msg.GetBucketId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	if err := h.bucket.SyncBucket(ctx, bucketID, req.Msg.GetRelativePath()); err != nil {
		if errors.Is(err, usecase.ErrBucketNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&bucketv1.SyncBucketResponse{}), nil
}
