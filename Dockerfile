FROM golang:1.22.1-bookworm as build
WORKDIR /go/src
COPY . .
RUN go build -o /go/bin/app ./cmd/server/main.go

FROM gcr.io/distroless/base-debian12
ENV TZ="Asia/Tokyo"
COPY --from=build /go/bin/app /
CMD ["/app"]
