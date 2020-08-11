# 客户端服务端双向证书校验
服务端可以要求对客户端的证书进行校验，以更严格识别客户端的身份，限制客户端的访问。

```
# 查看证书文件
openssl x509 -in server.crt -noout -text
```

# client
- 私钥文件：ca.key
- 数字证书：ca.crt

```
// 生成一个CA私钥
openssl genrsa -out ca.key 2048
// 使用CA私钥生成客户端的数字证书
openssl req -x509 -new -nodes -key ca.key -days 365 -out ca.crt

// 生成客户端私钥
openssl genrsa -out client.key 2048
// 使用客户端私钥生成数字证书
openssl req -new -key client.key -out client.csr
// 创建客户端扩展配置信息
echo "extendedKeyUsage=clientAuth" >> client.ext
// 使用客户端ca私钥和客户端扩展配置信息签发客户端的数字证书
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extfile client.ext -out client.crt -days 365
```


# server
- 私钥文件   server.key
- 数字证书   server.crt

```
cp ca.crt ca.key ./server/

// 生成服务器端的私钥
openssl genrsa -out server.key 2048
// 使用服务器端的私钥生成一个数字证书
openssl req -new -key server.key -out server.csr
// 使用客户端ca私钥签发服务器端的数字证书
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365
```

# 参考
https://blog.csdn.net/zhengzizhi/article/details/73720069
https://blog.csdn.net/zhengzizhi/article/details/73699154
