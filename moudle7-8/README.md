---

---

# [cncamp04](https://github.com/realpeiqi/cncamp04/)

## [Module8](https://github.com/realpeiqi/cncamp04/tree/main/moudle8)

作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 Kubernetes 集群，以下是你可以思考的维度。

------

第一部分

- 优雅启动
- 优雅终止
- 资源需求和 QoS 保证
- 探活
- 日常运维需求，日志等级
- 配置和代码分离

------

第二部分

- Service
- Ingress
- 如何确保整个应用的高可用
- 如何通过证书保证 httpServer 的通讯安全

作业：

1	创建deployment

```sh
kubectl apply -f deployment.yaml
```

2	创建创建configmap

```sh
kubectl apply -f configmap.yaml
```

3	创建创建service 

```
kubectl apply -f service.yaml
```

4	查看pod

```
root@master:~/cncamp04/moudle7-8# kubectl get pod 
NAME                                    READY   STATUS             RESTARTS       AGE
httpserver-5f87c874bd-bm7mb             1/1     Running            0              34m
httpserver-5f87c874bd-drv56             1/1     Running            0              34m
httpserver-5f87c874bd-rq5lz             1/1     Running            0              34m

```

5	查看endpoints

```
root@master:~/cncamp04/moudle7-8# kubectl get ep
NAME          ENDPOINTS                                                       AGE
httpsvc       192.168.104.19:9999,192.168.166.146:9999,192.168.166.147:9999   11m

```

6	查看service

```
root@master:~/cncamp04/moudle7-8# kubectl get svc 
NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
httpsvc      NodePort   10.96.134.80   <none>        80/TCP    16m
```

7	查看curl +NodePort

```sh
root@master:~/cncamp04/moudle7-8# curl 10.96.134.80/healthz
200
```

8	安装ingress

```
kubectl apply -f ingress-controller.yaml
```

9	查看ingress安装情况 

```
root@master:~# kubectl get pod -n ingress-nginx
NAME                                      READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create--1-bz72g   0/1     Completed   0          3h25m
ingress-nginx-admission-patch--1-xdth8    0/1     Completed   0          3h25m
ingress-nginx-controller-8m5gg            1/1     Running     0          3h16m
ingress-nginx-controller-l5hdf            1/1     Running     0          3h25m
ingress-nginx-controller-rzbsz            1/1     Running     0          3h25m
```

10	创建ingress

```
kubectl apply -f ingress.yaml
```

11 查看ingress

```
root@master:~# kubectl get ingress
NAME      CLASS   HOSTS          ADDRESS                            PORTS   AGE
ingress   nginx   www.dada.com   10.0.0.130,10.0.0.131,10.0.0.132   80      3h3m
```

12 测试ingress(做本地hosts解析)

```
root@master:~#curl  http://www.fengwei.com/healthz
200
```

