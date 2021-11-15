Make your own Dockerfile

```shell
aboubakar@ismael:~$ mkdir mydockerfile

aboubakar@ismael:~/mydockerfile$ cat mydockerfile
FROM ubuntu:18.04

MAINTAINER ismael<ismael@gmail.com>

ENV MYPATH /usr/local

WORKDIR $MYPATH


RUN apt-get update && apt-get -y install vim 
RUN apt-get install net-tools 

EXPOSE 80

CMD echo $MYPATH 
CMD echo "end...."
CMD /bin/sh



aboubakar@ismael:~/mydockerfile$ sudo docker build -f mydockerfile -t myubuntu:0.1 .
...
 ---> 56c54d402e79
Successfully built 56c54d402e79
Successfully tagged myubuntu:0.1

REPOSITORY            TAG       IMAGE ID       CREATED              SIZE
myubuntu              0.1       56c54d402e79   About a minute ago   163MB
aboubakar@ismael:~/mydockerfile$ sudo docker run -it myubuntu:0.1
root@56c54d402e79# vim
root@56c54d402e79# pwd
/usr/local
root@56c54d402e79# cd /usr/local
root@56c54d402e79# ls
bin  etc  games  include  lib  man  sbin  share  src
root@56c54d402e79# ifconfig


aboubakar@ismael:~/mydockerfile$ sudo docker history 56c54d402e79
IMAGE          CREATED         CREATED BY                                      SIZE      COMMENT
56c54d402e79   6 minutes ago   /bin/sh -c #(nop)  CMD ["/bin/sh" "-c" "/bin…   0B        
c7d7b8c83d4b   6 minutes ago   /bin/sh -c #(nop)  CMD ["/bin/sh" "-c" "echo…   0B        
6bb8cd3916a4   6 minutes ago   /bin/sh -c #(nop)  CMD ["/bin/sh" "-c" "echo…   0B        
6f31c79a6af4   6 minutes ago   /bin/sh -c #(nop)  EXPOSE 80                    0B        
e805f8366f9a   6 minutes ago   /bin/sh -c apt-get install net-tools            1.47MB    
3df2ebaf31cd   6 minutes ago   /bin/sh -c apt-get update && apt-get -y inst…   98.3MB    
4b54d8780fc3   7 minutes ago   /bin/sh -c #(nop) WORKDIR /usr/local            0B        
35b5f18a965f   7 minutes ago   /bin/sh -c #(nop)  ENV MYPATH=/usr/local        0B        
a1343e418fa4   7 minutes ago   /bin/sh -c #(nop)  MAINTAINER ismael<ismael@…   0B        
5a214d77f5d7   6 weeks ago     /bin/sh -c #(nop)  CMD ["bash"]                 0B        
<missing>      6 weeks ago     /bin/sh -c #(nop) ADD file:0d82cd095966e8ee7…   63.1MB 
```

