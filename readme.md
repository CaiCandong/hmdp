# 基于黑马点评的Golang实现
## 前端服务启动
```shell
    cd ./views/nginx-1.23.4
    nginx.exe
```

``` 


.
├── cmd
│   ├── web           # Gin框架相关代码
│   │   ├── main.go   # 应用程序入口
│   │   └── router.go # 路由配置
│   └── cli           # 命令行工具相关代码
│       ├── main.go   # 命令行工具入口
│       └── ...
├── internal
│   ├── app           # 应用层
│   │   ├── cart      # 购物车相关服务
│   │   ├── order     # 订单相关服务
│   │   └── user      # 用户相关服务
│   ├── domain        # 领域层
│   │   ├── cart      # 购物车领域模型和服务接口
│   │   ├── order     # 订单领域模型和服务接口
│   │   ├── product   # 商品领域模型和服务接口
│   │   └── user      # 用户领域模型和服务接口
│   ├── infrastructure # 基础设施层
│   │   ├── database  # 数据库相关代码
│   │   ├── event     # 领域事件相关代码
│   │   └── logging   # 日志相关代码
│   └── interfaces    # 表示层
│       ├── api       # RESTful API相关代码
│       └── web       # Web页面相关代码
└── pkg               # 公共代码和工具包
    ├── auth          # 身份验证相关代码
    ├── config        # 配置相关代码
    ├── errors        # 错误处理相关代码
    ├── logger        # 日志相关代码
    └── utils         # 公共工具函数


```

