# 研学管理系统

![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)

基于gin + gorm 管理平台，部署方便快捷，目前支持的模块如下

### 用户模块

* 学校管理
* 用户管理
    * 结合JWT鉴权方式，用户登录
    * 手机验证码注册
    * 退出登录
* 角色管理
    * 角色列表
    * 角色详情
    * 角色编辑
    * 创建角色
    * 角色分配
* 权限管理 & 资源管理
    * 资源管理
    * 权限分配

### 课程管理

* 课程列表
* 课程发布

### 资讯管理

* 资讯列表
* 资讯发布

### 路线管理

* 路线列表
* 路线发布

### 基地管理

* 基地管理
* 基地发布

### 运行流程

#### Mysql相关

1. 安装 `brew install mysql`
2. 运行 `mysql.server start`
3. 创建数据库，用户和分配权限

```shell
   echo 'CREATE DATABASE IF NOT EXISTS yanxue DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;'|mysql -uroot -p
   echo "CREATE USER 'admin'@'%' IDENTIFIED BY 'admin';"|mysql -uroot -p
   echo "GRANT ALL PRIVILEGES on yanxue.* to 'admin'@'%';"|mysql -uroot -p
   echo "FLUSH PRIVILEGES;"|mysql -uroot -p
```

4. 验证结果

```
    mysql -uadmin -p
    show databases;
```

5. 如果忘记root密码的话，及时更改

```
mysqld_safe --skip-grant-tables

mysql -uroot
```

```
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'XXX';

如果遇到密码不合规
SET GLOBAL validate_password.LENGTH = 4;
SET GLOBAL validate_password.policy = 0;
SET GLOBAL validate_password.mixed_case_count = 0;
SET GLOBAL validate_password.number_count = 0;
SET GLOBAL validate_password.special_char_count = 0;
SET GLOBAL validate_password.check_user_name = 0;
ALTER USER 'root'@'localhost' IDENTIFIED BY 'XXX';
FLUSH PRIVILEGES;

创建用户
CREATE USER 'admin'@'%' IDENTIFIED BY 'XXXX';

分配权限
GRANT ALL PRIVILEGES on yanxue.* to 'admin'@'%'; 指定数据库
GRANT ALL ON *.* TO 'admin'@'%'; 所有权限

如果端口号没有的话
vim /etc/my.cnf
输入以下内容
[mysqld]
port=3306
```

#### linux Mysql安装

1. 安装 `yum -y install mysql-server`
2. 运行mysql `systemctl start mysqld`
3. 查看状态 `systemctl status mysqld`

#### linux supervisor安装

1. 安装 `yum install supervisor`
2. 运行 `supervisord -c /etc/supervisord.conf`
3. 检查线程状态 `supervisorctl status`
4. 加新线程 `reload`