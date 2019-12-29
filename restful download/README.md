## restful download

> 简介

演示了如何创建 restful 风格的 Go 服务，实现文件下载到本地客户端的功能。

- 运行 main.go
- 运行 client 目录下的 dotnet core 程序

client 端程序将请求 Go 服务器 images 目录下的指定文件。 Go程序读取指定文件，将字节流写入 channel。 client 读取到后，将之写入到 client 所在目录。