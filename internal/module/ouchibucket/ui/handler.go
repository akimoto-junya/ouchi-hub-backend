package ui

import (
	"net/http"

	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/bucket/v1/bucketv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/item/v1/itemv1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/storage/v1/storagev1connect"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/di"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/ui/handler"
)

func RegisterHandlers(s *http.ServeMux, d *di.Dependencies) {
	var p string
	var h http.Handler
	// storage
	p, h = storagev1connect.NewStorageServiceHandler(handler.NewStorageHandler(d.Storage))
	s.Handle(p, h)
	// bucket
	p, h = bucketv1connect.NewBucketServiceHandler(handler.NewBucketHandler(d.Bucket))
	s.Handle(p, h)
	// item
	p, h = itemv1connect.NewItemServiceHandler(handler.NewItemHandler(d.Item))
	s.Handle(p, h)
}
