
FROM golang:latest

#MAINTAINER 873421427@qq.com

WORKDIR $GOPATH/src/github.com/earnsparemoney/backend

COPY . .
RUN go get -u github.com/gpmgo/gopm
RUN gopm get -d -v ./...
RUN gopm install -v ./...



EXPOSE 443


CMD go run main.go

