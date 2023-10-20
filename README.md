# GOScanner

## 组件
1. Log组件
   - 自定义log组件
   - 提供不同级别的日志等级（DEBUG, INFO, NOTICE, WARNING, ERROR, CRITICAL, FATAL）
   - 不同颜色
   - 提供简单易用的输出接口

2. 命令解释组件
   - 支持命令行参数的自定义
   - 支持命令行参数的解析

3. 进度条组件
   - 支持自定义进度条
   - 支持自定义进度条长度
   - 支持自定义进度条颜色

4. AES加解密组件
   - 支持AES加密
   - 支持AES解密
  
5. 文件MD5值的计算功能
   - 支持文件MD5的计算
## 功能
1. 并发端口扫描器
    - 并发端口扫描器
    - 协程数量自定义
    - 端口范围自定义

2. Ping网段IP扫描器
    - 并发Ping扫描器
    - 网段自定义

3. Mysql数据表元信息扫描器
   - 并发执行
   - 原生SQL语句
   - 指定数据库查询

4. 单一文件加密/解密
   - 支持AES加密
   - 支持AES解密
   - 支持自定义密钥
   - 支持文件级别的加解密

5. 文件压缩/解压缩
   - 支持tar-gzip压缩
   - 并发压缩

6. 文件MD5码计算
   - 指定文件路径
   - 内存占用小
## 构建过程
```shell
git clone https://github.com/HE-DE/GOScanner.git
cd GOScanner
mkdir bin
go build -o "./bin/gs.exe" .
cd bin
./gs --help
```