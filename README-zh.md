# Red Team TL;DR

[English](./index.md) | [中文简体](./README-zh.md)

## 什么是 Red Team TL;DR ?

red-tldr 是一个轻量级的文本搜索工具，它用于帮助红队工作人员快速的寻找到想要执行的命令、关键点，因此它比较适用于有一定经验的红队人员使用。

## 为什么选择 Red Team TL;DR ?

在我的日常工作中，需要记忆很多命令，我多数只知道它的开头，人的记忆是有限的，而通过搜索引擎寻找我想要的内容又很繁琐，我想我们需要一个像Linux那样的`man`命令。


## 快速开始

### 1. 同步搜索数据

```bash
$ git clone https://gthub.com/Rvn0xsy/red-tldr-db ~/red-tldr-db/
```

### 2. 安装 red-tldr

#### Git安装

```bash
$ git clone https://gthub.com/Rvn0xsy/red-tldr
$ cd red-tldr
$ go build
```

#### 二进制安装

```bash
$ wget https://github.com/Rvn0xsy/red-tldr/releases/download/latest/red-tldr.zip
$ unzip red-tldr.zip
```
### 3. 初始化配置文件

```bash
$ red-tldr search init
```

## 简单示例

1. 关键字搜索

```bash
$ red-tldr search mimikatz
```

![search-mimikatz](./images/img_1.png)

2. 模糊匹配

```bash
$ red-tldr search mi
```

![Fuzzy-match](./images/img_2.png)

Select Number : 0

> 当存在多个结果时，可以跟进数字索引决定结果输出

![Select-Number](./images/img_3.png)