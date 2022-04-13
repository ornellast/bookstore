FROM golang:1.18.0-alpine3.15  as builder
COPY go.mod go.sum /go/src/github.com/ornellast/bucketeer/
WORKDIR /go/src/github.com/ornellast/bucketeer
RUN go mod download
COPY . /go/src/github.com/ornellast/bucketeer/
RUN rm -rf ./.git
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/bucketeer github.com/ornellast/bucketeer

# RUN GOOS=linux go build -tags musl -a -installsuffix cgo -o build/bucketeer github.com/ornellast/bucketeer
# RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o build/bucketeer github.com/ornellast/bucketeer

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/ornellast/bucketeer/build/bucketeer /usr/bin/bucketeer
EXPOSE 8080 8080
ENTRYPOINT [ "/usr/bin/bucketeer" ]