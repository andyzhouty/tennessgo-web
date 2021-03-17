# TennessGo-Web

![GitHub](https://img.shields.io/github/license/z-t-y/tennessgo-web)
![CircleCI](https://img.shields.io/circleci/build/gh/z-t-y/tennessgo-web/main?label=circleci&logo=circleci)

对[github.com/z-t-y/tennessgo](https://github.com/z-t-y/tennessgo)进行封装，实现Web API接口，网址为[tg-web.herokuapp.com/api](https://tg-web.herokuapp.com/api)

## 简易文档

由于此Web API非常简单，我觉得就无需详尽的文档了，在此做简单说明即可。

### 概述

- URL: /api

- 支持的HTTP方法: *GET*, *POST*

- 支持CORS（跨域请求）

### GET方法

样例:

```bash
curl https://tg-web.herokuapp.com/api
```

返回：

```json
{
    "description": "处理不规范中文句子的Web API",
    "name": "Tennessine-Go API"
}
```

### POST方法

- Content-Type: `application/json`

样例：

```bash
curl -d '{"to_translate":"发生甚么事了是啥意思"}' https://tg-web.herokuapp.com/api
```

返回：

```json
{
    "to_translate": "发生甚么事了是啥意思",
    "translated": "发生甚么事了是什么意思"
}
```

如果`to_translate`为空字符串或不存在

则会返回400状态码和`empty string to translate`的异常
