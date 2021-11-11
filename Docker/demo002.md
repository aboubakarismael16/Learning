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