FROM golang:latest

RUN mkdir -p /go/src/web-app 
WORKDIR /go/src/web-app 
# 復制 web-app 目錄到容器中
COPY . /go/src/web-app 
# 下載並安裝第三方依賴到容器中 
RUN go-wrapper download 
RUN go-wrapper install 
# 設定 PORT 環境變數 
ENV PORT 8080 
# 給主機暴露 8080 埠，這樣外部網路可以訪問你的應用 
EXPOSE 8080 
# 告訴 Docker 啟動容器運行的命令 
CMD ["go-wrapper", "run"]
