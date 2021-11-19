使用multipass搭建k8s多节点集群和Dashboard

```shell
multipass launch -n master -c 1 -m 3G -d 20G
multipass launch -n node1 -c 1 -m 3G -d 20G
multipass launch -n node2 -c 1 -m 3G -d 20G


aboubakar@ismael:~$ multipass list
Name                    State             IPv4             Image
master                  Running           192.168.105.5    Ubuntu 20.04 LTS
node1                   Running           192.168.105.6    Ubuntu 20.04 LTS
node2                   Running           192.168.105.7    Ubuntu 20.04 LTS

ubuntu@node1:~$ sudo mv /etc/apt/sources.list /etc/apt/sources.list.bak
ubuntu@node1:~$ sudo vim /etc/apt/sources.list

# ubuntu 20.04(focal) 配置如下
deb http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse


ubuntu@master:~$ sudo apt-get update && sudo apt-get upgrade -y

```
Ubuntu 安装 Docker

```shell
# 使用官方安装脚本自动安装
ubuntu@master:~$ curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun

ubuntu@master:~$ docker images
Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Get "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/images/json": dial unix /var/run/docker.sock: connect: permission denied

ubuntu@master:~$ sudo groupadd docker
groupadd: group 'docker' already exists
ubuntu@master:~$ sudo gpasswd -a ubuntu docker
Adding user ubuntu to group docker
ubuntu@master:~$ sudo service docker restart
ubuntu@master:~$ sudo vim /etc/docker/daemon.json

{ "registry-mirrors": [
    "https://hkaofvr0.mirror.aliyuncs.com"
  ]
 }

ubuntu@master:~$ sudo systemctl daemon-reload
ubuntu@master:~$ sudo systemctl restart docker
# 重启 iTerm2
ubuntu@node1:~$ exit
logout
➜  ~ multipass shell node1

ubuntu@master:~$ docker info

 Registry Mirrors:
  https://hkaofvr0.mirror.aliyuncs.com/

# Install Compose on Linux systems

sudo apt install docker-compose -y

ubuntu@master:~$ sudo curl -L "https://github.com/docker/compose/releases/download/v2.0.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

ubuntu@master:~$ sudo chmod +x /usr/local/bin/docker-compose
ubuntu@master:~$ sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
ubuntu@master:~$ docker-compose --version

Docker Compose version v2.0.1

```

root设置

```shell
ubuntu@master:~$ sudo passwd root
ubuntu@master:~$ sudo passwd -dl root

ubuntu@master:~$ su root
Password:
root@master:/home/ubuntu#

root@master:/home/ubuntu# su ubuntu
ubuntu@master:~$

```

ubuntu 安装 microk8s

```shell
ubuntu@master:~$ sudo snap install microk8s --classic
microk8s (1.22/stable) v1.22.2 from Canonical✓ installed

ubuntu@master:~$ microk8s status
Insufficient permissions to access MicroK8s.
You can either try again with sudo or add the user ubuntu to the 'microk8s' group:

    sudo usermod -a -G microk8s ubuntu
    sudo chown -f -R ubuntu ~/.kube

After this, reload the user groups either via a reboot or by running 'newgrp microk8s'.
ubuntu@master:~$ sudo usermod -a -G microk8s ubuntu
ubuntu@master:~$ sudo chown -f -R ubuntu ~/.kube
ubuntu@master:~$ newgrp microk8s

ubuntu@master:~$ microk8s status
microk8s is not running. Use microk8s inspect for a deeper inspection.
ubuntu@master:~$ microk8s inspect

Inspecting Certificates
Inspecting services
  Service snap.microk8s.daemon-cluster-agent is running
  Service snap.microk8s.daemon-containerd is running
  Service snap.microk8s.daemon-apiserver-kicker is running
  Service snap.microk8s.daemon-kubelite is running
  Copy service arguments to the final report tarball
Inspecting AppArmor configuration
Gathering system information
  Copy processes list to the final report tarball
  Copy snap list to the final report tarball
  Copy VM name (or none) to the final report tarball
  Copy disk usage information to the final report tarball
  Copy memory usage information to the final report tarball
  Copy server uptime to the final report tarball
  Copy current linux distribution to the final report tarball
  Copy openSSL information to the final report tarball
  Copy network configuration to the final report tarball
Inspecting kubernetes cluster
  Inspect kubernetes cluster
Inspecting juju
  Inspect Juju
Inspecting kubeflow
  Inspect Kubeflow


The change can be made persistent with: sudo apt-get install iptables-persistent
WARNING:  Docker is installed.
Add the following lines to /etc/docker/daemon.json:
{
    "insecure-registries" : ["localhost:32000"]
}
and then restart docker with: sudo systemctl restart docker
Building the report tarball
  Report tarball is at /var/snap/microk8s/2551/inspection-report-20211031_124330.tar.gz



```