# golang-api-app

To running app without docker: 
```
go run .
```

To build app :
```
go build .
```

To build docker :
```
docker build -t golang-api-app .
```


TO run docker :
```
docker run -it --net="host" golang-api-app
```

config:
--net="host" is used if you want connect to local db

to fix permission denied when build:
```
sudo groupadd docker
```
```
sudo usermod -aG docker ${USER}
```
```
su ${USER}
```

```
sudo systemctl restart docker
```