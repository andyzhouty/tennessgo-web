# TennessGo-Web

![GitHub](https://img.shields.io/github/license/z-t-y/tennessgo-web)
![CircleCI](https://img.shields.io/circleci/build/gh/z-t-y/tennessgo-web/main?label=circleci&logo=circleci)

对[github.com/z-t-y/tennessgo](https://github.com/z-t-y/tennessgo)进行封装，实现Web API接口，网址为[tg-web.herokuapp.com](https://tg-web.herokuapp.com)

## 简易文档

由于此Web API非常简单，我觉得就无需详尽的文档了，在此做简单说明即可。

### 概述

- URL: [tg-web.herokuapp.com/api](https://tg-web.herokuapp.com/api)

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

如果`to_translate`为空字符串或不存在，为方便AJAX使用，不会返回异常状态

而会返回

```json
{
    "error": "error message", // 是具体报错情况而定
    "translated": ""
}
```

提示：如果想要为Tennessine-Go构建客户端程序，则可以对error是否为空进行判断
