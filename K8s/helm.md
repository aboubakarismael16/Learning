### 安装

我们可以在Helm Realese页面下载二进制文件，这里下载的v2.10.0版本，解压后将可执行文件helm拷贝到/usr/local/bin目录下即可，这样Helm客户端就在这台机器上安装完成了。

现在我们可以使用Helm命令查看版本了，会提示无法连接到服务端Tiller：

```shell
$ helm version
Client: &version.Version{SemVer:"v2.10.0", GitCommit:"9ad53aac42165a5fadc6c87be0dea6b115f93090", GitTreeState:"clean"}
Error: could not find tiller
```

要安装 Helm 的服务端程序，我们需要使用到kubectl工具，所以先确保kubectl工具能够正常的访问 kubernetes 集群的apiserver哦。

然后我们在命令行中执行初始化操作：
```shell
$ helm init
```

由于 Helm 默认会去gcr.io拉取镜像，所以如果你当前执行的机器没有配置科学上网的话可以实现下面的命令代替：

```shell
$ helm init --upgrade --tiller-image cnych/tiller:v2.10.0
$HELM_HOME has been configured at /root/.helm.

Tiller (the Helm server-side component) has been installed into your Kubernetes Cluster.

Please note: by default, Tiller is deployed with an insecure 'allow unauthenticated users' policy.
To prevent this, run `helm init` with the --tiller-tls-verify flag.
For more information on securing your installation see: https://docs.helm.sh/using_helm/#securing-your-helm-installation
Happy Helming!
```

如果一直卡住或者报 google api 之类的错误，可以使用下面的命令进行初始化：

```shell
$ helm init --upgrade --tiller-image cnych/tiller:v2.10.0 --stable-repo-url https://cnych.github.io/kube-charts-mirror/

```

这个命令会把默认的 google 的仓库地址替换成我同步的一个镜像地址。

如果在安装过程中遇到了一些其他问题，比如初始化的时候出现了如下错误：

```shell
E0125 14:03:19.093131   56246 portforward.go:331] an error occurred forwarding 55943 -> 44134: error forwarding port 44134 to pod d01941068c9dfea1c9e46127578994d1cf8bc34c971ff109dc6faa4c05043a6e, uid : unable to do port forwarding: socat not found.
2018/01/25 14:03:19 (0xc420476210) (0xc4203ae1e0) Stream removed, broadcasting: 3
2018/01/25 14:03:19 (0xc4203ae1e0) (3) Writing data frame
2018/01/25 14:03:19 (0xc420476210) (0xc4200c3900) Create stream
2018/01/25 14:03:19 (0xc420476210) (0xc4200c3900) Stream added, broadcasting: 5
Error: cannot connect to Tiller

```
解决方案：在节点上安装socat可以解决

```shell
$ sudo snap install  socat
```

Helm 服务端正常安装完成后，Tiller默认被部署在kubernetes集群的kube-system命名空间下：

```shell
$ kubectl get pod -n kube-system -l app=helm
NAME                             READY     STATUS    RESTARTS   AGE
tiller-deploy-86b844d8c6-44fpq   1/1       Running   0          7m
```

此时，我们查看 Helm 版本就都正常了：

```shell
$ helm version
Client: &version.Version{SemVer:"v2.10.0", GitCommit:"9ad53aac42165a5fadc6c87be0dea6b115f93090", GitTreeState:"clean"}
Server: &version.Version{SemVer:"v2.10.0", GitCommit:"9ad53aac42165a5fadc6c87be0dea6b115f93090", GitTreeState:"clean"}
```
另外一个值得注意的问题是RBAC，我们的 kubernetes 集群是1.10.0版本的，默认开启了RBAC访问控制，所以我们需要为Tiller创建一个ServiceAccount，让他拥有执行的权限，详细内容可以查看 Helm 文档中的Role-based Access Control。 创建rbac.yaml文件：

```shell
apiVersion: v1
kind: ServiceAccount
metadata:
  name: tiller
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: tiller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: tiller
    namespace: kube-system
```