```shell
aboubakar@ismael:~/mydockerfile$ vim docker-cmd
aboubakar@ismael:~/mydockerfile$ cat docker-cmd
FROM ubuntu:18.04

CMD ["ls", "-s"]


aboubakar@ismael:~/mydockerfile$ sudo docker build -f docker-cmd -t cmd .
Sending build context to Docker daemon  3.072kB
Step 1/2 : FROM ubuntu:18.04
 ---> 5a214d77f5d7
Step 2/2 : CMD ["ls", "-s"]
 ---> Running in 08753a22f377
Removing intermediate container 08753a22f377
 ---> b8418fd1f086
Successfully built b8418fd1f086
Successfully tagged cmd:latest

aboubakar@ismael:~/mydockerfile$ sudo docker run b8418fd1f086
.
..
total 64
4 bin
4 boot
0 dev
4 etc
4 home
4 lib
4 lib64
4 media
4 mnt
...


# 不能追加，必须全部替换

aboubakar@ismael:~/mydockerfile$ sudo docker run b8418fd1f086 -1
docker: Error response from daemon: OCI runtime create failed: container_linux.go:380: starting container process caused: exec: "-1": executable file not found in $PATH: unknown.

aboubakar@ismael:~/mydockerfile$ sudo docker run b8418fd1f086 ls -al
total 72
drwxr-xr-x   1 root root 4096 Nov 15 04:23 .
drwxr-xr-x   1 root root 4096 Nov 15 04:23 ..
-rwxr-xr-x   1 root root    0 Nov 15 04:23 .dockerenv
drwxr-xr-x   2 root root 4096 Sep 30 12:34 bin
drwxr-xr-x   2 root root 4096 Apr 24  2018 boot
drwxr-xr-x   5 root root  340 Nov 15 04:23 dev
drwxr-xr-x   1 root root 4096 Nov 15 04:23 etc
drwxr-xr-x   2 root root 4096 Apr 24  2018 home
drwxr-xr-x   8 root root 4096 May 23  2017 lib
...



# 可以追加命令

aboubakar@ismael:~/mydockerfile$ vim docker-entrypoint
aboubakar@ismael:~/mydockerfile$ cat docker-entrypoint
FROM ubuntu:18.04

ENTRYPOINT ["ls","-a"]


aboubakar@ismael:~/mydockerfile$ sudo docker build -f docker-entrypoint -t entrypoint .
Sending build context to Docker daemon  4.096kB
Step 1/2 : FROM ubuntu:18.04
 ---> 5a214d77f5d7
Step 2/2 : ENTRYPOINT ["ls","-s"]
 ---> Running in c80702a69937
Removing intermediate container c80702a69937
 ---> 49d9121165b9
Successfully built 49d9121165b9
Successfully tagged entrypoint:latest


aboubakar@ismael:~/mydockerfile$ sudo docker run 49d9121165b9
ltotal 64
4 bin
4 boot
0 dev
4 etc
4 home
4 lib
4 lib64
4 media
4 mnt
...


# 可以追加命令

aboubakar@ismael:~/mydockerfile$ sudo docker run 22a3899f045d -l
total 64
4 drwxr-xr-x   2 root root 4096 Dec 17  2019 bin
4 drwxr-xr-x   2 root root 4096 Apr 10  2014 boot
0 drwxr-xr-x   5 root root  340 Nov 15 04:18 dev
4 drwxr-xr-x   1 root root 4096 Nov 15 04:18 etc
4 drwxr-xr-x   2 root root 4096 Apr 10  2014 home
4 drwxr-xr-x  12 root root 4096 Dec 17  2019 lib
4 drwxr-xr-x   2 root root 4096 Dec 17  2019 lib64
...

```

镜像发布

````shell
aboubakar@ismael:~$ sudo docker login
Authenticating with existing credentials...
WARNING! Your password will be stored unencrypted in /root/snap/docker/1125/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded


aboubakar@ismael:~/MyDockerfile$ cat Dockerfile
FROM ubuntu:18.04

MAINTAINER ismael<aboubakarismael16@gmail.com>

ENV MYPATH /usr/local

WORKDIR $MYPATH

RUN apt-get update && apt-get -y install vim

RUN apt-get install net-tools


EXPOSE 80

CMD echo MYPATH
CMD echo "end ..."
CMD /bin/sh


aboubakar@ismael:~/MyDockerfile$ sudo docker build -f Dockerfile -t ismael/ubuntu .
Successfully built bfd63f524382
Successfully tagged ismael/ubuntu:latest

aboubakar@ismael:~/MyDockerfile$ sudo docker images
REPOSITORY            TAG       IMAGE ID       CREATED              SIZE
ismael/ubuntu         latest    bfd63f524382   About a minute ago   163MB

aboubakar@ismael:~/MyDockerfile$ sudo docker push ismael/ubuntu
Using default tag: latest
The push refers to repository [docker.io/ismael/ubuntu]
82e18c734c40: Pushed
b68cb5c1d13d: Pushed
9f54eef41275: Mounted from library/ubuntu
latest: digest: sha256:dc71bd31b77150560a90d0b7faaaecc9b37977df22f9c90ecbfe838e4452e977 size: 951


```