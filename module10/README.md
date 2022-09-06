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

### helm安装prometheus

```
#
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
# 安装
helm -n prometheus-stack install  kube-prometheus-stack prometheus-community/kube-prometheus-stack 
```

### 配置prometheus

``` 
root@master:~# kubectl get prometheuses -n prometheus-stack
NAME                               VERSION   REPLICAS   AGE
kube-prometheus-stack-prometheus   v2.37.0   1          9h
```

如果改动的配置不生效，查看opeator log

``` 
root@master:~# kubectl get pod -n prometheus-stack kube-prometheus-stack-operator-7bf84dc5bb-dl8gt
NAME                                              READY   STATUS    RESTARTS   AGE
kube-prometheus-stack-operator-7bf84dc5bb-dl8gt   1/1     Running   0          9h
```



### 创建发现规则

```
见prometheus-additional.yaml文件内容
```

### 存入集群的secret

```
root@master:~/cncamp04/module10# kubectl create secret generic additional-configs --from-file=prometheus-additional.yaml -n  prometheus-stack
secret/additional-configs created
```

### 查看所有secret

``` 
root@master:~/cncamp04/module10# kubectl get secret -n prometheus-stack
NAME                                                           TYPE                                  DATA   AGE
additional-configs                                             Opaque                                1      8m32s
alertmanager-kube-prometheus-stack-alertmanager                Opaque                                1      9h
alertmanager-kube-prometheus-stack-alertmanager-generated      Opaque                                1      9h
alertmanager-kube-prometheus-stack-alertmanager-tls-assets-0   Opaque                                0      9h
alertmanager-kube-prometheus-stack-alertmanager-web-config     Opaque                                1      9h
default-token-smq8m                                            kubernetes.io/service-account-token   3      21h
kube-prometheus-stack-admission                                Opaque                                3      9h
kube-prometheus-stack-alertmanager-token-4mb46                 kubernetes.io/service-account-token   3      9h
kube-prometheus-stack-grafana                                  Opaque                                3      9h
```



### 注入Prometheus

然后我们需要在声明 prometheus 的资源对象文件中通过 additionalScrapeConfigs 属性添加上这个额外的配置：

```
root@master:~# kubectl get prometheuses -n prometheus-stack -oyaml
  ...
  spec:
    additionalScrapeConfigs:			# 未注入内容
      key: prometheus-additional.yaml 	# 未注入内容
      name: additional-configs 			# 未注入内容
    alerting:
      alertmanagers:
   ...
```



#### 理论上pod会重启，这里没有重启，我们通过删除重建pod (达到重启prometheus目的)

```
root@master:~/prometheus# kubectl delete  pod -n prometheus-stack prometheus-kube-prometheus-stack-prometheus-0 
pod "prometheus-kube-prometheus-stack-prometheus-0" deleted
```

### 配置RBAC

```
root@master:~/cncamp04/module10# kubectl create -f clusterrole.yaml 
clusterrole.rbac.authorization.k8s.io/prometheus-k8s created
```

### 在prometheus的svc上打上annoation

```
kind: Service
metadata:
  annotations:  # annoation
    prometheus.io/port: "8080" # annoation
    prometheus.io/scrape: "true" # annoation
  creationTimestamp: "2022-03-12T03:59:18Z"
  labels:
    app: httpserver
```



### 在grafana的secret配置查看用户名和密码

``` 
root@master:~/cncamp04/module10/https-server# kubectl edit secret  -n prometheus-stack   kube-prometheus-stack-grafana

# Please edit the object below. Lines beginning with a '#' will be ignored,
# and an empty file will abort the edit. If an error occurs while saving this file will be
# reopened with the relevant failures.
#
apiVersion: v1
data:
  admin-password: cHJvbS1vcGVyYXRvcg==
  admin-user: YWRtaW4=
  ldap-toml: ""
kind: Secret
metadata:
  annotations:
    meta.helm.sh/release-name: kube-prometheus-stack
    meta.helm.sh/release-namespace: prometheus-stack
  creationTimestamp: "2022-09-05T03:50:21Z"
  labels:
    app.kubernetes.io/instance: kube-prometheus-stack
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: grafana
    app.kubernetes.io/version: 9.0.5
    helm.sh/chart: grafana-6.32.10
  name: kube-prometheus-stack-grafana
  namespace: prometheus-stack
  resourceVersion: "10177"
  uid: e9c72068-5174-4dd3-8499-a74f8b07bbc8
type: Opaque

```

### 解码grafana的secret用户名和密码

``` 
root@master:~/cncamp04/module10/https-server# echo cHJvbS1vcGVyYXRvcg== | base64 -d
prom-operatorroot@master:~/cncamp04/module10/https-server# 
```

### 查看Prometheus原生监控图(也可以修改svc把公网地址加入到externallPs中)

```
kubectl port-forward -n prometheus-stack  --address 0.0.0.0 svc/kube-prometheus-stack-prometheus 9090:9090

###externallPs
  externalIPs: # 添加一下两个IP地址，一个是内网地址，一个是公网地址
  - 172.20.36.160  
  - 8.218.13.225
```



### 服务发现 

``` 
看到kubernetes-endponnts 表明 additional中监控 注入成功
```



### PromQL基本语法

```PromQL
histogram_quantile(0.95,sum(rate(cloudnative_execution_latency_seconds_bucket[5m])) by (le))
```



###  grafana

port forward 到本地

```
kubectl port-forward -n prometheus-stack  --address 0.0.0.0 svc/kube-prometheus-stack-grafana 3000:80
```



### import 选项中 导入json 