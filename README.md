# go
golang

## go-api
- golang:1.18.2

#### Kernel
```
root@c4a2f69e1ec7:/home# uname -a
Linux c4a2f69e1ec7 5.10.104-linuxkit #1 SMP Thu Mar 17 17:08:06 UTC 2022 x86_64 GNU/Linux
```

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
