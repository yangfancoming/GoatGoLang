# 安装 beego（是源码） bee  （是工具用于构建框架）
    1、beego 的安装  下载缓慢
        go get github.com/astaxie/beego 
    2、bee 工具的安装
         go get github.com/beego/bee
         
    
    func main() {
        beego.Run("localhost:8080")
    }

    运行后打印日志：  http server Running on http://localhost:8080 
    浏览器 输入  http://localhost:8080 
 
 
 # bee
    切换目录到  E:\Code\Golang\workspace\GoatGoLang\chapter6-2-1
    运行  bee new item01
    结论： 无论切换到哪个目录，bee new 只会在工程控件的src目录下创建 web框架
    即： E:\Code\Golang\workspace\src 下 创建 item01 目录
   

# beego 安装报错： package github.com/astaxie/beego: exit status 128

    解决方法1：   
    如果在网络上面都没法解决，最笨的办法就是下载源码安装，
    下载链接：https://github.com/astaxie/beego
           https://github.com/beego/bee
    
    把两个源码放在下面对应的目录
        src/github.com/astaxie/beego/
    然后执行命令：
    src/github.com/beego/bee/
    go install github.com/astaxie/beego
    go install src/github.com/beego/bee
    
    解决方法2：
        下载时候 右键对应目录下的 .git 目录 查看大小。。。
    
    解决方法3：   
        git config --global url."git://github.com/astaxie/beego".insteadOf "https://github.com/astaxie/beego"    //命令修复
        再次运行  go get github.com/astaxie/beego 


