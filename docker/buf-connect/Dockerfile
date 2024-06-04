FROM golang:1.22.4-alpine3.19 as prepare
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

FROM bufbuild/buf:1.32.2
COPY --from=prepare /go/bin/grpcurl /usr/local/bin/
COPY --from=prepare /go/bin/protoc-gen-go /usr/local/bin/
COPY --from=prepare /go/bin/protoc-gen-connect-go /usr/local/bin/
