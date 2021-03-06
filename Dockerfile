FROM golang:latest

RUN mkdir -p /go/src/web-app
WORKDIR /go/src/web-app
# 復制 web-app 目錄到容器中
COPY main /go/src/web-app/main
COPY go-sql-driver /go/src/go-sql-driver
# 設定 PORT 環境變數
ENV PORT 8080
# 給主機暴露 8080 埠，這樣外部網路可以訪問你的應用
EXPOSE 8080

RUN go build -o /go/src/main /go/src/web-app/main/main.go

WORKDIR /go/src

CMD ["./main"]
