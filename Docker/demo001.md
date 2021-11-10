镜像命令

```shell
aboubakar@ismael:~$ sudo docker images --help

Usage:  docker images [OPTIONS] [REPOSITORY[:TAG]]

List images

Options:
  -a, --all             Show all images (default hides intermediate images)
      --digests         Show digests
  -f, --filter filter   Filter output based on conditions provided
      --format string   Pretty-print images using a Go template
      --no-trunc        Don't truncate output
  -q, --quiet           Only show image IDs
  

aboubakar@ismael:~$ sudo docker images
REPOSITORY   TAG       IMAGE ID       CREATED       SIZE
mysql        latest    ecac195d15af   3 weeks ago   516MB


aboubakar@ismael:~$ sudo docker images -q
ecac195d15af

```

实战

```shell
docker pull redis
Using default tag: latest  
latest: Pulling from library/redis
7d63c13d9b9b: Pull complete # 分层下载，docker image的核心，联合文件系统
a2c3b174c5ad: Pull complete
283a10257b0f: Pull complete
7a08c63a873a: Pull complete
0531663a7f55: Pull complete
9bf50efb265c: Pull complete
Digest: sha256:a89cb097693dd354de598d279c304a1c73ee550fbfff6d9ee515568e0c749cfe
Status: Downloaded newer image for redis:latest
docker.io/library/redis:latest  #真实地址



# 等价
docker pull redis
docker pull docker.io/library/redis:latest



~ docker pull mysql:5.7
5.7: Pulling from library/mysql
b380bbd43752: Already exists
f23cbf2ecc5d: Already exists
30cfc6c29c0a: Already exists
b38609286cbe: Already exists
8211d9e66cd6: Already exists
2313f9eeca4a: Already exists
7eb487d00da0: Already exists
a71aacf913e7: Pull complete
393153c555df: Pull complete
06628e2290d7: Pull complete
ff2ab8dac9ac: Pull complete
Digest: sha256:2db8bfd2656b51ded5d938abcded8d32ec6181a9eae8dfc7ddf87a656ef97e97
Status: Downloaded newer image for mysql:5.7
docker.io/library/mysql:5.7



# 删除 images id=938b57d64674 

➜  ~ docker rmi -f 938b57d64674
Untagged: mysql:5.7
Untagged: mysql@sha256:2db8bfd2656b51ded5d938abcded8d32ec6181a9eae8dfc7ddf87a656ef97e97
Deleted: sha256:938b57d64674c4a123bf8bed384e5e057be77db934303b3023d9be331398b761
Deleted: sha256:d81fc74bcfc422d67d8507aa0688160bc4ca6515e0a1c8edcdb54f89a0376ff1
Deleted: sha256:a6a530ba6d8591630a1325b53ef2404b8ab593a0775441b716ac4175c14463e6
Deleted: sha256:2a503984330e2cec317bc2ef793f5d4d7b3fd8d50009a4f673026c3195460200
Deleted: sha256:e2a4585c625da1cf4909cdf89b8433dd89ed5c90ebdb3a979d068b161513de90

# 全部删除
➜  ~ docker rmi -f $(docker images -aq)
```


容器命令

```shell
➜  ~ docker pull ubuntu

# 启动并进入容器
➜  ~ docker run -it ubuntu /bin/bash
root@658d9ac2bb23:/#

root@658d9ac2bb23:/# exit
exit
➜  ~

# 查看运行的容器

➜  ~ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED       STATUS       PORTS     NAMES
440e5b6b3a30   ubuntu    "bash"    9 hours ago   Up 2 hours             ubuntu

# 曾经运行过的

➜  ~ docker ps -a
CONTAINER ID   IMAGE          COMMAND                  CREATED         STATUS                          PORTS                                        NAMES
658d9ac2bb23   ubuntu         "/bin/bash"              3 minutes ago   Exited (0) About a minute ago                                                strange_pascal
440e5b6b3a30   ubuntu         "bash"                   9 hours ago     Up 2 hours                                                                   ubuntu
c96cefc26fb1   10d7504ea271   "/whoami"                20 hours ago    Exited (255) 4 hours ago        80/tcp                                       demo-whoami-1
937064d706ca   traefik:v2.5   "/entrypoint.sh --ap…"   20 hours ago    Exited (255) 4 hours ago        0.0.0.0:80->80/tcp, 0.0.0.0:8080->8080/tcp   demo-reverse-proxy-1
073ed0c2fc09   mysql          "docker-entrypoint.s…"   43 hours ago    Exited (255) 4 hours ago        0.0.0.0:3306->3306/tcp, 33060/tcp            mysql


➜  ~ docker ps -n=1
CONTAINER ID   IMAGE     COMMAND       CREATED          STATUS                      PORTS     NAMES
658d9ac2bb23   ubuntu    "/bin/bash"   16 minutes ago   Exited (0) 14 minutes ago             strange_pascal

➜  ~ docker ps -aq
658d9ac2bb23
440e5b6b3a30
c96cefc26fb1
937064d706ca
073ed0c2fc09


# 停止容器并退出
➜  ~ docker run -it ubuntu /bin/bash
root@0e8a4299ceb0:/# exit
exit
➜  ~ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES

# contrl + P + Q 容器不停止退出
```

删除容器

```shell
docker rm 容器id


# 删除所有容器

docker rm -f $(docker ps -aq)

docker ps -a -q | xargs docker rm 

```
常用命令

启动停止容器

```shell
docker run -d ubuntu
fef5ab64b693ad7084987563c845feae76415d524bf422a85d1e3c5fd6523b81
docker ps               
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES

# docker 发现没有提供服务，立即停止
```

查看日志

