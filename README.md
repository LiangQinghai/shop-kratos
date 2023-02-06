# shop-kratos

> shop application with go-kratos

# 开发步骤

- 环境准备

  1. kratos-cli
  
  ```shell
  go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
  ```
  
  2. protoc(https://github.com/protocolbuffers/protobuf/releases)

- 新增一个业务域（如果需要的话）

```shell
kratos new app/user --nomod -r https://gitee.com/go-kratos/kratos-layout.git   
```

- 创建一个`.proto`文件

```shell
kratos proto add api/user/v1/user.proto 
```

- 编辑`.proto`文件, 定义`rpc/http`接口

```protobuf
syntax = "proto3";

package api.user.v1;

import "google/api/annotations.proto";

option go_package = "shop-kratos/api/user/v1;v1";
option java_multiple_files = true;
option java_package = "api.user.v1";

service User {
	rpc RegisterUser(RegisterRequest) returns (RegisterReply) {
		option (google.api.http) = {
			post: "/api/user/v1/register",
			body: "*"
		};
	}
}

message RegisterRequest {
	string username = 1;
	string password = 2;
}
message RegisterReply {
	int64 userId = 1;
}
```

- 生成`client`端代码

```shell
kratos proto client api\user\v1\user.proto
```

- 生成service端代码（非覆盖性，新增`.proto`接口后请手动新增service端方法）

```shell
kratos proto server api\user\v1\user.proto -t app\user\internal\service
```

- `service` `wire`注入编辑，确保`service/service.go`中`ProviderSet`添加了新增的`service`业务类

- 编写业务代码`biz/user.go`，定义`UserRepo`接口和`UserUseCase`业务类，并且确保在`biz/biz.go`中注入`UserCase`业务类实例

- 编写`repository`代码`data/user.go`, 实现`biz`中的`UserRepo`接口，调用数据库/其他业务存储库，确保在`data/data.go`中注入`UserRepo`实例

- 注册 `http/rpc` 服务

`http`类服务`server/http.go`:
```go
// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, UserService *service.UserService, logger log.Logger) *http.Server {
    // ... 省略
	v1.RegisterUserHTTPServer(srv, UserService)
	// ... 省略
}
```

`rpc`类服务`server/grpc.go`:
```go
// NewHTTPServer new an HTTP server.
func NewGRPCServer(c *conf.Server, UserService *service.UserService, logger log.Logger) *http.Server {
    // ... 省略
	v1.RegisterUserServer(srv, UserService)
	// ... 省略
}
```

- 生成`google-wire`代码

```shell
go generate ./...
```

- 运行

```shell
kratos run
```