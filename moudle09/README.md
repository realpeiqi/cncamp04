## [Module09](https://github.com/realpeiqi/cncamp04/tree/main/moudle09)
------
## 课后练习 9.1
### 测试对 CPU 的校验和准入行为
- 定义一个 Pod，并将该 Pod 中的 nodeName 属性直接写成集群中的节点名。
- 将 Pod 的 CPU 的资源设置为超出计算节点的 CPU 的值。
- 创建该 Pod。
- 观察行为并思考。
------
1	创建Pod

```sh
kubectl apply -f pod-nodename.yaml
```
------
2	查看pod
```sh
root@k8s-master1:~/cncamp04/moudle09# kubectl get pod
NAME                          READY   STATUS     RESTARTS   AGE
cm-acme-http-solver-h9h6h     1/1     Running    0          18h
httpserver-5f87c874bd-5bmnw   1/1     Running    0          25h
httpserver-5f87c874bd-dx5fk   1/1     Running    0          25h
httpserver-5f87c874bd-tnqhw   1/1     Running    0          25h
pod-node2name                 0/1     OutOfcpu   0          7m27s
```
------
3	思考
```sh
OutOfcpu :表明cpu个数只有4个。而pod资源需求中request需要5个。
表示cpu资源不满足pod的创建条件。
```
------
