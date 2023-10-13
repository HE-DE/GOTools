# GOScanner

## 组件
1. Log组件
   - 自定义log组件
   - 提供不同级别的日志等级（DEBUG, INFO, NOTICE, WARNING, ERROR, CRITICAL, FATAL）
   - 不同颜色
   - 提供简单易用的输出接口
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

## 构建过程
```shell
git clone https://github.com/HE-DE/GOScanner.git
cd GOScanner
mkdir bin
go build -o "./bin/gs.exe" .
cd bin
./gs --help
```