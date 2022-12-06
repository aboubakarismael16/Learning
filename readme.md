# Learning note 🐣🐣🐣


## Kubernetes img src="/docs/imgs/kubernetes-icon.svg" style="zoom:50%;" />
[Kubernetes](https://kubernetes.io/docs/concepts/overview/), also known as K8s, is an open-source system for automating deployment, scaling, and management of containerized applications.


---

| contents                      |                detail                                              |
| ----------------------------- | ------------------------------------------------------------ |
| install-cluster-Ubuntu20      | [setup cluster](https://github.com/aboubakarismael16/Learning/blob/main/K8s/install-cluster-Ubuntu20.md) |
| k8s-1.24.sh                   | [install kubernetes 1.24](https://github.com/aboubakarismael16/Learning/blob/main/K8s/k8s-1.24.sh) |
| kuberntes in action           | [Kubernetes核心实战](https://github.com/aboubakarismael16/Learning/blob/main/K8s/k8s-in-Action.md) |
| kubernetes_setup_using_eksctl | [Setup Kubernetes on Amazon EKS](https://github.com/aboubakarismael16/Learning/blob/main/K8s/kubernetes_setup_using_eksctl.md) |
| vm                            | [vm is use to monitoring](https://github.com/aboubakarismael16/Learning/tree/main/K8s/vm) |
| openkruise                    | [openkruise is open-source project](https://github.com/aboubakarismael16/Learning/tree/main/K8s/openkruise) |
| multipass                     | [multipass is bunch of files that content a lot of projects](https://github.com/aboubakarismael16/Learning/tree/main/K8s/multipass) |
| efk                           | [elasticsearch-fluentd-kibana](https://github.com/aboubakarismael16/Learning/tree/main/K8s/efk) |
| prometheus                    | [prometheus and grafana for monitoring](https://github.com/aboubakarismael16/Learning/tree/main/K8s/prometheus) |



---------------------------------------
## Mysql


### [Demo1](https://github.com/aboubakarismael16/Learning/blob/main/Mysql/database/demo001.txt)
- 1、什么是数据库？什么是数据库管理系统？什么是SQL？他们之间的关系是什么？
- 2、安装MySQL数据库管理系统。
- 3、MySQL数据库的完美卸载！
- 4、看一下计算机上的服务，找一找MySQL的服务在哪里？
- 5、在windows操作系统当中，怎么使用命令来启动和关闭mysql服务呢？
- 6、mysql安装了，服务启动了，怎么使用客户端登录mysql数据库呢？
- 7、mysql常用命令：
- 8、数据库当中最基本的单元是表：table
- 9、关于SQL语句的分类？
- 10、导入一下提前准备好的数据

### [Demo2](https://github.com/aboubakarismael16/Learning/blob/main/Mysql/database/demo002.txt)
- 11、关于导入的这几张表？
- 12、不看表中的数据，只看表的结构，有一个命令：
- 13、简单查询
  - 13.1、查询一个字段？
  - 13.2、查询两个字段，或者多个字段怎么办？
  - 13.3、查询所有字段怎么办？
  - 13.4、给查询的列起别名？
  - 13.5、计算员工年薪？sal * 12

### [Demo3](https://github.com/aboubakarismael16/Learning/blob/main/Mysql/database/demo003.txt)
- 14、条件查询
- 14.1、什么是条件查询？
  - 14.2、都有哪些条件？
    - `=` 等于
    - `<>`或 `!=` 不等于
    - `<` 小于
    - `<=` 小于等于
    - `>` 大于
    - `>=` 大于等于
    - `is null` 为 null（is not null 不为空）
    - 查询哪些员工的津贴/补助不为null？
    - `and` 并且
    - `or` 或者
    - `and`和`or`同时出现的话，有优先级问题吗？
    - 查询工作岗位是MANAGER和SALESMAN的员工？
    - 查询薪资是800和5000的员工信息？
    - `not` 可以取非，主要用在 `is` 或 `in` 中
    - `like`
    - 找出名字中含有O的？
    - 找出名字以T结尾的？
    - 找出名字以K开始的？
    - 找出第二个字每是A的？
    - 找出第三个字母是R的？
    - 找出名字中有“_”的？

### [Demo4](https://github.com/aboubakarismael16/Learning/blob/main/Mysql/database/demo004.txt)
- 15、排序
  - 15.1、查询所有员工薪资，排序？
  - 15.2、怎么降序？
  - 15.3、可以两个字段排序吗？或者说按照多个字段排序？
  - 15.4、了解：根据字段的位置也可以排序
- 16、综合一点的案例
- 17、数据处理函数
  - 17.1、数据处理函数又被称为单行处理函数
  - 17.2、单行处理函数常见的有哪些？
     - `lower` 转换小写
     - `upper` 转换大写
     - `substr` 取子串（substr( 被截取的字符串, 起始下标,截取的长度)）
     - `concat`函数进行字符串的拼接
     - `length` 取长度
     - `trim` 去空格
     - `round` 四舍五入
     - `rand()` 生成随机数
     - `ifnull` 可以将 `null` 转换成一个具体值

### [Demo5](https://github.com/aboubakarismael16/Learning/blob/main/Mysql/database/demo005.txt)
- 18、分组函数（多行处理函数）
- 19、分组查询（非常重要：五颗星*****）
	- 19.1、什么是分组查询？
	- 19.2、将之前的关键字全部组合在一起，来看一下他们的执行顺序？
	- 19.3、找出每个工作岗位的工资和？
	- 19.4、找出每个部门的最高薪资
	- 19.5、找出“每个部门，不同工作岗位”的最高薪资？
	- 19.6、使用having可以对分完组之后的数据进一步过滤。
	- 19.7、where没办法的？？？？
- 20、大总结（单表的查询学完了）
	- 执行顺序？
		- 1. `from`
		- 2. `where`
		- 3. `group by`
		- 4. `having`
		- 5. `select`
		- 6. `order by`

### [Demo6](https://github.com/aboubakarismael16/Learning/blob/main/Mysql/database/demo006.txt)
- 1、把查询结果去除重复记录【distinct】
- 2、连接查询
  - 2.1、什么是连接查询？
  - 2.2、连接查询的分类？
  - 2.3、当两张表进行连接查询时，没有任何条件的限制会发生什么现象？
  - 2.4、怎么避免笛卡尔积现象？
  - 2.5、内连接之等值连接。
  - 2.6、内连接之非等值连接
  - 2.7、内连接之自连接
  - 2.8、外连接
  - 2.9、三张表，四张表怎么连接？

### [Demo7](https://github.com/aboubakarismael16/Learning/blob/main/Mysql/database/demo007.txt)
- 3、子查询？
  - 3.1、什么是子查询？
  - 3.2、子查询都可以出现在哪里呢？
  - 3.3、`where`子句中的子查询
  - 3.4、`from`子句中的子查询
  - 3.5、`select`后面出现的子查询（这个内容不需要掌握，了解即可！！！）
- 4、`union`合并查询结果集
- 5、`limit`（非常重要）
  - 5.1、`limit`作用：将查询结果集的一部分取出来。通常使用在分页查询当中。
  - 5.2、`limit`怎么用呢？
  - 5.3、注意：mysql当中`limit`在`order by`之后执行！！！！！！
  - 5.4、取出工资排名在[3-5]名的员工？
  - 5.5、取出工资排名在[5-9]名的员工？
  - 5.6、分页
- 6、关于DQL语句的大总结：
- 7、表的创建（建表）
  - 7.1、建表的语法格式：(建表属于DDL语句，DDL包括：create drop alter)
  - 7.2、关于mysql中的数据类型？
  - 7.3、创建一个学生表？
  - 7.4、插入数据`insert` （DML）
  - 7.5、`insert`插入日期
  - 7.6、`date`和`datetime`两个类型的区别？
  - 7.7、修改`update`（DML）
  - 7.8、删除数据 `delete` （DML）

### [Demo8]()
- 1、查询每一个员工的所在部门名称？要求显示员工名和部门名。
- 2、`insert`语句可以一次插入多条记录吗？【掌握】
- 3、快速创建表？【了解内容】
- 4、将查询结果插入到一张表当中？`insert`相关的！！！【了解内容】
- 5、快速删除表中的数据？【`truncate`比较重要，必须掌握】
- 6、对表结构的增删改？
- 7、约束（非常重要，五颗星*****）
  - 7.1、什么是约束？
  - 7.2、约束包括哪些？
  - 7.3、非空约束：`not null`
  - 7.4、唯一性约束: `unique`
  - 7.5、主键约束（`primary key`，简称PK）非常重要五颗星*****
  - 7.6、外键约束（`foreign key`，简称FK）非常重要五颗星*****
- 8、存储引擎（了解内容）
  - 8.1、什么是存储引擎，有什么用呢？
  - 8.2、怎么给表添加/指定“存储引擎”呢？
  - 8.3、怎么查看mysql支持哪些存储引擎呢？
  - 8.4、关于mysql常用的存储引擎介绍一下
- 9、事务（重点：五颗星*****，必须理解，必须掌握）
  - 9.1、什么是事务？
  - 9.2、只有DML语句才会有事务这一说，其它语句和事务无关！！！
  - 9.3、假设所有的业务，只要一条DML语句就能完成，还有必要存在事务机制吗？
  - 9.4、事务是怎么做到多条DML语句同时成功和同时失败的呢？
  - 9.5、怎么提交事务，怎么回滚事务？
  - 9.6、事务包括4个特性？
  - 9.7、重点研究一下事务的隔离性！！！
  - 9.8、验证各种隔离级别

### [Demo9](https://github.com/aboubakarismael16/Learning/blob/main/Mysql/database/demo009.txt)
- 1、索引（index）
  - 1.1、什么是索引？
  - 1.2、索引的实现原理？
  - 1.3、在mysql当中，主键上，以及unique字段上都会自动添加索引的！！！！
  - 1.4、索引怎么创建？怎么删除？语法是什么？
  - 1.5、在mysql当中，怎么查看一个SQL语句是否使用了索引进行检索？
  - 1.6、索引有失效的时候，什么时候索引失效呢？
  - 1.7、索引是各种数据库进行优化的重要手段。优化的时候优先考虑的因素就是索引。
- 2、视图(view)
  - 2.1、什么是视图？
  - 2.2、怎么创建视图对象？怎么删除视图对象？
  - 2.3、用视图做什么？
  - 2.4、视图对象在实际开发中到底有什么用？《方便，简化开发，利于维护》
- 3、DBA常用命令？
- 4、数据库设计三范式
  - 4.1、什么是数据库设计范式？
  - 4.2、数据库设计范式共有？
  - 4.3、第一范式
  - 4.4、第二范式：
  - 4.5、第三范式
  - 4.6、总结表的设计？
  - 4.7、嘱咐一句话
## Docker

## K8s