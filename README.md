# Quickstart
```
git clone https://github.com/Sarmirim/gofukurokuju/
cd gofukurokuju
go build .
docker-compose up
```

Check environment
(*system = windows/linux/...)
```
go tool dist list
go env GOOS GOARCH
go env -w GOOS=*system
```

Use in local terminal
```bash
go build .
gofokurokuju
```

Use in docker
```
docker build .
docker images
docker run -p 9876:9876 --name gofukurokuju -d id
docker run -p 9876:9876 --name gofukurokuju -it id
```
Docker-compose
```
docker-compose up 
```
with custom .env
```
docker-compose --env-file ./env  up 
```

```
docker logs gofukurokuju
docker start gofukurokuju 
```


docker cp gofukurokuju:/etc/ssl/certs/ca-certificates.crt ./etc/ssl/certs/