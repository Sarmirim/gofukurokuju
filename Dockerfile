FROM golang:1.19 AS build
WORKDIR /src/
COPY main.go go.* /src/
COPY ./reddit /src/reddit
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /src/out/gofukurokuju .

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/out/gofukurokuju /
ENTRYPOINT [ "./gofukurokuju" ]
