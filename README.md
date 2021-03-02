# Quickstart
```
git clone https://github.com/Sarmirim/gofukurokuju/
cd gofukurokuju
go build .
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
gofokurokuju
```

Use in docker
```
docker build .
docker images
docker run -p 9876:9876 --name gofukurokuju -d id
docker run -p 9876:9876 --name gofukurokuju -it id
```

```
docker logs gofukurokuju
docker start gofukurokuju 
```