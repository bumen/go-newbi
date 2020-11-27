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
   
 > [图片](Assets/Res/image/go-newbie-GOPATH.png)
   
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
        > 这个可以通过设置idea中的project gopath，然后启动idea后，自动把project gopath加入到gopath中。
        > 而不会影响系统环境变量的GOPATH，所以如果在idea中开发，然后使用go命令打包其实也不需要手动设置了
   
   > [图片](Assets/Res/image/go-newbie-hello-GOPATH.png)
      
 8. 在idea 中执行Hello.og
   + 配置GOPATH：见图
      - [x] Use GOPATH that's defined in system environment. 这个选择Global GOPATH框里是否使用系统环境变量设置的GOPATH
      - [x] Index entire GOPATH. 这个选项是否索引搜索所有GOPATH中的包，如果不选则只搜索当前项目vendored目录
      - 这两个选项都可以不使用
   
   + 在
   
   > [图片](Assets/Res/image/go-bewbie-idea-GOPATH.png)
   > [图片](Assets/Res/image/go-newbie-idea-hello-result.png)

### windows软链接遇到的问题，（linux不确定）
 * 使用go-newbie/.godeps作用为项目GOPATH时，然后将hello链接到src子目录后的代码同步问题
   + 链接完之后，由于修改了原hello包中的代码后，而链接的hello中代码没有及时更新，导致新加的接口或修改不能及时使用或生效
   + 所以解决
     1. 不使用go-newbie/.godeps作为项目GOPATH, 而将go-newbie作为GOPATH目录，而创建一个src子目录。将包都放到src目录下，可能使用引用时需要加上go-newbie
     2. 每个修改完使用文件reload 快捷键 Ctrl + Alt + Y （目前我使用的是这种）