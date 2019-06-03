
FROM golang:latest

#MAINTAINER 873421427@qq.com

WORKDIR $GOPATH/src/github.com/earnsparemoney/backend

COPY . .
RUN go get -d -v ./...
RUN go install -v ./...


EXPOSE 443


CMD go run main.go