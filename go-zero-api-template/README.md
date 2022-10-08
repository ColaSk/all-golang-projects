# go-zero 单体http服务进阶

## TREE

```text
.admin
    └── api
        ├── api.api
        ├── api.go
        ├── etc
        │   └── api-api.yaml // 服务配置文件
        └── internal
            ├── config // 配置文件
            │   └── config.go
            ├── handler // route与view
            │   ├── apihandler.go
            │   └── routes.go
            ├── logic // 服务处理程序
            │   └── apilogic.go
            ├── svc // 服务上下文
            │   └── servicecontext.go
            └── types
                └── types.go
```
