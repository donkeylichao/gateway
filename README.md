# gateway
beego框架开发的网关

#### beego gateway

1.基础版本
  - 功能
  - - 后台网关配置（CURD）
  - - 网关配置添加缓存
  - - 网关转发

网关转发功能流程|描述
--|--
获取配置数据|获取service的转发地址配置
解析路由|解析请求路由匹配到转发地址
拼装转发数据|整理调用接口需要的数据请求接口并返回结果


2.完善版本
  - 功能添加
  - - 支持上传文件接口转发



---
安装说明
1. 创建数据库gateway,倒入sql/gateway.sql数据
    - 默认账号 123456@qq.com
    - 密码 123456
2. route配置
    - 路由redis缓存key名称 cache = routeCache
    - 路由参数替换标志 parser_placeholder = [_stub_]
    - - 例如访问地址为/v1/api/demo/3  api_alias配置/v1/api/demo/[_stub_] api_path配置/api/demo/[_stub_]

