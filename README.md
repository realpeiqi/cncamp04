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

------


## [Module3](https://github.com/realpeiqi/cncamp04/tree/main/moudle3)

构建本地镜像 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化 将镜像推送至 docker 官方镜像仓库 通过 docker 命令本地启动 httpserver 通过 nsenter 进入容器查看 

------

## [Module2](https://github.com/realpeiqi/cncamp04/tree/main/moudle2)

编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完： 接收客户端 request，并将 request 中带的 header 写入 response header 读取当前系统的环境变量中的 VERSION 配置，并写入 response header Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出 当访问 localhost/healthz 时，应返回 200

------

## [Module1](https://github.com/realpeiqi/cncamp04/tree/main/moudle1)

基于 Channel 编写一个简单的单线程生产者消费者模型： 队列： 队列长度 10，队列元素类型为 int 生产者： 每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞 消费者： 每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
