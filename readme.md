# GO-BBS Backend Server
### 当前集成了哪些golang开发组件
 - gin http框架
 - gorm 数据库操作
 - redis 缓存组件
 - zap 日志
 - lancet开发工具函数库
 - air 热加载
 - ozzo-validation 验证库
 - ginview 模板库
 - robfig-cron 定时任务
 - gohouse/converter gorm结构体生成工具
 - jordan-wright/email  邮件服务
 - oss云存储
 - prometheus 性能监控
 - olivere/elastic ElasticSearch服务

### 目录说明
- app
    - aggregate 聚合根
    - console 定时任务脚本目录
    - constants 常量
    - entity 实体
    - event  事件
    - exceptions 错误处理，错误信息
    - https 网络请求入口
      - controllers 控制器目录
      - middleware 中间件
      - model 模型层 只负责数据表的字段、属性、查询条件、返回值的定义
        - response 响应结构体VO 
        - requests 请求结构体PO
    - repository 仓库，负责对外的业务逻辑处理
    - service 服务层
    - transform model与vo的转换层
- autocode 代码生成器 repository model requests entity
- config 配置
- core 核心功能库
- deploy 容器化部署配置文件存放目录
- docs 接口文档目录
- global 全局变量
- plugin 插件库
- initialize 初始化程序
- resource 前端资源库
- router 路由
- storage 缓存目录、日志保存、文件上传下载目录
- test 单元测试
- utils 工具库及助手函数

### 项目启动
- 方式一：原始启动
```
go run main.go
```
- 方式二：air热加载启动
```
下载air
go install github.com/cosmtrek/air@latest
然后
air
```

### 项目数据库
本项目中的demo数据库采用的是修罗xiuno-bbs的数据库。而本项目本质就是个脚手架，所以对于使用什么样的数据库什么样的表完全由使用者来决定。