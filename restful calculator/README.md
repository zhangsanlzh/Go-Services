## restful calculator

> 简介

这个模块演示了如何使用 Go语言编写 restful 风格的 http 服务器。

需要预先安装 `mux`. 

``` shell
go get github.com/gorilla/mux
```

之后，切换到 main.go 文件所在的目录，执行

``` go
go run .
```

或 

``` go
go run main.go
```

然后，在浏览器或 `postman` 工具中输入 `localhost:8080/3/add/5`, 即可看到 go 版 http 服务执行的结果。

可将 URL 中的动词更改为 `sub`、`mult`、`div`。分别代表 减、乘，除操作。