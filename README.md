docker build .


environment
```
go tool dist list
go env GOOS GOARCH
go env -w GOOS=linux
```

docker
```
docker build .
docker images
docker run -p 9876:9876 --name gofukurokuju -d id
docker run -p 9876:9876 --name gofukurokuju -it id
docker logs gofukurokuju
docker start gofukurokuju 
```