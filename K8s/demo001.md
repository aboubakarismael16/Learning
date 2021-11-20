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

ubuntu@master:~$ sudo iptables -P FORWARD ACCEPT
ubuntu@master:~$ sudo vim /etc/docker/daemon.json

{   "exec-opts": ["native.cgroupdriver=systemd"],
    "registry-mirrors":[
        "https://hkaofvr0.mirror.aliyuncs.com",
        "http://docker.mirrors.ustc.edu.cn"
    ],
    "insecure-registries" : ["localhost:32000"]
}

ubuntu@master:~$ sudo systemctl restart docker


ubuntu@master:~$ sudo vim /var/snap/microk8s/current/args/kubelet

--pod-infra-container-image=s7799653/pause:3.1

ubuntu@master:~$ sudo vim /var/snap/microk8s/current/args/containerd-template.toml

sandbox_image = "s7799653/pause:3.1"

 "registry-mirrors":[
        "https://hkaofvr0.mirror.aliyuncs.com",
    ]

ubuntu@master:~$ microk8s.stop && microk8s.start



设置kubectl别名：

sudo snap alias microk8s.kubectl kubectl

ubuntu@master:~$ kubectl get po -n kube-system
NAME                                      READY   STATUS     RESTARTS   AGE
calico-node-kxs9s                         0/1     Init:0/3   0          6m40s
calico-kube-controllers-c478f855f-q5xfg   0/1     Pending    0          6m40s

```


```shell
ubuntu@master:~$ kubectl describe pods calico-node-kxs9s -n kube-system
Name:                 calico-node-kxs9s
Namespace:            kube-system
Priority:             2000001000
Priority Class Name:  system-node-critical
Node:                 master/10.218.50.30
Start Time:           Sat, 20 Nov 2021 20:52:13 +0800
Labels:               controller-revision-hash=74d696bc89
                      k8s-app=calico-node
                      pod-template-generation=2
Annotations:          kubectl.kubernetes.io/restartedAt: 2021-11-20T20:51:08+08:00
                      scheduler.alpha.kubernetes.io/critical-pod: 
Status:               Pending
IP:                   10.218.50.30
IPs:
  IP:           10.218.50.30
