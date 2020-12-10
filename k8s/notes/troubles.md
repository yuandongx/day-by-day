# 1 容器一直处于`ContainerCreating`
2020年12月10日
## 状态
`buntu@ubuntu-master:~/workspace/sample/web$ kubectl  get pods
NAME                     READY   STATUS              RESTARTS   AGE
mysql-86b9547d6c-tkz8t   0/1     ContainerCreating   0          9s
redis-5487f8f5f6-fz82n   0/1     ContainerCreating   0          5m50s
redis-5487f8f5f6-zngqb   0/1     ContainerCreating   0          5m50s
`
## kubectl describe deployment redis
`
ubuntu@ubuntu-master:~/workspace/sample/web$ kubectl describe deployment redis
Name:                   redis
Namespace:              default
CreationTimestamp:      Thu, 10 Dec 2020 09:04:08 +0000
Labels:                 app.kubernetes.io/name=redis
Annotations:            deployment.kubernetes.io/revision: 1
Selector:               app.kubernetes.io/name=redis
Replicas:               2 desired | 2 updated | 2 total | 0 available | 2 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  25% max unavailable, 25% max surge
Pod Template:
  Labels:  app.kubernetes.io/name=redis
  Containers:
   redis:
    Image:        registry.cn-hangzhou.aliyuncs.com/yuandongx/redis
    Port:         6379/TCP
    Host Port:    0/TCP
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      False   MinimumReplicasUnavailable
  Progressing    True    ReplicaSetUpdated
OldReplicaSets:  <none>
NewReplicaSet:   redis-5487f8f5f6 (2/2 replicas created)
Events:
  Type    Reason             Age    From                   Message
  ----    ------             ----   ----                   -------
  Normal  ScalingReplicaSet  2m22s  deployment-controller  Scaled up replica set redis-5487f8f5f6 to 2

`
## 解决问题思路
 - 查看pods状态一直是`ContainerCreating`状态，说明容器没起来，下一步查看 `kubectl describe deployment mysql`
 - 查看不到更多可用的信息，再查看事件 `kubectl get events --all-namespaces  --sort-by='.metadata.creationTimestamp'`
 - 发现 `
 default       28m         Warning   FailedCreatePodSandBox   pod/redis-7bb6944b8f-xxwb7     Failed to create pod sandbox: rpc error: code = Unknown desc = failed to set up sandbox container "ab83b168fe2ee19013778bdb4f1ee0fea4c5efed096834f72f16cce55c7b6a2a" network for pod "redis-7bb6944b8f-xxwb7": networkPlugin cni failed to set up pod "redis-7bb6944b8f-xxwb7_default" network: open /run/flannel/subnet.env: no such file or directory
 `
 - 排查node的 `/run/flannel/subnet.env`,现发工作结点没有这个文件，需要修改
 ## 最终解决
 添加文件 `/run/flannel/subnet.env`
 内容为
 `
 ubuntu@ubuntu-master:/run/flannel$ cat subnet.env
FLANNEL_NETWORK=192.168.96.0/24
FLANNEL_SUBNET=1.0.0.1/24
FLANNEL_MTU=1450
FLANNEL_IPMASQ=true
 `
 
