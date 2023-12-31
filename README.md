﻿# GOScanner

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

6. 扫描系统信息

7. 图片拼接
   - 支持图片拼接
   - 图片路径和拼接次序自定义
   - 自定义宽度
   - 支持jepg格式图片拼接

8. 视频帧率修改
   - 支持视频帧率修改
   - 支持修改为指定帧率
   - 支持输出到指定目录
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

4. 单一文件加密/解密器
   - 支持AES加密
   - 支持AES解密
   - 支持自定义密钥
   - 支持文件级别的加解密

5. 文件压缩/解压缩器
   - 支持tar-gzip压缩
   - 并发压缩

6. 文件MD5码计算器
   - 指定文件路径
   - 内存占用小

7. 系统信息器
   - 查询CPU,内存,硬盘,主机,系统负载

8. 图片拼接器
   - 支持图片拼接
   - 图片路径和拼接次序自定义
   - 自定义宽度
   - 支持jepg格式图片拼接

9. 视频帧率修改器
   - 支持视频帧率修改
   - 支持修改为指定帧率
   - **注意：本地需要安装ffmpeg**
## 构建过程
```shell
git clone https://github.com/HE-DE/GOScanner.git
cd GOScanner
mkdir bin
go build -o "./bin/gs.exe" .
cd bin
./gs --help
```

## 相关命令

| 命令 | 说明                        |
| ---- | --------------------------- |
| help | 命令帮助提示                |
| ps   | ping扫描网段下的存活IP      |
| sc   | 扫描端口                    |
| ms   | 扫描mysql数据库中表的元信息 |
| fe   | 文件加密                    |
| fd   | 文件解密                    |
| tg   | 打包压缩                    |
| ug   | 解包解压缩                  |
| md5  | 计算文件的MD5码             |
| sys  | 扫描系统信息                |
| is   | 图片拼接                    |
| vc   | 视频帧率修改                  |
