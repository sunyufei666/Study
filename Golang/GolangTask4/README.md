## 运行环境

| 环境 | 版本 | 说明 |
|------|------|------|
| Go | 1.16+ | 编译和运行 |
| MySQL | 5.7+ | 数据存储 |

## 依赖管理

使用 Go Modules 进行依赖管理，初始化项目：
```bash
go mod init
go mod tidy
```

使用gin和gorm分别作为web框架和orm框架：
```bash
go get -u "gorm.io/driver/mysql"
go get -u "github.com/gin-gonic/gin"
```

使用JWT来创建和验证token：
```bash
go get -u "github.com/dgrijalva/jwt-go"
```

## 启动方式

服务启动：
```bash
go run main.go
```