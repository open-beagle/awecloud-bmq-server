# awecloud-bmq-server

多云消息服务

## 端口设计

| 端口 | 协议      | 用途       | URL                    |
| ---- | --------- | ---------- | ---------------------- |
| 8081 | h2c       | grpc 服务  | /awecloud/bmq/registry |
| 8082 | websocket | 消息服务   | /awecloud/bmq/message  |
| 8080 | http      | 控制台服务 | /awecloud/bmq/api      |
