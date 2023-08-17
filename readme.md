#FreeBns-BBS
###当前集成了哪些golang开发组件
 - gin
 - gorm
 - redis
 - zap
 - lancet开发工具函数库
 - air 热加载
 - ozzo-validation 验证库
 - ginview 模板库

### 目录说明
- app
    - aggregate 聚合根
    - console 定时任务脚本目录
    - consts 常量
    - entity  实体
    - exceptions 错误处理，错误信息
    - https 网络请求入口
      - controllers 控制器目录
      - middleware 中间件
      - model 模型层 只负责数据表的字段、属性、查询条件、返回值的定义
      - requests 请求结构体
    - respository 仓库，负责对外的业务逻辑处理
    - service 服务层
    
- config 配置
- core 核心功能库
- deploy 容器化部署配置文件存放目录
- docs swager目录
- global 全局变量
- plugin 插件库
- initialize 初始化程序
- resource 前端资源库
- router 路由
- storage 缓存目录、日志保存、文件上传下载目录
- test 单元测试
- utils 工具库