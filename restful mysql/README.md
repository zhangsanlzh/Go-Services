## restful mysql

> 简介

此模块演示了如何使用 Go 创建 restful 风格的 http 服务程序，以操作 mysql 数据库。

除需要预先安装 `mux` 外，还需要先安装 mysql 驱动程序。

``` go
go get github.com/go-sql-driver/mysql
```

然后执行 `main.go`。

利用 `Postman` 工具，分别测试各个 restful 接口功能是否正常。

> 注意

- 请先保证 mysql 服务正常运行。用户名和密码均为 `root`，数据库名为 `testGo`。当然，也可以修改代码来自定义测试库。

- 不同 `url` 对应的 http 动作也不同。修改为对应的 http 动作。