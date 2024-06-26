// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ouchibucket/bucket/v1/bucket.proto

package bucketv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/bucket/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// BucketServiceName is the fully-qualified name of the BucketService service.
	BucketServiceName = "ouchibucket.bucket.v1.BucketService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BucketServiceCreateBucketProcedure is the fully-qualified name of the BucketService's
	// CreateBucket RPC.
	BucketServiceCreateBucketProcedure = "/ouchibucket.bucket.v1.BucketService/CreateBucket"
	// BucketServiceListBucketsProcedure is the fully-qualified name of the BucketService's ListBuckets
	// RPC.
	BucketServiceListBucketsProcedure = "/ouchibucket.bucket.v1.BucketService/ListBuckets"
	// BucketServiceSyncBucketProcedure is the fully-qualified name of the BucketService's SyncBucket
	// RPC.
	BucketServiceSyncBucketProcedure = "/ouchibucket.bucket.v1.BucketService/SyncBucket"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	bucketServiceServiceDescriptor            = v1.File_ouchibucket_bucket_v1_bucket_proto.Services().ByName("BucketService")
	bucketServiceCreateBucketMethodDescriptor = bucketServiceServiceDescriptor.Methods().ByName("CreateBucket")
	bucketServiceListBucketsMethodDescriptor  = bucketServiceServiceDescriptor.Methods().ByName("ListBuckets")
	bucketServiceSyncBucketMethodDescriptor   = bucketServiceServiceDescriptor.Methods().ByName("SyncBucket")
)

// BucketServiceClient is a client for the ouchibucket.bucket.v1.BucketService service.
type BucketServiceClient interface {
	CreateBucket(context.Context, *connect.Request[v1.CreateBucketRequest]) (*connect.Response[v1.CreateBucketResponse], error)
	ListBuckets(context.Context, *connect.Request[v1.ListBucketsRequest]) (*connect.Response[v1.ListBucketsResponse], error)
	SyncBucket(context.Context, *connect.Request[v1.SyncBucketRequest]) (*connect.Response[v1.SyncBucketResponse], error)
}

// NewBucketServiceClient constructs a client for the ouchibucket.bucket.v1.BucketService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBucketServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BucketServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &bucketServiceClient{
		createBucket: connect.NewClient[v1.CreateBucketRequest, v1.CreateBucketResponse](
			httpClient,
			baseURL+BucketServiceCreateBucketProcedure,
			connect.WithSchema(bucketServiceCreateBucketMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listBuckets: connect.NewClient[v1.ListBucketsRequest, v1.ListBucketsResponse](
			httpClient,
			baseURL+BucketServiceListBucketsProcedure,
			connect.WithSchema(bucketServiceListBucketsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		syncBucket: connect.NewClient[v1.SyncBucketRequest, v1.SyncBucketResponse](
			httpClient,
			baseURL+BucketServiceSyncBucketProcedure,
			connect.WithSchema(bucketServiceSyncBucketMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// bucketServiceClient implements BucketServiceClient.
type bucketServiceClient struct {
	createBucket *connect.Client[v1.CreateBucketRequest, v1.CreateBucketResponse]
	listBuckets  *connect.Client[v1.ListBucketsRequest, v1.ListBucketsResponse]
	syncBucket   *connect.Client[v1.SyncBucketRequest, v1.SyncBucketResponse]
}

// CreateBucket calls ouchibucket.bucket.v1.BucketService.CreateBucket.
func (c *bucketServiceClient) CreateBucket(ctx context.Context, req *connect.Request[v1.CreateBucketRequest]) (*connect.Response[v1.CreateBucketResponse], error) {
	return c.createBucket.CallUnary(ctx, req)
}

// ListBuckets calls ouchibucket.bucket.v1.BucketService.ListBuckets.
func (c *bucketServiceClient) ListBuckets(ctx context.Context, req *connect.Request[v1.ListBucketsRequest]) (*connect.Response[v1.ListBucketsResponse], error) {
	return c.listBuckets.CallUnary(ctx, req)
}

// SyncBucket calls ouchibucket.bucket.v1.BucketService.SyncBucket.
func (c *bucketServiceClient) SyncBucket(ctx context.Context, req *connect.Request[v1.SyncBucketRequest]) (*connect.Response[v1.SyncBucketResponse], error) {
	return c.syncBucket.CallUnary(ctx, req)
}

// BucketServiceHandler is an implementation of the ouchibucket.bucket.v1.BucketService service.
type BucketServiceHandler interface {
	CreateBucket(context.Context, *connect.Request[v1.CreateBucketRequest]) (*connect.Response[v1.CreateBucketResponse], error)
	ListBuckets(context.Context, *connect.Request[v1.ListBucketsRequest]) (*connect.Response[v1.ListBucketsResponse], error)
	SyncBucket(context.Context, *connect.Request[v1.SyncBucketRequest]) (*connect.Response[v1.SyncBucketResponse], error)
}

// NewBucketServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBucketServiceHandler(svc BucketServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	bucketServiceCreateBucketHandler := connect.NewUnaryHandler(
		BucketServiceCreateBucketProcedure,
		svc.CreateBucket,
		connect.WithSchema(bucketServiceCreateBucketMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	bucketServiceListBucketsHandler := connect.NewUnaryHandler(
		BucketServiceListBucketsProcedure,
		svc.ListBuckets,
		connect.WithSchema(bucketServiceListBucketsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	bucketServiceSyncBucketHandler := connect.NewUnaryHandler(
		BucketServiceSyncBucketProcedure,
		svc.SyncBucket,
		connect.WithSchema(bucketServiceSyncBucketMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/ouchibucket.bucket.v1.BucketService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BucketServiceCreateBucketProcedure:
			bucketServiceCreateBucketHandler.ServeHTTP(w, r)
		case BucketServiceListBucketsProcedure:
			bucketServiceListBucketsHandler.ServeHTTP(w, r)
		case BucketServiceSyncBucketProcedure:
			bucketServiceSyncBucketHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBucketServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBucketServiceHandler struct{}

func (UnimplementedBucketServiceHandler) CreateBucket(context.Context, *connect.Request[v1.CreateBucketRequest]) (*connect.Response[v1.CreateBucketResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ouchibucket.bucket.v1.BucketService.CreateBucket is not implemented"))
}

func (UnimplementedBucketServiceHandler) ListBuckets(context.Context, *connect.Request[v1.ListBucketsRequest]) (*connect.Response[v1.ListBucketsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ouchibucket.bucket.v1.BucketService.ListBuckets is not implemented"))
}

func (UnimplementedBucketServiceHandler) SyncBucket(context.Context, *connect.Request[v1.SyncBucketRequest]) (*connect.Response[v1.SyncBucketResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ouchibucket.bucket.v1.BucketService.SyncBucket is not implemented"))
}
