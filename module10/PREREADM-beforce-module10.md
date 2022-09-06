---

---

# [cncamp04](https://github.com/realpeiqi/cncamp04/)

## [Module10](https://github.com/realpeiqi/cncamp04/tree/main/moudle10)

作业要求：

------

- 1.为 HTTPServer 添加 0-2 秒的随机延时；
- 2.为 HTTPServer 项目添加延时 Metric；
- 3.将 HTTPServer 部署至测试集群，并完成 Prometheus 配置；
- 4.从 Promethus 界面中查询延时指标数据；
- 5.（可选）创建一个 Grafana Dashboard 展现延时分配情况。

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

------



作业：

1	创建deployment

```sh
kubectl apply -f deployment.yaml
```

2	创建创建configmap

```sh
kubectl apply -f portconfigmap.yaml
```

3	创建创建service 

```sh
kubectl apply -f service.yaml
```

4	查看pod

```sh
root@master:~/cncamp04/module10# kubectl get pod 
NAME                          READY   STATUS    RESTARTS   AGE
httpserver-66ddbb4d56-99xq8   1/1     Running   0          4h35m
httpserver-66ddbb4d56-bsms9   1/1     Running   0          4h35m
root@master:~/cncamp04/module10#
```

5	查看endpoints

```sh
root@master:~/cncamp04/module10# kubectl get ep
NAME         ENDPOINTS                                 AGE
httpsvc      192.168.219.93:8080,192.168.219.94:8080   4h5m
```

6	查看service

```sh
root@master:~/cncamp04/module10# kubectl get svc
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
httpsvc      ClusterIP   10.110.26.171   <none>        80/TCP    4h5m
```

7	测试svc

```sh
root@master:~/cncamp04/module10# curl 10.110.26.171/healthz
working
root@master:~/cncamp04/module10# curl 10.110.26.171/images
<h1>540<h1>
```

8	helm安装ingress

```sh
root@master:~#helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
root@master:~#helm install ingress-nginx ingress-nginx/ingress-nginx --create-namespace --namespace ingress
```

9	查看ingress之pod情况

```sh
root@master:~# kubectl get pod -n ingress-nginx
NAME                                      READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create--1-bz72g   0/1     Completed   0          3h25m
ingress-nginx-admission-patch--1-xdth8    0/1     Completed   0          3h25m
ingress-nginx-controller-8m5gg            1/1     Running     0          3h16m
ingress-nginx-controller-l5hdf            1/1     Running     0          3h25m
ingress-nginx-controller-rzbsz            1/1     Running     0          3h25m
```

10	查看ingress之svc情况(EXTERNAL-IP 为<pending> 状态，需要修改ingress-nginx-controller配置，手动添加nodeIP地址)

``` 
root@master:~/cncamp04/module10# kubectl get svc  -n ingress
NAME                                 TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
ingress-nginx-controller             LoadBalancer   10.103.54.246   <pending>     80:31784/TCP,443:31854/TCP   101s
ingress-nginx-controller-admission   ClusterIP      10.97.25.245    <none>        443/TCP                      101s
```

11	创建http-ingress

```sh
kubectl apply -f http-ingress.yaml
```



12 查看ingress(ADDRESS 为空，需要修改ingress-nginx-controller配置，手动添加nodeIP地址)

```sh
root@master:~/cncamp04/module10# kubectl get ingress httpserver
NAME      CLASS   HOSTS               ADDRESS   PORTS   AGE
ingress   nginx   www.fengwei.space             80      12m
```



13	创建https-ingress(首先先删除http-ingress  kubectl delete -f http-ingress.yaml)

```sh
kubectl apply -f https-ingress.yaml

```



14 查看ingress(ADDRESS 为空，需要修改ingress-nginx-controller配置，手动添加nodeIP地址)

