# heytom-go

基于 [Kratos](https://go-kratos.dev/) 框架构建的微服务项目。

## 项目结构

```
heytom-go/
├── api/                    # API 定义（protobuf）
│   └── heytom/v1/         # v1 版本 API
├── cmd/                    # 应用入口
│   └── server/            # 服务端主程序
├── configs/               # 配置文件
├── internal/              # 内部代码
│   ├── biz/              # 业务逻辑层
│   ├── conf/             # 配置结构定义
│   ├── data/             # 数据访问层
│   ├── server/           # 服务器配置（HTTP/gRPC）
│   └── service/          # 服务实现层
├── third_party/          # 第三方 proto 文件
├── go.mod
└── Makefile
```

## 快速开始

### 1. 安装依赖工具

```bash
make init
```

这将安装以下工具：
- protoc-gen-go
- protoc-gen-go-grpc
- protoc-gen-go-http
- protoc-gen-go-errors
- kratos CLI
- wire

### 2. 下载第三方 proto 文件

```bash
# 克隆 googleapis
git clone https://github.com/googleapis/googleapis.git third_party/googleapis

# 或者从 Kratos 仓库复制
git clone https://github.com/go-kratos/kratos.git /tmp/kratos
cp -r /tmp/kratos/third_party/* third_party/
```

### 3. 生成代码

```bash
# 生成 proto 代码
make api

# 生成 wire 依赖注入代码
make wire
```

### 4. 运行服务

```bash
# 下载 Go 依赖
go mod tidy

# 运行服务
make run
```

服务将在以下端口启动：
- HTTP: `http://localhost:8000`
- gRPC: `localhost:9000`

## 开发指南

### 添加新的 API

1. 在 `api/heytom/v1/` 目录下定义 proto 文件
2. 运行 `make api` 生成代码
3. 在 `internal/service/` 实现服务接口
4. 在 `internal/biz/` 实现业务逻辑
5. 在 `internal/data/` 实现数据访问

### 配置管理

配置文件位于 `configs/config.yaml`，可以通过环境变量或命令行参数指定配置路径：

```bash
./bin/server -conf ./configs/config.yaml
```

### 依赖注入

项目使用 [Wire](https://github.com/google/wire) 进行依赖注入。修改依赖关系后，运行：

```bash
make wire
```

## 构建

```bash
# 构建二进制文件
make build

# 指定版本号
VERSION=v1.0.0 make build
```

## 测试

```bash
make test
```

## 技术栈

- [Kratos](https://go-kratos.dev/) - 微服务框架
- [Wire](https://github.com/google/wire) - 依赖注入
- [Protobuf](https://developers.google.com/protocol-buffers) - API 定义
- [gRPC](https://grpc.io/) - RPC 框架

## License

MIT