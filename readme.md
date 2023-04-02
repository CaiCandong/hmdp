# 基于黑马点评的Golang实现
## 前端服务启动
```shell
    cd ./views/nginx-1.23.4
    ./nginx.exe
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
│   ├── app           # 应用层：连接 domain 和 interfaces 层
│   │   ├── dto       # DTO(Data Transfer Object)数据传输对象
│   │   ├── service     # 很薄的一层，没有业务逻辑, 不应该存在if/else这种判断
│   │   └── user      # 用户相关服务
│   ├── domain        # 领域层
│   │   ├── aggregate     # 把需要"一起操作"的实体 放到一起
│   │   ├── repository    # 形态上就是一个接口定义
│   │   ├── entity        # 实体必须有唯一标识
│   │   ├── valueobject   # 除了唯一标志，其他任何字段都可以是值对象,只需根据“值”就能判断两者是否相等
│   │   └── serivce      # 用户领域模型和服务接口
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

# 参考github
https://github.com/victorsteven/food-app-server/tree/master/infrastructure

# 相关概念介绍
- Entity - is a class with an ID. In the case of relational DB it's usually a class that's mapped to a DB table with some primary key.
- DTO (Data Transfer Object) - is a class that maps well on what you're sending over the network. E.g. if you exchange JSON or XML data, it usually has fields just enough to fill those requests/responses. Note, that it may have fewer or more fields than Entity.
- VO (Value Object) is a class-value. E.g. you could create class like Grams or Money - it will contain some primitives (e.g. some double value) and it's possible to compare Value Objects using these primitives. They don't have a database ID. They help replacing primitives with more object-oriented classes related to our particular domain.
- Domain Model contains all Entities and Value Objects. And some other types of classes depending on the classification you use.