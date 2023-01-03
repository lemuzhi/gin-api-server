# gin-project-template
基于gin的通用项目模板，包含基本的项目目录结构，日志处理、swagger文档，jwt鉴权，时区处理，命令行，配置文件，注册中心和sql语句自动转换结构体等
### 目录结构
- cmd 启动入口 
- conf 配置文件 
- docs 文档集合 
- global 全局变量
- internal 内部模块
  - dao 数据访问层，所有与数据相关的操作都会在这层，mysql,redis等
  - middleware 中间件
  - model 模型层，用于存放数据库表映射的model对象
  - routers 路由相关逻辑
  - service 项目核心业务逻辑
- pkg 项目相关的模块包
- storage 项目生成的临时文件夹
- scripts 各类构建、安装、分析等操作的脚本
- third_party 第三方的资源工具
