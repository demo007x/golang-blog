## **golang blog**
> golang 快速构建自己的博客。基于gin， gorm实现。

<p class="center">
<img src="https://img.shields.io/github/issues/anziguoer/golang-blog" alt="GitHub issues badge">
<img src="https://img.shields.io/github/forks/anziguoer/golang-blog" alt="GitHub forks badge">
<img src="https://img.shields.io/github/stars/anziguoer/golang-blog" alt="GitHub stars badge">
<img src="https://img.shields.io/github/license/anziguoer/golang-blog" alt="GitHub license badge">
<img src="https://img.shields.io/twitter/url?url=https%3A%2F%2Fgithub.com%2Fanziguoer%2Fgolang-blog" alt="Twitter badge">
</p>


### 说明

golang blog 基于`gin`，`gorm`实现。你可以快速构建自己的博客程序。前端使用`Bootstrap4`构建的模板文件（也可以自己构建模板）。

### 安装

- 设置配置文件。将`.env.example`复制一份， 保存为 .env
    - env 块
        ```ini
        [env]
        Mode=debug //可选参数`debug`, `test`, `release` `debug` 会打印日志文件， `sql`语句到控制台。`test` 测试使用 `release` 线上使用， 不会打印相关的日志，`sql`语句
        Port= //配置运行端口
        AppName= //app名称
        AppUrl= // 项目运行地址， 会自动判定链接跳转地址
        ```
  - database 块
  
    ```ini
    [database]
    Type = mysql // 数据库类型
    User = root  // 数据库账户
    Password = root // 数据库密码
    Host = 127.0.0.1:3306 // 数据库链接地址
    Name = go_blog   // 数据库名
    TablePrefix =    // 数据库前缀
    ```
- 导入数据库sql文件 `sql\sql.sql`

### 启动项目

- `go run main.go`
- `go build -o filename` && `./filename`
    
### **已完成的功能**
- [√] markdown 富文本编辑器
- [√] 自动生成目录

### **计划功能**
- [×] 标签功能优化
- [×] markdown 文件转html文件
- [×] 创建markdown文章,并发布文章
- [×] 命令安装blog
- [×] 深色主题
- [×] 优雅重启

### 感谢
- https://github.com/gin-gonic/gin
- https://gorm.io/zh_CN/
- https://github.com/russross/blackfriday