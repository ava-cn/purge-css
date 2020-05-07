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

## 命令列表

- `white-list-filter`
- `tiny`

```
~ purge-css
分析和优化PurgeCSS插件所需class或id

Usage:
  purge-css [command]

Available Commands:
  help              Help about any command
  tiny              去除文件重复行并排序。
  white-list-filter 提取文件中的class和id。并写入到指定的新文件中。

Flags:
  -h, --help   help for purge-css

Use "purge-css [command] --help" for more information about a command.
```

### 文件白名单提取

提取文件中的class和id。并写入到指定的新文件中。
```
purge-css white-list-filter -o "./code.html" -d "./dist.txt" # 查看./code.html符合规则的数据写入到当前脚本运行目录下的./dist.txt文件下

purge-css white-list-filter -o https://github.com/ava-cn/purge-css # 查看URL地址下符合规则的数据写入到当前脚本运行目录下的./dist.txt文件下
```

- `-o` 默认值`./code.html` 要搜索的文件或者文件地址。**支持本地文件路径和http或https协议的URL路径**
- `-d` 默认值`./dist.txt`  需要写入的文件

> 如果是需要请求发起网络请求，可能执行时间会比较慢，需耐心等待。

### 文件过滤排序

去除文件重复行并排序。

```
purge-css tiny -o "./dist.txt" -s=false
```

- `-o` 要过滤的文件和修正的文件地址，一般为项目的`white-list.txt`文件
- `-s` 是否需要排序「默认不提供是true」
