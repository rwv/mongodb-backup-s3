ARG GO_VERSION=1.18

FROM golang:${GO_VERSION} AS builder

WORKDIR $GOPATH/src/github.com/rwv/mongodb-backup-s3

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . $GOPATH/src/github.com/rwv/mongodb-backup-s3

RUN go build -o /mongodb-backup-s3

# Second Stage
FROM mongo

COPY --from=builder /mongodb-backup-s3 /mongodb-backup-s3

CMD ["/mongodb-backup-s3"]
