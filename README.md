# gin-project-template
基于gin的通用项目模板，包含基本的项目目录结构，日志处理、swagger文档，jwt鉴权，时区处理，命令行，配置文件，注册中心和sql语句自动转换结构体等
### 目录结构
- cmd main函数所在文件夹 
- conf 配置文件 
- deploy 容器编排部署配置目录
- docs 文档集合，存放设计文档、开发文档和用户文档等
  - devel/{en-US,zh-CN} 存放开发文档、hack 文档等
  - guide/{en-US,zh-CN} 存放用户手册，安装、quickstart、产品文档等
- global 全局变量
- internal 内部模块
  - dao 数据访问层，所有与数据相关的操作都会在这层，mysql,redis等
  - middleware 中间件
  - model 模型层，用于存放数据库表映射的model对象
  - routers 路由相关逻辑
  - service 项目核心业务逻辑
  - pkg 内部贡献包存放目录
    - errcode 公共错误码
    - validation 内部通用的验证函数
    - middleware 内部HTTP处理链
    - 
- pkg 外部应用可使用的代码，即其他应用可以通过import导入目录下面的代码
- scripts 用于存放各类构建、安装、分析等操作的脚本文件
  - lib 存放 shell 脚本
- storage 项目生成的临时文件夹
- test 测试相关的代码或文件
- third_party 外部帮助工具，分支代码或其他第三方应用，例如Swagger UI
  - forked 例如我们fork一个包，并做了一些修改，也可以将改包放在此目录下
- tools 存放项目的支持工具，可导入来自/pkg 和/internal 目录的代码