Controlled By:  DaemonSet/calico-node
Init Containers:
  upgrade-ipam:
    Container ID:  
    Image:         docker.io/calico/cni:v3.19.1
    Image ID:      
    Port:          <none>
    Host Port:     <none>
    Command:
      /opt/cni/bin/calico-ipam
      -upgrade
    State:          Waiting
      Reason:       PodInitializing
    Ready:          False
    Restart Count:  0
    Environment Variables from:
      kubernetes-services-endpoint  ConfigMap  Optional: true
    Environment:
      KUBERNETES_NODE_NAME:        (v1:spec.nodeName)
      CALICO_NETWORKING_BACKEND:  <set to the key 'calico_backend' of config map 'calico-config'>  Optional: false
    Mounts:
      /host/opt/cni/bin from cni-bin-dir (rw)
      /var/lib/cni/networks from host-local-net-dir (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-mpznt (ro)
  install-cni:
    Container ID:  
    Image:         docker.io/calico/cni:v3.19.1
    Image ID:      
    Port:          <none>
    Host Port:     <none>
    Command:
      /opt/cni/bin/install
    State:          Waiting
      Reason:       PodInitializing
    Ready:          False
    Restart Count:  0
    Environment Variables from:
      kubernetes-services-endpoint  ConfigMap  Optional: true
    Environment:
      CNI_CONF_NAME:         10-calico.conflist
      CNI_NETWORK_CONFIG:    <set to the key 'cni_network_config' of config map 'calico-config'>  Optional: false
      KUBERNETES_NODE_NAME:   (v1:spec.nodeName)
      CNI_MTU:               <set to the key 'veth_mtu' of config map 'calico-config'>  Optional: false
      SLEEP:                 false
      CNI_NET_DIR:           /var/snap/microk8s/current/args/cni-network
    Mounts:
      /host/etc/cni/net.d from cni-net-dir (rw)
      /host/opt/cni/bin from cni-bin-dir (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-mpznt (ro)
  flexvol-driver:
    Container ID:   
    Image:          docker.io/calico/pod2daemon-flexvol:v3.19.1
    Image ID:       
    Port:           <none>
    Host Port:      <none>
    State:          Waiting
      Reason:       PodInitializing
    Ready:          False
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /host/driver from flexvol-driver-host (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-mpznt (ro)
Containers:
  calico-node:
    Container ID:   
    Image:          docker.io/calico/node:v3.19.1
    Image ID:       
    Port:           <none>
    Host Port:      <none>
    State:          Waiting
      Reason:       PodInitializing
    Ready:          False
    Restart Count:  0
    Requests:
      cpu:      250m
    Liveness:   exec [/bin/calico-node -felix-live] delay=10s timeout=1s period=10s #success=1 #failure=6
    Readiness:  exec [/bin/calico-node -felix-ready] delay=0s timeout=1s period=10s #success=1 #failure=3
    Environment Variables from:
      kubernetes-services-endpoint  ConfigMap  Optional: true
    Environment:
      DATASTORE_TYPE:                     kubernetes
      WAIT_FOR_DATASTORE:                 true
      NODENAME:                            (v1:spec.nodeName)
      CALICO_NETWORKING_BACKEND:          <set to the key 'calico_backend' of config map 'calico-config'>  Optional: false
      CLUSTER_TYPE:                       k8s,bgp
      IP:                                 autodetect
      IP_AUTODETECTION_METHOD:            first-found
      CALICO_IPV4POOL_VXLAN:              Always
      FELIX_IPINIPMTU:                    <set to the key 'veth_mtu' of config map 'calico-config'>  Optional: false
      FELIX_VXLANMTU:                     <set to the key 'veth_mtu' of config map 'calico-config'>  Optional: false
      FELIX_WIREGUARDMTU:                 <set to the key 'veth_mtu' of config map 'calico-config'>  Optional: false
      CALICO_IPV4POOL_CIDR:               10.1.0.0/16
      CALICO_DISABLE_FILE_LOGGING:        true
      FELIX_DEFAULTENDPOINTTOHOSTACTION:  ACCEPT
      FELIX_IPV6SUPPORT:                  false
      FELIX_LOGSEVERITYSCREEN:            error
      FELIX_HEALTHENABLED:                true
    Mounts:
      /lib/modules from lib-modules (ro)
      /run/xtables.lock from xtables-lock (rw)
      /var/lib/calico from var-lib-calico (rw)
      /var/run/calico from var-run-calico (rw)
      /var/run/nodeagent from policysync (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-mpznt (ro)
Conditions:
  Type              Status
  Initialized       False 
  Ready             False 
  ContainersReady   False 
  PodScheduled      True 
Volumes:
  lib-modules:
    Type:          HostPath (bare host directory volume)
    Path:          /lib/modules
    HostPathType:  
  var-run-calico:
    Type:          HostPath (bare host directory volume)
    Path:          /var/snap/microk8s/current/var/run/calico
    HostPathType:  
  var-lib-calico:
    Type:          HostPath (bare host directory volume)
    Path:          /var/snap/microk8s/current/var/lib/calico
    HostPathType:  
  xtables-lock:
    Type:          HostPath (bare host directory volume)
    Path:          /run/xtables.lock
    HostPathType:  FileOrCreate
  cni-bin-dir:
    Type:          HostPath (bare host directory volume)
    Path:          /var/snap/microk8s/current/opt/cni/bin
    HostPathType:  
  cni-net-dir:
    Type:          HostPath (bare host directory volume)
    Path:          /var/snap/microk8s/current/args/cni-network
    HostPathType:  
  host-local-net-dir:
    Type:          HostPath (bare host directory volume)
    Path:          /var/snap/microk8s/current/var/lib/cni/networks
    HostPathType:  
  policysync:
    Type:          HostPath (bare host directory volume)
    Path:          /var/snap/microk8s/current/var/run/nodeagent
    HostPathType:  DirectoryOrCreate
  flexvol-driver-host:
    Type:          HostPath (bare host directory volume)
    Path:          /usr/libexec/kubernetes/kubelet-plugins/volume/exec/nodeagent~uds
    HostPathType:  DirectoryOrCreate
  kube-api-access-mpznt:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   Burstable
Node-Selectors:              kubernetes.io/os=linux
Tolerations:                 :NoSchedule op=Exists
                             :NoExecute op=Exists
                             CriticalAddonsOnly op=Exists
                             node.kubernetes.io/disk-pressure:NoSchedule op=Exists
                             node.kubernetes.io/memory-pressure:NoSchedule op=Exists
                             node.kubernetes.io/network-unavailable:NoSchedule op=Exists
                             node.kubernetes.io/not-ready:NoExecute op=Exists
                             node.kubernetes.io/pid-pressure:NoSchedule op=Exists
                             node.kubernetes.io/unreachable:NoExecute op=Exists
                             node.kubernetes.io/unschedulable:NoSchedule op=Exists
Events:
  Type     Reason                  Age                  From               Message
  ----     ------                  ----                 ----               -------
  Normal   Scheduled               8m25s                default-scheduler  Successfully assigned kube-system/calico-node-kxs9s to master
  Warning  FailedCreatePodSandBox  7m56s                kubelet            Failed to create pod sandbox: rpc error: code = Unknown desc = failed to get sandbox image "k8s.gcr.io/pause:3.1": failed to pull image "k8s.gcr.io/pause:3.1": failed to pull and unpack image "k8s.gcr.io/pause:3.1": failed to resolve reference "k8s.gcr.io/pause:3.1": failed to do request: Head "https://k8s.gcr.io/v2/pause/manifests/3.1": dial tcp 74.125.23.82:443: i/o timeout
  Warning  FailedCreatePodSandBox  5m27s                kubelet            Failed to create pod sandbox: rpc error: code = Unknown desc = failed to get sandbox image "k8s.gcr.io/pause:3.1": failed to pull image "k8s.gcr.io/pause:3.1": failed to pull and unpack image "k8s.gcr.io/pause:3.1": failed to resolve reference "k8s.gcr.io/pause:3.1": failed to do request: Head "https://k8s.gcr.io/v2/pause/manifests/3.1": dial tcp 74.125.23.82:443: i/o timeout
  Warning  FailedCreatePodSandBox  4m55s                kubelet            Failed to create pod sandbox: rpc error: code = Unavailable desc = error reading from server: EOF
  Warning  FailedCreatePodSandBox  17s (x6 over 3m57s)  kubelet            Failed to create pod sandbox: rpc error: code = Unknown desc = failed to get sandbox image "k8s.gcr.io/pause:3.1": failed to pull image "k8s.gcr.io/pause:3.1": failed to pull and unpack image "k8s.gcr.io/pause:3.1": failed to resolve reference "k8s.gcr.io/pause:3.1": failed to do request: Head "https://k8s.gcr.io/v2/pause/manifests/3.1": dial tcp 74.125.23.82:443: i/o timeout

```

```shell
ubuntu@master:~$ kubectl config view
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: DATA+OMITTED
    server: https://127.0.0.1:16443
  name: microk8s-cluster
contexts:
- context:
    cluster: microk8s-cluster
    user: admin
  name: microk8s
current-context: microk8s
kind: Config
preferences: {}
users:
- name: admin
  user:
    token: REDACTED


ubuntu@master:~$ kubectl config get-contexts
CURRENT   NAME       CLUSTER            AUTHINFO   NAMESPACE
*         microk8s   microk8s-cluster   admin      

13628


```