## Crypto的工厂模式

通过对`cryptoInterface`的实现，来解耦密钥的生成和加密算法的选择。

通过`cryptoFactory.getCryptoInstance()`来得到`cryptoInterface`的实例。

在工厂内部，通过懒加载的方式来获取`cryptoInterface`的实例。