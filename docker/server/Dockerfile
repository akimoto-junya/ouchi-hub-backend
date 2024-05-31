FROM golang:1.22.3-bookworm as build
WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /go/bin/app ./cmd/server/main.go

FROM gcr.io/distroless/base-debian12
ENV TZ="Asia/Tokyo"
COPY --from=build /go/bin/app /
CMD ["/app"]
