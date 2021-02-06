### go使用protobuf
  * 有两种使用方式 
     1. 使用golang的protobuf
     2. 使用gogo的protobuf
        - gogo的比golang的快一些
        
        
### 编译proto
 * 首先设置好自己的GOPATH
    + 通过go get 命令安装的包都会放到gopath下，使用时也会去gopath查找
    
 * 编译proto需要安装go的编译插件
    + 分别对应gogo的和golang的
 * gogo编译命令
    + protoc -I xxx/xxx/ --gofast_out= xxxx/xxx/  xxx.proto
    > gogo的go_out有多种，自己看使用哪种
 * golang
    + protoc -I xxx/xx/ --go_out=xxx/xx/ xxx.proto
    
 * go的proto文件注意
    + 需要设置 option go_package=",;xxx", xxx来指定生成的go的包名
    
 
    
 1. 使用golang
    + 代码中需要使用库文件：go get github.com/golang/protobuf/proto
    + protoc的插件使用：go get github.com/golang/protobuf/protoc-gen-go
    
    > 对于内网无法下载时，需要手动把golang的protobuf解压到GOPATH/src目录下
    >> 然后go get github.com/golang/protobuf/protoc-gen-go 安装插件，会在GOPATH/bin下生成 protoc-gen-go.exe
    >> 需要把GOPATH/bin加到环境变量中，才能找到此插件
 
 2. 使用gogo
    + 代码库：go get github.com/gogo/protobuf/proto（其实代码库中包括了插件，所以在安装插件时不会在下载了，直接编译了）
    + 插件安装：go get github.com/gogo/protobuf/protoc-gen-gofast 
    + 内网：先安装源码，然后在go get 安装插件
                                                                                                                                                                                                                                                                                                                                                                                                       