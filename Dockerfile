FROM golang

MAINTAINER TomaChen

RUN go env -w GO111MODULE=on

RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /home/workspace

ADD . /home/workspace

# CMD go mod init berpar

# CMD go mod tidy

CMD go get -u github.com/gin-gonic/gin

RUN go build main.go

EXPOSE 8080

ENTRYPOINT ["./main"]