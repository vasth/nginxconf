# nginxconf
web页面自由配置nginx虚拟主机

该项目基于golang的 beego web框架 主要实现了一下功能
创建新的nginx代理配置文件 并自动调用 nginx reload 命令
暂时没有实现删除配置文件命令

#应用场景
nginx需要实现代理访问多个主机，操作人员无需懂得nginx配置只需要在web页面点击下鼠标即可
多个 golang 应用程序在服务器后台运行在非80端口，nginx更具不同的域名代理访问不同的golang应用程序

#改造后可以做什么？
增加设置多个nginx配置模板，可以实现负载均衡、自动配置虚拟主机等高级的nginx的功能
nginx配置模板为nginxconf/template/domainlocalconf.tpl

#该项目使用帮助
由于我在windows上编译的，暂时没有找到交叉编译sqlite到linux的方法，该程序的数据存储于sqlite，虽然我在linux环境下编译出来了32位的编译文件，但是由于不会上传到github上，如果需要可以发送邮件到1935873589@qq.com. 运行成功后输入用户名 admin
密码 adminpwd，如果需要更改密码请在/conf/app.conf配置文件里面更改 
user = admin
pwd = adminpwd

#该项目用到的库
beego  "github.com/astaxie/beego"
go-sh  "github.com/codeskyblue/go-sh"