```shell
sudo docker ps
CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS         PORTS      NAMES
3fcd02749b23   redis     "docker-entrypoint.s…"   10 seconds ago   Up 7 seconds   6379/tcp   fervent_elbakyan
aboubakar@ismael:~$ sudo docker logs -tf --tail 10 3fcd02749b23
2021-11-10T13:29:38.964025213Z isamel
2021-11-10T13:29:39.965421949Z isamel
2021-11-10T13:29:40.967395853Z isamel
2021-11-10T13:29:41.968841369Z isamel
2021-11-10T13:29:42.970435198Z isamel
2021-11-10T13:29:43.972006160Z isamel
2021-11-10T13:29:44.973353064Z isamel
2021-11-10T13:29:45.974592921Z isamel
2021-11-10T13:29:46.975870350Z isamel
2021-11-10T13:29:47.977282013Z isamel
2021-11-10T13:29:48.979206674Z isamel
2021-11-10T13:29:49.980747911Z isamel
2021-11-10T13:29:50.982444150Z isamel
2021-11-10T13:29:51.983626233Z isamel
2021-11-10T13:29:52.985514496Z isamel


```

查看容器中进程信息

```shell
sudo docker top 3fcd02749b23
UID                 PID                 PPID                C                   STIME               TTY                 TIME                CMD
root                23905               23884               0                   21:28               ?                   00:00:00            /bin/sh -c while true ; do echo isamel; sleep 1; done
root                24405               23905               0                   21:31               ?                   00:00:00            sleep 1

```
查看镜像元数据

```shell

sudo docker inspect 3fcd02749b23
[
    {
        "Id": "3fcd02749b233f9a25a59aa112ae4e8599c2170e05b00cd20fe72c3611ccf37f",
        "Created": "2021-11-10T13:28:56.354617346Z",
        "Path": "docker-entrypoint.sh",
        "Args": [
            "/bin/sh",
            "-c",
            "while true ; do echo isamel; sleep 1; done"
        ],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 23905,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2021-11-10T13:28:58.886843516Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        ...

```
进入当前正在运行的容器

```shell
➜  ~ docker start fd68d763d455
fd68d763d455
➜  ~ docker ps
CONTAINER ID   IMAGE     COMMAND       CREATED          STATUS         PORTS     NAMES
fd68d763d455   ubuntu    "/bin/bash"   58 seconds ago   Up 3 seconds             lucid_newton


# 进入容器打开新终端
➜  ~ docker exec -it fd68d763d455 /bin/bash
root@fd68d763d455:/#


# 进入容器正在执行的终端，不会打开新终端
➜  ~ docker attach fd68d763d455
root@fd68d763d455:/#

```

从容器内拷贝文件到主机

```shell
root@625456221cb4:/home# touch x.go
root@625456221cb4:/home# ls
x.go
root@625456221cb4:/home# exit
exit

➜  demo docker cp 625456221cb4:/home/x.go /Users/X/demo
➜  demo ls
docker-compose.yml mojo.go            x.go

```


部署 nginx


```shell
 docker pull nginx
Using default tag: latest
latest: Pulling from library/nginx
b380bbd43752: Already exists
fca7e12d1754: Pull complete
745ab57616cb: Pull complete
a4723e260b6f: Pull complete
1c84ebdff681: Pull complete
858292fd2e56: Pull complete
Digest: sha256:644a70516a26004c97d0d85c7fe1d0c3a67ea8ab7ddf4aff193d9f301670cf36
Status: Downloaded newer image for nginx:latest
docker.io/library/nginx:latest


➜  / docker run -d -p 3344:80 nginx
de55b21c29a68f0591ffea5bf760cb08bf2eb40aa54e95ba45709bf558d8f0cb

➜  / docker ps
CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS         PORTS                  NAMES
de55b21c29a6   nginx     "/docker-entrypoint.…"   7 seconds ago   Up 5 seconds   0.0.0.0:3344->80/tcp   pedantic_khorana

➜  / curl  0.0.0.0:3344


➜  / docker exec -it de55b21c29a6 /bin/bash
root@de55b21c29a6:/# whereis nginx
nginx: /usr/sbin/nginx /usr/lib/nginx /etc/nginx /usr/share/nginx
root@de55b21c29a6:/# /etc/nginx
bash: /etc/nginx: Is a directory
root@de55b21c29a6:/# ls
bin   dev		   docker-entrypoint.sh  home  lib64  mnt  proc  run   srv  tmp  var
boot  docker-entrypoint.d  etc			 lib   media  opt  root  sbin  sys  usr
root@de55b21c29a6:/# cd /etc/nginx
root@de55b21c29a6:/etc/nginx# ls
conf.d	fastcgi_params	mime.types  modules  nginx.conf  scgi_params  uwsgi_params

root@de55b21c29a6:/etc/nginx# exit
exit
➜  / docker ps
CONTAINER ID   IMAGE     COMMAND                  CREATED        STATUS        PORTS                  NAMES
de55b21c29a6   nginx     "/docker-entrypoint.…"   11 hours ago   Up 11 hours   0.0.0.0:3344->80/tcp   pedantic_khorana

➜  / docker stop de55b21c29a6
de55b21c29a6

➜  / curl  0.0.0.0:3344
curl: (7) Failed to connect to 0.0.0.0 port 3344: Connection refused

```

docker web可视化管理工具

安装portainer

```shell

docker run -d -p 8088:9000 --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v --privileged=true portainer/portainer

```

commit 镜像

```shell
➜  ~ docker commit --help

Usage:  docker commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]

Create a new image from a container's changes

Options:
  -a, --author string    Author (e.g., "John Hannibal
                         Smith <hannibal@a-team.com>")
  -c, --change list      Apply Dockerfile instruction
                         to the created image
  -m, --message string   Commit message
  -p, --pause            Pause container during
                         commit (default true)

```