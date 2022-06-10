# go
golang

## go-api
- golang:1.18.2

- Loacl Test
````
go run main.go
````
```
curl http://localhost:8080
# res: {"message":"Hello World!"}
```

- Create Docker image & container
```
docker build ./ -t go-api-image
docker images | grep go-api-image # Create OK
# go-sample-image latest
```
```
docker run --name go-api -d -p 80:8080 go-api-image
```
- Operation confirmation
```
curl http://localhost:8080/
# {"message":"Hello World!"}
```
