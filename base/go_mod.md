## Go Mod 

### 初始化
```
cd project_name;
go mod init project_name
``` 

> 将会在project_name目录下生成go.mod文件 \
 文件内的内容有project_name为根目录
go build 会查找文件依赖并下载

### 下载依赖
```
go mod download //下载依赖包
go mod tidy //更新依赖包，移除不用的模块
```

### 添加依赖包
#### 方法1:
先修改go.mod文件，再下载 \
``` 
vim go.mod 
go mod download
```

#### 方法2：
go get 会自动更新go.mod \
`go get -u github.com/jinzhu/gorm`

#### 方法3：
go run ,go build 会自动下载依赖

### go.mod 文件
``` golang
module gotest

go 1.13

require (
	github.com/jinzhu/gorm v1.9.12
	github.com/kr/pretty v0.2.0 // indirect
	github.com/satori/go.uuid v1.2.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace (
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
```

>>>
  module 语句指定包的名字（路径）\
  require 语句指定的依赖项模块 \
  replace 语句可以替换依赖项模块 \
  exclude 语句可以忽略依赖项模块 
>>>

### GoLand 设置 (注意)
>1. Preferences -> Go -> Go Modules (vgo)，勾选 Enable Go Modules (vgo) integration 以启用 Go Modules，并在 Proxy 输入框中输入 https://goproxy.io,direct
>2. Preferences -> Go -> GOPATH，勾选上 Index entire GOPATH 以索引整个 GOPATH，不然无法导入包。