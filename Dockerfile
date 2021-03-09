FROM scratch
# Import the Certificate-Authority certificates for enabling HTTPS. (Certificates=(9/march/2021))
COPY ./etc/ssl/certs/ /etc/ssl/certs/
COPY gofukurokuju gofukurokuju
ENTRYPOINT [ "./gofukurokuju" ]

# # To use with newest certificates
# FROM golang:alpine AS builder
# RUN apk update
# RUN apk add -U --no-cache ca-certificates && update-ca-certificates

# FROM scratch
# # Import the Certificate-Authority certificates for enabling HTTPS.
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY gofukurokuju gofukurokuju
# ENTRYPOINT [ "./gofukurokuju" ]