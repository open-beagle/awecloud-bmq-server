# SSL

生成非对称RES密钥对，来对关键信息进行加解密保护。前端js加密，后端go解密。

```bash
# pri
openssl genrsa -out ssl/bmq.key 1024

# pub
openssl rsa -pubout -in ssl/bmq.key -out ssl/bmq.crt
```
