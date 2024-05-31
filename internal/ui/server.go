package ui

import (
	"net/http"

	"connectrpc.com/grpcreflect"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/bucket/v1/bucketv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/item/v1/itemv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/storage/v1/storagev1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchihub/category/v1/categoryv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchihub/content/v1/contentv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchihub/maker/v1/makerv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchihub/work/v1/workv1connect"
)

func RegisterServerReflection(s *http.ServeMux) {
	// server reflection
	reflecter := grpcreflect.NewStaticReflector(
		categoryv1connect.CategoryServiceName,
		makerv1connect.MakerServiceName,
		workv1connect.WorkServiceName,

		// ouchibucket
		storagev1connect.StorageServiceName,
		bucketv1connect.BucketServiceName,
		itemv1connect.ItemServiceName,
		contentv1connect.ContentServiceName,
	)
	s.Handle(grpcreflect.NewHandlerV1(reflecter))
	s.Handle(grpcreflect.NewHandlerV1Alpha(reflecter))
}
