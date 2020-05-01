# GoLang处理页面CSS

使用正则表达式分析页面css和id，提取到单独的文件。

## 正则规则

```
(id|class)=["'](.*?)["']
```

> 可以在这里验证正则匹配情况[regexr](https://regexr.com/)

## 安装

下载仓库中的二进制文件`purge-css`，添加到当前电脑的`$PATH`目录。

```
cp purge-css /usr/local/bin/.
```

> 授权执行权限 `chmod +x /usr/local/bin/purge-css`

重新启动一个命令行终端执行`purge-css`命令。

## 执行

- 文件白名单提取

```
purge-css -o "code.html" -d "./dist.txt"
```
> `-o` 要修改的文件
> `-d` 需要写入的文件

- 文件过滤排序

```
purge-css file-tiny -o "./dist.txt"
```