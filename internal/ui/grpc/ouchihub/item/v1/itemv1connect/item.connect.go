// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ouchihub/item/v1/item.proto

package itemv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/akimoto-junya/ouchi-hub-backend/internal/ui/grpc/ouchihub/item/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	// ItemServiceName is the fully-qualified name of the ItemService service.
	ItemServiceName = "ouchihub.item.v1.ItemService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ItemServiceGetDirectoryProcedure is the fully-qualified name of the ItemService's GetDirectory
	// RPC.
	ItemServiceGetDirectoryProcedure = "/ouchihub.item.v1.ItemService/GetDirectory"
	// ItemServiceMoveItemProcedure is the fully-qualified name of the ItemService's MoveItem RPC.
	ItemServiceMoveItemProcedure = "/ouchihub.item.v1.ItemService/MoveItem"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	itemServiceServiceDescriptor            = v1.File_ouchihub_item_v1_item_proto.Services().ByName("ItemService")
	itemServiceGetDirectoryMethodDescriptor = itemServiceServiceDescriptor.Methods().ByName("GetDirectory")
	itemServiceMoveItemMethodDescriptor     = itemServiceServiceDescriptor.Methods().ByName("MoveItem")
)

// ItemServiceClient is a client for the ouchihub.item.v1.ItemService service.
type ItemServiceClient interface {
	GetDirectory(context.Context, *connect.Request[v1.GetDirectoryRequest]) (*connect.Response[v1.GetDirectoryResponse], error)
	MoveItem(context.Context, *connect.Request[v1.MoveItemRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewItemServiceClient constructs a client for the ouchihub.item.v1.ItemService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewItemServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ItemServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &itemServiceClient{
		getDirectory: connect.NewClient[v1.GetDirectoryRequest, v1.GetDirectoryResponse](
			httpClient,
			baseURL+ItemServiceGetDirectoryProcedure,
			connect.WithSchema(itemServiceGetDirectoryMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		moveItem: connect.NewClient[v1.MoveItemRequest, emptypb.Empty](
			httpClient,
			baseURL+ItemServiceMoveItemProcedure,
			connect.WithSchema(itemServiceMoveItemMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// itemServiceClient implements ItemServiceClient.
type itemServiceClient struct {
	getDirectory *connect.Client[v1.GetDirectoryRequest, v1.GetDirectoryResponse]
	moveItem     *connect.Client[v1.MoveItemRequest, emptypb.Empty]
}

// GetDirectory calls ouchihub.item.v1.ItemService.GetDirectory.
func (c *itemServiceClient) GetDirectory(ctx context.Context, req *connect.Request[v1.GetDirectoryRequest]) (*connect.Response[v1.GetDirectoryResponse], error) {
	return c.getDirectory.CallUnary(ctx, req)
}

// MoveItem calls ouchihub.item.v1.ItemService.MoveItem.
func (c *itemServiceClient) MoveItem(ctx context.Context, req *connect.Request[v1.MoveItemRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.moveItem.CallUnary(ctx, req)
}

// ItemServiceHandler is an implementation of the ouchihub.item.v1.ItemService service.
type ItemServiceHandler interface {
	GetDirectory(context.Context, *connect.Request[v1.GetDirectoryRequest]) (*connect.Response[v1.GetDirectoryResponse], error)
	MoveItem(context.Context, *connect.Request[v1.MoveItemRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewItemServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewItemServiceHandler(svc ItemServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	itemServiceGetDirectoryHandler := connect.NewUnaryHandler(
		ItemServiceGetDirectoryProcedure,
		svc.GetDirectory,
		connect.WithSchema(itemServiceGetDirectoryMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	itemServiceMoveItemHandler := connect.NewUnaryHandler(
		ItemServiceMoveItemProcedure,
		svc.MoveItem,
		connect.WithSchema(itemServiceMoveItemMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/ouchihub.item.v1.ItemService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ItemServiceGetDirectoryProcedure:
			itemServiceGetDirectoryHandler.ServeHTTP(w, r)
		case ItemServiceMoveItemProcedure:
			itemServiceMoveItemHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedItemServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedItemServiceHandler struct{}

func (UnimplementedItemServiceHandler) GetDirectory(context.Context, *connect.Request[v1.GetDirectoryRequest]) (*connect.Response[v1.GetDirectoryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ouchihub.item.v1.ItemService.GetDirectory is not implemented"))
}

func (UnimplementedItemServiceHandler) MoveItem(context.Context, *connect.Request[v1.MoveItemRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ouchihub.item.v1.ItemService.MoveItem is not implemented"))
}
