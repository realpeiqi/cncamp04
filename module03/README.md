* **构建本地镜像**
* **编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化**
* **将镜像推送至 docker 官方镜像仓库**
* **通过 docker 命令本地启动 httpserver**
* **通过 nsenter 进入容器查看 IP 配置**

## Dockerfile

```dockerfile
FROM golang:1.18 AS build
WORKDIR /httpserver/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROM busybox
COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE 8080
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]
```

## Docker build

1. docker build

```bash
sudo docker build -f Dockerfile . -t httpserver:v2.0.1
```

2. docker build 结果

```bash
root@master:~/cncamp04/module03# sudo docker build -f Dockerfile . -t httpserver:v1.0.1
Sending build context to Docker daemon   12.8kB
Step 1/12 : FROM golang:1.18 AS build
1.18: Pulling from library/golang
1671565cc8df: Pull complete 
3e94d13e55e7: Pull complete 
fa9c7528c685: Pull complete 
53ad072f9cd1: Pull complete 
dacb042d4b55: Pull complete 
1b8b7b0817d4: Pull complete 
17ab5c0fa3a1: Pull complete 
Digest: sha256:5540a6a6b3b612c382accc545b3f6702de21e77b15d89ad947116c94b5f42993
Status: Downloaded newer image for golang:1.18
 ---> 06c366130191
Step 2/12 : WORKDIR /httpserver/
 ---> Running in b3c615aa36fe
Removing intermediate container b3c615aa36fe
 ---> bf2312b636ed
Step 3/12 : COPY . .
 ---> d1f60a044520
Step 4/12 : ENV CGO_ENABLED=0
 ---> Running in 9edd2b8b16e7
Removing intermediate container 9edd2b8b16e7
 ---> ea64943e485d
Step 5/12 : ENV GO111MODULE=on
 ---> Running in ddc5aa510a97
Removing intermediate container ddc5aa510a97
 ---> 3190f83bc425
Step 6/12 : ENV GOPROXY=https://goproxy.cn,direct
 ---> Running in 09df3e4bcebb
Removing intermediate container 09df3e4bcebb
 ---> fcb8faaebf22
Step 7/12 : RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go
 ---> Running in 13656e4b03c5
Removing intermediate container 13656e4b03c5
 ---> 72c6139cb709
Step 8/12 : FROM busybox
latest: Pulling from library/busybox
2c39bef88607: Pull complete 
Digest: sha256:20142e89dab967c01765b0aea3be4cec3a5957cc330f061e5503ef6168ae6613
Status: Downloaded newer image for busybox:latest
 ---> c98db043bed9
Step 9/12 : COPY --from=build /httpserver/httpserver /httpserver/httpserver
 ---> 087fd4f78ba2
Step 10/12 : EXPOSE 8080
 ---> Running in ad35a34dbc8d
Removing intermediate container ad35a34dbc8d
 ---> 483a00916cfa
Step 11/12 : WORKDIR /httpserver/
 ---> Running in 6aa58cc43fc4
Removing intermediate container 6aa58cc43fc4
 ---> 480d8b1722e1
Step 12/12 : ENTRYPOINT ["./httpserver"]
 ---> Running in 15304452b594
Removing intermediate container 15304452b594
 ---> cbbecd2596ec
Successfully built cbbecd2596ec
Successfully tagged httpserver:v1.0.1
```

## docker images | grep  httpserver

```bash
REPOSITORY                                      TAG                 IMAGE ID            CREATED             SIZE
httpserver                                      v1.0.1              853abc99be5d        3 minutes ago       8.58MB
```

## docker run

```plain
root@master:~/cncamp04/module03# sudo docker run -d httpserver:v1.0.1
992117b7b664369b838cfdaa42389196398b75e60e330478c1518d3e39b6d2e2
```

## docker ps

```bash
CONTAINER ID        IMAGE                                                 COMMAND                  CREATED             STATUS              PORTS                    NAMES
992117b7b664        httpserver:v1.0.1                                     "./httpserver"           43 seconds ago      Up 42 seconds       8082/tcp                 suspicious_euclid
```

## lsns -t net

```plain
root@master:~/cncamp04/module03# lsns -t net
        NS TYPE NPROCS   PID USER    COMMAND
4026531956 net     104     1 root    /usr/lib/systemd/systemd --switched-root --system --deserialize 22
4026532170 net       1 28554 admin   /usr/share/kibana/bin/../node/bin/node --no-warnings --max-http-header-size=65536 /usr/share/kibana/bin
4026532231 net       1 28594 admin   /docker-java-home/bin/java -Duser.dir=/opt/cerebro -Dhosts.0.host=http://elasticsearch:9200 -cp  -jar /
4026532292 net       3 29126 polkitd redis-server *:637
4026532353 net       2 28621 admin   /usr/share/elasticsearch/jdk/bin/java -Xms1g -Xmx1g -XX:+UseConcMarkSweepGC -XX:CMSInitiatingOccupancyF
4026532416 net       3 21409 root    nginx: master process nginx -g daemon off
4026532479 net       2  7627 root    /bin/sh -c /httpserver
4026532605 net       1  3196 root    ./httpserver
```

## nsenter -t 3196 -n ip a

```plain
root@master:~/cncamp04/module03# nsenter -t 3196 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
207: eth0@if208: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:06 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.6/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```

## sudo docker exec -it 992117b7b664 ip a

```plain
root@master:~/cncamp04/module03# docker exec -it 992117b7b664 ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
207: eth0@if208: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue
    link/ether 02:42:ac:11:00:06 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.6/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```

## docker tag

```bash
sudo docker tag httpserver:v1.0.1 dynine/httpserver:v1.0.1

REPOSITORY                                      TAG                 IMAGE ID            CREATED             SIZE
httpserver                                      v1.0.1              853abc99be5d        30 minutes ago      8.58MB
dynine/httpserver                              v1.0.1              853abc99be5d        30 minutes ago      8.58MB
```

## docker push

```bash
sudo docker push dynine/httpserver:v1.0.1
The push refers to repository [docker.io/dynine/httpserver]
214a69e5c125: Pushed
7ad00cd55506: Pushed
v1.0.1: digest: sha256:2932824ab46fdf6a7bce180b622c1bbc3db2a6b9f0252c9e35ec0739a7fb8e17 size: 738

```

#### 测试httpserver容器是否正常

##### 测试本地

```sh
root@master:~/cncamp04/module03# curl 172.17.0.2:8080/healthz
OK,working
```

#### 新建一个名称为web的容器做端口转发做测试通过浏览器测试

```sh
docker container run -d --rm --name web -p 8080:8080 httpserver:v1.0.1
```

## docker kill container_id

```bash
sudo docker kill f56c4993d1ea
```


