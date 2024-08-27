## 使用

自己测试使用的windows系统切换php版本工具

**1、配置PHPV_HOME环境变量**

PHPV_HOME作为php软链的目录

**2、安装phpv**

go版本1.22

源码安装

```shell
go isntall github.com/zhang-jianqiang/phpv@latest
```

二进制安装

下载二进制并将二进制所在目录添加到path环境变量

**3、配置**

配置存放多个php版本的目录

```shell
phpv config D:/phpstudy/php
```

列出可用的php版本

```shell
phpv ls
```

切换php版本

```shell
phpv use php7.3.4nts
```
