# 使用Alpine镜像作为基础镜像
FROM alpine:latest

# 创建应用工作目录和配置目录
RUN mkdir -p /app && mkdir -p /app/other/etc

# 将本地编译好的二进制文件复制到容器中
COPY ./other-api /app/other-api

# 设置工作目录
WORKDIR /app

# 设置容器启动时执行的命令
CMD ["./other-api"]

# 暴露容器服务端口
EXPOSE 8090