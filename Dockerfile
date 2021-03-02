# Multistage for certificates
# And build (but container immediately exit if copy build from "builder")
# AS builder
FROM golang:alpine AS builder
RUN apk update
RUN apk add -U --no-cache ca-certificates && update-ca-certificates
# WORKDIR /go/pkg/mod/github.com/Sarmirim/gofukurokuju
# COPY . .
# RUN go install .
# RUN go install github.com/Sarmirim/gofukurokuju@latest

FROM scratch
# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /go/bin/gofukurokuju ./gofukurokuju
COPY gofukurokuju gofukurokuju

ENTRYPOINT [ "./gofukurokuju" ]