```sh
root@master:~/cncamp04/module10# kubectl get ingress httpserver
NAME                        CLASS    HOSTS                      ADDRESS   PORTS     AGE
cm-acme-http-solver-8sqtb   <none>   httpserver.femgwei.space             80        15m
httpserver                  nginx    httpserver.femgwei.space             80, 443   20m
```



15 便捷ingress-nginx-controller，添加 externalIPs参数和他们的内容，即内网IP地址和公网IP地址，也指添加内网地址IP或者公网IP地址

``` 
kubectl edit  svc ingress-nginx-controller -n ingress-nginx
##############################
spec:
  allocateLoadBalancerNodePorts: true
  clusterIP: 10.103.54.246
  clusterIPs:
  - 10.103.54.246
  externalIPs: # 添加一下两个IP地址，一个是内网地址，一个是公网地址
  - 172.20.36.160  
  - 8.218.13.225
  externalTrafficPolicy: Cluster
  internalTrafficPolicy: Cluster
```



16 修改ingress-nginx-controller之后再次查看https-ingress (ADDRESS,获取到IP地址)

```
root@master:~# kubectl get ingress 
NAME         CLASS   HOSTS                      ADDRESS                      PORTS     AGE
httpserver   nginx   httpserver.fengwei.space   172.20.36.160,8.218.13.225   80, 443   154m
```



17  修改ingress-nginx-controller之后再次查看ingress之svc情况(EXTERNAL-IP 获取到IP地址: 172.20.36.160,8.218.13.225)

``` 
root@master:~# kubectl get svc -n ingress
NAME                                 TYPE           CLUSTER-IP      EXTERNAL-IP                  PORT(S)                      AGE
ingress-nginx-controller             LoadBalancer   10.103.54.246   172.20.36.160,8.218.13.225   80:31784/TCP,443:31854/TCP   4h57m
ingress-nginx-controller-admission   ClusterIP      10.97.25.245    <none>                       443/TCP                      4h57m
```

18 install cert-manager

```
helm repo add jetstack https://charts.jetstack.io
helm repo update

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.0/cert-manager.crds.yaml

helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.7.1
```



19 签发证书CA的配置 （即创建issuer）

```  sh
kubectl apply -f issuer.yaml
```

20 查看issuer

``` sh
root@master:~# kubectl get cert 
NAME         READY   SECRET       AGE
httpserver   True    httpserver   164m
```

21 查看CertificateRequest 

```
root@master:~# kubectl get CertificateRequest
NAME               APPROVED   DENIED   READY   ISSUER             REQUESTOR                                         AGE
httpserver-cfdwx   True                True    letsencrypt-prod   system:serviceaccount:cert-manager:cert-manager   163m
```



22	重建https-ingress

```  sh
kubectl delete -f https-ingress.yaml && kubectl apply -f https-ingress.yaml
```

23	查看ingress

```  sh
root@k8s-master1:~/cncamp04/module10# kubectl get ingress
NAME                        CLASS    HOSTS             ADDRESS                 PORTS     AGE
cm-acme-http-solver-vp6xp   <none>   www.fengwei.com   10.0.0.141,10.0.0.142   80        5m
https-ingress               nginx    www.fengwei.com   10.0.0.141,10.0.0.142   80, 443   5m9s
```

24	查看secret

``` sh
root@k8s-master1:~/cncamp04/module10# kubectl get secret 
NAME                  TYPE                                  DATA   AGE
default-token-n7gdd   kubernetes.io/service-account-token   3      8h
fengwei-tls-r8svq     Opaque                                1      8m58s
letsencrypt-prod      Opaque                                1      13
```



25	测试 证书是否生效(在浏览器上查看证书，这书信息正常)

``` sh
root@master:~# curl  https://httpserver.fengwei.space/healthz
working
root@master:~# curl  https://httpserver.fengwei.space/images
<h1>59<h1>]
root@master:~# curl  https://httpserver.fengwei.space/images
<h1>511<h1>
```

