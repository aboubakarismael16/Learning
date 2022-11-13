  cat > k8s-1.24.2-master.sh << 'eof'  #!/bin/bash  startTime=`date +%Y%m%d-%H:%M:%S`  startTime_s=`date +%s`  
  ​  # 环境准备  
  ​  # 1、关闭防火墙  ufw disable  
  ​  # 2、关闭swap  swapoff -a  # 临时关闭  sed -i 's/.*swap.*/#&/g' /etc/fstab # 永久关闭  
  ​  # 3、加载IPVS模块  apt install ipset ipvsadm -y  
  ​  cat > /etc/modules-load.d/ipvs.conf << EOF  modprobe -- ip_vs  modprobe -- ip_vs_rr  modprobe -- ip_vs_wrr  modprobe -- ip_vs_sh  modprobe -- nf_conntrack  EOF  
  ​  modprobe -- ip_vs  modprobe -- ip_vs_rr  modprobe -- ip_vs_wrr  modprobe -- ip_vs_sh  
  ​  kernel_version=$(uname -r | cut -d- -f1)  echo $kernel_version  
  ​  if [ `expr $kernel_version \> 4.19` -eq 1 ]      then          modprobe -- nf_conntrack      else          modprobe -- nf_conntrack_ipv4  fi  
  ​  bash /etc/modules-load.d/ipvs.conf && lsmod | grep -e ip_vs -e nf_conntrack  
  ​  # 4、安装container  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg  
  ​  echo \    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://mirrors.tuna.tsinghua.edu.cn/docker-ce/linux/ubuntu \    $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null      apt-get update  
  ​  apt-get install containerd.io=1.6.4-1  
  ​  cat <<EOF | sudo tee /etc/modules-load.d/containerd.conf  overlay  br_netfilter  EOF  ​
  # 1.20+需要开启br_netfilter  sudo modprobe overlay  sudo modprobe br_netfilter  
 ​  cat <<EOF | sudo tee /etc/sysctl.d/99-kubernetes-cri.conf  net.bridge.bridge-nf-call-iptables  = 1  net.ipv4.ip_forward                 = 1  net.bridge.bridge-nf-call-ip6tables = 1  EOF  
 ​  sudo sysctl --system  
 ​  mkdir -p /etc/containerd  containerd config default | sudo tee /etc/containerd/config.toml  
 ​  # 修改cgroup Driver为systemd  sed -ri 's#SystemdCgroup = false#SystemdCgroup = true#' /etc/containerd/config.toml  
 ​  # 更改sandbox_image  sed -ri 's#k8s.gcr.io\/pause:3.6#registry.aliyuncs.com\/google_containers\/pause:3.7#' /etc/containerd/config.toml  
 ​  # endpoint位置添加阿里云的镜像源  sed -ri 's#https:\/\/registry-1.docker.io#https:\/\/registry.aliyuncs.com#' /etc/containerd/config.toml  
 ​  systemctl daemon-reload  
 ​  systemctl restart containerd  
 ​  # 5、安装k8s-1.24.2  curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -  
 ​  cat <<EOF >/etc/apt/sources.list.d/kubernetes.list  deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main  EOF  
 ​  apt-get update  
 ​  apt-get install -y kubelet=1.24.2-00 kubeadm=1.24.2-00 kubectl=1.24.2-00  
 ​  # 设置crictl  cat << EOF >> /etc/crictl.yaml  runtime-endpoint: unix:///var/run/containerd/containerd.sock  image-endpoint: unix:///var/run/containerd/containerd.sock  timeout: 10   debug: false  EOF  
 ​  mkdir ~/kubeadm_init && cd ~/kubeadm_init  
 ​  kubeadm config print init-defaults > kubeadm-init.yaml  
 ​  cat > kubeadm-init.yaml << EOF  apiVersion: kubeadm.k8s.io/v1beta3  bootstrapTokens:  - groups:    - system:bootstrappers:kubeadm:default-node-token    token: abcdef.0123456789abcdef    ttl: 24h0m0s    usages:    - signing    - authentication  kind: InitConfiguration  localAPIEndpoint:    advertiseAddress: `hostname -i`  #master_ip    bindPort: 6443  nodeRegistration:    criSocket: unix:///var/run/containerd/containerd.sock    imagePullPolicy: IfNotPresent    name: master    taints:    - effect: "NoSchedule"      key: "node-role.kubernetes.io/master"  ---  apiServer:    timeoutForControlPlane: 4m0s  apiVersion: kubeadm.k8s.io/v1beta3  certificatesDir: /etc/kubernetes/pki  clusterName: kubernetes  controllerManager: {}  dns: {}  etcd:    local:      dataDir: /var/lib/etcd  imageRepository: registry.aliyuncs.com/google_containers  kind: ClusterConfiguration  kubernetesVersion: v1.24.2  networking:    dnsDomain: cluster.local    serviceSubnet: 10.96.0.0/12    podSubnet: 10.244.0.0/16  scheduler: {}  ---  apiVersion: kubeproxy.config.k8s.io/v1alpha1  kind: KubeProxyConfiguration  mode: ipvs  ---  apiVersion: kubelet.config.k8s.io/v1beta1  kind: KubeletConfiguration  cgroupDriver: systemd  EOF  ​  # 预拉取镜像  kubeadm config images pull --config kubeadm-init.yaml  ​
   # 初始化集群  kubeadm init --config=kubeadm-init.yaml | tee kubeadm-init.log  
 ​  mkdir -p $HOME/.kube  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config  sudo chown $(id -u):$(id -g) $HOME/.kube/config  
 ​  kubectl get node -owide   
 ​  # 脚本执行时间  endTime=`date +%Y%m%d-%H:%M:%S`  endTime_s=`date +%s`  sumTime=$[ $endTime_s - $startTime_s ]  echo "$startTime ---> $endTime" "Total:$sumTime seconds"  eof
