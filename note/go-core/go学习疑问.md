# go 学习疑问记录

## GOPATH 的易错点

- gopath 可以有多个，在mac 中":"分隔

- 在用 goland 开发时，需要在设置中设置 project path,并重启，这时 go build/install 才能生效

- go build 默认在 gopath 路径 + src + 指定路径（其中 src是编译器加上的，自己再加就会报错）

