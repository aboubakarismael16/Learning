Dockerfile base 


```shell
aboubakar@ismael:~$ mkdir docker-volumes
aboubakar@ismael:~$ cd docker-volumes
aboubakar@ismael:~/docker-volumes$ pwd
/home/aboubakar/docker-volumes
aboubakar@ismael:~/docker-volumes$ touch dockerfile01
aboubakar@ismael:~/docker-volumes$ ls
dockerfile01

aboubakar@ismael:~/docker-volumes$ vim dockerfile01
aboubakar@ismael:~/docker-volumes$ cat dockerfile01
FROM centos

VOLUME ["volume01","volume02"]

CMD  ech "____ismael-centos_____"


CMD /bin/sh


aboubakar@ismael:~/docker-volumes$ sudo docker build -f dockerfile01 -t isamel/centos .
Sending build context to Docker daemon  1.536kB
Error response from daemon: the Dockerfile (dockerfile01) cannot be empty
aboubakar@ismael:~/docker-volumes$ sudo docker build -f dockerfile01 -t isamel/centos .
Sending build context to Docker daemon  3.072kB
Step 1/4 : FROM centos
 ---> 5d0da3dc9764
Step 2/4 : VOLUME ["volume01","volume02"]
 ---> Running in 0c3bdcf987aa
Removing intermediate container 0c3bdcf987aa
 ---> bb1b125d8bc2
Step 3/4 : CMD  ech "____ismael-centos_____"
 ---> Running in 28ff69e2150d
Removing intermediate container 28ff69e2150d
 ---> 25fbff97ac6b
Step 4/4 : CMD /bin/sh
 ---> Running in 9dc97b39f055
Removing intermediate container 9dc97b39f055
 ---> e894f1865326
Successfully built e894f1865326
Successfully tagged isamel/centos:latest


aboubakar@ismael:~/docker-volumes$ sudo docker images
REPOSITORY            TAG       IMAGE ID       CREATED              SIZE
isamel/centos         latest    e894f1865326   About a minute ago   231MB
centos                latest    5d0da3dc9764   8 weeks ago     231MB


## version:1.0

aboubakar@ismael:~/docker-volumes$ sudo docker build -f /home/aboubakar/docker-volumes/dockerfile01 -t ismael/centos:1.0 .
Sending build context to Docker daemon  3.072kB
Step 1/4 : FROM centos
 ---> 5d0da3dc9764
Step 2/4 : VOLUME ["volume01","volume02"]
 ---> Using cache
 ---> bb1b125d8bc2
Step 3/4 : CMD  ech "____ismael-centos_____"
 ---> Using cache
 ---> 25fbff97ac6b
Step 4/4 : CMD /bin/sh
 ---> Using cache
 ---> e894f1865326
Successfully built e894f1865326
Successfully tagged ismael/centos:1.0


aboubakar@ismael:~/docker-volumes$ sudo docker images
REPOSITORY            TAG       IMAGE ID       CREATED         SIZE
isamel/centos         latest    e894f1865326   6 minutes ago   231MB
ismael/centos         1.0       e894f1865326   6 minutes ago   231MB
centos                latest    5d0da3dc9764   8 weeks ago     231MB


```

启动自己镜像

```shell
aboubakar@ismael:~/docker-volumes$ sudo docker images
REPOSITORY            TAG       IMAGE ID       CREATED         SIZE
isamel/centos         latest    e894f1865326   6 minutes ago   231MB
ismael/centos         1.0       e894f1865326   6 minutes ago   231MB
centos                latest    5d0da3dc9764   8 weeks ago     231MB

aboubakar@ismael:~/docker-volumes$ sudo docker run -it e894f1865326 /bin/sh
sh-4.4# ls
bin  dev  etc  home  lib  lib64  lost+found  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var  volume01	volume02

```

找到volume目录

```shell
aboubakar@ismael:~/docker-volumes$ sudo docker inspect e894f1865326
 
...

"Image": "sha256:25fbff97ac6b37a785192f7298e37dd533013fedce57b66458231ea6b29dd316",
            "Volumes": {
                "volume01": {},
                "volume02": {}
            },
 ....

```

查看数据同步成功

```shell

```