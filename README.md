# [cncamp04](https://github.com/realpeiqi/cncamp04/)
------
## [Module15](https://github.com/realpeiqi/cncamp04/tree/main/module15)

------

## 结业总结

------

- 详见[毕业总结.md文档](https://github.com/realpeiqi/cncamp04/tree/main/module15/毕业总结.md)

------

## 结业总结

------

- 详见[毕业总结.md文档](

------
## [Module12](https://github.com/realpeiqi/cncamp04/tree/main/module12)
  
------

## 作业内容

------
- 把我们的 httpserver 服务以 Istio Ingress Gateway 的形式发布出来。以下是你需要考虑的几点：

1. 如何实现安全保证；
2. 七层路由规则；
3. 考虑 open tracing 的接入。

------

## [Module10](https://github.com/realpeiqi/cncamp04/tree/main/module10)

##作业要求：

------

- 1.为 HTTPServer 添加 0-2 秒的随机延时；
- 2.为 HTTPServer 项目添加延时 Metric；
- 3.将 HTTPServer 部署至测试集群，并完成 Prometheus 配置；
- 4.从 Promethus 界面中查询延时指标数据；
- 5.（可选）创建一个 Grafana Dashboard 展现延时分配情况。

------

## [Module09](https://github.com/realpeiqi/cncamp04/tree/main/module09)

------

## 课后练习 9.1
### 测试对 CPU 的校验和准入行为
- 定义一个 Pod，并将该 Pod 中的 nodeName 属性直接写成集群中的节点名。
- 将 Pod 的 CPU 的资源设置为超出计算节点的 CPU 的值。
- 创建该 Pod。
- 观察行为并思考。

-----

## [Module08](https://github.com/realpeiqi/cncamp04/tree/main/module08)

## 作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 Kubernetes 集群，以下是你可以思考的维度。

------

## 第一部分

- 优雅启动
- 优雅终止
- 资源需求和 QoS 保证
- 探活
- 日常运维需求，日志等级
- 配置和代码分离

------

## 第二部分

- Service
- Ingress
- 如何确保整个应用的高可用
- 如何通过证书保证 httpServer 的通讯安全

------


## [Module03](https://github.com/realpeiqi/cncamp04/tree/main/moudle03)

构建本地镜像 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化 将镜像推送至 docker 官方镜像仓库 通过 docker 命令本地启动 httpserver 通过 nsenter 进入容器查看 

------

## [Module02](https://github.com/realpeiqi/cncamp04/tree/main/moudle02)

编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完： 接收客户端 request，并将 request 中带的 header 写入 response header 读取当前系统的环境变量中的 VERSION 配置，并写入 response header Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出 当访问 localhost/healthz 时，应返回 200

------

## [Module01](https://github.com/realpeiqi/cncamp04/tree/main/moudle01)

基于 Channel 编写一个简单的单线程生产者消费者模型： 队列： 队列长度 10，队列元素类型为 int 生产者： 每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞 消费者： 每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
