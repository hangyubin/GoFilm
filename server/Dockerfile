# 使用golang:1.18-alpine作为基础镜像构建后端
FROM golang:1.18-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制所有文件
COPY . .

# 构建后端项目
RUN go build -o gofilm main.go

# 使用alpine:latest作为运行时镜像
FROM alpine:latest

# 安装tzdata包，用于处理时区
RUN apk add --no-cache tzdata

# 设置工作目录
WORKDIR /app

# 复制构建好的二进制文件
COPY --from=builder /app/gofilm .

# 配置文件通过环境变量加载，无需复制配置文件

# 设置时区
ENV TZ=Asia/Shanghai

# 暴露3600端口
EXPOSE 3600

# 启动后端服务
CMD ["./gofilm"]
