
FROM golang:latest

#MAINTAINER 873421427@qq.com

WORKDIR $GOPATH/src/github.com/earnsparemoney/backend

COPY . .
RUN go get github.com/gpmgo/gopm
RUN gopm get -g 
RUN gopm install 


EXPOSE 443


CMD go run main.go

