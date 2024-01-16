FROM golang:1.21.6-bookworm as build
WORKDIR /go/src
COPY . .
RUN go build -o /go/bin/app main.go

FROM gcr.io/distroless/base-debian-12
ENV TZ="Asia/Tokyo"
COPY --from=build /go/bin/app /
CMD ["/app"]
