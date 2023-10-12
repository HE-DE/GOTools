# GOScanner

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