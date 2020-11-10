# Go-DB-Doc-Generator

### 一、项目介绍
此项目主要是读取项目项目目录文件进行分析，提取出schema的定义文件，并写到markdown文件中

### 二、使用方法
1. 改下sourceConfig中定义的源目路径
2. 注意看下代码中是否有不符合要求的部分

### 三、参考：

- [GeeORM](https://geektutu.com/post/geeorm.html)
- [GoImports](https://github.com/golang/tools/blob/master/cmd/goimports/goimports.go)
- [GORM](https://github.com/go-gorm/gorm)
    - [GORM源码解读1](https://juejin.im/post/6844904025763086350)
    - [GORM源码解读2](https://juejin.im/post/6844904033648394254)
    - [GORM源码解读3](https://juejin.im/post/6844904047774793735)
- [Go Ast包](https://golang.org/pkg/go/ast/)
- [Go Reflect包](https://golang.org/pkg/reflect/)
- [Beego中实现的swagger](https://github.com/beego/bee/blob/develop/generate/swaggergen/g_docs.go)


