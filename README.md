# earn spare money backend
A go implementation of backend service


## how to run in docker 
Due to the fucking GFW, I try many times to solve the golang packages dependency in my own US server and finnaly build the docker successfully.
1. first you need to register an account in docker hub 
2. sign in docker hub 
```
docker login --username=yourusername --password=yourpassword
```
3. pull the docker image from docker hub
```
docker pull zzt1234/godocker:version1.0
```
4. run the image 
```
docker run -d -p 443:443 zzt1234/godocker:version1.0
```

## how to run in your own computer 
1. make sure you have install go and configure your go path correctly
2. download gopm, which is a package manager
```
go get -u github.com/gpmgo/gopm
```
3. get packages needed in .gopmfile using gopm command 
```
gopm get -g
```
4. run
```
go run main.go
```
