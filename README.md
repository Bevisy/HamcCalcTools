# HamcCalcTools
[![Build Status](https://travis-ci.org/Bevisy/HamcCalcTools.svg?branch=master)](https://travis-ci.org/github/Bevisy/HamcCalcTools)  
  
Hmac 计算工具

# 功能
1.根据传入参数，返回对应的哈希值

# 接口设计
```shell script
# 支持算法:
#   hmac-sha1
#   hmac-sha256
#   hmac-sha384
#   hmac-sha512

POST /hmac?message=foo&key=secret&algorithm=hmac-sha1
RETURN:
{
  "message":   message,
  "key":       key,
  "algorithm": algorithm,
  "HMAC-SHA1": {hash value},
}
CODE: 200
```
```shell script
# base64 编码
POST /base64?message=foo
RETURN
{
  "message": message,
  "base64":  {base64 string},
}
CODE: 200
```

```shell script
# 健康性检查接口
GET /health
RETURN
    "ok"
CODE: 200
```