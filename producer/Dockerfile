FROM golang:1.18.0-alpine3.15  as builder
COPY go.mod go.sum /go/src/github.com/ornellast/bookstore/
WORKDIR /go/src/github.com/ornellast/bookstore
RUN go mod download
COPY . /go/src/github.com/ornellast/bookstore/
RUN rm -rf ./.git
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/bookstore github.com/ornellast/bookstore

# RUN GOOS=linux go build -tags musl -a -installsuffix cgo -o build/bookstore github.com/ornellast/bookstore
# RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o build/bookstore github.com/ornellast/bookstore

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/ornellast/bookstore/build/bookstore /usr/bin/bookstore
EXPOSE 8080 8080
ENTRYPOINT [ "/usr/bin/bookstore" ]