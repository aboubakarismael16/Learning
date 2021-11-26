```shell
ubuntu@master:~$ kubectl get ns
NAME                   STATUS   AGE
default                Active   2d1h
kube-node-lease        Active   2d1h
kube-public            Active   2d1h
kube-system            Active   2d1h
kubernetes-dashboard   Active   20h



ubuntu@master:~$ kubectl -n kubernetes-dashboard get all
NAME                                        READY   STATUS    RESTARTS   AGE
pod/kubernetes-dashboard-5c6dff6c6f-v8fvz   1/1     Running   1          20h

NAME                                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/dashboard-metrics-scraper   ClusterIP   10.98.249.56     <none>        8000/TCP   20h
service/kubernetes-dashboard        ClusterIP   10.100.178.150   <none>        443/TCP    20h

NAME                                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/kubernetes-dashboard   1/1     1            1           20h

NAME                                              DESIRED   CURRENT   READY   AGE
replicaset.apps/kubernetes-dashboard-5c6dff6c6f   1         1         1       20h



ubuntu@master:~$ kubectl -n kubernetes-dashboard describe service kubernetes-dashboard
Name:              kubernetes-dashboard
Namespace:         kubernetes-dashboard
Labels:            k8s-app=kubernetes-dashboard
Annotations:       Selector:  k8s-app=kubernetes-dashboard
Type:              ClusterIP
IP:                10.100.178.150
Port:              <unset>  443/TCP
TargetPort:        8443/TCP
Endpoints:         192.168.166.130:8443
Session Affinity:  None
Events:            <none>

ubuntu@master:~/kubernetes/dashboard$ kubectl -n kubernetes-dashboard port-forward kubernetes-dashboard-5c6dff6c6f-v8fvz 8000:8443
Forwarding from 127.0.0.1:8000 -> 8443
Forwarding from [::1]:8000 -> 8443

```