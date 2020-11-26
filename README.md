# go-newbie
go anything is here

### windows idea 2020.1.3中创建hello package
 1. git clone go-newbie
 2. 使用idea 打开项目go-newbie
 3. 创建hello文件夹
 4. 在go-newbie目录中打开git bash 窗口
   + source gvp
      - 初始化.godeps
   + gpm-local to /xxx/go-newbie/hello
      - 将hello 加入到项目GOPATH中
   + gvp out
   
 > 
   
 5. cd hello
 6. 创建hello.go, world/world.go， Godep-Git依赖
   + hello.go 中import 相关包
 7. 在hello目录中打开git bash 窗口
   + source gvp
      - 初始化.godeps
   + gpm-git
      - 下载依赖
   + go run Hello.go
      - 使用命令行方式执行
      - 报了一个hello/world包找不到，因为GOPATH中没有go-newbie/.godeps目录。
      - 解决方法，需要手动把go-newbie/.godeps目录加载到GOPATH环境变量中（如果通过idea执行，则可以配置GOPATH）
      
 8. 在idea 中执行Hello.og
   + 配置GOPATH：见图
      - [x] Use GOPATH that's defined in system environment. 这个选择Global GOPATH框里是否使用系统环境变量设置的GOPATH
      - [x] Index entire GOPATH. 这个选项是否索引搜索所有GOPATH中的包，如果不选则只搜索当前项目vendored目录
      - 这两个选项都可以不使用
   
