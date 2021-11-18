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

网络连通

```shell
aboubakar@ismael:~$ sudo docker network --help

Usage:  docker network COMMAND  [OPTIONS] NETWORK CONTAINER

Manage networks

Commands:
  connect     Connect a container to a network
  create      Create a network
  disconnect  Disconnect a container from a network
  inspect     Display detailed information on one or more networks
  ls          List networks
  prune       Remove all unused networks
  rm          Remove one or more networks

aboubakar@ismael:~$ sudo docker network inspect mynet
"Containers": {
            "5c100e6f67ef9330453c1bf25128b846197b3d02f2aa80d6c54a07927cd34aa8": {
                "Name": "myubuntu01",
                "EndpointID": "65865842d40d3c67983adf64eb4c2407b7ff5a0f807ba31412b3a1799fe63fe1",
                "MacAddress": "02:42:c4:c6:00:02",
                "IPv4Address": "196.198.0.2/16",
                "IPv6Address": ""
            },
            "c1b397d3feca48cb5ea0406d5c72c78cef19e825325379f8bfbdd1896aefedd0": {
                "Name": "myubuntu02",
                "EndpointID": "2bd6b8ff2c9ebf163d71101506a27c92de2996a0ae59760e43ee8af9efe4826c",
                "MacAddress": "02:42:c4:c6:00:03",
                "IPv4Address": "196.198.0.3/16",
                "IPv6Address": ""
            }
        },

```

redis 集群

```shell
aboubakar@ismael:~$ sudo docker network create redis-net --subnet 172.38.0.0/16
3ade65949b3e41d4291198109f1b7d41ebf26e00ed920d89ab8174be81b6af67
aboubakar@ismael:~$ sudo docker network ls
NETWORK ID     NAME                                          DRIVER    SCOPE
9c58db000380   bridge                                        bridge    local
f180033fac2f   docker_gwbridge                               bridge    local
f03013318046   flask-app                                     bridge    local
efb50ee870e6   host                                          host      local
801625b712bd   mongo-net                                     bridge    local
b02cb5e6f62e   mongo-network                                 bridge    local
0824ebb88b73   mynet                                         bridge    local
d007db4b0977   none                                          null      local
2d51b88ea376   rabbitmq_default                              bridge    local
3ade65949b3e   redis-net                                     bridge    local
aboubakar@ismael:~$ sudo docker network inspect redis-net
[
    {
        "Name": "redis-net",
        "Id": "3ade65949b3e41d4291198109f1b7d41ebf26e00ed920d89ab8174be81b6af67",
        "Created": "2021-11-18T21:28:08.188864747+08:00",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "172.38.0.0/16"
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
for port in $(seq 1 6); \
do \
mkdir -p /mydata/redis-net/node-${port}/conf
touch /mydata/redis-net/node-${port}/conf/redis-net.conf
cat << EOF >/mydata/redis-net/node-${port}/conf/redis-net.conf
port 6379
bind 0.0.0.0
cluster-enabled yes 
cluster-config-file nodes.conf
cluster-node-timeout 5000
cluster-announce-ip 172.38.0.1${port}
cluster-announce-port 6379
cluster-announce-bus-port 16379
appendonly yes
EOF
done

root@ismael:/mydata/redis-net# ls
node-1  node-2  node-3  node-4  node-5  node-6
root@ismael:/mydata/redis-net# cd node-1
root@ismael:/mydata/redis-net/node-1# ls
conf
root@ismael:/mydata/redis-net/node-1# cd conf/
root@ismael:/mydata/redis-net/node-1/conf# ls
redis-net.conf
root@ismael:/mydata/redis-net/node-1/conf# cat redis-net.conf
port 6379
bind 0.0.0.0
cluster-enabled yes 
cluster-config-file nodes.conf
cluster-node-timeout 5000
cluster-announce-ip 172.38.0.11
cluster-announce-port 6379
cluster-announce-bus-port 16379
appendonly yes



docker run -p 6371:6379 -p 16371:16379 --name redis-1 \
-v /mydata/redis-net/node-1/data:/data \
-v /mydata/redis-net/node-1/conf/redis-net.conf:/etc/redis-net/redis.conf \
-d --net redis-net --ip 172.38.0.11 redis:latest redis-server /etc/redis-net/redis-net.conf

docker run -p 6372:6379 -p 16372:16379 --name redis-2 \
-v /mydata/redis-net/node-1/data:/data \
-v /mydata/redis-net/node-1/conf/redis-net.conf:/etc/redis-net/redis.conf \
-d --net redis-net --ip 172.38.0.12 redis:latest redis-server /etc/redis-net/redis-net.conf

docker run -p 6373:6379 -p 16373:16379 --name redis-3 \
-v /mydata/redis-net/node-1/data:/data \
-v /mydata/redis-net/node-1/conf/redis-net.conf:/etc/redis-net/redis.conf \
-d --net redis-net --ip 172.38.0.13 redis:latest redis-server /etc/redis-net/redis-net.conf

docker run -p 6374:6379 -p 16374:16379 --name redis-4 \
-v /mydata/redis-net/node-1/data:/data \
-v /mydata/redis-net/node-1/conf/redis-net.conf:/etc/redis-net/redis.conf \
-d --net redis-net --ip 172.38.0.14 redis:latest redis-server /etc/redis-net/redis-net.conf

docker run -p 6375:6379 -p 16375:16379 --name redis-5 \
-v /mydata/redis-net/node-1/data:/data \
-v /mydata/redis-net/node-1/conf/redis-net.conf:/etc/redis-net/redis.conf \
-d --net redis-net --ip 172.38.0.15 redis:latest redis-server /etc/redis-net/redis-net.conf

docker run -p 6376:6379 -p 16376:16379 --name redis-6 \
-v /mydata/redis-net/node-1/data:/data \
-v /mydata/redis-net/node-1/conf/redis-net.conf:/etc/redis-net/redis.conf \
-d --net redis-net --ip 172.38.0.16 redis:latest redis-server /etc/redis-net/redis-net.conf

```