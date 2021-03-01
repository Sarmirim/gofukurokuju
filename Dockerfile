# FROM scratch
# COPY gofukurokuju gofukurokuju
# EXPOSE 9876
# ENTRYPOINT [ "gofukurokuju" ]

# AS builder
FROM golang:alpine AS builder
RUN apk update
RUN apk add -U --no-cache ca-certificates && update-ca-certificates
# WORKDIR /go/pkg/mod/github.com/Sarmirim/gofukurokuju
# COPY . .
# RUN go install .

# RUN go install github.com/Sarmirim/gofukurokuju@develop

# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
# WORKDIR /gofukurokuju
# WORKDIR $GOPATH/src/myapp/
# COPY . $GOPATH/src/myapp/
# RUN go install && go build -o go/bin/fuku
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
#     -ldflags="-w -s" -o $PROJ_BIN_PATH ./cmd/app
# RUN go get ./gofukurokuju && go install && go build .
# COPY . .

# COPY . gofukurokuju

FROM scratch
# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /go/src/myapp/cmd/app gofukurokuju
COPY gofukurokuju gofukurokuju

ENTRYPOINT [ "./gofukurokuju" ]