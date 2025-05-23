# PKI

~~~ shell
go run register.go
go run ra.go
go run ca.go
go run http.go
cd ft
npm run dev
~~~

## 启动后路径说明

<http://localhost:3001/servers> 服务注册中心写入的信息说明
<http://localhost:8082/ca-cert> ca初始化后的根证书
<http://localhost:8082/ca-key> ca生成的密钥

<http://localhost:8080/key> 客户端本地生成的密钥
<http://localhost:8080/csr> 客户端本地生成的csr
<http://localhost:8080/cert> 客户端从ca得到的证书

## 环境要求

在ubuntu环境下，安装OpenSSL与go环境，同时要有gcc编译器

基于OpenSSL的RSA使用的CGO开发，调用OpenSSL在c语言中的库函数

## tree

~~~
pki/
├── core/
│   ├── pki/                  # PKI核心逻辑模块
│   │   ├── cert_manager.go   # 证书管理逻辑
│   │   ├── ca.go             # CA相关逻辑
│   │   └── revocation.go     # 证书吊销逻辑
│   └── crypto/               # 加密工具模块
│       ├── interface.go      # 定义加密工具的接口
│       ├── go_crypto/        # 使用Go标准库实现的加密工具
│       │   └── impl.go       # 具体实现
│       ├── openssl/          # 使用OpenSSL实现的加密工具
│       │   └── impl.go       # 具体实现
│       └── cryptoFactory.go  # 加密工具工厂模式实现
├── config/                   # 配置模块
│   └── config.go             # 系统配置加载与解析
├── utils/                    # 工具模块
│   └── helper.go             # 辅助函数
├── main.go                   # 主程序入口
└── README.md                 # 项目说明文档
~~~

~~~cmd
go run ./register.go
go run ./ca.go
go run ./ra.go
go run ./http.go
~~~

## 项目说明

1. 基于OpenSSL，支持生成自签名证书，并加入额外的信任链信息；
2. 实现证书的生命周期管理，包括颁发、更新、吊销等操作；
3. 开发并集成内部信任链，通过内部CA为自签名证书提供信任链，确保证书可追溯；
4. 开发一个简单的透明日志记录系统，记录所有签发、更新、吊销的证书。、

提高要求：

1. 设计引用更高安全性的签名算法（如ECC、SHA-256、SHA-3等）；
2. 构建一个基于内部CA的多级信任体系，扩展证书的适用范围和可信度；
3. 设计一个自动更新和替换过期自签名证书的系统，避免因证书过期导致的安全风险；
4. 加强证书验证或使用双向SSL/TLS，防范中间人攻击。

## 说明

这里要共享的服务地址只要注册服务发现中心的地址就够了，通过注册服务发现中心的地址就可得到CA与RA的地址
请默认使用`3001`端口来启动服务注册发现中心

## 流程说明

主要有以下4个部分组成

1. 客户端
2. 注册中心（RA）
3. 认证中心（CA）
4. 证书存储

## 典型流程

- 客户端生成密钥对，并创建证书签名请求（CSR）。
- 客户端通过RA的API提交CSR，RA验证用户身份。
- RA将验证通过的CSR转发给CA。
- CA签发证书，并通过API将证书返回给客户端。
- 客户端使用证书进行加密、签名等操作。

![alt text](https://ucc.alicdn.com/pic/developer-ecology/89c9c3f19e994b36bb1ecbe2594fed82.png?x-oss-process=image%2Fresize%2Cw_1400%2Fformat%2Cwebp)

### 两个客户端之间的通信流程

如上，客户端通过ca获取证书

- 步骤1：建立连接

    客户端A向客户端B发起连接请求。

    客户端B响应请求，并准备进行安全通信。

- 步骤2：交换证书

    客户端A将自己的数字证书发送给客户端B。

    客户端B将自己的数字证书发送给客户端A。

- 步骤3：验证证书

    客户端A使用CA的根证书验证客户端B的证书：

    检查证书是否由可信的CA签发。

    检查证书是否在有效期内。

    检查证书是否被吊销（通过CRL或OCSP）。

    客户端B同样验证客户端A的证书。

    如果证书验证失败，通信终止；如果验证成功，继续下一步。

- 步骤4：密钥交换

    客户端A生成一个随机的对称密钥（用于加密通信数据）。

    客户端A使用客户端B的公钥（从客户端B的证书中获取）加密这个对称密钥，并发送给客户端B。

    客户端B使用自己的私钥解密，得到对称密钥。

- 步骤5：加密通信

    客户端A和客户端B使用对称密钥对通信数据进行加密和解密。

    对称加密算法（如AES）通常用于加密通信内容，因为它的效率比非对称加密高。

通信完成后，对称密钥被丢弃，以确保每次会话的密钥唯一性。

如果需要再次通信，重新执行上述流程。