然后使用kubectl创建：

```shell
$ kubectl create -f rbac-config.yaml
serviceaccount "tiller" created
clusterrolebinding.rbac.authorization.k8s.io "tiller" created

```

创建了tiller的 ServceAccount 后还没完，因为我们的 Tiller 之前已经就部署成功了，而且是没有指定 ServiceAccount 的，所以我们需要给 Tiller 打上一个 ServiceAccount 的补丁：

```shell
$ kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'

```
NB:上面这一步非常重要，不然后面在使用 Helm 的过程中可能出现Error: no available release name found的错误信息。

至此, Helm客户端和服务端都配置完成了，接下来我们看看如何使用吧。

使用

我们现在了尝试创建一个 Chart：

```shell
$ helm create hello-helm
Creating hello-helm
$ tree hello-helm
hello-helm
├── charts
├── Chart.yaml
├── templates
│   ├── deployment.yaml
│   ├── _helpers.tpl
│   ├── ingress.yaml
│   ├── NOTES.txt
│   └── service.yaml
└── values.yaml

2 directories, 7 files
```

我们通过查看templates目录下面的deployment.yaml文件可以看出默认创建的 Chart 是一个 nginx 服务，具体的每个文件是干什么用的，我们可以前往 Helm 官方文档进行查看，后面会和大家详细讲解的。比如这里我们来安装 1.7.9 这个版本的 nginx，则我们更改 value.yaml 文件下面的 image tag 即可，将默认的 stable 更改为 1.7.9，为了测试方便，我们把 Service 的类型也改成 NodePort

```shell
...
image:
  repository: nginx
  tag: 1.7.9
  pullPolicy: IfNotPresent

nameOverride: ""
fullnameOverride: ""

service:
  type: NodePort
  port: 80
...

```

现在我们来尝试安装下这个 Chart :

```shell
$ helm install ./hello-helm
NAME:   iced-ferret
LAST DEPLOYED: Thu Aug 30 23:39:45 2018
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/Service
NAME                    TYPE       CLUSTER-IP     EXTERNAL-IP  PORT(S)  AGE
iced-ferret-hello-helm  ClusterIP  10.100.118.77  <none>       80/TCP   0s

==> v1beta2/Deployment
NAME                    DESIRED  CURRENT  UP-TO-DATE  AVAILABLE  AGE
iced-ferret-hello-helm  1        0        0           0          0s

==> v1/Pod(related)
NAME                                     READY  STATUS   RESTARTS  AGE
iced-ferret-hello-helm-58cb69d5bb-s9f2m  0/1    Pending  0         0s


NOTES:
1. Get the application URL by running these commands:
  export POD_NAME=$(kubectl get pods --namespace default -l "app=hello-helm,release=iced-ferret" -o jsonpath="{.items[0].metadata.name}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl port-forward $POD_NAME 8080:80

$ kubectl get pods -l app=hello-helm
NAME                                      READY     STATUS    RESTARTS   AGE
iced-ferret-hello-helm-58cb69d5bb-s9f2m   1/1       Running   0          2m
$ kubectl get svc -l app=hello-helm
NAME                       TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
iced-ferret-hello-helm   NodePort   10.104.127.141   <none>        80:31236/TCP   3m
```

查看release：

```shell
$ helm list
NAME             REVISION    UPDATED                     STATUS      CHART               APP VERSION    NAMESPACE
winning-zebra    1           Thu Aug 30 23:50:29 2018    DEPLOYED    hello-helm-0.1.0    1.0            default
```

打包chart：

```shell
$ helm package hello-helm
Successfully packaged chart and saved it to: /root/course/kubeadm/helm/hello-helm-0.1.0.tgz
```

然后我们就可以将打包的tgz文件分发到任意的服务器上，通过helm fetch就可以获取到该 Chart 了。

删除release：

```shell
$ helm delete winning-zebra
release "winning-zebra" deleted
```

然后我们看到kubernetes集群上的该 nginx 服务也已经被删除了。

```shell
$ kubectl get pods -l app=hello-helm
No resources found.
```