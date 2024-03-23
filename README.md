# Go Connector for TDengine

## 改动说明
原driver-go对于跨平台的编译环境不够友好。
我们需要搭建arm64/amd64/darwin多个单独的编译环境。
TDengine版本迭代又非常频繁，导致我们时不时要维护不止一个环境。（加之我们其他一些C库）
这里改了下编译的兼容性：

```sh
./wrapper/taosc.go
```

```sh
#cgo CFLAGS: -IC:/TDengine/include -I/usr/include
#cgo linux LDFLAGS: -L/usr/lib -ltaos
#cgo windows LDFLAGS: -LC:/TDengine/driver -ltaos
#cgo darwin LDFLAGS: -L/usr/local/lib -ltaos
```

## todo
接下来考虑用dlopen的方式，进一步降低对编译环境的依赖。

## 跳转至官方文档
driver-go其他更详细使用看官方文档：
[driver-go](https://github.com/taosdata/driver-go)

