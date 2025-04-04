## Crypto的工厂模式

通过对`cryptoInterface`的实现，来解耦密钥的生成和加密算法的选择。

通过`cryptoFactory.getCryptoInstance()`来得到`cryptoInterface`的实例。

在工厂内部，通过懒加载的方式来获取`cryptoInterface`的实例。

## PEM类型

算法|公钥解析函数|公钥 PEM 类型|私钥解析函数|私钥 PEM 类型
-|-|-|-|-
RSA|x509.ParsePKCS1PublicKey|RSA PUBLIC KEY|x509.ParsePKCS1PrivateKey|RSA PRIVATE KEY
ECDSA|x509.ParsePKIXPublicKey|PUBLIC KEY|x509.ParseECPrivateKey|EC PRIVATE KEY
Ed25519|x509.ParsePKIXPublicKey|PUBLIC KEY|x509.ParsePKCS8PrivateKey|PRIVATE KEY
DSA|x509.ParsePKIXPublicKey|PUBLIC KEY|x509.ParsePKCS8PrivateKey|PRIVATE KEY
通用(需要类型断言获取具体类型的密钥)|x509.ParsePKIXPublicKey|PUBLIC KEY|x509.ParsePKCS8PrivateKey|PRIVATE KEY

因为PKI系统更多的专注于公钥的分发，所以这里建议以`PKIX`的格式来解析与序列化公钥，私钥则自行管理。如果想统一管理私钥的话，则使用`PKCS8`规范