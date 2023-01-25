# Install Kubernetes Cluster using kubeadm

[The Easy Way to Install Kubernetes 1.22 and containerd](https://medium.com/itnext/the-easy-way-to-install-kubernetes-1-22-and-containerd-fec2d07912bd)

```shell
curl -L https://github.com/kubesphere/kubekey/releases/download/v1.2.0-alpha.4/kubekey-v1.2.0-alpha.4-linux-amd64.tar.gz > installer.tar.gz && tar -zxf installer.tar.gz
export KKZONE=cn # Run the following command first to make sure you download KubeKey from the correct zone
apt install conntrack
apt install socat
## cluster only with one node
./kk create cluster --with-kubernetes v1.22.1 --yes --skip-pull-images

## cluster with multiple nodes
##1. First connect to master node and worker node
sudo su
## 2. generate ssh  in each node
ssh-keygen -t rsa 
##3. share the publickey and connect the worker node to master node
ssh <worker_node_username>@<worker_node_ip>
## 4. Create an example configuration file
Create an example configuration file
## 5. Create an example configuration file
## 6. Create a cluster using the configuration file
./kk create cluster -f config-sample.yaml --with-kubernetes v1.22.1 --yes --skip-pull-images

```

Follow this documentation to set up a Kubernetes cluster on __Ubuntu 20.04 LTS__.

This documentation guides you in setting up a cluster with one master node and one worker node.

## Assumptions

```shell
multipass launch -n master -c 2 -m 2G -d 10G
multipass launch -n node1 -c 1 -m 1G -d 10G
```

## On both master and node

##### Login as `root` user

```shell
sudo su -
```

Perform all the commands as root user unless otherwise specified

##### Disable Firewall

```shell
ufw disable
```

##### Disable swap

```shell
swapoff -a; sed -i '/swap/d' /etc/fstab
```

##### Update sysctl settings for Kubernetes networking

```shell
cat >>/etc/sysctl.d/kubernetes.conf<<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sysctl --system
```

##### Install docker engine

```shell
{
  apt install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common
  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
  add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
  apt update
  apt install -y docker-ce=5:19.03.10~3-0~ubuntu-focal containerd.io
  
}
```

```shell
~$ sudo groupadd docker
groupadd: group 'docker' already exists
~$ sudo gpasswd -a ubuntu docker
Adding user ubuntu to group docker
~$ sudo service docker restart
~$ sudo vim /etc/docker/daemon.json

{ "registry-mirrors": [
    "https://hkaofvr0.mirror.aliyuncs.com"
  ]
 }

~$ sudo systemctl daemon-reload
~$ sudo systemctl restart docker

```

```shell
~$ sudo mv /etc/apt/sources.list /etc/apt/sources.list.bak
~$ sudo vim /etc/apt/sources.list

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

~$ sudo apt-get update && sudo apt-get upgrade -y

```

### Kubernetes Setup

##### Add Apt repository

```shell
{
    echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list
}
```

### apt 官方源 有时下载不下来  并添加Kubernetes安装的密钥

```shell
sudo apt update && sudo apt install -y apt-transport-https curl
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -
 
sudo touch /etc/apt/sources.list.d/kubernetes.list 
sudo echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" >> /etc/apt/sources.list.d/kubernetes.list

```

### 配置国内中科大的源 并添加Kubernetes安装的密钥

```shell
curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add - 
sudo echo "deb http://mirrors.ustc.edu.cn/kubernetes/apt kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list
sudo apt update
```

##### Install Kubernetes components

```shell
apt update && apt install -y kubeadm=1.18.5-00 kubelet=1.18.5-00 kubectl=1.18.5-00
```

### 保持版本，取消自动更新

```shell
sudo apt-mark hold kubelet kubeadm kubectl
```

##### In case you are using LXC containers for Kubernetes nodes

Hack required to provision K8s v1.15+ in LXC containers

```shell
{
  mknod /dev/kmsg c 1 11
  echo '#!/bin/sh -e' >> /etc/rc.local
  echo 'mknod /dev/kmsg c 1 11' >> /etc/rc.local
  chmod +x /etc/rc.local
}
```

## On master

##### Initialize Kubernetes Cluster

Update the below command with the ip address of kmaster

```shell
root@master:~# sudo kubeadm init --image-repository=registry.aliyuncs.com/google_containers  --pod-network-cidr=192.168.0.0/16 --apiserver-advertise-address=10.218.50.66  --kubernetes-version=v1.18.5

W1122 17:11:15.053634   16191 configset.go:202] WARNING: kubeadm cannot validate component configs for API groups [kubelet.config.k8s.io kubeproxy.config.k8s.io]
[init] Using Kubernetes version: v1.18.5
[preflight] Running pre-flight checks
 [WARNING IsDockerSystemdCheck]: detected "cgroupfs" as the Docker cgroup driver. The recommended driver is "systemd". Please follow the guide at https://kubernetes.io/docs/setup/cri/
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Starting the kubelet
[certs] Using certificateDir folder "/etc/kubernetes/pki"
...
Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 10.218.50.66:6443 --token mt0bg6.114do561l08t13ok \
    --discovery-token-ca-cert-hash sha256:ca67dee7ff72399b6006b6e1db985fd730cb3cd2d2047149d2b7a132df6d672c 


```

##### Deploy Calico network

```shell
kubectl --kubeconfig=/etc/kubernetes/admin.conf create -f https://docs.projectcalico.org/v3.14/manifests/calico.yaml
```

##### Cluster join command

```shell
root@master:~# kubeadm token create --print-join-command
W1122 17:15:24.495946   22552 configset.go:202] WARNING: kubeadm cannot validate component configs for API groups [kubelet.config.k8s.io kubeproxy.config.k8s.io]
kubeadm join 10.218.50.66:6443 --token xclp7y.eolz1mh3ytgqftby     --discovery-token-ca-cert-hash sha256:ca67dee7ff72399b6006b6e1db985fd730cb3cd2d2047149d2b7a132df6d672c 
```

##### To be able to run kubectl commands as non-root user

If you want to be able to run kubectl commands as non-root user, then as a non-root user perform these

```shell
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

## On node

##### Join the cluster

Use the output from __kubeadm token create__ command in previous step from the master server and run here.

## Verifying the cluster (On master)

##### Get Nodes status

```shell
ubuntu@master:~$ kubectl get nodes
NAME     STATUS     ROLES    AGE     VERSION
master   Ready      master   7m10s   v1.18.5
node1    Ready   <none>   71s     v1.18.5
```

##### Get component status

```shell
ubuntu@master:~$ kubectl get cs
NAME                 STATUS    MESSAGE             ERROR
scheduler            Healthy   ok            
controller-manager   Healthy   ok            
etcd-0               Healthy   {"health":"true"}   

```

Have Fun Journey 🏄🏄
