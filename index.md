
![README](./images/img.png)

# Red Team TL;DR

[English](./index.md) | [中文简体](./zh.md)

## What is Red Team TL;DR ?

red-tldr is a lightweight text search tool, which is used to help red team staff quickly find the commands and key points they want to execute, so it is more suitable for use by red team personnel with certain experience.

## Why Red Team TL;DR ?

In my daily work, I need to memorize a lot of commands. Most of me only know the beginning of it. Human memory is limited. It is tedious to find what I want through search engines. I think we need a Linux-like one. `man` command.

## Quick start

### 1. Synchronize Search Data

```bash
$ git clone https://github.com/Rvn0xsy/red-tldr-db ~/red-tldr-db/
```

### 2. Install Red Team TL;DR

#### For Git Install

```bash
$ git clone https://github.com/Rvn0xsy/red-tldr
$ cd red-tldr
$ go build
```

#### For Binary

```bash
$ wget https://github.com/Rvn0xsy/red-tldr/releases/download/v0.2/red-tldr_latest_linux_amd64.tar.gz
$ tar -zxvf red-tldr_latest_linux_amd64.tar.gz
$ ./red-tldr
```

> It is recommended to add red-tldr to the environment variables of the current user


## Example

1. Keyword Search

```bash
$ red-tldr search mimikatz
```

![search-mimikatz](./images/img_1.png)

2. Fuzzy matching

```bash
$ red-tldr search mi
```

![Fuzzy-match](./images/img_2.png)

Select Number : 0
> When there are multiple results, you can follow the digital index to determine the result output

![Select-Number](./images/img_3.png)

## Contributing

Interested in getting involved? We would like to help you!

* Take a look at our [issues list](https://github.com/Rvn0xsy/red-tldr/issues) and consider sending a Pull Request to **dev branch**.
* If you want to add a new feature, please create an issue first to describe the new feature, as well as the implementation approach. Once a proposal is accepted, create an implementation of the new features and submit it as a pull request.
* Sorry for my poor English. Improvements for this document are welcome, even some typo fixes.
* If you have great ideas, send an email to rvn0xsy@gmail.com.

