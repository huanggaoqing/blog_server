# 依赖于golang sdk
FROM golang:alpine AS builder
# 设置容器里golang的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
# 移动到容器里的工作目录：/build
WORKDIR /build
# 将代码复制到容器的工作目录中
COPY . .
# 将我们的代码编译成二进制可执行文件 app
RUN go build -o app .

###################
# 接下来创建一个小镜像
###################
FROM scratch

# 创建容器工作目录
WORKDIR /blog_server

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/app /
COPY --from=builder /build/conf /

# 需要运行的命令
ENTRYPOINT ["./blog_server/app"]
