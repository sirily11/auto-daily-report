# SME

### 项目描述

sme-server 是为中小型企业提供科技转型的后端项目。主要提供 Web3.0 钱包，AI 点单，线上支付等服务，帮助中小企完成科技升级。

### 项目结构

SME 项目是一个基于 Golang 的項目，采用 Gin 框架以提高 HTTP 服务的性能。它使用 Wire進行依賴注入。Wire 完全基于代码生成，在开发阶段，wire 会自动生成组件的初始化代码，生成代码人类可读，可以提交仓库，也可以正常编译。并通过 JWT 技术进行用户认证，确保了安全性和高效的身份验证。

## 项目启动

### 安装二方库依赖

创建一个 `GitHub Personal Access Token`，添加权限 `read:packages`, 然后将 `Token` 添加到环境变量中，名称为 `GITHUB_TOKEN`

```bash
export GITHUB_TOKEN=your_token
```

### 安装wire和mockgen

wire是一个依赖注入工具，用于生成依赖注入代码。mockgen是一个mock工具，用于生成mock代码。

```bash
go install github.com/google/wire/cmd/wire@latest && go install go.uber.org/mock/mockgen@latest
```


### 启动 go generate 自动生成 wire 代码

```bash
go generate ./...
```

### 启动项目

注意：需要环境变量 `GOEXPERIMENT=rangefunc`

build 项目

```bash
go build
```

如果沒有在terminal中添加環境變量`GOEXPERIMENT=rangefunc`，則需要build時附帶上該環境變量

```bash
GOEXPERIMENT=rangefunc go build
```

启动项目

```bash
go run main.go
```

如果沒有在terminal中添加環境變量`GOEXPERIMENT=rangefunc`，則需要run時附帶上該環境變量

```bash
GOEXPERIMENT=rangefunc go run main.go
```

### lint 代码规范

采用 golangci-lint 进行代码规范检查。其中，具体的 lint 配置在`.golangci.yml`文件中。

下载 golangci-lint

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2
```

```bash
golangci-lint run
```

### go format

```bash
go fmt ./...
```

## 项目测试

每一个功能模块都有对应的测试用例，可以通过以下命令进行测试。测试文件的命名规则为 `*_test.go`。

### 单元测试

对所有的 controller，service，repository 进行单元测试。

测试启动

```bash
go test -v ./...
```

### e2e 测试

尚待完善
