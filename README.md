# HeyTom

基于 [Kratos](https://go-kratos.dev/) 框架构建的 Go 微服务项目。

## 项目简介

HeyTom 是一个使用 Kratos v2 框架开发的微服务应用，提供 HTTP 和 gRPC 双协议支持，包含完整的问候服务示例。

## 技术栈

- **框架**: Kratos v2.8.0
- **语言**: Go 1.21+
- **依赖注入**: Google Wire
- **协议**: HTTP + gRPC
- **配置**: YAML
- **API 定义**: Protocol Buffers

## 项目结构

```
.
├── api/                    # API 定义（protobuf）
│   └── helloworld/v1/     # 问候服务 API
├── cmd/heytom/            # 应用入口
│   ├── main.go           # 主程序
│   ├── wire.go           # Wire 依赖注入配置
│   └── wire_gen.go       # Wire 生成代码
├── configs/               # 配置文件
│   └── config.yaml       # 应用配置
├── internal/              # 内部代码
│   ├── biz/              # 业务逻辑层
│   ├── data/             # 数据访问层
│   ├── server/           # 服务器配置（HTTP/gRPC）
│   ├── service/          # 服务实现层
│   └── conf/             # 配置结构定义
├── third_party/           # 第三方 proto 文件
├── Dockerfile            # Docker 构建文件
├── Makefile              # 构建脚本
└── openapi.yaml          # OpenAPI 规范
```

## 快速开始

### 环境要求

- Go 1.21 或更高版本
- Protocol Buffers 编译器（protoc）
- Make 工具

### 安装依赖工具

```bash
make init
```

此命令会安装以下工具：
- protoc-gen-go
- protoc-gen-go-grpc
- protoc-gen-go-http
- protoc-gen-openapi
- kratos CLI
- wire

### 生成代码

```bash
# 生成所有代码（API + 配置 + Wire）
make all

# 或分别生成
make api        # 生成 API 代码
make config     # 生成配置代码
make generate   # 生成 Wire 依赖注入代码
```

### 编译项目

```bash
make build
```

编译后的二进制文件位于 `./bin/` 目录。

### 运行服务

```bash
# 使用默认配置路径
./bin/heytom -conf ./configs

# 或指定配置文件
./bin/heytom -conf /path/to/config.yaml
```

## 配置说明

配置文件位于 `configs/config.yaml`：

```yaml
server:
  http:
    addr: 0.0.0.0:8000      # HTTP 服务地址
    timeout: 1s
  grpc:
    addr: 0.0.0.0:9000      # gRPC 服务地址
    timeout: 1s
data:
  database:
    driver: mysql
    source: root:root@tcp(127.0.0.1:3306)/test
  redis:
    addr: 127.0.0.1:6379
    read_timeout: 0.2s
    write_timeout: 0.2s
```

## API 接口

### HTTP API

**问候接口**
```
GET /helloworld/{name}
```

示例：
```bash
curl http://localhost:8000/helloworld/Tom
```

响应：
```json
{
  "message": "Hello Tom"
}
```

### gRPC API

使用 gRPC 客户端连接到 `localhost:9000` 调用 `Greeter.SayHello` 方法。

完整的 API 文档参见 `openapi.yaml`。

## Docker 部署

### 构建镜像

```bash
docker build -t heytom:latest .
```

### 运行容器

```bash
docker run -d \
  -p 8000:8000 \
  -p 9000:9000 \
  -v /path/to/config:/data/conf \
  heytom:latest
```

## 开发指南

### 添加新的 API

1. 在 `api/` 目录下定义 protobuf 文件
2. 运行 `make api` 生成代码
3. 在 `internal/service/` 实现服务逻辑
4. 在 `internal/biz/` 实现业务逻辑
5. 在 `internal/data/` 实现数据访问

### 架构分层

- **Service 层**: 处理 API 请求，参数校验
- **Biz 层**: 业务逻辑处理
- **Data 层**: 数据访问，外部服务调用

### 依赖注入

项目使用 Wire 进行依赖注入，修改依赖关系后需要：

```bash
make generate
```

## 常用命令

```bash
make help       # 查看所有可用命令
make init       # 初始化开发环境
make api        # 生成 API 代码
make config     # 生成配置代码
make generate   # 生成 Wire 代码
make build      # 编译项目
make all        # 生成所有代码
```

## 许可证

详见 [LICENSE](LICENSE) 文件。

## 相关链接

- [Kratos 官方文档](https://go-kratos.dev/)
- [Protocol Buffers](https://protobuf.dev/)
- [Google Wire](https://github.com/google/wire)
