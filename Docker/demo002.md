已入门 Docker

容器数据卷

使用数据卷

方式一：直接使用命令来挂载 -v

双向同步数据 好处：只需本地修改、容器内会自动同步

```shell
docker run it -v 主机目录：容器目录

aboubakar@ismael:~$ sudo docker run -it -v /home/aboubakar/go/src/Learning/Docker:/home centos /bin/sh
sh-4.4# ls
bin  dev  etc  home  lib  lib64  lost+found  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
sh-4.4# cd home
sh-4.4# ls
demo001.md  demo002.md
sh-4.4# touch test.go
sh-4.4# ls
demo001.md  demo002.md	test.go


##Source
aboubakar@ismael:~/go/src/Learning/Docker$ ls
demo001.md  demo002.md	test.go


"Mounts": [
            {
                "Type": "bind",
                "Source": "/home/aboubakar/go/src/Learning/Docker",
                "Destination": "/home",
                "Mode": "",
                "RW": true,
                "Propagation": "rprivate"
            }



sh-4.4# exit
exit
aboubakar@ismael:~$ sudo docker ps -a
CONTAINER ID   IMAGE                 COMMAND                  CREATED             STATUS                         PORTS     NAMES
65001becf498   centos                "/bin/sh"                6 minutes ago       Exited (0) 29 seconds ago                romantic_brown

## after exit from the destination and then add test1.go to source

aboubakar@ismael:~/go/src/Learning/Docker$ touch test1.go
aboubakar@ismael:~/go/src/Learning/Docker$ ls
demo001.md  demo002.md	test1.go  test.go


## destination restart

aboubakar@ismael:~$ sudo docker start 65001becf498
65001becf498
aboubakar@ismael:~$ sudo docker attach 65001becf498
sh-4.4# ls
bin  dev  etc  home  lib  lib64  lost+found  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
sh-4.4# cd home
sh-4.4# ls
demo001.md  demo002.md	test.go  test1.go

```

安装mysql同步数据

```shell
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:tag

docker exec -it some-mysql bash

mysql -uroot -p123456


➜  ~ docker pull mysql

➜  ~ docker images
REPOSITORY            TAG       IMAGE ID       CREATED        SIZE
mysql                 latest    ecac195d15af   3 days ago     516MB


➜  ~ docker run -d -p 3310:3306 -v /Users/Shared/mysql/conf:/etc/mysql/conf.d -v /Users/Shared/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 --name mysql01 mysql:latest
ac84769de66024ca380c4336b0e77dba130e7463d73d414afa742a63f444d1de
 
➜  ~ docker exec -it mysql01 bash
root@ac84769de660:/# mysql -uroot -p123456
mysql> create database xn;
Query OK, 1 row affected (0.04 sec)


➜  ~ /Users/Shared/mysql/data
➜  data ls
#ib_16384_0.dblwr  binlog.index       ib_logfile0        performance_schema undo_001
#ib_16384_1.dblwr  ca-key.pem         ib_logfile1        private_key.pem    undo_002

➜  data ls
 xn


➜  ~ docker rm -f mysql01
mysql01
➜  ~ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES

```

数据已同步到本地，将mysql删除、本地数据不会丢失，这就实现容器数据持久化

匿名卷挂载

```shell
aboubakar@ismael:~$ sudo docker run -d -P --name nginx01 -v /ect/nginx nginx
Unable to find image 'nginx:latest' locally
latest: Pulling from library/nginx
7d63c13d9b9b: Already exists 
15641ef07d80: Already exists 
392f7fc44052: Already exists 
8765c7b04ad8: Already exists 
8ddffa52b5c7: Already exists 
353f1054328a: Already exists 
docker: error pulling image configuration: Get "https://production.cloudflare.docker.com/registry-v2/docker/registry/v2/blobs/sha256/04/04661cdce5812210bac48a8af672915d0719e745414b4c322719ff48c7da5b83/data?verify=1636637745-OE9olqWiMsWing%2Fjp32%2Feh0vfVU%3D": net/http: TLS handshake timeout.
See 'docker run --help'.

aboubakar@ismael:~$ sudo docker volume ls
DRIVER    VOLUME NAME
local     0b5995e6443047662f95b21f57a216cae8afd4fe823a416fdb71bc304aff9da1

```

具名卷挂载  

```shell
aboubakar@ismael:~$ sudo docker run -d -P --name nginx02 -v juming-nginx:/ect/nginx nginx
Unable to find image 'nginx:latest' locally
latest: Pulling from library/nginx
7d63c13d9b9b: Already exists 
15641ef07d80: Already exists 
392f7fc44052: Already exists 
8765c7b04ad8: Already exists 
8ddffa52b5c7: Already exists 
353f1054328a: Already exists 
Digest: sha256:dfef797ddddfc01645503cef9036369f03ae920cac82d344d58b637ee861fda1
Status: Downloaded newer image for nginx:latest
57f34a7effa08f8d8fc74f6c2a5aa166ff5a52ae45673b55010852d2b440b082

aboubakar@ismael:~$ sudo docker volume ls
DRIVER    VOLUME NAME
local     juming-nginx


aboubakar@ismael:~$ sudo docker volume  inspect juming-nginx
[
    {
        "CreatedAt": "2021-11-11T20:50:41+08:00",
        "Driver": "local",
        "Labels": null,
        "Mountpoint": "/var/snap/docker/common/var-lib-docker/volumes/juming-nginx/_data",
        "Name": "juming-nginx",
        "Options": null,
        "Scope": "local"
    }
]

```