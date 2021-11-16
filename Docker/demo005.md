Docker NetwWorking 

```shell
aboubakar@ismael:~$ ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: wlp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 3c:f0:11:13:d9:40 brd ff:ff:ff:ff:ff:ff
    inet 192.168.101.23/24 brd 192.168.101.255 scope global dynamic noprefixroute wlp1s0
       valid_lft 84965sec preferred_lft 84965sec
    inet6 fe80::f40b:65a7:b37:ca5/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
3: mpqemubr0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether 52:54:00:57:8b:16 brd ff:ff:ff:ff:ff:ff
    inet 10.218.50.1/24 brd 10.218.50.255 scope global mpqemubr0
       valid_lft forever preferred_lft forever
4: br-2d51b88ea376: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
    link/ether 02:42:43:ee:0f:c0 brd ff:ff:ff:ff:ff:ff
    inet 172.21.0.1/16 brd 172.21.255.255 scope global br-2d51b88ea376
       valid_lft forever preferred_lft forever
5: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
    link/ether 02:42:52:6f:27:01 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever

```

自定义网络

```shell
aboubakar@ismael:~$ sudo docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
18d7aa9c30f4   bridge    bridge    local
9f67d7222dee   host      host      local
23dd5a314f78   none      null      local
```

网络模式：

    -bridge: 桥接 docker 默认
    -host：和宿主机共享网络
    -none: 不配置网络
    -container: 容器内网络连通（用的少），局限很大


```shell
aboubakar@ismael:~$ sudo docker network create --driver bridge --subnet 196.198.0.0/16 --gateway 196.198.0.1 mynet
0824ebb88b732209760e35874ef3c53b7a3dbbf68c4a8d5ab7d4c5c12ae9affa

aboubakar@ismael:~$ sudo docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
18d7aa9c30f4   bridge    bridge    local
9f67d7222dee   host      host      local
23dd5a314f78   none      null      local
0824ebb88b73   mynet     bridge    local


aboubakar@ismael:~$ sudo docker inspect mynet
[
    {
        "Name": "mynet",
        "Id": "0824ebb88b732209760e35874ef3c53b7a3dbbf68c4a8d5ab7d4c5c12ae9affa",
        "Created": "2021-11-16T22:10:52.890149153+08:00",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "196.198.0.0/16",
                    "Gateway": "196.198.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {},
        "Options": {},
        "Labels": {}
    }
]

```

```shell
aboubakar@ismael:~$ sudo docker run -itd --name myubuntu01 --network mynet ubuntu /bin/sh
69c362628b5fe15f0ed2b3a27f596683513fb2f91c60affff57fa1ab5888f862
aboubakar@ismael:~$ sudo docker run -itd --name myubuntu02 --network mynet ubuntu /bin/sh
d384bab45b49bec8ea4e783fab57e3f62edcff34a510ae91e6606659630aa5cd
aboubakar@ismael:~$ sudo docker ps
CONTAINER ID   IMAGE     COMMAND     CREATED          STATUS          PORTS     NAMES
d384bab45b49   ubuntu    "/bin/sh"   11 seconds ago   Up 10 seconds             myubuntu02
69c362628b5f   ubuntu    "/bin/sh"   43 seconds ago   Up 42 seconds             myubuntu01


aboubakar@ismael:~$ sudo docker exec -it myubuntu01 /bin/bash
root@40c30f43fef3:/# apt-get update
root@40c30f43fef3:/# apt install iputils-ping -y

# ping 196.198.0.2
PING 196.198.0.2 (196.198.0.2) 56(84) bytes of data.
64 bytes from 196.198.0.2: icmp_seq=1 ttl=64 time=0.076 ms
64 bytes from 196.198.0.2: icmp_seq=2 ttl=64 time=0.070 ms
64 bytes from 196.198.0.2: icmp_seq=3 ttl=64 time=0.071 ms
64 bytes from 196.198.0.2: icmp_seq=4 ttl=64 time=0.071 ms

# ping myubuntu02
PING myubuntu02 (196.198.0.3) 56(84) bytes of data.
64 bytes from myubuntu02.mynet (196.198.0.3): icmp_seq=1 ttl=64 time=0.189 ms
64 bytes from myubuntu02.mynet (196.198.0.3): icmp_seq=2 ttl=64 time=0.121 ms
64 bytes from myubuntu02.mynet (196.198.0.3): icmp_seq=3 ttl=64 time=0.111 ms
64 bytes from myubuntu02.mynet (196.198.0.3): icmp_seq=4 ttl=64 time=0.111 ms
64 bytes from myubuntu02.mynet (196.198.0.3): icmp_seq=5 ttl=64 time=0.115 ms


aboubakar@ismael:~$ sudo docker exec -it myubuntu01 ping 196.198.0.2
PING 196.198.0.2 (196.198.0.2) 56(84) bytes of data.
64 bytes from 196.198.0.2: icmp_seq=1 ttl=64 time=0.049 ms
64 bytes from 196.198.0.2: icmp_seq=2 ttl=64 time=0.068 ms
64 bytes from 196.198.0.2: icmp_seq=3 ttl=64 time=0.071 ms


aboubakar@ismael:~$ sudo docker exec -ti myubuntu01 ping myubuntu02
PING myubuntu02 (196.198.0.3) 56(84) bytes of data.
64 bytes from myubuntu02.mynet (196.198.0.3): icmp_seq=1 ttl=64 time=0.120 ms
64 bytes from myubuntu02.mynet (196.198.0.3): icmp_seq=2 ttl=64 time=0.119 ms
64 bytes from myubuntu02.mynet (196.198.0.3): icmp_seq=3 ttl=64 time=0.120 ms
64 bytes from myubuntu02.mynet (196.198.0.3): icmp_seq=4 ttl=64 time=0.119 ms


```