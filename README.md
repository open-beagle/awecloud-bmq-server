# awecloud-bmq-server

多云消息服务

## 端口设计

| 端口 | 协议      | 用途       | URL                    |
| ---- | --------- | ---------- | ---------------------- |
| 81 | h2c       | grpc 服务  | /awecloud/bmq/registry |
| 82 | websocket | 消息服务   | /awecloud/bmq/message  |
| 83 | http      | 控制台服务 | /awecloud/bmq/api      |

## grpc 服务

1. Login 登录服务

提供服务用于认证客户端，经认证的客户端可以访问服务。

2. Listen 侦听服务

客户端连接后，通过侦听服务接受服务消息，按照服务器的记录.

## 消息服务

## 控制台服务
